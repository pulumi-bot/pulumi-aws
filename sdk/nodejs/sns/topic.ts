// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {ARN} from "..";

/**
 * Provides an SNS topic resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const userUpdates = new aws.sns.Topic("user_updates", {});
 * ```
 * ## Example with Delivery Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const userUpdates = new aws.sns.Topic("user_updates", {
 *     deliveryPolicy: `{
 *   "http": {
 *     "defaultHealthyRetryPolicy": {
 *       "minDelayTarget": 20,
 *       "maxDelayTarget": 20,
 *       "numRetries": 3,
 *       "numMaxDelayRetries": 0,
 *       "numNoDelayRetries": 0,
 *       "numMinDelayRetries": 0,
 *       "backoffFunction": "linear"
 *     },
 *     "disableSubscriptionOverrides": false,
 *     "defaultThrottlePolicy": {
 *       "maxReceivesPerSecond": 1
 *     }
 *   }
 * }
 * `,
 * });
 * ```
 *
 * ## Example with Server-side encryption (SSE)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const userUpdates = new aws.sns.Topic("user_updates", {
 *     kmsMasterKeyId: "alias/aws/sns",
 * });
 * ```
 *
 * ## Example with First-In-First-Out (FIFO)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const userUpdates = new aws.sns.Topic("user_updates", {
 *     contentBasedDeduplication: true,
 *     fifoTopic: true,
 * });
 * ```
 *
 * ## Message Delivery Status Arguments
 *
 * The `<endpoint>_success_feedback_role_arn` and `<endpoint>_failure_feedback_role_arn` arguments are used to give Amazon SNS write access to use CloudWatch Logs on your behalf. The `<endpoint>_success_feedback_sample_rate` argument is for specifying the sample rate percentage (0-100) of successfully delivered messages. After you configure the  `<endpoint>_failure_feedback_role_arn` argument, then all failed message deliveries generate CloudWatch Logs.
 *
 * ## Import
 *
 * SNS Topics can be imported using the `topic arn`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:sns/topic:Topic user_updates arn:aws:sns:us-west-2:0123456789012:my-topic
 * ```
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicState, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sns/topic:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * IAM role for failure feedback
     */
    public readonly applicationFailureFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    public readonly applicationSuccessFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Percentage of success to sample
     */
    public readonly applicationSuccessFeedbackSampleRate!: pulumi.Output<number | undefined>;
    /**
     * The ARN of the SNS topic, as a more obvious property (clone of id)
     */
    public /*out*/ readonly arn!: pulumi.Output<ARN>;
    /**
     * Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
     */
    public readonly contentBasedDeduplication!: pulumi.Output<boolean | undefined>;
    /**
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     */
    public readonly deliveryPolicy!: pulumi.Output<string | undefined>;
    /**
     * The display name for the topic
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
     */
    public readonly fifoTopic!: pulumi.Output<boolean | undefined>;
    /**
     * IAM role for failure feedback
     */
    public readonly httpFailureFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    public readonly httpSuccessFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Percentage of success to sample
     */
    public readonly httpSuccessFeedbackSampleRate!: pulumi.Output<number | undefined>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     */
    public readonly kmsMasterKeyId!: pulumi.Output<string | undefined>;
    /**
     * IAM role for failure feedback
     */
    public readonly lambdaFailureFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    public readonly lambdaSuccessFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Percentage of success to sample
     */
    public readonly lambdaSuccessFeedbackSampleRate!: pulumi.Output<number | undefined>;
    /**
     * The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`
     */
    public readonly namePrefix!: pulumi.Output<string>;
    /**
     * The fully-formed AWS policy as JSON.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * IAM role for failure feedback
     */
    public readonly sqsFailureFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    public readonly sqsSuccessFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * Percentage of success to sample
     */
    public readonly sqsSuccessFeedbackSampleRate!: pulumi.Output<number | undefined>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs | TopicState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicState | undefined;
            inputs["applicationFailureFeedbackRoleArn"] = state ? state.applicationFailureFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackRoleArn"] = state ? state.applicationSuccessFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackSampleRate"] = state ? state.applicationSuccessFeedbackSampleRate : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["contentBasedDeduplication"] = state ? state.contentBasedDeduplication : undefined;
            inputs["deliveryPolicy"] = state ? state.deliveryPolicy : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["fifoTopic"] = state ? state.fifoTopic : undefined;
            inputs["httpFailureFeedbackRoleArn"] = state ? state.httpFailureFeedbackRoleArn : undefined;
            inputs["httpSuccessFeedbackRoleArn"] = state ? state.httpSuccessFeedbackRoleArn : undefined;
            inputs["httpSuccessFeedbackSampleRate"] = state ? state.httpSuccessFeedbackSampleRate : undefined;
            inputs["kmsMasterKeyId"] = state ? state.kmsMasterKeyId : undefined;
            inputs["lambdaFailureFeedbackRoleArn"] = state ? state.lambdaFailureFeedbackRoleArn : undefined;
            inputs["lambdaSuccessFeedbackRoleArn"] = state ? state.lambdaSuccessFeedbackRoleArn : undefined;
            inputs["lambdaSuccessFeedbackSampleRate"] = state ? state.lambdaSuccessFeedbackSampleRate : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["sqsFailureFeedbackRoleArn"] = state ? state.sqsFailureFeedbackRoleArn : undefined;
            inputs["sqsSuccessFeedbackRoleArn"] = state ? state.sqsSuccessFeedbackRoleArn : undefined;
            inputs["sqsSuccessFeedbackSampleRate"] = state ? state.sqsSuccessFeedbackSampleRate : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as TopicArgs | undefined;
            inputs["applicationFailureFeedbackRoleArn"] = args ? args.applicationFailureFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackRoleArn"] = args ? args.applicationSuccessFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackSampleRate"] = args ? args.applicationSuccessFeedbackSampleRate : undefined;
            inputs["contentBasedDeduplication"] = args ? args.contentBasedDeduplication : undefined;
            inputs["deliveryPolicy"] = args ? args.deliveryPolicy : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["fifoTopic"] = args ? args.fifoTopic : undefined;
            inputs["httpFailureFeedbackRoleArn"] = args ? args.httpFailureFeedbackRoleArn : undefined;
            inputs["httpSuccessFeedbackRoleArn"] = args ? args.httpSuccessFeedbackRoleArn : undefined;
            inputs["httpSuccessFeedbackSampleRate"] = args ? args.httpSuccessFeedbackSampleRate : undefined;
            inputs["kmsMasterKeyId"] = args ? args.kmsMasterKeyId : undefined;
            inputs["lambdaFailureFeedbackRoleArn"] = args ? args.lambdaFailureFeedbackRoleArn : undefined;
            inputs["lambdaSuccessFeedbackRoleArn"] = args ? args.lambdaSuccessFeedbackRoleArn : undefined;
            inputs["lambdaSuccessFeedbackSampleRate"] = args ? args.lambdaSuccessFeedbackSampleRate : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["sqsFailureFeedbackRoleArn"] = args ? args.sqsFailureFeedbackRoleArn : undefined;
            inputs["sqsSuccessFeedbackRoleArn"] = args ? args.sqsSuccessFeedbackRoleArn : undefined;
            inputs["sqsSuccessFeedbackSampleRate"] = args ? args.sqsSuccessFeedbackSampleRate : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Topic resources.
 */
export interface TopicState {
    /**
     * IAM role for failure feedback
     */
    applicationFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    applicationSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    applicationSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The ARN of the SNS topic, as a more obvious property (clone of id)
     */
    arn?: pulumi.Input<ARN>;
    /**
     * Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
     */
    contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     */
    deliveryPolicy?: pulumi.Input<string>;
    /**
     * The display name for the topic
     */
    displayName?: pulumi.Input<string>;
    /**
     * Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
     */
    fifoTopic?: pulumi.Input<boolean>;
    /**
     * IAM role for failure feedback
     */
    httpFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    httpSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    httpSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     */
    kmsMasterKeyId?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    lambdaFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    lambdaSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    lambdaSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The fully-formed AWS policy as JSON.
     */
    policy?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    sqsFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    sqsSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    sqsSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * IAM role for failure feedback
     */
    applicationFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    applicationSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    applicationSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * Enables content-based deduplication for FIFO topics. For more information, see the [related documentation](https://docs.aws.amazon.com/sns/latest/dg/fifo-message-dedup.html)
     */
    contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     */
    deliveryPolicy?: pulumi.Input<string>;
    /**
     * The display name for the topic
     */
    displayName?: pulumi.Input<string>;
    /**
     * Boolean indicating whether or not to create a FIFO (first-in-first-out) topic (default is `false`).
     */
    fifoTopic?: pulumi.Input<boolean>;
    /**
     * IAM role for failure feedback
     */
    httpFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    httpSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    httpSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     */
    kmsMasterKeyId?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    lambdaFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    lambdaSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    lambdaSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The name of the topic. Topic names must be made up of only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. For a FIFO (first-in-first-out) topic, the name must end with the `.fifo` suffix. If omitted, this provider will assign a random, unique name. Conflicts with `namePrefix`
     */
    name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The fully-formed AWS policy as JSON.
     */
    policy?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    sqsFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    sqsSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    sqsSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
