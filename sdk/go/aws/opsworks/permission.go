// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an OpsWorks permission resource.
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
// 		_, err := opsworks.NewPermission(ctx, "myStackPermission", &opsworks.PermissionArgs{
// 			AllowSsh:  pulumi.Bool(true),
// 			AllowSudo: pulumi.Bool(true),
// 			Level:     pulumi.String("iam_only"),
// 			UserArn:   pulumi.Any(aws_iam_user.User.Arn),
// 			StackId:   pulumi.Any(aws_opsworks_stack.Stack.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Permission struct {
	pulumi.CustomResourceState

	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSsh pulumi.BoolOutput `pulumi:"allowSsh"`
	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo pulumi.BoolOutput `pulumi:"allowSudo"`
	// The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iamOnly`
	Level pulumi.StringOutput `pulumi:"level"`
	// The stack to set the permissions for
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// The user's IAM ARN to set permissions for
	UserArn pulumi.StringOutput `pulumi:"userArn"`
}

// NewPermission registers a new resource with the given unique name, arguments, and options.
func NewPermission(ctx *pulumi.Context,
	name string, args *PermissionArgs, opts ...pulumi.ResourceOption) (*Permission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.UserArn == nil {
		return nil, errors.New("invalid value for required argument 'UserArn'")
	}
	var resource Permission
	err := ctx.RegisterResource("aws:opsworks/permission:Permission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermission gets an existing Permission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionState, opts ...pulumi.ResourceOption) (*Permission, error) {
	var resource Permission
	err := ctx.ReadResource("aws:opsworks/permission:Permission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Permission resources.
type permissionState struct {
	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSsh *bool `pulumi:"allowSsh"`
	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo *bool `pulumi:"allowSudo"`
	// The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iamOnly`
	Level *string `pulumi:"level"`
	// The stack to set the permissions for
	StackId *string `pulumi:"stackId"`
	// The user's IAM ARN to set permissions for
	UserArn *string `pulumi:"userArn"`
}

type PermissionState struct {
	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSsh pulumi.BoolPtrInput
	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo pulumi.BoolPtrInput
	// The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iamOnly`
	Level pulumi.StringPtrInput
	// The stack to set the permissions for
	StackId pulumi.StringPtrInput
	// The user's IAM ARN to set permissions for
	UserArn pulumi.StringPtrInput
}

func (PermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionState)(nil)).Elem()
}

type permissionArgs struct {
	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSsh *bool `pulumi:"allowSsh"`
	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo *bool `pulumi:"allowSudo"`
	// The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iamOnly`
	Level *string `pulumi:"level"`
	// The stack to set the permissions for
	StackId *string `pulumi:"stackId"`
	// The user's IAM ARN to set permissions for
	UserArn string `pulumi:"userArn"`
}

// The set of arguments for constructing a Permission resource.
type PermissionArgs struct {
	// Whether the user is allowed to use SSH to communicate with the instance
	AllowSsh pulumi.BoolPtrInput
	// Whether the user is allowed to use sudo to elevate privileges
	AllowSudo pulumi.BoolPtrInput
	// The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iamOnly`
	Level pulumi.StringPtrInput
	// The stack to set the permissions for
	StackId pulumi.StringPtrInput
	// The user's IAM ARN to set permissions for
	UserArn pulumi.StringInput
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionArgs)(nil)).Elem()
}
