// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an ElastiCache parameter group resource.
 * 
 * > **NOTE:** Attempting to remove the `reserved-memory` parameter when `family` is set to `redis2.6` or `redis2.8` may show a perpetual difference in Terraform due to an Elasticache API limitation. Leave that parameter configured with any value to workaround the issue.
 */
export class ParameterGroup extends pulumi.CustomResource {
    /**
     * Get an existing ParameterGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ParameterGroupState, opts?: pulumi.CustomResourceOptions): ParameterGroup {
        return new ParameterGroup(name, <any>state, { ...opts, id: id });
    }

    /**
     * The description of the ElastiCache parameter group. Defaults to "Managed by Terraform".
     */
    public readonly description: pulumi.Output<string>;
    /**
     * The family of the ElastiCache parameter group.
     */
    public readonly family: pulumi.Output<string>;
    /**
     * The name of the ElastiCache parameter.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * A list of ElastiCache parameters to apply.
     */
    public readonly parameters: pulumi.Output<{ name: string, value: string }[] | undefined>;

    /**
     * Create a ParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ParameterGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ParameterGroupArgs | ParameterGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ParameterGroupState = argsOrState as ParameterGroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["family"] = state ? state.family : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
        } else {
            const args = argsOrState as ParameterGroupArgs | undefined;
            if (!args || args.family === undefined) {
                throw new Error("Missing required property 'family'");
            }
            inputs["description"] = (args ? args.description : undefined) || "Managed by Pulumi";
            inputs["family"] = args ? args.family : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
        }
        super("aws:elasticache/parameterGroup:ParameterGroup", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ParameterGroup resources.
 */
export interface ParameterGroupState {
    /**
     * The description of the ElastiCache parameter group. Defaults to "Managed by Terraform".
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The family of the ElastiCache parameter group.
     */
    readonly family?: pulumi.Input<string>;
    /**
     * The name of the ElastiCache parameter.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of ElastiCache parameters to apply.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
}

/**
 * The set of arguments for constructing a ParameterGroup resource.
 */
export interface ParameterGroupArgs {
    /**
     * The description of the ElastiCache parameter group. Defaults to "Managed by Terraform".
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The family of the ElastiCache parameter group.
     */
    readonly family: pulumi.Input<string>;
    /**
     * The name of the ElastiCache parameter.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of ElastiCache parameters to apply.
     */
    readonly parameters?: pulumi.Input<pulumi.Input<{ name: pulumi.Input<string>, value: pulumi.Input<string> }>[]>;
}
