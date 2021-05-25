// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Glue Catalog Table Resource. You can refer to the [Glue Developer Guide](http://docs.aws.amazon.com/glue/latest/dg/populate-data-catalog.html) for a full explanation of the Glue Data Catalog functionality.
 *
 * ## Example Usage
 * ### Basic Table
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const awsGlueCatalogTable = new aws.glue.CatalogTable("aws_glue_catalog_table", {
 *     databaseName: "MyCatalogDatabase",
 *     name: "MyCatalogTable",
 * });
 * ```
 * ### Parquet Table for Athena
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const awsGlueCatalogTable = new aws.glue.CatalogTable("aws_glue_catalog_table", {
 *     databaseName: "MyCatalogDatabase",
 *     name: "MyCatalogTable",
 *     parameters: {
 *         EXTERNAL: "TRUE",
 *         "parquet.compression": "SNAPPY",
 *     },
 *     storageDescriptor: {
 *         columns: [
 *             {
 *                 name: "my_string",
 *                 type: "string",
 *             },
 *             {
 *                 name: "my_double",
 *                 type: "double",
 *             },
 *             {
 *                 comment: "",
 *                 name: "my_date",
 *                 type: "date",
 *             },
 *             {
 *                 comment: "",
 *                 name: "my_bigint",
 *                 type: "bigint",
 *             },
 *             {
 *                 comment: "",
 *                 name: "my_struct",
 *                 type: "struct<my_nested_string:string>",
 *             },
 *         ],
 *         inputFormat: "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat",
 *         location: "s3://my-bucket/event-streams/my-stream",
 *         outputFormat: "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat",
 *         serDeInfo: {
 *             name: "my-stream",
 *             parameters: {
 *                 "serialization.format": 1,
 *             },
 *             serializationLibrary: "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe",
 *         },
 *     },
 *     tableType: "EXTERNAL_TABLE",
 * });
 * ```
 *
 * ## Import
 *
 * Glue Tables can be imported with their catalog ID (usually AWS account ID), database name, and table name, e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/catalogTable:CatalogTable MyTable 123456789012:MyDatabase:MyTable
 * ```
 */
export class CatalogTable extends pulumi.CustomResource {
    /**
     * Get an existing CatalogTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogTableState, opts?: pulumi.CustomResourceOptions): CatalogTable {
        return new CatalogTable(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/catalogTable:CatalogTable';

    /**
     * Returns true if the given object is an instance of CatalogTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogTable.__pulumiType;
    }

    /**
     * The ARN of the Glue Table.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ID of the Data Catalog in which the table resides.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * Name of the catalog database that contains the target table.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * Description of the table.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the target table.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Owner of the table.
     */
    public readonly owner!: pulumi.Output<string | undefined>;
    /**
     * Map of initialization parameters for the SerDe, in key-value form.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
     */
    public readonly partitionIndices!: pulumi.Output<outputs.glue.CatalogTablePartitionIndex[] | undefined>;
    /**
     * Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
     */
    public readonly partitionKeys!: pulumi.Output<outputs.glue.CatalogTablePartitionKey[] | undefined>;
    /**
     * Retention time for this table.
     */
    public readonly retention!: pulumi.Output<number | undefined>;
    /**
     * Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables. html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
     */
    public readonly storageDescriptor!: pulumi.Output<outputs.glue.CatalogTableStorageDescriptor | undefined>;
    /**
     * Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
     */
    public readonly tableType!: pulumi.Output<string | undefined>;
    /**
     * Configuration block of a target table for resource linking. See `targetTable` below.
     */
    public readonly targetTable!: pulumi.Output<outputs.glue.CatalogTableTargetTable | undefined>;
    /**
     * If the table is a view, the expanded text of the view; otherwise null.
     */
    public readonly viewExpandedText!: pulumi.Output<string | undefined>;
    /**
     * If the table is a view, the original text of the view; otherwise null.
     */
    public readonly viewOriginalText!: pulumi.Output<string | undefined>;

    /**
     * Create a CatalogTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogTableArgs | CatalogTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogTableState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["owner"] = state ? state.owner : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["partitionIndices"] = state ? state.partitionIndices : undefined;
            inputs["partitionKeys"] = state ? state.partitionKeys : undefined;
            inputs["retention"] = state ? state.retention : undefined;
            inputs["storageDescriptor"] = state ? state.storageDescriptor : undefined;
            inputs["tableType"] = state ? state.tableType : undefined;
            inputs["targetTable"] = state ? state.targetTable : undefined;
            inputs["viewExpandedText"] = state ? state.viewExpandedText : undefined;
            inputs["viewOriginalText"] = state ? state.viewOriginalText : undefined;
        } else {
            const args = argsOrState as CatalogTableArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["owner"] = args ? args.owner : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["partitionIndices"] = args ? args.partitionIndices : undefined;
            inputs["partitionKeys"] = args ? args.partitionKeys : undefined;
            inputs["retention"] = args ? args.retention : undefined;
            inputs["storageDescriptor"] = args ? args.storageDescriptor : undefined;
            inputs["tableType"] = args ? args.tableType : undefined;
            inputs["targetTable"] = args ? args.targetTable : undefined;
            inputs["viewExpandedText"] = args ? args.viewExpandedText : undefined;
            inputs["viewOriginalText"] = args ? args.viewOriginalText : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CatalogTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogTable resources.
 */
export interface CatalogTableState {
    /**
     * The ARN of the Glue Table.
     */
    arn?: pulumi.Input<string>;
    /**
     * ID of the Data Catalog in which the table resides.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Name of the catalog database that contains the target table.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * Description of the table.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the target table.
     */
    name?: pulumi.Input<string>;
    /**
     * Owner of the table.
     */
    owner?: pulumi.Input<string>;
    /**
     * Map of initialization parameters for the SerDe, in key-value form.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
     */
    partitionIndices?: pulumi.Input<pulumi.Input<inputs.glue.CatalogTablePartitionIndex>[]>;
    /**
     * Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
     */
    partitionKeys?: pulumi.Input<pulumi.Input<inputs.glue.CatalogTablePartitionKey>[]>;
    /**
     * Retention time for this table.
     */
    retention?: pulumi.Input<number>;
    /**
     * Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables. html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
     */
    storageDescriptor?: pulumi.Input<inputs.glue.CatalogTableStorageDescriptor>;
    /**
     * Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
     */
    tableType?: pulumi.Input<string>;
    /**
     * Configuration block of a target table for resource linking. See `targetTable` below.
     */
    targetTable?: pulumi.Input<inputs.glue.CatalogTableTargetTable>;
    /**
     * If the table is a view, the expanded text of the view; otherwise null.
     */
    viewExpandedText?: pulumi.Input<string>;
    /**
     * If the table is a view, the original text of the view; otherwise null.
     */
    viewOriginalText?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CatalogTable resource.
 */
export interface CatalogTableArgs {
    /**
     * ID of the Data Catalog in which the table resides.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Name of the catalog database that contains the target table.
     */
    databaseName: pulumi.Input<string>;
    /**
     * Description of the table.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the target table.
     */
    name?: pulumi.Input<string>;
    /**
     * Owner of the table.
     */
    owner?: pulumi.Input<string>;
    /**
     * Map of initialization parameters for the SerDe, in key-value form.
     */
    parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Configuration block for a maximum of 3 partition indexes. See `partitionIndex` below.
     */
    partitionIndices?: pulumi.Input<pulumi.Input<inputs.glue.CatalogTablePartitionIndex>[]>;
    /**
     * Configuration block of columns by which the table is partitioned. Only primitive types are supported as partition keys. See `partitionKeys` below.
     */
    partitionKeys?: pulumi.Input<pulumi.Input<inputs.glue.CatalogTablePartitionKey>[]>;
    /**
     * Retention time for this table.
     */
    retention?: pulumi.Input<number>;
    /**
     * Configuration block for information about the physical storage of this table. For more information, refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables. html#aws-glue-api-catalog-tables-StorageDescriptor). See `storageDescriptor` below.
     */
    storageDescriptor?: pulumi.Input<inputs.glue.CatalogTableStorageDescriptor>;
    /**
     * Type of this table (EXTERNAL_TABLE, VIRTUAL_VIEW, etc.). While optional, some Athena DDL queries such as `ALTER TABLE` and `SHOW CREATE TABLE` will fail if this argument is empty.
     */
    tableType?: pulumi.Input<string>;
    /**
     * Configuration block of a target table for resource linking. See `targetTable` below.
     */
    targetTable?: pulumi.Input<inputs.glue.CatalogTableTargetTable>;
    /**
     * If the table is a view, the expanded text of the view; otherwise null.
     */
    viewExpandedText?: pulumi.Input<string>;
    /**
     * If the table is a view, the original text of the view; otherwise null.
     */
    viewOriginalText?: pulumi.Input<string>;
}
