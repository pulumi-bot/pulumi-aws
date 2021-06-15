// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

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
	case "aws:apigatewayv2/api:Api":
		r = &Api{}
	case "aws:apigatewayv2/apiMapping:ApiMapping":
		r = &ApiMapping{}
	case "aws:apigatewayv2/authorizer:Authorizer":
		r = &Authorizer{}
	case "aws:apigatewayv2/deployment:Deployment":
		r = &Deployment{}
	case "aws:apigatewayv2/domainName:DomainName":
		r = &DomainName{}
	case "aws:apigatewayv2/integration:Integration":
		r = &Integration{}
	case "aws:apigatewayv2/integrationResponse:IntegrationResponse":
		r = &IntegrationResponse{}
	case "aws:apigatewayv2/model:Model":
		r = &Model{}
	case "aws:apigatewayv2/route:Route":
		r = &Route{}
	case "aws:apigatewayv2/routeResponse:RouteResponse":
		r = &RouteResponse{}
	case "aws:apigatewayv2/stage:Stage":
		r = &Stage{}
	case "aws:apigatewayv2/vpcLink:VpcLink":
		r = &VpcLink{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/api",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/apiMapping",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/authorizer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/deployment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/domainName",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/integration",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/integrationResponse",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/model",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/routeResponse",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/stage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"aws",
		"apigatewayv2/vpcLink",
		&module{version},
	)
}
