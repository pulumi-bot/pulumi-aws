# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['FunctionEventInvokeConfig']


class FunctionEventInvokeConfig(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_config: Optional[pulumi.Input[pulumi.InputType['FunctionEventInvokeConfigDestinationConfigArgs']]] = None,
                 function_name: Optional[pulumi.Input[str]] = None,
                 maximum_event_age_in_seconds: Optional[pulumi.Input[int]] = None,
                 maximum_retry_attempts: Optional[pulumi.Input[int]] = None,
                 qualifier: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an asynchronous invocation configuration for a Lambda Function or Alias. More information about asynchronous invocations and the configurable values can be found in the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html).

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['FunctionEventInvokeConfigDestinationConfigArgs']] destination_config: Configuration block with destination configuration. See below for details.
        :param pulumi.Input[str] function_name: Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
        :param pulumi.Input[int] maximum_event_age_in_seconds: Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
        :param pulumi.Input[int] maximum_retry_attempts: Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
        :param pulumi.Input[str] qualifier: Lambda Function published version, `$LATEST`, or Lambda Alias name.
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

            __props__['destination_config'] = destination_config
            if function_name is None:
                raise TypeError("Missing required property 'function_name'")
            __props__['function_name'] = function_name
            __props__['maximum_event_age_in_seconds'] = maximum_event_age_in_seconds
            __props__['maximum_retry_attempts'] = maximum_retry_attempts
            __props__['qualifier'] = qualifier
        super(FunctionEventInvokeConfig, __self__).__init__(
            'aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_config: Optional[pulumi.Input[pulumi.InputType['FunctionEventInvokeConfigDestinationConfigArgs']]] = None,
            function_name: Optional[pulumi.Input[str]] = None,
            maximum_event_age_in_seconds: Optional[pulumi.Input[int]] = None,
            maximum_retry_attempts: Optional[pulumi.Input[int]] = None,
            qualifier: Optional[pulumi.Input[str]] = None) -> 'FunctionEventInvokeConfig':
        """
        Get an existing FunctionEventInvokeConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['FunctionEventInvokeConfigDestinationConfigArgs']] destination_config: Configuration block with destination configuration. See below for details.
        :param pulumi.Input[str] function_name: Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
        :param pulumi.Input[int] maximum_event_age_in_seconds: Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
        :param pulumi.Input[int] maximum_retry_attempts: Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
        :param pulumi.Input[str] qualifier: Lambda Function published version, `$LATEST`, or Lambda Alias name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["destination_config"] = destination_config
        __props__["function_name"] = function_name
        __props__["maximum_event_age_in_seconds"] = maximum_event_age_in_seconds
        __props__["maximum_retry_attempts"] = maximum_retry_attempts
        __props__["qualifier"] = qualifier
        return FunctionEventInvokeConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> pulumi.Output[Optional['outputs.FunctionEventInvokeConfigDestinationConfig']]:
        """
        Configuration block with destination configuration. See below for details.
        """
        return pulumi.get(self, "destination_config")

    @property
    @pulumi.getter(name="functionName")
    def function_name(self) -> pulumi.Output[str]:
        """
        Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
        """
        return pulumi.get(self, "function_name")

    @property
    @pulumi.getter(name="maximumEventAgeInSeconds")
    def maximum_event_age_in_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
        """
        return pulumi.get(self, "maximum_event_age_in_seconds")

    @property
    @pulumi.getter(name="maximumRetryAttempts")
    def maximum_retry_attempts(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
        """
        return pulumi.get(self, "maximum_retry_attempts")

    @property
    @pulumi.getter
    def qualifier(self) -> pulumi.Output[Optional[str]]:
        """
        Lambda Function published version, `$LATEST`, or Lambda Alias name.
        """
        return pulumi.get(self, "qualifier")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

