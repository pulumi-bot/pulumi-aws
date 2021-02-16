// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {Topic} from "./index";

/**
 * Provides a resource for subscribing to SNS topics. Requires that an SNS topic exist for the subscription to attach to.
 * This resource allows you to automatically place messages sent to SNS topics in SQS queues, send them as HTTP(S) POST requests
 * to a given endpoint, send SMS messages, or notify devices / applications. The most likely use case will
 * probably be SQS queues.
 *
 * > **NOTE:** If the SNS topic and SQS queue are in different AWS regions, it is important for the "aws.sns.TopicSubscription" to use an AWS provider that is in the same region of the SNS topic. If the "aws.sns.TopicSubscription" is using a provider with a different region than the SNS topic, the subscription will fail to create.
 *
 * > **NOTE:** Setup of cross-account subscriptions from SNS topics to SQS queues requires the provider to have access to BOTH accounts.
 *
 * > **NOTE:** If SNS topic and SQS queue are in different AWS accounts but the same region it is important for the "aws.sns.TopicSubscription" to use the AWS provider of the account with the SQS queue. If "aws.sns.TopicSubscription" is using a Provider with a different account than the SQS queue, the provider creates the subscriptions but does not keep state and tries to re-create the subscription at every apply.
 *
 * > **NOTE:** If SNS topic and SQS queue are in different AWS accounts and different AWS regions it is important to recognize that the subscription needs to be initiated from the account with the SQS queue but in the region of the SNS topic.
 *
 * ## Example Usage
 *
 * You can directly supply a topic and ARN by hand in the `topicArn` property along with the queue ARN:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const userUpdatesSqsTarget = new aws.sns.TopicSubscription("user_updates_sqs_target", {
 *     endpoint: "arn:aws:sqs:us-west-2:432981146916:queue-too",
 *     protocol: "sqs",
 *     topic: "arn:aws:sns:us-west-2:432981146916:user-updates-topic",
 * });
 * ```
 *
 * Alternatively you can use the ARN properties of a managed SNS topic and SQS queue:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const userUpdates = new aws.sns.Topic("userUpdates", {});
 * const userUpdatesQueue = new aws.sqs.Queue("userUpdatesQueue", {});
 * const userUpdatesSqsTarget = new aws.sns.TopicSubscription("userUpdatesSqsTarget", {
 *     topic: userUpdates.arn,
 *     protocol: "sqs",
 *     endpoint: userUpdatesQueue.arn,
 * });
 * ```
 *
 * You can subscribe SNS topics to SQS queues in different Amazon accounts and regions:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const sns = config.getObject("sns") || {
 *     "account-id": "111111111111",
 *     "role-name": "service/service",
 *     name: "example-sns-topic",
 *     display_name: "example",
 *     region: "us-west-1",
 * };
 * const sqs = config.getObject("sqs") || {
 *     "account-id": "222222222222",
 *     "role-name": "service/service",
 *     name: "example-sqs-queue",
 *     region: "us-east-1",
 * };
 * const sns-topic-policy = aws.iam.getPolicyDocument({
 *     policyId: "__default_policy_ID",
 *     statements: [
 *         {
 *             actions: [
 *                 "SNS:Subscribe",
 *                 "SNS:SetTopicAttributes",
 *                 "SNS:RemovePermission",
 *                 "SNS:Publish",
 *                 "SNS:ListSubscriptionsByTopic",
 *                 "SNS:GetTopicAttributes",
 *                 "SNS:DeleteTopic",
 *                 "SNS:AddPermission",
 *             ],
 *             conditions: [{
 *                 test: "StringEquals",
 *                 variable: "AWS:SourceOwner",
 *                 values: [sns["account-id"]],
 *             }],
 *             effect: "Allow",
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: ["*"],
 *             }],
 *             resources: [`arn:aws:sns:${sns.region}:${sns["account-id"]}:${sns.name}`],
 *             sid: "__default_statement_ID",
 *         },
 *         {
 *             actions: [
 *                 "SNS:Subscribe",
 *                 "SNS:Receive",
 *             ],
 *             conditions: [{
 *                 test: "StringLike",
 *                 variable: "SNS:Endpoint",
 *                 values: [`arn:aws:sqs:${sqs.region}:${sqs["account-id"]}:${sqs.name}`],
 *             }],
 *             effect: "Allow",
 *             principals: [{
 *                 type: "AWS",
 *                 identifiers: ["*"],
 *             }],
 *             resources: [`arn:aws:sns:${sns.region}:${sns["account-id"]}:${sns.name}`],
 *             sid: "__console_sub_0",
 *         },
 *     ],
 * });
 * const sqs-queue-policy = aws.iam.getPolicyDocument({
 *     policyId: `arn:aws:sqs:${sqs.region}:${sqs["account-id"]}:${sqs.name}/SQSDefaultPolicy`,
 *     statements: [{
 *         sid: "example-sns-topic",
 *         effect: "Allow",
 *         principals: [{
 *             type: "AWS",
 *             identifiers: ["*"],
 *         }],
 *         actions: ["SQS:SendMessage"],
 *         resources: [`arn:aws:sqs:${sqs.region}:${sqs["account-id"]}:${sqs.name}`],
 *         conditions: [{
 *             test: "ArnEquals",
 *             variable: "aws:SourceArn",
 *             values: [`arn:aws:sns:${sns.region}:${sns["account-id"]}:${sns.name}`],
 *         }],
 *     }],
 * });
 * // provider to manage SNS topics
 * const awsSns = new aws.Provider("awsSns", {
 *     region: sns.region,
 *     assumeRole: {
 *         roleArn: `arn:aws:iam::${sns["account-id"]}:role/${sns["role-name"]}`,
 *         sessionName: `sns-${sns.region}`,
 *     },
 * });
 * // provider to manage SQS queues
 * const awsSqs = new aws.Provider("awsSqs", {
 *     region: sqs.region,
 *     assumeRole: {
 *         roleArn: `arn:aws:iam::${sqs["account-id"]}:role/${sqs["role-name"]}`,
 *         sessionName: `sqs-${sqs.region}`,
 *     },
 * });
 * // provider to subscribe SQS to SNS (using the SQS account but the SNS region)
 * const sns2sqs = new aws.Provider("sns2sqs", {
 *     region: sns.region,
 *     assumeRole: {
 *         roleArn: `arn:aws:iam::${sqs["account-id"]}:role/${sqs["role-name"]}`,
 *         sessionName: `sns2sqs-${sns.region}`,
 *     },
 * });
 * const sns_topicTopic = new aws.sns.Topic("sns-topicTopic", {
 *     displayName: sns.display_name,
 *     policy: sns_topic_policy.then(sns_topic_policy => sns_topic_policy.json),
 * }, {
 *     provider: "aws.sns",
 * });
 * const sqs_queue = new aws.sqs.Queue("sqs-queue", {policy: sqs_queue_policy.then(sqs_queue_policy => sqs_queue_policy.json)}, {
 *     provider: "aws.sqs",
 * });
 * const sns_topicTopicSubscription = new aws.sns.TopicSubscription("sns-topicTopicSubscription", {
 *     topic: sns_topicTopic.arn,
 *     protocol: "sqs",
 *     endpoint: sqs_queue.arn,
 * }, {
 *     provider: "aws.sns2sqs",
 * });
 * ```
 *
 * ## Import
 *
 * SNS Topic Subscriptions can be imported using the `subscription arn`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:sns/topicSubscription:TopicSubscription user_updates_sqs_target arn:aws:sns:us-west-2:0123456789012:my-topic:8a21d249-4329-4871-acc6-7be709c6ea7f
 * ```
 */
export class TopicSubscription extends pulumi.CustomResource {
    /**
     * Get an existing TopicSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicSubscriptionState, opts?: pulumi.CustomResourceOptions): TopicSubscription {
        return new TopicSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sns/topicSubscription:TopicSubscription';

    /**
     * Returns true if the given object is an instance of TopicSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicSubscription.__pulumiType;
    }

    /**
     * The ARN of the subscription stored as a more user-friendly property
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
     */
    public readonly confirmationTimeoutInMinutes!: pulumi.Output<number | undefined>;
    /**
     * JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
     */
    public readonly deliveryPolicy!: pulumi.Output<string | undefined>;
    /**
     * The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
     */
    public readonly endpoint!: pulumi.Output<string>;
    /**
     * Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
     */
    public readonly endpointAutoConfirms!: pulumi.Output<boolean | undefined>;
    /**
     * JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
     */
    public readonly filterPolicy!: pulumi.Output<string | undefined>;
    /**
     * The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
     */
    public readonly rawMessageDelivery!: pulumi.Output<boolean | undefined>;
    /**
     * JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
     */
    public readonly redrivePolicy!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the SNS topic to subscribe to
     */
    public readonly topic!: pulumi.Output<string>;

    /**
     * Create a TopicSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicSubscriptionArgs | TopicSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicSubscriptionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["confirmationTimeoutInMinutes"] = state ? state.confirmationTimeoutInMinutes : undefined;
            inputs["deliveryPolicy"] = state ? state.deliveryPolicy : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["endpointAutoConfirms"] = state ? state.endpointAutoConfirms : undefined;
            inputs["filterPolicy"] = state ? state.filterPolicy : undefined;
            inputs["protocol"] = state ? state.protocol : undefined;
            inputs["rawMessageDelivery"] = state ? state.rawMessageDelivery : undefined;
            inputs["redrivePolicy"] = state ? state.redrivePolicy : undefined;
            inputs["topic"] = state ? state.topic : undefined;
        } else {
            const args = argsOrState as TopicSubscriptionArgs | undefined;
            if ((!args || args.endpoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoint'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.topic === undefined) && !opts.urn) {
                throw new Error("Missing required property 'topic'");
            }
            inputs["confirmationTimeoutInMinutes"] = args ? args.confirmationTimeoutInMinutes : undefined;
            inputs["deliveryPolicy"] = args ? args.deliveryPolicy : undefined;
            inputs["endpoint"] = args ? args.endpoint : undefined;
            inputs["endpointAutoConfirms"] = args ? args.endpointAutoConfirms : undefined;
            inputs["filterPolicy"] = args ? args.filterPolicy : undefined;
            inputs["protocol"] = args ? args.protocol : undefined;
            inputs["rawMessageDelivery"] = args ? args.rawMessageDelivery : undefined;
            inputs["redrivePolicy"] = args ? args.redrivePolicy : undefined;
            inputs["topic"] = args ? args.topic : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(TopicSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TopicSubscription resources.
 */
export interface TopicSubscriptionState {
    /**
     * The ARN of the subscription stored as a more user-friendly property
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
     */
    readonly confirmationTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
     */
    readonly deliveryPolicy?: pulumi.Input<string>;
    /**
     * The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
     */
    readonly endpointAutoConfirms?: pulumi.Input<boolean>;
    /**
     * JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
     */
    readonly filterPolicy?: pulumi.Input<string>;
    /**
     * The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
     */
    readonly protocol?: pulumi.Input<string>;
    /**
     * Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
     */
    readonly rawMessageDelivery?: pulumi.Input<boolean>;
    /**
     * JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
     */
    readonly redrivePolicy?: pulumi.Input<string>;
    /**
     * The ARN of the SNS topic to subscribe to
     */
    readonly topic?: pulumi.Input<string | Topic>;
}

/**
 * The set of arguments for constructing a TopicSubscription resource.
 */
export interface TopicSubscriptionArgs {
    /**
     * Integer indicating number of minutes to wait in retying mode for fetching subscription arn before marking it as failure. Only applicable for http and https protocols (default is 1 minute).
     */
    readonly confirmationTimeoutInMinutes?: pulumi.Input<number>;
    /**
     * JSON String with the delivery policy (retries, backoff, etc.) that will be used in the subscription - this only applies to HTTP/S subscriptions. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html) for more details.
     */
    readonly deliveryPolicy?: pulumi.Input<string>;
    /**
     * The endpoint to send data to, the contents will vary with the protocol. (see below for more information)
     */
    readonly endpoint: pulumi.Input<string>;
    /**
     * Boolean indicating whether the end point is capable of [auto confirming subscription](http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.html#SendMessageToHttp.prepare) e.g., PagerDuty (default is false)
     */
    readonly endpointAutoConfirms?: pulumi.Input<boolean>;
    /**
     * JSON String with the filter policy that will be used in the subscription to filter messages seen by the target resource. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/message-filtering.html) for more details.
     */
    readonly filterPolicy?: pulumi.Input<string>;
    /**
     * The protocol to use. The possible values for this are: `sqs`, `sms`, `lambda`, `application`. (`http` or `https` are partially supported, see below) (`email` is an option but is unsupported, see below).
     */
    readonly protocol: pulumi.Input<string>;
    /**
     * Boolean indicating whether or not to enable raw message delivery (the original message is directly passed, not wrapped in JSON with the original message in the message property) (default is false).
     */
    readonly rawMessageDelivery?: pulumi.Input<boolean>;
    /**
     * JSON String with the redrive policy that will be used in the subscription. Refer to the [SNS docs](https://docs.aws.amazon.com/sns/latest/dg/sns-dead-letter-queues.html#how-messages-moved-into-dead-letter-queue) for more details.
     */
    readonly redrivePolicy?: pulumi.Input<string>;
    /**
     * The ARN of the SNS topic to subscribe to
     */
    readonly topic: pulumi.Input<string | Topic>;
}
