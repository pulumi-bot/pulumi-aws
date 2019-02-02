# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IpSet(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the WAF IPSet.
    """
    ip_set_descriptors: pulumi.Output[list]
    """
    One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
    """
    name: pulumi.Output[str]
    """
    The name or description of the IPSet.
    """
    def __init__(__self__, __name__, __opts__=None, ip_set_descriptors=None, name=None):
        """
        Provides a WAF IPSet Resource
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] ip_set_descriptors: One or more pairs specifying the IP address type (IPV4 or IPV6) and the IP address range (in CIDR format) from which web requests originate.
        :param pulumi.Input[str] name: The name or description of the IPSet.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['ip_set_descriptors'] = ip_set_descriptors

        __props__['name'] = name

        __props__['arn'] = None

        super(IpSet, __self__).__init__(
            'aws:waf/ipSet:IpSet',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

