// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a security group resource.
 * 
 * > **NOTE on Security Groups and Security Group Rules:** Terraform currently
 * provides both a standalone Security Group Rule resource (a single `ingress` or
 * `egress` rule), and a Security Group resource with `ingress` and `egress` rules
 * defined in-line. At this time you cannot use a Security Group with in-line rules
 * in conjunction with any Security Group Rule resources. Doing so will cause
 * a conflict of rule settings and will overwrite rules.
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
 * const allowTls = new aws.ec2.SecurityGroup("allow_tls", {
 *     description: "Allow TLS inbound traffic",
 *     egress: [{
 *         cidrBlocks: ["0.0.0.0/0"],
 *         fromPort: 0,
 *         prefixListIds: ["pl-12c4e678"],
 *         protocol: "-1",
 *         toPort: 0,
 *     }],
 *     ingress: [{
 *         // Please restrict your ingress to only necessary IPs and ports.
 *         // Opening to 0.0.0.0/0 can lead to security vulnerabilities.
 *         cidrBlocks: "", // add a CIDR block here
 *         // TLS (change to whatever ports you need)
 *         fromPort: 443,
 *         protocol: "-1",
 *         toPort: 443,
 *     }],
 *     vpcId: aws_vpc_main.id,
 * });
 * ```
 * 
 * Basic usage with tags:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const allowTls = new aws.ec2.SecurityGroup("allow_tls", {
 *     description: "Allow TLS inbound traffic",
 *     ingress: [{
 *         // Please restrict your ingress to only necessary IPs and ports.
 *         // Opening to 0.0.0.0/0 can lead to security vulnerabilities.
 *         cidrBlocks: "", // add your IP address here
 *         // TLS (change to whatever ports you need)
 *         fromPort: 443,
 *         protocol: "tcp",
 *         toPort: 443,
 *     }],
 *     tags: {
 *         Name: "allow_all",
 *     },
 * });
 * ```
 * 
 * ## Usage with prefix list IDs
 * 
 * Prefix list IDs are managed by AWS internally. Prefix list IDs
 * are associated with a prefix list name, or service name, that is linked to a specific region.
 * Prefix list IDs are exported on VPC Endpoints, so you can use this format:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * // ...
 * const myEndpoint = new aws.ec2.VpcEndpoint("my_endpoint", {});
 * ```
 */
export class SecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing SecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecurityGroupState, opts?: pulumi.CustomResourceOptions): SecurityGroup {
        return new SecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/securityGroup:SecurityGroup';

    /**
     * Returns true if the given object is an instance of SecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecurityGroup.__pulumiType;
    }

    /**
     * The ARN of the security group
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The security group description. Defaults to
     * "Managed by Terraform". Cannot be "". __NOTE__: This field maps to the AWS
     * `GroupDescription` attribute, for which there is no Update API. If you'd like
     * to classify your security groups in a way that can be updated, use `tags`.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Can be specified multiple times for each
     * egress rule. Each egress block supports fields documented below.
     * This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
     */
    public readonly egress!: pulumi.Output<{ cidrBlocks?: string[], description?: string, fromPort: number, ipv6CidrBlocks?: string[], prefixListIds?: string[], protocol: string, securityGroups?: string[], self?: boolean, toPort: number }[]>;
    /**
     * Can be specified multiple times for each
     * ingress rule. Each ingress block supports fields documented below.
     * This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
     */
    public readonly ingress!: pulumi.Output<{ cidrBlocks?: string[], description?: string, fromPort: number, ipv6CidrBlocks?: string[], prefixListIds?: string[], protocol: string, securityGroups?: string[], self?: boolean, toPort: number }[]>;
    /**
     * The name of the security group. If omitted, Terraform will
     * assign a random, unique name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The owner ID.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * Instruct Terraform to revoke all of the
     * Security Groups attached ingress and egress rules before deleting the rule
     * itself. This is normally not needed, however certain AWS services such as
     * Elastic Map Reduce may automatically add required rules to security groups used
     * with the service, and those rules may contain a cyclic dependency that prevent
     * the security groups from being destroyed without removing the dependency first.
     * Default `false`
     */
    public readonly revokeRulesOnDelete!: pulumi.Output<boolean | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a SecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecurityGroupArgs | SecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecurityGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["egress"] = state ? state.egress : undefined;
            inputs["ingress"] = state ? state.ingress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["revokeRulesOnDelete"] = state ? state.revokeRulesOnDelete : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as SecurityGroupArgs | undefined;
            inputs["description"] = (args ? args.description : undefined) || "Managed by Pulumi";
            inputs["egress"] = args ? args.egress : undefined;
            inputs["ingress"] = args ? args.ingress : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["revokeRulesOnDelete"] = args ? args.revokeRulesOnDelete : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
        }
        super(SecurityGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecurityGroup resources.
 */
export interface SecurityGroupState {
    /**
     * The ARN of the security group
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The security group description. Defaults to
     * "Managed by Terraform". Cannot be "". __NOTE__: This field maps to the AWS
     * `GroupDescription` attribute, for which there is no Update API. If you'd like
     * to classify your security groups in a way that can be updated, use `tags`.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Can be specified multiple times for each
     * egress rule. Each egress block supports fields documented below.
     * This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
     */
    readonly egress?: pulumi.Input<pulumi.Input<{ cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, description?: pulumi.Input<string>, fromPort: pulumi.Input<number>, ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, prefixListIds?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string>, securityGroups?: pulumi.Input<pulumi.Input<string>[]>, self?: pulumi.Input<boolean>, toPort: pulumi.Input<number> }>[]>;
    /**
     * Can be specified multiple times for each
     * ingress rule. Each ingress block supports fields documented below.
     * This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
     */
    readonly ingress?: pulumi.Input<pulumi.Input<{ cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, description?: pulumi.Input<string>, fromPort: pulumi.Input<number>, ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, prefixListIds?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string>, securityGroups?: pulumi.Input<pulumi.Input<string>[]>, self?: pulumi.Input<boolean>, toPort: pulumi.Input<number> }>[]>;
    /**
     * The name of the security group. If omitted, Terraform will
     * assign a random, unique name
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The owner ID.
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * Instruct Terraform to revoke all of the
     * Security Groups attached ingress and egress rules before deleting the rule
     * itself. This is normally not needed, however certain AWS services such as
     * Elastic Map Reduce may automatically add required rules to security groups used
     * with the service, and those rules may contain a cyclic dependency that prevent
     * the security groups from being destroyed without removing the dependency first.
     * Default `false`
     */
    readonly revokeRulesOnDelete?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VPC ID.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecurityGroup resource.
 */
export interface SecurityGroupArgs {
    /**
     * The security group description. Defaults to
     * "Managed by Terraform". Cannot be "". __NOTE__: This field maps to the AWS
     * `GroupDescription` attribute, for which there is no Update API. If you'd like
     * to classify your security groups in a way that can be updated, use `tags`.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Can be specified multiple times for each
     * egress rule. Each egress block supports fields documented below.
     * This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
     */
    readonly egress?: pulumi.Input<pulumi.Input<{ cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, description?: pulumi.Input<string>, fromPort: pulumi.Input<number>, ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, prefixListIds?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string>, securityGroups?: pulumi.Input<pulumi.Input<string>[]>, self?: pulumi.Input<boolean>, toPort: pulumi.Input<number> }>[]>;
    /**
     * Can be specified multiple times for each
     * ingress rule. Each ingress block supports fields documented below.
     * This argument is processed in [attribute-as-blocks mode](https://www.terraform.io/docs/configuration/attr-as-blocks.html).
     */
    readonly ingress?: pulumi.Input<pulumi.Input<{ cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, description?: pulumi.Input<string>, fromPort: pulumi.Input<number>, ipv6CidrBlocks?: pulumi.Input<pulumi.Input<string>[]>, prefixListIds?: pulumi.Input<pulumi.Input<string>[]>, protocol: pulumi.Input<string>, securityGroups?: pulumi.Input<pulumi.Input<string>[]>, self?: pulumi.Input<boolean>, toPort: pulumi.Input<number> }>[]>;
    /**
     * The name of the security group. If omitted, Terraform will
     * assign a random, unique name
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified
     * prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * Instruct Terraform to revoke all of the
     * Security Groups attached ingress and egress rules before deleting the rule
     * itself. This is normally not needed, however certain AWS services such as
     * Elastic Map Reduce may automatically add required rules to security groups used
     * with the service, and those rules may contain a cyclic dependency that prevent
     * the security groups from being destroyed without removing the dependency first.
     * Default `false`
     */
    readonly revokeRulesOnDelete?: pulumi.Input<boolean>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VPC ID.
     */
    readonly vpcId?: pulumi.Input<string>;
}
