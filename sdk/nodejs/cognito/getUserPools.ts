// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of cognito user pools.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const selectedRestApi = aws.apigateway.getRestApi({
 *     name: _var.api_gateway_name,
 * });
 * const selectedUserPools = aws.cognito.getUserPools({
 *     name: _var.cognito_user_pool_name,
 * });
 * const cognito = new aws.apigateway.Authorizer("cognito", {
 *     type: "COGNITO_USER_POOLS",
 *     restApi: selectedRestApi.then(selectedRestApi => selectedRestApi.id),
 *     providerArns: selectedUserPools.then(selectedUserPools => selectedUserPools.arns),
 * });
 * ```
 */
export function getUserPools(args: GetUserPoolsArgs, opts?: pulumi.InvokeOptions): Promise<GetUserPoolsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:cognito/getUserPools:getUserPools", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getUserPools.
 */
export interface GetUserPoolsArgs {
    /**
     * Name of the cognito user pools. Name is not a unique attribute for cognito user pool, so multiple pools might be returned with given name.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getUserPools.
 */
export interface GetUserPoolsResult {
    readonly arns: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The list of cognito user pool ids.
     */
    readonly ids: string[];
    readonly name: string;
}
