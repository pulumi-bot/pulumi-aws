# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UsagePlanKey(pulumi.CustomResource):
    key_id: pulumi.Output[str]
    """
    The identifier of the API key resource.
    """
    key_type: pulumi.Output[str]
    """
    The type of the API key resource. Currently, the valid key type is API_KEY.
    """
    name: pulumi.Output[str]
    """
    The name of a usage plan key.
    """
    usage_plan_id: pulumi.Output[str]
    """
    The Id of the usage plan resource representing to associate the key to.
    """
    value: pulumi.Output[str]
    """
    The value of a usage plan key.
    """
    def __init__(__self__, resource_name, opts=None, key_id=None, key_type=None, usage_plan_id=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway Usage Plan Key.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key_id: The identifier of the API key resource.
        :param pulumi.Input[str] key_type: The type of the API key resource. Currently, the valid key type is API_KEY.
        :param pulumi.Input[str] usage_plan_id: The Id of the usage plan resource representing to associate the key to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/api_gateway_usage_plan_key.html.markdown.
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

        if key_id is None:
            raise TypeError("Missing required property 'key_id'")
        __props__['key_id'] = key_id

        if key_type is None:
            raise TypeError("Missing required property 'key_type'")
        __props__['key_type'] = key_type

        if usage_plan_id is None:
            raise TypeError("Missing required property 'usage_plan_id'")
        __props__['usage_plan_id'] = usage_plan_id

        __props__['name'] = None
        __props__['value'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(UsagePlanKey, __self__).__init__(
            'aws:apigateway/usagePlanKey:UsagePlanKey',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

