// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to get information about an EBS Snapshot for use when provisioning EBS Volumes
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_snapshot.html.markdown.
func LookupSnapshot(ctx *pulumi.Context, args *GetSnapshotArgs) (*GetSnapshotResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["filters"] = args.Filters
		inputs["mostRecent"] = args.MostRecent
		inputs["owners"] = args.Owners
		inputs["restorableByUserIds"] = args.RestorableByUserIds
		inputs["snapshotIds"] = args.SnapshotIds
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("aws:ebs/getSnapshot:getSnapshot", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSnapshotResult{
		DataEncryptionKeyId: outputs["dataEncryptionKeyId"],
		Description: outputs["description"],
		Encrypted: outputs["encrypted"],
		Filters: outputs["filters"],
		KmsKeyId: outputs["kmsKeyId"],
		MostRecent: outputs["mostRecent"],
		OwnerAlias: outputs["ownerAlias"],
		OwnerId: outputs["ownerId"],
		Owners: outputs["owners"],
		RestorableByUserIds: outputs["restorableByUserIds"],
		SnapshotId: outputs["snapshotId"],
		SnapshotIds: outputs["snapshotIds"],
		State: outputs["state"],
		Tags: outputs["tags"],
		VolumeId: outputs["volumeId"],
		VolumeSize: outputs["volumeSize"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSnapshot.
type GetSnapshotArgs struct {
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-snapshots in the AWS CLI reference][1].
	Filters interface{}
	// If more than one result is returned, use the most recent snapshot.
	MostRecent interface{}
	// Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
	Owners interface{}
	// One or more AWS accounts IDs that can create volumes from the snapshot.
	RestorableByUserIds interface{}
	// Returns information on a specific snapshot_id.
	SnapshotIds interface{}
	Tags interface{}
}

// A collection of values returned by getSnapshot.
type GetSnapshotResult struct {
	// The data encryption key identifier for the snapshot.
	DataEncryptionKeyId interface{}
	// A description for the snapshot
	Description interface{}
	// Whether the snapshot is encrypted.
	Encrypted interface{}
	Filters interface{}
	// The ARN for the KMS encryption key.
	KmsKeyId interface{}
	MostRecent interface{}
	// Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
	OwnerAlias interface{}
	// The AWS account ID of the EBS snapshot owner.
	OwnerId interface{}
	Owners interface{}
	RestorableByUserIds interface{}
	// The snapshot ID (e.g. snap-59fcb34e).
	SnapshotId interface{}
	SnapshotIds interface{}
	// The snapshot state.
	State interface{}
	// A mapping of tags for the resource.
	Tags interface{}
	// The volume ID (e.g. vol-59fcb34e).
	VolumeId interface{}
	// The size of the drive in GiBs.
	VolumeSize interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
