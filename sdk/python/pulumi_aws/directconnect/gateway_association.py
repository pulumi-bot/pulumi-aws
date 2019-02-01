# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GatewayAssociation(pulumi.CustomResource):
    dx_gateway_id: pulumi.Output[str]
    """
    The ID of the Direct Connect Gateway.
    """
    vpn_gateway_id: pulumi.Output[str]
    """
    The ID of the VGW with which to associate the gateway.
    """
    def __init__(__self__, __name__, __opts__=None, dx_gateway_id=None, vpn_gateway_id=None):
        """
        Associates a Direct Connect Gateway with a VGW.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] dx_gateway_id: The ID of the Direct Connect Gateway.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VGW with which to associate the gateway.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not dx_gateway_id:
            raise TypeError('Missing required property dx_gateway_id')
        __props__['dx_gateway_id'] = dx_gateway_id

        if not vpn_gateway_id:
            raise TypeError('Missing required property vpn_gateway_id')
        __props__['vpn_gateway_id'] = vpn_gateway_id

        super(GatewayAssociation, __self__).__init__(
            'aws:directconnect/gatewayAssociation:GatewayAssociation',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

