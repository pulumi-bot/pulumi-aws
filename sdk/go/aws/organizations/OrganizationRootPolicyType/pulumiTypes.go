// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package OrganizationRootPolicyType

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type OrganizationRootPolicyType struct {
	// The status of the policy type as it relates to the associated root
	Status *string `pulumi:"status"`
	Type *string `pulumi:"type"`
}

type OrganizationRootPolicyTypeInput interface {
	pulumi.Input

	ToOrganizationRootPolicyTypeOutput() OrganizationRootPolicyTypeOutput
	ToOrganizationRootPolicyTypeOutputWithContext(context.Context) OrganizationRootPolicyTypeOutput
}

type OrganizationRootPolicyTypeArgs struct {
	// The status of the policy type as it relates to the associated root
	Status pulumi.StringPtrInput `pulumi:"status"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (OrganizationRootPolicyTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationRootPolicyType)(nil)).Elem()
}

func (i OrganizationRootPolicyTypeArgs) ToOrganizationRootPolicyTypeOutput() OrganizationRootPolicyTypeOutput {
	return i.ToOrganizationRootPolicyTypeOutputWithContext(context.Background())
}

func (i OrganizationRootPolicyTypeArgs) ToOrganizationRootPolicyTypeOutputWithContext(ctx context.Context) OrganizationRootPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationRootPolicyTypeOutput)
}

type OrganizationRootPolicyTypeArrayInput interface {
	pulumi.Input

	ToOrganizationRootPolicyTypeArrayOutput() OrganizationRootPolicyTypeArrayOutput
	ToOrganizationRootPolicyTypeArrayOutputWithContext(context.Context) OrganizationRootPolicyTypeArrayOutput
}

type OrganizationRootPolicyTypeArray []OrganizationRootPolicyTypeInput

func (OrganizationRootPolicyTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationRootPolicyType)(nil)).Elem()
}

func (i OrganizationRootPolicyTypeArray) ToOrganizationRootPolicyTypeArrayOutput() OrganizationRootPolicyTypeArrayOutput {
	return i.ToOrganizationRootPolicyTypeArrayOutputWithContext(context.Background())
}

func (i OrganizationRootPolicyTypeArray) ToOrganizationRootPolicyTypeArrayOutputWithContext(ctx context.Context) OrganizationRootPolicyTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationRootPolicyTypeArrayOutput)
}

type OrganizationRootPolicyTypeOutput struct { *pulumi.OutputState }

func (OrganizationRootPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationRootPolicyType)(nil)).Elem()
}

func (o OrganizationRootPolicyTypeOutput) ToOrganizationRootPolicyTypeOutput() OrganizationRootPolicyTypeOutput {
	return o
}

func (o OrganizationRootPolicyTypeOutput) ToOrganizationRootPolicyTypeOutputWithContext(ctx context.Context) OrganizationRootPolicyTypeOutput {
	return o
}

// The status of the policy type as it relates to the associated root
func (o OrganizationRootPolicyTypeOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func (v OrganizationRootPolicyType) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o OrganizationRootPolicyTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func (v OrganizationRootPolicyType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type OrganizationRootPolicyTypeArrayOutput struct { *pulumi.OutputState}

func (OrganizationRootPolicyTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OrganizationRootPolicyType)(nil)).Elem()
}

func (o OrganizationRootPolicyTypeArrayOutput) ToOrganizationRootPolicyTypeArrayOutput() OrganizationRootPolicyTypeArrayOutput {
	return o
}

func (o OrganizationRootPolicyTypeArrayOutput) ToOrganizationRootPolicyTypeArrayOutputWithContext(ctx context.Context) OrganizationRootPolicyTypeArrayOutput {
	return o
}

func (o OrganizationRootPolicyTypeArrayOutput) Index(i pulumi.IntInput) OrganizationRootPolicyTypeOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) OrganizationRootPolicyType {
		return vs[0].([]OrganizationRootPolicyType)[vs[1].(int)]
	}).(OrganizationRootPolicyTypeOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationRootPolicyTypeOutput{})
	pulumi.RegisterOutputType(OrganizationRootPolicyTypeArrayOutput{})
}
