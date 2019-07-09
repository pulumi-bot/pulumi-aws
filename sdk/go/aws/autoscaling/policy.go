// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AutoScaling Scaling Policy resource.
// 
// > **NOTE:** You may want to omit `desired_capacity` attribute from attached `aws_autoscaling_group`
// when using autoscaling policies. It's good practice to pick either
// [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
// or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
// (policy-based) scaling.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/autoscaling_policy.html.markdown.
type Policy struct {
	s *pulumi.ResourceState
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOpt) (*Policy, error) {
	if args == nil || args.AutoscalingGroupName == nil {
		return nil, errors.New("missing required argument 'AutoscalingGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["adjustmentType"] = nil
		inputs["autoscalingGroupName"] = nil
		inputs["cooldown"] = nil
		inputs["estimatedInstanceWarmup"] = nil
		inputs["metricAggregationType"] = nil
		inputs["minAdjustmentMagnitude"] = nil
		inputs["name"] = nil
		inputs["policyType"] = nil
		inputs["scalingAdjustment"] = nil
		inputs["stepAdjustments"] = nil
		inputs["targetTrackingConfiguration"] = nil
	} else {
		inputs["adjustmentType"] = args.AdjustmentType
		inputs["autoscalingGroupName"] = args.AutoscalingGroupName
		inputs["cooldown"] = args.Cooldown
		inputs["estimatedInstanceWarmup"] = args.EstimatedInstanceWarmup
		inputs["metricAggregationType"] = args.MetricAggregationType
		inputs["minAdjustmentMagnitude"] = args.MinAdjustmentMagnitude
		inputs["name"] = args.Name
		inputs["policyType"] = args.PolicyType
		inputs["scalingAdjustment"] = args.ScalingAdjustment
		inputs["stepAdjustments"] = args.StepAdjustments
		inputs["targetTrackingConfiguration"] = args.TargetTrackingConfiguration
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:autoscaling/policy:Policy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyState, opts ...pulumi.ResourceOpt) (*Policy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["adjustmentType"] = state.AdjustmentType
		inputs["arn"] = state.Arn
		inputs["autoscalingGroupName"] = state.AutoscalingGroupName
		inputs["cooldown"] = state.Cooldown
		inputs["estimatedInstanceWarmup"] = state.EstimatedInstanceWarmup
		inputs["metricAggregationType"] = state.MetricAggregationType
		inputs["minAdjustmentMagnitude"] = state.MinAdjustmentMagnitude
		inputs["name"] = state.Name
		inputs["policyType"] = state.PolicyType
		inputs["scalingAdjustment"] = state.ScalingAdjustment
		inputs["stepAdjustments"] = state.StepAdjustments
		inputs["targetTrackingConfiguration"] = state.TargetTrackingConfiguration
	}
	s, err := ctx.ReadResource("aws:autoscaling/policy:Policy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Policy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Policy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
func (r *Policy) AdjustmentType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["adjustmentType"])
}

// The ARN assigned by AWS to the scaling policy.
func (r *Policy) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The name of the autoscaling group.
func (r *Policy) AutoscalingGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["autoscalingGroupName"])
}

// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
func (r *Policy) Cooldown() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["cooldown"])
}

// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
func (r *Policy) EstimatedInstanceWarmup() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["estimatedInstanceWarmup"])
}

// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
func (r *Policy) MetricAggregationType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["metricAggregationType"])
}

func (r *Policy) MinAdjustmentMagnitude() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["minAdjustmentMagnitude"])
}

// The name of the dimension.
func (r *Policy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
func (r *Policy) PolicyType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policyType"])
}

// The number of members by which to
// scale, when the adjustment bounds are breached. A positive value scales
// up. A negative value scales down.
func (r *Policy) ScalingAdjustment() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["scalingAdjustment"])
}

func (r *Policy) StepAdjustments() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["stepAdjustments"])
}

// A target tracking policy. These have the following structure:
func (r *Policy) TargetTrackingConfiguration() *pulumi.Output {
	return r.s.State["targetTrackingConfiguration"]
}

// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType interface{}
	// The ARN assigned by AWS to the scaling policy.
	Arn interface{}
	// The name of the autoscaling group.
	AutoscalingGroupName interface{}
	// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown interface{}
	// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup interface{}
	// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType interface{}
	MinAdjustmentMagnitude interface{}
	// The name of the dimension.
	Name interface{}
	// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType interface{}
	// The number of members by which to
	// scale, when the adjustment bounds are breached. A positive value scales
	// up. A negative value scales down.
	ScalingAdjustment interface{}
	StepAdjustments interface{}
	// A target tracking policy. These have the following structure:
	TargetTrackingConfiguration interface{}
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
	AdjustmentType interface{}
	// The name of the autoscaling group.
	AutoscalingGroupName interface{}
	// The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
	Cooldown interface{}
	// The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
	EstimatedInstanceWarmup interface{}
	// The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
	MetricAggregationType interface{}
	MinAdjustmentMagnitude interface{}
	// The name of the dimension.
	Name interface{}
	// The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
	PolicyType interface{}
	// The number of members by which to
	// scale, when the adjustment bounds are breached. A positive value scales
	// up. A negative value scales down.
	ScalingAdjustment interface{}
	StepAdjustments interface{}
	// A target tracking policy. These have the following structure:
	TargetTrackingConfiguration interface{}
}
