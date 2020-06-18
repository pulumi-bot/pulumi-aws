// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ses
{
    /// <summary>
    /// Provides an SES receipt filter resource
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
    ///         var filter = new Aws.Ses.ReceiptFilter("filter", new Aws.Ses.ReceiptFilterArgs
    ///         {
    ///             Cidr = "10.10.10.10",
    ///             Policy = "Block",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ReceiptFilter : Pulumi.CustomResource
    {
        /// <summary>
        /// The IP address or address range to filter, in CIDR notation
        /// </summary>
        [Output("cidr")]
        public Output<string> Cidr { get; private set; } = null!;

        /// <summary>
        /// The name of the filter
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Block or Allow
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a ReceiptFilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReceiptFilter(string name, ReceiptFilterArgs args, CustomResourceOptions? options = null)
            : base("aws:ses/receiptFilter:ReceiptFilter", name, args ?? new ReceiptFilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ReceiptFilter(string name, Input<string> id, ReceiptFilterState? state = null, CustomResourceOptions? options = null)
            : base("aws:ses/receiptFilter:ReceiptFilter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ReceiptFilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ReceiptFilter Get(string name, Input<string> id, ReceiptFilterState? state = null, CustomResourceOptions? options = null)
        {
            return new ReceiptFilter(name, id, state, options);
        }
    }

    public sealed class ReceiptFilterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address or address range to filter, in CIDR notation
        /// </summary>
        [Input("cidr", required: true)]
        public Input<string> Cidr { get; set; } = null!;

        /// <summary>
        /// The name of the filter
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Block or Allow
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public ReceiptFilterArgs()
        {
        }
    }

    public sealed class ReceiptFilterState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP address or address range to filter, in CIDR notation
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// The name of the filter
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Block or Allow
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public ReceiptFilterState()
        {
        }
    }
}
