// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Cognito User Pool resource.
type UserPool struct {
	pulumi.CustomResourceState

	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig UserPoolAdminCreateUserConfigOutput `pulumi:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes pulumi.StringArrayOutput `pulumi:"aliasAttributes"`
	// The ARN of the user pool.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes pulumi.StringArrayOutput `pulumi:"autoVerifiedAttributes"`
	// The date the user pool was created.
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The configuration for the user pool's device tracking.
	DeviceConfiguration UserPoolDeviceConfigurationPtrOutput `pulumi:"deviceConfiguration"`
	// The Email Configuration.
	EmailConfiguration UserPoolEmailConfigurationPtrOutput `pulumi:"emailConfiguration"`
	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage pulumi.StringOutput `pulumi:"emailVerificationMessage"`
	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject pulumi.StringOutput `pulumi:"emailVerificationSubject"`
	// The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig UserPoolLambdaConfigOutput `pulumi:"lambdaConfig"`
	// The date the user pool was last modified.
	LastModifiedDate pulumi.StringOutput `pulumi:"lastModifiedDate"`
	// Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
	MfaConfiguration pulumi.StringPtrOutput `pulumi:"mfaConfiguration"`
	// The name of the attribute.
	Name pulumi.StringOutput `pulumi:"name"`
	// A container for information about the user pool password policy.
	PasswordPolicy UserPoolPasswordPolicyOutput `pulumi:"passwordPolicy"`
	// A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
	Schemas UserPoolSchemaArrayOutput `pulumi:"schemas"`
	// A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
	SmsAuthenticationMessage pulumi.StringPtrOutput `pulumi:"smsAuthenticationMessage"`
	// Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
	SmsConfiguration UserPoolSmsConfigurationOutput `pulumi:"smsConfiguration"`
	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage pulumi.StringOutput `pulumi:"smsVerificationMessage"`
	// Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
	SoftwareTokenMfaConfiguration UserPoolSoftwareTokenMfaConfigurationPtrOutput `pulumi:"softwareTokenMfaConfiguration"`
	// A mapping of tags to assign to the User Pool.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns UserPoolUserPoolAddOnsPtrOutput `pulumi:"userPoolAddOns"`
	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes pulumi.StringArrayOutput `pulumi:"usernameAttributes"`
	// The Username Configuration.
	UsernameConfiguration UserPoolUsernameConfigurationPtrOutput `pulumi:"usernameConfiguration"`
	// The verification message templates configuration.
	VerificationMessageTemplate UserPoolVerificationMessageTemplateOutput `pulumi:"verificationMessageTemplate"`
}

// NewUserPool registers a new resource with the given unique name, arguments, and options.
func NewUserPool(ctx *pulumi.Context,
	name string, args *UserPoolArgs, opts ...pulumi.ResourceOption) (*UserPool, error) {
	if args == nil {
		args = &UserPoolArgs{}
	}
	var resource UserPool
	err := ctx.RegisterResource("aws:cognito/userPool:UserPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPool gets an existing UserPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPoolState, opts ...pulumi.ResourceOption) (*UserPool, error) {
	var resource UserPool
	err := ctx.ReadResource("aws:cognito/userPool:UserPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPool resources.
type userPoolState struct {
	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *UserPoolAdminCreateUserConfig `pulumi:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes []string `pulumi:"aliasAttributes"`
	// The ARN of the user pool.
	Arn *string `pulumi:"arn"`
	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes []string `pulumi:"autoVerifiedAttributes"`
	// The date the user pool was created.
	CreationDate *string `pulumi:"creationDate"`
	// The configuration for the user pool's device tracking.
	DeviceConfiguration *UserPoolDeviceConfiguration `pulumi:"deviceConfiguration"`
	// The Email Configuration.
	EmailConfiguration *UserPoolEmailConfiguration `pulumi:"emailConfiguration"`
	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage *string `pulumi:"emailVerificationMessage"`
	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject *string `pulumi:"emailVerificationSubject"`
	// The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
	Endpoint *string `pulumi:"endpoint"`
	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig *UserPoolLambdaConfig `pulumi:"lambdaConfig"`
	// The date the user pool was last modified.
	LastModifiedDate *string `pulumi:"lastModifiedDate"`
	// Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
	MfaConfiguration *string `pulumi:"mfaConfiguration"`
	// The name of the attribute.
	Name *string `pulumi:"name"`
	// A container for information about the user pool password policy.
	PasswordPolicy *UserPoolPasswordPolicy `pulumi:"passwordPolicy"`
	// A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
	Schemas []UserPoolSchema `pulumi:"schemas"`
	// A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
	SmsAuthenticationMessage *string `pulumi:"smsAuthenticationMessage"`
	// Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
	SmsConfiguration *UserPoolSmsConfiguration `pulumi:"smsConfiguration"`
	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage *string `pulumi:"smsVerificationMessage"`
	// Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
	SoftwareTokenMfaConfiguration *UserPoolSoftwareTokenMfaConfiguration `pulumi:"softwareTokenMfaConfiguration"`
	// A mapping of tags to assign to the User Pool.
	Tags map[string]interface{} `pulumi:"tags"`
	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns *UserPoolUserPoolAddOns `pulumi:"userPoolAddOns"`
	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes []string `pulumi:"usernameAttributes"`
	// The Username Configuration.
	UsernameConfiguration *UserPoolUsernameConfiguration `pulumi:"usernameConfiguration"`
	// The verification message templates configuration.
	VerificationMessageTemplate *UserPoolVerificationMessageTemplate `pulumi:"verificationMessageTemplate"`
}

type UserPoolState struct {
	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig UserPoolAdminCreateUserConfigPtrInput
	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes pulumi.StringArrayInput
	// The ARN of the user pool.
	Arn pulumi.StringPtrInput
	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes pulumi.StringArrayInput
	// The date the user pool was created.
	CreationDate pulumi.StringPtrInput
	// The configuration for the user pool's device tracking.
	DeviceConfiguration UserPoolDeviceConfigurationPtrInput
	// The Email Configuration.
	EmailConfiguration UserPoolEmailConfigurationPtrInput
	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage pulumi.StringPtrInput
	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject pulumi.StringPtrInput
	// The endpoint name of the user pool. Example format: cognito-idp.REGION.amazonaws.com/xxxx_yyyyy
	Endpoint pulumi.StringPtrInput
	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig UserPoolLambdaConfigPtrInput
	// The date the user pool was last modified.
	LastModifiedDate pulumi.StringPtrInput
	// Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
	MfaConfiguration pulumi.StringPtrInput
	// The name of the attribute.
	Name pulumi.StringPtrInput
	// A container for information about the user pool password policy.
	PasswordPolicy UserPoolPasswordPolicyPtrInput
	// A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
	Schemas UserPoolSchemaArrayInput
	// A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
	SmsAuthenticationMessage pulumi.StringPtrInput
	// Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
	SmsConfiguration UserPoolSmsConfigurationPtrInput
	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage pulumi.StringPtrInput
	// Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
	SoftwareTokenMfaConfiguration UserPoolSoftwareTokenMfaConfigurationPtrInput
	// A mapping of tags to assign to the User Pool.
	Tags pulumi.MapInput
	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns UserPoolUserPoolAddOnsPtrInput
	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes pulumi.StringArrayInput
	// The Username Configuration.
	UsernameConfiguration UserPoolUsernameConfigurationPtrInput
	// The verification message templates configuration.
	VerificationMessageTemplate UserPoolVerificationMessageTemplatePtrInput
}

func (UserPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolState)(nil)).Elem()
}

type userPoolArgs struct {
	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig *UserPoolAdminCreateUserConfigArgs `pulumi:"adminCreateUserConfig"`
	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes []string `pulumi:"aliasAttributes"`
	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes []string `pulumi:"autoVerifiedAttributes"`
	// The configuration for the user pool's device tracking.
	DeviceConfiguration *UserPoolDeviceConfigurationArgs `pulumi:"deviceConfiguration"`
	// The Email Configuration.
	EmailConfiguration *UserPoolEmailConfigurationArgs `pulumi:"emailConfiguration"`
	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage *string `pulumi:"emailVerificationMessage"`
	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject *string `pulumi:"emailVerificationSubject"`
	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig *UserPoolLambdaConfigArgs `pulumi:"lambdaConfig"`
	// Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
	MfaConfiguration *string `pulumi:"mfaConfiguration"`
	// The name of the attribute.
	Name *string `pulumi:"name"`
	// A container for information about the user pool password policy.
	PasswordPolicy *UserPoolPasswordPolicyArgs `pulumi:"passwordPolicy"`
	// A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
	Schemas []UserPoolSchemaArgs `pulumi:"schemas"`
	// A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
	SmsAuthenticationMessage *string `pulumi:"smsAuthenticationMessage"`
	// Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
	SmsConfiguration *UserPoolSmsConfigurationArgs `pulumi:"smsConfiguration"`
	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage *string `pulumi:"smsVerificationMessage"`
	// Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
	SoftwareTokenMfaConfiguration *UserPoolSoftwareTokenMfaConfigurationArgs `pulumi:"softwareTokenMfaConfiguration"`
	// A mapping of tags to assign to the User Pool.
	Tags map[string]interface{} `pulumi:"tags"`
	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns *UserPoolUserPoolAddOnsArgs `pulumi:"userPoolAddOns"`
	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes []string `pulumi:"usernameAttributes"`
	// The Username Configuration.
	UsernameConfiguration *UserPoolUsernameConfigurationArgs `pulumi:"usernameConfiguration"`
	// The verification message templates configuration.
	VerificationMessageTemplate *UserPoolVerificationMessageTemplateArgs `pulumi:"verificationMessageTemplate"`
}

// The set of arguments for constructing a UserPool resource.
type UserPoolArgs struct {
	// The configuration for AdminCreateUser requests.
	AdminCreateUserConfig UserPoolAdminCreateUserConfigArgsPtrInput
	// Attributes supported as an alias for this user pool. Possible values: phone_number, email, or preferred_username. Conflicts with `usernameAttributes`.
	AliasAttributes pulumi.StringArrayInput
	// The attributes to be auto-verified. Possible values: email, phone_number.
	AutoVerifiedAttributes pulumi.StringArrayInput
	// The configuration for the user pool's device tracking.
	DeviceConfiguration UserPoolDeviceConfigurationArgsPtrInput
	// The Email Configuration.
	EmailConfiguration UserPoolEmailConfigurationArgsPtrInput
	// A string representing the email verification message. Conflicts with `verificationMessageTemplate` configuration block `emailMessage` argument.
	EmailVerificationMessage pulumi.StringPtrInput
	// A string representing the email verification subject. Conflicts with `verificationMessageTemplate` configuration block `emailSubject` argument.
	EmailVerificationSubject pulumi.StringPtrInput
	// A container for the AWS Lambda triggers associated with the user pool.
	LambdaConfig UserPoolLambdaConfigArgsPtrInput
	// Multi-Factor Authentication (MFA) configuration for the User Pool. Defaults of `OFF`. Valid values:
	MfaConfiguration pulumi.StringPtrInput
	// The name of the attribute.
	Name pulumi.StringPtrInput
	// A container for information about the user pool password policy.
	PasswordPolicy UserPoolPasswordPolicyArgsPtrInput
	// A container with the schema attributes of a user pool. Schema attributes from the [standard attribute set](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#cognito-user-pools-standard-attributes) only need to be specified if they are different from the default configuration. Maximum of 50 attributes.
	Schemas UserPoolSchemaArgsArrayInput
	// A string representing the SMS authentication message. The message must contain the `{####}` placeholder, which will be replaced with the code.
	SmsAuthenticationMessage pulumi.StringPtrInput
	// Configuration block for Short Message Service (SMS) settings. Detailed below. These settings apply to SMS user verification and SMS Multi-Factor Authentication (MFA). Due to Cognito API restrictions, the SMS configuration cannot be removed without recreating the Cognito User Pool. For user data safety, this resource will ignore the removal of this configuration by disabling drift detection. To force resource recreation after this configuration has been applied, see the [`up` command and use --replace](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
	SmsConfiguration UserPoolSmsConfigurationArgsPtrInput
	// A string representing the SMS verification message. Conflicts with `verificationMessageTemplate` configuration block `smsMessage` argument.
	SmsVerificationMessage pulumi.StringPtrInput
	// Configuration block for software token Mult-Factor Authentication (MFA) settings. Detailed below.
	SoftwareTokenMfaConfiguration UserPoolSoftwareTokenMfaConfigurationArgsPtrInput
	// A mapping of tags to assign to the User Pool.
	Tags pulumi.MapInput
	// Configuration block for user pool add-ons to enable user pool advanced security mode features.
	UserPoolAddOns UserPoolUserPoolAddOnsArgsPtrInput
	// Specifies whether email addresses or phone numbers can be specified as usernames when a user signs up. Conflicts with `aliasAttributes`.
	UsernameAttributes pulumi.StringArrayInput
	// The Username Configuration.
	UsernameConfiguration UserPoolUsernameConfigurationArgsPtrInput
	// The verification message templates configuration.
	VerificationMessageTemplate UserPoolVerificationMessageTemplateArgsPtrInput
}

func (UserPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolArgs)(nil)).Elem()
}
