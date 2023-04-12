// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ssmcontacts

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceContact,
			TypeName: "aws_ssmcontacts_contact",
		},
		{
			Factory:  DataSourceContactChannel,
			TypeName: "aws_ssmcontacts_contact_channel",
		},
		{
			Factory:  DataSourcePlan,
			TypeName: "aws_ssmcontacts_plan",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceContact,
			TypeName: "aws_ssmcontacts_contact",
		},
		{
			Factory:  ResourceContactChannel,
			TypeName: "aws_ssmcontacts_contact_channel",
		},
		{
			Factory:  ResourcePlan,
			TypeName: "aws_ssmcontacts_plan",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SSMContacts
}

var ServicePackage = &servicePackage{}
