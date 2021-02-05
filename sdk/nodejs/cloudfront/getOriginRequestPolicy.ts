// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * The following example below creates a CloudFront origin request policy.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.cloudfront.getOriginRequestPolicy({
 *     name: "example-policy",
 * }, { async: true }));
 * ```
 */
export function getOriginRequestPolicy(args?: GetOriginRequestPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetOriginRequestPolicyResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:cloudfront/getOriginRequestPolicy:getOriginRequestPolicy", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOriginRequestPolicy.
 */
export interface GetOriginRequestPolicyArgs {
    /**
     * The identifier for the origin request policy.
     */
    readonly id?: string;
    /**
     * Unique name to identify the origin request policy.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getOriginRequestPolicy.
 */
export interface GetOriginRequestPolicyResult {
    /**
     * Comment to describe the origin request policy.
     */
    readonly comment: string;
    /**
     * Object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Cookies Config for more information.
     */
    readonly cookiesConfigs: outputs.cloudfront.GetOriginRequestPolicyCookiesConfig[];
    /**
     * The current version of the origin request policy.
     */
    readonly etag: string;
    /**
     * Object that determines whether any HTTP headers (and if so, which headers) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Headers Config for more information.
     */
    readonly headersConfigs: outputs.cloudfront.GetOriginRequestPolicyHeadersConfig[];
    readonly id?: string;
    readonly name?: string;
    /**
     * Object that determines whether any URL query strings in viewer requests (and if so, which query strings) are included in the origin request key and automatically included in requests that CloudFront sends to the origin. See Query Strings Config for more information.
     */
    readonly queryStringsConfigs: outputs.cloudfront.GetOriginRequestPolicyQueryStringsConfig[];
}
