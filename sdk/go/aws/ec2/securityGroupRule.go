// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/security_group_rule.html.markdown.
type SecurityGroupRule struct {
	s *pulumi.ResourceState
}

// NewSecurityGroupRule registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroupRule(ctx *pulumi.Context,
	name string, args *SecurityGroupRuleArgs, opts ...pulumi.ResourceOpt) (*SecurityGroupRule, error) {
	if args == nil || args.FromPort == nil {
		return nil, errors.New("missing required argument 'FromPort'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil || args.SecurityGroupId == nil {
		return nil, errors.New("missing required argument 'SecurityGroupId'")
	}
	if args == nil || args.ToPort == nil {
		return nil, errors.New("missing required argument 'ToPort'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cidrBlocks"] = nil
		inputs["description"] = nil
		inputs["fromPort"] = nil
		inputs["ipv6CidrBlocks"] = nil
		inputs["prefixListIds"] = nil
		inputs["protocol"] = nil
		inputs["securityGroupId"] = nil
		inputs["self"] = nil
		inputs["sourceSecurityGroupId"] = nil
		inputs["toPort"] = nil
		inputs["type"] = nil
	} else {
		inputs["cidrBlocks"] = args.CidrBlocks
		inputs["description"] = args.Description
		inputs["fromPort"] = args.FromPort
		inputs["ipv6CidrBlocks"] = args.Ipv6CidrBlocks
		inputs["prefixListIds"] = args.PrefixListIds
		inputs["protocol"] = args.Protocol
		inputs["securityGroupId"] = args.SecurityGroupId
		inputs["self"] = args.Self
		inputs["sourceSecurityGroupId"] = args.SourceSecurityGroupId
		inputs["toPort"] = args.ToPort
		inputs["type"] = args.Type
	}
	s, err := ctx.RegisterResource("aws:ec2/securityGroupRule:SecurityGroupRule", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityGroupRule{s: s}, nil
}

// GetSecurityGroupRule gets an existing SecurityGroupRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroupRule(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecurityGroupRuleState, opts ...pulumi.ResourceOpt) (*SecurityGroupRule, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["cidrBlocks"] = state.CidrBlocks
		inputs["description"] = state.Description
		inputs["fromPort"] = state.FromPort
		inputs["ipv6CidrBlocks"] = state.Ipv6CidrBlocks
		inputs["prefixListIds"] = state.PrefixListIds
		inputs["protocol"] = state.Protocol
		inputs["securityGroupId"] = state.SecurityGroupId
		inputs["self"] = state.Self
		inputs["sourceSecurityGroupId"] = state.SourceSecurityGroupId
		inputs["toPort"] = state.ToPort
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("aws:ec2/securityGroupRule:SecurityGroupRule", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecurityGroupRule{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecurityGroupRule) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecurityGroupRule) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
func (r *SecurityGroupRule) CidrBlocks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["cidrBlocks"])
}

// Description of the rule.
func (r *SecurityGroupRule) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The start port (or ICMP type number if protocol is "icmp").
func (r *SecurityGroupRule) FromPort() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["fromPort"])
}

// List of IPv6 CIDR blocks.
func (r *SecurityGroupRule) Ipv6CidrBlocks() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["ipv6CidrBlocks"])
}

// List of prefix list IDs (for allowing access to VPC endpoints).
// Only valid with `egress`.
func (r *SecurityGroupRule) PrefixListIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["prefixListIds"])
}

// The protocol. If not icmp, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
func (r *SecurityGroupRule) Protocol() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["protocol"])
}

// The security group to apply this rule to.
func (r *SecurityGroupRule) SecurityGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["securityGroupId"])
}

// If true, the security group itself will be added as
// a source to this ingress rule.
func (r *SecurityGroupRule) Self() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["self"])
}

// The security group id to allow access to/from,
// depending on the `type`. Cannot be specified with `cidr_blocks`.
func (r *SecurityGroupRule) SourceSecurityGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sourceSecurityGroupId"])
}

// The end port (or ICMP code if protocol is "icmp").
func (r *SecurityGroupRule) ToPort() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["toPort"])
}

// The type of rule being created. Valid options are `ingress` (inbound)
// or `egress` (outbound).
func (r *SecurityGroupRule) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering SecurityGroupRule resources.
type SecurityGroupRuleState struct {
	// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
	CidrBlocks interface{}
	// Description of the rule.
	Description interface{}
	// The start port (or ICMP type number if protocol is "icmp").
	FromPort interface{}
	// List of IPv6 CIDR blocks.
	Ipv6CidrBlocks interface{}
	// List of prefix list IDs (for allowing access to VPC endpoints).
	// Only valid with `egress`.
	PrefixListIds interface{}
	// The protocol. If not icmp, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
	Protocol interface{}
	// The security group to apply this rule to.
	SecurityGroupId interface{}
	// If true, the security group itself will be added as
	// a source to this ingress rule.
	Self interface{}
	// The security group id to allow access to/from,
	// depending on the `type`. Cannot be specified with `cidr_blocks`.
	SourceSecurityGroupId interface{}
	// The end port (or ICMP code if protocol is "icmp").
	ToPort interface{}
	// The type of rule being created. Valid options are `ingress` (inbound)
	// or `egress` (outbound).
	Type interface{}
}

// The set of arguments for constructing a SecurityGroupRule resource.
type SecurityGroupRuleArgs struct {
	// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
	CidrBlocks interface{}
	// Description of the rule.
	Description interface{}
	// The start port (or ICMP type number if protocol is "icmp").
	FromPort interface{}
	// List of IPv6 CIDR blocks.
	Ipv6CidrBlocks interface{}
	// List of prefix list IDs (for allowing access to VPC endpoints).
	// Only valid with `egress`.
	PrefixListIds interface{}
	// The protocol. If not icmp, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
	Protocol interface{}
	// The security group to apply this rule to.
	SecurityGroupId interface{}
	// If true, the security group itself will be added as
	// a source to this ingress rule.
	Self interface{}
	// The security group id to allow access to/from,
	// depending on the `type`. Cannot be specified with `cidr_blocks`.
	SourceSecurityGroupId interface{}
	// The end port (or ICMP code if protocol is "icmp").
	ToPort interface{}
	// The type of rule being created. Valid options are `ingress` (inbound)
	// or `egress` (outbound).
	Type interface{}
}
