// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElastiCache
{
    public partial class ReplicationGroup : Pulumi.CustomResource
    {
        [Output("applyImmediately")]
        public Output<bool> ApplyImmediately { get; private set; } = null!;

        [Output("atRestEncryptionEnabled")]
        public Output<bool?> AtRestEncryptionEnabled { get; private set; } = null!;

        [Output("authToken")]
        public Output<string?> AuthToken { get; private set; } = null!;

        [Output("autoMinorVersionUpgrade")]
        public Output<bool?> AutoMinorVersionUpgrade { get; private set; } = null!;

        [Output("automaticFailoverEnabled")]
        public Output<bool?> AutomaticFailoverEnabled { get; private set; } = null!;

        [Output("availabilityZones")]
        public Output<ImmutableArray<string>> AvailabilityZones { get; private set; } = null!;

        [Output("clusterMode")]
        public Output<Outputs.ReplicationGroupClusterMode> ClusterMode { get; private set; } = null!;

        [Output("configurationEndpointAddress")]
        public Output<string> ConfigurationEndpointAddress { get; private set; } = null!;

        [Output("engine")]
        public Output<string?> Engine { get; private set; } = null!;

        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        [Output("kmsKeyId")]
        public Output<string?> KmsKeyId { get; private set; } = null!;

        [Output("maintenanceWindow")]
        public Output<string> MaintenanceWindow { get; private set; } = null!;

        [Output("memberClusters")]
        public Output<ImmutableArray<string>> MemberClusters { get; private set; } = null!;

        [Output("nodeType")]
        public Output<string> NodeType { get; private set; } = null!;

        [Output("notificationTopicArn")]
        public Output<string?> NotificationTopicArn { get; private set; } = null!;

        [Output("numberCacheClusters")]
        public Output<int> NumberCacheClusters { get; private set; } = null!;

        [Output("parameterGroupName")]
        public Output<string> ParameterGroupName { get; private set; } = null!;

        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        [Output("primaryEndpointAddress")]
        public Output<string> PrimaryEndpointAddress { get; private set; } = null!;

        [Output("replicationGroupDescription")]
        public Output<string> ReplicationGroupDescription { get; private set; } = null!;

        [Output("replicationGroupId")]
        public Output<string> ReplicationGroupId { get; private set; } = null!;

        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        [Output("securityGroupNames")]
        public Output<ImmutableArray<string>> SecurityGroupNames { get; private set; } = null!;

        [Output("snapshotArns")]
        public Output<ImmutableArray<string>> SnapshotArns { get; private set; } = null!;

        [Output("snapshotName")]
        public Output<string?> SnapshotName { get; private set; } = null!;

        [Output("snapshotRetentionLimit")]
        public Output<int?> SnapshotRetentionLimit { get; private set; } = null!;

        [Output("snapshotWindow")]
        public Output<string> SnapshotWindow { get; private set; } = null!;

        [Output("subnetGroupName")]
        public Output<string> SubnetGroupName { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("transitEncryptionEnabled")]
        public Output<bool?> TransitEncryptionEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicationGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicationGroup(string name, ReplicationGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticache/replicationGroup:ReplicationGroup", name, args ?? new ReplicationGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReplicationGroup(string name, Input<string> id, ReplicationGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticache/replicationGroup:ReplicationGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ReplicationGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReplicationGroup Get(string name, Input<string> id, ReplicationGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ReplicationGroup(name, id, state, options);
        }
    }

    public sealed class ReplicationGroupArgs : Pulumi.ResourceArgs
    {
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        [Input("atRestEncryptionEnabled")]
        public Input<bool>? AtRestEncryptionEnabled { get; set; }

        [Input("authToken")]
        public Input<string>? AuthToken { get; set; }

        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        [Input("automaticFailoverEnabled")]
        public Input<bool>? AutomaticFailoverEnabled { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("clusterMode")]
        public Input<Inputs.ReplicationGroupClusterModeArgs>? ClusterMode { get; set; }

        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("maintenanceWindow")]
        public Input<string>? MaintenanceWindow { get; set; }

        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("notificationTopicArn")]
        public Input<string>? NotificationTopicArn { get; set; }

        [Input("numberCacheClusters")]
        public Input<int>? NumberCacheClusters { get; set; }

        [Input("parameterGroupName")]
        public Input<string>? ParameterGroupName { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("replicationGroupDescription", required: true)]
        public Input<string> ReplicationGroupDescription { get; set; } = null!;

        [Input("replicationGroupId")]
        public Input<string>? ReplicationGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroupNames")]
        private InputList<string>? _securityGroupNames;
        public InputList<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new InputList<string>());
            set => _securityGroupNames = value;
        }

        [Input("snapshotArns")]
        private InputList<string>? _snapshotArns;
        public InputList<string> SnapshotArns
        {
            get => _snapshotArns ?? (_snapshotArns = new InputList<string>());
            set => _snapshotArns = value;
        }

        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        [Input("snapshotRetentionLimit")]
        public Input<int>? SnapshotRetentionLimit { get; set; }

        [Input("snapshotWindow")]
        public Input<string>? SnapshotWindow { get; set; }

        [Input("subnetGroupName")]
        public Input<string>? SubnetGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("transitEncryptionEnabled")]
        public Input<bool>? TransitEncryptionEnabled { get; set; }

        public ReplicationGroupArgs()
        {
        }
    }

    public sealed class ReplicationGroupState : Pulumi.ResourceArgs
    {
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        [Input("atRestEncryptionEnabled")]
        public Input<bool>? AtRestEncryptionEnabled { get; set; }

        [Input("authToken")]
        public Input<string>? AuthToken { get; set; }

        [Input("autoMinorVersionUpgrade")]
        public Input<bool>? AutoMinorVersionUpgrade { get; set; }

        [Input("automaticFailoverEnabled")]
        public Input<bool>? AutomaticFailoverEnabled { get; set; }

        [Input("availabilityZones")]
        private InputList<string>? _availabilityZones;
        public InputList<string> AvailabilityZones
        {
            get => _availabilityZones ?? (_availabilityZones = new InputList<string>());
            set => _availabilityZones = value;
        }

        [Input("clusterMode")]
        public Input<Inputs.ReplicationGroupClusterModeGetArgs>? ClusterMode { get; set; }

        [Input("configurationEndpointAddress")]
        public Input<string>? ConfigurationEndpointAddress { get; set; }

        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("maintenanceWindow")]
        public Input<string>? MaintenanceWindow { get; set; }

        [Input("memberClusters")]
        private InputList<string>? _memberClusters;
        public InputList<string> MemberClusters
        {
            get => _memberClusters ?? (_memberClusters = new InputList<string>());
            set => _memberClusters = value;
        }

        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        [Input("notificationTopicArn")]
        public Input<string>? NotificationTopicArn { get; set; }

        [Input("numberCacheClusters")]
        public Input<int>? NumberCacheClusters { get; set; }

        [Input("parameterGroupName")]
        public Input<string>? ParameterGroupName { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("primaryEndpointAddress")]
        public Input<string>? PrimaryEndpointAddress { get; set; }

        [Input("replicationGroupDescription")]
        public Input<string>? ReplicationGroupDescription { get; set; }

        [Input("replicationGroupId")]
        public Input<string>? ReplicationGroupId { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityGroupNames")]
        private InputList<string>? _securityGroupNames;
        public InputList<string> SecurityGroupNames
        {
            get => _securityGroupNames ?? (_securityGroupNames = new InputList<string>());
            set => _securityGroupNames = value;
        }

        [Input("snapshotArns")]
        private InputList<string>? _snapshotArns;
        public InputList<string> SnapshotArns
        {
            get => _snapshotArns ?? (_snapshotArns = new InputList<string>());
            set => _snapshotArns = value;
        }

        [Input("snapshotName")]
        public Input<string>? SnapshotName { get; set; }

        [Input("snapshotRetentionLimit")]
        public Input<int>? SnapshotRetentionLimit { get; set; }

        [Input("snapshotWindow")]
        public Input<string>? SnapshotWindow { get; set; }

        [Input("subnetGroupName")]
        public Input<string>? SubnetGroupName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("transitEncryptionEnabled")]
        public Input<bool>? TransitEncryptionEnabled { get; set; }

        public ReplicationGroupState()
        {
        }
    }
}
