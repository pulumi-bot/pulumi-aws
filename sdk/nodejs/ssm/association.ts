// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Associates an SSM Document to an instance or EC2 tag.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ssm.Association("example", {targets: [{
 *     key: "InstanceIds",
 *     values: [aws_instance.example.id],
 * }]});
 * ```
 */
export class Association extends pulumi.CustomResource {
    /**
     * Get an existing Association resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AssociationState, opts?: pulumi.CustomResourceOptions): Association {
        return new Association(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ssm/association:Association';

    /**
     * Returns true if the given object is an instance of Association.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Association {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Association.__pulumiType;
    }

    /**
     * The ID of the SSM association.
     */
    public /*out*/ readonly associationId!: pulumi.Output<string>;
    /**
     * The descriptive name for the association.
     */
    public readonly associationName!: pulumi.Output<string | undefined>;
    /**
     * Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
     */
    public readonly automationTargetParameterName!: pulumi.Output<string | undefined>;
    /**
     * The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
     */
    public readonly complianceSeverity!: pulumi.Output<string | undefined>;
    /**
     * The document version you want to associate with the target(s). Can be a specific version or the default version.
     */
    public readonly documentVersion!: pulumi.Output<string>;
    /**
     * The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
     */
    public readonly instanceId!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
     */
    public readonly maxConcurrency!: pulumi.Output<string | undefined>;
    /**
     * The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
     */
    public readonly maxErrors!: pulumi.Output<string | undefined>;
    /**
     * The name of the SSM document to apply.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An output location block. Output Location is documented below.
     */
    public readonly outputLocation!: pulumi.Output<outputs.ssm.AssociationOutputLocation | undefined>;
    /**
     * A block of arbitrary string parameters to pass to the SSM document.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: string}>;
    /**
     * A cron expression when the association will be applied to the target(s).
     */
    public readonly scheduleExpression!: pulumi.Output<string | undefined>;
    /**
     * A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
     */
    public readonly targets!: pulumi.Output<outputs.ssm.AssociationTarget[]>;

    /**
     * Create a Association resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssociationArgs | AssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AssociationState | undefined;
            inputs["associationId"] = state ? state.associationId : undefined;
            inputs["associationName"] = state ? state.associationName : undefined;
            inputs["automationTargetParameterName"] = state ? state.automationTargetParameterName : undefined;
            inputs["complianceSeverity"] = state ? state.complianceSeverity : undefined;
            inputs["documentVersion"] = state ? state.documentVersion : undefined;
            inputs["instanceId"] = state ? state.instanceId : undefined;
            inputs["maxConcurrency"] = state ? state.maxConcurrency : undefined;
            inputs["maxErrors"] = state ? state.maxErrors : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["outputLocation"] = state ? state.outputLocation : undefined;
            inputs["parameters"] = state ? state.parameters : undefined;
            inputs["scheduleExpression"] = state ? state.scheduleExpression : undefined;
            inputs["targets"] = state ? state.targets : undefined;
        } else {
            const args = argsOrState as AssociationArgs | undefined;
            inputs["associationName"] = args ? args.associationName : undefined;
            inputs["automationTargetParameterName"] = args ? args.automationTargetParameterName : undefined;
            inputs["complianceSeverity"] = args ? args.complianceSeverity : undefined;
            inputs["documentVersion"] = args ? args.documentVersion : undefined;
            inputs["instanceId"] = args ? args.instanceId : undefined;
            inputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            inputs["maxErrors"] = args ? args.maxErrors : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["outputLocation"] = args ? args.outputLocation : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["scheduleExpression"] = args ? args.scheduleExpression : undefined;
            inputs["targets"] = args ? args.targets : undefined;
            inputs["associationId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Association.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Association resources.
 */
export interface AssociationState {
    /**
     * The ID of the SSM association.
     */
    readonly associationId?: pulumi.Input<string>;
    /**
     * The descriptive name for the association.
     */
    readonly associationName?: pulumi.Input<string>;
    /**
     * Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
     */
    readonly automationTargetParameterName?: pulumi.Input<string>;
    /**
     * The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
     */
    readonly complianceSeverity?: pulumi.Input<string>;
    /**
     * The document version you want to associate with the target(s). Can be a specific version or the default version.
     */
    readonly documentVersion?: pulumi.Input<string>;
    /**
     * The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
     */
    readonly maxConcurrency?: pulumi.Input<string>;
    /**
     * The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
     */
    readonly maxErrors?: pulumi.Input<string>;
    /**
     * The name of the SSM document to apply.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An output location block. Output Location is documented below.
     */
    readonly outputLocation?: pulumi.Input<inputs.ssm.AssociationOutputLocation>;
    /**
     * A block of arbitrary string parameters to pass to the SSM document.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A cron expression when the association will be applied to the target(s).
     */
    readonly scheduleExpression?: pulumi.Input<string>;
    /**
     * A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
     */
    readonly targets?: pulumi.Input<pulumi.Input<inputs.ssm.AssociationTarget>[]>;
}

/**
 * The set of arguments for constructing a Association resource.
 */
export interface AssociationArgs {
    /**
     * The descriptive name for the association.
     */
    readonly associationName?: pulumi.Input<string>;
    /**
     * Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
     */
    readonly automationTargetParameterName?: pulumi.Input<string>;
    /**
     * The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
     */
    readonly complianceSeverity?: pulumi.Input<string>;
    /**
     * The document version you want to associate with the target(s). Can be a specific version or the default version.
     */
    readonly documentVersion?: pulumi.Input<string>;
    /**
     * The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
     */
    readonly instanceId?: pulumi.Input<string>;
    /**
     * The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
     */
    readonly maxConcurrency?: pulumi.Input<string>;
    /**
     * The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
     */
    readonly maxErrors?: pulumi.Input<string>;
    /**
     * The name of the SSM document to apply.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * An output location block. Output Location is documented below.
     */
    readonly outputLocation?: pulumi.Input<inputs.ssm.AssociationOutputLocation>;
    /**
     * A block of arbitrary string parameters to pass to the SSM document.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A cron expression when the association will be applied to the target(s).
     */
    readonly scheduleExpression?: pulumi.Input<string>;
    /**
     * A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
     */
    readonly targets?: pulumi.Input<pulumi.Input<inputs.ssm.AssociationTarget>[]>;
}
