// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ebs
{
    /// <summary>
    /// Manages a single EBS volume.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Ebs.Volume("example", new Aws.Ebs.VolumeArgs
    ///         {
    ///             AvailabilityZone = "us-west-2a",
    ///             Size = 40,
    ///             Tags = 
    ///             {
    ///                 { "Name", "HelloWorld" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// &gt; **NOTE**: One of `size` or `snapshot_id` is required when specifying an EBS volume
    /// </summary>
    public partial class Volume : Pulumi.CustomResource
    {
        /// <summary>
        /// The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The AZ where the EBS volume will exist.
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Output("encrypted")]
        public Output<bool> Encrypted { get; private set; } = null!;

        /// <summary>
        /// The amount of IOPS to provision for the disk.
        /// </summary>
        [Output("iops")]
        public Output<int> Iops { get; private set; } = null!;

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported exclusively on `io1` volumes.
        /// </summary>
        [Output("multiAttachEnabled")]
        public Output<bool?> MultiAttachEnabled { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Output("outpostArn")]
        public Output<string?> OutpostArn { get; private set; } = null!;

        /// <summary>
        /// The size of the drive in GiBs.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// A snapshot to base the EBS volume off of.
        /// </summary>
        [Output("snapshotId")]
        public Output<string> SnapshotId { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of EBS volume. Can be "standard", "gp2", "io1", "sc1" or "st1" (Default: "gp2").
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Volume resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Volume(string name, VolumeArgs args, CustomResourceOptions? options = null)
            : base("aws:ebs/volume:Volume", name, args ?? new VolumeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Volume(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
            : base("aws:ebs/volume:Volume", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Volume resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Volume Get(string name, Input<string> id, VolumeState? state = null, CustomResourceOptions? options = null)
        {
            return new Volume(name, id, state, options);
        }
    }

    public sealed class VolumeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AZ where the EBS volume will exist.
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The amount of IOPS to provision for the disk.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported exclusively on `io1` volumes.
        /// </summary>
        [Input("multiAttachEnabled")]
        public Input<bool>? MultiAttachEnabled { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The size of the drive in GiBs.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// A snapshot to base the EBS volume off of.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of EBS volume. Can be "standard", "gp2", "io1", "sc1" or "st1" (Default: "gp2").
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VolumeArgs()
        {
        }
    }

    public sealed class VolumeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The AZ where the EBS volume will exist.
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// If true, the disk will be encrypted.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// The amount of IOPS to provision for the disk.
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true.
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported exclusively on `io1` volumes.
        /// </summary>
        [Input("multiAttachEnabled")]
        public Input<bool>? MultiAttachEnabled { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) of the Outpost.
        /// </summary>
        [Input("outpostArn")]
        public Input<string>? OutpostArn { get; set; }

        /// <summary>
        /// The size of the drive in GiBs.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// A snapshot to base the EBS volume off of.
        /// </summary>
        [Input("snapshotId")]
        public Input<string>? SnapshotId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The type of EBS volume. Can be "standard", "gp2", "io1", "sc1" or "st1" (Default: "gp2").
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VolumeState()
        {
        }
    }
}
