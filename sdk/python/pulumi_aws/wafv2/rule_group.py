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

__all__ = ['RuleGroup']


class RuleGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 capacity: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupRuleArgs']]]]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 visibility_config: Optional[pulumi.Input[pulumi.InputType['RuleGroupVisibilityConfigArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a WAFv2 Rule Group resource.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] capacity: The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
        :param pulumi.Input[str] description: A friendly description of the rule group.
        :param pulumi.Input[str] name: A friendly name of the rule group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupRuleArgs']]]] rules: The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
        :param pulumi.Input[pulumi.InputType['RuleGroupVisibilityConfigArgs']] visibility_config: Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
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

            if capacity is None:
                raise TypeError("Missing required property 'capacity'")
            __props__['capacity'] = capacity
            __props__['description'] = description
            __props__['name'] = name
            __props__['rules'] = rules
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['tags'] = tags
            if visibility_config is None:
                raise TypeError("Missing required property 'visibility_config'")
            __props__['visibility_config'] = visibility_config
            __props__['arn'] = None
            __props__['lock_token'] = None
        super(RuleGroup, __self__).__init__(
            'aws:wafv2/ruleGroup:RuleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            capacity: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lock_token: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupRuleArgs']]]]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            visibility_config: Optional[pulumi.Input[pulumi.InputType['RuleGroupVisibilityConfigArgs']]] = None) -> 'RuleGroup':
        """
        Get an existing RuleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the IP Set that this statement references.
        :param pulumi.Input[int] capacity: The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
        :param pulumi.Input[str] description: A friendly description of the rule group.
        :param pulumi.Input[str] name: A friendly name of the rule group.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupRuleArgs']]]] rules: The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
        :param pulumi.Input[pulumi.InputType['RuleGroupVisibilityConfigArgs']] visibility_config: Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["capacity"] = capacity
        __props__["description"] = description
        __props__["lock_token"] = lock_token
        __props__["name"] = name
        __props__["rules"] = rules
        __props__["scope"] = scope
        __props__["tags"] = tags
        __props__["visibility_config"] = visibility_config
        return RuleGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the IP Set that this statement references.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Output[int]:
        """
        The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
        """
        return pulumi.get(self, "capacity")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly description of the rule group.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lockToken")
    def lock_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "lock_token")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A friendly name of the rule group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.RuleGroupRule']]]:
        """
        The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An array of key:value pairs to associate with the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="visibilityConfig")
    def visibility_config(self) -> pulumi.Output['outputs.RuleGroupVisibilityConfig']:
        """
        Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
        """
        return pulumi.get(self, "visibility_config")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

