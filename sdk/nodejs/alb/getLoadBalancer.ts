// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.
 *
 * Provides information about a Load Balancer.
 *
 * This data source can prove useful when a module accepts an LB as an input
 * variable and needs to, for example, determine the security groups associated
 * with it, etc.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const lbArn = config.get("lbArn") || "";
 * const lbName = config.get("lbName") || "";
 *
 * const test = pulumi.output(aws.lb.getLoadBalancer({
 *     arn: lbArn,
 *     name: lbName,
 * }, { async: true }));
 * ```
 */
export function getLoadBalancer(args?: GetLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:alb/getLoadBalancer:getLoadBalancer", {
        "arn": args.arn,
        "name": args.name,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLoadBalancer.
 */
export interface GetLoadBalancerArgs {
    /**
     * The full ARN of the load balancer.
     */
    readonly arn?: string;
    /**
     * The unique name of the load balancer.
     */
    readonly name?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLoadBalancer.
 */
export interface GetLoadBalancerResult {
    readonly accessLogs: outputs.alb.GetLoadBalancerAccessLogs;
    readonly arn: string;
    readonly arnSuffix: string;
    readonly dnsName: string;
    readonly dropInvalidHeaderFields: boolean;
    readonly enableDeletionProtection: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly idleTimeout: number;
    readonly internal: boolean;
    readonly ipAddressType: string;
    readonly loadBalancerType: string;
    readonly name: string;
    readonly securityGroups: string[];
    readonly subnetMappings: outputs.alb.GetLoadBalancerSubnetMapping[];
    readonly subnets: string[];
    readonly tags: {[key: string]: string};
    readonly vpcId: string;
    readonly zoneId: string;
}
