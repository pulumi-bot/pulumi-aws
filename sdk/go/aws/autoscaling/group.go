// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// AutoScaling Groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:autoscaling/group:Group web web-asg
// ```
type Group struct {
	pulumi.CustomResourceState

	// The ARN for this AutoScaling Group
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A list of one or more availability zones for the group. Used for EC2-Classic and default subnets when not specified with `vpcZoneIdentifier` argument. Conflicts with `vpcZoneIdentifier`.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	DefaultCooldown pulumi.IntOutput `pulumi:"defaultCooldown"`
	// The number of Amazon EC2 instances that
	// should be running in the group. (See also Waiting for
	// Capacity below.)
	DesiredCapacity pulumi.IntOutput `pulumi:"desiredCapacity"`
	// A list of metrics to collect. The allowed values are `GroupDesiredCapacity`, `GroupInServiceCapacity`, `GroupPendingCapacity`, `GroupMinSize`, `GroupMaxSize`, `GroupInServiceInstances`, `GroupPendingInstances`, `GroupStandbyInstances`, `GroupStandbyCapacity`, `GroupTerminatingCapacity`, `GroupTerminatingInstances`, `GroupTotalCapacity`, `GroupTotalInstances`.
	EnabledMetrics pulumi.StringArrayOutput `pulumi:"enabledMetrics"`
	// Allows deleting the autoscaling group without waiting
	// for all instances in the pool to terminate.  You can force an autoscaling group to delete
	// even if it's in the process of scaling a resource. Normally, this provider
	// drains all the instances before deleting the group.  This bypasses that
	// behavior and potentially leaves resources dangling.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// Time (in seconds) after instance comes into service before checking health.
	HealthCheckGracePeriod pulumi.IntPtrOutput `pulumi:"healthCheckGracePeriod"`
	// "EC2" or "ELB". Controls how health checking is done.
	HealthCheckType pulumi.StringOutput `pulumi:"healthCheckType"`
	// One or more
	// [Lifecycle Hooks](http://docs.aws.amazon.com/autoscaling/latest/userguide/lifecycle-hooks.html)
	// to attach to the autoscaling group **before** instances are launched. The
	// syntax is exactly the same as the separate
	// `autoscaling.LifecycleHook`
	// resource, without the `autoscalingGroupName` attribute. Please note that this will only work when creating
	// a new autoscaling group. For all other use-cases, please use `autoscaling.LifecycleHook` resource.
	InitialLifecycleHooks GroupInitialLifecycleHookArrayOutput `pulumi:"initialLifecycleHooks"`
	// The name of the launch configuration to use.
	LaunchConfiguration pulumi.StringPtrOutput `pulumi:"launchConfiguration"`
	// Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
	LaunchTemplate GroupLaunchTemplatePtrOutput `pulumi:"launchTemplate"`
	// A list of elastic load balancer names to add to the autoscaling
	// group names. Only valid for classic load balancers. For ALBs, use `targetGroupArns` instead.
	LoadBalancers pulumi.StringArrayOutput `pulumi:"loadBalancers"`
	// The maximum amount of time, in seconds, that an instance can be in service, values must be either equal to 0 or between 604800 and 31536000 seconds.
	MaxInstanceLifetime pulumi.IntPtrOutput `pulumi:"maxInstanceLifetime"`
	// The maximum size of the auto scale group.
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// The granularity to associate with the metrics to collect. The only valid value is `1Minute`. Default is `1Minute`.
	MetricsGranularity pulumi.StringPtrOutput `pulumi:"metricsGranularity"`
	// Setting this causes this provider to wait for
	// this number of instances from this autoscaling group to show up healthy in the
	// ELB only on creation. Updates will not wait on ELB instance number changes.
	// (See also Waiting for Capacity below.)
	MinElbCapacity pulumi.IntPtrOutput `pulumi:"minElbCapacity"`
	// The minimum size of the auto scale group.
	// (See also Waiting for Capacity below.)
	MinSize pulumi.IntOutput `pulumi:"minSize"`
	// Configuration block containing settings to define launch targets for Auto Scaling groups. Defined below.
	MixedInstancesPolicy GroupMixedInstancesPolicyPtrOutput `pulumi:"mixedInstancesPolicy"`
	// The name of the auto scaling group. By default generated by this provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The name of the placement group into which you'll launch your instances, if any.
	PlacementGroup pulumi.StringPtrOutput `pulumi:"placementGroup"`
	// Allows setting instance protection. The
	// autoscaling group will not select instances with this setting for termination
	// during scale in events.
	ProtectFromScaleIn pulumi.BoolPtrOutput `pulumi:"protectFromScaleIn"`
	// The ARN of the service-linked role that the ASG will use to call other AWS services
	ServiceLinkedRoleArn pulumi.StringOutput `pulumi:"serviceLinkedRoleArn"`
	// A list of processes to suspend for the AutoScaling Group. The allowed values are `Launch`, `Terminate`, `HealthCheck`, `ReplaceUnhealthy`, `AZRebalance`, `AlarmNotification`, `ScheduledActions`, `AddToLoadBalancer`.
	// Note that if you suspend either the `Launch` or `Terminate` process types, it can prevent your autoscaling group from functioning properly.
	SuspendedProcesses pulumi.StringArrayOutput `pulumi:"suspendedProcesses"`
	// Configuration block(s) containing resource tags. Conflicts with `tags`. Documented below.
	Tags GroupTagArrayOutput `pulumi:"tags"`
	// Set of maps containing resource tags. Conflicts with `tag`. Documented below.
	TagsCollection pulumi.StringMapArrayOutput `pulumi:"tagsCollection"`
	// A set of `alb.TargetGroup` ARNs, for use with Application or Network Load Balancing.
	TargetGroupArns pulumi.StringArrayOutput `pulumi:"targetGroupArns"`
	// A list of policies to decide how the instances in the auto scale group should be terminated. The allowed values are `OldestInstance`, `NewestInstance`, `OldestLaunchConfiguration`, `ClosestToNextInstanceHour`, `OldestLaunchTemplate`, `AllocationStrategy`, `Default`.
	TerminationPolicies pulumi.StringArrayOutput `pulumi:"terminationPolicies"`
	// A list of subnet IDs to launch resources in. Subnets automatically determine which availability zones the group will reside. Conflicts with `availabilityZones`.
	VpcZoneIdentifiers pulumi.StringArrayOutput `pulumi:"vpcZoneIdentifiers"`
	// A maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for ASG instances to be healthy before timing out.  (See also Waiting
	// for Capacity below.) Setting this to "0" causes
	// this provider to skip all Capacity Waiting behavior.
	WaitForCapacityTimeout pulumi.StringPtrOutput `pulumi:"waitForCapacityTimeout"`
	// Setting this will cause this provider to wait
	// for exactly this number of healthy instances from this autoscaling group in
	// all attached load balancers on both create and update operations. (Takes
	// precedence over `minElbCapacity` behavior.)
	// (See also Waiting for Capacity below.)
	WaitForElbCapacity pulumi.IntPtrOutput `pulumi:"waitForElbCapacity"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.MaxSize == nil {
		return nil, errors.New("missing required argument 'MaxSize'")
	}
	if args == nil || args.MinSize == nil {
		return nil, errors.New("missing required argument 'MinSize'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("aws:autoscaling/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws:autoscaling/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The ARN for this AutoScaling Group
	Arn *string `pulumi:"arn"`
	// A list of one or more availability zones for the group. Used for EC2-Classic and default subnets when not specified with `vpcZoneIdentifier` argument. Conflicts with `vpcZoneIdentifier`.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// The number of Amazon EC2 instances that
	// should be running in the group. (See also Waiting for
	// Capacity below.)
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// A list of metrics to collect. The allowed values are `GroupDesiredCapacity`, `GroupInServiceCapacity`, `GroupPendingCapacity`, `GroupMinSize`, `GroupMaxSize`, `GroupInServiceInstances`, `GroupPendingInstances`, `GroupStandbyInstances`, `GroupStandbyCapacity`, `GroupTerminatingCapacity`, `GroupTerminatingInstances`, `GroupTotalCapacity`, `GroupTotalInstances`.
	EnabledMetrics []string `pulumi:"enabledMetrics"`
	// Allows deleting the autoscaling group without waiting
	// for all instances in the pool to terminate.  You can force an autoscaling group to delete
	// even if it's in the process of scaling a resource. Normally, this provider
	// drains all the instances before deleting the group.  This bypasses that
	// behavior and potentially leaves resources dangling.
	ForceDelete *bool `pulumi:"forceDelete"`
	// Time (in seconds) after instance comes into service before checking health.
	HealthCheckGracePeriod *int `pulumi:"healthCheckGracePeriod"`
	// "EC2" or "ELB". Controls how health checking is done.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// One or more
	// [Lifecycle Hooks](http://docs.aws.amazon.com/autoscaling/latest/userguide/lifecycle-hooks.html)
	// to attach to the autoscaling group **before** instances are launched. The
	// syntax is exactly the same as the separate
	// `autoscaling.LifecycleHook`
	// resource, without the `autoscalingGroupName` attribute. Please note that this will only work when creating
	// a new autoscaling group. For all other use-cases, please use `autoscaling.LifecycleHook` resource.
	InitialLifecycleHooks []GroupInitialLifecycleHook `pulumi:"initialLifecycleHooks"`
	// The name of the launch configuration to use.
	LaunchConfiguration *string `pulumi:"launchConfiguration"`
	// Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
	LaunchTemplate *GroupLaunchTemplate `pulumi:"launchTemplate"`
	// A list of elastic load balancer names to add to the autoscaling
	// group names. Only valid for classic load balancers. For ALBs, use `targetGroupArns` instead.
	LoadBalancers []string `pulumi:"loadBalancers"`
	// The maximum amount of time, in seconds, that an instance can be in service, values must be either equal to 0 or between 604800 and 31536000 seconds.
	MaxInstanceLifetime *int `pulumi:"maxInstanceLifetime"`
	// The maximum size of the auto scale group.
	MaxSize *int `pulumi:"maxSize"`
	// The granularity to associate with the metrics to collect. The only valid value is `1Minute`. Default is `1Minute`.
	MetricsGranularity *string `pulumi:"metricsGranularity"`
	// Setting this causes this provider to wait for
	// this number of instances from this autoscaling group to show up healthy in the
	// ELB only on creation. Updates will not wait on ELB instance number changes.
	// (See also Waiting for Capacity below.)
	MinElbCapacity *int `pulumi:"minElbCapacity"`
	// The minimum size of the auto scale group.
	// (See also Waiting for Capacity below.)
	MinSize *int `pulumi:"minSize"`
	// Configuration block containing settings to define launch targets for Auto Scaling groups. Defined below.
	MixedInstancesPolicy *GroupMixedInstancesPolicy `pulumi:"mixedInstancesPolicy"`
	// The name of the auto scaling group. By default generated by this provider.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The name of the placement group into which you'll launch your instances, if any.
	PlacementGroup *string `pulumi:"placementGroup"`
	// Allows setting instance protection. The
	// autoscaling group will not select instances with this setting for termination
	// during scale in events.
	ProtectFromScaleIn *bool `pulumi:"protectFromScaleIn"`
	// The ARN of the service-linked role that the ASG will use to call other AWS services
	ServiceLinkedRoleArn *string `pulumi:"serviceLinkedRoleArn"`
	// A list of processes to suspend for the AutoScaling Group. The allowed values are `Launch`, `Terminate`, `HealthCheck`, `ReplaceUnhealthy`, `AZRebalance`, `AlarmNotification`, `ScheduledActions`, `AddToLoadBalancer`.
	// Note that if you suspend either the `Launch` or `Terminate` process types, it can prevent your autoscaling group from functioning properly.
	SuspendedProcesses []string `pulumi:"suspendedProcesses"`
	// Configuration block(s) containing resource tags. Conflicts with `tags`. Documented below.
	Tags []GroupTag `pulumi:"tags"`
	// Set of maps containing resource tags. Conflicts with `tag`. Documented below.
	TagsCollection []map[string]string `pulumi:"tagsCollection"`
	// A set of `alb.TargetGroup` ARNs, for use with Application or Network Load Balancing.
	TargetGroupArns []string `pulumi:"targetGroupArns"`
	// A list of policies to decide how the instances in the auto scale group should be terminated. The allowed values are `OldestInstance`, `NewestInstance`, `OldestLaunchConfiguration`, `ClosestToNextInstanceHour`, `OldestLaunchTemplate`, `AllocationStrategy`, `Default`.
	TerminationPolicies []string `pulumi:"terminationPolicies"`
	// A list of subnet IDs to launch resources in. Subnets automatically determine which availability zones the group will reside. Conflicts with `availabilityZones`.
	VpcZoneIdentifiers []string `pulumi:"vpcZoneIdentifiers"`
	// A maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for ASG instances to be healthy before timing out.  (See also Waiting
	// for Capacity below.) Setting this to "0" causes
	// this provider to skip all Capacity Waiting behavior.
	WaitForCapacityTimeout *string `pulumi:"waitForCapacityTimeout"`
	// Setting this will cause this provider to wait
	// for exactly this number of healthy instances from this autoscaling group in
	// all attached load balancers on both create and update operations. (Takes
	// precedence over `minElbCapacity` behavior.)
	// (See also Waiting for Capacity below.)
	WaitForElbCapacity *int `pulumi:"waitForElbCapacity"`
}

type GroupState struct {
	// The ARN for this AutoScaling Group
	Arn pulumi.StringPtrInput
	// A list of one or more availability zones for the group. Used for EC2-Classic and default subnets when not specified with `vpcZoneIdentifier` argument. Conflicts with `vpcZoneIdentifier`.
	AvailabilityZones pulumi.StringArrayInput
	// The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	DefaultCooldown pulumi.IntPtrInput
	// The number of Amazon EC2 instances that
	// should be running in the group. (See also Waiting for
	// Capacity below.)
	DesiredCapacity pulumi.IntPtrInput
	// A list of metrics to collect. The allowed values are `GroupDesiredCapacity`, `GroupInServiceCapacity`, `GroupPendingCapacity`, `GroupMinSize`, `GroupMaxSize`, `GroupInServiceInstances`, `GroupPendingInstances`, `GroupStandbyInstances`, `GroupStandbyCapacity`, `GroupTerminatingCapacity`, `GroupTerminatingInstances`, `GroupTotalCapacity`, `GroupTotalInstances`.
	EnabledMetrics pulumi.StringArrayInput
	// Allows deleting the autoscaling group without waiting
	// for all instances in the pool to terminate.  You can force an autoscaling group to delete
	// even if it's in the process of scaling a resource. Normally, this provider
	// drains all the instances before deleting the group.  This bypasses that
	// behavior and potentially leaves resources dangling.
	ForceDelete pulumi.BoolPtrInput
	// Time (in seconds) after instance comes into service before checking health.
	HealthCheckGracePeriod pulumi.IntPtrInput
	// "EC2" or "ELB". Controls how health checking is done.
	HealthCheckType pulumi.StringPtrInput
	// One or more
	// [Lifecycle Hooks](http://docs.aws.amazon.com/autoscaling/latest/userguide/lifecycle-hooks.html)
	// to attach to the autoscaling group **before** instances are launched. The
	// syntax is exactly the same as the separate
	// `autoscaling.LifecycleHook`
	// resource, without the `autoscalingGroupName` attribute. Please note that this will only work when creating
	// a new autoscaling group. For all other use-cases, please use `autoscaling.LifecycleHook` resource.
	InitialLifecycleHooks GroupInitialLifecycleHookArrayInput
	// The name of the launch configuration to use.
	LaunchConfiguration pulumi.StringPtrInput
	// Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
	LaunchTemplate GroupLaunchTemplatePtrInput
	// A list of elastic load balancer names to add to the autoscaling
	// group names. Only valid for classic load balancers. For ALBs, use `targetGroupArns` instead.
	LoadBalancers pulumi.StringArrayInput
	// The maximum amount of time, in seconds, that an instance can be in service, values must be either equal to 0 or between 604800 and 31536000 seconds.
	MaxInstanceLifetime pulumi.IntPtrInput
	// The maximum size of the auto scale group.
	MaxSize pulumi.IntPtrInput
	// The granularity to associate with the metrics to collect. The only valid value is `1Minute`. Default is `1Minute`.
	MetricsGranularity pulumi.StringPtrInput
	// Setting this causes this provider to wait for
	// this number of instances from this autoscaling group to show up healthy in the
	// ELB only on creation. Updates will not wait on ELB instance number changes.
	// (See also Waiting for Capacity below.)
	MinElbCapacity pulumi.IntPtrInput
	// The minimum size of the auto scale group.
	// (See also Waiting for Capacity below.)
	MinSize pulumi.IntPtrInput
	// Configuration block containing settings to define launch targets for Auto Scaling groups. Defined below.
	MixedInstancesPolicy GroupMixedInstancesPolicyPtrInput
	// The name of the auto scaling group. By default generated by this provider.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The name of the placement group into which you'll launch your instances, if any.
	PlacementGroup pulumi.StringPtrInput
	// Allows setting instance protection. The
	// autoscaling group will not select instances with this setting for termination
	// during scale in events.
	ProtectFromScaleIn pulumi.BoolPtrInput
	// The ARN of the service-linked role that the ASG will use to call other AWS services
	ServiceLinkedRoleArn pulumi.StringPtrInput
	// A list of processes to suspend for the AutoScaling Group. The allowed values are `Launch`, `Terminate`, `HealthCheck`, `ReplaceUnhealthy`, `AZRebalance`, `AlarmNotification`, `ScheduledActions`, `AddToLoadBalancer`.
	// Note that if you suspend either the `Launch` or `Terminate` process types, it can prevent your autoscaling group from functioning properly.
	SuspendedProcesses pulumi.StringArrayInput
	// Configuration block(s) containing resource tags. Conflicts with `tags`. Documented below.
	Tags GroupTagArrayInput
	// Set of maps containing resource tags. Conflicts with `tag`. Documented below.
	TagsCollection pulumi.StringMapArrayInput
	// A set of `alb.TargetGroup` ARNs, for use with Application or Network Load Balancing.
	TargetGroupArns pulumi.StringArrayInput
	// A list of policies to decide how the instances in the auto scale group should be terminated. The allowed values are `OldestInstance`, `NewestInstance`, `OldestLaunchConfiguration`, `ClosestToNextInstanceHour`, `OldestLaunchTemplate`, `AllocationStrategy`, `Default`.
	TerminationPolicies pulumi.StringArrayInput
	// A list of subnet IDs to launch resources in. Subnets automatically determine which availability zones the group will reside. Conflicts with `availabilityZones`.
	VpcZoneIdentifiers pulumi.StringArrayInput
	// A maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for ASG instances to be healthy before timing out.  (See also Waiting
	// for Capacity below.) Setting this to "0" causes
	// this provider to skip all Capacity Waiting behavior.
	WaitForCapacityTimeout pulumi.StringPtrInput
	// Setting this will cause this provider to wait
	// for exactly this number of healthy instances from this autoscaling group in
	// all attached load balancers on both create and update operations. (Takes
	// precedence over `minElbCapacity` behavior.)
	// (See also Waiting for Capacity below.)
	WaitForElbCapacity pulumi.IntPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// A list of one or more availability zones for the group. Used for EC2-Classic and default subnets when not specified with `vpcZoneIdentifier` argument. Conflicts with `vpcZoneIdentifier`.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	DefaultCooldown *int `pulumi:"defaultCooldown"`
	// The number of Amazon EC2 instances that
	// should be running in the group. (See also Waiting for
	// Capacity below.)
	DesiredCapacity *int `pulumi:"desiredCapacity"`
	// A list of metrics to collect. The allowed values are `GroupDesiredCapacity`, `GroupInServiceCapacity`, `GroupPendingCapacity`, `GroupMinSize`, `GroupMaxSize`, `GroupInServiceInstances`, `GroupPendingInstances`, `GroupStandbyInstances`, `GroupStandbyCapacity`, `GroupTerminatingCapacity`, `GroupTerminatingInstances`, `GroupTotalCapacity`, `GroupTotalInstances`.
	EnabledMetrics []string `pulumi:"enabledMetrics"`
	// Allows deleting the autoscaling group without waiting
	// for all instances in the pool to terminate.  You can force an autoscaling group to delete
	// even if it's in the process of scaling a resource. Normally, this provider
	// drains all the instances before deleting the group.  This bypasses that
	// behavior and potentially leaves resources dangling.
	ForceDelete *bool `pulumi:"forceDelete"`
	// Time (in seconds) after instance comes into service before checking health.
	HealthCheckGracePeriod *int `pulumi:"healthCheckGracePeriod"`
	// "EC2" or "ELB". Controls how health checking is done.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// One or more
	// [Lifecycle Hooks](http://docs.aws.amazon.com/autoscaling/latest/userguide/lifecycle-hooks.html)
	// to attach to the autoscaling group **before** instances are launched. The
	// syntax is exactly the same as the separate
	// `autoscaling.LifecycleHook`
	// resource, without the `autoscalingGroupName` attribute. Please note that this will only work when creating
	// a new autoscaling group. For all other use-cases, please use `autoscaling.LifecycleHook` resource.
	InitialLifecycleHooks []GroupInitialLifecycleHook `pulumi:"initialLifecycleHooks"`
	// The name of the launch configuration to use.
	LaunchConfiguration interface{} `pulumi:"launchConfiguration"`
	// Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
	LaunchTemplate *GroupLaunchTemplate `pulumi:"launchTemplate"`
	// A list of elastic load balancer names to add to the autoscaling
	// group names. Only valid for classic load balancers. For ALBs, use `targetGroupArns` instead.
	LoadBalancers []string `pulumi:"loadBalancers"`
	// The maximum amount of time, in seconds, that an instance can be in service, values must be either equal to 0 or between 604800 and 31536000 seconds.
	MaxInstanceLifetime *int `pulumi:"maxInstanceLifetime"`
	// The maximum size of the auto scale group.
	MaxSize int `pulumi:"maxSize"`
	// The granularity to associate with the metrics to collect. The only valid value is `1Minute`. Default is `1Minute`.
	MetricsGranularity interface{} `pulumi:"metricsGranularity"`
	// Setting this causes this provider to wait for
	// this number of instances from this autoscaling group to show up healthy in the
	// ELB only on creation. Updates will not wait on ELB instance number changes.
	// (See also Waiting for Capacity below.)
	MinElbCapacity *int `pulumi:"minElbCapacity"`
	// The minimum size of the auto scale group.
	// (See also Waiting for Capacity below.)
	MinSize int `pulumi:"minSize"`
	// Configuration block containing settings to define launch targets for Auto Scaling groups. Defined below.
	MixedInstancesPolicy *GroupMixedInstancesPolicy `pulumi:"mixedInstancesPolicy"`
	// The name of the auto scaling group. By default generated by this provider.
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The name of the placement group into which you'll launch your instances, if any.
	PlacementGroup interface{} `pulumi:"placementGroup"`
	// Allows setting instance protection. The
	// autoscaling group will not select instances with this setting for termination
	// during scale in events.
	ProtectFromScaleIn *bool `pulumi:"protectFromScaleIn"`
	// The ARN of the service-linked role that the ASG will use to call other AWS services
	ServiceLinkedRoleArn *string `pulumi:"serviceLinkedRoleArn"`
	// A list of processes to suspend for the AutoScaling Group. The allowed values are `Launch`, `Terminate`, `HealthCheck`, `ReplaceUnhealthy`, `AZRebalance`, `AlarmNotification`, `ScheduledActions`, `AddToLoadBalancer`.
	// Note that if you suspend either the `Launch` or `Terminate` process types, it can prevent your autoscaling group from functioning properly.
	SuspendedProcesses []string `pulumi:"suspendedProcesses"`
	// Configuration block(s) containing resource tags. Conflicts with `tags`. Documented below.
	Tags []GroupTag `pulumi:"tags"`
	// Set of maps containing resource tags. Conflicts with `tag`. Documented below.
	TagsCollection []map[string]string `pulumi:"tagsCollection"`
	// A set of `alb.TargetGroup` ARNs, for use with Application or Network Load Balancing.
	TargetGroupArns []string `pulumi:"targetGroupArns"`
	// A list of policies to decide how the instances in the auto scale group should be terminated. The allowed values are `OldestInstance`, `NewestInstance`, `OldestLaunchConfiguration`, `ClosestToNextInstanceHour`, `OldestLaunchTemplate`, `AllocationStrategy`, `Default`.
	TerminationPolicies []string `pulumi:"terminationPolicies"`
	// A list of subnet IDs to launch resources in. Subnets automatically determine which availability zones the group will reside. Conflicts with `availabilityZones`.
	VpcZoneIdentifiers []string `pulumi:"vpcZoneIdentifiers"`
	// A maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for ASG instances to be healthy before timing out.  (See also Waiting
	// for Capacity below.) Setting this to "0" causes
	// this provider to skip all Capacity Waiting behavior.
	WaitForCapacityTimeout *string `pulumi:"waitForCapacityTimeout"`
	// Setting this will cause this provider to wait
	// for exactly this number of healthy instances from this autoscaling group in
	// all attached load balancers on both create and update operations. (Takes
	// precedence over `minElbCapacity` behavior.)
	// (See also Waiting for Capacity below.)
	WaitForElbCapacity *int `pulumi:"waitForElbCapacity"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// A list of one or more availability zones for the group. Used for EC2-Classic and default subnets when not specified with `vpcZoneIdentifier` argument. Conflicts with `vpcZoneIdentifier`.
	AvailabilityZones pulumi.StringArrayInput
	// The amount of time, in seconds, after a scaling activity completes before another scaling activity can start.
	DefaultCooldown pulumi.IntPtrInput
	// The number of Amazon EC2 instances that
	// should be running in the group. (See also Waiting for
	// Capacity below.)
	DesiredCapacity pulumi.IntPtrInput
	// A list of metrics to collect. The allowed values are `GroupDesiredCapacity`, `GroupInServiceCapacity`, `GroupPendingCapacity`, `GroupMinSize`, `GroupMaxSize`, `GroupInServiceInstances`, `GroupPendingInstances`, `GroupStandbyInstances`, `GroupStandbyCapacity`, `GroupTerminatingCapacity`, `GroupTerminatingInstances`, `GroupTotalCapacity`, `GroupTotalInstances`.
	EnabledMetrics pulumi.StringArrayInput
	// Allows deleting the autoscaling group without waiting
	// for all instances in the pool to terminate.  You can force an autoscaling group to delete
	// even if it's in the process of scaling a resource. Normally, this provider
	// drains all the instances before deleting the group.  This bypasses that
	// behavior and potentially leaves resources dangling.
	ForceDelete pulumi.BoolPtrInput
	// Time (in seconds) after instance comes into service before checking health.
	HealthCheckGracePeriod pulumi.IntPtrInput
	// "EC2" or "ELB". Controls how health checking is done.
	HealthCheckType pulumi.StringPtrInput
	// One or more
	// [Lifecycle Hooks](http://docs.aws.amazon.com/autoscaling/latest/userguide/lifecycle-hooks.html)
	// to attach to the autoscaling group **before** instances are launched. The
	// syntax is exactly the same as the separate
	// `autoscaling.LifecycleHook`
	// resource, without the `autoscalingGroupName` attribute. Please note that this will only work when creating
	// a new autoscaling group. For all other use-cases, please use `autoscaling.LifecycleHook` resource.
	InitialLifecycleHooks GroupInitialLifecycleHookArrayInput
	// The name of the launch configuration to use.
	LaunchConfiguration pulumi.Input
	// Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
	LaunchTemplate GroupLaunchTemplatePtrInput
	// A list of elastic load balancer names to add to the autoscaling
	// group names. Only valid for classic load balancers. For ALBs, use `targetGroupArns` instead.
	LoadBalancers pulumi.StringArrayInput
	// The maximum amount of time, in seconds, that an instance can be in service, values must be either equal to 0 or between 604800 and 31536000 seconds.
	MaxInstanceLifetime pulumi.IntPtrInput
	// The maximum size of the auto scale group.
	MaxSize pulumi.IntInput
	// The granularity to associate with the metrics to collect. The only valid value is `1Minute`. Default is `1Minute`.
	MetricsGranularity pulumi.Input
	// Setting this causes this provider to wait for
	// this number of instances from this autoscaling group to show up healthy in the
	// ELB only on creation. Updates will not wait on ELB instance number changes.
	// (See also Waiting for Capacity below.)
	MinElbCapacity pulumi.IntPtrInput
	// The minimum size of the auto scale group.
	// (See also Waiting for Capacity below.)
	MinSize pulumi.IntInput
	// Configuration block containing settings to define launch targets for Auto Scaling groups. Defined below.
	MixedInstancesPolicy GroupMixedInstancesPolicyPtrInput
	// The name of the auto scaling group. By default generated by this provider.
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The name of the placement group into which you'll launch your instances, if any.
	PlacementGroup pulumi.Input
	// Allows setting instance protection. The
	// autoscaling group will not select instances with this setting for termination
	// during scale in events.
	ProtectFromScaleIn pulumi.BoolPtrInput
	// The ARN of the service-linked role that the ASG will use to call other AWS services
	ServiceLinkedRoleArn pulumi.StringPtrInput
	// A list of processes to suspend for the AutoScaling Group. The allowed values are `Launch`, `Terminate`, `HealthCheck`, `ReplaceUnhealthy`, `AZRebalance`, `AlarmNotification`, `ScheduledActions`, `AddToLoadBalancer`.
	// Note that if you suspend either the `Launch` or `Terminate` process types, it can prevent your autoscaling group from functioning properly.
	SuspendedProcesses pulumi.StringArrayInput
	// Configuration block(s) containing resource tags. Conflicts with `tags`. Documented below.
	Tags GroupTagArrayInput
	// Set of maps containing resource tags. Conflicts with `tag`. Documented below.
	TagsCollection pulumi.StringMapArrayInput
	// A set of `alb.TargetGroup` ARNs, for use with Application or Network Load Balancing.
	TargetGroupArns pulumi.StringArrayInput
	// A list of policies to decide how the instances in the auto scale group should be terminated. The allowed values are `OldestInstance`, `NewestInstance`, `OldestLaunchConfiguration`, `ClosestToNextInstanceHour`, `OldestLaunchTemplate`, `AllocationStrategy`, `Default`.
	TerminationPolicies pulumi.StringArrayInput
	// A list of subnet IDs to launch resources in. Subnets automatically determine which availability zones the group will reside. Conflicts with `availabilityZones`.
	VpcZoneIdentifiers pulumi.StringArrayInput
	// A maximum
	// [duration](https://golang.org/pkg/time/#ParseDuration) that this provider should
	// wait for ASG instances to be healthy before timing out.  (See also Waiting
	// for Capacity below.) Setting this to "0" causes
	// this provider to skip all Capacity Waiting behavior.
	WaitForCapacityTimeout pulumi.StringPtrInput
	// Setting this will cause this provider to wait
	// for exactly this number of healthy instances from this autoscaling group in
	// all attached load balancers on both create and update operations. (Takes
	// precedence over `minElbCapacity` behavior.)
	// (See also Waiting for Capacity below.)
	WaitForElbCapacity pulumi.IntPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
