// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterMasterInstanceGroup

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/emr/ClusterMasterInstanceGroupEbsConfig"
)

type ClusterMasterInstanceGroup struct {
	BidPrice *string `pulumi:"bidPrice"`
	EbsConfigs []emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfig `pulumi:"ebsConfigs"`
	// The ID of the EMR Cluster
	Id *string `pulumi:"id"`
	InstanceCount *int `pulumi:"instanceCount"`
	InstanceType string `pulumi:"instanceType"`
	// The name of the job flow
	Name *string `pulumi:"name"`
}

type ClusterMasterInstanceGroupInput interface {
	pulumi.Input

	ToClusterMasterInstanceGroupOutput() ClusterMasterInstanceGroupOutput
	ToClusterMasterInstanceGroupOutputWithContext(context.Context) ClusterMasterInstanceGroupOutput
}

type ClusterMasterInstanceGroupArgs struct {
	BidPrice pulumi.StringPtrInput `pulumi:"bidPrice"`
	EbsConfigs emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfigArrayInput `pulumi:"ebsConfigs"`
	// The ID of the EMR Cluster
	Id pulumi.StringPtrInput `pulumi:"id"`
	InstanceCount pulumi.IntPtrInput `pulumi:"instanceCount"`
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// The name of the job flow
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ClusterMasterInstanceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterMasterInstanceGroup)(nil)).Elem()
}

func (i ClusterMasterInstanceGroupArgs) ToClusterMasterInstanceGroupOutput() ClusterMasterInstanceGroupOutput {
	return i.ToClusterMasterInstanceGroupOutputWithContext(context.Background())
}

func (i ClusterMasterInstanceGroupArgs) ToClusterMasterInstanceGroupOutputWithContext(ctx context.Context) ClusterMasterInstanceGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMasterInstanceGroupOutput)
}

func (i ClusterMasterInstanceGroupArgs) ToClusterMasterInstanceGroupPtrOutput() ClusterMasterInstanceGroupPtrOutput {
	return i.ToClusterMasterInstanceGroupPtrOutputWithContext(context.Background())
}

func (i ClusterMasterInstanceGroupArgs) ToClusterMasterInstanceGroupPtrOutputWithContext(ctx context.Context) ClusterMasterInstanceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMasterInstanceGroupOutput).ToClusterMasterInstanceGroupPtrOutputWithContext(ctx)
}

type ClusterMasterInstanceGroupPtrInput interface {
	pulumi.Input

	ToClusterMasterInstanceGroupPtrOutput() ClusterMasterInstanceGroupPtrOutput
	ToClusterMasterInstanceGroupPtrOutputWithContext(context.Context) ClusterMasterInstanceGroupPtrOutput
}

type clusterMasterInstanceGroupPtrType ClusterMasterInstanceGroupArgs

func ClusterMasterInstanceGroupPtr(v *ClusterMasterInstanceGroupArgs) ClusterMasterInstanceGroupPtrInput {	return (*clusterMasterInstanceGroupPtrType)(v)
}

func (*clusterMasterInstanceGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterMasterInstanceGroup)(nil)).Elem()
}

func (i *clusterMasterInstanceGroupPtrType) ToClusterMasterInstanceGroupPtrOutput() ClusterMasterInstanceGroupPtrOutput {
	return i.ToClusterMasterInstanceGroupPtrOutputWithContext(context.Background())
}

func (i *clusterMasterInstanceGroupPtrType) ToClusterMasterInstanceGroupPtrOutputWithContext(ctx context.Context) ClusterMasterInstanceGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterMasterInstanceGroupPtrOutput)
}

type ClusterMasterInstanceGroupOutput struct { *pulumi.OutputState }

func (ClusterMasterInstanceGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterMasterInstanceGroup)(nil)).Elem()
}

func (o ClusterMasterInstanceGroupOutput) ToClusterMasterInstanceGroupOutput() ClusterMasterInstanceGroupOutput {
	return o
}

func (o ClusterMasterInstanceGroupOutput) ToClusterMasterInstanceGroupOutputWithContext(ctx context.Context) ClusterMasterInstanceGroupOutput {
	return o
}

func (o ClusterMasterInstanceGroupOutput) ToClusterMasterInstanceGroupPtrOutput() ClusterMasterInstanceGroupPtrOutput {
	return o.ToClusterMasterInstanceGroupPtrOutputWithContext(context.Background())
}

func (o ClusterMasterInstanceGroupOutput) ToClusterMasterInstanceGroupPtrOutputWithContext(ctx context.Context) ClusterMasterInstanceGroupPtrOutput {
	return o.ApplyT(func(v ClusterMasterInstanceGroup) *ClusterMasterInstanceGroup {
		return &v
	}).(ClusterMasterInstanceGroupPtrOutput)
}
func (o ClusterMasterInstanceGroupOutput) BidPrice() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *string { return v.BidPrice }).(pulumi.StringPtrOutput)
}

func (o ClusterMasterInstanceGroupOutput) EbsConfigs() emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfigArrayOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) []emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfig { return v.EbsConfigs }).(emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfigArrayOutput)
}

// The ID of the EMR Cluster
func (o ClusterMasterInstanceGroupOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ClusterMasterInstanceGroupOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func (o ClusterMasterInstanceGroupOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) string { return v.InstanceType }).(pulumi.StringOutput)
}

// The name of the job flow
func (o ClusterMasterInstanceGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ClusterMasterInstanceGroupPtrOutput struct { *pulumi.OutputState}

func (ClusterMasterInstanceGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterMasterInstanceGroup)(nil)).Elem()
}

func (o ClusterMasterInstanceGroupPtrOutput) ToClusterMasterInstanceGroupPtrOutput() ClusterMasterInstanceGroupPtrOutput {
	return o
}

func (o ClusterMasterInstanceGroupPtrOutput) ToClusterMasterInstanceGroupPtrOutputWithContext(ctx context.Context) ClusterMasterInstanceGroupPtrOutput {
	return o
}

func (o ClusterMasterInstanceGroupPtrOutput) Elem() ClusterMasterInstanceGroupOutput {
	return o.ApplyT(func (v *ClusterMasterInstanceGroup) ClusterMasterInstanceGroup { return *v }).(ClusterMasterInstanceGroupOutput)
}

func (o ClusterMasterInstanceGroupPtrOutput) BidPrice() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *string { return v.BidPrice }).(pulumi.StringPtrOutput)
}

func (o ClusterMasterInstanceGroupPtrOutput) EbsConfigs() emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfigArrayOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) []emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfig { return v.EbsConfigs }).(emrClusterMasterInstanceGroupEbsConfig.ClusterMasterInstanceGroupEbsConfigArrayOutput)
}

// The ID of the EMR Cluster
func (o ClusterMasterInstanceGroupPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ClusterMasterInstanceGroupPtrOutput) InstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *int { return v.InstanceCount }).(pulumi.IntPtrOutput)
}

func (o ClusterMasterInstanceGroupPtrOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) string { return v.InstanceType }).(pulumi.StringOutput)
}

// The name of the job flow
func (o ClusterMasterInstanceGroupPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterMasterInstanceGroup) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterMasterInstanceGroupOutput{})
	pulumi.RegisterOutputType(ClusterMasterInstanceGroupPtrOutput{})
}
