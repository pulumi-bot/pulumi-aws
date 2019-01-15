# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class LoadBalancerBackendServerPolicy(pulumi.CustomResource):
    """
    Attaches a load balancer policy to an ELB backend server.
    
    """
    def __init__(__self__, __name__, __opts__=None, instance_port=None, load_balancer_name=None, policy_names=None):
        """Create a LoadBalancerBackendServerPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not instance_port:
            raise TypeError('Missing required property instance_port')
        __props__['instance_port'] = instance_port

        if not load_balancer_name:
            raise TypeError('Missing required property load_balancer_name')
        __props__['load_balancer_name'] = load_balancer_name

        __props__['policy_names'] = policy_names

        super(LoadBalancerBackendServerPolicy, __self__).__init__(
            'aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

