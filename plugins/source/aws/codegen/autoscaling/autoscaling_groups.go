// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package autoscaling

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func AutoscalingGroups() *schema.Table {
	return &schema.Table{
		Name:      "aws_autoscaling_groups",
		Resolver:  fetchAutoscalingGroups,
		Multiplex: client.ServiceAccountRegionMultiplexer("autoscaling"),
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
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "availability_zones",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AvailabilityZones"),
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "default_cooldown",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DefaultCooldown"),
			},
			{
				Name:     "desired_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DesiredCapacity"),
			},
			{
				Name:     "health_check_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HealthCheckType"),
			},
			{
				Name:     "max_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSize"),
			},
			{
				Name:     "min_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinSize"),
			},
			{
				Name:     "auto_scaling_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupARN"),
			},
			{
				Name:     "capacity_rebalance",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CapacityRebalance"),
			},
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Context"),
			},
			{
				Name:     "default_instance_warmup",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DefaultInstanceWarmup"),
			},
			{
				Name:     "desired_capacity_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DesiredCapacityType"),
			},
			{
				Name:     "enabled_metrics",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnabledMetrics"),
			},
			{
				Name:     "health_check_grace_period",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("HealthCheckGracePeriod"),
			},
			{
				Name:     "instances",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Instances"),
			},
			{
				Name:     "launch_configuration_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchConfigurationName"),
			},
			{
				Name:     "launch_template",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LaunchTemplate"),
			},
			{
				Name:     "load_balancer_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("LoadBalancerNames"),
			},
			{
				Name:     "max_instance_lifetime",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxInstanceLifetime"),
			},
			{
				Name:     "mixed_instances_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MixedInstancesPolicy"),
			},
			{
				Name:     "new_instances_protected_from_scale_in",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("NewInstancesProtectedFromScaleIn"),
			},
			{
				Name:     "placement_group",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlacementGroup"),
			},
			{
				Name:     "predicted_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("PredictedCapacity"),
			},
			{
				Name:     "service_linked_role_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ServiceLinkedRoleARN"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "suspended_processes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SuspendedProcesses"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "target_group_ar_ns",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("TargetGroupARNs"),
			},
			{
				Name:     "termination_policies",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("TerminationPolicies"),
			},
			{
				Name:     "vpc_zone_identifier",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VPCZoneIdentifier"),
			},
			{
				Name:     "warm_pool_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("WarmPoolConfiguration"),
			},
			{
				Name:     "warm_pool_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("WarmPoolSize"),
			},
		},
	}
}

func fetchAutoscalingGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling

	input := autoscaling.DescribeAutoScalingGroupsInput{}

	for {
		response, err := svc.DescribeAutoScalingGroups(ctx, &input)
		if err != nil {

			return diag.WrapError(err)
		}

		res <- response.AutoScalingGroups

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
