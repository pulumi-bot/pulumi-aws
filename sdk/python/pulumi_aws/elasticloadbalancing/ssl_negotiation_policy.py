# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class SslNegotiationPolicy(pulumi.CustomResource):
    """
    Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.
    """
    def __init__(__self__, __name__, __opts__=None, attributes=None, lb_port=None, load_balancer=None, name=None):
        """Create a SslNegotiationPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['attributes'] = attributes

        if not lb_port:
            raise TypeError('Missing required property lb_port')
        __props__['lbPort'] = lb_port

        if not load_balancer:
            raise TypeError('Missing required property load_balancer')
        __props__['loadBalancer'] = load_balancer

        __props__['name'] = name

        super(SslNegotiationPolicy, __self__).__init__(
            'aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy',
            __name__,
            __props__,
            __opts__)

