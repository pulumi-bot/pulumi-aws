// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an SSM Maintenance Window Task resource
 * 
 * ## Example Usage
 * 
 * ### Automation Tasks
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ssm.MaintenanceWindowTask("example", {
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     priority: 1,
 *     serviceRoleArn: aws_iam_role_example.arn,
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [aws_instance_example.id],
 *     }],
 *     taskArn: "AWS-RestartEC2Instance",
 *     taskInvocationParameters: {
 *         automationParameters: {
 *             documentVersion: "$LATEST",
 *             parameters: [{
 *                 name: "InstanceId",
 *                 values: [aws_instance_example.id],
 *             }],
 *         },
 *     },
 *     taskType: "AUTOMATION",
 *     windowId: aws_ssm_maintenance_window_example.id,
 * });
 * ```
 * 
 * ### Lambda Tasks
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ssm.MaintenanceWindowTask("example", {
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     priority: 1,
 *     serviceRoleArn: aws_iam_role_example.arn,
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [aws_instance_example.id],
 *     }],
 *     taskArn: aws_lambda_function_example.arn,
 *     taskInvocationParameters: {
 *         lambdaParameters: {
 *             clientContext: Buffer.from("{\"key1\":\"value1\"}").toString("base64"),
 *             payload: "{\"key1\":\"value1\"}",
 *         },
 *     },
 *     taskType: "LAMBDA",
 *     windowId: aws_ssm_maintenance_window_example.id,
 * });
 * ```
 * 
 * ### Run Command Tasks
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ssm.MaintenanceWindowTask("example", {
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     priority: 1,
 *     serviceRoleArn: aws_iam_role_example.arn,
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [aws_instance_example.id],
 *     }],
 *     taskArn: "AWS-RunShellScript",
 *     taskInvocationParameters: {
 *         runCommandParameters: {
 *             notificationConfig: {
 *                 notificationArn: aws_sns_topic_example.arn,
 *                 notificationEvents: ["All"],
 *                 notificationType: ["Command"],
 *             },
 *             outputS3Bucket: aws_s3_bucket_example.bucket,
 *             outputS3KeyPrefix: "output",
 *             parameters: [{
 *                 name: "commands",
 *                 values: ["date"],
 *             }],
 *             serviceRoleArn: aws_iam_role_example.arn,
 *             timeoutSeconds: 600,
 *         },
 *     },
 *     taskType: "RUN_COMMAND",
 *     windowId: aws_ssm_maintenance_window_example.id,
 * });
 * ```
 * 
 * ### Step Function Tasks
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ssm.MaintenanceWindowTask("example", {
 *     maxConcurrency: "2",
 *     maxErrors: "1",
 *     priority: 1,
 *     serviceRoleArn: aws_iam_role_example.arn,
 *     targets: [{
 *         key: "InstanceIds",
 *         values: [aws_instance_example.id],
 *     }],
 *     taskArn: aws_sfn_activity_example.id,
 *     taskInvocationParameters: {
 *         stepFunctionsParameters: {
 *             input: "{\"key1\":\"value1\"}",
 *             name: "example",
 *         },
 *     },
 *     taskType: "STEP_FUNCTIONS",
 *     windowId: aws_ssm_maintenance_window_example.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_maintenance_window_task.html.markdown.
 */
export class MaintenanceWindowTask extends pulumi.CustomResource {
    /**
     * Get an existing MaintenanceWindowTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MaintenanceWindowTaskState, opts?: pulumi.CustomResourceOptions): MaintenanceWindowTask {
        return new MaintenanceWindowTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/maintenanceWindowTask:MaintenanceWindowTask';

    /**
     * Returns true if the given object is an instance of MaintenanceWindowTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MaintenanceWindowTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MaintenanceWindowTask.__pulumiType;
    }

    /**
     * The description of the maintenance window task.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A structure containing information about an Amazon S3 bucket to write instance-level logs to. Use `taskInvocationParameters` configuration block `runCommandParameters` configuration block `output_s3_*` arguments instead. Conflicts with `taskInvocationParameters`. Documented below.
     */
    public readonly loggingInfo!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskLoggingInfo | undefined>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    public readonly maxConcurrency!: pulumi.Output<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    public readonly maxErrors!: pulumi.Output<string>;
    /**
     * The parameter name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The IAM service role to assume during task execution.
     */
    public readonly serviceRoleArn!: pulumi.Output<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    public readonly targets!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskTarget[]>;
    /**
     * The ARN of the task to execute.
     */
    public readonly taskArn!: pulumi.Output<string>;
    /**
     * The parameters for task execution. This argument is conflict with `taskParameters` and `loggingInfo`.
     */
    public readonly taskInvocationParameters!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskTaskInvocationParameters | undefined>;
    /**
     * A structure containing information about parameters required by the particular `taskArn`. Use `parameter` configuration blocks under the `taskInvocationParameters` configuration block instead. Conflicts with `taskInvocationParameters`. Documented below.
     */
    public readonly taskParameters!: pulumi.Output<outputs.ssm.MaintenanceWindowTaskTaskParameter[] | undefined>;
    /**
     * The type of task being registered. The only allowed value is `RUN_COMMAND`.
     */
    public readonly taskType!: pulumi.Output<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    public readonly windowId!: pulumi.Output<string>;

    /**
     * Create a MaintenanceWindowTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MaintenanceWindowTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MaintenanceWindowTaskArgs | MaintenanceWindowTaskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MaintenanceWindowTaskState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["loggingInfo"] = state ? state.loggingInfo : undefined;
            inputs["maxConcurrency"] = state ? state.maxConcurrency : undefined;
            inputs["maxErrors"] = state ? state.maxErrors : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["serviceRoleArn"] = state ? state.serviceRoleArn : undefined;
            inputs["targets"] = state ? state.targets : undefined;
            inputs["taskArn"] = state ? state.taskArn : undefined;
            inputs["taskInvocationParameters"] = state ? state.taskInvocationParameters : undefined;
            inputs["taskParameters"] = state ? state.taskParameters : undefined;
            inputs["taskType"] = state ? state.taskType : undefined;
            inputs["windowId"] = state ? state.windowId : undefined;
        } else {
            const args = argsOrState as MaintenanceWindowTaskArgs | undefined;
            if (!args || args.maxConcurrency === undefined) {
                throw new Error("Missing required property 'maxConcurrency'");
            }
            if (!args || args.maxErrors === undefined) {
                throw new Error("Missing required property 'maxErrors'");
            }
            if (!args || args.serviceRoleArn === undefined) {
                throw new Error("Missing required property 'serviceRoleArn'");
            }
            if (!args || args.targets === undefined) {
                throw new Error("Missing required property 'targets'");
            }
            if (!args || args.taskArn === undefined) {
                throw new Error("Missing required property 'taskArn'");
            }
            if (!args || args.taskType === undefined) {
                throw new Error("Missing required property 'taskType'");
            }
            if (!args || args.windowId === undefined) {
                throw new Error("Missing required property 'windowId'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["loggingInfo"] = args ? args.loggingInfo : undefined;
            inputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            inputs["maxErrors"] = args ? args.maxErrors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["serviceRoleArn"] = args ? args.serviceRoleArn : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["taskArn"] = args ? args.taskArn : undefined;
            inputs["taskInvocationParameters"] = args ? args.taskInvocationParameters : undefined;
            inputs["taskParameters"] = args ? args.taskParameters : undefined;
            inputs["taskType"] = args ? args.taskType : undefined;
            inputs["windowId"] = args ? args.windowId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MaintenanceWindowTask.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MaintenanceWindowTask resources.
 */
export interface MaintenanceWindowTaskState {
    /**
     * The description of the maintenance window task.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A structure containing information about an Amazon S3 bucket to write instance-level logs to. Use `taskInvocationParameters` configuration block `runCommandParameters` configuration block `output_s3_*` arguments instead. Conflicts with `taskInvocationParameters`. Documented below.
     *      * @deprecated use 'task_invocation_parameters' argument instead     */
    readonly loggingInfo?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskLoggingInfo>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    readonly maxConcurrency?: pulumi.Input<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    readonly maxErrors?: pulumi.Input<string>;
    /**
     * The parameter name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The IAM service role to assume during task execution.
     */
    readonly serviceRoleArn?: pulumi.Input<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    readonly targets?: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTaskTarget>[]>;
    /**
     * The ARN of the task to execute.
     */
    readonly taskArn?: pulumi.Input<string>;
    /**
     * The parameters for task execution. This argument is conflict with `taskParameters` and `loggingInfo`.
     */
    readonly taskInvocationParameters?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskTaskInvocationParameters>;
    /**
     * A structure containing information about parameters required by the particular `taskArn`. Use `parameter` configuration blocks under the `taskInvocationParameters` configuration block instead. Conflicts with `taskInvocationParameters`. Documented below.
     *      * @deprecated use 'task_invocation_parameters' argument instead     */
    readonly taskParameters?: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTaskTaskParameter>[]>;
    /**
     * The type of task being registered. The only allowed value is `RUN_COMMAND`.
     */
    readonly taskType?: pulumi.Input<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    readonly windowId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MaintenanceWindowTask resource.
 */
export interface MaintenanceWindowTaskArgs {
    /**
     * The description of the maintenance window task.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A structure containing information about an Amazon S3 bucket to write instance-level logs to. Use `taskInvocationParameters` configuration block `runCommandParameters` configuration block `output_s3_*` arguments instead. Conflicts with `taskInvocationParameters`. Documented below.
     *      * @deprecated use 'task_invocation_parameters' argument instead     */
    readonly loggingInfo?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskLoggingInfo>;
    /**
     * The maximum number of targets this task can be run for in parallel.
     */
    readonly maxConcurrency: pulumi.Input<string>;
    /**
     * The maximum number of errors allowed before this task stops being scheduled.
     */
    readonly maxErrors: pulumi.Input<string>;
    /**
     * The parameter name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The priority of the task in the Maintenance Window, the lower the number the higher the priority. Tasks in a Maintenance Window are scheduled in priority order with tasks that have the same priority scheduled in parallel.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The IAM service role to assume during task execution.
     */
    readonly serviceRoleArn: pulumi.Input<string>;
    /**
     * The targets (either instances or window target ids). Instances are specified using Key=InstanceIds,Values=instanceid1,instanceid2. Window target ids are specified using Key=WindowTargetIds,Values=window target id1, window target id2.
     */
    readonly targets: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTaskTarget>[]>;
    /**
     * The ARN of the task to execute.
     */
    readonly taskArn: pulumi.Input<string>;
    /**
     * The parameters for task execution. This argument is conflict with `taskParameters` and `loggingInfo`.
     */
    readonly taskInvocationParameters?: pulumi.Input<inputs.ssm.MaintenanceWindowTaskTaskInvocationParameters>;
    /**
     * A structure containing information about parameters required by the particular `taskArn`. Use `parameter` configuration blocks under the `taskInvocationParameters` configuration block instead. Conflicts with `taskInvocationParameters`. Documented below.
     *      * @deprecated use 'task_invocation_parameters' argument instead     */
    readonly taskParameters?: pulumi.Input<pulumi.Input<inputs.ssm.MaintenanceWindowTaskTaskParameter>[]>;
    /**
     * The type of task being registered. The only allowed value is `RUN_COMMAND`.
     */
    readonly taskType: pulumi.Input<string>;
    /**
     * The Id of the maintenance window to register the task with.
     */
    readonly windowId: pulumi.Input<string>;
}
