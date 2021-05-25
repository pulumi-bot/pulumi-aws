// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * This data source can be used to fetch information about a specific
 * IAM policy.
 *
 * ## Example Usage
 * ### By ARN
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.iam.getPolicy({
 *     arn: "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
 * }, { async: true }));
 * ```
 * ### By Name
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.iam.getPolicy({
 *     name: "test_policy",
 * }, { async: true }));
 * ```
 */
export function getPolicy(args?: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:iam/getPolicy:getPolicy", {
        "arn": args.arn,
        "name": args.name,
        "pathPrefix": args.pathPrefix,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getPolicy.
 */
export interface GetPolicyArgs {
    /**
     * The ARN of the IAM policy.
     */
    arn?: string;
    /**
     * The name of the IAM policy.
     */
    name?: string;
    /**
     * The prefix of the path to the IAM policy. Defaults to a slash (`/`).
     */
    pathPrefix?: string;
    /**
     * Key-value mapping of tags for the IAM Policy.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getPolicy.
 */
export interface GetPolicyResult {
    readonly arn: string;
    /**
     * The description of the policy.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The path to the policy.
     */
    readonly path: string;
    readonly pathPrefix?: string;
    /**
     * The policy document of the policy.
     */
    readonly policy: string;
    /**
     * The policy's ID.
     */
    readonly policyId: string;
    /**
     * Key-value mapping of tags for the IAM Policy.
     */
    readonly tags: {[key: string]: string};
}
