// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IoT policy attachment.
type PolicyAttachment struct {
	pulumi.CustomResourceState

	// The name of the policy to attach.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The identity to which the policy is attached.
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAttachment(ctx *pulumi.Context,
	name string, args *PolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	var resource PolicyAttachment
	err := ctx.RegisterResource("aws:iot/policyAttachment:PolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAttachment gets an existing PolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAttachmentState, opts ...pulumi.ResourceOption) (*PolicyAttachment, error) {
	var resource PolicyAttachment
	err := ctx.ReadResource("aws:iot/policyAttachment:PolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAttachment resources.
type policyAttachmentState struct {
	// The name of the policy to attach.
	Policy interface{} `pulumi:"policy"`
	// The identity to which the policy is attached.
	Target *string `pulumi:"target"`
}

type PolicyAttachmentState struct {
	// The name of the policy to attach.
	Policy pulumi.Input
	// The identity to which the policy is attached.
	Target pulumi.StringPtrInput
}

func (PolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentState)(nil)).Elem()
}

type policyAttachmentArgs struct {
	// The name of the policy to attach.
	Policy interface{} `pulumi:"policy"`
	// The identity to which the policy is attached.
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a PolicyAttachment resource.
type PolicyAttachmentArgs struct {
	// The name of the policy to attach.
	Policy pulumi.Input
	// The identity to which the policy is attached.
	Target pulumi.StringInput
}

func (PolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAttachmentArgs)(nil)).Elem()
}

type PolicyAttachmentInput interface {
	pulumi.Input

	ToPolicyAttachmentOutput() PolicyAttachmentOutput
	ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput
}

func (*PolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAttachment)(nil))
}

func (i *PolicyAttachment) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return i.ToPolicyAttachmentOutputWithContext(context.Background())
}

func (i *PolicyAttachment) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentOutput)
}

func (i *PolicyAttachment) ToPolicyAttachmentPtrOutput() PolicyAttachmentPtrOutput {
	return i.ToPolicyAttachmentPtrOutputWithContext(context.Background())
}

func (i *PolicyAttachment) ToPolicyAttachmentPtrOutputWithContext(ctx context.Context) PolicyAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentPtrOutput)
}

type PolicyAttachmentPtrInput interface {
	pulumi.Input

	ToPolicyAttachmentPtrOutput() PolicyAttachmentPtrOutput
	ToPolicyAttachmentPtrOutputWithContext(ctx context.Context) PolicyAttachmentPtrOutput
}

type policyAttachmentPtrType PolicyAttachmentArgs

func (*policyAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAttachment)(nil))
}

func (i *policyAttachmentPtrType) ToPolicyAttachmentPtrOutput() PolicyAttachmentPtrOutput {
	return i.ToPolicyAttachmentPtrOutputWithContext(context.Background())
}

func (i *policyAttachmentPtrType) ToPolicyAttachmentPtrOutputWithContext(ctx context.Context) PolicyAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentPtrOutput)
}

// PolicyAttachmentArrayInput is an input type that accepts PolicyAttachmentArray and PolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `PolicyAttachmentArrayInput` via:
//
//          PolicyAttachmentArray{ PolicyAttachmentArgs{...} }
type PolicyAttachmentArrayInput interface {
	pulumi.Input

	ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput
	ToPolicyAttachmentArrayOutputWithContext(context.Context) PolicyAttachmentArrayOutput
}

type PolicyAttachmentArray []PolicyAttachmentInput

func (PolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PolicyAttachment)(nil))
}

func (i PolicyAttachmentArray) ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput {
	return i.ToPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i PolicyAttachmentArray) ToPolicyAttachmentArrayOutputWithContext(ctx context.Context) PolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentArrayOutput)
}

// PolicyAttachmentMapInput is an input type that accepts PolicyAttachmentMap and PolicyAttachmentMapOutput values.
// You can construct a concrete instance of `PolicyAttachmentMapInput` via:
//
//          PolicyAttachmentMap{ "key": PolicyAttachmentArgs{...} }
type PolicyAttachmentMapInput interface {
	pulumi.Input

	ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput
	ToPolicyAttachmentMapOutputWithContext(context.Context) PolicyAttachmentMapOutput
}

type PolicyAttachmentMap map[string]PolicyAttachmentInput

func (PolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PolicyAttachment)(nil))
}

func (i PolicyAttachmentMap) ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput {
	return i.ToPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i PolicyAttachmentMap) ToPolicyAttachmentMapOutputWithContext(ctx context.Context) PolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyAttachmentMapOutput)
}

type PolicyAttachmentOutput struct {
	*pulumi.OutputState
}

func (PolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyAttachment)(nil))
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutput() PolicyAttachmentOutput {
	return o
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentOutputWithContext(ctx context.Context) PolicyAttachmentOutput {
	return o
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentPtrOutput() PolicyAttachmentPtrOutput {
	return o.ToPolicyAttachmentPtrOutputWithContext(context.Background())
}

func (o PolicyAttachmentOutput) ToPolicyAttachmentPtrOutputWithContext(ctx context.Context) PolicyAttachmentPtrOutput {
	return o.ApplyT(func(v PolicyAttachment) *PolicyAttachment {
		return &v
	}).(PolicyAttachmentPtrOutput)
}

type PolicyAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyAttachment)(nil))
}

func (o PolicyAttachmentPtrOutput) ToPolicyAttachmentPtrOutput() PolicyAttachmentPtrOutput {
	return o
}

func (o PolicyAttachmentPtrOutput) ToPolicyAttachmentPtrOutputWithContext(ctx context.Context) PolicyAttachmentPtrOutput {
	return o
}

type PolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PolicyAttachment)(nil))
}

func (o PolicyAttachmentArrayOutput) ToPolicyAttachmentArrayOutput() PolicyAttachmentArrayOutput {
	return o
}

func (o PolicyAttachmentArrayOutput) ToPolicyAttachmentArrayOutputWithContext(ctx context.Context) PolicyAttachmentArrayOutput {
	return o
}

func (o PolicyAttachmentArrayOutput) Index(i pulumi.IntInput) PolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PolicyAttachment {
		return vs[0].([]PolicyAttachment)[vs[1].(int)]
	}).(PolicyAttachmentOutput)
}

type PolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (PolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PolicyAttachment)(nil))
}

func (o PolicyAttachmentMapOutput) ToPolicyAttachmentMapOutput() PolicyAttachmentMapOutput {
	return o
}

func (o PolicyAttachmentMapOutput) ToPolicyAttachmentMapOutputWithContext(ctx context.Context) PolicyAttachmentMapOutput {
	return o
}

func (o PolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) PolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PolicyAttachment {
		return vs[0].(map[string]PolicyAttachment)[vs[1].(string)]
	}).(PolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyAttachmentOutput{})
	pulumi.RegisterOutputType(PolicyAttachmentPtrOutput{})
	pulumi.RegisterOutputType(PolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(PolicyAttachmentMapOutput{})
}
