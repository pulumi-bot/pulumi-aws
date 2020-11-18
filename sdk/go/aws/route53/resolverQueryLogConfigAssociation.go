// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Route 53 Resolver query logging configuration association resource.
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
// 		_, err := route53.NewResolverQueryLogConfigAssociation(ctx, "example", &route53.ResolverQueryLogConfigAssociationArgs{
// 			ResolverQueryLogConfigId: pulumi.Any(aws_route53_resolver_query_log_config.Example.Id),
// 			ResourceId:               pulumi.Any(aws_vpc.Example.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ResolverQueryLogConfigAssociation struct {
	pulumi.CustomResourceState

	// The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
	ResolverQueryLogConfigId pulumi.StringOutput `pulumi:"resolverQueryLogConfigId"`
	// The ID of a VPC that you want this query logging configuration to log queries for.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
}

// NewResolverQueryLogConfigAssociation registers a new resource with the given unique name, arguments, and options.
func NewResolverQueryLogConfigAssociation(ctx *pulumi.Context,
	name string, args *ResolverQueryLogConfigAssociationArgs, opts ...pulumi.ResourceOption) (*ResolverQueryLogConfigAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResolverQueryLogConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ResolverQueryLogConfigId'")
	}
	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	var resource ResolverQueryLogConfigAssociation
	err := ctx.RegisterResource("aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverQueryLogConfigAssociation gets an existing ResolverQueryLogConfigAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverQueryLogConfigAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverQueryLogConfigAssociationState, opts ...pulumi.ResourceOption) (*ResolverQueryLogConfigAssociation, error) {
	var resource ResolverQueryLogConfigAssociation
	err := ctx.ReadResource("aws:route53/resolverQueryLogConfigAssociation:ResolverQueryLogConfigAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverQueryLogConfigAssociation resources.
type resolverQueryLogConfigAssociationState struct {
	// The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
	ResolverQueryLogConfigId *string `pulumi:"resolverQueryLogConfigId"`
	// The ID of a VPC that you want this query logging configuration to log queries for.
	ResourceId *string `pulumi:"resourceId"`
}

type ResolverQueryLogConfigAssociationState struct {
	// The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
	ResolverQueryLogConfigId pulumi.StringPtrInput
	// The ID of a VPC that you want this query logging configuration to log queries for.
	ResourceId pulumi.StringPtrInput
}

func (ResolverQueryLogConfigAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverQueryLogConfigAssociationState)(nil)).Elem()
}

type resolverQueryLogConfigAssociationArgs struct {
	// The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
	ResolverQueryLogConfigId string `pulumi:"resolverQueryLogConfigId"`
	// The ID of a VPC that you want this query logging configuration to log queries for.
	ResourceId string `pulumi:"resourceId"`
}

// The set of arguments for constructing a ResolverQueryLogConfigAssociation resource.
type ResolverQueryLogConfigAssociationArgs struct {
	// The ID of the Route 53 Resolver query logging configuration that you want to associate a VPC with.
	ResolverQueryLogConfigId pulumi.StringInput
	// The ID of a VPC that you want this query logging configuration to log queries for.
	ResourceId pulumi.StringInput
}

func (ResolverQueryLogConfigAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverQueryLogConfigAssociationArgs)(nil)).Elem()
}

type ResolverQueryLogConfigAssociationInput interface {
	pulumi.Input

	ToResolverQueryLogConfigAssociationOutput() ResolverQueryLogConfigAssociationOutput
	ToResolverQueryLogConfigAssociationOutputWithContext(ctx context.Context) ResolverQueryLogConfigAssociationOutput
}

func (ResolverQueryLogConfigAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverQueryLogConfigAssociation)(nil)).Elem()
}

func (i ResolverQueryLogConfigAssociation) ToResolverQueryLogConfigAssociationOutput() ResolverQueryLogConfigAssociationOutput {
	return i.ToResolverQueryLogConfigAssociationOutputWithContext(context.Background())
}

func (i ResolverQueryLogConfigAssociation) ToResolverQueryLogConfigAssociationOutputWithContext(ctx context.Context) ResolverQueryLogConfigAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverQueryLogConfigAssociationOutput)
}

type ResolverQueryLogConfigAssociationOutput struct {
	*pulumi.OutputState
}

func (ResolverQueryLogConfigAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverQueryLogConfigAssociationOutput)(nil)).Elem()
}

func (o ResolverQueryLogConfigAssociationOutput) ToResolverQueryLogConfigAssociationOutput() ResolverQueryLogConfigAssociationOutput {
	return o
}

func (o ResolverQueryLogConfigAssociationOutput) ToResolverQueryLogConfigAssociationOutputWithContext(ctx context.Context) ResolverQueryLogConfigAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResolverQueryLogConfigAssociationOutput{})
}
