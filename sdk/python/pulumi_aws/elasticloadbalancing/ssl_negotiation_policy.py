# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SslNegotiationPolicyArgs', 'SslNegotiationPolicy']

@pulumi.input_type
class SslNegotiationPolicyArgs:
    def __init__(__self__, *,
                 lb_port: pulumi.Input[int],
                 load_balancer: pulumi.Input[str],
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SslNegotiationPolicy resource.
        :param pulumi.Input[int] lb_port: The load balancer port to which the policy
               should be applied. This must be an active listener on the load
               balancer.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        :param pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]] attributes: An SSL Negotiation policy attribute. Each has two properties:
        :param pulumi.Input[str] name: The name of the attribute
        """
        pulumi.set(__self__, "lb_port", lb_port)
        pulumi.set(__self__, "load_balancer", load_balancer)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="lbPort")
    def lb_port(self) -> pulumi.Input[int]:
        """
        The load balancer port to which the policy
        should be applied. This must be an active listener on the load
        balancer.
        """
        return pulumi.get(self, "lb_port")

    @lb_port.setter
    def lb_port(self, value: pulumi.Input[int]):
        pulumi.set(self, "lb_port", value)

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> pulumi.Input[str]:
        """
        The load balancer to which the policy
        should be attached.
        """
        return pulumi.get(self, "load_balancer")

    @load_balancer.setter
    def load_balancer(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]]]:
        """
        An SSL Negotiation policy attribute. Each has two properties:
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the attribute
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _SslNegotiationPolicyState:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]]] = None,
                 lb_port: Optional[pulumi.Input[int]] = None,
                 load_balancer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SslNegotiationPolicy resources.
        :param pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]] attributes: An SSL Negotiation policy attribute. Each has two properties:
        :param pulumi.Input[int] lb_port: The load balancer port to which the policy
               should be applied. This must be an active listener on the load
               balancer.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        :param pulumi.Input[str] name: The name of the attribute
        """
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if lb_port is not None:
            pulumi.set(__self__, "lb_port", lb_port)
        if load_balancer is not None:
            pulumi.set(__self__, "load_balancer", load_balancer)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]]]:
        """
        An SSL Negotiation policy attribute. Each has two properties:
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SslNegotiationPolicyAttributeArgs']]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter(name="lbPort")
    def lb_port(self) -> Optional[pulumi.Input[int]]:
        """
        The load balancer port to which the policy
        should be applied. This must be an active listener on the load
        balancer.
        """
        return pulumi.get(self, "lb_port")

    @lb_port.setter
    def lb_port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "lb_port", value)

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> Optional[pulumi.Input[str]]:
        """
        The load balancer to which the policy
        should be attached.
        """
        return pulumi.get(self, "load_balancer")

    @load_balancer.setter
    def load_balancer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "load_balancer", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the attribute
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


warnings.warn("""aws.elasticloadbalancing.SslNegotiationPolicy has been deprecated in favor of aws.elb.SslNegotiationPolicy""", DeprecationWarning)


class SslNegotiationPolicy(pulumi.CustomResource):
    warnings.warn("""aws.elasticloadbalancing.SslNegotiationPolicy has been deprecated in favor of aws.elb.SslNegotiationPolicy""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]]] = None,
                 lb_port: Optional[pulumi.Input[int]] = None,
                 load_balancer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
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

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]] attributes: An SSL Negotiation policy attribute. Each has two properties:
        :param pulumi.Input[int] lb_port: The load balancer port to which the policy
               should be applied. This must be an active listener on the load
               balancer.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        :param pulumi.Input[str] name: The name of the attribute
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: SslNegotiationPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        :param str resource_name_: The name of the resource.
        :param SslNegotiationPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SslNegotiationPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]]] = None,
                 lb_port: Optional[pulumi.Input[int]] = None,
                 load_balancer: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""SslNegotiationPolicy is deprecated: aws.elasticloadbalancing.SslNegotiationPolicy has been deprecated in favor of aws.elb.SslNegotiationPolicy""")
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SslNegotiationPolicyArgs.__new__(SslNegotiationPolicyArgs)

            __props__.__dict__["attributes"] = attributes
            if lb_port is None and not opts.urn:
                raise TypeError("Missing required property 'lb_port'")
            __props__.__dict__["lb_port"] = lb_port
            if load_balancer is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer'")
            __props__.__dict__["load_balancer"] = load_balancer
            __props__.__dict__["name"] = name
        super(SslNegotiationPolicy, __self__).__init__(
            'aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SslNegotiationPolicyAttributeArgs']]]]] = None,
            lb_port: Optional[pulumi.Input[int]] = None,
            load_balancer: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'SslNegotiationPolicy':
        """
        Get an existing SslNegotiationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
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

        __props__ = _SslNegotiationPolicyState.__new__(_SslNegotiationPolicyState)

        __props__.__dict__["attributes"] = attributes
        __props__.__dict__["lb_port"] = lb_port
        __props__.__dict__["load_balancer"] = load_balancer
        __props__.__dict__["name"] = name
        return SslNegotiationPolicy(resource_name_, opts=opts, __props__=__props__)

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

