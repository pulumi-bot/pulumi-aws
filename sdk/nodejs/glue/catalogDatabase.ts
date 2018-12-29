// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class CatalogDatabase extends pulumi.CustomResource {
    /**
     * Get an existing CatalogDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogDatabaseState, opts?: pulumi.CustomResourceOptions): CatalogDatabase {
        return new CatalogDatabase(name, <any>state, { ...opts, id: id });
    }

    public readonly catalogId: pulumi.Output<string>;
    public readonly description: pulumi.Output<string | undefined>;
    public readonly locationUri: pulumi.Output<string | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly parameters: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a CatalogDatabase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CatalogDatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogDatabaseArgs | CatalogDatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CatalogDatabaseState = argsOrState as CatalogDatabaseState | undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["locationUri"] = state ? state.locationUri : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
        } else {
            const args = argsOrState as CatalogDatabaseArgs | undefined;
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["locationUri"] = args ? args.locationUri : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
        }
        super("aws:glue/catalogDatabase:CatalogDatabase", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogDatabase resources.
 */
export interface CatalogDatabaseState {
    readonly catalogId?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly locationUri?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a CatalogDatabase resource.
 */
export interface CatalogDatabaseArgs {
    readonly catalogId?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly locationUri?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
