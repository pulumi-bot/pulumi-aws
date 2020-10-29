// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about an Elastic File System (EFS) Access Point.
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
