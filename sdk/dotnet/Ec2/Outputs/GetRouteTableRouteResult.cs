// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2.Outputs
{

    [OutputType]
    public sealed class GetRouteTableRouteResult
    {
        public readonly string CidrBlock;
        public readonly string EgressOnlyGatewayId;
        public readonly string GatewayId;
        public readonly string InstanceId;
        public readonly string Ipv6CidrBlock;
        public readonly string NatGatewayId;
        public readonly string NetworkInterfaceId;
        public readonly string TransitGatewayId;
        public readonly string VpcPeeringConnectionId;

        [OutputConstructor]
        private GetRouteTableRouteResult(
            string cidrBlock,

            string egressOnlyGatewayId,

            string gatewayId,

            string instanceId,

            string ipv6CidrBlock,

            string natGatewayId,

            string networkInterfaceId,

            string transitGatewayId,

            string vpcPeeringConnectionId)
        {
            CidrBlock = cidrBlock;
            EgressOnlyGatewayId = egressOnlyGatewayId;
            GatewayId = gatewayId;
            InstanceId = instanceId;
            Ipv6CidrBlock = ipv6CidrBlock;
            NatGatewayId = natGatewayId;
            NetworkInterfaceId = networkInterfaceId;
            TransitGatewayId = transitGatewayId;
            VpcPeeringConnectionId = vpcPeeringConnectionId;
        }
    }
}
