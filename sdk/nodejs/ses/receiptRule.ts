// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides an SES receipt rule resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Add a header to the email and store it in S3
 * const store = new aws.ses.ReceiptRule("store", {
 *     addHeaderActions: [{
 *         headerName: "Custom-Header",
 *         headerValue: "Added by SES",
 *         position: 1,
 *     }],
 *     enabled: true,
 *     recipients: ["karen@example.com"],
 *     ruleSetName: "default-rule-set",
 *     s3Actions: [{
 *         bucketName: "emails",
 *         position: 2,
 *     }],
 *     scanEnabled: true,
 * });
 * ```
 */
export class ReceiptRule extends pulumi.CustomResource {
    /**
     * Get an existing ReceiptRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReceiptRuleState, opts?: pulumi.CustomResourceOptions): ReceiptRule {
        return new ReceiptRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ses/receiptRule:ReceiptRule';

    /**
     * Returns true if the given object is an instance of ReceiptRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReceiptRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReceiptRule.__pulumiType;
    }

    /**
     * A list of Add Header Action blocks. Documented below.
     */
    public readonly addHeaderActions!: pulumi.Output<outputs.ses.ReceiptRuleAddHeaderAction[] | undefined>;
    /**
     * The name of the rule to place this rule after
     */
    public readonly after!: pulumi.Output<string | undefined>;
    /**
     * A list of Bounce Action blocks. Documented below.
     */
    public readonly bounceActions!: pulumi.Output<outputs.ses.ReceiptRuleBounceAction[] | undefined>;
    /**
     * If true, the rule will be enabled
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * A list of Lambda Action blocks. Documented below.
     */
    public readonly lambdaActions!: pulumi.Output<outputs.ses.ReceiptRuleLambdaAction[] | undefined>;
    /**
     * The name of the rule
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of email addresses
     */
    public readonly recipients!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the rule set
     */
    public readonly ruleSetName!: pulumi.Output<string>;
    /**
     * A list of S3 Action blocks. Documented below.
     */
    public readonly s3Actions!: pulumi.Output<outputs.ses.ReceiptRuleS3Action[] | undefined>;
    /**
     * If true, incoming emails will be scanned for spam and viruses
     */
    public readonly scanEnabled!: pulumi.Output<boolean>;
    /**
     * A list of SNS Action blocks. Documented below.
     */
    public readonly snsActions!: pulumi.Output<outputs.ses.ReceiptRuleSnsAction[] | undefined>;
    /**
     * A list of Stop Action blocks. Documented below.
     */
    public readonly stopActions!: pulumi.Output<outputs.ses.ReceiptRuleStopAction[] | undefined>;
    /**
     * Require or Optional
     */
    public readonly tlsPolicy!: pulumi.Output<string>;
    /**
     * A list of WorkMail Action blocks. Documented below.
     */
    public readonly workmailActions!: pulumi.Output<outputs.ses.ReceiptRuleWorkmailAction[] | undefined>;

    /**
     * Create a ReceiptRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReceiptRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReceiptRuleArgs | ReceiptRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ReceiptRuleState | undefined;
            inputs["addHeaderActions"] = state ? state.addHeaderActions : undefined;
            inputs["after"] = state ? state.after : undefined;
            inputs["bounceActions"] = state ? state.bounceActions : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["lambdaActions"] = state ? state.lambdaActions : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["recipients"] = state ? state.recipients : undefined;
            inputs["ruleSetName"] = state ? state.ruleSetName : undefined;
            inputs["s3Actions"] = state ? state.s3Actions : undefined;
            inputs["scanEnabled"] = state ? state.scanEnabled : undefined;
            inputs["snsActions"] = state ? state.snsActions : undefined;
            inputs["stopActions"] = state ? state.stopActions : undefined;
            inputs["tlsPolicy"] = state ? state.tlsPolicy : undefined;
            inputs["workmailActions"] = state ? state.workmailActions : undefined;
        } else {
            const args = argsOrState as ReceiptRuleArgs | undefined;
            if (!args || args.ruleSetName === undefined) {
                throw new Error("Missing required property 'ruleSetName'");
            }
            inputs["addHeaderActions"] = args ? args.addHeaderActions : undefined;
            inputs["after"] = args ? args.after : undefined;
            inputs["bounceActions"] = args ? args.bounceActions : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["lambdaActions"] = args ? args.lambdaActions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["recipients"] = args ? args.recipients : undefined;
            inputs["ruleSetName"] = args ? args.ruleSetName : undefined;
            inputs["s3Actions"] = args ? args.s3Actions : undefined;
            inputs["scanEnabled"] = args ? args.scanEnabled : undefined;
            inputs["snsActions"] = args ? args.snsActions : undefined;
            inputs["stopActions"] = args ? args.stopActions : undefined;
            inputs["tlsPolicy"] = args ? args.tlsPolicy : undefined;
            inputs["workmailActions"] = args ? args.workmailActions : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ReceiptRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReceiptRule resources.
 */
export interface ReceiptRuleState {
    /**
     * A list of Add Header Action blocks. Documented below.
     */
    readonly addHeaderActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleAddHeaderAction>[]>;
    /**
     * The name of the rule to place this rule after
     */
    readonly after?: pulumi.Input<string>;
    /**
     * A list of Bounce Action blocks. Documented below.
     */
    readonly bounceActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleBounceAction>[]>;
    /**
     * If true, the rule will be enabled
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * A list of Lambda Action blocks. Documented below.
     */
    readonly lambdaActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleLambdaAction>[]>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of email addresses
     */
    readonly recipients?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the rule set
     */
    readonly ruleSetName?: pulumi.Input<string>;
    /**
     * A list of S3 Action blocks. Documented below.
     */
    readonly s3Actions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleS3Action>[]>;
    /**
     * If true, incoming emails will be scanned for spam and viruses
     */
    readonly scanEnabled?: pulumi.Input<boolean>;
    /**
     * A list of SNS Action blocks. Documented below.
     */
    readonly snsActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleSnsAction>[]>;
    /**
     * A list of Stop Action blocks. Documented below.
     */
    readonly stopActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleStopAction>[]>;
    /**
     * Require or Optional
     */
    readonly tlsPolicy?: pulumi.Input<string>;
    /**
     * A list of WorkMail Action blocks. Documented below.
     */
    readonly workmailActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleWorkmailAction>[]>;
}

/**
 * The set of arguments for constructing a ReceiptRule resource.
 */
export interface ReceiptRuleArgs {
    /**
     * A list of Add Header Action blocks. Documented below.
     */
    readonly addHeaderActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleAddHeaderAction>[]>;
    /**
     * The name of the rule to place this rule after
     */
    readonly after?: pulumi.Input<string>;
    /**
     * A list of Bounce Action blocks. Documented below.
     */
    readonly bounceActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleBounceAction>[]>;
    /**
     * If true, the rule will be enabled
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * A list of Lambda Action blocks. Documented below.
     */
    readonly lambdaActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleLambdaAction>[]>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of email addresses
     */
    readonly recipients?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the rule set
     */
    readonly ruleSetName: pulumi.Input<string>;
    /**
     * A list of S3 Action blocks. Documented below.
     */
    readonly s3Actions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleS3Action>[]>;
    /**
     * If true, incoming emails will be scanned for spam and viruses
     */
    readonly scanEnabled?: pulumi.Input<boolean>;
    /**
     * A list of SNS Action blocks. Documented below.
     */
    readonly snsActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleSnsAction>[]>;
    /**
     * A list of Stop Action blocks. Documented below.
     */
    readonly stopActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleStopAction>[]>;
    /**
     * Require or Optional
     */
    readonly tlsPolicy?: pulumi.Input<string>;
    /**
     * A list of WorkMail Action blocks. Documented below.
     */
    readonly workmailActions?: pulumi.Input<pulumi.Input<inputs.ses.ReceiptRuleWorkmailAction>[]>;
}
