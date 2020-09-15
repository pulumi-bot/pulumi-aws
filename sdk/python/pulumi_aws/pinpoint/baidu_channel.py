# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['BaiduChannel']


class BaiduChannel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_key: Optional[pulumi.Input[str]] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Pinpoint Baidu Channel resource.

        > **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        app = aws.pinpoint.App("app")
        channel = aws.pinpoint.BaiduChannel("channel",
            api_key="",
            application_id=app.application_id,
            secret_key="")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: Platform credential API key from Baidu.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the channel. Defaults to `true`.
        :param pulumi.Input[str] secret_key: Platform credential Secret key from Baidu.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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
        super(BaiduChannel, __self__).__init__(
            'aws:pinpoint/baiduChannel:BaiduChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_key: Optional[pulumi.Input[str]] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            secret_key: Optional[pulumi.Input[str]] = None) -> 'BaiduChannel':
        """
        Get an existing BaiduChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_key: Platform credential API key from Baidu.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Specifies whether to enable the channel. Defaults to `true`.
        :param pulumi.Input[str] secret_key: Platform credential Secret key from Baidu.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["api_key"] = api_key
        __props__["application_id"] = application_id
        __props__["enabled"] = enabled
        __props__["secret_key"] = secret_key
        return BaiduChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> pulumi.Output[str]:
        """
        Platform credential API key from Baidu.
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        """
        The application ID.
        """
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether to enable the channel. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        Platform credential Secret key from Baidu.
        """
        return pulumi.get(self, "secret_key")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

