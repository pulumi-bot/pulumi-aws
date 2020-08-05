# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GetVolumeResult:
    """
    A collection of values returned by getVolume.
    """
    def __init__(__self__, arn=None, availability_zone=None, encrypted=None, filters=None, id=None, iops=None, kms_key_id=None, most_recent=None, multi_attach_enabled=None, outpost_arn=None, size=None, snapshot_id=None, tags=None, volume_id=None, volume_type=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The volume ARN (e.g. arn:aws:ec2:us-east-1:0123456789012:volume/vol-59fcb34e).
        """
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        __self__.availability_zone = availability_zone
        """
        The AZ where the EBS volume exists.
        """
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        __self__.encrypted = encrypted
        """
        Whether the disk is encrypted.
        """
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if iops and not isinstance(iops, float):
            raise TypeError("Expected argument 'iops' to be a float")
        __self__.iops = iops
        """
        The amount of IOPS for the disk.
        """
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        __self__.kms_key_id = kms_key_id
        """
        The ARN for the KMS encryption key.
        """
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        __self__.most_recent = most_recent
        if multi_attach_enabled and not isinstance(multi_attach_enabled, bool):
            raise TypeError("Expected argument 'multi_attach_enabled' to be a bool")
        __self__.multi_attach_enabled = multi_attach_enabled
        """
        (Optional) Specifies whether Amazon EBS Multi-Attach is enabled.
        """
        if outpost_arn and not isinstance(outpost_arn, str):
            raise TypeError("Expected argument 'outpost_arn' to be a str")
        __self__.outpost_arn = outpost_arn
        """
        The Amazon Resource Name (ARN) of the Outpost.
        """
        if size and not isinstance(size, float):
            raise TypeError("Expected argument 'size' to be a float")
        __self__.size = size
        """
        The size of the drive in GiBs.
        """
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        __self__.snapshot_id = snapshot_id
        """
        The snapshot_id the EBS volume is based off.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        A map of tags for the resource.
        """
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        __self__.volume_id = volume_id
        """
        The volume ID (e.g. vol-59fcb34e).
        """
        if volume_type and not isinstance(volume_type, str):
            raise TypeError("Expected argument 'volume_type' to be a str")
        __self__.volume_type = volume_type
        """
        The type of EBS volume.
        """


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


def get_volume(filters=None, most_recent=None, tags=None, opts=None):
    """
    Use this data source to get information about an EBS volume for use in other
    resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    ebs_volume = aws.ebs.get_volume(filters=[
            {
                "name": "volume-type",
                "values": ["gp2"],
            },
            {
                "name": "tag:Name",
                "values": ["Example"],
            },
        ],
        most_recent=True)
    ```


    :param list filters: One or more name/value pairs to filter off of. There are
           several valid keys, for a full reference, check out
           [describe-volumes in the AWS CLI reference][1].
    :param bool most_recent: If more than one result is returned, use the most
           recent Volume.
    :param dict tags: A map of tags for the resource.

    The **filters** object supports the following:

      * `name` (`str`)
      * `values` (`list`)
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['mostRecent'] = most_recent
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ebs/getVolume:getVolume', __args__, opts=opts).value

    return AwaitableGetVolumeResult(
        arn=__ret__.get('arn'),
        availability_zone=__ret__.get('availabilityZone'),
        encrypted=__ret__.get('encrypted'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        iops=__ret__.get('iops'),
        kms_key_id=__ret__.get('kmsKeyId'),
        most_recent=__ret__.get('mostRecent'),
        multi_attach_enabled=__ret__.get('multiAttachEnabled'),
        outpost_arn=__ret__.get('outpostArn'),
        size=__ret__.get('size'),
        snapshot_id=__ret__.get('snapshotId'),
        tags=__ret__.get('tags'),
        volume_id=__ret__.get('volumeId'),
        volume_type=__ret__.get('volumeType'))
