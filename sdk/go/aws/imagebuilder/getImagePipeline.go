// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about an Image Builder Image Pipeline.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/imagebuilder"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := imagebuilder.LookupImagePipeline(ctx, &imagebuilder.LookupImagePipelineArgs{
// 			Arn: "arn:aws:imagebuilder:us-west-2:aws:image-pipeline/example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupImagePipeline(ctx *pulumi.Context, args *LookupImagePipelineArgs, opts ...pulumi.InvokeOption) (*LookupImagePipelineResult, error) {
	var rv LookupImagePipelineResult
	err := ctx.Invoke("aws:imagebuilder/getImagePipeline:getImagePipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImagePipeline.
type LookupImagePipelineArgs struct {
	// Amazon Resource Name (ARN) of the image pipeline.
	Arn string `pulumi:"arn"`
	// Key-value map of resource tags for the image pipeline.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getImagePipeline.
type LookupImagePipelineResult struct {
	Arn string `pulumi:"arn"`
	// Date the image pipeline was created.
	DateCreated string `pulumi:"dateCreated"`
	// Date the image pipeline was last run.
	DateLastRun string `pulumi:"dateLastRun"`
	// Date the image pipeline will run next.
	DateNextRun string `pulumi:"dateNextRun"`
	// Date the image pipeline was updated.
	DateUpdated string `pulumi:"dateUpdated"`
	// Description of the image pipeline.
	Description string `pulumi:"description"`
	// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
	DistributionConfigurationArn string `pulumi:"distributionConfigurationArn"`
	// Whether additional information about the image being created is collected.
	EnhancedImageMetadataEnabled bool `pulumi:"enhancedImageMetadataEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
	ImageRecipeArn string `pulumi:"imageRecipeArn"`
	// List of an object with image tests configuration.
	ImageTestsConfigurations []GetImagePipelineImageTestsConfiguration `pulumi:"imageTestsConfigurations"`
	// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
	InfrastructureConfigurationArn string `pulumi:"infrastructureConfigurationArn"`
	// Name of the image pipeline.
	Name string `pulumi:"name"`
	// Platform of the image pipeline.
	Platform string `pulumi:"platform"`
	// List of an object with schedule settings.
	Schedules []GetImagePipelineSchedule `pulumi:"schedules"`
	// Status of the image pipeline.
	Status string `pulumi:"status"`
	// Key-value map of resource tags for the image pipeline.
	Tags map[string]string `pulumi:"tags"`
}

func LookupImagePipelineApply(ctx *pulumi.Context, args LookupImagePipelineApplyInput, opts ...pulumi.InvokeOption) LookupImagePipelineResultOutput {
	return args.ToLookupImagePipelineApplyOutput().ApplyT(func(v LookupImagePipelineArgs) (LookupImagePipelineResult, error) {
		r, err := LookupImagePipeline(ctx, &v, opts...)
		return *r, err

	}).(LookupImagePipelineResultOutput)
}

// LookupImagePipelineApplyInput is an input type that accepts LookupImagePipelineApplyArgs and LookupImagePipelineApplyOutput values.
// You can construct a concrete instance of `LookupImagePipelineApplyInput` via:
//
//          LookupImagePipelineApplyArgs{...}
type LookupImagePipelineApplyInput interface {
	pulumi.Input

	ToLookupImagePipelineApplyOutput() LookupImagePipelineApplyOutput
	ToLookupImagePipelineApplyOutputWithContext(context.Context) LookupImagePipelineApplyOutput
}

// A collection of arguments for invoking getImagePipeline.
type LookupImagePipelineApplyArgs struct {
	// Amazon Resource Name (ARN) of the image pipeline.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Key-value map of resource tags for the image pipeline.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupImagePipelineApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImagePipelineArgs)(nil)).Elem()
}

func (i LookupImagePipelineApplyArgs) ToLookupImagePipelineApplyOutput() LookupImagePipelineApplyOutput {
	return i.ToLookupImagePipelineApplyOutputWithContext(context.Background())
}

func (i LookupImagePipelineApplyArgs) ToLookupImagePipelineApplyOutputWithContext(ctx context.Context) LookupImagePipelineApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupImagePipelineApplyOutput)
}

// A collection of arguments for invoking getImagePipeline.
type LookupImagePipelineApplyOutput struct{ *pulumi.OutputState }

func (LookupImagePipelineApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImagePipelineArgs)(nil)).Elem()
}

func (o LookupImagePipelineApplyOutput) ToLookupImagePipelineApplyOutput() LookupImagePipelineApplyOutput {
	return o
}

func (o LookupImagePipelineApplyOutput) ToLookupImagePipelineApplyOutputWithContext(ctx context.Context) LookupImagePipelineApplyOutput {
	return o
}

// Amazon Resource Name (ARN) of the image pipeline.
func (o LookupImagePipelineApplyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineArgs) string { return v.Arn }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the image pipeline.
func (o LookupImagePipelineApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImagePipelineArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getImagePipeline.
type LookupImagePipelineResultOutput struct{ *pulumi.OutputState }

func (LookupImagePipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImagePipelineResult)(nil)).Elem()
}

func (o LookupImagePipelineResultOutput) ToLookupImagePipelineResultOutput() LookupImagePipelineResultOutput {
	return o
}

func (o LookupImagePipelineResultOutput) ToLookupImagePipelineResultOutputWithContext(ctx context.Context) LookupImagePipelineResultOutput {
	return o
}

func (o LookupImagePipelineResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Date the image pipeline was created.
func (o LookupImagePipelineResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// Date the image pipeline was last run.
func (o LookupImagePipelineResultOutput) DateLastRun() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.DateLastRun }).(pulumi.StringOutput)
}

// Date the image pipeline will run next.
func (o LookupImagePipelineResultOutput) DateNextRun() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.DateNextRun }).(pulumi.StringOutput)
}

// Date the image pipeline was updated.
func (o LookupImagePipelineResultOutput) DateUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.DateUpdated }).(pulumi.StringOutput)
}

// Description of the image pipeline.
func (o LookupImagePipelineResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.Description }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the Image Builder Distribution Configuration.
func (o LookupImagePipelineResultOutput) DistributionConfigurationArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.DistributionConfigurationArn }).(pulumi.StringOutput)
}

// Whether additional information about the image being created is collected.
func (o LookupImagePipelineResultOutput) EnhancedImageMetadataEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) bool { return v.EnhancedImageMetadataEnabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupImagePipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

// Amazon Resource Name (ARN) of the Image Builder Infrastructure Recipe.
func (o LookupImagePipelineResultOutput) ImageRecipeArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.ImageRecipeArn }).(pulumi.StringOutput)
}

// List of an object with image tests configuration.
func (o LookupImagePipelineResultOutput) ImageTestsConfigurations() GetImagePipelineImageTestsConfigurationArrayOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) []GetImagePipelineImageTestsConfiguration {
		return v.ImageTestsConfigurations
	}).(GetImagePipelineImageTestsConfigurationArrayOutput)
}

// Amazon Resource Name (ARN) of the Image Builder Infrastructure Configuration.
func (o LookupImagePipelineResultOutput) InfrastructureConfigurationArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.InfrastructureConfigurationArn }).(pulumi.StringOutput)
}

// Name of the image pipeline.
func (o LookupImagePipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

// Platform of the image pipeline.
func (o LookupImagePipelineResultOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.Platform }).(pulumi.StringOutput)
}

// List of an object with schedule settings.
func (o LookupImagePipelineResultOutput) Schedules() GetImagePipelineScheduleArrayOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) []GetImagePipelineSchedule { return v.Schedules }).(GetImagePipelineScheduleArrayOutput)
}

// Status of the image pipeline.
func (o LookupImagePipelineResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) string { return v.Status }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the image pipeline.
func (o LookupImagePipelineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImagePipelineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImagePipelineApplyOutput{})
	pulumi.RegisterOutputType(LookupImagePipelineResultOutput{})
}
