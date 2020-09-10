// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type BgpPeer struct {
	pulumi.CustomResourceState

	AddressFamily      pulumi.StringOutput `pulumi:"addressFamily"`
	AmazonAddress      pulumi.StringOutput `pulumi:"amazonAddress"`
	AwsDevice          pulumi.StringOutput `pulumi:"awsDevice"`
	BgpAsn             pulumi.IntOutput    `pulumi:"bgpAsn"`
	BgpAuthKey         pulumi.StringOutput `pulumi:"bgpAuthKey"`
	BgpPeerId          pulumi.StringOutput `pulumi:"bgpPeerId"`
	BgpStatus          pulumi.StringOutput `pulumi:"bgpStatus"`
	CustomerAddress    pulumi.StringOutput `pulumi:"customerAddress"`
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
	AddressFamily      *string `pulumi:"addressFamily"`
	AmazonAddress      *string `pulumi:"amazonAddress"`
	AwsDevice          *string `pulumi:"awsDevice"`
	BgpAsn             *int    `pulumi:"bgpAsn"`
	BgpAuthKey         *string `pulumi:"bgpAuthKey"`
	BgpPeerId          *string `pulumi:"bgpPeerId"`
	BgpStatus          *string `pulumi:"bgpStatus"`
	CustomerAddress    *string `pulumi:"customerAddress"`
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
}

type BgpPeerState struct {
	AddressFamily      pulumi.StringPtrInput
	AmazonAddress      pulumi.StringPtrInput
	AwsDevice          pulumi.StringPtrInput
	BgpAsn             pulumi.IntPtrInput
	BgpAuthKey         pulumi.StringPtrInput
	BgpPeerId          pulumi.StringPtrInput
	BgpStatus          pulumi.StringPtrInput
	CustomerAddress    pulumi.StringPtrInput
	VirtualInterfaceId pulumi.StringPtrInput
}

func (BgpPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerState)(nil)).Elem()
}

type bgpPeerArgs struct {
	AddressFamily      string  `pulumi:"addressFamily"`
	AmazonAddress      *string `pulumi:"amazonAddress"`
	BgpAsn             int     `pulumi:"bgpAsn"`
	BgpAuthKey         *string `pulumi:"bgpAuthKey"`
	CustomerAddress    *string `pulumi:"customerAddress"`
	VirtualInterfaceId string  `pulumi:"virtualInterfaceId"`
}

// The set of arguments for constructing a BgpPeer resource.
type BgpPeerArgs struct {
	AddressFamily      pulumi.StringInput
	AmazonAddress      pulumi.StringPtrInput
	BgpAsn             pulumi.IntInput
	BgpAuthKey         pulumi.StringPtrInput
	CustomerAddress    pulumi.StringPtrInput
	VirtualInterfaceId pulumi.StringInput
}

func (BgpPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bgpPeerArgs)(nil)).Elem()
}
