// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Route53
{
    /// <summary>
    /// Provides a Route53 Resolver rule association.
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
    ///         var example = new Aws.Route53.ResolverRuleAssociation("example", new Aws.Route53.ResolverRuleAssociationArgs
    ///         {
    ///             ResolverRuleId = aws_route53_resolver_rule.Sys.Id,
    ///             VpcId = aws_vpc.Foo.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Route53 Resolver rule associations can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:route53/resolverRuleAssociation:ResolverRuleAssociation example rslvr-rrassoc-97242eaf88example
    /// ```
    /// </summary>
    [AwsResourceType("aws:route53/resolverRuleAssociation:ResolverRuleAssociation")]
    public partial class ResolverRuleAssociation : Pulumi.CustomResource
    {
        /// <summary>
        /// A name for the association that you're creating between a resolver rule and a VPC.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the resolver rule that you want to associate with the VPC.
        /// </summary>
        [Output("resolverRuleId")]
        public Output<string> ResolverRuleId { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC that you want to associate the resolver rule with.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a ResolverRuleAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResolverRuleAssociation(string name, ResolverRuleAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:route53/resolverRuleAssociation:ResolverRuleAssociation", name, args ?? new ResolverRuleAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResolverRuleAssociation(string name, Input<string> id, ResolverRuleAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/resolverRuleAssociation:ResolverRuleAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResolverRuleAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResolverRuleAssociation Get(string name, Input<string> id, ResolverRuleAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new ResolverRuleAssociation(name, id, state, options);
        }
    }

    public sealed class ResolverRuleAssociationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for the association that you're creating between a resolver rule and a VPC.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the resolver rule that you want to associate with the VPC.
        /// </summary>
        [Input("resolverRuleId", required: true)]
        public Input<string> ResolverRuleId { get; set; } = null!;

        /// <summary>
        /// The ID of the VPC that you want to associate the resolver rule with.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public ResolverRuleAssociationArgs()
        {
        }
    }

    public sealed class ResolverRuleAssociationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for the association that you're creating between a resolver rule and a VPC.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the resolver rule that you want to associate with the VPC.
        /// </summary>
        [Input("resolverRuleId")]
        public Input<string>? ResolverRuleId { get; set; }

        /// <summary>
        /// The ID of the VPC that you want to associate the resolver rule with.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public ResolverRuleAssociationState()
        {
        }
    }
}
