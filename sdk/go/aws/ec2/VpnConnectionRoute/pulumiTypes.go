// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package VpnConnectionRoute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type VpnConnectionRoute struct {
	DestinationCidrBlock *string `pulumi:"destinationCidrBlock"`
	Source *string `pulumi:"source"`
	State *string `pulumi:"state"`
}

type VpnConnectionRouteInput interface {
	pulumi.Input

	ToVpnConnectionRouteOutput() VpnConnectionRouteOutput
	ToVpnConnectionRouteOutputWithContext(context.Context) VpnConnectionRouteOutput
}

type VpnConnectionRouteArgs struct {
	DestinationCidrBlock pulumi.StringPtrInput `pulumi:"destinationCidrBlock"`
	Source pulumi.StringPtrInput `pulumi:"source"`
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (VpnConnectionRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionRoute)(nil)).Elem()
}

func (i VpnConnectionRouteArgs) ToVpnConnectionRouteOutput() VpnConnectionRouteOutput {
	return i.ToVpnConnectionRouteOutputWithContext(context.Background())
}

func (i VpnConnectionRouteArgs) ToVpnConnectionRouteOutputWithContext(ctx context.Context) VpnConnectionRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionRouteOutput)
}

type VpnConnectionRouteArrayInput interface {
	pulumi.Input

	ToVpnConnectionRouteArrayOutput() VpnConnectionRouteArrayOutput
	ToVpnConnectionRouteArrayOutputWithContext(context.Context) VpnConnectionRouteArrayOutput
}

type VpnConnectionRouteArray []VpnConnectionRouteInput

func (VpnConnectionRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionRoute)(nil)).Elem()
}

func (i VpnConnectionRouteArray) ToVpnConnectionRouteArrayOutput() VpnConnectionRouteArrayOutput {
	return i.ToVpnConnectionRouteArrayOutputWithContext(context.Background())
}

func (i VpnConnectionRouteArray) ToVpnConnectionRouteArrayOutputWithContext(ctx context.Context) VpnConnectionRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnConnectionRouteArrayOutput)
}

type VpnConnectionRouteOutput struct { *pulumi.OutputState }

func (VpnConnectionRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnConnectionRoute)(nil)).Elem()
}

func (o VpnConnectionRouteOutput) ToVpnConnectionRouteOutput() VpnConnectionRouteOutput {
	return o
}

func (o VpnConnectionRouteOutput) ToVpnConnectionRouteOutputWithContext(ctx context.Context) VpnConnectionRouteOutput {
	return o
}

func (o VpnConnectionRouteOutput) DestinationCidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VpnConnectionRoute) *string { return v.DestinationCidrBlock }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionRouteOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VpnConnectionRoute) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o VpnConnectionRouteOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func (v VpnConnectionRoute) *string { return v.State }).(pulumi.StringPtrOutput)
}

type VpnConnectionRouteArrayOutput struct { *pulumi.OutputState}

func (VpnConnectionRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpnConnectionRoute)(nil)).Elem()
}

func (o VpnConnectionRouteArrayOutput) ToVpnConnectionRouteArrayOutput() VpnConnectionRouteArrayOutput {
	return o
}

func (o VpnConnectionRouteArrayOutput) ToVpnConnectionRouteArrayOutputWithContext(ctx context.Context) VpnConnectionRouteArrayOutput {
	return o
}

func (o VpnConnectionRouteArrayOutput) Index(i pulumi.IntInput) VpnConnectionRouteOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) VpnConnectionRoute {
		return vs[0].([]VpnConnectionRoute)[vs[1].(int)]
	}).(VpnConnectionRouteOutput)
}

func init() {
	pulumi.RegisterOutputType(VpnConnectionRouteOutput{})
	pulumi.RegisterOutputType(VpnConnectionRouteArrayOutput{})
}
