# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class LoadBalancerPolicy(pulumi.CustomResource):
    load_balancer_name: pulumi.Output[str]
    """
    The load balancer on which the policy is defined.
    """
    policy_attributes: pulumi.Output[list]
    """
    Policy attribute to apply to the policy.
    """
    policy_name: pulumi.Output[str]
    """
    The name of the load balancer policy.
    """
    policy_type_name: pulumi.Output[str]
    """
    The policy type.
    """
    def __init__(__self__, resource_name, opts=None, load_balancer_name=None, policy_attributes=None, policy_name=None, policy_type_name=None, __name__=None, __opts__=None):
        """
        Provides a load balancer policy, which can be attached to an ELB listener or backend server.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] load_balancer_name: The load balancer on which the policy is defined.
        :param pulumi.Input[list] policy_attributes: Policy attribute to apply to the policy.
        :param pulumi.Input[str] policy_name: The name of the load balancer policy.
        :param pulumi.Input[str] policy_type_name: The policy type.
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

        if load_balancer_name is None:
            raise TypeError("Missing required property 'load_balancer_name'")
        __props__['load_balancer_name'] = load_balancer_name

        __props__['policy_attributes'] = policy_attributes

        if policy_name is None:
            raise TypeError("Missing required property 'policy_name'")
        __props__['policy_name'] = policy_name

        if policy_type_name is None:
            raise TypeError("Missing required property 'policy_type_name'")
        __props__['policy_type_name'] = policy_type_name

        super(LoadBalancerPolicy, __self__).__init__(
            'aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

