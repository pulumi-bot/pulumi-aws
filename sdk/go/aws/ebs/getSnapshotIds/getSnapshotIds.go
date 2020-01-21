// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getSnapshotIds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ebs/getSnapshotIdsFilter"
)

// Use this data source to get a list of EBS Snapshot IDs matching the specified
// criteria.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_snapshot_ids.html.markdown.
func GetSnapshotIds(ctx *pulumi.Context, args *GetSnapshotIdsArgs, opts ...pulumi.InvokeOption) (*GetSnapshotIdsResult, error) {
	var rv GetSnapshotIdsResult
	err := ctx.Invoke("aws:ebs/getSnapshotIds:getSnapshotIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshotIds.
type GetSnapshotIdsArgs struct {
	// One or more name/value pairs to filter off of. There are
	// several valid keys, for a full reference, check out
	// [describe-volumes in the AWS CLI reference][1].
	Filters []ebsgetSnapshotIdsFilter.GetSnapshotIdsFilter `pulumi:"filters"`
	// Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
	Owners []string `pulumi:"owners"`
	// One or more AWS accounts IDs that can create volumes from the snapshot.
	RestorableByUserIds []string `pulumi:"restorableByUserIds"`
}


// A collection of values returned by getSnapshotIds.
type GetSnapshotIdsResult struct {
	Filters []ebsgetSnapshotIdsFilter.GetSnapshotIdsFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	Owners []string `pulumi:"owners"`
	RestorableByUserIds []string `pulumi:"restorableByUserIds"`
}

