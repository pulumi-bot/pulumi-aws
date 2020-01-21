// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterCacheNode

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterCacheNode struct {
	Address *string `pulumi:"address"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	Id *string `pulumi:"id"`
	// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
	Port *int `pulumi:"port"`
}

type ClusterCacheNodeInput interface {
	pulumi.Input

	ToClusterCacheNodeOutput() ClusterCacheNodeOutput
	ToClusterCacheNodeOutputWithContext(context.Context) ClusterCacheNodeOutput
}

type ClusterCacheNodeArgs struct {
	Address pulumi.StringPtrInput `pulumi:"address"`
	// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	Id pulumi.StringPtrInput `pulumi:"id"`
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

type ClusterCacheNodeOutput struct { *pulumi.OutputState }

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
	return o.ApplyT(func (v ClusterCacheNode) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferredAvailabilityZones` instead. Default: System chosen Availability Zone.
func (o ClusterCacheNodeOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterCacheNode) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o ClusterCacheNodeOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterCacheNode) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replicationGroupId`.
func (o ClusterCacheNodeOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func (v ClusterCacheNode) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type ClusterCacheNodeArrayOutput struct { *pulumi.OutputState}

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
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ClusterCacheNode {
		return vs[0].([]ClusterCacheNode)[vs[1].(int)]
	}).(ClusterCacheNodeOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterCacheNodeOutput{})
	pulumi.RegisterOutputType(ClusterCacheNodeArrayOutput{})
}
