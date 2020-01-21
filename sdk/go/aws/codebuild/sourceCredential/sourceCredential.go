// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package sourceCredential

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodeBuild Source Credentials Resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codebuild_source_credential.html.markdown.
type SourceCredential struct {
	pulumi.CustomResourceState

	// The ARN of Source Credential.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType pulumi.StringOutput `pulumi:"authType"`
	// The source provider used for this project.
	ServerType pulumi.StringOutput `pulumi:"serverType"`
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token pulumi.StringOutput `pulumi:"token"`
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewSourceCredential registers a new resource with the given unique name, arguments, and options.
func NewSourceCredential(ctx *pulumi.Context,
	name string, args *SourceCredentialArgs, opts ...pulumi.ResourceOption) (*SourceCredential, error) {
	if args == nil || args.AuthType == nil {
		return nil, errors.New("missing required argument 'AuthType'")
	}
	if args == nil || args.ServerType == nil {
		return nil, errors.New("missing required argument 'ServerType'")
	}
	if args == nil || args.Token == nil {
		return nil, errors.New("missing required argument 'Token'")
	}
	if args == nil {
		args = &SourceCredentialArgs{}
	}
	var resource SourceCredential
	err := ctx.RegisterResource("aws:codebuild/sourceCredential:SourceCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSourceCredential gets an existing SourceCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSourceCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SourceCredentialState, opts ...pulumi.ResourceOption) (*SourceCredential, error) {
	var resource SourceCredential
	err := ctx.ReadResource("aws:codebuild/sourceCredential:SourceCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SourceCredential resources.
type sourceCredentialState struct {
	// The ARN of Source Credential.
	Arn *string `pulumi:"arn"`
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType *string `pulumi:"authType"`
	// The source provider used for this project.
	ServerType *string `pulumi:"serverType"`
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token *string `pulumi:"token"`
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName *string `pulumi:"userName"`
}

type SourceCredentialState struct {
	// The ARN of Source Credential.
	Arn pulumi.StringPtrInput
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType pulumi.StringPtrInput
	// The source provider used for this project.
	ServerType pulumi.StringPtrInput
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token pulumi.StringPtrInput
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName pulumi.StringPtrInput
}

func (SourceCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCredentialState)(nil)).Elem()
}

type sourceCredentialArgs struct {
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType string `pulumi:"authType"`
	// The source provider used for this project.
	ServerType string `pulumi:"serverType"`
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token string `pulumi:"token"`
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a SourceCredential resource.
type SourceCredentialArgs struct {
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType pulumi.StringInput
	// The source provider used for this project.
	ServerType pulumi.StringInput
	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token pulumi.StringInput
	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName pulumi.StringPtrInput
}

func (SourceCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sourceCredentialArgs)(nil)).Elem()
}

