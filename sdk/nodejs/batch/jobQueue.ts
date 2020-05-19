// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Batch Job Queue resource.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testQueue = new aws.batch.JobQueue("testQueue", {
 *     computeEnvironments: [
 *         aws_batch_compute_environment_test_environment_1.arn,
 *         aws_batch_compute_environment_test_environment_2.arn,
 *     ],
 *     priority: 1,
 *     state: "ENABLED",
 * });
 * ```
 */
export class JobQueue extends pulumi.CustomResource {
    /**
     * Get an existing JobQueue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: JobQueueState, opts?: pulumi.CustomResourceOptions): JobQueue {
        return new JobQueue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:batch/jobQueue:JobQueue';

    /**
     * Returns true if the given object is an instance of JobQueue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is JobQueue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === JobQueue.__pulumiType;
    }

    /**
     * The Amazon Resource Name of the job queue.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specifies the set of compute environments
     * mapped to a job queue and their order.  The position of the compute environments
     * in the list will dictate the order. You can associate up to 3 compute environments
     * with a job queue.
     */
    public readonly computeEnvironments!: pulumi.Output<string[]>;
    /**
     * Specifies the name of the job queue.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The priority of the job queue. Job queues with a higher priority
     * are evaluated first when associated with the same compute environment.
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
     */
    public readonly state!: pulumi.Output<string>;

    /**
     * Create a JobQueue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: JobQueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: JobQueueArgs | JobQueueState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as JobQueueState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["computeEnvironments"] = state ? state.computeEnvironments : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as JobQueueArgs | undefined;
            if (!args || args.computeEnvironments === undefined) {
                throw new Error("Missing required property 'computeEnvironments'");
            }
            if (!args || args.priority === undefined) {
                throw new Error("Missing required property 'priority'");
            }
            if (!args || args.state === undefined) {
                throw new Error("Missing required property 'state'");
            }
            inputs["computeEnvironments"] = args ? args.computeEnvironments : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(JobQueue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering JobQueue resources.
 */
export interface JobQueueState {
    /**
     * The Amazon Resource Name of the job queue.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Specifies the set of compute environments
     * mapped to a job queue and their order.  The position of the compute environments
     * in the list will dictate the order. You can associate up to 3 compute environments
     * with a job queue.
     */
    readonly computeEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the job queue.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The priority of the job queue. Job queues with a higher priority
     * are evaluated first when associated with the same compute environment.
     */
    readonly priority?: pulumi.Input<number>;
    /**
     * The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
     */
    readonly state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a JobQueue resource.
 */
export interface JobQueueArgs {
    /**
     * Specifies the set of compute environments
     * mapped to a job queue and their order.  The position of the compute environments
     * in the list will dictate the order. You can associate up to 3 compute environments
     * with a job queue.
     */
    readonly computeEnvironments: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the name of the job queue.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The priority of the job queue. Job queues with a higher priority
     * are evaluated first when associated with the same compute environment.
     */
    readonly priority: pulumi.Input<number>;
    /**
     * The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
     */
    readonly state: pulumi.Input<string>;
}
