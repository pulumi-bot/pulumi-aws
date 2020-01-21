// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package hostedPrivateVirtualInterfaceAccepter

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the accepter's side of a Direct Connect hosted private virtual interface.
// This resource accepts ownership of a private virtual interface created by another AWS account.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_hosted_private_virtual_interface_accepter.html.markdown.
type HostedPrivateVirtualInterfaceAccepter struct {
	pulumi.CustomResourceState

	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrOutput `pulumi:"dxGatewayId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringOutput `pulumi:"virtualInterfaceId"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrOutput `pulumi:"vpnGatewayId"`
}

// NewHostedPrivateVirtualInterfaceAccepter registers a new resource with the given unique name, arguments, and options.
func NewHostedPrivateVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, args *HostedPrivateVirtualInterfaceAccepterArgs, opts ...pulumi.ResourceOption) (*HostedPrivateVirtualInterfaceAccepter, error) {
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	if args == nil {
		args = &HostedPrivateVirtualInterfaceAccepterArgs{}
	}
	var resource HostedPrivateVirtualInterfaceAccepter
	err := ctx.RegisterResource("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedPrivateVirtualInterfaceAccepter gets an existing HostedPrivateVirtualInterfaceAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPrivateVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedPrivateVirtualInterfaceAccepterState, opts ...pulumi.ResourceOption) (*HostedPrivateVirtualInterfaceAccepter, error) {
	var resource HostedPrivateVirtualInterfaceAccepter
	err := ctx.ReadResource("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedPrivateVirtualInterfaceAccepter resources.
type hostedPrivateVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

type HostedPrivateVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringPtrInput
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrInput
}

func (HostedPrivateVirtualInterfaceAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPrivateVirtualInterfaceAccepterState)(nil)).Elem()
}

type hostedPrivateVirtualInterfaceAccepterArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId string `pulumi:"virtualInterfaceId"`
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
}

// The set of arguments for constructing a HostedPrivateVirtualInterfaceAccepter resource.
type HostedPrivateVirtualInterfaceAccepterArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringInput
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId pulumi.StringPtrInput
}

func (HostedPrivateVirtualInterfaceAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedPrivateVirtualInterfaceAccepterArgs)(nil)).Elem()
}

