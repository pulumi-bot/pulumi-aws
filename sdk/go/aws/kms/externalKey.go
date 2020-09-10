// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ExternalKey struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput    `pulumi:"arn"`
	DeletionWindowInDays pulumi.IntPtrOutput    `pulumi:"deletionWindowInDays"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	Enabled              pulumi.BoolOutput      `pulumi:"enabled"`
	ExpirationModel      pulumi.StringOutput    `pulumi:"expirationModel"`
	KeyMaterialBase64    pulumi.StringPtrOutput `pulumi:"keyMaterialBase64"`
	KeyState             pulumi.StringOutput    `pulumi:"keyState"`
	KeyUsage             pulumi.StringOutput    `pulumi:"keyUsage"`
	Policy               pulumi.StringOutput    `pulumi:"policy"`
	Tags                 pulumi.StringMapOutput `pulumi:"tags"`
	ValidTo              pulumi.StringPtrOutput `pulumi:"validTo"`
}

// NewExternalKey registers a new resource with the given unique name, arguments, and options.
func NewExternalKey(ctx *pulumi.Context,
	name string, args *ExternalKeyArgs, opts ...pulumi.ResourceOption) (*ExternalKey, error) {
	if args == nil {
		args = &ExternalKeyArgs{}
	}
	var resource ExternalKey
	err := ctx.RegisterResource("aws:kms/externalKey:ExternalKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalKey gets an existing ExternalKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalKeyState, opts ...pulumi.ResourceOption) (*ExternalKey, error) {
	var resource ExternalKey
	err := ctx.ReadResource("aws:kms/externalKey:ExternalKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalKey resources.
type externalKeyState struct {
	Arn                  *string           `pulumi:"arn"`
	DeletionWindowInDays *int              `pulumi:"deletionWindowInDays"`
	Description          *string           `pulumi:"description"`
	Enabled              *bool             `pulumi:"enabled"`
	ExpirationModel      *string           `pulumi:"expirationModel"`
	KeyMaterialBase64    *string           `pulumi:"keyMaterialBase64"`
	KeyState             *string           `pulumi:"keyState"`
	KeyUsage             *string           `pulumi:"keyUsage"`
	Policy               *string           `pulumi:"policy"`
	Tags                 map[string]string `pulumi:"tags"`
	ValidTo              *string           `pulumi:"validTo"`
}

type ExternalKeyState struct {
	Arn                  pulumi.StringPtrInput
	DeletionWindowInDays pulumi.IntPtrInput
	Description          pulumi.StringPtrInput
	Enabled              pulumi.BoolPtrInput
	ExpirationModel      pulumi.StringPtrInput
	KeyMaterialBase64    pulumi.StringPtrInput
	KeyState             pulumi.StringPtrInput
	KeyUsage             pulumi.StringPtrInput
	Policy               pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	ValidTo              pulumi.StringPtrInput
}

func (ExternalKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalKeyState)(nil)).Elem()
}

type externalKeyArgs struct {
	DeletionWindowInDays *int              `pulumi:"deletionWindowInDays"`
	Description          *string           `pulumi:"description"`
	Enabled              *bool             `pulumi:"enabled"`
	KeyMaterialBase64    *string           `pulumi:"keyMaterialBase64"`
	Policy               *string           `pulumi:"policy"`
	Tags                 map[string]string `pulumi:"tags"`
	ValidTo              *string           `pulumi:"validTo"`
}

// The set of arguments for constructing a ExternalKey resource.
type ExternalKeyArgs struct {
	DeletionWindowInDays pulumi.IntPtrInput
	Description          pulumi.StringPtrInput
	Enabled              pulumi.BoolPtrInput
	KeyMaterialBase64    pulumi.StringPtrInput
	Policy               pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	ValidTo              pulumi.StringPtrInput
}

func (ExternalKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalKeyArgs)(nil)).Elem()
}
