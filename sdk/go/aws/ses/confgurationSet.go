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

type ConfgurationSetPtrInput interface {
	pulumi.Input

	ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput
	ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput
}

func (ConfgurationSet) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfgurationSet)(nil)).Elem()
}

func (i ConfgurationSet) ToConfgurationSetOutput() ConfgurationSetOutput {
	return i.ToConfgurationSetOutputWithContext(context.Background())
}

func (i ConfgurationSet) ToConfgurationSetOutputWithContext(ctx context.Context) ConfgurationSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfgurationSetOutput)
}

func (i ConfgurationSet) ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput {
	return i.ToConfgurationSetPtrOutputWithContext(context.Background())
}

func (i ConfgurationSet) ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfgurationSetPtrOutput)
}

type ConfgurationSetOutput struct {
	*pulumi.OutputState
}

func (ConfgurationSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfgurationSetOutput)(nil)).Elem()
}

func (o ConfgurationSetOutput) ToConfgurationSetOutput() ConfgurationSetOutput {
	return o
}

func (o ConfgurationSetOutput) ToConfgurationSetOutputWithContext(ctx context.Context) ConfgurationSetOutput {
	return o
}

type ConfgurationSetPtrOutput struct {
	*pulumi.OutputState
}

func (ConfgurationSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfgurationSet)(nil)).Elem()
}

func (o ConfgurationSetPtrOutput) ToConfgurationSetPtrOutput() ConfgurationSetPtrOutput {
	return o
}

func (o ConfgurationSetPtrOutput) ToConfgurationSetPtrOutputWithContext(ctx context.Context) ConfgurationSetPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfgurationSetOutput{})
	pulumi.RegisterOutputType(ConfgurationSetPtrOutput{})
}
