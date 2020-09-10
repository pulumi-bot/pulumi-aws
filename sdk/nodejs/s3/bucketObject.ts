// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Bucket} from "./index";

export class BucketObject extends pulumi.CustomResource {
    /**
     * Get an existing BucketObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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

    public readonly acl!: pulumi.Output<string | undefined>;
    public readonly bucket!: pulumi.Output<string>;
    public readonly cacheControl!: pulumi.Output<string | undefined>;
    public readonly content!: pulumi.Output<string | undefined>;
    public readonly contentBase64!: pulumi.Output<string | undefined>;
    public readonly contentDisposition!: pulumi.Output<string | undefined>;
    public readonly contentEncoding!: pulumi.Output<string | undefined>;
    public readonly contentLanguage!: pulumi.Output<string | undefined>;
    public readonly contentType!: pulumi.Output<string>;
    public readonly etag!: pulumi.Output<string>;
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    public readonly key!: pulumi.Output<string>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly objectLockLegalHoldStatus!: pulumi.Output<string | undefined>;
    public readonly objectLockMode!: pulumi.Output<string | undefined>;
    public readonly objectLockRetainUntilDate!: pulumi.Output<string | undefined>;
    public readonly serverSideEncryption!: pulumi.Output<string>;
    public readonly source!: pulumi.Output<pulumi.asset.Asset | pulumi.asset.Archive | undefined>;
    public readonly storageClass!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly versionId!: pulumi.Output<string>;
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
    readonly acl?: pulumi.Input<string>;
    readonly bucket?: pulumi.Input<string | Bucket>;
    readonly cacheControl?: pulumi.Input<string>;
    readonly content?: pulumi.Input<string>;
    readonly contentBase64?: pulumi.Input<string>;
    readonly contentDisposition?: pulumi.Input<string>;
    readonly contentEncoding?: pulumi.Input<string>;
    readonly contentLanguage?: pulumi.Input<string>;
    readonly contentType?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    readonly forceDestroy?: pulumi.Input<boolean>;
    readonly key?: pulumi.Input<string>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly objectLockLegalHoldStatus?: pulumi.Input<string>;
    readonly objectLockMode?: pulumi.Input<string>;
    readonly objectLockRetainUntilDate?: pulumi.Input<string>;
    readonly serverSideEncryption?: pulumi.Input<string>;
    readonly source?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    readonly storageClass?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly versionId?: pulumi.Input<string>;
    readonly websiteRedirect?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketObject resource.
 */
export interface BucketObjectArgs {
    readonly acl?: pulumi.Input<string>;
    readonly bucket: pulumi.Input<string | Bucket>;
    readonly cacheControl?: pulumi.Input<string>;
    readonly content?: pulumi.Input<string>;
    readonly contentBase64?: pulumi.Input<string>;
    readonly contentDisposition?: pulumi.Input<string>;
    readonly contentEncoding?: pulumi.Input<string>;
    readonly contentLanguage?: pulumi.Input<string>;
    readonly contentType?: pulumi.Input<string>;
    readonly etag?: pulumi.Input<string>;
    readonly forceDestroy?: pulumi.Input<boolean>;
    readonly key?: pulumi.Input<string>;
    readonly kmsKeyId?: pulumi.Input<string>;
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly objectLockLegalHoldStatus?: pulumi.Input<string>;
    readonly objectLockMode?: pulumi.Input<string>;
    readonly objectLockRetainUntilDate?: pulumi.Input<string>;
    readonly serverSideEncryption?: pulumi.Input<string>;
    readonly source?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    readonly storageClass?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly websiteRedirect?: pulumi.Input<string>;
}
