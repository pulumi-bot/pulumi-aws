// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information on a AWS Certificate Manager Private Certificate Authority (ACM PCA Certificate Authority).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = aws.acmpca.getCertificateAuthority({
 *     arn: "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/acmpca_certificate_authority.html.markdown.
 */
export function getCertificateAuthority(args: GetCertificateAuthorityArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateAuthorityResult> & GetCertificateAuthorityResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetCertificateAuthorityResult> = pulumi.runtime.invoke("aws:acmpca/getCertificateAuthority:getCertificateAuthority", {
        "arn": args.arn,
        "revocationConfigurations": args.revocationConfigurations,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getCertificateAuthority.
 */
export interface GetCertificateAuthorityArgs {
    /**
     * Amazon Resource Name (ARN) of the certificate authority.
     */
    readonly arn: string;
    readonly revocationConfigurations?: inputs.acmpca.GetCertificateAuthorityRevocationConfiguration[];
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getCertificateAuthority.
 */
export interface GetCertificateAuthorityResult {
    readonly arn: string;
    /**
     * Base64-encoded certificate authority (CA) certificate. Only available after the certificate authority certificate has been imported.
     */
    readonly certificate: string;
    /**
     * Base64-encoded certificate chain that includes any intermediate certificates and chains up to root on-premises certificate that you used to sign your private CA certificate. The chain does not include your private CA certificate. Only available after the certificate authority certificate has been imported.
     */
    readonly certificateChain: string;
    /**
     * The base64 PEM-encoded certificate signing request (CSR) for your private CA certificate.
     */
    readonly certificateSigningRequest: string;
    /**
     * Date and time after which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    readonly notAfter: string;
    /**
     * Date and time before which the certificate authority is not valid. Only available after the certificate authority certificate has been imported.
     */
    readonly notBefore: string;
    /**
     * Nested attribute containing revocation configuration.
     * * `revocation_configuration.0.crl_configuration` - Nested attribute containing configuration of the certificate revocation list (CRL), if any, maintained by the certificate authority.
     * * `revocation_configuration.0.crl_configuration.0.custom_cname` - Name inserted into the certificate CRL Distribution Points extension that enables the use of an alias for the CRL distribution point.
     * * `revocation_configuration.0.crl_configuration.0.enabled` - Boolean value that specifies whether certificate revocation lists (CRLs) are enabled.
     * * `revocation_configuration.0.crl_configuration.0.expiration_in_days` - Number of days until a certificate expires.
     * * `revocation_configuration.0.crl_configuration.0.s3_bucket_name` - Name of the S3 bucket that contains the CRL.
     */
    readonly revocationConfigurations: outputs.acmpca.GetCertificateAuthorityRevocationConfiguration[];
    /**
     * Serial number of the certificate authority. Only available after the certificate authority certificate has been imported.
     */
    readonly serial: string;
    /**
     * Status of the certificate authority.
     */
    readonly status: string;
    /**
     * Specifies a key-value map of user-defined tags that are attached to the certificate authority.
     */
    readonly tags: {[key: string]: string};
    /**
     * The type of the certificate authority.
     */
    readonly type: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
