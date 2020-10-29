// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information about a DB Cluster Snapshot for use when provisioning DB clusters.
//
// > **NOTE:** This data source does not apply to snapshots created on DB Instances.
// See the `rds.Snapshot` data source for DB Instance snapshots.
func LookupClusterSnapshot(ctx *pulumi.Context, args *LookupClusterSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupClusterSnapshotResult, error) {
	var rv LookupClusterSnapshotResult
	err := ctx.Invoke("aws:rds/getClusterSnapshot:getClusterSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClusterSnapshot.
type LookupClusterSnapshotArgs struct {
	// Returns the list of snapshots created by the specific db_cluster
	DbClusterIdentifier *string `pulumi:"dbClusterIdentifier"`
	// Returns information on a specific snapshot_id.
	DbClusterSnapshotIdentifier *string `pulumi:"dbClusterSnapshotIdentifier"`
	// Set this value to true to include manual DB Cluster Snapshots that are public and can be
	// copied or restored by any AWS account, otherwise set this value to false. The default is `false`.
	IncludePublic *bool `pulumi:"includePublic"`
	// Set this value to true to include shared manual DB Cluster Snapshots from other
	// AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false.
	// The default is `false`.
	IncludeShared *bool `pulumi:"includeShared"`
	// If more than one result is returned, use the most recent Snapshot.
	MostRecent *bool `pulumi:"mostRecent"`
	// The type of snapshots to be returned. If you don't specify a SnapshotType
	// value, then both automated and manual DB cluster snapshots are returned. Shared and public DB Cluster Snapshots are not
	// included in the returned results by default. Possible values are, `automated`, `manual`, `shared` and `public`.
	SnapshotType *string `pulumi:"snapshotType"`
	// A map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getClusterSnapshot.
type LookupClusterSnapshotResult struct {
	// Specifies the allocated storage size in gigabytes (GB).
	AllocatedStorage int `pulumi:"allocatedStorage"`
	// List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Specifies the DB cluster identifier of the DB cluster that this DB cluster snapshot was created from.
	DbClusterIdentifier *string `pulumi:"dbClusterIdentifier"`
	// The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
	DbClusterSnapshotArn        string  `pulumi:"dbClusterSnapshotArn"`
	DbClusterSnapshotIdentifier *string `pulumi:"dbClusterSnapshotIdentifier"`
	// Specifies the name of the database engine.
	Engine string `pulumi:"engine"`
	// Version of the database engine for this DB cluster snapshot.
	EngineVersion string `pulumi:"engineVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id            string `pulumi:"id"`
	IncludePublic *bool  `pulumi:"includePublic"`
	IncludeShared *bool  `pulumi:"includeShared"`
	// If storageEncrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
	KmsKeyId string `pulumi:"kmsKeyId"`
	// License model information for the restored DB cluster.
	LicenseModel string `pulumi:"licenseModel"`
	MostRecent   *bool  `pulumi:"mostRecent"`
	// Port that the DB cluster was listening on at the time of the snapshot.
	Port int `pulumi:"port"`
	// Time when the snapshot was taken, in Universal Coordinated Time (UTC).
	SnapshotCreateTime         string  `pulumi:"snapshotCreateTime"`
	SnapshotType               *string `pulumi:"snapshotType"`
	SourceDbClusterSnapshotArn string  `pulumi:"sourceDbClusterSnapshotArn"`
	// The status of this DB Cluster Snapshot.
	Status string `pulumi:"status"`
	// Specifies whether the DB cluster snapshot is encrypted.
	StorageEncrypted bool `pulumi:"storageEncrypted"`
	// A map of tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID associated with the DB cluster snapshot.
	VpcId string `pulumi:"vpcId"`
}
