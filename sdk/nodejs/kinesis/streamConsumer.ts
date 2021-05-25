// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a Kinesis Stream Consumer.
 *
 * > **Note:** You can register up to 20 consumers per stream. A given consumer can only be registered with one stream at a time.
 *
 * For more details, see the [Amazon Kinesis Stream Consumer Documentation](https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleStream = new aws.kinesis.Stream("exampleStream", {shardCount: 1});
 * const exampleStreamConsumer = new aws.kinesis.StreamConsumer("exampleStreamConsumer", {streamArn: exampleStream.arn});
 * ```
 *
 * ## Import
 *
 * Kinesis Stream Consumers can be imported using the Amazon Resource Name (ARN) e.g.
 *
 * ```sh
 *  $ pulumi import aws:kinesis/streamConsumer:StreamConsumer example arn:aws:kinesis:us-west-2:123456789012:stream/example/consumer/example:1616044553
 * ```
 *
 *  [1]https://docs.aws.amazon.com/streams/latest/dev/amazon-kinesis-consumers.html
 */
export class StreamConsumer extends pulumi.CustomResource {
    /**
     * Get an existing StreamConsumer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamConsumerState, opts?: pulumi.CustomResourceOptions): StreamConsumer {
        return new StreamConsumer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:kinesis/streamConsumer:StreamConsumer';

    /**
     * Returns true if the given object is an instance of StreamConsumer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StreamConsumer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StreamConsumer.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the stream consumer.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Name of the stream consumer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the data stream the consumer is registered with.
     */
    public readonly streamArn!: pulumi.Output<string>;

    /**
     * Create a StreamConsumer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamConsumerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamConsumerArgs | StreamConsumerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamConsumerState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["streamArn"] = state ? state.streamArn : undefined;
        } else {
            const args = argsOrState as StreamConsumerArgs | undefined;
            if ((!args || args.streamArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamArn'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["streamArn"] = args ? args.streamArn : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(StreamConsumer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StreamConsumer resources.
 */
export interface StreamConsumerState {
    /**
     * Amazon Resource Name (ARN) of the stream consumer.
     */
    arn?: pulumi.Input<string>;
    /**
     * Approximate timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of when the stream consumer was created.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * Name of the stream consumer.
     */
    name?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the data stream the consumer is registered with.
     */
    streamArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a StreamConsumer resource.
 */
export interface StreamConsumerArgs {
    /**
     * Name of the stream consumer.
     */
    name?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the data stream the consumer is registered with.
     */
    streamArn: pulumi.Input<string>;
}
