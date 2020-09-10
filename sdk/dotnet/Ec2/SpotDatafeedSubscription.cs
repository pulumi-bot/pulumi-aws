// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public partial class SpotDatafeedSubscription : Pulumi.CustomResource
    {
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("prefix")]
        public Output<string?> Prefix { get; private set; } = null!;


        /// <summary>
        /// Create a SpotDatafeedSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SpotDatafeedSubscription(string name, SpotDatafeedSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, args ?? new SpotDatafeedSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SpotDatafeedSubscription(string name, Input<string> id, SpotDatafeedSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SpotDatafeedSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SpotDatafeedSubscription Get(string name, Input<string> id, SpotDatafeedSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new SpotDatafeedSubscription(name, id, state, options);
        }
    }

    public sealed class SpotDatafeedSubscriptionArgs : Pulumi.ResourceArgs
    {
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public SpotDatafeedSubscriptionArgs()
        {
        }
    }

    public sealed class SpotDatafeedSubscriptionState : Pulumi.ResourceArgs
    {
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public SpotDatafeedSubscriptionState()
        {
        }
    }
}
