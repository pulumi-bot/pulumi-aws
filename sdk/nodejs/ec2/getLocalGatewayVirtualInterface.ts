// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides details about an EC2 Local Gateway Virtual Interface. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = .map(([, ]) => aws.ec2.getLocalGatewayVirtualInterface({
 *     id: __value,
 * }));
 * ```
 */
export function getLocalGatewayVirtualInterface(args?: GetLocalGatewayVirtualInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayVirtualInterfaceResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGatewayVirtualInterface:getLocalGatewayVirtualInterface", {
        "filters": args.filters,
        "id": args.id,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGatewayVirtualInterface.
 */
export interface GetLocalGatewayVirtualInterfaceArgs {
    /**
     * One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeLocalGatewayVirtualInterfaces.html) for supported filters. Detailed below.
     */
    readonly filters?: inputs.ec2.GetLocalGatewayVirtualInterfaceFilter[];
    /**
     * Identifier of EC2 Local Gateway Virtual Interface.
     */
    readonly id?: string;
    /**
     * Key-value map of resource tags, each pair of which must exactly match a pair on the desired local gateway route table.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGatewayVirtualInterface.
 */
export interface GetLocalGatewayVirtualInterfaceResult {
    readonly filters?: outputs.ec2.GetLocalGatewayVirtualInterfaceFilter[];
    readonly id: string;
    /**
     * Local address.
     */
    readonly localAddress: string;
    /**
     * Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the EC2 Local Gateway.
     */
    readonly localBgpAsn: number;
    /**
     * Identifier of the EC2 Local Gateway.
     */
    readonly localGatewayId: string;
    readonly localGatewayVirtualInterfaceIds: string[];
    /**
     * Peer address.
     */
    readonly peerAddress: string;
    /**
     * Border Gateway Protocol (BGP) Autonomous System Number (ASN) of the peer.
     */
    readonly peerBgpAsn: number;
    readonly tags: {[key: string]: string};
    /**
     * Virtual Local Area Network.
     */
    readonly vlan: number;
}
