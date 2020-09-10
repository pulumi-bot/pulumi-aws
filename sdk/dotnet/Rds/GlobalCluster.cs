// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Rds
{
    public partial class GlobalCluster : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        [Output("deletionProtection")]
        public Output<bool?> DeletionProtection { get; private set; } = null!;

        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        [Output("globalClusterIdentifier")]
        public Output<string> GlobalClusterIdentifier { get; private set; } = null!;

        [Output("globalClusterMembers")]
        public Output<ImmutableArray<Outputs.GlobalClusterGlobalClusterMember>> GlobalClusterMembers { get; private set; } = null!;

        [Output("globalClusterResourceId")]
        public Output<string> GlobalClusterResourceId { get; private set; } = null!;

        [Output("sourceDbClusterIdentifier")]
        public Output<string> SourceDbClusterIdentifier { get; private set; } = null!;

        [Output("storageEncrypted")]
        public Output<bool?> StorageEncrypted { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalCluster resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalCluster(string name, GlobalClusterArgs args, CustomResourceOptions? options = null)
            : base("aws:rds/globalCluster:GlobalCluster", name, args ?? new GlobalClusterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GlobalCluster(string name, Input<string> id, GlobalClusterState? state = null, CustomResourceOptions? options = null)
            : base("aws:rds/globalCluster:GlobalCluster", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GlobalCluster resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalCluster Get(string name, Input<string> id, GlobalClusterState? state = null, CustomResourceOptions? options = null)
        {
            return new GlobalCluster(name, id, state, options);
        }
    }

    public sealed class GlobalClusterArgs : Pulumi.ResourceArgs
    {
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("globalClusterIdentifier", required: true)]
        public Input<string> GlobalClusterIdentifier { get; set; } = null!;

        [Input("sourceDbClusterIdentifier")]
        public Input<string>? SourceDbClusterIdentifier { get; set; }

        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        public GlobalClusterArgs()
        {
        }
    }

    public sealed class GlobalClusterState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        [Input("engine")]
        public Input<string>? Engine { get; set; }

        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("globalClusterIdentifier")]
        public Input<string>? GlobalClusterIdentifier { get; set; }

        [Input("globalClusterMembers")]
        private InputList<Inputs.GlobalClusterGlobalClusterMemberGetArgs>? _globalClusterMembers;
        public InputList<Inputs.GlobalClusterGlobalClusterMemberGetArgs> GlobalClusterMembers
        {
            get => _globalClusterMembers ?? (_globalClusterMembers = new InputList<Inputs.GlobalClusterGlobalClusterMemberGetArgs>());
            set => _globalClusterMembers = value;
        }

        [Input("globalClusterResourceId")]
        public Input<string>? GlobalClusterResourceId { get; set; }

        [Input("sourceDbClusterIdentifier")]
        public Input<string>? SourceDbClusterIdentifier { get; set; }

        [Input("storageEncrypted")]
        public Input<bool>? StorageEncrypted { get; set; }

        public GlobalClusterState()
        {
        }
    }
}
