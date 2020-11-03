// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AutoScaling Scaling Policy resource.
//
// > **NOTE:** You may want to omit `desiredCapacity` attribute from attached `autoscaling.Group`
// when using autoscaling policies. It's good practice to pick either
// [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
// or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
// (policy-based) scaling.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/autoscaling"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bar, err := autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-east-1a"),
// 			},
// 			MaxSize:                pulumi.Int(5),
// 			MinSize:                pulumi.Int(2),
// 			HealthCheckGracePeriod: pulumi.Int(300),
// 			HealthCheckType:        pulumi.String("ELB"),
// 			ForceDelete:            pulumi.Bool(true),
// 			LaunchConfiguration:    pulumi.Any(aws_launch_configuration.Foo.Name),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = autoscaling.NewPolicy(ctx, "bat", &autoscaling.PolicyArgs{
// 			ScalingAdjustment:    pulumi.Int(4),
// 			AdjustmentType:       pulumi.String("ChangeInCapacity"),
// 			Cooldown:             pulumi.Int(300),
// 			AutoscalingGroupName: bar.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType pulumi.StringPtrOutput `pulumi:"adjustmentType"`
	// The ARN assigned by AWS to the scaling policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name of the autoscaling group.
	AutoscalingGroupName pulumi.StringOutput `pulumi:"autoscalingGroupName"`
	// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown pulumi.IntPtrOutput `pulumi:"cooldown"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup pulumi.IntPtrOutput `pulumi:"estimatedInstanceWarmup"`
	// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType pulumi.StringOutput `pulumi:"metricAggregationType"`
	// Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
	MinAdjustmentMagnitude pulumi.IntPtrOutput `pulumi:"minAdjustmentMagnitude"`
	// The name of the dimension.
	Name pulumi.StringOutput `pulumi:"name"`
	// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType pulumi.StringPtrOutput `pulumi:"policyType"`
	// The number of members by which to
	// scale, when the adjustment bounds are breached. A positive value scales
	// up. A negative value scales down.
	ScalingAdjustment pulumi.IntPtrOutput `pulumi:"scalingAdjustment"`
	// A set of adjustments that manage
	// group scaling. These have the following structure:
	StepAdjustments PolicyStepAdjustmentArrayOutput `pulumi:"stepAdjustments"`
	// A target tracking policy. These have the following structure:
	TargetTrackingConfiguration PolicyTargetTrackingConfigurationPtrOutput `pulumi:"targetTrackingConfiguration"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AutoscalingGroupName == nil {
		return nil, errors.New("invalid value for required argument 'AutoscalingGroupName'")
	}
	var resource Policy
	err := ctx.RegisterResource("aws:autoscaling/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("aws:autoscaling/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType *string `pulumi:"adjustmentType"`
	// The ARN assigned by AWS to the scaling policy.
	Arn *string `pulumi:"arn"`
	// The name of the autoscaling group.
	AutoscalingGroupName *string `pulumi:"autoscalingGroupName"`
	// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown *int `pulumi:"cooldown"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup *int `pulumi:"estimatedInstanceWarmup"`
	// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType *string `pulumi:"metricAggregationType"`
	// Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
	MinAdjustmentMagnitude *int `pulumi:"minAdjustmentMagnitude"`
	// The name of the dimension.
	Name *string `pulumi:"name"`
	// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType *string `pulumi:"policyType"`
	// The number of members by which to
	// scale, when the adjustment bounds are breached. A positive value scales
	// up. A negative value scales down.
	ScalingAdjustment *int `pulumi:"scalingAdjustment"`
	// A set of adjustments that manage
	// group scaling. These have the following structure:
	StepAdjustments []PolicyStepAdjustment `pulumi:"stepAdjustments"`
	// A target tracking policy. These have the following structure:
	TargetTrackingConfiguration *PolicyTargetTrackingConfiguration `pulumi:"targetTrackingConfiguration"`
}

type PolicyState struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType pulumi.StringPtrInput
	// The ARN assigned by AWS to the scaling policy.
	Arn pulumi.StringPtrInput
	// The name of the autoscaling group.
	AutoscalingGroupName pulumi.StringPtrInput
	// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown pulumi.IntPtrInput
	// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup pulumi.IntPtrInput
	// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType pulumi.StringPtrInput
	// Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
	MinAdjustmentMagnitude pulumi.IntPtrInput
	// The name of the dimension.
	Name pulumi.StringPtrInput
	// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType pulumi.StringPtrInput
	// The number of members by which to
	// scale, when the adjustment bounds are breached. A positive value scales
	// up. A negative value scales down.
	ScalingAdjustment pulumi.IntPtrInput
	// A set of adjustments that manage
	// group scaling. These have the following structure:
	StepAdjustments PolicyStepAdjustmentArrayInput
	// A target tracking policy. These have the following structure:
	TargetTrackingConfiguration PolicyTargetTrackingConfigurationPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType *string `pulumi:"adjustmentType"`
	// The name of the autoscaling group.
	AutoscalingGroupName string `pulumi:"autoscalingGroupName"`
	// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown *int `pulumi:"cooldown"`
	// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup *int `pulumi:"estimatedInstanceWarmup"`
	// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType *string `pulumi:"metricAggregationType"`
	// Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
	MinAdjustmentMagnitude *int `pulumi:"minAdjustmentMagnitude"`
	// The name of the dimension.
	Name *string `pulumi:"name"`
	// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType *string `pulumi:"policyType"`
	// The number of members by which to
	// scale, when the adjustment bounds are breached. A positive value scales
	// up. A negative value scales down.
	ScalingAdjustment *int `pulumi:"scalingAdjustment"`
	// A set of adjustments that manage
	// group scaling. These have the following structure:
	StepAdjustments []PolicyStepAdjustment `pulumi:"stepAdjustments"`
	// A target tracking policy. These have the following structure:
	TargetTrackingConfiguration *PolicyTargetTrackingConfiguration `pulumi:"targetTrackingConfiguration"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType pulumi.StringPtrInput
	// The name of the autoscaling group.
	AutoscalingGroupName pulumi.StringInput
	// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown pulumi.IntPtrInput
	// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup pulumi.IntPtrInput
	// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType pulumi.StringPtrInput
	// Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
	MinAdjustmentMagnitude pulumi.IntPtrInput
	// The name of the dimension.
	Name pulumi.StringPtrInput
	// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType pulumi.StringPtrInput
	// The number of members by which to
	// scale, when the adjustment bounds are breached. A positive value scales
	// up. A negative value scales down.
	ScalingAdjustment pulumi.IntPtrInput
	// A set of adjustments that manage
	// group scaling. These have the following structure:
	StepAdjustments PolicyStepAdjustmentArrayInput
	// A target tracking policy. These have the following structure:
	TargetTrackingConfiguration PolicyTargetTrackingConfigurationPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
