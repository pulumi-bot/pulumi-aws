// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export class InstanceFleet extends pulumi.CustomResource {
    /**
     * Get an existing InstanceFleet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceFleetState, opts?: pulumi.CustomResourceOptions): InstanceFleet {
        return new InstanceFleet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:emr/instanceFleet:InstanceFleet';

    /**
     * Returns true if the given object is an instance of InstanceFleet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceFleet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceFleet.__pulumiType;
    }

    /**
     * ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Configuration block for instance fleet
     */
    public readonly instanceTypeConfigs!: pulumi.Output<outputs.emr.InstanceFleetInstanceTypeConfig[] | undefined>;
    /**
     * Configuration block for launch specification
     */
    public readonly launchSpecifications!: pulumi.Output<outputs.emr.InstanceFleetLaunchSpecifications | undefined>;
    /**
     * Friendly name given to the instance fleet.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly provisionedOnDemandCapacity!: pulumi.Output<number>;
    public /*out*/ readonly provisionedSpotCapacity!: pulumi.Output<number>;
    /**
     * The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
     */
    public readonly targetOnDemandCapacity!: pulumi.Output<number | undefined>;
    /**
     * The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
     */
    public readonly targetSpotCapacity!: pulumi.Output<number | undefined>;

    /**
     * Create a InstanceFleet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceFleetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceFleetArgs | InstanceFleetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InstanceFleetState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["instanceTypeConfigs"] = state ? state.instanceTypeConfigs : undefined;
            inputs["launchSpecifications"] = state ? state.launchSpecifications : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["provisionedOnDemandCapacity"] = state ? state.provisionedOnDemandCapacity : undefined;
            inputs["provisionedSpotCapacity"] = state ? state.provisionedSpotCapacity : undefined;
            inputs["targetOnDemandCapacity"] = state ? state.targetOnDemandCapacity : undefined;
            inputs["targetSpotCapacity"] = state ? state.targetSpotCapacity : undefined;
        } else {
            const args = argsOrState as InstanceFleetArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["instanceTypeConfigs"] = args ? args.instanceTypeConfigs : undefined;
            inputs["launchSpecifications"] = args ? args.launchSpecifications : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["targetOnDemandCapacity"] = args ? args.targetOnDemandCapacity : undefined;
            inputs["targetSpotCapacity"] = args ? args.targetSpotCapacity : undefined;
            inputs["provisionedOnDemandCapacity"] = undefined /*out*/;
            inputs["provisionedSpotCapacity"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(InstanceFleet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceFleet resources.
 */
export interface InstanceFleetState {
    /**
     * ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * Configuration block for instance fleet
     */
    readonly instanceTypeConfigs?: pulumi.Input<pulumi.Input<inputs.emr.InstanceFleetInstanceTypeConfig>[]>;
    /**
     * Configuration block for launch specification
     */
    readonly launchSpecifications?: pulumi.Input<inputs.emr.InstanceFleetLaunchSpecifications>;
    /**
     * Friendly name given to the instance fleet.
     */
    readonly name?: pulumi.Input<string>;
    readonly provisionedOnDemandCapacity?: pulumi.Input<number>;
    readonly provisionedSpotCapacity?: pulumi.Input<number>;
    /**
     * The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
     */
    readonly targetOnDemandCapacity?: pulumi.Input<number>;
    /**
     * The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
     */
    readonly targetSpotCapacity?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a InstanceFleet resource.
 */
export interface InstanceFleetArgs {
    /**
     * ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * Configuration block for instance fleet
     */
    readonly instanceTypeConfigs?: pulumi.Input<pulumi.Input<inputs.emr.InstanceFleetInstanceTypeConfig>[]>;
    /**
     * Configuration block for launch specification
     */
    readonly launchSpecifications?: pulumi.Input<inputs.emr.InstanceFleetLaunchSpecifications>;
    /**
     * Friendly name given to the instance fleet.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
     */
    readonly targetOnDemandCapacity?: pulumi.Input<number>;
    /**
     * The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
     */
    readonly targetSpotCapacity?: pulumi.Input<number>;
}
