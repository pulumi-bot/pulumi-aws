// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Step Functions Activity data source
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const sfnActivity = pulumi.output(aws.sfn.getActivity({
 *     name: "my-activity",
 * }, { async: true }));
 * ```
 */
export function getActivity(args?: GetActivityArgs, opts?: pulumi.InvokeOptions): Promise<GetActivityResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:sfn/getActivity:getActivity", {
        "arn": args.arn,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getActivity.
 */
export interface GetActivityArgs {
    /**
     * The Amazon Resource Name (ARN) that identifies the activity.
     */
    readonly arn?: string;
    /**
     * The name that identifies the activity.
     */
    readonly name?: string;
}

/**
 * A collection of values returned by getActivity.
 */
export interface GetActivityResult {
    readonly arn: string;
    /**
     * The date the activity was created.
     */
    readonly creationDate: string;
    readonly name: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
