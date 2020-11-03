// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates an entry (a rule) in a network ACL with the specified rule number.
//
// > **NOTE on Network ACLs and Network ACL Rules:** This provider currently
// provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
// defined in-line. At this time you cannot use a Network ACL with in-line rules
// in conjunction with any Network ACL Rule resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		barNetworkAcl, err := ec2.NewNetworkAcl(ctx, "barNetworkAcl", &ec2.NetworkAclArgs{
// 			VpcId: pulumi.Any(aws_vpc.Foo.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewNetworkAclRule(ctx, "barNetworkAclRule", &ec2.NetworkAclRuleArgs{
// 			NetworkAclId: barNetworkAcl.ID(),
// 			RuleNumber:   pulumi.Int(200),
// 			Egress:       pulumi.Bool(false),
// 			Protocol:     pulumi.String("tcp"),
// 			RuleAction:   pulumi.String("allow"),
// 			CidrBlock:    pulumi.Any(aws_vpc.Foo.Cidr_block),
// 			FromPort:     pulumi.Int(22),
// 			ToPort:       pulumi.Int(22),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **Note:** One of either `cidrBlock` or `ipv6CidrBlock` is required.
type NetworkAclRule struct {
	pulumi.CustomResourceState

	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock pulumi.StringPtrOutput `pulumi:"cidrBlock"`
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress pulumi.BoolPtrOutput `pulumi:"egress"`
	// The from port to match.
	FromPort pulumi.IntPtrOutput `pulumi:"fromPort"`
	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
	IcmpCode pulumi.StringPtrOutput `pulumi:"icmpCode"`
	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
	IcmpType pulumi.StringPtrOutput `pulumi:"icmpType"`
	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock pulumi.StringPtrOutput `pulumi:"ipv6CidrBlock"`
	// The ID of the network ACL.
	NetworkAclId pulumi.StringOutput `pulumi:"networkAclId"`
	// The protocol. A value of -1 means all protocols.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber pulumi.IntOutput `pulumi:"ruleNumber"`
	// The to port to match.
	ToPort pulumi.IntPtrOutput `pulumi:"toPort"`
}

// NewNetworkAclRule registers a new resource with the given unique name, arguments, and options.
func NewNetworkAclRule(ctx *pulumi.Context,
	name string, args *NetworkAclRuleArgs, opts ...pulumi.ResourceOption) (*NetworkAclRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.NetworkAclId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkAclId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.RuleAction == nil {
		return nil, errors.New("invalid value for required argument 'RuleAction'")
	}
	if args.RuleNumber == nil {
		return nil, errors.New("invalid value for required argument 'RuleNumber'")
	}
	var resource NetworkAclRule
	err := ctx.RegisterResource("aws:ec2/networkAclRule:NetworkAclRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAclRule gets an existing NetworkAclRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAclRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAclRuleState, opts ...pulumi.ResourceOption) (*NetworkAclRule, error) {
	var resource NetworkAclRule
	err := ctx.ReadResource("aws:ec2/networkAclRule:NetworkAclRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAclRule resources.
type networkAclRuleState struct {
	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock *string `pulumi:"cidrBlock"`
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress *bool `pulumi:"egress"`
	// The from port to match.
	FromPort *int `pulumi:"fromPort"`
	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
	IcmpCode *string `pulumi:"icmpCode"`
	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
	IcmpType *string `pulumi:"icmpType"`
	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The ID of the network ACL.
	NetworkAclId *string `pulumi:"networkAclId"`
	// The protocol. A value of -1 means all protocols.
	Protocol *string `pulumi:"protocol"`
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction *string `pulumi:"ruleAction"`
	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber *int `pulumi:"ruleNumber"`
	// The to port to match.
	ToPort *int `pulumi:"toPort"`
}

type NetworkAclRuleState struct {
	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock pulumi.StringPtrInput
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress pulumi.BoolPtrInput
	// The from port to match.
	FromPort pulumi.IntPtrInput
	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
	IcmpCode pulumi.StringPtrInput
	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
	IcmpType pulumi.StringPtrInput
	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock pulumi.StringPtrInput
	// The ID of the network ACL.
	NetworkAclId pulumi.StringPtrInput
	// The protocol. A value of -1 means all protocols.
	Protocol pulumi.StringPtrInput
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction pulumi.StringPtrInput
	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber pulumi.IntPtrInput
	// The to port to match.
	ToPort pulumi.IntPtrInput
}

func (NetworkAclRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclRuleState)(nil)).Elem()
}

type networkAclRuleArgs struct {
	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock *string `pulumi:"cidrBlock"`
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress *bool `pulumi:"egress"`
	// The from port to match.
	FromPort *int `pulumi:"fromPort"`
	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
	IcmpCode *string `pulumi:"icmpCode"`
	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
	IcmpType *string `pulumi:"icmpType"`
	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The ID of the network ACL.
	NetworkAclId string `pulumi:"networkAclId"`
	// The protocol. A value of -1 means all protocols.
	Protocol string `pulumi:"protocol"`
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction string `pulumi:"ruleAction"`
	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber int `pulumi:"ruleNumber"`
	// The to port to match.
	ToPort *int `pulumi:"toPort"`
}

// The set of arguments for constructing a NetworkAclRule resource.
type NetworkAclRuleArgs struct {
	// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
	CidrBlock pulumi.StringPtrInput
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
	Egress pulumi.BoolPtrInput
	// The from port to match.
	FromPort pulumi.IntPtrInput
	// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
	IcmpCode pulumi.StringPtrInput
	// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
	IcmpType pulumi.StringPtrInput
	// The IPv6 CIDR block to allow or deny.
	Ipv6CidrBlock pulumi.StringPtrInput
	// The ID of the network ACL.
	NetworkAclId pulumi.StringInput
	// The protocol. A value of -1 means all protocols.
	Protocol pulumi.StringInput
	// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
	RuleAction pulumi.StringInput
	// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
	RuleNumber pulumi.IntInput
	// The to port to match.
	ToPort pulumi.IntPtrInput
}

func (NetworkAclRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclRuleArgs)(nil)).Elem()
}
