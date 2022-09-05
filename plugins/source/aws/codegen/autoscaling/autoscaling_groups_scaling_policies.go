// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package autoscaling

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	resolvers "github.com/cloudquery/cloudquery/plugins/source/aws/codegenmain/resolvers/autoscaling"
)

func AutoscalingGroupsScalingPolicies() *schema.Table {
	return &schema.Table{
		Name:      "aws_autoscaling_groups_scaling_policies",
		Resolver:  fetchAutoscalingGroupsScalingPolicies,
		Multiplex: client.ServiceAccountRegionMultiplexer("autoscaling"),
		Columns: []schema.Column{
			{
				Name:     "adjustment_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AdjustmentType"),
			},
			{
				Name:     "alarms",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Alarms"),
			},
			{
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "cooldown",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Cooldown"),
			},
			{
				Name:     "enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Enabled"),
			},
			{
				Name:     "estimated_instance_warmup",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EstimatedInstanceWarmup"),
			},
			{
				Name:     "metric_aggregation_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetricAggregationType"),
			},
			{
				Name:     "min_adjustment_magnitude",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinAdjustmentMagnitude"),
			},
			{
				Name:     "min_adjustment_step",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinAdjustmentStep"),
			},
			{
				Name:     "policy_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyARN"),
			},
			{
				Name:     "policy_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyName"),
			},
			{
				Name:     "policy_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PolicyType"),
			},
			{
				Name:     "predictive_scaling_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PredictiveScalingConfiguration"),
			},
			{
				Name:     "scaling_adjustment",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ScalingAdjustment"),
			},
			{
				Name:     "step_adjustments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StepAdjustments"),
			},
			{
				Name:     "target_tracking_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("TargetTrackingConfiguration"),
			},
		},
	}
}

func fetchAutoscalingGroupsScalingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling

	r1 := parent.Item.(types.AutoScalingGroup)

	input := autoscaling.DescribePoliciesInput{
		AutoScalingGroupName: r1.AutoScalingGroupName,
	}

	for {
		response, err := svc.DescribePolicies(ctx, &input)
		if err != nil {

			if resolvers.IsGroupNotExistsError(err) {
				return nil
			}
			return diag.WrapError(err)
		}

		res <- response.ScalingPolicies

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}
