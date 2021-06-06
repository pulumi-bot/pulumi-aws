// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancingv2

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
	case "aws:elasticloadbalancingv2/listener:Listener":
		r = &Listener{}
	case "aws:elasticloadbalancingv2/listenerCertificate:ListenerCertificate":
		r = &ListenerCertificate{}
	case "aws:elasticloadbalancingv2/listenerRule:ListenerRule":
		r = &ListenerRule{}
	case "aws:elasticloadbalancingv2/loadBalancer:LoadBalancer":
		r = &LoadBalancer{}
	case "aws:elasticloadbalancingv2/targetGroup:TargetGroup":
		r = &TargetGroup{}
	case "aws:elasticloadbalancingv2/targetGroupAttachment:TargetGroupAttachment":
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
		"elasticloadbalancingv2/listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticloadbalancingv2/listenerCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticloadbalancingv2/listenerRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticloadbalancingv2/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticloadbalancingv2/targetGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elasticloadbalancingv2/targetGroupAttachment",
		&module{version},
	)
}
