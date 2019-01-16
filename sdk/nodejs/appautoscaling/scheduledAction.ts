// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ScheduledAction extends pulumi.CustomResource {
    /**
     * Get an existing ScheduledAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduledActionState, opts?: pulumi.CustomResourceOptions): ScheduledAction {
        return new ScheduledAction(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly arn: pulumi.Output<string>;
    public readonly endTime: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly resourceId: pulumi.Output<string>;
    public readonly scalableDimension: pulumi.Output<string | undefined>;
    public readonly scalableTargetAction: pulumi.Output<{ maxCapacity?: number, minCapacity?: number } | undefined>;
    public readonly schedule: pulumi.Output<string | undefined>;
    public readonly serviceNamespace: pulumi.Output<string>;
    public readonly startTime: pulumi.Output<string | undefined>;

    /**
     * Create a ScheduledAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduledActionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduledActionArgs | ScheduledActionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ScheduledActionState = argsOrState as ScheduledActionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["scalableDimension"] = state ? state.scalableDimension : undefined;
            inputs["scalableTargetAction"] = state ? state.scalableTargetAction : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["serviceNamespace"] = state ? state.serviceNamespace : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as ScheduledActionArgs | undefined;
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.serviceNamespace === undefined) {
                throw new Error("Missing required property 'serviceNamespace'");
            }
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["scalableDimension"] = args ? args.scalableDimension : undefined;
            inputs["scalableTargetAction"] = args ? args.scalableTargetAction : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["serviceNamespace"] = args ? args.serviceNamespace : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:appautoscaling/scheduledAction:ScheduledAction", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScheduledAction resources.
 */
export interface ScheduledActionState {
    readonly arn?: pulumi.Input<string>;
    readonly endTime?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly resourceId?: pulumi.Input<string>;
    readonly scalableDimension?: pulumi.Input<string>;
    readonly scalableTargetAction?: pulumi.Input<{ maxCapacity?: pulumi.Input<number>, minCapacity?: pulumi.Input<number> }>;
    readonly schedule?: pulumi.Input<string>;
    readonly serviceNamespace?: pulumi.Input<string>;
    readonly startTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScheduledAction resource.
 */
export interface ScheduledActionArgs {
    readonly endTime?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly resourceId: pulumi.Input<string>;
    readonly scalableDimension?: pulumi.Input<string>;
    readonly scalableTargetAction?: pulumi.Input<{ maxCapacity?: pulumi.Input<number>, minCapacity?: pulumi.Input<number> }>;
    readonly schedule?: pulumi.Input<string>;
    readonly serviceNamespace: pulumi.Input<string>;
    readonly startTime?: pulumi.Input<string>;
}
