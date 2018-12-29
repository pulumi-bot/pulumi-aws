// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    public readonly catalogId: pulumi.Output<string>;
    public readonly connectionProperties: pulumi.Output<{[key: string]: any}>;
    public readonly connectionType: pulumi.Output<string | undefined>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly matchCriterias: pulumi.Output<string[] | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly physicalConnectionRequirements: pulumi.Output<{ availabilityZone?: string, securityGroupIdLists?: string[], subnetId?: string } | undefined>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: ConnectionState = argsOrState as ConnectionState | undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["connectionProperties"] = state ? state.connectionProperties : undefined;
            inputs["connectionType"] = state ? state.connectionType : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["matchCriterias"] = state ? state.matchCriterias : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["physicalConnectionRequirements"] = state ? state.physicalConnectionRequirements : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if (!args || args.connectionProperties === undefined) {
                throw new Error("Missing required property 'connectionProperties'");
            }
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["connectionProperties"] = args ? args.connectionProperties : undefined;
            inputs["connectionType"] = args ? args.connectionType : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["matchCriterias"] = args ? args.matchCriterias : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["physicalConnectionRequirements"] = args ? args.physicalConnectionRequirements : undefined;
        }
        super("aws:glue/connection:Connection", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    readonly catalogId?: pulumi.Input<string>;
    readonly connectionProperties?: pulumi.Input<{[key: string]: any}>;
    readonly connectionType?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly matchCriterias?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly physicalConnectionRequirements?: pulumi.Input<{ availabilityZone?: pulumi.Input<string>, securityGroupIdLists?: pulumi.Input<pulumi.Input<string>[]>, subnetId?: pulumi.Input<string> }>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    readonly catalogId?: pulumi.Input<string>;
    readonly connectionProperties: pulumi.Input<{[key: string]: any}>;
    readonly connectionType?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly matchCriterias?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly physicalConnectionRequirements?: pulumi.Input<{ availabilityZone?: pulumi.Input<string>, securityGroupIdLists?: pulumi.Input<pulumi.Input<string>[]>, subnetId?: pulumi.Input<string> }>;
}
