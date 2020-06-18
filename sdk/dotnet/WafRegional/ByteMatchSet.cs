// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafRegional
{
    /// <summary>
    /// Provides a WAF Regional Byte Match Set Resource for use with Application Load Balancer.
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
    ///         var byteSet = new Aws.WafRegional.ByteMatchSet("byteSet", new Aws.WafRegional.ByteMatchSetArgs
    ///         {
    ///             ByteMatchTuples = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.ByteMatchSetByteMatchTupleArgs
    ///                 {
    ///                     FieldToMatch = new Aws.WafRegional.Inputs.ByteMatchSetByteMatchTupleFieldToMatchArgs
    ///                     {
    ///                         Data = "referer",
    ///                         Type = "HEADER",
    ///                     },
    ///                     PositionalConstraint = "CONTAINS",
    ///                     TargetString = "badrefer1",
    ///                     TextTransformation = "NONE",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ByteMatchSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        /// </summary>
        [Output("byteMatchTuples")]
        public Output<ImmutableArray<Outputs.ByteMatchSetByteMatchTuple>> ByteMatchTuples { get; private set; } = null!;

        /// <summary>
        /// The name or description of the ByteMatchSet.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ByteMatchSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ByteMatchSet(string name, ByteMatchSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:wafregional/byteMatchSet:ByteMatchSet", name, args ?? new ByteMatchSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ByteMatchSet(string name, Input<string> id, ByteMatchSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:wafregional/byteMatchSet:ByteMatchSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ByteMatchSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ByteMatchSet Get(string name, Input<string> id, ByteMatchSetState? state = null, CustomResourceOptions? options = null)
        {
            return new ByteMatchSet(name, id, state, options);
        }
    }

    public sealed class ByteMatchSetArgs : Pulumi.ResourceArgs
    {
        [Input("byteMatchTuples")]
        private InputList<Inputs.ByteMatchSetByteMatchTupleArgs>? _byteMatchTuples;

        /// <summary>
        /// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        /// </summary>
        public InputList<Inputs.ByteMatchSetByteMatchTupleArgs> ByteMatchTuples
        {
            get => _byteMatchTuples ?? (_byteMatchTuples = new InputList<Inputs.ByteMatchSetByteMatchTupleArgs>());
            set => _byteMatchTuples = value;
        }

        /// <summary>
        /// The name or description of the ByteMatchSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ByteMatchSetArgs()
        {
        }
    }

    public sealed class ByteMatchSetState : Pulumi.ResourceArgs
    {
        [Input("byteMatchTuples")]
        private InputList<Inputs.ByteMatchSetByteMatchTupleGetArgs>? _byteMatchTuples;

        /// <summary>
        /// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        /// </summary>
        public InputList<Inputs.ByteMatchSetByteMatchTupleGetArgs> ByteMatchTuples
        {
            get => _byteMatchTuples ?? (_byteMatchTuples = new InputList<Inputs.ByteMatchSetByteMatchTupleGetArgs>());
            set => _byteMatchTuples = value;
        }

        /// <summary>
        /// The name or description of the ByteMatchSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ByteMatchSetState()
        {
        }
    }
}
