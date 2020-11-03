// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource for managing QuickSight User
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/quicksight"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := quicksight.NewUser(ctx, "example", &quicksight.UserArgs{
// 			Email:        pulumi.String("author@example.com"),
// 			IdentityType: pulumi.String("IAM"),
// 			UserName:     pulumi.String("an-author"),
// 			UserRole:     pulumi.String("AUTHOR"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type User struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the user
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringOutput `pulumi:"awsAccountId"`
	// The email address of the user that you want to register.
	Email pulumi.StringOutput `pulumi:"email"`
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn pulumi.StringPtrOutput `pulumi:"iamArn"`
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType pulumi.StringOutput `pulumi:"identityType"`
	// The namespace. Currently, you should set this to `default`.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName pulumi.StringPtrOutput `pulumi:"sessionName"`
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole pulumi.StringOutput `pulumi:"userRole"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.IdentityType == nil {
		return nil, errors.New("invalid value for required argument 'IdentityType'")
	}
	if args.UserRole == nil {
		return nil, errors.New("invalid value for required argument 'UserRole'")
	}
	var resource User
	err := ctx.RegisterResource("aws:quicksight/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws:quicksight/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Amazon Resource Name (ARN) of the user
	Arn *string `pulumi:"arn"`
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The email address of the user that you want to register.
	Email *string `pulumi:"email"`
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn *string `pulumi:"iamArn"`
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType *string `pulumi:"identityType"`
	// The namespace. Currently, you should set this to `default`.
	Namespace *string `pulumi:"namespace"`
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName *string `pulumi:"sessionName"`
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName *string `pulumi:"userName"`
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole *string `pulumi:"userRole"`
}

type UserState struct {
	// Amazon Resource Name (ARN) of the user
	Arn pulumi.StringPtrInput
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringPtrInput
	// The email address of the user that you want to register.
	Email pulumi.StringPtrInput
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn pulumi.StringPtrInput
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType pulumi.StringPtrInput
	// The namespace. Currently, you should set this to `default`.
	Namespace pulumi.StringPtrInput
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName pulumi.StringPtrInput
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName pulumi.StringPtrInput
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId *string `pulumi:"awsAccountId"`
	// The email address of the user that you want to register.
	Email string `pulumi:"email"`
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn *string `pulumi:"iamArn"`
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType string `pulumi:"identityType"`
	// The namespace. Currently, you should set this to `default`.
	Namespace *string `pulumi:"namespace"`
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName *string `pulumi:"sessionName"`
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName *string `pulumi:"userName"`
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole string `pulumi:"userRole"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
	AwsAccountId pulumi.StringPtrInput
	// The email address of the user that you want to register.
	Email pulumi.StringInput
	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn pulumi.StringPtrInput
	// Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
	IdentityType pulumi.StringInput
	// The namespace. Currently, you should set this to `default`.
	Namespace pulumi.StringPtrInput
	// The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
	SessionName pulumi.StringPtrInput
	// The Amazon QuickSight user name that you want to create for the user you are registering.
	UserName pulumi.StringPtrInput
	// The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
	UserRole pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}
