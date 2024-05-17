// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package route53

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newCIDRCollectionResource,
		},
		{
			Factory: newCIDRLocationResource,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceDelegationSet,
			TypeName: "aws_route53_delegation_set",
			Name:     "Reusable Delegation Set",
		},
		{
			Factory:  dataSourceTrafficPolicyDocument,
			TypeName: "aws_route53_traffic_policy_document",
			Name:     "Traffic Policy Document",
		},
		{
			Factory:  dataSourceZone,
			TypeName: "aws_route53_zone",
			Name:     "Hosted Zone",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceDelegationSet,
			TypeName: "aws_route53_delegation_set",
			Name:     "Reusable Delegation Set",
		},
		{
			Factory:  resourceHealthCheck,
			TypeName: "aws_route53_health_check",
			Name:     "Health Check",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "healthcheck",
			},
		},
		{
			Factory:  resourceHostedZoneDNSSEC,
			TypeName: "aws_route53_hosted_zone_dnssec",
			Name:     "Hosted Zone DNSSEC",
		},
		{
			Factory:  resourceKeySigningKey,
			TypeName: "aws_route53_key_signing_key",
			Name:     "Key Signing Key",
		},
		{
			Factory:  resourceQueryLog,
			TypeName: "aws_route53_query_log",
			Name:     "Query Logging Config",
		},
		{
			Factory:  ResourceRecord,
			TypeName: "aws_route53_record",
		},
		{
			Factory:  resourceTrafficPolicy,
			TypeName: "aws_route53_traffic_policy",
			Name:     "Traffic Policy",
		},
		{
			Factory:  resourceTrafficPolicyInstance,
			TypeName: "aws_route53_traffic_policy_instance",
			Name:     "Traffic Policy Instance",
		},
		{
			Factory:  resourceVPCAssociationAuthorization,
			TypeName: "aws_route53_vpc_association_authorization",
			Name:     "VPC Association Authorization",
		},
		{
			Factory:  resourceZone,
			TypeName: "aws_route53_zone",
			Name:     "Hosted Zone",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
				ResourceType:        "hostedzone",
			},
		},
		{
			Factory:  resourceZoneAssociation,
			TypeName: "aws_route53_zone_association",
			Name:     "Zone Association",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Route53
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
