// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the ID of an [elastic beanstalk hosted zone](http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region).
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const current = aws.elasticbeanstalk.getHostedZone();
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elastic_beanstalk_hosted_zone.html.markdown.
 */
export function getHostedZone(args?: GetHostedZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetHostedZoneResult> & GetHostedZoneResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetHostedZoneResult> = pulumi.runtime.invoke("aws:elasticbeanstalk/getHostedZone:getHostedZone", {
        "region": args.region,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getHostedZone.
 */
export interface GetHostedZoneArgs {
    /**
     * The region you'd like the zone for. By default, fetches the current region.
     */
    readonly region?: string;
}

/**
 * A collection of values returned by getHostedZone.
 */
export interface GetHostedZoneResult {
    /**
     * The region of the hosted zone.
     */
    readonly region?: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
