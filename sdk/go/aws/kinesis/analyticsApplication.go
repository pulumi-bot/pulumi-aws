// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
// allows processing and analyzing streaming data using standard SQL.
// 
// For more details, see the [Amazon Kinesis Analytics Documentation][1].
type AnalyticsApplication struct {
	s *pulumi.ResourceState
}

// NewAnalyticsApplication registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsApplication(ctx *pulumi.Context,
	name string, args *AnalyticsApplicationArgs, opts ...pulumi.ResourceOpt) (*AnalyticsApplication, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["cloudwatchLoggingOptions"] = nil
		inputs["code"] = nil
		inputs["description"] = nil
		inputs["inputs"] = nil
		inputs["name"] = nil
		inputs["outputs"] = nil
		inputs["referenceDataSources"] = nil
		inputs["tags"] = nil
	} else {
		inputs["cloudwatchLoggingOptions"] = args.CloudwatchLoggingOptions
		inputs["code"] = args.Code
		inputs["description"] = args.Description
		inputs["inputs"] = args.Inputs
		inputs["name"] = args.Name
		inputs["outputs"] = args.Outputs
		inputs["referenceDataSources"] = args.ReferenceDataSources
		inputs["tags"] = args.Tags
	}
	inputs["arn"] = nil
	inputs["createTimestamp"] = nil
	inputs["lastUpdateTimestamp"] = nil
	inputs["status"] = nil
	inputs["version"] = nil
	s, err := ctx.RegisterResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsApplication{s: s}, nil
}

// GetAnalyticsApplication gets an existing AnalyticsApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsApplication(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AnalyticsApplicationState, opts ...pulumi.ResourceOpt) (*AnalyticsApplication, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["cloudwatchLoggingOptions"] = state.CloudwatchLoggingOptions
		inputs["code"] = state.Code
		inputs["createTimestamp"] = state.CreateTimestamp
		inputs["description"] = state.Description
		inputs["inputs"] = state.Inputs
		inputs["lastUpdateTimestamp"] = state.LastUpdateTimestamp
		inputs["name"] = state.Name
		inputs["outputs"] = state.Outputs
		inputs["referenceDataSources"] = state.ReferenceDataSources
		inputs["status"] = state.Status
		inputs["tags"] = state.Tags
		inputs["version"] = state.Version
	}
	s, err := ctx.ReadResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AnalyticsApplication{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AnalyticsApplication) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AnalyticsApplication) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the Kinesis Analytics Appliation.
func (r *AnalyticsApplication) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The CloudWatch log stream options to monitor application errors.
// See CloudWatch Logging Options below for more details.
func (r *AnalyticsApplication) CloudwatchLoggingOptions() *pulumi.Output {
	return r.s.State["cloudwatchLoggingOptions"]
}

// SQL Code to transform input data, and generate output.
func (r *AnalyticsApplication) Code() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["code"])
}

// The Timestamp when the application version was created.
func (r *AnalyticsApplication) CreateTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["createTimestamp"])
}

// Description of the application.
func (r *AnalyticsApplication) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// Input configuration of the application. See Inputs below for more details.
func (r *AnalyticsApplication) Inputs() *pulumi.Output {
	return r.s.State["inputs"]
}

// The Timestamp when the application was last updated.
func (r *AnalyticsApplication) LastUpdateTimestamp() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["lastUpdateTimestamp"])
}

// Name of the Kinesis Analytics Application.
func (r *AnalyticsApplication) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Output destination configuration of the application. See Outputs below for more details.
func (r *AnalyticsApplication) Outputs() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["outputs"])
}

// An S3 Reference Data Source for the application.
// See Reference Data Sources below for more details.
func (r *AnalyticsApplication) ReferenceDataSources() *pulumi.Output {
	return r.s.State["referenceDataSources"]
}

// The Status of the application.
func (r *AnalyticsApplication) Status() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["status"])
}

// Key-value mapping of tags for the Kinesis Analytics Application.
func (r *AnalyticsApplication) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// The Version of the application.
func (r *AnalyticsApplication) Version() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["version"])
}

// Input properties used for looking up and filtering AnalyticsApplication resources.
type AnalyticsApplicationState struct {
	// The ARN of the Kinesis Analytics Appliation.
	Arn interface{}
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions interface{}
	// SQL Code to transform input data, and generate output.
	Code interface{}
	// The Timestamp when the application version was created.
	CreateTimestamp interface{}
	// Description of the application.
	Description interface{}
	// Input configuration of the application. See Inputs below for more details.
	Inputs interface{}
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp interface{}
	// Name of the Kinesis Analytics Application.
	Name interface{}
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs interface{}
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources interface{}
	// The Status of the application.
	Status interface{}
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags interface{}
	// The Version of the application.
	Version interface{}
}

// The set of arguments for constructing a AnalyticsApplication resource.
type AnalyticsApplicationArgs struct {
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions interface{}
	// SQL Code to transform input data, and generate output.
	Code interface{}
	// Description of the application.
	Description interface{}
	// Input configuration of the application. See Inputs below for more details.
	Inputs interface{}
	// Name of the Kinesis Analytics Application.
	Name interface{}
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs interface{}
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources interface{}
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags interface{}
}
