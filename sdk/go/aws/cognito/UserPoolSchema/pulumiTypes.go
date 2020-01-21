// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package UserPoolSchema

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/cognito/UserPoolSchemaNumberAttributeConstraints"
	"https:/github.com/pulumi/pulumi-aws/cognito/UserPoolSchemaStringAttributeConstraints"
)

type UserPoolSchema struct {
	// The attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
	AttributeDataType string `pulumi:"attributeDataType"`
	// Specifies whether the attribute type is developer only.
	DeveloperOnlyAttribute *bool `pulumi:"developerOnlyAttribute"`
	// Specifies whether the attribute can be changed once it has been created.
	Mutable *bool `pulumi:"mutable"`
	// The name of the attribute.
	Name string `pulumi:"name"`
	// Specifies the constraints for an attribute of the number type.
	NumberAttributeConstraints *cognitoUserPoolSchemaNumberAttributeConstraints.UserPoolSchemaNumberAttributeConstraints `pulumi:"numberAttributeConstraints"`
	// Specifies whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
	Required *bool `pulumi:"required"`
	// -Specifies the constraints for an attribute of the string type.
	StringAttributeConstraints *cognitoUserPoolSchemaStringAttributeConstraints.UserPoolSchemaStringAttributeConstraints `pulumi:"stringAttributeConstraints"`
}

type UserPoolSchemaInput interface {
	pulumi.Input

	ToUserPoolSchemaOutput() UserPoolSchemaOutput
	ToUserPoolSchemaOutputWithContext(context.Context) UserPoolSchemaOutput
}

type UserPoolSchemaArgs struct {
	// The attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
	AttributeDataType pulumi.StringInput `pulumi:"attributeDataType"`
	// Specifies whether the attribute type is developer only.
	DeveloperOnlyAttribute pulumi.BoolPtrInput `pulumi:"developerOnlyAttribute"`
	// Specifies whether the attribute can be changed once it has been created.
	Mutable pulumi.BoolPtrInput `pulumi:"mutable"`
	// The name of the attribute.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies the constraints for an attribute of the number type.
	NumberAttributeConstraints cognitoUserPoolSchemaNumberAttributeConstraints.UserPoolSchemaNumberAttributeConstraintsPtrInput `pulumi:"numberAttributeConstraints"`
	// Specifies whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
	Required pulumi.BoolPtrInput `pulumi:"required"`
	// -Specifies the constraints for an attribute of the string type.
	StringAttributeConstraints cognitoUserPoolSchemaStringAttributeConstraints.UserPoolSchemaStringAttributeConstraintsPtrInput `pulumi:"stringAttributeConstraints"`
}

func (UserPoolSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolSchema)(nil)).Elem()
}

func (i UserPoolSchemaArgs) ToUserPoolSchemaOutput() UserPoolSchemaOutput {
	return i.ToUserPoolSchemaOutputWithContext(context.Background())
}

func (i UserPoolSchemaArgs) ToUserPoolSchemaOutputWithContext(ctx context.Context) UserPoolSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolSchemaOutput)
}

type UserPoolSchemaArrayInput interface {
	pulumi.Input

	ToUserPoolSchemaArrayOutput() UserPoolSchemaArrayOutput
	ToUserPoolSchemaArrayOutputWithContext(context.Context) UserPoolSchemaArrayOutput
}

type UserPoolSchemaArray []UserPoolSchemaInput

func (UserPoolSchemaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserPoolSchema)(nil)).Elem()
}

func (i UserPoolSchemaArray) ToUserPoolSchemaArrayOutput() UserPoolSchemaArrayOutput {
	return i.ToUserPoolSchemaArrayOutputWithContext(context.Background())
}

func (i UserPoolSchemaArray) ToUserPoolSchemaArrayOutputWithContext(ctx context.Context) UserPoolSchemaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPoolSchemaArrayOutput)
}

type UserPoolSchemaOutput struct { *pulumi.OutputState }

func (UserPoolSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserPoolSchema)(nil)).Elem()
}

func (o UserPoolSchemaOutput) ToUserPoolSchemaOutput() UserPoolSchemaOutput {
	return o
}

func (o UserPoolSchemaOutput) ToUserPoolSchemaOutputWithContext(ctx context.Context) UserPoolSchemaOutput {
	return o
}

// The attribute data type. Must be one of `Boolean`, `Number`, `String`, `DateTime`.
func (o UserPoolSchemaOutput) AttributeDataType() pulumi.StringOutput {
	return o.ApplyT(func (v UserPoolSchema) string { return v.AttributeDataType }).(pulumi.StringOutput)
}

// Specifies whether the attribute type is developer only.
func (o UserPoolSchemaOutput) DeveloperOnlyAttribute() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v UserPoolSchema) *bool { return v.DeveloperOnlyAttribute }).(pulumi.BoolPtrOutput)
}

// Specifies whether the attribute can be changed once it has been created.
func (o UserPoolSchemaOutput) Mutable() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v UserPoolSchema) *bool { return v.Mutable }).(pulumi.BoolPtrOutput)
}

// The name of the attribute.
func (o UserPoolSchemaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v UserPoolSchema) string { return v.Name }).(pulumi.StringOutput)
}

// Specifies the constraints for an attribute of the number type.
func (o UserPoolSchemaOutput) NumberAttributeConstraints() cognitoUserPoolSchemaNumberAttributeConstraints.UserPoolSchemaNumberAttributeConstraintsPtrOutput {
	return o.ApplyT(func (v UserPoolSchema) *cognitoUserPoolSchemaNumberAttributeConstraints.UserPoolSchemaNumberAttributeConstraints { return v.NumberAttributeConstraints }).(cognitoUserPoolSchemaNumberAttributeConstraints.UserPoolSchemaNumberAttributeConstraintsPtrOutput)
}

// Specifies whether a user pool attribute is required. If the attribute is required and the user does not provide a value, registration or sign-in will fail.
func (o UserPoolSchemaOutput) Required() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v UserPoolSchema) *bool { return v.Required }).(pulumi.BoolPtrOutput)
}

// -Specifies the constraints for an attribute of the string type.
func (o UserPoolSchemaOutput) StringAttributeConstraints() cognitoUserPoolSchemaStringAttributeConstraints.UserPoolSchemaStringAttributeConstraintsPtrOutput {
	return o.ApplyT(func (v UserPoolSchema) *cognitoUserPoolSchemaStringAttributeConstraints.UserPoolSchemaStringAttributeConstraints { return v.StringAttributeConstraints }).(cognitoUserPoolSchemaStringAttributeConstraints.UserPoolSchemaStringAttributeConstraintsPtrOutput)
}

type UserPoolSchemaArrayOutput struct { *pulumi.OutputState}

func (UserPoolSchemaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UserPoolSchema)(nil)).Elem()
}

func (o UserPoolSchemaArrayOutput) ToUserPoolSchemaArrayOutput() UserPoolSchemaArrayOutput {
	return o
}

func (o UserPoolSchemaArrayOutput) ToUserPoolSchemaArrayOutputWithContext(ctx context.Context) UserPoolSchemaArrayOutput {
	return o
}

func (o UserPoolSchemaArrayOutput) Index(i pulumi.IntInput) UserPoolSchemaOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) UserPoolSchema {
		return vs[0].([]UserPoolSchema)[vs[1].(int)]
	}).(UserPoolSchemaOutput)
}

func init() {
	pulumi.RegisterOutputType(UserPoolSchemaOutput{})
	pulumi.RegisterOutputType(UserPoolSchemaArrayOutput{})
}
