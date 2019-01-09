# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventRule(pulumi.CustomResource):
    """
    Provides a CloudWatch Event Rule resource.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, event_pattern=None, is_enabled=None, name=None, name_prefix=None, role_arn=None, schedule_expression=None):
        """Create a EventRule resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['event_pattern'] = event_pattern

        __props__['is_enabled'] = is_enabled

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        __props__['role_arn'] = role_arn

        __props__['schedule_expression'] = schedule_expression

        __props__['arn'] = None

        super(EventRule, __self__).__init__(
            'aws:cloudwatch/eventRule:EventRule',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

