// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Batch Job Definition resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/batch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := batch.NewJobDefinition(ctx, "test", &batch.JobDefinitionArgs{
// 			ContainerProperties: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "	\"command\": [\"ls\", \"-la\"],\n", "	\"image\": \"busybox\",\n", "	\"memory\": 1024,\n", "	\"vcpus\": 1,\n", "	\"volumes\": [\n", "      {\n", "        \"host\": {\n", "          \"sourcePath\": \"/tmp\"\n", "        },\n", "        \"name\": \"tmp\"\n", "      }\n", "    ],\n", "	\"environment\": [\n", "		{\"name\": \"VARNAME\", \"value\": \"VARVAL\"}\n", "	],\n", "	\"mountPoints\": [\n", "		{\n", "          \"sourceVolume\": \"tmp\",\n", "          \"containerPath\": \"/tmp\",\n", "          \"readOnly\": false\n", "        }\n", "	],\n", "    \"ulimits\": [\n", "      {\n", "        \"hardLimit\": 1024,\n", "        \"name\": \"nofile\",\n", "        \"softLimit\": 1024\n", "      }\n", "    ]\n", "}\n", "\n")),
// 			Type: pulumi.String("container"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Batch Job Definition can be imported using the `arn`, e.g.
//
// ```sh
//  $ pulumi import aws:batch/jobDefinition:JobDefinition test arn:aws:batch:us-east-1:123456789012:job-definition/sample
// ```
type JobDefinition struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the job definition.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties pulumi.StringPtrOutput `pulumi:"containerProperties"`
	// Specifies the name of the job definition.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy JobDefinitionRetryStrategyPtrOutput `pulumi:"retryStrategy"`
	// The revision of the job definition.
	Revision pulumi.IntOutput `pulumi:"revision"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout JobDefinitionTimeoutPtrOutput `pulumi:"timeout"`
	// The type of job definition.  Must be `container`
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource JobDefinition
	err := ctx.RegisterResource("aws:batch/jobDefinition:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("aws:batch/jobDefinition:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobDefinition resources.
type jobDefinitionState struct {
	// The Amazon Resource Name of the job definition.
	Arn *string `pulumi:"arn"`
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties *string `pulumi:"containerProperties"`
	// Specifies the name of the job definition.
	Name *string `pulumi:"name"`
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters map[string]string `pulumi:"parameters"`
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy *JobDefinitionRetryStrategy `pulumi:"retryStrategy"`
	// The revision of the job definition.
	Revision *int `pulumi:"revision"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout *JobDefinitionTimeout `pulumi:"timeout"`
	// The type of job definition.  Must be `container`
	Type *string `pulumi:"type"`
}

type JobDefinitionState struct {
	// The Amazon Resource Name of the job definition.
	Arn pulumi.StringPtrInput
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties pulumi.StringPtrInput
	// Specifies the name of the job definition.
	Name pulumi.StringPtrInput
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters pulumi.StringMapInput
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy JobDefinitionRetryStrategyPtrInput
	// The revision of the job definition.
	Revision pulumi.IntPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout JobDefinitionTimeoutPtrInput
	// The type of job definition.  Must be `container`
	Type pulumi.StringPtrInput
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties *string `pulumi:"containerProperties"`
	// Specifies the name of the job definition.
	Name *string `pulumi:"name"`
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters map[string]string `pulumi:"parameters"`
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy *JobDefinitionRetryStrategy `pulumi:"retryStrategy"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout *JobDefinitionTimeout `pulumi:"timeout"`
	// The type of job definition.  Must be `container`
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties pulumi.StringPtrInput
	// Specifies the name of the job definition.
	Name pulumi.StringPtrInput
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters pulumi.StringMapInput
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy JobDefinitionRetryStrategyPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout JobDefinitionTimeoutPtrInput
	// The type of job definition.  Must be `container`
	Type pulumi.StringInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDefinition)(nil)).Elem()
}

func (i JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

type JobDefinitionOutput struct {
	*pulumi.OutputState
}

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobDefinitionOutput)(nil)).Elem()
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobDefinitionOutput{})
}
