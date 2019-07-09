// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ARN of a topic in AWS Simple Notification
 * Service (SNS). By using this data source, you can reference SNS topics
 * without having to hard code the ARNs as input.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = pulumi.output(aws.sns.getTopic({
 *     name: "an_example_topic",
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/sns_topic.html.markdown.
 */
export function getTopic(args: GetTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicResult> {
    return pulumi.runtime.invoke("aws:sns/getTopic:getTopic", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTopic.
 */
export interface GetTopicArgs {
    /**
     * The friendly name of the topic to match.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getTopic.
 */
export interface GetTopicResult {
    /**
     * Set to the ARN of the found topic, suitable for referencing in other resources that support SNS topics.
     */
    readonly arn: string;
    readonly name: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
