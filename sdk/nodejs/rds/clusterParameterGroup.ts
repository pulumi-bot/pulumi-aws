// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ClusterParameterGroup extends pulumi.CustomResource {
    /**
     * Get an existing ClusterParameterGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterParameterGroupState, opts?: pulumi.CustomResourceOptions): ClusterParameterGroup {
        return new ClusterParameterGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/clusterParameterGroup:ClusterParameterGroup';

    /**
     * Returns true if the given object is an instance of ClusterParameterGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterParameterGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterParameterGroup.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string>;
    public readonly family!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string>;
    public readonly parameters!: pulumi.Output<outputs.rds.ClusterParameterGroupParameter[] | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a ClusterParameterGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterParameterGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterParameterGroupArgs | ClusterParameterGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterParameterGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["family"] = state ? state.family : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ClusterParameterGroupArgs | undefined;
            if (!args || args.family === undefined) {
                throw new Error("Missing required property 'family'");
            }
            inputs["description"] = (args ? args.description : undefined) || "Managed by Pulumi";
            inputs["family"] = args ? args.family : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClusterParameterGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterParameterGroup resources.
 */
export interface ClusterParameterGroupState {
    readonly arn?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly family?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.rds.ClusterParameterGroupParameter>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a ClusterParameterGroup resource.
 */
export interface ClusterParameterGroupArgs {
    readonly description?: pulumi.Input<string>;
    readonly family: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly parameters?: pulumi.Input<pulumi.Input<inputs.rds.ClusterParameterGroupParameter>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
