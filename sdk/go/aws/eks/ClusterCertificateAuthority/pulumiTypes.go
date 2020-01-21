// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ClusterCertificateAuthority

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ClusterCertificateAuthority struct {
	// The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
	Data *string `pulumi:"data"`
}

type ClusterCertificateAuthorityInput interface {
	pulumi.Input

	ToClusterCertificateAuthorityOutput() ClusterCertificateAuthorityOutput
	ToClusterCertificateAuthorityOutputWithContext(context.Context) ClusterCertificateAuthorityOutput
}

type ClusterCertificateAuthorityArgs struct {
	// The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
	Data pulumi.StringPtrInput `pulumi:"data"`
}

func (ClusterCertificateAuthorityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCertificateAuthority)(nil)).Elem()
}

func (i ClusterCertificateAuthorityArgs) ToClusterCertificateAuthorityOutput() ClusterCertificateAuthorityOutput {
	return i.ToClusterCertificateAuthorityOutputWithContext(context.Background())
}

func (i ClusterCertificateAuthorityArgs) ToClusterCertificateAuthorityOutputWithContext(ctx context.Context) ClusterCertificateAuthorityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCertificateAuthorityOutput)
}

func (i ClusterCertificateAuthorityArgs) ToClusterCertificateAuthorityPtrOutput() ClusterCertificateAuthorityPtrOutput {
	return i.ToClusterCertificateAuthorityPtrOutputWithContext(context.Background())
}

func (i ClusterCertificateAuthorityArgs) ToClusterCertificateAuthorityPtrOutputWithContext(ctx context.Context) ClusterCertificateAuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCertificateAuthorityOutput).ToClusterCertificateAuthorityPtrOutputWithContext(ctx)
}

type ClusterCertificateAuthorityPtrInput interface {
	pulumi.Input

	ToClusterCertificateAuthorityPtrOutput() ClusterCertificateAuthorityPtrOutput
	ToClusterCertificateAuthorityPtrOutputWithContext(context.Context) ClusterCertificateAuthorityPtrOutput
}

type clusterCertificateAuthorityPtrType ClusterCertificateAuthorityArgs

func ClusterCertificateAuthorityPtr(v *ClusterCertificateAuthorityArgs) ClusterCertificateAuthorityPtrInput {	return (*clusterCertificateAuthorityPtrType)(v)
}

func (*clusterCertificateAuthorityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCertificateAuthority)(nil)).Elem()
}

func (i *clusterCertificateAuthorityPtrType) ToClusterCertificateAuthorityPtrOutput() ClusterCertificateAuthorityPtrOutput {
	return i.ToClusterCertificateAuthorityPtrOutputWithContext(context.Background())
}

func (i *clusterCertificateAuthorityPtrType) ToClusterCertificateAuthorityPtrOutputWithContext(ctx context.Context) ClusterCertificateAuthorityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterCertificateAuthorityPtrOutput)
}

type ClusterCertificateAuthorityOutput struct { *pulumi.OutputState }

func (ClusterCertificateAuthorityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterCertificateAuthority)(nil)).Elem()
}

func (o ClusterCertificateAuthorityOutput) ToClusterCertificateAuthorityOutput() ClusterCertificateAuthorityOutput {
	return o
}

func (o ClusterCertificateAuthorityOutput) ToClusterCertificateAuthorityOutputWithContext(ctx context.Context) ClusterCertificateAuthorityOutput {
	return o
}

func (o ClusterCertificateAuthorityOutput) ToClusterCertificateAuthorityPtrOutput() ClusterCertificateAuthorityPtrOutput {
	return o.ToClusterCertificateAuthorityPtrOutputWithContext(context.Background())
}

func (o ClusterCertificateAuthorityOutput) ToClusterCertificateAuthorityPtrOutputWithContext(ctx context.Context) ClusterCertificateAuthorityPtrOutput {
	return o.ApplyT(func(v ClusterCertificateAuthority) *ClusterCertificateAuthority {
		return &v
	}).(ClusterCertificateAuthorityPtrOutput)
}
// The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
func (o ClusterCertificateAuthorityOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterCertificateAuthority) *string { return v.Data }).(pulumi.StringPtrOutput)
}

type ClusterCertificateAuthorityPtrOutput struct { *pulumi.OutputState}

func (ClusterCertificateAuthorityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterCertificateAuthority)(nil)).Elem()
}

func (o ClusterCertificateAuthorityPtrOutput) ToClusterCertificateAuthorityPtrOutput() ClusterCertificateAuthorityPtrOutput {
	return o
}

func (o ClusterCertificateAuthorityPtrOutput) ToClusterCertificateAuthorityPtrOutputWithContext(ctx context.Context) ClusterCertificateAuthorityPtrOutput {
	return o
}

func (o ClusterCertificateAuthorityPtrOutput) Elem() ClusterCertificateAuthorityOutput {
	return o.ApplyT(func (v *ClusterCertificateAuthority) ClusterCertificateAuthority { return *v }).(ClusterCertificateAuthorityOutput)
}

// The base64 encoded certificate data required to communicate with your cluster. Add this to the `certificate-authority-data` section of the `kubeconfig` file for your cluster.
func (o ClusterCertificateAuthorityPtrOutput) Data() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ClusterCertificateAuthority) *string { return v.Data }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterCertificateAuthorityOutput{})
	pulumi.RegisterOutputType(ClusterCertificateAuthorityPtrOutput{})
}
