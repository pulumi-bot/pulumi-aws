# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpnConnectionRoute(pulumi.CustomResource):
    destination_cidr_block: pulumi.Output[str]
    """
    The CIDR block associated with the local subnet of the customer network.
    """
    vpn_connection_id: pulumi.Output[str]
    """
    The ID of the VPN connection.
    """
    def __init__(__self__, resource_name, opts=None, destination_cidr_block=None, vpn_connection_id=None, __name__=None, __opts__=None):
        """
        Provides a static route between a VPN connection and a customer gateway.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The CIDR block associated with the local subnet of the customer network.
        :param pulumi.Input[str] vpn_connection_id: The ID of the VPN connection.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpn_connection_route.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if destination_cidr_block is None:
            raise TypeError("Missing required property 'destination_cidr_block'")
        __props__['destination_cidr_block'] = destination_cidr_block
        if vpn_connection_id is None:
            raise TypeError("Missing required property 'vpn_connection_id'")
        __props__['vpn_connection_id'] = vpn_connection_id
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(VpnConnectionRoute, __self__).__init__(
            'aws:ec2/vpnConnectionRoute:VpnConnectionRoute',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

