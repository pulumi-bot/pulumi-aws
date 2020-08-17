# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Snapshot']


class Snapshot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 volume_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a Snapshot of an EBS Volume.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ebs.Volume("example",
            availability_zone="us-west-2a",
            size=40,
            tags={
                "Name": "HelloWorld",
            })
        example_snapshot = aws.ebs.Snapshot("exampleSnapshot",
            volume_id=example.id,
            tags={
                "Name": "HelloWorld_snap",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of what the snapshot is.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the snapshot
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['tags'] = tags
            if volume_id is None:
                raise TypeError("Missing required property 'volume_id'")
            __props__['volume_id'] = volume_id
            __props__['arn'] = None
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
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            data_encryption_key_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            encrypted: Optional[pulumi.Input[bool]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            owner_alias: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            volume_id: Optional[pulumi.Input[str]] = None,
            volume_size: Optional[pulumi.Input[float]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the EBS Snapshot.
        :param pulumi.Input[str] data_encryption_key_id: The data encryption key identifier for the snapshot.
        :param pulumi.Input[str] description: A description of what the snapshot is.
        :param pulumi.Input[bool] encrypted: Whether the snapshot is encrypted.
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key.
        :param pulumi.Input[str] owner_alias: Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
        :param pulumi.Input[str] owner_id: The AWS account ID of the EBS snapshot owner.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the snapshot
        :param pulumi.Input[str] volume_id: The Volume ID of which to make a snapshot.
        :param pulumi.Input[float] volume_size: The size of the drive in GiBs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
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

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) of the EBS Snapshot.
        """
        ...

    @property
    @pulumi.getter(name="dataEncryptionKeyId")
    def data_encryption_key_id(self) -> str:
        """
        The data encryption key identifier for the snapshot.
        """
        ...

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        A description of what the snapshot is.
        """
        ...

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        """
        Whether the snapshot is encrypted.
        """
        ...

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> str:
        """
        The ARN for the KMS encryption key.
        """
        ...

    @property
    @pulumi.getter(name="ownerAlias")
    def owner_alias(self) -> str:
        """
        Value from an Amazon-maintained list (`amazon`, `aws-marketplace`, `microsoft`) of snapshot owners.
        """
        ...

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        The AWS account ID of the EBS snapshot owner.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the snapshot
        """
        ...

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> str:
        """
        The Volume ID of which to make a snapshot.
        """
        ...

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> float:
        """
        The size of the drive in GiBs.
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

