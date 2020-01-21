// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package GatewaySmbActiveDirectorySettings

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GatewaySmbActiveDirectorySettings struct {
	// The name of the domain that you want the gateway to join.
	DomainName string `pulumi:"domainName"`
	// The password of the user who has permission to add the gateway to the Active Directory domain.
	Password string `pulumi:"password"`
	// The user name of user who has permission to add the gateway to the Active Directory domain.
	Username string `pulumi:"username"`
}

type GatewaySmbActiveDirectorySettingsInput interface {
	pulumi.Input

	ToGatewaySmbActiveDirectorySettingsOutput() GatewaySmbActiveDirectorySettingsOutput
	ToGatewaySmbActiveDirectorySettingsOutputWithContext(context.Context) GatewaySmbActiveDirectorySettingsOutput
}

type GatewaySmbActiveDirectorySettingsArgs struct {
	// The name of the domain that you want the gateway to join.
	DomainName pulumi.StringInput `pulumi:"domainName"`
	// The password of the user who has permission to add the gateway to the Active Directory domain.
	Password pulumi.StringInput `pulumi:"password"`
	// The user name of user who has permission to add the gateway to the Active Directory domain.
	Username pulumi.StringInput `pulumi:"username"`
}

func (GatewaySmbActiveDirectorySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewaySmbActiveDirectorySettings)(nil)).Elem()
}

func (i GatewaySmbActiveDirectorySettingsArgs) ToGatewaySmbActiveDirectorySettingsOutput() GatewaySmbActiveDirectorySettingsOutput {
	return i.ToGatewaySmbActiveDirectorySettingsOutputWithContext(context.Background())
}

func (i GatewaySmbActiveDirectorySettingsArgs) ToGatewaySmbActiveDirectorySettingsOutputWithContext(ctx context.Context) GatewaySmbActiveDirectorySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySmbActiveDirectorySettingsOutput)
}

func (i GatewaySmbActiveDirectorySettingsArgs) ToGatewaySmbActiveDirectorySettingsPtrOutput() GatewaySmbActiveDirectorySettingsPtrOutput {
	return i.ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(context.Background())
}

func (i GatewaySmbActiveDirectorySettingsArgs) ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) GatewaySmbActiveDirectorySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySmbActiveDirectorySettingsOutput).ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(ctx)
}

type GatewaySmbActiveDirectorySettingsPtrInput interface {
	pulumi.Input

	ToGatewaySmbActiveDirectorySettingsPtrOutput() GatewaySmbActiveDirectorySettingsPtrOutput
	ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(context.Context) GatewaySmbActiveDirectorySettingsPtrOutput
}

type gatewaySmbActiveDirectorySettingsPtrType GatewaySmbActiveDirectorySettingsArgs

func GatewaySmbActiveDirectorySettingsPtr(v *GatewaySmbActiveDirectorySettingsArgs) GatewaySmbActiveDirectorySettingsPtrInput {	return (*gatewaySmbActiveDirectorySettingsPtrType)(v)
}

func (*gatewaySmbActiveDirectorySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewaySmbActiveDirectorySettings)(nil)).Elem()
}

func (i *gatewaySmbActiveDirectorySettingsPtrType) ToGatewaySmbActiveDirectorySettingsPtrOutput() GatewaySmbActiveDirectorySettingsPtrOutput {
	return i.ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(context.Background())
}

func (i *gatewaySmbActiveDirectorySettingsPtrType) ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) GatewaySmbActiveDirectorySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewaySmbActiveDirectorySettingsPtrOutput)
}

type GatewaySmbActiveDirectorySettingsOutput struct { *pulumi.OutputState }

func (GatewaySmbActiveDirectorySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewaySmbActiveDirectorySettings)(nil)).Elem()
}

func (o GatewaySmbActiveDirectorySettingsOutput) ToGatewaySmbActiveDirectorySettingsOutput() GatewaySmbActiveDirectorySettingsOutput {
	return o
}

func (o GatewaySmbActiveDirectorySettingsOutput) ToGatewaySmbActiveDirectorySettingsOutputWithContext(ctx context.Context) GatewaySmbActiveDirectorySettingsOutput {
	return o
}

func (o GatewaySmbActiveDirectorySettingsOutput) ToGatewaySmbActiveDirectorySettingsPtrOutput() GatewaySmbActiveDirectorySettingsPtrOutput {
	return o.ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(context.Background())
}

func (o GatewaySmbActiveDirectorySettingsOutput) ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) GatewaySmbActiveDirectorySettingsPtrOutput {
	return o.ApplyT(func(v GatewaySmbActiveDirectorySettings) *GatewaySmbActiveDirectorySettings {
		return &v
	}).(GatewaySmbActiveDirectorySettingsPtrOutput)
}
// The name of the domain that you want the gateway to join.
func (o GatewaySmbActiveDirectorySettingsOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func (v GatewaySmbActiveDirectorySettings) string { return v.DomainName }).(pulumi.StringOutput)
}

// The password of the user who has permission to add the gateway to the Active Directory domain.
func (o GatewaySmbActiveDirectorySettingsOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func (v GatewaySmbActiveDirectorySettings) string { return v.Password }).(pulumi.StringOutput)
}

// The user name of user who has permission to add the gateway to the Active Directory domain.
func (o GatewaySmbActiveDirectorySettingsOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func (v GatewaySmbActiveDirectorySettings) string { return v.Username }).(pulumi.StringOutput)
}

type GatewaySmbActiveDirectorySettingsPtrOutput struct { *pulumi.OutputState}

func (GatewaySmbActiveDirectorySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewaySmbActiveDirectorySettings)(nil)).Elem()
}

func (o GatewaySmbActiveDirectorySettingsPtrOutput) ToGatewaySmbActiveDirectorySettingsPtrOutput() GatewaySmbActiveDirectorySettingsPtrOutput {
	return o
}

func (o GatewaySmbActiveDirectorySettingsPtrOutput) ToGatewaySmbActiveDirectorySettingsPtrOutputWithContext(ctx context.Context) GatewaySmbActiveDirectorySettingsPtrOutput {
	return o
}

func (o GatewaySmbActiveDirectorySettingsPtrOutput) Elem() GatewaySmbActiveDirectorySettingsOutput {
	return o.ApplyT(func (v *GatewaySmbActiveDirectorySettings) GatewaySmbActiveDirectorySettings { return *v }).(GatewaySmbActiveDirectorySettingsOutput)
}

// The name of the domain that you want the gateway to join.
func (o GatewaySmbActiveDirectorySettingsPtrOutput) DomainName() pulumi.StringOutput {
	return o.ApplyT(func (v GatewaySmbActiveDirectorySettings) string { return v.DomainName }).(pulumi.StringOutput)
}

// The password of the user who has permission to add the gateway to the Active Directory domain.
func (o GatewaySmbActiveDirectorySettingsPtrOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func (v GatewaySmbActiveDirectorySettings) string { return v.Password }).(pulumi.StringOutput)
}

// The user name of user who has permission to add the gateway to the Active Directory domain.
func (o GatewaySmbActiveDirectorySettingsPtrOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func (v GatewaySmbActiveDirectorySettings) string { return v.Username }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewaySmbActiveDirectorySettingsOutput{})
	pulumi.RegisterOutputType(GatewaySmbActiveDirectorySettingsPtrOutput{})
}
