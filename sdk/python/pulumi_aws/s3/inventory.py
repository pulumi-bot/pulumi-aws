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

__all__ = ['Inventory']


class Inventory(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[pulumi.InputType['InventoryDestinationArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 filter: Optional[pulumi.Input[pulumi.InputType['InventoryFilterArgs']]] = None,
                 included_object_versions: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 optional_fields: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 schedule: Optional[pulumi.Input[pulumi.InputType['InventoryScheduleArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a S3 bucket [inventory configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) resource.

        ## Example Usage
        ### Add inventory configuration

        ```python
        import pulumi
        import pulumi_aws as aws

        test_bucket = aws.s3.Bucket("testBucket")
        inventory = aws.s3.Bucket("inventory")
        test_inventory = aws.s3.Inventory("testInventory",
            bucket=test_bucket.id,
            included_object_versions="All",
            schedule={
                "frequency": "Daily",
            },
            destination={
                "bucket": {
                    "format": "ORC",
                    "bucketArn": inventory.arn,
                },
            })
        ```
        ### Add inventory configuration with S3 bucket object prefix

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.s3.Bucket("test")
        inventory = aws.s3.Bucket("inventory")
        test_prefix = aws.s3.Inventory("test-prefix",
            bucket=test.id,
            included_object_versions="All",
            schedule={
                "frequency": "Daily",
            },
            filter={
                "prefix": "documents/",
            },
            destination={
                "bucket": {
                    "format": "ORC",
                    "bucketArn": inventory.arn,
                    "prefix": "inventory",
                },
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket where the inventory configuration will be stored.
        :param pulumi.Input[pulumi.InputType['InventoryDestinationArgs']] destination: Contains information about where to publish the inventory results (documented below).
        :param pulumi.Input[bool] enabled: Specifies whether the inventory is enabled or disabled.
        :param pulumi.Input[pulumi.InputType['InventoryFilterArgs']] filter: Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
        :param pulumi.Input[str] included_object_versions: Object versions to include in the inventory list. Valid values: `All`, `Current`.
        :param pulumi.Input[str] name: Unique identifier of the inventory configuration for the bucket.
        :param pulumi.Input[List[pulumi.Input[str]]] optional_fields: List of optional fields that are included in the inventory results.
               Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
        :param pulumi.Input[pulumi.InputType['InventoryScheduleArgs']] schedule: Specifies the schedule for generating inventory results (documented below).
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

            if bucket is None:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            if destination is None:
                raise TypeError("Missing required property 'destination'")
            __props__['destination'] = destination
            __props__['enabled'] = enabled
            __props__['filter'] = filter
            if included_object_versions is None:
                raise TypeError("Missing required property 'included_object_versions'")
            __props__['included_object_versions'] = included_object_versions
            __props__['name'] = name
            __props__['optional_fields'] = optional_fields
            if schedule is None:
                raise TypeError("Missing required property 'schedule'")
            __props__['schedule'] = schedule
        super(Inventory, __self__).__init__(
            'aws:s3/inventory:Inventory',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            destination: Optional[pulumi.Input[pulumi.InputType['InventoryDestinationArgs']]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            filter: Optional[pulumi.Input[pulumi.InputType['InventoryFilterArgs']]] = None,
            included_object_versions: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            optional_fields: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            schedule: Optional[pulumi.Input[pulumi.InputType['InventoryScheduleArgs']]] = None) -> 'Inventory':
        """
        Get an existing Inventory resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket where the inventory configuration will be stored.
        :param pulumi.Input[pulumi.InputType['InventoryDestinationArgs']] destination: Contains information about where to publish the inventory results (documented below).
        :param pulumi.Input[bool] enabled: Specifies whether the inventory is enabled or disabled.
        :param pulumi.Input[pulumi.InputType['InventoryFilterArgs']] filter: Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
        :param pulumi.Input[str] included_object_versions: Object versions to include in the inventory list. Valid values: `All`, `Current`.
        :param pulumi.Input[str] name: Unique identifier of the inventory configuration for the bucket.
        :param pulumi.Input[List[pulumi.Input[str]]] optional_fields: List of optional fields that are included in the inventory results.
               Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
        :param pulumi.Input[pulumi.InputType['InventoryScheduleArgs']] schedule: Specifies the schedule for generating inventory results (documented below).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket"] = bucket
        __props__["destination"] = destination
        __props__["enabled"] = enabled
        __props__["filter"] = filter
        __props__["included_object_versions"] = included_object_versions
        __props__["name"] = name
        __props__["optional_fields"] = optional_fields
        __props__["schedule"] = schedule
        return Inventory(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        The name of the bucket where the inventory configuration will be stored.
        """
        ...

    @property
    @pulumi.getter
    def destination(self) -> 'outputs.InventoryDestination':
        """
        Contains information about where to publish the inventory results (documented below).
        """
        ...

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Specifies whether the inventory is enabled or disabled.
        """
        ...

    @property
    @pulumi.getter
    def filter(self) -> Optional['outputs.InventoryFilter']:
        """
        Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
        """
        ...

    @property
    @pulumi.getter(name="includedObjectVersions")
    def included_object_versions(self) -> str:
        """
        Object versions to include in the inventory list. Valid values: `All`, `Current`.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique identifier of the inventory configuration for the bucket.
        """
        ...

    @property
    @pulumi.getter(name="optionalFields")
    def optional_fields(self) -> Optional[List[str]]:
        """
        List of optional fields that are included in the inventory results.
        Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
        """
        ...

    @property
    @pulumi.getter
    def schedule(self) -> 'outputs.InventorySchedule':
        """
        Specifies the schedule for generating inventory results (documented below).
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

