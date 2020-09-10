# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AppCampaignHookArgs',
    'AppLimitsArgs',
    'AppQuietTimeArgs',
]

@pulumi.input_type
class AppCampaignHookArgs:
    def __init__(__self__, *,
                 lambda_function_name: Optional[pulumi.Input[str]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 web_url: Optional[pulumi.Input[str]] = None):
        if lambda_function_name is not None:
            pulumi.set(__self__, "lambda_function_name", lambda_function_name)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if web_url is not None:
            pulumi.set(__self__, "web_url", web_url)

    @property
    @pulumi.getter(name="lambdaFunctionName")
    def lambda_function_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "lambda_function_name")

    @lambda_function_name.setter
    def lambda_function_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lambda_function_name", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter(name="webUrl")
    def web_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "web_url")

    @web_url.setter
    def web_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "web_url", value)


@pulumi.input_type
class AppLimitsArgs:
    def __init__(__self__, *,
                 daily: Optional[pulumi.Input[float]] = None,
                 maximum_duration: Optional[pulumi.Input[float]] = None,
                 messages_per_second: Optional[pulumi.Input[float]] = None,
                 total: Optional[pulumi.Input[float]] = None):
        if daily is not None:
            pulumi.set(__self__, "daily", daily)
        if maximum_duration is not None:
            pulumi.set(__self__, "maximum_duration", maximum_duration)
        if messages_per_second is not None:
            pulumi.set(__self__, "messages_per_second", messages_per_second)
        if total is not None:
            pulumi.set(__self__, "total", total)

    @property
    @pulumi.getter
    def daily(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "daily")

    @daily.setter
    def daily(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "daily", value)

    @property
    @pulumi.getter(name="maximumDuration")
    def maximum_duration(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "maximum_duration")

    @maximum_duration.setter
    def maximum_duration(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "maximum_duration", value)

    @property
    @pulumi.getter(name="messagesPerSecond")
    def messages_per_second(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "messages_per_second")

    @messages_per_second.setter
    def messages_per_second(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "messages_per_second", value)

    @property
    @pulumi.getter
    def total(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "total")

    @total.setter
    def total(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "total", value)


@pulumi.input_type
class AppQuietTimeArgs:
    def __init__(__self__, *,
                 end: Optional[pulumi.Input[str]] = None,
                 start: Optional[pulumi.Input[str]] = None):
        if end is not None:
            pulumi.set(__self__, "end", end)
        if start is not None:
            pulumi.set(__self__, "start", start)

    @property
    @pulumi.getter
    def end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "end")

    @end.setter
    def end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end", value)

    @property
    @pulumi.getter
    def start(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "start")

    @start.setter
    def start(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start", value)


