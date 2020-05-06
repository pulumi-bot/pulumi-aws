// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Redshift event subscription resource.
 * 
 * 
 * ## Attributes
 * 
 * The following additional atttributes are provided:
 * 
 * * `arn` - Amazon Resource Name (ARN) of the Redshift event notification subscription
 * * `id` - The name of the Redshift event notification subscription
 * * `customerAwsId` - The AWS customer account associated with the Redshift event notification subscription
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/redshift_event_subscription.html.markdown.
 */
export class EventSubscription extends pulumi.CustomResource {
    /**
     * Get an existing EventSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventSubscriptionState, opts?: pulumi.CustomResourceOptions): EventSubscription {
        return new EventSubscription(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:redshift/eventSubscription:EventSubscription';

    /**
     * Returns true if the given object is an instance of EventSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSubscription.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly customerAwsId!: pulumi.Output<string>;
    /**
     * A boolean flag to enable/disable the subscription. Defaults to true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
     */
    public readonly eventCategories!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the Redshift event subscription.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
     */
    public readonly severity!: pulumi.Output<string | undefined>;
    /**
     * The ARN of the SNS topic to send events to.
     */
    public readonly snsTopicArn!: pulumi.Output<string>;
    /**
     * A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
     */
    public readonly sourceIds!: pulumi.Output<string[] | undefined>;
    /**
     * The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
     */
    public readonly sourceType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a EventSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSubscriptionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventSubscriptionArgs | EventSubscriptionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventSubscriptionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["customerAwsId"] = state ? state.customerAwsId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["eventCategories"] = state ? state.eventCategories : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["severity"] = state ? state.severity : undefined;
            inputs["snsTopicArn"] = state ? state.snsTopicArn : undefined;
            inputs["sourceIds"] = state ? state.sourceIds : undefined;
            inputs["sourceType"] = state ? state.sourceType : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EventSubscriptionArgs | undefined;
            if (!args || args.snsTopicArn === undefined) {
                throw new Error("Missing required property 'snsTopicArn'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["eventCategories"] = args ? args.eventCategories : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["severity"] = args ? args.severity : undefined;
            inputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
            inputs["sourceIds"] = args ? args.sourceIds : undefined;
            inputs["sourceType"] = args ? args.sourceType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["customerAwsId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventSubscription.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventSubscription resources.
 */
export interface EventSubscriptionState {
    readonly arn?: pulumi.Input<string>;
    readonly customerAwsId?: pulumi.Input<string>;
    /**
     * A boolean flag to enable/disable the subscription. Defaults to true.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
     */
    readonly eventCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Redshift event subscription.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
     */
    readonly severity?: pulumi.Input<string>;
    /**
     * The ARN of the SNS topic to send events to.
     */
    readonly snsTopicArn?: pulumi.Input<string>;
    /**
     * A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
     */
    readonly sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
     */
    readonly sourceType?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EventSubscription resource.
 */
export interface EventSubscriptionArgs {
    /**
     * A boolean flag to enable/disable the subscription. Defaults to true.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
     */
    readonly eventCategories?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Redshift event subscription.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`.
     */
    readonly severity?: pulumi.Input<string>;
    /**
     * The ARN of the SNS topic to send events to.
     */
    readonly snsTopicArn: pulumi.Input<string>;
    /**
     * A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a sourceType must also be specified.
     */
    readonly sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, or `cluster-snapshot`. If not set, all sources will be subscribed to.
     */
    readonly sourceType?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
