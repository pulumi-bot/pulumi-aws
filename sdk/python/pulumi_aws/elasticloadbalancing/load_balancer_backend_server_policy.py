# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['LoadBalancerBackendServerPolicy']

warnings.warn("aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favor of aws.elb.LoadBalancerBackendServerPolicy", DeprecationWarning)


class LoadBalancerBackendServerPolicy(pulumi.CustomResource):
    warnings.warn("aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favor of aws.elb.LoadBalancerBackendServerPolicy", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 instance_port: Optional[pulumi.Input[int]] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 policy_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Attaches a load balancer policy to an ELB backend server.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_port: The instance port to apply the policy to.
        :param pulumi.Input[str] load_balancer_name: The load balancer to attach the policy to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_names: List of Policy Names to apply to the backend server.
        """
        pulumi.log.warn("LoadBalancerBackendServerPolicy is deprecated: aws.elasticloadbalancing.LoadBalancerBackendServerPolicy has been deprecated in favor of aws.elb.LoadBalancerBackendServerPolicy")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if instance_port is None:
                raise TypeError("Missing required property 'instance_port'")
            __props__['instance_port'] = instance_port
            if load_balancer_name is None:
                raise TypeError("Missing required property 'load_balancer_name'")
            __props__['load_balancer_name'] = load_balancer_name
            __props__['policy_names'] = policy_names
        super(LoadBalancerBackendServerPolicy, __self__).__init__(
            'aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            instance_port: Optional[pulumi.Input[int]] = None,
            load_balancer_name: Optional[pulumi.Input[str]] = None,
            policy_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'LoadBalancerBackendServerPolicy':
        """
        Get an existing LoadBalancerBackendServerPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] instance_port: The instance port to apply the policy to.
        :param pulumi.Input[str] load_balancer_name: The load balancer to attach the policy to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_names: List of Policy Names to apply to the backend server.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["instance_port"] = instance_port
        __props__["load_balancer_name"] = load_balancer_name
        __props__["policy_names"] = policy_names
        return LoadBalancerBackendServerPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="instancePort")
    def instance_port(self) -> pulumi.Output[int]:
        """
        The instance port to apply the policy to.
        """
        return pulumi.get(self, "instance_port")

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> pulumi.Output[str]:
        """
        The load balancer to attach the policy to.
        """
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter(name="policyNames")
    def policy_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of Policy Names to apply to the backend server.
        """
        return pulumi.get(self, "policy_names")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

