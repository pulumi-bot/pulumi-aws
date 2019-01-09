// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    public readonly actions: pulumi.Output<{ arguments?: {[key: string]: any}, jobName: string, timeout?: number }[]>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly enabled: pulumi.Output<boolean | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly predicate: pulumi.Output<{ conditions: { jobName: string, logicalOperator?: string, state: string }[], logical?: string } | undefined>;
    public readonly schedule: pulumi.Output<string | undefined>;
    public readonly type: pulumi.Output<string>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TriggerState = argsOrState as TriggerState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["predicate"] = state ? state.predicate : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            if (!args || args.actions === undefined) {
                throw new Error("Missing required property 'actions'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["predicate"] = args ? args.predicate : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        super("aws:glue/trigger:Trigger", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    readonly actions?: pulumi.Input<pulumi.Input<{ arguments?: pulumi.Input<{[key: string]: any}>, jobName: pulumi.Input<string>, timeout?: pulumi.Input<number> }>[]>;
    readonly description?: pulumi.Input<string>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly predicate?: pulumi.Input<{ conditions: pulumi.Input<pulumi.Input<{ jobName: pulumi.Input<string>, logicalOperator?: pulumi.Input<string>, state: pulumi.Input<string> }>[]>, logical?: pulumi.Input<string> }>;
    readonly schedule?: pulumi.Input<string>;
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    readonly actions: pulumi.Input<pulumi.Input<{ arguments?: pulumi.Input<{[key: string]: any}>, jobName: pulumi.Input<string>, timeout?: pulumi.Input<number> }>[]>;
    readonly description?: pulumi.Input<string>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly predicate?: pulumi.Input<{ conditions: pulumi.Input<pulumi.Input<{ jobName: pulumi.Input<string>, logicalOperator?: pulumi.Input<string>, state: pulumi.Input<string> }>[]>, logical?: pulumi.Input<string> }>;
    readonly schedule?: pulumi.Input<string>;
    readonly type: pulumi.Input<string>;
}
