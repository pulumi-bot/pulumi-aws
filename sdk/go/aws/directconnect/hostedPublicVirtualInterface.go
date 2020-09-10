// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type HostedPublicVirtualInterface struct {
	pulumi.CustomResourceState

	AddressFamily       pulumi.StringOutput      `pulumi:"addressFamily"`
	AmazonAddress       pulumi.StringOutput      `pulumi:"amazonAddress"`
	AmazonSideAsn       pulumi.StringOutput      `pulumi:"amazonSideAsn"`
	Arn                 pulumi.StringOutput      `pulumi:"arn"`
	AwsDevice           pulumi.StringOutput      `pulumi:"awsDevice"`
	BgpAsn              pulumi.IntOutput         `pulumi:"bgpAsn"`
	BgpAuthKey          pulumi.StringOutput      `pulumi:"bgpAuthKey"`
	ConnectionId        pulumi.StringOutput      `pulumi:"connectionId"`
	CustomerAddress     pulumi.StringOutput      `pulumi:"customerAddress"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	OwnerAccountId      pulumi.StringOutput      `pulumi:"ownerAccountId"`
	RouteFilterPrefixes pulumi.StringArrayOutput `pulumi:"routeFilterPrefixes"`
	Vlan                pulumi.IntOutput         `pulumi:"vlan"`
}

// NewHostedPublicVirtualInterface registers a new resource with the given unique name, arguments, and options.
func NewHostedPublicVirtualInterface(ctx *pulumi.Context,
	name string, args *HostedPublicVirtualInterfaceArgs, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterface, error) {
	if args == nil || args.AddressFamily == nil {
		return nil, errors.New("missing required argument 'AddressFamily'")
	}
	if args == nil || args.BgpAsn == nil {
		return nil, errors.New("missing required argument 'BgpAsn'")
	}
	if args == nil || args.ConnectionId == nil {
		return nil, errors.New("missing required argument 'ConnectionId'")
	}
	if args == nil || args.OwnerAccountId == nil {
		return nil, errors.New("missing required argument 'OwnerAccountId'")
	}
	if args == nil || args.RouteFilterPrefixes == nil {
		return nil, errors.New("missing required argument 'RouteFilterPrefixes'")
	}
	if args == nil || args.Vlan == nil {
		return nil, errors.New("missing required argument 'Vlan'")
	}
	if args == nil {
		args = &HostedPublicVirtualInterfaceArgs{}
	}
	var resource HostedPublicVirtualInterface
	err := ctx.RegisterResource("aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPublicVirtualInterface gets an existing HostedPublicVirtualInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPublicVirtualInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPublicVirtualInterfaceState, opts ...pulumi.ResourceOption) (*HostedPublicVirtualInterface, error) {
	var resource HostedPublicVirtualInterface
	err := ctx.ReadResource("aws:directconnect/hostedPublicVirtualInterface:HostedPublicVirtualInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPublicVirtualInterface resources.
type hostedPublicVirtualInterfaceState struct {
	AddressFamily       *string  `pulumi:"addressFamily"`
	AmazonAddress       *string  `pulumi:"amazonAddress"`
	AmazonSideAsn       *string  `pulumi:"amazonSideAsn"`
	Arn                 *string  `pulumi:"arn"`
	AwsDevice           *string  `pulumi:"awsDevice"`
	BgpAsn              *int     `pulumi:"bgpAsn"`
	BgpAuthKey          *string  `pulumi:"bgpAuthKey"`
	ConnectionId        *string  `pulumi:"connectionId"`
	CustomerAddress     *string  `pulumi:"customerAddress"`
	Name                *string  `pulumi:"name"`
	OwnerAccountId      *string  `pulumi:"ownerAccountId"`
	RouteFilterPrefixes []string `pulumi:"routeFilterPrefixes"`
	Vlan                *int     `pulumi:"vlan"`
}

type HostedPublicVirtualInterfaceState struct {
	AddressFamily       pulumi.StringPtrInput
	AmazonAddress       pulumi.StringPtrInput
	AmazonSideAsn       pulumi.StringPtrInput
	Arn                 pulumi.StringPtrInput
	AwsDevice           pulumi.StringPtrInput
	BgpAsn              pulumi.IntPtrInput
	BgpAuthKey          pulumi.StringPtrInput
	ConnectionId        pulumi.StringPtrInput
	CustomerAddress     pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	OwnerAccountId      pulumi.StringPtrInput
	RouteFilterPrefixes pulumi.StringArrayInput
	Vlan                pulumi.IntPtrInput
}

func (HostedPublicVirtualInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceState)(nil)).Elem()
}

type hostedPublicVirtualInterfaceArgs struct {
	AddressFamily       string   `pulumi:"addressFamily"`
	AmazonAddress       *string  `pulumi:"amazonAddress"`
	BgpAsn              int      `pulumi:"bgpAsn"`
	BgpAuthKey          *string  `pulumi:"bgpAuthKey"`
	ConnectionId        string   `pulumi:"connectionId"`
	CustomerAddress     *string  `pulumi:"customerAddress"`
	Name                *string  `pulumi:"name"`
	OwnerAccountId      string   `pulumi:"ownerAccountId"`
	RouteFilterPrefixes []string `pulumi:"routeFilterPrefixes"`
	Vlan                int      `pulumi:"vlan"`
}

// The set of arguments for constructing a HostedPublicVirtualInterface resource.
type HostedPublicVirtualInterfaceArgs struct {
	AddressFamily       pulumi.StringInput
	AmazonAddress       pulumi.StringPtrInput
	BgpAsn              pulumi.IntInput
	BgpAuthKey          pulumi.StringPtrInput
	ConnectionId        pulumi.StringInput
	CustomerAddress     pulumi.StringPtrInput
	Name                pulumi.StringPtrInput
	OwnerAccountId      pulumi.StringInput
	RouteFilterPrefixes pulumi.StringArrayInput
	Vlan                pulumi.IntInput
}

func (HostedPublicVirtualInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPublicVirtualInterfaceArgs)(nil)).Elem()
}
