// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Xray
{
    public partial class SamplingRule : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("attributes")]
        public Output<ImmutableDictionary<string, string>?> Attributes { get; private set; } = null!;

        [Output("fixedRate")]
        public Output<double> FixedRate { get; private set; } = null!;

        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        [Output("httpMethod")]
        public Output<string> HttpMethod { get; private set; } = null!;

        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        [Output("reservoirSize")]
        public Output<int> ReservoirSize { get; private set; } = null!;

        [Output("resourceArn")]
        public Output<string> ResourceArn { get; private set; } = null!;

        [Output("ruleName")]
        public Output<string?> RuleName { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        [Output("serviceType")]
        public Output<string> ServiceType { get; private set; } = null!;

        [Output("urlPath")]
        public Output<string> UrlPath { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a SamplingRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SamplingRule(string name, SamplingRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:xray/samplingRule:SamplingRule", name, args ?? new SamplingRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SamplingRule(string name, Input<string> id, SamplingRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:xray/samplingRule:SamplingRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SamplingRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SamplingRule Get(string name, Input<string> id, SamplingRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SamplingRule(name, id, state, options);
        }
    }

    public sealed class SamplingRuleArgs : Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("fixedRate", required: true)]
        public Input<double> FixedRate { get; set; } = null!;

        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        [Input("httpMethod", required: true)]
        public Input<string> HttpMethod { get; set; } = null!;

        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("reservoirSize", required: true)]
        public Input<int> ReservoirSize { get; set; } = null!;

        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("serviceType", required: true)]
        public Input<string> ServiceType { get; set; } = null!;

        [Input("urlPath", required: true)]
        public Input<string> UrlPath { get; set; } = null!;

        [Input("version", required: true)]
        public Input<int> Version { get; set; } = null!;

        public SamplingRuleArgs()
        {
        }
    }

    public sealed class SamplingRuleState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("fixedRate")]
        public Input<double>? FixedRate { get; set; }

        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("reservoirSize")]
        public Input<int>? ReservoirSize { get; set; }

        [Input("resourceArn")]
        public Input<string>? ResourceArn { get; set; }

        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        [Input("urlPath")]
        public Input<string>? UrlPath { get; set; }

        [Input("version")]
        public Input<int>? Version { get; set; }

        public SamplingRuleState()
        {
        }
    }
}
