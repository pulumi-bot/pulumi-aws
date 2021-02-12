// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a Signer Signing Job.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testSp = new aws.signer.SigningProfile("testSp", {platformId: "AWSLambda-SHA384-ECDSA"});
 * const buildSigningJob = new aws.signer.SigningJob("buildSigningJob", {
 *     profileName: testSp.name,
 *     source: {
 *         s3: {
 *             bucket: "s3-bucket-name",
 *             key: "object-to-be-signed.zip",
 *             version: "jADjFYYYEXAMPLETszPjOmCMFDzd9dN1",
 *         },
 *     },
 *     destination: {
 *         s3: {
 *             bucket: "s3-bucket-name",
 *             prefix: "signed/",
 *         },
 *     },
 *     ignoreSigningJobFailure: true,
 * });
 * ```
 *
 * ## Import
 *
 * Signer signing jobs can be imported using the `job_id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:signer/signingJob:SigningJob test_signer_signing_job 9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee
 * ```
 */
export class SigningJob extends pulumi.CustomResource {
    /**
     * Get an existing SigningJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SigningJobState, opts?: pulumi.CustomResourceOptions): SigningJob {
        return new SigningJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:signer/signingJob:SigningJob';

    /**
     * Returns true if the given object is an instance of SigningJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SigningJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SigningJob.__pulumiType;
    }

    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was completed.
     */
    public /*out*/ readonly completedAt!: pulumi.Output<string>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The S3 bucket in which to save your signed object. See Destination below for details.
     */
    public readonly destination!: pulumi.Output<outputs.signer.SigningJobDestination>;
    /**
     * Set this argument to `true` to ignore signing job failures and retrieve failed status and reason. Default `false`.
     */
    public readonly ignoreSigningJobFailure!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the signing job on output.
     */
    public /*out*/ readonly jobId!: pulumi.Output<string>;
    /**
     * The IAM entity that initiated the signing job.
     */
    public /*out*/ readonly jobInvoker!: pulumi.Output<string>;
    /**
     * The AWS account ID of the job owner.
     */
    public /*out*/ readonly jobOwner!: pulumi.Output<string>;
    /**
     * A human-readable name for the signing platform associated with the signing job.
     */
    public /*out*/ readonly platformDisplayName!: pulumi.Output<string>;
    /**
     * The platform to which your signed code image will be distributed.
     */
    public /*out*/ readonly platformId!: pulumi.Output<string>;
    /**
     * The name of the profile to initiate the signing operation.
     */
    public readonly profileName!: pulumi.Output<string>;
    /**
     * The version of the signing profile used to initiate the signing job.
     */
    public /*out*/ readonly profileVersion!: pulumi.Output<string>;
    /**
     * The IAM principal that requested the signing job.
     */
    public /*out*/ readonly requestedBy!: pulumi.Output<string>;
    /**
     * A revocation record if the signature generated by the signing job has been revoked. Contains a timestamp and the ID of the IAM entity that revoked the signature.
     */
    public /*out*/ readonly revocationRecords!: pulumi.Output<outputs.signer.SigningJobRevocationRecord[]>;
    /**
     * The time when the signature of a signing job expires.
     */
    public /*out*/ readonly signatureExpiresAt!: pulumi.Output<string>;
    /**
     * Name of the S3 bucket where the signed code image is saved by code signing.
     */
    public /*out*/ readonly signedObjects!: pulumi.Output<outputs.signer.SigningJobSignedObject[]>;
    /**
     * The S3 bucket that contains the object to sign. See Source below for details.
     */
    public readonly source!: pulumi.Output<outputs.signer.SigningJobSource>;
    /**
     * Status of the signing job.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * String value that contains the status reason.
     */
    public /*out*/ readonly statusReason!: pulumi.Output<string>;

    /**
     * Create a SigningJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SigningJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SigningJobArgs | SigningJobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SigningJobState | undefined;
            inputs["completedAt"] = state ? state.completedAt : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["ignoreSigningJobFailure"] = state ? state.ignoreSigningJobFailure : undefined;
            inputs["jobId"] = state ? state.jobId : undefined;
            inputs["jobInvoker"] = state ? state.jobInvoker : undefined;
            inputs["jobOwner"] = state ? state.jobOwner : undefined;
            inputs["platformDisplayName"] = state ? state.platformDisplayName : undefined;
            inputs["platformId"] = state ? state.platformId : undefined;
            inputs["profileName"] = state ? state.profileName : undefined;
            inputs["profileVersion"] = state ? state.profileVersion : undefined;
            inputs["requestedBy"] = state ? state.requestedBy : undefined;
            inputs["revocationRecords"] = state ? state.revocationRecords : undefined;
            inputs["signatureExpiresAt"] = state ? state.signatureExpiresAt : undefined;
            inputs["signedObjects"] = state ? state.signedObjects : undefined;
            inputs["source"] = state ? state.source : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["statusReason"] = state ? state.statusReason : undefined;
        } else {
            const args = argsOrState as SigningJobArgs | undefined;
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.profileName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'profileName'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            inputs["destination"] = args ? args.destination : undefined;
            inputs["ignoreSigningJobFailure"] = args ? args.ignoreSigningJobFailure : undefined;
            inputs["profileName"] = args ? args.profileName : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["completedAt"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["jobId"] = undefined /*out*/;
            inputs["jobInvoker"] = undefined /*out*/;
            inputs["jobOwner"] = undefined /*out*/;
            inputs["platformDisplayName"] = undefined /*out*/;
            inputs["platformId"] = undefined /*out*/;
            inputs["profileVersion"] = undefined /*out*/;
            inputs["requestedBy"] = undefined /*out*/;
            inputs["revocationRecords"] = undefined /*out*/;
            inputs["signatureExpiresAt"] = undefined /*out*/;
            inputs["signedObjects"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusReason"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SigningJob.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SigningJob resources.
 */
export interface SigningJobState {
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was completed.
     */
    readonly completedAt?: pulumi.Input<string>;
    /**
     * Date and time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) that the signing job was created.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * The S3 bucket in which to save your signed object. See Destination below for details.
     */
    readonly destination?: pulumi.Input<inputs.signer.SigningJobDestination>;
    /**
     * Set this argument to `true` to ignore signing job failures and retrieve failed status and reason. Default `false`.
     */
    readonly ignoreSigningJobFailure?: pulumi.Input<boolean>;
    /**
     * The ID of the signing job on output.
     */
    readonly jobId?: pulumi.Input<string>;
    /**
     * The IAM entity that initiated the signing job.
     */
    readonly jobInvoker?: pulumi.Input<string>;
    /**
     * The AWS account ID of the job owner.
     */
    readonly jobOwner?: pulumi.Input<string>;
    /**
     * A human-readable name for the signing platform associated with the signing job.
     */
    readonly platformDisplayName?: pulumi.Input<string>;
    /**
     * The platform to which your signed code image will be distributed.
     */
    readonly platformId?: pulumi.Input<string>;
    /**
     * The name of the profile to initiate the signing operation.
     */
    readonly profileName?: pulumi.Input<string>;
    /**
     * The version of the signing profile used to initiate the signing job.
     */
    readonly profileVersion?: pulumi.Input<string>;
    /**
     * The IAM principal that requested the signing job.
     */
    readonly requestedBy?: pulumi.Input<string>;
    /**
     * A revocation record if the signature generated by the signing job has been revoked. Contains a timestamp and the ID of the IAM entity that revoked the signature.
     */
    readonly revocationRecords?: pulumi.Input<pulumi.Input<inputs.signer.SigningJobRevocationRecord>[]>;
    /**
     * The time when the signature of a signing job expires.
     */
    readonly signatureExpiresAt?: pulumi.Input<string>;
    /**
     * Name of the S3 bucket where the signed code image is saved by code signing.
     */
    readonly signedObjects?: pulumi.Input<pulumi.Input<inputs.signer.SigningJobSignedObject>[]>;
    /**
     * The S3 bucket that contains the object to sign. See Source below for details.
     */
    readonly source?: pulumi.Input<inputs.signer.SigningJobSource>;
    /**
     * Status of the signing job.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * String value that contains the status reason.
     */
    readonly statusReason?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SigningJob resource.
 */
export interface SigningJobArgs {
    /**
     * The S3 bucket in which to save your signed object. See Destination below for details.
     */
    readonly destination: pulumi.Input<inputs.signer.SigningJobDestination>;
    /**
     * Set this argument to `true` to ignore signing job failures and retrieve failed status and reason. Default `false`.
     */
    readonly ignoreSigningJobFailure?: pulumi.Input<boolean>;
    /**
     * The name of the profile to initiate the signing operation.
     */
    readonly profileName: pulumi.Input<string>;
    /**
     * The S3 bucket that contains the object to sign. See Source below for details.
     */
    readonly source: pulumi.Input<inputs.signer.SigningJobSource>;
}
