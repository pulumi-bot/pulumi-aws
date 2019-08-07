# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
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
    def __init__(__self__, resource_name, opts=None, api_stages=None, description=None, name=None, product_code=None, quota_settings=None, throttle_settings=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Usage Plan.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] api_stages: The associated API stages of the usage plan.
        :param pulumi.Input[str] description: The description of a usage plan.
        :param pulumi.Input[str] name: The name of the usage plan.
        :param pulumi.Input[str] product_code: The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.
        :param pulumi.Input[dict] quota_settings: The quota settings of the usage plan.
        :param pulumi.Input[dict] throttle_settings: The throttling limits of the usage plan.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_usage_plan.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['api_stages'] = api_stages
        __props__['description'] = description
        __props__['name'] = name
        __props__['product_code'] = product_code
        __props__['quota_settings'] = quota_settings
        __props__['throttle_settings'] = throttle_settings
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(UsagePlan, __self__).__init__(
            'aws:apigateway/usagePlan:UsagePlan',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

