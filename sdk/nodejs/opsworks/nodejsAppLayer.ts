// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class NodejsAppLayer extends pulumi.CustomResource {
    /**
     * Get an existing NodejsAppLayer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NodejsAppLayerState, opts?: pulumi.CustomResourceOptions): NodejsAppLayer {
        return new NodejsAppLayer(name, <any>state, { ...opts, id: id });
    }

    public readonly autoAssignElasticIps: pulumi.Output<boolean | undefined>;
    public readonly autoAssignPublicIps: pulumi.Output<boolean | undefined>;
    public readonly autoHealing: pulumi.Output<boolean | undefined>;
    public readonly customConfigureRecipes: pulumi.Output<string[] | undefined>;
    public readonly customDeployRecipes: pulumi.Output<string[] | undefined>;
    public readonly customInstanceProfileArn: pulumi.Output<string | undefined>;
    public readonly customJson: pulumi.Output<string | undefined>;
    public readonly customSecurityGroupIds: pulumi.Output<string[] | undefined>;
    public readonly customSetupRecipes: pulumi.Output<string[] | undefined>;
    public readonly customShutdownRecipes: pulumi.Output<string[] | undefined>;
    public readonly customUndeployRecipes: pulumi.Output<string[] | undefined>;
    public readonly drainElbOnShutdown: pulumi.Output<boolean | undefined>;
    public readonly ebsVolumes: pulumi.Output<{ iops?: number, mountPoint: string, numberOfDisks: number, raidLevel?: string, size: number, type?: string }[] | undefined>;
    public readonly elasticLoadBalancer: pulumi.Output<string | undefined>;
    public readonly installUpdatesOnBoot: pulumi.Output<boolean | undefined>;
    public readonly instanceShutdownTimeout: pulumi.Output<number | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly nodejsVersion: pulumi.Output<string | undefined>;
    public readonly stackId: pulumi.Output<string>;
    public readonly systemPackages: pulumi.Output<string[] | undefined>;
    public readonly useEbsOptimizedInstances: pulumi.Output<boolean | undefined>;

    /**
     * Create a NodejsAppLayer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NodejsAppLayerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NodejsAppLayerArgs | NodejsAppLayerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: NodejsAppLayerState = argsOrState as NodejsAppLayerState | undefined;
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
            inputs["nodejsVersion"] = state ? state.nodejsVersion : undefined;
            inputs["stackId"] = state ? state.stackId : undefined;
            inputs["systemPackages"] = state ? state.systemPackages : undefined;
            inputs["useEbsOptimizedInstances"] = state ? state.useEbsOptimizedInstances : undefined;
        } else {
            const args = argsOrState as NodejsAppLayerArgs | undefined;
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
            inputs["nodejsVersion"] = args ? args.nodejsVersion : undefined;
            inputs["stackId"] = args ? args.stackId : undefined;
            inputs["systemPackages"] = args ? args.systemPackages : undefined;
            inputs["useEbsOptimizedInstances"] = args ? args.useEbsOptimizedInstances : undefined;
        }
        super("aws:opsworks/nodejsAppLayer:NodejsAppLayer", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NodejsAppLayer resources.
 */
export interface NodejsAppLayerState {
    readonly autoAssignElasticIps?: pulumi.Input<boolean>;
    readonly autoAssignPublicIps?: pulumi.Input<boolean>;
    readonly autoHealing?: pulumi.Input<boolean>;
    readonly customConfigureRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customDeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customInstanceProfileArn?: pulumi.Input<string>;
    readonly customJson?: pulumi.Input<string>;
    readonly customSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customSetupRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customShutdownRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customUndeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly drainElbOnShutdown?: pulumi.Input<boolean>;
    readonly ebsVolumes?: pulumi.Input<pulumi.Input<{ iops?: pulumi.Input<number>, mountPoint: pulumi.Input<string>, numberOfDisks: pulumi.Input<number>, raidLevel?: pulumi.Input<string>, size: pulumi.Input<number>, type?: pulumi.Input<string> }>[]>;
    readonly elasticLoadBalancer?: pulumi.Input<string>;
    readonly installUpdatesOnBoot?: pulumi.Input<boolean>;
    readonly instanceShutdownTimeout?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly nodejsVersion?: pulumi.Input<string>;
    readonly stackId?: pulumi.Input<string>;
    readonly systemPackages?: pulumi.Input<pulumi.Input<string>[]>;
    readonly useEbsOptimizedInstances?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a NodejsAppLayer resource.
 */
export interface NodejsAppLayerArgs {
    readonly autoAssignElasticIps?: pulumi.Input<boolean>;
    readonly autoAssignPublicIps?: pulumi.Input<boolean>;
    readonly autoHealing?: pulumi.Input<boolean>;
    readonly customConfigureRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customDeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customInstanceProfileArn?: pulumi.Input<string>;
    readonly customJson?: pulumi.Input<string>;
    readonly customSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customSetupRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customShutdownRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly customUndeployRecipes?: pulumi.Input<pulumi.Input<string>[]>;
    readonly drainElbOnShutdown?: pulumi.Input<boolean>;
    readonly ebsVolumes?: pulumi.Input<pulumi.Input<{ iops?: pulumi.Input<number>, mountPoint: pulumi.Input<string>, numberOfDisks: pulumi.Input<number>, raidLevel?: pulumi.Input<string>, size: pulumi.Input<number>, type?: pulumi.Input<string> }>[]>;
    readonly elasticLoadBalancer?: pulumi.Input<string>;
    readonly installUpdatesOnBoot?: pulumi.Input<boolean>;
    readonly instanceShutdownTimeout?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly nodejsVersion?: pulumi.Input<string>;
    readonly stackId: pulumi.Input<string>;
    readonly systemPackages?: pulumi.Input<pulumi.Input<string>[]>;
    readonly useEbsOptimizedInstances?: pulumi.Input<boolean>;
}
