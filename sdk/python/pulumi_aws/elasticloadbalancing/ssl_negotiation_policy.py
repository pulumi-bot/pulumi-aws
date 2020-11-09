# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SslNegotiationPolicy']

warnings.warn("""aws.elasticloadbalancing.SslNegotiationPolicy has been deprecated in favor of aws.elb.SslNegotiationPolicy""", DeprecationWarning)


class SslNegotiationPolicy(pulumi.CustomResource):
    warnings.warn("""aws.elasticloadbalancing.SslNegotiationPolicy has been deprecated in favor of aws.elb.SslNegotiationPolicy""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]]] = None,
                 lb_port: Optional[pulumi.Input[int]] = None,
                 load_balancer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        lb = aws.elb.LoadBalancer("lb",
            availability_zones=["us-east-1a"],
            listeners=[aws.elb.LoadBalancerListenerArgs(
                instance_port=8000,
                instance_protocol="https",
                lb_port=443,
                lb_protocol="https",
                ssl_certificate_id="arn:aws:iam::123456789012:server-certificate/certName",
            )])
        foo = aws.elb.SslNegotiationPolicy("foo",
            load_balancer=lb.id,
            lb_port=443,
            attributes=[
                aws.elb.SslNegotiationPolicyAttributeArgs(
                    name="Protocol-TLSv1",
                    value="false",
                ),
                aws.elb.SslNegotiationPolicyAttributeArgs(
                    name="Protocol-TLSv1.1",
                    value="false",
                ),
                aws.elb.SslNegotiationPolicyAttributeArgs(
                    name="Protocol-TLSv1.2",
                    value="true",
                ),
                aws.elb.SslNegotiationPolicyAttributeArgs(
                    name="Server-Defined-Cipher-Order",
                    value="true",
                ),
                aws.elb.SslNegotiationPolicyAttributeArgs(
                    name="ECDHE-RSA-AES128-GCM-SHA256",
                    value="true",
                ),
                aws.elb.SslNegotiationPolicyAttributeArgs(
                    name="AES128-GCM-SHA256",
                    value="true",
                ),
                aws.elb.SslNegotiationPolicyAttributeArgs(
                    name="EDH-RSA-DES-CBC3-SHA",
                    value="false",
                ),
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]] attributes: An SSL Negotiation policy attribute. Each has two properties:
        :param pulumi.Input[int] lb_port: The load balancer port to which the policy
               should be applied. This must be an active listener on the load
               balancer.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        :param pulumi.Input[str] name: The name of the attribute
        """
        pulumi.log.warn("SslNegotiationPolicy is deprecated: aws.elasticloadbalancing.SslNegotiationPolicy has been deprecated in favor of aws.elb.SslNegotiationPolicy")
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

            __props__['attributes'] = attributes
            if lb_port is None:
                raise TypeError("Missing required property 'lb_port'")
            __props__['lb_port'] = lb_port
            if load_balancer is None:
                raise TypeError("Missing required property 'load_balancer'")
            __props__['load_balancer'] = load_balancer
            __props__['name'] = name
        super(SslNegotiationPolicy, __self__).__init__(
            'aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]]] = None,
            lb_port: Optional[pulumi.Input[int]] = None,
            load_balancer: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'SslNegotiationPolicy':
        """
        Get an existing SslNegotiationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]] attributes: An SSL Negotiation policy attribute. Each has two properties:
        :param pulumi.Input[int] lb_port: The load balancer port to which the policy
               should be applied. This must be an active listener on the load
               balancer.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        :param pulumi.Input[str] name: The name of the attribute
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["lb_port"] = lb_port
        __props__["load_balancer"] = load_balancer
        __props__["name"] = name
        return SslNegotiationPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Sequence['outputs.SslNegotiationPolicyAttribute']]]:
        """
        An SSL Negotiation policy attribute. Each has two properties:
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="lbPort")
    def lb_port(self) -> pulumi.Output[int]:
        """
        The load balancer port to which the policy
        should be applied. This must be an active listener on the load
        balancer.
        """
        return pulumi.get(self, "lb_port")

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> pulumi.Output[str]:
        """
        The load balancer to which the policy
        should be attached.
        """
        return pulumi.get(self, "load_balancer")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the attribute
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

