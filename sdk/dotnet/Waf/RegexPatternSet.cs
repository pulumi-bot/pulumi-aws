// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Waf
{
    /// <summary>
    /// Provides a WAF Regex Pattern Set Resource
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_regex_pattern_set.html.markdown.
    /// </summary>
    public partial class RegexPatternSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name or description of the Regex Pattern Set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
        /// </summary>
        [Output("regexPatternStrings")]
        public Output<ImmutableArray<string>> RegexPatternStrings { get; private set; } = null!;


        /// <summary>
        /// Create a RegexPatternSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegexPatternSet(string name, RegexPatternSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:waf/regexPatternSet:RegexPatternSet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private RegexPatternSet(string name, Input<string> id, RegexPatternSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:waf/regexPatternSet:RegexPatternSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RegexPatternSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegexPatternSet Get(string name, Input<string> id, RegexPatternSetState? state = null, CustomResourceOptions? options = null)
        {
            return new RegexPatternSet(name, id, state, options);
        }
    }

    public sealed class RegexPatternSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or description of the Regex Pattern Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regexPatternStrings")]
        private InputList<string>? _regexPatternStrings;

        /// <summary>
        /// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
        /// </summary>
        public InputList<string> RegexPatternStrings
        {
            get => _regexPatternStrings ?? (_regexPatternStrings = new InputList<string>());
            set => _regexPatternStrings = value;
        }

        public RegexPatternSetArgs()
        {
        }
    }

    public sealed class RegexPatternSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name or description of the Regex Pattern Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regexPatternStrings")]
        private InputList<string>? _regexPatternStrings;

        /// <summary>
        /// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
        /// </summary>
        public InputList<string> RegexPatternStrings
        {
            get => _regexPatternStrings ?? (_regexPatternStrings = new InputList<string>());
            set => _regexPatternStrings = value;
        }

        public RegexPatternSetState()
        {
        }
    }
}
