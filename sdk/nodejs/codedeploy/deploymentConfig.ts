// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a CodeDeploy deployment config for an application
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codedeploy_deployment_config.html.markdown.
 */
export class DeploymentConfig extends pulumi.CustomResource {
    /**
     * Get an existing DeploymentConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeploymentConfigState, opts?: pulumi.CustomResourceOptions): DeploymentConfig {
        return new DeploymentConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codedeploy/deploymentConfig:DeploymentConfig';

    /**
     * Returns true if the given object is an instance of DeploymentConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeploymentConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeploymentConfig.__pulumiType;
    }

    /**
     * The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
     */
    public readonly computePlatform!: pulumi.Output<string | undefined>;
    /**
     * The AWS Assigned deployment config id
     */
    public /*out*/ readonly deploymentConfigId!: pulumi.Output<string>;
    /**
     * The name of the deployment config.
     */
    public readonly deploymentConfigName!: pulumi.Output<string>;
    /**
     * A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
     */
    public readonly minimumHealthyHosts!: pulumi.Output<outputs.codedeploy.DeploymentConfigMinimumHealthyHosts | undefined>;
    /**
     * A trafficRoutingConfig block. Traffic Routing Config is documented below.
     */
    public readonly trafficRoutingConfig!: pulumi.Output<outputs.codedeploy.DeploymentConfigTrafficRoutingConfig | undefined>;

    /**
     * Create a DeploymentConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeploymentConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeploymentConfigArgs | DeploymentConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DeploymentConfigState | undefined;
            inputs["computePlatform"] = state ? state.computePlatform : undefined;
            inputs["deploymentConfigId"] = state ? state.deploymentConfigId : undefined;
            inputs["deploymentConfigName"] = state ? state.deploymentConfigName : undefined;
            inputs["minimumHealthyHosts"] = state ? state.minimumHealthyHosts : undefined;
            inputs["trafficRoutingConfig"] = state ? state.trafficRoutingConfig : undefined;
        } else {
            const args = argsOrState as DeploymentConfigArgs | undefined;
            if (!args || args.deploymentConfigName === undefined) {
                throw new Error("Missing required property 'deploymentConfigName'");
            }
            inputs["computePlatform"] = args ? args.computePlatform : undefined;
            inputs["deploymentConfigName"] = args ? args.deploymentConfigName : undefined;
            inputs["minimumHealthyHosts"] = args ? args.minimumHealthyHosts : undefined;
            inputs["trafficRoutingConfig"] = args ? args.trafficRoutingConfig : undefined;
            inputs["deploymentConfigId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DeploymentConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeploymentConfig resources.
 */
export interface DeploymentConfigState {
    /**
     * The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
     */
    readonly computePlatform?: pulumi.Input<string>;
    /**
     * The AWS Assigned deployment config id
     */
    readonly deploymentConfigId?: pulumi.Input<string>;
    /**
     * The name of the deployment config.
     */
    readonly deploymentConfigName?: pulumi.Input<string>;
    /**
     * A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
     */
    readonly minimumHealthyHosts?: pulumi.Input<inputs.codedeploy.DeploymentConfigMinimumHealthyHosts>;
    /**
     * A trafficRoutingConfig block. Traffic Routing Config is documented below.
     */
    readonly trafficRoutingConfig?: pulumi.Input<inputs.codedeploy.DeploymentConfigTrafficRoutingConfig>;
}

/**
 * The set of arguments for constructing a DeploymentConfig resource.
 */
export interface DeploymentConfigArgs {
    /**
     * The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
     */
    readonly computePlatform?: pulumi.Input<string>;
    /**
     * The name of the deployment config.
     */
    readonly deploymentConfigName: pulumi.Input<string>;
    /**
     * A minimumHealthyHosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
     */
    readonly minimumHealthyHosts?: pulumi.Input<inputs.codedeploy.DeploymentConfigMinimumHealthyHosts>;
    /**
     * A trafficRoutingConfig block. Traffic Routing Config is documented below.
     */
    readonly trafficRoutingConfig?: pulumi.Input<inputs.codedeploy.DeploymentConfigTrafficRoutingConfig>;
}
