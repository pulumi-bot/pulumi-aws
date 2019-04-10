// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ARN of a certificate in AWS Certificate
 * Manager (ACM), you can reference
 * it by domain without having to hard code the ARNs as input.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = pulumi.output(aws.acm.getCertificate({
 *     domain: "tf.example.com",
 *     mostRecent: true,
 *     types: ["AMAZON_ISSUED"],
 * }));
 * ```
 */
export function getCertificate(args: GetCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateResult> {
    return pulumi.runtime.invoke("aws:acm/getCertificate:getCertificate", {
        "domain": args.domain,
        "mostRecent": args.mostRecent,
        "statuses": args.statuses,
        "types": args.types,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificate.
 */
export interface GetCertificateArgs {
    /**
     * The domain of the certificate to look up. If no certificate is found with this name, an error will be returned.
     */
    readonly domain: string;
    /**
     * If set to true, it sorts the certificates matched by previous criteria by the NotBefore field, returning only the most recent one. If set to false, it returns an error if more than one certificate is found. Defaults to false.
     */
    readonly mostRecent?: boolean;
    /**
     * A list of statuses on which to filter the returned list. Valid values are `PENDING_VALIDATION`, `ISSUED`,
     * `INACTIVE`, `EXPIRED`, `VALIDATION_TIMED_OUT`, `REVOKED` and `FAILED`. If no value is specified, only certificates in the `ISSUED` state
     * are returned.
     */
    readonly statuses?: string[];
    /**
     * A list of types on which to filter the returned list. Valid values are `AMAZON_ISSUED` and `IMPORTED`.
     */
    readonly types?: string[];
}

/**
 * A collection of values returned by getCertificate.
 */
export interface GetCertificateResult {
    /**
     * Set to the ARN of the found certificate, suitable for referencing in other resources that support ACM certificates.
     */
    readonly arn: string;
    readonly domain: string;
    readonly mostRecent?: boolean;
    readonly statuses?: string[];
    readonly types?: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
