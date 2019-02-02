# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SnapshotCopy(pulumi.CustomResource):
    data_encryption_key_id: pulumi.Output[str]
    """
    The data encryption key identifier for the snapshot.
    * `source_snapshot_id` The ARN of the copied snapshot.
    * `source_region` The region of the source snapshot.
    """
    description: pulumi.Output[str]
    """
    A description of what the snapshot is.
    """
    encrypted: pulumi.Output[bool]
    """
    Whether the snapshot is encrypted.
    """
    kms_key_id: pulumi.Output[str]
    """
    The ARN for the KMS encryption key.
    * `source_snapshot_id` The ARN for the snapshot to be copied.
    * `source_region` The region of the source snapshot.
    """
    owner_alias: pulumi.Output[str]
    """
    Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
    """
    owner_id: pulumi.Output[str]
    """
    The AWS account ID of the snapshot owner.
    """
    source_region: pulumi.Output[str]
    source_snapshot_id: pulumi.Output[str]
    tags: pulumi.Output[dict]
    """
    A mapping of tags for the snapshot.
    """
    volume_id: pulumi.Output[str]
    volume_size: pulumi.Output[int]
    """
    The size of the drive in GiBs.
    """
    def __init__(__self__, __name__, __opts__=None, description=None, encrypted=None, kms_key_id=None, source_region=None, source_snapshot_id=None, tags=None):
        """
        Creates a Snapshot of a snapshot.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] description: A description of what the snapshot is.
        :param pulumi.Input[bool] encrypted: Whether the snapshot is encrypted.
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key.
               * `source_snapshot_id` The ARN for the snapshot to be copied.
               * `source_region` The region of the source snapshot.
        :param pulumi.Input[str] source_region
        :param pulumi.Input[str] source_snapshot_id
        :param pulumi.Input[dict] tags: A mapping of tags for the snapshot.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['description'] = description

        __props__['encrypted'] = encrypted

        __props__['kms_key_id'] = kms_key_id

        if not source_region:
            raise TypeError('Missing required property source_region')
        __props__['source_region'] = source_region

        if not source_snapshot_id:
            raise TypeError('Missing required property source_snapshot_id')
        __props__['source_snapshot_id'] = source_snapshot_id

        __props__['tags'] = tags

        __props__['data_encryption_key_id'] = None
        __props__['owner_alias'] = None
        __props__['owner_id'] = None
        __props__['volume_id'] = None
        __props__['volume_size'] = None

        super(SnapshotCopy, __self__).__init__(
            'aws:ebs/snapshotCopy:SnapshotCopy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

