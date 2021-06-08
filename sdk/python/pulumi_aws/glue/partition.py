# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PartitionArgs', 'Partition']

@pulumi.input_type
class PartitionArgs:
    def __init__(__self__, *,
                 database_name: pulumi.Input[str],
                 partition_values: pulumi.Input[Sequence[pulumi.Input[str]]],
                 table_name: pulumi.Input[str],
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 storage_descriptor: Optional[pulumi.Input['PartitionStorageDescriptorArgs']] = None):
        """
        The set of arguments for constructing a Partition resource.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] partition_values: The values that define the partition.
        :param pulumi.Input[str] catalog_id: ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of initialization parameters for the SerDe, in key-value form.
        :param pulumi.Input['PartitionStorageDescriptorArgs'] storage_descriptor: A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "partition_values", partition_values)
        pulumi.set(__self__, "table_name", table_name)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if storage_descriptor is not None:
            pulumi.set(__self__, "storage_descriptor", storage_descriptor)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Input[str]:
        """
        Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="partitionValues")
    def partition_values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The values that define the partition.
        """
        return pulumi.get(self, "partition_values")

    @partition_values.setter
    def partition_values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "partition_values", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of initialization parameters for the SerDe, in key-value form.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="storageDescriptor")
    def storage_descriptor(self) -> Optional[pulumi.Input['PartitionStorageDescriptorArgs']]:
        """
        A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
        """
        return pulumi.get(self, "storage_descriptor")

    @storage_descriptor.setter
    def storage_descriptor(self, value: Optional[pulumi.Input['PartitionStorageDescriptorArgs']]):
        pulumi.set(self, "storage_descriptor", value)


@pulumi.input_type
class _PartitionState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 last_accessed_time: Optional[pulumi.Input[str]] = None,
                 last_analyzed_time: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 partition_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 storage_descriptor: Optional[pulumi.Input['PartitionStorageDescriptorArgs']] = None,
                 table_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Partition resources.
        :param pulumi.Input[str] catalog_id: ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
        :param pulumi.Input[str] creation_time: The time at which the partition was created.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input[str] last_accessed_time: The last time at which the partition was accessed.
        :param pulumi.Input[str] last_analyzed_time: The last time at which column statistics were computed for this partition.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of initialization parameters for the SerDe, in key-value form.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] partition_values: The values that define the partition.
        :param pulumi.Input['PartitionStorageDescriptorArgs'] storage_descriptor: A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
        """
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if database_name is not None:
            pulumi.set(__self__, "database_name", database_name)
        if last_accessed_time is not None:
            pulumi.set(__self__, "last_accessed_time", last_accessed_time)
        if last_analyzed_time is not None:
            pulumi.set(__self__, "last_analyzed_time", last_analyzed_time)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if partition_values is not None:
            pulumi.set(__self__, "partition_values", partition_values)
        if storage_descriptor is not None:
            pulumi.set(__self__, "storage_descriptor", storage_descriptor)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The time at which the partition was created.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        """
        return pulumi.get(self, "database_name")

    @database_name.setter
    def database_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "database_name", value)

    @property
    @pulumi.getter(name="lastAccessedTime")
    def last_accessed_time(self) -> Optional[pulumi.Input[str]]:
        """
        The last time at which the partition was accessed.
        """
        return pulumi.get(self, "last_accessed_time")

    @last_accessed_time.setter
    def last_accessed_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_accessed_time", value)

    @property
    @pulumi.getter(name="lastAnalyzedTime")
    def last_analyzed_time(self) -> Optional[pulumi.Input[str]]:
        """
        The last time at which column statistics were computed for this partition.
        """
        return pulumi.get(self, "last_analyzed_time")

    @last_analyzed_time.setter
    def last_analyzed_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_analyzed_time", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of initialization parameters for the SerDe, in key-value form.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter(name="partitionValues")
    def partition_values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The values that define the partition.
        """
        return pulumi.get(self, "partition_values")

    @partition_values.setter
    def partition_values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "partition_values", value)

    @property
    @pulumi.getter(name="storageDescriptor")
    def storage_descriptor(self) -> Optional[pulumi.Input['PartitionStorageDescriptorArgs']]:
        """
        A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
        """
        return pulumi.get(self, "storage_descriptor")

    @storage_descriptor.setter
    def storage_descriptor(self, value: Optional[pulumi.Input['PartitionStorageDescriptorArgs']]):
        pulumi.set(self, "storage_descriptor", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)


class Partition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 partition_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 storage_descriptor: Optional[pulumi.Input[pulumi.InputType['PartitionStorageDescriptorArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a Glue Partition Resource.

        ## Import

        Glue Partitions can be imported with their catalog ID (usually AWS account ID), database name, table name and partition values e.g.

        ```sh
         $ pulumi import aws:glue/partition:Partition part 123456789012:MyDatabase:MyTable:val1#val2
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of initialization parameters for the SerDe, in key-value form.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] partition_values: The values that define the partition.
        :param pulumi.Input[pulumi.InputType['PartitionStorageDescriptorArgs']] storage_descriptor: A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: PartitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a Glue Partition Resource.

        ## Import

        Glue Partitions can be imported with their catalog ID (usually AWS account ID), database name, table name and partition values e.g.

        ```sh
         $ pulumi import aws:glue/partition:Partition part 123456789012:MyDatabase:MyTable:val1#val2
        ```

        :param str resource_name_: The name of the resource.
        :param PartitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PartitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 partition_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 storage_descriptor: Optional[pulumi.Input[pulumi.InputType['PartitionStorageDescriptorArgs']]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PartitionArgs.__new__(PartitionArgs)

            __props__.__dict__["catalog_id"] = catalog_id
            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__.__dict__["database_name"] = database_name
            __props__.__dict__["parameters"] = parameters
            if partition_values is None and not opts.urn:
                raise TypeError("Missing required property 'partition_values'")
            __props__.__dict__["partition_values"] = partition_values
            __props__.__dict__["storage_descriptor"] = storage_descriptor
            if table_name is None and not opts.urn:
                raise TypeError("Missing required property 'table_name'")
            __props__.__dict__["table_name"] = table_name
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["last_accessed_time"] = None
            __props__.__dict__["last_analyzed_time"] = None
        super(Partition, __self__).__init__(
            'aws:glue/partition:Partition',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            database_name: Optional[pulumi.Input[str]] = None,
            last_accessed_time: Optional[pulumi.Input[str]] = None,
            last_analyzed_time: Optional[pulumi.Input[str]] = None,
            parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            partition_values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            storage_descriptor: Optional[pulumi.Input[pulumi.InputType['PartitionStorageDescriptorArgs']]] = None,
            table_name: Optional[pulumi.Input[str]] = None) -> 'Partition':
        """
        Get an existing Partition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
        :param pulumi.Input[str] creation_time: The time at which the partition was created.
        :param pulumi.Input[str] database_name: Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        :param pulumi.Input[str] last_accessed_time: The last time at which the partition was accessed.
        :param pulumi.Input[str] last_analyzed_time: The last time at which column statistics were computed for this partition.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] parameters: A map of initialization parameters for the SerDe, in key-value form.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] partition_values: The values that define the partition.
        :param pulumi.Input[pulumi.InputType['PartitionStorageDescriptorArgs']] storage_descriptor: A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PartitionState.__new__(_PartitionState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["database_name"] = database_name
        __props__.__dict__["last_accessed_time"] = last_accessed_time
        __props__.__dict__["last_analyzed_time"] = last_analyzed_time
        __props__.__dict__["parameters"] = parameters
        __props__.__dict__["partition_values"] = partition_values
        __props__.__dict__["storage_descriptor"] = storage_descriptor
        __props__.__dict__["table_name"] = table_name
        return Partition(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        ID of the Glue Catalog and database to create the table in. If omitted, this defaults to the AWS Account ID plus the database name.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The time at which the partition was created.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> pulumi.Output[str]:
        """
        Name of the metadata database where the table metadata resides. For Hive compatibility, this must be all lowercase.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="lastAccessedTime")
    def last_accessed_time(self) -> pulumi.Output[str]:
        """
        The last time at which the partition was accessed.
        """
        return pulumi.get(self, "last_accessed_time")

    @property
    @pulumi.getter(name="lastAnalyzedTime")
    def last_analyzed_time(self) -> pulumi.Output[str]:
        """
        The last time at which column statistics were computed for this partition.
        """
        return pulumi.get(self, "last_analyzed_time")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of initialization parameters for the SerDe, in key-value form.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="partitionValues")
    def partition_values(self) -> pulumi.Output[Sequence[str]]:
        """
        The values that define the partition.
        """
        return pulumi.get(self, "partition_values")

    @property
    @pulumi.getter(name="storageDescriptor")
    def storage_descriptor(self) -> pulumi.Output[Optional['outputs.PartitionStorageDescriptor']]:
        """
        A storage descriptor object containing information about the physical storage of this table. You can refer to the [Glue Developer Guide](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-catalog-tables.html#aws-glue-api-catalog-tables-StorageDescriptor) for a full explanation of this object.
        """
        return pulumi.get(self, "storage_descriptor")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "table_name")

