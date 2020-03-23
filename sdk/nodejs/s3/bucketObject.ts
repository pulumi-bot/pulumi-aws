// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

import {Bucket} from "./bucket";

/**
 * Provides a S3 bucket object resource.
 * 
 * {{% examples %}}
 * ## Example Usage
 * 
 * {{% example %}}
 * ### Encrypting with KMS Key
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const examplekms = new aws.kms.Key("examplekms", {
 *     deletionWindowInDays: 7,
 *     description: "KMS key 1",
 * });
 * const examplebucket = new aws.s3.Bucket("examplebucket", {
 *     acl: "private",
 * });
 * const examplebucketObject = new aws.s3.BucketObject("examplebucketObject", {
 *     bucket: examplebucket.id,
 *     key: "someobject",
 *     kmsKeyId: examplekms.arn,
 *     source: new pulumi.asset.FileAsset("index.html"),
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### Server Side Encryption with S3 Default Master Key
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const examplebucket = new aws.s3.Bucket("examplebucket", {
 *     acl: "private",
 * });
 * const examplebucketObject = new aws.s3.BucketObject("examplebucketObject", {
 *     bucket: examplebucket.id,
 *     key: "someobject",
 *     serverSideEncryption: "aws:kms",
 *     source: new pulumi.asset.FileAsset("index.html"),
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### Server Side Encryption with AWS-Managed Key
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const examplebucket = new aws.s3.Bucket("examplebucket", {
 *     acl: "private",
 * });
 * const examplebucketObject = new aws.s3.BucketObject("examplebucketObject", {
 *     bucket: examplebucket.id,
 *     key: "someobject",
 *     serverSideEncryption: "AES256",
 *     source: new pulumi.asset.FileAsset("index.html"),
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% example %}}
 * ### S3 Object Lock
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const examplebucket = new aws.s3.Bucket("examplebucket", {
 *     acl: "private",
 *     objectLockConfiguration: {
 *         objectLockEnabled: "Enabled",
 *     },
 *     versioning: {
 *         enabled: true,
 *     },
 * });
 * const examplebucketObject = new aws.s3.BucketObject("examplebucketObject", {
 *     bucket: examplebucket.id,
 *     forceDestroy: true,
 *     key: "someobject",
 *     objectLockLegalHoldStatus: "ON",
 *     objectLockMode: "GOVERNANCE",
 *     objectLockRetainUntilDate: "2021-12-31T23:59:60Z",
 *     source: new pulumi.asset.FileAsset("important.txt"),
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_bucket_object.html.markdown.
 */
export class BucketObject extends pulumi.CustomResource {
    /**
     * Get an existing BucketObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketObjectState, opts?: pulumi.CustomResourceOptions): BucketObject {
        return new BucketObject(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/bucketObject:BucketObject';

    /**
     * Returns true if the given object is an instance of BucketObject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketObject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketObject.__pulumiType;
    }

    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * The name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
     */
    public readonly cacheControl!: pulumi.Output<string | undefined>;
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
     */
    public readonly content!: pulumi.Output<string | undefined>;
    /**
     * Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
     */
    public readonly contentBase64!: pulumi.Output<string | undefined>;
    /**
     * Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
     */
    public readonly contentDisposition!: pulumi.Output<string | undefined>;
    /**
     * Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
     */
    public readonly contentEncoding!: pulumi.Output<string | undefined>;
    /**
     * The language the content is in e.g. en-US or en-GB.
     */
    public readonly contentLanguage!: pulumi.Output<string | undefined>;
    /**
     * A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
     * This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Allow the object to be deleted by removing any legal hold on any object version.
     * Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the object once it is in the bucket.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Specifies the AWS KMS Key ARN to use for object encryption.
     * This value is a fully qualified **ARN** of the KMS Key. If using `aws.kms.Key`,
     * use the exported `arn` attribute:
     * `kmsKeyId = "${aws_kms_key.foo.arn}"`
     */
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    /**
     * A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
     */
    public readonly objectLockLegalHoldStatus!: pulumi.Output<string | undefined>;
    /**
     * The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
     */
    public readonly objectLockMode!: pulumi.Output<string | undefined>;
    /**
     * The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
     */
    public readonly objectLockRetainUntilDate!: pulumi.Output<string | undefined>;
    /**
     * Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
     */
    public readonly serverSideEncryption!: pulumi.Output<string>;
    /**
     * The path to a file that will be read and uploaded as raw bytes for the object content.
     */
    public readonly source!: pulumi.Output<pulumi.asset.Asset | pulumi.asset.Archive | undefined>;
    /**
     * Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
     * for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
     */
    public readonly storageClass!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the object.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A unique version ID value for the object, if bucket versioning
     * is enabled.
     */
    public /*out*/ readonly versionId!: pulumi.Output<string>;
    /**
     * Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
     */
    public readonly websiteRedirect!: pulumi.Output<string | undefined>;

    /**
     * Create a BucketObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketObjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketObjectArgs | BucketObjectState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketObjectState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["cacheControl"] = state ? state.cacheControl : undefined;
            inputs["content"] = state ? state.content : undefined;
            inputs["contentBase64"] = state ? state.contentBase64 : undefined;
            inputs["contentDisposition"] = state ? state.contentDisposition : undefined;
            inputs["contentEncoding"] = state ? state.contentEncoding : undefined;
            inputs["contentLanguage"] = state ? state.contentLanguage : undefined;
            inputs["contentType"] = state ? state.contentType : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["key"] = state ? state.key : undefined;
            inputs["kmsKeyId"] = state ? state.kmsKeyId : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["objectLockLegalHoldStatus"] = state ? state.objectLockLegalHoldStatus : undefined;
            inputs["objectLockMode"] = state ? state.objectLockMode : undefined;
            inputs["objectLockRetainUntilDate"] = state ? state.objectLockRetainUntilDate : undefined;
            inputs["serverSideEncryption"] = state ? state.serverSideEncryption : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["storageClass"] = state ? state.storageClass : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["versionId"] = state ? state.versionId : undefined;
            inputs["websiteRedirect"] = state ? state.websiteRedirect : undefined;
        } else {
            const args = argsOrState as BucketObjectArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["acl"] = args ? args.acl : undefined;
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["cacheControl"] = args ? args.cacheControl : undefined;
            inputs["content"] = args ? args.content : undefined;
            inputs["contentBase64"] = args ? args.contentBase64 : undefined;
            inputs["contentDisposition"] = args ? args.contentDisposition : undefined;
            inputs["contentEncoding"] = args ? args.contentEncoding : undefined;
            inputs["contentLanguage"] = args ? args.contentLanguage : undefined;
            inputs["contentType"] = args ? args.contentType : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["key"] = args ? args.key : undefined;
            inputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["objectLockLegalHoldStatus"] = args ? args.objectLockLegalHoldStatus : undefined;
            inputs["objectLockMode"] = args ? args.objectLockMode : undefined;
            inputs["objectLockRetainUntilDate"] = args ? args.objectLockRetainUntilDate : undefined;
            inputs["serverSideEncryption"] = args ? args.serverSideEncryption : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["storageClass"] = args ? args.storageClass : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["websiteRedirect"] = args ? args.websiteRedirect : undefined;
            inputs["versionId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BucketObject.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketObject resources.
 */
export interface BucketObjectState {
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
     */
    readonly acl?: pulumi.Input<string>;
    /**
     * The name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
     */
    readonly bucket?: pulumi.Input<string | Bucket>;
    /**
     * Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
     */
    readonly cacheControl?: pulumi.Input<string>;
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
     */
    readonly contentBase64?: pulumi.Input<string>;
    /**
     * Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * The language the content is in e.g. en-US or en-GB.
     */
    readonly contentLanguage?: pulumi.Input<string>;
    /**
     * A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
     * This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Allow the object to be deleted by removing any legal hold on any object version.
     * Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The name of the object once it is in the bucket.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * Specifies the AWS KMS Key ARN to use for object encryption.
     * This value is a fully qualified **ARN** of the KMS Key. If using `aws.kms.Key`,
     * use the exported `arn` attribute:
     * `kmsKeyId = "${aws_kms_key.foo.arn}"`
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
     */
    readonly objectLockLegalHoldStatus?: pulumi.Input<string>;
    /**
     * The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
     */
    readonly objectLockMode?: pulumi.Input<string>;
    /**
     * The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
     */
    readonly objectLockRetainUntilDate?: pulumi.Input<string>;
    /**
     * Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
     */
    readonly serverSideEncryption?: pulumi.Input<string>;
    /**
     * The path to a file that will be read and uploaded as raw bytes for the object content.
     */
    readonly source?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
     * for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
     */
    readonly storageClass?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * A unique version ID value for the object, if bucket versioning
     * is enabled.
     */
    readonly versionId?: pulumi.Input<string>;
    /**
     * Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
     */
    readonly websiteRedirect?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketObject resource.
 */
export interface BucketObjectArgs {
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
     */
    readonly acl?: pulumi.Input<string>;
    /**
     * The name of the bucket to put the file in. Alternatively, an [S3 access point](https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) ARN can be specified.
     */
    readonly bucket: pulumi.Input<string | Bucket>;
    /**
     * Specifies caching behavior along the request/reply chain Read [w3c cacheControl](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
     */
    readonly cacheControl?: pulumi.Input<string>;
    /**
     * Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
     */
    readonly content?: pulumi.Input<string>;
    /**
     * Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
     */
    readonly contentBase64?: pulumi.Input<string>;
    /**
     * Specifies presentational information for the object. Read [w3c contentDisposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
     */
    readonly contentDisposition?: pulumi.Input<string>;
    /**
     * Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
     */
    readonly contentEncoding?: pulumi.Input<string>;
    /**
     * The language the content is in e.g. en-US or en-GB.
     */
    readonly contentLanguage?: pulumi.Input<string>;
    /**
     * A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
     */
    readonly contentType?: pulumi.Input<string>;
    /**
     * Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (this provider 0.11.12 or later) or `${md5(file("path/to/file"))}` (this provider 0.11.11 or earlier).
     * This attribute is not compatible with KMS encryption, `kmsKeyId` or `serverSideEncryption = "aws:kms"`.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Allow the object to be deleted by removing any legal hold on any object version.
     * Default is `false`. This value should be set to `true` only if the bucket has S3 object lock enabled.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The name of the object once it is in the bucket.
     */
    readonly key?: pulumi.Input<string>;
    /**
     * Specifies the AWS KMS Key ARN to use for object encryption.
     * This value is a fully qualified **ARN** of the KMS Key. If using `aws.kms.Key`,
     * use the exported `arn` attribute:
     * `kmsKeyId = "${aws_kms_key.foo.arn}"`
     */
    readonly kmsKeyId?: pulumi.Input<string>;
    /**
     * A mapping of keys/values to provision metadata (will be automatically prefixed by `x-amz-meta-`, note that only lowercase label are currently supported by the AWS Go API).
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [legal hold](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-legal-holds) status that you want to apply to the specified object. Valid values are `ON` and `OFF`.
     */
    readonly objectLockLegalHoldStatus?: pulumi.Input<string>;
    /**
     * The object lock [retention mode](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-modes) that you want to apply to this object. Valid values are `GOVERNANCE` and `COMPLIANCE`.
     */
    readonly objectLockMode?: pulumi.Input<string>;
    /**
     * The date and time, in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8), when this object's object lock will [expire](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html#object-lock-retention-periods).
     */
    readonly objectLockRetainUntilDate?: pulumi.Input<string>;
    /**
     * Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
     */
    readonly serverSideEncryption?: pulumi.Input<string>;
    /**
     * The path to a file that will be read and uploaded as raw bytes for the object content.
     */
    readonly source?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
     * for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
     */
    readonly storageClass?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the object.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
     */
    readonly websiteRedirect?: pulumi.Input<string>;
}
