# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class SslNegotiationPolicy(pulumi.CustomResource):
    attributes: pulumi.Output[list]
    """
    An SSL Negotiation policy attribute. Each has two properties:

      * `name` (`str`) - The name of the attribute
      * `value` (`str`) - The value of the attribute
    """
    lb_port: pulumi.Output[float]
    """
    The load balancer port to which the policy
    should be applied. This must be an active listener on the load
    balancer.
    """
    load_balancer: pulumi.Output[str]
    """
    The load balancer to which the policy
    should be attached.
    """
    name: pulumi.Output[str]
    """
    The name of the attribute
    """
    def __init__(__self__, resource_name, opts=None, attributes=None, lb_port=None, load_balancer=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        lb = aws.elb.LoadBalancer("lb",
            availability_zones=["us-east-1a"],
            listeners=[{
                "instance_port": 8000,
                "instanceProtocol": "https",
                "lb_port": 443,
                "lbProtocol": "https",
                "sslCertificateId": "arn:aws:iam::123456789012:server-certificate/certName",
            }])
        foo = aws.elb.SslNegotiationPolicy("foo",
            attributes=[
                {
                    "name": "Protocol-TLSv1",
                    "value": "false",
                },
                {
                    "name": "Protocol-TLSv1.1",
                    "value": "false",
                },
                {
                    "name": "Protocol-TLSv1.2",
                    "value": "true",
                },
                {
                    "name": "Server-Defined-Cipher-Order",
                    "value": "true",
                },
                {
                    "name": "ECDHE-RSA-AES128-GCM-SHA256",
                    "value": "true",
                },
                {
                    "name": "AES128-GCM-SHA256",
                    "value": "true",
                },
                {
                    "name": "EDH-RSA-DES-CBC3-SHA",
                    "value": "false",
                },
            ],
            lb_port=443,
            load_balancer=lb.id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] attributes: An SSL Negotiation policy attribute. Each has two properties:
        :param pulumi.Input[float] lb_port: The load balancer port to which the policy
               should be applied. This must be an active listener on the load
               balancer.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        :param pulumi.Input[str] name: The name of the attribute

        The **attributes** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the attribute
          * `value` (`pulumi.Input[str]`) - The value of the attribute
        """
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
            opts.version = utilities.get_version()
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
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SslNegotiationPolicy, __self__).__init__(
            'aws:elb/sslNegotiationPolicy:SslNegotiationPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, attributes=None, lb_port=None, load_balancer=None, name=None):
        """
        Get an existing SslNegotiationPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] attributes: An SSL Negotiation policy attribute. Each has two properties:
        :param pulumi.Input[float] lb_port: The load balancer port to which the policy
               should be applied. This must be an active listener on the load
               balancer.
        :param pulumi.Input[str] load_balancer: The load balancer to which the policy
               should be attached.
        :param pulumi.Input[str] name: The name of the attribute

        The **attributes** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the attribute
          * `value` (`pulumi.Input[str]`) - The value of the attribute
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["lb_port"] = lb_port
        __props__["load_balancer"] = load_balancer
        __props__["name"] = name
        return SslNegotiationPolicy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
