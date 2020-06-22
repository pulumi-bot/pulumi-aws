// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > **NOTE on `maxKeys`:** Retrieving very large numbers of keys can adversely affect this provider's performance.
 *
 * The bucket-objects data source returns keys (i.e., file names) and other metadata about objects in an S3 bucket.
 *
 * ## Example Usage
 *
 * The following example retrieves a list of all object keys in an S3 bucket and creates corresponding object data sources:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myObjects = pulumi.output(aws.s3.getBucketObjects({
 *     bucket: "ourcorp",
 * }, { async: true }));
 * const objectInfo: pulumi.Output<aws.s3.GetBucketObjectResult>[] = [];
 * for (let i = 0; i < myObjects.apply(myObjects => myObjects.keys.length); i++) {
 *     objectInfo.push(pulumi.all([myObjects, myObjects]).apply(([myObjects, myObjects1]) => aws.s3.getBucketObject({
 *         bucket: myObjects.bucket,
 *         key: myObjects1.keys[i],
 *     }, { async: true })));
 * }
 * ```
 */
export function getBucketObjects(args: GetBucketObjectsArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketObjectsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:s3/getBucketObjects:getBucketObjects", {
        "bucket": args.bucket,
        "delimiter": args.delimiter,
        "encodingType": args.encodingType,
        "fetchOwner": args.fetchOwner,
        "maxKeys": args.maxKeys,
        "prefix": args.prefix,
        "startAfter": args.startAfter,
    }, opts);
}

/**
 * A collection of arguments for invoking getBucketObjects.
 */
export interface GetBucketObjectsArgs {
    /**
     * Lists object keys in this S3 bucket. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified
     */
    readonly bucket: string;
    /**
     * A character used to group keys (Default: none)
     */
    readonly delimiter?: string;
    /**
     * Encodes keys using this method (Default: none; besides none, only "url" can be used)
     */
    readonly encodingType?: string;
    /**
     * Boolean specifying whether to populate the owner list (Default: false)
     */
    readonly fetchOwner?: boolean;
    /**
     * Maximum object keys to return (Default: 1000)
     */
    readonly maxKeys?: number;
    /**
     * Limits results to object keys with this prefix (Default: none)
     */
    readonly prefix?: string;
    /**
     * Returns key names lexicographically after a specific object key in your bucket (Default: none; S3 lists object keys in UTF-8 character encoding in lexicographical order)
     */
    readonly startAfter?: string;
}

/**
 * A collection of values returned by getBucketObjects.
 */
export interface GetBucketObjectsResult {
    readonly bucket: string;
    /**
     * List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
     */
    readonly commonPrefixes: string[];
    readonly delimiter?: string;
    readonly encodingType?: string;
    readonly fetchOwner?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of strings representing object keys
     */
    readonly keys: string[];
    readonly maxKeys?: number;
    /**
     * List of strings representing object owner IDs (see `fetchOwner` above)
     */
    readonly owners: string[];
    readonly prefix?: string;
    readonly startAfter?: string;
}
