// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Use this data source to lookup information about IAM Server Certificates.
 */
export function getServerCertificate(args?: GetServerCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetServerCertificateResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getServerCertificate:getServerCertificate", {
        "latest": args.latest,
        "name": args.name,
        "namePrefix": args.namePrefix,
        "pathPrefix": args.pathPrefix,
    }, opts);
}

/**
 * A collection of arguments for invoking getServerCertificate.
 */
export interface GetServerCertificateArgs {
    /**
     * sort results by expiration date. returns the certificate with expiration date in furthest in the future.
     */
    readonly latest?: boolean;
    /**
     * exact name of the cert to lookup
     */
    readonly name?: string;
    /**
     * prefix of cert to filter by
     */
    readonly namePrefix?: string;
    /**
     * prefix of path to filter by
     */
    readonly pathPrefix?: string;
}

/**
 * A collection of values returned by getServerCertificate.
 */
export interface GetServerCertificateResult {
    readonly arn: string;
    readonly certificateBody: string;
    readonly certificateChain: string;
    readonly expirationDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly latest?: boolean;
    readonly name: string;
    readonly namePrefix?: string;
    readonly path: string;
    readonly pathPrefix?: string;
    readonly uploadDate: string;
}
