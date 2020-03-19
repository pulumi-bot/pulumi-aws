// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an OpsWorks User Profile resource.
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const myProfile = new aws.opsworks.UserProfile("myProfile", {
 *     sshUsername: "myUser",
 *     userArn: aws_iam_user_user.arn,
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/opsworks_user_profile.html.markdown.
 */
export class UserProfile extends pulumi.CustomResource {
    /**
     * Get an existing UserProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserProfileState, opts?: pulumi.CustomResourceOptions): UserProfile {
        return new UserProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opsworks/userProfile:UserProfile';

    /**
     * Returns true if the given object is an instance of UserProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserProfile.__pulumiType;
    }

    /**
     * Whether users can specify their own SSH public key through the My Settings page
     */
    public readonly allowSelfManagement!: pulumi.Output<boolean | undefined>;
    /**
     * The users public key
     */
    public readonly sshPublicKey!: pulumi.Output<string | undefined>;
    /**
     * The ssh username, with witch this user wants to log in
     */
    public readonly sshUsername!: pulumi.Output<string>;
    /**
     * The user's IAM ARN
     */
    public readonly userArn!: pulumi.Output<string>;

    /**
     * Create a UserProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserProfileArgs | UserProfileState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserProfileState | undefined;
            inputs["allowSelfManagement"] = state ? state.allowSelfManagement : undefined;
            inputs["sshPublicKey"] = state ? state.sshPublicKey : undefined;
            inputs["sshUsername"] = state ? state.sshUsername : undefined;
            inputs["userArn"] = state ? state.userArn : undefined;
        } else {
            const args = argsOrState as UserProfileArgs | undefined;
            if (!args || args.sshUsername === undefined) {
                throw new Error("Missing required property 'sshUsername'");
            }
            if (!args || args.userArn === undefined) {
                throw new Error("Missing required property 'userArn'");
            }
            inputs["allowSelfManagement"] = args ? args.allowSelfManagement : undefined;
            inputs["sshPublicKey"] = args ? args.sshPublicKey : undefined;
            inputs["sshUsername"] = args ? args.sshUsername : undefined;
            inputs["userArn"] = args ? args.userArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserProfile.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserProfile resources.
 */
export interface UserProfileState {
    /**
     * Whether users can specify their own SSH public key through the My Settings page
     */
    readonly allowSelfManagement?: pulumi.Input<boolean>;
    /**
     * The users public key
     */
    readonly sshPublicKey?: pulumi.Input<string>;
    /**
     * The ssh username, with witch this user wants to log in
     */
    readonly sshUsername?: pulumi.Input<string>;
    /**
     * The user's IAM ARN
     */
    readonly userArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserProfile resource.
 */
export interface UserProfileArgs {
    /**
     * Whether users can specify their own SSH public key through the My Settings page
     */
    readonly allowSelfManagement?: pulumi.Input<boolean>;
    /**
     * The users public key
     */
    readonly sshPublicKey?: pulumi.Input<string>;
    /**
     * The ssh username, with witch this user wants to log in
     */
    readonly sshUsername: pulumi.Input<string>;
    /**
     * The user's IAM ARN
     */
    readonly userArn: pulumi.Input<string>;
}
