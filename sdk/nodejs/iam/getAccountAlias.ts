// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The IAM Account Alias data source allows access to the account alias
 * for the effective account in which this provider is working.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const current = pulumi.output(aws.iam.getAccountAlias({ async: true }));
 *
 * export const accountId = current.accountAlias;
 * ```
 */
export function getAccountAlias(opts?: pulumi.InvokeOptions): Promise<GetAccountAliasResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getAccountAlias:getAccountAlias", {
    }, opts);
}

/**
 * A collection of values returned by getAccountAlias.
 */
export interface GetAccountAliasResult {
    /**
     * The alias associated with the AWS account.
     */
    readonly accountAlias: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
