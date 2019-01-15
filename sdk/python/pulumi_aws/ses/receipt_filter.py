# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ReceiptFilter(pulumi.CustomResource):
    cidr: pulumi.Output[str]
    """
    The IP address or address range to filter, in CIDR notation
    """
    name: pulumi.Output[str]
    """
    The name of the filter
    """
    policy: pulumi.Output[str]
    """
    Block or Allow
    """
    def __init__(__self__, __name__, __opts__=None, cidr=None, name=None, policy=None):
        """
        Provides an SES receipt filter resource
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] cidr: The IP address or address range to filter, in CIDR notation
        :param pulumi.Input[str] name: The name of the filter
        :param pulumi.Input[str] policy: Block or Allow
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not cidr:
            raise TypeError('Missing required property cidr')
        __props__['cidr'] = cidr

        __props__['name'] = name

        if not policy:
            raise TypeError('Missing required property policy')
        __props__['policy'] = policy

        super(ReceiptFilter, __self__).__init__(
            'aws:ses/receiptFilter:ReceiptFilter',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

