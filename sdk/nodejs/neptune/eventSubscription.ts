// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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

    public /*out*/ readonly arn: pulumi.Output<string>;
    public /*out*/ readonly customerAwsId: pulumi.Output<string>;
    public readonly enabled: pulumi.Output<boolean | undefined>;
    public readonly eventCategories: pulumi.Output<string[] | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly namePrefix: pulumi.Output<string>;
    public readonly snsTopicArn: pulumi.Output<string>;
    public readonly sourceIds: pulumi.Output<string[] | undefined>;
    public readonly sourceType: pulumi.Output<string | undefined>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;

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
            const state: EventSubscriptionState = argsOrState as EventSubscriptionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["customerAwsId"] = state ? state.customerAwsId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["eventCategories"] = state ? state.eventCategories : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["snsTopicArn"] = state ? state.snsTopicArn : undefined;
            inputs["sourceIds"] = state ? state.sourceIds : undefined;
            inputs["sourceType"] = state ? state.sourceType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EventSubscriptionArgs | undefined;
            if (!args || args.snsTopicArn === undefined) {
                throw new Error("Missing required property 'snsTopicArn'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["eventCategories"] = args ? args.eventCategories : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["snsTopicArn"] = args ? args.snsTopicArn : undefined;
            inputs["sourceIds"] = args ? args.sourceIds : undefined;
            inputs["sourceType"] = args ? args.sourceType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["customerAwsId"] = undefined /*out*/;
        }
        super("aws:neptune/eventSubscription:EventSubscription", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventSubscription resources.
 */
export interface EventSubscriptionState {
    readonly arn?: pulumi.Input<string>;
    readonly customerAwsId?: pulumi.Input<string>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly eventCategories?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly snsTopicArn?: pulumi.Input<string>;
    readonly sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sourceType?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a EventSubscription resource.
 */
export interface EventSubscriptionArgs {
    readonly enabled?: pulumi.Input<boolean>;
    readonly eventCategories?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly snsTopicArn: pulumi.Input<string>;
    readonly sourceIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sourceType?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
