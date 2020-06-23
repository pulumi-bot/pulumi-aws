// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Batch Job Queue data source allows access to details of a specific
// job queue within AWS Batch.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/batch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.LookupJobQueue(ctx, &batch.LookupJobQueueArgs{
// 			Name: "tf-test-batch-job-queue",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupJobQueue(ctx *pulumi.Context, args *LookupJobQueueArgs, opts ...pulumi.InvokeOption) (*LookupJobQueueResult, error) {
	var rv LookupJobQueueResult
	err := ctx.Invoke("aws:batch/getJobQueue:getJobQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getJobQueue.
type LookupJobQueueArgs struct {
	// The name of the job queue.
	Name string `pulumi:"name"`
}

// A collection of values returned by getJobQueue.
type LookupJobQueueResult struct {
	// The ARN of the job queue.
	Arn string `pulumi:"arn"`
	// The compute environments that are attached to the job queue and the order in
	// which job placement is preferred. Compute environments are selected for job placement in ascending order.
	// * `compute_environment_order.#.order` - The order of the compute environment.
	// * `compute_environment_order.#.compute_environment` - The ARN of the compute environment.
	ComputeEnvironmentOrders []GetJobQueueComputeEnvironmentOrder `pulumi:"computeEnvironmentOrders"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The priority of the job queue. Job queues with a higher priority are evaluated first when
	// associated with the same compute environment.
	Priority int `pulumi:"priority"`
	// Describes the ability of the queue to accept new jobs (for example, `ENABLED` or `DISABLED`).
	State string `pulumi:"state"`
	// The current status of the job queue (for example, `CREATING` or `VALID`).
	Status string `pulumi:"status"`
	// A short, human-readable string to provide additional details about the current status
	// of the job queue.
	StatusReason string `pulumi:"statusReason"`
}
