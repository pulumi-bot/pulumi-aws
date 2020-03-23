// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecs

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The ECS Cluster data source allows access to details of a specific
// cluster within an AWS ECS service.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecs_cluster.html.markdown.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("aws:ecs/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// The name of the ECS Cluster
	ClusterName string `pulumi:"clusterName"`
}


// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// The ARN of the ECS Cluster
	Arn string `pulumi:"arn"`
	ClusterName string `pulumi:"clusterName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The number of pending tasks for the ECS Cluster
	PendingTasksCount int `pulumi:"pendingTasksCount"`
	// The number of registered container instances for the ECS Cluster
	RegisteredContainerInstancesCount int `pulumi:"registeredContainerInstancesCount"`
	// The number of running tasks for the ECS Cluster
	RunningTasksCount int `pulumi:"runningTasksCount"`
	// The settings associated with the ECS Cluster.
	Settings []GetClusterSetting `pulumi:"settings"`
	// The status of the ECS Cluster
	Status string `pulumi:"status"`
}

