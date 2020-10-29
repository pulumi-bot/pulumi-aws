// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS App Mesh gateway route resource.
type GatewayRoute struct {
	pulumi.CustomResourceState

	// The ARN of the gateway route.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the gateway route.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the gateway route.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the gateway route.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// The name to use for the gateway route.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The gateway route specification to apply.
	Spec GatewayRouteSpecOutput `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
	VirtualGatewayName pulumi.StringOutput `pulumi:"virtualGatewayName"`
}

// NewGatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewGatewayRoute(ctx *pulumi.Context,
	name string, args *GatewayRouteArgs, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil || args.VirtualGatewayName == nil {
		return nil, errors.New("missing required argument 'VirtualGatewayName'")
	}
	if args == nil {
		args = &GatewayRouteArgs{}
	}
	var resource GatewayRoute
	err := ctx.RegisterResource("aws:appmesh/gatewayRoute:GatewayRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayRoute gets an existing GatewayRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayRouteState, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	var resource GatewayRoute
	err := ctx.ReadResource("aws:appmesh/gatewayRoute:GatewayRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayRoute resources.
type gatewayRouteState struct {
	// The ARN of the gateway route.
	Arn *string `pulumi:"arn"`
	// The creation date of the gateway route.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the gateway route.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the gateway route.
	MeshName *string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the gateway route.
	Name *string `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The gateway route specification to apply.
	Spec *GatewayRouteSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
	VirtualGatewayName *string `pulumi:"virtualGatewayName"`
}

type GatewayRouteState struct {
	// The ARN of the gateway route.
	Arn pulumi.StringPtrInput
	// The creation date of the gateway route.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the gateway route.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the gateway route.
	MeshName pulumi.StringPtrInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the gateway route.
	Name pulumi.StringPtrInput
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// The gateway route specification to apply.
	Spec GatewayRouteSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
	VirtualGatewayName pulumi.StringPtrInput
}

func (GatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteState)(nil)).Elem()
}

type gatewayRouteArgs struct {
	// The name of the service mesh in which to create the gateway route.
	MeshName string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the gateway route.
	Name *string `pulumi:"name"`
	// The gateway route specification to apply.
	Spec GatewayRouteSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
	VirtualGatewayName string `pulumi:"virtualGatewayName"`
}

// The set of arguments for constructing a GatewayRoute resource.
type GatewayRouteArgs struct {
	// The name of the service mesh in which to create the gateway route.
	MeshName pulumi.StringInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the gateway route.
	Name pulumi.StringPtrInput
	// The gateway route specification to apply.
	Spec GatewayRouteSpecInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
	VirtualGatewayName pulumi.StringInput
}

func (GatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteArgs)(nil)).Elem()
}
