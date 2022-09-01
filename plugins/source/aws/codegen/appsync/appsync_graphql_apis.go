// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package appsync

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/appsync"
)

func AppSyncGraphqlApis() *schema.Table {
	return &schema.Table{
		Name:      "aws_appsync_graphql_apis",
		Resolver:  fetchAppSyncGraphqlApis,
		Multiplex: client.ServiceAccountRegionMultiplexer("appsync"),
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
				Name:        "namespace",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSNamespace,
				Description: `The AWS Service Namespace of the resource.`,
			},
			{
				Name:     "additional_authentication_providers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AdditionalAuthenticationProviders"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiId"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "authentication_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AuthenticationType"),
			},
			{
				Name:     "lambda_authorizer_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LambdaAuthorizerConfig"),
			},
			{
				Name:     "log_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogConfig"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "open_id_connect_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OpenIDConnectConfig"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "uris",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Uris"),
			},
			{
				Name:     "user_pool_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("UserPoolConfig"),
			},
			{
				Name:     "waf_web_acl_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WafWebAclArn"),
			},
			{
				Name:     "xray_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("XrayEnabled"),
			},
		},
	}
}

func fetchAppSyncGraphqlApis(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().AppSync

	input := appsync.ListGraphqlApisInput{}

	for {
		response, err := svc.ListGraphqlApis(ctx, &input)
		if err != nil {

			return diag.WrapError(err)
		}
		res <- response.GraphqlApis
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
