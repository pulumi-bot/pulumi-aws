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
    'GetSnapshotResult',
    'AwaitableGetSnapshotResult',
    'get_snapshot',
]

@pulumi.output_type
class GetSnapshotResult:
    """
    A collection of values returned by getSnapshot.
    """
    def __init__(__self__, arn=None, data_encryption_key_id=None, description=None, encrypted=None, filters=None, id=None, kms_key_id=None, most_recent=None, owner_alias=None, owner_id=None, owners=None, restorable_by_user_ids=None, snapshot_id=None, snapshot_ids=None, state=None, tags=None, volume_id=None, volume_size=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if data_encryption_key_id and not isinstance(data_encryption_key_id, str):
            raise TypeError("Expected argument 'data_encryption_key_id' to be a str")
        pulumi.set(__self__, "data_encryption_key_id", data_encryption_key_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if most_recent and not isinstance(most_recent, bool):
            raise TypeError("Expected argument 'most_recent' to be a bool")
        pulumi.set(__self__, "most_recent", most_recent)
        if owner_alias and not isinstance(owner_alias, str):
            raise TypeError("Expected argument 'owner_alias' to be a str")
        pulumi.set(__self__, "owner_alias", owner_alias)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if owners and not isinstance(owners, list):
            raise TypeError("Expected argument 'owners' to be a list")
        pulumi.set(__self__, "owners", owners)
        if restorable_by_user_ids and not isinstance(restorable_by_user_ids, list):
            raise TypeError("Expected argument 'restorable_by_user_ids' to be a list")
        pulumi.set(__self__, "restorable_by_user_ids", restorable_by_user_ids)
        if snapshot_id and not isinstance(snapshot_id, str):
            raise TypeError("Expected argument 'snapshot_id' to be a str")
        pulumi.set(__self__, "snapshot_id", snapshot_id)
        if snapshot_ids and not isinstance(snapshot_ids, list):
            raise TypeError("Expected argument 'snapshot_ids' to be a list")
        pulumi.set(__self__, "snapshot_ids", snapshot_ids)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if volume_id and not isinstance(volume_id, str):
            raise TypeError("Expected argument 'volume_id' to be a str")
        pulumi.set(__self__, "volume_id", volume_id)
        if volume_size and not isinstance(volume_size, int):
            raise TypeError("Expected argument 'volume_size' to be a int")
        pulumi.set(__self__, "volume_size", volume_size)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) of the EBS Snapshot.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="dataEncryptionKeyId")
    def data_encryption_key_id(self) -> str:
        """
        The data encryption key identifier for the snapshot.
        """
        return pulumi.get(self, "data_encryption_key_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description for the snapshot
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        """
        Whether the snapshot is encrypted.
        """
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetSnapshotFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

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
    @pulumi.getter(name="ownerAlias")
    def owner_alias(self) -> str:
        """
        Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
        """
        return pulumi.get(self, "owner_alias")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        The AWS account ID of the EBS snapshot owner.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def owners(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "owners")

    @property
    @pulumi.getter(name="restorableByUserIds")
    def restorable_by_user_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "restorable_by_user_ids")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> str:
        """
        The snapshot ID (e.g. snap-59fcb34e).
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter(name="snapshotIds")
    def snapshot_ids(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "snapshot_ids")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The snapshot state.
        """
        return pulumi.get(self, "state")

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
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> int:
        """
        The size of the drive in GiBs.
        """
        return pulumi.get(self, "volume_size")


class AwaitableGetSnapshotResult(GetSnapshotResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSnapshotResult(
            arn=self.arn,
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


def get_snapshot(filters: Optional[Sequence[pulumi.InputType['GetSnapshotFilterArgs']]] = None,
                 most_recent: Optional[bool] = None,
                 owners: Optional[Sequence[str]] = None,
                 restorable_by_user_ids: Optional[Sequence[str]] = None,
                 snapshot_ids: Optional[Sequence[str]] = None,
                 tags: Optional[Mapping[str, str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSnapshotResult:
    """
    Use this data source to get information about an EBS Snapshot for use when provisioning EBS Volumes


    :param Sequence[pulumi.InputType['GetSnapshotFilterArgs']] filters: One or more name/value pairs to filter off of. There are
           several valid keys, for a full reference, check out
           [describe-snapshots in the AWS CLI reference][1].
    :param bool most_recent: If more than one result is returned, use the most recent snapshot.
    :param Sequence[str] owners: Returns the snapshots owned by the specified owner id. Multiple owners can be specified.
    :param Sequence[str] restorable_by_user_ids: One or more AWS accounts IDs that can create volumes from the snapshot.
    :param Sequence[str] snapshot_ids: Returns information on a specific snapshot_id.
    :param Mapping[str, str] tags: A map of tags for the resource.
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
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ebs/getSnapshot:getSnapshot', __args__, opts=opts, typ=GetSnapshotResult).value

    return AwaitableGetSnapshotResult(
        arn=__ret__.arn,
        data_encryption_key_id=__ret__.data_encryption_key_id,
        description=__ret__.description,
        encrypted=__ret__.encrypted,
        filters=__ret__.filters,
        id=__ret__.id,
        kms_key_id=__ret__.kms_key_id,
        most_recent=__ret__.most_recent,
        owner_alias=__ret__.owner_alias,
        owner_id=__ret__.owner_id,
        owners=__ret__.owners,
        restorable_by_user_ids=__ret__.restorable_by_user_ids,
        snapshot_id=__ret__.snapshot_id,
        snapshot_ids=__ret__.snapshot_ids,
        state=__ret__.state,
        tags=__ret__.tags,
        volume_id=__ret__.volume_id,
        volume_size=__ret__.volume_size)
