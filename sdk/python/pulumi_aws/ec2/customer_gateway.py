# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CustomerGateway(pulumi.CustomResource):
    bgp_asn: pulumi.Output[float]
    """
    The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
    """
    ip_address: pulumi.Output[str]
    """
    The IP address of the gateway's Internet-routable external interface.
    """
    tags: pulumi.Output[dict]
    """
    Tags to apply to the gateway.
    """
    type: pulumi.Output[str]
    """
    The type of customer gateway. The only type AWS
    supports at this time is "ipsec.1".
    """
    def __init__(__self__, resource_name, opts=None, bgp_asn=None, ip_address=None, tags=None, type=None, __name__=None, __opts__=None):
        """
        Provides a customer gateway inside a VPC. These objects can be connected to VPN gateways via VPN connections, and allow you to establish tunnels between your network and the VPC.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bgp_asn: The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        :param pulumi.Input[str] ip_address: The IP address of the gateway's Internet-routable external interface.
        :param pulumi.Input[dict] tags: Tags to apply to the gateway.
        :param pulumi.Input[str] type: The type of customer gateway. The only type AWS
               supports at this time is "ipsec.1".

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/customer_gateway.html.markdown.
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

        if bgp_asn is None:
            raise TypeError("Missing required property 'bgp_asn'")
        __props__['bgp_asn'] = bgp_asn
        if ip_address is None:
            raise TypeError("Missing required property 'ip_address'")
        __props__['ip_address'] = ip_address
        __props__['tags'] = tags
        if type is None:
            raise TypeError("Missing required property 'type'")
        __props__['type'] = type
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(CustomerGateway, __self__).__init__(
            'aws:ec2/customerGateway:CustomerGateway',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

