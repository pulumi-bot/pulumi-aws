// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class IdentityPoolRoleAttachment extends pulumi.CustomResource {
    /**
     * Get an existing IdentityPoolRoleAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityPoolRoleAttachmentState, opts?: pulumi.CustomResourceOptions): IdentityPoolRoleAttachment {
        return new IdentityPoolRoleAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment';

    /**
     * Returns true if the given object is an instance of IdentityPoolRoleAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityPoolRoleAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityPoolRoleAttachment.__pulumiType;
    }

    public readonly identityPoolId!: pulumi.Output<string>;
    public readonly roleMappings!: pulumi.Output<outputs.cognito.IdentityPoolRoleAttachmentRoleMapping[] | undefined>;
    public readonly roles!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a IdentityPoolRoleAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityPoolRoleAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityPoolRoleAttachmentArgs | IdentityPoolRoleAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IdentityPoolRoleAttachmentState | undefined;
            inputs["identityPoolId"] = state ? state.identityPoolId : undefined;
            inputs["roleMappings"] = state ? state.roleMappings : undefined;
            inputs["roles"] = state ? state.roles : undefined;
        } else {
            const args = argsOrState as IdentityPoolRoleAttachmentArgs | undefined;
            if (!args || args.identityPoolId === undefined) {
                throw new Error("Missing required property 'identityPoolId'");
            }
            if (!args || args.roles === undefined) {
                throw new Error("Missing required property 'roles'");
            }
            inputs["identityPoolId"] = args ? args.identityPoolId : undefined;
            inputs["roleMappings"] = args ? args.roleMappings : undefined;
            inputs["roles"] = args ? args.roles : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IdentityPoolRoleAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityPoolRoleAttachment resources.
 */
export interface IdentityPoolRoleAttachmentState {
    readonly identityPoolId?: pulumi.Input<string>;
    readonly roleMappings?: pulumi.Input<pulumi.Input<inputs.cognito.IdentityPoolRoleAttachmentRoleMapping>[]>;
    readonly roles?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a IdentityPoolRoleAttachment resource.
 */
export interface IdentityPoolRoleAttachmentArgs {
    readonly identityPoolId: pulumi.Input<string>;
    readonly roleMappings?: pulumi.Input<pulumi.Input<inputs.cognito.IdentityPoolRoleAttachmentRoleMapping>[]>;
    readonly roles: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
