// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"testing"

{{if .HasTags}}
  "github.com/aws/aws-sdk-go-v2/aws"
{{end}}
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"{{.TypesImport}}"
{{range .MockImports}}	{{.}}
{{end}}
)

func {{.MockFuncName}}(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMock{{.AWSService}}Client(ctrl)

{{if .MockRawPaginatorListType}}
	var item {{.MockRawPaginatorListType}}
{{else}}
	var item types.{{.PaginatorListName}}
{{end}}
	if err := faker.FakeData(&item); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().{{.ListVerb | Coalesce "List"}}{{.AWSSubService}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.{{.ListVerb | Coalesce "List"}}{{.AWSSubService}}Input{},
		gomock.Any(),
	).Return(
		&{{.AWSService | ToLower}}.{{.ListVerb | Coalesce "List"}}{{.AWSSubService}}Output{
{{if .MockRawPaginatorListType}}
		  {{.PaginatorListName}}: []{{.MockRawPaginatorListType}}{item},
{{else}}
		  {{.PaginatorListName}}: []types.{{.PaginatorListName}}{item},
{{end}}
    },
		nil,
	)

{{if .MockRawListDetailType}}
  var detail {{.MockRawListDetailType}}
{{else}}
	var detail types.{{.ItemName}}Detail
{{end}}
	if err := faker.FakeData(&detail); err != nil {
		t.Fatal(err)
	}
{{if .RawDescribeFieldValue}}
	detail.{{.ListFieldName}} = {{.RawDescribeFieldValue}}
{{else}}
	detail.{{.ListFieldName}} = item.{{.ListFieldName}}
{{end}}
	mock.EXPECT().{{.Verb | Coalesce "Describe"}}{{.ItemName}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.{{.Verb | Coalesce "Describe"}}{{.ItemName}}Input{
{{if .RawDescribeFieldValue}}
		  {{.ListFieldName}}: {{.RawDescribeFieldValue}},
{{else}}
	{{.ListFieldName}}: item.{{.ListFieldName}},
{{end}}
		},
		gomock.Any(),
	).Return(
		&{{.AWSService | ToLower}}.{{.Verb | Coalesce "Describe"}}{{.ItemName}}Output{
		  {{.ItemName}}: &detail,
    },
		nil,
	)

{{if .HasTags}}
	mock.EXPECT().ListTagsFor{{.ItemName}}(
		gomock.Any(),
		&{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Input{
		  {{.ListFieldName}}: detail.{{.ListFieldName}},
    },
	).Return(
		&{{.AWSService | ToLower}}.ListTagsFor{{.ItemName}}Output{
			Tags: []types.Tag{
				{Key: aws.String("key"), Value: aws.String("value")},
			},
		},
		nil,
	)
{{end}}
	return client.Services{
	  {{.AWSService}}: mock,
  }
}

func {{.TestFuncName}}(t *testing.T) {
	client.AwsMockTestHelper(t, {{.TableFuncName}}(), {{.MockFuncName}}, client.TestOptions{})
}
