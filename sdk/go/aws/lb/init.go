// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lb

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
	case "aws:lb/listener:Listener":
		r = &Listener{}
	case "aws:lb/listenerCertificate:ListenerCertificate":
		r = &ListenerCertificate{}
	case "aws:lb/listenerRule:ListenerRule":
		r = &ListenerRule{}
	case "aws:lb/loadBalancer:LoadBalancer":
		r = &LoadBalancer{}
	case "aws:lb/targetGroup:TargetGroup":
		r = &TargetGroup{}
	case "aws:lb/targetGroupAttachment:TargetGroupAttachment":
		r = &TargetGroupAttachment{}
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
		"lb/listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lb/listenerCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lb/listenerRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lb/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lb/targetGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"lb/targetGroupAttachment",
		&module{version},
	)
}
