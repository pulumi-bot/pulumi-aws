// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Load Balancer resource.
 *
 * > **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.
 *
 * ## Example Usage
 * ### Application Load Balancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.lb.LoadBalancer("test", {
 *     accessLogs: {
 *         bucket: aws_s3_bucket_lb_logs.bucket,
 *         enabled: true,
 *         prefix: "test-lb",
 *     },
 *     enableDeletionProtection: true,
 *     internal: false,
 *     loadBalancerType: "application",
 *     securityGroups: [aws_security_group_lb_sg.id],
 *     subnets: [aws_subnet_public.map(v => v.id)],
 *     tags: {
 *         Environment: "production",
 *     },
 * });
 * ```
 * ### Network Load Balancer
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.lb.LoadBalancer("test", {
 *     enableDeletionProtection: true,
 *     internal: false,
 *     loadBalancerType: "network",
 *     subnets: [aws_subnet_public.map(v => v.id)],
 *     tags: {
 *         Environment: "production",
 *     },
 * });
 * ```
 * ### Specifying Elastic IPs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lb.LoadBalancer("example", {
 *     loadBalancerType: "network",
 *     subnetMappings: [
 *         {
 *             allocationId: aws_eip_example1.id,
 *             subnetId: aws_subnet_example1.id,
 *         },
 *         {
 *             allocationId: aws_eip_example2.id,
 *             subnetId: aws_subnet_example2.id,
 *         },
 *     ],
 * });
 * ```
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
        return new LoadBalancer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lb/loadBalancer:LoadBalancer';

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

    /**
     * An Access Logs block. Access Logs documented below.
     */
    public readonly accessLogs!: pulumi.Output<outputs.lb.LoadBalancerAccessLogs | undefined>;
    /**
     * The ARN of the load balancer (matches `id`).
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The ARN suffix for use with CloudWatch Metrics.
     */
    public /*out*/ readonly arnSuffix!: pulumi.Output<string>;
    /**
     * The DNS name of the load balancer.
     */
    public /*out*/ readonly dnsName!: pulumi.Output<string>;
    /**
     * Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     */
    public readonly dropInvalidHeaderFields!: pulumi.Output<boolean | undefined>;
    /**
     * If true, cross-zone load balancing of the load balancer will be enabled.
     * This is a `network` load balancer feature. Defaults to `false`.
     */
    public readonly enableCrossZoneLoadBalancing!: pulumi.Output<boolean | undefined>;
    /**
     * If true, deletion of the load balancer will be disabled via
     * the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     */
    public readonly enableDeletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     */
    public readonly enableHttp2!: pulumi.Output<boolean | undefined>;
    /**
     * The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     */
    public readonly idleTimeout!: pulumi.Output<number | undefined>;
    /**
     * If true, the LB will be internal.
     */
    public readonly internal!: pulumi.Output<boolean>;
    /**
     * The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
     */
    public readonly ipAddressType!: pulumi.Output<string>;
    /**
     * The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
     */
    public readonly loadBalancerType!: pulumi.Output<string | undefined>;
    /**
     * The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
     */
    public readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * A subnet mapping block as documented below.
     */
    public readonly subnetMappings!: pulumi.Output<outputs.lb.LoadBalancerSubnetMapping[]>;
    /**
     * A list of subnet IDs to attach to the LB. Subnets
     * cannot be updated for Load Balancers of type `network`. Changing this value
     * for load balancers of type `network` will force a recreation of the resource.
     */
    public readonly subnets!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
     */
    public /*out*/ readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a LoadBalancer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: LoadBalancerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoadBalancerArgs | LoadBalancerState, opts?: pulumi.CustomResourceOptions) {
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
        const aliasOpts = { aliases: [{ type: "aws:elasticloadbalancingv2/loadBalancer:LoadBalancer" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(LoadBalancer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoadBalancer resources.
 */
export interface LoadBalancerState {
    /**
     * An Access Logs block. Access Logs documented below.
     */
    readonly accessLogs?: pulumi.Input<inputs.lb.LoadBalancerAccessLogs>;
    /**
     * The ARN of the load balancer (matches `id`).
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The ARN suffix for use with CloudWatch Metrics.
     */
    readonly arnSuffix?: pulumi.Input<string>;
    /**
     * The DNS name of the load balancer.
     */
    readonly dnsName?: pulumi.Input<string>;
    /**
     * Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     */
    readonly dropInvalidHeaderFields?: pulumi.Input<boolean>;
    /**
     * If true, cross-zone load balancing of the load balancer will be enabled.
     * This is a `network` load balancer feature. Defaults to `false`.
     */
    readonly enableCrossZoneLoadBalancing?: pulumi.Input<boolean>;
    /**
     * If true, deletion of the load balancer will be disabled via
     * the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     */
    readonly enableDeletionProtection?: pulumi.Input<boolean>;
    /**
     * Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     */
    readonly enableHttp2?: pulumi.Input<boolean>;
    /**
     * The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     */
    readonly idleTimeout?: pulumi.Input<number>;
    /**
     * If true, the LB will be internal.
     */
    readonly internal?: pulumi.Input<boolean>;
    /**
     * The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
     */
    readonly ipAddressType?: pulumi.Input<string>;
    /**
     * The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
     */
    readonly loadBalancerType?: pulumi.Input<string>;
    /**
     * The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A subnet mapping block as documented below.
     */
    readonly subnetMappings?: pulumi.Input<pulumi.Input<inputs.lb.LoadBalancerSubnetMapping>[]>;
    /**
     * A list of subnet IDs to attach to the LB. Subnets
     * cannot be updated for Load Balancers of type `network`. Changing this value
     * for load balancers of type `network` will force a recreation of the resource.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly vpcId?: pulumi.Input<string>;
    /**
     * The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
     */
    readonly zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoadBalancer resource.
 */
export interface LoadBalancerArgs {
    /**
     * An Access Logs block. Access Logs documented below.
     */
    readonly accessLogs?: pulumi.Input<inputs.lb.LoadBalancerAccessLogs>;
    /**
     * Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer (true) or routed to targets (false). The default is false. Elastic Load Balancing requires that message header names contain only alphanumeric characters and hyphens. Only valid for Load Balancers of type `application`.
     */
    readonly dropInvalidHeaderFields?: pulumi.Input<boolean>;
    /**
     * If true, cross-zone load balancing of the load balancer will be enabled.
     * This is a `network` load balancer feature. Defaults to `false`.
     */
    readonly enableCrossZoneLoadBalancing?: pulumi.Input<boolean>;
    /**
     * If true, deletion of the load balancer will be disabled via
     * the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
     */
    readonly enableDeletionProtection?: pulumi.Input<boolean>;
    /**
     * Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
     */
    readonly enableHttp2?: pulumi.Input<boolean>;
    /**
     * The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
     */
    readonly idleTimeout?: pulumi.Input<number>;
    /**
     * If true, the LB will be internal.
     */
    readonly internal?: pulumi.Input<boolean>;
    /**
     * The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
     */
    readonly ipAddressType?: pulumi.Input<string>;
    /**
     * The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
     */
    readonly loadBalancerType?: pulumi.Input<string>;
    /**
     * The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
     * must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
     * this provider will autogenerate a name beginning with `tf-lb`.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the specified prefix. Conflicts with `name`.
     */
    readonly namePrefix?: pulumi.Input<string>;
    /**
     * A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A subnet mapping block as documented below.
     */
    readonly subnetMappings?: pulumi.Input<pulumi.Input<inputs.lb.LoadBalancerSubnetMapping>[]>;
    /**
     * A list of subnet IDs to attach to the LB. Subnets
     * cannot be updated for Load Balancers of type `network`. Changing this value
     * for load balancers of type `network` will force a recreation of the resource.
     */
    readonly subnets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
