// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type UserLoginProfile struct {
	pulumi.CustomResourceState

	EncryptedPassword     pulumi.StringOutput  `pulumi:"encryptedPassword"`
	KeyFingerprint        pulumi.StringOutput  `pulumi:"keyFingerprint"`
	PasswordLength        pulumi.IntPtrOutput  `pulumi:"passwordLength"`
	PasswordResetRequired pulumi.BoolPtrOutput `pulumi:"passwordResetRequired"`
	PgpKey                pulumi.StringOutput  `pulumi:"pgpKey"`
	User                  pulumi.StringOutput  `pulumi:"user"`
}

// NewUserLoginProfile registers a new resource with the given unique name, arguments, and options.
func NewUserLoginProfile(ctx *pulumi.Context,
	name string, args *UserLoginProfileArgs, opts ...pulumi.ResourceOption) (*UserLoginProfile, error) {
	if args == nil || args.PgpKey == nil {
		return nil, errors.New("missing required argument 'PgpKey'")
	}
	if args == nil || args.User == nil {
		return nil, errors.New("missing required argument 'User'")
	}
	if args == nil {
		args = &UserLoginProfileArgs{}
	}
	var resource UserLoginProfile
	err := ctx.RegisterResource("aws:iam/userLoginProfile:UserLoginProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserLoginProfile gets an existing UserLoginProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserLoginProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserLoginProfileState, opts ...pulumi.ResourceOption) (*UserLoginProfile, error) {
	var resource UserLoginProfile
	err := ctx.ReadResource("aws:iam/userLoginProfile:UserLoginProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserLoginProfile resources.
type userLoginProfileState struct {
	EncryptedPassword     *string `pulumi:"encryptedPassword"`
	KeyFingerprint        *string `pulumi:"keyFingerprint"`
	PasswordLength        *int    `pulumi:"passwordLength"`
	PasswordResetRequired *bool   `pulumi:"passwordResetRequired"`
	PgpKey                *string `pulumi:"pgpKey"`
	User                  *string `pulumi:"user"`
}

type UserLoginProfileState struct {
	EncryptedPassword     pulumi.StringPtrInput
	KeyFingerprint        pulumi.StringPtrInput
	PasswordLength        pulumi.IntPtrInput
	PasswordResetRequired pulumi.BoolPtrInput
	PgpKey                pulumi.StringPtrInput
	User                  pulumi.StringPtrInput
}

func (UserLoginProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*userLoginProfileState)(nil)).Elem()
}

type userLoginProfileArgs struct {
	PasswordLength        *int   `pulumi:"passwordLength"`
	PasswordResetRequired *bool  `pulumi:"passwordResetRequired"`
	PgpKey                string `pulumi:"pgpKey"`
	User                  string `pulumi:"user"`
}

// The set of arguments for constructing a UserLoginProfile resource.
type UserLoginProfileArgs struct {
	PasswordLength        pulumi.IntPtrInput
	PasswordResetRequired pulumi.BoolPtrInput
	PgpKey                pulumi.StringInput
	User                  pulumi.StringInput
}

func (UserLoginProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userLoginProfileArgs)(nil)).Elem()
}
