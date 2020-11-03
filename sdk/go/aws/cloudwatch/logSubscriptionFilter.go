// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudWatch Logs subscription filter resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudwatch.NewLogSubscriptionFilter(ctx, "testLambdafunctionLogfilter", &cloudwatch.LogSubscriptionFilterArgs{
// 			RoleArn:        pulumi.Any(aws_iam_role.Iam_for_lambda.Arn),
// 			LogGroup:       pulumi.String("/aws/lambda/example_lambda_name"),
// 			FilterPattern:  pulumi.String("logtype test"),
// 			DestinationArn: pulumi.Any(aws_kinesis_stream.Test_logstream.Arn),
// 			Distribution:   pulumi.String("Random"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LogSubscriptionFilter struct {
	pulumi.CustomResourceState

	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn pulumi.StringOutput `pulumi:"destinationArn"`
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution pulumi.StringPtrOutput `pulumi:"distribution"`
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern pulumi.StringOutput `pulumi:"filterPattern"`
	// The name of the log group to associate the subscription filter with
	LogGroup pulumi.StringOutput `pulumi:"logGroup"`
	// A name for the subscription filter
	Name pulumi.StringOutput `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewLogSubscriptionFilter registers a new resource with the given unique name, arguments, and options.
func NewLogSubscriptionFilter(ctx *pulumi.Context,
	name string, args *LogSubscriptionFilterArgs, opts ...pulumi.ResourceOption) (*LogSubscriptionFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.DestinationArn == nil {
		return nil, errors.New("invalid value for required argument 'DestinationArn'")
	}
	if args.FilterPattern == nil {
		return nil, errors.New("invalid value for required argument 'FilterPattern'")
	}
	if args.LogGroup == nil {
		return nil, errors.New("invalid value for required argument 'LogGroup'")
	}
	var resource LogSubscriptionFilter
	err := ctx.RegisterResource("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogSubscriptionFilter gets an existing LogSubscriptionFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogSubscriptionFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogSubscriptionFilterState, opts ...pulumi.ResourceOption) (*LogSubscriptionFilter, error) {
	var resource LogSubscriptionFilter
	err := ctx.ReadResource("aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogSubscriptionFilter resources.
type logSubscriptionFilterState struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn *string `pulumi:"destinationArn"`
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution *string `pulumi:"distribution"`
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern *string `pulumi:"filterPattern"`
	// The name of the log group to associate the subscription filter with
	LogGroup *string `pulumi:"logGroup"`
	// A name for the subscription filter
	Name *string `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
	RoleArn *string `pulumi:"roleArn"`
}

type LogSubscriptionFilterState struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn pulumi.StringPtrInput
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution pulumi.StringPtrInput
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern pulumi.StringPtrInput
	// The name of the log group to associate the subscription filter with
	LogGroup pulumi.StringPtrInput
	// A name for the subscription filter
	Name pulumi.StringPtrInput
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
	RoleArn pulumi.StringPtrInput
}

func (LogSubscriptionFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*logSubscriptionFilterState)(nil)).Elem()
}

type logSubscriptionFilterArgs struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn string `pulumi:"destinationArn"`
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution *string `pulumi:"distribution"`
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern string `pulumi:"filterPattern"`
	// The name of the log group to associate the subscription filter with
	LogGroup interface{} `pulumi:"logGroup"`
	// A name for the subscription filter
	Name *string `pulumi:"name"`
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
	RoleArn *string `pulumi:"roleArn"`
}

// The set of arguments for constructing a LogSubscriptionFilter resource.
type LogSubscriptionFilterArgs struct {
	// The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
	DestinationArn pulumi.StringInput
	// The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
	Distribution pulumi.StringPtrInput
	// A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
	FilterPattern pulumi.StringInput
	// The name of the log group to associate the subscription filter with
	LogGroup pulumi.Input
	// A name for the subscription filter
	Name pulumi.StringPtrInput
	// The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
	RoleArn pulumi.StringPtrInput
}

func (LogSubscriptionFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logSubscriptionFilterArgs)(nil)).Elem()
}
