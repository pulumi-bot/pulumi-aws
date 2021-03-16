# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RuleGroupArgs', 'RuleGroup']

@pulumi.input_type
class RuleGroupArgs:
    def __init__(__self__, *,
                 metric_name: pulumi.Input[str],
                 activated_rules: Optional[pulumi.Input[Sequence[pulumi.Input['RuleGroupActivatedRuleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RuleGroup resource.
        :param pulumi.Input[str] metric_name: A friendly name for the metrics from the rule group
        :param pulumi.Input[Sequence[pulumi.Input['RuleGroupActivatedRuleArgs']]] activated_rules: A list of activated rules, see below
        :param pulumi.Input[str] name: A friendly name of the rule group
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        pulumi.set(__self__, "metric_name", metric_name)
        if activated_rules is not None:
            pulumi.set(__self__, "activated_rules", activated_rules)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Input[str]:
        """
        A friendly name for the metrics from the rule group
        """
        return pulumi.get(self, "metric_name")

    @metric_name.setter
    def metric_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "metric_name", value)

    @property
    @pulumi.getter(name="activatedRules")
    def activated_rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RuleGroupActivatedRuleArgs']]]]:
        """
        A list of activated rules, see below
        """
        return pulumi.get(self, "activated_rules")

    @activated_rules.setter
    def activated_rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RuleGroupActivatedRuleArgs']]]]):
        pulumi.set(self, "activated_rules", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A friendly name of the rule group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class RuleGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activated_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]]] = None,
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
            activated_rules=[aws.waf.RuleGroupActivatedRuleArgs(
                action=aws.waf.RuleGroupActivatedRuleActionArgs(
                    type="COUNT",
                ),
                priority=50,
                rule_id=example_rule.id,
            )])
        ```

        ## Import

        WAF Rule Group can be imported using the id, e.g.

        ```sh
         $ pulumi import aws:waf/ruleGroup:RuleGroup example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]] activated_rules: A list of activated rules, see below
        :param pulumi.Input[str] metric_name: A friendly name for the metrics from the rule group
        :param pulumi.Input[str] name: A friendly name of the rule group
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RuleGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Rule Group Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_rule = aws.waf.Rule("exampleRule", metric_name="example")
        example_rule_group = aws.waf.RuleGroup("exampleRuleGroup",
            metric_name="example",
            activated_rules=[aws.waf.RuleGroupActivatedRuleArgs(
                action=aws.waf.RuleGroupActivatedRuleActionArgs(
                    type="COUNT",
                ),
                priority=50,
                rule_id=example_rule.id,
            )])
        ```

        ## Import

        WAF Rule Group can be imported using the id, e.g.

        ```sh
         $ pulumi import aws:waf/ruleGroup:RuleGroup example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param RuleGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RuleGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activated_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]]] = None,
                 metric_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            if metric_name is None and not opts.urn:
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
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            activated_rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            metric_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'RuleGroup':
        """
        Get an existing RuleGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RuleGroupActivatedRuleArgs']]]] activated_rules: A list of activated rules, see below
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

    @property
    @pulumi.getter(name="activatedRules")
    def activated_rules(self) -> pulumi.Output[Optional[Sequence['outputs.RuleGroupActivatedRule']]]:
        """
        A list of activated rules, see below
        """
        return pulumi.get(self, "activated_rules")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the WAF rule group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> pulumi.Output[str]:
        """
        A friendly name for the metrics from the rule group
        """
        return pulumi.get(self, "metric_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A friendly name of the rule group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

