// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Step Function State Machine resource
 *
 * ## Example Usage
 * ### Basic (Standard Workflow)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ...
 * const sfnStateMachine = new aws.sfn.StateMachine("sfnStateMachine", {
 *     roleArn: aws_iam_role.iam_for_sfn.arn,
 *     definition: `{
 *   "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
 *   "StartAt": "HelloWorld",
 *   "States": {
 *     "HelloWorld": {
 *       "Type": "Task",
 *       "Resource": "${aws_lambda_function.lambda.arn}",
 *       "End": true
 *     }
 *   }
 * }
 * `,
 * });
 * ```
 * ### Basic (Express Workflow)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ...
 * const sfnStateMachine = new aws.sfn.StateMachine("sfnStateMachine", {
 *     roleArn: aws_iam_role.iam_for_sfn.arn,
 *     type: "EXPRESS",
 *     definition: `{
 *   "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
 *   "StartAt": "HelloWorld",
 *   "States": {
 *     "HelloWorld": {
 *       "Type": "Task",
 *       "Resource": "${aws_lambda_function.lambda.arn}",
 *       "End": true
 *     }
 *   }
 * }
 * `,
 * });
 * ```
 * ### Logging
 *
 * > *NOTE:* See the [AWS Step Functions Developer Guide](https://docs.aws.amazon.com/step-functions/latest/dg/welcome.html) for more information about enabling Step Function logging.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ...
 * const sfnStateMachine = new aws.sfn.StateMachine("sfnStateMachine", {
 *     roleArn: aws_iam_role.iam_for_sfn.arn,
 *     definition: `{
 *   "Comment": "A Hello World example of the Amazon States Language using an AWS Lambda Function",
 *   "StartAt": "HelloWorld",
 *   "States": {
 *     "HelloWorld": {
 *       "Type": "Task",
 *       "Resource": "${aws_lambda_function.lambda.arn}",
 *       "End": true
 *     }
 *   }
 * }
 * `,
 *     loggingConfiguration: {
 *         logDestination: aws_cloudwatch_log_group.log_group_for_sfn.arn,
 *         includeExecutionData: true,
 *         level: "ERROR",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * State Machines can be imported using the `arn`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:sfn/stateMachine:StateMachine foo arn:aws:states:eu-west-1:123456789098:stateMachine:bar
 * ```
 */
export class StateMachine extends pulumi.CustomResource {
    /**
     * Get an existing StateMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StateMachineState, opts?: pulumi.CustomResourceOptions): StateMachine {
        return new StateMachine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sfn/stateMachine:StateMachine';

    /**
     * Returns true if the given object is an instance of StateMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StateMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StateMachine.__pulumiType;
    }

    /**
     * The ARN of the state machine.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The date the state machine was created.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The Amazon States Language definition of the state machine.
     */
    public readonly definition!: pulumi.Output<string>;
    /**
     * Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
     */
    public readonly loggingConfiguration!: pulumi.Output<outputs.sfn.StateMachineLoggingConfiguration>;
    /**
     * The name of the state machine.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The current status of the state machine. Either "ACTIVE" or "DELETING".
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a StateMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StateMachineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StateMachineArgs | StateMachineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as StateMachineState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["creationDate"] = state ? state.creationDate : undefined;
            inputs["definition"] = state ? state.definition : undefined;
            inputs["loggingConfiguration"] = state ? state.loggingConfiguration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as StateMachineArgs | undefined;
            if ((!args || args.definition === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'definition'");
            }
            if ((!args || args.roleArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["loggingConfiguration"] = args ? args.loggingConfiguration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(StateMachine.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StateMachine resources.
 */
export interface StateMachineState {
    /**
     * The ARN of the state machine.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The date the state machine was created.
     */
    readonly creationDate?: pulumi.Input<string>;
    /**
     * The Amazon States Language definition of the state machine.
     */
    readonly definition?: pulumi.Input<string>;
    /**
     * Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
     */
    readonly loggingConfiguration?: pulumi.Input<inputs.sfn.StateMachineLoggingConfiguration>;
    /**
     * The name of the state machine.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The current status of the state machine. Either "ACTIVE" or "DELETING".
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StateMachine resource.
 */
export interface StateMachineArgs {
    /**
     * The Amazon States Language definition of the state machine.
     */
    readonly definition: pulumi.Input<string>;
    /**
     * Defines what execution history events are logged and where they are logged. The `loggingConfiguration` parameter is only valid when `type` is set to `EXPRESS`. Defaults to `OFF`. For more information see [Logging Express Workflows](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html) and [Log Levels](https://docs.aws.amazon.com/step-functions/latest/dg/cloudwatch-log-level.html) in the AWS Step Functions User Guide.
     */
    readonly loggingConfiguration?: pulumi.Input<inputs.sfn.StateMachineLoggingConfiguration>;
    /**
     * The name of the state machine.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM role to use for this state machine.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Determines whether a Standard or Express state machine is created. The default is STANDARD. You cannot update the type of a state machine once it has been created. Valid Values: STANDARD | EXPRESS
     */
    readonly type?: pulumi.Input<string>;
}
