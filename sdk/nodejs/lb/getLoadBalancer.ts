// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
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
 * const test = aws.lb.getLoadBalancer({
 *     arn: lbArn,
 *     name: lbName,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/lb.html.markdown.
 */
export function getLoadBalancer(args?: GetLoadBalancerArgs, opts?: pulumi.InvokeOptions): Promise<GetLoadBalancerResult> & GetLoadBalancerResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetLoadBalancerResult> = pulumi.runtime.invoke("aws:lb/getLoadBalancer:getLoadBalancer", {
        "arn": args.arn,
        "name": args.name,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
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
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getLoadBalancer.
 */
export interface GetLoadBalancerResult {
    readonly accessLogs: { bucket: string, enabled: boolean, prefix: string };
    readonly arn: string;
    readonly arnSuffix: string;
    readonly dnsName: string;
    readonly enableDeletionProtection: boolean;
    readonly idleTimeout: number;
    readonly internal: boolean;
    readonly loadBalancerType: string;
    readonly name: string;
    readonly securityGroups: string[];
    readonly subnetMappings: { allocationId?: string, subnetId: string }[];
    readonly subnets: string[];
    readonly tags: {[key: string]: any};
    readonly vpcId: string;
    readonly zoneId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
