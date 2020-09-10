// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class EventRule extends pulumi.CustomResource {
    /**
     * Get an existing EventRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventRuleState, opts?: pulumi.CustomResourceOptions): EventRule {
        return new EventRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventRule:EventRule';

    /**
     * Returns true if the given object is an instance of EventRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventRule.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly eventPattern!: pulumi.Output<string | undefined>;
    public readonly isEnabled!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly scheduleExpression!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a EventRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EventRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventRuleArgs | EventRuleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventRuleState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["eventPattern"] = state ? state.eventPattern : undefined;
            inputs["isEnabled"] = state ? state.isEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["scheduleExpression"] = state ? state.scheduleExpression : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as EventRuleArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["eventPattern"] = args ? args.eventPattern : undefined;
            inputs["isEnabled"] = args ? args.isEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["scheduleExpression"] = args ? args.scheduleExpression : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventRule resources.
 */
export interface EventRuleState {
    readonly arn?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly eventPattern?: pulumi.Input<string>;
    readonly isEnabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly roleArn?: pulumi.Input<string>;
    readonly scheduleExpression?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a EventRule resource.
 */
export interface EventRuleArgs {
    readonly description?: pulumi.Input<string>;
    readonly eventPattern?: pulumi.Input<string>;
    readonly isEnabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly roleArn?: pulumi.Input<string>;
    readonly scheduleExpression?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
