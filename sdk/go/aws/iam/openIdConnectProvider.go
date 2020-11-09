// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an IAM OpenID Connect provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.NewOpenIdConnectProvider(ctx, "_default", &iam.OpenIdConnectProviderArgs{
// 			ClientIdLists: pulumi.StringArray{
// 				pulumi.String("266362248691-342342xasdasdasda-apps.googleusercontent.com"),
// 			},
// 			ThumbprintLists: []interface{}{},
// 			Url:             pulumi.String("https://accounts.google.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// IAM OpenID Connect Providers can be imported using the `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:iam/openIdConnectProvider:OpenIdConnectProvider default arn:aws:iam::123456789012:oidc-provider/accounts.google.com
// ```
type OpenIdConnectProvider struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS for this provider.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists pulumi.StringArrayOutput `pulumi:"clientIdLists"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists pulumi.StringArrayOutput `pulumi:"thumbprintLists"`
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewOpenIdConnectProvider registers a new resource with the given unique name, arguments, and options.
func NewOpenIdConnectProvider(ctx *pulumi.Context,
	name string, args *OpenIdConnectProviderArgs, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	if args == nil || args.ClientIdLists == nil {
		return nil, errors.New("missing required argument 'ClientIdLists'")
	}
	if args == nil || args.ThumbprintLists == nil {
		return nil, errors.New("missing required argument 'ThumbprintLists'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	if args == nil {
		args = &OpenIdConnectProviderArgs{}
	}
	var resource OpenIdConnectProvider
	err := ctx.RegisterResource("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOpenIdConnectProvider gets an existing OpenIdConnectProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenIdConnectProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OpenIdConnectProviderState, opts ...pulumi.ResourceOption) (*OpenIdConnectProvider, error) {
	var resource OpenIdConnectProvider
	err := ctx.ReadResource("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OpenIdConnectProvider resources.
type openIdConnectProviderState struct {
	// The ARN assigned by AWS for this provider.
	Arn *string `pulumi:"arn"`
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists []string `pulumi:"clientIdLists"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists []string `pulumi:"thumbprintLists"`
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url *string `pulumi:"url"`
}

type OpenIdConnectProviderState struct {
	// The ARN assigned by AWS for this provider.
	Arn pulumi.StringPtrInput
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists pulumi.StringArrayInput
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists pulumi.StringArrayInput
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url pulumi.StringPtrInput
}

func (OpenIdConnectProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderState)(nil)).Elem()
}

type openIdConnectProviderArgs struct {
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists []string `pulumi:"clientIdLists"`
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists []string `pulumi:"thumbprintLists"`
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a OpenIdConnectProvider resource.
type OpenIdConnectProviderArgs struct {
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists pulumi.StringArrayInput
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s).
	ThumbprintLists pulumi.StringArrayInput
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url pulumi.StringInput
}

func (OpenIdConnectProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*openIdConnectProviderArgs)(nil)).Elem()
}
