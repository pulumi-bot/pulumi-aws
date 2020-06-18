// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The VPC Peering Connection data source provides details about
 * a specific VPC peering connection.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // Declare the data source
 * const pc = aws_vpc_foo.id.apply(id => aws.ec2.getVpcPeeringConnection({
 *     peerCidrBlock: "10.0.1.0/22",
 *     vpcId: id,
 * }, { async: true }));
 * // Create a route table
 * const rt = new aws.ec2.RouteTable("rt", {
 *     vpcId: aws_vpc_foo.id,
 * });
 * // Create a route
 * const route = new aws.ec2.Route("r", {
 *     destinationCidrBlock: pc.peerCidrBlock!,
 *     routeTableId: rt.id,
 *     vpcPeeringConnectionId: pc.id!,
 * });
 * ```
 */
export function getVpcPeeringConnection(args?: GetVpcPeeringConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcPeeringConnectionResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getVpcPeeringConnection:getVpcPeeringConnection", {
        "cidrBlock": args.cidrBlock,
        "filters": args.filters,
        "id": args.id,
        "ownerId": args.ownerId,
        "peerCidrBlock": args.peerCidrBlock,
        "peerOwnerId": args.peerOwnerId,
        "peerRegion": args.peerRegion,
        "peerVpcId": args.peerVpcId,
        "region": args.region,
        "status": args.status,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcPeeringConnection.
 */
export interface GetVpcPeeringConnectionArgs {
    /**
     * The CIDR block of the requester VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly cidrBlock?: string;
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetVpcPeeringConnectionFilter[];
    /**
     * The ID of the specific VPC Peering Connection to retrieve.
     */
    readonly id?: string;
    /**
     * The AWS account ID of the owner of the requester VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly ownerId?: string;
    /**
     * The CIDR block of the accepter VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly peerCidrBlock?: string;
    /**
     * The AWS account ID of the owner of the accepter VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly peerOwnerId?: string;
    /**
     * The region of the accepter VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly peerRegion?: string;
    /**
     * The ID of the accepter VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly peerVpcId?: string;
    /**
     * The region of the requester VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly region?: string;
    /**
     * The status of the specific VPC Peering Connection to retrieve.
     */
    readonly status?: string;
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired VPC Peering Connection.
     */
    readonly tags?: {[key: string]: any};
    /**
     * The ID of the requester VPC of the specific VPC Peering Connection to retrieve.
     */
    readonly vpcId?: string;
}

/**
 * A collection of values returned by getVpcPeeringConnection.
 */
export interface GetVpcPeeringConnectionResult {
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the accepter VPC.
     */
    readonly accepter: {[key: string]: boolean};
    readonly cidrBlock: string;
    readonly filters?: outputs.ec2.GetVpcPeeringConnectionFilter[];
    readonly id: string;
    readonly ownerId: string;
    readonly peerCidrBlock: string;
    readonly peerOwnerId: string;
    readonly peerRegion: string;
    readonly peerVpcId: string;
    readonly region: string;
    /**
     * A configuration block that describes [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options set for the requester VPC.
     */
    readonly requester: {[key: string]: boolean};
    readonly status: string;
    readonly tags: {[key: string]: any};
    readonly vpcId: string;
}
