// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ProjectEnvironmentEnvironmentVariable

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ProjectEnvironmentEnvironmentVariable struct {
	// The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
	Name string `pulumi:"name"`
	// The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
	Type *string `pulumi:"type"`
	// The environment variable's value.
	Value string `pulumi:"value"`
}

type ProjectEnvironmentEnvironmentVariableInput interface {
	pulumi.Input

	ToProjectEnvironmentEnvironmentVariableOutput() ProjectEnvironmentEnvironmentVariableOutput
	ToProjectEnvironmentEnvironmentVariableOutputWithContext(context.Context) ProjectEnvironmentEnvironmentVariableOutput
}

type ProjectEnvironmentEnvironmentVariableArgs struct {
	// The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
	Name pulumi.StringInput `pulumi:"name"`
	// The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// The environment variable's value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ProjectEnvironmentEnvironmentVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectEnvironmentEnvironmentVariable)(nil)).Elem()
}

func (i ProjectEnvironmentEnvironmentVariableArgs) ToProjectEnvironmentEnvironmentVariableOutput() ProjectEnvironmentEnvironmentVariableOutput {
	return i.ToProjectEnvironmentEnvironmentVariableOutputWithContext(context.Background())
}

func (i ProjectEnvironmentEnvironmentVariableArgs) ToProjectEnvironmentEnvironmentVariableOutputWithContext(ctx context.Context) ProjectEnvironmentEnvironmentVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentEnvironmentVariableOutput)
}

type ProjectEnvironmentEnvironmentVariableArrayInput interface {
	pulumi.Input

	ToProjectEnvironmentEnvironmentVariableArrayOutput() ProjectEnvironmentEnvironmentVariableArrayOutput
	ToProjectEnvironmentEnvironmentVariableArrayOutputWithContext(context.Context) ProjectEnvironmentEnvironmentVariableArrayOutput
}

type ProjectEnvironmentEnvironmentVariableArray []ProjectEnvironmentEnvironmentVariableInput

func (ProjectEnvironmentEnvironmentVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectEnvironmentEnvironmentVariable)(nil)).Elem()
}

func (i ProjectEnvironmentEnvironmentVariableArray) ToProjectEnvironmentEnvironmentVariableArrayOutput() ProjectEnvironmentEnvironmentVariableArrayOutput {
	return i.ToProjectEnvironmentEnvironmentVariableArrayOutputWithContext(context.Background())
}

func (i ProjectEnvironmentEnvironmentVariableArray) ToProjectEnvironmentEnvironmentVariableArrayOutputWithContext(ctx context.Context) ProjectEnvironmentEnvironmentVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectEnvironmentEnvironmentVariableArrayOutput)
}

type ProjectEnvironmentEnvironmentVariableOutput struct { *pulumi.OutputState }

func (ProjectEnvironmentEnvironmentVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectEnvironmentEnvironmentVariable)(nil)).Elem()
}

func (o ProjectEnvironmentEnvironmentVariableOutput) ToProjectEnvironmentEnvironmentVariableOutput() ProjectEnvironmentEnvironmentVariableOutput {
	return o
}

func (o ProjectEnvironmentEnvironmentVariableOutput) ToProjectEnvironmentEnvironmentVariableOutputWithContext(ctx context.Context) ProjectEnvironmentEnvironmentVariableOutput {
	return o
}

// The name of the project. If `type` is set to `S3`, this is the name of the output artifact object
func (o ProjectEnvironmentEnvironmentVariableOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v ProjectEnvironmentEnvironmentVariable) string { return v.Name }).(pulumi.StringOutput)
}

// The type of repository that contains the source code to be built. Valid values for this parameter are: `CODECOMMIT`, `CODEPIPELINE`, `GITHUB`, `GITHUB_ENTERPRISE`, `BITBUCKET` or `S3`.
func (o ProjectEnvironmentEnvironmentVariableOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ProjectEnvironmentEnvironmentVariable) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The environment variable's value.
func (o ProjectEnvironmentEnvironmentVariableOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func (v ProjectEnvironmentEnvironmentVariable) string { return v.Value }).(pulumi.StringOutput)
}

type ProjectEnvironmentEnvironmentVariableArrayOutput struct { *pulumi.OutputState}

func (ProjectEnvironmentEnvironmentVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProjectEnvironmentEnvironmentVariable)(nil)).Elem()
}

func (o ProjectEnvironmentEnvironmentVariableArrayOutput) ToProjectEnvironmentEnvironmentVariableArrayOutput() ProjectEnvironmentEnvironmentVariableArrayOutput {
	return o
}

func (o ProjectEnvironmentEnvironmentVariableArrayOutput) ToProjectEnvironmentEnvironmentVariableArrayOutputWithContext(ctx context.Context) ProjectEnvironmentEnvironmentVariableArrayOutput {
	return o
}

func (o ProjectEnvironmentEnvironmentVariableArrayOutput) Index(i pulumi.IntInput) ProjectEnvironmentEnvironmentVariableOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ProjectEnvironmentEnvironmentVariable {
		return vs[0].([]ProjectEnvironmentEnvironmentVariable)[vs[1].(int)]
	}).(ProjectEnvironmentEnvironmentVariableOutput)
}

func init() {
	pulumi.RegisterOutputType(ProjectEnvironmentEnvironmentVariableOutput{})
	pulumi.RegisterOutputType(ProjectEnvironmentEnvironmentVariableArrayOutput{})
}
