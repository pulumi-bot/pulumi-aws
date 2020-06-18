// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the name and value of a pre-existing API Key, for
 * example to supply credentials for a dependency microservice.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const myApiKey = pulumi.output(aws.apigateway.getKey({
 *     id: "ru3mpjgse6",
 * }, { async: true }));
 * ```
 */
export function getKey(args: GetKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:apigateway/getKey:getKey", {
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getKey.
 */
export interface GetKeyArgs {
    /**
     * The ID of the API Key to look up.
     */
    readonly id: string;
    /**
     * A map of tags for the resource.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getKey.
 */
export interface GetKeyResult {
    /**
     * The date and time when the API Key was created.
     */
    readonly createdDate: string;
    /**
     * The description of the API Key.
     */
    readonly description: string;
    /**
     * Specifies whether the API Key is enabled.
     */
    readonly enabled: boolean;
    /**
     * Set to the ID of the API Key.
     */
    readonly id: string;
    /**
     * The date and time when the API Key was last updated.
     */
    readonly lastUpdatedDate: string;
    /**
     * Set to the name of the API Key.
     */
    readonly name: string;
    /**
     * A map of tags for the resource.
     */
    readonly tags: {[key: string]: any};
    /**
     * Set to the value of the API Key.
     */
    readonly value: string;
}
