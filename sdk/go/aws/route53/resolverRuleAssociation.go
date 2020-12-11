// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Route53 Resolver rule association.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := route53.NewResolverRuleAssociation(ctx, "example", &route53.ResolverRuleAssociationArgs{
// 			ResolverRuleId: pulumi.Any(aws_route53_resolver_rule.Sys.Id),
// 			VpcId:          pulumi.Any(aws_vpc.Foo.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Route53 Resolver rule associations can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:route53/resolverRuleAssociation:ResolverRuleAssociation example rslvr-rrassoc-97242eaf88example
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResolverRuleId == nil {
		return nil, errors.New("invalid value for required argument 'ResolverRuleId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
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

type ResolverRuleAssociationInput interface {
	pulumi.Input

	ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput
	ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput
}

type ResolverRuleAssociationPtrInput interface {
	pulumi.Input

	ToResolverRuleAssociationPtrOutput() ResolverRuleAssociationPtrOutput
	ToResolverRuleAssociationPtrOutputWithContext(ctx context.Context) ResolverRuleAssociationPtrOutput
}

func (ResolverRuleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRuleAssociation)(nil)).Elem()
}

func (i ResolverRuleAssociation) ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput {
	return i.ToResolverRuleAssociationOutputWithContext(context.Background())
}

func (i ResolverRuleAssociation) ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleAssociationOutput)
}

func (i ResolverRuleAssociation) ToResolverRuleAssociationPtrOutput() ResolverRuleAssociationPtrOutput {
	return i.ToResolverRuleAssociationPtrOutputWithContext(context.Background())
}

func (i ResolverRuleAssociation) ToResolverRuleAssociationPtrOutputWithContext(ctx context.Context) ResolverRuleAssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleAssociationPtrOutput)
}

type ResolverRuleAssociationOutput struct {
	*pulumi.OutputState
}

func (ResolverRuleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRuleAssociationOutput)(nil)).Elem()
}

func (o ResolverRuleAssociationOutput) ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput {
	return o
}

func (o ResolverRuleAssociationOutput) ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput {
	return o
}

type ResolverRuleAssociationPtrOutput struct {
	*pulumi.OutputState
}

func (ResolverRuleAssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResolverRuleAssociation)(nil)).Elem()
}

func (o ResolverRuleAssociationPtrOutput) ToResolverRuleAssociationPtrOutput() ResolverRuleAssociationPtrOutput {
	return o
}

func (o ResolverRuleAssociationPtrOutput) ToResolverRuleAssociationPtrOutputWithContext(ctx context.Context) ResolverRuleAssociationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResolverRuleAssociationOutput{})
	pulumi.RegisterOutputType(ResolverRuleAssociationPtrOutput{})
}
