// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Direct Connect public virtual interface resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/directconnect"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = directconnect.NewPublicVirtualInterface(ctx, "foo", &directconnect.PublicVirtualInterfaceArgs{
// 			AddressFamily:   pulumi.String("ipv4"),
// 			AmazonAddress:   pulumi.String("175.45.176.2/30"),
// 			BgpAsn:          pulumi.Int(65352),
// 			ConnectionId:    pulumi.String("dxcon-zzzzzzzz"),
// 			CustomerAddress: pulumi.String("175.45.176.1/30"),
// 			RouteFilterPrefixes: pulumi.StringArray{
// 				pulumi.String("210.52.109.0/24"),
// 				pulumi.String("175.45.176.0/22"),
// 			},
// 			Vlan: pulumi.Int(4094),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type PublicVirtualInterface struct {
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
	// The name for the virtual interface.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayOutput `pulumi:"routeFilterPrefixes"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VLAN ID.
	Vlan pulumi.IntOutput `pulumi:"vlan"`
}

// NewPublicVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewPublicVirtualInterface(ctx *pulumi.Context,
	name string, args *PublicVirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*PublicVirtualInterface, error) {
	if args == nil || args.AddressFamily == nil {
		return nil, errors.New("missing required argument 'AddressFamily'")
	}
	if args == nil || args.BgpAsn == nil {
		return nil, errors.New("missing required argument 'BgpAsn'")
	}
	if args == nil || args.ConnectionId == nil {
		return nil, errors.New("missing required argument 'ConnectionId'")
	}
	if args == nil || args.RouteFilterPrefixes == nil {
		return nil, errors.New("missing required argument 'RouteFilterPrefixes'")
	}
	if args == nil || args.Vlan == nil {
		return nil, errors.New("missing required argument 'Vlan'")
	}
	if args == nil {
		args = &PublicVirtualInterfaceArgs{}
	}
	var resource PublicVirtualInterface
	err := ctx.RegisterResource("aws:directconnect/publicVirtualInterface:PublicVirtualInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicVirtualInterface gets an existing PublicVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicVirtualInterfaceState, opts ...pulumi.ResourceOption) (*PublicVirtualInterface, error) {
	var resource PublicVirtualInterface
	err := ctx.ReadResource("aws:directconnect/publicVirtualInterface:PublicVirtualInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicVirtualInterface resources.
type publicVirtualInterfaceState struct {
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
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []string `pulumi:"routeFilterPrefixes"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VLAN ID.
	Vlan *int `pulumi:"vlan"`
}

type PublicVirtualInterfaceState struct {
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
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VLAN ID.
	Vlan pulumi.IntPtrInput
}

func (PublicVirtualInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicVirtualInterfaceState)(nil)).Elem()
}

type publicVirtualInterfaceArgs struct {
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
	// The name for the virtual interface.
	Name *string `pulumi:"name"`
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes []string `pulumi:"routeFilterPrefixes"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VLAN ID.
	Vlan int `pulumi:"vlan"`
}

// The set of arguments for constructing a PublicVirtualInterface resource.
type PublicVirtualInterfaceArgs struct {
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
	// The name for the virtual interface.
	Name pulumi.StringPtrInput
	// A list of routes to be advertised to the AWS network in this region.
	RouteFilterPrefixes pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VLAN ID.
	Vlan pulumi.IntInput
}

func (PublicVirtualInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicVirtualInterfaceArgs)(nil)).Elem()
}
