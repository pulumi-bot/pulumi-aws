# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ApiKey(pulumi.CustomResource):
    created_date: pulumi.Output[str]
    """
    The creation date of the API key
    """
    description: pulumi.Output[str]
    """
    The API key description. Defaults to "Managed by Terraform".
    """
    enabled: pulumi.Output[bool]
    """
    Specifies whether the API key can be used by callers. Defaults to `true`.
    """
    last_updated_date: pulumi.Output[str]
    """
    The last update date of the API key
    """
    name: pulumi.Output[str]
    """
    The name of the API key
    """
    value: pulumi.Output[str]
    """
    The value of the API key. If not specified, it will be automatically generated by AWS on creation.
    """
    def __init__(__self__, resource_name, opts=None, description=None, enabled=None, name=None, value=None, __name__=None, __opts__=None):
        """
        Provides an API Gateway API Key.
        
        > **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The API key description. Defaults to "Managed by Terraform".
        :param pulumi.Input[bool] enabled: Specifies whether the API key can be used by callers. Defaults to `true`.
        :param pulumi.Input[str] name: The name of the API key
        :param pulumi.Input[str] value: The value of the API key. If not specified, it will be automatically generated by AWS on creation.
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

        if description is None:
            description = 'Managed by Pulumi'
        __props__['description'] = description

        __props__['enabled'] = enabled

        __props__['name'] = name

        __props__['value'] = value

        __props__['created_date'] = None
        __props__['last_updated_date'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ApiKey, __self__).__init__(
            'aws:apigateway/apiKey:ApiKey',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

