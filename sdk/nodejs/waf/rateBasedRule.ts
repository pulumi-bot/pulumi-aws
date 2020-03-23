// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a WAF Rate Based Rule Resource
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ipset = new aws.waf.IpSet("ipset", {
 *     ipSetDescriptors: [{
 *         type: "IPV4",
 *         value: "192.0.7.0/24",
 *     }],
 * });
 * const wafrule = new aws.waf.RateBasedRule("wafrule", {
 *     metricName: "tfWAFRule",
 *     predicates: [{
 *         dataId: ipset.id,
 *         negated: false,
 *         type: "IPMatch",
 *     }],
 *     rateKey: "IP",
 *     rateLimit: 100,
 * }, {dependsOn: [ipset]});
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_rate_based_rule.html.markdown.
 */
export class RateBasedRule extends pulumi.CustomResource {
    /**
     * Get an existing RateBasedRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RateBasedRuleState, opts?: pulumi.CustomResourceOptions): RateBasedRule {
        return new RateBasedRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:waf/rateBasedRule:RateBasedRule';

    /**
     * Returns true if the given object is an instance of RateBasedRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RateBasedRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RateBasedRule.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name or description for the Amazon CloudWatch metric of this rule.
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * The name or description of the rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    public readonly predicates!: pulumi.Output<outputs.waf.RateBasedRulePredicate[] | undefined>;
    /**
     * Valid value is IP.
     */
    public readonly rateKey!: pulumi.Output<string>;
    /**
     * The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
     */
    public readonly rateLimit!: pulumi.Output<number>;
    /**
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a RateBasedRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RateBasedRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RateBasedRuleArgs | RateBasedRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RateBasedRuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["predicates"] = state ? state.predicates : undefined;
            inputs["rateKey"] = state ? state.rateKey : undefined;
            inputs["rateLimit"] = state ? state.rateLimit : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RateBasedRuleArgs | undefined;
            if (!args || args.metricName === undefined) {
                throw new Error("Missing required property 'metricName'");
            }
            if (!args || args.rateKey === undefined) {
                throw new Error("Missing required property 'rateKey'");
            }
            if (!args || args.rateLimit === undefined) {
                throw new Error("Missing required property 'rateLimit'");
            }
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["predicates"] = args ? args.predicates : undefined;
            inputs["rateKey"] = args ? args.rateKey : undefined;
            inputs["rateLimit"] = args ? args.rateLimit : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RateBasedRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RateBasedRule resources.
 */
export interface RateBasedRuleState {
    /**
     * Amazon Resource Name (ARN)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name or description for the Amazon CloudWatch metric of this rule.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * The name or description of the rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    readonly predicates?: pulumi.Input<pulumi.Input<inputs.waf.RateBasedRulePredicate>[]>;
    /**
     * Valid value is IP.
     */
    readonly rateKey?: pulumi.Input<string>;
    /**
     * The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
     */
    readonly rateLimit?: pulumi.Input<number>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a RateBasedRule resource.
 */
export interface RateBasedRuleArgs {
    /**
     * The name or description for the Amazon CloudWatch metric of this rule.
     */
    readonly metricName: pulumi.Input<string>;
    /**
     * The name or description of the rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The objects to include in a rule (documented below).
     */
    readonly predicates?: pulumi.Input<pulumi.Input<inputs.waf.RateBasedRulePredicate>[]>;
    /**
     * Valid value is IP.
     */
    readonly rateKey: pulumi.Input<string>;
    /**
     * The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
     */
    readonly rateLimit: pulumi.Input<number>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
