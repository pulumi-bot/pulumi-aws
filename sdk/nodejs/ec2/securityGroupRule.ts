// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a security group rule resource. Represents a single `ingress` or
 * `egress` group rule, which can be added to external Security Groups.
 * 
 * > **NOTE on Security Groups and Security Group Rules:** Terraform currently
 * provides both a standalone Security Group Rule resource (a single `ingress` or
 * `egress` rule), and a Security Group resource with `ingress` and `egress` rules
 * defined in-line. At this time you cannot use a Security Group with in-line rules
 * in conjunction with any Security Group Rule resources. Doing so will cause
 * a conflict of rule settings and will overwrite rules.
 * 
 * > **NOTE:** Setting `protocol = "all"` or `protocol = -1` with `from_port` and `to_port` will result in the EC2 API creating a security group rule with all ports open. This API behavior cannot be controlled by Terraform and may generate warnings in the future.
 * 
 * > **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).
 * 
 * ## Example Usage
 * 
 * Basic usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_security_group_rule_allow_all = new aws.ec2.SecurityGroupRule("allow_all", {
 *     cidrBlocks: ["0.0.0.0/0"],
 *     fromPort: 0,
 *     prefixListIds: ["pl-12c4e678"],
 *     protocol: "tcp",
 *     securityGroupId: "sg-123456",
 *     toPort: 65535,
 *     type: "ingress",
 * });
 * ```
 * 
 * ## Usage with prefix list IDs
 * 
 * Prefix list IDs are manged by AWS internally. Prefix list IDs
 * are associated with a prefix list name, or service name, that is linked to a specific region.
 * Prefix list IDs are exported on VPC Endpoints, so you can use this format:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_vpc_endpoint_my_endpoint = new aws.ec2.VpcEndpoint("my_endpoint", {});
 * const aws_security_group_rule_allow_all = new aws.ec2.SecurityGroupRule("allow_all", {
 *     fromPort: 0,
 *     prefixListIds: [aws_vpc_endpoint_my_endpoint.prefixListId],
 *     protocol: "-1",
 *     securityGroupId: "sg-123456",
 *     toPort: 0,
 *     type: "egress",
 * });
 * ```
 */
export class SecurityGroupRule extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroupRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions): SecurityGroupRule {
        return new SecurityGroupRule(name, <any>state, { ...opts, id: id });
    }

    /**
     * List of CIDR blocks. Cannot be specified with `source_security_group_id`.
     */
    public readonly cidrBlocks: pulumi.Output<string[] | undefined>;
    /**
     * Description of the rule.
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * The start port (or ICMP type number if protocol is "icmp").
     */
    public readonly fromPort: pulumi.Output<number>;
    /**
     * List of IPv6 CIDR blocks.
     */
    public readonly ipv6CidrBlocks: pulumi.Output<string[] | undefined>;
    /**
     * List of prefix list IDs (for allowing access to VPC endpoints).
     * Only valid with `egress`.
     */
    public readonly prefixListIds: pulumi.Output<string[] | undefined>;
    /**
     * The protocol. If not icmp, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
     */
    public readonly protocol: pulumi.Output<string>;
    /**
     * The security group to apply this rule to.
     */
    public readonly securityGroupId: pulumi.Output<string>;
    /**
     * If true, the security group itself will be added as
     * a source to this ingress rule.
     */
    public readonly self: pulumi.Output<boolean | undefined>;
    /**
     * The security group id to allow access to/from,
     * depending on the `type`. Cannot be specified with `cidr_blocks`.
     */
    public readonly sourceSecurityGroupId: pulumi.Output<string>;
    /**
     * The end port (or ICMP code if protocol is "icmp").
     */
    public readonly toPort: pulumi.Output<number>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound)
     * or `egress` (outbound).
     */
    public readonly type: pulumi.Output<string>;

    /**
     * Create a SecurityGroupRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecurityGroupRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupRuleArgs | SecurityGroupRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: SecurityGroupRuleState = argsOrState as SecurityGroupRuleState | undefined;
            inputs["cidrBlocks"] = state ? state.cidrBlocks : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["fromPort"] = state ? state.fromPort : undefined;
            inputs["ipv6CidrBlocks"] = state ? state.ipv6CidrBlocks : undefined;
            inputs["prefixListIds"] = state ? state.prefixListIds : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["self"] = state ? state.self : undefined;
            inputs["sourceSecurityGroupId"] = state ? state.sourceSecurityGroupId : undefined;
            inputs["toPort"] = state ? state.toPort : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SecurityGroupRuleArgs | undefined;
            if (!args || args.fromPort === undefined) {
                throw new Error("Missing required property 'fromPort'");
            }
            if (!args || args.protocol === undefined) {
                throw new Error("Missing required property 'protocol'");
            }
            if (!args || args.securityGroupId === undefined) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if (!args || args.toPort === undefined) {
                throw new Error("Missing required property 'toPort'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["cidrBlocks"] = args ? args.cidrBlocks : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["fromPort"] = args ? args.fromPort : undefined;
            inputs["ipv6CidrBlocks"] = args ? args.ipv6CidrBlocks : undefined;
            inputs["prefixListIds"] = args ? args.prefixListIds : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["self"] = args ? args.self : undefined;
            inputs["sourceSecurityGroupId"] = args ? args.sourceSecurityGroupId : undefined;
            inputs["toPort"] = args ? args.toPort : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        super("aws:ec2/securityGroupRule:SecurityGroupRule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroupRule resources.
 */
export interface SecurityGroupRuleState {
    /**
     * List of CIDR blocks. Cannot be specified with `source_security_group_id`.
     */
    readonly cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The start port (or ICMP type number if protocol is "icmp").
     */
    readonly fromPort?: pulumi.Input<number>;
    /**
     * List of IPv6 CIDR blocks.
     */
    readonly ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of prefix list IDs (for allowing access to VPC endpoints).
     * Only valid with `egress`.
     */
    readonly prefixListIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The protocol. If not icmp, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * The security group to apply this rule to.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * If true, the security group itself will be added as
     * a source to this ingress rule.
     */
    readonly self?: pulumi.Input<boolean>;
    /**
     * The security group id to allow access to/from,
     * depending on the `type`. Cannot be specified with `cidr_blocks`.
     */
    readonly sourceSecurityGroupId?: pulumi.Input<string>;
    /**
     * The end port (or ICMP code if protocol is "icmp").
     */
    readonly toPort?: pulumi.Input<number>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound)
     * or `egress` (outbound).
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroupRule resource.
 */
export interface SecurityGroupRuleArgs {
    /**
     * List of CIDR blocks. Cannot be specified with `source_security_group_id`.
     */
    readonly cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the rule.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The start port (or ICMP type number if protocol is "icmp").
     */
    readonly fromPort: pulumi.Input<number>;
    /**
     * List of IPv6 CIDR blocks.
     */
    readonly ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of prefix list IDs (for allowing access to VPC endpoints).
     * Only valid with `egress`.
     */
    readonly prefixListIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The protocol. If not icmp, tcp, udp, or all use the [protocol number](https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * The security group to apply this rule to.
     */
    readonly securityGroupId: pulumi.Input<string>;
    /**
     * If true, the security group itself will be added as
     * a source to this ingress rule.
     */
    readonly self?: pulumi.Input<boolean>;
    /**
     * The security group id to allow access to/from,
     * depending on the `type`. Cannot be specified with `cidr_blocks`.
     */
    readonly sourceSecurityGroupId?: pulumi.Input<string>;
    /**
     * The end port (or ICMP code if protocol is "icmp").
     */
    readonly toPort: pulumi.Input<number>;
    /**
     * The type of rule being created. Valid options are `ingress` (inbound)
     * or `egress` (outbound).
     */
    readonly type: pulumi.Input<string>;
}
