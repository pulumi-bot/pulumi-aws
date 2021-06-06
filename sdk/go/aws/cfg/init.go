// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:cfg/aggregateAuthorization:AggregateAuthorization":
		r = &AggregateAuthorization{}
	case "aws:cfg/configurationAggregator:ConfigurationAggregator":
		r = &ConfigurationAggregator{}
	case "aws:cfg/conformancePack:ConformancePack":
		r = &ConformancePack{}
	case "aws:cfg/deliveryChannel:DeliveryChannel":
		r = &DeliveryChannel{}
	case "aws:cfg/organizationCustomRule:OrganizationCustomRule":
		r = &OrganizationCustomRule{}
	case "aws:cfg/organizationManagedRule:OrganizationManagedRule":
		r = &OrganizationManagedRule{}
	case "aws:cfg/recorder:Recorder":
		r = &Recorder{}
	case "aws:cfg/recorderStatus:RecorderStatus":
		r = &RecorderStatus{}
	case "aws:cfg/remediationConfiguration:RemediationConfiguration":
		r = &RemediationConfiguration{}
	case "aws:cfg/rule:Rule":
		r = &Rule{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version := aws.PkgVersion()
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/aggregateAuthorization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/configurationAggregator",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/conformancePack",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/deliveryChannel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/organizationCustomRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/organizationManagedRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/recorder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/recorderStatus",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/remediationConfiguration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"cfg/rule",
		&module{version},
	)
}
