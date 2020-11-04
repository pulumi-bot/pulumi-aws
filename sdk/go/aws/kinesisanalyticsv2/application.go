// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisanalyticsv2

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Kinesis Analytics v2 Application.
// This resource can be used to manage both Kinesis Data Analytics for SQL applications and Kinesis Data Analytics for Apache Flink applications.
//
// > **Note:** Kinesis Data Analytics for SQL applications created using this resource cannot currently be viewed in the AWS Console. To manage Kinesis Data Analytics for SQL applications that can also be viewed in the AWS Console, use the [`kinesis.AnalyticsApplication`](https://www.terraform.io/docs/providers/aws/r/kinesis_analytics_application.html) resource.
//
// ## Example Usage
// ### Apache Flink Application
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kinesisanalyticsv2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleBucket, err := s3.NewBucket(ctx, "exampleBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleBucketObject, err := s3.NewBucketObject(ctx, "exampleBucketObject", &s3.BucketObjectArgs{
// 			Bucket: exampleBucket.Bucket,
// 			Key:    pulumi.String("example-flink-application"),
// 			Source: pulumi.NewFileAsset("flink-app.jar"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kinesisanalyticsv2.NewApplication(ctx, "exampleApplication", &kinesisanalyticsv2.ApplicationArgs{
// 			RuntimeEnvironment:   pulumi.String("FLINK-1_8"),
// 			ServiceExecutionRole: pulumi.Any(aws_iam_role.Example.Arn),
// 			ApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationArgs{
// 				ApplicationCodeConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs{
// 					CodeContent: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs{
// 						S3ContentLocation: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentS3ContentLocationArgs{
// 							BucketArn: exampleBucket.Arn,
// 							FileKey:   exampleBucketObject.Key,
// 						},
// 					},
// 					CodeContentType: pulumi.String("ZIPFILE"),
// 				},
// 				EnvironmentProperties: &kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesArgs{
// 					PropertyGroups: kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArray{
// 						&kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs{
// 							PropertyGroupId: pulumi.String("PROPERTY-GROUP-1"),
// 							PropertyMap: pulumi.StringMap{
// 								"Key1": pulumi.String("Value1"),
// 							},
// 						},
// 						&kinesisanalyticsv2.ApplicationApplicationConfigurationEnvironmentPropertiesPropertyGroupArgs{
// 							PropertyGroupId: pulumi.String("PROPERTY-GROUP-2"),
// 							PropertyMap: pulumi.StringMap{
// 								"KeyA": pulumi.String("ValueA"),
// 								"KeyB": pulumi.String("ValueB"),
// 							},
// 						},
// 					},
// 				},
// 				FlinkApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationArgs{
// 					CheckpointConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationCheckpointConfigurationArgs{
// 						ConfigurationType: pulumi.String("DEFAULT"),
// 					},
// 					MonitoringConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationMonitoringConfigurationArgs{
// 						ConfigurationType: pulumi.String("CUSTOM"),
// 						LogLevel:          pulumi.String("DEBUG"),
// 						MetricsLevel:      pulumi.String("TASK"),
// 					},
// 					ParallelismConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationFlinkApplicationConfigurationParallelismConfigurationArgs{
// 						AutoScalingEnabled: pulumi.Bool(true),
// 						ConfigurationType:  pulumi.String("CUSTOM"),
// 						Parallelism:        pulumi.Int(10),
// 						ParallelismPerKpu:  pulumi.Int(4),
// 					},
// 				},
// 			},
// 			Tags: pulumi.StringMap{
// 				"Environment": pulumi.String("test"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### SQL Application
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/kinesisanalyticsv2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleLogStream, err := cloudwatch.NewLogStream(ctx, "exampleLogStream", &cloudwatch.LogStreamArgs{
// 			LogGroupName: exampleLogGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kinesisanalyticsv2.NewApplication(ctx, "exampleApplication", &kinesisanalyticsv2.ApplicationArgs{
// 			RuntimeEnvironment:   pulumi.String("SQL-1.0"),
// 			ServiceExecutionRole: pulumi.Any(aws_iam_role.Example.Arn),
// 			ApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationArgs{
// 				ApplicationCodeConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationArgs{
// 					CodeContent: &kinesisanalyticsv2.ApplicationApplicationConfigurationApplicationCodeConfigurationCodeContentArgs{
// 						TextContent: pulumi.String("SELECT 1;\n"),
// 					},
// 					CodeContentType: pulumi.String("PLAINTEXT"),
// 				},
// 				SqlApplicationConfiguration: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationArgs{
// 					Input: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputArgs{
// 						NamePrefix: pulumi.String("PREFIX_1"),
// 						InputParallelism: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputParallelismArgs{
// 							Count: pulumi.Int(3),
// 						},
// 						InputSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaArgs{
// 							RecordColumns: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArray{
// 								&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs{
// 									Name:    pulumi.String("COLUMN_1"),
// 									SqlType: pulumi.String("VARCHAR(8)"),
// 									Mapping: pulumi.String("MAPPING-1"),
// 								},
// 								&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordColumnArgs{
// 									Name:    pulumi.String("COLUMN_2"),
// 									SqlType: pulumi.String("DOUBLE"),
// 								},
// 							},
// 							RecordEncoding: pulumi.String("UTF-8"),
// 							RecordFormat: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatArgs{
// 								RecordFormatType: pulumi.String("CSV"),
// 								MappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersArgs{
// 									CsvMappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputSchemaRecordFormatMappingParametersCsvMappingParametersArgs{
// 										RecordColumnDelimiter: pulumi.String(","),
// 										RecordRowDelimiter:    pulumi.String("\n"),
// 									},
// 								},
// 							},
// 						},
// 						KinesisStreamsInput: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationInputKinesisStreamsInputArgs{
// 							ResourceArn: pulumi.Any(aws_kinesis_stream.Example.Arn),
// 						},
// 					},
// 					Outputs: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArray{
// 						&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs{
// 							Name: pulumi.String("OUTPUT_1"),
// 							DestinationSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs{
// 								RecordFormatType: pulumi.String("JSON"),
// 							},
// 							LambdaOutput: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputLambdaOutputArgs{
// 								ResourceArn: pulumi.Any(aws_lambda_function.Example.Arn),
// 							},
// 						},
// 						&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputArgs{
// 							Name: pulumi.String("OUTPUT_2"),
// 							DestinationSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputDestinationSchemaArgs{
// 								RecordFormatType: pulumi.String("CSV"),
// 							},
// 							KinesisFirehoseOutput: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationOutputKinesisFirehoseOutputArgs{
// 								ResourceArn: pulumi.Any(aws_kinesis_firehose_delivery_stream.Example.Arn),
// 							},
// 						},
// 					},
// 					ReferenceDataSource: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceArgs{
// 						TableName: pulumi.String("TABLE-1"),
// 						ReferenceSchema: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaArgs{
// 							RecordColumns: kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArray{
// 								&kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordColumnArgs{
// 									Name:    pulumi.String("COLUMN_1"),
// 									SqlType: pulumi.String("INTEGER"),
// 								},
// 							},
// 							RecordFormat: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatArgs{
// 								RecordFormatType: pulumi.String("JSON"),
// 								MappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersArgs{
// 									JsonMappingParameters: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceReferenceSchemaRecordFormatMappingParametersJsonMappingParametersArgs{
// 										RecordRowPath: pulumi.String("$"),
// 									},
// 								},
// 							},
// 						},
// 						S3ReferenceDataSource: &kinesisanalyticsv2.ApplicationApplicationConfigurationSqlApplicationConfigurationReferenceDataSourceS3ReferenceDataSourceArgs{
// 							BucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
// 							FileKey:   pulumi.String("KEY-1"),
// 						},
// 					},
// 				},
// 			},
// 			CloudwatchLoggingOptions: &kinesisanalyticsv2.ApplicationCloudwatchLoggingOptionsArgs{
// 				LogStreamArn: exampleLogStream.Arn,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Application struct {
	pulumi.CustomResourceState

	// The application's configuration
	ApplicationConfiguration ApplicationApplicationConfigurationOutput `pulumi:"applicationConfiguration"`
	// The ARN of the application.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
	CloudwatchLoggingOptions ApplicationCloudwatchLoggingOptionsPtrOutput `pulumi:"cloudwatchLoggingOptions"`
	// The current timestamp when the application was created.
	CreateTimestamp pulumi.StringOutput `pulumi:"createTimestamp"`
	// A summary description of the application.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The current timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringOutput `pulumi:"lastUpdateTimestamp"`
	// The name of the application.
	Name pulumi.StringOutput `pulumi:"name"`
	// The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
	RuntimeEnvironment pulumi.StringOutput `pulumi:"runtimeEnvironment"`
	// The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	ServiceExecutionRole pulumi.StringOutput `pulumi:"serviceExecutionRole"`
	// The status of the application.
	Status pulumi.StringOutput `pulumi:"status"`
	// A map of tags to assign to the application.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The current application version. Kinesis Data Analytics updates the `versionId` each time the application is updated.
	VersionId pulumi.IntOutput `pulumi:"versionId"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.RuntimeEnvironment == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeEnvironment'")
	}
	if args.ServiceExecutionRole == nil {
		return nil, errors.New("invalid value for required argument 'ServiceExecutionRole'")
	}
	var resource Application
	err := ctx.RegisterResource("aws:kinesisanalyticsv2/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws:kinesisanalyticsv2/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The application's configuration
	ApplicationConfiguration *ApplicationApplicationConfiguration `pulumi:"applicationConfiguration"`
	// The ARN of the application.
	Arn *string `pulumi:"arn"`
	// A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
	CloudwatchLoggingOptions *ApplicationCloudwatchLoggingOptions `pulumi:"cloudwatchLoggingOptions"`
	// The current timestamp when the application was created.
	CreateTimestamp *string `pulumi:"createTimestamp"`
	// A summary description of the application.
	Description *string `pulumi:"description"`
	// The current timestamp when the application was last updated.
	LastUpdateTimestamp *string `pulumi:"lastUpdateTimestamp"`
	// The name of the application.
	Name *string `pulumi:"name"`
	// The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
	RuntimeEnvironment *string `pulumi:"runtimeEnvironment"`
	// The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	ServiceExecutionRole *string `pulumi:"serviceExecutionRole"`
	// The status of the application.
	Status *string `pulumi:"status"`
	// A map of tags to assign to the application.
	Tags map[string]string `pulumi:"tags"`
	// The current application version. Kinesis Data Analytics updates the `versionId` each time the application is updated.
	VersionId *int `pulumi:"versionId"`
}

type ApplicationState struct {
	// The application's configuration
	ApplicationConfiguration ApplicationApplicationConfigurationPtrInput
	// The ARN of the application.
	Arn pulumi.StringPtrInput
	// A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
	CloudwatchLoggingOptions ApplicationCloudwatchLoggingOptionsPtrInput
	// The current timestamp when the application was created.
	CreateTimestamp pulumi.StringPtrInput
	// A summary description of the application.
	Description pulumi.StringPtrInput
	// The current timestamp when the application was last updated.
	LastUpdateTimestamp pulumi.StringPtrInput
	// The name of the application.
	Name pulumi.StringPtrInput
	// The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
	RuntimeEnvironment pulumi.StringPtrInput
	// The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	ServiceExecutionRole pulumi.StringPtrInput
	// The status of the application.
	Status pulumi.StringPtrInput
	// A map of tags to assign to the application.
	Tags pulumi.StringMapInput
	// The current application version. Kinesis Data Analytics updates the `versionId` each time the application is updated.
	VersionId pulumi.IntPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The application's configuration
	ApplicationConfiguration *ApplicationApplicationConfiguration `pulumi:"applicationConfiguration"`
	// A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
	CloudwatchLoggingOptions *ApplicationCloudwatchLoggingOptions `pulumi:"cloudwatchLoggingOptions"`
	// A summary description of the application.
	Description *string `pulumi:"description"`
	// The name of the application.
	Name *string `pulumi:"name"`
	// The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
	RuntimeEnvironment string `pulumi:"runtimeEnvironment"`
	// The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	ServiceExecutionRole string `pulumi:"serviceExecutionRole"`
	// A map of tags to assign to the application.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The application's configuration
	ApplicationConfiguration ApplicationApplicationConfigurationPtrInput
	// A [CloudWatch log stream](https://www.terraform.io/docs/providers/aws/r/cloudwatch_log_stream.html) to monitor application configuration errors.
	CloudwatchLoggingOptions ApplicationCloudwatchLoggingOptionsPtrInput
	// A summary description of the application.
	Description pulumi.StringPtrInput
	// The name of the application.
	Name pulumi.StringPtrInput
	// The runtime environment for the application. Valid values: `SQL-1_0`, `FLINK-1_6`, `FLINK-1_8`.
	RuntimeEnvironment pulumi.StringInput
	// The ARN of the [IAM role](https://www.terraform.io/docs/providers/aws/r/iam_role.html) used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	ServiceExecutionRole pulumi.StringInput
	// A map of tags to assign to the application.
	Tags pulumi.StringMapInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}
