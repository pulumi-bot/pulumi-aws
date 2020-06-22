// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Lambda event source mapping. This allows Lambda functions to get events from Kinesis, DynamoDB and SQS.
//
// For information about Lambda and how to use it, see [What is AWS Lambda?](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html).
// For information about event source mappings, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the API docs.
//
// ## Example Usage
// ### DynamoDB
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
// 			EventSourceArn:   pulumi.String(aws_dynamodb_table.Example.Stream_arn),
// 			FunctionName:     pulumi.String(aws_lambda_function.Example.Arn),
// 			StartingPosition: pulumi.String("LATEST"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Kinesis
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
// 			EventSourceArn:   pulumi.String(aws_kinesis_stream.Example.Arn),
// 			FunctionName:     pulumi.String(aws_lambda_function.Example.Arn),
// 			StartingPosition: pulumi.String("LATEST"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### SQS
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
// 			EventSourceArn: pulumi.String(aws_sqs_queue.Sqs_queue_test.Arn),
// 			FunctionName:   pulumi.String(aws_lambda_function.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type EventSourceMapping struct {
	pulumi.CustomResourceState

	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
	BatchSize                  pulumi.IntPtrOutput                          `pulumi:"batchSize"`
	BisectBatchOnFunctionError pulumi.BoolPtrOutput                         `pulumi:"bisectBatchOnFunctionError"`
	DestinationConfig          EventSourceMappingDestinationConfigPtrOutput `pulumi:"destinationConfig"`
	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
	EventSourceArn pulumi.StringOutput `pulumi:"eventSourceArn"`
	// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `functionName` above.)
	FunctionArn pulumi.StringOutput `pulumi:"functionArn"`
	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// The date this resource was last modified.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult pulumi.StringOutput `pulumi:"lastProcessingResult"`
	// The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
	MaximumBatchingWindowInSeconds pulumi.IntPtrOutput `pulumi:"maximumBatchingWindowInSeconds"`
	MaximumRecordAgeInSeconds      pulumi.IntOutput    `pulumi:"maximumRecordAgeInSeconds"`
	MaximumRetryAttempts           pulumi.IntOutput    `pulumi:"maximumRetryAttempts"`
	ParallelizationFactor          pulumi.IntOutput    `pulumi:"parallelizationFactor"`
	// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
	StartingPosition pulumi.StringPtrOutput `pulumi:"startingPosition"`
	// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	// * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	// * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
	// * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
	// * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
	// * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
	StartingPositionTimestamp pulumi.StringPtrOutput `pulumi:"startingPositionTimestamp"`
	// The state of the event source mapping.
	State pulumi.StringOutput `pulumi:"state"`
	// The reason the event source mapping is in its current state.
	StateTransitionReason pulumi.StringOutput `pulumi:"stateTransitionReason"`
	// The UUID of the created event source mapping.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewEventSourceMapping registers a new resource with the given unique name, arguments, and options.
func NewEventSourceMapping(ctx *pulumi.Context,
	name string, args *EventSourceMappingArgs, opts ...pulumi.ResourceOption) (*EventSourceMapping, error) {
	if args == nil || args.EventSourceArn == nil {
		return nil, errors.New("missing required argument 'EventSourceArn'")
	}
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil {
		args = &EventSourceMappingArgs{}
	}
	var resource EventSourceMapping
	err := ctx.RegisterResource("aws:lambda/eventSourceMapping:EventSourceMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSourceMapping gets an existing EventSourceMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSourceMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSourceMappingState, opts ...pulumi.ResourceOption) (*EventSourceMapping, error) {
	var resource EventSourceMapping
	err := ctx.ReadResource("aws:lambda/eventSourceMapping:EventSourceMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSourceMapping resources.
type eventSourceMappingState struct {
	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
	BatchSize                  *int                                 `pulumi:"batchSize"`
	BisectBatchOnFunctionError *bool                                `pulumi:"bisectBatchOnFunctionError"`
	DestinationConfig          *EventSourceMappingDestinationConfig `pulumi:"destinationConfig"`
	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
	EventSourceArn *string `pulumi:"eventSourceArn"`
	// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `functionName` above.)
	FunctionArn *string `pulumi:"functionArn"`
	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName *string `pulumi:"functionName"`
	// The date this resource was last modified.
	LastModified *string `pulumi:"lastModified"`
	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string `pulumi:"lastProcessingResult"`
	// The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
	MaximumBatchingWindowInSeconds *int `pulumi:"maximumBatchingWindowInSeconds"`
	MaximumRecordAgeInSeconds      *int `pulumi:"maximumRecordAgeInSeconds"`
	MaximumRetryAttempts           *int `pulumi:"maximumRetryAttempts"`
	ParallelizationFactor          *int `pulumi:"parallelizationFactor"`
	// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
	StartingPosition *string `pulumi:"startingPosition"`
	// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	// * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	// * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
	// * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
	// * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
	// * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
	StartingPositionTimestamp *string `pulumi:"startingPositionTimestamp"`
	// The state of the event source mapping.
	State *string `pulumi:"state"`
	// The reason the event source mapping is in its current state.
	StateTransitionReason *string `pulumi:"stateTransitionReason"`
	// The UUID of the created event source mapping.
	Uuid *string `pulumi:"uuid"`
}

type EventSourceMappingState struct {
	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
	BatchSize                  pulumi.IntPtrInput
	BisectBatchOnFunctionError pulumi.BoolPtrInput
	DestinationConfig          EventSourceMappingDestinationConfigPtrInput
	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
	EventSourceArn pulumi.StringPtrInput
	// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `functionName` above.)
	FunctionArn pulumi.StringPtrInput
	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName pulumi.StringPtrInput
	// The date this resource was last modified.
	LastModified pulumi.StringPtrInput
	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult pulumi.StringPtrInput
	// The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
	MaximumBatchingWindowInSeconds pulumi.IntPtrInput
	MaximumRecordAgeInSeconds      pulumi.IntPtrInput
	MaximumRetryAttempts           pulumi.IntPtrInput
	ParallelizationFactor          pulumi.IntPtrInput
	// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
	StartingPosition pulumi.StringPtrInput
	// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	// * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	// * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
	// * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
	// * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
	// * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
	StartingPositionTimestamp pulumi.StringPtrInput
	// The state of the event source mapping.
	State pulumi.StringPtrInput
	// The reason the event source mapping is in its current state.
	StateTransitionReason pulumi.StringPtrInput
	// The UUID of the created event source mapping.
	Uuid pulumi.StringPtrInput
}

func (EventSourceMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSourceMappingState)(nil)).Elem()
}

type eventSourceMappingArgs struct {
	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
	BatchSize                  *int                                 `pulumi:"batchSize"`
	BisectBatchOnFunctionError *bool                                `pulumi:"bisectBatchOnFunctionError"`
	DestinationConfig          *EventSourceMappingDestinationConfig `pulumi:"destinationConfig"`
	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
	EventSourceArn string `pulumi:"eventSourceArn"`
	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName string `pulumi:"functionName"`
	// The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
	MaximumBatchingWindowInSeconds *int `pulumi:"maximumBatchingWindowInSeconds"`
	MaximumRecordAgeInSeconds      *int `pulumi:"maximumRecordAgeInSeconds"`
	MaximumRetryAttempts           *int `pulumi:"maximumRetryAttempts"`
	ParallelizationFactor          *int `pulumi:"parallelizationFactor"`
	// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
	StartingPosition *string `pulumi:"startingPosition"`
	// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	// * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	// * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
	// * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
	// * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
	// * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
	StartingPositionTimestamp *string `pulumi:"startingPositionTimestamp"`
}

// The set of arguments for constructing a EventSourceMapping resource.
type EventSourceMappingArgs struct {
	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
	BatchSize                  pulumi.IntPtrInput
	BisectBatchOnFunctionError pulumi.BoolPtrInput
	DestinationConfig          EventSourceMappingDestinationConfigPtrInput
	// Determines if the mapping will be enabled on creation. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
	EventSourceArn pulumi.StringInput
	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName pulumi.StringInput
	// The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
	MaximumBatchingWindowInSeconds pulumi.IntPtrInput
	MaximumRecordAgeInSeconds      pulumi.IntPtrInput
	MaximumRetryAttempts           pulumi.IntPtrInput
	ParallelizationFactor          pulumi.IntPtrInput
	// The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
	StartingPosition pulumi.StringPtrInput
	// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	// * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	// * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
	// * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
	// * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
	// * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
	StartingPositionTimestamp pulumi.StringPtrInput
}

func (EventSourceMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSourceMappingArgs)(nil)).Elem()
}
