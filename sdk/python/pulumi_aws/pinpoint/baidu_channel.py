# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
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
    def __init__(__self__, resource_name, opts=None, api_key=None, application_id=None, enabled=None, secret_key=None, __name__=None, __opts__=None):
        """
        Provides a Pinpoint Baidu Channel resource.
        
        > **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.
        [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: Platform credential API key from Baidu.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the channel. Defaults to `true`.
        :param pulumi.Input[str] secret_key: Platform credential Secret key from Baidu.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_baidu_channel.html.markdown.
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

        if api_key is None:
            raise TypeError("Missing required property 'api_key'")
        __props__['api_key'] = api_key
        if application_id is None:
            raise TypeError("Missing required property 'application_id'")
        __props__['application_id'] = application_id
        __props__['enabled'] = enabled
        if secret_key is None:
            raise TypeError("Missing required property 'secret_key'")
        __props__['secret_key'] = secret_key
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(BaiduChannel, __self__).__init__(
            'aws:pinpoint/baiduChannel:BaiduChannel',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

