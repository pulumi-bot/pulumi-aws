# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ResolverRule']


class ResolverRule(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN (Amazon Resource Name) for the resolver rule.
    """

    domain_name: pulumi.Output[str] = pulumi.property("domainName")
    """
    DNS queries for this domain name are forwarded to the IP addresses that are specified using `target_ip`.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
    """

    owner_id: pulumi.Output[str] = pulumi.property("ownerId")
    """
    When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
    """

    resolver_endpoint_id: pulumi.Output[Optional[str]] = pulumi.property("resolverEndpointId")
    """
    The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `target_ip`.
    This argument should only be specified for `FORWARD` type rules.
    """

    rule_type: pulumi.Output[str] = pulumi.property("ruleType")
    """
    The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
    """

    share_status: pulumi.Output[str] = pulumi.property("shareStatus")
    """
    Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
    Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A map of tags to assign to the resource.
    """

    target_ips: pulumi.Output[Optional[List['outputs.ResolverRuleTargetIp']]] = pulumi.property("targetIps")
    """
    Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
    This argument should only be specified for `FORWARD` type rules.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resolver_endpoint_id: Optional[pulumi.Input[str]] = None,
                 rule_type: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_ips: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ResolverRuleTargetIpArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Route53 Resolver rule.

        ## Example Usage
        ### System rule

        ```python
        import pulumi
        import pulumi_aws as aws

        sys = aws.route53.ResolverRule("sys",
            domain_name="subdomain.example.com",
            rule_type="SYSTEM")
        ```
        ### Forward rule

        ```python
        import pulumi
        import pulumi_aws as aws

        fwd = aws.route53.ResolverRule("fwd",
            domain_name="example.com",
            rule_type="FORWARD",
            resolver_endpoint_id=aws_route53_resolver_endpoint["foo"]["id"],
            target_ips=[{
                "ip": "123.45.67.89",
            }],
            tags={
                "Environment": "Prod",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] domain_name: DNS queries for this domain name are forwarded to the IP addresses that are specified using `target_ip`.
        :param pulumi.Input[str] name: A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
        :param pulumi.Input[str] resolver_endpoint_id: The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `target_ip`.
               This argument should only be specified for `FORWARD` type rules.
        :param pulumi.Input[str] rule_type: The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ResolverRuleTargetIpArgs']]]] target_ips: Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
               This argument should only be specified for `FORWARD` type rules.
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

            if domain_name is None:
                raise TypeError("Missing required property 'domain_name'")
            __props__['domain_name'] = domain_name
            __props__['name'] = name
            __props__['resolver_endpoint_id'] = resolver_endpoint_id
            if rule_type is None:
                raise TypeError("Missing required property 'rule_type'")
            __props__['rule_type'] = rule_type
            __props__['tags'] = tags
            __props__['target_ips'] = target_ips
            __props__['arn'] = None
            __props__['owner_id'] = None
            __props__['share_status'] = None
        super(ResolverRule, __self__).__init__(
            'aws:route53/resolverRule:ResolverRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            resolver_endpoint_id: Optional[pulumi.Input[str]] = None,
            rule_type: Optional[pulumi.Input[str]] = None,
            share_status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_ips: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ResolverRuleTargetIpArgs']]]]] = None) -> 'ResolverRule':
        """
        Get an existing ResolverRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN (Amazon Resource Name) for the resolver rule.
        :param pulumi.Input[str] domain_name: DNS queries for this domain name are forwarded to the IP addresses that are specified using `target_ip`.
        :param pulumi.Input[str] name: A friendly name that lets you easily find a rule in the Resolver dashboard in the Route 53 console.
        :param pulumi.Input[str] owner_id: When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
        :param pulumi.Input[str] resolver_endpoint_id: The ID of the outbound resolver endpoint that you want to use to route DNS queries to the IP addresses that you specify using `target_ip`.
               This argument should only be specified for `FORWARD` type rules.
        :param pulumi.Input[str] rule_type: The rule type. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
        :param pulumi.Input[str] share_status: Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
               Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ResolverRuleTargetIpArgs']]]] target_ips: Configuration block(s) indicating the IPs that you want Resolver to forward DNS queries to (documented below).
               This argument should only be specified for `FORWARD` type rules.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["domain_name"] = domain_name
        __props__["name"] = name
        __props__["owner_id"] = owner_id
        __props__["resolver_endpoint_id"] = resolver_endpoint_id
        __props__["rule_type"] = rule_type
        __props__["share_status"] = share_status
        __props__["tags"] = tags
        __props__["target_ips"] = target_ips
        return ResolverRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

