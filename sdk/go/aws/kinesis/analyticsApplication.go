// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Kinesis Analytics Application resource. Kinesis Analytics is a managed service that
// allows processing and analyzing streaming data using standard SQL.
//
// For more details, see the [Amazon Kinesis Analytics Documentation](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/what-is.html).
//
// > **Note:** To manage Amazon Kinesis Data Analytics for Apache Flink applications, use the [`kinesisanalyticsv2.Application`](https://www.terraform.io/docs/providers/aws/r/kinesisanalyticsv2_application.html) resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kinesis"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
// 			ShardCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kinesis.NewAnalyticsApplication(ctx, "testApplication", &kinesis.AnalyticsApplicationArgs{
// 			Inputs: &kinesis.AnalyticsApplicationInputsArgs{
// 				NamePrefix: pulumi.String("test_prefix"),
// 				KinesisStream: &kinesis.AnalyticsApplicationInputsKinesisStreamArgs{
// 					ResourceArn: testStream.Arn,
// 					RoleArn:     pulumi.Any(aws_iam_role.Test.Arn),
// 				},
// 				Parallelism: &kinesis.AnalyticsApplicationInputsParallelismArgs{
// 					Count: pulumi.Int(1),
// 				},
// 				Schema: &kinesis.AnalyticsApplicationInputsSchemaArgs{
// 					RecordColumns: kinesis.AnalyticsApplicationInputsSchemaRecordColumnArray{
// 						&kinesis.AnalyticsApplicationInputsSchemaRecordColumnArgs{
// 							Mapping: pulumi.String(fmt.Sprintf("%v%v", "$", ".test")),
// 							Name:    pulumi.String("test"),
// 							SqlType: pulumi.String("VARCHAR(8)"),
// 						},
// 					},
// 					RecordEncoding: pulumi.String("UTF-8"),
// 					RecordFormat: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatArgs{
// 						MappingParameters: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersArgs{
// 							Json: &kinesis.AnalyticsApplicationInputsSchemaRecordFormatMappingParametersJsonArgs{
// 								RecordRowPath: pulumi.String("$"),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	// Key-value map of tags for the Kinesis Analytics Application.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
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
	// Key-value map of tags for the Kinesis Analytics Application.
	Tags map[string]string `pulumi:"tags"`
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
	// Key-value map of tags for the Kinesis Analytics Application.
	Tags pulumi.StringMapInput
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
	// Key-value map of tags for the Kinesis Analytics Application.
	Tags map[string]string `pulumi:"tags"`
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
	// Key-value map of tags for the Kinesis Analytics Application.
	Tags pulumi.StringMapInput
}

func (AnalyticsApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsApplicationArgs)(nil)).Elem()
}

type AnalyticsApplicationInput interface {
	pulumi.Input

	ToAnalyticsApplicationOutput() AnalyticsApplicationOutput
	ToAnalyticsApplicationOutputWithContext(ctx context.Context) AnalyticsApplicationOutput
}

func (AnalyticsApplication) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsApplication)(nil)).Elem()
}

func (i AnalyticsApplication) ToAnalyticsApplicationOutput() AnalyticsApplicationOutput {
	return i.ToAnalyticsApplicationOutputWithContext(context.Background())
}

func (i AnalyticsApplication) ToAnalyticsApplicationOutputWithContext(ctx context.Context) AnalyticsApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsApplicationOutput)
}

type AnalyticsApplicationOutput struct {
	*pulumi.OutputState
}

func (AnalyticsApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsApplicationOutput)(nil)).Elem()
}

func (o AnalyticsApplicationOutput) ToAnalyticsApplicationOutput() AnalyticsApplicationOutput {
	return o
}

func (o AnalyticsApplicationOutput) ToAnalyticsApplicationOutputWithContext(ctx context.Context) AnalyticsApplicationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AnalyticsApplicationOutput{})
}
