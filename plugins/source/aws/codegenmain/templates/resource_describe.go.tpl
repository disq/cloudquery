// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

{{range .Imports}}	"{{.}}"
{{end}}
)

func {{.TableFuncName}}() *schema.Table {
	return &schema.Table{{template "table.go.tpl" .Table}}
}

func {{.Table.Resolver}}(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService}}

	var input {{.AWSService | ToLower}}.{{.ListFunctionName}}Input
	paginator := {{.AWSService | ToLower}}.New{{.ListFunctionName}}Paginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, item := range output.{{.ItemName}}SummaryList {
			do, err := svc.{{.DescribeFunctionName}}(ctx, &{{.AWSService | ToLower}}.{{.DescribeFunctionName}}Input{
			  {{.DescribeFieldName}}: item.{{.DescribeFieldName}},
			})
			if err != nil {
				return diag.WrapError(err)
			}
			res <- do.{{.ItemName}}
		}
	}
	return nil
}

{{if .HasTags}}
func resolve{{.AWSService | ToCamel}}{{.AWSSubService | ToCamel}}Tags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cert := resource.Item.(*types.{{.AWSStructName}})
	cl := meta.(*client.Client)
	svc := cl.Services().{{.AWSService}}
	out, err := svc.ListTagsFor{{.ItemName}}(ctx, &{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Input{
	  {{.DescribeFieldName}}: cert.{{.DescribeFieldName}},
  })
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, client.TagsToMap(out.Tags)))
}
{{end}}

{{range .CustomResolvers}}{{.}}
{{end}}
