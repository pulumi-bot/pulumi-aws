// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Route 53 Resolver endpoint resource.
type ResolverEndpoint struct {
	pulumi.CustomResourceState

	// The ARN of the Route 53 Resolver endpoint.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringOutput `pulumi:"direction"`
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId pulumi.StringOutput `pulumi:"hostVpcId"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressArrayOutput `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewResolverEndpoint registers a new resource with the given unique name, arguments, and options.
func NewResolverEndpoint(ctx *pulumi.Context,
	name string, args *ResolverEndpointArgs, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	if args == nil || args.Direction == nil {
		return nil, errors.New("missing required argument 'Direction'")
	}
	if args == nil || args.IpAddresses == nil {
		return nil, errors.New("missing required argument 'IpAddresses'")
	}
	if args == nil || args.SecurityGroupIds == nil {
		return nil, errors.New("missing required argument 'SecurityGroupIds'")
	}
	if args == nil {
		args = &ResolverEndpointArgs{}
	}
	var resource ResolverEndpoint
	err := ctx.RegisterResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverEndpoint gets an existing ResolverEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverEndpointState, opts ...pulumi.ResourceOption) (*ResolverEndpoint, error) {
	var resource ResolverEndpoint
	err := ctx.ReadResource("aws:route53/resolverEndpoint:ResolverEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverEndpoint resources.
type resolverEndpointState struct {
	// The ARN of the Route 53 Resolver endpoint.
	Arn *string `pulumi:"arn"`
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction *string `pulumi:"direction"`
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId *string `pulumi:"hostVpcId"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses []ResolverEndpointIpAddress `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name *string `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ResolverEndpointState struct {
	// The ARN of the Route 53 Resolver endpoint.
	Arn pulumi.StringPtrInput
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringPtrInput
	// The ID of the VPC that you want to create the resolver endpoint in.
	HostVpcId pulumi.StringPtrInput
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressArrayInput
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringPtrInput
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ResolverEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointState)(nil)).Elem()
}

type resolverEndpointArgs struct {
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction string `pulumi:"direction"`
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses []ResolverEndpointIpAddress `pulumi:"ipAddresses"`
	// The friendly name of the Route 53 Resolver endpoint.
	Name *string `pulumi:"name"`
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ResolverEndpoint resource.
type ResolverEndpointArgs struct {
	// The direction of DNS queries to or from the Route 53 Resolver endpoint.
	// Valid values are `INBOUND` (resolver forwards DNS queries to the DNS service for a VPC from your network or another VPC)
	// or `OUTBOUND` (resolver forwards DNS queries from the DNS service for a VPC to your network or another VPC).
	Direction pulumi.StringInput
	// The subnets and IP addresses in your VPC that you want DNS queries to pass through on the way from your VPCs
	// to your network (for outbound endpoints) or on the way from your network to your VPCs (for inbound endpoints). Described below.
	IpAddresses ResolverEndpointIpAddressArrayInput
	// The friendly name of the Route 53 Resolver endpoint.
	Name pulumi.StringPtrInput
	// The ID of one or more security groups that you want to use to control access to this VPC.
	SecurityGroupIds pulumi.StringArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ResolverEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverEndpointArgs)(nil)).Elem()
}
