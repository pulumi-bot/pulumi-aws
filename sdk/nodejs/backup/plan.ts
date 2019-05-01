// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an AWS Backup plan resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.backup.Plan("example", {
 *     rules: [{
 *         ruleName: "tf_example_backup_rule",
 *         schedule: "cron(0 12 * * ? *)",
 *         targetVaultName: aws_backup_vault_test.name,
 *     }],
 * });
 * ```
 */
export class Plan extends pulumi.CustomResource {
    /**
     * Get an existing Plan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PlanState, opts?: pulumi.CustomResourceOptions): Plan {
        return new Plan(name, <any>state, { ...opts, id: id });
    }

    /**
     * The ARN of the backup plan.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The display name of a backup plan.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A rule object that specifies a scheduled task that is used to back up a selection of resources.
     */
    public readonly rules!: pulumi.Output<{ completionWindow?: number, lifecycle?: { coldStorageAfter?: number, deleteAfter?: number }, recoveryPointTags?: {[key: string]: string}, ruleName: string, schedule?: string, startWindow?: number, targetVaultName: string }[]>;
    /**
     * Metadata that you can assign to help organize the plans you create.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Plan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PlanArgs | PlanState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PlanState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as PlanArgs | undefined;
            if (!args || args.rules === undefined) {
                throw new Error("Missing required property 'rules'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("aws:backup/plan:Plan", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Plan resources.
 */
export interface PlanState {
    /**
     * The ARN of the backup plan.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The display name of a backup plan.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A rule object that specifies a scheduled task that is used to back up a selection of resources.
     */
    readonly rules?: pulumi.Input<pulumi.Input<{ completionWindow?: pulumi.Input<number>, lifecycle?: pulumi.Input<{ coldStorageAfter?: pulumi.Input<number>, deleteAfter?: pulumi.Input<number> }>, recoveryPointTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, ruleName: pulumi.Input<string>, schedule?: pulumi.Input<string>, startWindow?: pulumi.Input<number>, targetVaultName: pulumi.Input<string> }>[]>;
    /**
     * Metadata that you can assign to help organize the plans you create.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Plan resource.
 */
export interface PlanArgs {
    /**
     * The display name of a backup plan.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A rule object that specifies a scheduled task that is used to back up a selection of resources.
     */
    readonly rules: pulumi.Input<pulumi.Input<{ completionWindow?: pulumi.Input<number>, lifecycle?: pulumi.Input<{ coldStorageAfter?: pulumi.Input<number>, deleteAfter?: pulumi.Input<number> }>, recoveryPointTags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, ruleName: pulumi.Input<string>, schedule?: pulumi.Input<string>, startWindow?: pulumi.Input<number>, targetVaultName: pulumi.Input<string> }>[]>;
    /**
     * Metadata that you can assign to help organize the plans you create.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
