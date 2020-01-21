// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterClientAuthentication

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/msk/ClusterClientAuthenticationTls"
)

type ClusterClientAuthentication struct {
	// Configuration block for specifying TLS client authentication. See below.
	Tls *mskClusterClientAuthenticationTls.ClusterClientAuthenticationTls `pulumi:"tls"`
}

type ClusterClientAuthenticationInput interface {
	pulumi.Input

	ToClusterClientAuthenticationOutput() ClusterClientAuthenticationOutput
	ToClusterClientAuthenticationOutputWithContext(context.Context) ClusterClientAuthenticationOutput
}

type ClusterClientAuthenticationArgs struct {
	// Configuration block for specifying TLS client authentication. See below.
	Tls mskClusterClientAuthenticationTls.ClusterClientAuthenticationTlsPtrInput `pulumi:"tls"`
}

func (ClusterClientAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClientAuthentication)(nil)).Elem()
}

func (i ClusterClientAuthenticationArgs) ToClusterClientAuthenticationOutput() ClusterClientAuthenticationOutput {
	return i.ToClusterClientAuthenticationOutputWithContext(context.Background())
}

func (i ClusterClientAuthenticationArgs) ToClusterClientAuthenticationOutputWithContext(ctx context.Context) ClusterClientAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClientAuthenticationOutput)
}

func (i ClusterClientAuthenticationArgs) ToClusterClientAuthenticationPtrOutput() ClusterClientAuthenticationPtrOutput {
	return i.ToClusterClientAuthenticationPtrOutputWithContext(context.Background())
}

func (i ClusterClientAuthenticationArgs) ToClusterClientAuthenticationPtrOutputWithContext(ctx context.Context) ClusterClientAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClientAuthenticationOutput).ToClusterClientAuthenticationPtrOutputWithContext(ctx)
}

type ClusterClientAuthenticationPtrInput interface {
	pulumi.Input

	ToClusterClientAuthenticationPtrOutput() ClusterClientAuthenticationPtrOutput
	ToClusterClientAuthenticationPtrOutputWithContext(context.Context) ClusterClientAuthenticationPtrOutput
}

type clusterClientAuthenticationPtrType ClusterClientAuthenticationArgs

func ClusterClientAuthenticationPtr(v *ClusterClientAuthenticationArgs) ClusterClientAuthenticationPtrInput {	return (*clusterClientAuthenticationPtrType)(v)
}

func (*clusterClientAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClientAuthentication)(nil)).Elem()
}

func (i *clusterClientAuthenticationPtrType) ToClusterClientAuthenticationPtrOutput() ClusterClientAuthenticationPtrOutput {
	return i.ToClusterClientAuthenticationPtrOutputWithContext(context.Background())
}

func (i *clusterClientAuthenticationPtrType) ToClusterClientAuthenticationPtrOutputWithContext(ctx context.Context) ClusterClientAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterClientAuthenticationPtrOutput)
}

type ClusterClientAuthenticationOutput struct { *pulumi.OutputState }

func (ClusterClientAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterClientAuthentication)(nil)).Elem()
}

func (o ClusterClientAuthenticationOutput) ToClusterClientAuthenticationOutput() ClusterClientAuthenticationOutput {
	return o
}

func (o ClusterClientAuthenticationOutput) ToClusterClientAuthenticationOutputWithContext(ctx context.Context) ClusterClientAuthenticationOutput {
	return o
}

func (o ClusterClientAuthenticationOutput) ToClusterClientAuthenticationPtrOutput() ClusterClientAuthenticationPtrOutput {
	return o.ToClusterClientAuthenticationPtrOutputWithContext(context.Background())
}

func (o ClusterClientAuthenticationOutput) ToClusterClientAuthenticationPtrOutputWithContext(ctx context.Context) ClusterClientAuthenticationPtrOutput {
	return o.ApplyT(func(v ClusterClientAuthentication) *ClusterClientAuthentication {
		return &v
	}).(ClusterClientAuthenticationPtrOutput)
}
// Configuration block for specifying TLS client authentication. See below.
func (o ClusterClientAuthenticationOutput) Tls() mskClusterClientAuthenticationTls.ClusterClientAuthenticationTlsPtrOutput {
	return o.ApplyT(func (v ClusterClientAuthentication) *mskClusterClientAuthenticationTls.ClusterClientAuthenticationTls { return v.Tls }).(mskClusterClientAuthenticationTls.ClusterClientAuthenticationTlsPtrOutput)
}

type ClusterClientAuthenticationPtrOutput struct { *pulumi.OutputState}

func (ClusterClientAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterClientAuthentication)(nil)).Elem()
}

func (o ClusterClientAuthenticationPtrOutput) ToClusterClientAuthenticationPtrOutput() ClusterClientAuthenticationPtrOutput {
	return o
}

func (o ClusterClientAuthenticationPtrOutput) ToClusterClientAuthenticationPtrOutputWithContext(ctx context.Context) ClusterClientAuthenticationPtrOutput {
	return o
}

func (o ClusterClientAuthenticationPtrOutput) Elem() ClusterClientAuthenticationOutput {
	return o.ApplyT(func (v *ClusterClientAuthentication) ClusterClientAuthentication { return *v }).(ClusterClientAuthenticationOutput)
}

// Configuration block for specifying TLS client authentication. See below.
func (o ClusterClientAuthenticationPtrOutput) Tls() mskClusterClientAuthenticationTls.ClusterClientAuthenticationTlsPtrOutput {
	return o.ApplyT(func (v ClusterClientAuthentication) *mskClusterClientAuthenticationTls.ClusterClientAuthenticationTls { return v.Tls }).(mskClusterClientAuthenticationTls.ClusterClientAuthenticationTlsPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterClientAuthenticationOutput{})
	pulumi.RegisterOutputType(ClusterClientAuthenticationPtrOutput{})
}
