// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package job

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/glue/JobCommand"
	"https:/github.com/pulumi/pulumi-aws/glue/JobExecutionProperty"
)

// Provides a Glue Job resource.
// 
// > Glue functionality, such as monitoring and logging of jobs, is typically managed with the `defaultArguments` argument. See the [Special Parameters Used by AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-etl-glue-arguments.html) topic in the Glue developer guide for additional information.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/glue_job.html.markdown.
type Job struct {
	pulumi.CustomResourceState

	// **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	AllocatedCapacity pulumi.IntOutput `pulumi:"allocatedCapacity"`
	// Amazon Resource Name (ARN) of Glue Job
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The command of the job. Defined below.
	Command glueJobCommand.JobCommandOutput `pulumi:"command"`
	// The list of connections used for this job.
	Connections pulumi.StringArrayOutput `pulumi:"connections"`
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments pulumi.MapOutput `pulumi:"defaultArguments"`
	// Description of the job.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Execution property of the job. Defined below.
	ExecutionProperty glueJobExecutionProperty.JobExecutionPropertyOutput `pulumi:"executionProperty"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringOutput `pulumi:"glueVersion"`
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	MaxCapacity pulumi.Float64Output `pulumi:"maxCapacity"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries pulumi.IntPtrOutput `pulumi:"maxRetries"`
	// The name of the job command. Defaults to `glueetl`
	Name pulumi.StringOutput `pulumi:"name"`
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers pulumi.IntPtrOutput `pulumi:"numberOfWorkers"`
	// The ARN of the IAM role associated with this job.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration pulumi.StringPtrOutput `pulumi:"securityConfiguration"`
	// Key-value mapping of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrOutput `pulumi:"workerType"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil || args.Command == nil {
		return nil, errors.New("missing required argument 'Command'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil {
		args = &JobArgs{}
	}
	var resource Job
	err := ctx.RegisterResource("aws:glue/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("aws:glue/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	AllocatedCapacity *int `pulumi:"allocatedCapacity"`
	// Amazon Resource Name (ARN) of Glue Job
	Arn *string `pulumi:"arn"`
	// The command of the job. Defined below.
	Command *glueJobCommand.JobCommand `pulumi:"command"`
	// The list of connections used for this job.
	Connections []string `pulumi:"connections"`
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments map[string]interface{} `pulumi:"defaultArguments"`
	// Description of the job.
	Description *string `pulumi:"description"`
	// Execution property of the job. Defined below.
	ExecutionProperty *glueJobExecutionProperty.JobExecutionProperty `pulumi:"executionProperty"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion *string `pulumi:"glueVersion"`
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries *int `pulumi:"maxRetries"`
	// The name of the job command. Defaults to `glueetl`
	Name *string `pulumi:"name"`
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// The ARN of the IAM role associated with this job.
	RoleArn *string `pulumi:"roleArn"`
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout *int `pulumi:"timeout"`
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType *string `pulumi:"workerType"`
}

type JobState struct {
	// **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	AllocatedCapacity pulumi.IntPtrInput
	// Amazon Resource Name (ARN) of Glue Job
	Arn pulumi.StringPtrInput
	// The command of the job. Defined below.
	Command glueJobCommand.JobCommandPtrInput
	// The list of connections used for this job.
	Connections pulumi.StringArrayInput
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments pulumi.MapInput
	// Description of the job.
	Description pulumi.StringPtrInput
	// Execution property of the job. Defined below.
	ExecutionProperty glueJobExecutionProperty.JobExecutionPropertyPtrInput
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringPtrInput
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	MaxCapacity pulumi.Float64PtrInput
	// The maximum number of times to retry this job if it fails.
	MaxRetries pulumi.IntPtrInput
	// The name of the job command. Defaults to `glueetl`
	Name pulumi.StringPtrInput
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers pulumi.IntPtrInput
	// The ARN of the IAM role associated with this job.
	RoleArn pulumi.StringPtrInput
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrInput
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	AllocatedCapacity *int `pulumi:"allocatedCapacity"`
	// The command of the job. Defined below.
	Command glueJobCommand.JobCommand `pulumi:"command"`
	// The list of connections used for this job.
	Connections []string `pulumi:"connections"`
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments map[string]interface{} `pulumi:"defaultArguments"`
	// Description of the job.
	Description *string `pulumi:"description"`
	// Execution property of the job. Defined below.
	ExecutionProperty *glueJobExecutionProperty.JobExecutionProperty `pulumi:"executionProperty"`
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion *string `pulumi:"glueVersion"`
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	MaxCapacity *float64 `pulumi:"maxCapacity"`
	// The maximum number of times to retry this job if it fails.
	MaxRetries *int `pulumi:"maxRetries"`
	// The name of the job command. Defaults to `glueetl`
	Name *string `pulumi:"name"`
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers *int `pulumi:"numberOfWorkers"`
	// The ARN of the IAM role associated with this job.
	RoleArn string `pulumi:"roleArn"`
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration *string `pulumi:"securityConfiguration"`
	// Key-value mapping of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout *int `pulumi:"timeout"`
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType *string `pulumi:"workerType"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// **DEPRECATED** (Optional) The number of AWS Glue data processing units (DPUs) to allocate to this Job. At least 2 DPUs need to be allocated; the default is 10. A DPU is a relative measure of processing power that consists of 4 vCPUs of compute capacity and 16 GB of memory.
	AllocatedCapacity pulumi.IntPtrInput
	// The command of the job. Defined below.
	Command glueJobCommand.JobCommandInput
	// The list of connections used for this job.
	Connections pulumi.StringArrayInput
	// The map of default arguments for this job. You can specify arguments here that your own job-execution script consumes, as well as arguments that AWS Glue itself consumes. For information about how to specify and consume your own Job arguments, see the [Calling AWS Glue APIs in Python](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-calling.html) topic in the developer guide. For information about the key-value pairs that AWS Glue consumes to set up your job, see the [Special Parameters Used by AWS Glue](http://docs.aws.amazon.com/glue/latest/dg/aws-glue-programming-python-glue-arguments.html) topic in the developer guide.
	DefaultArguments pulumi.MapInput
	// Description of the job.
	Description pulumi.StringPtrInput
	// Execution property of the job. Defined below.
	ExecutionProperty glueJobExecutionProperty.JobExecutionPropertyPtrInput
	// The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
	GlueVersion pulumi.StringPtrInput
	// The maximum number of AWS Glue data processing units (DPUs) that can be allocated when this job runs.
	MaxCapacity pulumi.Float64PtrInput
	// The maximum number of times to retry this job if it fails.
	MaxRetries pulumi.IntPtrInput
	// The name of the job command. Defaults to `glueetl`
	Name pulumi.StringPtrInput
	// The number of workers of a defined workerType that are allocated when a job runs.
	NumberOfWorkers pulumi.IntPtrInput
	// The ARN of the IAM role associated with this job.
	RoleArn pulumi.StringInput
	// The name of the Security Configuration to be associated with the job.
	SecurityConfiguration pulumi.StringPtrInput
	// Key-value mapping of resource tags
	Tags pulumi.MapInput
	// The job timeout in minutes. The default is 2880 minutes (48 hours).
	Timeout pulumi.IntPtrInput
	// The type of predefined worker that is allocated when a job runs. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType pulumi.StringPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

