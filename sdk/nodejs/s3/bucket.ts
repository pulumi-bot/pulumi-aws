// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

import {PolicyDocument} from "../iam/documents";
import {CannedAcl} from "./cannedAcl";
import {RoutingRule} from "./routingRules";

/**
 * Provides a S3 bucket resource.
 * 
 * ## Example Usage
 * 
 * ### Private Bucket w/ Tags
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const bucket = new aws.s3.Bucket("b", {
 *     acl: "private",
 *     tags: {
 *         Environment: "Dev",
 *         Name: "My bucket",
 *     },
 * });
 * ```
 * 
 * ### Static Website Hosting
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as fs from "fs";
 * 
 * const bucket = new aws.s3.Bucket("b", {
 *     acl: "public-read",
 *     policy: fs.readFileSync("policy.json", "utf-8"),
 *     website: {
 *         errorDocument: "error.html",
 *         indexDocument: "index.html",
 *         routingRules: `[{
 *     "Condition": {
 *         "KeyPrefixEquals": "docs/"
 *     },
 *     "Redirect": {
 *         "ReplaceKeyPrefixWith": "documents/"
 *     }
 * }]
 * `,
 *     },
 * });
 * ```
 * 
 * ### Using CORS
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const bucket = new aws.s3.Bucket("b", {
 *     acl: "public-read",
 *     corsRules: [{
 *         allowedHeaders: ["*"],
 *         allowedMethods: [
 *             "PUT",
 *             "POST",
 *         ],
 *         allowedOrigins: ["https://s3-website-test.mydomain.com"],
 *         exposeHeaders: ["ETag"],
 *         maxAgeSeconds: 3000,
 *     }],
 * });
 * ```
 * 
 * ### Using versioning
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const bucket = new aws.s3.Bucket("b", {
 *     acl: "private",
 *     versioning: {
 *         enabled: true,
 *     },
 * });
 * ```
 * 
 * ### Enable Logging
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const logBucket = new aws.s3.Bucket("logBucket", {
 *     acl: "log-delivery-write",
 * });
 * const bucket = new aws.s3.Bucket("b", {
 *     acl: "private",
 *     loggings: [{
 *         targetBucket: logBucket.id,
 *         targetPrefix: "log/",
 *     }],
 * });
 * ```
 * 
 * ### Using object lifecycle
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 *     lifecycleRules: [
 *         {
 *             enabled: true,
 *             expiration: {
 *                 days: 90,
 *             },
 *             id: "log",
 *             prefix: "log/",
 *             tags: {
 *                 autoclean: "true",
 *                 rule: "log",
 *             },
 *             transitions: [
 *                 {
 *                     days: 30,
 *                     storageClass: "STANDARD_IA", // or "ONEZONE_IA"
 *                 },
 *                 {
 *                     days: 60,
 *                     storageClass: "GLACIER",
 *                 },
 *             ],
 *         },
 *         {
 *             enabled: true,
 *             expiration: {
 *                 date: "2016-01-12",
 *             },
 *             id: "tmp",
 *             prefix: "tmp/",
 *         },
 *     ],
 * });
 * const versioningBucket = new aws.s3.Bucket("versioningBucket", {
 *     acl: "private",
 *     lifecycleRules: [{
 *         enabled: true,
 *         noncurrentVersionExpiration: {
 *             days: 90,
 *         },
 *         noncurrentVersionTransitions: [
 *             {
 *                 days: 30,
 *                 storageClass: "STANDARD_IA",
 *             },
 *             {
 *                 days: 60,
 *                 storageClass: "GLACIER",
 *             },
 *         ],
 *         prefix: "config/",
 *     }],
 *     versioning: {
 *         enabled: true,
 *     },
 * });
 * ```
 * 
 * ### Using replication configuration
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const central = new aws.Provider("central", {
 *     region: "eu-central-1",
 * });
 * const replicationRole = new aws.iam.Role("replication", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "s3.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const destination = new aws.s3.Bucket("destination", {
 *     region: "eu-west-1",
 *     versioning: {
 *         enabled: true,
 *     },
 * });
 * const bucket = new aws.s3.Bucket("bucket", {
 *     acl: "private",
 *     region: "eu-central-1",
 *     replicationConfiguration: {
 *         role: replicationRole.arn,
 *         rules: [{
 *             destination: {
 *                 bucket: destination.arn,
 *                 storageClass: "STANDARD",
 *             },
 *             id: "foobar",
 *             prefix: "foo",
 *             status: "Enabled",
 *         }],
 *     },
 *     versioning: {
 *         enabled: true,
 *     },
 * }, {provider: central});
 * const replicationPolicy = new aws.iam.Policy("replication", {
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "s3:GetReplicationConfiguration",
 *         "s3:ListBucket"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": [
 *         "${bucket.arn}"
 *       ]
 *     },
 *     {
 *       "Action": [
 *         "s3:GetObjectVersion",
 *         "s3:GetObjectVersionAcl"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": [
 *         "${bucket.arn}/*"
 *       ]
 *     },
 *     {
 *       "Action": [
 *         "s3:ReplicateObject",
 *         "s3:ReplicateDelete"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "${destination.arn}/*"
 *     }
 *   ]
 * }
 * `,
 * });
 * const replicationPolicyAttachment = new aws.iam.PolicyAttachment("replication", {
 *     policyArn: replicationPolicy.arn,
 *     roles: [replicationRole.name],
 * });
 * ```
 * 
 * ### Enable Default Server Side Encryption
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const mykey = new aws.kms.Key("mykey", {
 *     deletionWindowInDays: 10,
 *     description: "This key is used to encrypt bucket objects",
 * });
 * const mybucket = new aws.s3.Bucket("mybucket", {
 *     serverSideEncryptionConfiguration: {
 *         rule: {
 *             applyServerSideEncryptionByDefault: {
 *                 kmsMasterKeyId: mykey.arn,
 *                 sseAlgorithm: "aws:kms",
 *             },
 *         },
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/s3_bucket.html.markdown.
 */
export class Bucket extends pulumi.CustomResource {
    /**
     * Get an existing Bucket resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketState, opts?: pulumi.CustomResourceOptions): Bucket {
        return new Bucket(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/bucket:Bucket';

    /**
     * Returns true if the given object is an instance of Bucket.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bucket {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bucket.__pulumiType;
    }

    /**
     * Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
     */
    public readonly accelerationStatus!: pulumi.Output<string>;
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
     */
    public readonly acl!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    public readonly arn!: pulumi.Output<string>;
    /**
     * The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
     */
    public /*out*/ readonly bucketDomainName!: pulumi.Output<string>;
    /**
     * Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
     */
    public readonly bucketPrefix!: pulumi.Output<string | undefined>;
    /**
     * The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
     */
    public /*out*/ readonly bucketRegionalDomainName!: pulumi.Output<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    public readonly corsRules!: pulumi.Output<outputs.s3.BucketCorsRule[] | undefined>;
    /**
     * A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    public readonly hostedZoneId!: pulumi.Output<string>;
    /**
     * A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
     */
    public readonly lifecycleRules!: pulumi.Output<outputs.s3.BucketLifecycleRule[] | undefined>;
    /**
     * A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
     */
    public readonly loggings!: pulumi.Output<outputs.s3.BucketLogging[] | undefined>;
    /**
     * A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
     */
    public readonly objectLockConfiguration!: pulumi.Output<outputs.s3.BucketObjectLockConfiguration | undefined>;
    /**
     * A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a deployment. In this case, please make sure you use the verbose/specific version of the policy.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
     */
    public readonly replicationConfiguration!: pulumi.Output<outputs.s3.BucketReplicationConfiguration | undefined>;
    /**
     * Specifies who should bear the cost of Amazon S3 data transfer.
     * Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
     * the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
     * developer guide for more information.
     */
    public readonly requestPayer!: pulumi.Output<string>;
    /**
     * A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
     */
    public readonly serverSideEncryptionConfiguration!: pulumi.Output<outputs.s3.BucketServerSideEncryptionConfiguration | undefined>;
    /**
     * A mapping of tags that identifies subset of objects to which the rule applies.
     * The rule applies only to objects having all the tags in its tagset.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    public readonly versioning!: pulumi.Output<outputs.s3.BucketVersioning>;
    /**
     * A website object (documented below).
     */
    public readonly website!: pulumi.Output<outputs.s3.BucketWebsite | undefined>;
    /**
     * The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    public readonly websiteDomain!: pulumi.Output<string>;
    /**
     * The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    public readonly websiteEndpoint!: pulumi.Output<string>;

    /**
     * Create a Bucket resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: BucketArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketArgs | BucketState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketState | undefined;
            inputs["accelerationStatus"] = state ? state.accelerationStatus : undefined;
            inputs["acl"] = state ? state.acl : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["bucketDomainName"] = state ? state.bucketDomainName : undefined;
            inputs["bucketPrefix"] = state ? state.bucketPrefix : undefined;
            inputs["bucketRegionalDomainName"] = state ? state.bucketRegionalDomainName : undefined;
            inputs["corsRules"] = state ? state.corsRules : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["hostedZoneId"] = state ? state.hostedZoneId : undefined;
            inputs["lifecycleRules"] = state ? state.lifecycleRules : undefined;
            inputs["loggings"] = state ? state.loggings : undefined;
            inputs["objectLockConfiguration"] = state ? state.objectLockConfiguration : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["replicationConfiguration"] = state ? state.replicationConfiguration : undefined;
            inputs["requestPayer"] = state ? state.requestPayer : undefined;
            inputs["serverSideEncryptionConfiguration"] = state ? state.serverSideEncryptionConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["versioning"] = state ? state.versioning : undefined;
            inputs["website"] = state ? state.website : undefined;
            inputs["websiteDomain"] = state ? state.websiteDomain : undefined;
            inputs["websiteEndpoint"] = state ? state.websiteEndpoint : undefined;
        } else {
            const args = argsOrState as BucketArgs | undefined;
            inputs["accelerationStatus"] = args ? args.accelerationStatus : undefined;
            inputs["acl"] = args ? args.acl : undefined;
            inputs["arn"] = args ? args.arn : undefined;
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["bucketPrefix"] = args ? args.bucketPrefix : undefined;
            inputs["corsRules"] = args ? args.corsRules : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["hostedZoneId"] = args ? args.hostedZoneId : undefined;
            inputs["lifecycleRules"] = args ? args.lifecycleRules : undefined;
            inputs["loggings"] = args ? args.loggings : undefined;
            inputs["objectLockConfiguration"] = args ? args.objectLockConfiguration : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["replicationConfiguration"] = args ? args.replicationConfiguration : undefined;
            inputs["requestPayer"] = args ? args.requestPayer : undefined;
            inputs["serverSideEncryptionConfiguration"] = args ? args.serverSideEncryptionConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["versioning"] = args ? args.versioning : undefined;
            inputs["website"] = args ? args.website : undefined;
            inputs["websiteDomain"] = args ? args.websiteDomain : undefined;
            inputs["websiteEndpoint"] = args ? args.websiteEndpoint : undefined;
            inputs["bucketDomainName"] = undefined /*out*/;
            inputs["bucketRegionalDomainName"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Bucket.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bucket resources.
 */
export interface BucketState {
    /**
     * Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
     */
    readonly accelerationStatus?: pulumi.Input<string>;
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
     */
    readonly acl?: pulumi.Input<string | CannedAcl>;
    /**
     * The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
     */
    readonly bucketDomainName?: pulumi.Input<string>;
    /**
     * Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
     */
    readonly bucketPrefix?: pulumi.Input<string>;
    /**
     * The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
     */
    readonly bucketRegionalDomainName?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    readonly corsRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketCorsRule>[]>;
    /**
     * A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    readonly hostedZoneId?: pulumi.Input<string>;
    /**
     * A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
     */
    readonly lifecycleRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketLifecycleRule>[]>;
    /**
     * A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
     */
    readonly loggings?: pulumi.Input<pulumi.Input<inputs.s3.BucketLogging>[]>;
    /**
     * A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
     */
    readonly objectLockConfiguration?: pulumi.Input<inputs.s3.BucketObjectLockConfiguration>;
    /**
     * A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a deployment. In this case, please make sure you use the verbose/specific version of the policy.
     */
    readonly policy?: pulumi.Input<string | PolicyDocument>;
    /**
     * If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
     */
    readonly replicationConfiguration?: pulumi.Input<inputs.s3.BucketReplicationConfiguration>;
    /**
     * Specifies who should bear the cost of Amazon S3 data transfer.
     * Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
     * the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
     * developer guide for more information.
     */
    readonly requestPayer?: pulumi.Input<string>;
    /**
     * A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
     */
    readonly serverSideEncryptionConfiguration?: pulumi.Input<inputs.s3.BucketServerSideEncryptionConfiguration>;
    /**
     * A mapping of tags that identifies subset of objects to which the rule applies.
     * The rule applies only to objects having all the tags in its tagset.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    readonly versioning?: pulumi.Input<inputs.s3.BucketVersioning>;
    /**
     * A website object (documented below).
     */
    readonly website?: pulumi.Input<inputs.s3.BucketWebsite>;
    /**
     * The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    readonly websiteDomain?: pulumi.Input<string>;
    /**
     * The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    readonly websiteEndpoint?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Bucket resource.
 */
export interface BucketArgs {
    /**
     * Sets the accelerate configuration of an existing bucket. Can be `Enabled` or `Suspended`.
     */
    readonly accelerationStatus?: pulumi.Input<string>;
    /**
     * The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
     */
    readonly acl?: pulumi.Input<string | CannedAcl>;
    /**
     * The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ARN of the S3 bucket where you want Amazon S3 to store replicas of the object identified by the rule.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * Creates a unique bucket name beginning with the specified prefix. Conflicts with `bucket`.
     */
    readonly bucketPrefix?: pulumi.Input<string>;
    /**
     * A rule of [Cross-Origin Resource Sharing](https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) (documented below).
     */
    readonly corsRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketCorsRule>[]>;
    /**
     * A boolean that indicates all objects (including any [locked objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock-overview.html)) should be deleted from the bucket so that the bucket can be destroyed without error. These objects are *not* recoverable.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
     */
    readonly hostedZoneId?: pulumi.Input<string>;
    /**
     * A configuration of [object lifecycle management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) (documented below).
     */
    readonly lifecycleRules?: pulumi.Input<pulumi.Input<inputs.s3.BucketLifecycleRule>[]>;
    /**
     * A settings of [bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/UG/ManagingBucketLogging.html) (documented below).
     */
    readonly loggings?: pulumi.Input<pulumi.Input<inputs.s3.BucketLogging>[]>;
    /**
     * A configuration of [S3 object locking](https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html) (documented below)
     */
    readonly objectLockConfiguration?: pulumi.Input<inputs.s3.BucketObjectLockConfiguration>;
    /**
     * A valid [bucket policy](https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html) JSON document. Note that if the policy document is not specific enough (but still valid), this provider may view the policy as constantly changing in a deployment. In this case, please make sure you use the verbose/specific version of the policy.
     */
    readonly policy?: pulumi.Input<string | PolicyDocument>;
    /**
     * If specified, the AWS region this bucket should reside in. Otherwise, the region used by the callee.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * A configuration of [replication configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/crr.html) (documented below).
     */
    readonly replicationConfiguration?: pulumi.Input<inputs.s3.BucketReplicationConfiguration>;
    /**
     * Specifies who should bear the cost of Amazon S3 data transfer.
     * Can be either `BucketOwner` or `Requester`. By default, the owner of the S3 bucket would incur
     * the costs of any data transfer. See [Requester Pays Buckets](http://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
     * developer guide for more information.
     */
    readonly requestPayer?: pulumi.Input<string>;
    /**
     * A configuration of [server-side encryption configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) (documented below)
     */
    readonly serverSideEncryptionConfiguration?: pulumi.Input<inputs.s3.BucketServerSideEncryptionConfiguration>;
    /**
     * A mapping of tags that identifies subset of objects to which the rule applies.
     * The rule applies only to objects having all the tags in its tagset.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A state of [versioning](https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html) (documented below)
     */
    readonly versioning?: pulumi.Input<inputs.s3.BucketVersioning>;
    /**
     * A website object (documented below).
     */
    readonly website?: pulumi.Input<inputs.s3.BucketWebsite>;
    /**
     * The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
     */
    readonly websiteDomain?: pulumi.Input<string>;
    /**
     * The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
     */
    readonly websiteEndpoint?: pulumi.Input<string>;
}
