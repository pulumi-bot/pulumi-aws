// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Retrieve information about an EC2 DHCP Options configuration.
 *
 * ## Example Usage
 * ### Lookup by DHCP Options ID
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ec2.getVpcDhcpOptions({
 *     dhcpOptionsId: "dopts-12345678",
 * }));
 * ```
 * ### Lookup by Filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ec2.getVpcDhcpOptions({
 *     filters: [
 *         {
 *             name: "key",
 *             values: ["domain-name"],
 *         },
 *         {
 *             name: "value",
 *             values: ["example.com"],
 *         },
 *     ],
 * }));
 * ```
 */
export function getVpcDhcpOptions(args?: GetVpcDhcpOptionsArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcDhcpOptionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getVpcDhcpOptions:getVpcDhcpOptions", {
        "dhcpOptionsId": args.dhcpOptionsId,
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcDhcpOptions.
 */
export interface GetVpcDhcpOptionsArgs {
    /**
     * The EC2 DHCP Options ID.
     */
    dhcpOptionsId?: string;
    /**
     * List of custom filters as described below.
     */
    filters?: inputs.ec2.GetVpcDhcpOptionsFilter[];
    /**
     * A map of tags assigned to the resource.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVpcDhcpOptions.
 */
export interface GetVpcDhcpOptionsResult {
    /**
     * The ARN of the DHCP Options Set.
     */
    readonly arn: string;
    /**
     * EC2 DHCP Options ID
     */
    readonly dhcpOptionsId: string;
    /**
     * The suffix domain name to used when resolving non Fully Qualified Domain Names. e.g. the `search` value in the `/etc/resolv.conf` file.
     */
    readonly domainName: string;
    /**
     * List of name servers.
     */
    readonly domainNameServers: string[];
    readonly filters?: outputs.ec2.GetVpcDhcpOptionsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of NETBIOS name servers.
     */
    readonly netbiosNameServers: string[];
    /**
     * The NetBIOS node type (1, 2, 4, or 8). For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
     */
    readonly netbiosNodeType: string;
    /**
     * List of NTP servers.
     */
    readonly ntpServers: string[];
    /**
     * The ID of the AWS account that owns the DHCP options set.
     */
    readonly ownerId: string;
    /**
     * A map of tags assigned to the resource.
     */
    readonly tags: {[key: string]: string};
}

export function getVpcDhcpOptionsApply(args?: GetVpcDhcpOptionsApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcDhcpOptionsResult> {
    return pulumi.output(args).apply(a => getVpcDhcpOptions(a, opts))
}

/**
 * A collection of arguments for invoking getVpcDhcpOptions.
 */
export interface GetVpcDhcpOptionsApplyArgs {
    /**
     * The EC2 DHCP Options ID.
     */
    dhcpOptionsId?: pulumi.Input<string>;
    /**
     * List of custom filters as described below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.ec2.GetVpcDhcpOptionsFilter>[]>;
    /**
     * A map of tags assigned to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
