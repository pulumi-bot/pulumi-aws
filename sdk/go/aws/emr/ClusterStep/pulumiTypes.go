// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterStep

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/emr/ClusterStepHadoopJarStep"
)

type ClusterStep struct {
	ActionOnFailure string `pulumi:"actionOnFailure"`
	HadoopJarStep emrClusterStepHadoopJarStep.ClusterStepHadoopJarStep `pulumi:"hadoopJarStep"`
	// The name of the job flow
	Name string `pulumi:"name"`
}

type ClusterStepInput interface {
	pulumi.Input

	ToClusterStepOutput() ClusterStepOutput
	ToClusterStepOutputWithContext(context.Context) ClusterStepOutput
}

type ClusterStepArgs struct {
	ActionOnFailure pulumi.StringInput `pulumi:"actionOnFailure"`
	HadoopJarStep emrClusterStepHadoopJarStep.ClusterStepHadoopJarStepInput `pulumi:"hadoopJarStep"`
	// The name of the job flow
	Name pulumi.StringInput `pulumi:"name"`
}

func (ClusterStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterStep)(nil)).Elem()
}

func (i ClusterStepArgs) ToClusterStepOutput() ClusterStepOutput {
	return i.ToClusterStepOutputWithContext(context.Background())
}

func (i ClusterStepArgs) ToClusterStepOutputWithContext(ctx context.Context) ClusterStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterStepOutput)
}

type ClusterStepArrayInput interface {
	pulumi.Input

	ToClusterStepArrayOutput() ClusterStepArrayOutput
	ToClusterStepArrayOutputWithContext(context.Context) ClusterStepArrayOutput
}

type ClusterStepArray []ClusterStepInput

func (ClusterStepArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterStep)(nil)).Elem()
}

func (i ClusterStepArray) ToClusterStepArrayOutput() ClusterStepArrayOutput {
	return i.ToClusterStepArrayOutputWithContext(context.Background())
}

func (i ClusterStepArray) ToClusterStepArrayOutputWithContext(ctx context.Context) ClusterStepArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterStepArrayOutput)
}

type ClusterStepOutput struct { *pulumi.OutputState }

func (ClusterStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterStep)(nil)).Elem()
}

func (o ClusterStepOutput) ToClusterStepOutput() ClusterStepOutput {
	return o
}

func (o ClusterStepOutput) ToClusterStepOutputWithContext(ctx context.Context) ClusterStepOutput {
	return o
}

func (o ClusterStepOutput) ActionOnFailure() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterStep) string { return v.ActionOnFailure }).(pulumi.StringOutput)
}

func (o ClusterStepOutput) HadoopJarStep() emrClusterStepHadoopJarStep.ClusterStepHadoopJarStepOutput {
	return o.ApplyT(func (v ClusterStep) emrClusterStepHadoopJarStep.ClusterStepHadoopJarStep { return v.HadoopJarStep }).(emrClusterStepHadoopJarStep.ClusterStepHadoopJarStepOutput)
}

// The name of the job flow
func (o ClusterStepOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterStep) string { return v.Name }).(pulumi.StringOutput)
}

type ClusterStepArrayOutput struct { *pulumi.OutputState}

func (ClusterStepArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterStep)(nil)).Elem()
}

func (o ClusterStepArrayOutput) ToClusterStepArrayOutput() ClusterStepArrayOutput {
	return o
}

func (o ClusterStepArrayOutput) ToClusterStepArrayOutputWithContext(ctx context.Context) ClusterStepArrayOutput {
	return o
}

func (o ClusterStepArrayOutput) Index(i pulumi.IntInput) ClusterStepOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ClusterStep {
		return vs[0].([]ClusterStep)[vs[1].(int)]
	}).(ClusterStepOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterStepOutput{})
	pulumi.RegisterOutputType(ClusterStepArrayOutput{})
}
