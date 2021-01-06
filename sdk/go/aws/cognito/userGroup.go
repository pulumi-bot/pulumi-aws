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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserPoolId == nil {
		return nil, errors.New("invalid value for required argument 'UserPoolId'")
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

func (*UserGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*UserGroup)(nil))
}

func (i *UserGroup) ToUserGroupOutput() UserGroupOutput {
	return i.ToUserGroupOutputWithContext(context.Background())
}

func (i *UserGroup) ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupOutput)
}

func (i *UserGroup) ToUserGroupPtrOutput() UserGroupPtrOutput {
	return i.ToUserGroupPtrOutputWithContext(context.Background())
}

func (i *UserGroup) ToUserGroupPtrOutputWithContext(ctx context.Context) UserGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupPtrOutput)
}

type UserGroupPtrInput interface {
	pulumi.Input

	ToUserGroupPtrOutput() UserGroupPtrOutput
	ToUserGroupPtrOutputWithContext(ctx context.Context) UserGroupPtrOutput
}

type userGroupPtrType UserGroupArgs

func (*userGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroup)(nil))
}

func (i *userGroupPtrType) ToUserGroupPtrOutput() UserGroupPtrOutput {
	return i.ToUserGroupPtrOutputWithContext(context.Background())
}

func (i *userGroupPtrType) ToUserGroupPtrOutputWithContext(ctx context.Context) UserGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupOutput).ToUserGroupPtrOutput()
}

// UserGroupArrayInput is an input type that accepts UserGroupArray and UserGroupArrayOutput values.
// You can construct a concrete instance of `UserGroupArrayInput` via:
//
//          UserGroupArray{ UserGroupArgs{...} }
type UserGroupArrayInput interface {
	pulumi.Input

	ToUserGroupArrayOutput() UserGroupArrayOutput
	ToUserGroupArrayOutputWithContext(context.Context) UserGroupArrayOutput
}

type UserGroupArray []UserGroupInput

func (UserGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserGroup)(nil))
}

func (i UserGroupArray) ToUserGroupArrayOutput() UserGroupArrayOutput {
	return i.ToUserGroupArrayOutputWithContext(context.Background())
}

func (i UserGroupArray) ToUserGroupArrayOutputWithContext(ctx context.Context) UserGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupArrayOutput)
}

// UserGroupMapInput is an input type that accepts UserGroupMap and UserGroupMapOutput values.
// You can construct a concrete instance of `UserGroupMapInput` via:
//
//          UserGroupMap{ "key": UserGroupArgs{...} }
type UserGroupMapInput interface {
	pulumi.Input

	ToUserGroupMapOutput() UserGroupMapOutput
	ToUserGroupMapOutputWithContext(context.Context) UserGroupMapOutput
}

type UserGroupMap map[string]UserGroupInput

func (UserGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserGroup)(nil))
}

func (i UserGroupMap) ToUserGroupMapOutput() UserGroupMapOutput {
	return i.ToUserGroupMapOutputWithContext(context.Background())
}

func (i UserGroupMap) ToUserGroupMapOutputWithContext(ctx context.Context) UserGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupMapOutput)
}

type UserGroupOutput struct {
	*pulumi.OutputState
}

func (UserGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserGroup)(nil))
}

func (o UserGroupOutput) ToUserGroupOutput() UserGroupOutput {
	return o
}

func (o UserGroupOutput) ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput {
	return o
}

func (o UserGroupOutput) ToUserGroupPtrOutput() UserGroupPtrOutput {
	return o.ToUserGroupPtrOutputWithContext(context.Background())
}

func (o UserGroupOutput) ToUserGroupPtrOutputWithContext(ctx context.Context) UserGroupPtrOutput {
	return o.ApplyT(func(v UserGroup) *UserGroup {
		return &v
	}).(UserGroupPtrOutput)
}

type UserGroupPtrOutput struct {
	*pulumi.OutputState
}

func (UserGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroup)(nil))
}

func (o UserGroupPtrOutput) ToUserGroupPtrOutput() UserGroupPtrOutput {
	return o
}

func (o UserGroupPtrOutput) ToUserGroupPtrOutputWithContext(ctx context.Context) UserGroupPtrOutput {
	return o
}

type UserGroupArrayOutput struct{ *pulumi.OutputState }

func (UserGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserGroup)(nil))
}

func (o UserGroupArrayOutput) ToUserGroupArrayOutput() UserGroupArrayOutput {
	return o
}

func (o UserGroupArrayOutput) ToUserGroupArrayOutputWithContext(ctx context.Context) UserGroupArrayOutput {
	return o
}

func (o UserGroupArrayOutput) Index(i pulumi.IntInput) UserGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UserGroup {
		return vs[0].([]UserGroup)[vs[1].(int)]
	}).(UserGroupOutput)
}

type UserGroupMapOutput struct{ *pulumi.OutputState }

func (UserGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserGroup)(nil))
}

func (o UserGroupMapOutput) ToUserGroupMapOutput() UserGroupMapOutput {
	return o
}

func (o UserGroupMapOutput) ToUserGroupMapOutputWithContext(ctx context.Context) UserGroupMapOutput {
	return o
}

func (o UserGroupMapOutput) MapIndex(k pulumi.StringInput) UserGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserGroup {
		return vs[0].(map[string]UserGroup)[vs[1].(string)]
	}).(UserGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(UserGroupOutput{})
	pulumi.RegisterOutputType(UserGroupPtrOutput{})
	pulumi.RegisterOutputType(UserGroupArrayOutput{})
	pulumi.RegisterOutputType(UserGroupMapOutput{})
}
