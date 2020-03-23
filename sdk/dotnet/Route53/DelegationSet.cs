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
    /// Provides a [Route53 Delegation Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API-actions-by-function.html#actions-by-function-reusable-delegation-sets) resource.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/route53_delegation_set.html.markdown.
    /// </summary>
    public partial class DelegationSet : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of authoritative name servers for the hosted zone
        /// (effectively a list of NS records).
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// This is a reference name used in Caller Reference
        /// (helpful for identifying single delegation set amongst others)
        /// </summary>
        [Output("referenceName")]
        public Output<string?> ReferenceName { get; private set; } = null!;


        /// <summary>
        /// Create a DelegationSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DelegationSet(string name, DelegationSetArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:route53/delegationSet:DelegationSet", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private DelegationSet(string name, Input<string> id, DelegationSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/delegationSet:DelegationSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DelegationSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DelegationSet Get(string name, Input<string> id, DelegationSetState? state = null, CustomResourceOptions? options = null)
        {
            return new DelegationSet(name, id, state, options);
        }
    }

    public sealed class DelegationSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// This is a reference name used in Caller Reference
        /// (helpful for identifying single delegation set amongst others)
        /// </summary>
        [Input("referenceName")]
        public Input<string>? ReferenceName { get; set; }

        public DelegationSetArgs()
        {
        }
    }

    public sealed class DelegationSetState : Pulumi.ResourceArgs
    {
        [Input("nameServers")]
        private InputList<string>? _nameServers;

        /// <summary>
        /// A list of authoritative name servers for the hosted zone
        /// (effectively a list of NS records).
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        /// <summary>
        /// This is a reference name used in Caller Reference
        /// (helpful for identifying single delegation set amongst others)
        /// </summary>
        [Input("referenceName")]
        public Input<string>? ReferenceName { get; set; }

        public DelegationSetState()
        {
        }
    }
}
