# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ListenerRule(pulumi.CustomResource):
    actions: pulumi.Output[list]
    """
    An Action block. Action blocks are documented below.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the rule (matches `id`)
    """
    conditions: pulumi.Output[list]
    """
    A Condition block. Condition blocks are documented below.
    """
    listener_arn: pulumi.Output[str]
    """
    The ARN of the listener to which to attach the rule.
    """
    priority: pulumi.Output[float]
    """
    The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
    """
    def __init__(__self__, resource_name, opts=None, actions=None, conditions=None, listener_arn=None, priority=None, __name__=None, __opts__=None):
        """
        Provides a Load Balancer Listener Rule resource.
        
        > **Note:** `aws_alb_listener_rule` is known as `aws_lb_listener_rule`. The functionality is identical.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[list] conditions: A Condition block. Condition blocks are documented below.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the rule.
        :param pulumi.Input[float] priority: The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener_rule.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if actions is None:
            raise TypeError("Missing required property 'actions'")
        __props__['actions'] = actions

        if conditions is None:
            raise TypeError("Missing required property 'conditions'")
        __props__['conditions'] = conditions

        if listener_arn is None:
            raise TypeError("Missing required property 'listener_arn'")
        __props__['listener_arn'] = listener_arn

        __props__['priority'] = priority

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:elasticloadbalancingv2/listenerRule:ListenerRule")])
        opts = alias_opts if opts is None else opts.merge(alias_opts)
        super(ListenerRule, __self__).__init__(
            'aws:lb/listenerRule:ListenerRule',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

