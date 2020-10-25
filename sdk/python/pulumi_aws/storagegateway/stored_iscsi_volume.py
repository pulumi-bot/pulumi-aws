# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['StoredIscsiVolume']


class StoredIscsiVolume(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 gateway_arn: Optional[pulumi.Input[str]] = None,
                 kms_encrypted: Optional[pulumi.Input[bool]] = None,
                 kms_key: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 preserve_existing_data: Optional[pulumi.Input[bool]] = None,
                 snapshot_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an AWS Storage Gateway stored iSCSI volume.

        > **NOTE:** The gateway must have a working storage added (e.g. via the [`storagegateway.WorkingStorage`](https://www.terraform.io/docs/providers/aws/r/storagegateway_working_storage.html) resource) before the volume is operational to clients, however the Storage Gateway API will allow volume creation without error in that case and return volume status as `WORKING STORAGE NOT CONFIGURED`.

        ## Example Usage
        ### Create Empty Stored iSCSI Volume

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.storagegateway.StoredIscsiVolume("example",
            gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
            network_interface_id=aws_instance["example"]["private_ip"],
            target_name="example",
            preserve_existing_data=False,
            disk_id=data["aws_storagegateway_local_disk"]["test"]["id"])
        ```
        ### Create Stored iSCSI Volume From Snapshot

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.storagegateway.StoredIscsiVolume("example",
            gateway_arn=aws_storagegateway_cache["example"]["gateway_arn"],
            network_interface_id=aws_instance["example"]["private_ip"],
            snapshot_id=aws_ebs_snapshot["example"]["id"],
            target_name="example",
            preserve_existing_data=False,
            disk_id=data["aws_storagegateway_local_disk"]["test"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: The unique identifier for the gateway local disk that is configured as a stored volume.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[bool] kms_encrypted: `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
        :param pulumi.Input[str] kms_key: The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is `true`.
        :param pulumi.Input[str] network_interface_id: The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
        :param pulumi.Input[bool] preserve_existing_data: Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] target_name: The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
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

            if disk_id is None:
                raise TypeError("Missing required property 'disk_id'")
            __props__['disk_id'] = disk_id
            if gateway_arn is None:
                raise TypeError("Missing required property 'gateway_arn'")
            __props__['gateway_arn'] = gateway_arn
            __props__['kms_encrypted'] = kms_encrypted
            __props__['kms_key'] = kms_key
            if network_interface_id is None:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__['network_interface_id'] = network_interface_id
            if preserve_existing_data is None:
                raise TypeError("Missing required property 'preserve_existing_data'")
            __props__['preserve_existing_data'] = preserve_existing_data
            __props__['snapshot_id'] = snapshot_id
            __props__['tags'] = tags
            if target_name is None:
                raise TypeError("Missing required property 'target_name'")
            __props__['target_name'] = target_name
            __props__['arn'] = None
            __props__['chap_enabled'] = None
            __props__['lun_number'] = None
            __props__['network_interface_port'] = None
            __props__['target_arn'] = None
            __props__['volume_attachment_status'] = None
            __props__['volume_id'] = None
            __props__['volume_size_in_bytes'] = None
            __props__['volume_status'] = None
            __props__['volume_type'] = None
        super(StoredIscsiVolume, __self__).__init__(
            'aws:storagegateway/storedIscsiVolume:StoredIscsiVolume',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            chap_enabled: Optional[pulumi.Input[bool]] = None,
            disk_id: Optional[pulumi.Input[str]] = None,
            gateway_arn: Optional[pulumi.Input[str]] = None,
            kms_encrypted: Optional[pulumi.Input[bool]] = None,
            kms_key: Optional[pulumi.Input[str]] = None,
            lun_number: Optional[pulumi.Input[int]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            network_interface_port: Optional[pulumi.Input[int]] = None,
            preserve_existing_data: Optional[pulumi.Input[bool]] = None,
            snapshot_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_arn: Optional[pulumi.Input[str]] = None,
            target_name: Optional[pulumi.Input[str]] = None,
            volume_attachment_status: Optional[pulumi.Input[str]] = None,
            volume_id: Optional[pulumi.Input[str]] = None,
            volume_size_in_bytes: Optional[pulumi.Input[int]] = None,
            volume_status: Optional[pulumi.Input[str]] = None,
            volume_type: Optional[pulumi.Input[str]] = None) -> 'StoredIscsiVolume':
        """
        Get an existing StoredIscsiVolume resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
        :param pulumi.Input[bool] chap_enabled: Whether mutual CHAP is enabled for the iSCSI target.
        :param pulumi.Input[str] disk_id: The unique identifier for the gateway local disk that is configured as a stored volume.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        :param pulumi.Input[bool] kms_encrypted: `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
        :param pulumi.Input[str] kms_key: The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is `true`.
        :param pulumi.Input[int] lun_number: Logical disk number.
        :param pulumi.Input[str] network_interface_id: The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
        :param pulumi.Input[int] network_interface_port: The port used to communicate with iSCSI targets.
        :param pulumi.Input[bool] preserve_existing_data: Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
        :param pulumi.Input[str] snapshot_id: The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value mapping of resource tags
        :param pulumi.Input[str] target_arn: Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
        :param pulumi.Input[str] target_name: The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
        :param pulumi.Input[str] volume_attachment_status: A value that indicates whether a storage volume is attached to, detached from, or is in the process of detaching from a gateway.
        :param pulumi.Input[str] volume_id: Volume ID, e.g. `vol-12345678`.
        :param pulumi.Input[int] volume_size_in_bytes: The size of the data stored on the volume in bytes.
        :param pulumi.Input[str] volume_status: indicates the state of the storage volume.
        :param pulumi.Input[str] volume_type: indicates the type of the volume.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["chap_enabled"] = chap_enabled
        __props__["disk_id"] = disk_id
        __props__["gateway_arn"] = gateway_arn
        __props__["kms_encrypted"] = kms_encrypted
        __props__["kms_key"] = kms_key
        __props__["lun_number"] = lun_number
        __props__["network_interface_id"] = network_interface_id
        __props__["network_interface_port"] = network_interface_port
        __props__["preserve_existing_data"] = preserve_existing_data
        __props__["snapshot_id"] = snapshot_id
        __props__["tags"] = tags
        __props__["target_arn"] = target_arn
        __props__["target_name"] = target_name
        __props__["volume_attachment_status"] = volume_attachment_status
        __props__["volume_id"] = volume_id
        __props__["volume_size_in_bytes"] = volume_size_in_bytes
        __props__["volume_status"] = volume_status
        __props__["volume_type"] = volume_type
        return StoredIscsiVolume(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Volume Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/volume/vol-12345678`.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="chapEnabled")
    def chap_enabled(self) -> pulumi.Output[bool]:
        """
        Whether mutual CHAP is enabled for the iSCSI target.
        """
        return pulumi.get(self, "chap_enabled")

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> pulumi.Output[str]:
        """
        The unique identifier for the gateway local disk that is configured as a stored volume.
        """
        return pulumi.get(self, "disk_id")

    @property
    @pulumi.getter(name="gatewayArn")
    def gateway_arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the gateway.
        """
        return pulumi.get(self, "gateway_arn")

    @property
    @pulumi.getter(name="kmsEncrypted")
    def kms_encrypted(self) -> pulumi.Output[Optional[bool]]:
        """
        `true` to use Amazon S3 server side encryption with your own AWS KMS key, or `false` to use a key managed by Amazon S3. Optional.
        """
        return pulumi.get(self, "kms_encrypted")

    @property
    @pulumi.getter(name="kmsKey")
    def kms_key(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server side encryption. This value can only be set when `kms_encrypted` is `true`.
        """
        return pulumi.get(self, "kms_key")

    @property
    @pulumi.getter(name="lunNumber")
    def lun_number(self) -> pulumi.Output[int]:
        """
        Logical disk number.
        """
        return pulumi.get(self, "lun_number")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        The network interface of the gateway on which to expose the iSCSI target. Only IPv4 addresses are accepted.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="networkInterfacePort")
    def network_interface_port(self) -> pulumi.Output[int]:
        """
        The port used to communicate with iSCSI targets.
        """
        return pulumi.get(self, "network_interface_port")

    @property
    @pulumi.getter(name="preserveExistingData")
    def preserve_existing_data(self) -> pulumi.Output[bool]:
        """
        Specify this field as `true` if you want to preserve the data on the local disk. Otherwise, specifying this field as false creates an empty volume.
        """
        return pulumi.get(self, "preserve_existing_data")

    @property
    @pulumi.getter(name="snapshotId")
    def snapshot_id(self) -> pulumi.Output[Optional[str]]:
        """
        The snapshot ID of the snapshot to restore as the new stored volume. e.g. `snap-1122aabb`.
        """
        return pulumi.get(self, "snapshot_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value mapping of resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> pulumi.Output[str]:
        """
        Target Amazon Resource Name (ARN), e.g. `arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678/target/iqn.1997-05.com.amazon:TargetName`.
        """
        return pulumi.get(self, "target_arn")

    @property
    @pulumi.getter(name="targetName")
    def target_name(self) -> pulumi.Output[str]:
        """
        The name of the iSCSI target used by initiators to connect to the target and as a suffix for the target ARN. The target name must be unique across all volumes of a gateway.
        """
        return pulumi.get(self, "target_name")

    @property
    @pulumi.getter(name="volumeAttachmentStatus")
    def volume_attachment_status(self) -> pulumi.Output[str]:
        """
        A value that indicates whether a storage volume is attached to, detached from, or is in the process of detaching from a gateway.
        """
        return pulumi.get(self, "volume_attachment_status")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> pulumi.Output[str]:
        """
        Volume ID, e.g. `vol-12345678`.
        """
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter(name="volumeSizeInBytes")
    def volume_size_in_bytes(self) -> pulumi.Output[int]:
        """
        The size of the data stored on the volume in bytes.
        """
        return pulumi.get(self, "volume_size_in_bytes")

    @property
    @pulumi.getter(name="volumeStatus")
    def volume_status(self) -> pulumi.Output[str]:
        """
        indicates the state of the storage volume.
        """
        return pulumi.get(self, "volume_status")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> pulumi.Output[str]:
        """
        indicates the type of the volume.
        """
        return pulumi.get(self, "volume_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

