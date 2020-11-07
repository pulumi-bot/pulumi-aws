// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get the Account ID of the [AWS Billing and Cost Management Service Account](http://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-getting-started.html#step-2) for the purpose of permitting in S3 bucket policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = pulumi.output(aws.getBillingServiceAccount({ async: true }));
 * const billingLogs = new aws.s3.Bucket("billing_logs", {
 *     acl: "private",
 *     policy: pulumi.interpolate`{
 *   "Id": "Policy",
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "s3:GetBucketAcl", "s3:GetBucketPolicy"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "arn:aws:s3:::my-billing-tf-test-bucket",
 *       "Principal": {
 *         "AWS": [
 *           "${main.arn}"
 *         ]
 *       }
 *     },
 *     {
 *       "Action": [
 *         "s3:PutObject"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "arn:aws:s3:::my-billing-tf-test-bucket/*",
 *       "Principal": {
 *         "AWS": [
 *           "${main.arn}"
 *         ]
 *       }
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 */
export function getBillingServiceAccount(opts?: pulumi.InvokeOptions): Promise<GetBillingServiceAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getBillingServiceAccount:getBillingServiceAccount", {
    }, opts);
}

/**
 * A collection of values returned by getBillingServiceAccount.
 */
export interface GetBillingServiceAccountResult {
    /**
     * The ARN of the AWS billing service account.
     */
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
