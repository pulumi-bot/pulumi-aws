// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Glue Job resource.
 *
 * > Glue functionality, such as monitoring and logging of jobs, is typically managed with the `defaultArguments` argument. See the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the Glue developer guide for additional information.
 *
 * ## Example Usage
 * ### Python Job
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Job("example", {
 *     roleArn: aws_iam_role.example.arn,
 *     command: {
 *         scriptLocation: `s3://${aws_s3_bucket.example.bucket}/example.py`,
 *     },
 * });
 * ```
 * ### Scala Job
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Job("example", {
 *     roleArn: aws_iam_role.example.arn,
 *     command: {
 *         scriptLocation: `s3://${aws_s3_bucket.example.bucket}/example.scala`,
 *     },
 *     defaultArguments: {
 *         "--job-language": "scala",
 *     },
 * });
 * ```
 * ### Enabling CloudWatch Logs and Metrics
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {retentionInDays: 14});
 * // ... other configuration ...
 * const exampleJob = new aws.glue.Job("exampleJob", {defaultArguments: {
 *     "--continuous-log-logGroup": exampleLogGroup.name,
 *     "--enable-continuous-cloudwatch-log": "true",
 *     "--enable-continuous-log-filter": "true",
 *     "--enable-metrics": "",
 * }});
 * ```
 *
 * ## Import
 *
 * Glue Jobs can be imported using `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/job:Job MyJob MyJob
 * ```
 */
export class Job extends pulumi.CustomResource {
    /**
     * Get an existing Job resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobState, opts?: pulumi.CustomResourceOptions): Job {
        return new Job(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/job:Job';

    /**
     * Returns true if the given object is an instance of Job.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Job {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Job.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of Glue Job
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The command of the job. Defined below.
     */
    public readonly command!: pulumi.Output<outputs.glue.JobCommand>;
    /**
     * The list of connections used for this job.
     */
    public readonly connections!: pulumi.Output<string[] | undefined>;
    /**
     * The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     */
    public readonly defaultArguments!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Description of the job.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Execution property of the job. Defined below.
     */
    public readonly executionProperty!: pulumi.Output<outputs.glue.JobExecutionProperty>;
    /**
     * The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     */
    public readonly glueVersion!: pulumi.Output<string>;
    /**
     * The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
     */
    public readonly maxCapacity!: pulumi.Output<number>;
    /**
     * The maximum number of times to retry this job if it fails.
     */
    public readonly maxRetries!: pulumi.Output<number | undefined>;
    /**
     * The name you assign to this job. It must be unique in your account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Non-overridable arguments for this job, specified as name-value pairs.
     */
    public readonly nonOverridableArguments!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Notification property of the job. Defined below.
     */
    public readonly notificationProperty!: pulumi.Output<outputs.glue.JobNotificationProperty>;
    /**
     * The number of workers of a defined workerType that are allocated when a job runs.
     */
    public readonly numberOfWorkers!: pulumi.Output<number | undefined>;
    /**
     * The ARN of the IAM role associated with this job.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The name of the Security Configuration to be associated with the job.
     */
    public readonly securityConfiguration!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The job timeout in minutes. The default is 2880 minutes (48 hours).
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
     */
    public readonly workerType!: pulumi.Output<string | undefined>;

    /**
     * Create a Job resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobArgs | JobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as JobState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["command"] = state ? state.command : undefined;
            inputs["connections"] = state ? state.connections : undefined;
            inputs["defaultArguments"] = state ? state.defaultArguments : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["executionProperty"] = state ? state.executionProperty : undefined;
            inputs["glueVersion"] = state ? state.glueVersion : undefined;
            inputs["maxCapacity"] = state ? state.maxCapacity : undefined;
            inputs["maxRetries"] = state ? state.maxRetries : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nonOverridableArguments"] = state ? state.nonOverridableArguments : undefined;
            inputs["notificationProperty"] = state ? state.notificationProperty : undefined;
            inputs["numberOfWorkers"] = state ? state.numberOfWorkers : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["securityConfiguration"] = state ? state.securityConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["workerType"] = state ? state.workerType : undefined;
        } else {
            const args = argsOrState as JobArgs | undefined;
            if ((!args || args.command === undefined) && !opts.urn) {
                throw new Error("Missing required property 'command'");
            }
            if ((!args || args.roleArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["command"] = args ? args.command : undefined;
            inputs["connections"] = args ? args.connections : undefined;
            inputs["defaultArguments"] = args ? args.defaultArguments : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["executionProperty"] = args ? args.executionProperty : undefined;
            inputs["glueVersion"] = args ? args.glueVersion : undefined;
            inputs["maxCapacity"] = args ? args.maxCapacity : undefined;
            inputs["maxRetries"] = args ? args.maxRetries : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nonOverridableArguments"] = args ? args.nonOverridableArguments : undefined;
            inputs["notificationProperty"] = args ? args.notificationProperty : undefined;
            inputs["numberOfWorkers"] = args ? args.numberOfWorkers : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["securityConfiguration"] = args ? args.securityConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["workerType"] = args ? args.workerType : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Job.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Job resources.
 */
export interface JobState {
    /**
     * Amazon Resource Name (ARN) of Glue Job
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The command of the job. Defined below.
     */
    readonly command?: pulumi.Input<inputs.glue.JobCommand>;
    /**
     * The list of connections used for this job.
     */
    readonly connections?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     */
    readonly defaultArguments?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Description of the job.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Execution property of the job. Defined below.
     */
    readonly executionProperty?: pulumi.Input<inputs.glue.JobExecutionProperty>;
    /**
     * The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     */
    readonly glueVersion?: pulumi.Input<string>;
    /**
     * The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
     */
    readonly maxCapacity?: pulumi.Input<number>;
    /**
     * The maximum number of times to retry this job if it fails.
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The name you assign to this job. It must be unique in your account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Non-overridable arguments for this job, specified as name-value pairs.
     */
    readonly nonOverridableArguments?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Notification property of the job. Defined below.
     */
    readonly notificationProperty?: pulumi.Input<inputs.glue.JobNotificationProperty>;
    /**
     * The number of workers of a defined workerType that are allocated when a job runs.
     */
    readonly numberOfWorkers?: pulumi.Input<number>;
    /**
     * The ARN of the IAM role associated with this job.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The name of the Security Configuration to be associated with the job.
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The job timeout in minutes. The default is 2880 minutes (48 hours).
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
     */
    readonly workerType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Job resource.
 */
export interface JobArgs {
    /**
     * The command of the job. Defined below.
     */
    readonly command: pulumi.Input<inputs.glue.JobCommand>;
    /**
     * The list of connections used for this job.
     */
    readonly connections?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
     */
    readonly defaultArguments?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Description of the job.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Execution property of the job. Defined below.
     */
    readonly executionProperty?: pulumi.Input<inputs.glue.JobExecutionProperty>;
    /**
     * The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     */
    readonly glueVersion?: pulumi.Input<string>;
    /**
     * The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs. `Required` when `pythonshell` is set, accept either `0.0625` or `1.0`. Use `numberOfWorkers` and `workerType` arguments instead with `glueVersion` `2.0` and above.
     */
    readonly maxCapacity?: pulumi.Input<number>;
    /**
     * The maximum number of times to retry this job if it fails.
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The name you assign to this job. It must be unique in your account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Non-overridable arguments for this job, specified as name-value pairs.
     */
    readonly nonOverridableArguments?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Notification property of the job. Defined below.
     */
    readonly notificationProperty?: pulumi.Input<inputs.glue.JobNotificationProperty>;
    /**
     * The number of workers of a defined workerType that are allocated when a job runs.
     */
    readonly numberOfWorkers?: pulumi.Input<number>;
    /**
     * The ARN of the IAM role associated with this job.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * The name of the Security Configuration to be associated with the job.
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The job timeout in minutes. The default is 2880 minutes (48 hours).
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
     */
    readonly workerType?: pulumi.Input<string>;
}
