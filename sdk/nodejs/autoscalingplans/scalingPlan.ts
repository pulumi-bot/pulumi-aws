// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an AWS Auto Scaling scaling plan.
 * More information can be found in the [AWS Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/plans/userguide/what-is-aws-auto-scaling.html).
 *
 * > **NOTE:** The AWS Auto Scaling service uses an AWS IAM service-linked role to manage predictive scaling of Amazon EC2 Auto Scaling groups. The service attempts to automatically create this role the first time a scaling plan with predictive scaling enabled is created.
 * An [`aws.iam.ServiceLinkedRole`](https://www.terraform.io/docs/providers/aws/r/iam_service_linked_role.html) resource can be used to manually manage this role.
 * See the [AWS documentation](https://docs.aws.amazon.com/autoscaling/plans/userguide/aws-auto-scaling-service-linked-roles.html#create-service-linked-role-manual) for more details.
 *
 * ## Example Usage
 */
export class ScalingPlan extends pulumi.CustomResource {
    /**
     * Get an existing ScalingPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScalingPlanState, opts?: pulumi.CustomResourceOptions): ScalingPlan {
        return new ScalingPlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:autoscalingplans/scalingPlan:ScalingPlan';

    /**
     * Returns true if the given object is an instance of ScalingPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScalingPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScalingPlan.__pulumiType;
    }

    /**
     * A CloudFormation stack or set of tags. You can create one scaling plan per application source.
     */
    public readonly applicationSource!: pulumi.Output<outputs.autoscalingplans.ScalingPlanApplicationSource>;
    /**
     * The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
     */
    public readonly scalingInstructions!: pulumi.Output<outputs.autoscalingplans.ScalingPlanScalingInstruction[]>;
    /**
     * The version number of the scaling plan. This value is always 1.
     */
    public /*out*/ readonly scalingPlanVersion!: pulumi.Output<number>;

    /**
     * Create a ScalingPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScalingPlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScalingPlanArgs | ScalingPlanState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScalingPlanState | undefined;
            inputs["applicationSource"] = state ? state.applicationSource : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["scalingInstructions"] = state ? state.scalingInstructions : undefined;
            inputs["scalingPlanVersion"] = state ? state.scalingPlanVersion : undefined;
        } else {
            const args = argsOrState as ScalingPlanArgs | undefined;
            if (!args || args.applicationSource === undefined) {
                throw new Error("Missing required property 'applicationSource'");
            }
            if (!args || args.scalingInstructions === undefined) {
                throw new Error("Missing required property 'scalingInstructions'");
            }
            inputs["applicationSource"] = args ? args.applicationSource : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["scalingInstructions"] = args ? args.scalingInstructions : undefined;
            inputs["scalingPlanVersion"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ScalingPlan.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScalingPlan resources.
 */
export interface ScalingPlanState {
    /**
     * A CloudFormation stack or set of tags. You can create one scaling plan per application source.
     */
    readonly applicationSource?: pulumi.Input<inputs.autoscalingplans.ScalingPlanApplicationSource>;
    /**
     * The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
     */
    readonly scalingInstructions?: pulumi.Input<pulumi.Input<inputs.autoscalingplans.ScalingPlanScalingInstruction>[]>;
    /**
     * The version number of the scaling plan. This value is always 1.
     */
    readonly scalingPlanVersion?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a ScalingPlan resource.
 */
export interface ScalingPlanArgs {
    /**
     * A CloudFormation stack or set of tags. You can create one scaling plan per application source.
     */
    readonly applicationSource: pulumi.Input<inputs.autoscalingplans.ScalingPlanApplicationSource>;
    /**
     * The name of the scaling plan. Names cannot contain vertical bars, colons, or forward slashes.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The scaling instructions. More details can be found in the [AWS Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/plans/APIReference/API_ScalingInstruction.html).
     */
    readonly scalingInstructions: pulumi.Input<pulumi.Input<inputs.autoscalingplans.ScalingPlanScalingInstruction>[]>;
}
