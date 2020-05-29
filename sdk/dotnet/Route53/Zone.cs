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
    /// Manages a Route53 Hosted Zone.
    /// 
    /// ## Example Usage
    /// 
    /// ### Public Zone
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var primary = new Aws.Route53.Zone("primary", new Aws.Route53.ZoneArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Public Subdomain Zone
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var main = new Aws.Route53.Zone("main", new Aws.Route53.ZoneArgs
    ///         {
    ///         });
    ///         var dev = new Aws.Route53.Zone("dev", new Aws.Route53.ZoneArgs
    ///         {
    ///             Tags = 
    ///             {
    ///                 { "Environment", "dev" },
    ///             },
    ///         });
    ///         var dev_ns = new Aws.Route53.Record("dev-ns", new Aws.Route53.RecordArgs
    ///         {
    ///             Name = "dev.example.com",
    ///             Records = 
    ///             {
    ///                 dev.NameServers.Apply(nameServers =&gt; nameServers[0]),
    ///                 dev.NameServers.Apply(nameServers =&gt; nameServers[1]),
    ///                 dev.NameServers.Apply(nameServers =&gt; nameServers[2]),
    ///                 dev.NameServers.Apply(nameServers =&gt; nameServers[3]),
    ///             },
    ///             Ttl = 30,
    ///             Type = "NS",
    ///             ZoneId = main.ZoneId,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Private Zone
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @private = new Aws.Route53.Zone("private", new Aws.Route53.ZoneArgs
    ///         {
    ///             Vpcs = 
    ///             {
    ///                 new Aws.Route53.Inputs.ZoneVpcArgs
    ///                 {
    ///                     VpcId = aws_vpc.Example.Id,
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Zone : Pulumi.CustomResource
    {
        /// <summary>
        /// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        /// </summary>
        [Output("delegationSetId")]
        public Output<string?> DelegationSetId { get; private set; } = null!;

        /// <summary>
        /// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// This is the name of the hosted zone.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of name servers in associated (or default) delegation set.
        /// Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
        /// </summary>
        [Output("nameServers")]
        public Output<ImmutableArray<string>> NameServers { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the zone.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any [`aws.route53.ZoneAssociation` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone_association.html) specifying the same zone ID. Detailed below.
        /// </summary>
        [Output("vpcs")]
        public Output<ImmutableArray<Outputs.ZoneVpc>> Vpcs { get; private set; } = null!;

        /// <summary>
        /// The Hosted Zone ID. This can be referenced by zone records.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:route53/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("aws:route53/zone:Zone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        /// </summary>
        [Input("delegationSetId")]
        public Input<string>? DelegationSetId { get; set; }

        /// <summary>
        /// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// This is the name of the hosted zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the zone.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcs")]
        private InputList<Inputs.ZoneVpcArgs>? _vpcs;

        /// <summary>
        /// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any [`aws.route53.ZoneAssociation` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone_association.html) specifying the same zone ID. Detailed below.
        /// </summary>
        public InputList<Inputs.ZoneVpcArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ZoneVpcArgs>());
            set => _vpcs = value;
        }

        public ZoneArgs()
        {
            Comment = "Managed by Pulumi";
        }
    }

    public sealed class ZoneState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
        /// </summary>
        [Input("delegationSetId")]
        public Input<string>? DelegationSetId { get; set; }

        /// <summary>
        /// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// This is the name of the hosted zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nameServers")]
        private InputList<string>? _nameServers;

        /// <summary>
        /// A list of name servers in associated (or default) delegation set.
        /// Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
        /// </summary>
        public InputList<string> NameServers
        {
            get => _nameServers ?? (_nameServers = new InputList<string>());
            set => _nameServers = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the zone.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("vpcs")]
        private InputList<Inputs.ZoneVpcGetArgs>? _vpcs;

        /// <summary>
        /// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any [`aws.route53.ZoneAssociation` resource](https://www.terraform.io/docs/providers/aws/r/route53_zone_association.html) specifying the same zone ID. Detailed below.
        /// </summary>
        public InputList<Inputs.ZoneVpcGetArgs> Vpcs
        {
            get => _vpcs ?? (_vpcs = new InputList<Inputs.ZoneVpcGetArgs>());
            set => _vpcs = value;
        }

        /// <summary>
        /// The Hosted Zone ID. This can be referenced by zone records.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ZoneState()
        {
            Comment = "Managed by Pulumi";
        }
    }
}
