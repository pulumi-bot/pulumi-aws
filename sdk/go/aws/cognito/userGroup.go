// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Cognito User Group resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cognito"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mainUserPool, err := cognito.NewUserPool(ctx, "mainUserPool", nil)
// 		if err != nil {
// 			return err
// 		}
// 		groupRole, err := iam.NewRole(ctx, "groupRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Federated\": \"cognito-identity.amazonaws.com\"\n", "      },\n", "      \"Action\": \"sts:AssumeRoleWithWebIdentity\",\n", "      \"Condition\": {\n", "        \"StringEquals\": {\n", "          \"cognito-identity.amazonaws.com:aud\": \"us-east-1:12345678-dead-beef-cafe-123456790ab\"\n", "        },\n", "        \"ForAnyValue:StringLike\": {\n", "          \"cognito-identity.amazonaws.com:amr\": \"authenticated\"\n", "        }\n", "      }\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cognito.NewUserGroup(ctx, "mainUserGroup", &cognito.UserGroupArgs{
// 			UserPoolId:  mainUserPool.ID(),
// 			Description: pulumi.String("Managed by Pulumi"),
// 			Precedence:  pulumi.Int(42),
// 			RoleArn:     groupRole.Arn,
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
// Cognito User Groups can be imported using the `user_pool_id`/`name` attributes concatenated, e.g.
//
// ```sh
//  $ pulumi import aws:cognito/userGroup:UserGroup group us-east-1_vG78M4goG/user-group
// ```
type UserGroup struct {
	pulumi.CustomResourceState

	// The description of the user group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the user group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The precedence of the user group.
	Precedence pulumi.IntPtrOutput `pulumi:"precedence"`
	// The ARN of the IAM role to be associated with the user group.
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// The user pool ID.
	UserPoolId pulumi.StringOutput `pulumi:"userPoolId"`
}

// NewUserGroup registers a new resource with the given unique name, arguments, and options.
func NewUserGroup(ctx *pulumi.Context,
	name string, args *UserGroupArgs, opts ...pulumi.ResourceOption) (*UserGroup, error) {
	if args == nil || args.UserPoolId == nil {
		return nil, errors.New("missing required argument 'UserPoolId'")
	}
	if args == nil {
		args = &UserGroupArgs{}
	}
	var resource UserGroup
	err := ctx.RegisterResource("aws:cognito/userGroup:UserGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGroup gets an existing UserGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGroupState, opts ...pulumi.ResourceOption) (*UserGroup, error) {
	var resource UserGroup
	err := ctx.ReadResource("aws:cognito/userGroup:UserGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGroup resources.
type userGroupState struct {
	// The description of the user group.
	Description *string `pulumi:"description"`
	// The name of the user group.
	Name *string `pulumi:"name"`
	// The precedence of the user group.
	Precedence *int `pulumi:"precedence"`
	// The ARN of the IAM role to be associated with the user group.
	RoleArn *string `pulumi:"roleArn"`
	// The user pool ID.
	UserPoolId *string `pulumi:"userPoolId"`
}

type UserGroupState struct {
	// The description of the user group.
	Description pulumi.StringPtrInput
	// The name of the user group.
	Name pulumi.StringPtrInput
	// The precedence of the user group.
	Precedence pulumi.IntPtrInput
	// The ARN of the IAM role to be associated with the user group.
	RoleArn pulumi.StringPtrInput
	// The user pool ID.
	UserPoolId pulumi.StringPtrInput
}

func (UserGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupState)(nil)).Elem()
}

type userGroupArgs struct {
	// The description of the user group.
	Description *string `pulumi:"description"`
	// The name of the user group.
	Name *string `pulumi:"name"`
	// The precedence of the user group.
	Precedence *int `pulumi:"precedence"`
	// The ARN of the IAM role to be associated with the user group.
	RoleArn *string `pulumi:"roleArn"`
	// The user pool ID.
	UserPoolId string `pulumi:"userPoolId"`
}

// The set of arguments for constructing a UserGroup resource.
type UserGroupArgs struct {
	// The description of the user group.
	Description pulumi.StringPtrInput
	// The name of the user group.
	Name pulumi.StringPtrInput
	// The precedence of the user group.
	Precedence pulumi.IntPtrInput
	// The ARN of the IAM role to be associated with the user group.
	RoleArn pulumi.StringPtrInput
	// The user pool ID.
	UserPoolId pulumi.StringInput
}

func (UserGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupArgs)(nil)).Elem()
}

type UserGroupInput interface {
	pulumi.Input

	ToUserGroupOutput() UserGroupOutput
	ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput
}

func (UserGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*UserGroup)(nil)).Elem()
}

func (i UserGroup) ToUserGroupOutput() UserGroupOutput {
	return i.ToUserGroupOutputWithContext(context.Background())
}

func (i UserGroup) ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupOutput)
}

type UserGroupOutput struct {
	*pulumi.OutputState
}

func (UserGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserGroupOutput)(nil)).Elem()
}

func (o UserGroupOutput) ToUserGroupOutput() UserGroupOutput {
	return o
}

func (o UserGroupOutput) ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserGroupOutput{})
}
