// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

import {PolicyDocument} from "./index";

/**
 * Provides an IAM role.
 *
 * > **NOTE:** If policies are attached to the role via the `aws.iam.PolicyAttachment` resource and you are modifying the role `name` or `path`, the `forceDetachPolicies` argument must be set to `true` and applied before attempting the operation otherwise you will encounter a `DeleteConflict` error. The `aws.iam.RolePolicyAttachment` resource does not have this requirement.
 *
 * > **NOTE:** If you use this resource's `managedPolicyArns` argument or `inlinePolicy` configuration blocks, this resource will take over exclusive management of the role's respective policy types (e.g., both policy types if both arguments are used). These arguments are incompatible with other ways of managing a role's policies, such as `aws.iam.PolicyAttachment`, `aws.iam.RolePolicyAttachment`, and `aws.iam.RolePolicy`. If you attempt to manage a role's policies by multiple means, you will get resource cycling and/or errors.
 *
 * ## Example Usage
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testRole = new aws.iam.Role("testRole", {
 *     assumeRolePolicy: JSON.stringify({
 *         Version: "2012-10-17",
 *         Statement: [{
 *             Action: "sts:AssumeRole",
 *             Effect: "Allow",
 *             Sid: "",
 *             Principal: {
 *                 Service: "ec2.amazonaws.com",
 *             },
 *         }],
 *     }),
 *     tags: {
 *         "tag-key": "tag-value",
 *     },
 * });
 * ```
 * ### Example of Using Data Source for Assume Role Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const instance-assume-role-policy = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["ec2.amazonaws.com"],
 *         }],
 *     }],
 * });
 * const instance = new aws.iam.Role("instance", {
 *     path: "/system/",
 *     assumeRolePolicy: instance_assume_role_policy.then(instance_assume_role_policy => instance_assume_role_policy.json),
 * });
 * ```
 * ### Example of Exclusive Inline Policies
 *
 * This example creates an IAM role with two inline IAM policies. If someone adds another inline policy out-of-band, on the next apply, the provider will remove that policy. If someone deletes these policies out-of-band, the provider will recreate them.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const inlinePolicy = aws.iam.getPolicyDocument({
 *     statements: [{
 *         actions: ["ec2:DescribeAccountAttributes"],
 *         resources: ["*"],
 *     }],
 * });
 * const example = new aws.iam.Role("example", {
 *     assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
 *     inlinePolicies: [
 *         {
 *             name: "my_inline_policy",
 *             policy: JSON.stringify({
 *                 Version: "2012-10-17",
 *                 Statement: [{
 *                     Action: ["ec2:Describe*"],
 *                     Effect: "Allow",
 *                     Resource: "*",
 *                 }],
 *             }),
 *         },
 *         {
 *             name: "policy-8675309",
 *             policy: inlinePolicy.then(inlinePolicy => inlinePolicy.json),
 *         },
 *     ],
 * });
 * ```
 * ### Example of Removing Inline Policies
 *
 * This example creates an IAM role with what appears to be empty IAM `inlinePolicy` argument instead of using `inlinePolicy` as a configuration block. The result is that if someone were to add an inline policy out-of-band, on the next apply, the provider will remove that policy.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.iam.Role("example", {
 *     assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
 *     inlinePolicies: [{}],
 * });
 * ```
 * ### Example of Exclusive Managed Policies
 *
 * This example creates an IAM role and attaches two managed IAM policies. If someone attaches another managed policy out-of-band, on the next apply, the provider will detach that policy. If someone detaches these policies out-of-band, the provider will attach them again.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const policyOne = new aws.iam.Policy("policyOne", {policy: JSON.stringify({
 *     Version: "2012-10-17",
 *     Statement: [{
 *         Action: ["ec2:Describe*"],
 *         Effect: "Allow",
 *         Resource: "*",
 *     }],
 * })});
 * const policyTwo = new aws.iam.Policy("policyTwo", {policy: JSON.stringify({
 *     Version: "2012-10-17",
 *     Statement: [{
 *         Action: [
 *             "s3:ListAllMyBuckets",
 *             "s3:ListBucket",
 *             "s3:HeadBucket",
 *         ],
 *         Effect: "Allow",
 *         Resource: "*",
 *     }],
 * })});
 * const example = new aws.iam.Role("example", {
 *     assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
 *     managedPolicyArns: [
 *         policyOne.arn,
 *         policyTwo.arn,
 *     ],
 * });
 * ```
 * ### Example of Removing Managed Policies
 *
 * This example creates an IAM role with an empty `managedPolicyArns` argument. If someone attaches a policy out-of-band, on the next apply, the provider will detach that policy.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.iam.Role("example", {
 *     assumeRolePolicy: data.aws_iam_policy_document.instance_assume_role_policy.json,
 *     managedPolicyArns: [],
 * });
 * ```
 *
 * ## Import
 *
 * IAM Roles can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:iam/role:Role developer developer_name
 * ```
 */
export class Role extends pulumi.CustomResource {
    /**
     * Get an existing Role resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoleState, opts?: pulumi.CustomResourceOptions): Role {
        return new Role(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/role:Role';

    /**
     * Returns true if the given object is an instance of Role.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Role {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Role.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) specifying the role.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Policy that grants an entity permission to assume the role.
     */
    public readonly assumeRolePolicy!: pulumi.Output<string>;
    /**
     * Creation date of the IAM role.
     */
    public /*out*/ readonly createDate!: pulumi.Output<string>;
    /**
     * Description of the role.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
     */
    public readonly forceDetachPolicies!: pulumi.Output<boolean | undefined>;
    /**
     * Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
     */
    public readonly inlinePolicies!: pulumi.Output<outputs.iam.RoleInlinePolicy[]>;
    /**
     * Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
     */
    public readonly managedPolicyArns!: pulumi.Output<string[]>;
    /**
     * Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
     */
    public readonly maxSessionDuration!: pulumi.Output<number | undefined>;
    /**
     * Name of the role policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * ARN of the policy that is used to set the permissions boundary for the role.
     */
    public readonly permissionsBoundary!: pulumi.Output<string | undefined>;
    /**
     * Key-value mapping of tags for the IAM role. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Stable and unique string identifying the role.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;

    /**
     * Create a Role resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoleArgs | RoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["assumeRolePolicy"] = state ? state.assumeRolePolicy : undefined;
            inputs["createDate"] = state ? state.createDate : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["forceDetachPolicies"] = state ? state.forceDetachPolicies : undefined;
            inputs["inlinePolicies"] = state ? state.inlinePolicies : undefined;
            inputs["managedPolicyArns"] = state ? state.managedPolicyArns : undefined;
            inputs["maxSessionDuration"] = state ? state.maxSessionDuration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["permissionsBoundary"] = state ? state.permissionsBoundary : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["uniqueId"] = state ? state.uniqueId : undefined;
        } else {
            const args = argsOrState as RoleArgs | undefined;
            if ((!args || args.assumeRolePolicy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'assumeRolePolicy'");
            }
            inputs["assumeRolePolicy"] = args ? args.assumeRolePolicy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["forceDetachPolicies"] = args ? args.forceDetachPolicies : undefined;
            inputs["inlinePolicies"] = args ? args.inlinePolicies : undefined;
            inputs["managedPolicyArns"] = args ? args.managedPolicyArns : undefined;
            inputs["maxSessionDuration"] = args ? args.maxSessionDuration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["permissionsBoundary"] = args ? args.permissionsBoundary : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createDate"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Role.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Role resources.
 */
export interface RoleState {
    /**
     * Amazon Resource Name (ARN) specifying the role.
     */
    arn?: pulumi.Input<string>;
    /**
     * Policy that grants an entity permission to assume the role.
     */
    assumeRolePolicy?: pulumi.Input<string | PolicyDocument>;
    /**
     * Creation date of the IAM role.
     */
    createDate?: pulumi.Input<string>;
    /**
     * Description of the role.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
     */
    forceDetachPolicies?: pulumi.Input<boolean>;
    /**
     * Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
     */
    inlinePolicies?: pulumi.Input<pulumi.Input<inputs.iam.RoleInlinePolicy>[]>;
    /**
     * Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
     */
    managedPolicyArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
     */
    maxSessionDuration?: pulumi.Input<number>;
    /**
     * Name of the role policy.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
     */
    path?: pulumi.Input<string>;
    /**
     * ARN of the policy that is used to set the permissions boundary for the role.
     */
    permissionsBoundary?: pulumi.Input<string>;
    /**
     * Key-value mapping of tags for the IAM role. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Stable and unique string identifying the role.
     */
    uniqueId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Role resource.
 */
export interface RoleArgs {
    /**
     * Policy that grants an entity permission to assume the role.
     */
    assumeRolePolicy: pulumi.Input<string | PolicyDocument>;
    /**
     * Description of the role.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether to force detaching any policies the role has before destroying it. Defaults to `false`.
     */
    forceDetachPolicies?: pulumi.Input<boolean>;
    /**
     * Configuration block defining an exclusive set of IAM inline policies associated with the IAM role. Defined below. If no blocks are configured, the provider will ignore any managing any inline policies in this resource. Configuring one empty block (i.e., `inlinePolicy {}`) will cause the provider to remove _all_ inline policies.
     */
    inlinePolicies?: pulumi.Input<pulumi.Input<inputs.iam.RoleInlinePolicy>[]>;
    /**
     * Set of exclusive IAM managed policy ARNs to attach to the IAM role. If this attribute is not configured, the provider will ignore policy attachments to this resource. When configured, the provider will align the role's managed policy attachments with this set by attaching or detaching managed policies. Configuring an empty set (i.e., `managedPolicyArns = []`) will cause the provider to remove _all_ managed policy attachments.
     */
    managedPolicyArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum session duration (in seconds) that you want to set for the specified role. If you do not specify a value for this setting, the default maximum of one hour is applied. This setting can have a value from 1 hour to 12 hours.
     */
    maxSessionDuration?: pulumi.Input<number>;
    /**
     * Name of the role policy.
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique friendly name beginning with the specified prefix. Conflicts with `name`.
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * Path to the role. See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
     */
    path?: pulumi.Input<string>;
    /**
     * ARN of the policy that is used to set the permissions boundary for the role.
     */
    permissionsBoundary?: pulumi.Input<string>;
    /**
     * Key-value mapping of tags for the IAM role. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
