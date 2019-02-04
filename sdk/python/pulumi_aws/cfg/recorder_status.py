# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RecorderStatus(pulumi.CustomResource):
    is_enabled: pulumi.Output[bool]
    """
    Whether the configuration recorder should be enabled or disabled.
    """
    name: pulumi.Output[str]
    """
    The name of the recorder
    """
    def __init__(__self__, __name__, __opts__=None, is_enabled=None, name=None):
        """
        Manages status (recording / stopped) of an AWS Config Configuration Recorder.
        
        > **Note:** Starting Configuration Recorder requires a [Delivery Channel](https://www.terraform.io/docs/providers/aws/r/config_delivery_channel.html) to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] is_enabled: Whether the configuration recorder should be enabled or disabled.
        :param pulumi.Input[str] name: The name of the recorder
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if is_enabled is None:
            raise TypeError('Missing required property is_enabled')
        __props__['is_enabled'] = is_enabled

        __props__['name'] = name

        super(RecorderStatus, __self__).__init__(
            'aws:cfg/recorderStatus:RecorderStatus',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

