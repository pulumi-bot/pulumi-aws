// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package acmpca

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
	case "aws:acmpca/certificate:Certificate":
		r = &Certificate{}
	case "aws:acmpca/certificateAuthority:CertificateAuthority":
		r = &CertificateAuthority{}
	case "aws:acmpca/certificateAuthorityCertificate:CertificateAuthorityCertificate":
		r = &CertificateAuthorityCertificate{}
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
		"acmpca/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"acmpca/certificateAuthority",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"acmpca/certificateAuthorityCertificate",
		&module{version},
	)
}
