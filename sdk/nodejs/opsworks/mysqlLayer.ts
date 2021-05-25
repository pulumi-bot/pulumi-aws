// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides an OpsWorks MySQL layer resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const db = new aws.opsworks.MysqlLayer("db", {stackId: aws_opsworks_stack.main.id});
 * ```
 */
export class MysqlLayer extends pulumi.CustomResource {
    /**
     * Get an existing MysqlLayer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MysqlLayerState, opts?: pulumi.CustomResourceOptions): MysqlLayer {
        return new MysqlLayer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:opsworks/mysqlLayer:MysqlLayer';

    /**
     * Returns true if the given object is an instance of MysqlLayer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MysqlLayer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MysqlLayer.__pulumiType;
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
    public readonly ebsVolumes!: pulumi.Output<outputs.opsworks.MysqlLayerEbsVolume[] | undefined>;
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
     * Root password to use for MySQL.
     */
    public readonly rootPassword!: pulumi.Output<string | undefined>;
    /**
     * Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
     */
    public readonly rootPasswordOnAllInstances!: pulumi.Output<boolean | undefined>;
    /**
     * The id of the stack the layer will belong to.
     */
    public readonly stackId!: pulumi.Output<string>;
    /**
     * Names of a set of system packages to install on the layer's instances.
     */
    public readonly systemPackages!: pulumi.Output<string[] | undefined>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;
    /**
     * Whether to use EBS-optimized instances.
     */
    public readonly useEbsOptimizedInstances!: pulumi.Output<boolean | undefined>;

    /**
     * Create a MysqlLayer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MysqlLayerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MysqlLayerArgs | MysqlLayerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MysqlLayerState | undefined;
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
            inputs["rootPassword"] = state ? state.rootPassword : undefined;
            inputs["rootPasswordOnAllInstances"] = state ? state.rootPasswordOnAllInstances : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["systemPackages"] = state ? state.systemPackages : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
            inputs["useEbsOptimizedInstances"] = state ? state.useEbsOptimizedInstances : undefined;
        } else {
            const args = argsOrState as MysqlLayerArgs | undefined;
            if ((!args || args.stackId === undefined) && !opts.urn) {
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
            inputs["rootPassword"] = args ? args.rootPassword : undefined;
            inputs["rootPasswordOnAllInstances"] = args ? args.rootPasswordOnAllInstances : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["systemPackages"] = args ? args.systemPackages : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["useEbsOptimizedInstances"] = args ? args.useEbsOptimizedInstances : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MysqlLayer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MysqlLayer resources.
 */
export interface MysqlLayerState {
    /**
     * The Amazon Resource Name(ARN) of the layer.
     */
    arn?: pulumi.Input<string>;
    /**
     * Whether to automatically assign an elastic IP address to the layer's instances.
     */
    autoAssignElasticIps?: pulumi.Input<boolean>;
    /**
     * For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
     */
    autoAssignPublicIps?: pulumi.Input<boolean>;
    /**
     * Whether to enable auto-healing for the layer.
     */
    autoHealing?: pulumi.Input<boolean>;
    customConfigureRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    customDeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of an IAM profile that will be used for the layer's instances.
     */
    customInstanceProfileArn?: pulumi.Input<string>;
    /**
     * Custom JSON attributes to apply to the layer.
     */
    customJson?: pulumi.Input<string>;
    /**
     * Ids for a set of security groups to apply to the layer's instances.
     */
    customSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    customSetupRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    customShutdownRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    customUndeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable Elastic Load Balancing connection draining.
     */
    drainElbOnShutdown?: pulumi.Input<boolean>;
    /**
     * `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
     */
    ebsVolumes?: pulumi.Input<pulumi.Input<inputs.opsworks.MysqlLayerEbsVolume>[]>;
    /**
     * Name of an Elastic Load Balancer to attach to this layer
     */
    elasticLoadBalancer?: pulumi.Input<string>;
    /**
     * Whether to install OS and package updates on each instance when it boots.
     */
    installUpdatesOnBoot?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
     */
    instanceShutdownTimeout?: pulumi.Input<number>;
    /**
     * A human-readable name for the layer.
     */
    name?: pulumi.Input<string>;
    /**
     * Root password to use for MySQL.
     */
    rootPassword?: pulumi.Input<string>;
    /**
     * Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
     */
    rootPasswordOnAllInstances?: pulumi.Input<boolean>;
    /**
     * The id of the stack the layer will belong to.
     */
    stackId?: pulumi.Input<string>;
    /**
     * Names of a set of system packages to install on the layer's instances.
     */
    systemPackages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to use EBS-optimized instances.
     */
    useEbsOptimizedInstances?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a MysqlLayer resource.
 */
export interface MysqlLayerArgs {
    /**
     * Whether to automatically assign an elastic IP address to the layer's instances.
     */
    autoAssignElasticIps?: pulumi.Input<boolean>;
    /**
     * For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
     */
    autoAssignPublicIps?: pulumi.Input<boolean>;
    /**
     * Whether to enable auto-healing for the layer.
     */
    autoHealing?: pulumi.Input<boolean>;
    customConfigureRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    customDeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ARN of an IAM profile that will be used for the layer's instances.
     */
    customInstanceProfileArn?: pulumi.Input<string>;
    /**
     * Custom JSON attributes to apply to the layer.
     */
    customJson?: pulumi.Input<string>;
    /**
     * Ids for a set of security groups to apply to the layer's instances.
     */
    customSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    customSetupRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    customShutdownRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    customUndeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to enable Elastic Load Balancing connection draining.
     */
    drainElbOnShutdown?: pulumi.Input<boolean>;
    /**
     * `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
     */
    ebsVolumes?: pulumi.Input<pulumi.Input<inputs.opsworks.MysqlLayerEbsVolume>[]>;
    /**
     * Name of an Elastic Load Balancer to attach to this layer
     */
    elasticLoadBalancer?: pulumi.Input<string>;
    /**
     * Whether to install OS and package updates on each instance when it boots.
     */
    installUpdatesOnBoot?: pulumi.Input<boolean>;
    /**
     * The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
     */
    instanceShutdownTimeout?: pulumi.Input<number>;
    /**
     * A human-readable name for the layer.
     */
    name?: pulumi.Input<string>;
    /**
     * Root password to use for MySQL.
     */
    rootPassword?: pulumi.Input<string>;
    /**
     * Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
     */
    rootPasswordOnAllInstances?: pulumi.Input<boolean>;
    /**
     * The id of the stack the layer will belong to.
     */
    stackId: pulumi.Input<string>;
    /**
     * Names of a set of system packages to install on the layer's instances.
     */
    systemPackages?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to use EBS-optimized instances.
     */
    useEbsOptimizedInstances?: pulumi.Input<boolean>;
}
