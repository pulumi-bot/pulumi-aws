# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class VpnGateway(pulumi.CustomResource):
    amazon_side_asn: pulumi.Output[str]
    """
    The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
    """
    availability_zone: pulumi.Output[str]
    """
    The Availability Zone for the virtual private gateway.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The VPC ID to create in.
    """
    def __init__(__self__, __name__, __opts__=None, amazon_side_asn=None, availability_zone=None, tags=None, vpc_id=None):
        """
        Provides a resource to create a VPC VPN Gateway.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] amazon_side_asn: The Autonomous System Number (ASN) for the Amazon side of the gateway. If you don't specify an ASN, the virtual private gateway is created with the default ASN.
        :param pulumi.Input[str] availability_zone: The Availability Zone for the virtual private gateway.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID to create in.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['amazon_side_asn'] = amazon_side_asn

        __props__['availability_zone'] = availability_zone

        __props__['tags'] = tags

        __props__['vpc_id'] = vpc_id

        super(VpnGateway, __self__).__init__(
            'aws:ec2/vpnGateway:VpnGateway',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

