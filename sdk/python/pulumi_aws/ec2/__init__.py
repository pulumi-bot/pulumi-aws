# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .ami import *
from .ami_copy import *
from .ami_from_instance import *
from .ami_launch_permission import *
from .availability_zone_group import *
from .capacity_reservation import *
from .customer_gateway import *
from .default_network_acl import *
from .default_route_table import *
from .default_security_group import *
from .default_subnet import *
from .default_vpc import *
from .default_vpc_dhcp_options import *
from .egress_only_internet_gateway import *
from .eip import *
from .eip_association import *
from .fleet import *
from .flow_log import *
from .get_coip_pool import *
from .get_coip_pools import *
from .get_customer_gateway import *
from .get_instance import *
from .get_instance_type import *
from .get_instance_type_offering import *
from .get_instance_type_offerings import *
from .get_instances import *
from .get_internet_gateway import *
from .get_launch_configuration import *
from .get_launch_template import *
from .get_local_gateway import *
from .get_local_gateway_route_table import *
from .get_local_gateway_route_tables import *
from .get_local_gateway_virtual_interface import *
from .get_local_gateway_virtual_interface_group import *
from .get_local_gateway_virtual_interface_groups import *
from .get_local_gateways import *
from .get_nat_gateway import *
from .get_network_acls import *
from .get_network_interface import *
from .get_network_interfaces import *
from .get_route import *
from .get_route_table import *
from .get_route_tables import *
from .get_security_group import *
from .get_security_groups import *
from .get_spot_price import *
from .get_subnet import *
from .get_subnet_ids import *
from .get_vpc import *
from .get_vpc_dhcp_options import *
from .get_vpc_endpoint import *
from .get_vpc_endpoint_service import *
from .get_vpc_peering_connection import *
from .get_vpc_peering_connections import *
from .get_vpcs import *
from .get_vpn_gateway import *
from .instance import *
from .internet_gateway import *
from .key_pair import *
from .launch_configuration import *
from .launch_template import *
from .local_gateway_route import *
from .local_gateway_route_table_vpc_association import *
from .main_route_table_association import *
from .nat_gateway import *
from .network_acl import *
from .network_acl_rule import *
from .network_interface import *
from .network_interface_attachment import *
from .network_interface_security_group_attachment import *
from .peering_connection_options import *
from .placement_group import *
from .proxy_protocol_policy import *
from .route import *
from .route_table import *
from .route_table_association import *
from .security_group import *
from .security_group_rule import *
from .snapshot_create_volume_permission import *
from .spot_datafeed_subscription import *
from .spot_fleet_request import *
from .spot_instance_request import *
from .subnet import *
from .tag import *
from .traffic_mirror_filter import *
from .traffic_mirror_filter_rule import *
from .traffic_mirror_session import *
from .traffic_mirror_target import *
from .transit_gateway_peering_attachment_accepter import *
from .volume_attachment import *
from .vpc import *
from .vpc_dhcp_options import *
from .vpc_dhcp_options_association import *
from .vpc_endpoint import *
from .vpc_endpoint_connection_notification import *
from .vpc_endpoint_route_table_association import *
from .vpc_endpoint_service import *
from .vpc_endpoint_service_allowed_principle import *
from .vpc_endpoint_subnet_association import *
from .vpc_ipv4_cidr_block_association import *
from .vpc_peering_connection import *
from .vpc_peering_connection_accepter import *
from .vpn_connection import *
from .vpn_connection_route import *
from .vpn_gateway import *
from .vpn_gateway_attachment import *
from .vpn_gateway_route_propagation import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi

    class Module(pulumi.runtime.ResourceModule):
        def version(self):
            return None

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "aws:ec2/ami:Ami":
                return Ami(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/amiCopy:AmiCopy":
                return AmiCopy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/amiFromInstance:AmiFromInstance":
                return AmiFromInstance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/amiLaunchPermission:AmiLaunchPermission":
                return AmiLaunchPermission(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup":
                return AvailabilityZoneGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/capacityReservation:CapacityReservation":
                return CapacityReservation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/customerGateway:CustomerGateway":
                return CustomerGateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/defaultNetworkAcl:DefaultNetworkAcl":
                return DefaultNetworkAcl(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/defaultRouteTable:DefaultRouteTable":
                return DefaultRouteTable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/defaultSecurityGroup:DefaultSecurityGroup":
                return DefaultSecurityGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/defaultSubnet:DefaultSubnet":
                return DefaultSubnet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/defaultVpc:DefaultVpc":
                return DefaultVpc(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions":
                return DefaultVpcDhcpOptions(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway":
                return EgressOnlyInternetGateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/eip:Eip":
                return Eip(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/eipAssociation:EipAssociation":
                return EipAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/fleet:Fleet":
                return Fleet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/flowLog:FlowLog":
                return FlowLog(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/instance:Instance":
                return Instance(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/internetGateway:InternetGateway":
                return InternetGateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/keyPair:KeyPair":
                return KeyPair(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/launchConfiguration:LaunchConfiguration":
                return LaunchConfiguration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/launchTemplate:LaunchTemplate":
                return LaunchTemplate(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/localGatewayRoute:LocalGatewayRoute":
                return LocalGatewayRoute(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation":
                return LocalGatewayRouteTableVpcAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/mainRouteTableAssociation:MainRouteTableAssociation":
                return MainRouteTableAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/natGateway:NatGateway":
                return NatGateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/networkAcl:NetworkAcl":
                return NetworkAcl(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/networkAclRule:NetworkAclRule":
                return NetworkAclRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/networkInterface:NetworkInterface":
                return NetworkInterface(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment":
                return NetworkInterfaceAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/networkInterfaceSecurityGroupAttachment:NetworkInterfaceSecurityGroupAttachment":
                return NetworkInterfaceSecurityGroupAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/peeringConnectionOptions:PeeringConnectionOptions":
                return PeeringConnectionOptions(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/placementGroup:PlacementGroup":
                return PlacementGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/proxyProtocolPolicy:ProxyProtocolPolicy":
                return ProxyProtocolPolicy(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/route:Route":
                return Route(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/routeTable:RouteTable":
                return RouteTable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/routeTableAssociation:RouteTableAssociation":
                return RouteTableAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/securityGroup:SecurityGroup":
                return SecurityGroup(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/securityGroupRule:SecurityGroupRule":
                return SecurityGroupRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/snapshotCreateVolumePermission:SnapshotCreateVolumePermission":
                return SnapshotCreateVolumePermission(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription":
                return SpotDatafeedSubscription(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/spotFleetRequest:SpotFleetRequest":
                return SpotFleetRequest(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/spotInstanceRequest:SpotInstanceRequest":
                return SpotInstanceRequest(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/subnet:Subnet":
                return Subnet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/tag:Tag":
                return Tag(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/trafficMirrorFilter:TrafficMirrorFilter":
                return TrafficMirrorFilter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/trafficMirrorFilterRule:TrafficMirrorFilterRule":
                return TrafficMirrorFilterRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/trafficMirrorSession:TrafficMirrorSession":
                return TrafficMirrorSession(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/trafficMirrorTarget:TrafficMirrorTarget":
                return TrafficMirrorTarget(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/transitGatewayPeeringAttachmentAccepter:TransitGatewayPeeringAttachmentAccepter":
                return TransitGatewayPeeringAttachmentAccepter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/volumeAttachment:VolumeAttachment":
                return VolumeAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpc:Vpc":
                return Vpc(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcDhcpOptions:VpcDhcpOptions":
                return VpcDhcpOptions(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcDhcpOptionsAssociation:VpcDhcpOptionsAssociation":
                return VpcDhcpOptionsAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcEndpoint:VpcEndpoint":
                return VpcEndpoint(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification":
                return VpcEndpointConnectionNotification(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation":
                return VpcEndpointRouteTableAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcEndpointService:VpcEndpointService":
                return VpcEndpointService(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcEndpointServiceAllowedPrinciple:VpcEndpointServiceAllowedPrinciple":
                return VpcEndpointServiceAllowedPrinciple(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation":
                return VpcEndpointSubnetAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcIpv4CidrBlockAssociation:VpcIpv4CidrBlockAssociation":
                return VpcIpv4CidrBlockAssociation(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcPeeringConnection:VpcPeeringConnection":
                return VpcPeeringConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter":
                return VpcPeeringConnectionAccepter(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpnConnection:VpnConnection":
                return VpnConnection(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpnConnectionRoute:VpnConnectionRoute":
                return VpnConnectionRoute(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpnGateway:VpnGateway":
                return VpnGateway(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpnGatewayAttachment:VpnGatewayAttachment":
                return VpnGatewayAttachment(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation":
                return VpnGatewayRoutePropagation(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("aws", "ec2/ami", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/amiCopy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/amiFromInstance", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/amiLaunchPermission", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/availabilityZoneGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/capacityReservation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/customerGateway", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/defaultNetworkAcl", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/defaultRouteTable", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/defaultSecurityGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/defaultSubnet", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/defaultVpc", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/defaultVpcDhcpOptions", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/egressOnlyInternetGateway", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/eip", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/eipAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/fleet", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/flowLog", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/instance", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/internetGateway", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/keyPair", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/launchConfiguration", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/launchTemplate", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/localGatewayRoute", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/localGatewayRouteTableVpcAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/mainRouteTableAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/natGateway", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/networkAcl", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/networkAclRule", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/networkInterface", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/networkInterfaceAttachment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/networkInterfaceSecurityGroupAttachment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/peeringConnectionOptions", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/placementGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/proxyProtocolPolicy", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/route", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/routeTable", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/routeTableAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/securityGroup", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/securityGroupRule", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/snapshotCreateVolumePermission", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/spotDatafeedSubscription", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/spotFleetRequest", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/spotInstanceRequest", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/subnet", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/tag", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/trafficMirrorFilter", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/trafficMirrorFilterRule", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/trafficMirrorSession", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/trafficMirrorTarget", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/transitGatewayPeeringAttachmentAccepter", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/volumeAttachment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpc", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcDhcpOptions", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcDhcpOptionsAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcEndpoint", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcEndpointConnectionNotification", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcEndpointRouteTableAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcEndpointService", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcEndpointServiceAllowedPrinciple", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcEndpointSubnetAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcIpv4CidrBlockAssociation", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcPeeringConnection", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpcPeeringConnectionAccepter", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpnConnection", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpnConnectionRoute", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpnGateway", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpnGatewayAttachment", _module_instance)
    pulumi.runtime.register_resource_module("aws", "ec2/vpnGatewayRoutePropagation", _module_instance)

_register_module()
