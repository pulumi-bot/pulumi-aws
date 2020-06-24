// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > **WARNING:** Multiple aws.iam.GroupMembership resources with the same group name will produce inconsistent behavior!
 *
 * Provides a top level resource to manage IAM Group membership for IAM Users. For
 * more information on managing IAM Groups or IAM Users, see [IAM Groups](https://www.terraform.io/docs/providers/aws/r/iam_group.html) or
 * [IAM Users](https://www.terraform.io/docs/providers/aws/r/iam_user.html)
 *
 * > **Note:** `aws.iam.GroupMembership` will conflict with itself if used more than once with the same group. To non-exclusively manage the users in a group, see the
 * [`aws.iam.UserGroupMembership` resource][3].
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const group = new aws.iam.Group("group", {});
 * const userOne = new aws.iam.User("user_one", {});
 * const userTwo = new aws.iam.User("user_two", {});
 * const team = new aws.iam.GroupMembership("team", {
 *     group: group.name,
 *     users: [
 *         userOne.name,
 *         userTwo.name,
 *     ],
 * });
 * ```
 */
export class GroupMembership extends pulumi.CustomResource {
    /**
     * Get an existing GroupMembership resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupMembershipState, opts?: pulumi.CustomResourceOptions): GroupMembership {
        return new GroupMembership(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iam/groupMembership:GroupMembership';

    /**
     * Returns true if the given object is an instance of GroupMembership.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupMembership {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupMembership.__pulumiType;
    }

    /**
     * The IAM Group name to attach the list of `users` to
     */
    public readonly group!: pulumi.Output<string>;
    /**
     * The name to identify the Group Membership
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of IAM User names to associate with the Group
     */
    public readonly users!: pulumi.Output<string[]>;

    /**
     * Create a GroupMembership resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupMembershipArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupMembershipArgs | GroupMembershipState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupMembershipState | undefined;
            inputs["group"] = state ? state.group : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["users"] = state ? state.users : undefined;
        } else {
            const args = argsOrState as GroupMembershipArgs | undefined;
            if (!args || args.group === undefined) {
                throw new Error("Missing required property 'group'");
            }
            if (!args || args.users === undefined) {
                throw new Error("Missing required property 'users'");
            }
            inputs["group"] = args ? args.group : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["users"] = args ? args.users : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GroupMembership.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupMembership resources.
 */
export interface GroupMembershipState {
    /**
     * The IAM Group name to attach the list of `users` to
     */
    readonly group?: pulumi.Input<string>;
    /**
     * The name to identify the Group Membership
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of IAM User names to associate with the Group
     */
    readonly users?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a GroupMembership resource.
 */
export interface GroupMembershipArgs {
    /**
     * The IAM Group name to attach the list of `users` to
     */
    readonly group: pulumi.Input<string>;
    /**
     * The name to identify the Group Membership
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of IAM User names to associate with the Group
     */
    readonly users: pulumi.Input<pulumi.Input<string>[]>;
}
