// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

import {IpAddressType, LoadBalancerType} from "../alb";

/**
 * @deprecated aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer
 */
export class LoadBalancer extends pulumi.CustomResource {
    /**
     * Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoadBalancerState, opts?: pulumi.CustomResourceOptions): LoadBalancer {
        pulumi.log.warn("LoadBalancer is deprecated: aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer")
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:applicationloadbalancing/loadBalancer:LoadBalancer';

    /**
     * Returns true if the given object is an instance of LoadBalancer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoadBalancer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoadBalancer.__pulumiType;
    }

    public readonly accessLogs!: pulumi.Output<outputs.applicationloadbalancing.LoadBalancerAccessLogs | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly arnSuffix!: pulumi.Output<string>;
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    public readonly dropInvalidHeaderFields!: pulumi.Output<boolean | undefined>;
    public readonly enableCrossZoneLoadBalancing!: pulumi.Output<boolean | undefined>;
    public readonly enableDeletionProtection!: pulumi.Output<boolean | undefined>;
    public readonly enableHttp2!: pulumi.Output<boolean | undefined>;
    public readonly idleTimeout!: pulumi.Output<number | undefined>;
    public readonly internal!: pulumi.Output<boolean>;
    public readonly ipAddressType!: pulumi.Output<IpAddressType>;
    public readonly loadBalancerType!: pulumi.Output<LoadBalancerType | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    public readonly securityGroups!: pulumi.Output<string[]>;
    public readonly subnetMappings!: pulumi.Output<outputs.applicationloadbalancing.LoadBalancerSubnetMapping[]>;
    public readonly subnets!: pulumi.Output<string[]>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    public /*out*/ readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer */
    constructor(name: string, args?: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    /** @deprecated aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer */
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("LoadBalancer is deprecated: aws.applicationloadbalancing.LoadBalancer has been deprecated in favor of aws.alb.LoadBalancer")
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LoadBalancerState | undefined;
            inputs["accessLogs"] = state ? state.accessLogs : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["arnSuffix"] = state ? state.arnSuffix : undefined;
            inputs["dnsName"] = state ? state.dnsName : undefined;
            inputs["dropInvalidHeaderFields"] = state ? state.dropInvalidHeaderFields : undefined;
            inputs["enableCrossZoneLoadBalancing"] = state ? state.enableCrossZoneLoadBalancing : undefined;
            inputs["enableDeletionProtection"] = state ? state.enableDeletionProtection : undefined;
            inputs["enableHttp2"] = state ? state.enableHttp2 : undefined;
            inputs["idleTimeout"] = state ? state.idleTimeout : undefined;
            inputs["internal"] = state ? state.internal : undefined;
            inputs["ipAddressType"] = state ? state.ipAddressType : undefined;
            inputs["loadBalancerType"] = state ? state.loadBalancerType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["subnetMappings"] = state ? state.subnetMappings : undefined;
            inputs["subnets"] = state ? state.subnets : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as LoadBalancerArgs | undefined;
            inputs["accessLogs"] = args ? args.accessLogs : undefined;
            inputs["dropInvalidHeaderFields"] = args ? args.dropInvalidHeaderFields : undefined;
            inputs["enableCrossZoneLoadBalancing"] = args ? args.enableCrossZoneLoadBalancing : undefined;
            inputs["enableDeletionProtection"] = args ? args.enableDeletionProtection : undefined;
            inputs["enableHttp2"] = args ? args.enableHttp2 : undefined;
            inputs["idleTimeout"] = args ? args.idleTimeout : undefined;
            inputs["internal"] = args ? args.internal : undefined;
            inputs["ipAddressType"] = args ? args.ipAddressType : undefined;
            inputs["loadBalancerType"] = args ? args.loadBalancerType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["securityGroups"] = args ? args.securityGroups : undefined;
            inputs["subnetMappings"] = args ? args.subnetMappings : undefined;
            inputs["subnets"] = args ? args.subnets : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["arnSuffix"] = undefined /*out*/;
            inputs["dnsName"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
            inputs["zoneId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    readonly accessLogs?: pulumi.Input<inputs.applicationloadbalancing.LoadBalancerAccessLogs>;
    readonly arn?: pulumi.Input<string>;
    readonly arnSuffix?: pulumi.Input<string>;
    readonly dnsName?: pulumi.Input<string>;
    readonly dropInvalidHeaderFields?: pulumi.Input<boolean>;
    readonly enableCrossZoneLoadBalancing?: pulumi.Input<boolean>;
    readonly enableDeletionProtection?: pulumi.Input<boolean>;
    readonly enableHttp2?: pulumi.Input<boolean>;
    readonly idleTimeout?: pulumi.Input<number>;
    readonly internal?: pulumi.Input<boolean>;
    readonly ipAddressType?: pulumi.Input<IpAddressType>;
    readonly loadBalancerType?: pulumi.Input<LoadBalancerType>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    readonly subnetMappings?: pulumi.Input<pulumi.Input<inputs.applicationloadbalancing.LoadBalancerSubnetMapping>[]>;
    readonly subnets?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcId?: pulumi.Input<string>;
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    readonly accessLogs?: pulumi.Input<inputs.applicationloadbalancing.LoadBalancerAccessLogs>;
    readonly dropInvalidHeaderFields?: pulumi.Input<boolean>;
    readonly enableCrossZoneLoadBalancing?: pulumi.Input<boolean>;
    readonly enableDeletionProtection?: pulumi.Input<boolean>;
    readonly enableHttp2?: pulumi.Input<boolean>;
    readonly idleTimeout?: pulumi.Input<number>;
    readonly internal?: pulumi.Input<boolean>;
    readonly ipAddressType?: pulumi.Input<IpAddressType>;
    readonly loadBalancerType?: pulumi.Input<LoadBalancerType>;
    readonly name?: pulumi.Input<string>;
    readonly namePrefix?: pulumi.Input<string>;
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    readonly subnetMappings?: pulumi.Input<pulumi.Input<inputs.applicationloadbalancing.LoadBalancerSubnetMapping>[]>;
    readonly subnets?: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
