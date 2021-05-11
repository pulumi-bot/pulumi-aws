// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Returns a unique endpoint specific to the AWS account making the call.
 */
export function getEndpoint(args?: GetEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iot/getEndpoint:getEndpoint", {
        "endpointType": args.endpointType,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpoint.
 */
export interface GetEndpointArgs {
    /**
     * Endpoint type. Valid values: `iot:CredentialProvider`, `iot:Data`, `iot:Data-ATS`, `iot:Job`.
     */
    endpointType?: string;
}

/**
 * A collection of values returned by getEndpoint.
 */
export interface GetEndpointResult {
    /**
     * The endpoint based on `endpointType`:
     * * No `endpointType`: Either `iot:Data` or `iot:Data-ATS` [depending on region](https://aws.amazon.com/blogs/iot/aws-iot-core-ats-endpoints/)
     * * `iot:CredentialsProvider`: `IDENTIFIER.credentials.iot.REGION.amazonaws.com`
     * * `iot:Data`: `IDENTIFIER.iot.REGION.amazonaws.com`
     * * `iot:Data-ATS`: `IDENTIFIER-ats.iot.REGION.amazonaws.com`
     * * `iot:Job`: `IDENTIFIER.jobs.iot.REGION.amazonaws.com`
     */
    readonly endpointAddress: string;
    readonly endpointType?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
