// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Glue Catalog Database Resource. You can refer to the [Glue Developer Guide](http://docs.aws.amazon.com/glue/latest/dg/populate-data-catalog.html) for a full explanation of the Glue Data Catalog functionality
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const awsGlueCatalogDatabase = new aws.glue.CatalogDatabase("aws_glue_catalog_database", {
 *     name: "MyCatalogDatabase",
 * });
 * ```
 *
 * ## Import
 *
 * Glue Catalog Databases can be imported using the `catalog_id:name`. If you have not set a Catalog ID specify the AWS Account ID that the database is in, e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/catalogDatabase:CatalogDatabase database 123456789012:my_database
 * ```
 */
export class CatalogDatabase extends pulumi.CustomResource {
    /**
     * Get an existing CatalogDatabase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogDatabaseState, opts?: pulumi.CustomResourceOptions): CatalogDatabase {
        return new CatalogDatabase(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/catalogDatabase:CatalogDatabase';

    /**
     * Returns true if the given object is an instance of CatalogDatabase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogDatabase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogDatabase.__pulumiType;
    }

    /**
     * ARN of the Glue Catalog Database.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ID of the Data Catalog in which the database resides.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * Description of the database.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Location of the database (for example, an HDFS path).
     */
    public readonly locationUri!: pulumi.Output<string | undefined>;
    /**
     * Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of key-value pairs that define parameters and properties of the database.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Configuration block for a target database for resource linking. See `targetDatabase` below.
     */
    public readonly targetDatabase!: pulumi.Output<outputs.glue.CatalogDatabaseTargetDatabase | undefined>;

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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogDatabaseState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["locationUri"] = state ? state.locationUri : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["targetDatabase"] = state ? state.targetDatabase : undefined;
        } else {
            const args = argsOrState as CatalogDatabaseArgs | undefined;
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["locationUri"] = args ? args.locationUri : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["targetDatabase"] = args ? args.targetDatabase : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CatalogDatabase.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogDatabase resources.
 */
export interface CatalogDatabaseState {
    /**
     * ARN of the Glue Catalog Database.
     */
    arn?: pulumi.Input<string>;
    /**
     * ID of the Data Catalog in which the database resides.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Description of the database.
     */
    description?: pulumi.Input<string>;
    /**
     * Location of the database (for example, an HDFS path).
     */
    locationUri?: pulumi.Input<string>;
    /**
     * Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
     */
    name?: pulumi.Input<string>;
    /**
     * List of key-value pairs that define parameters and properties of the database.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for a target database for resource linking. See `targetDatabase` below.
     */
    targetDatabase?: pulumi.Input<inputs.glue.CatalogDatabaseTargetDatabase>;
}

/**
 * The set of arguments for constructing a CatalogDatabase resource.
 */
export interface CatalogDatabaseArgs {
    /**
     * ID of the Data Catalog in which the database resides.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Description of the database.
     */
    description?: pulumi.Input<string>;
    /**
     * Location of the database (for example, an HDFS path).
     */
    locationUri?: pulumi.Input<string>;
    /**
     * Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
     */
    name?: pulumi.Input<string>;
    /**
     * List of key-value pairs that define parameters and properties of the database.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for a target database for resource linking. See `targetDatabase` below.
     */
    targetDatabase?: pulumi.Input<inputs.glue.CatalogDatabaseTargetDatabase>;
}
