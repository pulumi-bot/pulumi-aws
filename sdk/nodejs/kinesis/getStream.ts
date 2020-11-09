// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Kinesis Stream for use in other
 * resources.
 *
 * For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const stream = pulumi.output(aws.kinesis.getStream({
 *     name: "stream-name",
 * }, { async: true }));
 * ```
 */
export function getStream(args: GetStreamArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:kinesis/getStream:getStream", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getStream.
 */
export interface GetStreamArgs {
    /**
     * The name of the Kinesis Stream.
     */
    readonly name: string;
    /**
     * A map of tags to assigned to the stream.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getStream.
 */
export interface GetStreamResult {
    /**
     * The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).
     */
    readonly arn: string;
    /**
     * The list of shard ids in the CLOSED state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
     */
    readonly closedShards: string[];
    /**
     * The approximate UNIX timestamp that the stream was created.
     */
    readonly creationTimestamp: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the Kinesis Stream.
     */
    readonly name: string;
    /**
     * The list of shard ids in the OPEN state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
     */
    readonly openShards: string[];
    /**
     * Length of time (in hours) data records are accessible after they are added to the stream.
     */
    readonly retentionPeriod: number;
    /**
     * A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more.
     */
    readonly shardLevelMetrics: string[];
    /**
     * The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.
     */
    readonly status: string;
    /**
     * A map of tags to assigned to the stream.
     */
    readonly tags: {[key: string]: string};
}
