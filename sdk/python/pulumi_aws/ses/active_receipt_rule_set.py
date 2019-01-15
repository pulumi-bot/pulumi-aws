# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ActiveReceiptRuleSet(pulumi.CustomResource):
    rule_set_name: pulumi.Output[str]
    """
    The name of the rule set
    """
    def __init__(__self__, __name__, __opts__=None, rule_set_name=None):
        """
        Provides a resource to designate the active SES receipt rule set
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] rule_set_name: The name of the rule set
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not rule_set_name:
            raise TypeError('Missing required property rule_set_name')
        __props__['rule_set_name'] = rule_set_name

        super(ActiveReceiptRuleSet, __self__).__init__(
            'aws:ses/activeReceiptRuleSet:ActiveReceiptRuleSet',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

