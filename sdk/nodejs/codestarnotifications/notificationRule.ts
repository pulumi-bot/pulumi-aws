// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class NotificationRule extends pulumi.CustomResource {
    /**
     * Get an existing NotificationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
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

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly detailType!: pulumi.Output<string>;
    public readonly eventTypeIds!: pulumi.Output<string[]>;
    public readonly name!: pulumi.Output<string>;
    public readonly resource!: pulumi.Output<string>;
    public readonly status!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
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
    readonly arn?: pulumi.Input<string>;
    readonly detailType?: pulumi.Input<string>;
    readonly eventTypeIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly resource?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly targets?: pulumi.Input<pulumi.Input<inputs.codestarnotifications.NotificationRuleTarget>[]>;
}

/**
 * The set of arguments for constructing a NotificationRule resource.
 */
export interface NotificationRuleArgs {
    readonly detailType: pulumi.Input<string>;
    readonly eventTypeIds: pulumi.Input<pulumi.Input<string>[]>;
    readonly name?: pulumi.Input<string>;
    readonly resource: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly targets?: pulumi.Input<pulumi.Input<inputs.codestarnotifications.NotificationRuleTarget>[]>;
}
