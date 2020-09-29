// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
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
 * const main = aws.elb.getHostedZoneId({});
 * const www = new aws.route53.Record("www", {
 *     zoneId: aws_route53_zone.primary.zone_id,
 *     name: "example.com",
 *     type: "A",
 *     aliases: [{
 *         name: aws_elb.main.dns_name,
 *         zoneId: main.then(main => main.id),
 *         evaluateTargetHealth: true,
 *     }],
 * });
 * ```
 */
/** @deprecated aws.elasticloadbalancing.getHostedZoneId has been deprecated in favor of aws.elb.getHostedZoneId */
export function getHostedZoneId(args?: GetHostedZoneIdArgs, opts?: pulumi.InvokeOptions): Promise<GetHostedZoneIdResult> {
    pulumi.log.warn("getHostedZoneId is deprecated: aws.elasticloadbalancing.getHostedZoneId has been deprecated in favor of aws.elb.getHostedZoneId")
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId", {
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
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly region?: string;
}
