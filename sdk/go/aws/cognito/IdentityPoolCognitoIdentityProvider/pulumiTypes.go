// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package IdentityPoolCognitoIdentityProvider

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type IdentityPoolCognitoIdentityProvider struct {
	// The client ID for the Amazon Cognito Identity User Pool.
	ClientId *string `pulumi:"clientId"`
	// The provider name for an Amazon Cognito Identity User Pool.
	ProviderName *string `pulumi:"providerName"`
	// Whether server-side token validation is enabled for the identity provider’s token or not.
	ServerSideTokenCheck *bool `pulumi:"serverSideTokenCheck"`
}

type IdentityPoolCognitoIdentityProviderInput interface {
	pulumi.Input

	ToIdentityPoolCognitoIdentityProviderOutput() IdentityPoolCognitoIdentityProviderOutput
	ToIdentityPoolCognitoIdentityProviderOutputWithContext(context.Context) IdentityPoolCognitoIdentityProviderOutput
}

type IdentityPoolCognitoIdentityProviderArgs struct {
	// The client ID for the Amazon Cognito Identity User Pool.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// The provider name for an Amazon Cognito Identity User Pool.
	ProviderName pulumi.StringPtrInput `pulumi:"providerName"`
	// Whether server-side token validation is enabled for the identity provider’s token or not.
	ServerSideTokenCheck pulumi.BoolPtrInput `pulumi:"serverSideTokenCheck"`
}

func (IdentityPoolCognitoIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPoolCognitoIdentityProvider)(nil)).Elem()
}

func (i IdentityPoolCognitoIdentityProviderArgs) ToIdentityPoolCognitoIdentityProviderOutput() IdentityPoolCognitoIdentityProviderOutput {
	return i.ToIdentityPoolCognitoIdentityProviderOutputWithContext(context.Background())
}

func (i IdentityPoolCognitoIdentityProviderArgs) ToIdentityPoolCognitoIdentityProviderOutputWithContext(ctx context.Context) IdentityPoolCognitoIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPoolCognitoIdentityProviderOutput)
}

type IdentityPoolCognitoIdentityProviderArrayInput interface {
	pulumi.Input

	ToIdentityPoolCognitoIdentityProviderArrayOutput() IdentityPoolCognitoIdentityProviderArrayOutput
	ToIdentityPoolCognitoIdentityProviderArrayOutputWithContext(context.Context) IdentityPoolCognitoIdentityProviderArrayOutput
}

type IdentityPoolCognitoIdentityProviderArray []IdentityPoolCognitoIdentityProviderInput

func (IdentityPoolCognitoIdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityPoolCognitoIdentityProvider)(nil)).Elem()
}

func (i IdentityPoolCognitoIdentityProviderArray) ToIdentityPoolCognitoIdentityProviderArrayOutput() IdentityPoolCognitoIdentityProviderArrayOutput {
	return i.ToIdentityPoolCognitoIdentityProviderArrayOutputWithContext(context.Background())
}

func (i IdentityPoolCognitoIdentityProviderArray) ToIdentityPoolCognitoIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityPoolCognitoIdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPoolCognitoIdentityProviderArrayOutput)
}

type IdentityPoolCognitoIdentityProviderOutput struct { *pulumi.OutputState }

func (IdentityPoolCognitoIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPoolCognitoIdentityProvider)(nil)).Elem()
}

func (o IdentityPoolCognitoIdentityProviderOutput) ToIdentityPoolCognitoIdentityProviderOutput() IdentityPoolCognitoIdentityProviderOutput {
	return o
}

func (o IdentityPoolCognitoIdentityProviderOutput) ToIdentityPoolCognitoIdentityProviderOutputWithContext(ctx context.Context) IdentityPoolCognitoIdentityProviderOutput {
	return o
}

// The client ID for the Amazon Cognito Identity User Pool.
func (o IdentityPoolCognitoIdentityProviderOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v IdentityPoolCognitoIdentityProvider) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The provider name for an Amazon Cognito Identity User Pool.
func (o IdentityPoolCognitoIdentityProviderOutput) ProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func (v IdentityPoolCognitoIdentityProvider) *string { return v.ProviderName }).(pulumi.StringPtrOutput)
}

// Whether server-side token validation is enabled for the identity provider’s token or not.
func (o IdentityPoolCognitoIdentityProviderOutput) ServerSideTokenCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v IdentityPoolCognitoIdentityProvider) *bool { return v.ServerSideTokenCheck }).(pulumi.BoolPtrOutput)
}

type IdentityPoolCognitoIdentityProviderArrayOutput struct { *pulumi.OutputState}

func (IdentityPoolCognitoIdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentityPoolCognitoIdentityProvider)(nil)).Elem()
}

func (o IdentityPoolCognitoIdentityProviderArrayOutput) ToIdentityPoolCognitoIdentityProviderArrayOutput() IdentityPoolCognitoIdentityProviderArrayOutput {
	return o
}

func (o IdentityPoolCognitoIdentityProviderArrayOutput) ToIdentityPoolCognitoIdentityProviderArrayOutputWithContext(ctx context.Context) IdentityPoolCognitoIdentityProviderArrayOutput {
	return o
}

func (o IdentityPoolCognitoIdentityProviderArrayOutput) Index(i pulumi.IntInput) IdentityPoolCognitoIdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) IdentityPoolCognitoIdentityProvider {
		return vs[0].([]IdentityPoolCognitoIdentityProvider)[vs[1].(int)]
	}).(IdentityPoolCognitoIdentityProviderOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityPoolCognitoIdentityProviderOutput{})
	pulumi.RegisterOutputType(IdentityPoolCognitoIdentityProviderArrayOutput{})
}
