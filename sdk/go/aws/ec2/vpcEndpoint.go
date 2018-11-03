// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a VPC Endpoint resource.
// 
// ~> **NOTE on VPC Endpoints and VPC Endpoint Associations:** Terraform provides both standalone VPC Endpoint Associations for
// Route Tables - (an association between a VPC endpoint and a single `route_table_id`) and
// Subnets - (an association between a VPC endpoint and a single `subnet_id`) and
// a VPC Endpoint resource with `route_table_ids` and `subnet_ids` attributes.
// Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
// Doing so will cause a conflict of associations and will overwrite the association.
type VpcEndpoint struct {
	s *pulumi.ResourceState
}

// NewVpcEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpoint(ctx *pulumi.Context,
	name string, args *VpcEndpointArgs, opts ...pulumi.ResourceOpt) (*VpcEndpoint, error) {
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["autoAccept"] = nil
		inputs["policy"] = nil
		inputs["privateDnsEnabled"] = nil
		inputs["routeTableIds"] = nil
		inputs["securityGroupIds"] = nil
		inputs["serviceName"] = nil
		inputs["subnetIds"] = nil
		inputs["vpcEndpointType"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["autoAccept"] = args.AutoAccept
		inputs["policy"] = args.Policy
		inputs["privateDnsEnabled"] = args.PrivateDnsEnabled
		inputs["routeTableIds"] = args.RouteTableIds
		inputs["securityGroupIds"] = args.SecurityGroupIds
		inputs["serviceName"] = args.ServiceName
		inputs["subnetIds"] = args.SubnetIds
		inputs["vpcEndpointType"] = args.VpcEndpointType
		inputs["vpcId"] = args.VpcId
	}
	inputs["cidrBlocks"] = nil
	inputs["dnsEntries"] = nil
	inputs["networkInterfaceIds"] = nil
	inputs["prefixListId"] = nil
	inputs["state"] = nil
	s, err := ctx.RegisterResource("aws:ec2/vpcEndpoint:VpcEndpoint", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcEndpoint{s: s}, nil
}

// GetVpcEndpoint gets an existing VpcEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpoint(ctx *pulumi.Context,
	name string, id pulumi.ID, state *VpcEndpointState, opts ...pulumi.ResourceOpt) (*VpcEndpoint, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["autoAccept"] = state.AutoAccept
		inputs["cidrBlocks"] = state.CidrBlocks
		inputs["dnsEntries"] = state.DnsEntries
		inputs["networkInterfaceIds"] = state.NetworkInterfaceIds
		inputs["policy"] = state.Policy
		inputs["prefixListId"] = state.PrefixListId
		inputs["privateDnsEnabled"] = state.PrivateDnsEnabled
		inputs["routeTableIds"] = state.RouteTableIds
		inputs["securityGroupIds"] = state.SecurityGroupIds
		inputs["serviceName"] = state.ServiceName
		inputs["state"] = state.State
		inputs["subnetIds"] = state.SubnetIds
		inputs["vpcEndpointType"] = state.VpcEndpointType
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/vpcEndpoint:VpcEndpoint", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &VpcEndpoint{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *VpcEndpoint) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *VpcEndpoint) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
func (r *VpcEndpoint) AutoAccept() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["autoAccept"])
}

// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
func (r *VpcEndpoint) CidrBlocks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["cidrBlocks"])
}

// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
func (r *VpcEndpoint) DnsEntries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["dnsEntries"])
}

// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
func (r *VpcEndpoint) NetworkInterfaceIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["networkInterfaceIds"])
}

// A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
func (r *VpcEndpoint) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
func (r *VpcEndpoint) PrefixListId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["prefixListId"])
}

// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
// Defaults to `false`.
func (r *VpcEndpoint) PrivateDnsEnabled() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["privateDnsEnabled"])
}

// One or more route table IDs. Applicable for endpoints of type `Gateway`.
func (r *VpcEndpoint) RouteTableIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["routeTableIds"])
}

// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
func (r *VpcEndpoint) SecurityGroupIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["securityGroupIds"])
}

// The service name, in the form `com.amazonaws.region.service` for AWS services.
func (r *VpcEndpoint) ServiceName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["serviceName"])
}

// The state of the VPC endpoint.
func (r *VpcEndpoint) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
func (r *VpcEndpoint) SubnetIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["subnetIds"])
}

// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
func (r *VpcEndpoint) VpcEndpointType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcEndpointType"])
}

// The ID of the VPC in which the endpoint will be used.
func (r *VpcEndpoint) VpcId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering VpcEndpoint resources.
type VpcEndpointState struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept interface{}
	// The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
	CidrBlocks interface{}
	// The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
	DnsEntries interface{}
	// One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
	NetworkInterfaceIds interface{}
	// A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
	Policy interface{}
	// The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
	PrefixListId interface{}
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled interface{}
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds interface{}
	// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
	SecurityGroupIds interface{}
	// The service name, in the form `com.amazonaws.region.service` for AWS services.
	ServiceName interface{}
	// The state of the VPC endpoint.
	State interface{}
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
	SubnetIds interface{}
	// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
	VpcEndpointType interface{}
	// The ID of the VPC in which the endpoint will be used.
	VpcId interface{}
}

// The set of arguments for constructing a VpcEndpoint resource.
type VpcEndpointArgs struct {
	// Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
	AutoAccept interface{}
	// A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
	Policy interface{}
	// Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
	// Defaults to `false`.
	PrivateDnsEnabled interface{}
	// One or more route table IDs. Applicable for endpoints of type `Gateway`.
	RouteTableIds interface{}
	// The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
	SecurityGroupIds interface{}
	// The service name, in the form `com.amazonaws.region.service` for AWS services.
	ServiceName interface{}
	// The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
	SubnetIds interface{}
	// The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
	VpcEndpointType interface{}
	// The ID of the VPC in which the endpoint will be used.
	VpcId interface{}
}
