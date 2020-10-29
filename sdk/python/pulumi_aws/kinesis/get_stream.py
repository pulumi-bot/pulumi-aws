# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetStreamResult',
    'AwaitableGetStreamResult',
    'get_stream',
]

@pulumi.output_type
class GetStreamResult:
    """
    A collection of values returned by getStream.
    """
    def __init__(__self__, arn=None, closed_shards=None, creation_timestamp=None, id=None, name=None, open_shards=None, retention_period=None, shard_level_metrics=None, status=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if closed_shards and not isinstance(closed_shards, list):
            raise TypeError("Expected argument 'closed_shards' to be a list")
        pulumi.set(__self__, "closed_shards", closed_shards)
        if creation_timestamp and not isinstance(creation_timestamp, int):
            raise TypeError("Expected argument 'creation_timestamp' to be a int")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if open_shards and not isinstance(open_shards, list):
            raise TypeError("Expected argument 'open_shards' to be a list")
        pulumi.set(__self__, "open_shards", open_shards)
        if retention_period and not isinstance(retention_period, int):
            raise TypeError("Expected argument 'retention_period' to be a int")
        pulumi.set(__self__, "retention_period", retention_period)
        if shard_level_metrics and not isinstance(shard_level_metrics, list):
            raise TypeError("Expected argument 'shard_level_metrics' to be a list")
        pulumi.set(__self__, "shard_level_metrics", shard_level_metrics)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="closedShards")
    def closed_shards(self) -> Sequence[str]:
        """
        The list of shard ids in the CLOSED state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
        """
        return pulumi.get(self, "closed_shards")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> int:
        """
        The approximate UNIX timestamp that the stream was created.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Kinesis Stream.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openShards")
    def open_shards(self) -> Sequence[str]:
        """
        The list of shard ids in the OPEN state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
        """
        return pulumi.get(self, "open_shards")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> int:
        """
        Length of time (in hours) data records are accessible after they are added to the stream.
        """
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter(name="shardLevelMetrics")
    def shard_level_metrics(self) -> Sequence[str]:
        """
        A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more.
        """
        return pulumi.get(self, "shard_level_metrics")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags to assigned to the stream.
        """
        return pulumi.get(self, "tags")


class AwaitableGetStreamResult(GetStreamResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStreamResult(
            arn=self.arn,
            closed_shards=self.closed_shards,
            creation_timestamp=self.creation_timestamp,
            id=self.id,
            name=self.name,
            open_shards=self.open_shards,
            retention_period=self.retention_period,
            shard_level_metrics=self.shard_level_metrics,
            status=self.status,
            tags=self.tags)


def get_stream(name: Optional[str] = None,
               tags: Optional[Mapping[str, str]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStreamResult:
    """
    Use this data source to get information about a Kinesis Stream for use in other
    resources.

    For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).


    :param str name: The name of the Kinesis Stream.
    :param Mapping[str, str] tags: A map of tags to assigned to the stream.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:kinesis/getStream:getStream', __args__, opts=opts, typ=GetStreamResult).value

    return AwaitableGetStreamResult(
        arn=__ret__.arn,
        closed_shards=__ret__.closed_shards,
        creation_timestamp=__ret__.creation_timestamp,
        id=__ret__.id,
        name=__ret__.name,
        open_shards=__ret__.open_shards,
        retention_period=__ret__.retention_period,
        shard_level_metrics=__ret__.shard_level_metrics,
        status=__ret__.status,
        tags=__ret__.tags)
