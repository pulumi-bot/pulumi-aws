// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {InstancePlatform} from "./instancePlatform";
import {InstanceType} from "./instanceType";
import {Tenancy} from "./tenancy";

export class CapacityReservation extends pulumi.CustomResource {
    /**
     * Get an existing CapacityReservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CapacityReservationState, opts?: pulumi.CustomResourceOptions): CapacityReservation {
        return new CapacityReservation(name, <any>state, { ...opts, id: id });
    }

    public readonly availabilityZone: pulumi.Output<string>;
    public readonly ebsOptimized: pulumi.Output<boolean | undefined>;
    public readonly endDate: pulumi.Output<string | undefined>;
    public readonly endDateType: pulumi.Output<string | undefined>;
    public readonly ephemeralStorage: pulumi.Output<boolean | undefined>;
    public readonly instanceCount: pulumi.Output<number>;
    public readonly instanceMatchCriteria: pulumi.Output<string | undefined>;
    public readonly instancePlatform: pulumi.Output<InstancePlatform>;
    public readonly instanceType: pulumi.Output<InstanceType>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly tenancy: pulumi.Output<Tenancy | undefined>;

    /**
     * Create a CapacityReservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CapacityReservationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CapacityReservationArgs | CapacityReservationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: CapacityReservationState = argsOrState as CapacityReservationState | undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["ebsOptimized"] = state ? state.ebsOptimized : undefined;
            inputs["endDate"] = state ? state.endDate : undefined;
            inputs["endDateType"] = state ? state.endDateType : undefined;
            inputs["ephemeralStorage"] = state ? state.ephemeralStorage : undefined;
            inputs["instanceCount"] = state ? state.instanceCount : undefined;
            inputs["instanceMatchCriteria"] = state ? state.instanceMatchCriteria : undefined;
            inputs["instancePlatform"] = state ? state.instancePlatform : undefined;
            inputs["instanceType"] = state ? state.instanceType : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tenancy"] = state ? state.tenancy : undefined;
        } else {
            const args = argsOrState as CapacityReservationArgs | undefined;
            if (!args || args.availabilityZone === undefined) {
                throw new Error("Missing required property 'availabilityZone'");
            }
            if (!args || args.instanceCount === undefined) {
                throw new Error("Missing required property 'instanceCount'");
            }
            if (!args || args.instancePlatform === undefined) {
                throw new Error("Missing required property 'instancePlatform'");
            }
            if (!args || args.instanceType === undefined) {
                throw new Error("Missing required property 'instanceType'");
            }
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["ebsOptimized"] = args ? args.ebsOptimized : undefined;
            inputs["endDate"] = args ? args.endDate : undefined;
            inputs["endDateType"] = args ? args.endDateType : undefined;
            inputs["ephemeralStorage"] = args ? args.ephemeralStorage : undefined;
            inputs["instanceCount"] = args ? args.instanceCount : undefined;
            inputs["instanceMatchCriteria"] = args ? args.instanceMatchCriteria : undefined;
            inputs["instancePlatform"] = args ? args.instancePlatform : undefined;
            inputs["instanceType"] = args ? args.instanceType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tenancy"] = args ? args.tenancy : undefined;
        }
        super("aws:ec2/capacityReservation:CapacityReservation", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CapacityReservation resources.
 */
export interface CapacityReservationState {
    readonly availabilityZone?: pulumi.Input<string>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly endDate?: pulumi.Input<string>;
    readonly endDateType?: pulumi.Input<string>;
    readonly ephemeralStorage?: pulumi.Input<boolean>;
    readonly instanceCount?: pulumi.Input<number>;
    readonly instanceMatchCriteria?: pulumi.Input<string>;
    readonly instancePlatform?: pulumi.Input<InstancePlatform>;
    readonly instanceType?: pulumi.Input<InstanceType>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly tenancy?: pulumi.Input<Tenancy>;
}

/**
 * The set of arguments for constructing a CapacityReservation resource.
 */
export interface CapacityReservationArgs {
    readonly availabilityZone: pulumi.Input<string>;
    readonly ebsOptimized?: pulumi.Input<boolean>;
    readonly endDate?: pulumi.Input<string>;
    readonly endDateType?: pulumi.Input<string>;
    readonly ephemeralStorage?: pulumi.Input<boolean>;
    readonly instanceCount: pulumi.Input<number>;
    readonly instanceMatchCriteria?: pulumi.Input<string>;
    readonly instancePlatform: pulumi.Input<InstancePlatform>;
    readonly instanceType: pulumi.Input<InstanceType>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly tenancy?: pulumi.Input<Tenancy>;
}
