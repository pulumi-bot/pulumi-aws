# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ReplicationGroup(pulumi.CustomResource):
    apply_immediately: pulumi.Output[bool]
    """
    Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
    """
    at_rest_encryption_enabled: pulumi.Output[bool]
    """
    Whether to enable encryption at rest.
    """
    auth_token: pulumi.Output[str]
    """
    The password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
    """
    auto_minor_version_upgrade: pulumi.Output[bool]
    """
    Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
    """
    automatic_failover_enabled: pulumi.Output[bool]
    """
    Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
    """
    availability_zones: pulumi.Output[list]
    """
    A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
    """
    cluster_mode: pulumi.Output[dict]
    """
    Create a native redis cluster. `automatic_failover_enabled` must be set to true. Cluster Mode documented below. Only 1 `cluster_mode` block is allowed.
    """
    configuration_endpoint_address: pulumi.Output[str]
    """
    The address of the replication group configuration endpoint when cluster mode is enabled.
    """
    engine: pulumi.Output[str]
    """
    The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
    """
    engine_version: pulumi.Output[str]
    """
    The version number of the cache engine to be used for the cache clusters in this replication group.
    """
    maintenance_window: pulumi.Output[str]
    """
    Specifies the weekly time range for when maintenance
    on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
    The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
    """
    member_clusters: pulumi.Output[list]
    """
    The identifiers of all the nodes that are part of this replication group.
    """
    node_type: pulumi.Output[str]
    """
    The compute and memory capacity of the nodes in the node group.
    """
    notification_topic_arn: pulumi.Output[str]
    """
    An Amazon Resource Name (ARN) of an
    SNS topic to send ElastiCache notifications to. Example:
    `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
    """
    number_cache_clusters: pulumi.Output[float]
    """
    The number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications.
    """
    parameter_group_name: pulumi.Output[str]
    """
    The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.
    """
    port: pulumi.Output[float]
    """
    The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
    """
    primary_endpoint_address: pulumi.Output[str]
    """
    (Redis only) The address of the endpoint for the primary node in the replication group, if the cluster mode is disabled.
    """
    replication_group_description: pulumi.Output[str]
    """
    A user-created description for the replication group.
    """
    replication_group_id: pulumi.Output[str]
    """
    The replication group identifier. This parameter is stored as a lowercase string.
    """
    security_group_ids: pulumi.Output[list]
    """
    One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
    """
    security_group_names: pulumi.Output[list]
    """
    A list of cache security group names to associate with this replication group.
    """
    snapshot_arns: pulumi.Output[list]
    """
    A single-element string list containing an
    Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
    Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
    """
    snapshot_name: pulumi.Output[str]
    """
    The name of a snapshot from which to restore data into the new node group. Changing the `snapshot_name` forces a new resource.
    """
    snapshot_retention_limit: pulumi.Output[float]
    """
    The number of days for which ElastiCache will
    retain automatic cache cluster snapshots before deleting them. For example, if you set
    SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
    before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
    Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
    """
    snapshot_window: pulumi.Output[str]
    """
    The daily time range (in UTC) during which ElastiCache will
    begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
    """
    subnet_group_name: pulumi.Output[str]
    """
    The name of the cache subnet group to be used for the replication group.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource
    """
    transit_encryption_enabled: pulumi.Output[bool]
    """
    Whether to enable encryption in transit.
    """
    def __init__(__self__, resource_name, opts=None, apply_immediately=None, at_rest_encryption_enabled=None, auth_token=None, auto_minor_version_upgrade=None, automatic_failover_enabled=None, availability_zones=None, cluster_mode=None, engine=None, engine_version=None, maintenance_window=None, node_type=None, notification_topic_arn=None, number_cache_clusters=None, parameter_group_name=None, port=None, replication_group_description=None, replication_group_id=None, security_group_ids=None, security_group_names=None, snapshot_arns=None, snapshot_name=None, snapshot_retention_limit=None, snapshot_window=None, subnet_group_name=None, tags=None, transit_encryption_enabled=None, __name__=None, __opts__=None):
        """
        Provides an ElastiCache Replication Group resource.
        For working with Memcached or single primary Redis instances (Cluster Mode Disabled), see the
        [`elasticache.Cluster` resource](https://www.terraform.io/docs/providers/aws/r/elasticache_cluster.html).
        
        > **Note:** When you change an attribute, such as `engine_version`, by
        default the ElastiCache API applies it in the next maintenance window. Because
        of this, this provider may report a difference in its planning phase because the
        actual modification has not yet taken place. You can use the
        `apply_immediately` flag to instruct the service to apply the change
        immediately. Using `apply_immediately` can result in a brief downtime as
        servers reboots.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
        :param pulumi.Input[bool] at_rest_encryption_enabled: Whether to enable encryption at rest.
        :param pulumi.Input[str] auth_token: The password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
        :param pulumi.Input[bool] automatic_failover_enabled: Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
        :param pulumi.Input[list] availability_zones: A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
        :param pulumi.Input[dict] cluster_mode: Create a native redis cluster. `automatic_failover_enabled` must be set to true. Cluster Mode documented below. Only 1 `cluster_mode` block is allowed.
        :param pulumi.Input[str] engine: The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
        :param pulumi.Input[str] engine_version: The version number of the cache engine to be used for the cache clusters in this replication group.
        :param pulumi.Input[str] maintenance_window: Specifies the weekly time range for when maintenance
               on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
               The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
        :param pulumi.Input[str] node_type: The compute and memory capacity of the nodes in the node group.
        :param pulumi.Input[str] notification_topic_arn: An Amazon Resource Name (ARN) of an
               SNS topic to send ElastiCache notifications to. Example:
               `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        :param pulumi.Input[float] number_cache_clusters: The number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications.
        :param pulumi.Input[str] parameter_group_name: The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.
        :param pulumi.Input[float] port: The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
        :param pulumi.Input[str] replication_group_description: A user-created description for the replication group.
        :param pulumi.Input[str] replication_group_id: The replication group identifier. This parameter is stored as a lowercase string.
        :param pulumi.Input[list] security_group_ids: One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
        :param pulumi.Input[list] security_group_names: A list of cache security group names to associate with this replication group.
        :param pulumi.Input[list] snapshot_arns: A single-element string list containing an
               Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
               Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
        :param pulumi.Input[str] snapshot_name: The name of a snapshot from which to restore data into the new node group. Changing the `snapshot_name` forces a new resource.
        :param pulumi.Input[float] snapshot_retention_limit: The number of days for which ElastiCache will
               retain automatic cache cluster snapshots before deleting them. For example, if you set
               SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
               before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
               Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
        :param pulumi.Input[str] snapshot_window: The daily time range (in UTC) during which ElastiCache will
               begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
        :param pulumi.Input[str] subnet_group_name: The name of the cache subnet group to be used for the replication group.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource
        :param pulumi.Input[bool] transit_encryption_enabled: Whether to enable encryption in transit.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticache_replication_group.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apply_immediately'] = apply_immediately
        __props__['at_rest_encryption_enabled'] = at_rest_encryption_enabled
        __props__['auth_token'] = auth_token
        __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade
        __props__['automatic_failover_enabled'] = automatic_failover_enabled
        __props__['availability_zones'] = availability_zones
        __props__['cluster_mode'] = cluster_mode
        __props__['engine'] = engine
        __props__['engine_version'] = engine_version
        __props__['maintenance_window'] = maintenance_window
        __props__['node_type'] = node_type
        __props__['notification_topic_arn'] = notification_topic_arn
        __props__['number_cache_clusters'] = number_cache_clusters
        __props__['parameter_group_name'] = parameter_group_name
        __props__['port'] = port
        if replication_group_description is None:
            raise TypeError("Missing required property 'replication_group_description'")
        __props__['replication_group_description'] = replication_group_description
        __props__['replication_group_id'] = replication_group_id
        __props__['security_group_ids'] = security_group_ids
        __props__['security_group_names'] = security_group_names
        __props__['snapshot_arns'] = snapshot_arns
        __props__['snapshot_name'] = snapshot_name
        __props__['snapshot_retention_limit'] = snapshot_retention_limit
        __props__['snapshot_window'] = snapshot_window
        __props__['subnet_group_name'] = subnet_group_name
        __props__['tags'] = tags
        __props__['transit_encryption_enabled'] = transit_encryption_enabled
        __props__['configuration_endpoint_address'] = None
        __props__['member_clusters'] = None
        __props__['primary_endpoint_address'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ReplicationGroup, __self__).__init__(
            'aws:elasticache/replicationGroup:ReplicationGroup',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

