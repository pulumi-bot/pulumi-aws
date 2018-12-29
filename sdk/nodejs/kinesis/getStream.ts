// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getStream(args: GetStreamArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamResult> {
    return pulumi.runtime.invoke("aws:kinesis/getStream:getStream", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getStream.
 */
export interface GetStreamArgs {
    readonly name: string;
}

/**
 * A collection of values returned by getStream.
 */
export interface GetStreamResult {
    readonly arn: string;
    readonly closedShards: string[];
    readonly creationTimestamp: number;
    readonly openShards: string[];
    readonly retentionPeriod: number;
    readonly shardLevelMetrics: string[];
    readonly status: string;
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
