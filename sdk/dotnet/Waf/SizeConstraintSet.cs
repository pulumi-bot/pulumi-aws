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
    /// Provides a WAF Size Constraint Set Resource
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var sizeConstraintSet = new Aws.Waf.SizeConstraintSet("sizeConstraintSet", new Aws.Waf.SizeConstraintSetArgs
    ///         {
    ///             SizeConstraints = 
    ///             {
    ///                 new Aws.Waf.Inputs.SizeConstraintSetSizeConstraintArgs
    ///                 {
    ///                     ComparisonOperator = "EQ",
    ///                     FieldToMatch = new Aws.Waf.Inputs.SizeConstraintSetSizeConstraintFieldToMatchArgs
    ///                     {
    ///                         Type = "BODY",
    ///                     },
    ///                     Size = 4096,
    ///                     TextTransformation = "NONE",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class SizeConstraintSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name or description of the Size Constraint Set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies the parts of web requests that you want to inspect the size of.
        /// </summary>
        [Output("sizeConstraints")]
        public Output<ImmutableArray<Outputs.SizeConstraintSetSizeConstraint>> SizeConstraints { get; private set; } = null!;


        /// <summary>
        /// Create a SizeConstraintSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SizeConstraintSet(string name, SizeConstraintSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:waf/sizeConstraintSet:SizeConstraintSet", name, args ?? new SizeConstraintSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SizeConstraintSet(string name, Input<string> id, SizeConstraintSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:waf/sizeConstraintSet:SizeConstraintSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SizeConstraintSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SizeConstraintSet Get(string name, Input<string> id, SizeConstraintSetState? state = null, CustomResourceOptions? options = null)
        {
            return new SizeConstraintSet(name, id, state, options);
        }
    }

    public sealed class SizeConstraintSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name or description of the Size Constraint Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("sizeConstraints")]
        private InputList<Inputs.SizeConstraintSetSizeConstraintArgs>? _sizeConstraints;

        /// <summary>
        /// Specifies the parts of web requests that you want to inspect the size of.
        /// </summary>
        public InputList<Inputs.SizeConstraintSetSizeConstraintArgs> SizeConstraints
        {
            get => _sizeConstraints ?? (_sizeConstraints = new InputList<Inputs.SizeConstraintSetSizeConstraintArgs>());
            set => _sizeConstraints = value;
        }

        public SizeConstraintSetArgs()
        {
        }
    }

    public sealed class SizeConstraintSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name or description of the Size Constraint Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("sizeConstraints")]
        private InputList<Inputs.SizeConstraintSetSizeConstraintGetArgs>? _sizeConstraints;

        /// <summary>
        /// Specifies the parts of web requests that you want to inspect the size of.
        /// </summary>
        public InputList<Inputs.SizeConstraintSetSizeConstraintGetArgs> SizeConstraints
        {
            get => _sizeConstraints ?? (_sizeConstraints = new InputList<Inputs.SizeConstraintSetSizeConstraintGetArgs>());
            set => _sizeConstraints = value;
        }

        public SizeConstraintSetState()
        {
        }
    }
}
