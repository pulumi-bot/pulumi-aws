// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getCredentials(args: GetCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetCredentialsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ecr/getCredentials:getCredentials", {
        "registryId": args.registryId,
    }, opts);
}

/**
 * A collection of arguments for invoking getCredentials.
 */
export interface GetCredentialsArgs {
    readonly registryId: string;
}

/**
 * A collection of values returned by getCredentials.
 */
export interface GetCredentialsResult {
    readonly authorizationToken: string;
    readonly expiresAt: string;
    readonly proxyEndpoint: string;
    readonly registryId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
