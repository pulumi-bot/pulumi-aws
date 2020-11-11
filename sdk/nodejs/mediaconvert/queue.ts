// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an AWS Elemental MediaConvert Queue.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.mediaconvert.Queue("test", {});
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
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueueState, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mediaconvert/queue:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * The Arn of the queue
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A description of the queue
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A unique identifier describing the queue
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
     */
    public readonly pricingPlan!: pulumi.Output<string | undefined>;
    /**
     * A detail pricing plan of the  reserved queue. See below.
     */
    public readonly reservationPlanSettings!: pulumi.Output<outputs.mediaconvert.QueueReservationPlanSettings>;
    /**
     * A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

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
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pricingPlan"] = state ? state.pricingPlan : undefined;
            inputs["reservationPlanSettings"] = state ? state.reservationPlanSettings : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as QueueArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pricingPlan"] = args ? args.pricingPlan : undefined;
            inputs["reservationPlanSettings"] = args ? args.reservationPlanSettings : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Queue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Queue resources.
 */
export interface QueueState {
    /**
     * The Arn of the queue
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A description of the queue
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique identifier describing the queue
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
     */
    readonly pricingPlan?: pulumi.Input<string>;
    /**
     * A detail pricing plan of the  reserved queue. See below.
     */
    readonly reservationPlanSettings?: pulumi.Input<inputs.mediaconvert.QueueReservationPlanSettings>;
    /**
     * A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * A description of the queue
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A unique identifier describing the queue
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies whether the pricing plan for the queue is on-demand or reserved. Valid values are `ON_DEMAND` or `RESERVED`. Default to `ON_DEMAND`.
     */
    readonly pricingPlan?: pulumi.Input<string>;
    /**
     * A detail pricing plan of the  reserved queue. See below.
     */
    readonly reservationPlanSettings?: pulumi.Input<inputs.mediaconvert.QueueReservationPlanSettings>;
    /**
     * A status of the queue. Valid values are `ACTIVE` or `RESERVED`. Default to `PAUSED`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
