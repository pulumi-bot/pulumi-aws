// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the accepter's side of a VPC Peering Connection.
// 
// When a cross-account (requester's AWS account differs from the accepter's AWS account) or an inter-region
// VPC Peering Connection is created, a VPC Peering Connection resource is automatically created in the
// accepter's account.
// The requester can use the `ec2.VpcPeeringConnection` resource to manage its side of the connection
// and the accepter can use the `ec2.VpcPeeringConnectionAccepter` resource to "adopt" its side of the
// connection into management.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_peering_connection_accepter.html.markdown.
type VpcPeeringConnectionAccepter struct {
	s *pulumi.ResourceState
}

// NewVpcPeeringConnectionAccepter registers a new resource with the given unique name, arguments, and options.
func NewVpcPeeringConnectionAccepter(ctx *pulumi.Context,
	name string, args *VpcPeeringConnectionAccepterArgs, opts ...pulumi.ResourceOpt) (*VpcPeeringConnectionAccepter, error) {
	if args == nil || args.VpcPeeringConnectionId == nil {
		return nil, errors.New("missing required argument 'VpcPeeringConnectionId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["accepter"] = nil
		inputs["autoAccept"] = nil
		inputs["requester"] = nil
		inputs["tags"] = nil
		inputs["vpcPeeringConnectionId"] = nil
	} else {
		inputs["accepter"] = args.Accepter
		inputs["autoAccept"] = args.AutoAccept
		inputs["requester"] = args.Requester
		inputs["tags"] = args.Tags
		inputs["vpcPeeringConnectionId"] = args.VpcPeeringConnectionId
	}
	inputs["acceptStatus"] = nil
	inputs["peerOwnerId"] = nil
	inputs["peerRegion"] = nil
	inputs["peerVpcId"] = nil
	inputs["vpcId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcPeeringConnectionAccepter{s: s}, nil
}

// GetVpcPeeringConnectionAccepter gets an existing VpcPeeringConnectionAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcPeeringConnectionAccepter(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcPeeringConnectionAccepterState, opts ...pulumi.ResourceOpt) (*VpcPeeringConnectionAccepter, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["acceptStatus"] = state.AcceptStatus
		inputs["accepter"] = state.Accepter
		inputs["autoAccept"] = state.AutoAccept
		inputs["peerOwnerId"] = state.PeerOwnerId
		inputs["peerRegion"] = state.PeerRegion
		inputs["peerVpcId"] = state.PeerVpcId
		inputs["requester"] = state.Requester
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
		inputs["vpcPeeringConnectionId"] = state.VpcPeeringConnectionId
	}
	s, err := ctx.ReadResource("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcPeeringConnectionAccepter{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VpcPeeringConnectionAccepter) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VpcPeeringConnectionAccepter) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The status of the VPC Peering Connection request.
func (r *VpcPeeringConnectionAccepter) AcceptStatus() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["acceptStatus"])
}

// A configuration block that describes [VPC Peering Connection]
// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
func (r *VpcPeeringConnectionAccepter) Accepter() *pulumi.Output {
	return r.s.State["accepter"]
}

// Whether or not to accept the peering request. Defaults to `false`.
func (r *VpcPeeringConnectionAccepter) AutoAccept() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoAccept"])
}

// The AWS account ID of the owner of the requester VPC.
func (r *VpcPeeringConnectionAccepter) PeerOwnerId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerOwnerId"])
}

// The region of the accepter VPC.
func (r *VpcPeeringConnectionAccepter) PeerRegion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerRegion"])
}

// The ID of the requester VPC.
func (r *VpcPeeringConnectionAccepter) PeerVpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["peerVpcId"])
}

// A configuration block that describes [VPC Peering Connection]
// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
func (r *VpcPeeringConnectionAccepter) Requester() *pulumi.Output {
	return r.s.State["requester"]
}

// A mapping of tags to assign to the resource.
func (r *VpcPeeringConnectionAccepter) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The ID of the accepter VPC.
func (r *VpcPeeringConnectionAccepter) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// The VPC Peering Connection ID to manage.
func (r *VpcPeeringConnectionAccepter) VpcPeeringConnectionId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcPeeringConnectionId"])
}

// Input properties used for looking up and filtering VpcPeeringConnectionAccepter resources.
type VpcPeeringConnectionAccepterState struct {
	// The status of the VPC Peering Connection request.
	AcceptStatus interface{}
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
	Accepter interface{}
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept interface{}
	// The AWS account ID of the owner of the requester VPC.
	PeerOwnerId interface{}
	// The region of the accepter VPC.
	PeerRegion interface{}
	// The ID of the requester VPC.
	PeerVpcId interface{}
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
	Requester interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The ID of the accepter VPC.
	VpcId interface{}
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId interface{}
}

// The set of arguments for constructing a VpcPeeringConnectionAccepter resource.
type VpcPeeringConnectionAccepterArgs struct {
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the accepter VPC.
	Accepter interface{}
	// Whether or not to accept the peering request. Defaults to `false`.
	AutoAccept interface{}
	// A configuration block that describes [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options set for the requester VPC.
	Requester interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The VPC Peering Connection ID to manage.
	VpcPeeringConnectionId interface{}
}
