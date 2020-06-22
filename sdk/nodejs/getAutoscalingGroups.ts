// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The Autoscaling Groups data source allows access to the list of AWS
 * ASGs within a specific region. This will allow you to pass a list of AutoScaling Groups to other resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const groups = pulumi.output(aws.getAutoscalingGroups({
 *     filters: [
 *         {
 *             name: "key",
 *             values: ["Team"],
 *         },
 *         {
 *             name: "value",
 *             values: ["Pets"],
 *         },
 *     ],
 * }, { async: true }));
 * const slackNotifications = new aws.autoscaling.Notification("slack_notifications", {
 *     groupNames: groups.names,
 *     notifications: [
 *         "autoscaling:EC2_INSTANCE_LAUNCH",
 *         "autoscaling:EC2_INSTANCE_TERMINATE",
 *         "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
 *         "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
 *     ],
 *     topicArn: "TOPIC ARN",
 * });
 * ```
 */
export function getAutoscalingGroups(args?: GetAutoscalingGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetAutoscalingGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getAutoscalingGroups:getAutoscalingGroups", {
        "filters": args.filters,
    }, opts);
}

/**
 * A collection of arguments for invoking getAutoscalingGroups.
 */
export interface GetAutoscalingGroupsArgs {
    /**
     * A filter used to scope the list e.g. by tags. See [related docs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Filter.html).
     */
    readonly filters?: inputs.GetAutoscalingGroupsFilter[];
}

/**
 * A collection of values returned by getAutoscalingGroups.
 */
export interface GetAutoscalingGroupsResult {
    /**
     * A list of the Autoscaling Groups Arns in the current region.
     */
    readonly arns: string[];
    readonly filters?: outputs.GetAutoscalingGroupsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the Autoscaling Groups in the current region.
     */
    readonly names: string[];
}
