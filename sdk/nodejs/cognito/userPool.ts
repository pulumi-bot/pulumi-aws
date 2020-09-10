// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class UserPool extends pulumi.CustomResource {
    /**
     * Get an existing UserPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPoolState, opts?: pulumi.CustomResourceOptions): UserPool {
        return new UserPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cognito/userPool:UserPool';

    /**
     * Returns true if the given object is an instance of UserPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPool.__pulumiType;
    }

    public readonly adminCreateUserConfig!: pulumi.Output<outputs.cognito.UserPoolAdminCreateUserConfig>;
    public readonly aliasAttributes!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly autoVerifiedAttributes!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    public readonly deviceConfiguration!: pulumi.Output<outputs.cognito.UserPoolDeviceConfiguration | undefined>;
    public readonly emailConfiguration!: pulumi.Output<outputs.cognito.UserPoolEmailConfiguration | undefined>;
    public readonly emailVerificationMessage!: pulumi.Output<string>;
    public readonly emailVerificationSubject!: pulumi.Output<string>;
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    public readonly lambdaConfig!: pulumi.Output<outputs.cognito.UserPoolLambdaConfig>;
    public /*out*/ readonly lastModifiedDate!: pulumi.Output<string>;
    public readonly mfaConfiguration!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly passwordPolicy!: pulumi.Output<outputs.cognito.UserPoolPasswordPolicy>;
    public readonly schemas!: pulumi.Output<outputs.cognito.UserPoolSchema[] | undefined>;
    public readonly smsAuthenticationMessage!: pulumi.Output<string | undefined>;
    public readonly smsConfiguration!: pulumi.Output<outputs.cognito.UserPoolSmsConfiguration>;
    public readonly smsVerificationMessage!: pulumi.Output<string>;
    public readonly softwareTokenMfaConfiguration!: pulumi.Output<outputs.cognito.UserPoolSoftwareTokenMfaConfiguration | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly userPoolAddOns!: pulumi.Output<outputs.cognito.UserPoolUserPoolAddOns | undefined>;
    public readonly usernameAttributes!: pulumi.Output<string[] | undefined>;
    public readonly usernameConfiguration!: pulumi.Output<outputs.cognito.UserPoolUsernameConfiguration | undefined>;
    public readonly verificationMessageTemplate!: pulumi.Output<outputs.cognito.UserPoolVerificationMessageTemplate>;

    /**
     * Create a UserPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPoolArgs | UserPoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserPoolState | undefined;
            inputs["adminCreateUserConfig"] = state ? state.adminCreateUserConfig : undefined;
            inputs["aliasAttributes"] = state ? state.aliasAttributes : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoVerifiedAttributes"] = state ? state.autoVerifiedAttributes : undefined;
            inputs["creationDate"] = state ? state.creationDate : undefined;
            inputs["deviceConfiguration"] = state ? state.deviceConfiguration : undefined;
            inputs["emailConfiguration"] = state ? state.emailConfiguration : undefined;
            inputs["emailVerificationMessage"] = state ? state.emailVerificationMessage : undefined;
            inputs["emailVerificationSubject"] = state ? state.emailVerificationSubject : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["lambdaConfig"] = state ? state.lambdaConfig : undefined;
            inputs["lastModifiedDate"] = state ? state.lastModifiedDate : undefined;
            inputs["mfaConfiguration"] = state ? state.mfaConfiguration : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["passwordPolicy"] = state ? state.passwordPolicy : undefined;
            inputs["schemas"] = state ? state.schemas : undefined;
            inputs["smsAuthenticationMessage"] = state ? state.smsAuthenticationMessage : undefined;
            inputs["smsConfiguration"] = state ? state.smsConfiguration : undefined;
            inputs["smsVerificationMessage"] = state ? state.smsVerificationMessage : undefined;
            inputs["softwareTokenMfaConfiguration"] = state ? state.softwareTokenMfaConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["userPoolAddOns"] = state ? state.userPoolAddOns : undefined;
            inputs["usernameAttributes"] = state ? state.usernameAttributes : undefined;
            inputs["usernameConfiguration"] = state ? state.usernameConfiguration : undefined;
            inputs["verificationMessageTemplate"] = state ? state.verificationMessageTemplate : undefined;
        } else {
            const args = argsOrState as UserPoolArgs | undefined;
            inputs["adminCreateUserConfig"] = args ? args.adminCreateUserConfig : undefined;
            inputs["aliasAttributes"] = args ? args.aliasAttributes : undefined;
            inputs["autoVerifiedAttributes"] = args ? args.autoVerifiedAttributes : undefined;
            inputs["deviceConfiguration"] = args ? args.deviceConfiguration : undefined;
            inputs["emailConfiguration"] = args ? args.emailConfiguration : undefined;
            inputs["emailVerificationMessage"] = args ? args.emailVerificationMessage : undefined;
            inputs["emailVerificationSubject"] = args ? args.emailVerificationSubject : undefined;
            inputs["lambdaConfig"] = args ? args.lambdaConfig : undefined;
            inputs["mfaConfiguration"] = args ? args.mfaConfiguration : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["passwordPolicy"] = args ? args.passwordPolicy : undefined;
            inputs["schemas"] = args ? args.schemas : undefined;
            inputs["smsAuthenticationMessage"] = args ? args.smsAuthenticationMessage : undefined;
            inputs["smsConfiguration"] = args ? args.smsConfiguration : undefined;
            inputs["smsVerificationMessage"] = args ? args.smsVerificationMessage : undefined;
            inputs["softwareTokenMfaConfiguration"] = args ? args.softwareTokenMfaConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["userPoolAddOns"] = args ? args.userPoolAddOns : undefined;
            inputs["usernameAttributes"] = args ? args.usernameAttributes : undefined;
            inputs["usernameConfiguration"] = args ? args.usernameConfiguration : undefined;
            inputs["verificationMessageTemplate"] = args ? args.verificationMessageTemplate : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["creationDate"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["lastModifiedDate"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserPool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPool resources.
 */
export interface UserPoolState {
    readonly adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfig>;
    readonly aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly arn?: pulumi.Input<string>;
    readonly autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly creationDate?: pulumi.Input<string>;
    readonly deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfiguration>;
    readonly emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfiguration>;
    readonly emailVerificationMessage?: pulumi.Input<string>;
    readonly emailVerificationSubject?: pulumi.Input<string>;
    readonly endpoint?: pulumi.Input<string>;
    readonly lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfig>;
    readonly lastModifiedDate?: pulumi.Input<string>;
    readonly mfaConfiguration?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly passwordPolicy?: pulumi.Input<inputs.cognito.UserPoolPasswordPolicy>;
    readonly schemas?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchema>[]>;
    readonly smsAuthenticationMessage?: pulumi.Input<string>;
    readonly smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfiguration>;
    readonly smsVerificationMessage?: pulumi.Input<string>;
    readonly softwareTokenMfaConfiguration?: pulumi.Input<inputs.cognito.UserPoolSoftwareTokenMfaConfiguration>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolUserPoolAddOns>;
    readonly usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfiguration>;
    readonly verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplate>;
}

/**
 * The set of arguments for constructing a UserPool resource.
 */
export interface UserPoolArgs {
    readonly adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfig>;
    readonly aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfiguration>;
    readonly emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfiguration>;
    readonly emailVerificationMessage?: pulumi.Input<string>;
    readonly emailVerificationSubject?: pulumi.Input<string>;
    readonly lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfig>;
    readonly mfaConfiguration?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly passwordPolicy?: pulumi.Input<inputs.cognito.UserPoolPasswordPolicy>;
    readonly schemas?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchema>[]>;
    readonly smsAuthenticationMessage?: pulumi.Input<string>;
    readonly smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfiguration>;
    readonly smsVerificationMessage?: pulumi.Input<string>;
    readonly softwareTokenMfaConfiguration?: pulumi.Input<inputs.cognito.UserPoolSoftwareTokenMfaConfiguration>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolUserPoolAddOns>;
    readonly usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfiguration>;
    readonly verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplate>;
}
