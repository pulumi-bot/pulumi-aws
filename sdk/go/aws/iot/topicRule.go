// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type TopicRule struct {
	pulumi.CustomResourceState

	// The ARN of the topic rule
	Arn pulumi.StringOutput `pulumi:"arn"`
	CloudwatchAlarm TopicRuleCloudwatchAlarmPtrOutput `pulumi:"cloudwatchAlarm"`
	CloudwatchMetric TopicRuleCloudwatchMetricPtrOutput `pulumi:"cloudwatchMetric"`
	// The description of the rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Dynamodb TopicRuleDynamodbPtrOutput `pulumi:"dynamodb"`
	Elasticsearch TopicRuleElasticsearchPtrOutput `pulumi:"elasticsearch"`
	// Specifies whether the rule is enabled.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	Firehose TopicRuleFirehosePtrOutput `pulumi:"firehose"`
	Kinesis TopicRuleKinesisPtrOutput `pulumi:"kinesis"`
	Lambda TopicRuleLambdaPtrOutput `pulumi:"lambda"`
	// The name of the rule.
	Name pulumi.StringOutput `pulumi:"name"`
	Republish TopicRuleRepublishPtrOutput `pulumi:"republish"`
	S3 TopicRuleS3PtrOutput `pulumi:"s3"`
	Sns TopicRuleSnsPtrOutput `pulumi:"sns"`
	// The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
	Sql pulumi.StringOutput `pulumi:"sql"`
	// The version of the SQL rules engine to use when evaluating the rule.
	SqlVersion pulumi.StringOutput `pulumi:"sqlVersion"`
	Sqs TopicRuleSqsPtrOutput `pulumi:"sqs"`
}

// NewTopicRule registers a new resource with the given unique name, arguments, and options.
func NewTopicRule(ctx *pulumi.Context,
	name string, args *TopicRuleArgs, opts ...pulumi.ResourceOption) (*TopicRule, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.Sql == nil {
		return nil, errors.New("missing required argument 'Sql'")
	}
	if args == nil || args.SqlVersion == nil {
		return nil, errors.New("missing required argument 'SqlVersion'")
	}
	if args == nil {
		args = &TopicRuleArgs{}
	}
	var resource TopicRule
	err := ctx.RegisterResource("aws:iot/topicRule:TopicRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicRule gets an existing TopicRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicRuleState, opts ...pulumi.ResourceOption) (*TopicRule, error) {
	var resource TopicRule
	err := ctx.ReadResource("aws:iot/topicRule:TopicRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicRule resources.
type topicRuleState struct {
	// The ARN of the topic rule
	Arn *string `pulumi:"arn"`
	CloudwatchAlarm *TopicRuleCloudwatchAlarm `pulumi:"cloudwatchAlarm"`
	CloudwatchMetric *TopicRuleCloudwatchMetric `pulumi:"cloudwatchMetric"`
	// The description of the rule.
	Description *string `pulumi:"description"`
	Dynamodb *TopicRuleDynamodb `pulumi:"dynamodb"`
	Elasticsearch *TopicRuleElasticsearch `pulumi:"elasticsearch"`
	// Specifies whether the rule is enabled.
	Enabled *bool `pulumi:"enabled"`
	Firehose *TopicRuleFirehose `pulumi:"firehose"`
	Kinesis *TopicRuleKinesis `pulumi:"kinesis"`
	Lambda *TopicRuleLambda `pulumi:"lambda"`
	// The name of the rule.
	Name *string `pulumi:"name"`
	Republish *TopicRuleRepublish `pulumi:"republish"`
	S3 *TopicRuleS3 `pulumi:"s3"`
	Sns *TopicRuleSns `pulumi:"sns"`
	// The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
	Sql *string `pulumi:"sql"`
	// The version of the SQL rules engine to use when evaluating the rule.
	SqlVersion *string `pulumi:"sqlVersion"`
	Sqs *TopicRuleSqs `pulumi:"sqs"`
}

type TopicRuleState struct {
	// The ARN of the topic rule
	Arn pulumi.StringPtrInput
	CloudwatchAlarm TopicRuleCloudwatchAlarmPtrInput
	CloudwatchMetric TopicRuleCloudwatchMetricPtrInput
	// The description of the rule.
	Description pulumi.StringPtrInput
	Dynamodb TopicRuleDynamodbPtrInput
	Elasticsearch TopicRuleElasticsearchPtrInput
	// Specifies whether the rule is enabled.
	Enabled pulumi.BoolPtrInput
	Firehose TopicRuleFirehosePtrInput
	Kinesis TopicRuleKinesisPtrInput
	Lambda TopicRuleLambdaPtrInput
	// The name of the rule.
	Name pulumi.StringPtrInput
	Republish TopicRuleRepublishPtrInput
	S3 TopicRuleS3PtrInput
	Sns TopicRuleSnsPtrInput
	// The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
	Sql pulumi.StringPtrInput
	// The version of the SQL rules engine to use when evaluating the rule.
	SqlVersion pulumi.StringPtrInput
	Sqs TopicRuleSqsPtrInput
}

func (TopicRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicRuleState)(nil)).Elem()
}

type topicRuleArgs struct {
	CloudwatchAlarm *TopicRuleCloudwatchAlarm `pulumi:"cloudwatchAlarm"`
	CloudwatchMetric *TopicRuleCloudwatchMetric `pulumi:"cloudwatchMetric"`
	// The description of the rule.
	Description *string `pulumi:"description"`
	Dynamodb *TopicRuleDynamodb `pulumi:"dynamodb"`
	Elasticsearch *TopicRuleElasticsearch `pulumi:"elasticsearch"`
	// Specifies whether the rule is enabled.
	Enabled bool `pulumi:"enabled"`
	Firehose *TopicRuleFirehose `pulumi:"firehose"`
	Kinesis *TopicRuleKinesis `pulumi:"kinesis"`
	Lambda *TopicRuleLambda `pulumi:"lambda"`
	// The name of the rule.
	Name *string `pulumi:"name"`
	Republish *TopicRuleRepublish `pulumi:"republish"`
	S3 *TopicRuleS3 `pulumi:"s3"`
	Sns *TopicRuleSns `pulumi:"sns"`
	// The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
	Sql string `pulumi:"sql"`
	// The version of the SQL rules engine to use when evaluating the rule.
	SqlVersion string `pulumi:"sqlVersion"`
	Sqs *TopicRuleSqs `pulumi:"sqs"`
}

// The set of arguments for constructing a TopicRule resource.
type TopicRuleArgs struct {
	CloudwatchAlarm TopicRuleCloudwatchAlarmPtrInput
	CloudwatchMetric TopicRuleCloudwatchMetricPtrInput
	// The description of the rule.
	Description pulumi.StringPtrInput
	Dynamodb TopicRuleDynamodbPtrInput
	Elasticsearch TopicRuleElasticsearchPtrInput
	// Specifies whether the rule is enabled.
	Enabled pulumi.BoolInput
	Firehose TopicRuleFirehosePtrInput
	Kinesis TopicRuleKinesisPtrInput
	Lambda TopicRuleLambdaPtrInput
	// The name of the rule.
	Name pulumi.StringPtrInput
	Republish TopicRuleRepublishPtrInput
	S3 TopicRuleS3PtrInput
	Sns TopicRuleSnsPtrInput
	// The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
	Sql pulumi.StringInput
	// The version of the SQL rules engine to use when evaluating the rule.
	SqlVersion pulumi.StringInput
	Sqs TopicRuleSqsPtrInput
}

func (TopicRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicRuleArgs)(nil)).Elem()
}

