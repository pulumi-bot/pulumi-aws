// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides details about a specific S3 bucket.
 * 
 * This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
 * Distribution.
 * 
 * ## Example Usage
 * 
 * ### Route53 Record
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const testZone = pulumi.output(aws.route53.getZone({
 *     name: "test.com.",
 * }));
 * const selected = pulumi.output(aws.s3.getBucket({
 *     bucket: "bucket.test.com",
 * }));
 * const example = new aws.route53.Record("example", {
 *     aliases: [{
 *         name: selected.websiteDomain,
 *         zoneId: selected.hostedZoneId,
 *     }],
 *     type: "A",
 *     zoneId: testZone.id,
 * });
 * ```
 * 
 * ### CloudFront Origin
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const selected = pulumi.output(aws.s3.getBucket({
 *     bucket: "a-test-bucket",
 * }));
 * const test = new aws.cloudfront.Distribution("test", {
 *     origins: [{
 *         domainName: selected.bucketDomainName,
 *         originId: "s3-selected-bucket",
 *     }],
 * });
 * ```
 */
export function getBucket(args: GetBucketArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:s3/getBucket:getBucket", {
        "bucket": args.bucket,
    }, opts);
}

/**
 * A collection of arguments for invoking getBucket.
 */
export interface GetBucketArgs {
    /**
     * The name of the bucket
     */
    readonly bucket: string;
}

/**
 * A collection of values returned by getBucket.
 */
export interface GetBucketResult {
    /**
     * The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    readonly arn: string;
    readonly bucket: string;
    /**
     * The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
     */
    readonly bucketDomainName: string;
    /**
     * The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
     */
    readonly bucketRegionalDomainName: string;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    readonly hostedZoneId: string;
    /**
     * The AWS region this bucket resides in.
     */
    readonly region: string;
    /**
     * The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    readonly websiteDomain: string;
    /**
     * The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    readonly websiteEndpoint: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
