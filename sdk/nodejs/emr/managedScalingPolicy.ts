// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a Managed Scaling policy for EMR Cluster. With Amazon EMR versions 5.30.0 and later (except for Amazon EMR 6.0.0), you can enable EMR managed scaling to automatically increase or decrease the number of instances or units in your cluster based on workload. See [Using EMR Managed Scaling in Amazon EMR](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-managed-scaling.html) for more information.
 */
export class ManagedScalingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ManagedScalingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ManagedScalingPolicyState, opts?: pulumi.CustomResourceOptions): ManagedScalingPolicy {
        return new ManagedScalingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:emr/managedScalingPolicy:ManagedScalingPolicy';

    /**
     * Returns true if the given object is an instance of ManagedScalingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedScalingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedScalingPolicy.__pulumiType;
    }

    /**
     * The id of the EMR cluster
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Configuration block with compute limit settings. Described below.
     */
    public readonly computeLimits!: pulumi.Output<outputs.emr.ManagedScalingPolicyComputeLimit[]>;

    /**
     * Create a ManagedScalingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedScalingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ManagedScalingPolicyArgs | ManagedScalingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ManagedScalingPolicyState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["computeLimits"] = state ? state.computeLimits : undefined;
        } else {
            const args = argsOrState as ManagedScalingPolicyArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.computeLimits === undefined) {
                throw new Error("Missing required property 'computeLimits'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["computeLimits"] = args ? args.computeLimits : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ManagedScalingPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ManagedScalingPolicy resources.
 */
export interface ManagedScalingPolicyState {
    /**
     * The id of the EMR cluster
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * Configuration block with compute limit settings. Described below.
     */
    readonly computeLimits?: pulumi.Input<pulumi.Input<inputs.emr.ManagedScalingPolicyComputeLimit>[]>;
}

/**
 * The set of arguments for constructing a ManagedScalingPolicy resource.
 */
export interface ManagedScalingPolicyArgs {
    /**
     * The id of the EMR cluster
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * Configuration block with compute limit settings. Described below.
     */
    readonly computeLimits: pulumi.Input<pulumi.Input<inputs.emr.ManagedScalingPolicyComputeLimit>[]>;
}
