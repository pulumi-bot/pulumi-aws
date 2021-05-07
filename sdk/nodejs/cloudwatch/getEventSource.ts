// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an EventBridge Partner Event Source. This data source will only return one partner event source. An error will be returned if multiple sources match the same name prefix.
 *
 * > **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const examplepartner = pulumi.output(aws.cloudwatch.getEventSource({
 *     namePrefix: "aws.partner/examplepartner.com",
 * }, { async: true }));
 * ```
 */
export function getEventSource(args?: GetEventSourceArgs, opts?: pulumi.InvokeOptions): Promise<GetEventSourceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:cloudwatch/getEventSource:getEventSource", {
        "namePrefix": args.namePrefix,
    }, opts);
}

/**
 * A collection of arguments for invoking getEventSource.
 */
export interface GetEventSourceArgs {
    /**
     * Specifying this limits the results to only those partner event sources with names that start with the specified prefix
     */
    readonly namePrefix?: string;
}

/**
 * A collection of values returned by getEventSource.
 */
export interface GetEventSourceResult {
    /**
     * The ARN of the partner event source
     */
    readonly arn: string;
    /**
     * The name of the SaaS partner that created the event source
     */
    readonly createdBy: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the event source
     */
    readonly name: string;
    readonly namePrefix?: string;
    /**
     * The state of the event source (`ACTIVE` or `PENDING`)
     */
    readonly state: string;
}
