// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ARN and URL of queue in AWS Simple Queue Service (SQS).
 * By using this data source, you can reference SQS queues without having to hardcode
 * the ARNs as input.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.sqs.getQueue({
 *     name: "queue",
 * }));
 * ```
 */
export function getQueue(args: GetQueueArgs, opts?: pulumi.InvokeOptions): Promise<GetQueueResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:sqs/getQueue:getQueue", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getQueue.
 */
export interface GetQueueArgs {
    /**
     * The name of the queue to match.
     */
    name: string;
    /**
     * A map of tags for the resource.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getQueue.
 */
export interface GetQueueResult {
    /**
     * The Amazon Resource Name (ARN) of the queue.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * A map of tags for the resource.
     */
    readonly tags: {[key: string]: string};
    /**
     * The URL of the queue.
     */
    readonly url: string;
}

export function getQueueApply(args: GetQueueApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueueResult> {
    return pulumi.output(args).apply(a => getQueue(a, opts))
}

/**
 * A collection of arguments for invoking getQueue.
 */
export interface GetQueueApplyArgs {
    /**
     * The name of the queue to match.
     */
    name: pulumi.Input<string>;
    /**
     * A map of tags for the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
