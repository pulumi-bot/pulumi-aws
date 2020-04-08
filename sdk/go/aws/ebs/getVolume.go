// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about an EBS volume for use in other
// resources.
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("aws:ebs/getVolume:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolume.
type LookupVolumeArgs struct {
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-volumes in the AWS CLI reference][1].
	Filters []GetVolumeFilterArgs `pulumi:"filters"`
	// If more than one result is returned, use the most
	// recent Volume.
	MostRecent *bool `pulumi:"mostRecent"`
	// A mapping of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getVolume.
type LookupVolumeResult struct {
	// The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
	Arn string `pulumi:"arn"`
	// The AZ where the EBS volume exists.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Whether the disk is encrypted.
	Encrypted bool              `pulumi:"encrypted"`
	Filters   []GetVolumeFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The amount of IOPS for the disk.
	Iops int `pulumi:"iops"`
	// The ARN for the KMS encryption key.
	KmsKeyId   string `pulumi:"kmsKeyId"`
	MostRecent *bool  `pulumi:"mostRecent"`
	// The size of the drive in GiBs.
	Size int `pulumi:"size"`
	// The snapshotId the EBS volume is based off.
	SnapshotId string `pulumi:"snapshotId"`
	// A mapping of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The volume ID (e.g. vol-59fcb34e).
	VolumeId string `pulumi:"volumeId"`
	// The type of EBS volume.
	VolumeType string `pulumi:"volumeType"`
}
