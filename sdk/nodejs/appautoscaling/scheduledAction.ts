// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Application AutoScaling ScheduledAction resource.
 * 
 * ## Example Usage
 * 
 * ### DynamoDB Table Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const dynamodbTarget = new aws.appautoscaling.Target("dynamodb", {
 *     maxCapacity: 100,
 *     minCapacity: 5,
 *     resourceId: "table/tableName",
 *     roleArn: aws_iam_role_DynamoDBAutoscaleRole.arn,
 *     scalableDimension: "dynamodb:table:ReadCapacityUnits",
 *     serviceNamespace: "dynamodb",
 * });
 * const dynamodbScheduledAction = new aws.appautoscaling.ScheduledAction("dynamodb", {
 *     resourceId: dynamodbTarget.resourceId,
 *     scalableDimension: dynamodbTarget.scalableDimension,
 *     scalableTargetAction: {
 *         maxCapacity: 200,
 *         minCapacity: 1,
 *     },
 *     schedule: "at(2006-01-02T15:04:05)",
 *     serviceNamespace: dynamodbTarget.serviceNamespace,
 * });
 * ```
 * 
 * ### ECS Service Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ecsTarget = new aws.appautoscaling.Target("ecs", {
 *     maxCapacity: 4,
 *     minCapacity: 1,
 *     resourceId: "service/clusterName/serviceName",
 *     roleArn: var_ecs_iam_role,
 *     scalableDimension: "ecs:service:DesiredCount",
 *     serviceNamespace: "ecs",
 * });
 * const ecsScheduledAction = new aws.appautoscaling.ScheduledAction("ecs", {
 *     resourceId: ecsTarget.resourceId,
 *     scalableDimension: ecsTarget.scalableDimension,
 *     scalableTargetAction: {
 *         maxCapacity: 10,
 *         minCapacity: 1,
 *     },
 *     schedule: "at(2006-01-02T15:04:05)",
 *     serviceNamespace: ecsTarget.serviceNamespace,
 * });
 * ```
 */
export class ScheduledAction extends pulumi.CustomResource {
    /**
     * Get an existing ScheduledAction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduledActionState, opts?: pulumi.CustomResourceOptions): ScheduledAction {
        return new ScheduledAction(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:appautoscaling/scheduledAction:ScheduledAction';

    /**
     * Returns true if the given object is an instance of ScheduledAction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScheduledAction {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === ScheduledAction.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the scheduled action.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The date and time for the scheduled action to end. Specify the following format: 2006-01-02T15:04:05Z
     */
    public readonly endTime!: pulumi.Output<string | undefined>;
    /**
     * The name of the scheduled action.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
     */
    public readonly scalableDimension!: pulumi.Output<string | undefined>;
    /**
     * The new minimum and maximum capacity. You can set both values or just one. See below
     */
    public readonly scalableTargetAction!: pulumi.Output<{ maxCapacity?: number, minCapacity?: number } | undefined>;
    /**
     * The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). In UTC. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
     */
    public readonly schedule!: pulumi.Output<string | undefined>;
    /**
     * The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
     */
    public readonly serviceNamespace!: pulumi.Output<string>;
    /**
     * The date and time for the scheduled action to start. Specify the following format: 2006-01-02T15:04:05Z
     */
    public readonly startTime!: pulumi.Output<string | undefined>;

    /**
     * Create a ScheduledAction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduledActionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduledActionArgs | ScheduledActionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScheduledActionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["endTime"] = state ? state.endTime : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["scalableDimension"] = state ? state.scalableDimension : undefined;
            inputs["scalableTargetAction"] = state ? state.scalableTargetAction : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
            inputs["serviceNamespace"] = state ? state.serviceNamespace : undefined;
            inputs["startTime"] = state ? state.startTime : undefined;
        } else {
            const args = argsOrState as ScheduledActionArgs | undefined;
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.serviceNamespace === undefined) {
                throw new Error("Missing required property 'serviceNamespace'");
            }
            inputs["endTime"] = args ? args.endTime : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["scalableDimension"] = args ? args.scalableDimension : undefined;
            inputs["scalableTargetAction"] = args ? args.scalableTargetAction : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["serviceNamespace"] = args ? args.serviceNamespace : undefined;
            inputs["startTime"] = args ? args.startTime : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super(ScheduledAction.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScheduledAction resources.
 */
export interface ScheduledActionState {
    /**
     * The Amazon Resource Name (ARN) of the scheduled action.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The date and time for the scheduled action to end. Specify the following format: 2006-01-02T15:04:05Z
     */
    readonly endTime?: pulumi.Input<string>;
    /**
     * The name of the scheduled action.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
     */
    readonly scalableDimension?: pulumi.Input<string>;
    /**
     * The new minimum and maximum capacity. You can set both values or just one. See below
     */
    readonly scalableTargetAction?: pulumi.Input<{ maxCapacity?: pulumi.Input<number>, minCapacity?: pulumi.Input<number> }>;
    /**
     * The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). In UTC. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
     */
    readonly serviceNamespace?: pulumi.Input<string>;
    /**
     * The date and time for the scheduled action to start. Specify the following format: 2006-01-02T15:04:05Z
     */
    readonly startTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ScheduledAction resource.
 */
export interface ScheduledActionArgs {
    /**
     * The date and time for the scheduled action to end. Specify the following format: 2006-01-02T15:04:05Z
     */
    readonly endTime?: pulumi.Input<string>;
    /**
     * The name of the scheduled action.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The identifier of the resource associated with the scheduled action. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ResourceId)
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The scalable dimension. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ScalableDimension) Example: ecs:service:DesiredCount
     */
    readonly scalableDimension?: pulumi.Input<string>;
    /**
     * The new minimum and maximum capacity. You can set both values or just one. See below
     */
    readonly scalableTargetAction?: pulumi.Input<{ maxCapacity?: pulumi.Input<number>, minCapacity?: pulumi.Input<number> }>;
    /**
     * The schedule for this action. The following formats are supported: At expressions - at(yyyy-mm-ddThh:mm:ss), Rate expressions - rate(valueunit), Cron expressions - cron(fields). In UTC. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-Schedule)
     */
    readonly schedule?: pulumi.Input<string>;
    /**
     * The namespace of the AWS service. Documentation can be found in the parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_PutScheduledAction.html#ApplicationAutoScaling-PutScheduledAction-request-ServiceNamespace) Example: ecs
     */
    readonly serviceNamespace: pulumi.Input<string>;
    /**
     * The date and time for the scheduled action to start. Specify the following format: 2006-01-02T15:04:05Z
     */
    readonly startTime?: pulumi.Input<string>;
}
