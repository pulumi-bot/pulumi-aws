// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pinpoint
{
    /// <summary>
    /// Provides a Pinpoint GCM Channel resource.
    /// 
    /// &gt; **Note:** Api Key argument will be stored in the raw state as plain-text.
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
    ///         var app = new Aws.Pinpoint.App("app", new Aws.Pinpoint.AppArgs
    ///         {
    ///         });
    ///         var gcm = new Aws.Pinpoint.GcmChannel("gcm", new Aws.Pinpoint.GcmChannelArgs
    ///         {
    ///             ApplicationId = app.ApplicationId,
    ///             ApiKey = "api_key",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Pinpoint GCM Channel can be imported using the `application-id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:pinpoint/gcmChannel:GcmChannel gcm application-id
    /// ```
    /// </summary>
    [AwsResourceType("aws:pinpoint/gcmChannel:GcmChannel")]
    public partial class GcmChannel : Pulumi.CustomResource
    {
        /// <summary>
        /// Platform credential API key from Google.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The application ID.
        /// </summary>
        [Output("applicationId")]
        public Output<string> ApplicationId { get; private set; } = null!;

        /// <summary>
        /// Whether the channel is enabled or disabled. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;


        /// <summary>
        /// Create a GcmChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GcmChannel(string name, GcmChannelArgs args, CustomResourceOptions? options = null)
            : base("aws:pinpoint/gcmChannel:GcmChannel", name, args ?? new GcmChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GcmChannel(string name, Input<string> id, GcmChannelState? state = null, CustomResourceOptions? options = null)
            : base("aws:pinpoint/gcmChannel:GcmChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GcmChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GcmChannel Get(string name, Input<string> id, GcmChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new GcmChannel(name, id, state, options);
        }
    }

    public sealed class GcmChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Platform credential API key from Google.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// The application ID.
        /// </summary>
        [Input("applicationId", required: true)]
        public Input<string> ApplicationId { get; set; } = null!;

        /// <summary>
        /// Whether the channel is enabled or disabled. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public GcmChannelArgs()
        {
        }
    }

    public sealed class GcmChannelState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Platform credential API key from Google.
        /// </summary>
        [Input("apiKey")]
        public Input<string>? ApiKey { get; set; }

        /// <summary>
        /// The application ID.
        /// </summary>
        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        /// <summary>
        /// Whether the channel is enabled or disabled. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public GcmChannelState()
        {
        }
    }
}
