# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TableItemArgs', 'TableItem']

@pulumi.input_type
class TableItemArgs:
    def __init__(__self__, *,
                 hash_key: pulumi.Input[str],
                 item: pulumi.Input[str],
                 table_name: pulumi.Input[str],
                 range_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TableItem resource.
        :param pulumi.Input[str] hash_key: Hash key to use for lookups and identification of the item
        :param pulumi.Input[str] item: JSON representation of a map of attribute name/value pairs, one for each attribute.
               Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        :param pulumi.Input[str] table_name: The name of the table to contain the item.
        :param pulumi.Input[str] range_key: Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        """
        pulumi.set(__self__, "hash_key", hash_key)
        pulumi.set(__self__, "item", item)
        pulumi.set(__self__, "table_name", table_name)
        if range_key is not None:
            pulumi.set(__self__, "range_key", range_key)

    @property
    @pulumi.getter(name="hashKey")
    def hash_key(self) -> pulumi.Input[str]:
        """
        Hash key to use for lookups and identification of the item
        """
        return pulumi.get(self, "hash_key")

    @hash_key.setter
    def hash_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "hash_key", value)

    @property
    @pulumi.getter
    def item(self) -> pulumi.Input[str]:
        """
        JSON representation of a map of attribute name/value pairs, one for each attribute.
        Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        """
        return pulumi.get(self, "item")

    @item.setter
    def item(self, value: pulumi.Input[str]):
        pulumi.set(self, "item", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Input[str]:
        """
        The name of the table to contain the item.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "table_name", value)

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> Optional[pulumi.Input[str]]:
        """
        Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        """
        return pulumi.get(self, "range_key")

    @range_key.setter
    def range_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range_key", value)


@pulumi.input_type
class _TableItemState:
    def __init__(__self__, *,
                 hash_key: Optional[pulumi.Input[str]] = None,
                 item: Optional[pulumi.Input[str]] = None,
                 range_key: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering TableItem resources.
        :param pulumi.Input[str] hash_key: Hash key to use for lookups and identification of the item
        :param pulumi.Input[str] item: JSON representation of a map of attribute name/value pairs, one for each attribute.
               Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        :param pulumi.Input[str] range_key: Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        :param pulumi.Input[str] table_name: The name of the table to contain the item.
        """
        if hash_key is not None:
            pulumi.set(__self__, "hash_key", hash_key)
        if item is not None:
            pulumi.set(__self__, "item", item)
        if range_key is not None:
            pulumi.set(__self__, "range_key", range_key)
        if table_name is not None:
            pulumi.set(__self__, "table_name", table_name)

    @property
    @pulumi.getter(name="hashKey")
    def hash_key(self) -> Optional[pulumi.Input[str]]:
        """
        Hash key to use for lookups and identification of the item
        """
        return pulumi.get(self, "hash_key")

    @hash_key.setter
    def hash_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hash_key", value)

    @property
    @pulumi.getter
    def item(self) -> Optional[pulumi.Input[str]]:
        """
        JSON representation of a map of attribute name/value pairs, one for each attribute.
        Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        """
        return pulumi.get(self, "item")

    @item.setter
    def item(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "item", value)

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> Optional[pulumi.Input[str]]:
        """
        Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        """
        return pulumi.get(self, "range_key")

    @range_key.setter
    def range_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "range_key", value)

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the table to contain the item.
        """
        return pulumi.get(self, "table_name")

    @table_name.setter
    def table_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "table_name", value)


class TableItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hash_key: Optional[pulumi.Input[str]] = None,
                 item: Optional[pulumi.Input[str]] = None,
                 range_key: Optional[pulumi.Input[str]] = None,
                 table_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a DynamoDB table item resource

        > **Note:** This resource is not meant to be used for managing large amounts of data in your table, it is not designed to scale.
          You should perform **regular backups** of all data in the table, see [AWS docs for more](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/BackupRestore.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_table = aws.dynamodb.Table("exampleTable",
            read_capacity=10,
            write_capacity=10,
            hash_key="exampleHashKey",
            attributes=[aws.dynamodb.TableAttributeArgs(
                name="exampleHashKey",
                type="S",
            )])
        example_table_item = aws.dynamodb.TableItem("exampleTableItem",
            table_name=example_table.name,
            hash_key=example_table.hash_key,
            item=\"\"\"{
          "exampleHashKey": {"S": "something"},
          "one": {"N": "11111"},
          "two": {"N": "22222"},
          "three": {"N": "33333"},
          "four": {"N": "44444"}
        }
        \"\"\")
        ```

        ## Import

        DynamoDB table items cannot be imported.

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hash_key: Hash key to use for lookups and identification of the item
        :param pulumi.Input[str] item: JSON representation of a map of attribute name/value pairs, one for each attribute.
               Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        :param pulumi.Input[str] range_key: Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        :param pulumi.Input[str] table_name: The name of the table to contain the item.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: TableItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a DynamoDB table item resource

        > **Note:** This resource is not meant to be used for managing large amounts of data in your table, it is not designed to scale.
          You should perform **regular backups** of all data in the table, see [AWS docs for more](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/BackupRestore.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_table = aws.dynamodb.Table("exampleTable",
            read_capacity=10,
            write_capacity=10,
            hash_key="exampleHashKey",
            attributes=[aws.dynamodb.TableAttributeArgs(
                name="exampleHashKey",
                type="S",
            )])
        example_table_item = aws.dynamodb.TableItem("exampleTableItem",
            table_name=example_table.name,
            hash_key=example_table.hash_key,
            item=\"\"\"{
          "exampleHashKey": {"S": "something"},
          "one": {"N": "11111"},
          "two": {"N": "22222"},
          "three": {"N": "33333"},
          "four": {"N": "44444"}
        }
        \"\"\")
        ```

        ## Import

        DynamoDB table items cannot be imported.

        :param str resource_name_: The name of the resource.
        :param TableItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TableItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hash_key: Optional[pulumi.Input[str]] = None,
                 item: Optional[pulumi.Input[str]] = None,
                 range_key: Optional[pulumi.Input[str]] = None,
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
            __props__ = TableItemArgs.__new__(TableItemArgs)

            if hash_key is None and not opts.urn:
                raise TypeError("Missing required property 'hash_key'")
            __props__.__dict__["hash_key"] = hash_key
            if item is None and not opts.urn:
                raise TypeError("Missing required property 'item'")
            __props__.__dict__["item"] = item
            __props__.__dict__["range_key"] = range_key
            if table_name is None and not opts.urn:
                raise TypeError("Missing required property 'table_name'")
            __props__.__dict__["table_name"] = table_name
        super(TableItem, __self__).__init__(
            'aws:dynamodb/tableItem:TableItem',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            hash_key: Optional[pulumi.Input[str]] = None,
            item: Optional[pulumi.Input[str]] = None,
            range_key: Optional[pulumi.Input[str]] = None,
            table_name: Optional[pulumi.Input[str]] = None) -> 'TableItem':
        """
        Get an existing TableItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hash_key: Hash key to use for lookups and identification of the item
        :param pulumi.Input[str] item: JSON representation of a map of attribute name/value pairs, one for each attribute.
               Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        :param pulumi.Input[str] range_key: Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        :param pulumi.Input[str] table_name: The name of the table to contain the item.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TableItemState.__new__(_TableItemState)

        __props__.__dict__["hash_key"] = hash_key
        __props__.__dict__["item"] = item
        __props__.__dict__["range_key"] = range_key
        __props__.__dict__["table_name"] = table_name
        return TableItem(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="hashKey")
    def hash_key(self) -> pulumi.Output[str]:
        """
        Hash key to use for lookups and identification of the item
        """
        return pulumi.get(self, "hash_key")

    @property
    @pulumi.getter
    def item(self) -> pulumi.Output[str]:
        """
        JSON representation of a map of attribute name/value pairs, one for each attribute.
        Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.
        """
        return pulumi.get(self, "item")

    @property
    @pulumi.getter(name="rangeKey")
    def range_key(self) -> pulumi.Output[Optional[str]]:
        """
        Range key to use for lookups and identification of the item. Required if there is range key defined in the table.
        """
        return pulumi.get(self, "range_key")

    @property
    @pulumi.getter(name="tableName")
    def table_name(self) -> pulumi.Output[str]:
        """
        The name of the table to contain the item.
        """
        return pulumi.get(self, "table_name")

