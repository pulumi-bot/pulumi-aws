// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Cognito Identity Pool.
//
// ## Example Usage
type IdentityPool struct {
	pulumi.CustomResourceState

	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities pulumi.BoolPtrOutput `pulumi:"allowUnauthenticatedIdentities"`
	// The ARN of the identity pool.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders IdentityPoolCognitoIdentityProviderArrayOutput `pulumi:"cognitoIdentityProviders"`
	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	DeveloperProviderName pulumi.StringPtrOutput `pulumi:"developerProviderName"`
	// The Cognito Identity Pool name.
	IdentityPoolName pulumi.StringOutput `pulumi:"identityPoolName"`
	// A list of OpendID Connect provider ARNs.
	OpenidConnectProviderArns pulumi.StringArrayOutput `pulumi:"openidConnectProviderArns"`
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	SamlProviderArns pulumi.StringArrayOutput `pulumi:"samlProviderArns"`
	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders pulumi.StringMapOutput `pulumi:"supportedLoginProviders"`
	// A map of tags to assign to the Identity Pool.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewIdentityPool registers a new resource with the given unique name, arguments, and options.
func NewIdentityPool(ctx *pulumi.Context,
	name string, args *IdentityPoolArgs, opts ...pulumi.ResourceOption) (*IdentityPool, error) {
	if args == nil || args.IdentityPoolName == nil {
		return nil, errors.New("missing required argument 'IdentityPoolName'")
	}
	if args == nil {
		args = &IdentityPoolArgs{}
	}
	var resource IdentityPool
	err := ctx.RegisterResource("aws:cognito/identityPool:IdentityPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityPool gets an existing IdentityPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityPoolState, opts ...pulumi.ResourceOption) (*IdentityPool, error) {
	var resource IdentityPool
	err := ctx.ReadResource("aws:cognito/identityPool:IdentityPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityPool resources.
type identityPoolState struct {
	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities *bool `pulumi:"allowUnauthenticatedIdentities"`
	// The ARN of the identity pool.
	Arn *string `pulumi:"arn"`
	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders []IdentityPoolCognitoIdentityProvider `pulumi:"cognitoIdentityProviders"`
	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	DeveloperProviderName *string `pulumi:"developerProviderName"`
	// The Cognito Identity Pool name.
	IdentityPoolName *string `pulumi:"identityPoolName"`
	// A list of OpendID Connect provider ARNs.
	OpenidConnectProviderArns []string `pulumi:"openidConnectProviderArns"`
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	SamlProviderArns []string `pulumi:"samlProviderArns"`
	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string `pulumi:"supportedLoginProviders"`
	// A map of tags to assign to the Identity Pool.
	Tags map[string]string `pulumi:"tags"`
}

type IdentityPoolState struct {
	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities pulumi.BoolPtrInput
	// The ARN of the identity pool.
	Arn pulumi.StringPtrInput
	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders IdentityPoolCognitoIdentityProviderArrayInput
	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	DeveloperProviderName pulumi.StringPtrInput
	// The Cognito Identity Pool name.
	IdentityPoolName pulumi.StringPtrInput
	// A list of OpendID Connect provider ARNs.
	OpenidConnectProviderArns pulumi.StringArrayInput
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	SamlProviderArns pulumi.StringArrayInput
	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders pulumi.StringMapInput
	// A map of tags to assign to the Identity Pool.
	Tags pulumi.StringMapInput
}

func (IdentityPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolState)(nil)).Elem()
}

type identityPoolArgs struct {
	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities *bool `pulumi:"allowUnauthenticatedIdentities"`
	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders []IdentityPoolCognitoIdentityProvider `pulumi:"cognitoIdentityProviders"`
	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	DeveloperProviderName *string `pulumi:"developerProviderName"`
	// The Cognito Identity Pool name.
	IdentityPoolName string `pulumi:"identityPoolName"`
	// A list of OpendID Connect provider ARNs.
	OpenidConnectProviderArns []string `pulumi:"openidConnectProviderArns"`
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	SamlProviderArns []string `pulumi:"samlProviderArns"`
	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders map[string]string `pulumi:"supportedLoginProviders"`
	// A map of tags to assign to the Identity Pool.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IdentityPool resource.
type IdentityPoolArgs struct {
	// Whether the identity pool supports unauthenticated logins or not.
	AllowUnauthenticatedIdentities pulumi.BoolPtrInput
	// An array of Amazon Cognito Identity user pools and their client IDs.
	CognitoIdentityProviders IdentityPoolCognitoIdentityProviderArrayInput
	// The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
	// backend and the Cognito service to communicate about the developer provider.
	DeveloperProviderName pulumi.StringPtrInput
	// The Cognito Identity Pool name.
	IdentityPoolName pulumi.StringInput
	// A list of OpendID Connect provider ARNs.
	OpenidConnectProviderArns pulumi.StringArrayInput
	// An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
	SamlProviderArns pulumi.StringArrayInput
	// Key-Value pairs mapping provider names to provider app IDs.
	SupportedLoginProviders pulumi.StringMapInput
	// A map of tags to assign to the Identity Pool.
	Tags pulumi.StringMapInput
}

func (IdentityPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolArgs)(nil)).Elem()
}
