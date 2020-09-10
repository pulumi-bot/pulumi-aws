# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['SmsChannel']


class SmsChannel(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 sender_id: Optional[pulumi.Input[str]] = None,
                 short_code: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a SmsChannel resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            if application_id is None:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['enabled'] = enabled
            __props__['sender_id'] = sender_id
            __props__['short_code'] = short_code
            __props__['promotional_messages_per_second'] = None
            __props__['transactional_messages_per_second'] = None
        super(SmsChannel, __self__).__init__(
            'aws:pinpoint/smsChannel:SmsChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            promotional_messages_per_second: Optional[pulumi.Input[float]] = None,
            sender_id: Optional[pulumi.Input[str]] = None,
            short_code: Optional[pulumi.Input[str]] = None,
            transactional_messages_per_second: Optional[pulumi.Input[float]] = None) -> 'SmsChannel':
        """
        Get an existing SmsChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["application_id"] = application_id
        __props__["enabled"] = enabled
        __props__["promotional_messages_per_second"] = promotional_messages_per_second
        __props__["sender_id"] = sender_id
        __props__["short_code"] = short_code
        __props__["transactional_messages_per_second"] = transactional_messages_per_second
        return SmsChannel(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationId")
    def application_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "application_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="promotionalMessagesPerSecond")
    def promotional_messages_per_second(self) -> pulumi.Output[float]:
        return pulumi.get(self, "promotional_messages_per_second")

    @property
    @pulumi.getter(name="senderId")
    def sender_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "sender_id")

    @property
    @pulumi.getter(name="shortCode")
    def short_code(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "short_code")

    @property
    @pulumi.getter(name="transactionalMessagesPerSecond")
    def transactional_messages_per_second(self) -> pulumi.Output[float]:
        return pulumi.get(self, "transactional_messages_per_second")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

