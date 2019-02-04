# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Table(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The arn of the table
    """
    attributes: pulumi.Output[list]
    """
    List of nested attribute definitions. Only required for `hash_key` and `range_key` attributes. Each attribute has two properties:
    """
    billing_mode: pulumi.Output[str]
    """
    Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
    """
    global_secondary_indexes: pulumi.Output[list]
    """
    Describe a GSO for the table;
    subject to the normal limits on the number of GSIs, projected
    attributes, etc.
    """
    hash_key: pulumi.Output[str]
    """
    The name of the hash key in the index; must be
    defined as an attribute in the resource.
    """
    local_secondary_indexes: pulumi.Output[list]
    """
    Describe an LSI on the table;
    these can only be allocated *at creation* so you cannot change this
    definition after you have created the resource.
    """
    name: pulumi.Output[str]
    """
    The name of the index
    """
    point_in_time_recovery: pulumi.Output[dict]
    """
    Point-in-time recovery options.
    """
    range_key: pulumi.Output[str]
    """
    The name of the range key; must be defined
    """
    read_capacity: pulumi.Output[int]
    """
    The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
    """
    server_side_encryption: pulumi.Output[dict]
    """
    Encrypt at rest options.
    """
    stream_arn: pulumi.Output[str]
    """
    The ARN of the Table Stream. Only available when `stream_enabled = true`
    """
    stream_enabled: pulumi.Output[bool]
    """
    Indicates whether Streams are to be enabled (true) or disabled (false).
    """
    stream_label: pulumi.Output[str]
    """
    A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
    a unique identifier for the stream on its own. However, the combination of AWS customer ID,
    table name and this field is guaranteed to be unique.
    It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`
    """
    stream_view_type: pulumi.Output[str]
    """
    When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to populate on the created table.
    """
    ttl: pulumi.Output[dict]
    """
    Defines ttl, has two properties, and can only be specified once:
    """
    write_capacity: pulumi.Output[int]
    """
    The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
    """
    def __init__(__self__, __name__, __opts__=None, attributes=None, billing_mode=None, global_secondary_indexes=None, hash_key=None, local_secondary_indexes=None, name=None, point_in_time_recovery=None, range_key=None, read_capacity=None, server_side_encryption=None, stream_enabled=None, stream_view_type=None, tags=None, ttl=None, write_capacity=None):
        """
        Provides a DynamoDB table resource
        
        > **Note:** It is recommended to use `lifecycle` [`ignore_changes`](https://www.terraform.io/docs/configuration/resources.html#ignore_changes) for `read_capacity` and/or `write_capacity` if there's [autoscaling policy](https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy.html) attached to the table.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[list] attributes: List of nested attribute definitions. Only required for `hash_key` and `range_key` attributes. Each attribute has two properties:
        :param pulumi.Input[str] billing_mode: Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
        :param pulumi.Input[list] global_secondary_indexes: Describe a GSO for the table;
               subject to the normal limits on the number of GSIs, projected
               attributes, etc.
        :param pulumi.Input[str] hash_key: The name of the hash key in the index; must be
               defined as an attribute in the resource.
        :param pulumi.Input[list] local_secondary_indexes: Describe an LSI on the table;
               these can only be allocated *at creation* so you cannot change this
               definition after you have created the resource.
        :param pulumi.Input[str] name: The name of the index
        :param pulumi.Input[dict] point_in_time_recovery: Point-in-time recovery options.
        :param pulumi.Input[str] range_key: The name of the range key; must be defined
        :param pulumi.Input[int] read_capacity: The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        :param pulumi.Input[dict] server_side_encryption: Encrypt at rest options.
        :param pulumi.Input[bool] stream_enabled: Indicates whether Streams are to be enabled (true) or disabled (false).
        :param pulumi.Input[str] stream_view_type: When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
        :param pulumi.Input[dict] tags: A map of tags to populate on the created table.
        :param pulumi.Input[dict] ttl: Defines ttl, has two properties, and can only be specified once:
        :param pulumi.Input[int] write_capacity: The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if attributes is None:
            raise TypeError('Missing required property attributes')
        __props__['attributes'] = attributes

        __props__['billing_mode'] = billing_mode

        __props__['global_secondary_indexes'] = global_secondary_indexes

        if hash_key is None:
            raise TypeError('Missing required property hash_key')
        __props__['hash_key'] = hash_key

        __props__['local_secondary_indexes'] = local_secondary_indexes

        __props__['name'] = name

        __props__['point_in_time_recovery'] = point_in_time_recovery

        __props__['range_key'] = range_key

        __props__['read_capacity'] = read_capacity

        __props__['server_side_encryption'] = server_side_encryption

        __props__['stream_enabled'] = stream_enabled

        __props__['stream_view_type'] = stream_view_type

        __props__['tags'] = tags

        __props__['ttl'] = ttl

        __props__['write_capacity'] = write_capacity

        __props__['arn'] = None
        __props__['stream_arn'] = None
        __props__['stream_label'] = None

        super(Table, __self__).__init__(
            'aws:dynamodb/table:Table',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

