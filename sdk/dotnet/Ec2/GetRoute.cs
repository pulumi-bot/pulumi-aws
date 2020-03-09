// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public static partial class GetRoute
    {
        /// <summary>
        /// `aws.ec2.Route` provides details about a specific Route.
        /// 
        /// This resource can prove useful when finding the resource
        /// associated with a CIDR. For example, finding the peering
        /// connection associated with a CIDR value.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route.html.markdown.
        /// </summary>
        public static Task<GetRouteResult> InvokeAsync(GetRouteArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRouteResult>("aws:ec2/getRoute:getRoute", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRouteArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The CIDR block of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationCidrBlock")]
        public string? DestinationCidrBlock { get; set; }

        /// <summary>
        /// The IPv6 CIDR block of the Route belonging to the Route Table.
        /// </summary>
        [Input("destinationIpv6CidrBlock")]
        public string? DestinationIpv6CidrBlock { get; set; }

        /// <summary>
        /// The Egress Only Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("egressOnlyGatewayId")]
        public string? EgressOnlyGatewayId { get; set; }

        /// <summary>
        /// The Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("gatewayId")]
        public string? GatewayId { get; set; }

        /// <summary>
        /// The Instance ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("instanceId")]
        public string? InstanceId { get; set; }

        /// <summary>
        /// The NAT Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("natGatewayId")]
        public string? NatGatewayId { get; set; }

        /// <summary>
        /// The Network Interface ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("networkInterfaceId")]
        public string? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The id of the specific Route Table containing the Route entry.
        /// </summary>
        [Input("routeTableId", required: true)]
        public string RouteTableId { get; set; } = null!;

        /// <summary>
        /// The EC2 Transit Gateway ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("transitGatewayId")]
        public string? TransitGatewayId { get; set; }

        /// <summary>
        /// The VPC Peering Connection ID of the Route belonging to the Route Table.
        /// </summary>
        [Input("vpcPeeringConnectionId")]
        public string? VpcPeeringConnectionId { get; set; }

        public GetRouteArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRouteResult
    {
        public readonly string DestinationCidrBlock;
        public readonly string DestinationIpv6CidrBlock;
        public readonly string EgressOnlyGatewayId;
        public readonly string GatewayId;
        public readonly string InstanceId;
        public readonly string NatGatewayId;
        public readonly string NetworkInterfaceId;
        public readonly string RouteTableId;
        public readonly string TransitGatewayId;
        public readonly string VpcPeeringConnectionId;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRouteResult(
            string destinationCidrBlock,
            string destinationIpv6CidrBlock,
            string egressOnlyGatewayId,
            string gatewayId,
            string instanceId,
            string natGatewayId,
            string networkInterfaceId,
            string routeTableId,
            string transitGatewayId,
            string vpcPeeringConnectionId,
            string id)
        {
            DestinationCidrBlock = destinationCidrBlock;
            DestinationIpv6CidrBlock = destinationIpv6CidrBlock;
            EgressOnlyGatewayId = egressOnlyGatewayId;
            GatewayId = gatewayId;
            InstanceId = instanceId;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            RouteTableId = routeTableId;
            TransitGatewayId = transitGatewayId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
            Id = id;
        }
    }
}
