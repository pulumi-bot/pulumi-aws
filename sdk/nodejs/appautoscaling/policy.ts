// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appautoscaling/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly policyType!: pulumi.Output<string | undefined>;
    public readonly resourceId!: pulumi.Output<string>;
    public readonly scalableDimension!: pulumi.Output<string>;
    public readonly serviceNamespace!: pulumi.Output<string>;
    public readonly stepScalingPolicyConfiguration!: pulumi.Output<outputs.appautoscaling.PolicyStepScalingPolicyConfiguration | undefined>;
    public readonly targetTrackingScalingPolicyConfiguration!: pulumi.Output<outputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyType"] = state ? state.policyType : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["scalableDimension"] = state ? state.scalableDimension : undefined;
            inputs["serviceNamespace"] = state ? state.serviceNamespace : undefined;
            inputs["stepScalingPolicyConfiguration"] = state ? state.stepScalingPolicyConfiguration : undefined;
            inputs["targetTrackingScalingPolicyConfiguration"] = state ? state.targetTrackingScalingPolicyConfiguration : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.scalableDimension === undefined) {
                throw new Error("Missing required property 'scalableDimension'");
            }
            if (!args || args.serviceNamespace === undefined) {
                throw new Error("Missing required property 'serviceNamespace'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["scalableDimension"] = args ? args.scalableDimension : undefined;
            inputs["serviceNamespace"] = args ? args.serviceNamespace : undefined;
            inputs["stepScalingPolicyConfiguration"] = args ? args.stepScalingPolicyConfiguration : undefined;
            inputs["targetTrackingScalingPolicyConfiguration"] = args ? args.targetTrackingScalingPolicyConfiguration : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    readonly arn?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly policyType?: pulumi.Input<string>;
    readonly resourceId?: pulumi.Input<string>;
    readonly scalableDimension?: pulumi.Input<string>;
    readonly serviceNamespace?: pulumi.Input<string>;
    readonly stepScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyStepScalingPolicyConfiguration>;
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    readonly name?: pulumi.Input<string>;
    readonly policyType?: pulumi.Input<string>;
    readonly resourceId: pulumi.Input<string>;
    readonly scalableDimension: pulumi.Input<string>;
    readonly serviceNamespace: pulumi.Input<string>;
    readonly stepScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyStepScalingPolicyConfiguration>;
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration>;
}
