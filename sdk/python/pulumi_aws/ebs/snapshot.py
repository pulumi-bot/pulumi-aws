# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Snapshot(pulumi.CustomResource):
    data_encryption_key_id: pulumi.Output[str]
    """
    The data encryption key identifier for the snapshot.
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
    """
    owner_alias: pulumi.Output[str]
    """
    Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
    """
    owner_id: pulumi.Output[str]
    """
    The AWS account ID of the EBS snapshot owner.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the snapshot
    """
    volume_id: pulumi.Output[str]
    """
    The Volume ID of which to make a snapshot.
    """
    volume_size: pulumi.Output[float]
    """
    The size of the drive in GiBs.
    """
    def __init__(__self__, resource_name, opts=None, description=None, tags=None, volume_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Snapshot of an EBS Volume.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ebs_snapshot.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of what the snapshot is.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the snapshot
        :param pulumi.Input[str] volume_id: The Volume ID of which to make a snapshot.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['tags'] = tags
            if volume_id is None:
                raise TypeError("Missing required property 'volume_id'")
            __props__['volume_id'] = volume_id
            __props__['data_encryption_key_id'] = None
            __props__['encrypted'] = None
            __props__['kms_key_id'] = None
            __props__['owner_alias'] = None
            __props__['owner_id'] = None
            __props__['volume_size'] = None
        super(Snapshot, __self__).__init__(
            'aws:ebs/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, data_encryption_key_id=None, description=None, encrypted=None, kms_key_id=None, owner_alias=None, owner_id=None, tags=None, volume_id=None, volume_size=None):
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_encryption_key_id: The data encryption key identifier for the snapshot.
        :param pulumi.Input[str] description: A description of what the snapshot is.
        :param pulumi.Input[bool] encrypted: Whether the snapshot is encrypted.
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key.
        :param pulumi.Input[str] owner_alias: Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
        :param pulumi.Input[str] owner_id: The AWS account ID of the EBS snapshot owner.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the snapshot
        :param pulumi.Input[str] volume_id: The Volume ID of which to make a snapshot.
        :param pulumi.Input[float] volume_size: The size of the drive in GiBs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_encryption_key_id"] = data_encryption_key_id
        __props__["description"] = description
        __props__["encrypted"] = encrypted
        __props__["kms_key_id"] = kms_key_id
        __props__["owner_alias"] = owner_alias
        __props__["owner_id"] = owner_id
        __props__["tags"] = tags
        __props__["volume_id"] = volume_id
        __props__["volume_size"] = volume_size
        return Snapshot(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

