// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * This data source can be used to fetch information about a specific
 * IAM group. By using this data source, you can reference IAM group
 * properties without having to hard code ARNs as input.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.iam.getGroup({
 *     groupName: "an_example_group_name",
 * }, { async: true }));
 * ```
 */
export function getGroup(args: GetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getGroup:getGroup", {
        "groupName": args.groupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getGroup.
 */
export interface GetGroupArgs {
    /**
     * The friendly IAM group name to match.
     */
    groupName: string;
}

/**
 * A collection of values returned by getGroup.
 */
export interface GetGroupResult {
    /**
     * The Amazon Resource Name (ARN) specifying the iam user.
     */
    readonly arn: string;
    /**
     * The stable and unique string identifying the group.
     */
    readonly groupId: string;
    readonly groupName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The path to the iam user.
     */
    readonly path: string;
    /**
     * List of objects containing group member information. See supported fields below.
     */
    readonly users: outputs.iam.GetGroupUser[];
}
