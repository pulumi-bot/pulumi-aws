// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Timestream table resource.
 *
 * ## Example Usage
 * ### Basic usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.timestreamwrite.Table("example", {
 *     databaseName: aws_timestreamwrite_database.example.database_name,
 *     tableName: "example",
 * });
 * ```
 * ### Full usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.timestreamwrite.Table("example", {
 *     databaseName: aws_timestreamwrite_database.example.database_name,
 *     tableName: "example",
 *     retentionProperties: {
 *         magneticStoreRetentionPeriodInDays: 30,
 *         memoryStoreRetentionPeriodInHours: 8,
 *     },
 *     tags: {
 *         Name: "example-timestream-table",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Timestream tables can be imported using the `table_name` and `database_name` separate by a colon (`:`), e.g.
 *
 * ```sh
 *  $ pulumi import aws:timestreamwrite/table:Table example ExampleTable:ExampleDatabase
 * ```
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:timestreamwrite/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * The ARN that uniquely identifies this table.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the Timestream database.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
     */
    public readonly retentionProperties!: pulumi.Output<outputs.timestreamwrite.TableRetentionProperties>;
    /**
     * The name of the Timestream table.
     */
    public readonly tableName!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["retentionProperties"] = state ? state.retentionProperties : undefined;
            inputs["tableName"] = state ? state.tableName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if ((!args || args.databaseName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'databaseName'");
            }
            if ((!args || args.tableName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableName'");
            }
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["retentionProperties"] = args ? args.retentionProperties : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Table.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * The ARN that uniquely identifies this table.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the Timestream database.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
     */
    retentionProperties?: pulumi.Input<inputs.timestreamwrite.TableRetentionProperties>;
    /**
     * The name of the Timestream table.
     */
    tableName?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * The name of the Timestream database.
     */
    databaseName: pulumi.Input<string>;
    /**
     * The retention duration for the memory store and magnetic store. See Retention Properties below for more details. If not provided, `magneticStoreRetentionPeriodInDays` default to 73000 and `memoryStoreRetentionPeriodInHours` defaults to 6.
     */
    retentionProperties?: pulumi.Input<inputs.timestreamwrite.TableRetentionProperties>;
    /**
     * The name of the Timestream table.
     */
    tableName: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block).
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
