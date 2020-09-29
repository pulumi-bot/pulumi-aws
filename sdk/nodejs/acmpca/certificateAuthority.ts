// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage AWS Certificate Manager Private Certificate Authorities (ACM PCA Certificate Authorities).
 *
 * > **NOTE:** Creating this resource will leave the certificate authority in a `PENDING_CERTIFICATE` status, which means it cannot yet issue certificates. To complete this setup, you must fully sign the certificate authority CSR available in the `certificateSigningRequest` attribute and import the signed certificate using the AWS SDK, CLI or Console. This provider can support another resource to manage that workflow automatically in the future.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.acmpca.CertificateAuthority("example", {
 *     certificateAuthorityConfiguration: {
 *         keyAlgorithm: "RSA_4096",
 *         signingAlgorithm: "SHA512WITHRSA",
 *         subject: {
 *             commonName: "example.com",
 *         },
 *     },
 *     permanentDeletionTimeInDays: 7,
 * });
 * ```
 * ### Enable Certificate Revocation List
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleBucket = new aws.s3.Bucket("exampleBucket", {});
 * const acmpcaBucketAccess = pulumi.all([exampleBucket.arn, exampleBucket.arn]).apply(([exampleBucketArn, exampleBucketArn1]) => aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: [
 *             "s3:GetBucketAcl",
 *             "s3:GetBucketLocation",
 *             "s3:PutObject",
 *             "s3:PutObjectAcl",
 *         ],
 *         resources: [
 *             exampleBucketArn,
 *             `${exampleBucketArn1}/*`,
 *         ],
 *         principals: [{
 *             identifiers: ["acm-pca.amazonaws.com"],
 *             type: "Service",
 *         }],
 *     }],
 * }));
 * const exampleBucketPolicy = new aws.s3.BucketPolicy("exampleBucketPolicy", {
 *     bucket: exampleBucket.id,
 *     policy: acmpcaBucketAccess.json,
 * });
 * const exampleCertificateAuthority = new aws.acmpca.CertificateAuthority("exampleCertificateAuthority", {
 *     certificateAuthorityConfiguration: {
 *         keyAlgorithm: "RSA_4096",
 *         signingAlgorithm: "SHA512WITHRSA",
 *         subject: {
 *             commonName: "example.com",
 *         },
 *     },
 *     revocationConfiguration: {
 *         crlConfiguration: {
 *             customCname: "crl.example.com",
 *             enabled: true,
 *             expirationInDays: 7,
 *             s3BucketName: exampleBucket.id,
 *         },
 *     },
 * }, {
 *     dependsOn: [exampleBucketPolicy],
 * });
 * ```
 */
export class CertificateAuthority extends pulumi.CustomResource {
    /**
     * Get an existing CertificateAuthority resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CertificateAuthorityState, opts?: pulumi.CustomResourceOptions): CertificateAuthority {
        return new CertificateAuthority(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:acmpca/certificateAuthority:CertificateAuthority';

    /**
     * Returns true if the given object is an instance of CertificateAuthority.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateAuthority {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateAuthority.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the certificate authority.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly certificate!: pulumi.Output<string>;
    /**
     * Nested argument containing algorithms and certificate subject information. Defined below.
     */
    public readonly certificateAuthorityConfiguration!: pulumi.Output<outputs.acmpca.CertificateAuthorityCertificateAuthorityConfiguration>;
    /**
     * Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly certificateChain!: pulumi.Output<string>;
    /**
     * The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
     */
    public /*out*/ readonly certificateSigningRequest!: pulumi.Output<string>;
    /**
     * Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly notAfter!: pulumi.Output<string>;
    /**
     * Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly notBefore!: pulumi.Output<string>;
    /**
     * The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
     */
    public readonly permanentDeletionTimeInDays!: pulumi.Output<number | undefined>;
    /**
     * Nested argument containing revocation configuration. Defined below.
     */
    public readonly revocationConfiguration!: pulumi.Output<outputs.acmpca.CertificateAuthorityRevocationConfiguration | undefined>;
    /**
     * Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
     */
    public /*out*/ readonly serial!: pulumi.Output<string>;
    /**
     * Status of the certificate authority.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the certificate authority.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a CertificateAuthority resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateAuthorityArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CertificateAuthorityArgs | CertificateAuthorityState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CertificateAuthorityState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["certificateAuthorityConfiguration"] = state ? state.certificateAuthorityConfiguration : undefined;
            inputs["certificateChain"] = state ? state.certificateChain : undefined;
            inputs["certificateSigningRequest"] = state ? state.certificateSigningRequest : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["notAfter"] = state ? state.notAfter : undefined;
            inputs["notBefore"] = state ? state.notBefore : undefined;
            inputs["permanentDeletionTimeInDays"] = state ? state.permanentDeletionTimeInDays : undefined;
            inputs["revocationConfiguration"] = state ? state.revocationConfiguration : undefined;
            inputs["serial"] = state ? state.serial : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as CertificateAuthorityArgs | undefined;
            if (!args || args.certificateAuthorityConfiguration === undefined) {
                throw new Error("Missing required property 'certificateAuthorityConfiguration'");
            }
            inputs["certificateAuthorityConfiguration"] = args ? args.certificateAuthorityConfiguration : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["permanentDeletionTimeInDays"] = args ? args.permanentDeletionTimeInDays : undefined;
            inputs["revocationConfiguration"] = args ? args.revocationConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["certificate"] = undefined /*out*/;
            inputs["certificateChain"] = undefined /*out*/;
            inputs["certificateSigningRequest"] = undefined /*out*/;
            inputs["notAfter"] = undefined /*out*/;
            inputs["notBefore"] = undefined /*out*/;
            inputs["serial"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CertificateAuthority.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CertificateAuthority resources.
 */
export interface CertificateAuthorityState {
    /**
     * Amazon Resource Name (ARN) of the certificate authority.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * Nested argument containing algorithms and certificate subject information. Defined below.
     */
    readonly certificateAuthorityConfiguration?: pulumi.Input<inputs.acmpca.CertificateAuthorityCertificateAuthorityConfiguration>;
    /**
     * Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
     */
    readonly certificateChain?: pulumi.Input<string>;
    /**
     * The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
     */
    readonly certificateSigningRequest?: pulumi.Input<string>;
    /**
     * Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    readonly notAfter?: pulumi.Input<string>;
    /**
     * Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    readonly notBefore?: pulumi.Input<string>;
    /**
     * The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
     */
    readonly permanentDeletionTimeInDays?: pulumi.Input<number>;
    /**
     * Nested argument containing revocation configuration. Defined below.
     */
    readonly revocationConfiguration?: pulumi.Input<inputs.acmpca.CertificateAuthorityRevocationConfiguration>;
    /**
     * Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
     */
    readonly serial?: pulumi.Input<string>;
    /**
     * Status of the certificate authority.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the certificate authority.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CertificateAuthority resource.
 */
export interface CertificateAuthorityArgs {
    /**
     * Nested argument containing algorithms and certificate subject information. Defined below.
     */
    readonly certificateAuthorityConfiguration: pulumi.Input<inputs.acmpca.CertificateAuthorityCertificateAuthorityConfiguration>;
    /**
     * Boolean value that specifies whether certificate revocation lists (CRLs) are enabled. Defaults to `false`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The number of days to make a CA restorable after it has been deleted, must be between 7 to 30 days, with default to 30 days.
     */
    readonly permanentDeletionTimeInDays?: pulumi.Input<number>;
    /**
     * Nested argument containing revocation configuration. Defined below.
     */
    readonly revocationConfiguration?: pulumi.Input<inputs.acmpca.CertificateAuthorityRevocationConfiguration>;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the certificate authority.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the certificate authority. Defaults to `SUBORDINATE`. Valid values: `ROOT` and `SUBORDINATE`.
     */
    readonly type?: pulumi.Input<string>;
}
