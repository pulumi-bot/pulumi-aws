// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package macie2

import (
	"fmt"
	"os"

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
	case "aws:macie2/account:Account":
		r = &Account{}
	case "aws:macie2/classificationJob:ClassificationJob":
		r = &ClassificationJob{}
	case "aws:macie2/invitationAccepter:InvitationAccepter":
		r = &InvitationAccepter{}
	case "aws:macie2/member:Member":
		r = &Member{}
	case "aws:macie2/organizationAdminAccount:OrganizationAdminAccount":
		r = &OrganizationAdminAccount{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"macie2/account",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"macie2/classificationJob",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"macie2/invitationAccepter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"macie2/member",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"macie2/organizationAdminAccount",
		&module{version},
	)
}
