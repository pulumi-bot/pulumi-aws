// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datapipeline

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Data Pipeline resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/datapipeline"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := datapipeline.NewPipeline(ctx, "_default", nil)
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
// `aws_datapipeline_pipeline` can be imported by using the id (Pipeline ID), e.g.
//
// ```sh
//  $ pulumi import aws:datapipeline/pipeline:Pipeline default df-1234567890
// ```
type Pipeline struct {
	pulumi.CustomResourceState

	// The description of Pipeline.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of Pipeline.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		args = &PipelineArgs{}
	}

	var resource Pipeline
	err := ctx.RegisterResource("aws:datapipeline/pipeline:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("aws:datapipeline/pipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	// The description of Pipeline.
	Description *string `pulumi:"description"`
	// The name of Pipeline.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type PipelineState struct {
	// The description of Pipeline.
	Description pulumi.StringPtrInput
	// The name of Pipeline.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// The description of Pipeline.
	Description *string `pulumi:"description"`
	// The name of Pipeline.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// The description of Pipeline.
	Description pulumi.StringPtrInput
	// The name of Pipeline.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

func (i *Pipeline) ToPipelinePtrOutput() PipelinePtrOutput {
	return i.ToPipelinePtrOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelinePtrOutput)
}

type PipelinePtrInput interface {
	pulumi.Input

	ToPipelinePtrOutput() PipelinePtrOutput
	ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput
}

type pipelinePtrType PipelineArgs

func (*pipelinePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil))
}

func (i *pipelinePtrType) ToPipelinePtrOutput() PipelinePtrOutput {
	return i.ToPipelinePtrOutputWithContext(context.Background())
}

func (i *pipelinePtrType) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput).ToPipelinePtrOutput()
}

// PipelineArrayInput is an input type that accepts PipelineArray and PipelineArrayOutput values.
// You can construct a concrete instance of `PipelineArrayInput` via:
//
//          PipelineArray{ PipelineArgs{...} }
type PipelineArrayInput interface {
	pulumi.Input

	ToPipelineArrayOutput() PipelineArrayOutput
	ToPipelineArrayOutputWithContext(context.Context) PipelineArrayOutput
}

type PipelineArray []PipelineInput

func (PipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pipeline)(nil))
}

func (i PipelineArray) ToPipelineArrayOutput() PipelineArrayOutput {
	return i.ToPipelineArrayOutputWithContext(context.Background())
}

func (i PipelineArray) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineArrayOutput)
}

// PipelineMapInput is an input type that accepts PipelineMap and PipelineMapOutput values.
// You can construct a concrete instance of `PipelineMapInput` via:
//
//          PipelineMap{ "key": PipelineArgs{...} }
type PipelineMapInput interface {
	pulumi.Input

	ToPipelineMapOutput() PipelineMapOutput
	ToPipelineMapOutputWithContext(context.Context) PipelineMapOutput
}

type PipelineMap map[string]PipelineInput

func (PipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pipeline)(nil))
}

func (i PipelineMap) ToPipelineMapOutput() PipelineMapOutput {
	return i.ToPipelineMapOutputWithContext(context.Background())
}

func (i PipelineMap) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineMapOutput)
}

type PipelineOutput struct {
	*pulumi.OutputState
}

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Pipeline)(nil))
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelinePtrOutput() PipelinePtrOutput {
	return o.ToPipelinePtrOutputWithContext(context.Background())
}

func (o PipelineOutput) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return o.ApplyT(func(v Pipeline) *Pipeline {
		return &v
	}).(PipelinePtrOutput)
}

type PipelinePtrOutput struct {
	*pulumi.OutputState
}

func (PipelinePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil))
}

func (o PipelinePtrOutput) ToPipelinePtrOutput() PipelinePtrOutput {
	return o
}

func (o PipelinePtrOutput) ToPipelinePtrOutputWithContext(ctx context.Context) PipelinePtrOutput {
	return o
}

type PipelineArrayOutput struct{ *pulumi.OutputState }

func (PipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Pipeline)(nil))
}

func (o PipelineArrayOutput) ToPipelineArrayOutput() PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) Index(i pulumi.IntInput) PipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Pipeline {
		return vs[0].([]Pipeline)[vs[1].(int)]
	}).(PipelineOutput)
}

type PipelineMapOutput struct{ *pulumi.OutputState }

func (PipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Pipeline)(nil))
}

func (o PipelineMapOutput) ToPipelineMapOutput() PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) MapIndex(k pulumi.StringInput) PipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Pipeline {
		return vs[0].(map[string]Pipeline)[vs[1].(string)]
	}).(PipelineOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineOutput{})
	pulumi.RegisterOutputType(PipelinePtrOutput{})
	pulumi.RegisterOutputType(PipelineArrayOutput{})
	pulumi.RegisterOutputType(PipelineMapOutput{})
}
