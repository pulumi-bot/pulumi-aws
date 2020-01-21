// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package SecurityGroupEgress

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type SecurityGroupEgress struct {
	CidrBlocks []string `pulumi:"cidrBlocks"`
	// The security group description. Defaults to
	// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
	// `GroupDescription` attribute, for which there is no Update API. If you'd like
	// to classify your security groups in a way that can be updated, use `tags`.
	Description *string `pulumi:"description"`
	FromPort int `pulumi:"fromPort"`
	Ipv6CidrBlocks []string `pulumi:"ipv6CidrBlocks"`
	PrefixListIds []string `pulumi:"prefixListIds"`
	Protocol string `pulumi:"protocol"`
	SecurityGroups []string `pulumi:"securityGroups"`
	Self *bool `pulumi:"self"`
	ToPort int `pulumi:"toPort"`
}

type SecurityGroupEgressInput interface {
	pulumi.Input

	ToSecurityGroupEgressOutput() SecurityGroupEgressOutput
	ToSecurityGroupEgressOutputWithContext(context.Context) SecurityGroupEgressOutput
}

type SecurityGroupEgressArgs struct {
	CidrBlocks pulumi.StringArrayInput `pulumi:"cidrBlocks"`
	// The security group description. Defaults to
	// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
	// `GroupDescription` attribute, for which there is no Update API. If you'd like
	// to classify your security groups in a way that can be updated, use `tags`.
	Description pulumi.StringPtrInput `pulumi:"description"`
	FromPort pulumi.IntInput `pulumi:"fromPort"`
	Ipv6CidrBlocks pulumi.StringArrayInput `pulumi:"ipv6CidrBlocks"`
	PrefixListIds pulumi.StringArrayInput `pulumi:"prefixListIds"`
	Protocol pulumi.StringInput `pulumi:"protocol"`
	SecurityGroups pulumi.StringArrayInput `pulumi:"securityGroups"`
	Self pulumi.BoolPtrInput `pulumi:"self"`
	ToPort pulumi.IntInput `pulumi:"toPort"`
}

func (SecurityGroupEgressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupEgress)(nil)).Elem()
}

func (i SecurityGroupEgressArgs) ToSecurityGroupEgressOutput() SecurityGroupEgressOutput {
	return i.ToSecurityGroupEgressOutputWithContext(context.Background())
}

func (i SecurityGroupEgressArgs) ToSecurityGroupEgressOutputWithContext(ctx context.Context) SecurityGroupEgressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupEgressOutput)
}

type SecurityGroupEgressArrayInput interface {
	pulumi.Input

	ToSecurityGroupEgressArrayOutput() SecurityGroupEgressArrayOutput
	ToSecurityGroupEgressArrayOutputWithContext(context.Context) SecurityGroupEgressArrayOutput
}

type SecurityGroupEgressArray []SecurityGroupEgressInput

func (SecurityGroupEgressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroupEgress)(nil)).Elem()
}

func (i SecurityGroupEgressArray) ToSecurityGroupEgressArrayOutput() SecurityGroupEgressArrayOutput {
	return i.ToSecurityGroupEgressArrayOutputWithContext(context.Background())
}

func (i SecurityGroupEgressArray) ToSecurityGroupEgressArrayOutputWithContext(ctx context.Context) SecurityGroupEgressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupEgressArrayOutput)
}

type SecurityGroupEgressOutput struct { *pulumi.OutputState }

func (SecurityGroupEgressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupEgress)(nil)).Elem()
}

func (o SecurityGroupEgressOutput) ToSecurityGroupEgressOutput() SecurityGroupEgressOutput {
	return o
}

func (o SecurityGroupEgressOutput) ToSecurityGroupEgressOutputWithContext(ctx context.Context) SecurityGroupEgressOutput {
	return o
}

func (o SecurityGroupEgressOutput) CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func (v SecurityGroupEgress) []string { return v.CidrBlocks }).(pulumi.StringArrayOutput)
}

// The security group description. Defaults to
// "Managed by Pulumi". Cannot be "". __NOTE__: This field maps to the AWS
// `GroupDescription` attribute, for which there is no Update API. If you'd like
// to classify your security groups in a way that can be updated, use `tags`.
func (o SecurityGroupEgressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func (v SecurityGroupEgress) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityGroupEgressOutput) FromPort() pulumi.IntOutput {
	return o.ApplyT(func (v SecurityGroupEgress) int { return v.FromPort }).(pulumi.IntOutput)
}

func (o SecurityGroupEgressOutput) Ipv6CidrBlocks() pulumi.StringArrayOutput {
	return o.ApplyT(func (v SecurityGroupEgress) []string { return v.Ipv6CidrBlocks }).(pulumi.StringArrayOutput)
}

func (o SecurityGroupEgressOutput) PrefixListIds() pulumi.StringArrayOutput {
	return o.ApplyT(func (v SecurityGroupEgress) []string { return v.PrefixListIds }).(pulumi.StringArrayOutput)
}

func (o SecurityGroupEgressOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func (v SecurityGroupEgress) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o SecurityGroupEgressOutput) SecurityGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func (v SecurityGroupEgress) []string { return v.SecurityGroups }).(pulumi.StringArrayOutput)
}

func (o SecurityGroupEgressOutput) Self() pulumi.BoolPtrOutput {
	return o.ApplyT(func (v SecurityGroupEgress) *bool { return v.Self }).(pulumi.BoolPtrOutput)
}

func (o SecurityGroupEgressOutput) ToPort() pulumi.IntOutput {
	return o.ApplyT(func (v SecurityGroupEgress) int { return v.ToPort }).(pulumi.IntOutput)
}

type SecurityGroupEgressArrayOutput struct { *pulumi.OutputState}

func (SecurityGroupEgressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityGroupEgress)(nil)).Elem()
}

func (o SecurityGroupEgressArrayOutput) ToSecurityGroupEgressArrayOutput() SecurityGroupEgressArrayOutput {
	return o
}

func (o SecurityGroupEgressArrayOutput) ToSecurityGroupEgressArrayOutputWithContext(ctx context.Context) SecurityGroupEgressArrayOutput {
	return o
}

func (o SecurityGroupEgressArrayOutput) Index(i pulumi.IntInput) SecurityGroupEgressOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) SecurityGroupEgress {
		return vs[0].([]SecurityGroupEgress)[vs[1].(int)]
	}).(SecurityGroupEgressOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityGroupEgressOutput{})
	pulumi.RegisterOutputType(SecurityGroupEgressArrayOutput{})
}
