// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudTrail resource.
//
// > *NOTE:* For a multi-region trail, this resource must be in the home region of the trail.
//
// > *NOTE:* For an organization trail, this resource must be in the master account of the organization.
//
// ## Example Usage
// ### Basic
//
// Enable CloudTrail to capture all compatible management events in region.
// For capturing events from services like IAM, `includeGlobalServiceEvents` must be enabled.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := s3.NewBucket(ctx, "foo", &s3.BucketArgs{
// 			ForceDestroy: pulumi.Bool(true),
// 			Policy:       pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Sid\": \"AWSCloudTrailAclCheck\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "              \"Service\": \"cloudtrail.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:GetBucketAcl\",\n", "            \"Resource\": \"arn:aws:s3:::tf-test-trail\"\n", "        },\n", "        {\n", "            \"Sid\": \"AWSCloudTrailWrite\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "              \"Service\": \"cloudtrail.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:PutObject\",\n", "            \"Resource\": \"arn:aws:s3:::tf-test-trail/prefix/AWSLogs/", current.AccountId, "/*\",\n", "            \"Condition\": {\n", "                \"StringEquals\": {\n", "                    \"s3:x-amz-acl\": \"bucket-owner-full-control\"\n", "                }\n", "            }\n", "        }\n", "    ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "foobar", &cloudtrail.TrailArgs{
// 			S3BucketName:               foo.ID(),
// 			S3KeyPrefix:                pulumi.String("prefix"),
// 			IncludeGlobalServiceEvents: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Data Event Logging
//
// CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the [CloudTrail API DataResource documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_DataResource.html).
// ### Logging All Lambda Function Invocations
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::Lambda::Function"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("arn:aws:lambda"),
// 							},
// 						},
// 					},
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					ReadWriteType:           pulumi.String("All"),
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
// ### Logging All S3 Bucket Object Events
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::S3::Object"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("arn:aws:s3:::"),
// 							},
// 						},
// 					},
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					ReadWriteType:           pulumi.String("All"),
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
// ### Logging Individual S3 Bucket Events
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		important_bucket, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "important-bucket",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::S3::Object"),
// 							Values: pulumi.StringArray{
// 								pulumi.String(fmt.Sprintf("%v%v", important_bucket.Arn, "/")),
// 							},
// 						},
// 					},
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					ReadWriteType:           pulumi.String("All"),
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
// ### Sending Events to CloudWatch Logs
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "exampleTrail", &cloudtrail.TrailArgs{
// 			CloudWatchLogsGroupArn: exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v", arn, ":*"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Trail struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name of the trail.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrOutput `pulumi:"cloudWatchLogsGroupArn"`
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrOutput `pulumi:"cloudWatchLogsRoleArn"`
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrOutput `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors TrailEventSelectorArrayOutput `pulumi:"eventSelectors"`
	// The region in which the trail was created.
	HomeRegion pulumi.StringOutput `pulumi:"homeRegion"`
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrOutput `pulumi:"includeGlobalServiceEvents"`
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors TrailInsightSelectorArrayOutput `pulumi:"insightSelectors"`
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrOutput `pulumi:"isMultiRegionTrail"`
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrOutput `pulumi:"isOrganizationTrail"`
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Specifies the name of the trail.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringOutput `pulumi:"s3BucketName"`
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrOutput `pulumi:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrOutput `pulumi:"snsTopicName"`
	// A map of tags to assign to the trail
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewTrail registers a new resource with the given unique name, arguments, and options.
func NewTrail(ctx *pulumi.Context,
	name string, args *TrailArgs, opts ...pulumi.ResourceOption) (*Trail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.S3BucketName == nil {
		return nil, errors.New("invalid value for required argument 'S3BucketName'")
	}
	var resource Trail
	err := ctx.RegisterResource("aws:cloudtrail/trail:Trail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrail gets an existing Trail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrailState, opts ...pulumi.ResourceOption) (*Trail, error) {
	var resource Trail
	err := ctx.ReadResource("aws:cloudtrail/trail:Trail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trail resources.
type trailState struct {
	// The Amazon Resource Name of the trail.
	Arn *string `pulumi:"arn"`
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `pulumi:"cloudWatchLogsGroupArn"`
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `pulumi:"cloudWatchLogsRoleArn"`
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation *bool `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors []TrailEventSelector `pulumi:"eventSelectors"`
	// The region in which the trail was created.
	HomeRegion *string `pulumi:"homeRegion"`
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents *bool `pulumi:"includeGlobalServiceEvents"`
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors []TrailInsightSelector `pulumi:"insightSelectors"`
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail *bool `pulumi:"isMultiRegionTrail"`
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail *bool `pulumi:"isOrganizationTrail"`
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the name of the trail.
	Name *string `pulumi:"name"`
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName *string `pulumi:"s3BucketName"`
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName *string `pulumi:"snsTopicName"`
	// A map of tags to assign to the trail
	Tags map[string]string `pulumi:"tags"`
}

type TrailState struct {
	// The Amazon Resource Name of the trail.
	Arn pulumi.StringPtrInput
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrInput
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrInput
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrInput
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrInput
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors TrailEventSelectorArrayInput
	// The region in which the trail was created.
	HomeRegion pulumi.StringPtrInput
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrInput
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors TrailInsightSelectorArrayInput
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrInput
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrInput
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the name of the trail.
	Name pulumi.StringPtrInput
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringPtrInput
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrInput
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrInput
	// A map of tags to assign to the trail
	Tags pulumi.StringMapInput
}

func (TrailState) ElementType() reflect.Type {
	return reflect.TypeOf((*trailState)(nil)).Elem()
}

type trailArgs struct {
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `pulumi:"cloudWatchLogsGroupArn"`
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `pulumi:"cloudWatchLogsRoleArn"`
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation *bool `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors []TrailEventSelector `pulumi:"eventSelectors"`
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents *bool `pulumi:"includeGlobalServiceEvents"`
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors []TrailInsightSelector `pulumi:"insightSelectors"`
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail *bool `pulumi:"isMultiRegionTrail"`
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail *bool `pulumi:"isOrganizationTrail"`
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the name of the trail.
	Name *string `pulumi:"name"`
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName string `pulumi:"s3BucketName"`
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName *string `pulumi:"snsTopicName"`
	// A map of tags to assign to the trail
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Trail resource.
type TrailArgs struct {
	// Specifies a log group name using an Amazon Resource Name (ARN),
	// that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrInput
	// Specifies the role for the CloudWatch Logs
	// endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrInput
	// Specifies whether log file integrity validation is enabled.
	// Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrInput
	// Enables logging for the trail. Defaults to `true`.
	// Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrInput
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
	EventSelectors TrailEventSelectorArrayInput
	// Specifies whether the trail is publishing events
	// from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrInput
	// Specifies an insight selector for identifying unusual operational activity. Fields documented below.
	InsightSelectors TrailInsightSelectorArrayInput
	// Specifies whether the trail is created in the current
	// region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrInput
	// Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrInput
	// Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the name of the trail.
	Name pulumi.StringPtrInput
	// Specifies the name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringInput
	// Specifies the S3 key prefix that follows
	// the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrInput
	// Specifies the name of the Amazon SNS topic
	// defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrInput
	// A map of tags to assign to the trail
	Tags pulumi.StringMapInput
}

func (TrailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trailArgs)(nil)).Elem()
}

type TrailInput interface {
	pulumi.Input

	ToTrailOutput() TrailOutput
	ToTrailOutputWithContext(ctx context.Context) TrailOutput
}

func (Trail) ElementType() reflect.Type {
	return reflect.TypeOf((*Trail)(nil)).Elem()
}

func (i Trail) ToTrailOutput() TrailOutput {
	return i.ToTrailOutputWithContext(context.Background())
}

func (i Trail) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailOutput)
}

type TrailOutput struct {
	*pulumi.OutputState
}

func (TrailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailOutput)(nil)).Elem()
}

func (o TrailOutput) ToTrailOutput() TrailOutput {
	return o
}

func (o TrailOutput) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TrailOutput{})
}
