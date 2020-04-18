// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to lookup current AWS partition in which this provider is working
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const current = aws.getPartition();
 * const s3Policy = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["s3:ListBucket"],
 *         resources: [`arn:${current.partition}:s3:::my-bucket`],
 *         sid: "1",
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/partition.html.markdown.
 */
export function getPartition(opts?: pulumi.InvokeOptions): Promise<GetPartitionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getPartition:getPartition", {
    }, opts);
}

/**
 * A collection of values returned by getPartition.
 */
export interface GetPartitionResult {
    readonly dnsSuffix: string;
    readonly partition: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
