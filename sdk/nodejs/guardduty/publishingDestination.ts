// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a GuardDuty PublishingDestination. Requires an existing GuardDuty Detector.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const currentCallerIdentity = aws.getCallerIdentity({});
 * const currentRegion = aws.getRegion({});
 * const gdBucket = new aws.s3.Bucket("gdBucket", {
 *     acl: "private",
 *     forceDestroy: true,
 * });
 * const bucketPol = pulumi.all([gdBucket.arn, gdBucket.arn]).apply(([gdBucketArn, gdBucketArn1]) => aws.iam.getPolicyDocument({
 *     statements: [
 *         {
 *             sid: "Allow PutObject",
 *             actions: ["s3:PutObject"],
 *             resources: [`${gdBucketArn}/*`],
 *             principals: [{
 *                 type: "Service",
 *                 identifiers: ["guardduty.amazonaws.com"],
 *             }],
 *         },
 *         {
 *             sid: "Allow GetBucketLocation",
 *             actions: ["s3:GetBucketLocation"],
 *             resources: [gdBucketArn1],
 *             principals: [{
 *                 type: "Service",
 *                 identifiers: ["guardduty.amazonaws.com"],
 *             }],
 *         },
 *     ],
 * }));
 * const kmsPol = Promise.all([currentRegion, currentCallerIdentity, currentRegion, currentCallerIdentity, currentCallerIdentity]).then(([currentRegion, currentCallerIdentity, currentRegion1, currentCallerIdentity1, currentCallerIdentity2]) => aws.iam.getPolicyDocument({
 *     statements: [
 *         {
 *             sid: "Allow GuardDuty to encrypt findings",
 *             actions: ["kms:GenerateDataKey"],
 *             resources: [`arn:aws:kms:${currentRegion.name}:${currentCallerIdentity.accountId}:key/*`],
 *             principals: [{
 *                 type: "Service",
 *                 identifiers: ["guardduty.amazonaws.com"],
 *             }],
 *         },
 *         {
 *             sid: "Allow all users to modify/delete key (test only)",
 *             actions: ["kms:*"],
 *             resources: [`arn:aws:kms:${currentRegion1.name}:${currentCallerIdentity1.accountId}:key/*`],
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: [`arn:aws:iam::${currentCallerIdentity2.accountId}:root`],
 *             }],
 *         },
 *     ],
 * }));
 * const testGd = new aws.guardduty.Detector("testGd", {enable: true});
 * const gdBucketPolicy = new aws.s3.BucketPolicy("gdBucketPolicy", {
 *     bucket: gdBucket.id,
 *     policy: bucketPol.json,
 * });
 * const gdKey = new aws.kms.Key("gdKey", {
 *     description: "Temporary key for AccTest of TF",
 *     deletionWindowInDays: 7,
 *     policy: kmsPol.then(kmsPol => kmsPol.json),
 * });
 * const test = new aws.guardduty.PublishingDestination("test", {
 *     detectorId: testGd.id,
 *     destinationArn: gdBucket.arn,
 *     kmsKeyArn: gdKey.arn,
 * }, {
 *     dependsOn: [gdBucketPolicy],
 * });
 * ```
 *
 * > **Note:** Please do not use this simple example for Bucket-Policy and KMS Key Policy in a production environment. It is much too open for such a use-case. Refer to the AWS documentation here: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_exportfindings.html
 *
 * ## Import
 *
 * GuardDuty PublishingDestination can be imported using the the master GuardDuty detector ID and PublishingDestinationID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:guardduty/publishingDestination:PublishingDestination test a4b86f26fa42e7e7cf0d1c333ea77777:a4b86f27a0e464e4a7e0516d242f1234
 * ```
 */
export class PublishingDestination extends pulumi.CustomResource {
    /**
     * Get an existing PublishingDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublishingDestinationState, opts?: pulumi.CustomResourceOptions): PublishingDestination {
        return new PublishingDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:guardduty/publishingDestination:PublishingDestination';

    /**
     * Returns true if the given object is an instance of PublishingDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublishingDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublishingDestination.__pulumiType;
    }

    /**
     * The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
     */
    public readonly destinationArn!: pulumi.Output<string>;
    /**
     * Currently there is only "S3" available as destination type which is also the default value
     */
    public readonly destinationType!: pulumi.Output<string | undefined>;
    /**
     * The detector ID of the GuardDuty.
     */
    public readonly detectorId!: pulumi.Output<string>;
    /**
     * The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
     */
    public readonly kmsKeyArn!: pulumi.Output<string>;

    /**
     * Create a PublishingDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublishingDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublishingDestinationArgs | PublishingDestinationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PublishingDestinationState | undefined;
            inputs["destinationArn"] = state ? state.destinationArn : undefined;
            inputs["destinationType"] = state ? state.destinationType : undefined;
            inputs["detectorId"] = state ? state.detectorId : undefined;
            inputs["kmsKeyArn"] = state ? state.kmsKeyArn : undefined;
        } else {
            const args = argsOrState as PublishingDestinationArgs | undefined;
            if (!args || args.destinationArn === undefined) {
                throw new Error("Missing required property 'destinationArn'");
            }
            if (!args || args.detectorId === undefined) {
                throw new Error("Missing required property 'detectorId'");
            }
            if (!args || args.kmsKeyArn === undefined) {
                throw new Error("Missing required property 'kmsKeyArn'");
            }
            inputs["destinationArn"] = args ? args.destinationArn : undefined;
            inputs["destinationType"] = args ? args.destinationType : undefined;
            inputs["detectorId"] = args ? args.detectorId : undefined;
            inputs["kmsKeyArn"] = args ? args.kmsKeyArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PublishingDestination.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublishingDestination resources.
 */
export interface PublishingDestinationState {
    /**
     * The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
     */
    readonly destinationArn?: pulumi.Input<string>;
    /**
     * Currently there is only "S3" available as destination type which is also the default value
     */
    readonly destinationType?: pulumi.Input<string>;
    /**
     * The detector ID of the GuardDuty.
     */
    readonly detectorId?: pulumi.Input<string>;
    /**
     * The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
     */
    readonly kmsKeyArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PublishingDestination resource.
 */
export interface PublishingDestinationArgs {
    /**
     * The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
     */
    readonly destinationArn: pulumi.Input<string>;
    /**
     * Currently there is only "S3" available as destination type which is also the default value
     */
    readonly destinationType?: pulumi.Input<string>;
    /**
     * The detector ID of the GuardDuty.
     */
    readonly detectorId: pulumi.Input<string>;
    /**
     * The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
     */
    readonly kmsKeyArn: pulumi.Input<string>;
}
