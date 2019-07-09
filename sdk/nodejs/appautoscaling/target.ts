// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Application AutoScaling ScalableTarget resource. To manage policies which get attached to the target, see the [`aws_appautoscaling_policy` resource](https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy.html).
 * 
 * ## Example Usage
 * 
 * ### DynamoDB Table Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const dynamodbTableReadTarget = new aws.appautoscaling.Target("dynamodb_table_read_target", {
 *     maxCapacity: 100,
 *     minCapacity: 5,
 *     resourceId: pulumi.interpolate`table/${aws_dynamodb_table_example.name}`,
 *     roleArn: aws_iam_role_DynamoDBAutoscaleRole.arn,
 *     scalableDimension: "dynamodb:table:ReadCapacityUnits",
 *     serviceNamespace: "dynamodb",
 * });
 * ```
 * 
 * ### DynamoDB Index Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const dynamodbIndexReadTarget = new aws.appautoscaling.Target("dynamodb_index_read_target", {
 *     maxCapacity: 100,
 *     minCapacity: 5,
 *     resourceId: pulumi.interpolate`table/${aws_dynamodb_table_example.name}/index/${var_index_name}`,
 *     roleArn: aws_iam_role_DynamoDBAutoscaleRole.arn,
 *     scalableDimension: "dynamodb:index:ReadCapacityUnits",
 *     serviceNamespace: "dynamodb",
 * });
 * ```
 * 
 * ### ECS Service Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ecsTarget = new aws.appautoscaling.Target("ecs_target", {
 *     maxCapacity: 4,
 *     minCapacity: 1,
 *     resourceId: pulumi.interpolate`service/${aws_ecs_cluster_example.name}/${aws_ecs_service_example.name}`,
 *     roleArn: var_ecs_iam_role,
 *     scalableDimension: "ecs:service:DesiredCount",
 *     serviceNamespace: "ecs",
 * });
 * ```
 * 
 * ### Aurora Read Replica Autoscaling
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const replicas = new aws.appautoscaling.Target("replicas", {
 *     maxCapacity: 15,
 *     minCapacity: 1,
 *     resourceId: pulumi.interpolate`cluster:${aws_rds_cluster_example.id}`,
 *     scalableDimension: "rds:cluster:ReadReplicaCount",
 *     serviceNamespace: "rds",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appautoscaling_target.html.markdown.
 */
export class Target extends pulumi.CustomResource {
    /**
     * Get an existing Target resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TargetState, opts?: pulumi.CustomResourceOptions): Target {
        return new Target(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appautoscaling/target:Target';

    /**
     * Returns true if the given object is an instance of Target.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Target {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Target.__pulumiType;
    }

    /**
     * The max capacity of the scalable target.
     */
    public readonly maxCapacity!: pulumi.Output<number>;
    /**
     * The min capacity of the scalable target.
     */
    public readonly minCapacity!: pulumi.Output<number>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly resourceId!: pulumi.Output<string>;
    /**
     * The ARN of the IAM role that allows Application
     * AutoScaling to modify your scalable target on your behalf.
     */
    public readonly roleArn!: pulumi.Output<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly scalableDimension!: pulumi.Output<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    public readonly serviceNamespace!: pulumi.Output<string>;

    /**
     * Create a Target resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TargetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TargetArgs | TargetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TargetState | undefined;
            inputs["maxCapacity"] = state ? state.maxCapacity : undefined;
            inputs["minCapacity"] = state ? state.minCapacity : undefined;
            inputs["resourceId"] = state ? state.resourceId : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["scalableDimension"] = state ? state.scalableDimension : undefined;
            inputs["serviceNamespace"] = state ? state.serviceNamespace : undefined;
        } else {
            const args = argsOrState as TargetArgs | undefined;
            if (!args || args.maxCapacity === undefined) {
                throw new Error("Missing required property 'maxCapacity'");
            }
            if (!args || args.minCapacity === undefined) {
                throw new Error("Missing required property 'minCapacity'");
            }
            if (!args || args.resourceId === undefined) {
                throw new Error("Missing required property 'resourceId'");
            }
            if (!args || args.scalableDimension === undefined) {
                throw new Error("Missing required property 'scalableDimension'");
            }
            if (!args || args.serviceNamespace === undefined) {
                throw new Error("Missing required property 'serviceNamespace'");
            }
            inputs["maxCapacity"] = args ? args.maxCapacity : undefined;
            inputs["minCapacity"] = args ? args.minCapacity : undefined;
            inputs["resourceId"] = args ? args.resourceId : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["scalableDimension"] = args ? args.scalableDimension : undefined;
            inputs["serviceNamespace"] = args ? args.serviceNamespace : undefined;
        }
        super(Target.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Target resources.
 */
export interface TargetState {
    /**
     * The max capacity of the scalable target.
     */
    readonly maxCapacity?: pulumi.Input<number>;
    /**
     * The min capacity of the scalable target.
     */
    readonly minCapacity?: pulumi.Input<number>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId?: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that allows Application
     * AutoScaling to modify your scalable target on your behalf.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension?: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Target resource.
 */
export interface TargetArgs {
    /**
     * The max capacity of the scalable target.
     */
    readonly maxCapacity: pulumi.Input<number>;
    /**
     * The min capacity of the scalable target.
     */
    readonly minCapacity: pulumi.Input<number>;
    /**
     * The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly resourceId: pulumi.Input<string>;
    /**
     * The ARN of the IAM role that allows Application
     * AutoScaling to modify your scalable target on your behalf.
     */
    readonly roleArn?: pulumi.Input<string>;
    /**
     * The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly scalableDimension: pulumi.Input<string>;
    /**
     * The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
     */
    readonly serviceNamespace: pulumi.Input<string>;
}
