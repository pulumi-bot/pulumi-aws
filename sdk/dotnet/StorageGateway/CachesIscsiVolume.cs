// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.StorageGateway
{
    public partial class CachesIscsiVolume : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("chapEnabled")]
        public Output<bool> ChapEnabled { get; private set; } = null!;

        [Output("gatewayArn")]
        public Output<string> GatewayArn { get; private set; } = null!;

        [Output("lunNumber")]
        public Output<int> LunNumber { get; private set; } = null!;

        [Output("networkInterfaceId")]
        public Output<string> NetworkInterfaceId { get; private set; } = null!;

        [Output("networkInterfacePort")]
        public Output<int> NetworkInterfacePort { get; private set; } = null!;

        [Output("snapshotId")]
        public Output<string?> SnapshotId { get; private set; } = null!;

        [Output("sourceVolumeArn")]
        public Output<string?> SourceVolumeArn { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("targetArn")]
        public Output<string> TargetArn { get; private set; } = null!;

        [Output("targetName")]
        public Output<string> TargetName { get; private set; } = null!;

        [Output("volumeArn")]
        public Output<string> VolumeArn { get; private set; } = null!;

        [Output("volumeId")]
        public Output<string> VolumeId { get; private set; } = null!;

        [Output("volumeSizeInBytes")]
        public Output<int> VolumeSizeInBytes { get; private set; } = null!;


        /// <summary>
        /// Create a CachesIscsiVolume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CachesIscsiVolume(string name, CachesIscsiVolumeArgs args, CustomResourceOptions? options = null)
            : base("aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume", name, args ?? new CachesIscsiVolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CachesIscsiVolume(string name, Input<string> id, CachesIscsiVolumeState? state = null, CustomResourceOptions? options = null)
            : base("aws:storagegateway/cachesIscsiVolume:CachesIscsiVolume", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CachesIscsiVolume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CachesIscsiVolume Get(string name, Input<string> id, CachesIscsiVolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new CachesIscsiVolume(name, id, state, options);
        }
    }

    public sealed class CachesIscsiVolumeArgs : Pulumi.ResourceArgs
    {
        [Input("gatewayArn", required: true)]
        public Input<string> GatewayArn { get; set; } = null!;

        [Input("networkInterfaceId", required: true)]
        public Input<string> NetworkInterfaceId { get; set; } = null!;

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("sourceVolumeArn")]
        public Input<string>? SourceVolumeArn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targetName", required: true)]
        public Input<string> TargetName { get; set; } = null!;

        [Input("volumeSizeInBytes", required: true)]
        public Input<int> VolumeSizeInBytes { get; set; } = null!;

        public CachesIscsiVolumeArgs()
        {
        }
    }

    public sealed class CachesIscsiVolumeState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("chapEnabled")]
        public Input<bool>? ChapEnabled { get; set; }

        [Input("gatewayArn")]
        public Input<string>? GatewayArn { get; set; }

        [Input("lunNumber")]
        public Input<int>? LunNumber { get; set; }

        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        [Input("networkInterfacePort")]
        public Input<int>? NetworkInterfacePort { get; set; }

        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("sourceVolumeArn")]
        public Input<string>? SourceVolumeArn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targetArn")]
        public Input<string>? TargetArn { get; set; }

        [Input("targetName")]
        public Input<string>? TargetName { get; set; }

        [Input("volumeArn")]
        public Input<string>? VolumeArn { get; set; }

        [Input("volumeId")]
        public Input<string>? VolumeId { get; set; }

        [Input("volumeSizeInBytes")]
        public Input<int>? VolumeSizeInBytes { get; set; }

        public CachesIscsiVolumeState()
        {
        }
    }
}
