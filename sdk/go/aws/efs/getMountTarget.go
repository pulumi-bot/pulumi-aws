// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides information about an Elastic File System Mount Target (EFS).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		mountTargetId := ""
// 		if param := cfg.Get("mountTargetId"); param != "" {
// 			mountTargetId = param
// 		}
// 		opt0 := mountTargetId
// 		_, err := efs.LookupMountTarget(ctx, &efs.LookupMountTargetArgs{
// 			MountTargetId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupMountTarget(ctx *pulumi.Context, args *LookupMountTargetArgs, opts ...pulumi.InvokeOption) (*LookupMountTargetResult, error) {
	var rv LookupMountTargetResult
	err := ctx.Invoke("aws:efs/getMountTarget:getMountTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMountTarget.
type LookupMountTargetArgs struct {
	// ID or ARN of the access point whose mount target that you want to find. It must be included if a `fileSystemId` and `mountTargetId` are not included.
	AccessPointId *string `pulumi:"accessPointId"`
	// ID or ARN of the file system whose mount target that you want to find. It must be included if an `accessPointId` and `mountTargetId` are not included.
	FileSystemId *string `pulumi:"fileSystemId"`
	// ID or ARN of the mount target that you want to find. It must be included in your request if an `accessPointId` and `fileSystemId` are not included.
	MountTargetId *string `pulumi:"mountTargetId"`
}

// A collection of values returned by getMountTarget.
type LookupMountTargetResult struct {
	AccessPointId *string `pulumi:"accessPointId"`
	// The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneId string `pulumi:"availabilityZoneId"`
	// The name of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneName string `pulumi:"availabilityZoneName"`
	// The DNS name for the EFS file system.
	DnsName string `pulumi:"dnsName"`
	// Amazon Resource Name of the file system for which the mount target is intended.
	FileSystemArn string `pulumi:"fileSystemArn"`
	FileSystemId  string `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Address at which the file system may be mounted via the mount target.
	IpAddress string `pulumi:"ipAddress"`
	// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
	MountTargetDnsName string `pulumi:"mountTargetDnsName"`
	MountTargetId      string `pulumi:"mountTargetId"`
	// The ID of the network interface that Amazon EFS created when it created the mount target.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// AWS account ID that owns the resource.
	OwnerId string `pulumi:"ownerId"`
	// List of VPC security group IDs attached to the mount target.
	SecurityGroups []string `pulumi:"securityGroups"`
	// ID of the mount target's subnet.
	SubnetId string `pulumi:"subnetId"`
}

func LookupMountTargetOutput(ctx *pulumi.Context, args LookupMountTargetOutputArgs, opts ...pulumi.InvokeOption) LookupMountTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMountTargetResult, error) {
			args := v.(LookupMountTargetArgs)
			r, err := LookupMountTarget(ctx, &args, opts...)
			return *r, err
		}).(LookupMountTargetResultOutput)
}

// A collection of arguments for invoking getMountTarget.
type LookupMountTargetOutputArgs struct {
	// ID or ARN of the access point whose mount target that you want to find. It must be included if a `fileSystemId` and `mountTargetId` are not included.
	AccessPointId pulumi.StringPtrInput `pulumi:"accessPointId"`
	// ID or ARN of the file system whose mount target that you want to find. It must be included if an `accessPointId` and `mountTargetId` are not included.
	FileSystemId pulumi.StringPtrInput `pulumi:"fileSystemId"`
	// ID or ARN of the mount target that you want to find. It must be included in your request if an `accessPointId` and `fileSystemId` are not included.
	MountTargetId pulumi.StringPtrInput `pulumi:"mountTargetId"`
}

func (LookupMountTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMountTargetArgs)(nil)).Elem()
}

// A collection of values returned by getMountTarget.
type LookupMountTargetResultOutput struct{ *pulumi.OutputState }

func (LookupMountTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMountTargetResult)(nil)).Elem()
}

func (o LookupMountTargetResultOutput) ToLookupMountTargetResultOutput() LookupMountTargetResultOutput {
	return o
}

func (o LookupMountTargetResultOutput) ToLookupMountTargetResultOutputWithContext(ctx context.Context) LookupMountTargetResultOutput {
	return o
}

func (o LookupMountTargetResultOutput) AccessPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMountTargetResult) *string { return v.AccessPointId }).(pulumi.StringPtrOutput)
}

// The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
func (o LookupMountTargetResultOutput) AvailabilityZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.AvailabilityZoneId }).(pulumi.StringOutput)
}

// The name of the Availability Zone (AZ) that the mount target resides in.
func (o LookupMountTargetResultOutput) AvailabilityZoneName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.AvailabilityZoneName }).(pulumi.StringOutput)
}

// The DNS name for the EFS file system.
func (o LookupMountTargetResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.DnsName }).(pulumi.StringOutput)
}

// Amazon Resource Name of the file system for which the mount target is intended.
func (o LookupMountTargetResultOutput) FileSystemArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.FileSystemArn }).(pulumi.StringOutput)
}

func (o LookupMountTargetResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMountTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Address at which the file system may be mounted via the mount target.
func (o LookupMountTargetResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
func (o LookupMountTargetResultOutput) MountTargetDnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.MountTargetDnsName }).(pulumi.StringOutput)
}

func (o LookupMountTargetResultOutput) MountTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.MountTargetId }).(pulumi.StringOutput)
}

// The ID of the network interface that Amazon EFS created when it created the mount target.
func (o LookupMountTargetResultOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

// AWS account ID that owns the resource.
func (o LookupMountTargetResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// List of VPC security group IDs attached to the mount target.
func (o LookupMountTargetResultOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMountTargetResult) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

// ID of the mount target's subnet.
func (o LookupMountTargetResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMountTargetResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMountTargetResultOutput{})
}
