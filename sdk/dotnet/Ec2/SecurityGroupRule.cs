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
    /// Provides a security group rule resource. Represents a single `ingress` or
    /// `egress` group rule, which can be added to external Security Groups.
    /// 
    /// &gt; **NOTE on Security Groups and Security Group Rules:** This provider currently
    /// provides both a standalone Security Group Rule resource (a single `ingress` or
    /// `egress` rule), and a Security Group resource with `ingress` and `egress` rules
    /// defined in-line. At this time you cannot use a Security Group with in-line rules
    /// in conjunction with any Security Group Rule resources. Doing so will cause
    /// a conflict of rule settings and will overwrite rules.
    /// 
    /// &gt; **NOTE:** Setting `protocol = "all"` or `protocol = -1` with `from_port` and `to_port` will result in the EC2 API creating a security group rule with all ports open. This API behavior cannot be controlled by this provider and may generate warnings in the future.
    /// 
    /// &gt; **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).
    /// </summary>
    public partial class SecurityGroupRule : Pulumi.CustomResource
    {
        /// <summary>
        /// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Output("cidrBlocks")]
        public Output<ImmutableArray<string>> CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// Description of the rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The start port (or ICMP type number if protocol is "icmp" or "icmpv6").
        /// </summary>
        [Output("fromPort")]
        public Output<int> FromPort { get; private set; } = null!;

        /// <summary>
        /// List of IPv6 CIDR blocks.
        /// </summary>
        [Output("ipv6CidrBlocks")]
        public Output<ImmutableArray<string>> Ipv6CidrBlocks { get; private set; } = null!;

        /// <summary>
        /// List of prefix list IDs (for allowing access to VPC endpoints).
        /// Only valid with `egress`.
        /// </summary>
        [Output("prefixListIds")]
        public Output<ImmutableArray<string>> PrefixListIds { get; private set; } = null!;

        /// <summary>
        /// The protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The security group to apply this rule to.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// If true, the security group itself will be added as
        /// a source to this ingress rule. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Output("self")]
        public Output<bool?> Self { get; private set; } = null!;

        /// <summary>
        /// The security group id to allow access to/from,
        /// depending on the `type`. Cannot be specified with `cidr_blocks` and `self`.
        /// </summary>
        [Output("sourceSecurityGroupId")]
        public Output<string> SourceSecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The end port (or ICMP code if protocol is "icmp").
        /// </summary>
        [Output("toPort")]
        public Output<int> ToPort { get; private set; } = null!;

        /// <summary>
        /// The type of rule being created. Valid options are `ingress` (inbound)
        /// or `egress` (outbound).
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityGroupRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityGroupRule(string name, SecurityGroupRuleArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/securityGroupRule:SecurityGroupRule", name, args ?? new SecurityGroupRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityGroupRule(string name, Input<string> id, SecurityGroupRuleState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/securityGroupRule:SecurityGroupRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityGroupRule Get(string name, Input<string> id, SecurityGroupRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityGroupRule(name, id, state, options);
        }
    }

    public sealed class SecurityGroupRuleArgs : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// Description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The start port (or ICMP type number if protocol is "icmp" or "icmpv6").
        /// </summary>
        [Input("fromPort", required: true)]
        public Input<int> FromPort { get; set; } = null!;

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;

        /// <summary>
        /// List of IPv6 CIDR blocks.
        /// </summary>
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;

        /// <summary>
        /// List of prefix list IDs (for allowing access to VPC endpoints).
        /// Only valid with `egress`.
        /// </summary>
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        /// <summary>
        /// The protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The security group to apply this rule to.
        /// </summary>
        [Input("securityGroupId", required: true)]
        public Input<string> SecurityGroupId { get; set; } = null!;

        /// <summary>
        /// If true, the security group itself will be added as
        /// a source to this ingress rule. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Input("self")]
        public Input<bool>? Self { get; set; }

        /// <summary>
        /// The security group id to allow access to/from,
        /// depending on the `type`. Cannot be specified with `cidr_blocks` and `self`.
        /// </summary>
        [Input("sourceSecurityGroupId")]
        public Input<string>? SourceSecurityGroupId { get; set; }

        /// <summary>
        /// The end port (or ICMP code if protocol is "icmp").
        /// </summary>
        [Input("toPort", required: true)]
        public Input<int> ToPort { get; set; } = null!;

        /// <summary>
        /// The type of rule being created. Valid options are `ingress` (inbound)
        /// or `egress` (outbound).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SecurityGroupRuleArgs()
        {
        }
    }

    public sealed class SecurityGroupRuleState : Pulumi.ResourceArgs
    {
        [Input("cidrBlocks")]
        private InputList<string>? _cidrBlocks;

        /// <summary>
        /// List of CIDR blocks. Cannot be specified with `source_security_group_id`.
        /// </summary>
        public InputList<string> CidrBlocks
        {
            get => _cidrBlocks ?? (_cidrBlocks = new InputList<string>());
            set => _cidrBlocks = value;
        }

        /// <summary>
        /// Description of the rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The start port (or ICMP type number if protocol is "icmp" or "icmpv6").
        /// </summary>
        [Input("fromPort")]
        public Input<int>? FromPort { get; set; }

        [Input("ipv6CidrBlocks")]
        private InputList<string>? _ipv6CidrBlocks;

        /// <summary>
        /// List of IPv6 CIDR blocks.
        /// </summary>
        public InputList<string> Ipv6CidrBlocks
        {
            get => _ipv6CidrBlocks ?? (_ipv6CidrBlocks = new InputList<string>());
            set => _ipv6CidrBlocks = value;
        }

        [Input("prefixListIds")]
        private InputList<string>? _prefixListIds;

        /// <summary>
        /// List of prefix list IDs (for allowing access to VPC endpoints).
        /// Only valid with `egress`.
        /// </summary>
        public InputList<string> PrefixListIds
        {
            get => _prefixListIds ?? (_prefixListIds = new InputList<string>());
            set => _prefixListIds = value;
        }

        /// <summary>
        /// The protocol. If not icmp, icmpv6, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The security group to apply this rule to.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// If true, the security group itself will be added as
        /// a source to this ingress rule. Cannot be specified with `source_security_group_id`.
        /// </summary>
        [Input("self")]
        public Input<bool>? Self { get; set; }

        /// <summary>
        /// The security group id to allow access to/from,
        /// depending on the `type`. Cannot be specified with `cidr_blocks` and `self`.
        /// </summary>
        [Input("sourceSecurityGroupId")]
        public Input<string>? SourceSecurityGroupId { get; set; }

        /// <summary>
        /// The end port (or ICMP code if protocol is "icmp").
        /// </summary>
        [Input("toPort")]
        public Input<int>? ToPort { get; set; }

        /// <summary>
        /// The type of rule being created. Valid options are `ingress` (inbound)
        /// or `egress` (outbound).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public SecurityGroupRuleState()
        {
        }
    }
}
