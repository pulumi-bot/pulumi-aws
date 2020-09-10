// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

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
    readonly bucket: string;
    readonly delimiter?: string;
    readonly encodingType?: string;
    readonly fetchOwner?: boolean;
    readonly maxKeys?: number;
    readonly prefix?: string;
    readonly startAfter?: string;
}

/**
 * A collection of values returned by getBucketObjects.
 */
export interface GetBucketObjectsResult {
    readonly bucket: string;
    readonly commonPrefixes: string[];
    readonly delimiter?: string;
    readonly encodingType?: string;
    readonly fetchOwner?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keys: string[];
    readonly maxKeys?: number;
    readonly owners: string[];
    readonly prefix?: string;
    readonly startAfter?: string;
}
