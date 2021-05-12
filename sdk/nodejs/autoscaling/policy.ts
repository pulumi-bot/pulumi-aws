// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an AutoScaling Scaling Policy resource.
 *
 * > **NOTE:** You may want to omit `desiredCapacity` attribute from attached `aws.autoscaling.Group`
 * when using autoscaling policies. It's good practice to pick either
 * [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
 * or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
 * (policy-based) scaling.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bar = new aws.autoscaling.Group("bar", {
 *     availabilityZones: ["us-east-1a"],
 *     maxSize: 5,
 *     minSize: 2,
 *     healthCheckGracePeriod: 300,
 *     healthCheckType: "ELB",
 *     forceDelete: true,
 *     launchConfiguration: aws_launch_configuration.foo.name,
 * });
 * const bat = new aws.autoscaling.Policy("bat", {
 *     scalingAdjustment: 4,
 *     adjustmentType: "ChangeInCapacity",
 *     cooldown: 300,
 *     autoscalingGroupName: bar.name,
 * });
 * ```
 *
 * ## Import
 *
 * AutoScaling scaling policy can be imported using the role autoscaling_group_name and name separated by `/`.
 *
 * ```sh
 *  $ pulumi import aws:autoscaling/policy:Policy test-policy asg-name/policy-name
 * ```
 */
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
    public static readonly __pulumiType = 'aws:autoscaling/policy:Policy';

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

    /**
     * Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     */
    public readonly adjustmentType!: pulumi.Output<string | undefined>;
    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the autoscaling group.
     */
    public readonly autoscalingGroupName!: pulumi.Output<string>;
    /**
     * The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     */
    public readonly cooldown!: pulumi.Output<number | undefined>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
     */
    public readonly estimatedInstanceWarmup!: pulumi.Output<number | undefined>;
    /**
     * The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
     */
    public readonly metricAggregationType!: pulumi.Output<string>;
    /**
     * Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
     */
    public readonly minAdjustmentMagnitude!: pulumi.Output<number | undefined>;
    /**
     * The name of the dimension.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
     */
    public readonly policyType!: pulumi.Output<string | undefined>;
    /**
     * The number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     */
    public readonly scalingAdjustment!: pulumi.Output<number | undefined>;
    /**
     * A set of adjustments that manage
     * group scaling. These have the following structure:
     */
    public readonly stepAdjustments!: pulumi.Output<outputs.autoscaling.PolicyStepAdjustment[] | undefined>;
    /**
     * A target tracking policy. These have the following structure:
     */
    public readonly targetTrackingConfiguration!: pulumi.Output<outputs.autoscaling.PolicyTargetTrackingConfiguration | undefined>;

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
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            inputs["adjustmentType"] = state ? state.adjustmentType : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoscalingGroupName"] = state ? state.autoscalingGroupName : undefined;
            inputs["cooldown"] = state ? state.cooldown : undefined;
            inputs["estimatedInstanceWarmup"] = state ? state.estimatedInstanceWarmup : undefined;
            inputs["metricAggregationType"] = state ? state.metricAggregationType : undefined;
            inputs["minAdjustmentMagnitude"] = state ? state.minAdjustmentMagnitude : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["policyType"] = state ? state.policyType : undefined;
            inputs["scalingAdjustment"] = state ? state.scalingAdjustment : undefined;
            inputs["stepAdjustments"] = state ? state.stepAdjustments : undefined;
            inputs["targetTrackingConfiguration"] = state ? state.targetTrackingConfiguration : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.autoscalingGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoscalingGroupName'");
            }
            inputs["adjustmentType"] = args ? args.adjustmentType : undefined;
            inputs["autoscalingGroupName"] = args ? args.autoscalingGroupName : undefined;
            inputs["cooldown"] = args ? args.cooldown : undefined;
            inputs["estimatedInstanceWarmup"] = args ? args.estimatedInstanceWarmup : undefined;
            inputs["metricAggregationType"] = args ? args.metricAggregationType : undefined;
            inputs["minAdjustmentMagnitude"] = args ? args.minAdjustmentMagnitude : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["policyType"] = args ? args.policyType : undefined;
            inputs["scalingAdjustment"] = args ? args.scalingAdjustment : undefined;
            inputs["stepAdjustments"] = args ? args.stepAdjustments : undefined;
            inputs["targetTrackingConfiguration"] = args ? args.targetTrackingConfiguration : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Policy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    arn?: pulumi.Input<string>;
    /**
     * The name of the autoscaling group.
     */
    autoscalingGroupName?: pulumi.Input<string>;
    /**
     * The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
     */
    estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
     */
    metricAggregationType?: pulumi.Input<string>;
    /**
     * Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
     */
    minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * The name of the dimension.
     */
    name?: pulumi.Input<string>;
    /**
     * The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
     */
    policyType?: pulumi.Input<string>;
    /**
     * The number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     */
    scalingAdjustment?: pulumi.Input<number>;
    /**
     * A set of adjustments that manage
     * group scaling. These have the following structure:
     */
    stepAdjustments?: pulumi.Input<pulumi.Input<inputs.autoscaling.PolicyStepAdjustment>[]>;
    /**
     * A target tracking policy. These have the following structure:
     */
    targetTrackingConfiguration?: pulumi.Input<inputs.autoscaling.PolicyTargetTrackingConfiguration>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
     */
    adjustmentType?: pulumi.Input<string>;
    /**
     * The name of the autoscaling group.
     */
    autoscalingGroupName: pulumi.Input<string>;
    /**
     * The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
     */
    cooldown?: pulumi.Input<number>;
    /**
     * The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
     */
    estimatedInstanceWarmup?: pulumi.Input<number>;
    /**
     * The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
     */
    metricAggregationType?: pulumi.Input<string>;
    /**
     * Minimum value to scale by when `adjustmentType` is set to `PercentChangeInCapacity`.
     */
    minAdjustmentMagnitude?: pulumi.Input<number>;
    /**
     * The name of the dimension.
     */
    name?: pulumi.Input<string>;
    /**
     * The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
     */
    policyType?: pulumi.Input<string>;
    /**
     * The number of members by which to
     * scale, when the adjustment bounds are breached. A positive value scales
     * up. A negative value scales down.
     */
    scalingAdjustment?: pulumi.Input<number>;
    /**
     * A set of adjustments that manage
     * group scaling. These have the following structure:
     */
    stepAdjustments?: pulumi.Input<pulumi.Input<inputs.autoscaling.PolicyStepAdjustment>[]>;
    /**
     * A target tracking policy. These have the following structure:
     */
    targetTrackingConfiguration?: pulumi.Input<inputs.autoscaling.PolicyTargetTrackingConfiguration>;
}
