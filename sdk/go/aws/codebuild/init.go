// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
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
	case "aws:codebuild/project:Project":
		r = &Project{}
	case "aws:codebuild/reportGroup:ReportGroup":
		r = &ReportGroup{}
	case "aws:codebuild/sourceCredential:SourceCredential":
		r = &SourceCredential{}
	case "aws:codebuild/webhook:Webhook":
		r = &Webhook{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"codebuild/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"codebuild/reportGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"codebuild/sourceCredential",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"codebuild/webhook",
		&module{version},
	)
}
