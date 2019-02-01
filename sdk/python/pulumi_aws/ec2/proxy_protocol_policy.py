# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ProxyProtocolPolicy(pulumi.CustomResource):
    instance_ports: pulumi.Output[list]
    """
    List of instance ports to which the policy
    should be applied. This can be specified if the protocol is SSL or TCP.
    """
    load_balancer: pulumi.Output[str]
    """
    The load balancer to which the policy
    should be attached.
    """
    def __init__(__self__, __name__, __opts__=None, instance_ports=None, load_balancer=None):
        """
        Provides a proxy protocol policy, which allows an ELB to carry a client connection information to a backend.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] instance_ports: List of instance ports to which the policy
               should be applied. This can be specified if the protocol is SSL or TCP.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not instance_ports:
            raise TypeError('Missing required property instance_ports')
        __props__['instance_ports'] = instance_ports

        if not load_balancer:
            raise TypeError('Missing required property load_balancer')
        __props__['load_balancer'] = load_balancer

        super(ProxyProtocolPolicy, __self__).__init__(
            'aws:ec2/proxyProtocolPolicy:ProxyProtocolPolicy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

