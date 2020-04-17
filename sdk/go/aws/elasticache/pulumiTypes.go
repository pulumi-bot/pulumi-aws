// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticache

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClusterCacheNode struct {
	Address *string `pulumi:"address"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	Id               *string `pulumi:"id"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port *int `pulumi:"port"`
}

// ClusterCacheNodeInput is an input type that accepts ClusterCacheNodeArgs and ClusterCacheNodeOutput values.
// You can construct a concrete instance of `ClusterCacheNodeInput` via:
//
// 		 ClusterCacheNodeArgs{...}
//
type ClusterCacheNodeInput interface {
	pulumi.Input

	ToClusterCacheNodeOutput() ClusterCacheNodeOutput
	ToClusterCacheNodeOutputWithContext(context.Context) ClusterCacheNodeOutput
}

type ClusterCacheNodeArgs struct {
	Address pulumi.StringPtrInput `pulumi:"address"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	Id               pulumi.StringPtrInput `pulumi:"id"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port pulumi.IntPtrInput `pulumi:"port"`
}

func (ClusterCacheNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCacheNode)(nil)).Elem()
}

func (i ClusterCacheNodeArgs) ToClusterCacheNodeOutput() ClusterCacheNodeOutput {
	return i.ToClusterCacheNodeOutputWithContext(context.Background())
}

func (i ClusterCacheNodeArgs) ToClusterCacheNodeOutputWithContext(ctx context.Context) ClusterCacheNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCacheNodeOutput)
}

// ClusterCacheNodeArrayInput is an input type that accepts ClusterCacheNodeArray and ClusterCacheNodeArrayOutput values.
// You can construct a concrete instance of `ClusterCacheNodeArrayInput` via:
//
// 		 ClusterCacheNodeArray{ ClusterCacheNodeArgs{...} }
//
type ClusterCacheNodeArrayInput interface {
	pulumi.Input

	ToClusterCacheNodeArrayOutput() ClusterCacheNodeArrayOutput
	ToClusterCacheNodeArrayOutputWithContext(context.Context) ClusterCacheNodeArrayOutput
}

type ClusterCacheNodeArray []ClusterCacheNodeInput

func (ClusterCacheNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterCacheNode)(nil)).Elem()
}

func (i ClusterCacheNodeArray) ToClusterCacheNodeArrayOutput() ClusterCacheNodeArrayOutput {
	return i.ToClusterCacheNodeArrayOutputWithContext(context.Background())
}

func (i ClusterCacheNodeArray) ToClusterCacheNodeArrayOutputWithContext(ctx context.Context) ClusterCacheNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCacheNodeArrayOutput)
}

type ClusterCacheNodeOutput struct{ *pulumi.OutputState }

func (ClusterCacheNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCacheNode)(nil)).Elem()
}

func (o ClusterCacheNodeOutput) ToClusterCacheNodeOutput() ClusterCacheNodeOutput {
	return o
}

func (o ClusterCacheNodeOutput) ToClusterCacheNodeOutputWithContext(ctx context.Context) ClusterCacheNodeOutput {
	return o
}

func (o ClusterCacheNodeOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterCacheNode) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
func (o ClusterCacheNodeOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterCacheNode) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o ClusterCacheNodeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterCacheNode) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
func (o ClusterCacheNodeOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterCacheNode) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ClusterCacheNodeArrayOutput struct{ *pulumi.OutputState }

func (ClusterCacheNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterCacheNode)(nil)).Elem()
}

func (o ClusterCacheNodeArrayOutput) ToClusterCacheNodeArrayOutput() ClusterCacheNodeArrayOutput {
	return o
}

func (o ClusterCacheNodeArrayOutput) ToClusterCacheNodeArrayOutputWithContext(ctx context.Context) ClusterCacheNodeArrayOutput {
	return o
}

func (o ClusterCacheNodeArrayOutput) Index(i pulumi.IntInput) ClusterCacheNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterCacheNode {
		return vs[0].([]ClusterCacheNode)[vs[1].(int)]
	}).(ClusterCacheNodeOutput)
}

type ParameterGroupParameter struct {
	// The name of the ElastiCache parameter group.
	Name string `pulumi:"name"`
	// The value of the ElastiCache parameter.
	Value string `pulumi:"value"`
}

// ParameterGroupParameterInput is an input type that accepts ParameterGroupParameterArgs and ParameterGroupParameterOutput values.
// You can construct a concrete instance of `ParameterGroupParameterInput` via:
//
// 		 ParameterGroupParameterArgs{...}
//
type ParameterGroupParameterInput interface {
	pulumi.Input

	ToParameterGroupParameterOutput() ParameterGroupParameterOutput
	ToParameterGroupParameterOutputWithContext(context.Context) ParameterGroupParameterOutput
}

type ParameterGroupParameterArgs struct {
	// The name of the ElastiCache parameter group.
	Name pulumi.StringInput `pulumi:"name"`
	// The value of the ElastiCache parameter.
	Value pulumi.StringInput `pulumi:"value"`
}

func (ParameterGroupParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return i.ToParameterGroupParameterOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArgs) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterOutput)
}

// ParameterGroupParameterArrayInput is an input type that accepts ParameterGroupParameterArray and ParameterGroupParameterArrayOutput values.
// You can construct a concrete instance of `ParameterGroupParameterArrayInput` via:
//
// 		 ParameterGroupParameterArray{ ParameterGroupParameterArgs{...} }
//
type ParameterGroupParameterArrayInput interface {
	pulumi.Input

	ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput
	ToParameterGroupParameterArrayOutputWithContext(context.Context) ParameterGroupParameterArrayOutput
}

type ParameterGroupParameterArray []ParameterGroupParameterInput

func (ParameterGroupParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return i.ToParameterGroupParameterArrayOutputWithContext(context.Background())
}

func (i ParameterGroupParameterArray) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterGroupParameterArrayOutput)
}

type ParameterGroupParameterOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutput() ParameterGroupParameterOutput {
	return o
}

func (o ParameterGroupParameterOutput) ToParameterGroupParameterOutputWithContext(ctx context.Context) ParameterGroupParameterOutput {
	return o
}

// The name of the ElastiCache parameter group.
func (o ParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

// The value of the ElastiCache parameter.
func (o ParameterGroupParameterOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Value }).(pulumi.StringOutput)
}

type ParameterGroupParameterArrayOutput struct{ *pulumi.OutputState }

func (ParameterGroupParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ParameterGroupParameter)(nil)).Elem()
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutput() ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) ToParameterGroupParameterArrayOutputWithContext(ctx context.Context) ParameterGroupParameterArrayOutput {
	return o
}

func (o ParameterGroupParameterArrayOutput) Index(i pulumi.IntInput) ParameterGroupParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ParameterGroupParameter {
		return vs[0].([]ParameterGroupParameter)[vs[1].(int)]
	}).(ParameterGroupParameterOutput)
}

type ReplicationGroupClusterMode struct {
	// Specify the number of node groups (shards) for this Redis replication group. Changing this number will trigger an online resizing operation before other settings modifications.
	NumNodeGroups int `pulumi:"numNodeGroups"`
	// Specify the number of replica nodes in each node group. Valid values are 0 to 5. Changing this number will force a new resource.
	ReplicasPerNodeGroup int `pulumi:"replicasPerNodeGroup"`
}

// ReplicationGroupClusterModeInput is an input type that accepts ReplicationGroupClusterModeArgs and ReplicationGroupClusterModeOutput values.
// You can construct a concrete instance of `ReplicationGroupClusterModeInput` via:
//
// 		 ReplicationGroupClusterModeArgs{...}
//
type ReplicationGroupClusterModeInput interface {
	pulumi.Input

	ToReplicationGroupClusterModeOutput() ReplicationGroupClusterModeOutput
	ToReplicationGroupClusterModeOutputWithContext(context.Context) ReplicationGroupClusterModeOutput
}

type ReplicationGroupClusterModeArgs struct {
	// Specify the number of node groups (shards) for this Redis replication group. Changing this number will trigger an online resizing operation before other settings modifications.
	NumNodeGroups pulumi.IntInput `pulumi:"numNodeGroups"`
	// Specify the number of replica nodes in each node group. Valid values are 0 to 5. Changing this number will force a new resource.
	ReplicasPerNodeGroup pulumi.IntInput `pulumi:"replicasPerNodeGroup"`
}

func (ReplicationGroupClusterModeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationGroupClusterMode)(nil)).Elem()
}

func (i ReplicationGroupClusterModeArgs) ToReplicationGroupClusterModeOutput() ReplicationGroupClusterModeOutput {
	return i.ToReplicationGroupClusterModeOutputWithContext(context.Background())
}

func (i ReplicationGroupClusterModeArgs) ToReplicationGroupClusterModeOutputWithContext(ctx context.Context) ReplicationGroupClusterModeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationGroupClusterModeOutput)
}

func (i ReplicationGroupClusterModeArgs) ToReplicationGroupClusterModePtrOutput() ReplicationGroupClusterModePtrOutput {
	return i.ToReplicationGroupClusterModePtrOutputWithContext(context.Background())
}

func (i ReplicationGroupClusterModeArgs) ToReplicationGroupClusterModePtrOutputWithContext(ctx context.Context) ReplicationGroupClusterModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationGroupClusterModeOutput).ToReplicationGroupClusterModePtrOutputWithContext(ctx)
}

// ReplicationGroupClusterModePtrInput is an input type that accepts ReplicationGroupClusterModeArgs, ReplicationGroupClusterModePtr and ReplicationGroupClusterModePtrOutput values.
// You can construct a concrete instance of `ReplicationGroupClusterModePtrInput` via:
//
// 		 ReplicationGroupClusterModeArgs{...}
//
//  or:
//
// 		 nil
//
type ReplicationGroupClusterModePtrInput interface {
	pulumi.Input

	ToReplicationGroupClusterModePtrOutput() ReplicationGroupClusterModePtrOutput
	ToReplicationGroupClusterModePtrOutputWithContext(context.Context) ReplicationGroupClusterModePtrOutput
}

type replicationGroupClusterModePtrType ReplicationGroupClusterModeArgs

func ReplicationGroupClusterModePtr(v *ReplicationGroupClusterModeArgs) ReplicationGroupClusterModePtrInput {
	return (*replicationGroupClusterModePtrType)(v)
}

func (*replicationGroupClusterModePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationGroupClusterMode)(nil)).Elem()
}

func (i *replicationGroupClusterModePtrType) ToReplicationGroupClusterModePtrOutput() ReplicationGroupClusterModePtrOutput {
	return i.ToReplicationGroupClusterModePtrOutputWithContext(context.Background())
}

func (i *replicationGroupClusterModePtrType) ToReplicationGroupClusterModePtrOutputWithContext(ctx context.Context) ReplicationGroupClusterModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationGroupClusterModePtrOutput)
}

type ReplicationGroupClusterModeOutput struct{ *pulumi.OutputState }

func (ReplicationGroupClusterModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationGroupClusterMode)(nil)).Elem()
}

func (o ReplicationGroupClusterModeOutput) ToReplicationGroupClusterModeOutput() ReplicationGroupClusterModeOutput {
	return o
}

func (o ReplicationGroupClusterModeOutput) ToReplicationGroupClusterModeOutputWithContext(ctx context.Context) ReplicationGroupClusterModeOutput {
	return o
}

func (o ReplicationGroupClusterModeOutput) ToReplicationGroupClusterModePtrOutput() ReplicationGroupClusterModePtrOutput {
	return o.ToReplicationGroupClusterModePtrOutputWithContext(context.Background())
}

func (o ReplicationGroupClusterModeOutput) ToReplicationGroupClusterModePtrOutputWithContext(ctx context.Context) ReplicationGroupClusterModePtrOutput {
	return o.ApplyT(func(v ReplicationGroupClusterMode) *ReplicationGroupClusterMode {
		return &v
	}).(ReplicationGroupClusterModePtrOutput)
}

// Specify the number of node groups (shards) for this Redis replication group. Changing this number will trigger an online resizing operation before other settings modifications.
func (o ReplicationGroupClusterModeOutput) NumNodeGroups() pulumi.IntOutput {
	return o.ApplyT(func(v ReplicationGroupClusterMode) int { return v.NumNodeGroups }).(pulumi.IntOutput)
}

// Specify the number of replica nodes in each node group. Valid values are 0 to 5. Changing this number will force a new resource.
func (o ReplicationGroupClusterModeOutput) ReplicasPerNodeGroup() pulumi.IntOutput {
	return o.ApplyT(func(v ReplicationGroupClusterMode) int { return v.ReplicasPerNodeGroup }).(pulumi.IntOutput)
}

type ReplicationGroupClusterModePtrOutput struct{ *pulumi.OutputState }

func (ReplicationGroupClusterModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationGroupClusterMode)(nil)).Elem()
}

func (o ReplicationGroupClusterModePtrOutput) ToReplicationGroupClusterModePtrOutput() ReplicationGroupClusterModePtrOutput {
	return o
}

func (o ReplicationGroupClusterModePtrOutput) ToReplicationGroupClusterModePtrOutputWithContext(ctx context.Context) ReplicationGroupClusterModePtrOutput {
	return o
}

func (o ReplicationGroupClusterModePtrOutput) Elem() ReplicationGroupClusterModeOutput {
	return o.ApplyT(func(v *ReplicationGroupClusterMode) ReplicationGroupClusterMode { return *v }).(ReplicationGroupClusterModeOutput)
}

// Specify the number of node groups (shards) for this Redis replication group. Changing this number will trigger an online resizing operation before other settings modifications.
func (o ReplicationGroupClusterModePtrOutput) NumNodeGroups() pulumi.IntOutput {
	return o.ApplyT(func(v ReplicationGroupClusterMode) int { return v.NumNodeGroups }).(pulumi.IntOutput)
}

// Specify the number of replica nodes in each node group. Valid values are 0 to 5. Changing this number will force a new resource.
func (o ReplicationGroupClusterModePtrOutput) ReplicasPerNodeGroup() pulumi.IntOutput {
	return o.ApplyT(func(v ReplicationGroupClusterMode) int { return v.ReplicasPerNodeGroup }).(pulumi.IntOutput)
}

type GetClusterCacheNode struct {
	Address string `pulumi:"address"`
	// The Availability Zone for the cache cluster.
	AvailabilityZone string `pulumi:"availabilityZone"`
	Id               string `pulumi:"id"`
	// The port number on which each of the cache nodes will
	// accept connections.
	Port int `pulumi:"port"`
}

// GetClusterCacheNodeInput is an input type that accepts GetClusterCacheNodeArgs and GetClusterCacheNodeOutput values.
// You can construct a concrete instance of `GetClusterCacheNodeInput` via:
//
// 		 GetClusterCacheNodeArgs{...}
//
type GetClusterCacheNodeInput interface {
	pulumi.Input

	ToGetClusterCacheNodeOutput() GetClusterCacheNodeOutput
	ToGetClusterCacheNodeOutputWithContext(context.Context) GetClusterCacheNodeOutput
}

type GetClusterCacheNodeArgs struct {
	Address pulumi.StringInput `pulumi:"address"`
	// The Availability Zone for the cache cluster.
	AvailabilityZone pulumi.StringInput `pulumi:"availabilityZone"`
	Id               pulumi.StringInput `pulumi:"id"`
	// The port number on which each of the cache nodes will
	// accept connections.
	Port pulumi.IntInput `pulumi:"port"`
}

func (GetClusterCacheNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCacheNode)(nil)).Elem()
}

func (i GetClusterCacheNodeArgs) ToGetClusterCacheNodeOutput() GetClusterCacheNodeOutput {
	return i.ToGetClusterCacheNodeOutputWithContext(context.Background())
}

func (i GetClusterCacheNodeArgs) ToGetClusterCacheNodeOutputWithContext(ctx context.Context) GetClusterCacheNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterCacheNodeOutput)
}

// GetClusterCacheNodeArrayInput is an input type that accepts GetClusterCacheNodeArray and GetClusterCacheNodeArrayOutput values.
// You can construct a concrete instance of `GetClusterCacheNodeArrayInput` via:
//
// 		 GetClusterCacheNodeArray{ GetClusterCacheNodeArgs{...} }
//
type GetClusterCacheNodeArrayInput interface {
	pulumi.Input

	ToGetClusterCacheNodeArrayOutput() GetClusterCacheNodeArrayOutput
	ToGetClusterCacheNodeArrayOutputWithContext(context.Context) GetClusterCacheNodeArrayOutput
}

type GetClusterCacheNodeArray []GetClusterCacheNodeInput

func (GetClusterCacheNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterCacheNode)(nil)).Elem()
}

func (i GetClusterCacheNodeArray) ToGetClusterCacheNodeArrayOutput() GetClusterCacheNodeArrayOutput {
	return i.ToGetClusterCacheNodeArrayOutputWithContext(context.Background())
}

func (i GetClusterCacheNodeArray) ToGetClusterCacheNodeArrayOutputWithContext(ctx context.Context) GetClusterCacheNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClusterCacheNodeArrayOutput)
}

type GetClusterCacheNodeOutput struct{ *pulumi.OutputState }

func (GetClusterCacheNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClusterCacheNode)(nil)).Elem()
}

func (o GetClusterCacheNodeOutput) ToGetClusterCacheNodeOutput() GetClusterCacheNodeOutput {
	return o
}

func (o GetClusterCacheNodeOutput) ToGetClusterCacheNodeOutputWithContext(ctx context.Context) GetClusterCacheNodeOutput {
	return o
}

func (o GetClusterCacheNodeOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCacheNode) string { return v.Address }).(pulumi.StringOutput)
}

// The Availability Zone for the cache cluster.
func (o GetClusterCacheNodeOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCacheNode) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o GetClusterCacheNodeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClusterCacheNode) string { return v.Id }).(pulumi.StringOutput)
}

// The port number on which each of the cache nodes will
// accept connections.
func (o GetClusterCacheNodeOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v GetClusterCacheNode) int { return v.Port }).(pulumi.IntOutput)
}

type GetClusterCacheNodeArrayOutput struct{ *pulumi.OutputState }

func (GetClusterCacheNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClusterCacheNode)(nil)).Elem()
}

func (o GetClusterCacheNodeArrayOutput) ToGetClusterCacheNodeArrayOutput() GetClusterCacheNodeArrayOutput {
	return o
}

func (o GetClusterCacheNodeArrayOutput) ToGetClusterCacheNodeArrayOutputWithContext(ctx context.Context) GetClusterCacheNodeArrayOutput {
	return o
}

func (o GetClusterCacheNodeArrayOutput) Index(i pulumi.IntInput) GetClusterCacheNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetClusterCacheNode {
		return vs[0].([]GetClusterCacheNode)[vs[1].(int)]
	}).(GetClusterCacheNodeOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterCacheNodeOutput{})
	pulumi.RegisterOutputType(ClusterCacheNodeArrayOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArrayOutput{})
	pulumi.RegisterOutputType(ReplicationGroupClusterModeOutput{})
	pulumi.RegisterOutputType(ReplicationGroupClusterModePtrOutput{})
	pulumi.RegisterOutputType(GetClusterCacheNodeOutput{})
	pulumi.RegisterOutputType(GetClusterCacheNodeArrayOutput{})
}
