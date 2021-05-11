// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Gamelift Fleet resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.gamelift.Fleet("example", {
 *     buildId: aws_gamelift_build.example.id,
 *     ec2InstanceType: "t2.micro",
 *     fleetType: "ON_DEMAND",
 *     runtimeConfiguration: {
 *         serverProcesses: [{
 *             concurrentExecutions: 1,
 *             launchPath: "C:\\game\\GomokuServer.exe",
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Gamelift Fleets cannot be imported at this time.
 */
export class Fleet extends pulumi.CustomResource {
    /**
     * Get an existing Fleet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FleetState, opts?: pulumi.CustomResourceOptions): Fleet {
        return new Fleet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:gamelift/fleet:Fleet';

    /**
     * Returns true if the given object is an instance of Fleet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Fleet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Fleet.__pulumiType;
    }

    /**
     * Fleet ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * ID of the Gamelift Build to be deployed on the fleet.
     */
    public readonly buildId!: pulumi.Output<string>;
    /**
     * Human-readable description of the fleet.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
     */
    public readonly ec2InboundPermissions!: pulumi.Output<outputs.gamelift.FleetEc2InboundPermission[] | undefined>;
    /**
     * Name of an EC2 instance type. e.g. `t2.micro`
     */
    public readonly ec2InstanceType!: pulumi.Output<string>;
    /**
     * Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
     */
    public readonly fleetType!: pulumi.Output<string | undefined>;
    /**
     * ARN of an IAM role that instances in the fleet can assume.
     */
    public readonly instanceRoleArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly logPaths!: pulumi.Output<string[]>;
    /**
     * List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
     */
    public readonly metricGroups!: pulumi.Output<string[]>;
    /**
     * The name of the fleet.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
     */
    public readonly newGameSessionProtectionPolicy!: pulumi.Output<string | undefined>;
    /**
     * Operating system of the fleet's computing resources.
     */
    public /*out*/ readonly operatingSystem!: pulumi.Output<string>;
    /**
     * Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
     */
    public readonly resourceCreationLimitPolicy!: pulumi.Output<outputs.gamelift.FleetResourceCreationLimitPolicy | undefined>;
    /**
     * Instructions for launching server processes on each instance in the fleet. See below.
     */
    public readonly runtimeConfiguration!: pulumi.Output<outputs.gamelift.FleetRuntimeConfiguration | undefined>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Fleet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FleetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FleetArgs | FleetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FleetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["buildId"] = state ? state.buildId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["ec2InboundPermissions"] = state ? state.ec2InboundPermissions : undefined;
            inputs["ec2InstanceType"] = state ? state.ec2InstanceType : undefined;
            inputs["fleetType"] = state ? state.fleetType : undefined;
            inputs["instanceRoleArn"] = state ? state.instanceRoleArn : undefined;
            inputs["logPaths"] = state ? state.logPaths : undefined;
            inputs["metricGroups"] = state ? state.metricGroups : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["newGameSessionProtectionPolicy"] = state ? state.newGameSessionProtectionPolicy : undefined;
            inputs["operatingSystem"] = state ? state.operatingSystem : undefined;
            inputs["resourceCreationLimitPolicy"] = state ? state.resourceCreationLimitPolicy : undefined;
            inputs["runtimeConfiguration"] = state ? state.runtimeConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as FleetArgs | undefined;
            if ((!args || args.buildId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'buildId'");
            }
            if ((!args || args.ec2InstanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ec2InstanceType'");
            }
            inputs["buildId"] = args ? args.buildId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["ec2InboundPermissions"] = args ? args.ec2InboundPermissions : undefined;
            inputs["ec2InstanceType"] = args ? args.ec2InstanceType : undefined;
            inputs["fleetType"] = args ? args.fleetType : undefined;
            inputs["instanceRoleArn"] = args ? args.instanceRoleArn : undefined;
            inputs["metricGroups"] = args ? args.metricGroups : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["newGameSessionProtectionPolicy"] = args ? args.newGameSessionProtectionPolicy : undefined;
            inputs["resourceCreationLimitPolicy"] = args ? args.resourceCreationLimitPolicy : undefined;
            inputs["runtimeConfiguration"] = args ? args.runtimeConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["logPaths"] = undefined /*out*/;
            inputs["operatingSystem"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Fleet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Fleet resources.
 */
export interface FleetState {
    /**
     * Fleet ARN.
     */
    arn?: pulumi.Input<string>;
    /**
     * ID of the Gamelift Build to be deployed on the fleet.
     */
    buildId?: pulumi.Input<string>;
    /**
     * Human-readable description of the fleet.
     */
    description?: pulumi.Input<string>;
    /**
     * Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
     */
    ec2InboundPermissions?: pulumi.Input<pulumi.Input<inputs.gamelift.FleetEc2InboundPermission>[]>;
    /**
     * Name of an EC2 instance type. e.g. `t2.micro`
     */
    ec2InstanceType?: pulumi.Input<string>;
    /**
     * Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
     */
    fleetType?: pulumi.Input<string>;
    /**
     * ARN of an IAM role that instances in the fleet can assume.
     */
    instanceRoleArn?: pulumi.Input<string>;
    logPaths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
     */
    metricGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the fleet.
     */
    name?: pulumi.Input<string>;
    /**
     * Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
     */
    newGameSessionProtectionPolicy?: pulumi.Input<string>;
    /**
     * Operating system of the fleet's computing resources.
     */
    operatingSystem?: pulumi.Input<string>;
    /**
     * Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
     */
    resourceCreationLimitPolicy?: pulumi.Input<inputs.gamelift.FleetResourceCreationLimitPolicy>;
    /**
     * Instructions for launching server processes on each instance in the fleet. See below.
     */
    runtimeConfiguration?: pulumi.Input<inputs.gamelift.FleetRuntimeConfiguration>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Fleet resource.
 */
export interface FleetArgs {
    /**
     * ID of the Gamelift Build to be deployed on the fleet.
     */
    buildId: pulumi.Input<string>;
    /**
     * Human-readable description of the fleet.
     */
    description?: pulumi.Input<string>;
    /**
     * Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
     */
    ec2InboundPermissions?: pulumi.Input<pulumi.Input<inputs.gamelift.FleetEc2InboundPermission>[]>;
    /**
     * Name of an EC2 instance type. e.g. `t2.micro`
     */
    ec2InstanceType: pulumi.Input<string>;
    /**
     * Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
     */
    fleetType?: pulumi.Input<string>;
    /**
     * ARN of an IAM role that instances in the fleet can assume.
     */
    instanceRoleArn?: pulumi.Input<string>;
    /**
     * List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
     */
    metricGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the fleet.
     */
    name?: pulumi.Input<string>;
    /**
     * Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
     */
    newGameSessionProtectionPolicy?: pulumi.Input<string>;
    /**
     * Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
     */
    resourceCreationLimitPolicy?: pulumi.Input<inputs.gamelift.FleetResourceCreationLimitPolicy>;
    /**
     * Instructions for launching server processes on each instance in the fleet. See below.
     */
    runtimeConfiguration?: pulumi.Input<inputs.gamelift.FleetRuntimeConfiguration>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
