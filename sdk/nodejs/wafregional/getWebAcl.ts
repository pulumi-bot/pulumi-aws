// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * `aws.wafregional.WebAcl` Retrieves a WAF Regional Web ACL Resource Id.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.wafregional.getWebAcl({
 *     name: "tfWAFRegionalWebACL",
 * }, { async: true }));
 * ```
 */
export function getWebAcl(args: GetWebAclArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAclResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:wafregional/getWebAcl:getWebAcl", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getWebAcl.
 */
export interface GetWebAclArgs {
    /**
     * The name of the WAF Regional Web ACL.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getWebAcl.
 */
export interface GetWebAclResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
