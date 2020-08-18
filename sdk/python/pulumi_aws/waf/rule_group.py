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

__all__ = ['RuleGroup']


class RuleGroup(pulumi.CustomResource):
    activated_rules: pulumi.Output[Optional[List['outputs.RuleGroupActivatedRule']]] = pulumi.property("activatedRules")
    """
    A list of activated rules, see below
    """

    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN of the WAF rule group.
    """

    metric_name: pulumi.Output[str] = pulumi.property("metricName")
    """
    A friendly name for the metrics from the rule group
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    A friendly name of the rule group
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    Key-value map of resource tags
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activated_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a WAF Rule Group Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_rule = aws.waf.Rule("exampleRule", metric_name="example")
        example_rule_group = aws.waf.RuleGroup("exampleRuleGroup",
            metric_name="example",
            activated_rules=[{
                "action": {
                    "type": "COUNT",
                },
                "priority": 50,
                "rule_id": example_rule.id,
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]] activated_rules: A list of activated rules, see below
        :param pulumi.Input[str] metric_name: A friendly name for the metrics from the rule group
        :param pulumi.Input[str] name: A friendly name of the rule group
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
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

            __props__['activated_rules'] = activated_rules
            if metric_name is None:
                raise TypeError("Missing required property 'metric_name'")
            __props__['metric_name'] = metric_name
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(RuleGroup, __self__).__init__(
            'aws:waf/ruleGroup:RuleGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            activated_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            metric_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'RuleGroup':
        """
        Get an existing RuleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]] activated_rules: A list of activated rules, see below
        :param pulumi.Input[str] arn: The ARN of the WAF rule group.
        :param pulumi.Input[str] metric_name: A friendly name for the metrics from the rule group
        :param pulumi.Input[str] name: A friendly name of the rule group
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["activated_rules"] = activated_rules
        __props__["arn"] = arn
        __props__["metric_name"] = metric_name
        __props__["name"] = name
        __props__["tags"] = tags
        return RuleGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

