// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Glue Trigger resource.
 *
 * ## Example Usage
 * ### Conditional Trigger
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Trigger("example", {
 *     type: "CONDITIONAL",
 *     actions: [{
 *         jobName: aws_glue_job.example1.name,
 *     }],
 *     predicate: {
 *         conditions: [{
 *             jobName: aws_glue_job.example2.name,
 *             state: "SUCCEEDED",
 *         }],
 *     },
 * });
 * ```
 * ### On-Demand Trigger
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Trigger("example", {
 *     type: "ON_DEMAND",
 *     actions: [{
 *         jobName: aws_glue_job.example.name,
 *     }],
 * });
 * ```
 * ### Scheduled Trigger
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Trigger("example", {
 *     schedule: "cron(15 12 * * ? *)",
 *     type: "SCHEDULED",
 *     actions: [{
 *         jobName: aws_glue_job.example.name,
 *     }],
 * });
 * ```
 * ### Conditional Trigger with Crawler Action
 *
 * **Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Trigger("example", {
 *     type: "CONDITIONAL",
 *     actions: [{
 *         crawlerName: aws_glue_crawler.example1.name,
 *     }],
 *     predicate: {
 *         conditions: [{
 *             jobName: aws_glue_job.example2.name,
 *             state: "SUCCEEDED",
 *         }],
 *     },
 * });
 * ```
 * ### Conditional Trigger with Crawler Condition
 *
 * **Note:** Triggers can have both a crawler action and a crawler condition, just no example provided.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.Trigger("example", {
 *     type: "CONDITIONAL",
 *     actions: [{
 *         jobName: aws_glue_job.example1.name,
 *     }],
 *     predicate: {
 *         conditions: [{
 *             crawlerName: aws_glue_crawler.example2.name,
 *             crawlState: "SUCCEEDED",
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Glue Triggers can be imported using `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/trigger:Trigger MyTrigger MyTrigger
 * ```
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/trigger:Trigger';

    /**
     * Returns true if the given object is an instance of Trigger.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Trigger {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Trigger.__pulumiType;
    }

    /**
     * List of actions initiated by this trigger when it fires. Defined below.
     */
    public readonly actions!: pulumi.Output<outputs.glue.TriggerAction[]>;
    /**
     * Amazon Resource Name (ARN) of Glue Trigger
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A description of the new trigger.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Start the trigger. Defaults to `true`. Not valid to disable for `ON_DEMAND` type.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the trigger.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. Defined below.
     */
    public readonly predicate!: pulumi.Output<outputs.glue.TriggerPredicate | undefined>;
    /**
     * A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
     */
    public readonly workflowName!: pulumi.Output<string | undefined>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TriggerState | undefined;
            inputs["actions"] = state ? state.actions : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["predicate"] = state ? state.predicate : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["workflowName"] = state ? state.workflowName : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["predicate"] = args ? args.predicate : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["workflowName"] = args ? args.workflowName : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Trigger.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    /**
     * List of actions initiated by this trigger when it fires. Defined below.
     */
    readonly actions?: pulumi.Input<pulumi.Input<inputs.glue.TriggerAction>[]>;
    /**
     * Amazon Resource Name (ARN) of Glue Trigger
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A description of the new trigger.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Start the trigger. Defaults to `true`. Not valid to disable for `ON_DEMAND` type.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The name of the trigger.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. Defined below.
     */
    readonly predicate?: pulumi.Input<inputs.glue.TriggerPredicate>;
    /**
     * A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
     */
    readonly workflowName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * List of actions initiated by this trigger when it fires. Defined below.
     */
    readonly actions: pulumi.Input<pulumi.Input<inputs.glue.TriggerAction>[]>;
    /**
     * A description of the new trigger.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Start the trigger. Defaults to `true`. Not valid to disable for `ON_DEMAND` type.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The name of the trigger.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A predicate to specify when the new trigger should fire. Required when trigger type is `CONDITIONAL`. Defined below.
     */
    readonly predicate?: pulumi.Input<inputs.glue.TriggerPredicate>;
    /**
     * A cron expression used to specify the schedule. [Time-Based Schedules for Jobs and Crawlers](https://docs.aws.amazon.com/glue/latest/dg/monitor-data-warehouse-schedule.html)
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of trigger. Valid values are `CONDITIONAL`, `ON_DEMAND`, and `SCHEDULED`.
     */
    readonly type: pulumi.Input<string>;
    /**
     * A workflow to which the trigger should be associated to. Every workflow graph (DAG) needs a starting trigger (`ON_DEMAND` or `SCHEDULED` type) and can contain multiple additional `CONDITIONAL` triggers.
     */
    readonly workflowName?: pulumi.Input<string>;
}
