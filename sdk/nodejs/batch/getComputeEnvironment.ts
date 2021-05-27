// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The Batch Compute Environment data source allows access to details of a specific
 * compute environment within AWS Batch.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const batch_mongo = pulumi.output(aws.batch.getComputeEnvironment({
 *     computeEnvironmentName: "batch-mongo-production",
 * }));
 * ```
 */
export function getComputeEnvironment(args: GetComputeEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetComputeEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:batch/getComputeEnvironment:getComputeEnvironment", {
        "computeEnvironmentName": args.computeEnvironmentName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getComputeEnvironment.
 */
export interface GetComputeEnvironmentArgs {
    /**
     * The name of the Batch Compute Environment
     */
    computeEnvironmentName: string;
    /**
     * Key-value map of resource tags
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getComputeEnvironment.
 */
export interface GetComputeEnvironmentResult {
    /**
     * The ARN of the compute environment.
     */
    readonly arn: string;
    readonly computeEnvironmentName: string;
    /**
     * The ARN of the underlying Amazon ECS cluster used by the compute environment.
     */
    readonly ecsClusterArn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ARN of the IAM role that allows AWS Batch to make calls to other AWS services on your behalf.
     */
    readonly serviceRole: string;
    /**
     * The state of the compute environment (for example, `ENABLED` or `DISABLED`). If the state is `ENABLED`, then the compute environment accepts jobs from a queue and can scale out automatically based on queues.
     */
    readonly state: string;
    /**
     * The current status of the compute environment (for example, `CREATING` or `VALID`).
     */
    readonly status: string;
    /**
     * A short, human-readable string to provide additional details about the current status of the compute environment.
     */
    readonly statusReason: string;
    /**
     * Key-value map of resource tags
     */
    readonly tags: {[key: string]: string};
    /**
     * The type of the compute environment (for example, `MANAGED` or `UNMANAGED`).
     */
    readonly type: string;
}

export function getComputeEnvironmentApply(args: GetComputeEnvironmentApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetComputeEnvironmentResult> {
    return pulumi.output(args).apply(a => getComputeEnvironment(a, opts))
}

/**
 * A collection of arguments for invoking getComputeEnvironment.
 */
export interface GetComputeEnvironmentApplyArgs {
    /**
     * The name of the Batch Compute Environment
     */
    computeEnvironmentName: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
