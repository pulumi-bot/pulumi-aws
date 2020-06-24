// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about an Elastic File System Mount Target (EFS).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/efs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := efs.LookupMountTarget(ctx, &efs.LookupMountTargetArgs{
// 			MountTargetId: mountTargetId,
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
	// ID of the mount target that you want to have described
	MountTargetId string `pulumi:"mountTargetId"`
}

// A collection of values returned by getMountTarget.
type LookupMountTargetResult struct {
	// The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneId string `pulumi:"availabilityZoneId"`
	// The name of the Availability Zone (AZ) that the mount target resides in.
	AvailabilityZoneName string `pulumi:"availabilityZoneName"`
	// The DNS name for the EFS file system.
	DnsName string `pulumi:"dnsName"`
	// Amazon Resource Name of the file system for which the mount target is intended.
	FileSystemArn string `pulumi:"fileSystemArn"`
	// ID of the file system for which the mount target is intended.
	FileSystemId string `pulumi:"fileSystemId"`
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
