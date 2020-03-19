# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetResolverRuleResult:
    """
    A collection of values returned by getResolverRule.
    """
    def __init__(__self__, arn=None, domain_name=None, id=None, name=None, owner_id=None, resolver_endpoint_id=None, resolver_rule_id=None, rule_type=None, share_status=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The ARN (Amazon Resource Name) for the resolver rule.
        """
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        __self__.domain_name = domain_name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        """
        When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
        """
        if resolver_endpoint_id and not isinstance(resolver_endpoint_id, str):
            raise TypeError("Expected argument 'resolver_endpoint_id' to be a str")
        __self__.resolver_endpoint_id = resolver_endpoint_id
        if resolver_rule_id and not isinstance(resolver_rule_id, str):
            raise TypeError("Expected argument 'resolver_rule_id' to be a str")
        __self__.resolver_rule_id = resolver_rule_id
        if rule_type and not isinstance(rule_type, str):
            raise TypeError("Expected argument 'rule_type' to be a str")
        __self__.rule_type = rule_type
        if share_status and not isinstance(share_status, str):
            raise TypeError("Expected argument 'share_status' to be a str")
        __self__.share_status = share_status
        """
        Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
        Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A mapping of tags assigned to the resolver rule.
        """
class AwaitableGetResolverRuleResult(GetResolverRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResolverRuleResult(
            arn=self.arn,
            domain_name=self.domain_name,
            id=self.id,
            name=self.name,
            owner_id=self.owner_id,
            resolver_endpoint_id=self.resolver_endpoint_id,
            resolver_rule_id=self.resolver_rule_id,
            rule_type=self.rule_type,
            share_status=self.share_status,
            tags=self.tags)

def get_resolver_rule(domain_name=None,name=None,resolver_endpoint_id=None,resolver_rule_id=None,rule_type=None,tags=None,opts=None):
    """
    `route53.ResolverRule` provides details about a specific Route53 Resolver rule.

    {{% examples %}}
    {{% /examples %}}

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/route53_resolver_rule.html.markdown.


    :param str domain_name: The domain name the desired resolver rule forwards DNS queries for. Conflicts with `resolver_rule_id`.
    :param str name: The friendly name of the desired resolver rule. Conflicts with `resolver_rule_id`.
    :param str resolver_endpoint_id: The ID of the outbound resolver endpoint of the desired resolver rule. Conflicts with `resolver_rule_id`.
    :param str resolver_rule_id: The ID of the desired resolver rule. Conflicts with `domain_name`, `name`, `resolver_endpoint_id` and `rule_type`.
    :param str rule_type: The rule type of the desired resolver rule. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`. Conflicts with `resolver_rule_id`.
    """
    __args__ = dict()


    __args__['domainName'] = domain_name
    __args__['name'] = name
    __args__['resolverEndpointId'] = resolver_endpoint_id
    __args__['resolverRuleId'] = resolver_rule_id
    __args__['ruleType'] = rule_type
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:route53/getResolverRule:getResolverRule', __args__, opts=opts).value

    return AwaitableGetResolverRuleResult(
        arn=__ret__.get('arn'),
        domain_name=__ret__.get('domainName'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        owner_id=__ret__.get('ownerId'),
        resolver_endpoint_id=__ret__.get('resolverEndpointId'),
        resolver_rule_id=__ret__.get('resolverRuleId'),
        rule_type=__ret__.get('ruleType'),
        share_status=__ret__.get('shareStatus'),
        tags=__ret__.get('tags'))
