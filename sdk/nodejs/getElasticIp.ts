// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * `aws.ec2.Eip` provides details about a specific Elastic IP.
 *
 * ## Example Usage
 * ### Search By Allocation ID (VPC only)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const byAllocationId = pulumi.output(aws.getElasticIp({
 *     id: "eipalloc-12345678",
 * }, { async: true }));
 * ```
 * ### Search By Filters (EC2-Classic or VPC)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const byFilter = pulumi.output(aws.getElasticIp({
 *     filters: [{
 *         name: "tag:Name",
 *         values: ["exampleNameTagValue"],
 *     }],
 * }, { async: true }));
 * ```
 * ### Search By Public IP (EC2-Classic or VPC)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const byPublicIp = pulumi.output(aws.getElasticIp({
 *     publicIp: "1.2.3.4",
 * }, { async: true }));
 * ```
 * ### Search By Tags (EC2-Classic or VPC)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const byTags = pulumi.output(aws.getElasticIp({
 *     tags: {
 *         Name: "exampleNameTagValue",
 *     },
 * }, { async: true }));
 * ```
 */
export function getElasticIp(args?: GetElasticIpArgs, opts?: pulumi.InvokeOptions): Promise<GetElasticIpResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getElasticIp:getElasticIp", {
        "filters": args.filters,
        "id": args.id,
        "publicIp": args.publicIp,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getElasticIp.
 */
export interface GetElasticIpArgs {
    /**
     * One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).
     */
    readonly filters?: inputs.GetElasticIpFilter[];
    /**
     * The allocation id of the specific VPC EIP to retrieve. If a classic EIP is required, do NOT set `id`, only set `publicIp`
     */
    readonly id?: string;
    /**
     * The public IP of the specific EIP to retrieve.
     */
    readonly publicIp?: string;
    /**
     * A map of tags, each pair of which must exactly match a pair on the desired Elastic IP
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getElasticIp.
 */
export interface GetElasticIpResult {
    /**
     * The ID representing the association of the address with an instance in a VPC.
     */
    readonly associationId: string;
    /**
     * Customer Owned IP.
     */
    readonly customerOwnedIp: string;
    /**
     * The ID of a Customer Owned IP Pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
     */
    readonly customerOwnedIpv4Pool: string;
    /**
     * Indicates whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).
     */
    readonly domain: string;
    readonly filters?: outputs.GetElasticIpFilter[];
    /**
     * If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.
     */
    readonly id: string;
    /**
     * The ID of the instance that the address is associated with (if any).
     */
    readonly instanceId: string;
    /**
     * The ID of the network interface.
     */
    readonly networkInterfaceId: string;
    /**
     * The ID of the AWS account that owns the network interface.
     */
    readonly networkInterfaceOwnerId: string;
    /**
     * The Private DNS associated with the Elastic IP address.
     */
    readonly privateDns: string;
    /**
     * The private IP address associated with the Elastic IP address.
     */
    readonly privateIp: string;
    /**
     * Public DNS associated with the Elastic IP address.
     */
    readonly publicDns: string;
    /**
     * Public IP address of Elastic IP.
     */
    readonly publicIp: string;
    /**
     * The ID of an address pool.
     */
    readonly publicIpv4Pool: string;
    /**
     * Key-value map of tags associated with Elastic IP.
     */
    readonly tags: {[key: string]: any};
}
