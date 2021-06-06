// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

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
	case "aws:emr/cluster:Cluster":
		r = &Cluster{}
	case "aws:emr/instanceFleet:InstanceFleet":
		r = &InstanceFleet{}
	case "aws:emr/instanceGroup:InstanceGroup":
		r = &InstanceGroup{}
	case "aws:emr/managedScalingPolicy:ManagedScalingPolicy":
		r = &ManagedScalingPolicy{}
	case "aws:emr/securityConfiguration:SecurityConfiguration":
		r = &SecurityConfiguration{}
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
		"emr/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"emr/instanceFleet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"emr/instanceGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"emr/managedScalingPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"emr/securityConfiguration",
		&module{version},
	)
}
