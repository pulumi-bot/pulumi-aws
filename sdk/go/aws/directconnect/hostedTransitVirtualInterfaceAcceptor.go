// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type HostedTransitVirtualInterfaceAcceptor struct {
	pulumi.CustomResourceState

	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringOutput `pulumi:"dxGatewayId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringOutput `pulumi:"virtualInterfaceId"`
}

// NewHostedTransitVirtualInterfaceAcceptor registers a new resource with the given unique name, arguments, and options.
func NewHostedTransitVirtualInterfaceAcceptor(ctx *pulumi.Context,
	name string, args *HostedTransitVirtualInterfaceAcceptorArgs, opts ...pulumi.ResourceOption) (*HostedTransitVirtualInterfaceAcceptor, error) {
	if args == nil || args.DxGatewayId == nil {
		return nil, errors.New("missing required argument 'DxGatewayId'")
	}
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	if args == nil {
		args = &HostedTransitVirtualInterfaceAcceptorArgs{}
	}
	var resource HostedTransitVirtualInterfaceAcceptor
	err := ctx.RegisterResource("aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedTransitVirtualInterfaceAcceptor gets an existing HostedTransitVirtualInterfaceAcceptor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedTransitVirtualInterfaceAcceptor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedTransitVirtualInterfaceAcceptorState, opts ...pulumi.ResourceOption) (*HostedTransitVirtualInterfaceAcceptor, error) {
	var resource HostedTransitVirtualInterfaceAcceptor
	err := ctx.ReadResource("aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedTransitVirtualInterfaceAcceptor resources.
type hostedTransitVirtualInterfaceAcceptorState struct {
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
}

type HostedTransitVirtualInterfaceAcceptorState struct {
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringPtrInput
}

func (HostedTransitVirtualInterfaceAcceptorState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedTransitVirtualInterfaceAcceptorState)(nil)).Elem()
}

type hostedTransitVirtualInterfaceAcceptorArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId string `pulumi:"dxGatewayId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId string `pulumi:"virtualInterfaceId"`
}

// The set of arguments for constructing a HostedTransitVirtualInterfaceAcceptor resource.
type HostedTransitVirtualInterfaceAcceptorArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringInput
}

func (HostedTransitVirtualInterfaceAcceptorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedTransitVirtualInterfaceAcceptorArgs)(nil)).Elem()
}

