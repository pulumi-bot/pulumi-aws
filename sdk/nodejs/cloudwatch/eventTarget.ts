// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class EventTarget extends pulumi.CustomResource {
    /**
     * Get an existing EventTarget resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventTargetState, opts?: pulumi.CustomResourceOptions): EventTarget {
        return new EventTarget(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cloudwatch/eventTarget:EventTarget';

    /**
     * Returns true if the given object is an instance of EventTarget.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventTarget {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventTarget.__pulumiType;
    }

    public readonly arn!: pulumi.Output<string>;
    public readonly batchTarget!: pulumi.Output<outputs.cloudwatch.EventTargetBatchTarget | undefined>;
    public readonly ecsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetEcsTarget | undefined>;
    public readonly input!: pulumi.Output<string | undefined>;
    public readonly inputPath!: pulumi.Output<string | undefined>;
    public readonly inputTransformer!: pulumi.Output<outputs.cloudwatch.EventTargetInputTransformer | undefined>;
    public readonly kinesisTarget!: pulumi.Output<outputs.cloudwatch.EventTargetKinesisTarget | undefined>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly rule!: pulumi.Output<string>;
    public readonly runCommandTargets!: pulumi.Output<outputs.cloudwatch.EventTargetRunCommandTarget[] | undefined>;
    public readonly sqsTarget!: pulumi.Output<outputs.cloudwatch.EventTargetSqsTarget | undefined>;
    public readonly targetId!: pulumi.Output<string>;

    /**
     * Create a EventTarget resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventTargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventTargetArgs | EventTargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EventTargetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["batchTarget"] = state ? state.batchTarget : undefined;
            inputs["ecsTarget"] = state ? state.ecsTarget : undefined;
            inputs["input"] = state ? state.input : undefined;
            inputs["inputPath"] = state ? state.inputPath : undefined;
            inputs["inputTransformer"] = state ? state.inputTransformer : undefined;
            inputs["kinesisTarget"] = state ? state.kinesisTarget : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["rule"] = state ? state.rule : undefined;
            inputs["runCommandTargets"] = state ? state.runCommandTargets : undefined;
            inputs["sqsTarget"] = state ? state.sqsTarget : undefined;
            inputs["targetId"] = state ? state.targetId : undefined;
        } else {
            const args = argsOrState as EventTargetArgs | undefined;
            if (!args || args.arn === undefined) {
                throw new Error("Missing required property 'arn'");
            }
            if (!args || args.rule === undefined) {
                throw new Error("Missing required property 'rule'");
            }
            inputs["arn"] = args ? args.arn : undefined;
            inputs["batchTarget"] = args ? args.batchTarget : undefined;
            inputs["ecsTarget"] = args ? args.ecsTarget : undefined;
            inputs["input"] = args ? args.input : undefined;
            inputs["inputPath"] = args ? args.inputPath : undefined;
            inputs["inputTransformer"] = args ? args.inputTransformer : undefined;
            inputs["kinesisTarget"] = args ? args.kinesisTarget : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["rule"] = args ? args.rule : undefined;
            inputs["runCommandTargets"] = args ? args.runCommandTargets : undefined;
            inputs["sqsTarget"] = args ? args.sqsTarget : undefined;
            inputs["targetId"] = args ? args.targetId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EventTarget.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventTarget resources.
 */
export interface EventTargetState {
    readonly arn?: pulumi.Input<string>;
    readonly batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    readonly ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    readonly input?: pulumi.Input<string>;
    readonly inputPath?: pulumi.Input<string>;
    readonly inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    readonly kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    readonly roleArn?: pulumi.Input<string>;
    readonly rule?: pulumi.Input<string>;
    readonly runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    readonly sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    readonly targetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventTarget resource.
 */
export interface EventTargetArgs {
    readonly arn: pulumi.Input<string>;
    readonly batchTarget?: pulumi.Input<inputs.cloudwatch.EventTargetBatchTarget>;
    readonly ecsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetEcsTarget>;
    readonly input?: pulumi.Input<string>;
    readonly inputPath?: pulumi.Input<string>;
    readonly inputTransformer?: pulumi.Input<inputs.cloudwatch.EventTargetInputTransformer>;
    readonly kinesisTarget?: pulumi.Input<inputs.cloudwatch.EventTargetKinesisTarget>;
    readonly roleArn?: pulumi.Input<string>;
    readonly rule: pulumi.Input<string>;
    readonly runCommandTargets?: pulumi.Input<pulumi.Input<inputs.cloudwatch.EventTargetRunCommandTarget>[]>;
    readonly sqsTarget?: pulumi.Input<inputs.cloudwatch.EventTargetSqsTarget>;
    readonly targetId?: pulumi.Input<string>;
}
