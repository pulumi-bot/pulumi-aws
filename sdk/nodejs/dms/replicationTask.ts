// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a DMS (Data Migration Service) replication task resource. DMS replication tasks can be created, updated, deleted, and imported.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Create a new replication task
 * const test = new aws.dms.ReplicationTask("test", {
 *     cdcStartTime: 1484346880,
 *     migrationType: "full-load",
 *     replicationInstanceArn: aws_dms_replication_instance["test-dms-replication-instance-tf"].replication_instance_arn,
 *     replicationTaskId: "test-dms-replication-task-tf",
 *     replicationTaskSettings: "...",
 *     sourceEndpointArn: aws_dms_endpoint["test-dms-source-endpoint-tf"].endpoint_arn,
 *     tableMappings: `{"rules":[{"rule-type":"selection","rule-id":"1","rule-name":"1","object-locator":{"schema-name":"%","table-name":"%"},"rule-action":"include"}]}`,
 *     tags: {
 *         Name: "test",
 *     },
 *     targetEndpointArn: aws_dms_endpoint["test-dms-target-endpoint-tf"].endpoint_arn,
 * });
 * ```
 *
 * ## Import
 *
 * Replication tasks can be imported using the `replication_task_id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:dms/replicationTask:ReplicationTask test test-dms-replication-task-tf
 * ```
 */
export class ReplicationTask extends pulumi.CustomResource {
    /**
     * Get an existing ReplicationTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReplicationTaskState, opts?: pulumi.CustomResourceOptions): ReplicationTask {
        return new ReplicationTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:dms/replicationTask:ReplicationTask';

    /**
     * Returns true if the given object is an instance of ReplicationTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReplicationTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReplicationTask.__pulumiType;
    }

    /**
     * The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
     */
    public readonly cdcStartTime!: pulumi.Output<string | undefined>;
    /**
     * The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
     */
    public readonly migrationType!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the replication instance.
     */
    public readonly replicationInstanceArn!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) for the replication task.
     */
    public /*out*/ readonly replicationTaskArn!: pulumi.Output<string>;
    /**
     * The replication task identifier.
     */
    public readonly replicationTaskId!: pulumi.Output<string>;
    /**
     * An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
     */
    public readonly replicationTaskSettings!: pulumi.Output<string | undefined>;
    /**
     * The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
     */
    public readonly sourceEndpointArn!: pulumi.Output<string>;
    /**
     * An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
     */
    public readonly tableMappings!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
     */
    public readonly targetEndpointArn!: pulumi.Output<string>;

    /**
     * Create a ReplicationTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReplicationTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReplicationTaskArgs | ReplicationTaskState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReplicationTaskState | undefined;
            inputs["cdcStartTime"] = state ? state.cdcStartTime : undefined;
            inputs["migrationType"] = state ? state.migrationType : undefined;
            inputs["replicationInstanceArn"] = state ? state.replicationInstanceArn : undefined;
            inputs["replicationTaskArn"] = state ? state.replicationTaskArn : undefined;
            inputs["replicationTaskId"] = state ? state.replicationTaskId : undefined;
            inputs["replicationTaskSettings"] = state ? state.replicationTaskSettings : undefined;
            inputs["sourceEndpointArn"] = state ? state.sourceEndpointArn : undefined;
            inputs["tableMappings"] = state ? state.tableMappings : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["targetEndpointArn"] = state ? state.targetEndpointArn : undefined;
        } else {
            const args = argsOrState as ReplicationTaskArgs | undefined;
            if ((!args || args.migrationType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'migrationType'");
            }
            if ((!args || args.replicationInstanceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicationInstanceArn'");
            }
            if ((!args || args.replicationTaskId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'replicationTaskId'");
            }
            if ((!args || args.sourceEndpointArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceEndpointArn'");
            }
            if ((!args || args.tableMappings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tableMappings'");
            }
            if ((!args || args.targetEndpointArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'targetEndpointArn'");
            }
            inputs["cdcStartTime"] = args ? args.cdcStartTime : undefined;
            inputs["migrationType"] = args ? args.migrationType : undefined;
            inputs["replicationInstanceArn"] = args ? args.replicationInstanceArn : undefined;
            inputs["replicationTaskId"] = args ? args.replicationTaskId : undefined;
            inputs["replicationTaskSettings"] = args ? args.replicationTaskSettings : undefined;
            inputs["sourceEndpointArn"] = args ? args.sourceEndpointArn : undefined;
            inputs["tableMappings"] = args ? args.tableMappings : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["targetEndpointArn"] = args ? args.targetEndpointArn : undefined;
            inputs["replicationTaskArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ReplicationTask.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ReplicationTask resources.
 */
export interface ReplicationTaskState {
    /**
     * The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
     */
    cdcStartTime?: pulumi.Input<string>;
    /**
     * The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
     */
    migrationType?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the replication instance.
     */
    replicationInstanceArn?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) for the replication task.
     */
    replicationTaskArn?: pulumi.Input<string>;
    /**
     * The replication task identifier.
     */
    replicationTaskId?: pulumi.Input<string>;
    /**
     * An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
     */
    replicationTaskSettings?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
     */
    sourceEndpointArn?: pulumi.Input<string>;
    /**
     * An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
     */
    tableMappings?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
     */
    targetEndpointArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ReplicationTask resource.
 */
export interface ReplicationTaskArgs {
    /**
     * The Unix timestamp integer for the start of the Change Data Capture (CDC) operation.
     */
    cdcStartTime?: pulumi.Input<string>;
    /**
     * The migration type. Can be one of `full-load | cdc | full-load-and-cdc`.
     */
    migrationType: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the replication instance.
     */
    replicationInstanceArn: pulumi.Input<string>;
    /**
     * The replication task identifier.
     */
    replicationTaskId: pulumi.Input<string>;
    /**
     * An escaped JSON string that contains the task settings. For a complete list of task settings, see [Task Settings for AWS Database Migration Service Tasks](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TaskSettings.html).
     */
    replicationTaskSettings?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) string that uniquely identifies the source endpoint.
     */
    sourceEndpointArn: pulumi.Input<string>;
    /**
     * An escaped JSON string that contains the table mappings. For information on table mapping see [Using Table Mapping with an AWS Database Migration Service Task to Select and Filter Data](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.CustomizingTasks.TableMapping.html)
     */
    tableMappings: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The Amazon Resource Name (ARN) string that uniquely identifies the target endpoint.
     */
    targetEndpointArn: pulumi.Input<string>;
}
