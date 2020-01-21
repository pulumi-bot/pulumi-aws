// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package WebhookAuthenticationConfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type WebhookAuthenticationConfiguration struct {
	AllowedIpRange *string `pulumi:"allowedIpRange"`
	SecretToken *string `pulumi:"secretToken"`
}

type WebhookAuthenticationConfigurationInput interface {
	pulumi.Input

	ToWebhookAuthenticationConfigurationOutput() WebhookAuthenticationConfigurationOutput
	ToWebhookAuthenticationConfigurationOutputWithContext(context.Context) WebhookAuthenticationConfigurationOutput
}

type WebhookAuthenticationConfigurationArgs struct {
	AllowedIpRange pulumi.StringPtrInput `pulumi:"allowedIpRange"`
	SecretToken pulumi.StringPtrInput `pulumi:"secretToken"`
}

func (WebhookAuthenticationConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookAuthenticationConfiguration)(nil)).Elem()
}

func (i WebhookAuthenticationConfigurationArgs) ToWebhookAuthenticationConfigurationOutput() WebhookAuthenticationConfigurationOutput {
	return i.ToWebhookAuthenticationConfigurationOutputWithContext(context.Background())
}

func (i WebhookAuthenticationConfigurationArgs) ToWebhookAuthenticationConfigurationOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookAuthenticationConfigurationOutput)
}

func (i WebhookAuthenticationConfigurationArgs) ToWebhookAuthenticationConfigurationPtrOutput() WebhookAuthenticationConfigurationPtrOutput {
	return i.ToWebhookAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (i WebhookAuthenticationConfigurationArgs) ToWebhookAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookAuthenticationConfigurationOutput).ToWebhookAuthenticationConfigurationPtrOutputWithContext(ctx)
}

type WebhookAuthenticationConfigurationPtrInput interface {
	pulumi.Input

	ToWebhookAuthenticationConfigurationPtrOutput() WebhookAuthenticationConfigurationPtrOutput
	ToWebhookAuthenticationConfigurationPtrOutputWithContext(context.Context) WebhookAuthenticationConfigurationPtrOutput
}

type webhookAuthenticationConfigurationPtrType WebhookAuthenticationConfigurationArgs

func WebhookAuthenticationConfigurationPtr(v *WebhookAuthenticationConfigurationArgs) WebhookAuthenticationConfigurationPtrInput {	return (*webhookAuthenticationConfigurationPtrType)(v)
}

func (*webhookAuthenticationConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookAuthenticationConfiguration)(nil)).Elem()
}

func (i *webhookAuthenticationConfigurationPtrType) ToWebhookAuthenticationConfigurationPtrOutput() WebhookAuthenticationConfigurationPtrOutput {
	return i.ToWebhookAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (i *webhookAuthenticationConfigurationPtrType) ToWebhookAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookAuthenticationConfigurationPtrOutput)
}

type WebhookAuthenticationConfigurationOutput struct { *pulumi.OutputState }

func (WebhookAuthenticationConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookAuthenticationConfiguration)(nil)).Elem()
}

func (o WebhookAuthenticationConfigurationOutput) ToWebhookAuthenticationConfigurationOutput() WebhookAuthenticationConfigurationOutput {
	return o
}

func (o WebhookAuthenticationConfigurationOutput) ToWebhookAuthenticationConfigurationOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationOutput {
	return o
}

func (o WebhookAuthenticationConfigurationOutput) ToWebhookAuthenticationConfigurationPtrOutput() WebhookAuthenticationConfigurationPtrOutput {
	return o.ToWebhookAuthenticationConfigurationPtrOutputWithContext(context.Background())
}

func (o WebhookAuthenticationConfigurationOutput) ToWebhookAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationPtrOutput {
	return o.ApplyT(func(v WebhookAuthenticationConfiguration) *WebhookAuthenticationConfiguration {
		return &v
	}).(WebhookAuthenticationConfigurationPtrOutput)
}
func (o WebhookAuthenticationConfigurationOutput) AllowedIpRange() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookAuthenticationConfiguration) *string { return v.AllowedIpRange }).(pulumi.StringPtrOutput)
}

func (o WebhookAuthenticationConfigurationOutput) SecretToken() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookAuthenticationConfiguration) *string { return v.SecretToken }).(pulumi.StringPtrOutput)
}

type WebhookAuthenticationConfigurationPtrOutput struct { *pulumi.OutputState}

func (WebhookAuthenticationConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookAuthenticationConfiguration)(nil)).Elem()
}

func (o WebhookAuthenticationConfigurationPtrOutput) ToWebhookAuthenticationConfigurationPtrOutput() WebhookAuthenticationConfigurationPtrOutput {
	return o
}

func (o WebhookAuthenticationConfigurationPtrOutput) ToWebhookAuthenticationConfigurationPtrOutputWithContext(ctx context.Context) WebhookAuthenticationConfigurationPtrOutput {
	return o
}

func (o WebhookAuthenticationConfigurationPtrOutput) Elem() WebhookAuthenticationConfigurationOutput {
	return o.ApplyT(func (v *WebhookAuthenticationConfiguration) WebhookAuthenticationConfiguration { return *v }).(WebhookAuthenticationConfigurationOutput)
}

func (o WebhookAuthenticationConfigurationPtrOutput) AllowedIpRange() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookAuthenticationConfiguration) *string { return v.AllowedIpRange }).(pulumi.StringPtrOutput)
}

func (o WebhookAuthenticationConfigurationPtrOutput) SecretToken() pulumi.StringPtrOutput {
	return o.ApplyT(func (v WebhookAuthenticationConfiguration) *string { return v.SecretToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebhookAuthenticationConfigurationOutput{})
	pulumi.RegisterOutputType(WebhookAuthenticationConfigurationPtrOutput{})
}
