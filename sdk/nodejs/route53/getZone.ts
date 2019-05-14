// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `aws_route53_zone` provides details about a specific Route 53 Hosted Zone.
 * 
 * This data source allows to find a Hosted Zone ID given Hosted Zone name and certain search criteria.
 * 
 * ## Example Usage
 * 
 * The following example shows how to get a Hosted Zone from its name and from this data how to create a Record Set.
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const selected = pulumi.output(aws.route53.getZone({
 *     name: "test.com.",
 *     privateZone: true,
 * }));
 * const www = new aws.route53.Record("www", {
 *     records: ["10.0.0.1"],
 *     ttl: 300,
 *     type: "A",
 *     zoneId: selected.zoneId,
 * });
 * ```
 */
export function getZone(args?: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:route53/getZone:getZone", {
        "callerReference": args.callerReference,
        "comment": args.comment,
        "name": args.name,
        "privateZone": args.privateZone,
        "resourceRecordSetCount": args.resourceRecordSetCount,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    readonly callerReference?: string;
    readonly comment?: string;
    /**
     * The Hosted Zone name of the desired Hosted Zone.
     */
    readonly name?: string;
    /**
     * Used with `name` field to get a private Hosted Zone.
     */
    readonly privateZone?: boolean;
    readonly resourceRecordSetCount?: number;
    /**
     * Used with `name` field. A mapping of tags, each pair of which must exactly match
     * a pair on the desired Hosted Zone.
     */
    readonly tags?: {[key: string]: any};
    /**
     * Used with `name` field to get a private Hosted Zone associated with the vpc_id (in this case, private_zone is not mandatory).
     */
    readonly vpcId?: string;
    /**
     * The Hosted Zone id of the desired Hosted Zone.
     */
    readonly zoneId?: string;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    /**
     * Caller Reference of the Hosted Zone.
     */
    readonly callerReference: string;
    /**
     * The comment field of the Hosted Zone.
     */
    readonly comment: string;
    readonly name: string;
    /**
     * The list of DNS name servers for the Hosted Zone.
     */
    readonly nameServers: string[];
    readonly privateZone?: boolean;
    /**
     * the number of Record Set in the Hosted Zone
     */
    readonly resourceRecordSetCount: number;
    readonly tags: {[key: string]: any};
    readonly vpcId: string;
    readonly zoneId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
