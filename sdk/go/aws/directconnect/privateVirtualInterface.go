// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Direct Connect private virtual interface resource.
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
// 		_, err := directconnect.NewPrivateVirtualInterface(ctx, "foo", &directconnect.PrivateVirtualInterfaceArgs{
// 			AddressFamily: pulumi.String("ipv4"),
// 			BgpAsn:        pulumi.Int(65352),
// 			ConnectionId:  pulumi.String("dxcon-zzzzzzzz"),
// 			Vlan:          pulumi.Int(4094),
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
// Direct Connect private virtual interfaces can be imported using the `vif id`, e.g.
//
// ```sh
//  $ pulumi import aws:directconnect/privateVirtualInterface:PrivateVirtualInterface test dxvif-33cc44dd
// ```
type PrivateVirtualInterface struct {
	pulumi.CustomResourceState

	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringOutput `pulumi:"amazonAddress"`
	AmazonSideAsn pulumi.StringOutput `pulumi:"amazonSideAsn"`
	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringOutput `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntOutput `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringOutput `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringOutput `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringOutput `pulumi:"customerAddress"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrOutput `pulumi:"dxGatewayId"`
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable pulumi.BoolOutput `pulumi:"jumboFrameCapable"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrOutput `pulumi:"mtu"`
	// The name for the virtual interface.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VLAN ID.
	Vlan pulumi.IntOutput `pulumi:"vlan"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrOutput `pulumi:"vpnGatewayId"`
}

// NewPrivateVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewPrivateVirtualInterface(ctx *pulumi.Context,
	name string, args *PrivateVirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*PrivateVirtualInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AddressFamily == nil {
		return nil, errors.New("invalid value for required argument 'AddressFamily'")
	}
	if args.BgpAsn == nil {
		return nil, errors.New("invalid value for required argument 'BgpAsn'")
	}
	if args.ConnectionId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionId'")
	}
	if args.Vlan == nil {
		return nil, errors.New("invalid value for required argument 'Vlan'")
	}
	var resource PrivateVirtualInterface
	err := ctx.RegisterResource("aws:directconnect/privateVirtualInterface:PrivateVirtualInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateVirtualInterface gets an existing PrivateVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateVirtualInterfaceState, opts ...pulumi.ResourceOption) (*PrivateVirtualInterface, error) {
	var resource PrivateVirtualInterface
	err := ctx.ReadResource("aws:directconnect/privateVirtualInterface:PrivateVirtualInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateVirtualInterface resources.
type privateVirtualInterfaceState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily *string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `pulumi:"amazonAddress"`
	AmazonSideAsn *string `pulumi:"amazonSideAsn"`
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice *string `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn *int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId *string `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `pulumi:"customerAddress"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `pulumi:"jumboFrameCapable"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VLAN ID.
	Vlan *int `pulumi:"vlan"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type PrivateVirtualInterfaceState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringPtrInput
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringPtrInput
	AmazonSideAsn pulumi.StringPtrInput
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDevice pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntPtrInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringPtrInput
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringPtrInput
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable pulumi.BoolPtrInput
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VLAN ID.
	Vlan pulumi.IntPtrInput
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrInput
}

func (PrivateVirtualInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateVirtualInterfaceState)(nil)).Elem()
}

type privateVirtualInterfaceArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress *string `pulumi:"amazonAddress"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId string `pulumi:"connectionId"`
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress *string `pulumi:"customerAddress"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu *int `pulumi:"mtu"`
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VLAN ID.
	Vlan int `pulumi:"vlan"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a PrivateVirtualInterface resource.
type PrivateVirtualInterfaceArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringInput
	// The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
	AmazonAddress pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
	ConnectionId pulumi.StringInput
	// The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
	CustomerAddress pulumi.StringPtrInput
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
	// The MTU of a virtual private interface can be either `1500` or `9001` (jumbo frames). Default is `1500`.
	Mtu pulumi.IntPtrInput
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VLAN ID.
	Vlan pulumi.IntInput
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrInput
}

func (PrivateVirtualInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateVirtualInterfaceArgs)(nil)).Elem()
}

type PrivateVirtualInterfaceInput interface {
	pulumi.Input

	ToPrivateVirtualInterfaceOutput() PrivateVirtualInterfaceOutput
	ToPrivateVirtualInterfaceOutputWithContext(ctx context.Context) PrivateVirtualInterfaceOutput
}

func (*PrivateVirtualInterface) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateVirtualInterface)(nil))
}

func (i *PrivateVirtualInterface) ToPrivateVirtualInterfaceOutput() PrivateVirtualInterfaceOutput {
	return i.ToPrivateVirtualInterfaceOutputWithContext(context.Background())
}

func (i *PrivateVirtualInterface) ToPrivateVirtualInterfaceOutputWithContext(ctx context.Context) PrivateVirtualInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateVirtualInterfaceOutput)
}

func (i *PrivateVirtualInterface) ToPrivateVirtualInterfacePtrOutput() PrivateVirtualInterfacePtrOutput {
	return i.ToPrivateVirtualInterfacePtrOutputWithContext(context.Background())
}

func (i *PrivateVirtualInterface) ToPrivateVirtualInterfacePtrOutputWithContext(ctx context.Context) PrivateVirtualInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateVirtualInterfacePtrOutput)
}

type PrivateVirtualInterfacePtrInput interface {
	pulumi.Input

	ToPrivateVirtualInterfacePtrOutput() PrivateVirtualInterfacePtrOutput
	ToPrivateVirtualInterfacePtrOutputWithContext(ctx context.Context) PrivateVirtualInterfacePtrOutput
}

type privateVirtualInterfacePtrType PrivateVirtualInterfaceArgs

func (*privateVirtualInterfacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateVirtualInterface)(nil))
}

func (i *privateVirtualInterfacePtrType) ToPrivateVirtualInterfacePtrOutput() PrivateVirtualInterfacePtrOutput {
	return i.ToPrivateVirtualInterfacePtrOutputWithContext(context.Background())
}

func (i *privateVirtualInterfacePtrType) ToPrivateVirtualInterfacePtrOutputWithContext(ctx context.Context) PrivateVirtualInterfacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateVirtualInterfaceOutput).ToPrivateVirtualInterfacePtrOutput()
}

// PrivateVirtualInterfaceArrayInput is an input type that accepts PrivateVirtualInterfaceArray and PrivateVirtualInterfaceArrayOutput values.
// You can construct a concrete instance of `PrivateVirtualInterfaceArrayInput` via:
//
//          PrivateVirtualInterfaceArray{ PrivateVirtualInterfaceArgs{...} }
type PrivateVirtualInterfaceArrayInput interface {
	pulumi.Input

	ToPrivateVirtualInterfaceArrayOutput() PrivateVirtualInterfaceArrayOutput
	ToPrivateVirtualInterfaceArrayOutputWithContext(context.Context) PrivateVirtualInterfaceArrayOutput
}

type PrivateVirtualInterfaceArray []PrivateVirtualInterfaceInput

func (PrivateVirtualInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateVirtualInterface)(nil))
}

func (i PrivateVirtualInterfaceArray) ToPrivateVirtualInterfaceArrayOutput() PrivateVirtualInterfaceArrayOutput {
	return i.ToPrivateVirtualInterfaceArrayOutputWithContext(context.Background())
}

func (i PrivateVirtualInterfaceArray) ToPrivateVirtualInterfaceArrayOutputWithContext(ctx context.Context) PrivateVirtualInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateVirtualInterfaceArrayOutput)
}

// PrivateVirtualInterfaceMapInput is an input type that accepts PrivateVirtualInterfaceMap and PrivateVirtualInterfaceMapOutput values.
// You can construct a concrete instance of `PrivateVirtualInterfaceMapInput` via:
//
//          PrivateVirtualInterfaceMap{ "key": PrivateVirtualInterfaceArgs{...} }
type PrivateVirtualInterfaceMapInput interface {
	pulumi.Input

	ToPrivateVirtualInterfaceMapOutput() PrivateVirtualInterfaceMapOutput
	ToPrivateVirtualInterfaceMapOutputWithContext(context.Context) PrivateVirtualInterfaceMapOutput
}

type PrivateVirtualInterfaceMap map[string]PrivateVirtualInterfaceInput

func (PrivateVirtualInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PrivateVirtualInterface)(nil))
}

func (i PrivateVirtualInterfaceMap) ToPrivateVirtualInterfaceMapOutput() PrivateVirtualInterfaceMapOutput {
	return i.ToPrivateVirtualInterfaceMapOutputWithContext(context.Background())
}

func (i PrivateVirtualInterfaceMap) ToPrivateVirtualInterfaceMapOutputWithContext(ctx context.Context) PrivateVirtualInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateVirtualInterfaceMapOutput)
}

type PrivateVirtualInterfaceOutput struct {
	*pulumi.OutputState
}

func (PrivateVirtualInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateVirtualInterface)(nil))
}

func (o PrivateVirtualInterfaceOutput) ToPrivateVirtualInterfaceOutput() PrivateVirtualInterfaceOutput {
	return o
}

func (o PrivateVirtualInterfaceOutput) ToPrivateVirtualInterfaceOutputWithContext(ctx context.Context) PrivateVirtualInterfaceOutput {
	return o
}

func (o PrivateVirtualInterfaceOutput) ToPrivateVirtualInterfacePtrOutput() PrivateVirtualInterfacePtrOutput {
	return o.ToPrivateVirtualInterfacePtrOutputWithContext(context.Background())
}

func (o PrivateVirtualInterfaceOutput) ToPrivateVirtualInterfacePtrOutputWithContext(ctx context.Context) PrivateVirtualInterfacePtrOutput {
	return o.ApplyT(func(v PrivateVirtualInterface) *PrivateVirtualInterface {
		return &v
	}).(PrivateVirtualInterfacePtrOutput)
}

type PrivateVirtualInterfacePtrOutput struct {
	*pulumi.OutputState
}

func (PrivateVirtualInterfacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateVirtualInterface)(nil))
}

func (o PrivateVirtualInterfacePtrOutput) ToPrivateVirtualInterfacePtrOutput() PrivateVirtualInterfacePtrOutput {
	return o
}

func (o PrivateVirtualInterfacePtrOutput) ToPrivateVirtualInterfacePtrOutputWithContext(ctx context.Context) PrivateVirtualInterfacePtrOutput {
	return o
}

type PrivateVirtualInterfaceArrayOutput struct{ *pulumi.OutputState }

func (PrivateVirtualInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateVirtualInterface)(nil))
}

func (o PrivateVirtualInterfaceArrayOutput) ToPrivateVirtualInterfaceArrayOutput() PrivateVirtualInterfaceArrayOutput {
	return o
}

func (o PrivateVirtualInterfaceArrayOutput) ToPrivateVirtualInterfaceArrayOutputWithContext(ctx context.Context) PrivateVirtualInterfaceArrayOutput {
	return o
}

func (o PrivateVirtualInterfaceArrayOutput) Index(i pulumi.IntInput) PrivateVirtualInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateVirtualInterface {
		return vs[0].([]PrivateVirtualInterface)[vs[1].(int)]
	}).(PrivateVirtualInterfaceOutput)
}

type PrivateVirtualInterfaceMapOutput struct{ *pulumi.OutputState }

func (PrivateVirtualInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PrivateVirtualInterface)(nil))
}

func (o PrivateVirtualInterfaceMapOutput) ToPrivateVirtualInterfaceMapOutput() PrivateVirtualInterfaceMapOutput {
	return o
}

func (o PrivateVirtualInterfaceMapOutput) ToPrivateVirtualInterfaceMapOutputWithContext(ctx context.Context) PrivateVirtualInterfaceMapOutput {
	return o
}

func (o PrivateVirtualInterfaceMapOutput) MapIndex(k pulumi.StringInput) PrivateVirtualInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PrivateVirtualInterface {
		return vs[0].(map[string]PrivateVirtualInterface)[vs[1].(string)]
	}).(PrivateVirtualInterfaceOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateVirtualInterfaceOutput{})
	pulumi.RegisterOutputType(PrivateVirtualInterfacePtrOutput{})
	pulumi.RegisterOutputType(PrivateVirtualInterfaceArrayOutput{})
	pulumi.RegisterOutputType(PrivateVirtualInterfaceMapOutput{})
}
