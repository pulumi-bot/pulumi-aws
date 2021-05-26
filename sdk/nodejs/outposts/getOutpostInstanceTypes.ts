// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Information about Outposts Instance Types.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.outposts.getOutpostInstanceTypes({
 *     arn: data.aws_outposts_outpost.example.arn,
 * });
 * ```
 */
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
    /**
     * Outpost Amazon Resource Name (ARN).
     */
    arn: string;
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
    /**
     * Set of instance types.
     */
    readonly instanceTypes: string[];
}

export function getOutpostInstanceTypesApply(args: GetOutpostInstanceTypesApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOutpostInstanceTypesResult> {
    return pulumi.output(args).apply(a => getOutpostInstanceTypes(a, opts))
}

/**
 * A collection of arguments for invoking getOutpostInstanceTypes.
 */
export interface GetOutpostInstanceTypesApplyArgs {
    /**
     * Outpost Amazon Resource Name (ARN).
     */
    arn: pulumi.Input<string>;
}
