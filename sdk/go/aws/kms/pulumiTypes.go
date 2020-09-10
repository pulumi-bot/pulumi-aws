// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GrantConstraint struct {
	EncryptionContextEquals map[string]string `pulumi:"encryptionContextEquals"`
	EncryptionContextSubset map[string]string `pulumi:"encryptionContextSubset"`
}

// GrantConstraintInput is an input type that accepts GrantConstraintArgs and GrantConstraintOutput values.
// You can construct a concrete instance of `GrantConstraintInput` via:
//
//          GrantConstraintArgs{...}
type GrantConstraintInput interface {
	pulumi.Input

	ToGrantConstraintOutput() GrantConstraintOutput
	ToGrantConstraintOutputWithContext(context.Context) GrantConstraintOutput
}

type GrantConstraintArgs struct {
	EncryptionContextEquals pulumi.StringMapInput `pulumi:"encryptionContextEquals"`
	EncryptionContextSubset pulumi.StringMapInput `pulumi:"encryptionContextSubset"`
}

func (GrantConstraintArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GrantConstraint)(nil)).Elem()
}

func (i GrantConstraintArgs) ToGrantConstraintOutput() GrantConstraintOutput {
	return i.ToGrantConstraintOutputWithContext(context.Background())
}

func (i GrantConstraintArgs) ToGrantConstraintOutputWithContext(ctx context.Context) GrantConstraintOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantConstraintOutput)
}

// GrantConstraintArrayInput is an input type that accepts GrantConstraintArray and GrantConstraintArrayOutput values.
// You can construct a concrete instance of `GrantConstraintArrayInput` via:
//
//          GrantConstraintArray{ GrantConstraintArgs{...} }
type GrantConstraintArrayInput interface {
	pulumi.Input

	ToGrantConstraintArrayOutput() GrantConstraintArrayOutput
	ToGrantConstraintArrayOutputWithContext(context.Context) GrantConstraintArrayOutput
}

type GrantConstraintArray []GrantConstraintInput

func (GrantConstraintArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GrantConstraint)(nil)).Elem()
}

func (i GrantConstraintArray) ToGrantConstraintArrayOutput() GrantConstraintArrayOutput {
	return i.ToGrantConstraintArrayOutputWithContext(context.Background())
}

func (i GrantConstraintArray) ToGrantConstraintArrayOutputWithContext(ctx context.Context) GrantConstraintArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantConstraintArrayOutput)
}

type GrantConstraintOutput struct{ *pulumi.OutputState }

func (GrantConstraintOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GrantConstraint)(nil)).Elem()
}

func (o GrantConstraintOutput) ToGrantConstraintOutput() GrantConstraintOutput {
	return o
}

func (o GrantConstraintOutput) ToGrantConstraintOutputWithContext(ctx context.Context) GrantConstraintOutput {
	return o
}

func (o GrantConstraintOutput) EncryptionContextEquals() pulumi.StringMapOutput {
	return o.ApplyT(func(v GrantConstraint) map[string]string { return v.EncryptionContextEquals }).(pulumi.StringMapOutput)
}

func (o GrantConstraintOutput) EncryptionContextSubset() pulumi.StringMapOutput {
	return o.ApplyT(func(v GrantConstraint) map[string]string { return v.EncryptionContextSubset }).(pulumi.StringMapOutput)
}

type GrantConstraintArrayOutput struct{ *pulumi.OutputState }

func (GrantConstraintArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GrantConstraint)(nil)).Elem()
}

func (o GrantConstraintArrayOutput) ToGrantConstraintArrayOutput() GrantConstraintArrayOutput {
	return o
}

func (o GrantConstraintArrayOutput) ToGrantConstraintArrayOutputWithContext(ctx context.Context) GrantConstraintArrayOutput {
	return o
}

func (o GrantConstraintArrayOutput) Index(i pulumi.IntInput) GrantConstraintOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GrantConstraint {
		return vs[0].([]GrantConstraint)[vs[1].(int)]
	}).(GrantConstraintOutput)
}

type GetSecretSecret struct {
	Context     map[string]string `pulumi:"context"`
	GrantTokens []string          `pulumi:"grantTokens"`
	Name        string            `pulumi:"name"`
	Payload     string            `pulumi:"payload"`
}

// GetSecretSecretInput is an input type that accepts GetSecretSecretArgs and GetSecretSecretOutput values.
// You can construct a concrete instance of `GetSecretSecretInput` via:
//
//          GetSecretSecretArgs{...}
type GetSecretSecretInput interface {
	pulumi.Input

	ToGetSecretSecretOutput() GetSecretSecretOutput
	ToGetSecretSecretOutputWithContext(context.Context) GetSecretSecretOutput
}

type GetSecretSecretArgs struct {
	Context     pulumi.StringMapInput   `pulumi:"context"`
	GrantTokens pulumi.StringArrayInput `pulumi:"grantTokens"`
	Name        pulumi.StringInput      `pulumi:"name"`
	Payload     pulumi.StringInput      `pulumi:"payload"`
}

func (GetSecretSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretSecret)(nil)).Elem()
}

func (i GetSecretSecretArgs) ToGetSecretSecretOutput() GetSecretSecretOutput {
	return i.ToGetSecretSecretOutputWithContext(context.Background())
}

func (i GetSecretSecretArgs) ToGetSecretSecretOutputWithContext(ctx context.Context) GetSecretSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretSecretOutput)
}

// GetSecretSecretArrayInput is an input type that accepts GetSecretSecretArray and GetSecretSecretArrayOutput values.
// You can construct a concrete instance of `GetSecretSecretArrayInput` via:
//
//          GetSecretSecretArray{ GetSecretSecretArgs{...} }
type GetSecretSecretArrayInput interface {
	pulumi.Input

	ToGetSecretSecretArrayOutput() GetSecretSecretArrayOutput
	ToGetSecretSecretArrayOutputWithContext(context.Context) GetSecretSecretArrayOutput
}

type GetSecretSecretArray []GetSecretSecretInput

func (GetSecretSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretSecret)(nil)).Elem()
}

func (i GetSecretSecretArray) ToGetSecretSecretArrayOutput() GetSecretSecretArrayOutput {
	return i.ToGetSecretSecretArrayOutputWithContext(context.Background())
}

func (i GetSecretSecretArray) ToGetSecretSecretArrayOutputWithContext(ctx context.Context) GetSecretSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretSecretArrayOutput)
}

type GetSecretSecretOutput struct{ *pulumi.OutputState }

func (GetSecretSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretSecret)(nil)).Elem()
}

func (o GetSecretSecretOutput) ToGetSecretSecretOutput() GetSecretSecretOutput {
	return o
}

func (o GetSecretSecretOutput) ToGetSecretSecretOutputWithContext(ctx context.Context) GetSecretSecretOutput {
	return o
}

func (o GetSecretSecretOutput) Context() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSecretSecret) map[string]string { return v.Context }).(pulumi.StringMapOutput)
}

func (o GetSecretSecretOutput) GrantTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecretSecret) []string { return v.GrantTokens }).(pulumi.StringArrayOutput)
}

func (o GetSecretSecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretSecret) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSecretSecretOutput) Payload() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretSecret) string { return v.Payload }).(pulumi.StringOutput)
}

type GetSecretSecretArrayOutput struct{ *pulumi.OutputState }

func (GetSecretSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretSecret)(nil)).Elem()
}

func (o GetSecretSecretArrayOutput) ToGetSecretSecretArrayOutput() GetSecretSecretArrayOutput {
	return o
}

func (o GetSecretSecretArrayOutput) ToGetSecretSecretArrayOutputWithContext(ctx context.Context) GetSecretSecretArrayOutput {
	return o
}

func (o GetSecretSecretArrayOutput) Index(i pulumi.IntInput) GetSecretSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecretSecret {
		return vs[0].([]GetSecretSecret)[vs[1].(int)]
	}).(GetSecretSecretOutput)
}

type GetSecretsSecret struct {
	Context     map[string]string `pulumi:"context"`
	GrantTokens []string          `pulumi:"grantTokens"`
	Name        string            `pulumi:"name"`
	Payload     string            `pulumi:"payload"`
}

// GetSecretsSecretInput is an input type that accepts GetSecretsSecretArgs and GetSecretsSecretOutput values.
// You can construct a concrete instance of `GetSecretsSecretInput` via:
//
//          GetSecretsSecretArgs{...}
type GetSecretsSecretInput interface {
	pulumi.Input

	ToGetSecretsSecretOutput() GetSecretsSecretOutput
	ToGetSecretsSecretOutputWithContext(context.Context) GetSecretsSecretOutput
}

type GetSecretsSecretArgs struct {
	Context     pulumi.StringMapInput   `pulumi:"context"`
	GrantTokens pulumi.StringArrayInput `pulumi:"grantTokens"`
	Name        pulumi.StringInput      `pulumi:"name"`
	Payload     pulumi.StringInput      `pulumi:"payload"`
}

func (GetSecretsSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsSecret)(nil)).Elem()
}

func (i GetSecretsSecretArgs) ToGetSecretsSecretOutput() GetSecretsSecretOutput {
	return i.ToGetSecretsSecretOutputWithContext(context.Background())
}

func (i GetSecretsSecretArgs) ToGetSecretsSecretOutputWithContext(ctx context.Context) GetSecretsSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretsSecretOutput)
}

// GetSecretsSecretArrayInput is an input type that accepts GetSecretsSecretArray and GetSecretsSecretArrayOutput values.
// You can construct a concrete instance of `GetSecretsSecretArrayInput` via:
//
//          GetSecretsSecretArray{ GetSecretsSecretArgs{...} }
type GetSecretsSecretArrayInput interface {
	pulumi.Input

	ToGetSecretsSecretArrayOutput() GetSecretsSecretArrayOutput
	ToGetSecretsSecretArrayOutputWithContext(context.Context) GetSecretsSecretArrayOutput
}

type GetSecretsSecretArray []GetSecretsSecretInput

func (GetSecretsSecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretsSecret)(nil)).Elem()
}

func (i GetSecretsSecretArray) ToGetSecretsSecretArrayOutput() GetSecretsSecretArrayOutput {
	return i.ToGetSecretsSecretArrayOutputWithContext(context.Background())
}

func (i GetSecretsSecretArray) ToGetSecretsSecretArrayOutputWithContext(ctx context.Context) GetSecretsSecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSecretsSecretArrayOutput)
}

type GetSecretsSecretOutput struct{ *pulumi.OutputState }

func (GetSecretsSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretsSecret)(nil)).Elem()
}

func (o GetSecretsSecretOutput) ToGetSecretsSecretOutput() GetSecretsSecretOutput {
	return o
}

func (o GetSecretsSecretOutput) ToGetSecretsSecretOutputWithContext(ctx context.Context) GetSecretsSecretOutput {
	return o
}

func (o GetSecretsSecretOutput) Context() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSecretsSecret) map[string]string { return v.Context }).(pulumi.StringMapOutput)
}

func (o GetSecretsSecretOutput) GrantTokens() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSecretsSecret) []string { return v.GrantTokens }).(pulumi.StringArrayOutput)
}

func (o GetSecretsSecretOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecret) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetSecretsSecretOutput) Payload() pulumi.StringOutput {
	return o.ApplyT(func(v GetSecretsSecret) string { return v.Payload }).(pulumi.StringOutput)
}

type GetSecretsSecretArrayOutput struct{ *pulumi.OutputState }

func (GetSecretsSecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetSecretsSecret)(nil)).Elem()
}

func (o GetSecretsSecretArrayOutput) ToGetSecretsSecretArrayOutput() GetSecretsSecretArrayOutput {
	return o
}

func (o GetSecretsSecretArrayOutput) ToGetSecretsSecretArrayOutputWithContext(ctx context.Context) GetSecretsSecretArrayOutput {
	return o
}

func (o GetSecretsSecretArrayOutput) Index(i pulumi.IntInput) GetSecretsSecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetSecretsSecret {
		return vs[0].([]GetSecretsSecret)[vs[1].(int)]
	}).(GetSecretsSecretOutput)
}

func init() {
	pulumi.RegisterOutputType(GrantConstraintOutput{})
	pulumi.RegisterOutputType(GrantConstraintArrayOutput{})
	pulumi.RegisterOutputType(GetSecretSecretOutput{})
	pulumi.RegisterOutputType(GetSecretSecretArrayOutput{})
	pulumi.RegisterOutputType(GetSecretsSecretOutput{})
	pulumi.RegisterOutputType(GetSecretsSecretArrayOutput{})
}
