# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VpnGatewayRoutePropagation(pulumi.CustomResource):
    route_table_id: pulumi.Output[str]
    """
    The id of the `ec2.RouteTable` to propagate routes into.
    """
    vpn_gateway_id: pulumi.Output[str]
    """
    The id of the `ec2.VpnGateway` to propagate routes from.
    """
    def __init__(__self__, resource_name, opts=None, route_table_id=None, vpn_gateway_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Requests automatic route propagation between a VPN gateway and a route table.

        > **Note:** This resource should not be used with a route table that has
        the `propagating_vgws` argument set. If that argument is set, any route
        propagation not explicitly listed in its value will be removed.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.VpnGatewayRoutePropagation("example",
            route_table_id=aws_route_table["example"]["id"],
            vpn_gateway_id=aws_vpn_gateway["example"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The id of the `ec2.RouteTable` to propagate routes into.
        :param pulumi.Input[str] vpn_gateway_id: The id of the `ec2.VpnGateway` to propagate routes from.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if route_table_id is None:
                raise TypeError("Missing required property 'route_table_id'")
            __props__['route_table_id'] = route_table_id
            if vpn_gateway_id is None:
                raise TypeError("Missing required property 'vpn_gateway_id'")
            __props__['vpn_gateway_id'] = vpn_gateway_id
        super(VpnGatewayRoutePropagation, __self__).__init__(
            'aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, route_table_id=None, vpn_gateway_id=None):
        """
        Get an existing VpnGatewayRoutePropagation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The id of the `ec2.RouteTable` to propagate routes into.
        :param pulumi.Input[str] vpn_gateway_id: The id of the `ec2.VpnGateway` to propagate routes from.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["route_table_id"] = route_table_id
        __props__["vpn_gateway_id"] = vpn_gateway_id
        return VpnGatewayRoutePropagation(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
