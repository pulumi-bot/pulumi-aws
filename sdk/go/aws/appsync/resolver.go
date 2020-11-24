// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AppSync Resolver.
//
// ## Import
//
// `aws_appsync_resolver` can be imported with their `api_id`, a hyphen, `type`, a hypen and `field` e.g.
//
// ```sh
//  $ pulumi import aws:appsync/resolver:Resolver example abcdef123456-exampleType-exampleField
// ```
type Resolver struct {
	pulumi.CustomResourceState

	// The API ID for the GraphQL API.
	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// The ARN
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The CachingConfig.
	CachingConfig ResolverCachingConfigPtrOutput `pulumi:"cachingConfig"`
	// The DataSource name.
	DataSource pulumi.StringPtrOutput `pulumi:"dataSource"`
	// The field name from the schema defined in the GraphQL API.
	Field pulumi.StringOutput `pulumi:"field"`
	// The resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The PipelineConfig.
	PipelineConfig ResolverPipelineConfigPtrOutput `pulumi:"pipelineConfig"`
	// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
	RequestTemplate pulumi.StringOutput `pulumi:"requestTemplate"`
	// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
	ResponseTemplate pulumi.StringOutput `pulumi:"responseTemplate"`
	// The type name from the schema defined in the GraphQL API.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewResolver registers a new resource with the given unique name, arguments, and options.
func NewResolver(ctx *pulumi.Context,
	name string, args *ResolverArgs, opts ...pulumi.ResourceOption) (*Resolver, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.Field == nil {
		return nil, errors.New("invalid value for required argument 'Field'")
	}
	if args.RequestTemplate == nil {
		return nil, errors.New("invalid value for required argument 'RequestTemplate'")
	}
	if args.ResponseTemplate == nil {
		return nil, errors.New("invalid value for required argument 'ResponseTemplate'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Resolver
	err := ctx.RegisterResource("aws:appsync/resolver:Resolver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolver gets an existing Resolver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverState, opts ...pulumi.ResourceOption) (*Resolver, error) {
	var resource Resolver
	err := ctx.ReadResource("aws:appsync/resolver:Resolver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resolver resources.
type resolverState struct {
	// The API ID for the GraphQL API.
	ApiId *string `pulumi:"apiId"`
	// The ARN
	Arn *string `pulumi:"arn"`
	// The CachingConfig.
	CachingConfig *ResolverCachingConfig `pulumi:"cachingConfig"`
	// The DataSource name.
	DataSource *string `pulumi:"dataSource"`
	// The field name from the schema defined in the GraphQL API.
	Field *string `pulumi:"field"`
	// The resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind *string `pulumi:"kind"`
	// The PipelineConfig.
	PipelineConfig *ResolverPipelineConfig `pulumi:"pipelineConfig"`
	// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
	RequestTemplate *string `pulumi:"requestTemplate"`
	// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
	ResponseTemplate *string `pulumi:"responseTemplate"`
	// The type name from the schema defined in the GraphQL API.
	Type *string `pulumi:"type"`
}

type ResolverState struct {
	// The API ID for the GraphQL API.
	ApiId pulumi.StringPtrInput
	// The ARN
	Arn pulumi.StringPtrInput
	// The CachingConfig.
	CachingConfig ResolverCachingConfigPtrInput
	// The DataSource name.
	DataSource pulumi.StringPtrInput
	// The field name from the schema defined in the GraphQL API.
	Field pulumi.StringPtrInput
	// The resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind pulumi.StringPtrInput
	// The PipelineConfig.
	PipelineConfig ResolverPipelineConfigPtrInput
	// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
	RequestTemplate pulumi.StringPtrInput
	// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
	ResponseTemplate pulumi.StringPtrInput
	// The type name from the schema defined in the GraphQL API.
	Type pulumi.StringPtrInput
}

func (ResolverState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverState)(nil)).Elem()
}

type resolverArgs struct {
	// The API ID for the GraphQL API.
	ApiId string `pulumi:"apiId"`
	// The CachingConfig.
	CachingConfig *ResolverCachingConfig `pulumi:"cachingConfig"`
	// The DataSource name.
	DataSource *string `pulumi:"dataSource"`
	// The field name from the schema defined in the GraphQL API.
	Field string `pulumi:"field"`
	// The resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind *string `pulumi:"kind"`
	// The PipelineConfig.
	PipelineConfig *ResolverPipelineConfig `pulumi:"pipelineConfig"`
	// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
	RequestTemplate string `pulumi:"requestTemplate"`
	// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
	ResponseTemplate string `pulumi:"responseTemplate"`
	// The type name from the schema defined in the GraphQL API.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Resolver resource.
type ResolverArgs struct {
	// The API ID for the GraphQL API.
	ApiId pulumi.StringInput
	// The CachingConfig.
	CachingConfig ResolverCachingConfigPtrInput
	// The DataSource name.
	DataSource pulumi.StringPtrInput
	// The field name from the schema defined in the GraphQL API.
	Field pulumi.StringInput
	// The resolver type. Valid values are `UNIT` and `PIPELINE`.
	Kind pulumi.StringPtrInput
	// The PipelineConfig.
	PipelineConfig ResolverPipelineConfigPtrInput
	// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
	RequestTemplate pulumi.StringInput
	// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
	ResponseTemplate pulumi.StringInput
	// The type name from the schema defined in the GraphQL API.
	Type pulumi.StringInput
}

func (ResolverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverArgs)(nil)).Elem()
}

type ResolverInput interface {
	pulumi.Input

	ToResolverOutput() ResolverOutput
	ToResolverOutputWithContext(ctx context.Context) ResolverOutput
}

func (Resolver) ElementType() reflect.Type {
	return reflect.TypeOf((*Resolver)(nil)).Elem()
}

func (i Resolver) ToResolverOutput() ResolverOutput {
	return i.ToResolverOutputWithContext(context.Background())
}

func (i Resolver) ToResolverOutputWithContext(ctx context.Context) ResolverOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverOutput)
}

type ResolverOutput struct {
	*pulumi.OutputState
}

func (ResolverOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverOutput)(nil)).Elem()
}

func (o ResolverOutput) ToResolverOutput() ResolverOutput {
	return o
}

func (o ResolverOutput) ToResolverOutputWithContext(ctx context.Context) ResolverOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResolverOutput{})
}
