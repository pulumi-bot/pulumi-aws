// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

import {Bucket} from "./bucket";

export class BucketObject extends fabric.Resource {
    public readonly acl?: fabric.Computed<string>;
    public readonly bucket: fabric.Computed<Bucket>;
    public readonly cacheControl?: fabric.Computed<string>;
    public readonly content?: fabric.Computed<string>;
    public readonly contentDisposition?: fabric.Computed<string>;
    public readonly contentEncoding?: fabric.Computed<string>;
    public readonly contentLanguage?: fabric.Computed<string>;
    public readonly contentType: fabric.Computed<string>;
    public readonly etag: fabric.Computed<string>;
    public readonly key: fabric.Computed<string>;
    public readonly kmsKeyId?: fabric.Computed<string>;
    public readonly serverSideEncryption: fabric.Computed<string>;
    public readonly source?: fabric.Computed<fabric.asset.Asset>;
    public readonly storageClass: fabric.Computed<string>;
    public readonly tags?: fabric.Computed<{[key: string]: any}>;
    public /*out*/ readonly versionId: fabric.Computed<string>;
    public readonly websiteRedirect?: fabric.Computed<string>;

    constructor(urnName: string, args: BucketObjectArgs, dependsOn?: fabric.Resource[]) {
        if (args.bucket === undefined) {
            throw new Error("Missing required property 'bucket'");
        }
        super("aws:s3/bucketObject:BucketObject", urnName, {
            "acl": args.acl,
            "bucket": args.bucket,
            "cacheControl": args.cacheControl,
            "content": args.content,
            "contentDisposition": args.contentDisposition,
            "contentEncoding": args.contentEncoding,
            "contentLanguage": args.contentLanguage,
            "contentType": args.contentType,
            "etag": args.etag,
            "key": args.key,
            "kmsKeyId": args.kmsKeyId,
            "serverSideEncryption": args.serverSideEncryption,
            "source": args.source,
            "storageClass": args.storageClass,
            "tags": args.tags,
            "websiteRedirect": args.websiteRedirect,
            "versionId": undefined,
        }, dependsOn);
    }
}

export interface BucketObjectArgs {
    readonly acl?: fabric.ComputedValue<string>;
    readonly bucket: fabric.ComputedValue<Bucket>;
    readonly cacheControl?: fabric.ComputedValue<string>;
    readonly content?: fabric.ComputedValue<string>;
    readonly contentDisposition?: fabric.ComputedValue<string>;
    readonly contentEncoding?: fabric.ComputedValue<string>;
    readonly contentLanguage?: fabric.ComputedValue<string>;
    readonly contentType?: fabric.ComputedValue<string>;
    readonly etag?: fabric.ComputedValue<string>;
    readonly key?: fabric.ComputedValue<string>;
    readonly kmsKeyId?: fabric.ComputedValue<string>;
    readonly serverSideEncryption?: fabric.ComputedValue<string>;
    readonly source?: fabric.asset.Asset;
    readonly storageClass?: fabric.ComputedValue<string>;
    readonly tags?: fabric.ComputedValue<{[key: string]: any}>;
    readonly websiteRedirect?: fabric.ComputedValue<string>;
}

