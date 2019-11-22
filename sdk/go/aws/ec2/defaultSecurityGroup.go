// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the default AWS Security Group.
// 
// For EC2 Classic accounts, each region comes with a Default Security Group.
// Additionally, each VPC created in AWS comes with a Default Security Group that can be managed, but not
// destroyed. **This is an advanced resource**, and has special caveats to be aware
// of when using it. Please read this document in its entirety before using this
// resource.
// 
// The `ec2.DefaultSecurityGroup` behaves differently from normal resources, in that
// this provider does not _create_ this resource, but instead "adopts" it
// into management. We can do this because these default security groups cannot be
// destroyed, and are created with a known set of default ingress/egress rules.
// 
// When this provider first adopts the Default Security Group, it **immediately removes all
// ingress and egress rules in the Security Group**. It then proceeds to create any rules specified in the
// configuration. This step is required so that only the rules specified in the
// configuration are created.
// 
// This resource treats its inline rules as absolute; only the rules defined
// inline are created, and any additions/removals external to this resource will
// result in diff shown. For these reasons, this resource is incompatible with the
// `ec2.SecurityGroupRule` resource.
// 
// For more information about Default Security Groups, see the AWS Documentation on
// [Default Security Groups][aws-default-security-groups].
// 
// ## Usage
// 
// With the exceptions mentioned above, `ec2.DefaultSecurityGroup` should
// identical behavior to `ec2.SecurityGroup`. Please consult [AWS_SECURITY_GROUP](https://www.terraform.io/docs/providers/aws/r/security_group.html)
// for further usage documentation.
// 
// ### Removing `ec2.DefaultSecurityGroup` from your configuration
// 
// Each AWS VPC (or region, if using EC2 Classic) comes with a Default Security
// Group that cannot be deleted. The `ec2.DefaultSecurityGroup` allows you to
// manage this Security Group, but this provider cannot destroy it. Removing this resource
// from your configuration will remove it from your statefile and management, but
// will not destroy the Security Group. All ingress or egress rules will be left as
// they are at the time of removal. You can resume managing them via the AWS Console.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/default_security_group.html.markdown.
type DefaultSecurityGroup struct {
	s *pulumi.ResourceState
}

// NewDefaultSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewDefaultSecurityGroup(ctx *pulumi.Context,
	name string, args *DefaultSecurityGroupArgs, opts ...pulumi.ResourceOpt) (*DefaultSecurityGroup, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["egress"] = nil
		inputs["ingress"] = nil
		inputs["revokeRulesOnDelete"] = nil
		inputs["tags"] = nil
		inputs["vpcId"] = nil
	} else {
		inputs["egress"] = args.Egress
		inputs["ingress"] = args.Ingress
		inputs["revokeRulesOnDelete"] = args.RevokeRulesOnDelete
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	inputs["arn"] = nil
	inputs["name"] = nil
	inputs["ownerId"] = nil
	s, err := ctx.RegisterResource("aws:ec2/defaultSecurityGroup:DefaultSecurityGroup", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultSecurityGroup{s: s}, nil
}

// GetDefaultSecurityGroup gets an existing DefaultSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultSecurityGroupState, opts ...pulumi.ResourceOpt) (*DefaultSecurityGroup, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["egress"] = state.Egress
		inputs["ingress"] = state.Ingress
		inputs["name"] = state.Name
		inputs["ownerId"] = state.OwnerId
		inputs["revokeRulesOnDelete"] = state.RevokeRulesOnDelete
		inputs["tags"] = state.Tags
		inputs["vpcId"] = state.VpcId
	}
	s, err := ctx.ReadResource("aws:ec2/defaultSecurityGroup:DefaultSecurityGroup", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultSecurityGroup{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DefaultSecurityGroup) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DefaultSecurityGroup) ID() pulumi.IDOutput {
	return r.s.ID()
}

func (r *DefaultSecurityGroup) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// Can be specified multiple times for each
// egress rule. Each egress block supports fields documented below.
func (r *DefaultSecurityGroup) Egress() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["egress"])
}

// Can be specified multiple times for each
// ingress rule. Each ingress block supports fields documented below.
func (r *DefaultSecurityGroup) Ingress() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["ingress"])
}

// The name of the security group
func (r *DefaultSecurityGroup) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// The owner ID.
func (r *DefaultSecurityGroup) OwnerId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["ownerId"])
}

func (r *DefaultSecurityGroup) RevokeRulesOnDelete() pulumi.BoolOutput {
	return (pulumi.BoolOutput)(r.s.State["revokeRulesOnDelete"])
}

// A mapping of tags to assign to the resource.
func (r *DefaultSecurityGroup) Tags() pulumi.MapOutput {
	return (pulumi.MapOutput)(r.s.State["tags"])
}

// The VPC ID. **Note that changing
// the `vpcId` will _not_ restore any default security group rules that were
// modified, added, or removed.** It will be left in its current state
func (r *DefaultSecurityGroup) VpcId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["vpcId"])
}

// Input properties used for looking up and filtering DefaultSecurityGroup resources.
type DefaultSecurityGroupState struct {
	Arn interface{}
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress interface{}
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress interface{}
	// The name of the security group
	Name interface{}
	// The owner ID.
	OwnerId interface{}
	RevokeRulesOnDelete interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The VPC ID. **Note that changing
	// the `vpcId` will _not_ restore any default security group rules that were
	// modified, added, or removed.** It will be left in its current state
	VpcId interface{}
}

// The set of arguments for constructing a DefaultSecurityGroup resource.
type DefaultSecurityGroupArgs struct {
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress interface{}
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress interface{}
	RevokeRulesOnDelete interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// The VPC ID. **Note that changing
	// the `vpcId` will _not_ restore any default security group rules that were
	// modified, added, or removed.** It will be left in its current state
	VpcId interface{}
}
