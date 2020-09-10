// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

import {LaunchConfiguration, PlacementGroup} from "../ec2";
import {Metric, MetricsGranularity} from "./index";

export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:autoscaling/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly availabilityZones!: pulumi.Output<string[]>;
    public readonly defaultCooldown!: pulumi.Output<number>;
    public readonly desiredCapacity!: pulumi.Output<number>;
    public readonly enabledMetrics!: pulumi.Output<Metric[] | undefined>;
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    public readonly healthCheckGracePeriod!: pulumi.Output<number | undefined>;
    public readonly healthCheckType!: pulumi.Output<string>;
    public readonly initialLifecycleHooks!: pulumi.Output<outputs.autoscaling.GroupInitialLifecycleHook[] | undefined>;
    public readonly launchConfiguration!: pulumi.Output<string | undefined>;
    public readonly launchTemplate!: pulumi.Output<outputs.autoscaling.GroupLaunchTemplate | undefined>;
    public readonly loadBalancers!: pulumi.Output<string[] | undefined>;
    public readonly maxInstanceLifetime!: pulumi.Output<number | undefined>;
    public readonly maxSize!: pulumi.Output<number>;
    public readonly metricsGranularity!: pulumi.Output<string | undefined>;
    public readonly minElbCapacity!: pulumi.Output<number | undefined>;
    public readonly minSize!: pulumi.Output<number>;
    public readonly mixedInstancesPolicy!: pulumi.Output<outputs.autoscaling.GroupMixedInstancesPolicy | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public readonly placementGroup!: pulumi.Output<string | undefined>;
    public readonly protectFromScaleIn!: pulumi.Output<boolean | undefined>;
    public readonly serviceLinkedRoleArn!: pulumi.Output<string>;
    public readonly suspendedProcesses!: pulumi.Output<string[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.autoscaling.GroupTag[] | undefined>;
    public readonly tagsCollection!: pulumi.Output<{[key: string]: string}[] | undefined>;
    public readonly targetGroupArns!: pulumi.Output<string[] | undefined>;
    public readonly terminationPolicies!: pulumi.Output<string[] | undefined>;
    public readonly vpcZoneIdentifiers!: pulumi.Output<string[]>;
    public readonly waitForCapacityTimeout!: pulumi.Output<string | undefined>;
    public readonly waitForElbCapacity!: pulumi.Output<number | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["availabilityZones"] = state ? state.availabilityZones : undefined;
            inputs["defaultCooldown"] = state ? state.defaultCooldown : undefined;
            inputs["desiredCapacity"] = state ? state.desiredCapacity : undefined;
            inputs["enabledMetrics"] = state ? state.enabledMetrics : undefined;
            inputs["forceDelete"] = state ? state.forceDelete : undefined;
            inputs["healthCheckGracePeriod"] = state ? state.healthCheckGracePeriod : undefined;
            inputs["healthCheckType"] = state ? state.healthCheckType : undefined;
            inputs["initialLifecycleHooks"] = state ? state.initialLifecycleHooks : undefined;
            inputs["launchConfiguration"] = state ? state.launchConfiguration : undefined;
            inputs["launchTemplate"] = state ? state.launchTemplate : undefined;
            inputs["loadBalancers"] = state ? state.loadBalancers : undefined;
            inputs["maxInstanceLifetime"] = state ? state.maxInstanceLifetime : undefined;
            inputs["maxSize"] = state ? state.maxSize : undefined;
            inputs["metricsGranularity"] = state ? state.metricsGranularity : undefined;
            inputs["minElbCapacity"] = state ? state.minElbCapacity : undefined;
            inputs["minSize"] = state ? state.minSize : undefined;
            inputs["mixedInstancesPolicy"] = state ? state.mixedInstancesPolicy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["placementGroup"] = state ? state.placementGroup : undefined;
            inputs["protectFromScaleIn"] = state ? state.protectFromScaleIn : undefined;
            inputs["serviceLinkedRoleArn"] = state ? state.serviceLinkedRoleArn : undefined;
            inputs["suspendedProcesses"] = state ? state.suspendedProcesses : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsCollection"] = state ? state.tagsCollection : undefined;
            inputs["targetGroupArns"] = state ? state.targetGroupArns : undefined;
            inputs["terminationPolicies"] = state ? state.terminationPolicies : undefined;
            inputs["vpcZoneIdentifiers"] = state ? state.vpcZoneIdentifiers : undefined;
            inputs["waitForCapacityTimeout"] = state ? state.waitForCapacityTimeout : undefined;
            inputs["waitForElbCapacity"] = state ? state.waitForElbCapacity : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if (!args || args.maxSize === undefined) {
                throw new Error("Missing required property 'maxSize'");
            }
            if (!args || args.minSize === undefined) {
                throw new Error("Missing required property 'minSize'");
            }
            inputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            inputs["defaultCooldown"] = args ? args.defaultCooldown : undefined;
            inputs["desiredCapacity"] = args ? args.desiredCapacity : undefined;
            inputs["enabledMetrics"] = args ? args.enabledMetrics : undefined;
            inputs["forceDelete"] = args ? args.forceDelete : undefined;
            inputs["healthCheckGracePeriod"] = args ? args.healthCheckGracePeriod : undefined;
            inputs["healthCheckType"] = args ? args.healthCheckType : undefined;
            inputs["initialLifecycleHooks"] = args ? args.initialLifecycleHooks : undefined;
            inputs["launchConfiguration"] = args ? args.launchConfiguration : undefined;
            inputs["launchTemplate"] = args ? args.launchTemplate : undefined;
            inputs["loadBalancers"] = args ? args.loadBalancers : undefined;
            inputs["maxInstanceLifetime"] = args ? args.maxInstanceLifetime : undefined;
            inputs["maxSize"] = args ? args.maxSize : undefined;
            inputs["metricsGranularity"] = args ? args.metricsGranularity : undefined;
            inputs["minElbCapacity"] = args ? args.minElbCapacity : undefined;
            inputs["minSize"] = args ? args.minSize : undefined;
            inputs["mixedInstancesPolicy"] = args ? args.mixedInstancesPolicy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["placementGroup"] = args ? args.placementGroup : undefined;
            inputs["protectFromScaleIn"] = args ? args.protectFromScaleIn : undefined;
            inputs["serviceLinkedRoleArn"] = args ? args.serviceLinkedRoleArn : undefined;
            inputs["suspendedProcesses"] = args ? args.suspendedProcesses : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsCollection"] = args ? args.tagsCollection : undefined;
            inputs["targetGroupArns"] = args ? args.targetGroupArns : undefined;
            inputs["terminationPolicies"] = args ? args.terminationPolicies : undefined;
            inputs["vpcZoneIdentifiers"] = args ? args.vpcZoneIdentifiers : undefined;
            inputs["waitForCapacityTimeout"] = args ? args.waitForCapacityTimeout : undefined;
            inputs["waitForElbCapacity"] = args ? args.waitForElbCapacity : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Group.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    readonly arn?: pulumi.Input<string>;
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly defaultCooldown?: pulumi.Input<number>;
    readonly desiredCapacity?: pulumi.Input<number>;
    readonly enabledMetrics?: pulumi.Input<pulumi.Input<Metric>[]>;
    readonly forceDelete?: pulumi.Input<boolean>;
    readonly healthCheckGracePeriod?: pulumi.Input<number>;
    readonly healthCheckType?: pulumi.Input<string>;
    readonly initialLifecycleHooks?: pulumi.Input<pulumi.Input<inputs.autoscaling.GroupInitialLifecycleHook>[]>;
    readonly launchConfiguration?: pulumi.Input<string | LaunchConfiguration>;
    readonly launchTemplate?: pulumi.Input<inputs.autoscaling.GroupLaunchTemplate>;
    readonly loadBalancers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly maxInstanceLifetime?: pulumi.Input<number>;
    readonly maxSize?: pulumi.Input<number>;
    readonly metricsGranularity?: pulumi.Input<string | MetricsGranularity>;
    readonly minElbCapacity?: pulumi.Input<number>;
    readonly minSize?: pulumi.Input<number>;
    readonly mixedInstancesPolicy?: pulumi.Input<inputs.autoscaling.GroupMixedInstancesPolicy>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly placementGroup?: pulumi.Input<string | PlacementGroup>;
    readonly protectFromScaleIn?: pulumi.Input<boolean>;
    readonly serviceLinkedRoleArn?: pulumi.Input<string>;
    readonly suspendedProcesses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<pulumi.Input<inputs.autoscaling.GroupTag>[]>;
    readonly tagsCollection?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    readonly targetGroupArns?: pulumi.Input<pulumi.Input<string>[]>;
    readonly terminationPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    readonly vpcZoneIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly waitForCapacityTimeout?: pulumi.Input<string>;
    readonly waitForElbCapacity?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    readonly availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    readonly defaultCooldown?: pulumi.Input<number>;
    readonly desiredCapacity?: pulumi.Input<number>;
    readonly enabledMetrics?: pulumi.Input<pulumi.Input<Metric>[]>;
    readonly forceDelete?: pulumi.Input<boolean>;
    readonly healthCheckGracePeriod?: pulumi.Input<number>;
    readonly healthCheckType?: pulumi.Input<string>;
    readonly initialLifecycleHooks?: pulumi.Input<pulumi.Input<inputs.autoscaling.GroupInitialLifecycleHook>[]>;
    readonly launchConfiguration?: pulumi.Input<string | LaunchConfiguration>;
    readonly launchTemplate?: pulumi.Input<inputs.autoscaling.GroupLaunchTemplate>;
    readonly loadBalancers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly maxInstanceLifetime?: pulumi.Input<number>;
    readonly maxSize: pulumi.Input<number>;
    readonly metricsGranularity?: pulumi.Input<string | MetricsGranularity>;
    readonly minElbCapacity?: pulumi.Input<number>;
    readonly minSize: pulumi.Input<number>;
    readonly mixedInstancesPolicy?: pulumi.Input<inputs.autoscaling.GroupMixedInstancesPolicy>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly placementGroup?: pulumi.Input<string | PlacementGroup>;
    readonly protectFromScaleIn?: pulumi.Input<boolean>;
    readonly serviceLinkedRoleArn?: pulumi.Input<string>;
    readonly suspendedProcesses?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<pulumi.Input<inputs.autoscaling.GroupTag>[]>;
    readonly tagsCollection?: pulumi.Input<pulumi.Input<{[key: string]: pulumi.Input<string>}>[]>;
    readonly targetGroupArns?: pulumi.Input<pulumi.Input<string>[]>;
    readonly terminationPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    readonly vpcZoneIdentifiers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly waitForCapacityTimeout?: pulumi.Input<string>;
    readonly waitForElbCapacity?: pulumi.Input<number>;
}
