// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get information on an existing autoscaling group.
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("aws:autoscaling/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// Specify the exact name of the desired autoscaling group.
	Name string `pulumi:"name"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// The Amazon Resource Name (ARN) of the Auto Scaling group.
	Arn string `pulumi:"arn"`
	// One or more Availability Zones for the group.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	DefaultCooldown   int      `pulumi:"defaultCooldown"`
	// The desired size of the group.
	DesiredCapacity int `pulumi:"desiredCapacity"`
	// The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
	HealthCheckGracePeriod int `pulumi:"healthCheckGracePeriod"`
	// The service to use for the health checks. The valid values are EC2 and ELB.
	HealthCheckType string `pulumi:"healthCheckType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the associated launch configuration.
	LaunchConfiguration string `pulumi:"launchConfiguration"`
	// One or more load balancers associated with the group.
	LoadBalancers []string `pulumi:"loadBalancers"`
	// The maximum size of the group.
	MaxSize int `pulumi:"maxSize"`
	// The minimum size of the group.
	MinSize int `pulumi:"minSize"`
	// Name of the Auto Scaling Group.
	Name                             string `pulumi:"name"`
	NewInstancesProtectedFromScaleIn bool   `pulumi:"newInstancesProtectedFromScaleIn"`
	// The name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.
	PlacementGroup string `pulumi:"placementGroup"`
	// The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
	ServiceLinkedRoleArn string `pulumi:"serviceLinkedRoleArn"`
	// The current state of the group when DeleteAutoScalingGroup is in progress.
	Status string `pulumi:"status"`
	// The Amazon Resource Names (ARN) of the target groups for your load balancer.
	TargetGroupArns []string `pulumi:"targetGroupArns"`
	// The termination policies for the group.
	TerminationPolicies []string `pulumi:"terminationPolicies"`
	// VPC ID for the group.
	VpcZoneIdentifier string `pulumi:"vpcZoneIdentifier"`
}
