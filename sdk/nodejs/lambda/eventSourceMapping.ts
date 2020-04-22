// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Lambda event source mapping. This allows Lambda functions to get events from Kinesis, DynamoDB and SQS.
 * 
 * For information about Lambda and how to use it, see [What is AWS Lambda?](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html).
 * For information about event source mappings, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the API docs.
 * 
 * ## Example Usage
 * 
 * ### DynamoDB
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_dynamodb_table_example.streamArn,
 *     functionName: aws_lambda_function_example.arn,
 *     startingPosition: "LATEST",
 * });
 * ```
 * 
 * ### Kinesis
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_kinesis_stream_example.arn,
 *     functionName: aws_lambda_function_example.arn,
 *     startingPosition: "LATEST",
 * });
 * ```
 * 
 * ### SQS
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_sqs_queue_sqs_queue_test.arn,
 *     functionName: aws_lambda_function_example.arn,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_event_source_mapping.html.markdown.
 */
export class EventSourceMapping extends pulumi.CustomResource {
    /**
     * Get an existing EventSourceMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventSourceMappingState, opts?: pulumi.CustomResourceOptions): EventSourceMapping {
        return new EventSourceMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/eventSourceMapping:EventSourceMapping';

    /**
     * Returns true if the given object is an instance of EventSourceMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSourceMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSourceMapping.__pulumiType;
    }

    /**
     * The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
     */
    public readonly batchSize!: pulumi.Output<number | undefined>;
    public readonly bisectBatchOnFunctionError!: pulumi.Output<boolean | undefined>;
    public readonly destinationConfig!: pulumi.Output<outputs.lambda.EventSourceMappingDestinationConfig | undefined>;
    /**
     * Determines if the mapping will be enabled on creation. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
     */
    public readonly eventSourceArn!: pulumi.Output<string>;
    /**
     * The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `functionName` above.)
     */
    public /*out*/ readonly functionArn!: pulumi.Output<string>;
    /**
     * The name or the ARN of the Lambda function that will be subscribing to events.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The result of the last AWS Lambda invocation of your Lambda function.
     */
    public /*out*/ readonly lastProcessingResult!: pulumi.Output<string>;
    /**
     * The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
     */
    public readonly maximumBatchingWindowInSeconds!: pulumi.Output<number | undefined>;
    public readonly maximumRecordAgeInSeconds!: pulumi.Output<number>;
    public readonly maximumRetryAttempts!: pulumi.Output<number>;
    public readonly parallelizationFactor!: pulumi.Output<number>;
    /**
     * The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
     */
    public readonly startingPosition!: pulumi.Output<string | undefined>;
    /**
     * A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
     * * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
     * * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
     * * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
     * * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
     * * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
     */
    public readonly startingPositionTimestamp!: pulumi.Output<string | undefined>;
    /**
     * The state of the event source mapping.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The reason the event source mapping is in its current state.
     */
    public /*out*/ readonly stateTransitionReason!: pulumi.Output<string>;
    /**
     * The UUID of the created event source mapping.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a EventSourceMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSourceMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventSourceMappingArgs | EventSourceMappingState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventSourceMappingState | undefined;
            inputs["batchSize"] = state ? state.batchSize : undefined;
            inputs["bisectBatchOnFunctionError"] = state ? state.bisectBatchOnFunctionError : undefined;
            inputs["destinationConfig"] = state ? state.destinationConfig : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["eventSourceArn"] = state ? state.eventSourceArn : undefined;
            inputs["functionArn"] = state ? state.functionArn : undefined;
            inputs["functionName"] = state ? state.functionName : undefined;
            inputs["lastModified"] = state ? state.lastModified : undefined;
            inputs["lastProcessingResult"] = state ? state.lastProcessingResult : undefined;
            inputs["maximumBatchingWindowInSeconds"] = state ? state.maximumBatchingWindowInSeconds : undefined;
            inputs["maximumRecordAgeInSeconds"] = state ? state.maximumRecordAgeInSeconds : undefined;
            inputs["maximumRetryAttempts"] = state ? state.maximumRetryAttempts : undefined;
            inputs["parallelizationFactor"] = state ? state.parallelizationFactor : undefined;
            inputs["startingPosition"] = state ? state.startingPosition : undefined;
            inputs["startingPositionTimestamp"] = state ? state.startingPositionTimestamp : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["stateTransitionReason"] = state ? state.stateTransitionReason : undefined;
            inputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as EventSourceMappingArgs | undefined;
            if (!args || args.eventSourceArn === undefined) {
                throw new Error("Missing required property 'eventSourceArn'");
            }
            if (!args || args.functionName === undefined) {
                throw new Error("Missing required property 'functionName'");
            }
            inputs["batchSize"] = args ? args.batchSize : undefined;
            inputs["bisectBatchOnFunctionError"] = args ? args.bisectBatchOnFunctionError : undefined;
            inputs["destinationConfig"] = args ? args.destinationConfig : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["eventSourceArn"] = args ? args.eventSourceArn : undefined;
            inputs["functionName"] = args ? args.functionName : undefined;
            inputs["maximumBatchingWindowInSeconds"] = args ? args.maximumBatchingWindowInSeconds : undefined;
            inputs["maximumRecordAgeInSeconds"] = args ? args.maximumRecordAgeInSeconds : undefined;
            inputs["maximumRetryAttempts"] = args ? args.maximumRetryAttempts : undefined;
            inputs["parallelizationFactor"] = args ? args.parallelizationFactor : undefined;
            inputs["startingPosition"] = args ? args.startingPosition : undefined;
            inputs["startingPositionTimestamp"] = args ? args.startingPositionTimestamp : undefined;
            inputs["functionArn"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["lastProcessingResult"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["stateTransitionReason"] = undefined /*out*/;
            inputs["uuid"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventSourceMapping.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventSourceMapping resources.
 */
export interface EventSourceMappingState {
    /**
     * The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
     */
    readonly batchSize?: pulumi.Input<number>;
    readonly bisectBatchOnFunctionError?: pulumi.Input<boolean>;
    readonly destinationConfig?: pulumi.Input<inputs.lambda.EventSourceMappingDestinationConfig>;
    /**
     * Determines if the mapping will be enabled on creation. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
     */
    readonly eventSourceArn?: pulumi.Input<string>;
    /**
     * The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `functionName` above.)
     */
    readonly functionArn?: pulumi.Input<string>;
    /**
     * The name or the ARN of the Lambda function that will be subscribing to events.
     */
    readonly functionName?: pulumi.Input<string>;
    /**
     * The date this resource was last modified.
     */
    readonly lastModified?: pulumi.Input<string>;
    /**
     * The result of the last AWS Lambda invocation of your Lambda function.
     */
    readonly lastProcessingResult?: pulumi.Input<string>;
    /**
     * The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
     */
    readonly maximumBatchingWindowInSeconds?: pulumi.Input<number>;
    readonly maximumRecordAgeInSeconds?: pulumi.Input<number>;
    readonly maximumRetryAttempts?: pulumi.Input<number>;
    readonly parallelizationFactor?: pulumi.Input<number>;
    /**
     * The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
     */
    readonly startingPosition?: pulumi.Input<string>;
    /**
     * A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
     * * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
     * * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
     * * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
     * * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
     * * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
     */
    readonly startingPositionTimestamp?: pulumi.Input<string>;
    /**
     * The state of the event source mapping.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The reason the event source mapping is in its current state.
     */
    readonly stateTransitionReason?: pulumi.Input<string>;
    /**
     * The UUID of the created event source mapping.
     */
    readonly uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventSourceMapping resource.
 */
export interface EventSourceMappingArgs {
    /**
     * The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB and Kinesis, `10` for SQS.
     */
    readonly batchSize?: pulumi.Input<number>;
    readonly bisectBatchOnFunctionError?: pulumi.Input<boolean>;
    readonly destinationConfig?: pulumi.Input<inputs.lambda.EventSourceMappingDestinationConfig>;
    /**
     * Determines if the mapping will be enabled on creation. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The event source ARN - can be a Kinesis stream, DynamoDB stream, or SQS queue.
     */
    readonly eventSourceArn: pulumi.Input<string>;
    /**
     * The name or the ARN of the Lambda function that will be subscribing to events.
     */
    readonly functionName: pulumi.Input<string>;
    /**
     * The maximum amount of time to gather records before invoking the function, in seconds.  Records will continue to buffer until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. Defaults to as soon as records are available in the stream. If the batch it reads from the stream only has one record in it, Lambda only sends one record to the function.
     */
    readonly maximumBatchingWindowInSeconds?: pulumi.Input<number>;
    readonly maximumRecordAgeInSeconds?: pulumi.Input<number>;
    readonly maximumRetryAttempts?: pulumi.Input<number>;
    readonly parallelizationFactor?: pulumi.Input<number>;
    /**
     * The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis or DynamoDB. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
     */
    readonly startingPosition?: pulumi.Input<string>;
    /**
     * A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
     * * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
     * * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum of 0, maximum and default of 10000.
     * * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Minimum of 60, maximum and default of 604800.
     * * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
     * * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
     */
    readonly startingPositionTimestamp?: pulumi.Input<string>;
}
