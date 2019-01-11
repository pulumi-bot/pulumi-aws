# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventPermission(pulumi.CustomResource):
    """
    Provides a resource to create a CloudWatch Events permission to support cross-account events in the current account default event bus.
    """
    def __init__(__self__, __name__, __opts__=None, action=None, condition=None, principal=None, statement_id=None):
        """Create a EventPermission resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['action'] = action

        __props__['condition'] = condition

        if not principal:
            raise TypeError('Missing required property principal')
        __props__['principal'] = principal

        if not statement_id:
            raise TypeError('Missing required property statement_id')
        __props__['statement_id'] = statement_id

        super(EventPermission, __self__).__init__(
            'aws:cloudwatch/eventPermission:EventPermission',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

