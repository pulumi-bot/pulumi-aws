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
    /// Provides a Pinpoint SMS Channel resource.
    /// </summary>
    public partial class SmsChannel : Pulumi.CustomResource
    {
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
        /// Promotional messages per second that can be sent.
        /// </summary>
        [Output("promotionalMessagesPerSecond")]
        public Output<int> PromotionalMessagesPerSecond { get; private set; } = null!;

        /// <summary>
        /// Sender identifier of your messages.
        /// </summary>
        [Output("senderId")]
        public Output<string?> SenderId { get; private set; } = null!;

        /// <summary>
        /// The Short Code registered with the phone provider.
        /// </summary>
        [Output("shortCode")]
        public Output<string?> ShortCode { get; private set; } = null!;

        /// <summary>
        /// Transactional messages per second that can be sent.
        /// </summary>
        [Output("transactionalMessagesPerSecond")]
        public Output<int> TransactionalMessagesPerSecond { get; private set; } = null!;


        /// <summary>
        /// Create a SmsChannel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmsChannel(string name, SmsChannelArgs args, CustomResourceOptions? options = null)
            : base("aws:pinpoint/smsChannel:SmsChannel", name, args ?? new SmsChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmsChannel(string name, Input<string> id, SmsChannelState? state = null, CustomResourceOptions? options = null)
            : base("aws:pinpoint/smsChannel:SmsChannel", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SmsChannel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmsChannel Get(string name, Input<string> id, SmsChannelState? state = null, CustomResourceOptions? options = null)
        {
            return new SmsChannel(name, id, state, options);
        }
    }

    public sealed class SmsChannelArgs : Pulumi.ResourceArgs
    {
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

        /// <summary>
        /// Sender identifier of your messages.
        /// </summary>
        [Input("senderId")]
        public Input<string>? SenderId { get; set; }

        /// <summary>
        /// The Short Code registered with the phone provider.
        /// </summary>
        [Input("shortCode")]
        public Input<string>? ShortCode { get; set; }

        public SmsChannelArgs()
        {
        }
    }

    public sealed class SmsChannelState : Pulumi.ResourceArgs
    {
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

        /// <summary>
        /// Promotional messages per second that can be sent.
        /// </summary>
        [Input("promotionalMessagesPerSecond")]
        public Input<int>? PromotionalMessagesPerSecond { get; set; }

        /// <summary>
        /// Sender identifier of your messages.
        /// </summary>
        [Input("senderId")]
        public Input<string>? SenderId { get; set; }

        /// <summary>
        /// The Short Code registered with the phone provider.
        /// </summary>
        [Input("shortCode")]
        public Input<string>? ShortCode { get; set; }

        /// <summary>
        /// Transactional messages per second that can be sent.
        /// </summary>
        [Input("transactionalMessagesPerSecond")]
        public Input<int>? TransactionalMessagesPerSecond { get; set; }

        public SmsChannelState()
        {
        }
    }
}
