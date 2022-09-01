package helpers

import (
	"log"
	"reflect"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/recipes"
	"github.com/iancoleman/strcase"
)

func TableAndFetcherNames(r *recipes.Resource) (string, string) {
	cqSubservice := Coalesce(r.CQSubserviceOverride, r.AWSSubService)

	tableNameFromSubService := strcase.ToSnake(cqSubservice)
	fetcherNameFromSubService := strcase.ToCamel(cqSubservice)
	{
		// Generate table and fetcher names using parent info

		prev := tableNameFromSubService
		var (
			preTableNames   []string
			preFetcherNames []string
		)
		rp := r.Parent
		for rp != nil {
			if rp.CQSubserviceOverride != "" {
				preTableNames = append(preTableNames, rp.CQSubserviceOverride)
				preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.CQSubserviceOverride))
			} else {
				ins := strcase.ToSnake(rp.ItemName)
				if !strings.HasPrefix(prev, ins) {
					preTableNames = append(preTableNames, ins)
					preFetcherNames = append(preFetcherNames, strcase.ToCamel(rp.ItemName))
					prev = ins
				}
			}
			rp = rp.Parent
		}
		if len(preTableNames) > 0 {
			tableNameFromSubService = strings.Join(ReverseStringSlice(preTableNames), "_") + "_" + tableNameFromSubService
			fetcherNameFromSubService = strings.Join(ReverseStringSlice(preFetcherNames), "") + fetcherNameFromSubService
		}
	}

	return tableNameFromSubService, fetcherNameFromSubService
}

type InferResult struct {
	Method     string
	SubService string

	ItemsFieldCandidates []reflect.StructField // Struct field candidates that contains the item or items. Only valid in Output type structs.
	PaginatorTokenField  *reflect.StructField

	Fields     map[string]reflect.Type
	FieldOrder []string
}

func (ir *InferResult) ItemsField() reflect.StructField {
	if len(ir.ItemsFieldCandidates) != 1 {
		log.Fatal("Could not determine ItemsName for ", ir.Method, ":", len(ir.ItemsFieldCandidates), " candidates")
	}
	return ir.ItemsFieldCandidates[0]
}

func InferFromStruct(s interface{}, expectSingular, expectInput bool) *InferResult {
	res := InferResult{
		Fields: make(map[string]reflect.Type),
	}

	ist := BareType(reflect.TypeOf(s))
	suffix := StringSwitch(expectInput, "Input", "Output")

	for _, verbCandidate := range []string{"Get", "Describe", "List"} {
		if !strings.HasPrefix(ist.Name(), verbCandidate) {
			continue
		}

		if !strings.HasSuffix(ist.Name(), suffix) {
			log.Fatal("Unhandled struct type (bad suffix): ", ist.Name())
		}

		res.Method = strings.TrimSuffix(ist.Name(), suffix)

		res.SubService = strings.TrimPrefix(res.Method, verbCandidate)
		if res.SubService == "" {
			log.Fatal("Failed calculating AWSSubService: empty output for", ist.Name())
		}

		break
	}

	for i := 0; i < ist.NumField(); i++ {
		f := ist.Field(i)
		if f.Name == "noSmithyDocumentSerde" || f.Type.String() == "document.NoSerde" || f.Type.String() == "middleware.Metadata" {
			continue
		}
		if f.Name == "NextToken" {
			res.PaginatorTokenField = &f
		}

		res.Fields[f.Name] = f.Type
		res.FieldOrder = append(res.FieldOrder, f.Name)

		if expectSingular && f.Type.Kind() != reflect.Slice {
			res.ItemsFieldCandidates = append(res.ItemsFieldCandidates, f)
			continue
		}
		if !expectSingular && f.Type.Kind() == reflect.Slice {
			res.ItemsFieldCandidates = append(res.ItemsFieldCandidates, f)
			continue
		}
	}

	return &res
}
