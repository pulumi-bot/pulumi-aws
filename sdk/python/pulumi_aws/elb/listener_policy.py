# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ListenerPolicy']


class ListenerPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 load_balancer_name: Optional[pulumi.Input[str]] = None,
                 load_balancer_port: Optional[pulumi.Input[int]] = None,
                 policy_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Attaches a load balancer policy to an ELB Listener.

        ## Example Usage
        ### Custom Policy

        ```python
        import pulumi
        import pulumi_aws as aws

        wu_tang = aws.elb.LoadBalancer("wu-tang",
            availability_zones=["us-east-1a"],
            listeners=[aws.elb.LoadBalancerListenerArgs(
                instance_port=443,
                instance_protocol="http",
                lb_port=443,
                lb_protocol="https",
                ssl_certificate_id="arn:aws:iam::000000000000:server-certificate/wu-tang.net",
            )],
            tags={
                "Name": "wu-tang",
            })
        wu_tang_ssl = aws.elb.LoadBalancerPolicy("wu-tang-ssl",
            load_balancer_name=wu_tang.name,
            policy_name="wu-tang-ssl",
            policy_type_name="SSLNegotiationPolicyType",
            policy_attributes=[
                aws.elb.LoadBalancerPolicyPolicyAttributeArgs(
                    name="ECDHE-ECDSA-AES128-GCM-SHA256",
                    value="true",
                ),
                aws.elb.LoadBalancerPolicyPolicyAttributeArgs(
                    name="Protocol-TLSv1.2",
                    value="true",
                ),
            ])
        wu_tang_listener_policies_443 = aws.elb.ListenerPolicy("wu-tang-listener-policies-443",
            load_balancer_name=wu_tang.name,
            load_balancer_port=443,
            policy_names=[wu_tang_ssl.policy_name])
        ```

        This example shows how to customize the TLS settings of an HTTPS listener.
        ### AWS Predefined Security Policy

        ```python
        import pulumi
        import pulumi_aws as aws

        wu_tang = aws.elb.LoadBalancer("wu-tang",
            availability_zones=["us-east-1a"],
            listeners=[aws.elb.LoadBalancerListenerArgs(
                instance_port=443,
                instance_protocol="http",
                lb_port=443,
                lb_protocol="https",
                ssl_certificate_id="arn:aws:iam::000000000000:server-certificate/wu-tang.net",
            )],
            tags={
                "Name": "wu-tang",
            })
        wu_tang_ssl_tls_1_1 = aws.elb.LoadBalancerPolicy("wu-tang-ssl-tls-1-1",
            load_balancer_name=wu_tang.name,
            policy_name="wu-tang-ssl",
            policy_type_name="SSLNegotiationPolicyType",
            policy_attributes=[aws.elb.LoadBalancerPolicyPolicyAttributeArgs(
                name="Reference-Security-Policy",
                value="ELBSecurityPolicy-TLS-1-1-2017-01",
            )])
        wu_tang_listener_policies_443 = aws.elb.ListenerPolicy("wu-tang-listener-policies-443",
            load_balancer_name=wu_tang.name,
            load_balancer_port=443,
            policy_names=[wu_tang_ssl_tls_1_1.policy_name])
        ```

        This example shows how to add a [Predefined Security Policy for ELBs](https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-security-policy-table.html)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] load_balancer_name: The load balancer to attach the policy to.
        :param pulumi.Input[int] load_balancer_port: The load balancer listener port to apply the policy to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_names: List of Policy Names to apply to the backend server.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if load_balancer_name is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer_name'")
            __props__['load_balancer_name'] = load_balancer_name
            if load_balancer_port is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer_port'")
            __props__['load_balancer_port'] = load_balancer_port
            __props__['policy_names'] = policy_names
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:elasticloadbalancing/listenerPolicy:ListenerPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ListenerPolicy, __self__).__init__(
            'aws:elb/listenerPolicy:ListenerPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            load_balancer_name: Optional[pulumi.Input[str]] = None,
            load_balancer_port: Optional[pulumi.Input[int]] = None,
            policy_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'ListenerPolicy':
        """
        Get an existing ListenerPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] load_balancer_name: The load balancer to attach the policy to.
        :param pulumi.Input[int] load_balancer_port: The load balancer listener port to apply the policy to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_names: List of Policy Names to apply to the backend server.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["load_balancer_name"] = load_balancer_name
        __props__["load_balancer_port"] = load_balancer_port
        __props__["policy_names"] = policy_names
        return ListenerPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="loadBalancerName")
    def load_balancer_name(self) -> pulumi.Output[str]:
        """
        The load balancer to attach the policy to.
        """
        return pulumi.get(self, "load_balancer_name")

    @property
    @pulumi.getter(name="loadBalancerPort")
    def load_balancer_port(self) -> pulumi.Output[int]:
        """
        The load balancer listener port to apply the policy to.
        """
        return pulumi.get(self, "load_balancer_port")

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

