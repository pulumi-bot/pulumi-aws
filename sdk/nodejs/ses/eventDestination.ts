// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an SES event destination
 * 
 * ## Example Usage
 * 
 * ### CloudWatch Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const cloudwatch = new aws.ses.EventDestination("cloudwatch", {
 *     cloudwatchDestinations: [{
 *         defaultValue: "default",
 *         dimensionName: "dimension",
 *         valueSource: "emailHeader",
 *     }],
 *     configurationSetName: aws_ses_configuration_set_example.name,
 *     enabled: true,
 *     matchingTypes: [
 *         "bounce",
 *         "send",
 *     ],
 * });
 * ```
 * 
 * ### Kinesis Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const kinesis = new aws.ses.EventDestination("kinesis", {
 *     configurationSetName: aws_ses_configuration_set_example.name,
 *     enabled: true,
 *     kinesisDestination: {
 *         roleArn: aws_iam_role_example.arn,
 *         streamArn: aws_kinesis_firehose_delivery_stream_example.arn,
 *     },
 *     matchingTypes: [
 *         "bounce",
 *         "send",
 *     ],
 * });
 * ```
 * 
 * ### SNS Destination
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const sns = new aws.ses.EventDestination("sns", {
 *     configurationSetName: aws_ses_configuration_set_example.name,
 *     enabled: true,
 *     matchingTypes: [
 *         "bounce",
 *         "send",
 *     ],
 *     snsDestination: {
 *         topicArn: aws_sns_topic_example.arn,
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_event_destination.html.markdown.
 */
export class EventDestination extends pulumi.CustomResource {
    /**
     * Get an existing EventDestination resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventDestinationState, opts?: pulumi.CustomResourceOptions): EventDestination {
        return new EventDestination(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ses/eventDestination:EventDestination';

    /**
     * Returns true if the given object is an instance of EventDestination.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventDestination {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventDestination.__pulumiType;
    }

    /**
     * CloudWatch destination for the events
     */
    public readonly cloudwatchDestinations!: pulumi.Output<outputs.ses.EventDestinationCloudwatchDestination[] | undefined>;
    /**
     * The name of the configuration set
     */
    public readonly configurationSetName!: pulumi.Output<string>;
    /**
     * If true, the event destination will be enabled
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Send the events to a kinesis firehose destination
     */
    public readonly kinesisDestination!: pulumi.Output<outputs.ses.EventDestinationKinesisDestination | undefined>;
    /**
     * A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
     */
    public readonly matchingTypes!: pulumi.Output<string[]>;
    /**
     * The name of the event destination
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Send the events to an SNS Topic destination
     */
    public readonly snsDestination!: pulumi.Output<outputs.ses.EventDestinationSnsDestination | undefined>;

    /**
     * Create a EventDestination resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventDestinationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventDestinationArgs | EventDestinationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventDestinationState | undefined;
            inputs["cloudwatchDestinations"] = state ? state.cloudwatchDestinations : undefined;
            inputs["configurationSetName"] = state ? state.configurationSetName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["kinesisDestination"] = state ? state.kinesisDestination : undefined;
            inputs["matchingTypes"] = state ? state.matchingTypes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["snsDestination"] = state ? state.snsDestination : undefined;
        } else {
            const args = argsOrState as EventDestinationArgs | undefined;
            if (!args || args.configurationSetName === undefined) {
                throw new Error("Missing required property 'configurationSetName'");
            }
            if (!args || args.matchingTypes === undefined) {
                throw new Error("Missing required property 'matchingTypes'");
            }
            inputs["cloudwatchDestinations"] = args ? args.cloudwatchDestinations : undefined;
            inputs["configurationSetName"] = args ? args.configurationSetName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["kinesisDestination"] = args ? args.kinesisDestination : undefined;
            inputs["matchingTypes"] = args ? args.matchingTypes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["snsDestination"] = args ? args.snsDestination : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventDestination.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventDestination resources.
 */
export interface EventDestinationState {
    /**
     * CloudWatch destination for the events
     */
    readonly cloudwatchDestinations?: pulumi.Input<pulumi.Input<inputs.ses.EventDestinationCloudwatchDestination>[]>;
    /**
     * The name of the configuration set
     */
    readonly configurationSetName?: pulumi.Input<string>;
    /**
     * If true, the event destination will be enabled
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Send the events to a kinesis firehose destination
     */
    readonly kinesisDestination?: pulumi.Input<inputs.ses.EventDestinationKinesisDestination>;
    /**
     * A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
     */
    readonly matchingTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the event destination
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Send the events to an SNS Topic destination
     */
    readonly snsDestination?: pulumi.Input<inputs.ses.EventDestinationSnsDestination>;
}

/**
 * The set of arguments for constructing a EventDestination resource.
 */
export interface EventDestinationArgs {
    /**
     * CloudWatch destination for the events
     */
    readonly cloudwatchDestinations?: pulumi.Input<pulumi.Input<inputs.ses.EventDestinationCloudwatchDestination>[]>;
    /**
     * The name of the configuration set
     */
    readonly configurationSetName: pulumi.Input<string>;
    /**
     * If true, the event destination will be enabled
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Send the events to a kinesis firehose destination
     */
    readonly kinesisDestination?: pulumi.Input<inputs.ses.EventDestinationKinesisDestination>;
    /**
     * A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
     */
    readonly matchingTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the event destination
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Send the events to an SNS Topic destination
     */
    readonly snsDestination?: pulumi.Input<inputs.ses.EventDestinationSnsDestination>;
}
