// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CodePipeline
{
    public partial class Webhook : Pulumi.CustomResource
    {
        [Output("authentication")]
        public Output<string> Authentication { get; private set; } = null!;

        [Output("authenticationConfiguration")]
        public Output<Outputs.WebhookAuthenticationConfiguration?> AuthenticationConfiguration { get; private set; } = null!;

        [Output("filters")]
        public Output<ImmutableArray<Outputs.WebhookFilter>> Filters { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("targetAction")]
        public Output<string> TargetAction { get; private set; } = null!;

        [Output("targetPipeline")]
        public Output<string> TargetPipeline { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a Webhook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Webhook(string name, WebhookArgs args, CustomResourceOptions? options = null)
            : base("aws:codepipeline/webhook:Webhook", name, args ?? new WebhookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Webhook(string name, Input<string> id, WebhookState? state = null, CustomResourceOptions? options = null)
            : base("aws:codepipeline/webhook:Webhook", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Webhook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Webhook Get(string name, Input<string> id, WebhookState? state = null, CustomResourceOptions? options = null)
        {
            return new Webhook(name, id, state, options);
        }
    }

    public sealed class WebhookArgs : Pulumi.ResourceArgs
    {
        [Input("authentication", required: true)]
        public Input<string> Authentication { get; set; } = null!;

        [Input("authenticationConfiguration")]
        public Input<Inputs.WebhookAuthenticationConfigurationArgs>? AuthenticationConfiguration { get; set; }

        [Input("filters", required: true)]
        private InputList<Inputs.WebhookFilterArgs>? _filters;
        public InputList<Inputs.WebhookFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.WebhookFilterArgs>());
            set => _filters = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targetAction", required: true)]
        public Input<string> TargetAction { get; set; } = null!;

        [Input("targetPipeline", required: true)]
        public Input<string> TargetPipeline { get; set; } = null!;

        public WebhookArgs()
        {
        }
    }

    public sealed class WebhookState : Pulumi.ResourceArgs
    {
        [Input("authentication")]
        public Input<string>? Authentication { get; set; }

        [Input("authenticationConfiguration")]
        public Input<Inputs.WebhookAuthenticationConfigurationGetArgs>? AuthenticationConfiguration { get; set; }

        [Input("filters")]
        private InputList<Inputs.WebhookFilterGetArgs>? _filters;
        public InputList<Inputs.WebhookFilterGetArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.WebhookFilterGetArgs>());
            set => _filters = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("targetAction")]
        public Input<string>? TargetAction { get; set; }

        [Input("targetPipeline")]
        public Input<string>? TargetPipeline { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        public WebhookState()
        {
        }
    }
}
