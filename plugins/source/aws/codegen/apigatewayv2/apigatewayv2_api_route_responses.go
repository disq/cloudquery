// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

func Apigatewayv2ApiRouteResponses() *schema.Table {
	return &schema.Table{
		Name:      "aws_apigatewayv2_api_route_responses",
		Resolver:  fetchApigatewayv2ApiRouteResponses,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigatewayv2"),
		Columns: []schema.Column{
			{
				Name:     "route_response_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteResponseKey"),
			},
			{
				Name:     "model_selection_expression",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ModelSelectionExpression"),
			},
			{
				Name:     "response_models",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResponseModels"),
			},
			{
				Name:     "response_parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResponseParameters"),
			},
			{
				Name:     "route_response_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RouteResponseId"),
			},
		},
	}
}

func fetchApigatewayv2ApiRouteResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	r2 := parent.Item.(types.Route)
	parent = parent.Parent

	r1 := parent.Item.(types.Api)

	input := apigatewayv2.GetRouteResponsesInput{
		RouteId: r2.RouteId,
		ApiId:   r1.ApiId,
	}

	for {
		response, err := svc.GetRouteResponses(ctx, &input)
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
