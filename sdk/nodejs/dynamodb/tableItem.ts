// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DynamoDB table item resource
 * 
 * > **Note:** This resource is not meant to be used for managing large amounts of data in your table, it is not designed to scale.
 *   You should perform **regular backups** of all data in the table, see [AWS docs for more](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/BackupRestore.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const aws_dynamodb_table_example = new aws.dynamodb.Table("example", {
 *     attributes: [{
 *         name: "exampleHashKey",
 *         type: "S",
 *     }],
 *     hashKey: "exampleHashKey",
 *     name: "example-name",
 *     readCapacity: 10,
 *     writeCapacity: 10,
 * });
 * const aws_dynamodb_table_item_example = new aws.dynamodb.TableItem("example", {
 *     hashKey: aws_dynamodb_table_example.hashKey,
 *     item: "{\n  \"exampleHashKey\": {\"S\": \"something\"},\n  \"one\": {\"N\": \"11111\"},\n  \"two\": {\"N\": \"22222\"},\n  \"three\": {\"N\": \"33333\"},\n  \"four\": {\"N\": \"44444\"}\n}\n",
 *     tableName: aws_dynamodb_table_example.name,
 * });
 * ```
 */
export class TableItem extends pulumi.CustomResource {
    /**
     * Get an existing TableItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableItemState, opts?: pulumi.CustomResourceOptions): TableItem {
        return new TableItem(name, <any>state, { ...opts, id: id });
    }

    /**
     * Hash key to use for lookups and identification of the item
     */
    public readonly hashKey: pulumi.Output<string>;
    /**
     * JSON representation of a map of attribute name/value pairs, one for each attribute.
     * Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
     */
    public readonly item: pulumi.Output<string>;
    /**
     * Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
     */
    public readonly rangeKey: pulumi.Output<string | undefined>;
    /**
     * The name of the table to contain the item.
     */
    public readonly tableName: pulumi.Output<string>;

    /**
     * Create a TableItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableItemArgs | TableItemState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: TableItemState = argsOrState as TableItemState | undefined;
            inputs["hashKey"] = state ? state.hashKey : undefined;
            inputs["item"] = state ? state.item : undefined;
            inputs["rangeKey"] = state ? state.rangeKey : undefined;
            inputs["tableName"] = state ? state.tableName : undefined;
        } else {
            const args = argsOrState as TableItemArgs | undefined;
            if (!args || args.hashKey === undefined) {
                throw new Error("Missing required property 'hashKey'");
            }
            if (!args || args.item === undefined) {
                throw new Error("Missing required property 'item'");
            }
            if (!args || args.tableName === undefined) {
                throw new Error("Missing required property 'tableName'");
            }
            inputs["hashKey"] = args ? args.hashKey : undefined;
            inputs["item"] = args ? args.item : undefined;
            inputs["rangeKey"] = args ? args.rangeKey : undefined;
            inputs["tableName"] = args ? args.tableName : undefined;
        }
        super("aws:dynamodb/tableItem:TableItem", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TableItem resources.
 */
export interface TableItemState {
    /**
     * Hash key to use for lookups and identification of the item
     */
    readonly hashKey?: pulumi.Input<string>;
    /**
     * JSON representation of a map of attribute name/value pairs, one for each attribute.
     * Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
     */
    readonly item?: pulumi.Input<string>;
    /**
     * Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
     */
    readonly rangeKey?: pulumi.Input<string>;
    /**
     * The name of the table to contain the item.
     */
    readonly tableName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TableItem resource.
 */
export interface TableItemArgs {
    /**
     * Hash key to use for lookups and identification of the item
     */
    readonly hashKey: pulumi.Input<string>;
    /**
     * JSON representation of a map of attribute name/value pairs, one for each attribute.
     * Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
     */
    readonly item: pulumi.Input<string>;
    /**
     * Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
     */
    readonly rangeKey?: pulumi.Input<string>;
    /**
     * The name of the table to contain the item.
     */
    readonly tableName: pulumi.Input<string>;
}
