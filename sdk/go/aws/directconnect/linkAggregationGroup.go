// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Direct Connect LAG. Connections can be added to the LAG via the `directconnect.Connection` and `directconnect.ConnectionAssociation` resources.
//
// > *NOTE:* When creating a LAG, Direct Connect requires creating a Connection. This provider will remove this unmanaged connection during resource creation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := directconnect.NewLinkAggregationGroup(ctx, "hoge", &directconnect.LinkAggregationGroupArgs{
// 			ConnectionsBandwidth: pulumi.String("1Gbps"),
// 			ForceDestroy:         pulumi.Bool(true),
// 			Location:             pulumi.String("EqDC2"),
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
// Direct Connect LAGs can be imported using the `lag id`, e.g.
//
// ```sh
//  $ pulumi import aws:directconnect/linkAggregationGroup:LinkAggregationGroup test_lag dxlag-fgnsp5rq
// ```
type LinkAggregationGroup struct {
	pulumi.CustomResourceState

	// The ARN of the LAG.
	// * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	ConnectionsBandwidth pulumi.StringOutput `pulumi:"connectionsBandwidth"`
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringOutput `pulumi:"hasLogicalRedundancy"`
	JumboFrameCapable    pulumi.BoolOutput   `pulumi:"jumboFrameCapable"`
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the LAG.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewLinkAggregationGroup registers a new resource with the given unique name, arguments, and options.
func NewLinkAggregationGroup(ctx *pulumi.Context,
	name string, args *LinkAggregationGroupArgs, opts ...pulumi.ResourceOption) (*LinkAggregationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionsBandwidth == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionsBandwidth'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource LinkAggregationGroup
	err := ctx.RegisterResource("aws:directconnect/linkAggregationGroup:LinkAggregationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkAggregationGroup gets an existing LinkAggregationGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkAggregationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkAggregationGroupState, opts ...pulumi.ResourceOption) (*LinkAggregationGroup, error) {
	var resource LinkAggregationGroup
	err := ctx.ReadResource("aws:directconnect/linkAggregationGroup:LinkAggregationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkAggregationGroup resources.
type linkAggregationGroupState struct {
	// The ARN of the LAG.
	// * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
	Arn *string `pulumi:"arn"`
	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	ConnectionsBandwidth *string `pulumi:"connectionsBandwidth"`
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy *string `pulumi:"hasLogicalRedundancy"`
	JumboFrameCapable    *bool   `pulumi:"jumboFrameCapable"`
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location *string `pulumi:"location"`
	// The name of the LAG.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type LinkAggregationGroupState struct {
	// The ARN of the LAG.
	// * `jumboFrameCapable` -Indicates whether jumbo frames (9001 MTU) are supported.
	Arn pulumi.StringPtrInput
	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	ConnectionsBandwidth pulumi.StringPtrInput
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// Indicates whether the LAG supports a secondary BGP peer in the same address family (IPv4/IPv6).
	HasLogicalRedundancy pulumi.StringPtrInput
	JumboFrameCapable    pulumi.BoolPtrInput
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringPtrInput
	// The name of the LAG.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LinkAggregationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAggregationGroupState)(nil)).Elem()
}

type linkAggregationGroupArgs struct {
	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	ConnectionsBandwidth string `pulumi:"connectionsBandwidth"`
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location string `pulumi:"location"`
	// The name of the LAG.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LinkAggregationGroup resource.
type LinkAggregationGroupArgs struct {
	// The bandwidth of the individual physical connections bundled by the LAG. Valid values: 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps and 10Gbps. Case sensitive.
	ConnectionsBandwidth pulumi.StringInput
	// A boolean that indicates all connections associated with the LAG should be deleted so that the LAG can be destroyed without error. These objects are *not* recoverable.
	ForceDestroy pulumi.BoolPtrInput
	// The AWS Direct Connect location in which the LAG should be allocated. See [DescribeLocations](https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DescribeLocations.html) for the list of AWS Direct Connect locations. Use `locationCode`.
	Location pulumi.StringInput
	// The name of the LAG.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (LinkAggregationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAggregationGroupArgs)(nil)).Elem()
}

type LinkAggregationGroupInput interface {
	pulumi.Input

	ToLinkAggregationGroupOutput() LinkAggregationGroupOutput
	ToLinkAggregationGroupOutputWithContext(ctx context.Context) LinkAggregationGroupOutput
}

func (*LinkAggregationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkAggregationGroup)(nil))
}

func (i *LinkAggregationGroup) ToLinkAggregationGroupOutput() LinkAggregationGroupOutput {
	return i.ToLinkAggregationGroupOutputWithContext(context.Background())
}

func (i *LinkAggregationGroup) ToLinkAggregationGroupOutputWithContext(ctx context.Context) LinkAggregationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAggregationGroupOutput)
}

func (i *LinkAggregationGroup) ToLinkAggregationGroupPtrOutput() LinkAggregationGroupPtrOutput {
	return i.ToLinkAggregationGroupPtrOutputWithContext(context.Background())
}

func (i *LinkAggregationGroup) ToLinkAggregationGroupPtrOutputWithContext(ctx context.Context) LinkAggregationGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAggregationGroupPtrOutput)
}

type LinkAggregationGroupPtrInput interface {
	pulumi.Input

	ToLinkAggregationGroupPtrOutput() LinkAggregationGroupPtrOutput
	ToLinkAggregationGroupPtrOutputWithContext(ctx context.Context) LinkAggregationGroupPtrOutput
}

type linkAggregationGroupPtrType LinkAggregationGroupArgs

func (*linkAggregationGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkAggregationGroup)(nil))
}

func (i *linkAggregationGroupPtrType) ToLinkAggregationGroupPtrOutput() LinkAggregationGroupPtrOutput {
	return i.ToLinkAggregationGroupPtrOutputWithContext(context.Background())
}

func (i *linkAggregationGroupPtrType) ToLinkAggregationGroupPtrOutputWithContext(ctx context.Context) LinkAggregationGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAggregationGroupPtrOutput)
}

// LinkAggregationGroupArrayInput is an input type that accepts LinkAggregationGroupArray and LinkAggregationGroupArrayOutput values.
// You can construct a concrete instance of `LinkAggregationGroupArrayInput` via:
//
//          LinkAggregationGroupArray{ LinkAggregationGroupArgs{...} }
type LinkAggregationGroupArrayInput interface {
	pulumi.Input

	ToLinkAggregationGroupArrayOutput() LinkAggregationGroupArrayOutput
	ToLinkAggregationGroupArrayOutputWithContext(context.Context) LinkAggregationGroupArrayOutput
}

type LinkAggregationGroupArray []LinkAggregationGroupInput

func (LinkAggregationGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinkAggregationGroup)(nil))
}

func (i LinkAggregationGroupArray) ToLinkAggregationGroupArrayOutput() LinkAggregationGroupArrayOutput {
	return i.ToLinkAggregationGroupArrayOutputWithContext(context.Background())
}

func (i LinkAggregationGroupArray) ToLinkAggregationGroupArrayOutputWithContext(ctx context.Context) LinkAggregationGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAggregationGroupArrayOutput)
}

// LinkAggregationGroupMapInput is an input type that accepts LinkAggregationGroupMap and LinkAggregationGroupMapOutput values.
// You can construct a concrete instance of `LinkAggregationGroupMapInput` via:
//
//          LinkAggregationGroupMap{ "key": LinkAggregationGroupArgs{...} }
type LinkAggregationGroupMapInput interface {
	pulumi.Input

	ToLinkAggregationGroupMapOutput() LinkAggregationGroupMapOutput
	ToLinkAggregationGroupMapOutputWithContext(context.Context) LinkAggregationGroupMapOutput
}

type LinkAggregationGroupMap map[string]LinkAggregationGroupInput

func (LinkAggregationGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinkAggregationGroup)(nil))
}

func (i LinkAggregationGroupMap) ToLinkAggregationGroupMapOutput() LinkAggregationGroupMapOutput {
	return i.ToLinkAggregationGroupMapOutputWithContext(context.Background())
}

func (i LinkAggregationGroupMap) ToLinkAggregationGroupMapOutputWithContext(ctx context.Context) LinkAggregationGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAggregationGroupMapOutput)
}

type LinkAggregationGroupOutput struct {
	*pulumi.OutputState
}

func (LinkAggregationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkAggregationGroup)(nil))
}

func (o LinkAggregationGroupOutput) ToLinkAggregationGroupOutput() LinkAggregationGroupOutput {
	return o
}

func (o LinkAggregationGroupOutput) ToLinkAggregationGroupOutputWithContext(ctx context.Context) LinkAggregationGroupOutput {
	return o
}

func (o LinkAggregationGroupOutput) ToLinkAggregationGroupPtrOutput() LinkAggregationGroupPtrOutput {
	return o.ToLinkAggregationGroupPtrOutputWithContext(context.Background())
}

func (o LinkAggregationGroupOutput) ToLinkAggregationGroupPtrOutputWithContext(ctx context.Context) LinkAggregationGroupPtrOutput {
	return o.ApplyT(func(v LinkAggregationGroup) *LinkAggregationGroup {
		return &v
	}).(LinkAggregationGroupPtrOutput)
}

type LinkAggregationGroupPtrOutput struct {
	*pulumi.OutputState
}

func (LinkAggregationGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkAggregationGroup)(nil))
}

func (o LinkAggregationGroupPtrOutput) ToLinkAggregationGroupPtrOutput() LinkAggregationGroupPtrOutput {
	return o
}

func (o LinkAggregationGroupPtrOutput) ToLinkAggregationGroupPtrOutputWithContext(ctx context.Context) LinkAggregationGroupPtrOutput {
	return o
}

type LinkAggregationGroupArrayOutput struct{ *pulumi.OutputState }

func (LinkAggregationGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkAggregationGroup)(nil))
}

func (o LinkAggregationGroupArrayOutput) ToLinkAggregationGroupArrayOutput() LinkAggregationGroupArrayOutput {
	return o
}

func (o LinkAggregationGroupArrayOutput) ToLinkAggregationGroupArrayOutputWithContext(ctx context.Context) LinkAggregationGroupArrayOutput {
	return o
}

func (o LinkAggregationGroupArrayOutput) Index(i pulumi.IntInput) LinkAggregationGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkAggregationGroup {
		return vs[0].([]LinkAggregationGroup)[vs[1].(int)]
	}).(LinkAggregationGroupOutput)
}

type LinkAggregationGroupMapOutput struct{ *pulumi.OutputState }

func (LinkAggregationGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkAggregationGroup)(nil))
}

func (o LinkAggregationGroupMapOutput) ToLinkAggregationGroupMapOutput() LinkAggregationGroupMapOutput {
	return o
}

func (o LinkAggregationGroupMapOutput) ToLinkAggregationGroupMapOutputWithContext(ctx context.Context) LinkAggregationGroupMapOutput {
	return o
}

func (o LinkAggregationGroupMapOutput) MapIndex(k pulumi.StringInput) LinkAggregationGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkAggregationGroup {
		return vs[0].(map[string]LinkAggregationGroup)[vs[1].(string)]
	}).(LinkAggregationGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkAggregationGroupOutput{})
	pulumi.RegisterOutputType(LinkAggregationGroupPtrOutput{})
	pulumi.RegisterOutputType(LinkAggregationGroupArrayOutput{})
	pulumi.RegisterOutputType(LinkAggregationGroupMapOutput{})
}
