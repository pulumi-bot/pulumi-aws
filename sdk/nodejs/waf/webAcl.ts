// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a WAF Web ACL Resource
 * 
 * ## Example Usage
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
 * const wafrule = new aws.waf.Rule("wafrule", {
 *     metricName: "tfWAFRule",
 *     predicates: [{
 *         dataId: ipset.id,
 *         negated: false,
 *         type: "IPMatch",
 *     }],
 * }, {dependsOn: [ipset]});
 * const wafAcl = new aws.waf.WebAcl("waf_acl", {
 *     defaultAction: {
 *         type: "ALLOW",
 *     },
 *     metricName: "tfWebACL",
 *     rules: [{
 *         action: {
 *             type: "BLOCK",
 *         },
 *         priority: 1,
 *         ruleId: wafrule.id,
 *         type: "REGULAR",
 *     }],
 * }, {dependsOn: [ipset, wafrule]});
 * ```
 * 
 * ### Logging
 * 
 * > *NOTE:* The Kinesis Firehose Delivery Stream name must begin with `aws-waf-logs-` and be located in `us-east-1` region. See the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/logging.html) for more information about enabling WAF logging.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.waf.WebAcl("example", {
 *     // ... other configuration ...
 *     loggingConfiguration: {
 *         logDestination: aws_kinesis_firehose_delivery_stream_example.arn,
 *         redactedFields: {
 *             fieldToMatches: [
 *                 {
 *                     type: "URI",
 *                 },
 *                 {
 *                     data: "referer",
 *                     type: "HEADER",
 *                 },
 *             ],
 *         },
 *     },
 * });
 * ```
 */
export class WebAcl extends pulumi.CustomResource {
    /**
     * Get an existing WebAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebAclState, opts?: pulumi.CustomResourceOptions): WebAcl {
        return new WebAcl(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:waf/webAcl:WebAcl';

    /**
     * Returns true if the given object is an instance of WebAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WebAcl {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === WebAcl.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
     */
    public readonly defaultAction!: pulumi.Output<{ type: string }>;
    /**
     * Configuration block to enable WAF logging. Detailed below.
     */
    public readonly loggingConfiguration!: pulumi.Output<{ logDestination: string, redactedFields?: { fieldToMatches: { data?: string, type: string }[] } } | undefined>;
    /**
     * The name or description for the Amazon CloudWatch metric of this web ACL.
     */
    public readonly metricName!: pulumi.Output<string>;
    /**
     * The name or description of the web ACL.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
     */
    public readonly rules!: pulumi.Output<{ action?: { type: string }, overrideAction?: { type: string }, priority: number, ruleId: string, type?: string }[] | undefined>;

    /**
     * Create a WebAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebAclArgs | WebAclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WebAclState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["defaultAction"] = state ? state.defaultAction : undefined;
            inputs["loggingConfiguration"] = state ? state.loggingConfiguration : undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as WebAclArgs | undefined;
            if (!args || args.defaultAction === undefined) {
                throw new Error("Missing required property 'defaultAction'");
            }
            if (!args || args.metricName === undefined) {
                throw new Error("Missing required property 'metricName'");
            }
            inputs["defaultAction"] = args ? args.defaultAction : undefined;
            inputs["loggingConfiguration"] = args ? args.loggingConfiguration : undefined;
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super(WebAcl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WebAcl resources.
 */
export interface WebAclState {
    readonly arn?: pulumi.Input<string>;
    /**
     * Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
     */
    readonly defaultAction?: pulumi.Input<{ type: pulumi.Input<string> }>;
    /**
     * Configuration block to enable WAF logging. Detailed below.
     */
    readonly loggingConfiguration?: pulumi.Input<{ logDestination: pulumi.Input<string>, redactedFields?: pulumi.Input<{ fieldToMatches: pulumi.Input<pulumi.Input<{ data?: pulumi.Input<string>, type: pulumi.Input<string> }>[]> }> }>;
    /**
     * The name or description for the Amazon CloudWatch metric of this web ACL.
     */
    readonly metricName?: pulumi.Input<string>;
    /**
     * The name or description of the web ACL.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
     */
    readonly rules?: pulumi.Input<pulumi.Input<{ action?: pulumi.Input<{ type: pulumi.Input<string> }>, overrideAction?: pulumi.Input<{ type: pulumi.Input<string> }>, priority: pulumi.Input<number>, ruleId: pulumi.Input<string>, type?: pulumi.Input<string> }>[]>;
}

/**
 * The set of arguments for constructing a WebAcl resource.
 */
export interface WebAclArgs {
    /**
     * Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
     */
    readonly defaultAction: pulumi.Input<{ type: pulumi.Input<string> }>;
    /**
     * Configuration block to enable WAF logging. Detailed below.
     */
    readonly loggingConfiguration?: pulumi.Input<{ logDestination: pulumi.Input<string>, redactedFields?: pulumi.Input<{ fieldToMatches: pulumi.Input<pulumi.Input<{ data?: pulumi.Input<string>, type: pulumi.Input<string> }>[]> }> }>;
    /**
     * The name or description for the Amazon CloudWatch metric of this web ACL.
     */
    readonly metricName: pulumi.Input<string>;
    /**
     * The name or description of the web ACL.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
     */
    readonly rules?: pulumi.Input<pulumi.Input<{ action?: pulumi.Input<{ type: pulumi.Input<string> }>, overrideAction?: pulumi.Input<{ type: pulumi.Input<string> }>, priority: pulumi.Input<number>, ruleId: pulumi.Input<string>, type?: pulumi.Input<string> }>[]>;
}
