// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get a list of EBS Snapshot IDs matching the specified
// criteria.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ebs.LookupSnapshotIds(ctx, &ebs.LookupSnapshotIdsArgs{
// 			Filters: ebs.getSnapshotIdsFilterArray{
// 				&ebs.LookupSnapshotIdsFilter{
// 					Name: "volume-size",
// 					Values: []string{
// 						"40",
// 					},
// 				},
// 				&ebs.LookupSnapshotIdsFilter{
// 					Name: "tag:Name",
// 					Values: []string{
// 						"Example",
// 					},
// 				},
// 			},
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
	Filters []GetSnapshotIdsFilter `pulumi:"filters"`
	// Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
	Owners []string `pulumi:"owners"`
	// One or more AWS accounts IDs that can create volumes from the snapshot.
	RestorableByUserIds []string `pulumi:"restorableByUserIds"`
}

// A collection of values returned by getSnapshotIds.
type GetSnapshotIdsResult struct {
	Filters []GetSnapshotIdsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	Ids                 []string `pulumi:"ids"`
	Owners              []string `pulumi:"owners"`
	RestorableByUserIds []string `pulumi:"restorableByUserIds"`
}
