// Code generated by codegen; DO NOT EDIT.

package redshift

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SubnetGroups() *schema.Table {
	return &schema.Table{
		Name:        "aws_redshift_subnet_groups",
		Description: "https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterSubnetGroup.html",
		Resolver:    fetchRedshiftSubnetGroups,
		Multiplex:   client.ServiceAccountRegionMultiplexer("redshift"),
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
				Name:        "arn",
				Type:        schema.TypeString,
				Resolver:    resolveSubnetGroupArn(),
				Description: `The Amazon Resource Name (ARN) for the resource.`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
				Description: `The list of tags for the cluster subnet group.`,
			},
			{
				Name:     "cluster_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClusterSubnetGroupName"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "subnet_group_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetGroupStatus"),
			},
			{
				Name:     "subnets",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Subnets"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
