// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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

    public readonly cluster: pulumi.Output<string>;
    public readonly deploymentController: pulumi.Output<{ type?: string } | undefined>;
    public readonly deploymentMaximumPercent: pulumi.Output<number | undefined>;
    public readonly deploymentMinimumHealthyPercent: pulumi.Output<number | undefined>;
    public readonly desiredCount: pulumi.Output<number | undefined>;
    public readonly enableEcsManagedTags: pulumi.Output<boolean | undefined>;
    public readonly healthCheckGracePeriodSeconds: pulumi.Output<number | undefined>;
    public readonly iamRole: pulumi.Output<string>;
    public readonly launchType: pulumi.Output<string | undefined>;
    public readonly loadBalancers: pulumi.Output<{ containerName: string, containerPort: number, elbName?: string, targetGroupArn?: string }[] | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly networkConfiguration: pulumi.Output<{ assignPublicIp?: boolean, securityGroups?: string[], subnets: string[] } | undefined>;
    public readonly orderedPlacementStrategies: pulumi.Output<{ field?: string, type: string }[] | undefined>;
    public readonly placementConstraints: pulumi.Output<{ expression?: string, type: string }[] | undefined>;
    public readonly placementStrategies: pulumi.Output<{ field?: string, type: string }[] | undefined>;
    public readonly platformVersion: pulumi.Output<string>;
    public readonly propagateTags: pulumi.Output<string | undefined>;
    public readonly schedulingStrategy: pulumi.Output<string | undefined>;
    public readonly serviceRegistries: pulumi.Output<{ containerName?: string, containerPort?: number, port?: number, registryArn: string } | undefined>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly taskDefinition: pulumi.Output<string>;
    public readonly waitForSteadyState: pulumi.Output<boolean | undefined>;

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
            const state: ServiceState = argsOrState as ServiceState | undefined;
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
            inputs["placementStrategies"] = state ? state.placementStrategies : undefined;
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
            inputs["placementStrategies"] = args ? args.placementStrategies : undefined;
            inputs["platformVersion"] = args ? args.platformVersion : undefined;
            inputs["propagateTags"] = args ? args.propagateTags : undefined;
            inputs["schedulingStrategy"] = args ? args.schedulingStrategy : undefined;
            inputs["serviceRegistries"] = args ? args.serviceRegistries : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["taskDefinition"] = args ? args.taskDefinition : undefined;
            inputs["waitForSteadyState"] = args ? args.waitForSteadyState : undefined;
        }
        super("aws:ecs/service:Service", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    readonly cluster?: pulumi.Input<string>;
    readonly deploymentController?: pulumi.Input<{ type?: pulumi.Input<string> }>;
    readonly deploymentMaximumPercent?: pulumi.Input<number>;
    readonly deploymentMinimumHealthyPercent?: pulumi.Input<number>;
    readonly desiredCount?: pulumi.Input<number>;
    readonly enableEcsManagedTags?: pulumi.Input<boolean>;
    readonly healthCheckGracePeriodSeconds?: pulumi.Input<number>;
    readonly iamRole?: pulumi.Input<string>;
    readonly launchType?: pulumi.Input<string>;
    readonly loadBalancers?: pulumi.Input<pulumi.Input<{ containerName: pulumi.Input<string>, containerPort: pulumi.Input<number>, elbName?: pulumi.Input<string>, targetGroupArn?: pulumi.Input<string> }>[]>;
    readonly name?: pulumi.Input<string>;
    readonly networkConfiguration?: pulumi.Input<{ assignPublicIp?: pulumi.Input<boolean>, securityGroups?: pulumi.Input<pulumi.Input<string>[]>, subnets: pulumi.Input<pulumi.Input<string>[]> }>;
    readonly orderedPlacementStrategies?: pulumi.Input<pulumi.Input<{ field?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    readonly placementConstraints?: pulumi.Input<pulumi.Input<{ expression?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    readonly placementStrategies?: pulumi.Input<pulumi.Input<{ field?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    readonly platformVersion?: pulumi.Input<string>;
    readonly propagateTags?: pulumi.Input<string>;
    readonly schedulingStrategy?: pulumi.Input<string>;
    readonly serviceRegistries?: pulumi.Input<{ containerName?: pulumi.Input<string>, containerPort?: pulumi.Input<number>, port?: pulumi.Input<number>, registryArn: pulumi.Input<string> }>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly taskDefinition?: pulumi.Input<string>;
    readonly waitForSteadyState?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    readonly cluster?: pulumi.Input<string>;
    readonly deploymentController?: pulumi.Input<{ type?: pulumi.Input<string> }>;
    readonly deploymentMaximumPercent?: pulumi.Input<number>;
    readonly deploymentMinimumHealthyPercent?: pulumi.Input<number>;
    readonly desiredCount?: pulumi.Input<number>;
    readonly enableEcsManagedTags?: pulumi.Input<boolean>;
    readonly healthCheckGracePeriodSeconds?: pulumi.Input<number>;
    readonly iamRole?: pulumi.Input<string>;
    readonly launchType?: pulumi.Input<string>;
    readonly loadBalancers?: pulumi.Input<pulumi.Input<{ containerName: pulumi.Input<string>, containerPort: pulumi.Input<number>, elbName?: pulumi.Input<string>, targetGroupArn?: pulumi.Input<string> }>[]>;
    readonly name?: pulumi.Input<string>;
    readonly networkConfiguration?: pulumi.Input<{ assignPublicIp?: pulumi.Input<boolean>, securityGroups?: pulumi.Input<pulumi.Input<string>[]>, subnets: pulumi.Input<pulumi.Input<string>[]> }>;
    readonly orderedPlacementStrategies?: pulumi.Input<pulumi.Input<{ field?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    readonly placementConstraints?: pulumi.Input<pulumi.Input<{ expression?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    readonly placementStrategies?: pulumi.Input<pulumi.Input<{ field?: pulumi.Input<string>, type: pulumi.Input<string> }>[]>;
    readonly platformVersion?: pulumi.Input<string>;
    readonly propagateTags?: pulumi.Input<string>;
    readonly schedulingStrategy?: pulumi.Input<string>;
    readonly serviceRegistries?: pulumi.Input<{ containerName?: pulumi.Input<string>, containerPort?: pulumi.Input<number>, port?: pulumi.Input<number>, registryArn: pulumi.Input<string> }>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly taskDefinition: pulumi.Input<string>;
    readonly waitForSteadyState?: pulumi.Input<boolean>;
}
