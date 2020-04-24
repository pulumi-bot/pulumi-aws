// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the Account ID of the [AWS CloudTrail Service Account](http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-supported-regions.html)
 * in a given region for the purpose of allowing CloudTrail to store trail data in S3.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const main = aws.cloudtrail.getServiceAccount();
 * const bucket = new aws.s3.Bucket("bucket", {
 *     forceDestroy: true,
 *     policy: `{
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
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/cloudtrail_service_account.html.markdown.
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
    readonly region?: string;
}

/**
 * A collection of values returned by getServiceAccount.
 */
export interface GetServiceAccountResult {
    /**
     * The ARN of the AWS CloudTrail service account in the selected region.
     */
    readonly arn: string;
    readonly region?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
