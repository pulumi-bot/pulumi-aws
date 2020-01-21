// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package NodeGroupResource

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/eks/NodeGroupResourceAutoscalingGroup"
)

type NodeGroupResource struct {
	// List of objects containing information about AutoScaling Groups.
	AutoscalingGroups []eksNodeGroupResourceAutoscalingGroup.NodeGroupResourceAutoscalingGroup `pulumi:"autoscalingGroups"`
	// Identifier of the remote access EC2 Security Group.
	RemoteAccessSecurityGroupId *string `pulumi:"remoteAccessSecurityGroupId"`
}

type NodeGroupResourceInput interface {
	pulumi.Input

	ToNodeGroupResourceOutput() NodeGroupResourceOutput
	ToNodeGroupResourceOutputWithContext(context.Context) NodeGroupResourceOutput
}

type NodeGroupResourceArgs struct {
	// List of objects containing information about AutoScaling Groups.
	AutoscalingGroups eksNodeGroupResourceAutoscalingGroup.NodeGroupResourceAutoscalingGroupArrayInput `pulumi:"autoscalingGroups"`
	// Identifier of the remote access EC2 Security Group.
	RemoteAccessSecurityGroupId pulumi.StringPtrInput `pulumi:"remoteAccessSecurityGroupId"`
}

func (NodeGroupResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeGroupResource)(nil)).Elem()
}

func (i NodeGroupResourceArgs) ToNodeGroupResourceOutput() NodeGroupResourceOutput {
	return i.ToNodeGroupResourceOutputWithContext(context.Background())
}

func (i NodeGroupResourceArgs) ToNodeGroupResourceOutputWithContext(ctx context.Context) NodeGroupResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupResourceOutput)
}

type NodeGroupResourceArrayInput interface {
	pulumi.Input

	ToNodeGroupResourceArrayOutput() NodeGroupResourceArrayOutput
	ToNodeGroupResourceArrayOutputWithContext(context.Context) NodeGroupResourceArrayOutput
}

type NodeGroupResourceArray []NodeGroupResourceInput

func (NodeGroupResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeGroupResource)(nil)).Elem()
}

func (i NodeGroupResourceArray) ToNodeGroupResourceArrayOutput() NodeGroupResourceArrayOutput {
	return i.ToNodeGroupResourceArrayOutputWithContext(context.Background())
}

func (i NodeGroupResourceArray) ToNodeGroupResourceArrayOutputWithContext(ctx context.Context) NodeGroupResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeGroupResourceArrayOutput)
}

type NodeGroupResourceOutput struct { *pulumi.OutputState }

func (NodeGroupResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodeGroupResource)(nil)).Elem()
}

func (o NodeGroupResourceOutput) ToNodeGroupResourceOutput() NodeGroupResourceOutput {
	return o
}

func (o NodeGroupResourceOutput) ToNodeGroupResourceOutputWithContext(ctx context.Context) NodeGroupResourceOutput {
	return o
}

// List of objects containing information about AutoScaling Groups.
func (o NodeGroupResourceOutput) AutoscalingGroups() eksNodeGroupResourceAutoscalingGroup.NodeGroupResourceAutoscalingGroupArrayOutput {
	return o.ApplyT(func (v NodeGroupResource) []eksNodeGroupResourceAutoscalingGroup.NodeGroupResourceAutoscalingGroup { return v.AutoscalingGroups }).(eksNodeGroupResourceAutoscalingGroup.NodeGroupResourceAutoscalingGroupArrayOutput)
}

// Identifier of the remote access EC2 Security Group.
func (o NodeGroupResourceOutput) RemoteAccessSecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func (v NodeGroupResource) *string { return v.RemoteAccessSecurityGroupId }).(pulumi.StringPtrOutput)
}

type NodeGroupResourceArrayOutput struct { *pulumi.OutputState}

func (NodeGroupResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NodeGroupResource)(nil)).Elem()
}

func (o NodeGroupResourceArrayOutput) ToNodeGroupResourceArrayOutput() NodeGroupResourceArrayOutput {
	return o
}

func (o NodeGroupResourceArrayOutput) ToNodeGroupResourceArrayOutputWithContext(ctx context.Context) NodeGroupResourceArrayOutput {
	return o
}

func (o NodeGroupResourceArrayOutput) Index(i pulumi.IntInput) NodeGroupResourceOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) NodeGroupResource {
		return vs[0].([]NodeGroupResource)[vs[1].(int)]
	}).(NodeGroupResourceOutput)
}

func init() {
	pulumi.RegisterOutputType(NodeGroupResourceOutput{})
	pulumi.RegisterOutputType(NodeGroupResourceArrayOutput{})
}
