// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can be used to fetch information about a specific
 * IAM role. By using this data source, you can reference IAM role
 * properties without having to hard code ARNs as input.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.iam.getRole({
 *     name: "anExampleRoleName",
 * }, { async: true }));
 * ```
 */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getRole:getRole", {
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleArgs {
    /**
     * The friendly IAM role name to match.
     */
    readonly name: string;
    /**
     * The tags attached to the role.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    /**
     * The Amazon Resource Name (ARN) specifying the role.
     */
    readonly arn: string;
    /**
     * The policy document associated with the role.
     */
    readonly assumeRolePolicy: string;
    /**
     * Creation date of the role in RFC 3339 format.
     */
    readonly createDate: string;
    /**
     * Description for the role.
     */
    readonly description: string;
    /**
     * Maximum session duration.
     */
    readonly maxSessionDuration: number;
    readonly name: string;
    /**
     * The path to the role.
     */
    readonly path: string;
    /**
     * The ARN of the policy that is used to set the permissions boundary for the role.
     */
    readonly permissionsBoundary: string;
    /**
     * The tags attached to the role.
     */
    readonly tags: {[key: string]: any};
    /**
     * The stable and unique string identifying the role.
     */
    readonly uniqueId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
