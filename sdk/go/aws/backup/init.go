// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

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
	case "aws:backup/globalSettings:GlobalSettings":
		r = &GlobalSettings{}
	case "aws:backup/plan:Plan":
		r = &Plan{}
	case "aws:backup/regionSettings:RegionSettings":
		r = &RegionSettings{}
	case "aws:backup/selection:Selection":
		r = &Selection{}
	case "aws:backup/vault:Vault":
		r = &Vault{}
	case "aws:backup/vaultNotifications:VaultNotifications":
		r = &VaultNotifications{}
	case "aws:backup/vaultPolicy:VaultPolicy":
		r = &VaultPolicy{}
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
		"backup/globalSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"backup/plan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"backup/regionSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"backup/selection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"backup/vault",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"backup/vaultNotifications",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"backup/vaultPolicy",
		&module{version},
	)
}
