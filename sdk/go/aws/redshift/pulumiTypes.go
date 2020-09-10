// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClusterLogging struct {
	BucketName  *string `pulumi:"bucketName"`
	Enable      bool    `pulumi:"enable"`
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
}

// ClusterLoggingInput is an input type that accepts ClusterLoggingArgs and ClusterLoggingOutput values.
// You can construct a concrete instance of `ClusterLoggingInput` via:
//
//          ClusterLoggingArgs{...}
type ClusterLoggingInput interface {
	pulumi.Input

	ToClusterLoggingOutput() ClusterLoggingOutput
	ToClusterLoggingOutputWithContext(context.Context) ClusterLoggingOutput
}

type ClusterLoggingArgs struct {
	BucketName  pulumi.StringPtrInput `pulumi:"bucketName"`
	Enable      pulumi.BoolInput      `pulumi:"enable"`
	S3KeyPrefix pulumi.StringPtrInput `pulumi:"s3KeyPrefix"`
}

func (ClusterLoggingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLogging)(nil)).Elem()
}

func (i ClusterLoggingArgs) ToClusterLoggingOutput() ClusterLoggingOutput {
	return i.ToClusterLoggingOutputWithContext(context.Background())
}

func (i ClusterLoggingArgs) ToClusterLoggingOutputWithContext(ctx context.Context) ClusterLoggingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingOutput)
}

func (i ClusterLoggingArgs) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return i.ToClusterLoggingPtrOutputWithContext(context.Background())
}

func (i ClusterLoggingArgs) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingOutput).ToClusterLoggingPtrOutputWithContext(ctx)
}

// ClusterLoggingPtrInput is an input type that accepts ClusterLoggingArgs, ClusterLoggingPtr and ClusterLoggingPtrOutput values.
// You can construct a concrete instance of `ClusterLoggingPtrInput` via:
//
//          ClusterLoggingArgs{...}
//
//  or:
//
//          nil
type ClusterLoggingPtrInput interface {
	pulumi.Input

	ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput
	ToClusterLoggingPtrOutputWithContext(context.Context) ClusterLoggingPtrOutput
}

type clusterLoggingPtrType ClusterLoggingArgs

func ClusterLoggingPtr(v *ClusterLoggingArgs) ClusterLoggingPtrInput {
	return (*clusterLoggingPtrType)(v)
}

func (*clusterLoggingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterLogging)(nil)).Elem()
}

func (i *clusterLoggingPtrType) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return i.ToClusterLoggingPtrOutputWithContext(context.Background())
}

func (i *clusterLoggingPtrType) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingPtrOutput)
}

type ClusterLoggingOutput struct{ *pulumi.OutputState }

func (ClusterLoggingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLogging)(nil)).Elem()
}

func (o ClusterLoggingOutput) ToClusterLoggingOutput() ClusterLoggingOutput {
	return o
}

func (o ClusterLoggingOutput) ToClusterLoggingOutputWithContext(ctx context.Context) ClusterLoggingOutput {
	return o
}

func (o ClusterLoggingOutput) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return o.ToClusterLoggingPtrOutputWithContext(context.Background())
}

func (o ClusterLoggingOutput) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return o.ApplyT(func(v ClusterLogging) *ClusterLogging {
		return &v
	}).(ClusterLoggingPtrOutput)
}
func (o ClusterLoggingOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterLogging) *string { return v.BucketName }).(pulumi.StringPtrOutput)
}

func (o ClusterLoggingOutput) Enable() pulumi.BoolOutput {
	return o.ApplyT(func(v ClusterLogging) bool { return v.Enable }).(pulumi.BoolOutput)
}

func (o ClusterLoggingOutput) S3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterLogging) *string { return v.S3KeyPrefix }).(pulumi.StringPtrOutput)
}

type ClusterLoggingPtrOutput struct{ *pulumi.OutputState }

func (ClusterLoggingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterLogging)(nil)).Elem()
}

func (o ClusterLoggingPtrOutput) ToClusterLoggingPtrOutput() ClusterLoggingPtrOutput {
	return o
}

func (o ClusterLoggingPtrOutput) ToClusterLoggingPtrOutputWithContext(ctx context.Context) ClusterLoggingPtrOutput {
	return o
}

func (o ClusterLoggingPtrOutput) Elem() ClusterLoggingOutput {
	return o.ApplyT(func(v *ClusterLogging) ClusterLogging { return *v }).(ClusterLoggingOutput)
}

func (o ClusterLoggingPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterLogging) *string {
		if v == nil {
			return nil
		}
		return v.BucketName
	}).(pulumi.StringPtrOutput)
}

func (o ClusterLoggingPtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterLogging) *bool {
		if v == nil {
			return nil
		}
		return &v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterLoggingPtrOutput) S3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterLogging) *string {
		if v == nil {
			return nil
		}
		return v.S3KeyPrefix
	}).(pulumi.StringPtrOutput)
}

type ClusterSnapshotCopy struct {
	DestinationRegion string  `pulumi:"destinationRegion"`
	GrantName         *string `pulumi:"grantName"`
	RetentionPeriod   *int    `pulumi:"retentionPeriod"`
}

// ClusterSnapshotCopyInput is an input type that accepts ClusterSnapshotCopyArgs and ClusterSnapshotCopyOutput values.
// You can construct a concrete instance of `ClusterSnapshotCopyInput` via:
//
//          ClusterSnapshotCopyArgs{...}
type ClusterSnapshotCopyInput interface {
	pulumi.Input

	ToClusterSnapshotCopyOutput() ClusterSnapshotCopyOutput
	ToClusterSnapshotCopyOutputWithContext(context.Context) ClusterSnapshotCopyOutput
}

type ClusterSnapshotCopyArgs struct {
	DestinationRegion pulumi.StringInput    `pulumi:"destinationRegion"`
	GrantName         pulumi.StringPtrInput `pulumi:"grantName"`
	RetentionPeriod   pulumi.IntPtrInput    `pulumi:"retentionPeriod"`
}

func (ClusterSnapshotCopyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSnapshotCopy)(nil)).Elem()
}

func (i ClusterSnapshotCopyArgs) ToClusterSnapshotCopyOutput() ClusterSnapshotCopyOutput {
	return i.ToClusterSnapshotCopyOutputWithContext(context.Background())
}

func (i ClusterSnapshotCopyArgs) ToClusterSnapshotCopyOutputWithContext(ctx context.Context) ClusterSnapshotCopyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSnapshotCopyOutput)
}

func (i ClusterSnapshotCopyArgs) ToClusterSnapshotCopyPtrOutput() ClusterSnapshotCopyPtrOutput {
	return i.ToClusterSnapshotCopyPtrOutputWithContext(context.Background())
}

func (i ClusterSnapshotCopyArgs) ToClusterSnapshotCopyPtrOutputWithContext(ctx context.Context) ClusterSnapshotCopyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSnapshotCopyOutput).ToClusterSnapshotCopyPtrOutputWithContext(ctx)
}

// ClusterSnapshotCopyPtrInput is an input type that accepts ClusterSnapshotCopyArgs, ClusterSnapshotCopyPtr and ClusterSnapshotCopyPtrOutput values.
// You can construct a concrete instance of `ClusterSnapshotCopyPtrInput` via:
//
//          ClusterSnapshotCopyArgs{...}
//
//  or:
//
//          nil
type ClusterSnapshotCopyPtrInput interface {
	pulumi.Input

	ToClusterSnapshotCopyPtrOutput() ClusterSnapshotCopyPtrOutput
	ToClusterSnapshotCopyPtrOutputWithContext(context.Context) ClusterSnapshotCopyPtrOutput
}

type clusterSnapshotCopyPtrType ClusterSnapshotCopyArgs

func ClusterSnapshotCopyPtr(v *ClusterSnapshotCopyArgs) ClusterSnapshotCopyPtrInput {
	return (*clusterSnapshotCopyPtrType)(v)
}

func (*clusterSnapshotCopyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSnapshotCopy)(nil)).Elem()
}

func (i *clusterSnapshotCopyPtrType) ToClusterSnapshotCopyPtrOutput() ClusterSnapshotCopyPtrOutput {
	return i.ToClusterSnapshotCopyPtrOutputWithContext(context.Background())
}

func (i *clusterSnapshotCopyPtrType) ToClusterSnapshotCopyPtrOutputWithContext(ctx context.Context) ClusterSnapshotCopyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterSnapshotCopyPtrOutput)
}

type ClusterSnapshotCopyOutput struct{ *pulumi.OutputState }

func (ClusterSnapshotCopyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSnapshotCopy)(nil)).Elem()
}

func (o ClusterSnapshotCopyOutput) ToClusterSnapshotCopyOutput() ClusterSnapshotCopyOutput {
	return o
}

func (o ClusterSnapshotCopyOutput) ToClusterSnapshotCopyOutputWithContext(ctx context.Context) ClusterSnapshotCopyOutput {
	return o
}

func (o ClusterSnapshotCopyOutput) ToClusterSnapshotCopyPtrOutput() ClusterSnapshotCopyPtrOutput {
	return o.ToClusterSnapshotCopyPtrOutputWithContext(context.Background())
}

func (o ClusterSnapshotCopyOutput) ToClusterSnapshotCopyPtrOutputWithContext(ctx context.Context) ClusterSnapshotCopyPtrOutput {
	return o.ApplyT(func(v ClusterSnapshotCopy) *ClusterSnapshotCopy {
		return &v
	}).(ClusterSnapshotCopyPtrOutput)
}
func (o ClusterSnapshotCopyOutput) DestinationRegion() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterSnapshotCopy) string { return v.DestinationRegion }).(pulumi.StringOutput)
}

func (o ClusterSnapshotCopyOutput) GrantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterSnapshotCopy) *string { return v.GrantName }).(pulumi.StringPtrOutput)
}

func (o ClusterSnapshotCopyOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterSnapshotCopy) *int { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

type ClusterSnapshotCopyPtrOutput struct{ *pulumi.OutputState }

func (ClusterSnapshotCopyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSnapshotCopy)(nil)).Elem()
}

func (o ClusterSnapshotCopyPtrOutput) ToClusterSnapshotCopyPtrOutput() ClusterSnapshotCopyPtrOutput {
	return o
}

func (o ClusterSnapshotCopyPtrOutput) ToClusterSnapshotCopyPtrOutputWithContext(ctx context.Context) ClusterSnapshotCopyPtrOutput {
	return o
}

func (o ClusterSnapshotCopyPtrOutput) Elem() ClusterSnapshotCopyOutput {
	return o.ApplyT(func(v *ClusterSnapshotCopy) ClusterSnapshotCopy { return *v }).(ClusterSnapshotCopyOutput)
}

func (o ClusterSnapshotCopyPtrOutput) DestinationRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSnapshotCopy) *string {
		if v == nil {
			return nil
		}
		return &v.DestinationRegion
	}).(pulumi.StringPtrOutput)
}

func (o ClusterSnapshotCopyPtrOutput) GrantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterSnapshotCopy) *string {
		if v == nil {
			return nil
		}
		return v.GrantName
	}).(pulumi.StringPtrOutput)
}

func (o ClusterSnapshotCopyPtrOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterSnapshotCopy) *int {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.IntPtrOutput)
}

type ParameterGroupParameter struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

// ParameterGroupParameterInput is an input type that accepts ParameterGroupParameterArgs and ParameterGroupParameterOutput values.
// You can construct a concrete instance of `ParameterGroupParameterInput` via:
//
//          ParameterGroupParameterArgs{...}
type ParameterGroupParameterInput interface {
	pulumi.Input

	ToParameterGroupParameterOutput() ParameterGroupParameterOutput
	ToParameterGroupParameterOutputWithContext(context.Context) ParameterGroupParameterOutput
}

type ParameterGroupParameterArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
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
//          ParameterGroupParameterArray{ ParameterGroupParameterArgs{...} }
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

func (o ParameterGroupParameterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ParameterGroupParameter) string { return v.Name }).(pulumi.StringOutput)
}

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

type SecurityGroupIngress struct {
	Cidr                 *string `pulumi:"cidr"`
	SecurityGroupName    *string `pulumi:"securityGroupName"`
	SecurityGroupOwnerId *string `pulumi:"securityGroupOwnerId"`
}

// SecurityGroupIngressInput is an input type that accepts SecurityGroupIngressArgs and SecurityGroupIngressOutput values.
// You can construct a concrete instance of `SecurityGroupIngressInput` via:
//
//          SecurityGroupIngressArgs{...}
type SecurityGroupIngressInput interface {
	pulumi.Input

	ToSecurityGroupIngressOutput() SecurityGroupIngressOutput
	ToSecurityGroupIngressOutputWithContext(context.Context) SecurityGroupIngressOutput
}

type SecurityGroupIngressArgs struct {
	Cidr                 pulumi.StringPtrInput `pulumi:"cidr"`
	SecurityGroupName    pulumi.StringPtrInput `pulumi:"securityGroupName"`
	SecurityGroupOwnerId pulumi.StringPtrInput `pulumi:"securityGroupOwnerId"`
}

func (SecurityGroupIngressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupIngress)(nil)).Elem()
}

func (i SecurityGroupIngressArgs) ToSecurityGroupIngressOutput() SecurityGroupIngressOutput {
	return i.ToSecurityGroupIngressOutputWithContext(context.Background())
}

func (i SecurityGroupIngressArgs) ToSecurityGroupIngressOutputWithContext(ctx context.Context) SecurityGroupIngressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupIngressOutput)
}

// SecurityGroupIngressArrayInput is an input type that accepts SecurityGroupIngressArray and SecurityGroupIngressArrayOutput values.
// You can construct a concrete instance of `SecurityGroupIngressArrayInput` via:
//
//          SecurityGroupIngressArray{ SecurityGroupIngressArgs{...} }
type SecurityGroupIngressArrayInput interface {
	pulumi.Input

	ToSecurityGroupIngressArrayOutput() SecurityGroupIngressArrayOutput
	ToSecurityGroupIngressArrayOutputWithContext(context.Context) SecurityGroupIngressArrayOutput
}

type SecurityGroupIngressArray []SecurityGroupIngressInput

func (SecurityGroupIngressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroupIngress)(nil)).Elem()
}

func (i SecurityGroupIngressArray) ToSecurityGroupIngressArrayOutput() SecurityGroupIngressArrayOutput {
	return i.ToSecurityGroupIngressArrayOutputWithContext(context.Background())
}

func (i SecurityGroupIngressArray) ToSecurityGroupIngressArrayOutputWithContext(ctx context.Context) SecurityGroupIngressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupIngressArrayOutput)
}

type SecurityGroupIngressOutput struct{ *pulumi.OutputState }

func (SecurityGroupIngressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupIngress)(nil)).Elem()
}

func (o SecurityGroupIngressOutput) ToSecurityGroupIngressOutput() SecurityGroupIngressOutput {
	return o
}

func (o SecurityGroupIngressOutput) ToSecurityGroupIngressOutputWithContext(ctx context.Context) SecurityGroupIngressOutput {
	return o
}

func (o SecurityGroupIngressOutput) Cidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupIngress) *string { return v.Cidr }).(pulumi.StringPtrOutput)
}

func (o SecurityGroupIngressOutput) SecurityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupIngress) *string { return v.SecurityGroupName }).(pulumi.StringPtrOutput)
}

func (o SecurityGroupIngressOutput) SecurityGroupOwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityGroupIngress) *string { return v.SecurityGroupOwnerId }).(pulumi.StringPtrOutput)
}

type SecurityGroupIngressArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupIngressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroupIngress)(nil)).Elem()
}

func (o SecurityGroupIngressArrayOutput) ToSecurityGroupIngressArrayOutput() SecurityGroupIngressArrayOutput {
	return o
}

func (o SecurityGroupIngressArrayOutput) ToSecurityGroupIngressArrayOutputWithContext(ctx context.Context) SecurityGroupIngressArrayOutput {
	return o
}

func (o SecurityGroupIngressArrayOutput) Index(i pulumi.IntInput) SecurityGroupIngressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityGroupIngress {
		return vs[0].([]SecurityGroupIngress)[vs[1].(int)]
	}).(SecurityGroupIngressOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterLoggingOutput{})
	pulumi.RegisterOutputType(ClusterLoggingPtrOutput{})
	pulumi.RegisterOutputType(ClusterSnapshotCopyOutput{})
	pulumi.RegisterOutputType(ClusterSnapshotCopyPtrOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterOutput{})
	pulumi.RegisterOutputType(ParameterGroupParameterArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupIngressOutput{})
	pulumi.RegisterOutputType(SecurityGroupIngressArrayOutput{})
}
