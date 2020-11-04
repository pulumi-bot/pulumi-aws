// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AutoScaling Schedule resource.
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
// 		foobarGroup, err := autoscaling.NewGroup(ctx, "foobarGroup", &autoscaling.GroupArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-west-2a"),
// 			},
// 			MaxSize:                pulumi.Int(1),
// 			MinSize:                pulumi.Int(1),
// 			HealthCheckGracePeriod: pulumi.Int(300),
// 			HealthCheckType:        pulumi.String("ELB"),
// 			ForceDelete:            pulumi.Bool(true),
// 			TerminationPolicies: pulumi.StringArray{
// 				pulumi.String("OldestInstance"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = autoscaling.NewSchedule(ctx, "foobarSchedule", &autoscaling.ScheduleArgs{
// 			ScheduledActionName:  pulumi.String("foobar"),
// 			MinSize:              pulumi.Int(0),
// 			MaxSize:              pulumi.Int(1),
// 			DesiredCapacity:      pulumi.Int(0),
// 			StartTime:            pulumi.String("2016-12-11T18:00:00Z"),
// 			EndTime:              pulumi.String("2016-12-12T06:00:00Z"),
// 			AutoscalingGroupName: foobarGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Schedule struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS to the autoscaling schedule.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name or Amazon Resource Name (ARN) of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringOutput `pulumi:"autoscalingGroupName"`
	// The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
	DesiredCapacity pulumi.IntOutput `pulumi:"desiredCapacity"`
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The maximum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the maximum size at the scheduled time.
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// The minimum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the minimum size at the scheduled time.
	MinSize pulumi.IntOutput `pulumi:"minSize"`
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
	Recurrence pulumi.StringOutput `pulumi:"recurrence"`
	// The name of this scaling action.
	ScheduledActionName pulumi.StringOutput `pulumi:"scheduledActionName"`
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewSchedule registers a new resource with the given unique name, arguments, and options.
func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.AutoscalingGroupName == nil {
		return nil, errors.New("invalid value for required argument 'AutoscalingGroupName'")
	}
	if args.ScheduledActionName == nil {
		return nil, errors.New("invalid value for required argument 'ScheduledActionName'")
	}
	var resource Schedule
	err := ctx.RegisterResource("aws:autoscaling/schedule:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSchedule gets an existing Schedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("aws:autoscaling/schedule:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Schedule resources.
type scheduleState struct {
	// The ARN assigned by AWS to the autoscaling schedule.
	Arn *string `pulumi:"arn"`
	// The name or Amazon Resource Name (ARN) of the Auto Scaling group.
	AutoscalingGroupName *string `pulumi:"autoscalingGroupName"`
	// The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	EndTime *string `pulumi:"endTime"`
	// The maximum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the maximum size at the scheduled time.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the minimum size at the scheduled time.
	MinSize *int `pulumi:"minSize"`
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
	Recurrence *string `pulumi:"recurrence"`
	// The name of this scaling action.
	ScheduledActionName *string `pulumi:"scheduledActionName"`
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	StartTime *string `pulumi:"startTime"`
}

type ScheduleState struct {
	// The ARN assigned by AWS to the autoscaling schedule.
	Arn pulumi.StringPtrInput
	// The name or Amazon Resource Name (ARN) of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringPtrInput
	// The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
	DesiredCapacity pulumi.IntPtrInput
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	EndTime pulumi.StringPtrInput
	// The maximum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the maximum size at the scheduled time.
	MaxSize pulumi.IntPtrInput
	// The minimum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the minimum size at the scheduled time.
	MinSize pulumi.IntPtrInput
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
	Recurrence pulumi.StringPtrInput
	// The name of this scaling action.
	ScheduledActionName pulumi.StringPtrInput
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	StartTime pulumi.StringPtrInput
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	// The name or Amazon Resource Name (ARN) of the Auto Scaling group.
	AutoscalingGroupName string `pulumi:"autoscalingGroupName"`
	// The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	EndTime *string `pulumi:"endTime"`
	// The maximum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the maximum size at the scheduled time.
	MaxSize *int `pulumi:"maxSize"`
	// The minimum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the minimum size at the scheduled time.
	MinSize *int `pulumi:"minSize"`
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
	Recurrence *string `pulumi:"recurrence"`
	// The name of this scaling action.
	ScheduledActionName string `pulumi:"scheduledActionName"`
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	StartTime *string `pulumi:"startTime"`
}

// The set of arguments for constructing a Schedule resource.
type ScheduleArgs struct {
	// The name or Amazon Resource Name (ARN) of the Auto Scaling group.
	AutoscalingGroupName pulumi.StringInput
	// The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
	DesiredCapacity pulumi.IntPtrInput
	// The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	EndTime pulumi.StringPtrInput
	// The maximum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the maximum size at the scheduled time.
	MaxSize pulumi.IntPtrInput
	// The minimum size for the Auto Scaling group. Default 0.
	// Set to -1 if you don't want to change the minimum size at the scheduled time.
	MinSize pulumi.IntPtrInput
	// The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
	Recurrence pulumi.StringPtrInput
	// The name of this scaling action.
	ScheduledActionName pulumi.StringInput
	// The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
	// If you try to schedule your action in the past, Auto Scaling returns an error message.
	StartTime pulumi.StringPtrInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}
