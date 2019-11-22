// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM OpenID Connect provider.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_openid_connect_provider.html.markdown.
type OpenIdConnectProvider struct {
	s *pulumi.ResourceState
}

// NewOpenIdConnectProvider registers a new resource with the given unique name, arguments, and options.
func NewOpenIdConnectProvider(ctx *pulumi.Context,
	name string, args *OpenIdConnectProviderArgs, opts ...pulumi.ResourceOpt) (*OpenIdConnectProvider, error) {
	if args == nil || args.ClientIdLists == nil {
		return nil, errors.New("missing required argument 'ClientIdLists'")
	}
	if args == nil || args.ThumbprintLists == nil {
		return nil, errors.New("missing required argument 'ThumbprintLists'")
	}
	if args == nil || args.Url == nil {
		return nil, errors.New("missing required argument 'Url'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["clientIdLists"] = nil
		inputs["thumbprintLists"] = nil
		inputs["url"] = nil
	} else {
		inputs["clientIdLists"] = args.ClientIdLists
		inputs["thumbprintLists"] = args.ThumbprintLists
		inputs["url"] = args.Url
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OpenIdConnectProvider{s: s}, nil
}

// GetOpenIdConnectProvider gets an existing OpenIdConnectProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOpenIdConnectProvider(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OpenIdConnectProviderState, opts ...pulumi.ResourceOpt) (*OpenIdConnectProvider, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["clientIdLists"] = state.ClientIdLists
		inputs["thumbprintLists"] = state.ThumbprintLists
		inputs["url"] = state.Url
	}
	s, err := ctx.ReadResource("aws:iam/openIdConnectProvider:OpenIdConnectProvider", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OpenIdConnectProvider{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OpenIdConnectProvider) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OpenIdConnectProvider) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The ARN assigned by AWS for this provider.
func (r *OpenIdConnectProvider) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
func (r *OpenIdConnectProvider) ClientIdLists() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["clientIdLists"])
}

// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s). 
func (r *OpenIdConnectProvider) ThumbprintLists() pulumi.ArrayOutput {
	return (pulumi.ArrayOutput)(r.s.State["thumbprintLists"])
}

// The URL of the identity provider. Corresponds to the _iss_ claim.
func (r *OpenIdConnectProvider) Url() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["url"])
}

// Input properties used for looking up and filtering OpenIdConnectProvider resources.
type OpenIdConnectProviderState struct {
	// The ARN assigned by AWS for this provider.
	Arn interface{}
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists interface{}
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s). 
	ThumbprintLists interface{}
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url interface{}
}

// The set of arguments for constructing a OpenIdConnectProvider resource.
type OpenIdConnectProviderArgs struct {
	// A list of client IDs (also known as audiences). When a mobile or web app registers with an OpenID Connect provider, they establish a value that identifies the application. (This is the value that's sent as the clientId parameter on OAuth requests.)
	ClientIdLists interface{}
	// A list of server certificate thumbprints for the OpenID Connect (OIDC) identity provider's server certificate(s). 
	ThumbprintLists interface{}
	// The URL of the identity provider. Corresponds to the _iss_ claim.
	Url interface{}
}
