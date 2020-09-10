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
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 db_instance_identifier: Optional[pulumi.Input[str]] = None,
                 db_snapshot_identifier: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Snapshot resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            if db_instance_identifier is None:
                raise TypeError("Missing required property 'db_instance_identifier'")
            __props__['db_instance_identifier'] = db_instance_identifier
            if db_snapshot_identifier is None:
                raise TypeError("Missing required property 'db_snapshot_identifier'")
            __props__['db_snapshot_identifier'] = db_snapshot_identifier
            __props__['tags'] = tags
            __props__['allocated_storage'] = None
            __props__['availability_zone'] = None
            __props__['db_snapshot_arn'] = None
            __props__['encrypted'] = None
            __props__['engine'] = None
            __props__['engine_version'] = None
            __props__['iops'] = None
            __props__['kms_key_id'] = None
            __props__['license_model'] = None
            __props__['option_group_name'] = None
            __props__['port'] = None
            __props__['snapshot_type'] = None
            __props__['source_db_snapshot_identifier'] = None
            __props__['source_region'] = None
            __props__['status'] = None
            __props__['storage_type'] = None
            __props__['vpc_id'] = None
        super(Snapshot, __self__).__init__(
            'aws:rds/snapshot:Snapshot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocated_storage: Optional[pulumi.Input[float]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            db_instance_identifier: Optional[pulumi.Input[str]] = None,
            db_snapshot_arn: Optional[pulumi.Input[str]] = None,
            db_snapshot_identifier: Optional[pulumi.Input[str]] = None,
            encrypted: Optional[pulumi.Input[bool]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            iops: Optional[pulumi.Input[float]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            license_model: Optional[pulumi.Input[str]] = None,
            option_group_name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[float]] = None,
            snapshot_type: Optional[pulumi.Input[str]] = None,
            source_db_snapshot_identifier: Optional[pulumi.Input[str]] = None,
            source_region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Snapshot':
        """
        Get an existing Snapshot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocated_storage"] = allocated_storage
        __props__["availability_zone"] = availability_zone
        __props__["db_instance_identifier"] = db_instance_identifier
        __props__["db_snapshot_arn"] = db_snapshot_arn
        __props__["db_snapshot_identifier"] = db_snapshot_identifier
        __props__["encrypted"] = encrypted
        __props__["engine"] = engine
        __props__["engine_version"] = engine_version
        __props__["iops"] = iops
        __props__["kms_key_id"] = kms_key_id
        __props__["license_model"] = license_model
        __props__["option_group_name"] = option_group_name
        __props__["port"] = port
        __props__["snapshot_type"] = snapshot_type
        __props__["source_db_snapshot_identifier"] = source_db_snapshot_identifier
        __props__["source_region"] = source_region
        __props__["status"] = status
        __props__["storage_type"] = storage_type
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return Snapshot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocatedStorage")
    def allocated_storage(self) -> pulumi.Output[float]:
        return pulumi.get(self, "allocated_storage")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="dbInstanceIdentifier")
    def db_instance_identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "db_instance_identifier")

    @property
    @pulumi.getter(name="dbSnapshotArn")
    def db_snapshot_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "db_snapshot_arn")

    @property
    @pulumi.getter(name="dbSnapshotIdentifier")
    def db_snapshot_identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "db_snapshot_identifier")

    @property
    @pulumi.getter
    def encrypted(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter
    def iops(self) -> pulumi.Output[float]:
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="licenseModel")
    def license_model(self) -> pulumi.Output[str]:
        return pulumi.get(self, "license_model")

    @property
    @pulumi.getter(name="optionGroupName")
    def option_group_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "option_group_name")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[float]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="snapshotType")
    def snapshot_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "snapshot_type")

    @property
    @pulumi.getter(name="sourceDbSnapshotIdentifier")
    def source_db_snapshot_identifier(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_db_snapshot_identifier")

    @property
    @pulumi.getter(name="sourceRegion")
    def source_region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "source_region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

