// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the accepter's side of a Direct Connect hosted private virtual interface.
// This resource accepts ownership of a private virtual interface created by another AWS account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_hosted_private_virtual_interface_accepter.html.markdown.
type HostedPrivateVirtualInterfaceAccepter struct {
	s *pulumi.ResourceState
}

// NewHostedPrivateVirtualInterfaceAccepter registers a new resource with the given unique name, arguments, and options.
func NewHostedPrivateVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, args *HostedPrivateVirtualInterfaceAccepterArgs, opts ...pulumi.ResourceOpt) (*HostedPrivateVirtualInterfaceAccepter, error) {
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["dxGatewayId"] = nil
		inputs["tags"] = nil
		inputs["virtualInterfaceId"] = nil
		inputs["vpnGatewayId"] = nil
	} else {
		inputs["dxGatewayId"] = args.DxGatewayId
		inputs["tags"] = args.Tags
		inputs["virtualInterfaceId"] = args.VirtualInterfaceId
		inputs["vpnGatewayId"] = args.VpnGatewayId
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HostedPrivateVirtualInterfaceAccepter{s: s}, nil
}

// GetHostedPrivateVirtualInterfaceAccepter gets an existing HostedPrivateVirtualInterfaceAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedPrivateVirtualInterfaceAccepter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *HostedPrivateVirtualInterfaceAccepterState, opts ...pulumi.ResourceOpt) (*HostedPrivateVirtualInterfaceAccepter, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["dxGatewayId"] = state.DxGatewayId
		inputs["tags"] = state.Tags
		inputs["virtualInterfaceId"] = state.VirtualInterfaceId
		inputs["vpnGatewayId"] = state.VpnGatewayId
	}
	s, err := ctx.ReadResource("aws:directconnect/hostedPrivateVirtualInterfaceAccepter:HostedPrivateVirtualInterfaceAccepter", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &HostedPrivateVirtualInterfaceAccepter{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *HostedPrivateVirtualInterfaceAccepter) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *HostedPrivateVirtualInterfaceAccepter) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the virtual interface.
func (r *HostedPrivateVirtualInterfaceAccepter) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The ID of the Direct Connect gateway to which to connect the virtual interface.
func (r *HostedPrivateVirtualInterfaceAccepter) DxGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["dxGatewayId"])
}

// A mapping of tags to assign to the resource.
func (r *HostedPrivateVirtualInterfaceAccepter) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The ID of the Direct Connect virtual interface to accept.
func (r *HostedPrivateVirtualInterfaceAccepter) VirtualInterfaceId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["virtualInterfaceId"])
}

// The ID of the virtual private gateway to which to connect the virtual interface.
func (r *HostedPrivateVirtualInterfaceAccepter) VpnGatewayId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpnGatewayId"])
}

// Input properties used for looking up and filtering HostedPrivateVirtualInterfaceAccepter resources.
type HostedPrivateVirtualInterfaceAccepterState struct {
	// The ARN of the virtual interface.
	Arn interface{}
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId interface{}
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId interface{}
}

// The set of arguments for constructing a HostedPrivateVirtualInterfaceAccepter resource.
type HostedPrivateVirtualInterfaceAccepterArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId interface{}
	// The ID of the virtual private gateway to which to connect the virtual interface.
	VpnGatewayId interface{}
}
