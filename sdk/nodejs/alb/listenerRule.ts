// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ListenerRule extends pulumi.CustomResource {
    /**
     * Get an existing ListenerRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ListenerRuleState, opts?: pulumi.CustomResourceOptions): ListenerRule {
        return new ListenerRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:alb/listenerRule:ListenerRule';

    /**
     * Returns true if the given object is an instance of ListenerRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ListenerRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ListenerRule.__pulumiType;
    }

    public readonly actions!: pulumi.Output<outputs.alb.ListenerRuleAction[]>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly conditions!: pulumi.Output<outputs.alb.ListenerRuleCondition[]>;
    public readonly listenerArn!: pulumi.Output<string>;
    public readonly priority!: pulumi.Output<number>;

    /**
     * Create a ListenerRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ListenerRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ListenerRuleArgs | ListenerRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ListenerRuleState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["conditions"] = state ? state.conditions : undefined;
            inputs["listenerArn"] = state ? state.listenerArn : undefined;
            inputs["priority"] = state ? state.priority : undefined;
        } else {
            const args = argsOrState as ListenerRuleArgs | undefined;
            if (!args || args.actions === undefined) {
                throw new Error("Missing required property 'actions'");
            }
            if (!args || args.conditions === undefined) {
                throw new Error("Missing required property 'conditions'");
            }
            if (!args || args.listenerArn === undefined) {
                throw new Error("Missing required property 'listenerArn'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["conditions"] = args ? args.conditions : undefined;
            inputs["listenerArn"] = args ? args.listenerArn : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "aws:applicationloadbalancing/listenerRule:ListenerRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ListenerRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ListenerRule resources.
 */
export interface ListenerRuleState {
    readonly actions?: pulumi.Input<pulumi.Input<inputs.alb.ListenerRuleAction>[]>;
    readonly arn?: pulumi.Input<string>;
    readonly conditions?: pulumi.Input<pulumi.Input<inputs.alb.ListenerRuleCondition>[]>;
    readonly listenerArn?: pulumi.Input<string>;
    readonly priority?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ListenerRule resource.
 */
export interface ListenerRuleArgs {
    readonly actions: pulumi.Input<pulumi.Input<inputs.alb.ListenerRuleAction>[]>;
    readonly conditions: pulumi.Input<pulumi.Input<inputs.alb.ListenerRuleCondition>[]>;
    readonly listenerArn: pulumi.Input<string>;
    readonly priority?: pulumi.Input<number>;
}
