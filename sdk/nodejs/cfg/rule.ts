// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an AWS Config Rule.
 *
 * > **Note:** Config Rule requires an existing `Configuration Recorder` to be present. Use of `dependsOn` is recommended (as shown below) to avoid race conditions.
 *
 * ## Example Usage
 * ### AWS Managed Rules
 *
 * AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const role = new aws.iam.Role("role", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "config.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const foo = new aws.cfg.Recorder("foo", {roleArn: role.arn});
 * const rule = new aws.cfg.Rule("rule", {source: {
 *     owner: "AWS",
 *     sourceIdentifier: "S3_BUCKET_VERSIONING_ENABLED",
 * }}, {
 *     dependsOn: [foo],
 * });
 * const rolePolicy = new aws.iam.RolePolicy("rolePolicy", {
 *     role: role.id,
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *   	{
 *   		"Action": "config:Put*",
 *   		"Effect": "Allow",
 *   		"Resource": "*"
 *
 *   	}
 *   ]
 * }
 * `,
 * });
 * ```
 * ### Custom Rules
 *
 * Custom rules can be used by setting the source owner to `CUSTOM_LAMBDA` and the source identifier to the Amazon Resource Name (ARN) of the Lambda Function. The AWS Config service must have permissions to invoke the Lambda Function, e.g. via the `aws.lambda.Permission` resource. More information about custom rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleRecorder = new aws.cfg.Recorder("exampleRecorder", {});
 * // ... other configuration ...
 * const exampleFunction = new aws.lambda.Function("exampleFunction", {});
 * // ... other configuration ...
 * const examplePermission = new aws.lambda.Permission("examplePermission", {
 *     action: "lambda:InvokeFunction",
 *     "function": exampleFunction.arn,
 *     principal: "config.amazonaws.com",
 * });
 * // ... other configuration ...
 * const exampleRule = new aws.cfg.Rule("exampleRule", {source: {
 *     owner: "CUSTOM_LAMBDA",
 *     sourceIdentifier: exampleFunction.arn,
 * }}, {
 *     dependsOn: [
 *         exampleRecorder,
 *         examplePermission,
 *     ],
 * });
 * ```
 */
export class Rule extends pulumi.CustomResource {
    /**
     * Get an existing Rule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RuleState, opts?: pulumi.CustomResourceOptions): Rule {
        return new Rule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cfg/rule:Rule';

    /**
     * Returns true if the given object is an instance of Rule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rule.__pulumiType;
    }

    /**
     * The ARN of the config rule
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the rule
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A string in JSON format that is passed to the AWS Config rule Lambda function.
     */
    public readonly inputParameters!: pulumi.Output<string | undefined>;
    /**
     * The frequency that you want AWS Config to run evaluations for a rule that
     * is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
     */
    public readonly maximumExecutionFrequency!: pulumi.Output<string | undefined>;
    /**
     * The name of the rule
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the config rule
     */
    public /*out*/ readonly ruleId!: pulumi.Output<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    public readonly scope!: pulumi.Output<outputs.cfg.RuleScope | undefined>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    public readonly source!: pulumi.Output<outputs.cfg.RuleSource>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

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
            const state = argsOrState as RuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["inputParameters"] = state ? state.inputParameters : undefined;
            inputs["maximumExecutionFrequency"] = state ? state.maximumExecutionFrequency : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["ruleId"] = state ? state.ruleId : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["tags"] = state ? state.tags : undefined;
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
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ruleId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Rule.__pulumiType, name, inputs, opts);
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
     * is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
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
    readonly scope?: pulumi.Input<inputs.cfg.RuleScope>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    readonly source?: pulumi.Input<inputs.cfg.RuleSource>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
     * is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
     */
    readonly maximumExecutionFrequency?: pulumi.Input<string>;
    /**
     * The name of the rule
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Scope defines which resources can trigger an evaluation for the rule as documented below.
     */
    readonly scope?: pulumi.Input<inputs.cfg.RuleScope>;
    /**
     * Source specifies the rule owner, the rule identifier, and the notifications that cause
     * the function to evaluate your AWS resources as documented below.
     */
    readonly source: pulumi.Input<inputs.cfg.RuleSource>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
