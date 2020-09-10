# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['WindowsFileSystem']


class WindowsFileSystem(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active_directory_id: Optional[pulumi.Input[str]] = None,
                 automatic_backup_retention_days: Optional[pulumi.Input[float]] = None,
                 copy_tags_to_backups: Optional[pulumi.Input[bool]] = None,
                 daily_automatic_backup_start_time: Optional[pulumi.Input[str]] = None,
                 deployment_type: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 preferred_subnet_id: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 self_managed_active_directory: Optional[pulumi.Input[pulumi.InputType['WindowsFileSystemSelfManagedActiveDirectoryArgs']]] = None,
                 skip_final_backup: Optional[pulumi.Input[bool]] = None,
                 storage_capacity: Optional[pulumi.Input[float]] = None,
                 storage_type: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 throughput_capacity: Optional[pulumi.Input[float]] = None,
                 weekly_maintenance_start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a WindowsFileSystem resource with the given unique name, props, and options.
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

            __props__['active_directory_id'] = active_directory_id
            __props__['automatic_backup_retention_days'] = automatic_backup_retention_days
            __props__['copy_tags_to_backups'] = copy_tags_to_backups
            __props__['daily_automatic_backup_start_time'] = daily_automatic_backup_start_time
            __props__['deployment_type'] = deployment_type
            __props__['kms_key_id'] = kms_key_id
            __props__['preferred_subnet_id'] = preferred_subnet_id
            __props__['security_group_ids'] = security_group_ids
            __props__['self_managed_active_directory'] = self_managed_active_directory
            __props__['skip_final_backup'] = skip_final_backup
            if storage_capacity is None:
                raise TypeError("Missing required property 'storage_capacity'")
            __props__['storage_capacity'] = storage_capacity
            __props__['storage_type'] = storage_type
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            if throughput_capacity is None:
                raise TypeError("Missing required property 'throughput_capacity'")
            __props__['throughput_capacity'] = throughput_capacity
            __props__['weekly_maintenance_start_time'] = weekly_maintenance_start_time
            __props__['arn'] = None
            __props__['dns_name'] = None
            __props__['network_interface_ids'] = None
            __props__['owner_id'] = None
            __props__['preferred_file_server_ip'] = None
            __props__['remote_administration_endpoint'] = None
            __props__['vpc_id'] = None
        super(WindowsFileSystem, __self__).__init__(
            'aws:fsx/windowsFileSystem:WindowsFileSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_directory_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            automatic_backup_retention_days: Optional[pulumi.Input[float]] = None,
            copy_tags_to_backups: Optional[pulumi.Input[bool]] = None,
            daily_automatic_backup_start_time: Optional[pulumi.Input[str]] = None,
            deployment_type: Optional[pulumi.Input[str]] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            network_interface_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            preferred_file_server_ip: Optional[pulumi.Input[str]] = None,
            preferred_subnet_id: Optional[pulumi.Input[str]] = None,
            remote_administration_endpoint: Optional[pulumi.Input[str]] = None,
            security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            self_managed_active_directory: Optional[pulumi.Input[pulumi.InputType['WindowsFileSystemSelfManagedActiveDirectoryArgs']]] = None,
            skip_final_backup: Optional[pulumi.Input[bool]] = None,
            storage_capacity: Optional[pulumi.Input[float]] = None,
            storage_type: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            throughput_capacity: Optional[pulumi.Input[float]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            weekly_maintenance_start_time: Optional[pulumi.Input[str]] = None) -> 'WindowsFileSystem':
        """
        Get an existing WindowsFileSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["active_directory_id"] = active_directory_id
        __props__["arn"] = arn
        __props__["automatic_backup_retention_days"] = automatic_backup_retention_days
        __props__["copy_tags_to_backups"] = copy_tags_to_backups
        __props__["daily_automatic_backup_start_time"] = daily_automatic_backup_start_time
        __props__["deployment_type"] = deployment_type
        __props__["dns_name"] = dns_name
        __props__["kms_key_id"] = kms_key_id
        __props__["network_interface_ids"] = network_interface_ids
        __props__["owner_id"] = owner_id
        __props__["preferred_file_server_ip"] = preferred_file_server_ip
        __props__["preferred_subnet_id"] = preferred_subnet_id
        __props__["remote_administration_endpoint"] = remote_administration_endpoint
        __props__["security_group_ids"] = security_group_ids
        __props__["self_managed_active_directory"] = self_managed_active_directory
        __props__["skip_final_backup"] = skip_final_backup
        __props__["storage_capacity"] = storage_capacity
        __props__["storage_type"] = storage_type
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["throughput_capacity"] = throughput_capacity
        __props__["vpc_id"] = vpc_id
        __props__["weekly_maintenance_start_time"] = weekly_maintenance_start_time
        return WindowsFileSystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeDirectoryId")
    def active_directory_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "active_directory_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="automaticBackupRetentionDays")
    def automatic_backup_retention_days(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "automatic_backup_retention_days")

    @property
    @pulumi.getter(name="copyTagsToBackups")
    def copy_tags_to_backups(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "copy_tags_to_backups")

    @property
    @pulumi.getter(name="dailyAutomaticBackupStartTime")
    def daily_automatic_backup_start_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "daily_automatic_backup_start_time")

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "deployment_type")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="networkInterfaceIds")
    def network_interface_ids(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "network_interface_ids")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="preferredFileServerIp")
    def preferred_file_server_ip(self) -> pulumi.Output[str]:
        return pulumi.get(self, "preferred_file_server_ip")

    @property
    @pulumi.getter(name="preferredSubnetId")
    def preferred_subnet_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "preferred_subnet_id")

    @property
    @pulumi.getter(name="remoteAdministrationEndpoint")
    def remote_administration_endpoint(self) -> pulumi.Output[str]:
        return pulumi.get(self, "remote_administration_endpoint")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Optional[List[str]]]:
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="selfManagedActiveDirectory")
    def self_managed_active_directory(self) -> pulumi.Output[Optional['outputs.WindowsFileSystemSelfManagedActiveDirectory']]:
        return pulumi.get(self, "self_managed_active_directory")

    @property
    @pulumi.getter(name="skipFinalBackup")
    def skip_final_backup(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "skip_final_backup")

    @property
    @pulumi.getter(name="storageCapacity")
    def storage_capacity(self) -> pulumi.Output[float]:
        return pulumi.get(self, "storage_capacity")

    @property
    @pulumi.getter(name="storageType")
    def storage_type(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "storage_type")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="throughputCapacity")
    def throughput_capacity(self) -> pulumi.Output[float]:
        return pulumi.get(self, "throughput_capacity")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="weeklyMaintenanceStartTime")
    def weekly_maintenance_start_time(self) -> pulumi.Output[str]:
        return pulumi.get(self, "weekly_maintenance_start_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

