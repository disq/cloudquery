// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
)

func Apigatewayv2Apis() *schema.Table {

	return &schema.Table{
		Name:      "aws_apigatewayv2_apis",
		Resolver:  fetchApigatewayv2Apis,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigatewayv2"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "protocol_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ProtocolType"),
			},
			{
				Name:     "route_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteSelectionExpression"),
			},
			{
				Name:     "api_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiEndpoint"),
			},
			{
				Name:     "api_gateway_managed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ApiGatewayManaged"),
			},
			{
				Name:     "api_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiId"),
			},
			{
				Name:     "api_key_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiKeySelectionExpression"),
			},
			{
				Name:     "cors_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CorsConfiguration"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disable_execute_api_endpoint",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableExecuteApiEndpoint"),
			},
			{
				Name:     "disable_schema_validation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableSchemaValidation"),
			},
			{
				Name:     "import_info",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ImportInfo"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "warnings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Warnings"),
			},
		},
	}
}

func fetchApigatewayv2Apis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	var input apigatewayv2.GetApisInput

	for {
		response, err := svc.GetApis(ctx, &input)
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.Items
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
