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
    /// Provides a WAF Regional Geo Match Set Resource
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
    ///         var geoMatchSet = new Aws.WafRegional.GeoMatchSet("geoMatchSet", new Aws.WafRegional.GeoMatchSetArgs
    ///         {
    ///             GeoMatchConstraints = 
    ///             {
    ///                 new Aws.WafRegional.Inputs.GeoMatchSetGeoMatchConstraintArgs
    ///                 {
    ///                     Type = "Country",
    ///                     Value = "US",
    ///                 },
    ///                 new Aws.WafRegional.Inputs.GeoMatchSetGeoMatchConstraintArgs
    ///                 {
    ///                     Type = "Country",
    ///                     Value = "CA",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// WAF Regional Geo Match Set can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:wafregional/geoMatchSet:GeoMatchSet geo_match_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
    /// ```
    /// </summary>
    [AwsResourceType("aws:wafregional/geoMatchSet:GeoMatchSet")]
    public partial class GeoMatchSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
        /// </summary>
        [Output("geoMatchConstraints")]
        public Output<ImmutableArray<Outputs.GeoMatchSetGeoMatchConstraint>> GeoMatchConstraints { get; private set; } = null!;

        /// <summary>
        /// The name or description of the Geo Match Set.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a GeoMatchSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GeoMatchSet(string name, GeoMatchSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:wafregional/geoMatchSet:GeoMatchSet", name, args ?? new GeoMatchSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GeoMatchSet(string name, Input<string> id, GeoMatchSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:wafregional/geoMatchSet:GeoMatchSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GeoMatchSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GeoMatchSet Get(string name, Input<string> id, GeoMatchSetState? state = null, CustomResourceOptions? options = null)
        {
            return new GeoMatchSet(name, id, state, options);
        }
    }

    public sealed class GeoMatchSetArgs : Pulumi.ResourceArgs
    {
        [Input("geoMatchConstraints")]
        private InputList<Inputs.GeoMatchSetGeoMatchConstraintArgs>? _geoMatchConstraints;

        /// <summary>
        /// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
        /// </summary>
        public InputList<Inputs.GeoMatchSetGeoMatchConstraintArgs> GeoMatchConstraints
        {
            get => _geoMatchConstraints ?? (_geoMatchConstraints = new InputList<Inputs.GeoMatchSetGeoMatchConstraintArgs>());
            set => _geoMatchConstraints = value;
        }

        /// <summary>
        /// The name or description of the Geo Match Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GeoMatchSetArgs()
        {
        }
    }

    public sealed class GeoMatchSetState : Pulumi.ResourceArgs
    {
        [Input("geoMatchConstraints")]
        private InputList<Inputs.GeoMatchSetGeoMatchConstraintGetArgs>? _geoMatchConstraints;

        /// <summary>
        /// The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
        /// </summary>
        public InputList<Inputs.GeoMatchSetGeoMatchConstraintGetArgs> GeoMatchConstraints
        {
            get => _geoMatchConstraints ?? (_geoMatchConstraints = new InputList<Inputs.GeoMatchSetGeoMatchConstraintGetArgs>());
            set => _geoMatchConstraints = value;
        }

        /// <summary>
        /// The name or description of the Geo Match Set.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GeoMatchSetState()
        {
        }
    }
}
