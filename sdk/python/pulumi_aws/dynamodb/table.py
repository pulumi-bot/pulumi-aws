# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Table(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The arn of the table
    """
    attributes: pulumi.Output[list]
    """
    List of nested attribute definitions. Only required for `hash_key` and `range_key` attributes. Each attribute has two properties:

      * `name` (`str`) - The name of the index
      * `type` (`str`) - Attribute type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data
    """
    billing_mode: pulumi.Output[str]
    """
    Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
    """
    global_secondary_indexes: pulumi.Output[list]
    """
    Describe a GSI for the table;
    subject to the normal limits on the number of GSIs, projected
    attributes, etc.

      * `hash_key` (`str`) - The name of the hash key in the index; must be
        defined as an attribute in the resource.
      * `name` (`str`) - The name of the index
      * `nonKeyAttributes` (`list`) - Only required with `INCLUDE` as a
        projection type; a list of attributes to project into the index. These
        do not need to be defined as attributes on the table.
      * `projectionType` (`str`) - One of `ALL`, `INCLUDE` or `KEYS_ONLY`
        where `ALL` projects every attribute into the index, `KEYS_ONLY`
        projects just the hash and range key into the index, and `INCLUDE`
        projects only the keys specified in the _non_key_attributes_
        parameter.
      * `range_key` (`str`) - The name of the range key; must be defined
      * `read_capacity` (`float`) - The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
      * `write_capacity` (`float`) - The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
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

      * `name` (`str`) - The name of the index
      * `nonKeyAttributes` (`list`) - Only required with `INCLUDE` as a
        projection type; a list of attributes to project into the index. These
        do not need to be defined as attributes on the table.
      * `projectionType` (`str`) - One of `ALL`, `INCLUDE` or `KEYS_ONLY`
        where `ALL` projects every attribute into the index, `KEYS_ONLY`
        projects just the hash and range key into the index, and `INCLUDE`
        projects only the keys specified in the _non_key_attributes_
        parameter.
      * `range_key` (`str`) - The name of the range key; must be defined
    """
    name: pulumi.Output[str]
    """
    The name of the index
    """
    point_in_time_recovery: pulumi.Output[dict]
    """
    Point-in-time recovery options.

      * `enabled` (`bool`) - Indicates whether ttl is enabled (true) or disabled (false).
    """
    range_key: pulumi.Output[str]
    """
    The name of the range key; must be defined
    """
    read_capacity: pulumi.Output[float]
    """
    The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
    """
    replicas: pulumi.Output[list]
    """
    Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.

      * `regionName` (`str`) - Region name of the replica.
    """
    server_side_encryption: pulumi.Output[dict]
    """
    Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.

      * `enabled` (`bool`) - Indicates whether ttl is enabled (true) or disabled (false).
      * `kms_key_arn` (`str`) - The ARN of the CMK that should be used for the AWS KMS encryption.
        This attribute should only be specified if the key is different from the default DynamoDB CMK, `alias/aws/dynamodb`.
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

      * `attributeName` (`str`) - The name of the table attribute to store the TTL timestamp in.
      * `enabled` (`bool`) - Indicates whether ttl is enabled (true) or disabled (false).
    """
    write_capacity: pulumi.Output[float]
    """
    The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.
    """
    def __init__(__self__, resource_name, opts=None, attributes=None, billing_mode=None, global_secondary_indexes=None, hash_key=None, local_secondary_indexes=None, name=None, point_in_time_recovery=None, range_key=None, read_capacity=None, replicas=None, server_side_encryption=None, stream_enabled=None, stream_view_type=None, tags=None, ttl=None, write_capacity=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DynamoDB table resource

        > **Note:** It is recommended to use [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) for `read_capacity` and/or `write_capacity` if there's `autoscaling policy` attached to the table.

        ## Example Usage

        The following dynamodb table description models the table and GSI shown
        in the [AWS SDK example documentation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html)

        ```python
        import pulumi
        import pulumi_aws as aws

        basic_dynamodb_table = aws.dynamodb.Table("basic-dynamodb-table",
            attributes=[
                {
                    "name": "UserId",
                    "type": "S",
                },
                {
                    "name": "GameTitle",
                    "type": "S",
                },
                {
                    "name": "TopScore",
                    "type": "N",
                },
            ],
            billing_mode="PROVISIONED",
            global_secondary_indexes=[{
                "hash_key": "GameTitle",
                "name": "GameTitleIndex",
                "nonKeyAttributes": ["UserId"],
                "projectionType": "INCLUDE",
                "range_key": "TopScore",
                "read_capacity": 10,
                "write_capacity": 10,
            }],
            hash_key="UserId",
            range_key="GameTitle",
            read_capacity=20,
            tags={
                "Environment": "production",
                "Name": "dynamodb-table-1",
            },
            ttl={
                "attributeName": "TimeToExist",
                "enabled": False,
            },
            write_capacity=20)
        ```
        ### Global Tables

        This resource implements support for [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) via `replica` configuration blocks. For working with [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html), see the `dynamodb.GlobalTable` resource.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.dynamodb.Table("example",
            attributes=[{
                "name": "TestTableHashKey",
                "type": "S",
            }],
            billing_mode="PAY_PER_REQUEST",
            hash_key="TestTableHashKey",
            replicas=[
                {
                    "regionName": "us-east-2",
                },
                {
                    "regionName": "us-west-2",
                },
            ],
            stream_enabled=True,
            stream_view_type="NEW_AND_OLD_IMAGES")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] attributes: List of nested attribute definitions. Only required for `hash_key` and `range_key` attributes. Each attribute has two properties:
        :param pulumi.Input[str] billing_mode: Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
        :param pulumi.Input[list] global_secondary_indexes: Describe a GSI for the table;
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
        :param pulumi.Input[float] read_capacity: The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        :param pulumi.Input[list] replicas: Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.
        :param pulumi.Input[dict] server_side_encryption: Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
        :param pulumi.Input[bool] stream_enabled: Indicates whether Streams are to be enabled (true) or disabled (false).
        :param pulumi.Input[str] stream_view_type: When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
        :param pulumi.Input[dict] tags: A map of tags to populate on the created table.
        :param pulumi.Input[dict] ttl: Defines ttl, has two properties, and can only be specified once:
        :param pulumi.Input[float] write_capacity: The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.

        The **attributes** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the index
          * `type` (`pulumi.Input[str]`) - Attribute type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data

        The **global_secondary_indexes** object supports the following:

          * `hash_key` (`pulumi.Input[str]`) - The name of the hash key in the index; must be
            defined as an attribute in the resource.
          * `name` (`pulumi.Input[str]`) - The name of the index
          * `nonKeyAttributes` (`pulumi.Input[list]`) - Only required with `INCLUDE` as a
            projection type; a list of attributes to project into the index. These
            do not need to be defined as attributes on the table.
          * `projectionType` (`pulumi.Input[str]`) - One of `ALL`, `INCLUDE` or `KEYS_ONLY`
            where `ALL` projects every attribute into the index, `KEYS_ONLY`
            projects just the hash and range key into the index, and `INCLUDE`
            projects only the keys specified in the _non_key_attributes_
            parameter.
          * `range_key` (`pulumi.Input[str]`) - The name of the range key; must be defined
          * `read_capacity` (`pulumi.Input[float]`) - The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
          * `write_capacity` (`pulumi.Input[float]`) - The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.

        The **local_secondary_indexes** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the index
          * `nonKeyAttributes` (`pulumi.Input[list]`) - Only required with `INCLUDE` as a
            projection type; a list of attributes to project into the index. These
            do not need to be defined as attributes on the table.
          * `projectionType` (`pulumi.Input[str]`) - One of `ALL`, `INCLUDE` or `KEYS_ONLY`
            where `ALL` projects every attribute into the index, `KEYS_ONLY`
            projects just the hash and range key into the index, and `INCLUDE`
            projects only the keys specified in the _non_key_attributes_
            parameter.
          * `range_key` (`pulumi.Input[str]`) - The name of the range key; must be defined

        The **point_in_time_recovery** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Indicates whether ttl is enabled (true) or disabled (false).

        The **replicas** object supports the following:

          * `regionName` (`pulumi.Input[str]`) - Region name of the replica.

        The **server_side_encryption** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Indicates whether ttl is enabled (true) or disabled (false).
          * `kms_key_arn` (`pulumi.Input[str]`) - The ARN of the CMK that should be used for the AWS KMS encryption.
            This attribute should only be specified if the key is different from the default DynamoDB CMK, `alias/aws/dynamodb`.

        The **ttl** object supports the following:

          * `attributeName` (`pulumi.Input[str]`) - The name of the table attribute to store the TTL timestamp in.
          * `enabled` (`pulumi.Input[bool]`) - Indicates whether ttl is enabled (true) or disabled (false).
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

            if attributes is None:
                raise TypeError("Missing required property 'attributes'")
            __props__['attributes'] = attributes
            __props__['billingMode'] = billing_mode
            __props__['globalSecondaryIndexes'] = global_secondary_indexes
            if hash_key is None:
                raise TypeError("Missing required property 'hash_key'")
            __props__['hashKey'] = hash_key
            __props__['localSecondaryIndexes'] = local_secondary_indexes
            __props__['name'] = name
            __props__['pointInTimeRecovery'] = point_in_time_recovery
            __props__['rangeKey'] = range_key
            __props__['readCapacity'] = read_capacity
            __props__['replicas'] = replicas
            __props__['serverSideEncryption'] = server_side_encryption
            __props__['streamEnabled'] = stream_enabled
            __props__['streamViewType'] = stream_view_type
            __props__['tags'] = tags
            __props__['ttl'] = ttl
            __props__['writeCapacity'] = write_capacity
            __props__['arn'] = None
            __props__['stream_arn'] = None
            __props__['stream_label'] = None
        super(Table, __self__).__init__(
            'aws:dynamodb/table:Table',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, attributes=None, billing_mode=None, global_secondary_indexes=None, hash_key=None, local_secondary_indexes=None, name=None, point_in_time_recovery=None, range_key=None, read_capacity=None, replicas=None, server_side_encryption=None, stream_arn=None, stream_enabled=None, stream_label=None, stream_view_type=None, tags=None, ttl=None, write_capacity=None):
        """
        Get an existing Table resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The arn of the table
        :param pulumi.Input[list] attributes: List of nested attribute definitions. Only required for `hash_key` and `range_key` attributes. Each attribute has two properties:
        :param pulumi.Input[str] billing_mode: Controls how you are charged for read and write throughput and how you manage capacity. The valid values are `PROVISIONED` and `PAY_PER_REQUEST`. Defaults to `PROVISIONED`.
        :param pulumi.Input[list] global_secondary_indexes: Describe a GSI for the table;
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
        :param pulumi.Input[float] read_capacity: The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
        :param pulumi.Input[list] replicas: Configuration block(s) with [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html) replication configurations. Detailed below.
        :param pulumi.Input[dict] server_side_encryption: Encryption at rest options. AWS DynamoDB tables are automatically encrypted at rest with an AWS owned Customer Master Key if this argument isn't specified.
        :param pulumi.Input[str] stream_arn: The ARN of the Table Stream. Only available when `stream_enabled = true`
        :param pulumi.Input[bool] stream_enabled: Indicates whether Streams are to be enabled (true) or disabled (false).
        :param pulumi.Input[str] stream_label: A timestamp, in ISO 8601 format, for this stream. Note that this timestamp is not
               a unique identifier for the stream on its own. However, the combination of AWS customer ID,
               table name and this field is guaranteed to be unique.
               It can be used for creating CloudWatch Alarms. Only available when `stream_enabled = true`
        :param pulumi.Input[str] stream_view_type: When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values are `KEYS_ONLY`, `NEW_IMAGE`, `OLD_IMAGE`, `NEW_AND_OLD_IMAGES`.
        :param pulumi.Input[dict] tags: A map of tags to populate on the created table.
        :param pulumi.Input[dict] ttl: Defines ttl, has two properties, and can only be specified once:
        :param pulumi.Input[float] write_capacity: The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.

        The **attributes** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the index
          * `type` (`pulumi.Input[str]`) - Attribute type, which must be a scalar type: `S`, `N`, or `B` for (S)tring, (N)umber or (B)inary data

        The **global_secondary_indexes** object supports the following:

          * `hash_key` (`pulumi.Input[str]`) - The name of the hash key in the index; must be
            defined as an attribute in the resource.
          * `name` (`pulumi.Input[str]`) - The name of the index
          * `nonKeyAttributes` (`pulumi.Input[list]`) - Only required with `INCLUDE` as a
            projection type; a list of attributes to project into the index. These
            do not need to be defined as attributes on the table.
          * `projectionType` (`pulumi.Input[str]`) - One of `ALL`, `INCLUDE` or `KEYS_ONLY`
            where `ALL` projects every attribute into the index, `KEYS_ONLY`
            projects just the hash and range key into the index, and `INCLUDE`
            projects only the keys specified in the _non_key_attributes_
            parameter.
          * `range_key` (`pulumi.Input[str]`) - The name of the range key; must be defined
          * `read_capacity` (`pulumi.Input[float]`) - The number of read units for this index. Must be set if billing_mode is set to PROVISIONED.
          * `write_capacity` (`pulumi.Input[float]`) - The number of write units for this index. Must be set if billing_mode is set to PROVISIONED.

        The **local_secondary_indexes** object supports the following:

          * `name` (`pulumi.Input[str]`) - The name of the index
          * `nonKeyAttributes` (`pulumi.Input[list]`) - Only required with `INCLUDE` as a
            projection type; a list of attributes to project into the index. These
            do not need to be defined as attributes on the table.
          * `projectionType` (`pulumi.Input[str]`) - One of `ALL`, `INCLUDE` or `KEYS_ONLY`
            where `ALL` projects every attribute into the index, `KEYS_ONLY`
            projects just the hash and range key into the index, and `INCLUDE`
            projects only the keys specified in the _non_key_attributes_
            parameter.
          * `range_key` (`pulumi.Input[str]`) - The name of the range key; must be defined

        The **point_in_time_recovery** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Indicates whether ttl is enabled (true) or disabled (false).

        The **replicas** object supports the following:

          * `regionName` (`pulumi.Input[str]`) - Region name of the replica.

        The **server_side_encryption** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Indicates whether ttl is enabled (true) or disabled (false).
          * `kms_key_arn` (`pulumi.Input[str]`) - The ARN of the CMK that should be used for the AWS KMS encryption.
            This attribute should only be specified if the key is different from the default DynamoDB CMK, `alias/aws/dynamodb`.

        The **ttl** object supports the following:

          * `attributeName` (`pulumi.Input[str]`) - The name of the table attribute to store the TTL timestamp in.
          * `enabled` (`pulumi.Input[bool]`) - Indicates whether ttl is enabled (true) or disabled (false).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["attributes"] = attributes
        __props__["billing_mode"] = billing_mode
        __props__["global_secondary_indexes"] = global_secondary_indexes
        __props__["hash_key"] = hash_key
        __props__["local_secondary_indexes"] = local_secondary_indexes
        __props__["name"] = name
        __props__["point_in_time_recovery"] = point_in_time_recovery
        __props__["range_key"] = range_key
        __props__["read_capacity"] = read_capacity
        __props__["replicas"] = replicas
        __props__["server_side_encryption"] = server_side_encryption
        __props__["stream_arn"] = stream_arn
        __props__["stream_enabled"] = stream_enabled
        __props__["stream_label"] = stream_label
        __props__["stream_view_type"] = stream_view_type
        __props__["tags"] = tags
        __props__["ttl"] = ttl
        __props__["write_capacity"] = write_capacity
        return Table(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
