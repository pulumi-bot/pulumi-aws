// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SES receipt rule set resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.ses.ReceiptRuleSet("main", {
 *     ruleSetName: "primary-rules",
 * });
 * ```
 *
 * ## Import
 *
 * SES receipt rule sets can be imported using the rule set name.
 *
 * ```sh
 *  $ pulumi import aws:ses/receiptRuleSet:ReceiptRuleSet my_rule_set my_rule_set_name
 * ```
 */
export class ReceiptRuleSet extends pulumi.CustomResource {
    /**
     * Get an existing ReceiptRuleSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReceiptRuleSetState, opts?: pulumi.CustomResourceOptions): ReceiptRuleSet {
        return new ReceiptRuleSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ses/receiptRuleSet:ReceiptRuleSet';

    /**
     * Returns true if the given object is an instance of ReceiptRuleSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReceiptRuleSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReceiptRuleSet.__pulumiType;
    }

    /**
     * SES receipt rule set ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Name of the rule set.
     */
    public readonly ruleSetName!: pulumi.Output<string>;

    /**
     * Create a ReceiptRuleSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReceiptRuleSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReceiptRuleSetArgs | ReceiptRuleSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReceiptRuleSetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["ruleSetName"] = state ? state.ruleSetName : undefined;
        } else {
            const args = argsOrState as ReceiptRuleSetArgs | undefined;
            if ((!args || args.ruleSetName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleSetName'");
            }
            inputs["ruleSetName"] = args ? args.ruleSetName : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ReceiptRuleSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReceiptRuleSet resources.
 */
export interface ReceiptRuleSetState {
    /**
     * SES receipt rule set ARN.
     */
    arn?: pulumi.Input<string>;
    /**
     * Name of the rule set.
     */
    ruleSetName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReceiptRuleSet resource.
 */
export interface ReceiptRuleSetArgs {
    /**
     * Name of the rule set.
     */
    ruleSetName: pulumi.Input<string>;
}
