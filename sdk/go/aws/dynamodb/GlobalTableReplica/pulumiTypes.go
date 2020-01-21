// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package GlobalTableReplica

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type GlobalTableReplica struct {
	// AWS region name of replica DynamoDB Table. e.g. `us-east-1`
	RegionName string `pulumi:"regionName"`
}

type GlobalTableReplicaInput interface {
	pulumi.Input

	ToGlobalTableReplicaOutput() GlobalTableReplicaOutput
	ToGlobalTableReplicaOutputWithContext(context.Context) GlobalTableReplicaOutput
}

type GlobalTableReplicaArgs struct {
	// AWS region name of replica DynamoDB Table. e.g. `us-east-1`
	RegionName pulumi.StringInput `pulumi:"regionName"`
}

func (GlobalTableReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalTableReplica)(nil)).Elem()
}

func (i GlobalTableReplicaArgs) ToGlobalTableReplicaOutput() GlobalTableReplicaOutput {
	return i.ToGlobalTableReplicaOutputWithContext(context.Background())
}

func (i GlobalTableReplicaArgs) ToGlobalTableReplicaOutputWithContext(ctx context.Context) GlobalTableReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalTableReplicaOutput)
}

type GlobalTableReplicaArrayInput interface {
	pulumi.Input

	ToGlobalTableReplicaArrayOutput() GlobalTableReplicaArrayOutput
	ToGlobalTableReplicaArrayOutputWithContext(context.Context) GlobalTableReplicaArrayOutput
}

type GlobalTableReplicaArray []GlobalTableReplicaInput

func (GlobalTableReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalTableReplica)(nil)).Elem()
}

func (i GlobalTableReplicaArray) ToGlobalTableReplicaArrayOutput() GlobalTableReplicaArrayOutput {
	return i.ToGlobalTableReplicaArrayOutputWithContext(context.Background())
}

func (i GlobalTableReplicaArray) ToGlobalTableReplicaArrayOutputWithContext(ctx context.Context) GlobalTableReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalTableReplicaArrayOutput)
}

type GlobalTableReplicaOutput struct { *pulumi.OutputState }

func (GlobalTableReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalTableReplica)(nil)).Elem()
}

func (o GlobalTableReplicaOutput) ToGlobalTableReplicaOutput() GlobalTableReplicaOutput {
	return o
}

func (o GlobalTableReplicaOutput) ToGlobalTableReplicaOutputWithContext(ctx context.Context) GlobalTableReplicaOutput {
	return o
}

// AWS region name of replica DynamoDB Table. e.g. `us-east-1`
func (o GlobalTableReplicaOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func (v GlobalTableReplica) string { return v.RegionName }).(pulumi.StringOutput)
}

type GlobalTableReplicaArrayOutput struct { *pulumi.OutputState}

func (GlobalTableReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GlobalTableReplica)(nil)).Elem()
}

func (o GlobalTableReplicaArrayOutput) ToGlobalTableReplicaArrayOutput() GlobalTableReplicaArrayOutput {
	return o
}

func (o GlobalTableReplicaArrayOutput) ToGlobalTableReplicaArrayOutputWithContext(ctx context.Context) GlobalTableReplicaArrayOutput {
	return o
}

func (o GlobalTableReplicaArrayOutput) Index(i pulumi.IntInput) GlobalTableReplicaOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) GlobalTableReplica {
		return vs[0].([]GlobalTableReplica)[vs[1].(int)]
	}).(GlobalTableReplicaOutput)
}

func init() {
	pulumi.RegisterOutputType(GlobalTableReplicaOutput{})
	pulumi.RegisterOutputType(GlobalTableReplicaArrayOutput{})
}
