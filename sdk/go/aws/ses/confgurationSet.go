// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: aws.ses.ConfgurationSet has been deprecated in favor of aws.ses.ConfigurationSet
type ConfgurationSet struct {
	pulumi.CustomResourceState

	Name pulumi.StringOutput `pulumi:"name"`
}

// NewConfgurationSet registers a new resource with the given unique name, arguments, and options.
func NewConfgurationSet(ctx *pulumi.Context,
	name string, args *ConfgurationSetArgs, opts ...pulumi.ResourceOption) (*ConfgurationSet, error) {
	if args == nil {
		args = &ConfgurationSetArgs{}
	}

	var resource ConfgurationSet
	err := ctx.RegisterResource("aws:ses/confgurationSet:ConfgurationSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfgurationSet gets an existing ConfgurationSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfgurationSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfgurationSetState, opts ...pulumi.ResourceOption) (*ConfgurationSet, error) {
	var resource ConfgurationSet
	err := ctx.ReadResource("aws:ses/confgurationSet:ConfgurationSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfgurationSet resources.
type confgurationSetState struct {
	Name *string `pulumi:"name"`
}

type ConfgurationSetState struct {
	Name pulumi.StringPtrInput
}

func (ConfgurationSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*confgurationSetState)(nil)).Elem()
}

type confgurationSetArgs struct {
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ConfgurationSet resource.
type ConfgurationSetArgs struct {
	Name pulumi.StringPtrInput
}

func (ConfgurationSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*confgurationSetArgs)(nil)).Elem()
}

type ConfgurationSetInput interface {
	pulumi.Input

	ToConfgurationSetOutput() ConfgurationSetOutput
	ToConfgurationSetOutputWithContext(ctx context.Context) ConfgurationSetOutput
}

func (*ConfgurationSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfgurationSet)(nil))
}

func (i *ConfgurationSet) ToConfgurationSetOutput() ConfgurationSetOutput {
	return i.ToConfgurationSetOutputWithContext(context.Background())
}

func (i *ConfgurationSet) ToConfgurationSetOutputWithContext(ctx context.Context) ConfgurationSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfgurationSetOutput)
}

func (i *ConfgurationSet) ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput {
	return i.ToConfgurationSetPtrOutputWithContext(context.Background())
}

func (i *ConfgurationSet) ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfgurationSetPtrOutput)
}

type ConfgurationSetPtrInput interface {
	pulumi.Input

	ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput
	ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput
}

type confgurationSetPtrType ConfgurationSetArgs

func (*confgurationSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfgurationSet)(nil))
}

func (i *confgurationSetPtrType) ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput {
	return i.ToConfgurationSetPtrOutputWithContext(context.Background())
}

func (i *confgurationSetPtrType) ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfgurationSetPtrOutput)
}

// ConfgurationSetArrayInput is an input type that accepts ConfgurationSetArray and ConfgurationSetArrayOutput values.
// You can construct a concrete instance of `ConfgurationSetArrayInput` via:
//
//          ConfgurationSetArray{ ConfgurationSetArgs{...} }
type ConfgurationSetArrayInput interface {
	pulumi.Input

	ToConfgurationSetArrayOutput() ConfgurationSetArrayOutput
	ToConfgurationSetArrayOutputWithContext(context.Context) ConfgurationSetArrayOutput
}

type ConfgurationSetArray []ConfgurationSetInput

func (ConfgurationSetArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ConfgurationSet)(nil))
}

func (i ConfgurationSetArray) ToConfgurationSetArrayOutput() ConfgurationSetArrayOutput {
	return i.ToConfgurationSetArrayOutputWithContext(context.Background())
}

func (i ConfgurationSetArray) ToConfgurationSetArrayOutputWithContext(ctx context.Context) ConfgurationSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfgurationSetArrayOutput)
}

// ConfgurationSetMapInput is an input type that accepts ConfgurationSetMap and ConfgurationSetMapOutput values.
// You can construct a concrete instance of `ConfgurationSetMapInput` via:
//
//          ConfgurationSetMap{ "key": ConfgurationSetArgs{...} }
type ConfgurationSetMapInput interface {
	pulumi.Input

	ToConfgurationSetMapOutput() ConfgurationSetMapOutput
	ToConfgurationSetMapOutputWithContext(context.Context) ConfgurationSetMapOutput
}

type ConfgurationSetMap map[string]ConfgurationSetInput

func (ConfgurationSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ConfgurationSet)(nil))
}

func (i ConfgurationSetMap) ToConfgurationSetMapOutput() ConfgurationSetMapOutput {
	return i.ToConfgurationSetMapOutputWithContext(context.Background())
}

func (i ConfgurationSetMap) ToConfgurationSetMapOutputWithContext(ctx context.Context) ConfgurationSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfgurationSetMapOutput)
}

type ConfgurationSetOutput struct {
	*pulumi.OutputState
}

func (ConfgurationSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfgurationSet)(nil))
}

func (o ConfgurationSetOutput) ToConfgurationSetOutput() ConfgurationSetOutput {
	return o
}

func (o ConfgurationSetOutput) ToConfgurationSetOutputWithContext(ctx context.Context) ConfgurationSetOutput {
	return o
}

func (o ConfgurationSetOutput) ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput {
	return o.ToConfgurationSetPtrOutputWithContext(context.Background())
}

func (o ConfgurationSetOutput) ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput {
	return o.ApplyT(func(v ConfgurationSet) *ConfgurationSet {
		return &v
	}).(ConfgurationSetPtrOutput)
}

type ConfgurationSetPtrOutput struct {
	*pulumi.OutputState
}

func (ConfgurationSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfgurationSet)(nil))
}

func (o ConfgurationSetPtrOutput) ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput {
	return o
}

func (o ConfgurationSetPtrOutput) ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput {
	return o
}

type ConfgurationSetArrayOutput struct{ *pulumi.OutputState }

func (ConfgurationSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfgurationSet)(nil))
}

func (o ConfgurationSetArrayOutput) ToConfgurationSetArrayOutput() ConfgurationSetArrayOutput {
	return o
}

func (o ConfgurationSetArrayOutput) ToConfgurationSetArrayOutputWithContext(ctx context.Context) ConfgurationSetArrayOutput {
	return o
}

func (o ConfgurationSetArrayOutput) Index(i pulumi.IntInput) ConfgurationSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfgurationSet {
		return vs[0].([]ConfgurationSet)[vs[1].(int)]
	}).(ConfgurationSetOutput)
}

type ConfgurationSetMapOutput struct{ *pulumi.OutputState }

func (ConfgurationSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConfgurationSet)(nil))
}

func (o ConfgurationSetMapOutput) ToConfgurationSetMapOutput() ConfgurationSetMapOutput {
	return o
}

func (o ConfgurationSetMapOutput) ToConfgurationSetMapOutputWithContext(ctx context.Context) ConfgurationSetMapOutput {
	return o
}

func (o ConfgurationSetMapOutput) MapIndex(k pulumi.StringInput) ConfgurationSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConfgurationSet {
		return vs[0].(map[string]ConfgurationSet)[vs[1].(string)]
	}).(ConfgurationSetOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfgurationSetOutput{})
	pulumi.RegisterOutputType(ConfgurationSetPtrOutput{})
	pulumi.RegisterOutputType(ConfgurationSetArrayOutput{})
	pulumi.RegisterOutputType(ConfgurationSetMapOutput{})
}
