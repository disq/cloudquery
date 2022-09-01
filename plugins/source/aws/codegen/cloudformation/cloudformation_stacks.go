// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package cloudformation

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func CloudformationStacks() *schema.Table {
	return &schema.Table{
		Name:      "aws_cloudformation_stacks",
		Resolver:  fetchCloudformationStacks,
		Multiplex: client.ServiceAccountRegionMultiplexer("cloudformation"),
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
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "stack_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackName"),
			},
			{
				Name:     "stack_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackStatus"),
			},
			{
				Name:     "capabilities",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Capabilities"),
			},
			{
				Name:     "change_set_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ChangeSetId"),
			},
			{
				Name:     "deletion_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeletionTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disable_rollback",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableRollback"),
			},
			{
				Name:     "drift_information",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DriftInformation"),
			},
			{
				Name:     "enable_termination_protection",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableTerminationProtection"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
			{
				Name:     "notification_ar_ns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("NotificationARNs"),
			},
			{
				Name:     "outputs",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Outputs"),
			},
			{
				Name:     "parameters",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Parameters"),
			},
			{
				Name:     "parent_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ParentId"),
			},
			{
				Name:     "role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RoleARN"),
			},
			{
				Name:     "rollback_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RollbackConfiguration"),
			},
			{
				Name:     "root_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RootId"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackId"),
			},
			{
				Name:     "stack_status_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackStatusReason"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "timeout_in_minutes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TimeoutInMinutes"),
			},
		},
	}
}

func fetchCloudformationStacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudformation

	input := cloudformation.DescribeStacksInput{}

	for {
		response, err := svc.DescribeStacks(ctx, &input)
		if err != nil {

			return diag.WrapError(err)
		}
		res <- response.Stacks
		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
