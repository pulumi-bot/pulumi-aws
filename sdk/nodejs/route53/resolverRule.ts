// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Route53 Resolver rule.
 *
 * ## Example Usage
 * ### System rule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const sys = new aws.route53.ResolverRule("sys", {
 *     domainName: "subdomain.example.com",
 *     ruleType: "SYSTEM",
 * });
 * ```
 * ### Forward rule
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const fwd = new aws.route53.ResolverRule("fwd", {
 *     domainName: "example.com",
 *     ruleType: "FORWARD",
 *     resolverEndpointId: aws_route53_resolver_endpoint.foo.id,
 *     targetIps: [{
 *         ip: "123.45.67.89",
 *     }],
 *     tags: {
 *         Environment: "Prod",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Route53 Resolver rules can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:route53/resolverRule:ResolverRule sys rslvr-rr-0123456789abcdef0
 * ```
 */
export class ResolverRule extends pulumi.CustomResource {
    /**
     * Get an existing ResolverRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResolverRuleState, opts?: pulumi.CustomResourceOptions): ResolverRule {
        return new ResolverRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:route53/resolverRule:ResolverRule';

    /**
     * Returns true if the given object is an instance of ResolverRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResolverRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResolverRule.__pulumiType;
    }

    /**
     * The ARN (Amazon Resource Name) for the resolver rule.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
     * This argument should only be specified for `FORWARD` type rules.
     */
    public readonly resolverEndpointId!: pulumi.Output<string | undefined>;
    /**
     * The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
     */
    public readonly ruleType!: pulumi.Output<string>;
    /**
     * Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
     * Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
     */
    public /*out*/ readonly shareStatus!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
     * This argument should only be specified for `FORWARD` type rules.
     */
    public readonly targetIps!: pulumi.Output<outputs.route53.ResolverRuleTargetIp[] | undefined>;

    /**
     * Create a ResolverRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResolverRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResolverRuleArgs | ResolverRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResolverRuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["resolverEndpointId"] = state ? state.resolverEndpointId : undefined;
            inputs["ruleType"] = state ? state.ruleType : undefined;
            inputs["shareStatus"] = state ? state.shareStatus : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["targetIps"] = state ? state.targetIps : undefined;
        } else {
            const args = argsOrState as ResolverRuleArgs | undefined;
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.ruleType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleType'");
            }
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resolverEndpointId"] = args ? args.resolverEndpointId : undefined;
            inputs["ruleType"] = args ? args.ruleType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["targetIps"] = args ? args.targetIps : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["shareStatus"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ResolverRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ResolverRule resources.
 */
export interface ResolverRuleState {
    /**
     * The ARN (Amazon Resource Name) for the resolver rule.
     */
    arn?: pulumi.Input<string>;
    /**
     * DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
     */
    domainName?: pulumi.Input<string>;
    /**
     * A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
     */
    name?: pulumi.Input<string>;
    /**
     * When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
     * This argument should only be specified for `FORWARD` type rules.
     */
    resolverEndpointId?: pulumi.Input<string>;
    /**
     * The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
     */
    ruleType?: pulumi.Input<string>;
    /**
     * Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
     * Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
     */
    shareStatus?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
     * This argument should only be specified for `FORWARD` type rules.
     */
    targetIps?: pulumi.Input<pulumi.Input<inputs.route53.ResolverRuleTargetIp>[]>;
}

/**
 * The set of arguments for constructing a ResolverRule resource.
 */
export interface ResolverRuleArgs {
    /**
     * DNS queries for this domain name are forwarded to the IP addresses that are specified using `targetIp`.
     */
    domainName: pulumi.Input<string>;
    /**
     * A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `targetIp`.
     * This argument should only be specified for `FORWARD` type rules.
     */
    resolverEndpointId?: pulumi.Input<string>;
    /**
     * The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
     */
    ruleType: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
     * This argument should only be specified for `FORWARD` type rules.
     */
    targetIps?: pulumi.Input<pulumi.Input<inputs.route53.ResolverRuleTargetIp>[]>;
}
