// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an AWS Config Delivery Channel.
 * 
 * > **Note:** Delivery Channel requires a [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `dependsOn` (as shown below) is recommended to avoid race conditions.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_delivery_channel.html.markdown.
 */
export class DeliveryChannel extends pulumi.CustomResource {
    /**
     * Get an existing DeliveryChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeliveryChannelState, opts?: pulumi.CustomResourceOptions): DeliveryChannel {
        return new DeliveryChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cfg/deliveryChannel:DeliveryChannel';

    /**
     * Returns true if the given object is an instance of DeliveryChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeliveryChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeliveryChannel.__pulumiType;
    }

    /**
     * The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the S3 bucket used to store the configuration history.
     */
    public readonly s3BucketName!: pulumi.Output<string>;
    /**
     * The prefix for the specified S3 bucket.
     */
    public readonly s3KeyPrefix!: pulumi.Output<string | undefined>;
    /**
     * Options for how AWS Config delivers configuration snapshots. See below
     */
    public readonly snapshotDeliveryProperties!: pulumi.Output<outputs.cfg.DeliveryChannelSnapshotDeliveryProperties | undefined>;
    /**
     * The ARN of the SNS topic that AWS Config delivers notifications to.
     */
    public readonly snsTopicArn!: pulumi.Output<string | undefined>;

    /**
     * Create a DeliveryChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeliveryChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeliveryChannelArgs | DeliveryChannelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DeliveryChannelState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["s3BucketName"] = state ? state.s3BucketName : undefined;
            inputs["s3KeyPrefix"] = state ? state.s3KeyPrefix : undefined;
            inputs["snapshotDeliveryProperties"] = state ? state.snapshotDeliveryProperties : undefined;
            inputs["snsTopicArn"] = state ? state.snsTopicArn : undefined;
        } else {
            const args = argsOrState as DeliveryChannelArgs | undefined;
            if (!args || args.s3BucketName === undefined) {
                throw new Error("Missing required property 's3BucketName'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["s3BucketName"] = args ? args.s3BucketName : undefined;
            inputs["s3KeyPrefix"] = args ? args.s3KeyPrefix : undefined;
            inputs["snapshotDeliveryProperties"] = args ? args.snapshotDeliveryProperties : undefined;
            inputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DeliveryChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeliveryChannel resources.
 */
export interface DeliveryChannelState {
    /**
     * The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the S3 bucket used to store the configuration history.
     */
    readonly s3BucketName?: pulumi.Input<string>;
    /**
     * The prefix for the specified S3 bucket.
     */
    readonly s3KeyPrefix?: pulumi.Input<string>;
    /**
     * Options for how AWS Config delivers configuration snapshots. See below
     */
    readonly snapshotDeliveryProperties?: pulumi.Input<inputs.cfg.DeliveryChannelSnapshotDeliveryProperties>;
    /**
     * The ARN of the SNS topic that AWS Config delivers notifications to.
     */
    readonly snsTopicArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeliveryChannel resource.
 */
export interface DeliveryChannelArgs {
    /**
     * The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the S3 bucket used to store the configuration history.
     */
    readonly s3BucketName: pulumi.Input<string>;
    /**
     * The prefix for the specified S3 bucket.
     */
    readonly s3KeyPrefix?: pulumi.Input<string>;
    /**
     * Options for how AWS Config delivers configuration snapshots. See below
     */
    readonly snapshotDeliveryProperties?: pulumi.Input<inputs.cfg.DeliveryChannelSnapshotDeliveryProperties>;
    /**
     * The ARN of the SNS topic that AWS Config delivers notifications to.
     */
    readonly snsTopicArn?: pulumi.Input<string>;
}
