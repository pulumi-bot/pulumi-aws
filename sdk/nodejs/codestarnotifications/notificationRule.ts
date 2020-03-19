// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a CodeStar Notifications Rule.
 * 
 * {{% examples %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codestarnotifications_notification_rule.markdown.
 */
export class NotificationRule extends pulumi.CustomResource {
    /**
     * Get an existing NotificationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotificationRuleState, opts?: pulumi.CustomResourceOptions): NotificationRule {
        return new NotificationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codestarnotifications/notificationRule:NotificationRule';

    /**
     * Returns true if the given object is an instance of NotificationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotificationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotificationRule.__pulumiType;
    }

    /**
     * The codestar notification rule ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
     */
    public readonly detailType!: pulumi.Output<string>;
    /**
     * A list of event types associated with this notification rule.
     * For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
     */
    public readonly eventTypeIds!: pulumi.Output<string[]>;
    /**
     * The name of notification rule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN of the resource to associate with the notification rule.
     */
    public readonly resource!: pulumi.Output<string>;
    /**
     * The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
     */
    public readonly targets!: pulumi.Output<outputs.codestarnotifications.NotificationRuleTarget[] | undefined>;

    /**
     * Create a NotificationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NotificationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotificationRuleArgs | NotificationRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NotificationRuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["detailType"] = state ? state.detailType : undefined;
            inputs["eventTypeIds"] = state ? state.eventTypeIds : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resource"] = state ? state.resource : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targets"] = state ? state.targets : undefined;
        } else {
            const args = argsOrState as NotificationRuleArgs | undefined;
            if (!args || args.detailType === undefined) {
                throw new Error("Missing required property 'detailType'");
            }
            if (!args || args.eventTypeIds === undefined) {
                throw new Error("Missing required property 'eventTypeIds'");
            }
            if (!args || args.resource === undefined) {
                throw new Error("Missing required property 'resource'");
            }
            inputs["detailType"] = args ? args.detailType : undefined;
            inputs["eventTypeIds"] = args ? args.eventTypeIds : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NotificationRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotificationRule resources.
 */
export interface NotificationRuleState {
    /**
     * The codestar notification rule ARN.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
     */
    readonly detailType?: pulumi.Input<string>;
    /**
     * A list of event types associated with this notification rule.
     * For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
     */
    readonly eventTypeIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of notification rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the resource to associate with the notification rule.
     */
    readonly resource?: pulumi.Input<string>;
    /**
     * The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
     */
    readonly targets?: pulumi.Input<pulumi.Input<inputs.codestarnotifications.NotificationRuleTarget>[]>;
}

/**
 * The set of arguments for constructing a NotificationRule resource.
 */
export interface NotificationRuleArgs {
    /**
     * The level of detail to include in the notifications for this resource. Possible values are `BASIC` and `FULL`.
     */
    readonly detailType: pulumi.Input<string>;
    /**
     * A list of event types associated with this notification rule.
     * For list of allowed events see [here](https://docs.aws.amazon.com/codestar-notifications/latest/userguide/concepts.html#concepts-api).
     */
    readonly eventTypeIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of notification rule.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN of the resource to associate with the notification rule.
     */
    readonly resource: pulumi.Input<string>;
    /**
     * The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Configuration blocks containing notification target information. Can be specified multiple times. At least one target must be specified on creation.
     */
    readonly targets?: pulumi.Input<pulumi.Input<inputs.codestarnotifications.NotificationRuleTarget>[]>;
}
