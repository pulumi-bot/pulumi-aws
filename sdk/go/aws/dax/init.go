// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws:dax/cluster:Cluster":
		r, err = NewCluster(ctx, name, nil, pulumi.URN_(urn))
	case "aws:dax/parameterGroup:ParameterGroup":
		r, err = NewParameterGroup(ctx, name, nil, pulumi.URN_(urn))
	case "aws:dax/subnetGroup:SubnetGroup":
		r, err = NewSubnetGroup(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"dax/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"dax/parameterGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"dax/subnetGroup",
		&module{version},
	)
}
