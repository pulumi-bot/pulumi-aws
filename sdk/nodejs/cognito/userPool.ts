// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Cognito User Pool resource.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cognito_user_pool.markdown.
 */
export class UserPool extends pulumi.CustomResource {
    /**
     * Get an existing UserPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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

    /**
     * The configuration for AdminCreateUser requests.
     */
    public readonly adminCreateUserConfig!: pulumi.Output<outputs.cognito.UserPoolAdminCreateUserConfig>;
    /**
     * Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
     */
    public readonly aliasAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * The ARN of the user pool.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The attributes to be auto-verified. Possible values: email, phone_number.
     */
    public readonly autoVerifiedAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * The date the user pool was created.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The configuration for the user pool's device tracking.
     */
    public readonly deviceConfiguration!: pulumi.Output<outputs.cognito.UserPoolDeviceConfiguration | undefined>;
    /**
     * The Email Configuration.
     */
    public readonly emailConfiguration!: pulumi.Output<outputs.cognito.UserPoolEmailConfiguration | undefined>;
    /**
     * A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
     */
    public readonly emailVerificationMessage!: pulumi.Output<string>;
    /**
     * A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
     */
    public readonly emailVerificationSubject!: pulumi.Output<string>;
    /**
     * The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * A container for the AWS Lambda triggers associated with the user pool.
     */
    public readonly lambdaConfig!: pulumi.Output<outputs.cognito.UserPoolLambdaConfig>;
    /**
     * The date the user pool was last modified.
     */
    public /*out*/ readonly lastModifiedDate!: pulumi.Output<string>;
    /**
     * Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
     */
    public readonly mfaConfiguration!: pulumi.Output<string | undefined>;
    /**
     * The name of the attribute.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A container for information about the user pool password policy.
     */
    public readonly passwordPolicy!: pulumi.Output<outputs.cognito.UserPoolPasswordPolicy>;
    /**
     * A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
     */
    public readonly schemas!: pulumi.Output<outputs.cognito.UserPoolSchema[] | undefined>;
    /**
     * A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
     */
    public readonly smsAuthenticationMessage!: pulumi.Output<string | undefined>;
    /**
     * Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
     */
    public readonly smsConfiguration!: pulumi.Output<outputs.cognito.UserPoolSmsConfiguration>;
    /**
     * A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
     */
    public readonly smsVerificationMessage!: pulumi.Output<string>;
    /**
     * Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
     */
    public readonly softwareTokenMfaConfiguration!: pulumi.Output<outputs.cognito.UserPoolSoftwareTokenMfaConfiguration | undefined>;
    /**
     * A map of tags to assign to the User Pool.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Configuration block for user pool add-ons to enable user pool advanced security mode features.
     */
    public readonly userPoolAddOns!: pulumi.Output<outputs.cognito.UserPoolUserPoolAddOns | undefined>;
    /**
     * Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
     */
    public readonly usernameAttributes!: pulumi.Output<string[] | undefined>;
    /**
     * The Username Configuration.
     */
    public readonly usernameConfiguration!: pulumi.Output<outputs.cognito.UserPoolUsernameConfiguration | undefined>;
    /**
     * The verification message templates configuration.
     */
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
    /**
     * The configuration for AdminCreateUser requests.
     */
    readonly adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfig>;
    /**
     * Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
     */
    readonly aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of the user pool.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The attributes to be auto-verified. Possible values: email, phone_number.
     */
    readonly autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date the user pool was created.
     */
    readonly creationDate?: pulumi.Input<string>;
    /**
     * The configuration for the user pool's device tracking.
     */
    readonly deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfiguration>;
    /**
     * The Email Configuration.
     */
    readonly emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfiguration>;
    /**
     * A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
     */
    readonly emailVerificationMessage?: pulumi.Input<string>;
    /**
     * A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
     */
    readonly emailVerificationSubject?: pulumi.Input<string>;
    /**
     * The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * A container for the AWS Lambda triggers associated with the user pool.
     */
    readonly lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfig>;
    /**
     * The date the user pool was last modified.
     */
    readonly lastModifiedDate?: pulumi.Input<string>;
    /**
     * Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
     */
    readonly mfaConfiguration?: pulumi.Input<string>;
    /**
     * The name of the attribute.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A container for information about the user pool password policy.
     */
    readonly passwordPolicy?: pulumi.Input<inputs.cognito.UserPoolPasswordPolicy>;
    /**
     * A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
     */
    readonly schemas?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchema>[]>;
    /**
     * A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
     */
    readonly smsAuthenticationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
     */
    readonly smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfiguration>;
    /**
     * A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
     */
    readonly smsVerificationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
     */
    readonly softwareTokenMfaConfiguration?: pulumi.Input<inputs.cognito.UserPoolSoftwareTokenMfaConfiguration>;
    /**
     * A map of tags to assign to the User Pool.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Configuration block for user pool add-ons to enable user pool advanced security mode features.
     */
    readonly userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolUserPoolAddOns>;
    /**
     * Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
     */
    readonly usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Username Configuration.
     */
    readonly usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfiguration>;
    /**
     * The verification message templates configuration.
     */
    readonly verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplate>;
}

/**
 * The set of arguments for constructing a UserPool resource.
 */
export interface UserPoolArgs {
    /**
     * The configuration for AdminCreateUser requests.
     */
    readonly adminCreateUserConfig?: pulumi.Input<inputs.cognito.UserPoolAdminCreateUserConfig>;
    /**
     * Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
     */
    readonly aliasAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The attributes to be auto-verified. Possible values: email, phone_number.
     */
    readonly autoVerifiedAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for the user pool's device tracking.
     */
    readonly deviceConfiguration?: pulumi.Input<inputs.cognito.UserPoolDeviceConfiguration>;
    /**
     * The Email Configuration.
     */
    readonly emailConfiguration?: pulumi.Input<inputs.cognito.UserPoolEmailConfiguration>;
    /**
     * A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
     */
    readonly emailVerificationMessage?: pulumi.Input<string>;
    /**
     * A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
     */
    readonly emailVerificationSubject?: pulumi.Input<string>;
    /**
     * A container for the AWS Lambda triggers associated with the user pool.
     */
    readonly lambdaConfig?: pulumi.Input<inputs.cognito.UserPoolLambdaConfig>;
    /**
     * Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
     */
    readonly mfaConfiguration?: pulumi.Input<string>;
    /**
     * The name of the attribute.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A container for information about the user pool password policy.
     */
    readonly passwordPolicy?: pulumi.Input<inputs.cognito.UserPoolPasswordPolicy>;
    /**
     * A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
     */
    readonly schemas?: pulumi.Input<pulumi.Input<inputs.cognito.UserPoolSchema>[]>;
    /**
     * A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
     */
    readonly smsAuthenticationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
     */
    readonly smsConfiguration?: pulumi.Input<inputs.cognito.UserPoolSmsConfiguration>;
    /**
     * A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
     */
    readonly smsVerificationMessage?: pulumi.Input<string>;
    /**
     * Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
     */
    readonly softwareTokenMfaConfiguration?: pulumi.Input<inputs.cognito.UserPoolSoftwareTokenMfaConfiguration>;
    /**
     * A map of tags to assign to the User Pool.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Configuration block for user pool add-ons to enable user pool advanced security mode features.
     */
    readonly userPoolAddOns?: pulumi.Input<inputs.cognito.UserPoolUserPoolAddOns>;
    /**
     * Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
     */
    readonly usernameAttributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Username Configuration.
     */
    readonly usernameConfiguration?: pulumi.Input<inputs.cognito.UserPoolUsernameConfiguration>;
    /**
     * The verification message templates configuration.
     */
    readonly verificationMessageTemplate?: pulumi.Input<inputs.cognito.UserPoolVerificationMessageTemplate>;
}
