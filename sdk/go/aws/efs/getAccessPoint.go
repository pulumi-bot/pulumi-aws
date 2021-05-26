// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about an Elastic File System (EFS) Access Point.
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
// 		_, err := efs.LookupAccessPoint(ctx, &efs.LookupAccessPointArgs{
// 			AccessPointId: "fsap-12345678",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupAccessPoint(ctx *pulumi.Context, args *LookupAccessPointArgs, opts ...pulumi.InvokeOption) (*LookupAccessPointResult, error) {
	var rv LookupAccessPointResult
	err := ctx.Invoke("aws:efs/getAccessPoint:getAccessPoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessPoint.
type LookupAccessPointArgs struct {
	// The ID that identifies the file system.
	AccessPointId string `pulumi:"accessPointId"`
	// Key-value mapping of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getAccessPoint.
type LookupAccessPointResult struct {
	AccessPointId string `pulumi:"accessPointId"`
	// Amazon Resource Name of the file system.
	Arn string `pulumi:"arn"`
	// Amazon Resource Name of the file system.
	FileSystemArn string `pulumi:"fileSystemArn"`
	// The ID of the file system for which the access point is intended.
	FileSystemId string `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	OwnerId string `pulumi:"ownerId"`
	// Single element list containing operating system user and group applied to all file system requests made using the access point.
	PosixUsers      []GetAccessPointPosixUser     `pulumi:"posixUsers"`
	RootDirectories []GetAccessPointRootDirectory `pulumi:"rootDirectories"`
	// Key-value mapping of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

func LookupAccessPointApply(ctx *pulumi.Context, args LookupAccessPointApplyInput, opts ...pulumi.InvokeOption) LookupAccessPointResultOutput {
	return args.ToLookupAccessPointApplyOutput().ApplyT(func(v LookupAccessPointArgs) (LookupAccessPointResult, error) {
		r, err := LookupAccessPoint(ctx, &v, opts...)
		return *r, err

	}).(LookupAccessPointResultOutput)
}

// LookupAccessPointApplyInput is an input type that accepts LookupAccessPointApplyArgs and LookupAccessPointApplyOutput values.
// You can construct a concrete instance of `LookupAccessPointApplyInput` via:
//
//          LookupAccessPointApplyArgs{...}
type LookupAccessPointApplyInput interface {
	pulumi.Input

	ToLookupAccessPointApplyOutput() LookupAccessPointApplyOutput
	ToLookupAccessPointApplyOutputWithContext(context.Context) LookupAccessPointApplyOutput
}

// A collection of arguments for invoking getAccessPoint.
type LookupAccessPointApplyArgs struct {
	// The ID that identifies the file system.
	AccessPointId pulumi.StringInput `pulumi:"accessPointId"`
	// Key-value mapping of resource tags.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupAccessPointApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointArgs)(nil)).Elem()
}

func (i LookupAccessPointApplyArgs) ToLookupAccessPointApplyOutput() LookupAccessPointApplyOutput {
	return i.ToLookupAccessPointApplyOutputWithContext(context.Background())
}

func (i LookupAccessPointApplyArgs) ToLookupAccessPointApplyOutputWithContext(ctx context.Context) LookupAccessPointApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupAccessPointApplyOutput)
}

// A collection of arguments for invoking getAccessPoint.
type LookupAccessPointApplyOutput struct{ *pulumi.OutputState }

func (LookupAccessPointApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointArgs)(nil)).Elem()
}

func (o LookupAccessPointApplyOutput) ToLookupAccessPointApplyOutput() LookupAccessPointApplyOutput {
	return o
}

func (o LookupAccessPointApplyOutput) ToLookupAccessPointApplyOutputWithContext(ctx context.Context) LookupAccessPointApplyOutput {
	return o
}

// The ID that identifies the file system.
func (o LookupAccessPointApplyOutput) AccessPointId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPointArgs) string { return v.AccessPointId }).(pulumi.StringOutput)
}

// Key-value mapping of resource tags.
func (o LookupAccessPointApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccessPointArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getAccessPoint.
type LookupAccessPointResultOutput struct{ *pulumi.OutputState }

func (LookupAccessPointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessPointResult)(nil)).Elem()
}

func (o LookupAccessPointResultOutput) ToLookupAccessPointResultOutput() LookupAccessPointResultOutput {
	return o
}

func (o LookupAccessPointResultOutput) ToLookupAccessPointResultOutputWithContext(ctx context.Context) LookupAccessPointResultOutput {
	return o
}

func (o LookupAccessPointResultOutput) AccessPointId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPointResult) string { return v.AccessPointId }).(pulumi.StringOutput)
}

// Amazon Resource Name of the file system.
func (o LookupAccessPointResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPointResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Amazon Resource Name of the file system.
func (o LookupAccessPointResultOutput) FileSystemArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPointResult) string { return v.FileSystemArn }).(pulumi.StringOutput)
}

// The ID of the file system for which the access point is intended.
func (o LookupAccessPointResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPointResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAccessPointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPointResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAccessPointResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessPointResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// Single element list containing operating system user and group applied to all file system requests made using the access point.
func (o LookupAccessPointResultOutput) PosixUsers() GetAccessPointPosixUserArrayOutput {
	return o.ApplyT(func(v LookupAccessPointResult) []GetAccessPointPosixUser { return v.PosixUsers }).(GetAccessPointPosixUserArrayOutput)
}

func (o LookupAccessPointResultOutput) RootDirectories() GetAccessPointRootDirectoryArrayOutput {
	return o.ApplyT(func(v LookupAccessPointResult) []GetAccessPointRootDirectory { return v.RootDirectories }).(GetAccessPointRootDirectoryArrayOutput)
}

// Key-value mapping of resource tags.
func (o LookupAccessPointResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAccessPointResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessPointApplyOutput{})
	pulumi.RegisterOutputType(LookupAccessPointResultOutput{})
}
