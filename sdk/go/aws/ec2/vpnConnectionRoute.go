// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a static route between a VPN connection and a customer gateway.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpn_connection_route.html.markdown.
type VpnConnectionRoute struct {
	pulumi.CustomResourceState

	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	// The ID of the VPN connection.
	VpnConnectionId pulumi.StringOutput `pulumi:"vpnConnectionId"`
}

// NewVpnConnectionRoute registers a new resource with the given unique name, arguments, and options.
func NewVpnConnectionRoute(ctx *pulumi.Context,
	name string, args *VpnConnectionRouteArgs, opts ...pulumi.ResourceOption) (*VpnConnectionRoute, error) {
	if args == nil || args.DestinationCidrBlock == nil {
		return nil, errors.New("missing required argument 'DestinationCidrBlock'")
	}
	if args == nil || args.VpnConnectionId == nil {
		return nil, errors.New("missing required argument 'VpnConnectionId'")
	}
	if args == nil {
		args = &VpnConnectionRouteArgs{}
	}
	var resource VpnConnectionRoute
	err := ctx.RegisterResource("aws:ec2/vpnConnectionRoute:VpnConnectionRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnConnectionRoute gets an existing VpnConnectionRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnConnectionRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnConnectionRouteState, opts ...pulumi.ResourceOption) (*VpnConnectionRoute, error) {
	var resource VpnConnectionRoute
	err := ctx.ReadResource("aws:ec2/vpnConnectionRoute:VpnConnectionRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnConnectionRoute resources.
type vpnConnectionRouteState struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	// The ID of the VPN connection.
	VpnConnectionId *string `pulumi:"vpnConnectionId"`
}

type VpnConnectionRouteState struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock pulumi.StringPtrInput
	// The ID of the VPN connection.
	VpnConnectionId pulumi.StringPtrInput
}

func (VpnConnectionRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionRouteState)(nil)).Elem()
}

type vpnConnectionRouteArgs struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock string `pulumi:"destinationCidrBlock"`
	// The ID of the VPN connection.
	VpnConnectionId string `pulumi:"vpnConnectionId"`
}

// The set of arguments for constructing a VpnConnectionRoute resource.
type VpnConnectionRouteArgs struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock pulumi.StringInput
	// The ID of the VPN connection.
	VpnConnectionId pulumi.StringInput
}

func (VpnConnectionRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnConnectionRouteArgs)(nil)).Elem()
}

