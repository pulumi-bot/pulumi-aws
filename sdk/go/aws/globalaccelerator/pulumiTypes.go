// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AcceleratorAttributes struct {
	FlowLogsEnabled  *bool   `pulumi:"flowLogsEnabled"`
	FlowLogsS3Bucket *string `pulumi:"flowLogsS3Bucket"`
	FlowLogsS3Prefix *string `pulumi:"flowLogsS3Prefix"`
}

// AcceleratorAttributesInput is an input type that accepts AcceleratorAttributesArgs and AcceleratorAttributesOutput values.
// You can construct a concrete instance of `AcceleratorAttributesInput` via:
//
//          AcceleratorAttributesArgs{...}
type AcceleratorAttributesInput interface {
	pulumi.Input

	ToAcceleratorAttributesOutput() AcceleratorAttributesOutput
	ToAcceleratorAttributesOutputWithContext(context.Context) AcceleratorAttributesOutput
}

type AcceleratorAttributesArgs struct {
	FlowLogsEnabled  pulumi.BoolPtrInput   `pulumi:"flowLogsEnabled"`
	FlowLogsS3Bucket pulumi.StringPtrInput `pulumi:"flowLogsS3Bucket"`
	FlowLogsS3Prefix pulumi.StringPtrInput `pulumi:"flowLogsS3Prefix"`
}

func (AcceleratorAttributesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorAttributes)(nil)).Elem()
}

func (i AcceleratorAttributesArgs) ToAcceleratorAttributesOutput() AcceleratorAttributesOutput {
	return i.ToAcceleratorAttributesOutputWithContext(context.Background())
}

func (i AcceleratorAttributesArgs) ToAcceleratorAttributesOutputWithContext(ctx context.Context) AcceleratorAttributesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorAttributesOutput)
}

func (i AcceleratorAttributesArgs) ToAcceleratorAttributesPtrOutput() AcceleratorAttributesPtrOutput {
	return i.ToAcceleratorAttributesPtrOutputWithContext(context.Background())
}

func (i AcceleratorAttributesArgs) ToAcceleratorAttributesPtrOutputWithContext(ctx context.Context) AcceleratorAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorAttributesOutput).ToAcceleratorAttributesPtrOutputWithContext(ctx)
}

// AcceleratorAttributesPtrInput is an input type that accepts AcceleratorAttributesArgs, AcceleratorAttributesPtr and AcceleratorAttributesPtrOutput values.
// You can construct a concrete instance of `AcceleratorAttributesPtrInput` via:
//
//          AcceleratorAttributesArgs{...}
//
//  or:
//
//          nil
type AcceleratorAttributesPtrInput interface {
	pulumi.Input

	ToAcceleratorAttributesPtrOutput() AcceleratorAttributesPtrOutput
	ToAcceleratorAttributesPtrOutputWithContext(context.Context) AcceleratorAttributesPtrOutput
}

type acceleratorAttributesPtrType AcceleratorAttributesArgs

func AcceleratorAttributesPtr(v *AcceleratorAttributesArgs) AcceleratorAttributesPtrInput {
	return (*acceleratorAttributesPtrType)(v)
}

func (*acceleratorAttributesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AcceleratorAttributes)(nil)).Elem()
}

func (i *acceleratorAttributesPtrType) ToAcceleratorAttributesPtrOutput() AcceleratorAttributesPtrOutput {
	return i.ToAcceleratorAttributesPtrOutputWithContext(context.Background())
}

func (i *acceleratorAttributesPtrType) ToAcceleratorAttributesPtrOutputWithContext(ctx context.Context) AcceleratorAttributesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorAttributesPtrOutput)
}

type AcceleratorAttributesOutput struct{ *pulumi.OutputState }

func (AcceleratorAttributesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorAttributes)(nil)).Elem()
}

func (o AcceleratorAttributesOutput) ToAcceleratorAttributesOutput() AcceleratorAttributesOutput {
	return o
}

func (o AcceleratorAttributesOutput) ToAcceleratorAttributesOutputWithContext(ctx context.Context) AcceleratorAttributesOutput {
	return o
}

func (o AcceleratorAttributesOutput) ToAcceleratorAttributesPtrOutput() AcceleratorAttributesPtrOutput {
	return o.ToAcceleratorAttributesPtrOutputWithContext(context.Background())
}

func (o AcceleratorAttributesOutput) ToAcceleratorAttributesPtrOutputWithContext(ctx context.Context) AcceleratorAttributesPtrOutput {
	return o.ApplyT(func(v AcceleratorAttributes) *AcceleratorAttributes {
		return &v
	}).(AcceleratorAttributesPtrOutput)
}
func (o AcceleratorAttributesOutput) FlowLogsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AcceleratorAttributes) *bool { return v.FlowLogsEnabled }).(pulumi.BoolPtrOutput)
}

func (o AcceleratorAttributesOutput) FlowLogsS3Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AcceleratorAttributes) *string { return v.FlowLogsS3Bucket }).(pulumi.StringPtrOutput)
}

func (o AcceleratorAttributesOutput) FlowLogsS3Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AcceleratorAttributes) *string { return v.FlowLogsS3Prefix }).(pulumi.StringPtrOutput)
}

type AcceleratorAttributesPtrOutput struct{ *pulumi.OutputState }

func (AcceleratorAttributesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AcceleratorAttributes)(nil)).Elem()
}

func (o AcceleratorAttributesPtrOutput) ToAcceleratorAttributesPtrOutput() AcceleratorAttributesPtrOutput {
	return o
}

func (o AcceleratorAttributesPtrOutput) ToAcceleratorAttributesPtrOutputWithContext(ctx context.Context) AcceleratorAttributesPtrOutput {
	return o
}

func (o AcceleratorAttributesPtrOutput) Elem() AcceleratorAttributesOutput {
	return o.ApplyT(func(v *AcceleratorAttributes) AcceleratorAttributes { return *v }).(AcceleratorAttributesOutput)
}

func (o AcceleratorAttributesPtrOutput) FlowLogsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AcceleratorAttributes) *bool {
		if v == nil {
			return nil
		}
		return v.FlowLogsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AcceleratorAttributesPtrOutput) FlowLogsS3Bucket() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AcceleratorAttributes) *string {
		if v == nil {
			return nil
		}
		return v.FlowLogsS3Bucket
	}).(pulumi.StringPtrOutput)
}

func (o AcceleratorAttributesPtrOutput) FlowLogsS3Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AcceleratorAttributes) *string {
		if v == nil {
			return nil
		}
		return v.FlowLogsS3Prefix
	}).(pulumi.StringPtrOutput)
}

type AcceleratorIpSet struct {
	IpAddresses []string `pulumi:"ipAddresses"`
	IpFamily    *string  `pulumi:"ipFamily"`
}

// AcceleratorIpSetInput is an input type that accepts AcceleratorIpSetArgs and AcceleratorIpSetOutput values.
// You can construct a concrete instance of `AcceleratorIpSetInput` via:
//
//          AcceleratorIpSetArgs{...}
type AcceleratorIpSetInput interface {
	pulumi.Input

	ToAcceleratorIpSetOutput() AcceleratorIpSetOutput
	ToAcceleratorIpSetOutputWithContext(context.Context) AcceleratorIpSetOutput
}

type AcceleratorIpSetArgs struct {
	IpAddresses pulumi.StringArrayInput `pulumi:"ipAddresses"`
	IpFamily    pulumi.StringPtrInput   `pulumi:"ipFamily"`
}

func (AcceleratorIpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorIpSet)(nil)).Elem()
}

func (i AcceleratorIpSetArgs) ToAcceleratorIpSetOutput() AcceleratorIpSetOutput {
	return i.ToAcceleratorIpSetOutputWithContext(context.Background())
}

func (i AcceleratorIpSetArgs) ToAcceleratorIpSetOutputWithContext(ctx context.Context) AcceleratorIpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorIpSetOutput)
}

// AcceleratorIpSetArrayInput is an input type that accepts AcceleratorIpSetArray and AcceleratorIpSetArrayOutput values.
// You can construct a concrete instance of `AcceleratorIpSetArrayInput` via:
//
//          AcceleratorIpSetArray{ AcceleratorIpSetArgs{...} }
type AcceleratorIpSetArrayInput interface {
	pulumi.Input

	ToAcceleratorIpSetArrayOutput() AcceleratorIpSetArrayOutput
	ToAcceleratorIpSetArrayOutputWithContext(context.Context) AcceleratorIpSetArrayOutput
}

type AcceleratorIpSetArray []AcceleratorIpSetInput

func (AcceleratorIpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AcceleratorIpSet)(nil)).Elem()
}

func (i AcceleratorIpSetArray) ToAcceleratorIpSetArrayOutput() AcceleratorIpSetArrayOutput {
	return i.ToAcceleratorIpSetArrayOutputWithContext(context.Background())
}

func (i AcceleratorIpSetArray) ToAcceleratorIpSetArrayOutputWithContext(ctx context.Context) AcceleratorIpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorIpSetArrayOutput)
}

type AcceleratorIpSetOutput struct{ *pulumi.OutputState }

func (AcceleratorIpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AcceleratorIpSet)(nil)).Elem()
}

func (o AcceleratorIpSetOutput) ToAcceleratorIpSetOutput() AcceleratorIpSetOutput {
	return o
}

func (o AcceleratorIpSetOutput) ToAcceleratorIpSetOutputWithContext(ctx context.Context) AcceleratorIpSetOutput {
	return o
}

func (o AcceleratorIpSetOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AcceleratorIpSet) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o AcceleratorIpSetOutput) IpFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AcceleratorIpSet) *string { return v.IpFamily }).(pulumi.StringPtrOutput)
}

type AcceleratorIpSetArrayOutput struct{ *pulumi.OutputState }

func (AcceleratorIpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AcceleratorIpSet)(nil)).Elem()
}

func (o AcceleratorIpSetArrayOutput) ToAcceleratorIpSetArrayOutput() AcceleratorIpSetArrayOutput {
	return o
}

func (o AcceleratorIpSetArrayOutput) ToAcceleratorIpSetArrayOutputWithContext(ctx context.Context) AcceleratorIpSetArrayOutput {
	return o
}

func (o AcceleratorIpSetArrayOutput) Index(i pulumi.IntInput) AcceleratorIpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AcceleratorIpSet {
		return vs[0].([]AcceleratorIpSet)[vs[1].(int)]
	}).(AcceleratorIpSetOutput)
}

type EndpointGroupEndpointConfiguration struct {
	EndpointId *string `pulumi:"endpointId"`
	Weight     *int    `pulumi:"weight"`
}

// EndpointGroupEndpointConfigurationInput is an input type that accepts EndpointGroupEndpointConfigurationArgs and EndpointGroupEndpointConfigurationOutput values.
// You can construct a concrete instance of `EndpointGroupEndpointConfigurationInput` via:
//
//          EndpointGroupEndpointConfigurationArgs{...}
type EndpointGroupEndpointConfigurationInput interface {
	pulumi.Input

	ToEndpointGroupEndpointConfigurationOutput() EndpointGroupEndpointConfigurationOutput
	ToEndpointGroupEndpointConfigurationOutputWithContext(context.Context) EndpointGroupEndpointConfigurationOutput
}

type EndpointGroupEndpointConfigurationArgs struct {
	EndpointId pulumi.StringPtrInput `pulumi:"endpointId"`
	Weight     pulumi.IntPtrInput    `pulumi:"weight"`
}

func (EndpointGroupEndpointConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointGroupEndpointConfiguration)(nil)).Elem()
}

func (i EndpointGroupEndpointConfigurationArgs) ToEndpointGroupEndpointConfigurationOutput() EndpointGroupEndpointConfigurationOutput {
	return i.ToEndpointGroupEndpointConfigurationOutputWithContext(context.Background())
}

func (i EndpointGroupEndpointConfigurationArgs) ToEndpointGroupEndpointConfigurationOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupEndpointConfigurationOutput)
}

// EndpointGroupEndpointConfigurationArrayInput is an input type that accepts EndpointGroupEndpointConfigurationArray and EndpointGroupEndpointConfigurationArrayOutput values.
// You can construct a concrete instance of `EndpointGroupEndpointConfigurationArrayInput` via:
//
//          EndpointGroupEndpointConfigurationArray{ EndpointGroupEndpointConfigurationArgs{...} }
type EndpointGroupEndpointConfigurationArrayInput interface {
	pulumi.Input

	ToEndpointGroupEndpointConfigurationArrayOutput() EndpointGroupEndpointConfigurationArrayOutput
	ToEndpointGroupEndpointConfigurationArrayOutputWithContext(context.Context) EndpointGroupEndpointConfigurationArrayOutput
}

type EndpointGroupEndpointConfigurationArray []EndpointGroupEndpointConfigurationInput

func (EndpointGroupEndpointConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointGroupEndpointConfiguration)(nil)).Elem()
}

func (i EndpointGroupEndpointConfigurationArray) ToEndpointGroupEndpointConfigurationArrayOutput() EndpointGroupEndpointConfigurationArrayOutput {
	return i.ToEndpointGroupEndpointConfigurationArrayOutputWithContext(context.Background())
}

func (i EndpointGroupEndpointConfigurationArray) ToEndpointGroupEndpointConfigurationArrayOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointGroupEndpointConfigurationArrayOutput)
}

type EndpointGroupEndpointConfigurationOutput struct{ *pulumi.OutputState }

func (EndpointGroupEndpointConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointGroupEndpointConfiguration)(nil)).Elem()
}

func (o EndpointGroupEndpointConfigurationOutput) ToEndpointGroupEndpointConfigurationOutput() EndpointGroupEndpointConfigurationOutput {
	return o
}

func (o EndpointGroupEndpointConfigurationOutput) ToEndpointGroupEndpointConfigurationOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationOutput {
	return o
}

func (o EndpointGroupEndpointConfigurationOutput) EndpointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointGroupEndpointConfiguration) *string { return v.EndpointId }).(pulumi.StringPtrOutput)
}

func (o EndpointGroupEndpointConfigurationOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointGroupEndpointConfiguration) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type EndpointGroupEndpointConfigurationArrayOutput struct{ *pulumi.OutputState }

func (EndpointGroupEndpointConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointGroupEndpointConfiguration)(nil)).Elem()
}

func (o EndpointGroupEndpointConfigurationArrayOutput) ToEndpointGroupEndpointConfigurationArrayOutput() EndpointGroupEndpointConfigurationArrayOutput {
	return o
}

func (o EndpointGroupEndpointConfigurationArrayOutput) ToEndpointGroupEndpointConfigurationArrayOutputWithContext(ctx context.Context) EndpointGroupEndpointConfigurationArrayOutput {
	return o
}

func (o EndpointGroupEndpointConfigurationArrayOutput) Index(i pulumi.IntInput) EndpointGroupEndpointConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointGroupEndpointConfiguration {
		return vs[0].([]EndpointGroupEndpointConfiguration)[vs[1].(int)]
	}).(EndpointGroupEndpointConfigurationOutput)
}

type ListenerPortRange struct {
	FromPort *int `pulumi:"fromPort"`
	ToPort   *int `pulumi:"toPort"`
}

// ListenerPortRangeInput is an input type that accepts ListenerPortRangeArgs and ListenerPortRangeOutput values.
// You can construct a concrete instance of `ListenerPortRangeInput` via:
//
//          ListenerPortRangeArgs{...}
type ListenerPortRangeInput interface {
	pulumi.Input

	ToListenerPortRangeOutput() ListenerPortRangeOutput
	ToListenerPortRangeOutputWithContext(context.Context) ListenerPortRangeOutput
}

type ListenerPortRangeArgs struct {
	FromPort pulumi.IntPtrInput `pulumi:"fromPort"`
	ToPort   pulumi.IntPtrInput `pulumi:"toPort"`
}

func (ListenerPortRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerPortRange)(nil)).Elem()
}

func (i ListenerPortRangeArgs) ToListenerPortRangeOutput() ListenerPortRangeOutput {
	return i.ToListenerPortRangeOutputWithContext(context.Background())
}

func (i ListenerPortRangeArgs) ToListenerPortRangeOutputWithContext(ctx context.Context) ListenerPortRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPortRangeOutput)
}

// ListenerPortRangeArrayInput is an input type that accepts ListenerPortRangeArray and ListenerPortRangeArrayOutput values.
// You can construct a concrete instance of `ListenerPortRangeArrayInput` via:
//
//          ListenerPortRangeArray{ ListenerPortRangeArgs{...} }
type ListenerPortRangeArrayInput interface {
	pulumi.Input

	ToListenerPortRangeArrayOutput() ListenerPortRangeArrayOutput
	ToListenerPortRangeArrayOutputWithContext(context.Context) ListenerPortRangeArrayOutput
}

type ListenerPortRangeArray []ListenerPortRangeInput

func (ListenerPortRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerPortRange)(nil)).Elem()
}

func (i ListenerPortRangeArray) ToListenerPortRangeArrayOutput() ListenerPortRangeArrayOutput {
	return i.ToListenerPortRangeArrayOutputWithContext(context.Background())
}

func (i ListenerPortRangeArray) ToListenerPortRangeArrayOutputWithContext(ctx context.Context) ListenerPortRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ListenerPortRangeArrayOutput)
}

type ListenerPortRangeOutput struct{ *pulumi.OutputState }

func (ListenerPortRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListenerPortRange)(nil)).Elem()
}

func (o ListenerPortRangeOutput) ToListenerPortRangeOutput() ListenerPortRangeOutput {
	return o
}

func (o ListenerPortRangeOutput) ToListenerPortRangeOutputWithContext(ctx context.Context) ListenerPortRangeOutput {
	return o
}

func (o ListenerPortRangeOutput) FromPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListenerPortRange) *int { return v.FromPort }).(pulumi.IntPtrOutput)
}

func (o ListenerPortRangeOutput) ToPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ListenerPortRange) *int { return v.ToPort }).(pulumi.IntPtrOutput)
}

type ListenerPortRangeArrayOutput struct{ *pulumi.OutputState }

func (ListenerPortRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ListenerPortRange)(nil)).Elem()
}

func (o ListenerPortRangeArrayOutput) ToListenerPortRangeArrayOutput() ListenerPortRangeArrayOutput {
	return o
}

func (o ListenerPortRangeArrayOutput) ToListenerPortRangeArrayOutputWithContext(ctx context.Context) ListenerPortRangeArrayOutput {
	return o
}

func (o ListenerPortRangeArrayOutput) Index(i pulumi.IntInput) ListenerPortRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ListenerPortRange {
		return vs[0].([]ListenerPortRange)[vs[1].(int)]
	}).(ListenerPortRangeOutput)
}

func init() {
	pulumi.RegisterOutputType(AcceleratorAttributesOutput{})
	pulumi.RegisterOutputType(AcceleratorAttributesPtrOutput{})
	pulumi.RegisterOutputType(AcceleratorIpSetOutput{})
	pulumi.RegisterOutputType(AcceleratorIpSetArrayOutput{})
	pulumi.RegisterOutputType(EndpointGroupEndpointConfigurationOutput{})
	pulumi.RegisterOutputType(EndpointGroupEndpointConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ListenerPortRangeOutput{})
	pulumi.RegisterOutputType(ListenerPortRangeArrayOutput{})
}
