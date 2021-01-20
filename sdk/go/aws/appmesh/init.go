// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

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
	case "aws:appmesh/gatewayRoute:GatewayRoute":
		r, err = NewGatewayRoute(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appmesh/mesh:Mesh":
		r, err = NewMesh(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appmesh/route:Route":
		r, err = NewRoute(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appmesh/virtualGateway:VirtualGateway":
		r, err = NewVirtualGateway(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appmesh/virtualNode:VirtualNode":
		r, err = NewVirtualNode(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appmesh/virtualRouter:VirtualRouter":
		r, err = NewVirtualRouter(ctx, name, nil, pulumi.URN_(urn))
	case "aws:appmesh/virtualService:VirtualService":
		r, err = NewVirtualService(ctx, name, nil, pulumi.URN_(urn))
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
		"appmesh/gatewayRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/mesh",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualGateway",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualNode",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualRouter",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"appmesh/virtualService",
		&module{version},
	)
}
