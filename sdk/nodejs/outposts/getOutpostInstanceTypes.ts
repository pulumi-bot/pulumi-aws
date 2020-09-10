// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getOutpostInstanceTypes(args: GetOutpostInstanceTypesArgs, opts?: pulumi.InvokeOptions): Promise<GetOutpostInstanceTypesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:outposts/getOutpostInstanceTypes:getOutpostInstanceTypes", {
        "arn": args.arn,
    }, opts);
}

/**
 * A collection of arguments for invoking getOutpostInstanceTypes.
 */
export interface GetOutpostInstanceTypesArgs {
    readonly arn: string;
}

/**
 * A collection of values returned by getOutpostInstanceTypes.
 */
export interface GetOutpostInstanceTypesResult {
    readonly arn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceTypes: string[];
}
