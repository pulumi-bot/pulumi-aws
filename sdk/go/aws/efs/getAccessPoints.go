// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about multiple Elastic File System (EFS) Access Points.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := efs.GetAccessPoints(ctx, &efs.GetAccessPointsArgs{
// 			FileSystemId: "fs-12345678",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetAccessPoints(ctx *pulumi.Context, args *GetAccessPointsArgs, opts ...pulumi.InvokeOption) (*GetAccessPointsResult, error) {
	var rv GetAccessPointsResult
	err := ctx.Invoke("aws:efs/getAccessPoints:getAccessPoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessPoints.
type GetAccessPointsArgs struct {
	// EFS File System identifier.
	FileSystemId string `pulumi:"fileSystemId"`
}

// A collection of values returned by getAccessPoints.
type GetAccessPointsResult struct {
	// Set of Amazon Resource Names (ARNs).
	Arns         []string `pulumi:"arns"`
	FileSystemId string   `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of identifiers.
	Ids []string `pulumi:"ids"`
}

func GetAccessPointsOutput(ctx *pulumi.Context, args GetAccessPointsOutputArgs, opts ...pulumi.InvokeOption) GetAccessPointsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAccessPointsResult, error) {
			args := v.(GetAccessPointsArgs)
			r, err := GetAccessPoints(ctx, &args, opts...)
			return *r, err
		}).(GetAccessPointsResultOutput)
}

// A collection of arguments for invoking getAccessPoints.
type GetAccessPointsOutputArgs struct {
	// EFS File System identifier.
	FileSystemId pulumi.StringInput `pulumi:"fileSystemId"`
}

func (GetAccessPointsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessPointsArgs)(nil)).Elem()
}

// A collection of values returned by getAccessPoints.
type GetAccessPointsResultOutput struct{ *pulumi.OutputState }

func (GetAccessPointsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAccessPointsResult)(nil)).Elem()
}

func (o GetAccessPointsResultOutput) ToGetAccessPointsResultOutput() GetAccessPointsResultOutput {
	return o
}

func (o GetAccessPointsResultOutput) ToGetAccessPointsResultOutputWithContext(ctx context.Context) GetAccessPointsResultOutput {
	return o
}

// Set of Amazon Resource Names (ARNs).
func (o GetAccessPointsResultOutput) Arns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccessPointsResult) []string { return v.Arns }).(pulumi.StringArrayOutput)
}

func (o GetAccessPointsResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessPointsResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAccessPointsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAccessPointsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Set of identifiers.
func (o GetAccessPointsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAccessPointsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAccessPointsResultOutput{})
}
