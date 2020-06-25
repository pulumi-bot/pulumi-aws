// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an SSM Patch Baseline data source. Useful if you wish to reuse the default baselines provided.
 *
 * ## Example Usage
 *
 * To retrieve a baseline provided by AWS:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const centos = pulumi.output(aws.ssm.getPatchBaseline({
 *     namePrefix: "AWS-",
 *     operatingSystem: "CENTOS",
 *     owner: "AWS",
 * }, { async: true }));
 * ```
 *
 * To retrieve a baseline on your account:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const defaultCustom = pulumi.output(aws.ssm.getPatchBaseline({
 *     defaultBaseline: true,
 *     namePrefix: "MyCustomBaseline",
 *     operatingSystem: "WINDOWS",
 *     owner: "Self",
 * }, { async: true }));
 * ```
 */
export function getPatchBaseline(args: GetPatchBaselineArgs, opts?: pulumi.InvokeOptions): Promise<GetPatchBaselineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ssm/getPatchBaseline:getPatchBaseline", {
        "defaultBaseline": args.defaultBaseline,
        "namePrefix": args.namePrefix,
        "operatingSystem": args.operatingSystem,
        "owner": args.owner,
    }, opts);
}

/**
 * A collection of arguments for invoking getPatchBaseline.
 */
export interface GetPatchBaselineArgs {
    /**
     * Filters the results against the baselines defaultBaseline field.
     */
    readonly defaultBaseline?: boolean;
    /**
     * Filter results by the baseline name prefix.
     */
    readonly namePrefix?: string;
    /**
     * The specified OS for the baseline.
     */
    readonly operatingSystem?: string;
    /**
     * The owner of the baseline. Valid values: `All`, `AWS`, `Self` (the current account).
     */
    readonly owner: string;
}

/**
 * A collection of values returned by getPatchBaseline.
 */
export interface GetPatchBaselineResult {
    readonly defaultBaseline?: boolean;
    /**
     * The description of the baseline.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the baseline.
     */
    readonly name: string;
    readonly namePrefix?: string;
    readonly operatingSystem?: string;
    readonly owner: string;
}
