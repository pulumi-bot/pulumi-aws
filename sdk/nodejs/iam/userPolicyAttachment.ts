// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {ARN} from "../index";
import {User} from "./user";

/**
 * Attaches a Managed IAM Policy to an IAM user
 * 
 * ~> **NOTE:** The usage of this resource conflicts with the `aws_iam_policy_attachment` resource and will permanently show a difference if both are defined.
 */
export class UserPolicyAttachment extends pulumi.CustomResource {
    /**
     * Get an existing UserPolicyAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPolicyAttachmentState): UserPolicyAttachment {
        return new UserPolicyAttachment(name, <any>state, { id });
    }

    /**
     * The ARN of the policy you want to apply
     */
    public readonly policyArn: pulumi.Output<ARN>;
    /**
     * The user the policy should be applied to
     */
    public readonly user: pulumi.Output<User>;

    /**
     * Create a UserPolicyAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPolicyAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPolicyAttachmentArgs | UserPolicyAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: UserPolicyAttachmentState = argsOrState as UserPolicyAttachmentState | undefined;
            inputs["policyArn"] = state ? state.policyArn : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as UserPolicyAttachmentArgs | undefined;
            if (!args || args.policyArn === undefined) {
                throw new Error("Missing required property 'policyArn'");
            }
            if (!args || args.user === undefined) {
                throw new Error("Missing required property 'user'");
            }
            inputs["policyArn"] = args ? args.policyArn : undefined;
            inputs["user"] = args ? args.user : undefined;
        }
        super("aws:iam/userPolicyAttachment:UserPolicyAttachment", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPolicyAttachment resources.
 */
export interface UserPolicyAttachmentState {
    /**
     * The ARN of the policy you want to apply
     */
    readonly policyArn?: pulumi.Input<ARN>;
    /**
     * The user the policy should be applied to
     */
    readonly user?: pulumi.Input<User>;
}

/**
 * The set of arguments for constructing a UserPolicyAttachment resource.
 */
export interface UserPolicyAttachmentArgs {
    /**
     * The ARN of the policy you want to apply
     */
    readonly policyArn: pulumi.Input<ARN>;
    /**
     * The user the policy should be applied to
     */
    readonly user: pulumi.Input<User>;
}
