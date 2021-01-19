// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscalingplans

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an AWS Auto Scaling scaling plan.
// More information can be found in the [AWS Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/plans/userguide/what-is-aws-auto-scaling.html).
//
// > **NOTE:** The AWS Auto Scaling service uses an AWS IAM service-linked role to manage predictive scaling of Amazon EC2 Auto Scaling groups. The service attempts to automatically create this role the first time a scaling plan with predictive scaling enabled is created.
// An [`iam.ServiceLinkedRole`](https://www.terraform.io/docs/providers/aws/r/iam_service_linked_role.html) resource can be used to manually manage this role.
// See the [AWS documentation](https://docs.aws.amazon.com/autoscaling/plans/userguide/aws-auto-scaling-service-linked-roles.html#create-service-linked-role-manual) for more details.
//
// ## Example Usage
//
// ## Import
//
// Auto Scaling scaling plans can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:autoscalingplans/scalingPlan:ScalingPlan example MyScale1
// ```
type ScalingPlan struct {
	pulumi.CustomResourceState

	// A CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource ScalingPlanApplicationSourceOutput `pulumi:"applicationSource"`
	// The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
	ScalingInstructions ScalingPlanScalingInstructionArrayOutput `pulumi:"scalingInstructions"`
	// The version number of the scaling plan. This value is always 1.
	ScalingPlanVersion pulumi.IntOutput `pulumi:"scalingPlanVersion"`
}

// NewScalingPlan registers a new resource with the given unique name, arguments, and options.
func NewScalingPlan(ctx *pulumi.Context,
	name string, args *ScalingPlanArgs, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationSource == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationSource'")
	}
	if args.ScalingInstructions == nil {
		return nil, errors.New("invalid value for required argument 'ScalingInstructions'")
	}
	var resource ScalingPlan
	err := ctx.RegisterResource("aws:autoscalingplans/scalingPlan:ScalingPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingPlan gets an existing ScalingPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingPlanState, opts ...pulumi.ResourceOption) (*ScalingPlan, error) {
	var resource ScalingPlan
	err := ctx.ReadResource("aws:autoscalingplans/scalingPlan:ScalingPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingPlan resources.
type scalingPlanState struct {
	// A CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource *ScalingPlanApplicationSource `pulumi:"applicationSource"`
	// The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name *string `pulumi:"name"`
	// The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
	ScalingInstructions []ScalingPlanScalingInstruction `pulumi:"scalingInstructions"`
	// The version number of the scaling plan. This value is always 1.
	ScalingPlanVersion *int `pulumi:"scalingPlanVersion"`
}

type ScalingPlanState struct {
	// A CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource ScalingPlanApplicationSourcePtrInput
	// The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name pulumi.StringPtrInput
	// The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
	ScalingInstructions ScalingPlanScalingInstructionArrayInput
	// The version number of the scaling plan. This value is always 1.
	ScalingPlanVersion pulumi.IntPtrInput
}

func (ScalingPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanState)(nil)).Elem()
}

type scalingPlanArgs struct {
	// A CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource ScalingPlanApplicationSource `pulumi:"applicationSource"`
	// The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name *string `pulumi:"name"`
	// The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
	ScalingInstructions []ScalingPlanScalingInstruction `pulumi:"scalingInstructions"`
}

// The set of arguments for constructing a ScalingPlan resource.
type ScalingPlanArgs struct {
	// A CloudFormation stack or set of tags. You can create one scaling plan per application source.
	ApplicationSource ScalingPlanApplicationSourceInput
	// The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
	Name pulumi.StringPtrInput
	// The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
	ScalingInstructions ScalingPlanScalingInstructionArrayInput
}

func (ScalingPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingPlanArgs)(nil)).Elem()
}

type ScalingPlanInput interface {
	pulumi.Input

	ToScalingPlanOutput() ScalingPlanOutput
	ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput
}

func (*ScalingPlan) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPlan)(nil))
}

func (i *ScalingPlan) ToScalingPlanOutput() ScalingPlanOutput {
	return i.ToScalingPlanOutputWithContext(context.Background())
}

func (i *ScalingPlan) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanOutput)
}

func (i *ScalingPlan) ToScalingPlanPtrOutput() ScalingPlanPtrOutput {
	return i.ToScalingPlanPtrOutputWithContext(context.Background())
}

func (i *ScalingPlan) ToScalingPlanPtrOutputWithContext(ctx context.Context) ScalingPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanPtrOutput)
}

type ScalingPlanPtrInput interface {
	pulumi.Input

	ToScalingPlanPtrOutput() ScalingPlanPtrOutput
	ToScalingPlanPtrOutputWithContext(ctx context.Context) ScalingPlanPtrOutput
}

type scalingPlanPtrType ScalingPlanArgs

func (*scalingPlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlan)(nil))
}

func (i *scalingPlanPtrType) ToScalingPlanPtrOutput() ScalingPlanPtrOutput {
	return i.ToScalingPlanPtrOutputWithContext(context.Background())
}

func (i *scalingPlanPtrType) ToScalingPlanPtrOutputWithContext(ctx context.Context) ScalingPlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanPtrOutput)
}

// ScalingPlanArrayInput is an input type that accepts ScalingPlanArray and ScalingPlanArrayOutput values.
// You can construct a concrete instance of `ScalingPlanArrayInput` via:
//
//          ScalingPlanArray{ ScalingPlanArgs{...} }
type ScalingPlanArrayInput interface {
	pulumi.Input

	ToScalingPlanArrayOutput() ScalingPlanArrayOutput
	ToScalingPlanArrayOutputWithContext(context.Context) ScalingPlanArrayOutput
}

type ScalingPlanArray []ScalingPlanInput

func (ScalingPlanArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ScalingPlan)(nil))
}

func (i ScalingPlanArray) ToScalingPlanArrayOutput() ScalingPlanArrayOutput {
	return i.ToScalingPlanArrayOutputWithContext(context.Background())
}

func (i ScalingPlanArray) ToScalingPlanArrayOutputWithContext(ctx context.Context) ScalingPlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanArrayOutput)
}

// ScalingPlanMapInput is an input type that accepts ScalingPlanMap and ScalingPlanMapOutput values.
// You can construct a concrete instance of `ScalingPlanMapInput` via:
//
//          ScalingPlanMap{ "key": ScalingPlanArgs{...} }
type ScalingPlanMapInput interface {
	pulumi.Input

	ToScalingPlanMapOutput() ScalingPlanMapOutput
	ToScalingPlanMapOutputWithContext(context.Context) ScalingPlanMapOutput
}

type ScalingPlanMap map[string]ScalingPlanInput

func (ScalingPlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ScalingPlan)(nil))
}

func (i ScalingPlanMap) ToScalingPlanMapOutput() ScalingPlanMapOutput {
	return i.ToScalingPlanMapOutputWithContext(context.Background())
}

func (i ScalingPlanMap) ToScalingPlanMapOutputWithContext(ctx context.Context) ScalingPlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingPlanMapOutput)
}

type ScalingPlanOutput struct {
	*pulumi.OutputState
}

func (ScalingPlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingPlan)(nil))
}

func (o ScalingPlanOutput) ToScalingPlanOutput() ScalingPlanOutput {
	return o
}

func (o ScalingPlanOutput) ToScalingPlanOutputWithContext(ctx context.Context) ScalingPlanOutput {
	return o
}

func (o ScalingPlanOutput) ToScalingPlanPtrOutput() ScalingPlanPtrOutput {
	return o.ToScalingPlanPtrOutputWithContext(context.Background())
}

func (o ScalingPlanOutput) ToScalingPlanPtrOutputWithContext(ctx context.Context) ScalingPlanPtrOutput {
	return o.ApplyT(func(v ScalingPlan) *ScalingPlan {
		return &v
	}).(ScalingPlanPtrOutput)
}

type ScalingPlanPtrOutput struct {
	*pulumi.OutputState
}

func (ScalingPlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingPlan)(nil))
}

func (o ScalingPlanPtrOutput) ToScalingPlanPtrOutput() ScalingPlanPtrOutput {
	return o
}

func (o ScalingPlanPtrOutput) ToScalingPlanPtrOutputWithContext(ctx context.Context) ScalingPlanPtrOutput {
	return o
}

type ScalingPlanArrayOutput struct{ *pulumi.OutputState }

func (ScalingPlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScalingPlan)(nil))
}

func (o ScalingPlanArrayOutput) ToScalingPlanArrayOutput() ScalingPlanArrayOutput {
	return o
}

func (o ScalingPlanArrayOutput) ToScalingPlanArrayOutputWithContext(ctx context.Context) ScalingPlanArrayOutput {
	return o
}

func (o ScalingPlanArrayOutput) Index(i pulumi.IntInput) ScalingPlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScalingPlan {
		return vs[0].([]ScalingPlan)[vs[1].(int)]
	}).(ScalingPlanOutput)
}

type ScalingPlanMapOutput struct{ *pulumi.OutputState }

func (ScalingPlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ScalingPlan)(nil))
}

func (o ScalingPlanMapOutput) ToScalingPlanMapOutput() ScalingPlanMapOutput {
	return o
}

func (o ScalingPlanMapOutput) ToScalingPlanMapOutputWithContext(ctx context.Context) ScalingPlanMapOutput {
	return o
}

func (o ScalingPlanMapOutput) MapIndex(k pulumi.StringInput) ScalingPlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ScalingPlan {
		return vs[0].(map[string]ScalingPlan)[vs[1].(string)]
	}).(ScalingPlanOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingPlanOutput{})
	pulumi.RegisterOutputType(ScalingPlanPtrOutput{})
	pulumi.RegisterOutputType(ScalingPlanArrayOutput{})
	pulumi.RegisterOutputType(ScalingPlanMapOutput{})
}
