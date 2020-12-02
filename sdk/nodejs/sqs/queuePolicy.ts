// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {PolicyDocument} from "../iam";

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
 * const queue = new aws.sqs.Queue("queue", {});
 * const test = new aws.sqs.QueuePolicy("test", {
 *     queueUrl: queue.id,
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Id": "sqspolicy",
 *   "Statement": [
 *     {
 *       "Sid": "First",
 *       "Effect": "Allow",
 *       "Principal": "*",
 *       "Action": "sqs:SendMessage",
 *       "Resource": "${queue.arn}",
 *       "Condition": {
 *         "ArnEquals": {
 *           "aws:SourceArn": "${aws_sns_topic.example.arn}"
 *         }
 *       }
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * SQS Queue Policies can be imported using the queue URL, e.g.
 *
 * ```sh
 *  $ pulumi import aws:sqs/queuePolicy:QueuePolicy test https://queue.amazonaws.com/0123456789012/myqueue
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
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueuePolicyState, opts?: pulumi.CustomResourceOptions): QueuePolicy {
        return new QueuePolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sqs/queuePolicy:QueuePolicy';

    /**
     * Returns true if the given object is an instance of QueuePolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueuePolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueuePolicy.__pulumiType;
    }

    /**
     * The JSON policy for the SQS queue.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * The URL of the SQS Queue to which to attach the policy
     */
    public readonly queueUrl!: pulumi.Output<string>;

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
            const state = argsOrState as QueuePolicyState | undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["queueUrl"] = state ? state.queueUrl : undefined;
        } else {
            const args = argsOrState as QueuePolicyArgs | undefined;
            if ((!args || args.policy === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'policy'");
            }
            if ((!args || args.queueUrl === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'queueUrl'");
            }
            inputs["policy"] = args ? args.policy : undefined;
            inputs["queueUrl"] = args ? args.queueUrl : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(QueuePolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QueuePolicy resources.
 */
export interface QueuePolicyState {
    /**
     * The JSON policy for the SQS queue.
     */
    readonly policy?: pulumi.Input<string | PolicyDocument>;
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
     * The JSON policy for the SQS queue.
     */
    readonly policy: pulumi.Input<string | PolicyDocument>;
    /**
     * The URL of the SQS Queue to which to attach the policy
     */
    readonly queueUrl: pulumi.Input<string>;
}
