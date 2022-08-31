// Code generated by codegen; DO NOT EDIT.

package apigatewayv2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/apigatewayv2"
)

func Apigatewayv2VpcLinks() *schema.Table {
	return &schema.Table{
		Name:      "aws_apigatewayv2_vpc_links",
		Resolver:  fetchApigatewayv2VpcLinks,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigatewayv2"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "security_group_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecurityGroupIds"),
			},
			{
				Name:     "subnet_ids",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubnetIds"),
			},
			{
				Name:     "vpc_link_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkId"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "vpc_link_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkStatus"),
			},
			{
				Name:     "vpc_link_status_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkStatusMessage"),
			},
			{
				Name:     "vpc_link_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcLinkVersion"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolvers.ResolveVPCLinkArn,
			},
		},
	}
}

func fetchApigatewayv2VpcLinks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Apigatewayv2

	input := apigatewayv2.GetVpcLinksInput{}

	for {
		response, err := svc.GetVpcLinks(ctx, &input)
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
