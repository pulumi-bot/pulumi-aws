// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getOrganizationAccount

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GetOrganizationAccount struct {
	// ARN of the root
	Arn string `pulumi:"arn"`
	// Email of the account
	Email string `pulumi:"email"`
	// Identifier of the root
	Id string `pulumi:"id"`
	// The name of the policy type
	Name string `pulumi:"name"`
}

type GetOrganizationAccountInput interface {
	pulumi.Input

	ToGetOrganizationAccountOutput() GetOrganizationAccountOutput
	ToGetOrganizationAccountOutputWithContext(context.Context) GetOrganizationAccountOutput
}

type GetOrganizationAccountArgs struct {
	// ARN of the root
	Arn pulumi.StringInput `pulumi:"arn"`
	// Email of the account
	Email pulumi.StringInput `pulumi:"email"`
	// Identifier of the root
	Id pulumi.StringInput `pulumi:"id"`
	// The name of the policy type
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetOrganizationAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationAccount)(nil)).Elem()
}

func (i GetOrganizationAccountArgs) ToGetOrganizationAccountOutput() GetOrganizationAccountOutput {
	return i.ToGetOrganizationAccountOutputWithContext(context.Background())
}

func (i GetOrganizationAccountArgs) ToGetOrganizationAccountOutputWithContext(ctx context.Context) GetOrganizationAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetOrganizationAccountOutput)
}

type GetOrganizationAccountArrayInput interface {
	pulumi.Input

	ToGetOrganizationAccountArrayOutput() GetOrganizationAccountArrayOutput
	ToGetOrganizationAccountArrayOutputWithContext(context.Context) GetOrganizationAccountArrayOutput
}

type GetOrganizationAccountArray []GetOrganizationAccountInput

func (GetOrganizationAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetOrganizationAccount)(nil)).Elem()
}

func (i GetOrganizationAccountArray) ToGetOrganizationAccountArrayOutput() GetOrganizationAccountArrayOutput {
	return i.ToGetOrganizationAccountArrayOutputWithContext(context.Background())
}

func (i GetOrganizationAccountArray) ToGetOrganizationAccountArrayOutputWithContext(ctx context.Context) GetOrganizationAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetOrganizationAccountArrayOutput)
}

type GetOrganizationAccountOutput struct { *pulumi.OutputState }

func (GetOrganizationAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationAccount)(nil)).Elem()
}

func (o GetOrganizationAccountOutput) ToGetOrganizationAccountOutput() GetOrganizationAccountOutput {
	return o
}

func (o GetOrganizationAccountOutput) ToGetOrganizationAccountOutputWithContext(ctx context.Context) GetOrganizationAccountOutput {
	return o
}

// ARN of the root
func (o GetOrganizationAccountOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func (v GetOrganizationAccount) string { return v.Arn }).(pulumi.StringOutput)
}

// Email of the account
func (o GetOrganizationAccountOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func (v GetOrganizationAccount) string { return v.Email }).(pulumi.StringOutput)
}

// Identifier of the root
func (o GetOrganizationAccountOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func (v GetOrganizationAccount) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the policy type
func (o GetOrganizationAccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v GetOrganizationAccount) string { return v.Name }).(pulumi.StringOutput)
}

type GetOrganizationAccountArrayOutput struct { *pulumi.OutputState}

func (GetOrganizationAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetOrganizationAccount)(nil)).Elem()
}

func (o GetOrganizationAccountArrayOutput) ToGetOrganizationAccountArrayOutput() GetOrganizationAccountArrayOutput {
	return o
}

func (o GetOrganizationAccountArrayOutput) ToGetOrganizationAccountArrayOutputWithContext(ctx context.Context) GetOrganizationAccountArrayOutput {
	return o
}

func (o GetOrganizationAccountArrayOutput) Index(i pulumi.IntInput) GetOrganizationAccountOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GetOrganizationAccount {
		return vs[0].([]GetOrganizationAccount)[vs[1].(int)]
	}).(GetOrganizationAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationAccountOutput{})
	pulumi.RegisterOutputType(GetOrganizationAccountArrayOutput{})
}
