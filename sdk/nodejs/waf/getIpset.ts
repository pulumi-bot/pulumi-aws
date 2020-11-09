// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * `aws.waf.IpSet` Retrieves a WAF IP Set Resource Id.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.waf.getIpset({
 *     name: "tfWAFIPSet",
 * }, { async: true }));
 * ```
 */
export function getIpset(args: GetIpsetArgs, opts?: pulumi.InvokeOptions): Promise<GetIpsetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:waf/getIpset:getIpset", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpset.
 */
export interface GetIpsetArgs {
    /**
     * The name of the WAF IP set.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getIpset.
 */
export interface GetIpsetResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
