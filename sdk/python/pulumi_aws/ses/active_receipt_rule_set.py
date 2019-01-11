# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ActiveReceiptRuleSet(pulumi.CustomResource):
    """
    Provides a resource to designate the active SES receipt rule set
    """
    def __init__(__self__, __name__, __opts__=None, rule_set_name=None):
        """Create a ActiveReceiptRuleSet resource with the given unique name, props, and options."""
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

