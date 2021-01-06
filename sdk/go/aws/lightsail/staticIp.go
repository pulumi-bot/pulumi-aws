// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lightsail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allocates a static IP address.
//
// > **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/lightsail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := lightsail.NewStaticIp(ctx, "test", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type StaticIp struct {
	pulumi.CustomResourceState

	// The ARN of the Lightsail static IP
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The allocated static IP address
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The name for the allocated static IP
	Name pulumi.StringOutput `pulumi:"name"`
	// The support code.
	SupportCode pulumi.StringOutput `pulumi:"supportCode"`
}

// NewStaticIp registers a new resource with the given unique name, arguments, and options.
func NewStaticIp(ctx *pulumi.Context,
	name string, args *StaticIpArgs, opts ...pulumi.ResourceOption) (*StaticIp, error) {
	if args == nil {
		args = &StaticIpArgs{}
	}

	var resource StaticIp
	err := ctx.RegisterResource("aws:lightsail/staticIp:StaticIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStaticIp gets an existing StaticIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStaticIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticIpState, opts ...pulumi.ResourceOption) (*StaticIp, error) {
	var resource StaticIp
	err := ctx.ReadResource("aws:lightsail/staticIp:StaticIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StaticIp resources.
type staticIpState struct {
	// The ARN of the Lightsail static IP
	Arn *string `pulumi:"arn"`
	// The allocated static IP address
	IpAddress *string `pulumi:"ipAddress"`
	// The name for the allocated static IP
	Name *string `pulumi:"name"`
	// The support code.
	SupportCode *string `pulumi:"supportCode"`
}

type StaticIpState struct {
	// The ARN of the Lightsail static IP
	Arn pulumi.StringPtrInput
	// The allocated static IP address
	IpAddress pulumi.StringPtrInput
	// The name for the allocated static IP
	Name pulumi.StringPtrInput
	// The support code.
	SupportCode pulumi.StringPtrInput
}

func (StaticIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpState)(nil)).Elem()
}

type staticIpArgs struct {
	// The name for the allocated static IP
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a StaticIp resource.
type StaticIpArgs struct {
	// The name for the allocated static IP
	Name pulumi.StringPtrInput
}

func (StaticIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticIpArgs)(nil)).Elem()
}

type StaticIpInput interface {
	pulumi.Input

	ToStaticIpOutput() StaticIpOutput
	ToStaticIpOutputWithContext(ctx context.Context) StaticIpOutput
}

func (*StaticIp) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticIp)(nil))
}

func (i *StaticIp) ToStaticIpOutput() StaticIpOutput {
	return i.ToStaticIpOutputWithContext(context.Background())
}

func (i *StaticIp) ToStaticIpOutputWithContext(ctx context.Context) StaticIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpOutput)
}

func (i *StaticIp) ToStaticIpPtrOutput() StaticIpPtrOutput {
	return i.ToStaticIpPtrOutputWithContext(context.Background())
}

func (i *StaticIp) ToStaticIpPtrOutputWithContext(ctx context.Context) StaticIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpPtrOutput)
}

type StaticIpPtrInput interface {
	pulumi.Input

	ToStaticIpPtrOutput() StaticIpPtrOutput
	ToStaticIpPtrOutputWithContext(ctx context.Context) StaticIpPtrOutput
}

type staticIpPtrType StaticIpArgs

func (*staticIpPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticIp)(nil))
}

func (i *staticIpPtrType) ToStaticIpPtrOutput() StaticIpPtrOutput {
	return i.ToStaticIpPtrOutputWithContext(context.Background())
}

func (i *staticIpPtrType) ToStaticIpPtrOutputWithContext(ctx context.Context) StaticIpPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpOutput).ToStaticIpPtrOutput()
}

// StaticIpArrayInput is an input type that accepts StaticIpArray and StaticIpArrayOutput values.
// You can construct a concrete instance of `StaticIpArrayInput` via:
//
//          StaticIpArray{ StaticIpArgs{...} }
type StaticIpArrayInput interface {
	pulumi.Input

	ToStaticIpArrayOutput() StaticIpArrayOutput
	ToStaticIpArrayOutputWithContext(context.Context) StaticIpArrayOutput
}

type StaticIpArray []StaticIpInput

func (StaticIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticIp)(nil))
}

func (i StaticIpArray) ToStaticIpArrayOutput() StaticIpArrayOutput {
	return i.ToStaticIpArrayOutputWithContext(context.Background())
}

func (i StaticIpArray) ToStaticIpArrayOutputWithContext(ctx context.Context) StaticIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpArrayOutput)
}

// StaticIpMapInput is an input type that accepts StaticIpMap and StaticIpMapOutput values.
// You can construct a concrete instance of `StaticIpMapInput` via:
//
//          StaticIpMap{ "key": StaticIpArgs{...} }
type StaticIpMapInput interface {
	pulumi.Input

	ToStaticIpMapOutput() StaticIpMapOutput
	ToStaticIpMapOutputWithContext(context.Context) StaticIpMapOutput
}

type StaticIpMap map[string]StaticIpInput

func (StaticIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StaticIp)(nil))
}

func (i StaticIpMap) ToStaticIpMapOutput() StaticIpMapOutput {
	return i.ToStaticIpMapOutputWithContext(context.Background())
}

func (i StaticIpMap) ToStaticIpMapOutputWithContext(ctx context.Context) StaticIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticIpMapOutput)
}

type StaticIpOutput struct {
	*pulumi.OutputState
}

func (StaticIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticIp)(nil))
}

func (o StaticIpOutput) ToStaticIpOutput() StaticIpOutput {
	return o
}

func (o StaticIpOutput) ToStaticIpOutputWithContext(ctx context.Context) StaticIpOutput {
	return o
}

func (o StaticIpOutput) ToStaticIpPtrOutput() StaticIpPtrOutput {
	return o.ToStaticIpPtrOutputWithContext(context.Background())
}

func (o StaticIpOutput) ToStaticIpPtrOutputWithContext(ctx context.Context) StaticIpPtrOutput {
	return o.ApplyT(func(v StaticIp) *StaticIp {
		return &v
	}).(StaticIpPtrOutput)
}

type StaticIpPtrOutput struct {
	*pulumi.OutputState
}

func (StaticIpPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticIp)(nil))
}

func (o StaticIpPtrOutput) ToStaticIpPtrOutput() StaticIpPtrOutput {
	return o
}

func (o StaticIpPtrOutput) ToStaticIpPtrOutputWithContext(ctx context.Context) StaticIpPtrOutput {
	return o
}

type StaticIpArrayOutput struct{ *pulumi.OutputState }

func (StaticIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticIp)(nil))
}

func (o StaticIpArrayOutput) ToStaticIpArrayOutput() StaticIpArrayOutput {
	return o
}

func (o StaticIpArrayOutput) ToStaticIpArrayOutputWithContext(ctx context.Context) StaticIpArrayOutput {
	return o
}

func (o StaticIpArrayOutput) Index(i pulumi.IntInput) StaticIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticIp {
		return vs[0].([]StaticIp)[vs[1].(int)]
	}).(StaticIpOutput)
}

type StaticIpMapOutput struct{ *pulumi.OutputState }

func (StaticIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StaticIp)(nil))
}

func (o StaticIpMapOutput) ToStaticIpMapOutput() StaticIpMapOutput {
	return o
}

func (o StaticIpMapOutput) ToStaticIpMapOutputWithContext(ctx context.Context) StaticIpMapOutput {
	return o
}

func (o StaticIpMapOutput) MapIndex(k pulumi.StringInput) StaticIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StaticIp {
		return vs[0].(map[string]StaticIp)[vs[1].(string)]
	}).(StaticIpOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticIpOutput{})
	pulumi.RegisterOutputType(StaticIpPtrOutput{})
	pulumi.RegisterOutputType(StaticIpArrayOutput{})
	pulumi.RegisterOutputType(StaticIpMapOutput{})
}
