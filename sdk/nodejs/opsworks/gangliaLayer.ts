// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides an OpsWorks Ganglia layer resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const monitor = new aws.opsworks.GangliaLayer("monitor", {
 *     stackId: aws_opsworks_stack.main.id,
 *     password: "foobarbaz",
 * });
 * ```
 */
export class GangliaLayer extends pulumi.CustomResource {
    /**
     * Get an existing GangliaLayer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GangliaLayerState, opts?: pulumi.CustomResourceOptions): GangliaLayer {
        return new GangliaLayer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opsworks/gangliaLayer:GangliaLayer';

    /**
     * Returns true if the given object is an instance of GangliaLayer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GangliaLayer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GangliaLayer.__pulumiType;
    }

    /**
     * The Amazon Resource Name(ARN) of the layer.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Whether to automatically assign an elastic IP address to the layer's instances.
     */
    public readonly autoAssignElasticIps!: pulumi.Output<boolean | undefined>;
    /**
     * For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
     */
    public readonly autoAssignPublicIps!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to enable auto-healing for the layer.
     */
    public readonly autoHealing!: pulumi.Output<boolean | undefined>;
    public readonly customConfigureRecipes!: pulumi.Output<string[] | undefined>;
    public readonly customDeployRecipes!: pulumi.Output<string[] | undefined>;
    /**
     * The ARN of an IAM profile that will be used for the layer's instances.
     */
    public readonly customInstanceProfileArn!: pulumi.Output<string | undefined>;
    /**
     * Custom JSON attributes to apply to the layer.
     */
    public readonly customJson!: pulumi.Output<string | undefined>;
    /**
     * Ids for a set of security groups to apply to the layer's instances.
     */
    public readonly customSecurityGroupIds!: pulumi.Output<string[] | undefined>;
    public readonly customSetupRecipes!: pulumi.Output<string[] | undefined>;
    public readonly customShutdownRecipes!: pulumi.Output<string[] | undefined>;
    public readonly customUndeployRecipes!: pulumi.Output<string[] | undefined>;
    /**
     * Whether to enable Elastic Load Balancing connection draining.
     */
    public readonly drainElbOnShutdown!: pulumi.Output<boolean | undefined>;
    /**
     * `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
     */
    public readonly ebsVolumes!: pulumi.Output<outputs.opsworks.GangliaLayerEbsVolume[] | undefined>;
    /**
     * Name of an Elastic Load Balancer to attach to this layer
     */
    public readonly elasticLoadBalancer!: pulumi.Output<string | undefined>;
    /**
     * Whether to install OS and package updates on each instance when it boots.
     */
    public readonly installUpdatesOnBoot!: pulumi.Output<boolean | undefined>;
    /**
     * The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
     */
    public readonly instanceShutdownTimeout!: pulumi.Output<number | undefined>;
    /**
     * A human-readable name for the layer.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The password to use for Ganglia.
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * The id of the stack the layer will belong to.
     */
    public readonly stackId!: pulumi.Output<string>;
    /**
     * Names of a set of system packages to install on the layer's instances.
     */
    public readonly systemPackages!: pulumi.Output<string[] | undefined>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The URL path to use for Ganglia. Defaults to "/ganglia".
     */
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * Whether to use EBS-optimized instances.
     */
    public readonly useEbsOptimizedInstances!: pulumi.Output<boolean | undefined>;
    /**
     * The username to use for Ganglia. Defaults to "opsworks".
     */
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a GangliaLayer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GangliaLayerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GangliaLayerArgs | GangliaLayerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GangliaLayerState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoAssignElasticIps"] = state ? state.autoAssignElasticIps : undefined;
            inputs["autoAssignPublicIps"] = state ? state.autoAssignPublicIps : undefined;
            inputs["autoHealing"] = state ? state.autoHealing : undefined;
            inputs["customConfigureRecipes"] = state ? state.customConfigureRecipes : undefined;
            inputs["customDeployRecipes"] = state ? state.customDeployRecipes : undefined;
            inputs["customInstanceProfileArn"] = state ? state.customInstanceProfileArn : undefined;
            inputs["customJson"] = state ? state.customJson : undefined;
            inputs["customSecurityGroupIds"] = state ? state.customSecurityGroupIds : undefined;
            inputs["customSetupRecipes"] = state ? state.customSetupRecipes : undefined;
            inputs["customShutdownRecipes"] = state ? state.customShutdownRecipes : undefined;
            inputs["customUndeployRecipes"] = state ? state.customUndeployRecipes : undefined;
            inputs["drainElbOnShutdown"] = state ? state.drainElbOnShutdown : undefined;
            inputs["ebsVolumes"] = state ? state.ebsVolumes : undefined;
            inputs["elasticLoadBalancer"] = state ? state.elasticLoadBalancer : undefined;
            inputs["installUpdatesOnBoot"] = state ? state.installUpdatesOnBoot : undefined;
            inputs["instanceShutdownTimeout"] = state ? state.instanceShutdownTimeout : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["systemPackages"] = state ? state.systemPackages : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["useEbsOptimizedInstances"] = state ? state.useEbsOptimizedInstances : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as GangliaLayerArgs | undefined;
            if (!args || args.password === undefined) {
                throw new Error("Missing required property 'password'");
            }
            if (!args || args.stackId === undefined) {
                throw new Error("Missing required property 'stackId'");
            }
            inputs["autoAssignElasticIps"] = args ? args.autoAssignElasticIps : undefined;
            inputs["autoAssignPublicIps"] = args ? args.autoAssignPublicIps : undefined;
            inputs["autoHealing"] = args ? args.autoHealing : undefined;
            inputs["customConfigureRecipes"] = args ? args.customConfigureRecipes : undefined;
            inputs["customDeployRecipes"] = args ? args.customDeployRecipes : undefined;
            inputs["customInstanceProfileArn"] = args ? args.customInstanceProfileArn : undefined;
            inputs["customJson"] = args ? args.customJson : undefined;
            inputs["customSecurityGroupIds"] = args ? args.customSecurityGroupIds : undefined;
            inputs["customSetupRecipes"] = args ? args.customSetupRecipes : undefined;
            inputs["customShutdownRecipes"] = args ? args.customShutdownRecipes : undefined;
            inputs["customUndeployRecipes"] = args ? args.customUndeployRecipes : undefined;
            inputs["drainElbOnShutdown"] = args ? args.drainElbOnShutdown : undefined;
            inputs["ebsVolumes"] = args ? args.ebsVolumes : undefined;
            inputs["elasticLoadBalancer"] = args ? args.elasticLoadBalancer : undefined;
            inputs["installUpdatesOnBoot"] = args ? args.installUpdatesOnBoot : undefined;
            inputs["instanceShutdownTimeout"] = args ? args.instanceShutdownTimeout : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["password"] = args ? args.password : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["systemPackages"] = args ? args.systemPackages : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["useEbsOptimizedInstances"] = args ? args.useEbsOptimizedInstances : undefined;
            inputs["username"] = args ? args.username : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GangliaLayer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GangliaLayer resources.
 */
export interface GangliaLayerState {
    /**
     * The Amazon Resource Name(ARN) of the layer.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Whether to automatically assign an elastic IP address to the layer's instances.
     */
    readonly autoAssignElasticIps?: pulumi.Input<boolean>;
    /**
     * For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
     */
    readonly autoAssignPublicIps?: pulumi.Input<boolean>;
    /**
     * Whether to enable auto-healing for the layer.
     */
    readonly autoHealing?: pulumi.Input<boolean>;
    readonly customConfigureRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customDeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of an IAM profile that will be used for the layer's instances.
     */
    readonly customInstanceProfileArn?: pulumi.Input<string>;
    /**
     * Custom JSON attributes to apply to the layer.
     */
    readonly customJson?: pulumi.Input<string>;
    /**
     * Ids for a set of security groups to apply to the layer's instances.
     */
    readonly customSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customSetupRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customShutdownRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customUndeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable Elastic Load Balancing connection draining.
     */
    readonly drainElbOnShutdown?: pulumi.Input<boolean>;
    /**
     * `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
     */
    readonly ebsVolumes?: pulumi.Input<pulumi.Input<inputs.opsworks.GangliaLayerEbsVolume>[]>;
    /**
     * Name of an Elastic Load Balancer to attach to this layer
     */
    readonly elasticLoadBalancer?: pulumi.Input<string>;
    /**
     * Whether to install OS and package updates on each instance when it boots.
     */
    readonly installUpdatesOnBoot?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
     */
    readonly instanceShutdownTimeout?: pulumi.Input<number>;
    /**
     * A human-readable name for the layer.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password to use for Ganglia.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * The id of the stack the layer will belong to.
     */
    readonly stackId?: pulumi.Input<string>;
    /**
     * Names of a set of system packages to install on the layer's instances.
     */
    readonly systemPackages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The URL path to use for Ganglia. Defaults to "/ganglia".
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Whether to use EBS-optimized instances.
     */
    readonly useEbsOptimizedInstances?: pulumi.Input<boolean>;
    /**
     * The username to use for Ganglia. Defaults to "opsworks".
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GangliaLayer resource.
 */
export interface GangliaLayerArgs {
    /**
     * Whether to automatically assign an elastic IP address to the layer's instances.
     */
    readonly autoAssignElasticIps?: pulumi.Input<boolean>;
    /**
     * For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
     */
    readonly autoAssignPublicIps?: pulumi.Input<boolean>;
    /**
     * Whether to enable auto-healing for the layer.
     */
    readonly autoHealing?: pulumi.Input<boolean>;
    readonly customConfigureRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customDeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of an IAM profile that will be used for the layer's instances.
     */
    readonly customInstanceProfileArn?: pulumi.Input<string>;
    /**
     * Custom JSON attributes to apply to the layer.
     */
    readonly customJson?: pulumi.Input<string>;
    /**
     * Ids for a set of security groups to apply to the layer's instances.
     */
    readonly customSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customSetupRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customShutdownRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customUndeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable Elastic Load Balancing connection draining.
     */
    readonly drainElbOnShutdown?: pulumi.Input<boolean>;
    /**
     * `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
     */
    readonly ebsVolumes?: pulumi.Input<pulumi.Input<inputs.opsworks.GangliaLayerEbsVolume>[]>;
    /**
     * Name of an Elastic Load Balancer to attach to this layer
     */
    readonly elasticLoadBalancer?: pulumi.Input<string>;
    /**
     * Whether to install OS and package updates on each instance when it boots.
     */
    readonly installUpdatesOnBoot?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
     */
    readonly instanceShutdownTimeout?: pulumi.Input<number>;
    /**
     * A human-readable name for the layer.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The password to use for Ganglia.
     */
    readonly password: pulumi.Input<string>;
    /**
     * The id of the stack the layer will belong to.
     */
    readonly stackId: pulumi.Input<string>;
    /**
     * Names of a set of system packages to install on the layer's instances.
     */
    readonly systemPackages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The URL path to use for Ganglia. Defaults to "/ganglia".
     */
    readonly url?: pulumi.Input<string>;
    /**
     * Whether to use EBS-optimized instances.
     */
    readonly useEbsOptimizedInstances?: pulumi.Input<boolean>;
    /**
     * The username to use for Ganglia. Defaults to "opsworks".
     */
    readonly username?: pulumi.Input<string>;
}
