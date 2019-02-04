# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Rule(pulumi.CustomResource):
    metric_name: pulumi.Output[str]
    """
    The name or description for the Amazon CloudWatch metric of this rule.
    """
    name: pulumi.Output[str]
    """
    The name or description of the rule.
    """
    predicates: pulumi.Output[list]
    """
    The objects to include in a rule.
    """
    def __init__(__self__, __name__, __opts__=None, metric_name=None, name=None, predicates=None):
        """
        Provides an WAF Regional Rule Resource for use with Application Load Balancer.
        
        ## Nested Fields
        
        ### `predicate`
        
        See the [WAF Documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_Predicate.html) for more information.
        
        #### Arguments
        
        * `type` - (Required) The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`
        * `data_id` - (Required) The unique identifier of a predicate, such as the ID of a `ByteMatchSet` or `IPSet`.
        * `negated` - (Required) Whether to use the settings or the negated settings that you specified in the objects.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this rule.
        :param pulumi.Input[str] name: The name or description of the rule.
        :param pulumi.Input[list] predicates: The objects to include in a rule.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if metric_name is None:
            raise TypeError('Missing required property metric_name')
        __props__['metric_name'] = metric_name

        __props__['name'] = name

        __props__['predicates'] = predicates

        super(Rule, __self__).__init__(
            'aws:wafregional/rule:Rule',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

