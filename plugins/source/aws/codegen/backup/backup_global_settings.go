// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package backup

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func BackupGlobalSettings() *schema.Table {
	return &schema.Table{
		Name:      "aws_backup_global_settings",
		Resolver:  fetchBackupGlobalSettings,
		Multiplex: client.ServiceAccountRegionMultiplexer("backup"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:     "global_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GlobalSettings"),
			},
			{
				Name:     "last_update_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastUpdateTime"),
			},
		},
	}
}

func fetchBackupGlobalSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup

	input := backup.DescribeGlobalSettingsInput{}

	{
		response, err := svc.DescribeGlobalSettings(ctx, &input)
		if err != nil {

			if client.IgnoreAccessDeniedServiceDisabled(err) || client.IsAWSError(err, "ERROR_9601") /* "Your account is not a member of an organization" */ {
				meta.Logger().Debug("received access denied on DescribeGlobalSettings", "err", err)
				return nil
			}
			if client.IsAWSError(err, "ERROR_2502") /* "Feature Cross Account Backup is not available in current region" */ {
				meta.Logger().Debug("Feature Cross Account Backup is not available in current region on DescribeGlobalSettings", "err", err)
				return nil
			}

			return diag.WrapError(err)
		}

		res <- response

	}
	return nil
}
