// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.SecurityHub
{
    public partial class StandardsSubscription : Pulumi.CustomResource
    {
        [Output("standardsArn")]
        public Output<string> StandardsArn { get; private set; } = null!;


        /// <summary>
        /// Create a StandardsSubscription resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StandardsSubscription(string name, StandardsSubscriptionArgs args, CustomResourceOptions? options = null)
            : base("aws:securityhub/standardsSubscription:StandardsSubscription", name, args ?? new StandardsSubscriptionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StandardsSubscription(string name, Input<string> id, StandardsSubscriptionState? state = null, CustomResourceOptions? options = null)
            : base("aws:securityhub/standardsSubscription:StandardsSubscription", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StandardsSubscription resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StandardsSubscription Get(string name, Input<string> id, StandardsSubscriptionState? state = null, CustomResourceOptions? options = null)
        {
            return new StandardsSubscription(name, id, state, options);
        }
    }

    public sealed class StandardsSubscriptionArgs : Pulumi.ResourceArgs
    {
        [Input("standardsArn", required: true)]
        public Input<string> StandardsArn { get; set; } = null!;

        public StandardsSubscriptionArgs()
        {
        }
    }

    public sealed class StandardsSubscriptionState : Pulumi.ResourceArgs
    {
        [Input("standardsArn")]
        public Input<string>? StandardsArn { get; set; }

        public StandardsSubscriptionState()
        {
        }
    }
}
