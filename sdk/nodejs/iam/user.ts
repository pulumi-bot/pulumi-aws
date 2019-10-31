// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an IAM user.
 * 
 * > *NOTE:* If policies are attached to the user via the [`aws.iam.PolicyAttachment` resource](https://www.terraform.io/docs/providers/aws/r/iam_policy_attachment.html) and you are modifying the user `name` or `path`, the `forceDestroy` argument must be set to `true` and applied before attempting the operation otherwise you will encounter a `DeleteConflict` error. The [`aws.iam.UserPolicyAttachment` resource (recommended)](https://www.terraform.io/docs/providers/aws/r/iam_user_policy_attachment.html) does not have this requirement.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const lbUser = new aws.iam.User("lb", {
 *     path: "/system/",
 *     tags: {
 *         "tag-key": "tag-value",
 *     },
 * });
 * const lbAccessKey = new aws.iam.AccessKey("lb", {
 *     user: lbUser.name,
 * });
 * const lbRo = new aws.iam.UserPolicy("lbRo", {
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "ec2:Describe*"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": "*"
 *     }
 *   ]
 * }
 * `,
 *     user: lbUser.name,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_user.html.markdown.
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * The ARN assigned by AWS for this user.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * When destroying this user, destroy even if it
     * has non-this provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
     * a user with non-this provider-managed access keys and login profile will fail to be destroyed.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Path in which to create the user.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the policy that is used to set the permissions boundary for the user.
     */
    public readonly permissionsBoundary!: pulumi.Output<string | undefined>;
    /**
     * Key-value mapping of tags for the IAM user
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The [unique ID][1] assigned by AWS.
     */
    public /*out*/ readonly uniqueId!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["permissionsBoundary"] = state ? state.permissionsBoundary : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uniqueId"] = state ? state.uniqueId : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["permissionsBoundary"] = args ? args.permissionsBoundary : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["uniqueId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * The ARN assigned by AWS for this user.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * When destroying this user, destroy even if it
     * has non-this provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
     * a user with non-this provider-managed access keys and login profile will fail to be destroyed.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Path in which to create the user.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The ARN of the policy that is used to set the permissions boundary for the user.
     */
    readonly permissionsBoundary?: pulumi.Input<string>;
    /**
     * Key-value mapping of tags for the IAM user
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The [unique ID][1] assigned by AWS.
     */
    readonly uniqueId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * When destroying this user, destroy even if it
     * has non-this provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
     * a user with non-this provider-managed access keys and login profile will fail to be destroyed.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Path in which to create the user.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * The ARN of the policy that is used to set the permissions boundary for the user.
     */
    readonly permissionsBoundary?: pulumi.Input<string>;
    /**
     * Key-value mapping of tags for the IAM user
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
