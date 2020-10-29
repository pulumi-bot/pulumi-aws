// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Direct Connect BGP peer resource.
type BgpPeer struct {
	pulumi.CustomResourceState

	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringOutput `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress pulumi.StringOutput `pulumi:"amazonAddress"`
	// The Direct Connect endpoint on which the BGP peer terminates.
	AwsDevice pulumi.StringOutput `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntOutput `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringOutput `pulumi:"bgpAuthKey"`
	// The ID of the BGP peer.
	BgpPeerId pulumi.StringOutput `pulumi:"bgpPeerId"`
	// The Up/Down state of the BGP peer.
	BgpStatus pulumi.StringOutput `pulumi:"bgpStatus"`
	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress pulumi.StringOutput `pulumi:"customerAddress"`
	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId pulumi.StringOutput `pulumi:"virtualInterfaceId"`
}

// NewBgpPeer registers a new resource with the given unique name, arguments, and options.
func NewBgpPeer(ctx *pulumi.Context,
	name string, args *BgpPeerArgs, opts ...pulumi.ResourceOption) (*BgpPeer, error) {
	if args == nil || args.AddressFamily == nil {
		return nil, errors.New("missing required argument 'AddressFamily'")
	}
	if args == nil || args.BgpAsn == nil {
		return nil, errors.New("missing required argument 'BgpAsn'")
	}
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	if args == nil {
		args = &BgpPeerArgs{}
	}
	var resource BgpPeer
	err := ctx.RegisterResource("aws:directconnect/bgpPeer:BgpPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBgpPeer gets an existing BgpPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBgpPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BgpPeerState, opts ...pulumi.ResourceOption) (*BgpPeer, error) {
	var resource BgpPeer
	err := ctx.ReadResource("aws:directconnect/bgpPeer:BgpPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BgpPeer resources.
type bgpPeerState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily *string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress *string `pulumi:"amazonAddress"`
	// The Direct Connect endpoint on which the BGP peer terminates.
	AwsDevice *string `pulumi:"awsDevice"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn *int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The ID of the BGP peer.
	BgpPeerId *string `pulumi:"bgpPeerId"`
	// The Up/Down state of the BGP peer.
	BgpStatus *string `pulumi:"bgpStatus"`
	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress *string `pulumi:"customerAddress"`
	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
}

type BgpPeerState struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringPtrInput
	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress pulumi.StringPtrInput
	// The Direct Connect endpoint on which the BGP peer terminates.
	AwsDevice pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntPtrInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The ID of the BGP peer.
	BgpPeerId pulumi.StringPtrInput
	// The Up/Down state of the BGP peer.
	BgpStatus pulumi.StringPtrInput
	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress pulumi.StringPtrInput
	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId pulumi.StringPtrInput
}

func (BgpPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerState)(nil)).Elem()
}

type bgpPeerArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily string `pulumi:"addressFamily"`
	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress *string `pulumi:"amazonAddress"`
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn int `pulumi:"bgpAsn"`
	// The authentication key for BGP configuration.
	BgpAuthKey *string `pulumi:"bgpAuthKey"`
	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress *string `pulumi:"customerAddress"`
	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId string `pulumi:"virtualInterfaceId"`
}

// The set of arguments for constructing a BgpPeer resource.
type BgpPeerArgs struct {
	// The address family for the BGP peer. ` ipv4  ` or `ipv6`.
	AddressFamily pulumi.StringInput
	// The IPv4 CIDR address to use to send traffic to Amazon.
	// Required for IPv4 BGP peers on public virtual interfaces.
	AmazonAddress pulumi.StringPtrInput
	// The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
	BgpAsn pulumi.IntInput
	// The authentication key for BGP configuration.
	BgpAuthKey pulumi.StringPtrInput
	// The IPv4 CIDR destination address to which Amazon should send traffic.
	// Required for IPv4 BGP peers on public virtual interfaces.
	CustomerAddress pulumi.StringPtrInput
	// The ID of the Direct Connect virtual interface on which to create the BGP peer.
	VirtualInterfaceId pulumi.StringInput
}

func (BgpPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerArgs)(nil)).Elem()
}
