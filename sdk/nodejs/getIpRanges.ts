// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "./types";
import * as utilities from "./utilities";

/**
 * Use this data source to get the IP ranges of various AWS products and services. For more information about the contents of this data source and required JSON syntax if referencing a custom URL, see the [AWS IP Address Ranges documentation](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const europeanEc2 = aws.getIpRanges({
 *     regions: [
 *         "eu-west-1",
 *         "eu-central-1",
 *     ],
 *     services: ["ec2"],
 * });
 * const fromEurope = new aws.ec2.SecurityGroup("fromEurope", {
 *     ingress: [{
 *         fromPort: "443",
 *         toPort: "443",
 *         protocol: "tcp",
 *         cidrBlocks: europeanEc2.then(europeanEc2 => europeanEc2.cidrBlocks),
 *         ipv6CidrBlocks: europeanEc2.then(europeanEc2 => europeanEc2.ipv6CidrBlocks),
 *     }],
 *     tags: {
 *         CreateDate: europeanEc2.then(europeanEc2 => europeanEc2.createDate),
 *         SyncToken: europeanEc2.then(europeanEc2 => europeanEc2.syncToken),
 *     },
 * });
 * ```
 */
export function getIpRanges(args: GetIpRangesArgs, opts?: pulumi.InvokeOptions): Promise<GetIpRangesResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getIpRanges:getIpRanges", {
        "regions": args.regions,
        "services": args.services,
        "url": args.url,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpRanges.
 */
export interface GetIpRangesArgs {
    /**
     * Filter IP ranges by regions (or include all regions, if
     * omitted). Valid items are `global` (for `cloudfront`) as well as all AWS regions
     * (e.g. `eu-central-1`)
     */
    regions?: string[];
    /**
     * Filter IP ranges by services. Valid items are `amazon`
     * (for amazon.com), `amazonConnect`, `apiGateway`, `cloud9`, `cloudfront`,
     * `codebuild`, `dynamodb`, `ec2`, `ec2InstanceConnect`, `globalaccelerator`,
     * `route53`, `route53Healthchecks`, `s3` and `workspacesGateways`. See the
     * [`service` attribute][2] documentation for other possible values.
     */
    services: string[];
    /**
     * Custom URL for source JSON file. Syntax must match [AWS IP Address Ranges documentation](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html). Defaults to `https://ip-ranges.amazonaws.com/ip-ranges.json`.
     */
    url?: string;
}

/**
 * A collection of values returned by getIpRanges.
 */
export interface GetIpRangesResult {
    /**
     * The lexically ordered list of CIDR blocks.
     */
    readonly cidrBlocks: string[];
    /**
     * The publication time of the IP ranges (e.g. `2016-08-03-23-46-05`).
     */
    readonly createDate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The lexically ordered list of IPv6 CIDR blocks.
     */
    readonly ipv6CidrBlocks: string[];
    readonly regions?: string[];
    readonly services: string[];
    /**
     * The publication time of the IP ranges, in Unix epoch time format
     * (e.g. `1470267965`).
     */
    readonly syncToken: number;
    readonly url?: string;
}

export function getIpRangesApply(args: GetIpRangesApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpRangesResult> {
    return pulumi.output(args).apply(a => getIpRanges(a, opts))
}

/**
 * A collection of arguments for invoking getIpRanges.
 */
export interface GetIpRangesApplyArgs {
    /**
     * Filter IP ranges by regions (or include all regions, if
     * omitted). Valid items are `global` (for `cloudfront`) as well as all AWS regions
     * (e.g. `eu-central-1`)
     */
    regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Filter IP ranges by services. Valid items are `amazon`
     * (for amazon.com), `amazonConnect`, `apiGateway`, `cloud9`, `cloudfront`,
     * `codebuild`, `dynamodb`, `ec2`, `ec2InstanceConnect`, `globalaccelerator`,
     * `route53`, `route53Healthchecks`, `s3` and `workspacesGateways`. See the
     * [`service` attribute][2] documentation for other possible values.
     */
    services: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Custom URL for source JSON file. Syntax must match [AWS IP Address Ranges documentation](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html). Defaults to `https://ip-ranges.amazonaws.com/ip-ranges.json`.
     */
    url?: pulumi.Input<string>;
}
