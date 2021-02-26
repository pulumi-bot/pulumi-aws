// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Autoscaling Groups data source allows access to the list of AWS
// ASGs within a specific region. This will allow you to pass a list of AutoScaling Groups to other resources.
//
// Deprecated: aws.getAutoscalingGroups has been deprecated in favor of aws.autoscaling.getAmiIds
func GetAutoscalingGroups(ctx *pulumi.Context, args *GetAutoscalingGroupsArgs, opts ...pulumi.InvokeOption) (*GetAutoscalingGroupsResult, error) {
	var rv GetAutoscalingGroupsResult
	err := ctx.Invoke("aws:index/getAutoscalingGroups:getAutoscalingGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoscalingGroups.
type GetAutoscalingGroupsArgs struct {
	// A filter used to scope the list e.g. by tags. See [related docs](http://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_Filter.html).
	Filters []GetAutoscalingGroupsFilter `pulumi:"filters"`
}

// A collection of values returned by getAutoscalingGroups.
type GetAutoscalingGroupsResult struct {
	// A list of the Autoscaling Groups Arns in the current region.
	Arns    []string                     `pulumi:"arns"`
	Filters []GetAutoscalingGroupsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the Autoscaling Groups in the current region.
	Names []string `pulumi:"names"`
}
