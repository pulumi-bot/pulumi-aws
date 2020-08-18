# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['LustreFileSystem']


class LustreFileSystem(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_type: Optional[pulumi.Input[str]] = None,
                 export_path: Optional[pulumi.Input[str]] = None,
                 import_path: Optional[pulumi.Input[str]] = None,
                 imported_file_chunk_size: Optional[pulumi.Input[float]] = None,
                 per_unit_storage_throughput: Optional[pulumi.Input[float]] = None,
                 security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 storage_capacity: Optional[pulumi.Input[float]] = None,
                 subnet_ids: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 weekly_maintenance_start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a FSx Lustre File System. See the [FSx Lustre Guide](https://docs.aws.amazon.com/fsx/latest/LustreGuide/what-is.html) for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.fsx.LustreFileSystem("example",
            import_path=f"s3://{aws_s3_bucket['example']['bucket']}",
            storage_capacity=1200,
            subnet_ids=[aws_subnet["example"]["id"]])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deployment_type: - The filesystem deployment type. One of: `SCRATCH_1`, `SCRATCH_2`, `PERSISTENT_1`.
        :param pulumi.Input[str] export_path: S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        :param pulumi.Input[str] import_path: S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        :param pulumi.Input[float] imported_file_chunk_size: For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        :param pulumi.Input[float] per_unit_storage_throughput: - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the `PERSISTENT_1` deployment_type. For valid values, see the [AWS documentation](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemLustreConfiguration.html).
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_ids: A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        :param pulumi.Input[float] storage_capacity: The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
        :param pulumi.Input[str] subnet_ids: A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the file system.
        :param pulumi.Input[str] weekly_maintenance_start_time: The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
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

            __props__['deployment_type'] = deployment_type
            __props__['export_path'] = export_path
            __props__['import_path'] = import_path
            __props__['imported_file_chunk_size'] = imported_file_chunk_size
            __props__['per_unit_storage_throughput'] = per_unit_storage_throughput
            __props__['security_group_ids'] = security_group_ids
            if storage_capacity is None:
                raise TypeError("Missing required property 'storage_capacity'")
            __props__['storage_capacity'] = storage_capacity
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            __props__['weekly_maintenance_start_time'] = weekly_maintenance_start_time
            __props__['arn'] = None
            __props__['dns_name'] = None
            __props__['network_interface_ids'] = None
            __props__['owner_id'] = None
            __props__['vpc_id'] = None
        super(LustreFileSystem, __self__).__init__(
            'aws:fsx/lustreFileSystem:LustreFileSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            deployment_type: Optional[pulumi.Input[str]] = None,
            dns_name: Optional[pulumi.Input[str]] = None,
            export_path: Optional[pulumi.Input[str]] = None,
            import_path: Optional[pulumi.Input[str]] = None,
            imported_file_chunk_size: Optional[pulumi.Input[float]] = None,
            network_interface_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            per_unit_storage_throughput: Optional[pulumi.Input[float]] = None,
            security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            storage_capacity: Optional[pulumi.Input[float]] = None,
            subnet_ids: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            weekly_maintenance_start_time: Optional[pulumi.Input[str]] = None) -> 'LustreFileSystem':
        """
        Get an existing LustreFileSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name of the file system.
        :param pulumi.Input[str] deployment_type: - The filesystem deployment type. One of: `SCRATCH_1`, `SCRATCH_2`, `PERSISTENT_1`.
        :param pulumi.Input[str] dns_name: DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
        :param pulumi.Input[str] export_path: S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        :param pulumi.Input[str] import_path: S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        :param pulumi.Input[float] imported_file_chunk_size: For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        :param pulumi.Input[List[pulumi.Input[str]]] network_interface_ids: Set of Elastic Network Interface identifiers from which the file system is accessible.
        :param pulumi.Input[str] owner_id: AWS account identifier that created the file system.
        :param pulumi.Input[float] per_unit_storage_throughput: - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the `PERSISTENT_1` deployment_type. For valid values, see the [AWS documentation](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemLustreConfiguration.html).
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_ids: A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        :param pulumi.Input[float] storage_capacity: The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
        :param pulumi.Input[str] subnet_ids: A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the file system.
        :param pulumi.Input[str] vpc_id: Identifier of the Virtual Private Cloud for the file system.
        :param pulumi.Input[str] weekly_maintenance_start_time: The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["deployment_type"] = deployment_type
        __props__["dns_name"] = dns_name
        __props__["export_path"] = export_path
        __props__["import_path"] = import_path
        __props__["imported_file_chunk_size"] = imported_file_chunk_size
        __props__["network_interface_ids"] = network_interface_ids
        __props__["owner_id"] = owner_id
        __props__["per_unit_storage_throughput"] = per_unit_storage_throughput
        __props__["security_group_ids"] = security_group_ids
        __props__["storage_capacity"] = storage_capacity
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        __props__["weekly_maintenance_start_time"] = weekly_maintenance_start_time
        return LustreFileSystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name of the file system.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="deploymentType")
    def deployment_type(self) -> Optional[str]:
        """
        - The filesystem deployment type. One of: `SCRATCH_1`, `SCRATCH_2`, `PERSISTENT_1`.
        """
        return pulumi.get(self, "deployment_type")

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        """
        DNS name for the file system, e.g. `fs-12345678.fsx.us-west-2.amazonaws.com`
        """
        return pulumi.get(self, "dns_name")

    @property
    @pulumi.getter(name="exportPath")
    def export_path(self) -> str:
        """
        S3 URI (with optional prefix) where the root of your Amazon FSx file system is exported. Can only be specified with `import_path` argument and the path must use the same Amazon S3 bucket as specified in `import_path`. Set equal to `import_path` to overwrite files on export. Defaults to `s3://{IMPORT BUCKET}/FSxLustre{CREATION TIMESTAMP}`.
        """
        return pulumi.get(self, "export_path")

    @property
    @pulumi.getter(name="importPath")
    def import_path(self) -> Optional[str]:
        """
        S3 URI (with optional prefix) that you're using as the data repository for your FSx for Lustre file system. For example, `s3://example-bucket/optional-prefix/`.
        """
        return pulumi.get(self, "import_path")

    @property
    @pulumi.getter(name="importedFileChunkSize")
    def imported_file_chunk_size(self) -> float:
        """
        For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk. Can only be specified with `import_path` argument. Defaults to `1024`. Minimum of `1` and maximum of `512000`.
        """
        return pulumi.get(self, "imported_file_chunk_size")

    @property
    @pulumi.getter(name="networkInterfaceIds")
    def network_interface_ids(self) -> List[str]:
        """
        Set of Elastic Network Interface identifiers from which the file system is accessible.
        """
        return pulumi.get(self, "network_interface_ids")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        AWS account identifier that created the file system.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="perUnitStorageThroughput")
    def per_unit_storage_throughput(self) -> Optional[float]:
        """
        - Describes the amount of read and write throughput for each 1 tebibyte of storage, in MB/s/TiB, required for the `PERSISTENT_1` deployment_type. For valid values, see the [AWS documentation](https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateFileSystemLustreConfiguration.html).
        """
        return pulumi.get(self, "per_unit_storage_throughput")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[List[str]]:
        """
        A list of IDs for the security groups that apply to the specified network interfaces created for file system access. These security groups will apply to all network interfaces.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="storageCapacity")
    def storage_capacity(self) -> float:
        """
        The storage capacity (GiB) of the file system. Minimum of `1200`. Storage capacity is provisioned in increments of 3,600 GiB.
        """
        return pulumi.get(self, "storage_capacity")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> str:
        """
        A list of IDs for the subnets that the file system will be accessible from. File systems currently support only one subnet. The file server is also launched in that subnet's Availability Zone.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the file system.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        Identifier of the Virtual Private Cloud for the file system.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="weeklyMaintenanceStartTime")
    def weekly_maintenance_start_time(self) -> str:
        """
        The preferred start time (in `d:HH:MM` format) to perform weekly maintenance, in the UTC time zone.
        """
        return pulumi.get(self, "weekly_maintenance_start_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

