// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package resolverRuleAssociation

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Route53 Resolver rule association.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_resolver_rule_association.html.markdown.
type ResolverRuleAssociation struct {
	pulumi.CustomResourceState

	// A name for the association that you're creating between a resolver rule and a VPC.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the resolver rule that you want to associate with the VPC.
	ResolverRuleId pulumi.StringOutput `pulumi:"resolverRuleId"`
	// The ID of the VPC that you want to associate the resolver rule with.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewResolverRuleAssociation registers a new resource with the given unique name, arguments, and options.
func NewResolverRuleAssociation(ctx *pulumi.Context,
	name string, args *ResolverRuleAssociationArgs, opts ...pulumi.ResourceOption) (*ResolverRuleAssociation, error) {
	if args == nil || args.ResolverRuleId == nil {
		return nil, errors.New("missing required argument 'ResolverRuleId'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &ResolverRuleAssociationArgs{}
	}
	var resource ResolverRuleAssociation
	err := ctx.RegisterResource("aws:route53/resolverRuleAssociation:ResolverRuleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverRuleAssociation gets an existing ResolverRuleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverRuleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverRuleAssociationState, opts ...pulumi.ResourceOption) (*ResolverRuleAssociation, error) {
	var resource ResolverRuleAssociation
	err := ctx.ReadResource("aws:route53/resolverRuleAssociation:ResolverRuleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverRuleAssociation resources.
type resolverRuleAssociationState struct {
	// A name for the association that you're creating between a resolver rule and a VPC.
	Name *string `pulumi:"name"`
	// The ID of the resolver rule that you want to associate with the VPC.
	ResolverRuleId *string `pulumi:"resolverRuleId"`
	// The ID of the VPC that you want to associate the resolver rule with.
	VpcId *string `pulumi:"vpcId"`
}

type ResolverRuleAssociationState struct {
	// A name for the association that you're creating between a resolver rule and a VPC.
	Name pulumi.StringPtrInput
	// The ID of the resolver rule that you want to associate with the VPC.
	ResolverRuleId pulumi.StringPtrInput
	// The ID of the VPC that you want to associate the resolver rule with.
	VpcId pulumi.StringPtrInput
}

func (ResolverRuleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleAssociationState)(nil)).Elem()
}

type resolverRuleAssociationArgs struct {
	// A name for the association that you're creating between a resolver rule and a VPC.
	Name *string `pulumi:"name"`
	// The ID of the resolver rule that you want to associate with the VPC.
	ResolverRuleId string `pulumi:"resolverRuleId"`
	// The ID of the VPC that you want to associate the resolver rule with.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a ResolverRuleAssociation resource.
type ResolverRuleAssociationArgs struct {
	// A name for the association that you're creating between a resolver rule and a VPC.
	Name pulumi.StringPtrInput
	// The ID of the resolver rule that you want to associate with the VPC.
	ResolverRuleId pulumi.StringInput
	// The ID of the VPC that you want to associate the resolver rule with.
	VpcId pulumi.StringInput
}

func (ResolverRuleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleAssociationArgs)(nil)).Elem()
}

