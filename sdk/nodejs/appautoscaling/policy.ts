// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an Application AutoScaling Policy resource.
 *
 * ## Example Usage
 * ### DynamoDB Table Autoscaling
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodbTableReadTarget", {
 *     maxCapacity: 100,
 *     minCapacity: 5,
 *     resourceId: "table/tableName",
 *     scalableDimension: "dynamodb:table:ReadCapacityUnits",
 *     serviceNamespace: "dynamodb",
 * });
 * const dynamodbTableReadPolicy = new aws.appautoscaling.Policy("dynamodbTableReadPolicy", {
 *     policyType: "TargetTrackingScaling",
 *     resourceId: dynamodbTableReadTarget.resourceId,
 *     scalableDimension: dynamodbTableReadTarget.scalableDimension,
 *     serviceNamespace: dynamodbTableReadTarget.serviceNamespace,
 *     targetTrackingScalingPolicyConfiguration: {
 *         predefinedMetricSpecification: {
 *             predefinedMetricType: "DynamoDBReadCapacityUtilization",
 *         },
 *         targetValue: 70,
 *     },
 * });
 * ```
 * ### ECS Service Autoscaling
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ecsTarget = new aws.appautoscaling.Target("ecsTarget", {
 *     maxCapacity: 4,
 *     minCapacity: 1,
 *     resourceId: "service/clusterName/serviceName",
 *     scalableDimension: "ecs:service:DesiredCount",
 *     serviceNamespace: "ecs",
 * });
 * const ecsPolicy = new aws.appautoscaling.Policy("ecsPolicy", {
 *     policyType: "StepScaling",
 *     resourceId: ecsTarget.resourceId,
 *     scalableDimension: ecsTarget.scalableDimension,
 *     serviceNamespace: ecsTarget.serviceNamespace,
 *     stepScalingPolicyConfiguration: {
 *         adjustmentType: "ChangeInCapacity",
 *         cooldown: 60,
 *         metricAggregationType: "Maximum",
 *         stepAdjustments: [{
 *             metricIntervalUpperBound: 0,
 *             scalingAdjustment: -1,
 *         }],
 *     },
 * });
 * ```
 * ### Preserve desired count when updating an autoscaled ECS Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ecsService = new aws.ecs.Service("ecsService", {
 *     cluster: "clusterName",
 *     taskDefinition: "taskDefinitionFamily:1",
 *     desiredCount: 2,
 * });
 * ```
 * ### Aurora Read Replica Autoscaling
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const replicasTarget = new aws.appautoscaling.Target("replicasTarget", {
 *     serviceNamespace: "rds",
 *     scalableDimension: "rds:cluster:ReadReplicaCount",
 *     resourceId: `cluster:${aws_rds_cluster.example.id}`,
 *     minCapacity: 1,
 *     maxCapacity: 15,
 * });
 * const replicasPolicy = new aws.appautoscaling.Policy("replicasPolicy", {
 *     serviceNamespace: replicasTarget.serviceNamespace,
 *     scalableDimension: replicasTarget.scalableDimension,
 *     resourceId: replicasTarget.resourceId,
 *     policyType: "TargetTrackingScaling",
 *     targetTrackingScalingPolicyConfiguration: {
 *         predefinedMetricSpecification: {
 *             predefinedMetricType: "RDSReaderAverageCPUUtilization",
 *         },
 *         targetValue: 75,
 *         scaleInCooldown: 300,
 *         scaleOutCooldown: 300,
 *     },
 * });
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

    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the policy.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
     */
    public readonly policyType!: pulumi.Output<string | undefined>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly scalableDimension!: pulumi.Output<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly serviceNamespace!: pulumi.Output<string>;
    /**
     * Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
     */
    public readonly stepScalingPolicyConfiguration!: pulumi.Output<outputs.appautoscaling.PolicyStepScalingPolicyConfiguration | undefined>;
    /**
     * A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
     */
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
    /**
     * The ARN assigned by AWS to the scaling policy.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
     */
    readonly policyType?: pulumi.Input<string>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension?: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace?: pulumi.Input<string>;
    /**
     * Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
     */
    readonly stepScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyStepScalingPolicyConfiguration>;
    /**
     * A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
     */
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * The name of the policy.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
     */
    readonly policyType?: pulumi.Input<string>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace: pulumi.Input<string>;
    /**
     * Step scaling policy configuration, requires `policyType = "StepScaling"` (default). See supported fields below.
     */
    readonly stepScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyStepScalingPolicyConfiguration>;
    /**
     * A target tracking policy, requires `policyType = "TargetTrackingScaling"`. See supported fields below.
     */
    readonly targetTrackingScalingPolicyConfiguration?: pulumi.Input<inputs.appautoscaling.PolicyTargetTrackingScalingPolicyConfiguration>;
}
