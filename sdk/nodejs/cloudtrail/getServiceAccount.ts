// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the Account ID of the [AWS CloudTrail Service Account](http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-supported-regions.html)
 * in a given region for the purpose of allowing CloudTrail to store trail data in S3.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = pulumi.output(aws.cloudtrail.getServiceAccount());
 * const bucket = new aws.s3.Bucket("bucket", {
 *     forceDestroy: true,
 *     policy: pulumi.interpolate`{
 *   "Version": "2008-10-17",
 *   "Statement": [
 *     {
 *       "Sid": "Put bucket policy needed for trails",
 *       "Effect": "Allow",
 *       "Principal": {
 *         "AWS": "${main.arn}"
 *       },
 *       "Action": "s3:PutObject",
 *       "Resource": "arn:aws:s3:::tf-cloudtrail-logging-test-bucket/*"
 *     },
 *     {
 *       "Sid": "Get bucket policy needed for trails",
 *       "Effect": "Allow",
 *       "Principal": {
 *         "AWS": "${main.arn}"
 *       },
 *       "Action": "s3:GetBucketAcl",
 *       "Resource": "arn:aws:s3:::tf-cloudtrail-logging-test-bucket"
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 */
export function getServiceAccount(args?: GetServiceAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceAccountResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:cloudtrail/getServiceAccount:getServiceAccount", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceAccount.
 */
export interface GetServiceAccountArgs {
    /**
     * Name of the region whose AWS CloudTrail account ID is desired.
     * Defaults to the region from the AWS provider configuration.
     */
    region?: string;
}

/**
 * A collection of values returned by getServiceAccount.
 */
export interface GetServiceAccountResult {
    /**
     * The ARN of the AWS CloudTrail service account in the selected region.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region?: string;
}

export function getServiceAccountApply(args?: GetServiceAccountApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceAccountResult> {
    return pulumi.output(args).apply(a => getServiceAccount(a, opts))
}

/**
 * A collection of arguments for invoking getServiceAccount.
 */
export interface GetServiceAccountApplyArgs {
    /**
     * Name of the region whose AWS CloudTrail account ID is desired.
     * Defaults to the region from the AWS provider configuration.
     */
    region?: pulumi.Input<string>;
}
