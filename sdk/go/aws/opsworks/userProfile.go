// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an OpsWorks User Profile resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewUserProfile(ctx, "myProfile", &opsworks.UserProfileArgs{
// 			UserArn:     pulumi.Any(aws_iam_user.User.Arn),
// 			SshUsername: pulumi.String("my_user"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type UserProfile struct {
	pulumi.CustomResourceState

	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement pulumi.BoolPtrOutput `pulumi:"allowSelfManagement"`
	// The users public key
	SshPublicKey pulumi.StringPtrOutput `pulumi:"sshPublicKey"`
	// The ssh username, with witch this user wants to log in
	SshUsername pulumi.StringOutput `pulumi:"sshUsername"`
	// The user's IAM ARN
	UserArn pulumi.StringOutput `pulumi:"userArn"`
}

// NewUserProfile registers a new resource with the given unique name, arguments, and options.
func NewUserProfile(ctx *pulumi.Context,
	name string, args *UserProfileArgs, opts ...pulumi.ResourceOption) (*UserProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SshUsername == nil {
		return nil, errors.New("invalid value for required argument 'SshUsername'")
	}
	if args.UserArn == nil {
		return nil, errors.New("invalid value for required argument 'UserArn'")
	}
	var resource UserProfile
	err := ctx.RegisterResource("aws:opsworks/userProfile:UserProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserProfile gets an existing UserProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserProfileState, opts ...pulumi.ResourceOption) (*UserProfile, error) {
	var resource UserProfile
	err := ctx.ReadResource("aws:opsworks/userProfile:UserProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserProfile resources.
type userProfileState struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement *bool `pulumi:"allowSelfManagement"`
	// The users public key
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// The ssh username, with witch this user wants to log in
	SshUsername *string `pulumi:"sshUsername"`
	// The user's IAM ARN
	UserArn *string `pulumi:"userArn"`
}

type UserProfileState struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement pulumi.BoolPtrInput
	// The users public key
	SshPublicKey pulumi.StringPtrInput
	// The ssh username, with witch this user wants to log in
	SshUsername pulumi.StringPtrInput
	// The user's IAM ARN
	UserArn pulumi.StringPtrInput
}

func (UserProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*userProfileState)(nil)).Elem()
}

type userProfileArgs struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement *bool `pulumi:"allowSelfManagement"`
	// The users public key
	SshPublicKey *string `pulumi:"sshPublicKey"`
	// The ssh username, with witch this user wants to log in
	SshUsername string `pulumi:"sshUsername"`
	// The user's IAM ARN
	UserArn string `pulumi:"userArn"`
}

// The set of arguments for constructing a UserProfile resource.
type UserProfileArgs struct {
	// Whether users can specify their own SSH public key through the My Settings page
	AllowSelfManagement pulumi.BoolPtrInput
	// The users public key
	SshPublicKey pulumi.StringPtrInput
	// The ssh username, with witch this user wants to log in
	SshUsername pulumi.StringInput
	// The user's IAM ARN
	UserArn pulumi.StringInput
}

func (UserProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userProfileArgs)(nil)).Elem()
}

type UserProfileInput interface {
	pulumi.Input

	ToUserProfileOutput() UserProfileOutput
	ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput
}

type UserProfilePtrInput interface {
	pulumi.Input

	ToUserProfilePtrOutput() UserProfilePtrOutput
	ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput
}

func (UserProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProfile)(nil)).Elem()
}

func (i UserProfile) ToUserProfileOutput() UserProfileOutput {
	return i.ToUserProfileOutputWithContext(context.Background())
}

func (i UserProfile) ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfileOutput)
}

func (i UserProfile) ToUserProfilePtrOutput() UserProfilePtrOutput {
	return i.ToUserProfilePtrOutputWithContext(context.Background())
}

func (i UserProfile) ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserProfilePtrOutput)
}

type UserProfileOutput struct {
	*pulumi.OutputState
}

func (UserProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserProfileOutput)(nil)).Elem()
}

func (o UserProfileOutput) ToUserProfileOutput() UserProfileOutput {
	return o
}

func (o UserProfileOutput) ToUserProfileOutputWithContext(ctx context.Context) UserProfileOutput {
	return o
}

type UserProfilePtrOutput struct {
	*pulumi.OutputState
}

func (UserProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserProfile)(nil)).Elem()
}

func (o UserProfilePtrOutput) ToUserProfilePtrOutput() UserProfilePtrOutput {
	return o
}

func (o UserProfilePtrOutput) ToUserProfilePtrOutputWithContext(ctx context.Context) UserProfilePtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserProfileOutput{})
	pulumi.RegisterOutputType(UserProfilePtrOutput{})
}
