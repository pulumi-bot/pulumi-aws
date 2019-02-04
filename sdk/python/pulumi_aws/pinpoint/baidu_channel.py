# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BaiduChannel(pulumi.CustomResource):
    api_key: pulumi.Output[str]
    """
    Platform credential API key from Baidu.
    """
    application_id: pulumi.Output[str]
    """
    The application ID.
    """
    enabled: pulumi.Output[bool]
    """
    Specifies whether to enable the channel. Defaults to `true`.
    """
    secret_key: pulumi.Output[str]
    """
    Platform credential Secret key from Baidu.
    """
    def __init__(__self__, __name__, __opts__=None, api_key=None, application_id=None, enabled=None, secret_key=None):
        """
        Provides a Pinpoint Baidu Channel resource.
        
        > **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] api_key: Platform credential API key from Baidu.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the channel. Defaults to `true`.
        :param pulumi.Input[str] secret_key: Platform credential Secret key from Baidu.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if api_key is None:
            raise TypeError('Missing required property api_key')
        __props__['api_key'] = api_key

        if application_id is None:
            raise TypeError('Missing required property application_id')
        __props__['application_id'] = application_id

        __props__['enabled'] = enabled

        if secret_key is None:
            raise TypeError('Missing required property secret_key')
        __props__['secret_key'] = secret_key

        super(BaiduChannel, __self__).__init__(
            'aws:pinpoint/baiduChannel:BaiduChannel',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

