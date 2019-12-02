// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
 * in a given region for the purpose of using in an AWS Route53 Alias.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const main = aws.elb.getHostedZoneId();
 * const www = new aws.route53.Record("www", {
 *     aliases: [{
 *         evaluateTargetHealth: true,
 *         name: aws_elb_main.dnsName,
 *         zoneId: main.id,
 *     }],
 *     type: "A",
 *     zoneId: aws_route53_zone_primary.zoneId,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_hosted_zone_id.html.markdown.
 */
export function getHostedZoneId(args?: GetHostedZoneIdArgs, opts?: pulumi.InvokeOptions): Promise<GetHostedZoneIdResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:elb/getHostedZoneId:getHostedZoneId", {
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getHostedZoneId.
 */
export interface GetHostedZoneIdArgs {
    /**
     * Name of the region whose AWS ELB HostedZoneId is desired.
     * Defaults to the region from the AWS provider configuration.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getHostedZoneId.
 */
export interface GetHostedZoneIdResult {
    readonly region?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
