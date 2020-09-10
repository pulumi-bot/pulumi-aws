// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getVpc(args?: GetVpcArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getVpc:getVpc", {
        "cidrBlock": args.cidrBlock,
        "default": args.default,
        "dhcpOptionsId": args.dhcpOptionsId,
        "filters": args.filters,
        "id": args.id,
        "state": args.state,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpc.
 */
export interface GetVpcArgs {
    readonly cidrBlock?: string;
    readonly default?: boolean;
    readonly dhcpOptionsId?: string;
    readonly filters?: inputs.ec2.GetVpcFilter[];
    readonly id?: string;
    readonly state?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getVpc.
 */
export interface GetVpcResult {
    readonly arn: string;
    readonly cidrBlock: string;
    readonly cidrBlockAssociations: outputs.ec2.GetVpcCidrBlockAssociation[];
    readonly default: boolean;
    readonly dhcpOptionsId: string;
    readonly enableDnsHostnames: boolean;
    readonly enableDnsSupport: boolean;
    readonly filters?: outputs.ec2.GetVpcFilter[];
    readonly id: string;
    readonly instanceTenancy: string;
    readonly ipv6AssociationId: string;
    readonly ipv6CidrBlock: string;
    readonly mainRouteTableId: string;
    readonly ownerId: string;
    readonly state: string;
    readonly tags: {[key: string]: string};
}
