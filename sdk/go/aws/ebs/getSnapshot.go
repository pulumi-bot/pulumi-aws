// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about an EBS Snapshot for use when provisioning EBS Volumes
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ebs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ebs.LookupSnapshot(ctx, &ebs.LookupSnapshotArgs{
// 			Filters: []ebs.GetSnapshotFilter{
// 				ebs.GetSnapshotFilter{
// 					Name: "volume-size",
// 					Values: []string{
// 						"40",
// 					},
// 				},
// 				ebs.GetSnapshotFilter{
// 					Name: "tag:Name",
// 					Values: []string{
// 						"Example",
// 					},
// 				},
// 			},
// 			MostRecent: true,
// 			Owners: []string{
// 				"self",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("aws:ebs/getSnapshot:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotArgs struct {
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-snapshots in the AWS CLI reference][1].
	Filters []GetSnapshotFilter `pulumi:"filters"`
	// If more than one result is returned, use the most recent snapshot.
	MostRecent *bool `pulumi:"mostRecent"`
	// Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
	Owners []string `pulumi:"owners"`
	// One or more AWS accounts IDs that can create volumes from the snapshot.
	RestorableByUserIds []string `pulumi:"restorableByUserIds"`
	// Returns information on a specific snapshot_id.
	SnapshotIds []string `pulumi:"snapshotIds"`
	// A map of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResult struct {
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId string `pulumi:"dataEncryptionKeyId"`
	// A description for the snapshot
	Description string `pulumi:"description"`
	// Whether the snapshot is encrypted.
	Encrypted bool                `pulumi:"encrypted"`
	Filters   []GetSnapshotFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ARN for the KMS encryption key.
	KmsKeyId   string `pulumi:"kmsKeyId"`
	MostRecent *bool  `pulumi:"mostRecent"`
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias string `pulumi:"ownerAlias"`
	// The AWS account ID of the EBS snapshot owner.
	OwnerId             string   `pulumi:"ownerId"`
	Owners              []string `pulumi:"owners"`
	RestorableByUserIds []string `pulumi:"restorableByUserIds"`
	// The snapshot ID (e.g. snap-59fcb34e).
	SnapshotId  string   `pulumi:"snapshotId"`
	SnapshotIds []string `pulumi:"snapshotIds"`
	// The snapshot state.
	State string `pulumi:"state"`
	// A map of tags for the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The volume ID (e.g. vol-59fcb34e).
	VolumeId string `pulumi:"volumeId"`
	// The size of the drive in GiBs.
	VolumeSize int `pulumi:"volumeSize"`
}
