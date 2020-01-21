// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package UserPoolAdminCreateUserConfigInviteMessageTemplate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type UserPoolAdminCreateUserConfigInviteMessageTemplate struct {
	// The email message template. Must contain the `{####}` placeholder. Conflicts with `emailVerificationMessage` argument.
	EmailMessage *string `pulumi:"emailMessage"`
	// The subject line for the email message template. Conflicts with `emailVerificationSubject` argument.
	EmailSubject *string `pulumi:"emailSubject"`
	// The SMS message template. Must contain the `{####}` placeholder. Conflicts with `smsVerificationMessage` argument.
	SmsMessage *string `pulumi:"smsMessage"`
}

type UserPoolAdminCreateUserConfigInviteMessageTemplateInput interface {
	pulumi.Input

	ToUserPoolAdminCreateUserConfigInviteMessageTemplateOutput() UserPoolAdminCreateUserConfigInviteMessageTemplateOutput
	ToUserPoolAdminCreateUserConfigInviteMessageTemplateOutputWithContext(context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplateOutput
}

type UserPoolAdminCreateUserConfigInviteMessageTemplateArgs struct {
	// The email message template. Must contain the `{####}` placeholder. Conflicts with `emailVerificationMessage` argument.
	EmailMessage pulumi.StringPtrInput `pulumi:"emailMessage"`
	// The subject line for the email message template. Conflicts with `emailVerificationSubject` argument.
	EmailSubject pulumi.StringPtrInput `pulumi:"emailSubject"`
	// The SMS message template. Must contain the `{####}` placeholder. Conflicts with `smsVerificationMessage` argument.
	SmsMessage pulumi.StringPtrInput `pulumi:"smsMessage"`
}

func (UserPoolAdminCreateUserConfigInviteMessageTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolAdminCreateUserConfigInviteMessageTemplate)(nil)).Elem()
}

func (i UserPoolAdminCreateUserConfigInviteMessageTemplateArgs) ToUserPoolAdminCreateUserConfigInviteMessageTemplateOutput() UserPoolAdminCreateUserConfigInviteMessageTemplateOutput {
	return i.ToUserPoolAdminCreateUserConfigInviteMessageTemplateOutputWithContext(context.Background())
}

func (i UserPoolAdminCreateUserConfigInviteMessageTemplateArgs) ToUserPoolAdminCreateUserConfigInviteMessageTemplateOutputWithContext(ctx context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolAdminCreateUserConfigInviteMessageTemplateOutput)
}

func (i UserPoolAdminCreateUserConfigInviteMessageTemplateArgs) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput() UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return i.ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(context.Background())
}

func (i UserPoolAdminCreateUserConfigInviteMessageTemplateArgs) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(ctx context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolAdminCreateUserConfigInviteMessageTemplateOutput).ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(ctx)
}

type UserPoolAdminCreateUserConfigInviteMessageTemplatePtrInput interface {
	pulumi.Input

	ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput() UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput
	ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput
}

type userPoolAdminCreateUserConfigInviteMessageTemplatePtrType UserPoolAdminCreateUserConfigInviteMessageTemplateArgs

func UserPoolAdminCreateUserConfigInviteMessageTemplatePtr(v *UserPoolAdminCreateUserConfigInviteMessageTemplateArgs) UserPoolAdminCreateUserConfigInviteMessageTemplatePtrInput {	return (*userPoolAdminCreateUserConfigInviteMessageTemplatePtrType)(v)
}

func (*userPoolAdminCreateUserConfigInviteMessageTemplatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPoolAdminCreateUserConfigInviteMessageTemplate)(nil)).Elem()
}

func (i *userPoolAdminCreateUserConfigInviteMessageTemplatePtrType) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput() UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return i.ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(context.Background())
}

func (i *userPoolAdminCreateUserConfigInviteMessageTemplatePtrType) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(ctx context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput)
}

type UserPoolAdminCreateUserConfigInviteMessageTemplateOutput struct { *pulumi.OutputState }

func (UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolAdminCreateUserConfigInviteMessageTemplate)(nil)).Elem()
}

func (o UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) ToUserPoolAdminCreateUserConfigInviteMessageTemplateOutput() UserPoolAdminCreateUserConfigInviteMessageTemplateOutput {
	return o
}

func (o UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) ToUserPoolAdminCreateUserConfigInviteMessageTemplateOutputWithContext(ctx context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplateOutput {
	return o
}

func (o UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput() UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return o.ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(context.Background())
}

func (o UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(ctx context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return o.ApplyT(func(v UserPoolAdminCreateUserConfigInviteMessageTemplate) *UserPoolAdminCreateUserConfigInviteMessageTemplate {
		return &v
	}).(UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput)
}
// The email message template. Must contain the `{####}` placeholder. Conflicts with `emailVerificationMessage` argument.
func (o UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) EmailMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func (v UserPoolAdminCreateUserConfigInviteMessageTemplate) *string { return v.EmailMessage }).(pulumi.StringPtrOutput)
}

// The subject line for the email message template. Conflicts with `emailVerificationSubject` argument.
func (o UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) EmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func (v UserPoolAdminCreateUserConfigInviteMessageTemplate) *string { return v.EmailSubject }).(pulumi.StringPtrOutput)
}

// The SMS message template. Must contain the `{####}` placeholder. Conflicts with `smsVerificationMessage` argument.
func (o UserPoolAdminCreateUserConfigInviteMessageTemplateOutput) SmsMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func (v UserPoolAdminCreateUserConfigInviteMessageTemplate) *string { return v.SmsMessage }).(pulumi.StringPtrOutput)
}

type UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput struct { *pulumi.OutputState}

func (UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPoolAdminCreateUserConfigInviteMessageTemplate)(nil)).Elem()
}

func (o UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput() UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return o
}

func (o UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput) ToUserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutputWithContext(ctx context.Context) UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput {
	return o
}

func (o UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput) Elem() UserPoolAdminCreateUserConfigInviteMessageTemplateOutput {
	return o.ApplyT(func (v *UserPoolAdminCreateUserConfigInviteMessageTemplate) UserPoolAdminCreateUserConfigInviteMessageTemplate { return *v }).(UserPoolAdminCreateUserConfigInviteMessageTemplateOutput)
}

// The email message template. Must contain the `{####}` placeholder. Conflicts with `emailVerificationMessage` argument.
func (o UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput) EmailMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func (v UserPoolAdminCreateUserConfigInviteMessageTemplate) *string { return v.EmailMessage }).(pulumi.StringPtrOutput)
}

// The subject line for the email message template. Conflicts with `emailVerificationSubject` argument.
func (o UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput) EmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func (v UserPoolAdminCreateUserConfigInviteMessageTemplate) *string { return v.EmailSubject }).(pulumi.StringPtrOutput)
}

// The SMS message template. Must contain the `{####}` placeholder. Conflicts with `smsVerificationMessage` argument.
func (o UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput) SmsMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func (v UserPoolAdminCreateUserConfigInviteMessageTemplate) *string { return v.SmsMessage }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(UserPoolAdminCreateUserConfigInviteMessageTemplateOutput{})
	pulumi.RegisterOutputType(UserPoolAdminCreateUserConfigInviteMessageTemplatePtrOutput{})
}
