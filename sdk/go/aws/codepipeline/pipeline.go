// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package codepipeline

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodePipeline.
// 
// > **NOTE on `codepipeline.Pipeline`:** - the `GITHUB_TOKEN` environment variable must be set if the GitHub provider is specified.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codepipeline.markdown.
type Pipeline struct {
	pulumi.CustomResourceState

	// The codepipeline ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// An artifactStore block. Artifact stores are documented below.
	// * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
	ArtifactStore PipelineArtifactStoreOutput `pulumi:"artifactStore"`
	// The name of the pipeline.
	Name pulumi.StringOutput `pulumi:"name"`
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	Stages PipelineStageArrayOutput `pulumi:"stages"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil || args.ArtifactStore == nil {
		return nil, errors.New("missing required argument 'ArtifactStore'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil || args.Stages == nil {
		return nil, errors.New("missing required argument 'Stages'")
	}
	if args == nil {
		args = &PipelineArgs{}
	}
	var resource Pipeline
	err := ctx.RegisterResource("aws:codepipeline/pipeline:Pipeline", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:codepipeline/pipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
	// The codepipeline ARN.
	Arn *string `pulumi:"arn"`
	// An artifactStore block. Artifact stores are documented below.
	// * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
	ArtifactStore *PipelineArtifactStore `pulumi:"artifactStore"`
	// The name of the pipeline.
	Name *string `pulumi:"name"`
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn *string `pulumi:"roleArn"`
	Stages []PipelineStage `pulumi:"stages"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

type PipelineState struct {
	// The codepipeline ARN.
	Arn pulumi.StringPtrInput
	// An artifactStore block. Artifact stores are documented below.
	// * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
	ArtifactStore PipelineArtifactStorePtrInput
	// The name of the pipeline.
	Name pulumi.StringPtrInput
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn pulumi.StringPtrInput
	Stages PipelineStageArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// An artifactStore block. Artifact stores are documented below.
	// * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
	ArtifactStore PipelineArtifactStore `pulumi:"artifactStore"`
	// The name of the pipeline.
	Name *string `pulumi:"name"`
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn string `pulumi:"roleArn"`
	Stages []PipelineStage `pulumi:"stages"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// An artifactStore block. Artifact stores are documented below.
	// * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
	ArtifactStore PipelineArtifactStoreInput
	// The name of the pipeline.
	Name pulumi.StringPtrInput
	// A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
	RoleArn pulumi.StringInput
	Stages PipelineStageArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

