// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Provides an AWS Cognito Identity Pool Roles Attachment.
 */
export class IdentityPoolRoleAttachment extends pulumi.CustomResource {
    /**
     * Get an existing IdentityPoolRoleAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityPoolRoleAttachmentState): IdentityPoolRoleAttachment {
        return new IdentityPoolRoleAttachment(name, <any>state, { id });
    }

    /**
     * An identity pool ID in the format REGION:GUID.
     */
    public readonly identityPoolId: pulumi.Output<string>;
    /**
     * A List of Role Mapping.
     */
    public readonly roleMappings: pulumi.Output<{ ambiguousRoleResolution?: string, identityProvider: string, mappingRules?: { claim: string, matchType: string, roleArn: string, value: string }[], type: string }[] | undefined>;
    /**
     * The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
     */
    public readonly roles: pulumi.Output<{ authenticated?: string, unauthenticated?: string }>;

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
            const state: IdentityPoolRoleAttachmentState = argsOrState as IdentityPoolRoleAttachmentState | undefined;
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
        super("aws:cognito/identityPoolRoleAttachment:IdentityPoolRoleAttachment", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityPoolRoleAttachment resources.
 */
export interface IdentityPoolRoleAttachmentState {
    /**
     * An identity pool ID in the format REGION:GUID.
     */
    readonly identityPoolId?: pulumi.Input<string>;
    /**
     * A List of Role Mapping.
     */
    readonly roleMappings?: pulumi.Input<pulumi.Input<{ ambiguousRoleResolution?: pulumi.Input<string>, identityProvider: pulumi.Input<string>, mappingRules?: pulumi.Input<pulumi.Input<{ claim: pulumi.Input<string>, matchType: pulumi.Input<string>, roleArn: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]>;
    /**
     * The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
     */
    readonly roles?: pulumi.Input<{ authenticated?: pulumi.Input<string>, unauthenticated?: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a IdentityPoolRoleAttachment resource.
 */
export interface IdentityPoolRoleAttachmentArgs {
    /**
     * An identity pool ID in the format REGION:GUID.
     */
    readonly identityPoolId: pulumi.Input<string>;
    /**
     * A List of Role Mapping.
     */
    readonly roleMappings?: pulumi.Input<pulumi.Input<{ ambiguousRoleResolution?: pulumi.Input<string>, identityProvider: pulumi.Input<string>, mappingRules?: pulumi.Input<pulumi.Input<{ claim: pulumi.Input<string>, matchType: pulumi.Input<string>, roleArn: pulumi.Input<string>, value: pulumi.Input<string> }>[]>, type: pulumi.Input<string> }>[]>;
    /**
     * The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
     */
    readonly roles: pulumi.Input<{ authenticated?: pulumi.Input<string>, unauthenticated?: pulumi.Input<string> }>;
}
