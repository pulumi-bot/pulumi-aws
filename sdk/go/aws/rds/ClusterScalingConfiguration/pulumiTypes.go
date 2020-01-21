// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterScalingConfiguration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterScalingConfiguration struct {
	// Whether to enable automatic pause. A DB cluster can be paused only when it's idle (it has no connections). If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it. Defaults to `true`.
	AutoPause *bool `pulumi:"autoPause"`
	// The maximum capacity. The maximum capacity must be greater than or equal to the minimum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `16`.
	MaxCapacity *int `pulumi:"maxCapacity"`
	// The minimum capacity. The minimum capacity must be lesser than or equal to the maximum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `2`.
	MinCapacity *int `pulumi:"minCapacity"`
	// The time, in seconds, before an Aurora DB cluster in serverless mode is paused. Valid values are `300` through `86400`. Defaults to `300`.
	SecondsUntilAutoPause *int `pulumi:"secondsUntilAutoPause"`
	// The action to take when the timeout is reached. Valid values: `ForceApplyCapacityChange`, `RollbackCapacityChange`. Defaults to `RollbackCapacityChange`. See [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.timeout-action).
	TimeoutAction *string `pulumi:"timeoutAction"`
}

type ClusterScalingConfigurationInput interface {
	pulumi.Input

	ToClusterScalingConfigurationOutput() ClusterScalingConfigurationOutput
	ToClusterScalingConfigurationOutputWithContext(context.Context) ClusterScalingConfigurationOutput
}

type ClusterScalingConfigurationArgs struct {
	// Whether to enable automatic pause. A DB cluster can be paused only when it's idle (it has no connections). If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it. Defaults to `true`.
	AutoPause pulumi.BoolPtrInput `pulumi:"autoPause"`
	// The maximum capacity. The maximum capacity must be greater than or equal to the minimum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `16`.
	MaxCapacity pulumi.IntPtrInput `pulumi:"maxCapacity"`
	// The minimum capacity. The minimum capacity must be lesser than or equal to the maximum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `2`.
	MinCapacity pulumi.IntPtrInput `pulumi:"minCapacity"`
	// The time, in seconds, before an Aurora DB cluster in serverless mode is paused. Valid values are `300` through `86400`. Defaults to `300`.
	SecondsUntilAutoPause pulumi.IntPtrInput `pulumi:"secondsUntilAutoPause"`
	// The action to take when the timeout is reached. Valid values: `ForceApplyCapacityChange`, `RollbackCapacityChange`. Defaults to `RollbackCapacityChange`. See [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.timeout-action).
	TimeoutAction pulumi.StringPtrInput `pulumi:"timeoutAction"`
}

func (ClusterScalingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterScalingConfiguration)(nil)).Elem()
}

func (i ClusterScalingConfigurationArgs) ToClusterScalingConfigurationOutput() ClusterScalingConfigurationOutput {
	return i.ToClusterScalingConfigurationOutputWithContext(context.Background())
}

func (i ClusterScalingConfigurationArgs) ToClusterScalingConfigurationOutputWithContext(ctx context.Context) ClusterScalingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterScalingConfigurationOutput)
}

func (i ClusterScalingConfigurationArgs) ToClusterScalingConfigurationPtrOutput() ClusterScalingConfigurationPtrOutput {
	return i.ToClusterScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i ClusterScalingConfigurationArgs) ToClusterScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterScalingConfigurationOutput).ToClusterScalingConfigurationPtrOutputWithContext(ctx)
}

type ClusterScalingConfigurationPtrInput interface {
	pulumi.Input

	ToClusterScalingConfigurationPtrOutput() ClusterScalingConfigurationPtrOutput
	ToClusterScalingConfigurationPtrOutputWithContext(context.Context) ClusterScalingConfigurationPtrOutput
}

type clusterScalingConfigurationPtrType ClusterScalingConfigurationArgs

func ClusterScalingConfigurationPtr(v *ClusterScalingConfigurationArgs) ClusterScalingConfigurationPtrInput {	return (*clusterScalingConfigurationPtrType)(v)
}

func (*clusterScalingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterScalingConfiguration)(nil)).Elem()
}

func (i *clusterScalingConfigurationPtrType) ToClusterScalingConfigurationPtrOutput() ClusterScalingConfigurationPtrOutput {
	return i.ToClusterScalingConfigurationPtrOutputWithContext(context.Background())
}

func (i *clusterScalingConfigurationPtrType) ToClusterScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterScalingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterScalingConfigurationPtrOutput)
}

type ClusterScalingConfigurationOutput struct { *pulumi.OutputState }

func (ClusterScalingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterScalingConfiguration)(nil)).Elem()
}

func (o ClusterScalingConfigurationOutput) ToClusterScalingConfigurationOutput() ClusterScalingConfigurationOutput {
	return o
}

func (o ClusterScalingConfigurationOutput) ToClusterScalingConfigurationOutputWithContext(ctx context.Context) ClusterScalingConfigurationOutput {
	return o
}

func (o ClusterScalingConfigurationOutput) ToClusterScalingConfigurationPtrOutput() ClusterScalingConfigurationPtrOutput {
	return o.ToClusterScalingConfigurationPtrOutputWithContext(context.Background())
}

func (o ClusterScalingConfigurationOutput) ToClusterScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterScalingConfigurationPtrOutput {
	return o.ApplyT(func(v ClusterScalingConfiguration) *ClusterScalingConfiguration {
		return &v
	}).(ClusterScalingConfigurationPtrOutput)
}
// Whether to enable automatic pause. A DB cluster can be paused only when it's idle (it has no connections). If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it. Defaults to `true`.
func (o ClusterScalingConfigurationOutput) AutoPause() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *bool { return v.AutoPause }).(pulumi.BoolPtrOutput)
}

// The maximum capacity. The maximum capacity must be greater than or equal to the minimum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `16`.
func (o ClusterScalingConfigurationOutput) MaxCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *int { return v.MaxCapacity }).(pulumi.IntPtrOutput)
}

// The minimum capacity. The minimum capacity must be lesser than or equal to the maximum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `2`.
func (o ClusterScalingConfigurationOutput) MinCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *int { return v.MinCapacity }).(pulumi.IntPtrOutput)
}

// The time, in seconds, before an Aurora DB cluster in serverless mode is paused. Valid values are `300` through `86400`. Defaults to `300`.
func (o ClusterScalingConfigurationOutput) SecondsUntilAutoPause() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *int { return v.SecondsUntilAutoPause }).(pulumi.IntPtrOutput)
}

// The action to take when the timeout is reached. Valid values: `ForceApplyCapacityChange`, `RollbackCapacityChange`. Defaults to `RollbackCapacityChange`. See [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.timeout-action).
func (o ClusterScalingConfigurationOutput) TimeoutAction() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *string { return v.TimeoutAction }).(pulumi.StringPtrOutput)
}

type ClusterScalingConfigurationPtrOutput struct { *pulumi.OutputState}

func (ClusterScalingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterScalingConfiguration)(nil)).Elem()
}

func (o ClusterScalingConfigurationPtrOutput) ToClusterScalingConfigurationPtrOutput() ClusterScalingConfigurationPtrOutput {
	return o
}

func (o ClusterScalingConfigurationPtrOutput) ToClusterScalingConfigurationPtrOutputWithContext(ctx context.Context) ClusterScalingConfigurationPtrOutput {
	return o
}

func (o ClusterScalingConfigurationPtrOutput) Elem() ClusterScalingConfigurationOutput {
	return o.ApplyT(func (v *ClusterScalingConfiguration) ClusterScalingConfiguration { return *v }).(ClusterScalingConfigurationOutput)
}

// Whether to enable automatic pause. A DB cluster can be paused only when it's idle (it has no connections). If a DB cluster is paused for more than seven days, the DB cluster might be backed up with a snapshot. In this case, the DB cluster is restored when there is a request to connect to it. Defaults to `true`.
func (o ClusterScalingConfigurationPtrOutput) AutoPause() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *bool { return v.AutoPause }).(pulumi.BoolPtrOutput)
}

// The maximum capacity. The maximum capacity must be greater than or equal to the minimum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `16`.
func (o ClusterScalingConfigurationPtrOutput) MaxCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *int { return v.MaxCapacity }).(pulumi.IntPtrOutput)
}

// The minimum capacity. The minimum capacity must be lesser than or equal to the maximum capacity. Valid capacity values are `1`, `2`, `4`, `8`, `16`, `32`, `64`, `128`, and `256`. Defaults to `2`.
func (o ClusterScalingConfigurationPtrOutput) MinCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *int { return v.MinCapacity }).(pulumi.IntPtrOutput)
}

// The time, in seconds, before an Aurora DB cluster in serverless mode is paused. Valid values are `300` through `86400`. Defaults to `300`.
func (o ClusterScalingConfigurationPtrOutput) SecondsUntilAutoPause() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *int { return v.SecondsUntilAutoPause }).(pulumi.IntPtrOutput)
}

// The action to take when the timeout is reached. Valid values: `ForceApplyCapacityChange`, `RollbackCapacityChange`. Defaults to `RollbackCapacityChange`. See [documentation](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.how-it-works.html#aurora-serverless.how-it-works.timeout-action).
func (o ClusterScalingConfigurationPtrOutput) TimeoutAction() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterScalingConfiguration) *string { return v.TimeoutAction }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterScalingConfigurationOutput{})
	pulumi.RegisterOutputType(ClusterScalingConfigurationPtrOutput{})
}
