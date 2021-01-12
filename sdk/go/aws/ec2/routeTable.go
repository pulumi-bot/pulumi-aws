// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a VPC routing table.
//
// > **NOTE on Route Tables and Routes:** This provider currently
// provides both a standalone Route resource and a Route Table resource with routes
// defined in-line. At this time you cannot use a Route Table with in-line routes
// in conjunction with any Route resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// > **NOTE on `gatewayId` and `natGatewayId`:** The AWS API is very forgiving with these two
// attributes and the `ec2.RouteTable` resource can be created with a NAT ID specified as a Gateway ID attribute.
// This _will_ lead to a permanent diff between your configuration and statefile, as the API returns the correct
// parameters in the returned route table. If you're experiencing constant diffs in your `ec2.RouteTable` resources,
// the first thing to check is whether or not you're specifying a NAT ID instead of a Gateway ID, or vice-versa.
//
// > **NOTE on `propagatingVgws` and the `ec2.VpnGatewayRoutePropagation` resource:**
// If the `propagatingVgws` argument is present, it's not supported to _also_
// define route propagations using `ec2.VpnGatewayRoutePropagation`, since
// this resource will delete any propagating gateways not explicitly listed in
// `propagatingVgws`. Omit this argument when defining route propagation using
// the separate resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2.NewRouteTable(ctx, "routeTable", &ec2.RouteTableArgs{
// 			VpcId: pulumi.Any(aws_vpc.Default.Id),
// 			Routes: ec2.RouteTableRouteArray{
// 				&ec2.RouteTableRouteArgs{
// 					CidrBlock: pulumi.String("10.0.1.0/24"),
// 					GatewayId: pulumi.Any(aws_internet_gateway.Main.Id),
// 				},
// 				&ec2.RouteTableRouteArgs{
// 					Ipv6CidrBlock:       pulumi.String("::/0"),
// 					EgressOnlyGatewayId: pulumi.Any(aws_egress_only_internet_gateway.Foo.Id),
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Name": pulumi.String("main"),
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
// Route Tables can be imported using the route table `id`. For example, to import route table `rtb-4e616f6d69`, use this command
//
// ```sh
//  $ pulumi import aws:ec2/routeTable:RouteTable public_rt rtb-4e616f6d69
// ```
type RouteTable struct {
	pulumi.CustomResourceState

	// The ID of the AWS account that owns the route table.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayOutput `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	Routes RouteTableRouteArrayOutput `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource RouteTable
	err := ctx.RegisterResource("aws:ec2/routeTable:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("aws:ec2/routeTable:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
	// The ID of the AWS account that owns the route table.
	OwnerId *string `pulumi:"ownerId"`
	// A list of virtual gateways for propagation.
	PropagatingVgws []string `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	Routes []RouteTableRoute `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type RouteTableState struct {
	// The ID of the AWS account that owns the route table.
	OwnerId pulumi.StringPtrInput
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayInput
	// A list of route objects. Their keys are documented below.
	Routes RouteTableRouteArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	// A list of virtual gateways for propagation.
	PropagatingVgws []string `pulumi:"propagatingVgws"`
	// A list of route objects. Their keys are documented below.
	Routes []RouteTableRoute `pulumi:"routes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	// A list of virtual gateways for propagation.
	PropagatingVgws pulumi.StringArrayInput
	// A list of route objects. Their keys are documented below.
	Routes RouteTableRouteArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTable)(nil))
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

func (i *RouteTable) ToRouteTablePtrOutput() RouteTablePtrOutput {
	return i.ToRouteTablePtrOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTablePtrOutputWithContext(ctx context.Context) RouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTablePtrOutput)
}

type RouteTablePtrInput interface {
	pulumi.Input

	ToRouteTablePtrOutput() RouteTablePtrOutput
	ToRouteTablePtrOutputWithContext(ctx context.Context) RouteTablePtrOutput
}

type routeTablePtrType RouteTableArgs

func (*routeTablePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil))
}

func (i *routeTablePtrType) ToRouteTablePtrOutput() RouteTablePtrOutput {
	return i.ToRouteTablePtrOutputWithContext(context.Background())
}

func (i *routeTablePtrType) ToRouteTablePtrOutputWithContext(ctx context.Context) RouteTablePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput).ToRouteTablePtrOutput()
}

// RouteTableArrayInput is an input type that accepts RouteTableArray and RouteTableArrayOutput values.
// You can construct a concrete instance of `RouteTableArrayInput` via:
//
//          RouteTableArray{ RouteTableArgs{...} }
type RouteTableArrayInput interface {
	pulumi.Input

	ToRouteTableArrayOutput() RouteTableArrayOutput
	ToRouteTableArrayOutputWithContext(context.Context) RouteTableArrayOutput
}

type RouteTableArray []RouteTableInput

func (RouteTableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteTable)(nil))
}

func (i RouteTableArray) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return i.ToRouteTableArrayOutputWithContext(context.Background())
}

func (i RouteTableArray) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableArrayOutput)
}

// RouteTableMapInput is an input type that accepts RouteTableMap and RouteTableMapOutput values.
// You can construct a concrete instance of `RouteTableMapInput` via:
//
//          RouteTableMap{ "key": RouteTableArgs{...} }
type RouteTableMapInput interface {
	pulumi.Input

	ToRouteTableMapOutput() RouteTableMapOutput
	ToRouteTableMapOutputWithContext(context.Context) RouteTableMapOutput
}

type RouteTableMap map[string]RouteTableInput

func (RouteTableMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RouteTable)(nil))
}

func (i RouteTableMap) ToRouteTableMapOutput() RouteTableMapOutput {
	return i.ToRouteTableMapOutputWithContext(context.Background())
}

func (i RouteTableMap) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableMapOutput)
}

type RouteTableOutput struct {
	*pulumi.OutputState
}

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTable)(nil))
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTablePtrOutput() RouteTablePtrOutput {
	return o.ToRouteTablePtrOutputWithContext(context.Background())
}

func (o RouteTableOutput) ToRouteTablePtrOutputWithContext(ctx context.Context) RouteTablePtrOutput {
	return o.ApplyT(func(v RouteTable) *RouteTable {
		return &v
	}).(RouteTablePtrOutput)
}

type RouteTablePtrOutput struct {
	*pulumi.OutputState
}

func (RouteTablePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouteTable)(nil))
}

func (o RouteTablePtrOutput) ToRouteTablePtrOutput() RouteTablePtrOutput {
	return o
}

func (o RouteTablePtrOutput) ToRouteTablePtrOutputWithContext(ctx context.Context) RouteTablePtrOutput {
	return o
}

type RouteTableArrayOutput struct{ *pulumi.OutputState }

func (RouteTableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteTable)(nil))
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutput() RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) ToRouteTableArrayOutputWithContext(ctx context.Context) RouteTableArrayOutput {
	return o
}

func (o RouteTableArrayOutput) Index(i pulumi.IntInput) RouteTableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteTable {
		return vs[0].([]RouteTable)[vs[1].(int)]
	}).(RouteTableOutput)
}

type RouteTableMapOutput struct{ *pulumi.OutputState }

func (RouteTableMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RouteTable)(nil))
}

func (o RouteTableMapOutput) ToRouteTableMapOutput() RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) ToRouteTableMapOutputWithContext(ctx context.Context) RouteTableMapOutput {
	return o
}

func (o RouteTableMapOutput) MapIndex(k pulumi.StringInput) RouteTableOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RouteTable {
		return vs[0].(map[string]RouteTable)[vs[1].(string)]
	}).(RouteTableOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteTableOutput{})
	pulumi.RegisterOutputType(RouteTablePtrOutput{})
	pulumi.RegisterOutputType(RouteTableArrayOutput{})
	pulumi.RegisterOutputType(RouteTableMapOutput{})
}
