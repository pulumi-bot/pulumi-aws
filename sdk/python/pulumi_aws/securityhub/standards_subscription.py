# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class StandardsSubscription(pulumi.CustomResource):
    standards_arn: pulumi.Output[str]
    """
    The ARN of a standard - see below.
    """
    def __init__(__self__, __name__, __opts__=None, standards_arn=None):
        """
        Subscribes to a Security Hub standard.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] standards_arn: The ARN of a standard - see below.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if standards_arn is None:
            raise TypeError('Missing required property standards_arn')
        __props__['standards_arn'] = standards_arn

        super(StandardsSubscription, __self__).__init__(
            'aws:securityhub/standardsSubscription:StandardsSubscription',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

