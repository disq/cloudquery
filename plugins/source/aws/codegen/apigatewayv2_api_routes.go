// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

func Apigatewayv2ApiRoutes() *schema.Table {
	return &schema.Table{
		Name:      "aws_apigatewayv2_api_routes",
		Resolver:  fetchApigatewayv2ApiRoutes,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigatewayv2"),
		Columns: []schema.Column{
			{
				Name:     "route_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteKey"),
			},
			{
				Name:     "api_gateway_managed",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ApiGatewayManaged"),
			},
			{
				Name:     "api_key_required",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ApiKeyRequired"),
			},
			{
				Name:     "authorization_scopes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AuthorizationScopes"),
			},
			{
				Name:     "authorization_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizationType"),
			},
			{
				Name:     "authorizer_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthorizerId"),
			},
			{
				Name:     "model_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelSelectionExpression"),
			},
			{
				Name:     "operation_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OperationName"),
			},
			{
				Name:     "request_models",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RequestModels"),
			},
			{
				Name:     "request_parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RequestParameters"),
			},
			{
				Name:     "route_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteId"),
			},
			{
				Name:     "route_response_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteResponseSelectionExpression"),
			},
			{
				Name:     "target",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Target"),
			},
		},
	}
}

func fetchApigatewayv2ApiRoutes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	r := parent.Item.(types.Api)

	input := apigatewayv2.GetRoutesInput{
		ApiId: r.ApiId,
	}

	for {
		response, err := svc.GetRoutes(ctx, &input)
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
