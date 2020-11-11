// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Batch Job Queue resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/batch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.NewJobQueue(ctx, "testQueue", &batch.JobQueueArgs{
// 			State:    pulumi.String("ENABLED"),
// 			Priority: pulumi.Int(1),
// 			ComputeEnvironments: pulumi.StringArray{
// 				pulumi.Any(aws_batch_compute_environment.Test_environment_1.Arn),
// 				pulumi.Any(aws_batch_compute_environment.Test_environment_2.Arn),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type JobQueue struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the job queue.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments pulumi.StringArrayOutput `pulumi:"computeEnvironments"`
	// Specifies the name of the job queue.
	Name pulumi.StringOutput `pulumi:"name"`
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State pulumi.StringOutput `pulumi:"state"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewJobQueue registers a new resource with the given unique name, arguments, and options.
func NewJobQueue(ctx *pulumi.Context,
	name string, args *JobQueueArgs, opts ...pulumi.ResourceOption) (*JobQueue, error) {
	if args == nil || args.ComputeEnvironments == nil {
		return nil, errors.New("missing required argument 'ComputeEnvironments'")
	}
	if args == nil || args.Priority == nil {
		return nil, errors.New("missing required argument 'Priority'")
	}
	if args == nil || args.State == nil {
		return nil, errors.New("missing required argument 'State'")
	}
	if args == nil {
		args = &JobQueueArgs{}
	}
	var resource JobQueue
	err := ctx.RegisterResource("aws:batch/jobQueue:JobQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobQueue gets an existing JobQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobQueueState, opts ...pulumi.ResourceOption) (*JobQueue, error) {
	var resource JobQueue
	err := ctx.ReadResource("aws:batch/jobQueue:JobQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobQueue resources.
type jobQueueState struct {
	// The Amazon Resource Name of the job queue.
	Arn *string `pulumi:"arn"`
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments []string `pulumi:"computeEnvironments"`
	// Specifies the name of the job queue.
	Name *string `pulumi:"name"`
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority *int `pulumi:"priority"`
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State *string `pulumi:"state"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type JobQueueState struct {
	// The Amazon Resource Name of the job queue.
	Arn pulumi.StringPtrInput
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments pulumi.StringArrayInput
	// Specifies the name of the job queue.
	Name pulumi.StringPtrInput
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority pulumi.IntPtrInput
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (JobQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobQueueState)(nil)).Elem()
}

type jobQueueArgs struct {
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments []string `pulumi:"computeEnvironments"`
	// Specifies the name of the job queue.
	Name *string `pulumi:"name"`
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority int `pulumi:"priority"`
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State string `pulumi:"state"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a JobQueue resource.
type JobQueueArgs struct {
	// Specifies the set of compute environments
	// mapped to a job queue and their order.  The position of the compute environments
	// in the list will dictate the order. You can associate up to 3 compute environments
	// with a job queue.
	ComputeEnvironments pulumi.StringArrayInput
	// Specifies the name of the job queue.
	Name pulumi.StringPtrInput
	// The priority of the job queue. Job queues with a higher priority
	// are evaluated first when associated with the same compute environment.
	Priority pulumi.IntInput
	// The state of the job queue. Must be one of: `ENABLED` or `DISABLED`
	State pulumi.StringInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (JobQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobQueueArgs)(nil)).Elem()
}

type JobQueueInput interface {
	pulumi.Input

	ToJobQueueOutput() JobQueueOutput
	ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput
}

func (JobQueue) ElementType() reflect.Type {
	return reflect.TypeOf((*JobQueue)(nil)).Elem()
}

func (i JobQueue) ToJobQueueOutput() JobQueueOutput {
	return i.ToJobQueueOutputWithContext(context.Background())
}

func (i JobQueue) ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobQueueOutput)
}

type JobQueueOutput struct {
	*pulumi.OutputState
}

func (JobQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobQueueOutput)(nil)).Elem()
}

func (o JobQueueOutput) ToJobQueueOutput() JobQueueOutput {
	return o
}

func (o JobQueueOutput) ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobQueueOutput{})
}
