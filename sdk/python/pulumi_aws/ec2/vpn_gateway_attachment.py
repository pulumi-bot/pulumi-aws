# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpnGatewayAttachment(pulumi.CustomResource):
    vpc_id: pulumi.Output[str]
    """
    The ID of the VPC.
    """
    vpn_gateway_id: pulumi.Output[str]
    """
    The ID of the Virtual Private Gateway.
    """
    def __init__(__self__, resource_name, opts=None, vpc_id=None, vpn_gateway_id=None, __name__=None, __opts__=None):
        """
        Provides a Virtual Private Gateway attachment resource, allowing for an existing
        hardware VPN gateway to be attached and/or detached from a VPC.
        
        > **Note:** The `aws_vpn_gateway`
        resource can also automatically attach the Virtual Private Gateway it creates
        to an existing VPC by setting the `vpc_id` attribute accordingly.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the Virtual Private Gateway.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if vpc_id is None:
            raise TypeError('Missing required property vpc_id')
        __props__['vpc_id'] = vpc_id

        if vpn_gateway_id is None:
            raise TypeError('Missing required property vpn_gateway_id')
        __props__['vpn_gateway_id'] = vpn_gateway_id

        super(VpnGatewayAttachment, __self__).__init__(
            'aws:ec2/vpnGatewayAttachment:VpnGatewayAttachment',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

