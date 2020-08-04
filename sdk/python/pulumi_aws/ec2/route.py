# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Route(pulumi.CustomResource):
    destination_cidr_block: pulumi.Output[str]
    """
    The destination CIDR block.
    """
    destination_ipv6_cidr_block: pulumi.Output[str]
    """
    The destination IPv6 CIDR block.
    """
    destination_prefix_list_id: pulumi.Output[str]
    egress_only_gateway_id: pulumi.Output[str]
    """
    Identifier of a VPC Egress Only Internet Gateway.
    """
    gateway_id: pulumi.Output[str]
    """
    Identifier of a VPC internet gateway or a virtual private gateway.
    """
    instance_id: pulumi.Output[str]
    """
    Identifier of an EC2 instance.
    """
    instance_owner_id: pulumi.Output[str]
    nat_gateway_id: pulumi.Output[str]
    """
    Identifier of a VPC NAT gateway.
    """
    network_interface_id: pulumi.Output[str]
    """
    Identifier of an EC2 network interface.
    """
    origin: pulumi.Output[str]
    route_table_id: pulumi.Output[str]
    """
    The ID of the routing table.
    """
    state: pulumi.Output[str]
    transit_gateway_id: pulumi.Output[str]
    """
    Identifier of an EC2 Transit Gateway.
    """
    vpc_peering_connection_id: pulumi.Output[str]
    """
    Identifier of a VPC peering connection.
    """
    def __init__(__self__, resource_name, opts=None, destination_cidr_block=None, destination_ipv6_cidr_block=None, egress_only_gateway_id=None, gateway_id=None, instance_id=None, nat_gateway_id=None, network_interface_id=None, route_table_id=None, transit_gateway_id=None, vpc_peering_connection_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a routing table entry (a route) in a VPC routing table.

        > **NOTE on Route Tables and Routes:** This provider currently
        provides both a standalone Route resource and a Route Table resource with routes
        defined in-line. At this time you cannot use a Route Table with in-line routes
        in conjunction with any Route resources. Doing so will cause
        a conflict of rule settings and will overwrite rules.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        route = aws.ec2.Route("route",
            route_table_id="rtb-4fbb3ac4",
            destination_cidr_block="10.0.1.0/22",
            vpc_peering_connection_id="pcx-45ff3dc1",
            opts=ResourceOptions(depends_on=["aws_route_table.testing"]))
        ```
        ## Example IPv6 Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        vpc = aws.ec2.Vpc("vpc",
            assign_generated_ipv6_cidr_block=True,
            cidr_block="10.1.0.0/16")
        egress = aws.ec2.EgressOnlyInternetGateway("egress", vpc_id=vpc.id)
        route = aws.ec2.Route("route",
            destination_ipv6_cidr_block="::/0",
            egress_only_gateway_id=egress.id,
            route_table_id="rtb-4fbb3ac4")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block.
        :param pulumi.Input[str] destination_ipv6_cidr_block: The destination IPv6 CIDR block.
        :param pulumi.Input[str] egress_only_gateway_id: Identifier of a VPC Egress Only Internet Gateway.
        :param pulumi.Input[str] gateway_id: Identifier of a VPC internet gateway or a virtual private gateway.
        :param pulumi.Input[str] instance_id: Identifier of an EC2 instance.
        :param pulumi.Input[str] nat_gateway_id: Identifier of a VPC NAT gateway.
        :param pulumi.Input[str] network_interface_id: Identifier of an EC2 network interface.
        :param pulumi.Input[str] route_table_id: The ID of the routing table.
        :param pulumi.Input[str] transit_gateway_id: Identifier of an EC2 Transit Gateway.
        :param pulumi.Input[str] vpc_peering_connection_id: Identifier of a VPC peering connection.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['destination_cidr_block'] = destination_cidr_block
            __props__['destination_ipv6_cidr_block'] = destination_ipv6_cidr_block
            __props__['egress_only_gateway_id'] = egress_only_gateway_id
            __props__['gateway_id'] = gateway_id
            __props__['instance_id'] = instance_id
            __props__['nat_gateway_id'] = nat_gateway_id
            __props__['network_interface_id'] = network_interface_id
            if route_table_id is None:
                raise TypeError("Missing required property 'route_table_id'")
            __props__['route_table_id'] = route_table_id
            __props__['transit_gateway_id'] = transit_gateway_id
            __props__['vpc_peering_connection_id'] = vpc_peering_connection_id
            __props__['destination_prefix_list_id'] = None
            __props__['instance_owner_id'] = None
            __props__['origin'] = None
            __props__['state'] = None
        super(Route, __self__).__init__(
            'aws:ec2/route:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, destination_cidr_block=None, destination_ipv6_cidr_block=None, destination_prefix_list_id=None, egress_only_gateway_id=None, gateway_id=None, instance_id=None, instance_owner_id=None, nat_gateway_id=None, network_interface_id=None, origin=None, route_table_id=None, state=None, transit_gateway_id=None, vpc_peering_connection_id=None):
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The destination CIDR block.
        :param pulumi.Input[str] destination_ipv6_cidr_block: The destination IPv6 CIDR block.
        :param pulumi.Input[str] egress_only_gateway_id: Identifier of a VPC Egress Only Internet Gateway.
        :param pulumi.Input[str] gateway_id: Identifier of a VPC internet gateway or a virtual private gateway.
        :param pulumi.Input[str] instance_id: Identifier of an EC2 instance.
        :param pulumi.Input[str] nat_gateway_id: Identifier of a VPC NAT gateway.
        :param pulumi.Input[str] network_interface_id: Identifier of an EC2 network interface.
        :param pulumi.Input[str] route_table_id: The ID of the routing table.
        :param pulumi.Input[str] transit_gateway_id: Identifier of an EC2 Transit Gateway.
        :param pulumi.Input[str] vpc_peering_connection_id: Identifier of a VPC peering connection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["destination_cidr_block"] = destination_cidr_block
        __props__["destination_ipv6_cidr_block"] = destination_ipv6_cidr_block
        __props__["destination_prefix_list_id"] = destination_prefix_list_id
        __props__["egress_only_gateway_id"] = egress_only_gateway_id
        __props__["gateway_id"] = gateway_id
        __props__["instance_id"] = instance_id
        __props__["instance_owner_id"] = instance_owner_id
        __props__["nat_gateway_id"] = nat_gateway_id
        __props__["network_interface_id"] = network_interface_id
        __props__["origin"] = origin
        __props__["route_table_id"] = route_table_id
        __props__["state"] = state
        __props__["transit_gateway_id"] = transit_gateway_id
        __props__["vpc_peering_connection_id"] = vpc_peering_connection_id
        return Route(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
