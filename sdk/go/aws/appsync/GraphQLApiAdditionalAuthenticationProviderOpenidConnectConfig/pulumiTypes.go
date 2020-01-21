// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig struct {
	// Number of milliseconds a token is valid after being authenticated.
	AuthTtl *int `pulumi:"authTtl"`
	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	ClientId *string `pulumi:"clientId"`
	// Number of milliseconds a token is valid after being issued to a user.
	IatTtl *int `pulumi:"iatTtl"`
	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	Issuer string `pulumi:"issuer"`
}

type GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigInput interface {
	pulumi.Input

	ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput
	ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutputWithContext(context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput
}

type GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs struct {
	// Number of milliseconds a token is valid after being authenticated.
	AuthTtl pulumi.IntPtrInput `pulumi:"authTtl"`
	// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// Number of milliseconds a token is valid after being issued to a user.
	IatTtl pulumi.IntPtrInput `pulumi:"iatTtl"`
	// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
	Issuer pulumi.StringInput `pulumi:"issuer"`
}

func (GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig)(nil)).Elem()
}

func (i GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput {
	return i.ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutputWithContext(context.Background())
}

func (i GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutputWithContext(ctx context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput)
}

func (i GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return i.ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(context.Background())
}

func (i GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(ctx context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput).ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(ctx)
}

type GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrInput interface {
	pulumi.Input

	ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput
	ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput
}

type graphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrType GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs

func GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtr(v *GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigArgs) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrInput {	return (*graphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrType)(v)
}

func (*graphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig)(nil)).Elem()
}

func (i *graphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrType) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return i.ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(context.Background())
}

func (i *graphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrType) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(ctx context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput)
}

type GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput struct { *pulumi.OutputState }

func (GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig)(nil)).Elem()
}

func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput {
	return o
}

func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutputWithContext(ctx context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput {
	return o
}

func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return o.ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(context.Background())
}

func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(ctx context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return o.ApplyT(func(v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) *GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig {
		return &v
	}).(GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput)
}
// Number of milliseconds a token is valid after being authenticated.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) AuthTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) *int { return v.AuthTtl }).(pulumi.IntPtrOutput)
}

// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// Number of milliseconds a token is valid after being issued to a user.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) IatTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) *int { return v.IatTtl }).(pulumi.IntPtrOutput)
}

// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) string { return v.Issuer }).(pulumi.StringOutput)
}

type GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput struct { *pulumi.OutputState}

func (GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig)(nil)).Elem()
}

func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return o
}

func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) ToGraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutputWithContext(ctx context.Context) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput {
	return o
}

func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) Elem() GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput {
	return o.ApplyT(func (v *GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig { return *v }).(GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput)
}

// Number of milliseconds a token is valid after being authenticated.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) AuthTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) *int { return v.AuthTtl }).(pulumi.IntPtrOutput)
}

// Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// Number of milliseconds a token is valid after being issued to a user.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) IatTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) *int { return v.IatTtl }).(pulumi.IntPtrOutput)
}

// Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
func (o GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func (v GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfig) string { return v.Issuer }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigOutput{})
	pulumi.RegisterOutputType(GraphQLApiAdditionalAuthenticationProviderOpenidConnectConfigPtrOutput{})
}
