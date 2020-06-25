// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
// identical behavior to `ec2.SecurityGroup`. Please consult `AWS_SECURITY_GROUP`
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
type DefaultSecurityGroup struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the security group
	Description pulumi.StringOutput `pulumi:"description"`
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress DefaultSecurityGroupEgressArrayOutput `pulumi:"egress"`
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress DefaultSecurityGroupIngressArrayOutput `pulumi:"ingress"`
	// The name of the security group
	Name pulumi.StringOutput `pulumi:"name"`
	// The owner ID.
	OwnerId             pulumi.StringOutput  `pulumi:"ownerId"`
	RevokeRulesOnDelete pulumi.BoolPtrOutput `pulumi:"revokeRulesOnDelete"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VPC ID. **Note that changing
	// the `vpcId` will _not_ restore any default security group rules that were
	// modified, added, or removed.** It will be left in its current state
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewDefaultSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewDefaultSecurityGroup(ctx *pulumi.Context,
	name string, args *DefaultSecurityGroupArgs, opts ...pulumi.ResourceOption) (*DefaultSecurityGroup, error) {
	if args == nil {
		args = &DefaultSecurityGroupArgs{}
	}
	var resource DefaultSecurityGroup
	err := ctx.RegisterResource("aws:ec2/defaultSecurityGroup:DefaultSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultSecurityGroup gets an existing DefaultSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultSecurityGroupState, opts ...pulumi.ResourceOption) (*DefaultSecurityGroup, error) {
	var resource DefaultSecurityGroup
	err := ctx.ReadResource("aws:ec2/defaultSecurityGroup:DefaultSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultSecurityGroup resources.
type defaultSecurityGroupState struct {
	Arn *string `pulumi:"arn"`
	// The description of the security group
	Description *string `pulumi:"description"`
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress []DefaultSecurityGroupEgress `pulumi:"egress"`
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress []DefaultSecurityGroupIngress `pulumi:"ingress"`
	// The name of the security group
	Name *string `pulumi:"name"`
	// The owner ID.
	OwnerId             *string `pulumi:"ownerId"`
	RevokeRulesOnDelete *bool   `pulumi:"revokeRulesOnDelete"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID. **Note that changing
	// the `vpcId` will _not_ restore any default security group rules that were
	// modified, added, or removed.** It will be left in its current state
	VpcId *string `pulumi:"vpcId"`
}

type DefaultSecurityGroupState struct {
	Arn pulumi.StringPtrInput
	// The description of the security group
	Description pulumi.StringPtrInput
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress DefaultSecurityGroupEgressArrayInput
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress DefaultSecurityGroupIngressArrayInput
	// The name of the security group
	Name pulumi.StringPtrInput
	// The owner ID.
	OwnerId             pulumi.StringPtrInput
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID. **Note that changing
	// the `vpcId` will _not_ restore any default security group rules that were
	// modified, added, or removed.** It will be left in its current state
	VpcId pulumi.StringPtrInput
}

func (DefaultSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSecurityGroupState)(nil)).Elem()
}

type defaultSecurityGroupArgs struct {
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress []DefaultSecurityGroupEgress `pulumi:"egress"`
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress             []DefaultSecurityGroupIngress `pulumi:"ingress"`
	RevokeRulesOnDelete *bool                         `pulumi:"revokeRulesOnDelete"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID. **Note that changing
	// the `vpcId` will _not_ restore any default security group rules that were
	// modified, added, or removed.** It will be left in its current state
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a DefaultSecurityGroup resource.
type DefaultSecurityGroupArgs struct {
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress DefaultSecurityGroupEgressArrayInput
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress             DefaultSecurityGroupIngressArrayInput
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID. **Note that changing
	// the `vpcId` will _not_ restore any default security group rules that were
	// modified, added, or removed.** It will be left in its current state
	VpcId pulumi.StringPtrInput
}

func (DefaultSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultSecurityGroupArgs)(nil)).Elem()
}
