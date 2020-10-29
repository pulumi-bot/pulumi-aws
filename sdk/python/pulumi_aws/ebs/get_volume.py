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

__all__ = [
    'GetVolumeResult',
    'AwaitableGetVolumeResult',
    'get_volume',
]

@pulumi.output_type
class GetVolumeResult:
    """
    A collection of values returned by getVolume.
    """
    def __init__(__self__, arn=None, availability_zone=None, encrypted=None, filters=None, id=None, iops=None, kms_key_id=None, most_recent=None, multi_attach_enabled=None, outpost_arn=None, size=None, snapshot_id=None, tags=None, volume_id=None, volume_type=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iops and not isinstance(iops, int):
            raise TypeError("Expected argument 'iops' to be a int")
        pulumi.set(__self__, "iops", iops)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        pulumi.set(__self__, "most_recent", most_recent)
        if multi_attach_enabled and not isinstance(multi_attach_enabled, bool):
            raise TypeError("Expected argument 'multi_attach_enabled' to be a bool")
        pulumi.set(__self__, "multi_attach_enabled", multi_attach_enabled)
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        pulumi.set(__self__, "outpost_arn", outpost_arn)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The AZ where the EBS volume exists.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        """
        Whether the disk is encrypted.
        """
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetVolumeFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def iops(self) -> int:
        """
        The amount of IOPS for the disk.
        """
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> str:
        """
        The ARN for the KMS encryption key.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="mostRecent")
    def most_recent(self) -> Optional[bool]:
        return pulumi.get(self, "most_recent")

    @property
    @pulumi.getter(name="multiAttachEnabled")
    def multi_attach_enabled(self) -> bool:
        """
        (Optional) Specifies whether Amazon EBS Multi-Attach is enabled.
        """
        return pulumi.get(self, "multi_attach_enabled")

    @property
    @pulumi.getter(name="outpostArn")
    def outpost_arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the Outpost.
        """
        return pulumi.get(self, "outpost_arn")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size of the drive in GiBs.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> str:
        """
        The snapshot_id the EBS volume is based off.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        A map of tags for the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> str:
        """
        The volume ID (e.g. vol-59fcb34e).
        """
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> str:
        """
        The type of EBS volume.
        """
        return pulumi.get(self, "volume_type")


class AwaitableGetVolumeResult(GetVolumeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVolumeResult(
            arn=self.arn,
            availability_zone=self.availability_zone,
            encrypted=self.encrypted,
            filters=self.filters,
            id=self.id,
            iops=self.iops,
            kms_key_id=self.kms_key_id,
            most_recent=self.most_recent,
            multi_attach_enabled=self.multi_attach_enabled,
            outpost_arn=self.outpost_arn,
            size=self.size,
            snapshot_id=self.snapshot_id,
            tags=self.tags,
            volume_id=self.volume_id,
            volume_type=self.volume_type)


def get_volume(filters: Optional[Sequence[pulumi.InputType['GetVolumeFilterArgs']]] = None,
               most_recent: Optional[bool] = None,
               tags: Optional[Mapping[str, str]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVolumeResult:
    """
    Use this data source to get information about an EBS volume for use in other
    resources.


    :param Sequence[pulumi.InputType['GetVolumeFilterArgs']] filters: One or more name/value pairs to filter off of. There are
           several valid keys, for a full reference, check out
           [describe-volumes in the AWS CLI reference][1].
    :param bool most_recent: If more than one result is returned, use the most
           recent Volume.
    :param Mapping[str, str] tags: A map of tags for the resource.
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['mostRecent'] = most_recent
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ebs/getVolume:getVolume', __args__, opts=opts, typ=GetVolumeResult).value

    return AwaitableGetVolumeResult(
        arn=__ret__.arn,
        availability_zone=__ret__.availability_zone,
        encrypted=__ret__.encrypted,
        filters=__ret__.filters,
        id=__ret__.id,
        iops=__ret__.iops,
        kms_key_id=__ret__.kms_key_id,
        most_recent=__ret__.most_recent,
        multi_attach_enabled=__ret__.multi_attach_enabled,
        outpost_arn=__ret__.outpost_arn,
        size=__ret__.size,
        snapshot_id=__ret__.snapshot_id,
        tags=__ret__.tags,
        volume_id=__ret__.volume_id,
        volume_type=__ret__.volume_type)
