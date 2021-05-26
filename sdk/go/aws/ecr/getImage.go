// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ECR Image data source allows the details of an image with a particular tag or digest to be retrieved.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "latest"
// 		_, err := ecr.GetImage(ctx, &ecr.GetImageArgs{
// 			ImageTag:       &opt0,
// 			RepositoryName: "my/service",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetImage(ctx *pulumi.Context, args *GetImageArgs, opts ...pulumi.InvokeOption) (*GetImageResult, error) {
	var rv GetImageResult
	err := ctx.Invoke("aws:ecr/getImage:getImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImage.
type GetImageArgs struct {
	// The sha256 digest of the image manifest. At least one of `imageDigest` or `imageTag` must be specified.
	ImageDigest *string `pulumi:"imageDigest"`
	// The tag associated with this image. At least one of `imageDigest` or `imageTag` must be specified.
	ImageTag *string `pulumi:"imageTag"`
	// The ID of the Registry where the repository resides.
	RegistryId *string `pulumi:"registryId"`
	// The name of the ECR Repository.
	RepositoryName string `pulumi:"repositoryName"`
}

// A collection of values returned by getImage.
type GetImageResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	ImageDigest string `pulumi:"imageDigest"`
	// The date and time, expressed as a unix timestamp, at which the current image was pushed to the repository.
	ImagePushedAt int `pulumi:"imagePushedAt"`
	// The size, in bytes, of the image in the repository.
	ImageSizeInBytes int     `pulumi:"imageSizeInBytes"`
	ImageTag         *string `pulumi:"imageTag"`
	// The list of tags associated with this image.
	ImageTags      []string `pulumi:"imageTags"`
	RegistryId     string   `pulumi:"registryId"`
	RepositoryName string   `pulumi:"repositoryName"`
}

func GetImageApply(ctx *pulumi.Context, args GetImageApplyInput, opts ...pulumi.InvokeOption) GetImageResultOutput {
	return args.ToGetImageApplyOutput().ApplyT(func(v GetImageArgs) (GetImageResult, error) {
		r, err := GetImage(ctx, &v, opts...)
		return *r, err

	}).(GetImageResultOutput)
}

// GetImageApplyInput is an input type that accepts GetImageApplyArgs and GetImageApplyOutput values.
// You can construct a concrete instance of `GetImageApplyInput` via:
//
//          GetImageApplyArgs{...}
type GetImageApplyInput interface {
	pulumi.Input

	ToGetImageApplyOutput() GetImageApplyOutput
	ToGetImageApplyOutputWithContext(context.Context) GetImageApplyOutput
}

// A collection of arguments for invoking getImage.
type GetImageApplyArgs struct {
	// The sha256 digest of the image manifest. At least one of `imageDigest` or `imageTag` must be specified.
	ImageDigest pulumi.StringPtrInput `pulumi:"imageDigest"`
	// The tag associated with this image. At least one of `imageDigest` or `imageTag` must be specified.
	ImageTag pulumi.StringPtrInput `pulumi:"imageTag"`
	// The ID of the Registry where the repository resides.
	RegistryId pulumi.StringPtrInput `pulumi:"registryId"`
	// The name of the ECR Repository.
	RepositoryName pulumi.StringInput `pulumi:"repositoryName"`
}

func (GetImageApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageArgs)(nil)).Elem()
}

func (i GetImageApplyArgs) ToGetImageApplyOutput() GetImageApplyOutput {
	return i.ToGetImageApplyOutputWithContext(context.Background())
}

func (i GetImageApplyArgs) ToGetImageApplyOutputWithContext(ctx context.Context) GetImageApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetImageApplyOutput)
}

// A collection of arguments for invoking getImage.
type GetImageApplyOutput struct{ *pulumi.OutputState }

func (GetImageApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageArgs)(nil)).Elem()
}

func (o GetImageApplyOutput) ToGetImageApplyOutput() GetImageApplyOutput {
	return o
}

func (o GetImageApplyOutput) ToGetImageApplyOutputWithContext(ctx context.Context) GetImageApplyOutput {
	return o
}

// The sha256 digest of the image manifest. At least one of `imageDigest` or `imageTag` must be specified.
func (o GetImageApplyOutput) ImageDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageArgs) *string { return v.ImageDigest }).(pulumi.StringPtrOutput)
}

// The tag associated with this image. At least one of `imageDigest` or `imageTag` must be specified.
func (o GetImageApplyOutput) ImageTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageArgs) *string { return v.ImageTag }).(pulumi.StringPtrOutput)
}

// The ID of the Registry where the repository resides.
func (o GetImageApplyOutput) RegistryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageArgs) *string { return v.RegistryId }).(pulumi.StringPtrOutput)
}

// The name of the ECR Repository.
func (o GetImageApplyOutput) RepositoryName() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageArgs) string { return v.RepositoryName }).(pulumi.StringOutput)
}

// A collection of values returned by getImage.
type GetImageResultOutput struct{ *pulumi.OutputState }

func (GetImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetImageResult)(nil)).Elem()
}

func (o GetImageResultOutput) ToGetImageResultOutput() GetImageResultOutput {
	return o
}

func (o GetImageResultOutput) ToGetImageResultOutputWithContext(ctx context.Context) GetImageResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetImageResultOutput) ImageDigest() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.ImageDigest }).(pulumi.StringOutput)
}

// The date and time, expressed as a unix timestamp, at which the current image was pushed to the repository.
func (o GetImageResultOutput) ImagePushedAt() pulumi.IntOutput {
	return o.ApplyT(func(v GetImageResult) int { return v.ImagePushedAt }).(pulumi.IntOutput)
}

// The size, in bytes, of the image in the repository.
func (o GetImageResultOutput) ImageSizeInBytes() pulumi.IntOutput {
	return o.ApplyT(func(v GetImageResult) int { return v.ImageSizeInBytes }).(pulumi.IntOutput)
}

func (o GetImageResultOutput) ImageTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetImageResult) *string { return v.ImageTag }).(pulumi.StringPtrOutput)
}

// The list of tags associated with this image.
func (o GetImageResultOutput) ImageTags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetImageResult) []string { return v.ImageTags }).(pulumi.StringArrayOutput)
}

func (o GetImageResultOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.RegistryId }).(pulumi.StringOutput)
}

func (o GetImageResultOutput) RepositoryName() pulumi.StringOutput {
	return o.ApplyT(func(v GetImageResult) string { return v.RepositoryName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetImageApplyOutput{})
	pulumi.RegisterOutputType(GetImageResultOutput{})
}
