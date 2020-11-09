// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Glue ML Transform resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testCatalogDatabase = new aws.glue.CatalogDatabase("testCatalogDatabase", {name: "example"});
 * const testCatalogTable = new aws.glue.CatalogTable("testCatalogTable", {
 *     name: "example",
 *     databaseName: testCatalogDatabase.name,
 *     owner: "my_owner",
 *     retention: 1,
 *     tableType: "VIRTUAL_VIEW",
 *     viewExpandedText: "view_expanded_text_1",
 *     viewOriginalText: "view_original_text_1",
 *     storageDescriptor: {
 *         bucketColumns: ["bucket_column_1"],
 *         compressed: false,
 *         inputFormat: "SequenceFileInputFormat",
 *         location: "my_location",
 *         numberOfBuckets: 1,
 *         outputFormat: "SequenceFileInputFormat",
 *         storedAsSubDirectories: false,
 *         parameters: {
 *             param1: "param1_val",
 *         },
 *         columns: [
 *             {
 *                 name: "my_column_1",
 *                 type: "int",
 *                 comment: "my_column1_comment",
 *             },
 *             {
 *                 name: "my_column_2",
 *                 type: "string",
 *                 comment: "my_column2_comment",
 *             },
 *         ],
 *         serDeInfo: {
 *             name: "ser_de_name",
 *             parameters: {
 *                 param1: "param_val_1",
 *             },
 *             serializationLibrary: "org.apache.hadoop.hive.serde2.columnar.ColumnarSerDe",
 *         },
 *         sortColumns: [{
 *             column: "my_column_1",
 *             sortOrder: 1,
 *         }],
 *         skewedInfo: {
 *             skewedColumnNames: ["my_column_1"],
 *             skewedColumnValueLocationMaps: {
 *                 my_column_1: "my_column_1_val_loc_map",
 *             },
 *             skewedColumnValues: ["skewed_val_1"],
 *         },
 *     },
 *     partitionKeys: [
 *         {
 *             name: "my_column_1",
 *             type: "int",
 *             comment: "my_column_1_comment",
 *         },
 *         {
 *             name: "my_column_2",
 *             type: "string",
 *             comment: "my_column_2_comment",
 *         },
 *     ],
 *     parameters: {
 *         param1: "param1_val",
 *     },
 * });
 * const testMLTransform = new aws.glue.MLTransform("testMLTransform", {
 *     roleArn: aws_iam_role.test.arn,
 *     inputRecordTables: [{
 *         databaseName: testCatalogTable.databaseName,
 *         tableName: testCatalogTable.name,
 *     }],
 *     parameters: {
 *         transformType: "FIND_MATCHES",
 *         findMatchesParameters: {
 *             primaryKeyColumnName: "my_column_1",
 *         },
 *     },
 * }, {
 *     dependsOn: [aws_iam_role_policy_attachment.test],
 * });
 * ```
 *
 * ## Import
 *
 * Glue ML Transforms can be imported using `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/mLTransform:MLTransform example tfm-c2cafbe83b1c575f49eaca9939220e2fcd58e2d5
 * ```
 */
export class MLTransform extends pulumi.CustomResource {
    /**
     * Get an existing MLTransform resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MLTransformState, opts?: pulumi.CustomResourceOptions): MLTransform {
        return new MLTransform(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/mLTransform:MLTransform';

    /**
     * Returns true if the given object is an instance of MLTransform.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MLTransform {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MLTransform.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of Glue ML Transform.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the ML Transform.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     */
    public readonly glueVersion!: pulumi.Output<string>;
    /**
     * A list of AWS Glue table definitions used by the transform. see Input Record Tables.
     */
    public readonly inputRecordTables!: pulumi.Output<outputs.glue.MLTransformInputRecordTable[]>;
    /**
     * The number of labels available for this transform.
     */
    public /*out*/ readonly labelCount!: pulumi.Output<number>;
    /**
     * The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
     */
    public readonly maxCapacity!: pulumi.Output<number>;
    /**
     * The maximum number of times to retry this ML Transform if it fails.
     */
    public readonly maxRetries!: pulumi.Output<number | undefined>;
    /**
     * The name you assign to this ML Transform. It must be unique in your account.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
     */
    public readonly numberOfWorkers!: pulumi.Output<number | undefined>;
    /**
     * The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
     */
    public readonly parameters!: pulumi.Output<outputs.glue.MLTransformParameters>;
    /**
     * The ARN of the IAM role associated with this ML Transform.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The object that represents the schema that this transform accepts. see Schema.
     */
    public /*out*/ readonly schemas!: pulumi.Output<outputs.glue.MLTransformSchema[]>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
     */
    public readonly workerType!: pulumi.Output<string | undefined>;

    /**
     * Create a MLTransform resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MLTransformArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MLTransformArgs | MLTransformState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as MLTransformState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["glueVersion"] = state ? state.glueVersion : undefined;
            inputs["inputRecordTables"] = state ? state.inputRecordTables : undefined;
            inputs["labelCount"] = state ? state.labelCount : undefined;
            inputs["maxCapacity"] = state ? state.maxCapacity : undefined;
            inputs["maxRetries"] = state ? state.maxRetries : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["numberOfWorkers"] = state ? state.numberOfWorkers : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["schemas"] = state ? state.schemas : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["workerType"] = state ? state.workerType : undefined;
        } else {
            const args = argsOrState as MLTransformArgs | undefined;
            if (!args || args.inputRecordTables === undefined) {
                throw new Error("Missing required property 'inputRecordTables'");
            }
            if (!args || args.parameters === undefined) {
                throw new Error("Missing required property 'parameters'");
            }
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["glueVersion"] = args ? args.glueVersion : undefined;
            inputs["inputRecordTables"] = args ? args.inputRecordTables : undefined;
            inputs["maxCapacity"] = args ? args.maxCapacity : undefined;
            inputs["maxRetries"] = args ? args.maxRetries : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["numberOfWorkers"] = args ? args.numberOfWorkers : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["workerType"] = args ? args.workerType : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["labelCount"] = undefined /*out*/;
            inputs["schemas"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(MLTransform.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MLTransform resources.
 */
export interface MLTransformState {
    /**
     * Amazon Resource Name (ARN) of Glue ML Transform.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Description of the ML Transform.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     */
    readonly glueVersion?: pulumi.Input<string>;
    /**
     * A list of AWS Glue table definitions used by the transform. see Input Record Tables.
     */
    readonly inputRecordTables?: pulumi.Input<pulumi.Input<inputs.glue.MLTransformInputRecordTable>[]>;
    /**
     * The number of labels available for this transform.
     */
    readonly labelCount?: pulumi.Input<number>;
    /**
     * The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
     */
    readonly maxCapacity?: pulumi.Input<number>;
    /**
     * The maximum number of times to retry this ML Transform if it fails.
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The name you assign to this ML Transform. It must be unique in your account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
     */
    readonly numberOfWorkers?: pulumi.Input<number>;
    /**
     * The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
     */
    readonly parameters?: pulumi.Input<inputs.glue.MLTransformParameters>;
    /**
     * The ARN of the IAM role associated with this ML Transform.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The object that represents the schema that this transform accepts. see Schema.
     */
    readonly schemas?: pulumi.Input<pulumi.Input<inputs.glue.MLTransformSchema>[]>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
     */
    readonly workerType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MLTransform resource.
 */
export interface MLTransformArgs {
    /**
     * Description of the ML Transform.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The version of glue to use, for example "1.0". For information about available versions, see the [AWS Glue Release Notes](https://docs.aws.amazon.com/glue/latest/dg/release-notes.html).
     */
    readonly glueVersion?: pulumi.Input<string>;
    /**
     * A list of AWS Glue table definitions used by the transform. see Input Record Tables.
     */
    readonly inputRecordTables: pulumi.Input<pulumi.Input<inputs.glue.MLTransformInputRecordTable>[]>;
    /**
     * The number of AWS Glue data processing units (DPUs) that are allocated to task runs for this transform. You can allocate from `2` to `100` DPUs; the default is `10`. `maxCapacity` is a mutually exclusive option with `numberOfWorkers` and `workerType`.
     */
    readonly maxCapacity?: pulumi.Input<number>;
    /**
     * The maximum number of times to retry this ML Transform if it fails.
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The name you assign to this ML Transform. It must be unique in your account.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of workers of a defined `workerType` that are allocated when an ML Transform runs. Required with `workerType`.
     */
    readonly numberOfWorkers?: pulumi.Input<number>;
    /**
     * The algorithmic parameters that are specific to the transform type used. Conditionally dependent on the transform type. see Parameters.
     */
    readonly parameters: pulumi.Input<inputs.glue.MLTransformParameters>;
    /**
     * The ARN of the IAM role associated with this ML Transform.
     */
    readonly roleArn: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ML Transform timeout in minutes. The default is 2880 minutes (48 hours).
     */
    readonly timeout?: pulumi.Input<number>;
    /**
     * The type of predefined worker that is allocated when an ML Transform runs. Accepts a value of `Standard`, `G.1X`, or `G.2X`. Required with `numberOfWorkers`.
     */
    readonly workerType?: pulumi.Input<string>;
}
