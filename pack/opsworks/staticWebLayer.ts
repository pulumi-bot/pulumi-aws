// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class StaticWebLayer extends fabric.Resource {
    public readonly autoAssignElasticIps?: fabric.Computed<boolean>;
    public readonly autoAssignPublicIps?: fabric.Computed<boolean>;
    public readonly autoHealing?: fabric.Computed<boolean>;
    public readonly customConfigureRecipes?: fabric.Computed<string[]>;
    public readonly customDeployRecipes?: fabric.Computed<string[]>;
    public readonly customInstanceProfileArn?: fabric.Computed<string>;
    public readonly customJson?: fabric.Computed<string>;
    public readonly customSecurityGroupIds?: fabric.Computed<string[]>;
    public readonly customSetupRecipes?: fabric.Computed<string[]>;
    public readonly customShutdownRecipes?: fabric.Computed<string[]>;
    public readonly customUndeployRecipes?: fabric.Computed<string[]>;
    public readonly drainElbOnShutdown?: fabric.Computed<boolean>;
    public readonly ebsVolume?: fabric.Computed<{ iops?: number, mountPoint: string, numberOfDisks: number, raidLevel?: string, size: number, type?: string }[]>;
    public readonly elasticLoadBalancer?: fabric.Computed<string>;
    public /*out*/ readonly layerId: fabric.Computed<string>;
    public readonly installUpdatesOnBoot?: fabric.Computed<boolean>;
    public readonly instanceShutdownTimeout?: fabric.Computed<number>;
    public readonly name: fabric.Computed<string>;
    public readonly stackId: fabric.Computed<string>;
    public readonly systemPackages?: fabric.Computed<string[]>;
    public readonly useEbsOptimizedInstances?: fabric.Computed<boolean>;

    constructor(urnName: string, args: StaticWebLayerArgs, dependsOn?: fabric.Resource[]) {
        if (args.stackId === undefined) {
            throw new Error("Missing required property 'stackId'");
        }
        super("aws:opsworks/staticWebLayer:StaticWebLayer", urnName, {
            "autoAssignElasticIps": args.autoAssignElasticIps,
            "autoAssignPublicIps": args.autoAssignPublicIps,
            "autoHealing": args.autoHealing,
            "customConfigureRecipes": args.customConfigureRecipes,
            "customDeployRecipes": args.customDeployRecipes,
            "customInstanceProfileArn": args.customInstanceProfileArn,
            "customJson": args.customJson,
            "customSecurityGroupIds": args.customSecurityGroupIds,
            "customSetupRecipes": args.customSetupRecipes,
            "customShutdownRecipes": args.customShutdownRecipes,
            "customUndeployRecipes": args.customUndeployRecipes,
            "drainElbOnShutdown": args.drainElbOnShutdown,
            "ebsVolume": args.ebsVolume,
            "elasticLoadBalancer": args.elasticLoadBalancer,
            "installUpdatesOnBoot": args.installUpdatesOnBoot,
            "instanceShutdownTimeout": args.instanceShutdownTimeout,
            "name": args.name,
            "stackId": args.stackId,
            "systemPackages": args.systemPackages,
            "useEbsOptimizedInstances": args.useEbsOptimizedInstances,
            "layerId": undefined,
        }, dependsOn);
    }
}

export interface StaticWebLayerArgs {
    readonly autoAssignElasticIps?: fabric.ComputedValue<boolean>;
    readonly autoAssignPublicIps?: fabric.ComputedValue<boolean>;
    readonly autoHealing?: fabric.ComputedValue<boolean>;
    readonly customConfigureRecipes?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly customDeployRecipes?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly customInstanceProfileArn?: fabric.ComputedValue<string>;
    readonly customJson?: fabric.ComputedValue<string>;
    readonly customSecurityGroupIds?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly customSetupRecipes?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly customShutdownRecipes?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly customUndeployRecipes?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly drainElbOnShutdown?: fabric.ComputedValue<boolean>;
    readonly ebsVolume?: fabric.ComputedValue<{ iops?: fabric.ComputedValue<number>, mountPoint: fabric.ComputedValue<string>, numberOfDisks: fabric.ComputedValue<number>, raidLevel?: fabric.ComputedValue<string>, size: fabric.ComputedValue<number>, type?: fabric.ComputedValue<string> }>[];
    readonly elasticLoadBalancer?: fabric.ComputedValue<string>;
    readonly installUpdatesOnBoot?: fabric.ComputedValue<boolean>;
    readonly instanceShutdownTimeout?: fabric.ComputedValue<number>;
    readonly name?: fabric.ComputedValue<string>;
    readonly stackId: fabric.ComputedValue<string>;
    readonly systemPackages?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly useEbsOptimizedInstances?: fabric.ComputedValue<boolean>;
}

