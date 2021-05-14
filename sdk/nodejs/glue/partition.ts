// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Glue Partition Resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Partition("example", {
 *     databaseName: "some-database",
 *     tableName: "some-table",
 *     values: ["some-value"],
 * });
 * ```
 *
 * ## Import
 *
 * Glue Partitions can be imported with their catalog ID (usually AWS account ID), database name, table name and partition values e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/partition:Partition part 123456789012:MyDatabase:MyTable:val1#val2
 * ```
 */
export class Partition extends pulumi.CustomResource {
    /**
     * Get an existing Partition resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PartitionState, opts?: pulumi.CustomResourceOptions): Partition {
        return new Partition(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/partition:Partition';

    /**
     * Returns true if the given object is an instance of Partition.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Partition {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Partition.__pulumiType;
    }

    /**
     * ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * The time at which the partition was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The last time at which the partition was accessed.
     */
    public /*out*/ readonly lastAccessedTime!: pulumi.Output<string>;
    /**
     * The last time at which column statistics were computed for this partition.
     */
    public /*out*/ readonly lastAnalyzedTime!: pulumi.Output<string>;
    /**
     * A map of initialization parameters for the SerDe, in key-value form.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The values that define the partition.
     */
    public readonly partitionValues!: pulumi.Output<string[]>;
    /**
     * A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
     */
    public readonly storageDescriptor!: pulumi.Output<outputs.glue.PartitionStorageDescriptor | undefined>;
    public readonly tableName!: pulumi.Output<string>;

    /**
     * Create a Partition resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PartitionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PartitionArgs | PartitionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PartitionState | undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["lastAccessedTime"] = state ? state.lastAccessedTime : undefined;
            inputs["lastAnalyzedTime"] = state ? state.lastAnalyzedTime : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["partitionValues"] = state ? state.partitionValues : undefined;
            inputs["storageDescriptor"] = state ? state.storageDescriptor : undefined;
            inputs["tableName"] = state ? state.tableName : undefined;
        } else {
            const args = argsOrState as PartitionArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.partitionValues === undefined) && !opts.urn) {
                throw new Error("Missing required property 'partitionValues'");
            }
            if ((!args || args.tableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableName'");
            }
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["partitionValues"] = args ? args.partitionValues : undefined;
            inputs["storageDescriptor"] = args ? args.storageDescriptor : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastAccessedTime"] = undefined /*out*/;
            inputs["lastAnalyzedTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Partition.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Partition resources.
 */
export interface PartitionState {
    /**
     * ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * The time at which the partition was created.
     */
    creationTime?: pulumi.Input<string>;
    /**
     * Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The last time at which the partition was accessed.
     */
    lastAccessedTime?: pulumi.Input<string>;
    /**
     * The last time at which column statistics were computed for this partition.
     */
    lastAnalyzedTime?: pulumi.Input<string>;
    /**
     * A map of initialization parameters for the SerDe, in key-value form.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The values that define the partition.
     */
    partitionValues?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
     */
    storageDescriptor?: pulumi.Input<inputs.glue.PartitionStorageDescriptor>;
    tableName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Partition resource.
 */
export interface PartitionArgs {
    /**
     * ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
     */
    databaseName: pulumi.Input<string>;
    /**
     * A map of initialization parameters for the SerDe, in key-value form.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The values that define the partition.
     */
    partitionValues: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
     */
    storageDescriptor?: pulumi.Input<inputs.glue.PartitionStorageDescriptor>;
    tableName: pulumi.Input<string>;
}
