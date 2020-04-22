// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public static class GetSnapshot
    {
        /// <summary>
        /// Use this data source to get information about a DB Snapshot for use when provisioning DB instances
        /// 
        /// &gt; **NOTE:** This data source does not apply to snapshots created on Aurora DB clusters.
        /// See the [`aws.rds.ClusterSnapshot` data source](https://www.terraform.io/docs/providers/aws/d/db_cluster_snapshot.html) for DB Cluster snapshots.
        /// 
        /// {{% examples %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSnapshotResult> InvokeAsync(GetSnapshotArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSnapshotResult>("aws:rds/getSnapshot:getSnapshot", args ?? new GetSnapshotArgs(), options.WithVersion());
    }


    public sealed class GetSnapshotArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Returns the list of snapshots created by the specific db_instance
        /// </summary>
        [Input("dbInstanceIdentifier")]
        public string? DbInstanceIdentifier { get; set; }

        /// <summary>
        /// Returns information on a specific snapshot_id.
        /// </summary>
        [Input("dbSnapshotIdentifier")]
        public string? DbSnapshotIdentifier { get; set; }

        /// <summary>
        /// Set this value to true to include manual DB snapshots that are public and can be
        /// copied or restored by any AWS account, otherwise set this value to false. The default is `false`.
        /// </summary>
        [Input("includePublic")]
        public bool? IncludePublic { get; set; }

        /// <summary>
        /// Set this value to true to include shared manual DB snapshots from other
        /// AWS accounts that this AWS account has been given permission to copy or restore, otherwise set this value to false.
        /// The default is `false`.
        /// </summary>
        [Input("includeShared")]
        public bool? IncludeShared { get; set; }

        /// <summary>
        /// If more than one result is returned, use the most
        /// recent Snapshot.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// The type of snapshots to be returned. If you don't specify a SnapshotType
        /// value, then both automated and manual snapshots are returned. Shared and public DB snapshots are not
        /// included in the returned results by default. Possible values are, `automated`, `manual`, `shared` and `public`.
        /// </summary>
        [Input("snapshotType")]
        public string? SnapshotType { get; set; }

        public GetSnapshotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSnapshotResult
    {
        /// <summary>
        /// Specifies the allocated storage size in gigabytes (GB).
        /// </summary>
        public readonly int AllocatedStorage;
        /// <summary>
        /// Specifies the name of the Availability Zone the DB instance was located in at the time of the DB snapshot.
        /// </summary>
        public readonly string AvailabilityZone;
        public readonly string? DbInstanceIdentifier;
        /// <summary>
        /// The Amazon Resource Name (ARN) for the DB snapshot.
        /// </summary>
        public readonly string DbSnapshotArn;
        public readonly string? DbSnapshotIdentifier;
        /// <summary>
        /// Specifies whether the DB snapshot is encrypted.
        /// </summary>
        public readonly bool Encrypted;
        /// <summary>
        /// Specifies the name of the database engine.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// Specifies the version of the database engine.
        /// </summary>
        public readonly string EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludePublic;
        public readonly bool? IncludeShared;
        /// <summary>
        /// Specifies the Provisioned IOPS (I/O operations per second) value of the DB instance at the time of the snapshot.
        /// </summary>
        public readonly int Iops;
        /// <summary>
        /// The ARN for the KMS encryption key.
        /// </summary>
        public readonly string KmsKeyId;
        /// <summary>
        /// License model information for the restored DB instance.
        /// </summary>
        public readonly string LicenseModel;
        public readonly bool? MostRecent;
        /// <summary>
        /// Provides the option group name for the DB snapshot.
        /// </summary>
        public readonly string OptionGroupName;
        public readonly int Port;
        /// <summary>
        /// Provides the time when the snapshot was taken, in Universal Coordinated Time (UTC).
        /// </summary>
        public readonly string SnapshotCreateTime;
        public readonly string? SnapshotType;
        /// <summary>
        /// The DB snapshot Arn that the DB snapshot was copied from. It only has value in case of cross customer or cross region copy.
        /// </summary>
        public readonly string SourceDbSnapshotIdentifier;
        /// <summary>
        /// The region that the DB snapshot was created in or copied from.
        /// </summary>
        public readonly string SourceRegion;
        /// <summary>
        /// Specifies the status of this DB snapshot.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Specifies the storage type associated with DB snapshot.
        /// </summary>
        public readonly string StorageType;
        /// <summary>
        /// Specifies the ID of the VPC associated with the DB snapshot.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private GetSnapshotResult(
            int allocatedStorage,

            string availabilityZone,

            string? dbInstanceIdentifier,

            string dbSnapshotArn,

            string? dbSnapshotIdentifier,

            bool encrypted,

            string engine,

            string engineVersion,

            string id,

            bool? includePublic,

            bool? includeShared,

            int iops,

            string kmsKeyId,

            string licenseModel,

            bool? mostRecent,

            string optionGroupName,

            int port,

            string snapshotCreateTime,

            string? snapshotType,

            string sourceDbSnapshotIdentifier,

            string sourceRegion,

            string status,

            string storageType,

            string vpcId)
        {
            AllocatedStorage = allocatedStorage;
            AvailabilityZone = availabilityZone;
            DbInstanceIdentifier = dbInstanceIdentifier;
            DbSnapshotArn = dbSnapshotArn;
            DbSnapshotIdentifier = dbSnapshotIdentifier;
            Encrypted = encrypted;
            Engine = engine;
            EngineVersion = engineVersion;
            Id = id;
            IncludePublic = includePublic;
            IncludeShared = includeShared;
            Iops = iops;
            KmsKeyId = kmsKeyId;
            LicenseModel = licenseModel;
            MostRecent = mostRecent;
            OptionGroupName = optionGroupName;
            Port = port;
            SnapshotCreateTime = snapshotCreateTime;
            SnapshotType = snapshotType;
            SourceDbSnapshotIdentifier = sourceDbSnapshotIdentifier;
            SourceRegion = sourceRegion;
            Status = status;
            StorageType = storageType;
            VpcId = vpcId;
        }
    }
}
