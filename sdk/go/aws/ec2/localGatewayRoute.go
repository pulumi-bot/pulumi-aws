// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type LocalGatewayRoute struct {
	pulumi.CustomResourceState

	DestinationCidrBlock                pulumi.StringOutput `pulumi:"destinationCidrBlock"`
	LocalGatewayRouteTableId            pulumi.StringOutput `pulumi:"localGatewayRouteTableId"`
	LocalGatewayVirtualInterfaceGroupId pulumi.StringOutput `pulumi:"localGatewayVirtualInterfaceGroupId"`
}

// NewLocalGatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewLocalGatewayRoute(ctx *pulumi.Context,
	name string, args *LocalGatewayRouteArgs, opts ...pulumi.ResourceOption) (*LocalGatewayRoute, error) {
	if args == nil || args.DestinationCidrBlock == nil {
		return nil, errors.New("missing required argument 'DestinationCidrBlock'")
	}
	if args == nil || args.LocalGatewayRouteTableId == nil {
		return nil, errors.New("missing required argument 'LocalGatewayRouteTableId'")
	}
	if args == nil || args.LocalGatewayVirtualInterfaceGroupId == nil {
		return nil, errors.New("missing required argument 'LocalGatewayVirtualInterfaceGroupId'")
	}
	if args == nil {
		args = &LocalGatewayRouteArgs{}
	}
	var resource LocalGatewayRoute
	err := ctx.RegisterResource("aws:ec2/localGatewayRoute:LocalGatewayRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGatewayRoute gets an existing LocalGatewayRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGatewayRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGatewayRouteState, opts ...pulumi.ResourceOption) (*LocalGatewayRoute, error) {
	var resource LocalGatewayRoute
	err := ctx.ReadResource("aws:ec2/localGatewayRoute:LocalGatewayRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGatewayRoute resources.
type localGatewayRouteState struct {
	DestinationCidrBlock                *string `pulumi:"destinationCidrBlock"`
	LocalGatewayRouteTableId            *string `pulumi:"localGatewayRouteTableId"`
	LocalGatewayVirtualInterfaceGroupId *string `pulumi:"localGatewayVirtualInterfaceGroupId"`
}

type LocalGatewayRouteState struct {
	DestinationCidrBlock                pulumi.StringPtrInput
	LocalGatewayRouteTableId            pulumi.StringPtrInput
	LocalGatewayVirtualInterfaceGroupId pulumi.StringPtrInput
}

func (LocalGatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteState)(nil)).Elem()
}

type localGatewayRouteArgs struct {
	DestinationCidrBlock                string `pulumi:"destinationCidrBlock"`
	LocalGatewayRouteTableId            string `pulumi:"localGatewayRouteTableId"`
	LocalGatewayVirtualInterfaceGroupId string `pulumi:"localGatewayVirtualInterfaceGroupId"`
}

// The set of arguments for constructing a LocalGatewayRoute resource.
type LocalGatewayRouteArgs struct {
	DestinationCidrBlock                pulumi.StringInput
	LocalGatewayRouteTableId            pulumi.StringInput
	LocalGatewayVirtualInterfaceGroupId pulumi.StringInput
}

func (LocalGatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteArgs)(nil)).Elem()
}
