// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Glue Crawler. More information can be found in the [AWS Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/add-crawler.html)
 *
 * ## Example Usage
 * ### DynamoDB Target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Crawler("example", {
 *     databaseName: aws_glue_catalog_database.example.name,
 *     role: aws_iam_role.example.arn,
 *     dynamodbTargets: [{
 *         path: "table-name",
 *     }],
 * });
 * ```
 * ### JDBC Target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Crawler("example", {
 *     databaseName: aws_glue_catalog_database.example.name,
 *     role: aws_iam_role.example.arn,
 *     jdbcTargets: [{
 *         connectionName: aws_glue_connection.example.name,
 *         path: `database-name/%`,
 *     }],
 * });
 * ```
 * ### S3 Target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Crawler("example", {
 *     databaseName: aws_glue_catalog_database.example.name,
 *     role: aws_iam_role.example.arn,
 *     s3Targets: [{
 *         path: `s3://${aws_s3_bucket.example.bucket}`,
 *     }],
 * });
 * ```
 * ### Catalog Target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Crawler("example", {
 *     databaseName: aws_glue_catalog_database.example.name,
 *     role: aws_iam_role.example.arn,
 *     catalogTargets: [{
 *         databaseName: aws_glue_catalog_database.example.name,
 *         tables: [aws_glue_catalog_table.example.name],
 *     }],
 *     schemaChangePolicy: {
 *         deleteBehavior: "LOG",
 *     },
 *     configuration: `{
 *   "Version":1.0,
 *   "Grouping": {
 *     "TableGroupingPolicy": "CombineCompatibleSchemas"
 *   }
 * }
 * `,
 * });
 * ```
 * ### MongoDB Target
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Crawler("example", {
 *     databaseName: aws_glue_catalog_database.example.name,
 *     role: aws_iam_role.example.arn,
 *     mongodbTargets: [{
 *         connectionName: aws_glue_connection.example.name,
 *         path: `database-name/%`,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * Glue Crawlers can be imported using `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/crawler:Crawler MyJob MyJob
 * ```
 */
export class Crawler extends pulumi.CustomResource {
    /**
     * Get an existing Crawler resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CrawlerState, opts?: pulumi.CustomResourceOptions): Crawler {
        return new Crawler(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/crawler:Crawler';

    /**
     * Returns true if the given object is an instance of Crawler.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Crawler {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Crawler.__pulumiType;
    }

    /**
     * The ARN of the crawler
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly catalogTargets!: pulumi.Output<outputs.glue.CrawlerCatalogTarget[] | undefined>;
    /**
     * List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
     */
    public readonly classifiers!: pulumi.Output<string[] | undefined>;
    /**
     * JSON string of configuration information.
     */
    public readonly configuration!: pulumi.Output<string | undefined>;
    /**
     * Glue database where results are written.
     */
    public readonly databaseName!: pulumi.Output<string>;
    /**
     * Description of the crawler.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of nested DynamoDB target arguments. See below.
     */
    public readonly dynamodbTargets!: pulumi.Output<outputs.glue.CrawlerDynamodbTarget[] | undefined>;
    /**
     * List of nested JBDC target arguments. See below.
     */
    public readonly jdbcTargets!: pulumi.Output<outputs.glue.CrawlerJdbcTarget[] | undefined>;
    /**
     * List nested MongoDB target arguments. See below.
     */
    public readonly mongodbTargets!: pulumi.Output<outputs.glue.CrawlerMongodbTarget[] | undefined>;
    /**
     * Name of the crawler.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * List nested Amazon S3 target arguments. See below.
     */
    public readonly s3Targets!: pulumi.Output<outputs.glue.CrawlerS3Target[] | undefined>;
    /**
     * A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * Policy for the crawler's update and deletion behavior.
     */
    public readonly schemaChangePolicy!: pulumi.Output<outputs.glue.CrawlerSchemaChangePolicy | undefined>;
    /**
     * The name of Security Configuration to be used by the crawler
     */
    public readonly securityConfiguration!: pulumi.Output<string | undefined>;
    /**
     * The table prefix used for catalog tables that are created.
     */
    public readonly tablePrefix!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Crawler resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CrawlerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CrawlerArgs | CrawlerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CrawlerState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["catalogTargets"] = state ? state.catalogTargets : undefined;
            inputs["classifiers"] = state ? state.classifiers : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["databaseName"] = state ? state.databaseName : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["dynamodbTargets"] = state ? state.dynamodbTargets : undefined;
            inputs["jdbcTargets"] = state ? state.jdbcTargets : undefined;
            inputs["mongodbTargets"] = state ? state.mongodbTargets : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["s3Targets"] = state ? state.s3Targets : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["schemaChangePolicy"] = state ? state.schemaChangePolicy : undefined;
            inputs["securityConfiguration"] = state ? state.securityConfiguration : undefined;
            inputs["tablePrefix"] = state ? state.tablePrefix : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as CrawlerArgs | undefined;
            if (!args || args.databaseName === undefined) {
                throw new Error("Missing required property 'databaseName'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["catalogTargets"] = args ? args.catalogTargets : undefined;
            inputs["classifiers"] = args ? args.classifiers : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["databaseName"] = args ? args.databaseName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["dynamodbTargets"] = args ? args.dynamodbTargets : undefined;
            inputs["jdbcTargets"] = args ? args.jdbcTargets : undefined;
            inputs["mongodbTargets"] = args ? args.mongodbTargets : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["s3Targets"] = args ? args.s3Targets : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["schemaChangePolicy"] = args ? args.schemaChangePolicy : undefined;
            inputs["securityConfiguration"] = args ? args.securityConfiguration : undefined;
            inputs["tablePrefix"] = args ? args.tablePrefix : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Crawler.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Crawler resources.
 */
export interface CrawlerState {
    /**
     * The ARN of the crawler
     */
    readonly arn?: pulumi.Input<string>;
    readonly catalogTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerCatalogTarget>[]>;
    /**
     * List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
     */
    readonly classifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * JSON string of configuration information.
     */
    readonly configuration?: pulumi.Input<string>;
    /**
     * Glue database where results are written.
     */
    readonly databaseName?: pulumi.Input<string>;
    /**
     * Description of the crawler.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of nested DynamoDB target arguments. See below.
     */
    readonly dynamodbTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerDynamodbTarget>[]>;
    /**
     * List of nested JBDC target arguments. See below.
     */
    readonly jdbcTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerJdbcTarget>[]>;
    /**
     * List nested MongoDB target arguments. See below.
     */
    readonly mongodbTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerMongodbTarget>[]>;
    /**
     * Name of the crawler.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * List nested Amazon S3 target arguments. See below.
     */
    readonly s3Targets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerS3Target>[]>;
    /**
     * A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Policy for the crawler's update and deletion behavior.
     */
    readonly schemaChangePolicy?: pulumi.Input<inputs.glue.CrawlerSchemaChangePolicy>;
    /**
     * The name of Security Configuration to be used by the crawler
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * The table prefix used for catalog tables that are created.
     */
    readonly tablePrefix?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Crawler resource.
 */
export interface CrawlerArgs {
    readonly catalogTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerCatalogTarget>[]>;
    /**
     * List of custom classifiers. By default, all AWS classifiers are included in a crawl, but these custom classifiers always override the default classifiers for a given classification.
     */
    readonly classifiers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * JSON string of configuration information.
     */
    readonly configuration?: pulumi.Input<string>;
    /**
     * Glue database where results are written.
     */
    readonly databaseName: pulumi.Input<string>;
    /**
     * Description of the crawler.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of nested DynamoDB target arguments. See below.
     */
    readonly dynamodbTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerDynamodbTarget>[]>;
    /**
     * List of nested JBDC target arguments. See below.
     */
    readonly jdbcTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerJdbcTarget>[]>;
    /**
     * List nested MongoDB target arguments. See below.
     */
    readonly mongodbTargets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerMongodbTarget>[]>;
    /**
     * Name of the crawler.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The IAM role friendly name (including path without leading slash), or ARN of an IAM role, used by the crawler to access other resources.
     */
    readonly role: pulumi.Input<string>;
    /**
     * List nested Amazon S3 target arguments. See below.
     */
    readonly s3Targets?: pulumi.Input<pulumi.Input<inputs.glue.CrawlerS3Target>[]>;
    /**
     * A cron expression used to specify the schedule. For more information, see [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html). For example, to run something every day at 12:15 UTC, you would specify: `cron(15 12 * * ? *)`.
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Policy for the crawler's update and deletion behavior.
     */
    readonly schemaChangePolicy?: pulumi.Input<inputs.glue.CrawlerSchemaChangePolicy>;
    /**
     * The name of Security Configuration to be used by the crawler
     */
    readonly securityConfiguration?: pulumi.Input<string>;
    /**
     * The table prefix used for catalog tables that are created.
     */
    readonly tablePrefix?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
