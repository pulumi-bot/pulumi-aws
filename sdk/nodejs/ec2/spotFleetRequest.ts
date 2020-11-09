// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an EC2 Spot Fleet Request resource. This allows a fleet of Spot
 * instances to be requested on the Spot market.
 *
 * ## Example Usage
 * ### Using launch specifications
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Request a Spot fleet
 * const cheapCompute = new aws.ec2.SpotFleetRequest("cheapCompute", {
 *     iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
 *     spotPrice: "0.03",
 *     allocationStrategy: "diversified",
 *     targetCapacity: 6,
 *     validUntil: "2019-11-04T20:44:20Z",
 *     launchSpecifications: [
 *         {
 *             instanceType: "m4.10xlarge",
 *             ami: "ami-1234",
 *             spotPrice: "2.793",
 *             placementTenancy: "dedicated",
 *             iamInstanceProfileArn: aws_iam_instance_profile.example.arn,
 *         },
 *         {
 *             instanceType: "m4.4xlarge",
 *             ami: "ami-5678",
 *             keyName: "my-key",
 *             spotPrice: "1.117",
 *             iamInstanceProfileArn: aws_iam_instance_profile.example.arn,
 *             availabilityZone: "us-west-1a",
 *             subnetId: "subnet-1234",
 *             weightedCapacity: 35,
 *             rootBlockDevices: [{
 *                 volumeSize: "300",
 *                 volumeType: "gp2",
 *             }],
 *             tags: {
 *                 Name: "spot-fleet-example",
 *             },
 *         },
 *     ],
 * });
 * ```
 * ### Using launch templates
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const fooLaunchTemplate = new aws.ec2.LaunchTemplate("fooLaunchTemplate", {
 *     imageId: "ami-516b9131",
 *     instanceType: "m1.small",
 *     keyName: "some-key",
 * });
 * const fooSpotFleetRequest = new aws.ec2.SpotFleetRequest("fooSpotFleetRequest", {
 *     iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
 *     spotPrice: "0.005",
 *     targetCapacity: 2,
 *     validUntil: "2019-11-04T20:44:20Z",
 *     launchTemplateConfigs: [{
 *         launchTemplateSpecification: {
 *             id: fooLaunchTemplate.id,
 *             version: fooLaunchTemplate.latestVersion,
 *         },
 *     }],
 * }, {
 *     dependsOn: [aws_iam_policy_attachment["test-attach"]],
 * });
 * ```
 *
 * > **NOTE:** This provider does not support the functionality where multiple `subnetId` or `availabilityZone` parameters can be specified in the same
 * launch configuration block. If you want to specify multiple values, then separate launch configuration blocks should be used:
 * ### Using multiple launch specifications
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ec2.SpotFleetRequest("foo", {
 *     iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
 *     launchSpecifications: [
 *         {
 *             ami: "ami-d06a90b0",
 *             availabilityZone: "us-west-2a",
 *             instanceType: "m1.small",
 *             keyName: "my-key",
 *         },
 *         {
 *             ami: "ami-d06a90b0",
 *             availabilityZone: "us-west-2a",
 *             instanceType: "m5.large",
 *             keyName: "my-key",
 *         },
 *     ],
 *     spotPrice: "0.005",
 *     targetCapacity: 2,
 *     validUntil: "2019-11-04T20:44:20Z",
 * });
 * ```
 * ### Using multiple launch configurations
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = aws.ec2.getSubnetIds({
 *     vpcId: _var.vpc_id,
 * });
 * const fooLaunchTemplate = new aws.ec2.LaunchTemplate("fooLaunchTemplate", {
 *     imageId: "ami-516b9131",
 *     instanceType: "m1.small",
 *     keyName: "some-key",
 * });
 * const fooSpotFleetRequest = new aws.ec2.SpotFleetRequest("fooSpotFleetRequest", {
 *     iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
 *     spotPrice: "0.005",
 *     targetCapacity: 2,
 *     validUntil: "2019-11-04T20:44:20Z",
 *     launchTemplateConfigs: [{
 *         launchTemplateSpecification: {
 *             id: fooLaunchTemplate.id,
 *             version: fooLaunchTemplate.latestVersion,
 *         },
 *         overrides: [
 *             {
 *                 subnetId: data.aws_subnets.example.ids[0],
 *             },
 *             {
 *                 subnetId: data.aws_subnets.example.ids[1],
 *             },
 *             {
 *                 subnetId: data.aws_subnets.example.ids[2],
 *             },
 *         ],
 *     }],
 * }, {
 *     dependsOn: [aws_iam_policy_attachment["test-attach"]],
 * });
 * ```
 */
export class SpotFleetRequest extends pulumi.CustomResource {
    /**
     * Get an existing SpotFleetRequest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpotFleetRequestState, opts?: pulumi.CustomResourceOptions): SpotFleetRequest {
        return new SpotFleetRequest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/spotFleetRequest:SpotFleetRequest';

    /**
     * Returns true if the given object is an instance of SpotFleetRequest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpotFleetRequest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpotFleetRequest.__pulumiType;
    }

    /**
     * Indicates how to allocate the target capacity across
     * the Spot pools specified by the Spot fleet request. The default is
     * `lowestPrice`.
     */
    public readonly allocationStrategy!: pulumi.Output<string | undefined>;
    public /*out*/ readonly clientToken!: pulumi.Output<string>;
    /**
     * Indicates whether running Spot
     * instances should be terminated if the target capacity of the Spot fleet
     * request is decreased below the current size of the Spot fleet.
     */
    public readonly excessCapacityTerminationPolicy!: pulumi.Output<string | undefined>;
    /**
     * The type of fleet request. Indicates whether the Spot Fleet only requests the target
     * capacity or also attempts to maintain it. Default is `maintain`.
     */
    public readonly fleetType!: pulumi.Output<string | undefined>;
    /**
     * Grants the Spot fleet permission to terminate
     * Spot instances on your behalf when you cancel its Spot fleet request using
     * CancelSpotFleetRequests or when the Spot fleet request expires, if you set
     * terminateInstancesWithExpiration.
     */
    public readonly iamFleetRole!: pulumi.Output<string>;
    /**
     * Indicates whether a Spot
     * instance stops or terminates when it is interrupted. Default is
     * `terminate`.
     */
    public readonly instanceInterruptionBehaviour!: pulumi.Output<string | undefined>;
    /**
     * The number of Spot pools across which to allocate your target Spot capacity.
     * Valid only when `allocationStrategy` is set to `lowestPrice`. Spot Fleet selects
     * the cheapest Spot pools and evenly allocates your target Spot capacity across
     * the number of Spot pools that you specify.
     */
    public readonly instancePoolsToUseCount!: pulumi.Output<number | undefined>;
    /**
     * Used to define the launch configuration of the
     * spot-fleet request. Can be specified multiple times to define different bids
     * across different markets and instance types. Conflicts with `launchTemplateConfig`. At least one of `launchSpecification` or `launchTemplateConfig` is required.
     */
    public readonly launchSpecifications!: pulumi.Output<outputs.ec2.SpotFleetRequestLaunchSpecification[] | undefined>;
    /**
     * Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launchSpecification`. At least one of `launchSpecification` or `launchTemplateConfig` is required.
     */
    public readonly launchTemplateConfigs!: pulumi.Output<outputs.ec2.SpotFleetRequestLaunchTemplateConfig[] | undefined>;
    /**
     * A list of elastic load balancer names to add to the Spot fleet.
     */
    public readonly loadBalancers!: pulumi.Output<string[]>;
    /**
     * Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
     */
    public readonly replaceUnhealthyInstances!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum spot bid for this override request.
     */
    public readonly spotPrice!: pulumi.Output<string | undefined>;
    /**
     * The state of the Spot fleet request.
     */
    public /*out*/ readonly spotRequestState!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The number of units to request. You can choose to set the
     * target capacity in terms of instances or a performance characteristic that is
     * important to your application workload, such as vCPUs, memory, or I/O.
     */
    public readonly targetCapacity!: pulumi.Output<number>;
    /**
     * A list of `aws.alb.TargetGroup` ARNs, for use with Application Load Balancing.
     */
    public readonly targetGroupArns!: pulumi.Output<string[]>;
    /**
     * Indicates whether running Spot
     * instances should be terminated when the Spot fleet request expires.
     */
    public readonly terminateInstancesWithExpiration!: pulumi.Output<boolean | undefined>;
    /**
     * The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
     */
    public readonly validFrom!: pulumi.Output<string | undefined>;
    /**
     * The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
     */
    public readonly validUntil!: pulumi.Output<string | undefined>;
    /**
     * If set, this provider will
     * wait for the Spot Request to be fulfilled, and will throw an error if the
     * timeout of 10m is reached.
     */
    public readonly waitForFulfillment!: pulumi.Output<boolean | undefined>;

    /**
     * Create a SpotFleetRequest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SpotFleetRequestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpotFleetRequestArgs | SpotFleetRequestState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SpotFleetRequestState | undefined;
            inputs["allocationStrategy"] = state ? state.allocationStrategy : undefined;
            inputs["clientToken"] = state ? state.clientToken : undefined;
            inputs["excessCapacityTerminationPolicy"] = state ? state.excessCapacityTerminationPolicy : undefined;
            inputs["fleetType"] = state ? state.fleetType : undefined;
            inputs["iamFleetRole"] = state ? state.iamFleetRole : undefined;
            inputs["instanceInterruptionBehaviour"] = state ? state.instanceInterruptionBehaviour : undefined;
            inputs["instancePoolsToUseCount"] = state ? state.instancePoolsToUseCount : undefined;
            inputs["launchSpecifications"] = state ? state.launchSpecifications : undefined;
            inputs["launchTemplateConfigs"] = state ? state.launchTemplateConfigs : undefined;
            inputs["loadBalancers"] = state ? state.loadBalancers : undefined;
            inputs["replaceUnhealthyInstances"] = state ? state.replaceUnhealthyInstances : undefined;
            inputs["spotPrice"] = state ? state.spotPrice : undefined;
            inputs["spotRequestState"] = state ? state.spotRequestState : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetCapacity"] = state ? state.targetCapacity : undefined;
            inputs["targetGroupArns"] = state ? state.targetGroupArns : undefined;
            inputs["terminateInstancesWithExpiration"] = state ? state.terminateInstancesWithExpiration : undefined;
            inputs["validFrom"] = state ? state.validFrom : undefined;
            inputs["validUntil"] = state ? state.validUntil : undefined;
            inputs["waitForFulfillment"] = state ? state.waitForFulfillment : undefined;
        } else {
            const args = argsOrState as SpotFleetRequestArgs | undefined;
            if (!args || args.iamFleetRole === undefined) {
                throw new Error("Missing required property 'iamFleetRole'");
            }
            if (!args || args.targetCapacity === undefined) {
                throw new Error("Missing required property 'targetCapacity'");
            }
            inputs["allocationStrategy"] = args ? args.allocationStrategy : undefined;
            inputs["excessCapacityTerminationPolicy"] = args ? args.excessCapacityTerminationPolicy : undefined;
            inputs["fleetType"] = args ? args.fleetType : undefined;
            inputs["iamFleetRole"] = args ? args.iamFleetRole : undefined;
            inputs["instanceInterruptionBehaviour"] = args ? args.instanceInterruptionBehaviour : undefined;
            inputs["instancePoolsToUseCount"] = args ? args.instancePoolsToUseCount : undefined;
            inputs["launchSpecifications"] = args ? args.launchSpecifications : undefined;
            inputs["launchTemplateConfigs"] = args ? args.launchTemplateConfigs : undefined;
            inputs["loadBalancers"] = args ? args.loadBalancers : undefined;
            inputs["replaceUnhealthyInstances"] = args ? args.replaceUnhealthyInstances : undefined;
            inputs["spotPrice"] = args ? args.spotPrice : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetCapacity"] = args ? args.targetCapacity : undefined;
            inputs["targetGroupArns"] = args ? args.targetGroupArns : undefined;
            inputs["terminateInstancesWithExpiration"] = args ? args.terminateInstancesWithExpiration : undefined;
            inputs["validFrom"] = args ? args.validFrom : undefined;
            inputs["validUntil"] = args ? args.validUntil : undefined;
            inputs["waitForFulfillment"] = args ? args.waitForFulfillment : undefined;
            inputs["clientToken"] = undefined /*out*/;
            inputs["spotRequestState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SpotFleetRequest.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpotFleetRequest resources.
 */
export interface SpotFleetRequestState {
    /**
     * Indicates how to allocate the target capacity across
     * the Spot pools specified by the Spot fleet request. The default is
     * `lowestPrice`.
     */
    readonly allocationStrategy?: pulumi.Input<string>;
    readonly clientToken?: pulumi.Input<string>;
    /**
     * Indicates whether running Spot
     * instances should be terminated if the target capacity of the Spot fleet
     * request is decreased below the current size of the Spot fleet.
     */
    readonly excessCapacityTerminationPolicy?: pulumi.Input<string>;
    /**
     * The type of fleet request. Indicates whether the Spot Fleet only requests the target
     * capacity or also attempts to maintain it. Default is `maintain`.
     */
    readonly fleetType?: pulumi.Input<string>;
    /**
     * Grants the Spot fleet permission to terminate
     * Spot instances on your behalf when you cancel its Spot fleet request using
     * CancelSpotFleetRequests or when the Spot fleet request expires, if you set
     * terminateInstancesWithExpiration.
     */
    readonly iamFleetRole?: pulumi.Input<string>;
    /**
     * Indicates whether a Spot
     * instance stops or terminates when it is interrupted. Default is
     * `terminate`.
     */
    readonly instanceInterruptionBehaviour?: pulumi.Input<string>;
    /**
     * The number of Spot pools across which to allocate your target Spot capacity.
     * Valid only when `allocationStrategy` is set to `lowestPrice`. Spot Fleet selects
     * the cheapest Spot pools and evenly allocates your target Spot capacity across
     * the number of Spot pools that you specify.
     */
    readonly instancePoolsToUseCount?: pulumi.Input<number>;
    /**
     * Used to define the launch configuration of the
     * spot-fleet request. Can be specified multiple times to define different bids
     * across different markets and instance types. Conflicts with `launchTemplateConfig`. At least one of `launchSpecification` or `launchTemplateConfig` is required.
     */
    readonly launchSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.SpotFleetRequestLaunchSpecification>[]>;
    /**
     * Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launchSpecification`. At least one of `launchSpecification` or `launchTemplateConfig` is required.
     */
    readonly launchTemplateConfigs?: pulumi.Input<pulumi.Input<inputs.ec2.SpotFleetRequestLaunchTemplateConfig>[]>;
    /**
     * A list of elastic load balancer names to add to the Spot fleet.
     */
    readonly loadBalancers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
     */
    readonly replaceUnhealthyInstances?: pulumi.Input<boolean>;
    /**
     * The maximum spot bid for this override request.
     */
    readonly spotPrice?: pulumi.Input<string>;
    /**
     * The state of the Spot fleet request.
     */
    readonly spotRequestState?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The number of units to request. You can choose to set the
     * target capacity in terms of instances or a performance characteristic that is
     * important to your application workload, such as vCPUs, memory, or I/O.
     */
    readonly targetCapacity?: pulumi.Input<number>;
    /**
     * A list of `aws.alb.TargetGroup` ARNs, for use with Application Load Balancing.
     */
    readonly targetGroupArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether running Spot
     * instances should be terminated when the Spot fleet request expires.
     */
    readonly terminateInstancesWithExpiration?: pulumi.Input<boolean>;
    /**
     * The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
     */
    readonly validFrom?: pulumi.Input<string>;
    /**
     * The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
     */
    readonly validUntil?: pulumi.Input<string>;
    /**
     * If set, this provider will
     * wait for the Spot Request to be fulfilled, and will throw an error if the
     * timeout of 10m is reached.
     */
    readonly waitForFulfillment?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a SpotFleetRequest resource.
 */
export interface SpotFleetRequestArgs {
    /**
     * Indicates how to allocate the target capacity across
     * the Spot pools specified by the Spot fleet request. The default is
     * `lowestPrice`.
     */
    readonly allocationStrategy?: pulumi.Input<string>;
    /**
     * Indicates whether running Spot
     * instances should be terminated if the target capacity of the Spot fleet
     * request is decreased below the current size of the Spot fleet.
     */
    readonly excessCapacityTerminationPolicy?: pulumi.Input<string>;
    /**
     * The type of fleet request. Indicates whether the Spot Fleet only requests the target
     * capacity or also attempts to maintain it. Default is `maintain`.
     */
    readonly fleetType?: pulumi.Input<string>;
    /**
     * Grants the Spot fleet permission to terminate
     * Spot instances on your behalf when you cancel its Spot fleet request using
     * CancelSpotFleetRequests or when the Spot fleet request expires, if you set
     * terminateInstancesWithExpiration.
     */
    readonly iamFleetRole: pulumi.Input<string>;
    /**
     * Indicates whether a Spot
     * instance stops or terminates when it is interrupted. Default is
     * `terminate`.
     */
    readonly instanceInterruptionBehaviour?: pulumi.Input<string>;
    /**
     * The number of Spot pools across which to allocate your target Spot capacity.
     * Valid only when `allocationStrategy` is set to `lowestPrice`. Spot Fleet selects
     * the cheapest Spot pools and evenly allocates your target Spot capacity across
     * the number of Spot pools that you specify.
     */
    readonly instancePoolsToUseCount?: pulumi.Input<number>;
    /**
     * Used to define the launch configuration of the
     * spot-fleet request. Can be specified multiple times to define different bids
     * across different markets and instance types. Conflicts with `launchTemplateConfig`. At least one of `launchSpecification` or `launchTemplateConfig` is required.
     */
    readonly launchSpecifications?: pulumi.Input<pulumi.Input<inputs.ec2.SpotFleetRequestLaunchSpecification>[]>;
    /**
     * Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launchSpecification`. At least one of `launchSpecification` or `launchTemplateConfig` is required.
     */
    readonly launchTemplateConfigs?: pulumi.Input<pulumi.Input<inputs.ec2.SpotFleetRequestLaunchTemplateConfig>[]>;
    /**
     * A list of elastic load balancer names to add to the Spot fleet.
     */
    readonly loadBalancers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
     */
    readonly replaceUnhealthyInstances?: pulumi.Input<boolean>;
    /**
     * The maximum spot bid for this override request.
     */
    readonly spotPrice?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The number of units to request. You can choose to set the
     * target capacity in terms of instances or a performance characteristic that is
     * important to your application workload, such as vCPUs, memory, or I/O.
     */
    readonly targetCapacity: pulumi.Input<number>;
    /**
     * A list of `aws.alb.TargetGroup` ARNs, for use with Application Load Balancing.
     */
    readonly targetGroupArns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Indicates whether running Spot
     * instances should be terminated when the Spot fleet request expires.
     */
    readonly terminateInstancesWithExpiration?: pulumi.Input<boolean>;
    /**
     * The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
     */
    readonly validFrom?: pulumi.Input<string>;
    /**
     * The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
     */
    readonly validUntil?: pulumi.Input<string>;
    /**
     * If set, this provider will
     * wait for the Spot Request to be fulfilled, and will throw an error if the
     * timeout of 10m is reached.
     */
    readonly waitForFulfillment?: pulumi.Input<boolean>;
}
