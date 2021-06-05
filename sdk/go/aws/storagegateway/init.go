// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storagegateway

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
	case "aws:storagegateway/cache:Cache":
		r = &Cache{}
	case "aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume":
		r = &CachesIscsiVolume{}
	case "aws:storagegateway/gateway:Gateway":
		r = &Gateway{}
	case "aws:storagegateway/nfsFileShare:NfsFileShare":
		r = &NfsFileShare{}
	case "aws:storagegateway/smbFileShare:SmbFileShare":
		r = &SmbFileShare{}
	case "aws:storagegateway/storedIscsiVolume:StoredIscsiVolume":
		r = &StoredIscsiVolume{}
	case "aws:storagegateway/tapePool:TapePool":
		r = &TapePool{}
	case "aws:storagegateway/uploadBuffer:UploadBuffer":
		r = &UploadBuffer{}
	case "aws:storagegateway/workingStorage:WorkingStorage":
		r = &WorkingStorage{}
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
		"storagegateway/cache",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/cachesIscsiVolume",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/gateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/nfsFileShare",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/smbFileShare",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/storedIscsiVolume",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/tapePool",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/uploadBuffer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"storagegateway/workingStorage",
		&module{version},
	)
}
