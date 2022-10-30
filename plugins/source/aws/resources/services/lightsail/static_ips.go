// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func StaticIps() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_static_ips",
		Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_StaticIp.html",
		Resolver:    fetchLightsailStaticIps,
		Multiplex:   client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name: "arn",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attached_to",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AttachedTo"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "ip_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpAddress"),
			},
			{
				Name:     "is_attached",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsAttached"),
			},
			{
				Name:     "location",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Location"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "resource_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceType"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
		},
	}
}
