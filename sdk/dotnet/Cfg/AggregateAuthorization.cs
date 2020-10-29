// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg
{
    /// <summary>
    /// Manages an AWS Config Aggregate Authorization
    /// </summary>
    public partial class AggregateAuthorization : Pulumi.CustomResource
    {
        /// <summary>
        /// Account ID
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the authorization
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Region
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a AggregateAuthorization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AggregateAuthorization(string name, AggregateAuthorizationArgs args, CustomResourceOptions? options = null)
            : base("aws:cfg/aggregateAuthorization:AggregateAuthorization", name, args ?? new AggregateAuthorizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AggregateAuthorization(string name, Input<string> id, AggregateAuthorizationState? state = null, CustomResourceOptions? options = null)
            : base("aws:cfg/aggregateAuthorization:AggregateAuthorization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AggregateAuthorization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AggregateAuthorization Get(string name, Input<string> id, AggregateAuthorizationState? state = null, CustomResourceOptions? options = null)
        {
            return new AggregateAuthorization(name, id, state, options);
        }
    }

    public sealed class AggregateAuthorizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account ID
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// Region
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

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

        public AggregateAuthorizationArgs()
        {
        }
    }

    public sealed class AggregateAuthorizationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Account ID
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The ARN of the authorization
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Region
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

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

        public AggregateAuthorizationState()
        {
        }
    }
}
