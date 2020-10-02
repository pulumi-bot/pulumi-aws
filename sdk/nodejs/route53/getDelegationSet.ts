// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * `aws.route53.DelegationSet` provides details about a specific Route 53 Delegation Set.
 *
 * This data source allows to find a list of name servers associated with a specific delegation set.
 *
 * ## Example Usage
 *
 * The following example shows how to get a delegation set from its id.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const dset = pulumi.output(aws.route53.getDelegationSet({
 *     id: "MQWGHCBFAKEID",
 * }, { async: true }));
 * ```
 */
export function getDelegationSet(args: GetDelegationSetArgs, opts?: pulumi.InvokeOptions): Promise<GetDelegationSetResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:route53/getDelegationSet:getDelegationSet", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getDelegationSet.
 */
export interface GetDelegationSetArgs {
    /**
     * The Hosted Zone id of the desired delegation set.
     */
    readonly id: string;
}

/**
 * A collection of values returned by getDelegationSet.
 */
export interface GetDelegationSetResult {
    readonly callerReference: string;
    readonly id: string;
    readonly nameServers: string[];
}
