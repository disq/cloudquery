// Code generated by codegen; DO NOT EDIT.

package lightsail

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Certificates() *schema.Table {
	return &schema.Table{
		Name:        "aws_lightsail_certificates",
		Description: "https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_Certificate.html",
		Resolver:    fetchLightsailCertificates,
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
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "domain_validation_records",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DomainValidationRecords"),
			},
			{
				Name:     "eligible_to_renew",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EligibleToRenew"),
			},
			{
				Name:     "in_use_resource_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InUseResourceCount"),
			},
			{
				Name:     "issued_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("IssuedAt"),
			},
			{
				Name:     "issuer_ca",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IssuerCA"),
			},
			{
				Name:     "key_algorithm",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyAlgorithm"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "not_after",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotAfter"),
			},
			{
				Name:     "not_before",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("NotBefore"),
			},
			{
				Name:     "renewal_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("RenewalSummary"),
			},
			{
				Name:     "request_failure_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RequestFailureReason"),
			},
			{
				Name:     "revocation_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RevocationReason"),
			},
			{
				Name:     "revoked_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("RevokedAt"),
			},
			{
				Name:     "serial_number",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SerialNumber"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "subject_alternative_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubjectAlternativeNames"),
			},
			{
				Name:     "support_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SupportCode"),
			},
		},
	}
}
