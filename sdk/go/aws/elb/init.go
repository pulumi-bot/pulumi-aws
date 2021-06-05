// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

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
	case "aws:elb/appCookieStickinessPolicy:AppCookieStickinessPolicy":
		r = &AppCookieStickinessPolicy{}
	case "aws:elb/attachment:Attachment":
		r = &Attachment{}
	case "aws:elb/listenerPolicy:ListenerPolicy":
		r = &ListenerPolicy{}
	case "aws:elb/loadBalancer:LoadBalancer":
		r = &LoadBalancer{}
	case "aws:elb/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy":
		r = &LoadBalancerBackendServerPolicy{}
	case "aws:elb/loadBalancerCookieStickinessPolicy:LoadBalancerCookieStickinessPolicy":
		r = &LoadBalancerCookieStickinessPolicy{}
	case "aws:elb/loadBalancerPolicy:LoadBalancerPolicy":
		r = &LoadBalancerPolicy{}
	case "aws:elb/sslNegotiationPolicy:SslNegotiationPolicy":
		r = &SslNegotiationPolicy{}
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
		"elb/appCookieStickinessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elb/attachment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elb/listenerPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elb/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elb/loadBalancerBackendServerPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elb/loadBalancerCookieStickinessPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elb/loadBalancerPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"elb/sslNegotiationPolicy",
		&module{version},
	)
}
