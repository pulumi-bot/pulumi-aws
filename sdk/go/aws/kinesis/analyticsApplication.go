// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kinesis

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
// allows processing and analyzing streaming data using standard SQL.
//
// For more details, see the [Amazon Kinesis Analytics Documentation][1].
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kinesis_analytics_application.html.markdown.
type AnalyticsApplication struct {
	pulumi.CustomResourceState

	// The ARN of the Kinesis Analytics Appliation.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions AnalyticsApplicationCloudwatchLoggingOptionsPtrOutput `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringPtrOutput `pulumi:"code"`
	// The Timestamp when the application version was created.
	CreateTimestamp pulumi.StringOutput `pulumi:"createTimestamp"`
	// Description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs AnalyticsApplicationInputsPtrOutput `pulumi:"inputs"`
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringOutput `pulumi:"lastUpdateTimestamp"`
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringOutput `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs AnalyticsApplicationOutputArrayOutput `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources AnalyticsApplicationReferenceDataSourcesPtrOutput `pulumi:"referenceDataSources"`
	// The Status of the application.
	Status pulumi.StringOutput `pulumi:"status"`
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The Version of the application.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewAnalyticsApplication registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsApplication(ctx *pulumi.Context,
	name string, args *AnalyticsApplicationArgs, opts ...pulumi.ResourceOption) (*AnalyticsApplication, error) {
	if args == nil {
		args = &AnalyticsApplicationArgs{}
	}
	var resource AnalyticsApplication
	err := ctx.RegisterResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsApplication gets an existing AnalyticsApplication resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsApplicationState, opts ...pulumi.ResourceOption) (*AnalyticsApplication, error) {
	var resource AnalyticsApplication
	err := ctx.ReadResource("aws:kinesis/analyticsApplication:AnalyticsApplication", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsApplication resources.
type analyticsApplicationState struct {
	// The ARN of the Kinesis Analytics Appliation.
	Arn *string `pulumi:"arn"`
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions *AnalyticsApplicationCloudwatchLoggingOptions `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code *string `pulumi:"code"`
	// The Timestamp when the application version was created.
	CreateTimestamp *string `pulumi:"createTimestamp"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs *AnalyticsApplicationInputs `pulumi:"inputs"`
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp *string `pulumi:"lastUpdateTimestamp"`
	// Name of the Kinesis Analytics Application.
	Name *string `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs []AnalyticsApplicationOutput `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources *AnalyticsApplicationReferenceDataSources `pulumi:"referenceDataSources"`
	// The Status of the application.
	Status *string `pulumi:"status"`
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags map[string]interface{} `pulumi:"tags"`
	// The Version of the application.
	Version *int `pulumi:"version"`
}

type AnalyticsApplicationState struct {
	// The ARN of the Kinesis Analytics Appliation.
	Arn pulumi.StringPtrInput
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions AnalyticsApplicationCloudwatchLoggingOptionsPtrInput
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringPtrInput
	// The Timestamp when the application version was created.
	CreateTimestamp pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// Input configuration of the application. See Inputs below for more details.
	Inputs AnalyticsApplicationInputsPtrInput
	// The Timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringPtrInput
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringPtrInput
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs AnalyticsApplicationOutputArrayInput
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources AnalyticsApplicationReferenceDataSourcesPtrInput
	// The Status of the application.
	Status pulumi.StringPtrInput
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags pulumi.MapInput
	// The Version of the application.
	Version pulumi.IntPtrInput
}

func (AnalyticsApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsApplicationState)(nil)).Elem()
}

type analyticsApplicationArgs struct {
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions *AnalyticsApplicationCloudwatchLoggingOptions `pulumi:"cloudwatchLoggingOptions"`
	// SQL Code to transform input data, and generate output.
	Code *string `pulumi:"code"`
	// Description of the application.
	Description *string `pulumi:"description"`
	// Input configuration of the application. See Inputs below for more details.
	Inputs *AnalyticsApplicationInputs `pulumi:"inputs"`
	// Name of the Kinesis Analytics Application.
	Name *string `pulumi:"name"`
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs []AnalyticsApplicationOutput `pulumi:"outputs"`
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources *AnalyticsApplicationReferenceDataSources `pulumi:"referenceDataSources"`
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a AnalyticsApplication resource.
type AnalyticsApplicationArgs struct {
	// The CloudWatch log stream options to monitor application errors.
	// See CloudWatch Logging Options below for more details.
	CloudwatchLoggingOptions AnalyticsApplicationCloudwatchLoggingOptionsPtrInput
	// SQL Code to transform input data, and generate output.
	Code pulumi.StringPtrInput
	// Description of the application.
	Description pulumi.StringPtrInput
	// Input configuration of the application. See Inputs below for more details.
	Inputs AnalyticsApplicationInputsPtrInput
	// Name of the Kinesis Analytics Application.
	Name pulumi.StringPtrInput
	// Output destination configuration of the application. See Outputs below for more details.
	Outputs AnalyticsApplicationOutputArrayInput
	// An S3 Reference Data Source for the application.
	// See Reference Data Sources below for more details.
	ReferenceDataSources AnalyticsApplicationReferenceDataSourcesPtrInput
	// Key-value mapping of tags for the Kinesis Analytics Application.
	Tags pulumi.MapInput
}

func (AnalyticsApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsApplicationArgs)(nil)).Elem()
}

