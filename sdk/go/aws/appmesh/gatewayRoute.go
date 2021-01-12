// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS App Mesh gateway route resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/appmesh"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appmesh.NewGatewayRoute(ctx, "example", &appmesh.GatewayRouteArgs{
// 			MeshName:           pulumi.String("example-service-mesh"),
// 			VirtualGatewayName: pulumi.Any(aws_appmesh_virtual_gateway.Example.Name),
// 			Spec: &appmesh.GatewayRouteSpecArgs{
// 				HttpRoute: &appmesh.GatewayRouteSpecHttpRouteArgs{
// 					Action: &appmesh.GatewayRouteSpecHttpRouteActionArgs{
// 						Target: &appmesh.GatewayRouteSpecHttpRouteActionTargetArgs{
// 							VirtualService: &appmesh.GatewayRouteSpecHttpRouteActionTargetVirtualServiceArgs{
// 								VirtualServiceName: pulumi.Any(aws_appmesh_virtual_service.Example.Name),
// 							},
// 						},
// 					},
// 					Match: &appmesh.GatewayRouteSpecHttpRouteMatchArgs{
// 						Prefix: pulumi.String("/"),
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// App Mesh gateway routes can be imported using `mesh_name` and `virtual_gateway_name` together with the gateway route's `name`, e.g.
//
// ```sh
//  $ pulumi import aws:appmesh/gatewayRoute:GatewayRoute example mesh/gw1/example-gateway-route
// ```
//
//  [1]/docs/providers/aws/index.html
type GatewayRoute struct {
	pulumi.CustomResourceState

	// The ARN of the gateway route.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the gateway route.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the gateway route.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// The name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The gateway route specification to apply.
	Spec GatewayRouteSpecOutput `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName pulumi.StringOutput `pulumi:"virtualGatewayName"`
}

// NewGatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewGatewayRoute(ctx *pulumi.Context,
	name string, args *GatewayRouteArgs, opts ...pulumi.ResourceOption) (*GatewayRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.VirtualGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualGatewayName'")
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
	// The name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName *string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The gateway route specification to apply.
	Spec *GatewayRouteSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName *string `pulumi:"virtualGatewayName"`
}

type GatewayRouteState struct {
	// The ARN of the gateway route.
	Arn pulumi.StringPtrInput
	// The creation date of the gateway route.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the gateway route.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringPtrInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// The gateway route specification to apply.
	Spec GatewayRouteSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName pulumi.StringPtrInput
}

func (GatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteState)(nil)).Elem()
}

type gatewayRouteArgs struct {
	// The name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The gateway route specification to apply.
	Spec GatewayRouteSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName string `pulumi:"virtualGatewayName"`
}

// The set of arguments for constructing a GatewayRoute resource.
type GatewayRouteArgs struct {
	// The name of the service mesh in which to create the gateway route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the gateway route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The gateway route specification to apply.
	Spec GatewayRouteSpecInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with. Must be between 1 and 255 characters in length.
	VirtualGatewayName pulumi.StringInput
}

func (GatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteArgs)(nil)).Elem()
}

type GatewayRouteInput interface {
	pulumi.Input

	ToGatewayRouteOutput() GatewayRouteOutput
	ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput
}

func (*GatewayRoute) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRoute)(nil))
}

func (i *GatewayRoute) ToGatewayRouteOutput() GatewayRouteOutput {
	return i.ToGatewayRouteOutputWithContext(context.Background())
}

func (i *GatewayRoute) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteOutput)
}

func (i *GatewayRoute) ToGatewayRoutePtrOutput() GatewayRoutePtrOutput {
	return i.ToGatewayRoutePtrOutputWithContext(context.Background())
}

func (i *GatewayRoute) ToGatewayRoutePtrOutputWithContext(ctx context.Context) GatewayRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRoutePtrOutput)
}

type GatewayRoutePtrInput interface {
	pulumi.Input

	ToGatewayRoutePtrOutput() GatewayRoutePtrOutput
	ToGatewayRoutePtrOutputWithContext(ctx context.Context) GatewayRoutePtrOutput
}

type gatewayRoutePtrType GatewayRouteArgs

func (*gatewayRoutePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil))
}

func (i *gatewayRoutePtrType) ToGatewayRoutePtrOutput() GatewayRoutePtrOutput {
	return i.ToGatewayRoutePtrOutputWithContext(context.Background())
}

func (i *gatewayRoutePtrType) ToGatewayRoutePtrOutputWithContext(ctx context.Context) GatewayRoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteOutput).ToGatewayRoutePtrOutput()
}

// GatewayRouteArrayInput is an input type that accepts GatewayRouteArray and GatewayRouteArrayOutput values.
// You can construct a concrete instance of `GatewayRouteArrayInput` via:
//
//          GatewayRouteArray{ GatewayRouteArgs{...} }
type GatewayRouteArrayInput interface {
	pulumi.Input

	ToGatewayRouteArrayOutput() GatewayRouteArrayOutput
	ToGatewayRouteArrayOutputWithContext(context.Context) GatewayRouteArrayOutput
}

type GatewayRouteArray []GatewayRouteInput

func (GatewayRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayRoute)(nil))
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return i.ToGatewayRouteArrayOutputWithContext(context.Background())
}

func (i GatewayRouteArray) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteArrayOutput)
}

// GatewayRouteMapInput is an input type that accepts GatewayRouteMap and GatewayRouteMapOutput values.
// You can construct a concrete instance of `GatewayRouteMapInput` via:
//
//          GatewayRouteMap{ "key": GatewayRouteArgs{...} }
type GatewayRouteMapInput interface {
	pulumi.Input

	ToGatewayRouteMapOutput() GatewayRouteMapOutput
	ToGatewayRouteMapOutputWithContext(context.Context) GatewayRouteMapOutput
}

type GatewayRouteMap map[string]GatewayRouteInput

func (GatewayRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GatewayRoute)(nil))
}

func (i GatewayRouteMap) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return i.ToGatewayRouteMapOutputWithContext(context.Background())
}

func (i GatewayRouteMap) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteMapOutput)
}

type GatewayRouteOutput struct {
	*pulumi.OutputState
}

func (GatewayRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayRoute)(nil))
}

func (o GatewayRouteOutput) ToGatewayRouteOutput() GatewayRouteOutput {
	return o
}

func (o GatewayRouteOutput) ToGatewayRouteOutputWithContext(ctx context.Context) GatewayRouteOutput {
	return o
}

func (o GatewayRouteOutput) ToGatewayRoutePtrOutput() GatewayRoutePtrOutput {
	return o.ToGatewayRoutePtrOutputWithContext(context.Background())
}

func (o GatewayRouteOutput) ToGatewayRoutePtrOutputWithContext(ctx context.Context) GatewayRoutePtrOutput {
	return o.ApplyT(func(v GatewayRoute) *GatewayRoute {
		return &v
	}).(GatewayRoutePtrOutput)
}

type GatewayRoutePtrOutput struct {
	*pulumi.OutputState
}

func (GatewayRoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRoute)(nil))
}

func (o GatewayRoutePtrOutput) ToGatewayRoutePtrOutput() GatewayRoutePtrOutput {
	return o
}

func (o GatewayRoutePtrOutput) ToGatewayRoutePtrOutputWithContext(ctx context.Context) GatewayRoutePtrOutput {
	return o
}

type GatewayRouteArrayOutput struct{ *pulumi.OutputState }

func (GatewayRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GatewayRoute)(nil))
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutput() GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) ToGatewayRouteArrayOutputWithContext(ctx context.Context) GatewayRouteArrayOutput {
	return o
}

func (o GatewayRouteArrayOutput) Index(i pulumi.IntInput) GatewayRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GatewayRoute {
		return vs[0].([]GatewayRoute)[vs[1].(int)]
	}).(GatewayRouteOutput)
}

type GatewayRouteMapOutput struct{ *pulumi.OutputState }

func (GatewayRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GatewayRoute)(nil))
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutput() GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) ToGatewayRouteMapOutputWithContext(ctx context.Context) GatewayRouteMapOutput {
	return o
}

func (o GatewayRouteMapOutput) MapIndex(k pulumi.StringInput) GatewayRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GatewayRoute {
		return vs[0].(map[string]GatewayRoute)[vs[1].(string)]
	}).(GatewayRouteOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayRouteOutput{})
	pulumi.RegisterOutputType(GatewayRoutePtrOutput{})
	pulumi.RegisterOutputType(GatewayRouteArrayOutput{})
	pulumi.RegisterOutputType(GatewayRouteMapOutput{})
}
