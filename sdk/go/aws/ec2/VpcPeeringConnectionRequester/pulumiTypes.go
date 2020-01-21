// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package VpcPeeringConnectionRequester

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type VpcPeeringConnectionRequester struct {
	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC.
	AllowClassicLinkToRemoteVpc *bool `pulumi:"allowClassicLinkToRemoteVpc"`
	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC. This is
	// [not supported](https://docs.aws.amazon.com/vpc/latest/peering/modify-peering-connections.html) for
	// inter-region VPC peering.
	AllowRemoteVpcDnsResolution *bool `pulumi:"allowRemoteVpcDnsResolution"`
	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection.
	AllowVpcToRemoteClassicLink *bool `pulumi:"allowVpcToRemoteClassicLink"`
}

type VpcPeeringConnectionRequesterInput interface {
	pulumi.Input

	ToVpcPeeringConnectionRequesterOutput() VpcPeeringConnectionRequesterOutput
	ToVpcPeeringConnectionRequesterOutputWithContext(context.Context) VpcPeeringConnectionRequesterOutput
}

type VpcPeeringConnectionRequesterArgs struct {
	// Allow a local linked EC2-Classic instance to communicate
	// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
	// to the remote VPC.
	AllowClassicLinkToRemoteVpc pulumi.BoolPtrInput `pulumi:"allowClassicLinkToRemoteVpc"`
	// Allow a local VPC to resolve public DNS hostnames to
	// private IP addresses when queried from instances in the peer VPC. This is
	// [not supported](https://docs.aws.amazon.com/vpc/latest/peering/modify-peering-connections.html) for
	// inter-region VPC peering.
	AllowRemoteVpcDnsResolution pulumi.BoolPtrInput `pulumi:"allowRemoteVpcDnsResolution"`
	// Allow a local VPC to communicate with a linked EC2-Classic
	// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
	// connection.
	AllowVpcToRemoteClassicLink pulumi.BoolPtrInput `pulumi:"allowVpcToRemoteClassicLink"`
}

func (VpcPeeringConnectionRequesterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcPeeringConnectionRequester)(nil)).Elem()
}

func (i VpcPeeringConnectionRequesterArgs) ToVpcPeeringConnectionRequesterOutput() VpcPeeringConnectionRequesterOutput {
	return i.ToVpcPeeringConnectionRequesterOutputWithContext(context.Background())
}

func (i VpcPeeringConnectionRequesterArgs) ToVpcPeeringConnectionRequesterOutputWithContext(ctx context.Context) VpcPeeringConnectionRequesterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringConnectionRequesterOutput)
}

func (i VpcPeeringConnectionRequesterArgs) ToVpcPeeringConnectionRequesterPtrOutput() VpcPeeringConnectionRequesterPtrOutput {
	return i.ToVpcPeeringConnectionRequesterPtrOutputWithContext(context.Background())
}

func (i VpcPeeringConnectionRequesterArgs) ToVpcPeeringConnectionRequesterPtrOutputWithContext(ctx context.Context) VpcPeeringConnectionRequesterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringConnectionRequesterOutput).ToVpcPeeringConnectionRequesterPtrOutputWithContext(ctx)
}

type VpcPeeringConnectionRequesterPtrInput interface {
	pulumi.Input

	ToVpcPeeringConnectionRequesterPtrOutput() VpcPeeringConnectionRequesterPtrOutput
	ToVpcPeeringConnectionRequesterPtrOutputWithContext(context.Context) VpcPeeringConnectionRequesterPtrOutput
}

type vpcPeeringConnectionRequesterPtrType VpcPeeringConnectionRequesterArgs

func VpcPeeringConnectionRequesterPtr(v *VpcPeeringConnectionRequesterArgs) VpcPeeringConnectionRequesterPtrInput {	return (*vpcPeeringConnectionRequesterPtrType)(v)
}

func (*vpcPeeringConnectionRequesterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPeeringConnectionRequester)(nil)).Elem()
}

func (i *vpcPeeringConnectionRequesterPtrType) ToVpcPeeringConnectionRequesterPtrOutput() VpcPeeringConnectionRequesterPtrOutput {
	return i.ToVpcPeeringConnectionRequesterPtrOutputWithContext(context.Background())
}

func (i *vpcPeeringConnectionRequesterPtrType) ToVpcPeeringConnectionRequesterPtrOutputWithContext(ctx context.Context) VpcPeeringConnectionRequesterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcPeeringConnectionRequesterPtrOutput)
}

type VpcPeeringConnectionRequesterOutput struct { *pulumi.OutputState }

func (VpcPeeringConnectionRequesterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcPeeringConnectionRequester)(nil)).Elem()
}

func (o VpcPeeringConnectionRequesterOutput) ToVpcPeeringConnectionRequesterOutput() VpcPeeringConnectionRequesterOutput {
	return o
}

func (o VpcPeeringConnectionRequesterOutput) ToVpcPeeringConnectionRequesterOutputWithContext(ctx context.Context) VpcPeeringConnectionRequesterOutput {
	return o
}

func (o VpcPeeringConnectionRequesterOutput) ToVpcPeeringConnectionRequesterPtrOutput() VpcPeeringConnectionRequesterPtrOutput {
	return o.ToVpcPeeringConnectionRequesterPtrOutputWithContext(context.Background())
}

func (o VpcPeeringConnectionRequesterOutput) ToVpcPeeringConnectionRequesterPtrOutputWithContext(ctx context.Context) VpcPeeringConnectionRequesterPtrOutput {
	return o.ApplyT(func(v VpcPeeringConnectionRequester) *VpcPeeringConnectionRequester {
		return &v
	}).(VpcPeeringConnectionRequesterPtrOutput)
}
// Allow a local linked EC2-Classic instance to communicate
// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
// to the remote VPC.
func (o VpcPeeringConnectionRequesterOutput) AllowClassicLinkToRemoteVpc() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v VpcPeeringConnectionRequester) *bool { return v.AllowClassicLinkToRemoteVpc }).(pulumi.BoolPtrOutput)
}

// Allow a local VPC to resolve public DNS hostnames to
// private IP addresses when queried from instances in the peer VPC. This is
// [not supported](https://docs.aws.amazon.com/vpc/latest/peering/modify-peering-connections.html) for
// inter-region VPC peering.
func (o VpcPeeringConnectionRequesterOutput) AllowRemoteVpcDnsResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v VpcPeeringConnectionRequester) *bool { return v.AllowRemoteVpcDnsResolution }).(pulumi.BoolPtrOutput)
}

// Allow a local VPC to communicate with a linked EC2-Classic
// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
// connection.
func (o VpcPeeringConnectionRequesterOutput) AllowVpcToRemoteClassicLink() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v VpcPeeringConnectionRequester) *bool { return v.AllowVpcToRemoteClassicLink }).(pulumi.BoolPtrOutput)
}

type VpcPeeringConnectionRequesterPtrOutput struct { *pulumi.OutputState}

func (VpcPeeringConnectionRequesterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcPeeringConnectionRequester)(nil)).Elem()
}

func (o VpcPeeringConnectionRequesterPtrOutput) ToVpcPeeringConnectionRequesterPtrOutput() VpcPeeringConnectionRequesterPtrOutput {
	return o
}

func (o VpcPeeringConnectionRequesterPtrOutput) ToVpcPeeringConnectionRequesterPtrOutputWithContext(ctx context.Context) VpcPeeringConnectionRequesterPtrOutput {
	return o
}

func (o VpcPeeringConnectionRequesterPtrOutput) Elem() VpcPeeringConnectionRequesterOutput {
	return o.ApplyT(func (v *VpcPeeringConnectionRequester) VpcPeeringConnectionRequester { return *v }).(VpcPeeringConnectionRequesterOutput)
}

// Allow a local linked EC2-Classic instance to communicate
// with instances in a peer VPC. This enables an outbound communication from the local ClassicLink connection
// to the remote VPC.
func (o VpcPeeringConnectionRequesterPtrOutput) AllowClassicLinkToRemoteVpc() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v VpcPeeringConnectionRequester) *bool { return v.AllowClassicLinkToRemoteVpc }).(pulumi.BoolPtrOutput)
}

// Allow a local VPC to resolve public DNS hostnames to
// private IP addresses when queried from instances in the peer VPC. This is
// [not supported](https://docs.aws.amazon.com/vpc/latest/peering/modify-peering-connections.html) for
// inter-region VPC peering.
func (o VpcPeeringConnectionRequesterPtrOutput) AllowRemoteVpcDnsResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v VpcPeeringConnectionRequester) *bool { return v.AllowRemoteVpcDnsResolution }).(pulumi.BoolPtrOutput)
}

// Allow a local VPC to communicate with a linked EC2-Classic
// instance in a peer VPC. This enables an outbound communication from the local VPC to the remote ClassicLink
// connection.
func (o VpcPeeringConnectionRequesterPtrOutput) AllowVpcToRemoteClassicLink() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v VpcPeeringConnectionRequester) *bool { return v.AllowVpcToRemoteClassicLink }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcPeeringConnectionRequesterOutput{})
	pulumi.RegisterOutputType(VpcPeeringConnectionRequesterPtrOutput{})
}
