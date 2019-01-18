// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows you to set a policy of an SQS Queue
 * while referencing ARN of the queue within the policy.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_sqs_queue_q = new aws.sqs.Queue("q", {
 *     name: "examplequeue",
 * });
 * const aws_sqs_queue_policy_test = new aws.sqs.QueuePolicy("test", {
 *     policy: pulumi.all([aws_sqs_queue_q.arn, aws_sqs_queue_q.arn]).apply(([__arg0, __arg1]) => `{
 *   "Version": "2012-10-17",
 *   "Id": "sqspolicy",
 *   "Statement": [
 *     {
 *       "Sid": "First",
 *       "Effect": "Allow",
 *       "Principal": "*",
 *       "Action": "sqs:SendMessage",
 *       "Resource": "${__arg0}",
 *       "Condition": {
 *         "ArnEquals": {
 *           "aws:SourceArn": "${__arg1}"
 *         }
 *       }
 *     }
 *   ]
 * }
 * `),
 *     queueUrl: aws_sqs_queue_q.id,
 * });
 * ```
 */
export class QueuePolicy extends pulumi.CustomResource {
    /**
     * Get an existing QueuePolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueuePolicyState, opts?: pulumi.CustomResourceOptions): QueuePolicy {
        return new QueuePolicy(name, <any>state, { ...opts, id: id });
    }

    /**
     * The JSON policy for the SQS queue. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    public readonly policy: pulumi.Output<string>;
    /**
     * The URL of the SQS Queue to which to attach the policy
     */
    public readonly queueUrl: pulumi.Output<string>;

    /**
     * Create a QueuePolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueuePolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueuePolicyArgs | QueuePolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: QueuePolicyState = argsOrState as QueuePolicyState | undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["queueUrl"] = state ? state.queueUrl : undefined;
        } else {
            const args = argsOrState as QueuePolicyArgs | undefined;
            if (!args || args.policy === undefined) {
                throw new Error("Missing required property 'policy'");
            }
            if (!args || args.queueUrl === undefined) {
                throw new Error("Missing required property 'queueUrl'");
            }
            inputs["policy"] = args ? args.policy : undefined;
            inputs["queueUrl"] = args ? args.queueUrl : undefined;
        }
        super("aws:sqs/queuePolicy:QueuePolicy", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QueuePolicy resources.
 */
export interface QueuePolicyState {
    /**
     * The JSON policy for the SQS queue. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The URL of the SQS Queue to which to attach the policy
     */
    readonly queueUrl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a QueuePolicy resource.
 */
export interface QueuePolicyArgs {
    /**
     * The JSON policy for the SQS queue. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
     */
    readonly policy: pulumi.Input<string>;
    /**
     * The URL of the SQS Queue to which to attach the policy
     */
    readonly queueUrl: pulumi.Input<string>;
}
