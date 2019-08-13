// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface TopicRuleCloudwatchAlarm {
    /**
     * The CloudWatch alarm name.
     */
    alarmName: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The reason for the alarm change.
     */
    stateReason: pulumi.Input<string>;
    /**
     * The value of the alarm state. Acceptable values are: OK, ALARM, INSUFFICIENT_DATA.
     */
    stateValue: pulumi.Input<string>;
}

export interface TopicRuleCloudwatchMetric {
    /**
     * The CloudWatch metric name.
     */
    metricName: pulumi.Input<string>;
    /**
     * The CloudWatch metric namespace name.
     */
    metricNamespace: pulumi.Input<string>;
    /**
     * An optional Unix timestamp (http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#about_timestamp).
     */
    metricTimestamp?: pulumi.Input<string>;
    /**
     * The metric unit (supported units can be found here: http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/cloudwatch_concepts.html#Unit)
     */
    metricUnit: pulumi.Input<string>;
    /**
     * The CloudWatch metric value.
     */
    metricValue: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
}

export interface TopicRuleDynamodb {
    /**
     * The hash key name.
     */
    hashKeyField: pulumi.Input<string>;
    /**
     * The hash key type. Valid values are "STRING" or "NUMBER".
     */
    hashKeyType?: pulumi.Input<string>;
    /**
     * The hash key value.
     */
    hashKeyValue: pulumi.Input<string>;
    /**
     * The action payload.
     */
    payloadField?: pulumi.Input<string>;
    /**
     * The range key name.
     */
    rangeKeyField?: pulumi.Input<string>;
    /**
     * The range key type. Valid values are "STRING" or "NUMBER".
     */
    rangeKeyType?: pulumi.Input<string>;
    /**
     * The range key value.
     */
    rangeKeyValue?: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The name of the DynamoDB table.
     */
    tableName: pulumi.Input<string>;
}

export interface TopicRuleElasticsearch {
    /**
     * The endpoint of your Elasticsearch domain.
     */
    endpoint: pulumi.Input<string>;
    /**
     * The unique identifier for the document you are storing.
     */
    id: pulumi.Input<string>;
    /**
     * The Elasticsearch index where you want to store your data.
     */
    index: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The type of document you are storing.
     */
    type: pulumi.Input<string>;
}

export interface TopicRuleFirehose {
    /**
     * The delivery stream name.
     */
    deliveryStreamName: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * A character separator that is used to separate records written to the Firehose stream. Valid values are: '\n' (newline), '\t' (tab), '\r\n' (Windows newline), ',' (comma).
     */
    separator?: pulumi.Input<string>;
}

export interface TopicRuleKinesis {
    /**
     * The partition key.
     */
    partitionKey?: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The name of the Amazon Kinesis stream.
     */
    streamName: pulumi.Input<string>;
}

export interface TopicRuleLambda {
    /**
     * The ARN of the Lambda function.
     */
    functionArn: pulumi.Input<string>;
}

export interface TopicRuleRepublish {
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The name of the MQTT topic the message should be republished to.
     */
    topic: pulumi.Input<string>;
}

export interface TopicRuleS3 {
    /**
     * The Amazon S3 bucket name.
     */
    bucketName: pulumi.Input<string>;
    /**
     * The object key.
     */
    key: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
}

export interface TopicRuleSns {
    /**
     * The message format of the message to publish. Accepted values are "JSON" and "RAW".
     */
    messageFormat?: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * The ARN of the SNS topic.
     */
    targetArn: pulumi.Input<string>;
}

export interface TopicRuleSqs {
    /**
     * The URL of the Amazon SQS queue.
     */
    queueUrl: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that grants access.
     */
    roleArn: pulumi.Input<string>;
    /**
     * Specifies whether to use Base64 encoding.
     */
    useBase64: pulumi.Input<boolean>;
}
