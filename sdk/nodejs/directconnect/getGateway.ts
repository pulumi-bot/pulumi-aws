// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Retrieve information about a Direct Connect Gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.directconnect.getGateway({
 *     name: "example",
 * }, { async: true }));
 * ```
 */
export function getGateway(args: GetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:directconnect/getGateway:getGateway", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getGateway.
 */
export interface GetGatewayArgs {
    /**
     * The name of the gateway to retrieve.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getGateway.
 */
export interface GetGatewayResult {
    /**
     * The ASN on the Amazon side of the connection.
     */
    readonly amazonSideAsn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * AWS Account ID of the gateway.
     */
    readonly ownerAccountId: string;
}
