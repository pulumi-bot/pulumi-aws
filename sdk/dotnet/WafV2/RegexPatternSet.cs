// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2
{
    /// <summary>
    /// Provides an AWS WAFv2 Regex Pattern Set Resource
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
    ///         var example = new Aws.WafV2.RegexPatternSet("example", new Aws.WafV2.RegexPatternSetArgs
    ///         {
    ///             Description = "Example regex pattern set",
    ///             RegularExpressions = 
    ///             {
    ///                 new Aws.WafV2.Inputs.RegexPatternSetRegularExpressionArgs
    ///                 {
    ///                     RegexString = "one",
    ///                 },
    ///                 new Aws.WafV2.Inputs.RegexPatternSetRegularExpressionArgs
    ///                 {
    ///                     RegexString = "two",
    ///                 },
    ///             },
    ///             Scope = "REGIONAL",
    ///             Tags = 
    ///             {
    ///                 { "Tag1", "Value1" },
    ///                 { "Tag2", "Value2" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class RegexPatternSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that identifies the cluster.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A friendly description of the regular expression pattern set.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("lockToken")]
        public Output<string> LockToken { get; private set; } = null!;

        /// <summary>
        /// A friendly name of the regular expression pattern set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        /// </summary>
        [Output("regularExpressions")]
        public Output<ImmutableArray<Outputs.RegexPatternSetRegularExpression>> RegularExpressions { get; private set; } = null!;

        /// <summary>
        /// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// An array of key:value pairs to associate with the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a RegexPatternSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegexPatternSet(string name, RegexPatternSetArgs args, CustomResourceOptions? options = null)
            : base("aws:wafv2/regexPatternSet:RegexPatternSet", name, args ?? new RegexPatternSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegexPatternSet(string name, Input<string> id, RegexPatternSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:wafv2/regexPatternSet:RegexPatternSet", name, state, MakeResourceOptions(options, id))
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
        /// A friendly description of the regular expression pattern set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A friendly name of the regular expression pattern set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regularExpressions")]
        private InputList<Inputs.RegexPatternSetRegularExpressionArgs>? _regularExpressions;

        /// <summary>
        /// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        /// </summary>
        public InputList<Inputs.RegexPatternSetRegularExpressionArgs> RegularExpressions
        {
            get => _regularExpressions ?? (_regularExpressions = new InputList<Inputs.RegexPatternSetRegularExpressionArgs>());
            set => _regularExpressions = value;
        }

        /// <summary>
        /// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// An array of key:value pairs to associate with the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RegexPatternSetArgs()
        {
        }
    }

    public sealed class RegexPatternSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) that identifies the cluster.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// A friendly description of the regular expression pattern set.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lockToken")]
        public Input<string>? LockToken { get; set; }

        /// <summary>
        /// A friendly name of the regular expression pattern set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regularExpressions")]
        private InputList<Inputs.RegexPatternSetRegularExpressionGetArgs>? _regularExpressions;

        /// <summary>
        /// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        /// </summary>
        public InputList<Inputs.RegexPatternSetRegularExpressionGetArgs> RegularExpressions
        {
            get => _regularExpressions ?? (_regularExpressions = new InputList<Inputs.RegexPatternSetRegularExpressionGetArgs>());
            set => _regularExpressions = value;
        }

        /// <summary>
        /// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// An array of key:value pairs to associate with the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public RegexPatternSetState()
        {
        }
    }
}
