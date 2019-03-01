// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AWS Config Rule.
 * 
 * > **Note:** Config Rule requires an existing [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `depends_on` is recommended (as shown below) to avoid race conditions.
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of the config rule
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * Description of the rule
     */
    public readonly description: pulumi.Output<string | undefined>;
    /**
     * A string in JSON format that is passed to the AWS Config rule Lambda function.
     */
    public readonly inputParameters: pulumi.Output<string | undefined>;
    /**
     * The frequency that you want AWS Config to run evaluations for a rule that
     * is triggered periodically. If specified, requires `message_type` to be `ScheduledNotification`.
     */
    public readonly maximumExecutionFrequency: pulumi.Output<string | undefined>;
    /**
     * The name of the rule
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the config rule
     */
    public /*out*/ readonly ruleId: pulumi.Output<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    public readonly scope: pulumi.Output<{ complianceResourceId?: string, complianceResourceTypes?: string[], tagKey?: string, tagValue?: string } | undefined>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    public readonly source: pulumi.Output<{ owner: string, sourceDetails?: { eventSource?: string, maximumExecutionFrequency?: string, messageType?: string }[], sourceIdentifier: string }>;

    /**
     * Create a Rule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RuleArgs | RuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RuleState = argsOrState as RuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["inputParameters"] = state ? state.inputParameters : undefined;
            inputs["maximumExecutionFrequency"] = state ? state.maximumExecutionFrequency : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ruleId"] = state ? state.ruleId : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["source"] = state ? state.source : undefined;
        } else {
            const args = argsOrState as RuleArgs | undefined;
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["inputParameters"] = args ? args.inputParameters : undefined;
            inputs["maximumExecutionFrequency"] = args ? args.maximumExecutionFrequency : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ruleId"] = undefined /*out*/;
        }
        super("aws:cfg/rule:Rule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rule resources.
 */
export interface RuleState {
    /**
     * The ARN of the config rule
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Description of the rule
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A string in JSON format that is passed to the AWS Config rule Lambda function.
     */
    readonly inputParameters?: pulumi.Input<string>;
    /**
     * The frequency that you want AWS Config to run evaluations for a rule that
     * is triggered periodically. If specified, requires `message_type` to be `ScheduledNotification`.
     */
    readonly maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the config rule
     */
    readonly ruleId?: pulumi.Input<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    readonly scope?: pulumi.Input<{ complianceResourceId?: pulumi.Input<string>, complianceResourceTypes?: pulumi.Input<pulumi.Input<string>[]>, tagKey?: pulumi.Input<string>, tagValue?: pulumi.Input<string> }>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    readonly source?: pulumi.Input<{ owner: pulumi.Input<string>, sourceDetails?: pulumi.Input<pulumi.Input<{ eventSource?: pulumi.Input<string>, maximumExecutionFrequency?: pulumi.Input<string>, messageType?: pulumi.Input<string> }>[]>, sourceIdentifier: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a Rule resource.
 */
export interface RuleArgs {
    /**
     * Description of the rule
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A string in JSON format that is passed to the AWS Config rule Lambda function.
     */
    readonly inputParameters?: pulumi.Input<string>;
    /**
     * The frequency that you want AWS Config to run evaluations for a rule that
     * is triggered periodically. If specified, requires `message_type` to be `ScheduledNotification`.
     */
    readonly maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    readonly scope?: pulumi.Input<{ complianceResourceId?: pulumi.Input<string>, complianceResourceTypes?: pulumi.Input<pulumi.Input<string>[]>, tagKey?: pulumi.Input<string>, tagValue?: pulumi.Input<string> }>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    readonly source: pulumi.Input<{ owner: pulumi.Input<string>, sourceDetails?: pulumi.Input<pulumi.Input<{ eventSource?: pulumi.Input<string>, maximumExecutionFrequency?: pulumi.Input<string>, messageType?: pulumi.Input<string> }>[]>, sourceIdentifier: pulumi.Input<string> }>;
}
