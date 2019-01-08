// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RateBasedRule extends pulumi.CustomResource {
    /**
     * Get an existing RateBasedRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RateBasedRuleState, opts?: pulumi.CustomResourceOptions): RateBasedRule {
        return new RateBasedRule(name, <any>state, { ...opts, id: id });
    }

    public readonly metricName: pulumi.Output<string>;
    public readonly name: pulumi.Output<string>;
    public readonly predicates: pulumi.Output<{ dataId: string, negated: boolean, type: string }[] | undefined>;
    public readonly rateKey: pulumi.Output<string>;
    public readonly rateLimit: pulumi.Output<number>;

    /**
     * Create a RateBasedRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RateBasedRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RateBasedRuleArgs | RateBasedRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RateBasedRuleState = argsOrState as RateBasedRuleState | undefined;
            inputs["metricName"] = state ? state.metricName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["predicates"] = state ? state.predicates : undefined;
            inputs["rateKey"] = state ? state.rateKey : undefined;
            inputs["rateLimit"] = state ? state.rateLimit : undefined;
        } else {
            const args = argsOrState as RateBasedRuleArgs | undefined;
            if (!args || args.metricName === undefined) {
                throw new Error("Missing required property 'metricName'");
            }
            if (!args || args.rateKey === undefined) {
                throw new Error("Missing required property 'rateKey'");
            }
            if (!args || args.rateLimit === undefined) {
                throw new Error("Missing required property 'rateLimit'");
            }
            inputs["metricName"] = args ? args.metricName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["predicates"] = args ? args.predicates : undefined;
            inputs["rateKey"] = args ? args.rateKey : undefined;
            inputs["rateLimit"] = args ? args.rateLimit : undefined;
        }
        super("aws:waf/rateBasedRule:RateBasedRule", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RateBasedRule resources.
 */
export interface RateBasedRuleState {
    readonly metricName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly predicates?: pulumi.Input<pulumi.Input<{ dataId: pulumi.Input<string>, negated: pulumi.Input<boolean>, type: pulumi.Input<string> }>[]>;
    readonly rateKey?: pulumi.Input<string>;
    readonly rateLimit?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a RateBasedRule resource.
 */
export interface RateBasedRuleArgs {
    readonly metricName: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly predicates?: pulumi.Input<pulumi.Input<{ dataId: pulumi.Input<string>, negated: pulumi.Input<boolean>, type: pulumi.Input<string> }>[]>;
    readonly rateKey: pulumi.Input<string>;
    readonly rateLimit: pulumi.Input<number>;
}
