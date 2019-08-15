// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * `aws.ec2.Route` provides details about a specific Route.
 * 
 * This resource can prove useful when finding the resource
 * associated with a CIDR. For example, finding the peering
 * connection associated with a CIDR value.
 * 
 * ## Example Usage
 * 
 * The following example shows how one might use a CIDR value to find a network interface id
 * and use this to create a data source of that network interface.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const subnetId = config.require("subnetId");
 * 
 * const route = aws_route_table_selected.id.apply(id => aws.ec2.getRoute({
 *     destinationCidrBlock: "10.0.1.0/24",
 *     routeTableId: id,
 * }));
 * const interfaceNetworkInterface = route.apply(route => aws.ec2.getNetworkInterface({
 *     networkInterfaceId: route.networkInterfaceId,
 * }));
 * const selected = aws.ec2.getRouteTable({
 *     subnetId: subnetId,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route.html.markdown.
 */
export function getRoute(args: GetRouteArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteResult> & GetRouteResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetRouteResult> = pulumi.runtime.invoke("aws:ec2/getRoute:getRoute", {
        "destinationCidrBlock": args.destinationCidrBlock,
        "destinationIpv6CidrBlock": args.destinationIpv6CidrBlock,
        "egressOnlyGatewayId": args.egressOnlyGatewayId,
        "gatewayId": args.gatewayId,
        "instanceId": args.instanceId,
        "natGatewayId": args.natGatewayId,
        "networkInterfaceId": args.networkInterfaceId,
        "routeTableId": args.routeTableId,
        "transitGatewayId": args.transitGatewayId,
        "vpcPeeringConnectionId": args.vpcPeeringConnectionId,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getRoute.
 */
export interface GetRouteArgs {
    /**
     * The CIDR block of the Route belonging to the Route Table.
     */
    readonly destinationCidrBlock?: string;
    /**
     * The IPv6 CIDR block of the Route belonging to the Route Table.
     */
    readonly destinationIpv6CidrBlock?: string;
    /**
     * The Egress Only Gateway ID of the Route belonging to the Route Table.
     */
    readonly egressOnlyGatewayId?: string;
    /**
     * The Gateway ID of the Route belonging to the Route Table.
     */
    readonly gatewayId?: string;
    /**
     * The Instance ID of the Route belonging to the Route Table.
     */
    readonly instanceId?: string;
    /**
     * The NAT Gateway ID of the Route belonging to the Route Table.
     */
    readonly natGatewayId?: string;
    /**
     * The Network Interface ID of the Route belonging to the Route Table.
     */
    readonly networkInterfaceId?: string;
    /**
     * The id of the specific Route Table containing the Route entry.
     */
    readonly routeTableId: string;
    /**
     * The EC2 Transit Gateway ID of the Route belonging to the Route Table.
     */
    readonly transitGatewayId?: string;
    /**
     * The VPC Peering Connection ID of the Route belonging to the Route Table.
     */
    readonly vpcPeeringConnectionId?: string;
}

/**
 * A collection of values returned by getRoute.
 */
export interface GetRouteResult {
    readonly destinationCidrBlock: string;
    readonly destinationIpv6CidrBlock: string;
    readonly egressOnlyGatewayId: string;
    readonly gatewayId: string;
    readonly instanceId: string;
    readonly natGatewayId: string;
    readonly networkInterfaceId: string;
    readonly routeTableId: string;
    readonly transitGatewayId: string;
    readonly vpcPeeringConnectionId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
