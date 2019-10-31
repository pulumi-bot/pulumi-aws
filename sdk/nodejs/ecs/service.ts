// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > **Note:** To prevent a race condition during service deletion, make sure to set `dependsOn` to the related `aws.iam.RolePolicy`; otherwise, the policy may be destroyed too soon and the ECS service will then get stuck in the `DRAINING` state.
 * 
 * Provides an ECS service - effectively a task that is expected to run until an error occurs or a user terminates it (typically a webserver or a database).
 * 
 * See [ECS Services section in AWS developer guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html).
 * 
 * ## Example Usage
 * 
 * ### Ignoring Changes to Desired Count
 * 
 * You can utilize the generic this provider resource [lifecycle configuration block](https://www.terraform.io/docs/configuration/resources.html#lifecycle) with `ignoreChanges` to create an ECS service with an initial count of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ecs.Service("example", {
 *     // Example: Create service with 2 instances to start
 *     desiredCount: 2,
 * }, {ignoreChanges: ["desiredCount"]});
 * ```
 * 
 * ### Daemon Scheduling Strategy
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const bar = new aws.ecs.Service("bar", {
 *     cluster: aws_ecs_cluster_foo.id,
 *     schedulingStrategy: "DAEMON",
 *     taskDefinition: aws_ecs_task_definition_bar.arn,
 * });
 * ```
 * 
 * ## deploymentController
 * 
 * The `deploymentController` configuration block supports the following:
 * 
 * * `type` - (Optional) Type of deployment controller. Valid values: `CODE_DEPLOY`, `ECS`. Default: `ECS`.
 * 
 * ## loadBalancer
 * 
 * `loadBalancer` supports the following:
 * 
 * * `elbName` - (Required for ELB Classic) The name of the ELB (Classic) to associate with the service.
 * * `targetGroupArn` - (Required for ALB/NLB) The ARN of the Load Balancer target group to associate with the service.
 * * `containerName` - (Required) The name of the container to associate with the load balancer (as it appears in a container definition).
 * * `containerPort` - (Required) The port on the container to associate with the load balancer.
 * 
 * > **Version note:** Multiple `loadBalancer` configuration block support was added in version 2.22.0 of the provider. This allows configuration of [ECS service support for multiple target groups](https://aws.amazon.com/about-aws/whats-new/2019/07/amazon-ecs-services-now-support-multiple-load-balancer-target-groups/).
 * 
 * ## orderedPlacementStrategy
 * 
 * `orderedPlacementStrategy` supports the following:
 * 
 * * `type` - (Required) The type of placement strategy. Must be one of: `binpack`, `random`, or `spread`
 * * `field` - (Optional) For the `spread` placement strategy, valid values are `instanceId` (or `host`,
 *  which has the same effect), or any platform or custom attribute that is applied to a container instance.
 *  For the `binpack` type, valid values are `memory` and `cpu`. For the `random` type, this attribute is not
 *  needed. For more information, see [Placement Strategy](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PlacementStrategy.html).
 * 
 * > **Note:** for `spread`, `host` and `instanceId` will be normalized, by AWS, to be `instanceId`. This means the statefile will show `instanceId` but your config will differ if you use `host`.
 * 
 * ## placementConstraints
 * 
 * `placementConstraints` support the following:
 * 
 * * `type` - (Required) The type of constraint. The only valid values at this time are `memberOf` and `distinctInstance`.
 * * `expression` -  (Optional) Cluster Query Language expression to apply to the constraint. Does not need to be specified
 * for the `distinctInstance` type.
 * For more information, see [Cluster Query Language in the Amazon EC2 Container
 * Service Developer
 * Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
 * 
 * ## networkConfiguration
 * 
 * `networkConfiguration` support the following:
 * 
 * * `subnets` - (Required) The subnets associated with the task or service.
 * * `securityGroups` - (Optional) The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
 * * `assignPublicIp` - (Optional) Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.
 * 
 * For more information, see [Task Networking](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-networking.html)
 * 
 * ## serviceRegistries
 * 
 * `serviceRegistries` support the following:
 * 
 * * `registryArn` - (Required) The ARN of the Service Registry. The currently supported service registry is Amazon Route 53 Auto Naming Service(`aws.servicediscovery.Service`). For more information, see [Service](https://docs.aws.amazon.com/Route53/latest/APIReference/API_autonaming_Service.html)
 * * `port` - (Optional) The port value used if your Service Discovery service specified an SRV record.
 * * `containerPort` - (Optional) The port value, already specified in the task definition, to be used for your service discovery service.
 * * `containerName` - (Optional) The container name value, already specified in the task definition, to be used for your service discovery service.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_service.html.markdown.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ecs/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * ARN of an ECS cluster
     */
    public readonly cluster!: pulumi.Output<string>;
    /**
     * Configuration block containing deployment controller configuration. Defined below.
     */
    public readonly deploymentController!: pulumi.Output<outputs.ecs.ServiceDeploymentController | undefined>;
    /**
     * The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
     */
    public readonly deploymentMaximumPercent!: pulumi.Output<number | undefined>;
    /**
     * The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
     */
    public readonly deploymentMinimumHealthyPercent!: pulumi.Output<number | undefined>;
    /**
     * The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
     */
    public readonly desiredCount!: pulumi.Output<number | undefined>;
    /**
     * Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
     */
    public readonly enableEcsManagedTags!: pulumi.Output<boolean | undefined>;
    /**
     * Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
     */
    public readonly healthCheckGracePeriodSeconds!: pulumi.Output<number | undefined>;
    /**
     * ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
     */
    public readonly iamRole!: pulumi.Output<string>;
    /**
     * The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
     */
    public readonly launchType!: pulumi.Output<string | undefined>;
    /**
     * A load balancer block. Load balancers documented below.
     */
    public readonly loadBalancers!: pulumi.Output<outputs.ecs.ServiceLoadBalancer[] | undefined>;
    /**
     * The name of the service (up to 255 letters, numbers, hyphens, and underscores)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
     */
    public readonly networkConfiguration!: pulumi.Output<outputs.ecs.ServiceNetworkConfiguration | undefined>;
    /**
     * Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
     */
    public readonly orderedPlacementStrategies!: pulumi.Output<outputs.ecs.ServiceOrderedPlacementStrategy[] | undefined>;
    /**
     * rules that are taken into consideration during task placement. Maximum number of
     * `placementConstraints` is `10`. Defined below.
     */
    public readonly placementConstraints!: pulumi.Output<outputs.ecs.ServicePlacementConstraint[] | undefined>;
    /**
     * The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
     */
    public readonly platformVersion!: pulumi.Output<string>;
    /**
     * Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
     */
    public readonly propagateTags!: pulumi.Output<string | undefined>;
    /**
     * The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Fargate tasks do not support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html).
     */
    public readonly schedulingStrategy!: pulumi.Output<string | undefined>;
    /**
     * The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
     */
    public readonly serviceRegistries!: pulumi.Output<outputs.ecs.ServiceServiceRegistries | undefined>;
    /**
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
     */
    public readonly taskDefinition!: pulumi.Output<string>;
    /**
     * If `true`, this provider will wait for the service to reach a steady state (like [`aws ecs wait services-stable`](https://docs.aws.amazon.com/cli/latest/reference/ecs/wait/services-stable.html)) before continuing. Default `false`.
     */
    public readonly waitForSteadyState!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServiceState | undefined;
            inputs["cluster"] = state ? state.cluster : undefined;
            inputs["deploymentController"] = state ? state.deploymentController : undefined;
            inputs["deploymentMaximumPercent"] = state ? state.deploymentMaximumPercent : undefined;
            inputs["deploymentMinimumHealthyPercent"] = state ? state.deploymentMinimumHealthyPercent : undefined;
            inputs["desiredCount"] = state ? state.desiredCount : undefined;
            inputs["enableEcsManagedTags"] = state ? state.enableEcsManagedTags : undefined;
            inputs["healthCheckGracePeriodSeconds"] = state ? state.healthCheckGracePeriodSeconds : undefined;
            inputs["iamRole"] = state ? state.iamRole : undefined;
            inputs["launchType"] = state ? state.launchType : undefined;
            inputs["loadBalancers"] = state ? state.loadBalancers : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkConfiguration"] = state ? state.networkConfiguration : undefined;
            inputs["orderedPlacementStrategies"] = state ? state.orderedPlacementStrategies : undefined;
            inputs["placementConstraints"] = state ? state.placementConstraints : undefined;
            inputs["platformVersion"] = state ? state.platformVersion : undefined;
            inputs["propagateTags"] = state ? state.propagateTags : undefined;
            inputs["schedulingStrategy"] = state ? state.schedulingStrategy : undefined;
            inputs["serviceRegistries"] = state ? state.serviceRegistries : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["taskDefinition"] = state ? state.taskDefinition : undefined;
            inputs["waitForSteadyState"] = state ? state.waitForSteadyState : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if (!args || args.taskDefinition === undefined) {
                throw new Error("Missing required property 'taskDefinition'");
            }
            inputs["cluster"] = args ? args.cluster : undefined;
            inputs["deploymentController"] = args ? args.deploymentController : undefined;
            inputs["deploymentMaximumPercent"] = args ? args.deploymentMaximumPercent : undefined;
            inputs["deploymentMinimumHealthyPercent"] = args ? args.deploymentMinimumHealthyPercent : undefined;
            inputs["desiredCount"] = args ? args.desiredCount : undefined;
            inputs["enableEcsManagedTags"] = args ? args.enableEcsManagedTags : undefined;
            inputs["healthCheckGracePeriodSeconds"] = args ? args.healthCheckGracePeriodSeconds : undefined;
            inputs["iamRole"] = args ? args.iamRole : undefined;
            inputs["launchType"] = args ? args.launchType : undefined;
            inputs["loadBalancers"] = args ? args.loadBalancers : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            inputs["orderedPlacementStrategies"] = args ? args.orderedPlacementStrategies : undefined;
            inputs["placementConstraints"] = args ? args.placementConstraints : undefined;
            inputs["platformVersion"] = args ? args.platformVersion : undefined;
            inputs["propagateTags"] = args ? args.propagateTags : undefined;
            inputs["schedulingStrategy"] = args ? args.schedulingStrategy : undefined;
            inputs["serviceRegistries"] = args ? args.serviceRegistries : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["taskDefinition"] = args ? args.taskDefinition : undefined;
            inputs["waitForSteadyState"] = args ? args.waitForSteadyState : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Service.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * ARN of an ECS cluster
     */
    readonly cluster?: pulumi.Input<string>;
    /**
     * Configuration block containing deployment controller configuration. Defined below.
     */
    readonly deploymentController?: pulumi.Input<inputs.ecs.ServiceDeploymentController>;
    /**
     * The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
     */
    readonly deploymentMaximumPercent?: pulumi.Input<number>;
    /**
     * The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
     */
    readonly deploymentMinimumHealthyPercent?: pulumi.Input<number>;
    /**
     * The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
     */
    readonly desiredCount?: pulumi.Input<number>;
    /**
     * Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
     */
    readonly enableEcsManagedTags?: pulumi.Input<boolean>;
    /**
     * Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
     */
    readonly healthCheckGracePeriodSeconds?: pulumi.Input<number>;
    /**
     * ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
     */
    readonly iamRole?: pulumi.Input<string>;
    /**
     * The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
     */
    readonly launchType?: pulumi.Input<string>;
    /**
     * A load balancer block. Load balancers documented below.
     */
    readonly loadBalancers?: pulumi.Input<pulumi.Input<inputs.ecs.ServiceLoadBalancer>[]>;
    /**
     * The name of the service (up to 255 letters, numbers, hyphens, and underscores)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
     */
    readonly networkConfiguration?: pulumi.Input<inputs.ecs.ServiceNetworkConfiguration>;
    /**
     * Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
     */
    readonly orderedPlacementStrategies?: pulumi.Input<pulumi.Input<inputs.ecs.ServiceOrderedPlacementStrategy>[]>;
    /**
     * rules that are taken into consideration during task placement. Maximum number of
     * `placementConstraints` is `10`. Defined below.
     */
    readonly placementConstraints?: pulumi.Input<pulumi.Input<inputs.ecs.ServicePlacementConstraint>[]>;
    /**
     * The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
     */
    readonly platformVersion?: pulumi.Input<string>;
    /**
     * Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
     */
    readonly propagateTags?: pulumi.Input<string>;
    /**
     * The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Fargate tasks do not support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html).
     */
    readonly schedulingStrategy?: pulumi.Input<string>;
    /**
     * The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
     */
    readonly serviceRegistries?: pulumi.Input<inputs.ecs.ServiceServiceRegistries>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
     */
    readonly taskDefinition?: pulumi.Input<string>;
    /**
     * If `true`, this provider will wait for the service to reach a steady state (like [`aws ecs wait services-stable`](https://docs.aws.amazon.com/cli/latest/reference/ecs/wait/services-stable.html)) before continuing. Default `false`.
     */
    readonly waitForSteadyState?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * ARN of an ECS cluster
     */
    readonly cluster?: pulumi.Input<string>;
    /**
     * Configuration block containing deployment controller configuration. Defined below.
     */
    readonly deploymentController?: pulumi.Input<inputs.ecs.ServiceDeploymentController>;
    /**
     * The upper limit (as a percentage of the service's desiredCount) of the number of running tasks that can be running in a service during a deployment. Not valid when using the `DAEMON` scheduling strategy.
     */
    readonly deploymentMaximumPercent?: pulumi.Input<number>;
    /**
     * The lower limit (as a percentage of the service's desiredCount) of the number of running tasks that must remain running and healthy in a service during a deployment.
     */
    readonly deploymentMinimumHealthyPercent?: pulumi.Input<number>;
    /**
     * The number of instances of the task definition to place and keep running. Defaults to 0. Do not specify if using the `DAEMON` scheduling strategy.
     */
    readonly desiredCount?: pulumi.Input<number>;
    /**
     * Specifies whether to enable Amazon ECS managed tags for the tasks within the service.
     */
    readonly enableEcsManagedTags?: pulumi.Input<boolean>;
    /**
     * Seconds to ignore failing load balancer health checks on newly instantiated tasks to prevent premature shutdown, up to 2147483647. Only valid for services configured to use load balancers.
     */
    readonly healthCheckGracePeriodSeconds?: pulumi.Input<number>;
    /**
     * ARN of the IAM role that allows Amazon ECS to make calls to your load balancer on your behalf. This parameter is required if you are using a load balancer with your service, but only if your task definition does not use the `awsvpc` network mode. If using `awsvpc` network mode, do not specify this role. If your account has already created the Amazon ECS service-linked role, that role is used by default for your service unless you specify a role here.
     */
    readonly iamRole?: pulumi.Input<string>;
    /**
     * The launch type on which to run your service. The valid values are `EC2` and `FARGATE`. Defaults to `EC2`.
     */
    readonly launchType?: pulumi.Input<string>;
    /**
     * A load balancer block. Load balancers documented below.
     */
    readonly loadBalancers?: pulumi.Input<pulumi.Input<inputs.ecs.ServiceLoadBalancer>[]>;
    /**
     * The name of the service (up to 255 letters, numbers, hyphens, and underscores)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The network configuration for the service. This parameter is required for task definitions that use the `awsvpc` network mode to receive their own Elastic Network Interface, and it is not supported for other network modes.
     */
    readonly networkConfiguration?: pulumi.Input<inputs.ecs.ServiceNetworkConfiguration>;
    /**
     * Service level strategy rules that are taken into consideration during task placement. List from top to bottom in order of precedence. The maximum number of `orderedPlacementStrategy` blocks is `5`. Defined below.
     */
    readonly orderedPlacementStrategies?: pulumi.Input<pulumi.Input<inputs.ecs.ServiceOrderedPlacementStrategy>[]>;
    /**
     * rules that are taken into consideration during task placement. Maximum number of
     * `placementConstraints` is `10`. Defined below.
     */
    readonly placementConstraints?: pulumi.Input<pulumi.Input<inputs.ecs.ServicePlacementConstraint>[]>;
    /**
     * The platform version on which to run your service. Only applicable for `launchType` set to `FARGATE`. Defaults to `LATEST`. More information about Fargate platform versions can be found in the [AWS ECS User Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html).
     */
    readonly platformVersion?: pulumi.Input<string>;
    /**
     * Specifies whether to propagate the tags from the task definition or the service to the tasks. The valid values are `SERVICE` and `TASK_DEFINITION`.
     */
    readonly propagateTags?: pulumi.Input<string>;
    /**
     * The scheduling strategy to use for the service. The valid values are `REPLICA` and `DAEMON`. Defaults to `REPLICA`. Note that [*Fargate tasks do not support the `DAEMON` scheduling strategy*](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/scheduling_tasks.html).
     */
    readonly schedulingStrategy?: pulumi.Input<string>;
    /**
     * The service discovery registries for the service. The maximum number of `serviceRegistries` blocks is `1`.
     */
    readonly serviceRegistries?: pulumi.Input<inputs.ecs.ServiceServiceRegistries>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The family and revision (`family:revision`) or full ARN of the task definition that you want to run in your service.
     */
    readonly taskDefinition: pulumi.Input<string>;
    /**
     * If `true`, this provider will wait for the service to reach a steady state (like [`aws ecs wait services-stable`](https://docs.aws.amazon.com/cli/latest/reference/ecs/wait/services-stable.html)) before continuing. Default `false`.
     */
    readonly waitForSteadyState?: pulumi.Input<boolean>;
}
