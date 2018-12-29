// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_iam_role_dlm_lifecycle_role = new aws.iam.Role("dlm_lifecycle_role", {
 *     assumeRolePolicy: "{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Action\": \"sts:AssumeRole\",\n      \"Principal\": {\n        \"Service\": \"dlm.amazonaws.com\"\n      },\n      \"Effect\": \"Allow\",\n      \"Sid\": \"\"\n    }\n  ]\n}\n",
 *     name: "dlm-lifecycle-role",
 * });
 * const aws_dlm_lifecycle_policy_example = new aws.dlm.LifecyclePolicy("example", {
 *     description: "example DLM lifecycle policy",
 *     executionRoleArn: aws_iam_role_dlm_lifecycle_role.arn,
 *     policyDetails: {
 *         resourceTypes: ["VOLUME"],
 *         schedules: [{
 *             copyTags: false,
 *             createRule: {
 *                 interval: 24,
 *                 intervalUnit: "HOURS",
 *                 times: "23:45",
 *             },
 *             name: "2 weeks of daily snapshots",
 *             retainRule: {
 *                 count: 14,
 *             },
 *             tagsToAdd: {
 *                 SnapshotCreator: "DLM",
 *             },
 *         }],
 *         targetTags: {
 *             Snapshot: "true",
 *         },
 *     },
 *     state: "ENABLED",
 * });
 * const aws_iam_role_policy_dlm_lifecycle = new aws.iam.RolePolicy("dlm_lifecycle", {
 *     name: "dlm-lifecycle-policy",
 *     policy: "{\n   \"Version\": \"2012-10-17\",\n   \"Statement\": [\n      {\n         \"Effect\": \"Allow\",\n         \"Action\": [\n            \"ec2:CreateSnapshot\",\n            \"ec2:DeleteSnapshot\",\n            \"ec2:DescribeVolumes\",\n            \"ec2:DescribeSnapshots\"\n         ],\n         \"Resource\": \"*\"\n      },\n      {\n         \"Effect\": \"Allow\",\n         \"Action\": [\n            \"ec2:CreateTags\"\n         ],\n         \"Resource\": \"arn:aws:ec2:*::snapshot/*\"\n      }\n   ]\n}\n",
 *     role: aws_iam_role_dlm_lifecycle_role.id,
 * });
 * ```
 */
export class LifecyclePolicy extends pulumi.CustomResource {
    /**
     * Get an existing LifecyclePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LifecyclePolicyState, opts?: pulumi.CustomResourceOptions): LifecyclePolicy {
        return new LifecyclePolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * A description for the DLM lifecycle policy.
     */
    public readonly description: pulumi.Output<string>;
    /**
     * The ARN of an IAM role that is able to be assumed by the DLM service.
     */
    public readonly executionRoleArn: pulumi.Output<string>;
    /**
     * See the `policy_details` configuration block. Max of 1.
     */
    public readonly policyDetails: pulumi.Output<{ resourceTypes: string[], schedules: { copyTags: boolean, createRule: { interval: number, intervalUnit?: string, times: string }, name: string, retainRule: { count: number }, tagsToAdd?: {[key: string]: any} }[], targetTags: {[key: string]: any} }>;
    /**
     * Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     */
    public readonly state: pulumi.Output<string | undefined>;

    /**
     * Create a LifecyclePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LifecyclePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LifecyclePolicyArgs | LifecyclePolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: LifecyclePolicyState = argsOrState as LifecyclePolicyState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["executionRoleArn"] = state ? state.executionRoleArn : undefined;
            inputs["policyDetails"] = state ? state.policyDetails : undefined;
            inputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as LifecyclePolicyArgs | undefined;
            if (!args || args.description === undefined) {
                throw new Error("Missing required property 'description'");
            }
            if (!args || args.executionRoleArn === undefined) {
                throw new Error("Missing required property 'executionRoleArn'");
            }
            if (!args || args.policyDetails === undefined) {
                throw new Error("Missing required property 'policyDetails'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            inputs["policyDetails"] = args ? args.policyDetails : undefined;
            inputs["state"] = args ? args.state : undefined;
        }
        super("aws:dlm/lifecyclePolicy:LifecyclePolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LifecyclePolicy resources.
 */
export interface LifecyclePolicyState {
    /**
     * A description for the DLM lifecycle policy.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that is able to be assumed by the DLM service.
     */
    readonly executionRoleArn?: pulumi.Input<string>;
    /**
     * See the `policy_details` configuration block. Max of 1.
     */
    readonly policyDetails?: pulumi.Input<{ resourceTypes: pulumi.Input<pulumi.Input<string>[]>, schedules: pulumi.Input<pulumi.Input<{ copyTags?: pulumi.Input<boolean>, createRule: pulumi.Input<{ interval: pulumi.Input<number>, intervalUnit?: pulumi.Input<string>, times?: pulumi.Input<string> }>, name: pulumi.Input<string>, retainRule: pulumi.Input<{ count: pulumi.Input<number> }>, tagsToAdd?: pulumi.Input<{[key: string]: any}> }>[]>, targetTags: pulumi.Input<{[key: string]: any}> }>;
    /**
     * Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LifecyclePolicy resource.
 */
export interface LifecyclePolicyArgs {
    /**
     * A description for the DLM lifecycle policy.
     */
    readonly description: pulumi.Input<string>;
    /**
     * The ARN of an IAM role that is able to be assumed by the DLM service.
     */
    readonly executionRoleArn: pulumi.Input<string>;
    /**
     * See the `policy_details` configuration block. Max of 1.
     */
    readonly policyDetails: pulumi.Input<{ resourceTypes: pulumi.Input<pulumi.Input<string>[]>, schedules: pulumi.Input<pulumi.Input<{ copyTags?: pulumi.Input<boolean>, createRule: pulumi.Input<{ interval: pulumi.Input<number>, intervalUnit?: pulumi.Input<string>, times?: pulumi.Input<string> }>, name: pulumi.Input<string>, retainRule: pulumi.Input<{ count: pulumi.Input<number> }>, tagsToAdd?: pulumi.Input<{[key: string]: any}> }>[]>, targetTags: pulumi.Input<{[key: string]: any}> }>;
    /**
     * Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
}
