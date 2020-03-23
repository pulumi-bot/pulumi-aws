// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {ARN} from "../index";

/**
 * Provides an SNS topic resource
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const userUpdates = new aws.sns.Topic("userUpdates", {});
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 * ## Example with Delivery Policy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const userUpdates = new aws.sns.Topic("userUpdates", {
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
 * const userUpdates = new aws.sns.Topic("userUpdates", {
 *     kmsMasterKeyId: "alias/aws/sns",
 * });
 * ```
 * 
 * ## Message Delivery Status Arguments
 * 
 * The `<endpoint>_success_feedback_role_arn` and `<endpoint>_failure_feedback_role_arn` arguments are used to give Amazon SNS write access to use CloudWatch Logs on your behalf. The `<endpoint>_success_feedback_sample_rate` argument is for specifying the sample rate percentage (0-100) of successfully delivered messages. After you configure the  `<endpoint>_failure_feedback_role_arn` argument, then all failed message deliveries generate CloudWatch Logs.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sns_topic.html.markdown.
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
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
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     */
    public readonly deliveryPolicy!: pulumi.Output<string | undefined>;
    /**
     * The display name for the SNS topic
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
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
     * The friendly name for the SNS topic. By default generated by this provider.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The friendly name for the SNS topic. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
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
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

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
        if (opts && opts.id) {
            const state = argsOrState as TopicState | undefined;
            inputs["applicationFailureFeedbackRoleArn"] = state ? state.applicationFailureFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackRoleArn"] = state ? state.applicationSuccessFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackSampleRate"] = state ? state.applicationSuccessFeedbackSampleRate : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["deliveryPolicy"] = state ? state.deliveryPolicy : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
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
        } else {
            const args = argsOrState as TopicArgs | undefined;
            inputs["applicationFailureFeedbackRoleArn"] = args ? args.applicationFailureFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackRoleArn"] = args ? args.applicationSuccessFeedbackRoleArn : undefined;
            inputs["applicationSuccessFeedbackSampleRate"] = args ? args.applicationSuccessFeedbackSampleRate : undefined;
            inputs["deliveryPolicy"] = args ? args.deliveryPolicy : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
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
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
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
    readonly applicationFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly applicationSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly applicationSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The ARN of the SNS topic, as a more obvious property (clone of id)
     */
    readonly arn?: pulumi.Input<ARN>;
    /**
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     */
    readonly deliveryPolicy?: pulumi.Input<string>;
    /**
     * The display name for the SNS topic
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    readonly httpFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly httpSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly httpSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     */
    readonly kmsMasterKeyId?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    readonly lambdaFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly lambdaSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly lambdaSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The friendly name for the SNS topic. By default generated by this provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The friendly name for the SNS topic. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    readonly policy?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    readonly sqsFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly sqsSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly sqsSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * IAM role for failure feedback
     */
    readonly applicationFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly applicationSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly applicationSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
     */
    readonly deliveryPolicy?: pulumi.Input<string>;
    /**
     * The display name for the SNS topic
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    readonly httpFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly httpSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly httpSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
     */
    readonly kmsMasterKeyId?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    readonly lambdaFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly lambdaSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly lambdaSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * The friendly name for the SNS topic. By default generated by this provider.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The friendly name for the SNS topic. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    readonly policy?: pulumi.Input<string>;
    /**
     * IAM role for failure feedback
     */
    readonly sqsFailureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this topic
     */
    readonly sqsSuccessFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * Percentage of success to sample
     */
    readonly sqsSuccessFeedbackSampleRate?: pulumi.Input<number>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
