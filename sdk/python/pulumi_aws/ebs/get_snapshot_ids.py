# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetSnapshotIdsResult(object):
    """
    A collection of values returned by getSnapshotIds.
    """
    def __init__(__self__, ids=None, id=None):
        if ids and not isinstance(ids, list):
            raise TypeError('Expected argument ids to be a list')
        __self__.ids = ids
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_snapshot_ids(filters=None, owners=None, restorable_by_user_ids=None):
    """
    Use this data source to get a list of EBS Snapshot IDs matching the specified
    criteria.
    """
    __args__ = dict()

    __args__['filters'] = filters
    __args__['owners'] = owners
    __args__['restorableByUserIds'] = restorable_by_user_ids
    __ret__ = await pulumi.runtime.invoke('aws:ebs/getSnapshotIds:getSnapshotIds', __args__)

    return GetSnapshotIdsResult(
        ids=__ret__.get('ids'),
        id=__ret__.get('id'))
