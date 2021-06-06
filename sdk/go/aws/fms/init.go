// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package fms

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
	case "aws:fms/adminAccount:AdminAccount":
		r = &AdminAccount{}
	case "aws:fms/policy:Policy":
		r = &Policy{}
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
		"fms/adminAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"fms/policy",
		&module{version},
	)
}
