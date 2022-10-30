// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SamlIdentityProviders() *schema.Table {
	return &schema.Table{
		Name:                "aws_iam_saml_identity_providers",
		Description:         "https://docs.aws.amazon.com/IAM/latest/APIReference/API_SAMLProviderListEntry.html",
		Resolver:            fetchIamSamlIdentityProviders,
		PreResourceResolver: getSamlIdentityProvider,
		Multiplex:           client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "create_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateDate"),
			},
			{
				Name:     "valid_until",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ValidUntil"),
			},
		},
	}
}
