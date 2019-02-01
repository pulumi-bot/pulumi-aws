# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UsagePlan(pulumi.CustomResource):
    api_stages: pulumi.Output[list]
    """
    The associated API stages of the usage plan.
    """
    description: pulumi.Output[str]
    """
    The description of a usage plan.
    """
    name: pulumi.Output[str]
    """
    The name of the usage plan.
    """
    product_code: pulumi.Output[str]
    """
    The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
    """
    quota_settings: pulumi.Output[dict]
    """
    The quota settings of the usage plan.
    """
    throttle_settings: pulumi.Output[dict]
    """
    The throttling limits of the usage plan.
    """
    def __init__(__self__, __name__, __opts__=None, api_stages=None, description=None, name=None, product_code=None, quota_settings=None, throttle_settings=None):
        """
        Provides an API Gateway Usage Plan.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] api_stages: The associated API stages of the usage plan.
        :param pulumi.Input[str] description: The description of a usage plan.
        :param pulumi.Input[str] name: The name of the usage plan.
        :param pulumi.Input[str] product_code: The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        :param pulumi.Input[dict] quota_settings: The quota settings of the usage plan.
        :param pulumi.Input[dict] throttle_settings: The throttling limits of the usage plan.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['api_stages'] = api_stages

        __props__['description'] = description

        __props__['name'] = name

        __props__['product_code'] = product_code

        __props__['quota_settings'] = quota_settings

        __props__['throttle_settings'] = throttle_settings

        super(UsagePlan, __self__).__init__(
            'aws:apigateway/usagePlan:UsagePlan',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

