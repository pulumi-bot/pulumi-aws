// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type UserPool struct {
	pulumi.CustomResourceState

	AdminCreateUserConfig         UserPoolAdminCreateUserConfigOutput            `pulumi:"adminCreateUserConfig"`
	AliasAttributes               pulumi.StringArrayOutput                       `pulumi:"aliasAttributes"`
	Arn                           pulumi.StringOutput                            `pulumi:"arn"`
	AutoVerifiedAttributes        pulumi.StringArrayOutput                       `pulumi:"autoVerifiedAttributes"`
	CreationDate                  pulumi.StringOutput                            `pulumi:"creationDate"`
	DeviceConfiguration           UserPoolDeviceConfigurationPtrOutput           `pulumi:"deviceConfiguration"`
	EmailConfiguration            UserPoolEmailConfigurationPtrOutput            `pulumi:"emailConfiguration"`
	EmailVerificationMessage      pulumi.StringOutput                            `pulumi:"emailVerificationMessage"`
	EmailVerificationSubject      pulumi.StringOutput                            `pulumi:"emailVerificationSubject"`
	Endpoint                      pulumi.StringOutput                            `pulumi:"endpoint"`
	LambdaConfig                  UserPoolLambdaConfigOutput                     `pulumi:"lambdaConfig"`
	LastModifiedDate              pulumi.StringOutput                            `pulumi:"lastModifiedDate"`
	MfaConfiguration              pulumi.StringPtrOutput                         `pulumi:"mfaConfiguration"`
	Name                          pulumi.StringOutput                            `pulumi:"name"`
	PasswordPolicy                UserPoolPasswordPolicyOutput                   `pulumi:"passwordPolicy"`
	Schemas                       UserPoolSchemaArrayOutput                      `pulumi:"schemas"`
	SmsAuthenticationMessage      pulumi.StringPtrOutput                         `pulumi:"smsAuthenticationMessage"`
	SmsConfiguration              UserPoolSmsConfigurationOutput                 `pulumi:"smsConfiguration"`
	SmsVerificationMessage        pulumi.StringOutput                            `pulumi:"smsVerificationMessage"`
	SoftwareTokenMfaConfiguration UserPoolSoftwareTokenMfaConfigurationPtrOutput `pulumi:"softwareTokenMfaConfiguration"`
	Tags                          pulumi.StringMapOutput                         `pulumi:"tags"`
	UserPoolAddOns                UserPoolUserPoolAddOnsPtrOutput                `pulumi:"userPoolAddOns"`
	UsernameAttributes            pulumi.StringArrayOutput                       `pulumi:"usernameAttributes"`
	UsernameConfiguration         UserPoolUsernameConfigurationPtrOutput         `pulumi:"usernameConfiguration"`
	VerificationMessageTemplate   UserPoolVerificationMessageTemplateOutput      `pulumi:"verificationMessageTemplate"`
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
	AdminCreateUserConfig         *UserPoolAdminCreateUserConfig         `pulumi:"adminCreateUserConfig"`
	AliasAttributes               []string                               `pulumi:"aliasAttributes"`
	Arn                           *string                                `pulumi:"arn"`
	AutoVerifiedAttributes        []string                               `pulumi:"autoVerifiedAttributes"`
	CreationDate                  *string                                `pulumi:"creationDate"`
	DeviceConfiguration           *UserPoolDeviceConfiguration           `pulumi:"deviceConfiguration"`
	EmailConfiguration            *UserPoolEmailConfiguration            `pulumi:"emailConfiguration"`
	EmailVerificationMessage      *string                                `pulumi:"emailVerificationMessage"`
	EmailVerificationSubject      *string                                `pulumi:"emailVerificationSubject"`
	Endpoint                      *string                                `pulumi:"endpoint"`
	LambdaConfig                  *UserPoolLambdaConfig                  `pulumi:"lambdaConfig"`
	LastModifiedDate              *string                                `pulumi:"lastModifiedDate"`
	MfaConfiguration              *string                                `pulumi:"mfaConfiguration"`
	Name                          *string                                `pulumi:"name"`
	PasswordPolicy                *UserPoolPasswordPolicy                `pulumi:"passwordPolicy"`
	Schemas                       []UserPoolSchema                       `pulumi:"schemas"`
	SmsAuthenticationMessage      *string                                `pulumi:"smsAuthenticationMessage"`
	SmsConfiguration              *UserPoolSmsConfiguration              `pulumi:"smsConfiguration"`
	SmsVerificationMessage        *string                                `pulumi:"smsVerificationMessage"`
	SoftwareTokenMfaConfiguration *UserPoolSoftwareTokenMfaConfiguration `pulumi:"softwareTokenMfaConfiguration"`
	Tags                          map[string]string                      `pulumi:"tags"`
	UserPoolAddOns                *UserPoolUserPoolAddOns                `pulumi:"userPoolAddOns"`
	UsernameAttributes            []string                               `pulumi:"usernameAttributes"`
	UsernameConfiguration         *UserPoolUsernameConfiguration         `pulumi:"usernameConfiguration"`
	VerificationMessageTemplate   *UserPoolVerificationMessageTemplate   `pulumi:"verificationMessageTemplate"`
}

type UserPoolState struct {
	AdminCreateUserConfig         UserPoolAdminCreateUserConfigPtrInput
	AliasAttributes               pulumi.StringArrayInput
	Arn                           pulumi.StringPtrInput
	AutoVerifiedAttributes        pulumi.StringArrayInput
	CreationDate                  pulumi.StringPtrInput
	DeviceConfiguration           UserPoolDeviceConfigurationPtrInput
	EmailConfiguration            UserPoolEmailConfigurationPtrInput
	EmailVerificationMessage      pulumi.StringPtrInput
	EmailVerificationSubject      pulumi.StringPtrInput
	Endpoint                      pulumi.StringPtrInput
	LambdaConfig                  UserPoolLambdaConfigPtrInput
	LastModifiedDate              pulumi.StringPtrInput
	MfaConfiguration              pulumi.StringPtrInput
	Name                          pulumi.StringPtrInput
	PasswordPolicy                UserPoolPasswordPolicyPtrInput
	Schemas                       UserPoolSchemaArrayInput
	SmsAuthenticationMessage      pulumi.StringPtrInput
	SmsConfiguration              UserPoolSmsConfigurationPtrInput
	SmsVerificationMessage        pulumi.StringPtrInput
	SoftwareTokenMfaConfiguration UserPoolSoftwareTokenMfaConfigurationPtrInput
	Tags                          pulumi.StringMapInput
	UserPoolAddOns                UserPoolUserPoolAddOnsPtrInput
	UsernameAttributes            pulumi.StringArrayInput
	UsernameConfiguration         UserPoolUsernameConfigurationPtrInput
	VerificationMessageTemplate   UserPoolVerificationMessageTemplatePtrInput
}

func (UserPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolState)(nil)).Elem()
}

type userPoolArgs struct {
	AdminCreateUserConfig         *UserPoolAdminCreateUserConfig         `pulumi:"adminCreateUserConfig"`
	AliasAttributes               []string                               `pulumi:"aliasAttributes"`
	AutoVerifiedAttributes        []string                               `pulumi:"autoVerifiedAttributes"`
	DeviceConfiguration           *UserPoolDeviceConfiguration           `pulumi:"deviceConfiguration"`
	EmailConfiguration            *UserPoolEmailConfiguration            `pulumi:"emailConfiguration"`
	EmailVerificationMessage      *string                                `pulumi:"emailVerificationMessage"`
	EmailVerificationSubject      *string                                `pulumi:"emailVerificationSubject"`
	LambdaConfig                  *UserPoolLambdaConfig                  `pulumi:"lambdaConfig"`
	MfaConfiguration              *string                                `pulumi:"mfaConfiguration"`
	Name                          *string                                `pulumi:"name"`
	PasswordPolicy                *UserPoolPasswordPolicy                `pulumi:"passwordPolicy"`
	Schemas                       []UserPoolSchema                       `pulumi:"schemas"`
	SmsAuthenticationMessage      *string                                `pulumi:"smsAuthenticationMessage"`
	SmsConfiguration              *UserPoolSmsConfiguration              `pulumi:"smsConfiguration"`
	SmsVerificationMessage        *string                                `pulumi:"smsVerificationMessage"`
	SoftwareTokenMfaConfiguration *UserPoolSoftwareTokenMfaConfiguration `pulumi:"softwareTokenMfaConfiguration"`
	Tags                          map[string]string                      `pulumi:"tags"`
	UserPoolAddOns                *UserPoolUserPoolAddOns                `pulumi:"userPoolAddOns"`
	UsernameAttributes            []string                               `pulumi:"usernameAttributes"`
	UsernameConfiguration         *UserPoolUsernameConfiguration         `pulumi:"usernameConfiguration"`
	VerificationMessageTemplate   *UserPoolVerificationMessageTemplate   `pulumi:"verificationMessageTemplate"`
}

// The set of arguments for constructing a UserPool resource.
type UserPoolArgs struct {
	AdminCreateUserConfig         UserPoolAdminCreateUserConfigPtrInput
	AliasAttributes               pulumi.StringArrayInput
	AutoVerifiedAttributes        pulumi.StringArrayInput
	DeviceConfiguration           UserPoolDeviceConfigurationPtrInput
	EmailConfiguration            UserPoolEmailConfigurationPtrInput
	EmailVerificationMessage      pulumi.StringPtrInput
	EmailVerificationSubject      pulumi.StringPtrInput
	LambdaConfig                  UserPoolLambdaConfigPtrInput
	MfaConfiguration              pulumi.StringPtrInput
	Name                          pulumi.StringPtrInput
	PasswordPolicy                UserPoolPasswordPolicyPtrInput
	Schemas                       UserPoolSchemaArrayInput
	SmsAuthenticationMessage      pulumi.StringPtrInput
	SmsConfiguration              UserPoolSmsConfigurationPtrInput
	SmsVerificationMessage        pulumi.StringPtrInput
	SoftwareTokenMfaConfiguration UserPoolSoftwareTokenMfaConfigurationPtrInput
	Tags                          pulumi.StringMapInput
	UserPoolAddOns                UserPoolUserPoolAddOnsPtrInput
	UsernameAttributes            pulumi.StringArrayInput
	UsernameConfiguration         UserPoolUsernameConfigurationPtrInput
	VerificationMessageTemplate   UserPoolVerificationMessageTemplatePtrInput
}

func (UserPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPoolArgs)(nil)).Elem()
}
