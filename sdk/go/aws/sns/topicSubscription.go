// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for subscribing to SNS topics. Requires that an SNS topic exist for the subscription to attach to. This resource allows you to automatically place messages sent to SNS topics in SQS queues, send them as HTTP(S) POST requests to a given endpoint, send SMS messages, or notify devices / applications. The most likely use case for provider users will probably be SQS queues.
//
// > **NOTE:** If the SNS topic and SQS queue are in different AWS regions, the `sns.TopicSubscription` must use an AWS provider that is in the same region as the SNS topic. If the `sns.TopicSubscription` uses a provider with a different region than the SNS topic, this provider will fail to create the subscription.
//
// > **NOTE:** Setup of cross-account subscriptions from SNS topics to SQS queues requires the provider to have access to BOTH accounts.
//
// > **NOTE:** If an SNS topic and SQS queue are in different AWS accounts but the same region, the `sns.TopicSubscription` must use the AWS provider for the account with the SQS queue. If `sns.TopicSubscription` uses a Provider with a different account than the SQS queue, this provider creates the subscription but does not keep state and tries to re-create the subscription at every `apply`.
//
// > **NOTE:** If an SNS topic and SQS queue are in different AWS accounts and different AWS regions, the subscription needs to be initiated from the account with the SQS queue but in the region of the SNS topic.
//
// > **NOTE:** You cannot unsubscribe to a subscription that is pending confirmation. If you use `email`, `email-json`, or `http`/`https` (without auto-confirmation enabled), until the subscription is confirmed (e.g., outside of this provider), AWS does not allow this provider to delete / unsubscribe the subscription. If you `destroy` an unconfirmed subscription, this provider will remove the subscription from its state but the subscription will still exist in AWS. However, if you delete an SNS topic, SNS [deletes all the subscriptions](https://docs.aws.amazon.com/sns/latest/dg/sns-delete-subscription-topic.html) associated with the topic. Also, you can import a subscription after confirmation and then have the capability to delete it.
//
// ## Example Usage
//
// You can directly supply a topic and ARN by hand in the `topicArn` property along with the queue ARN:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sns.NewTopicSubscription(ctx, "userUpdatesSqsTarget", &sns.TopicSubscriptionArgs{
// 			Endpoint: pulumi.String("arn:aws:sqs:us-west-2:432981146916:queue-too"),
// 			Protocol: pulumi.String("sqs"),
// 			Topic:    pulumi.Any("arn:aws:sns:us-west-2:432981146916:user-updates-topic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Alternatively you can use the ARN properties of a managed SNS topic and SQS queue:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sns"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		userUpdates, err := sns.NewTopic(ctx, "userUpdates", nil)
// 		if err != nil {
// 			return err
// 		}
// 		userUpdatesQueue, err := sqs.NewQueue(ctx, "userUpdatesQueue", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sns.NewTopicSubscription(ctx, "userUpdatesSqsTarget", &sns.TopicSubscriptionArgs{
// 			Topic:    userUpdates.Arn,
// 			Protocol: pulumi.String("sqs"),
// 			Endpoint: userUpdatesQueue.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// You can subscribe SNS topics to SQS queues in different Amazon accounts and regions:
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/config"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/providers"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sns"
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/sqs"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		sns := map[string]interface{}{
// 			"account-id":   "111111111111",
// 			"role-name":    "service/service",
// 			"name":         "example-sns-topic",
// 			"display_name": "example",
// 			"region":       "us-west-1",
// 		}
// 		if param := cfg.GetBool("sns"); param != nil {
// 			sns = param
// 		}
// 		sqs := map[string]interface{}{
// 			"account-id": "222222222222",
// 			"role-name":  "service/service",
// 			"name":       "example-sqs-queue",
// 			"region":     "us-east-1",
// 		}
// 		if param := cfg.GetBool("sqs"); param != nil {
// 			sqs = param
// 		}
// 		opt0 := "__default_policy_ID"
// 		sns_topic_policy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			PolicyId: &opt0,
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"SNS:Subscribe",
// 						"SNS:SetTopicAttributes",
// 						"SNS:RemovePermission",
// 						"SNS:Publish",
// 						"SNS:ListSubscriptionsByTopic",
// 						"SNS:GetTopicAttributes",
// 						"SNS:DeleteTopic",
// 						"SNS:AddPermission",
// 					},
// 					Conditions: []iam.GetPolicyDocumentStatementCondition{
// 						iam.GetPolicyDocumentStatementCondition{
// 							Test:     "StringEquals",
// 							Variable: "AWS:SourceOwner",
// 							Values: []string{
// 								sns.Account - id,
// 							},
// 						},
// 					},
// 					Effect: "Allow",
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "AWS",
// 							Identifiers: []string{
// 								"*",
// 							},
// 						},
// 					},
// 					Resources: []string{
// 						fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:sns:", sns.Region, ":", sns.Account-id, ":", sns.Name),
// 					},
// 					Sid: "__default_statement_ID",
// 				},
// 				iam.GetPolicyDocumentStatement{
// 					Actions: []string{
// 						"SNS:Subscribe",
// 						"SNS:Receive",
// 					},
// 					Conditions: []iam.GetPolicyDocumentStatementCondition{
// 						iam.GetPolicyDocumentStatementCondition{
// 							Test:     "StringLike",
// 							Variable: "SNS:Endpoint",
// 							Values: []string{
// 								fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:sqs:", sqs.Region, ":", sqs.Account-id, ":", sqs.Name),
// 							},
// 						},
// 					},
// 					Effect: "Allow",
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "AWS",
// 							Identifiers: []string{
// 								"*",
// 							},
// 						},
// 					},
// 					Resources: []string{
// 						fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:sns:", sns.Region, ":", sns.Account-id, ":", sns.Name),
// 					},
// 					Sid: "__console_sub_0",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt1 := fmt.Sprintf("%v%v%v%v%v%v%v", "arn:aws:sqs:", sqs.Region, ":", sqs.Account-id, ":", sqs.Name, "/SQSDefaultPolicy")
// 		sqs_queue_policy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
// 			PolicyId: &opt1,
// 			Statements: []iam.GetPolicyDocumentStatement{
// 				iam.GetPolicyDocumentStatement{
// 					Sid:    "example-sns-topic",
// 					Effect: "Allow",
// 					Principals: []iam.GetPolicyDocumentStatementPrincipal{
// 						iam.GetPolicyDocumentStatementPrincipal{
// 							Type: "AWS",
// 							Identifiers: []string{
// 								"*",
// 							},
// 						},
// 					},
// 					Actions: []string{
// 						"SQS:SendMessage",
// 					},
// 					Resources: []string{
// 						fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:sqs:", sqs.Region, ":", sqs.Account-id, ":", sqs.Name),
// 					},
// 					Conditions: []iam.GetPolicyDocumentStatementCondition{
// 						iam.GetPolicyDocumentStatementCondition{
// 							Test:     "ArnEquals",
// 							Variable: "aws:SourceArn",
// 							Values: []string{
// 								fmt.Sprintf("%v%v%v%v%v%v", "arn:aws:sns:", sns.Region, ":", sns.Account-id, ":", sns.Name),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = providers.Newaws(ctx, "awsSns", &providers.awsArgs{
// 			Region: sns.Region,
// 			AssumeRole: config.AssumeRole{
// 				RoleArn:     fmt.Sprintf("%v%v%v%v", "arn:aws:iam::", sns.Account-id, ":role/", sns.Role-name),
// 				SessionName: fmt.Sprintf("%v%v", "sns-", sns.Region),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = providers.Newaws(ctx, "awsSqs", &providers.awsArgs{
// 			Region: sqs.Region,
// 			AssumeRole: config.AssumeRole{
// 				RoleArn:     fmt.Sprintf("%v%v%v%v", "arn:aws:iam::", sqs.Account-id, ":role/", sqs.Role-name),
// 				SessionName: fmt.Sprintf("%v%v", "sqs-", sqs.Region),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = providers.Newaws(ctx, "sns2sqs", &providers.awsArgs{
// 			Region: sns.Region,
// 			AssumeRole: config.AssumeRole{
// 				RoleArn:     fmt.Sprintf("%v%v%v%v", "arn:aws:iam::", sqs.Account-id, ":role/", sqs.Role-name),
// 				SessionName: fmt.Sprintf("%v%v", "sns2sqs-", sns.Region),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sns.NewTopic(ctx, "sns_topicTopic", &sns.TopicArgs{
// 			DisplayName: pulumi.String(sns.Display_name),
// 			Policy:      pulumi.String(sns_topic_policy.Json),
// 		}, pulumi.Provider("aws.sns"))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sqs.NewQueue(ctx, "sqs_queue", &sqs.QueueArgs{
// 			Policy: pulumi.String(sqs_queue_policy.Json),
// 		}, pulumi.Provider("aws.sqs"))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sns.NewTopicSubscription(ctx, "sns_topicTopicSubscription", &sns.TopicSubscriptionArgs{
// 			Topic:    sns_topicTopic.Arn,
// 			Protocol: pulumi.String("sqs"),
// 			Endpoint: sqs_queue.Arn,
// 		}, pulumi.Provider("aws.sns2sqs"))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// SNS Topic Subscriptions can be imported using the `subscription arn`, e.g.
//
// ```sh
//  $ pulumi import aws:sns/topicSubscription:TopicSubscription user_updates_sqs_target arn:aws:sns:us-west-2:0123456789012:my-topic:8a21d249-4329-4871-acc6-7be709c6ea7f
// ```
type TopicSubscription struct {
	pulumi.CustomResourceState

	// ARN of the subscription.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
	ConfirmationTimeoutInMinutes pulumi.IntPtrOutput `pulumi:"confirmationTimeoutInMinutes"`
	// Whether the subscription confirmation request was authenticated.
	ConfirmationWasAuthenticated pulumi.BoolOutput `pulumi:"confirmationWasAuthenticated"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringPtrOutput `pulumi:"deliveryPolicy"`
	// Endpoint to send data to. The contents vary with the protocol. See details below.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
	EndpointAutoConfirms pulumi.BoolPtrOutput `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringPtrOutput `pulumi:"filterPolicy"`
	// AWS account ID of the subscription's owner.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Whether the subscription has not been confirmed.
	PendingConfirmation pulumi.BoolOutput `pulumi:"pendingConfirmation"`
	// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
	RawMessageDelivery pulumi.BoolPtrOutput `pulumi:"rawMessageDelivery"`
	// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
	RedrivePolicy pulumi.StringPtrOutput `pulumi:"redrivePolicy"`
	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
	SubscriptionRoleArn pulumi.StringPtrOutput `pulumi:"subscriptionRoleArn"`
	// ARN of the SNS topic to subscribe to.
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewTopicSubscription registers a new resource with the given unique name, arguments, and options.
func NewTopicSubscription(ctx *pulumi.Context,
	name string, args *TopicSubscriptionArgs, opts ...pulumi.ResourceOption) (*TopicSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	var resource TopicSubscription
	err := ctx.RegisterResource("aws:sns/topicSubscription:TopicSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopicSubscription gets an existing TopicSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopicSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicSubscriptionState, opts ...pulumi.ResourceOption) (*TopicSubscription, error) {
	var resource TopicSubscription
	err := ctx.ReadResource("aws:sns/topicSubscription:TopicSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TopicSubscription resources.
type topicSubscriptionState struct {
	// ARN of the subscription.
	Arn *string `pulumi:"arn"`
	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
	ConfirmationTimeoutInMinutes *int `pulumi:"confirmationTimeoutInMinutes"`
	// Whether the subscription confirmation request was authenticated.
	ConfirmationWasAuthenticated *bool `pulumi:"confirmationWasAuthenticated"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy *string `pulumi:"deliveryPolicy"`
	// Endpoint to send data to. The contents vary with the protocol. See details below.
	Endpoint *string `pulumi:"endpoint"`
	// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
	EndpointAutoConfirms *bool `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy *string `pulumi:"filterPolicy"`
	// AWS account ID of the subscription's owner.
	OwnerId *string `pulumi:"ownerId"`
	// Whether the subscription has not been confirmed.
	PendingConfirmation *bool `pulumi:"pendingConfirmation"`
	// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
	Protocol *string `pulumi:"protocol"`
	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
	RawMessageDelivery *bool `pulumi:"rawMessageDelivery"`
	// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
	RedrivePolicy *string `pulumi:"redrivePolicy"`
	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
	SubscriptionRoleArn *string `pulumi:"subscriptionRoleArn"`
	// ARN of the SNS topic to subscribe to.
	Topic interface{} `pulumi:"topic"`
}

type TopicSubscriptionState struct {
	// ARN of the subscription.
	Arn pulumi.StringPtrInput
	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
	ConfirmationTimeoutInMinutes pulumi.IntPtrInput
	// Whether the subscription confirmation request was authenticated.
	ConfirmationWasAuthenticated pulumi.BoolPtrInput
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringPtrInput
	// Endpoint to send data to. The contents vary with the protocol. See details below.
	Endpoint pulumi.StringPtrInput
	// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
	EndpointAutoConfirms pulumi.BoolPtrInput
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringPtrInput
	// AWS account ID of the subscription's owner.
	OwnerId pulumi.StringPtrInput
	// Whether the subscription has not been confirmed.
	PendingConfirmation pulumi.BoolPtrInput
	// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
	Protocol pulumi.StringPtrInput
	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
	RawMessageDelivery pulumi.BoolPtrInput
	// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
	RedrivePolicy pulumi.StringPtrInput
	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
	SubscriptionRoleArn pulumi.StringPtrInput
	// ARN of the SNS topic to subscribe to.
	Topic pulumi.Input
}

func (TopicSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSubscriptionState)(nil)).Elem()
}

type topicSubscriptionArgs struct {
	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
	ConfirmationTimeoutInMinutes *int `pulumi:"confirmationTimeoutInMinutes"`
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy *string `pulumi:"deliveryPolicy"`
	// Endpoint to send data to. The contents vary with the protocol. See details below.
	Endpoint string `pulumi:"endpoint"`
	// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
	EndpointAutoConfirms *bool `pulumi:"endpointAutoConfirms"`
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy *string `pulumi:"filterPolicy"`
	// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
	Protocol string `pulumi:"protocol"`
	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
	RawMessageDelivery *bool `pulumi:"rawMessageDelivery"`
	// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
	RedrivePolicy *string `pulumi:"redrivePolicy"`
	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
	SubscriptionRoleArn *string `pulumi:"subscriptionRoleArn"`
	// ARN of the SNS topic to subscribe to.
	Topic interface{} `pulumi:"topic"`
}

// The set of arguments for constructing a TopicSubscription resource.
type TopicSubscriptionArgs struct {
	// Integer indicating number of minutes to wait in retrying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols. Default is `1`.
	ConfirmationTimeoutInMinutes pulumi.IntPtrInput
	// JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
	DeliveryPolicy pulumi.StringPtrInput
	// Endpoint to send data to. The contents vary with the protocol. See details below.
	Endpoint pulumi.StringInput
	// Whether the endpoint is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) (e.g., PagerDuty). Default is `false`.
	EndpointAutoConfirms pulumi.BoolPtrInput
	// JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
	FilterPolicy pulumi.StringPtrInput
	// Protocol to use. Valid values are: `sqs`, `sms`, `lambda`, `firehose`, and `application`. Protocols `email`, `email-json`, `http` and `https` are also valid but partially supported. See details below.
	Protocol pulumi.StringInput
	// Whether to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property). Default is `false`.
	RawMessageDelivery pulumi.BoolPtrInput
	// JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
	RedrivePolicy pulumi.StringPtrInput
	// ARN of the IAM role to publish to Kinesis Data Firehose delivery stream. Refer to [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html).
	SubscriptionRoleArn pulumi.StringPtrInput
	// ARN of the SNS topic to subscribe to.
	Topic pulumi.Input
}

func (TopicSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicSubscriptionArgs)(nil)).Elem()
}

type TopicSubscriptionInput interface {
	pulumi.Input

	ToTopicSubscriptionOutput() TopicSubscriptionOutput
	ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput
}

func (*TopicSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil))
}

func (i *TopicSubscription) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return i.ToTopicSubscriptionOutputWithContext(context.Background())
}

func (i *TopicSubscription) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionOutput)
}

func (i *TopicSubscription) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return i.ToTopicSubscriptionPtrOutputWithContext(context.Background())
}

func (i *TopicSubscription) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionPtrOutput)
}

type TopicSubscriptionPtrInput interface {
	pulumi.Input

	ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput
	ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput
}

type topicSubscriptionPtrType TopicSubscriptionArgs

func (*topicSubscriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicSubscription)(nil))
}

func (i *topicSubscriptionPtrType) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return i.ToTopicSubscriptionPtrOutputWithContext(context.Background())
}

func (i *topicSubscriptionPtrType) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionPtrOutput)
}

// TopicSubscriptionArrayInput is an input type that accepts TopicSubscriptionArray and TopicSubscriptionArrayOutput values.
// You can construct a concrete instance of `TopicSubscriptionArrayInput` via:
//
//          TopicSubscriptionArray{ TopicSubscriptionArgs{...} }
type TopicSubscriptionArrayInput interface {
	pulumi.Input

	ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput
	ToTopicSubscriptionArrayOutputWithContext(context.Context) TopicSubscriptionArrayOutput
}

type TopicSubscriptionArray []TopicSubscriptionInput

func (TopicSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*TopicSubscription)(nil))
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return i.ToTopicSubscriptionArrayOutputWithContext(context.Background())
}

func (i TopicSubscriptionArray) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionArrayOutput)
}

// TopicSubscriptionMapInput is an input type that accepts TopicSubscriptionMap and TopicSubscriptionMapOutput values.
// You can construct a concrete instance of `TopicSubscriptionMapInput` via:
//
//          TopicSubscriptionMap{ "key": TopicSubscriptionArgs{...} }
type TopicSubscriptionMapInput interface {
	pulumi.Input

	ToTopicSubscriptionMapOutput() TopicSubscriptionMapOutput
	ToTopicSubscriptionMapOutputWithContext(context.Context) TopicSubscriptionMapOutput
}

type TopicSubscriptionMap map[string]TopicSubscriptionInput

func (TopicSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*TopicSubscription)(nil))
}

func (i TopicSubscriptionMap) ToTopicSubscriptionMapOutput() TopicSubscriptionMapOutput {
	return i.ToTopicSubscriptionMapOutputWithContext(context.Background())
}

func (i TopicSubscriptionMap) ToTopicSubscriptionMapOutputWithContext(ctx context.Context) TopicSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicSubscriptionMapOutput)
}

type TopicSubscriptionOutput struct {
	*pulumi.OutputState
}

func (TopicSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TopicSubscription)(nil))
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutput() TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionOutputWithContext(ctx context.Context) TopicSubscriptionOutput {
	return o
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return o.ToTopicSubscriptionPtrOutputWithContext(context.Background())
}

func (o TopicSubscriptionOutput) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return o.ApplyT(func(v TopicSubscription) *TopicSubscription {
		return &v
	}).(TopicSubscriptionPtrOutput)
}

type TopicSubscriptionPtrOutput struct {
	*pulumi.OutputState
}

func (TopicSubscriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicSubscription)(nil))
}

func (o TopicSubscriptionPtrOutput) ToTopicSubscriptionPtrOutput() TopicSubscriptionPtrOutput {
	return o
}

func (o TopicSubscriptionPtrOutput) ToTopicSubscriptionPtrOutputWithContext(ctx context.Context) TopicSubscriptionPtrOutput {
	return o
}

type TopicSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TopicSubscription)(nil))
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutput() TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) ToTopicSubscriptionArrayOutputWithContext(ctx context.Context) TopicSubscriptionArrayOutput {
	return o
}

func (o TopicSubscriptionArrayOutput) Index(i pulumi.IntInput) TopicSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TopicSubscription {
		return vs[0].([]TopicSubscription)[vs[1].(int)]
	}).(TopicSubscriptionOutput)
}

type TopicSubscriptionMapOutput struct{ *pulumi.OutputState }

func (TopicSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]TopicSubscription)(nil))
}

func (o TopicSubscriptionMapOutput) ToTopicSubscriptionMapOutput() TopicSubscriptionMapOutput {
	return o
}

func (o TopicSubscriptionMapOutput) ToTopicSubscriptionMapOutputWithContext(ctx context.Context) TopicSubscriptionMapOutput {
	return o
}

func (o TopicSubscriptionMapOutput) MapIndex(k pulumi.StringInput) TopicSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) TopicSubscription {
		return vs[0].(map[string]TopicSubscription)[vs[1].(string)]
	}).(TopicSubscriptionOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicSubscriptionOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionPtrOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(TopicSubscriptionMapOutput{})
}
