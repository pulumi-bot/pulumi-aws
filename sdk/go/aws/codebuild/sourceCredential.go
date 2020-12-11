// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CodeBuild Source Credentials Resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codebuild"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := codebuild.NewSourceCredential(ctx, "example", &codebuild.SourceCredentialArgs{
// 			AuthType:   pulumi.String("PERSONAL_ACCESS_TOKEN"),
// 			ServerType: pulumi.String("GITHUB"),
// 			Token:      pulumi.String("example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Bitbucket Server Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/codebuild"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := codebuild.NewSourceCredential(ctx, "example", &codebuild.SourceCredentialArgs{
// 			AuthType:   pulumi.String("BASIC_AUTH"),
// 			ServerType: pulumi.String("BITBUCKET"),
// 			Token:      pulumi.String("example"),
// 			UserName:   pulumi.String("test-user"),
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
// CodeBuild Source Credential can be imported using the CodeBuild Source Credential arn, e.g.
//
// ```sh
//  $ pulumi import aws:codebuild/sourceCredential:SourceCredential example arn:aws:codebuild:us-west-2:123456789:token:github
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthType == nil {
		return nil, errors.New("invalid value for required argument 'AuthType'")
	}
	if args.ServerType == nil {
		return nil, errors.New("invalid value for required argument 'ServerType'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
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

type SourceCredentialInput interface {
	pulumi.Input

	ToSourceCredentialOutput() SourceCredentialOutput
	ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput
}

type SourceCredentialPtrInput interface {
	pulumi.Input

	ToSourceCredentialPtrOutput() SourceCredentialPtrOutput
	ToSourceCredentialPtrOutputWithContext(ctx context.Context) SourceCredentialPtrOutput
}

func (SourceCredential) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCredential)(nil)).Elem()
}

func (i SourceCredential) ToSourceCredentialOutput() SourceCredentialOutput {
	return i.ToSourceCredentialOutputWithContext(context.Background())
}

func (i SourceCredential) ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCredentialOutput)
}

func (i SourceCredential) ToSourceCredentialPtrOutput() SourceCredentialPtrOutput {
	return i.ToSourceCredentialPtrOutputWithContext(context.Background())
}

func (i SourceCredential) ToSourceCredentialPtrOutputWithContext(ctx context.Context) SourceCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCredentialPtrOutput)
}

type SourceCredentialOutput struct {
	*pulumi.OutputState
}

func (SourceCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCredentialOutput)(nil)).Elem()
}

func (o SourceCredentialOutput) ToSourceCredentialOutput() SourceCredentialOutput {
	return o
}

func (o SourceCredentialOutput) ToSourceCredentialOutputWithContext(ctx context.Context) SourceCredentialOutput {
	return o
}

type SourceCredentialPtrOutput struct {
	*pulumi.OutputState
}

func (SourceCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceCredential)(nil)).Elem()
}

func (o SourceCredentialPtrOutput) ToSourceCredentialPtrOutput() SourceCredentialPtrOutput {
	return o
}

func (o SourceCredentialPtrOutput) ToSourceCredentialPtrOutputWithContext(ctx context.Context) SourceCredentialPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SourceCredentialOutput{})
	pulumi.RegisterOutputType(SourceCredentialPtrOutput{})
}
