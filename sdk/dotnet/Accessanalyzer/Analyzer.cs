// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AccessAnalyzer
{
    /// <summary>
    /// Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/accessanalyzer_analyzer.html.markdown.
    /// </summary>
    public partial class Analyzer : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Analyzer.
        /// </summary>
        [Output("analyzerName")]
        public Output<string> AnalyzerName { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Analyzer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Analyzer(string name, AnalyzerArgs args, CustomResourceOptions? options = null)
            : base("aws:accessanalyzer/analyzer:Analyzer", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Analyzer(string name, Input<string> id, AnalyzerState? state = null, CustomResourceOptions? options = null)
            : base("aws:accessanalyzer/analyzer:Analyzer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Analyzer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Analyzer Get(string name, Input<string> id, AnalyzerState? state = null, CustomResourceOptions? options = null)
        {
            return new Analyzer(name, id, state, options);
        }
    }

    public sealed class AnalyzerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Analyzer.
        /// </summary>
        [Input("analyzerName", required: true)]
        public Input<string> AnalyzerName { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AnalyzerArgs()
        {
        }
    }

    public sealed class AnalyzerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Analyzer.
        /// </summary>
        [Input("analyzerName")]
        public Input<string>? AnalyzerName { get; set; }

        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value mapping of resource tags.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Type of Analyzer. Valid value is currently only `ACCOUNT`. Defaults to `ACCOUNT`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AnalyzerState()
        {
        }
    }
}
