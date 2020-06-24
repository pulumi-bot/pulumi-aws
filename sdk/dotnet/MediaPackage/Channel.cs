// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.MediaPackage
{
    /// <summary>
    /// Provides an AWS Elemental MediaPackage Channel.
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
    ///         var kittens = new Aws.MediaPackage.Channel("kittens", new Aws.MediaPackage.ChannelArgs
    ///         {
    ///             ChannelId = "kitten-channel",
    ///             Description = "A channel dedicated to amusing videos of kittens.",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Channel : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the channel
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A unique identifier describing the channel
        /// </summary>
        [Output("channelId")]
        public Output<string> ChannelId { get; private set; } = null!;

        /// <summary>
        /// A description of the channel
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// A single item list of HLS ingest information
        /// </summary>
        [Output("hlsIngests")]
        public Output<ImmutableArray<Outputs.ChannelHlsIngest>> HlsIngests { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Channel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Channel(string name, ChannelArgs args, CustomResourceOptions? options = null)
            : base("aws:mediapackage/channel:Channel", name, args ?? new ChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Channel(string name, Input<string> id, ChannelState? state = null, CustomResourceOptions? options = null)
            : base("aws:mediapackage/channel:Channel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Channel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Channel Get(string name, Input<string> id, ChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new Channel(name, id, state, options);
        }
    }

    public sealed class ChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier describing the channel
        /// </summary>
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        /// <summary>
        /// A description of the channel
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ChannelArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class ChannelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the channel
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A unique identifier describing the channel
        /// </summary>
        [Input("channelId")]
        public Input<string>? ChannelId { get; set; }

        /// <summary>
        /// A description of the channel
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hlsIngests")]
        private InputList<Inputs.ChannelHlsIngestGetArgs>? _hlsIngests;

        /// <summary>
        /// A single item list of HLS ingest information
        /// </summary>
        public InputList<Inputs.ChannelHlsIngestGetArgs> HlsIngests
        {
            get => _hlsIngests ?? (_hlsIngests = new InputList<Inputs.ChannelHlsIngestGetArgs>());
            set => _hlsIngests = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ChannelState()
        {
            Description = "Managed by Pulumi";
        }
    }
}
