// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Batch Job Definition resource.
// 
// ## retryStrategy
// 
// `retryStrategy` supports the following:
// 
// * `attempts` - (Optional) The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
// 
// ## timeout
// 
// `timeout` supports the following:
// 
// * `attemptDurationSeconds` - (Optional) The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/batch_job_definition.html.markdown.
type JobDefinition struct {
	s *pulumi.ResourceState
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOpt) (*JobDefinition, error) {
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["containerProperties"] = nil
		inputs["name"] = nil
		inputs["parameters"] = nil
		inputs["retryStrategy"] = nil
		inputs["timeout"] = nil
		inputs["type"] = nil
	} else {
		inputs["containerProperties"] = args.ContainerProperties
		inputs["name"] = args.Name
		inputs["parameters"] = args.Parameters
		inputs["retryStrategy"] = args.RetryStrategy
		inputs["timeout"] = args.Timeout
		inputs["type"] = args.Type
	}
	inputs["arn"] = nil
	inputs["revision"] = nil
	s, err := ctx.RegisterResource("aws:batch/jobDefinition:JobDefinition", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobDefinition{s: s}, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JobDefinitionState, opts ...pulumi.ResourceOpt) (*JobDefinition, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["containerProperties"] = state.ContainerProperties
		inputs["name"] = state.Name
		inputs["parameters"] = state.Parameters
		inputs["retryStrategy"] = state.RetryStrategy
		inputs["revision"] = state.Revision
		inputs["timeout"] = state.Timeout
		inputs["type"] = state.Type
	}
	s, err := ctx.ReadResource("aws:batch/jobDefinition:JobDefinition", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobDefinition{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *JobDefinition) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *JobDefinition) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name of the job definition.
func (r *JobDefinition) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
func (r *JobDefinition) ContainerProperties() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["containerProperties"])
}

// Specifies the name of the job definition.
func (r *JobDefinition) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Specifies the parameter substitution placeholders to set in the job definition.
func (r *JobDefinition) Parameters() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["parameters"])
}

// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
// Maximum number of `retryStrategy` is `1`.  Defined below.
func (r *JobDefinition) RetryStrategy() *pulumi.Output {
	return r.s.State["retryStrategy"]
}

// The revision of the job definition.
func (r *JobDefinition) Revision() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["revision"])
}

// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
func (r *JobDefinition) Timeout() *pulumi.Output {
	return r.s.State["timeout"]
}

// The type of job definition.  Must be `container`
func (r *JobDefinition) Type() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["type"])
}

// Input properties used for looking up and filtering JobDefinition resources.
type JobDefinitionState struct {
	// The Amazon Resource Name of the job definition.
	Arn interface{}
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties interface{}
	// Specifies the name of the job definition.
	Name interface{}
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters interface{}
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy interface{}
	// The revision of the job definition.
	Revision interface{}
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout interface{}
	// The type of job definition.  Must be `container`
	Type interface{}
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	// A valid [container properties](http://docs.aws.amazon.com/batch/latest/APIReference/API_RegisterJobDefinition.html)
	// provided as a single valid JSON document. This parameter is required if the `type` parameter is `container`.
	ContainerProperties interface{}
	// Specifies the name of the job definition.
	Name interface{}
	// Specifies the parameter substitution placeholders to set in the job definition.
	Parameters interface{}
	// Specifies the retry strategy to use for failed jobs that are submitted with this job definition.
	// Maximum number of `retryStrategy` is `1`.  Defined below.
	RetryStrategy interface{}
	// Specifies the timeout for jobs so that if a job runs longer, AWS Batch terminates the job. Maximum number of `timeout` is `1`. Defined below.
	Timeout interface{}
	// The type of job definition.  Must be `container`
	Type interface{}
}
