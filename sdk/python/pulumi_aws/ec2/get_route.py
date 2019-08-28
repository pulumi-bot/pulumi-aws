# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetRouteResult:
    """
    A collection of values returned by getRoute.
    """
    def __init__(__self__, destination_cidr_block=None, destination_ipv6_cidr_block=None, egress_only_gateway_id=None, gateway_id=None, instance_id=None, nat_gateway_id=None, network_interface_id=None, route_table_id=None, transit_gateway_id=None, vpc_peering_connection_id=None, id=None):
        if destination_cidr_block and not isinstance(destination_cidr_block, str):
            raise TypeError("Expected argument 'destination_cidr_block' to be a str")
        __self__.destination_cidr_block = destination_cidr_block
        if destination_ipv6_cidr_block and not isinstance(destination_ipv6_cidr_block, str):
            raise TypeError("Expected argument 'destination_ipv6_cidr_block' to be a str")
        __self__.destination_ipv6_cidr_block = destination_ipv6_cidr_block
        if egress_only_gateway_id and not isinstance(egress_only_gateway_id, str):
            raise TypeError("Expected argument 'egress_only_gateway_id' to be a str")
        __self__.egress_only_gateway_id = egress_only_gateway_id
        if gateway_id and not isinstance(gateway_id, str):
            raise TypeError("Expected argument 'gateway_id' to be a str")
        __self__.gateway_id = gateway_id
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        __self__.instance_id = instance_id
        if nat_gateway_id and not isinstance(nat_gateway_id, str):
            raise TypeError("Expected argument 'nat_gateway_id' to be a str")
        __self__.nat_gateway_id = nat_gateway_id
        if network_interface_id and not isinstance(network_interface_id, str):
            raise TypeError("Expected argument 'network_interface_id' to be a str")
        __self__.network_interface_id = network_interface_id
        if route_table_id and not isinstance(route_table_id, str):
            raise TypeError("Expected argument 'route_table_id' to be a str")
        __self__.route_table_id = route_table_id
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        __self__.transit_gateway_id = transit_gateway_id
        if vpc_peering_connection_id and not isinstance(vpc_peering_connection_id, str):
            raise TypeError("Expected argument 'vpc_peering_connection_id' to be a str")
        __self__.vpc_peering_connection_id = vpc_peering_connection_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetRouteResult(GetRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteResult(
            destination_cidr_block=self.destination_cidr_block,
            destination_ipv6_cidr_block=self.destination_ipv6_cidr_block,
            egress_only_gateway_id=self.egress_only_gateway_id,
            gateway_id=self.gateway_id,
            instance_id=self.instance_id,
            nat_gateway_id=self.nat_gateway_id,
            network_interface_id=self.network_interface_id,
            route_table_id=self.route_table_id,
            transit_gateway_id=self.transit_gateway_id,
            vpc_peering_connection_id=self.vpc_peering_connection_id,
            id=self.id)

def get_route(destination_cidr_block=None,destination_ipv6_cidr_block=None,egress_only_gateway_id=None,gateway_id=None,instance_id=None,nat_gateway_id=None,network_interface_id=None,route_table_id=None,transit_gateway_id=None,vpc_peering_connection_id=None,opts=None):
    """
    `ec2.Route` provides details about a specific Route.
    
    This resource can prove useful when finding the resource
    associated with a CIDR. For example, finding the peering
    connection associated with a CIDR value.
    
    :param str destination_cidr_block: The CIDR block of the Route belonging to the Route Table.
    :param str destination_ipv6_cidr_block: The IPv6 CIDR block of the Route belonging to the Route Table.
    :param str egress_only_gateway_id: The Egress Only Gateway ID of the Route belonging to the Route Table.
    :param str gateway_id: The Gateway ID of the Route belonging to the Route Table.
    :param str instance_id: The Instance ID of the Route belonging to the Route Table.
    :param str nat_gateway_id: The NAT Gateway ID of the Route belonging to the Route Table.
    :param str network_interface_id: The Network Interface ID of the Route belonging to the Route Table.
    :param str route_table_id: The id of the specific Route Table containing the Route entry.
    :param str transit_gateway_id: The EC2 Transit Gateway ID of the Route belonging to the Route Table.
    :param str vpc_peering_connection_id: The VPC Peering Connection ID of the Route belonging to the Route Table.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route.html.markdown.
    """
    __args__ = dict()

    __args__['destinationCidrBlock'] = destination_cidr_block
    __args__['destinationIpv6CidrBlock'] = destination_ipv6_cidr_block
    __args__['egressOnlyGatewayId'] = egress_only_gateway_id
    __args__['gatewayId'] = gateway_id
    __args__['instanceId'] = instance_id
    __args__['natGatewayId'] = nat_gateway_id
    __args__['networkInterfaceId'] = network_interface_id
    __args__['routeTableId'] = route_table_id
    __args__['transitGatewayId'] = transit_gateway_id
    __args__['vpcPeeringConnectionId'] = vpc_peering_connection_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2/getRoute:getRoute', __args__, opts=opts).value

    return AwaitableGetRouteResult(
        destination_cidr_block=__ret__.get('destinationCidrBlock'),
        destination_ipv6_cidr_block=__ret__.get('destinationIpv6CidrBlock'),
        egress_only_gateway_id=__ret__.get('egressOnlyGatewayId'),
        gateway_id=__ret__.get('gatewayId'),
        instance_id=__ret__.get('instanceId'),
        nat_gateway_id=__ret__.get('natGatewayId'),
        network_interface_id=__ret__.get('networkInterfaceId'),
        route_table_id=__ret__.get('routeTableId'),
        transit_gateway_id=__ret__.get('transitGatewayId'),
        vpc_peering_connection_id=__ret__.get('vpcPeeringConnectionId'),
        id=__ret__.get('id'))
