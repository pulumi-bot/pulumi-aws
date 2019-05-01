// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const terraformQueue = new aws.sqs.Queue("terraform_queue", {
 *     delaySeconds: 90,
 *     maxMessageSize: 2048,
 *     messageRetentionSeconds: 86400,
 *     receiveWaitTimeSeconds: 10,
 *     redrivePolicy: aws_sqs_queue_terraform_queue_deadletter.arn.apply(arn => `{"deadLetterTargetArn":"${arn}","maxReceiveCount":4}`),
 *     tags: {
 *         Environment: "production",
 *     },
 * });
 * ```
 * 
 * ## FIFO queue
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const terraformQueue = new aws.sqs.Queue("terraform_queue", {
 *     contentBasedDeduplication: true,
 *     fifoQueue: true,
 * });
 * ```
 * 
 * ## Server-side encryption (SSE)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const terraformQueue = new aws.sqs.Queue("terraform_queue", {
 *     kmsDataKeyReusePeriodSeconds: 300,
 *     kmsMasterKeyId: "alias/aws/sqs",
 * });
 * ```
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of the SQS queue
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
     */
    public readonly contentBasedDeduplication!: pulumi.Output<boolean | undefined>;
    /**
     * The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
     */
    public readonly delaySeconds!: pulumi.Output<number | undefined>;
    /**
     * Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
     */
    public readonly fifoQueue!: pulumi.Output<boolean | undefined>;
    /**
     * The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
     */
    public readonly kmsDataKeyReusePeriodSeconds!: pulumi.Output<number>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
     */
    public readonly kmsMasterKeyId!: pulumi.Output<string | undefined>;
    /**
     * The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
     */
    public readonly maxMessageSize!: pulumi.Output<number | undefined>;
    /**
     * The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
     */
    public readonly messageRetentionSeconds!: pulumi.Output<number | undefined>;
    /**
     * This is the human-readable name of the queue. If omitted, Terraform will assign a random name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The JSON policy for the SQS queue. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
     */
    public readonly receiveWaitTimeSeconds!: pulumi.Output<number | undefined>;
    /**
     * The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
     */
    public readonly redrivePolicy!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the queue.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
     */
    public readonly visibilityTimeoutSeconds!: pulumi.Output<number | undefined>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: QueueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueueArgs | QueueState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as QueueState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["contentBasedDeduplication"] = state ? state.contentBasedDeduplication : undefined;
            inputs["delaySeconds"] = state ? state.delaySeconds : undefined;
            inputs["fifoQueue"] = state ? state.fifoQueue : undefined;
            inputs["kmsDataKeyReusePeriodSeconds"] = state ? state.kmsDataKeyReusePeriodSeconds : undefined;
            inputs["kmsMasterKeyId"] = state ? state.kmsMasterKeyId : undefined;
            inputs["maxMessageSize"] = state ? state.maxMessageSize : undefined;
            inputs["messageRetentionSeconds"] = state ? state.messageRetentionSeconds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["receiveWaitTimeSeconds"] = state ? state.receiveWaitTimeSeconds : undefined;
            inputs["redrivePolicy"] = state ? state.redrivePolicy : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["visibilityTimeoutSeconds"] = state ? state.visibilityTimeoutSeconds : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            inputs["contentBasedDeduplication"] = args ? args.contentBasedDeduplication : undefined;
            inputs["delaySeconds"] = args ? args.delaySeconds : undefined;
            inputs["fifoQueue"] = args ? args.fifoQueue : undefined;
            inputs["kmsDataKeyReusePeriodSeconds"] = args ? args.kmsDataKeyReusePeriodSeconds : undefined;
            inputs["kmsMasterKeyId"] = args ? args.kmsMasterKeyId : undefined;
            inputs["maxMessageSize"] = args ? args.maxMessageSize : undefined;
            inputs["messageRetentionSeconds"] = args ? args.messageRetentionSeconds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["receiveWaitTimeSeconds"] = args ? args.receiveWaitTimeSeconds : undefined;
            inputs["redrivePolicy"] = args ? args.redrivePolicy : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["visibilityTimeoutSeconds"] = args ? args.visibilityTimeoutSeconds : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:sqs/queue:Queue", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * The ARN of the SQS queue
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
     */
    readonly contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
     */
    readonly delaySeconds?: pulumi.Input<number>;
    /**
     * Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
     */
    readonly fifoQueue?: pulumi.Input<boolean>;
    /**
     * The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
     */
    readonly kmsDataKeyReusePeriodSeconds?: pulumi.Input<number>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
     */
    readonly kmsMasterKeyId?: pulumi.Input<string>;
    /**
     * The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
     */
    readonly maxMessageSize?: pulumi.Input<number>;
    /**
     * The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
     */
    readonly messageRetentionSeconds?: pulumi.Input<number>;
    /**
     * This is the human-readable name of the queue. If omitted, Terraform will assign a random name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The JSON policy for the SQS queue. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
     */
    readonly receiveWaitTimeSeconds?: pulumi.Input<number>;
    /**
     * The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
     */
    readonly redrivePolicy?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the queue.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
     */
    readonly visibilityTimeoutSeconds?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * Enables content-based deduplication for FIFO queues. For more information, see the [related documentation](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
     */
    readonly contentBasedDeduplication?: pulumi.Input<boolean>;
    /**
     * The time in seconds that the delivery of all messages in the queue will be delayed. An integer from 0 to 900 (15 minutes). The default for this attribute is 0 seconds.
     */
    readonly delaySeconds?: pulumi.Input<number>;
    /**
     * Boolean designating a FIFO queue. If not set, it defaults to `false` making it standard.
     */
    readonly fifoQueue?: pulumi.Input<boolean>;
    /**
     * The length of time, in seconds, for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. An integer representing seconds, between 60 seconds (1 minute) and 86,400 seconds (24 hours). The default is 300 (5 minutes).
     */
    readonly kmsDataKeyReusePeriodSeconds?: pulumi.Input<number>;
    /**
     * The ID of an AWS-managed customer master key (CMK) for Amazon SQS or a custom CMK. For more information, see [Key Terms](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
     */
    readonly kmsMasterKeyId?: pulumi.Input<string>;
    /**
     * The limit of how many bytes a message can contain before Amazon SQS rejects it. An integer from 1024 bytes (1 KiB) up to 262144 bytes (256 KiB). The default for this attribute is 262144 (256 KiB).
     */
    readonly maxMessageSize?: pulumi.Input<number>;
    /**
     * The number of seconds Amazon SQS retains a message. Integer representing seconds, from 60 (1 minute) to 1209600 (14 days). The default for this attribute is 345600 (4 days).
     */
    readonly messageRetentionSeconds?: pulumi.Input<number>;
    /**
     * This is the human-readable name of the queue. If omitted, Terraform will assign a random name.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * The JSON policy for the SQS queue. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The time for which a ReceiveMessage call will wait for a message to arrive (long polling) before returning. An integer from 0 to 20 (seconds). The default for this attribute is 0, meaning that the call will return immediately.
     */
    readonly receiveWaitTimeSeconds?: pulumi.Input<number>;
    /**
     * The JSON policy to set up the Dead Letter Queue, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/SQSDeadLetterQueue.html). **Note:** when specifying `maxReceiveCount`, you must specify it as an integer (`5`), and not a string (`"5"`).
     */
    readonly redrivePolicy?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the queue.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The visibility timeout for the queue. An integer from 0 to 43200 (12 hours). The default for this attribute is 30. For more information about visibility timeout, see [AWS docs](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/AboutVT.html).
     */
    readonly visibilityTimeoutSeconds?: pulumi.Input<number>;
}
