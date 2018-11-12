# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class UsagePlanKey(pulumi.CustomResource):
    """
    Provides an API Gateway Usage Plan Key.
    """
    def __init__(__self__, __name__, __opts__=None, key_id=None, key_type=None, usage_plan_id=None):
        """Create a UsagePlanKey resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not key_id:
            raise TypeError('Missing required property key_id')
        __props__['key_id'] = key_id

        if not key_type:
            raise TypeError('Missing required property key_type')
        __props__['key_type'] = key_type

        if not usage_plan_id:
            raise TypeError('Missing required property usage_plan_id')
        __props__['usage_plan_id'] = usage_plan_id

        __props__['name'] = None
        __props__['value'] = None

        super(UsagePlanKey, __self__).__init__(
            'aws:apigateway/usagePlanKey:UsagePlanKey',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

