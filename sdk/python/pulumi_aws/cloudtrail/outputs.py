# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'TrailEventSelector',
    'TrailEventSelectorDataResource',
    'TrailInsightSelector',
]

@pulumi.output_type
class TrailEventSelector(dict):
    def __init__(__self__, *,
                 data_resources: Optional[Sequence['outputs.TrailEventSelectorDataResource']] = None,
                 include_management_events: Optional[bool] = None,
                 read_write_type: Optional[str] = None):
        """
        :param Sequence['TrailEventSelectorDataResourceArgs'] data_resources: Specifies logging data events. Fields documented below.
        :param bool include_management_events: Specify if you want your event selector to include management events for your trail.
        :param str read_write_type: Specify if you want your trail to log read-only events, write-only events, or all. By default, the value is All. You can specify only the following value: "ReadOnly", "WriteOnly", "All". Defaults to `All`.
        """
        if data_resources is not None:
            pulumi.set(__self__, "data_resources", data_resources)
        if include_management_events is not None:
            pulumi.set(__self__, "include_management_events", include_management_events)
        if read_write_type is not None:
            pulumi.set(__self__, "read_write_type", read_write_type)

    @property
    @pulumi.getter(name="dataResources")
    def data_resources(self) -> Optional[Sequence['outputs.TrailEventSelectorDataResource']]:
        """
        Specifies logging data events. Fields documented below.
        """
        return pulumi.get(self, "data_resources")

    @property
    @pulumi.getter(name="includeManagementEvents")
    def include_management_events(self) -> Optional[bool]:
        """
        Specify if you want your event selector to include management events for your trail.
        """
        return pulumi.get(self, "include_management_events")

    @property
    @pulumi.getter(name="readWriteType")
    def read_write_type(self) -> Optional[str]:
        """
        Specify if you want your trail to log read-only events, write-only events, or all. By default, the value is All. You can specify only the following value: "ReadOnly", "WriteOnly", "All". Defaults to `All`.
        """
        return pulumi.get(self, "read_write_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TrailEventSelectorDataResource(dict):
    def __init__(__self__, *,
                 type: str,
                 values: Sequence[str]):
        """
        :param str type: The resource type in which you want to log data events. You can specify only the following value: "AWS::S3::Object", "AWS::Lambda::Function"
        :param Sequence[str] values: A list of ARN for the specified S3 buckets and object prefixes..
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type in which you want to log data events. You can specify only the following value: "AWS::S3::Object", "AWS::Lambda::Function"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        A list of ARN for the specified S3 buckets and object prefixes..
        """
        return pulumi.get(self, "values")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TrailInsightSelector(dict):
    def __init__(__self__, *,
                 insight_type: str):
        """
        :param str insight_type: The type of insights to log on a trail. In this release, only `ApiCallRateInsight` is supported as an insight type.
        """
        pulumi.set(__self__, "insight_type", insight_type)

    @property
    @pulumi.getter(name="insightType")
    def insight_type(self) -> str:
        """
        The type of insights to log on a trail. In this release, only `ApiCallRateInsight` is supported as an insight type.
        """
        return pulumi.get(self, "insight_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


