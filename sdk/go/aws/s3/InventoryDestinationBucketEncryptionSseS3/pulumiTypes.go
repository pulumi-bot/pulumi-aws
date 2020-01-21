// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package InventoryDestinationBucketEncryptionSseS3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type InventoryDestinationBucketEncryptionSseS3 struct {
}

type InventoryDestinationBucketEncryptionSseS3Input interface {
	pulumi.Input

	ToInventoryDestinationBucketEncryptionSseS3Output() InventoryDestinationBucketEncryptionSseS3Output
	ToInventoryDestinationBucketEncryptionSseS3OutputWithContext(context.Context) InventoryDestinationBucketEncryptionSseS3Output
}

type InventoryDestinationBucketEncryptionSseS3Args struct {
}

func (InventoryDestinationBucketEncryptionSseS3Args) ElementType() reflect.Type {
	return reflect.TypeOf((*InventoryDestinationBucketEncryptionSseS3)(nil)).Elem()
}

func (i InventoryDestinationBucketEncryptionSseS3Args) ToInventoryDestinationBucketEncryptionSseS3Output() InventoryDestinationBucketEncryptionSseS3Output {
	return i.ToInventoryDestinationBucketEncryptionSseS3OutputWithContext(context.Background())
}

func (i InventoryDestinationBucketEncryptionSseS3Args) ToInventoryDestinationBucketEncryptionSseS3OutputWithContext(ctx context.Context) InventoryDestinationBucketEncryptionSseS3Output {
	return pulumi.ToOutputWithContext(ctx, i).(InventoryDestinationBucketEncryptionSseS3Output)
}

func (i InventoryDestinationBucketEncryptionSseS3Args) ToInventoryDestinationBucketEncryptionSseS3PtrOutput() InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return i.ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(context.Background())
}

func (i InventoryDestinationBucketEncryptionSseS3Args) ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(ctx context.Context) InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InventoryDestinationBucketEncryptionSseS3Output).ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(ctx)
}

type InventoryDestinationBucketEncryptionSseS3PtrInput interface {
	pulumi.Input

	ToInventoryDestinationBucketEncryptionSseS3PtrOutput() InventoryDestinationBucketEncryptionSseS3PtrOutput
	ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(context.Context) InventoryDestinationBucketEncryptionSseS3PtrOutput
}

type inventoryDestinationBucketEncryptionSseS3PtrType InventoryDestinationBucketEncryptionSseS3Args

func InventoryDestinationBucketEncryptionSseS3Ptr(v *InventoryDestinationBucketEncryptionSseS3Args) InventoryDestinationBucketEncryptionSseS3PtrInput {	return (*inventoryDestinationBucketEncryptionSseS3PtrType)(v)
}

func (*inventoryDestinationBucketEncryptionSseS3PtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryDestinationBucketEncryptionSseS3)(nil)).Elem()
}

func (i *inventoryDestinationBucketEncryptionSseS3PtrType) ToInventoryDestinationBucketEncryptionSseS3PtrOutput() InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return i.ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(context.Background())
}

func (i *inventoryDestinationBucketEncryptionSseS3PtrType) ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(ctx context.Context) InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InventoryDestinationBucketEncryptionSseS3PtrOutput)
}

type InventoryDestinationBucketEncryptionSseS3Output struct { *pulumi.OutputState }

func (InventoryDestinationBucketEncryptionSseS3Output) ElementType() reflect.Type {
	return reflect.TypeOf((*InventoryDestinationBucketEncryptionSseS3)(nil)).Elem()
}

func (o InventoryDestinationBucketEncryptionSseS3Output) ToInventoryDestinationBucketEncryptionSseS3Output() InventoryDestinationBucketEncryptionSseS3Output {
	return o
}

func (o InventoryDestinationBucketEncryptionSseS3Output) ToInventoryDestinationBucketEncryptionSseS3OutputWithContext(ctx context.Context) InventoryDestinationBucketEncryptionSseS3Output {
	return o
}

func (o InventoryDestinationBucketEncryptionSseS3Output) ToInventoryDestinationBucketEncryptionSseS3PtrOutput() InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return o.ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(context.Background())
}

func (o InventoryDestinationBucketEncryptionSseS3Output) ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(ctx context.Context) InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return o.ApplyT(func(v InventoryDestinationBucketEncryptionSseS3) *InventoryDestinationBucketEncryptionSseS3 {
		return &v
	}).(InventoryDestinationBucketEncryptionSseS3PtrOutput)
}
type InventoryDestinationBucketEncryptionSseS3PtrOutput struct { *pulumi.OutputState}

func (InventoryDestinationBucketEncryptionSseS3PtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InventoryDestinationBucketEncryptionSseS3)(nil)).Elem()
}

func (o InventoryDestinationBucketEncryptionSseS3PtrOutput) ToInventoryDestinationBucketEncryptionSseS3PtrOutput() InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return o
}

func (o InventoryDestinationBucketEncryptionSseS3PtrOutput) ToInventoryDestinationBucketEncryptionSseS3PtrOutputWithContext(ctx context.Context) InventoryDestinationBucketEncryptionSseS3PtrOutput {
	return o
}

func (o InventoryDestinationBucketEncryptionSseS3PtrOutput) Elem() InventoryDestinationBucketEncryptionSseS3Output {
	return o.ApplyT(func (v *InventoryDestinationBucketEncryptionSseS3) InventoryDestinationBucketEncryptionSseS3 { return *v }).(InventoryDestinationBucketEncryptionSseS3Output)
}

func init() {
	pulumi.RegisterOutputType(InventoryDestinationBucketEncryptionSseS3Output{})
	pulumi.RegisterOutputType(InventoryDestinationBucketEncryptionSseS3PtrOutput{})
}
