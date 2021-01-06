// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Connects a custom domain name registered via `apigateway.DomainName`
// with a deployed API so that its methods can be called via the
// custom domain name.
//
// ## Import
//
// `aws_api_gateway_base_path_mapping` can be imported by using the domain name and base path, e.g. For empty `base_path` (e.g. root path (`/`))
//
// ```sh
//  $ pulumi import aws:apigateway/basePathMapping:BasePathMapping example example.com/
// ```
//
//  Otherwise
//
// ```sh
//  $ pulumi import aws:apigateway/basePathMapping:BasePathMapping example example.com/base-path
// ```
type BasePathMapping struct {
	pulumi.CustomResourceState

	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	BasePath pulumi.StringPtrOutput `pulumi:"basePath"`
	// The already-registered domain name to connect the API to.
	DomainName pulumi.StringOutput `pulumi:"domainName"`
	// The id of the API to connect.
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	StageName pulumi.StringPtrOutput `pulumi:"stageName"`
}

// NewBasePathMapping registers a new resource with the given unique name, arguments, and options.
func NewBasePathMapping(ctx *pulumi.Context,
	name string, args *BasePathMappingArgs, opts ...pulumi.ResourceOption) (*BasePathMapping, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	var resource BasePathMapping
	err := ctx.RegisterResource("aws:apigateway/basePathMapping:BasePathMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBasePathMapping gets an existing BasePathMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBasePathMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BasePathMappingState, opts ...pulumi.ResourceOption) (*BasePathMapping, error) {
	var resource BasePathMapping
	err := ctx.ReadResource("aws:apigateway/basePathMapping:BasePathMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BasePathMapping resources.
type basePathMappingState struct {
	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	BasePath *string `pulumi:"basePath"`
	// The already-registered domain name to connect the API to.
	DomainName *string `pulumi:"domainName"`
	// The id of the API to connect.
	RestApi *string `pulumi:"restApi"`
	// The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	StageName *string `pulumi:"stageName"`
}

type BasePathMappingState struct {
	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	BasePath pulumi.StringPtrInput
	// The already-registered domain name to connect the API to.
	DomainName pulumi.StringPtrInput
	// The id of the API to connect.
	RestApi pulumi.StringPtrInput
	// The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	StageName pulumi.StringPtrInput
}

func (BasePathMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*basePathMappingState)(nil)).Elem()
}

type basePathMappingArgs struct {
	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	BasePath *string `pulumi:"basePath"`
	// The already-registered domain name to connect the API to.
	DomainName string `pulumi:"domainName"`
	// The id of the API to connect.
	RestApi interface{} `pulumi:"restApi"`
	// The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	StageName *string `pulumi:"stageName"`
}

// The set of arguments for constructing a BasePathMapping resource.
type BasePathMappingArgs struct {
	// Path segment that must be prepended to the path when accessing the API via this mapping. If omitted, the API is exposed at the root of the given domain.
	BasePath pulumi.StringPtrInput
	// The already-registered domain name to connect the API to.
	DomainName pulumi.StringInput
	// The id of the API to connect.
	RestApi pulumi.Input
	// The name of a specific deployment stage to expose at the given path. If omitted, callers may select any stage by including its name as a path element after the base path.
	StageName pulumi.StringPtrInput
}

func (BasePathMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*basePathMappingArgs)(nil)).Elem()
}

type BasePathMappingInput interface {
	pulumi.Input

	ToBasePathMappingOutput() BasePathMappingOutput
	ToBasePathMappingOutputWithContext(ctx context.Context) BasePathMappingOutput
}

func (*BasePathMapping) ElementType() reflect.Type {
	return reflect.TypeOf((*BasePathMapping)(nil))
}

func (i *BasePathMapping) ToBasePathMappingOutput() BasePathMappingOutput {
	return i.ToBasePathMappingOutputWithContext(context.Background())
}

func (i *BasePathMapping) ToBasePathMappingOutputWithContext(ctx context.Context) BasePathMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasePathMappingOutput)
}

func (i *BasePathMapping) ToBasePathMappingPtrOutput() BasePathMappingPtrOutput {
	return i.ToBasePathMappingPtrOutputWithContext(context.Background())
}

func (i *BasePathMapping) ToBasePathMappingPtrOutputWithContext(ctx context.Context) BasePathMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasePathMappingPtrOutput)
}

type BasePathMappingPtrInput interface {
	pulumi.Input

	ToBasePathMappingPtrOutput() BasePathMappingPtrOutput
	ToBasePathMappingPtrOutputWithContext(ctx context.Context) BasePathMappingPtrOutput
}

type basePathMappingPtrType BasePathMappingArgs

func (*basePathMappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BasePathMapping)(nil))
}

func (i *basePathMappingPtrType) ToBasePathMappingPtrOutput() BasePathMappingPtrOutput {
	return i.ToBasePathMappingPtrOutputWithContext(context.Background())
}

func (i *basePathMappingPtrType) ToBasePathMappingPtrOutputWithContext(ctx context.Context) BasePathMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasePathMappingOutput).ToBasePathMappingPtrOutput()
}

// BasePathMappingArrayInput is an input type that accepts BasePathMappingArray and BasePathMappingArrayOutput values.
// You can construct a concrete instance of `BasePathMappingArrayInput` via:
//
//          BasePathMappingArray{ BasePathMappingArgs{...} }
type BasePathMappingArrayInput interface {
	pulumi.Input

	ToBasePathMappingArrayOutput() BasePathMappingArrayOutput
	ToBasePathMappingArrayOutputWithContext(context.Context) BasePathMappingArrayOutput
}

type BasePathMappingArray []BasePathMappingInput

func (BasePathMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BasePathMapping)(nil))
}

func (i BasePathMappingArray) ToBasePathMappingArrayOutput() BasePathMappingArrayOutput {
	return i.ToBasePathMappingArrayOutputWithContext(context.Background())
}

func (i BasePathMappingArray) ToBasePathMappingArrayOutputWithContext(ctx context.Context) BasePathMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasePathMappingArrayOutput)
}

// BasePathMappingMapInput is an input type that accepts BasePathMappingMap and BasePathMappingMapOutput values.
// You can construct a concrete instance of `BasePathMappingMapInput` via:
//
//          BasePathMappingMap{ "key": BasePathMappingArgs{...} }
type BasePathMappingMapInput interface {
	pulumi.Input

	ToBasePathMappingMapOutput() BasePathMappingMapOutput
	ToBasePathMappingMapOutputWithContext(context.Context) BasePathMappingMapOutput
}

type BasePathMappingMap map[string]BasePathMappingInput

func (BasePathMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BasePathMapping)(nil))
}

func (i BasePathMappingMap) ToBasePathMappingMapOutput() BasePathMappingMapOutput {
	return i.ToBasePathMappingMapOutputWithContext(context.Background())
}

func (i BasePathMappingMap) ToBasePathMappingMapOutputWithContext(ctx context.Context) BasePathMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BasePathMappingMapOutput)
}

type BasePathMappingOutput struct {
	*pulumi.OutputState
}

func (BasePathMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BasePathMapping)(nil))
}

func (o BasePathMappingOutput) ToBasePathMappingOutput() BasePathMappingOutput {
	return o
}

func (o BasePathMappingOutput) ToBasePathMappingOutputWithContext(ctx context.Context) BasePathMappingOutput {
	return o
}

func (o BasePathMappingOutput) ToBasePathMappingPtrOutput() BasePathMappingPtrOutput {
	return o.ToBasePathMappingPtrOutputWithContext(context.Background())
}

func (o BasePathMappingOutput) ToBasePathMappingPtrOutputWithContext(ctx context.Context) BasePathMappingPtrOutput {
	return o.ApplyT(func(v BasePathMapping) *BasePathMapping {
		return &v
	}).(BasePathMappingPtrOutput)
}

type BasePathMappingPtrOutput struct {
	*pulumi.OutputState
}

func (BasePathMappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BasePathMapping)(nil))
}

func (o BasePathMappingPtrOutput) ToBasePathMappingPtrOutput() BasePathMappingPtrOutput {
	return o
}

func (o BasePathMappingPtrOutput) ToBasePathMappingPtrOutputWithContext(ctx context.Context) BasePathMappingPtrOutput {
	return o
}

type BasePathMappingArrayOutput struct{ *pulumi.OutputState }

func (BasePathMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BasePathMapping)(nil))
}

func (o BasePathMappingArrayOutput) ToBasePathMappingArrayOutput() BasePathMappingArrayOutput {
	return o
}

func (o BasePathMappingArrayOutput) ToBasePathMappingArrayOutputWithContext(ctx context.Context) BasePathMappingArrayOutput {
	return o
}

func (o BasePathMappingArrayOutput) Index(i pulumi.IntInput) BasePathMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BasePathMapping {
		return vs[0].([]BasePathMapping)[vs[1].(int)]
	}).(BasePathMappingOutput)
}

type BasePathMappingMapOutput struct{ *pulumi.OutputState }

func (BasePathMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BasePathMapping)(nil))
}

func (o BasePathMappingMapOutput) ToBasePathMappingMapOutput() BasePathMappingMapOutput {
	return o
}

func (o BasePathMappingMapOutput) ToBasePathMappingMapOutputWithContext(ctx context.Context) BasePathMappingMapOutput {
	return o
}

func (o BasePathMappingMapOutput) MapIndex(k pulumi.StringInput) BasePathMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BasePathMapping {
		return vs[0].(map[string]BasePathMapping)[vs[1].(string)]
	}).(BasePathMappingOutput)
}

func init() {
	pulumi.RegisterOutputType(BasePathMappingOutput{})
	pulumi.RegisterOutputType(BasePathMappingPtrOutput{})
	pulumi.RegisterOutputType(BasePathMappingArrayOutput{})
	pulumi.RegisterOutputType(BasePathMappingMapOutput{})
}
