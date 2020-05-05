# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetSnapshotResult:
    """
    A collection of values returned by getSnapshot.
    """
    def __init__(__self__, data_encryption_key_id=None, description=None, encrypted=None, filters=None, id=None, kms_key_id=None, most_recent=None, owner_alias=None, owner_id=None, owners=None, restorable_by_user_ids=None, snapshot_id=None, snapshot_ids=None, state=None, tags=None, volume_id=None, volume_size=None):
        if data_encryption_key_id and not isinstance(data_encryption_key_id, str):
            raise TypeError("Expected argument 'data_encryption_key_id' to be a str")
        __self__.data_encryption_key_id = data_encryption_key_id
        """
        The data encryption key identifier for the snapshot.
        """
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        __self__.description = description
        """
        A description for the snapshot
        """
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        __self__.encrypted = encrypted
        """
        Whether the snapshot is encrypted.
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
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        __self__.kms_key_id = kms_key_id
        """
        The ARN for the KMS encryption key.
        """
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        __self__.most_recent = most_recent
        if owner_alias and not isinstance(owner_alias, str):
            raise TypeError("Expected argument 'owner_alias' to be a str")
        __self__.owner_alias = owner_alias
        """
        Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
        """
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        __self__.owner_id = owner_id
        """
        The AWS account ID of the EBS snapshot owner.
        """
        if owners and not isinstance(owners, list):
            raise TypeError("Expected argument 'owners' to be a list")
        __self__.owners = owners
        if restorable_by_user_ids and not isinstance(restorable_by_user_ids, list):
            raise TypeError("Expected argument 'restorable_by_user_ids' to be a list")
        __self__.restorable_by_user_ids = restorable_by_user_ids
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        __self__.snapshot_id = snapshot_id
        """
        The snapshot ID (e.g. snap-59fcb34e).
        """
        if snapshot_ids and not isinstance(snapshot_ids, list):
            raise TypeError("Expected argument 'snapshot_ids' to be a list")
        __self__.snapshot_ids = snapshot_ids
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        The snapshot state.
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
        if volume_size and not isinstance(volume_size, float):
            raise TypeError("Expected argument 'volume_size' to be a float")
        __self__.volume_size = volume_size
        """
        The size of the drive in GiBs.
        """
class AwaitableGetSnapshotResult(GetSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotResult(
            data_encryption_key_id=self.data_encryption_key_id,
            description=self.description,
            encrypted=self.encrypted,
            filters=self.filters,
            id=self.id,
            kms_key_id=self.kms_key_id,
            most_recent=self.most_recent,
            owner_alias=self.owner_alias,
            owner_id=self.owner_id,
            owners=self.owners,
            restorable_by_user_ids=self.restorable_by_user_ids,
            snapshot_id=self.snapshot_id,
            snapshot_ids=self.snapshot_ids,
            state=self.state,
            tags=self.tags,
            volume_id=self.volume_id,
            volume_size=self.volume_size)

def get_snapshot(filters=None,most_recent=None,owners=None,restorable_by_user_ids=None,snapshot_ids=None,tags=None,opts=None):
    """
    Use this data source to get information about an EBS Snapshot for use when provisioning EBS Volumes

    ## Example Usage



    ```python
    import pulumi
    import pulumi_aws as aws

    ebs_volume = aws.ebs.get_snapshot(filters=[
            {
                "name": "volume-size",
                "values": ["40"],
            },
            {
                "name": "tag:Name",
                "values": ["Example"],
            },
        ],
        most_recent=True,
        owners=["self"])
    ```



    :param list filters: One or more name/value pairs to filter off of. There are
           several valid keys, for a full reference, check out
           [describe-snapshots in the AWS CLI reference][1].
    :param bool most_recent: If more than one result is returned, use the most recent snapshot.
    :param list owners: Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
    :param list restorable_by_user_ids: One or more AWS accounts IDs that can create volumes from the snapshot.
    :param list snapshot_ids: Returns information on a specific snapshot_id.
    :param dict tags: A map of tags for the resource.

    The **filters** object supports the following:

      * `name` (`str`)
      * `values` (`list`)
    """
    __args__ = dict()


    __args__['filters'] = filters
    __args__['mostRecent'] = most_recent
    __args__['owners'] = owners
    __args__['restorableByUserIds'] = restorable_by_user_ids
    __args__['snapshotIds'] = snapshot_ids
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ebs/getSnapshot:getSnapshot', __args__, opts=opts).value

    return AwaitableGetSnapshotResult(
        data_encryption_key_id=__ret__.get('dataEncryptionKeyId'),
        description=__ret__.get('description'),
        encrypted=__ret__.get('encrypted'),
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        kms_key_id=__ret__.get('kmsKeyId'),
        most_recent=__ret__.get('mostRecent'),
        owner_alias=__ret__.get('ownerAlias'),
        owner_id=__ret__.get('ownerId'),
        owners=__ret__.get('owners'),
        restorable_by_user_ids=__ret__.get('restorableByUserIds'),
        snapshot_id=__ret__.get('snapshotId'),
        snapshot_ids=__ret__.get('snapshotIds'),
        state=__ret__.get('state'),
        tags=__ret__.get('tags'),
        volume_id=__ret__.get('volumeId'),
        volume_size=__ret__.get('volumeSize'))
