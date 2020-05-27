// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// Creates an entry (a rule) in a network ACL with the specified rule number.
    /// 
    /// &gt; **NOTE on Network ACLs and Network ACL Rules:** This provider currently
    /// provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
    /// defined in-line. At this time you cannot use a Network ACL with in-line rules
    /// in conjunction with any Network ACL Rule resources. Doing so will cause
    /// a conflict of rule settings and will overwrite rules.
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
    ///         var barNetworkAcl = new Aws.Ec2.NetworkAcl("barNetworkAcl", new Aws.Ec2.NetworkAclArgs
    ///         {
    ///             VpcId = aws_vpc.Foo.Id,
    ///         });
    ///         var barNetworkAclRule = new Aws.Ec2.NetworkAclRule("barNetworkAclRule", new Aws.Ec2.NetworkAclRuleArgs
    ///         {
    ///             NetworkAclId = barNetworkAcl.Id,
    ///             RuleNumber = 200,
    ///             Egress = false,
    ///             Protocol = "tcp",
    ///             RuleAction = "allow",
    ///             CidrBlock = aws_vpc.Foo.Cidr_block,
    ///             FromPort = 22,
    ///             ToPort = 22,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class NetworkAclRule : Pulumi.CustomResource
    {
        /// <summary>
        /// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        /// </summary>
        [Output("cidrBlock")]
        public Output<string?> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        /// </summary>
        [Output("egress")]
        public Output<bool?> Egress { get; private set; } = null!;

        /// <summary>
        /// The from port to match.
        /// </summary>
        [Output("fromPort")]
        public Output<int?> FromPort { get; private set; } = null!;

        /// <summary>
        /// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        /// </summary>
        [Output("icmpCode")]
        public Output<string?> IcmpCode { get; private set; } = null!;

        /// <summary>
        /// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        /// </summary>
        [Output("icmpType")]
        public Output<string?> IcmpType { get; private set; } = null!;

        /// <summary>
        /// The IPv6 CIDR block to allow or deny.
        /// </summary>
        [Output("ipv6CidrBlock")]
        public Output<string?> Ipv6CidrBlock { get; private set; } = null!;

        /// <summary>
        /// The ID of the network ACL.
        /// </summary>
        [Output("networkAclId")]
        public Output<string> NetworkAclId { get; private set; } = null!;

        /// <summary>
        /// The protocol. A value of -1 means all protocols.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        /// </summary>
        [Output("ruleAction")]
        public Output<string> RuleAction { get; private set; } = null!;

        /// <summary>
        /// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        /// </summary>
        [Output("ruleNumber")]
        public Output<int> RuleNumber { get; private set; } = null!;

        /// <summary>
        /// The to port to match.
        /// </summary>
        [Output("toPort")]
        public Output<int?> ToPort { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkAclRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkAclRule(string name, NetworkAclRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/networkAclRule:NetworkAclRule", name, args ?? new NetworkAclRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkAclRule(string name, Input<string> id, NetworkAclRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/networkAclRule:NetworkAclRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkAclRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkAclRule Get(string name, Input<string> id, NetworkAclRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkAclRule(name, id, state, options);
        }
    }

    public sealed class NetworkAclRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        /// </summary>
        [Input("egress")]
        public Input<bool>? Egress { get; set; }

        /// <summary>
        /// The from port to match.
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        /// <summary>
        /// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        /// </summary>
        [Input("icmpCode")]
        public Input<string>? IcmpCode { get; set; }

        /// <summary>
        /// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        /// </summary>
        [Input("icmpType")]
        public Input<string>? IcmpType { get; set; }

        /// <summary>
        /// The IPv6 CIDR block to allow or deny.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The ID of the network ACL.
        /// </summary>
        [Input("networkAclId", required: true)]
        public Input<string> NetworkAclId { get; set; } = null!;

        /// <summary>
        /// The protocol. A value of -1 means all protocols.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        /// </summary>
        [Input("ruleAction", required: true)]
        public Input<string> RuleAction { get; set; } = null!;

        /// <summary>
        /// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        /// </summary>
        [Input("ruleNumber", required: true)]
        public Input<int> RuleNumber { get; set; } = null!;

        /// <summary>
        /// The to port to match.
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        public NetworkAclRuleArgs()
        {
        }
    }

    public sealed class NetworkAclRuleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The network range to allow or deny, in CIDR notation (for example 172.16.0.0/24 ).
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet). Default `false`.
        /// </summary>
        [Input("egress")]
        public Input<bool>? Egress { get; set; }

        /// <summary>
        /// The from port to match.
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        /// <summary>
        /// ICMP protocol: The ICMP code. Required if specifying ICMP for the protocol. e.g. -1
        /// </summary>
        [Input("icmpCode")]
        public Input<string>? IcmpCode { get; set; }

        /// <summary>
        /// ICMP protocol: The ICMP type. Required if specifying ICMP for the protocol. e.g. -1
        /// </summary>
        [Input("icmpType")]
        public Input<string>? IcmpType { get; set; }

        /// <summary>
        /// The IPv6 CIDR block to allow or deny.
        /// </summary>
        [Input("ipv6CidrBlock")]
        public Input<string>? Ipv6CidrBlock { get; set; }

        /// <summary>
        /// The ID of the network ACL.
        /// </summary>
        [Input("networkAclId")]
        public Input<string>? NetworkAclId { get; set; }

        /// <summary>
        /// The protocol. A value of -1 means all protocols.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Indicates whether to allow or deny the traffic that matches the rule. Accepted values: `allow` | `deny`
        /// </summary>
        [Input("ruleAction")]
        public Input<string>? RuleAction { get; set; }

        /// <summary>
        /// The rule number for the entry (for example, 100). ACL entries are processed in ascending order by rule number.
        /// </summary>
        [Input("ruleNumber")]
        public Input<int>? RuleNumber { get; set; }

        /// <summary>
        /// The to port to match.
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        public NetworkAclRuleState()
        {
        }
    }
}
