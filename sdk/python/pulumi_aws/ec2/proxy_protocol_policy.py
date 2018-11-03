# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class ProxyProtocolPolicy(pulumi.CustomResource):
    """
    Provides a proxy protocol policy, which allows an ELB to carry a client connection information to a backend.
    """
    def __init__(__self__, __name__, __opts__=None, instance_ports=None, load_balancer=None):
        """Create a ProxyProtocolPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not instance_ports:
            raise TypeError('Missing required property instance_ports')
        __props__['instancePorts'] = instance_ports

        if not load_balancer:
            raise TypeError('Missing required property load_balancer')
        __props__['loadBalancer'] = load_balancer

        super(ProxyProtocolPolicy, __self__).__init__(
            'aws:ec2/proxyProtocolPolicy:ProxyProtocolPolicy',
            __name__,
            __props__,
            __opts__)

