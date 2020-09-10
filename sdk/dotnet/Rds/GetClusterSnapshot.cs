// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public static class GetClusterSnapshot
    {
        public static Task<GetClusterSnapshotResult> InvokeAsync(GetClusterSnapshotArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClusterSnapshotResult>("aws:rds/getClusterSnapshot:getClusterSnapshot", args ?? new GetClusterSnapshotArgs(), options.WithVersion());
    }


    public sealed class GetClusterSnapshotArgs : Pulumi.InvokeArgs
    {
        [Input("dbClusterIdentifier")]
        public string? DbClusterIdentifier { get; set; }

        [Input("dbClusterSnapshotIdentifier")]
        public string? DbClusterSnapshotIdentifier { get; set; }

        [Input("includePublic")]
        public bool? IncludePublic { get; set; }

        [Input("includeShared")]
        public bool? IncludeShared { get; set; }

        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        [Input("snapshotType")]
        public string? SnapshotType { get; set; }

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetClusterSnapshotArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClusterSnapshotResult
    {
        public readonly int AllocatedStorage;
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly string? DbClusterIdentifier;
        public readonly string DbClusterSnapshotArn;
        public readonly string? DbClusterSnapshotIdentifier;
        public readonly string Engine;
        public readonly string EngineVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool? IncludePublic;
        public readonly bool? IncludeShared;
        public readonly string KmsKeyId;
        public readonly string LicenseModel;
        public readonly bool? MostRecent;
        public readonly int Port;
        public readonly string SnapshotCreateTime;
        public readonly string? SnapshotType;
        public readonly string SourceDbClusterSnapshotArn;
        public readonly string Status;
        public readonly bool StorageEncrypted;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string VpcId;

        [OutputConstructor]
        private GetClusterSnapshotResult(
            int allocatedStorage,

            ImmutableArray<string> availabilityZones,

            string? dbClusterIdentifier,

            string dbClusterSnapshotArn,

            string? dbClusterSnapshotIdentifier,

            string engine,

            string engineVersion,

            string id,

            bool? includePublic,

            bool? includeShared,

            string kmsKeyId,

            string licenseModel,

            bool? mostRecent,

            int port,

            string snapshotCreateTime,

            string? snapshotType,

            string sourceDbClusterSnapshotArn,

            string status,

            bool storageEncrypted,

            ImmutableDictionary<string, string> tags,

            string vpcId)
        {
            AllocatedStorage = allocatedStorage;
            AvailabilityZones = availabilityZones;
            DbClusterIdentifier = dbClusterIdentifier;
            DbClusterSnapshotArn = dbClusterSnapshotArn;
            DbClusterSnapshotIdentifier = dbClusterSnapshotIdentifier;
            Engine = engine;
            EngineVersion = engineVersion;
            Id = id;
            IncludePublic = includePublic;
            IncludeShared = includeShared;
            KmsKeyId = kmsKeyId;
            LicenseModel = licenseModel;
            MostRecent = mostRecent;
            Port = port;
            SnapshotCreateTime = snapshotCreateTime;
            SnapshotType = snapshotType;
            SourceDbClusterSnapshotArn = sourceDbClusterSnapshotArn;
            Status = status;
            StorageEncrypted = storageEncrypted;
            Tags = tags;
            VpcId = vpcId;
        }
    }
}
