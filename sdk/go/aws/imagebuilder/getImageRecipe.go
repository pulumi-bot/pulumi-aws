// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about an Image Builder Image Recipe.
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
// 		_, err := imagebuilder.LookupImageRecipe(ctx, &imagebuilder.LookupImageRecipeArgs{
// 			Arn: "arn:aws:imagebuilder:us-east-1:aws:image-recipe/example/1.0.0",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupImageRecipe(ctx *pulumi.Context, args *LookupImageRecipeArgs, opts ...pulumi.InvokeOption) (*LookupImageRecipeResult, error) {
	var rv LookupImageRecipeResult
	err := ctx.Invoke("aws:imagebuilder/getImageRecipe:getImageRecipe", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageRecipe.
type LookupImageRecipeArgs struct {
	// Amazon Resource Name (ARN) of the image recipe.
	Arn string `pulumi:"arn"`
	// Key-value map of resource tags for the image recipe.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getImageRecipe.
type LookupImageRecipeResult struct {
	Arn string `pulumi:"arn"`
	// Set of objects with block device mappings for the the image recipe.
	BlockDeviceMappings []GetImageRecipeBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// List of objects with components for the image recipe.
	Components []GetImageRecipeComponent `pulumi:"components"`
	// Date the image recipe was created.
	DateCreated string `pulumi:"dateCreated"`
	// Description of the image recipe.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the image recipe.
	Name string `pulumi:"name"`
	// Owner of the image recipe.
	Owner string `pulumi:"owner"`
	// Platform of the image recipe.
	ParentImage string `pulumi:"parentImage"`
	// Platform of the image recipe.
	Platform string `pulumi:"platform"`
	// Key-value map of resource tags for the image recipe.
	Tags map[string]string `pulumi:"tags"`
	// Version of the image recipe.
	Version string `pulumi:"version"`
	// The working directory used during build and test workflows.
	WorkingDirectory string `pulumi:"workingDirectory"`
}

func LookupImageRecipeApply(ctx *pulumi.Context, args LookupImageRecipeApplyInput, opts ...pulumi.InvokeOption) LookupImageRecipeResultOutput {
	return args.ToLookupImageRecipeApplyOutput().ApplyT(func(v LookupImageRecipeArgs) (LookupImageRecipeResult, error) {
		r, err := LookupImageRecipe(ctx, &v, opts...)
		return *r, err

	}).(LookupImageRecipeResultOutput)
}

// LookupImageRecipeApplyInput is an input type that accepts LookupImageRecipeApplyArgs and LookupImageRecipeApplyOutput values.
// You can construct a concrete instance of `LookupImageRecipeApplyInput` via:
//
//          LookupImageRecipeApplyArgs{...}
type LookupImageRecipeApplyInput interface {
	pulumi.Input

	ToLookupImageRecipeApplyOutput() LookupImageRecipeApplyOutput
	ToLookupImageRecipeApplyOutputWithContext(context.Context) LookupImageRecipeApplyOutput
}

// A collection of arguments for invoking getImageRecipe.
type LookupImageRecipeApplyArgs struct {
	// Amazon Resource Name (ARN) of the image recipe.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Key-value map of resource tags for the image recipe.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupImageRecipeApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageRecipeArgs)(nil)).Elem()
}

func (i LookupImageRecipeApplyArgs) ToLookupImageRecipeApplyOutput() LookupImageRecipeApplyOutput {
	return i.ToLookupImageRecipeApplyOutputWithContext(context.Background())
}

func (i LookupImageRecipeApplyArgs) ToLookupImageRecipeApplyOutputWithContext(ctx context.Context) LookupImageRecipeApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupImageRecipeApplyOutput)
}

// A collection of arguments for invoking getImageRecipe.
type LookupImageRecipeApplyOutput struct{ *pulumi.OutputState }

func (LookupImageRecipeApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageRecipeArgs)(nil)).Elem()
}

func (o LookupImageRecipeApplyOutput) ToLookupImageRecipeApplyOutput() LookupImageRecipeApplyOutput {
	return o
}

func (o LookupImageRecipeApplyOutput) ToLookupImageRecipeApplyOutputWithContext(ctx context.Context) LookupImageRecipeApplyOutput {
	return o
}

// Amazon Resource Name (ARN) of the image recipe.
func (o LookupImageRecipeApplyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeArgs) string { return v.Arn }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the image recipe.
func (o LookupImageRecipeApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageRecipeArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getImageRecipe.
type LookupImageRecipeResultOutput struct{ *pulumi.OutputState }

func (LookupImageRecipeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupImageRecipeResult)(nil)).Elem()
}

func (o LookupImageRecipeResultOutput) ToLookupImageRecipeResultOutput() LookupImageRecipeResultOutput {
	return o
}

func (o LookupImageRecipeResultOutput) ToLookupImageRecipeResultOutputWithContext(ctx context.Context) LookupImageRecipeResultOutput {
	return o
}

func (o LookupImageRecipeResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Set of objects with block device mappings for the the image recipe.
func (o LookupImageRecipeResultOutput) BlockDeviceMappings() GetImageRecipeBlockDeviceMappingArrayOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) []GetImageRecipeBlockDeviceMapping { return v.BlockDeviceMappings }).(GetImageRecipeBlockDeviceMappingArrayOutput)
}

// List of objects with components for the image recipe.
func (o LookupImageRecipeResultOutput) Components() GetImageRecipeComponentArrayOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) []GetImageRecipeComponent { return v.Components }).(GetImageRecipeComponentArrayOutput)
}

// Date the image recipe was created.
func (o LookupImageRecipeResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// Description of the image recipe.
func (o LookupImageRecipeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupImageRecipeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the image recipe.
func (o LookupImageRecipeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Owner of the image recipe.
func (o LookupImageRecipeResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.Owner }).(pulumi.StringOutput)
}

// Platform of the image recipe.
func (o LookupImageRecipeResultOutput) ParentImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.ParentImage }).(pulumi.StringOutput)
}

// Platform of the image recipe.
func (o LookupImageRecipeResultOutput) Platform() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.Platform }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the image recipe.
func (o LookupImageRecipeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Version of the image recipe.
func (o LookupImageRecipeResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.Version }).(pulumi.StringOutput)
}

// The working directory used during build and test workflows.
func (o LookupImageRecipeResultOutput) WorkingDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupImageRecipeResult) string { return v.WorkingDirectory }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupImageRecipeApplyOutput{})
	pulumi.RegisterOutputType(LookupImageRecipeResultOutput{})
}
