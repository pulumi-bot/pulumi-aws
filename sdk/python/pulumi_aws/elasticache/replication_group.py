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

__all__ = ['ReplicationGroup']


class ReplicationGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_immediately: Optional[pulumi.Input[bool]] = None,
                 at_rest_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 automatic_failover_enabled: Optional[pulumi.Input[bool]] = None,
                 availability_zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 cluster_mode: Optional[pulumi.Input[pulumi.InputType['ReplicationGroupClusterModeArgs']]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 maintenance_window: Optional[pulumi.Input[str]] = None,
                 node_type: Optional[pulumi.Input[str]] = None,
                 notification_topic_arn: Optional[pulumi.Input[str]] = None,
                 number_cache_clusters: Optional[pulumi.Input[float]] = None,
                 parameter_group_name: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 replication_group_description: Optional[pulumi.Input[str]] = None,
                 replication_group_id: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 security_group_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 snapshot_arns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 snapshot_name: Optional[pulumi.Input[str]] = None,
                 snapshot_retention_limit: Optional[pulumi.Input[float]] = None,
                 snapshot_window: Optional[pulumi.Input[str]] = None,
                 subnet_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an ElastiCache Replication Group resource.
        For working with Memcached or single primary Redis instances (Cluster Mode Disabled), see the
        `elasticache.Cluster` resource.

        > **Note:** When you change an attribute, such as `engine_version`, by
        default the ElastiCache API applies it in the next maintenance window. Because
        of this, this provider may report a difference in its planning phase because the
        actual modification has not yet taken place. You can use the
        `apply_immediately` flag to instruct the service to apply the change
        immediately. Using `apply_immediately` can result in a brief downtime as
        servers reboots.

        ## Example Usage
        ### Redis Cluster Mode Disabled

        To create a single shard primary with single read replica:

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.elasticache.ReplicationGroup("example",
            automatic_failover_enabled=True,
            availability_zones=[
                "us-west-2a",
                "us-west-2b",
            ],
            node_type="cache.m4.large",
            number_cache_clusters=2,
            parameter_group_name="default.redis3.2",
            port=6379,
            replication_group_description="test description")
        ```

        You have two options for adjusting the number of replicas:

        * Adjusting `number_cache_clusters` directly. This will attempt to automatically add or remove replicas, but provides no granular control (e.g. preferred availability zone, cache cluster ID) for the added or removed replicas. This also currently expects cache cluster IDs in the form of `replication_group_id-00#`.
        * Otherwise for fine grained control of the underlying cache clusters, they can be added or removed with the `elasticache.Cluster` resource and its `replication_group_id` attribute. In this situation, you will need to utilize [`ignoreChanges`](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to prevent perpetual differences with the `number_cache_cluster` attribute.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.elasticache.ReplicationGroup("example",
            automatic_failover_enabled=True,
            availability_zones=[
                "us-west-2a",
                "us-west-2b",
            ],
            replication_group_description="test description",
            node_type="cache.m4.large",
            number_cache_clusters=2,
            parameter_group_name="default.redis3.2",
            port=6379)
        replica = None
        if 1 == True:
            replica = aws.elasticache.Cluster("replica", replication_group_id=example.id)
        ```
        ### Redis Cluster Mode Enabled

        To create two shards with a primary and a single read replica each:

        ```python
        import pulumi
        import pulumi_aws as aws

        baz = aws.elasticache.ReplicationGroup("baz",
            automatic_failover_enabled=True,
            cluster_mode={
                "numNodeGroups": 2,
                "replicasPerNodeGroup": 1,
            },
            node_type="cache.t2.small",
            parameter_group_name="default.redis3.2.cluster.on",
            port=6379,
            replication_group_description="test description")
        ```

        > **Note:** We currently do not support passing a `primary_cluster_id` in order to create the Replication Group.

        > **Note:** Automatic Failover is unavailable for Redis versions earlier than 2.8.6,
        and unavailable on T1 node types. For T2 node types, it is only available on Redis version 3.2.4 or later with cluster mode enabled. See the [High Availability Using Replication Groups](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Replication.html) guide
        for full details on using Replication Groups.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
        :param pulumi.Input[bool] at_rest_encryption_enabled: Whether to enable encryption at rest.
        :param pulumi.Input[str] auth_token: The password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
        :param pulumi.Input[bool] automatic_failover_enabled: Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] availability_zones: A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
        :param pulumi.Input[pulumi.InputType['ReplicationGroupClusterModeArgs']] cluster_mode: Create a native redis cluster. `automatic_failover_enabled` must be set to true. Cluster Mode documented below. Only 1 `cluster_mode` block is allowed.
        :param pulumi.Input[str] engine: The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
        :param pulumi.Input[str] engine_version: The version number of the cache engine to be used for the cache clusters in this replication group.
        :param pulumi.Input[str] kms_key_id: The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `at_rest_encryption_enabled = true`.
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
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_ids: One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_names: A list of cache security group names to associate with this replication group.
        :param pulumi.Input[List[pulumi.Input[str]]] snapshot_arns: A single-element string list containing an
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
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource
        :param pulumi.Input[bool] transit_encryption_enabled: Whether to enable encryption in transit.
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

            __props__['apply_immediately'] = apply_immediately
            __props__['at_rest_encryption_enabled'] = at_rest_encryption_enabled
            __props__['auth_token'] = auth_token
            __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade
            __props__['automatic_failover_enabled'] = automatic_failover_enabled
            __props__['availability_zones'] = availability_zones
            __props__['cluster_mode'] = cluster_mode
            __props__['engine'] = engine
            __props__['engine_version'] = engine_version
            __props__['kms_key_id'] = kms_key_id
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
        super(ReplicationGroup, __self__).__init__(
            'aws:elasticache/replicationGroup:ReplicationGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apply_immediately: Optional[pulumi.Input[bool]] = None,
            at_rest_encryption_enabled: Optional[pulumi.Input[bool]] = None,
            auth_token: Optional[pulumi.Input[str]] = None,
            auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
            automatic_failover_enabled: Optional[pulumi.Input[bool]] = None,
            availability_zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            cluster_mode: Optional[pulumi.Input[pulumi.InputType['ReplicationGroupClusterModeArgs']]] = None,
            configuration_endpoint_address: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            maintenance_window: Optional[pulumi.Input[str]] = None,
            member_clusters: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            node_type: Optional[pulumi.Input[str]] = None,
            notification_topic_arn: Optional[pulumi.Input[str]] = None,
            number_cache_clusters: Optional[pulumi.Input[float]] = None,
            parameter_group_name: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[float]] = None,
            primary_endpoint_address: Optional[pulumi.Input[str]] = None,
            replication_group_description: Optional[pulumi.Input[str]] = None,
            replication_group_id: Optional[pulumi.Input[str]] = None,
            security_group_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            security_group_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            snapshot_arns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            snapshot_name: Optional[pulumi.Input[str]] = None,
            snapshot_retention_limit: Optional[pulumi.Input[float]] = None,
            snapshot_window: Optional[pulumi.Input[str]] = None,
            subnet_group_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            transit_encryption_enabled: Optional[pulumi.Input[bool]] = None) -> 'ReplicationGroup':
        """
        Get an existing ReplicationGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
        :param pulumi.Input[bool] at_rest_encryption_enabled: Whether to enable encryption at rest.
        :param pulumi.Input[str] auth_token: The password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
        :param pulumi.Input[bool] automatic_failover_enabled: Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
        :param pulumi.Input[List[pulumi.Input[str]]] availability_zones: A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
        :param pulumi.Input[pulumi.InputType['ReplicationGroupClusterModeArgs']] cluster_mode: Create a native redis cluster. `automatic_failover_enabled` must be set to true. Cluster Mode documented below. Only 1 `cluster_mode` block is allowed.
        :param pulumi.Input[str] configuration_endpoint_address: The address of the replication group configuration endpoint when cluster mode is enabled.
        :param pulumi.Input[str] engine: The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
        :param pulumi.Input[str] engine_version: The version number of the cache engine to be used for the cache clusters in this replication group.
        :param pulumi.Input[str] kms_key_id: The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `at_rest_encryption_enabled = true`.
        :param pulumi.Input[str] maintenance_window: Specifies the weekly time range for when maintenance
               on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
               The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
        :param pulumi.Input[List[pulumi.Input[str]]] member_clusters: The identifiers of all the nodes that are part of this replication group.
        :param pulumi.Input[str] node_type: The compute and memory capacity of the nodes in the node group.
        :param pulumi.Input[str] notification_topic_arn: An Amazon Resource Name (ARN) of an
               SNS topic to send ElastiCache notifications to. Example:
               `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        :param pulumi.Input[float] number_cache_clusters: The number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications.
        :param pulumi.Input[str] parameter_group_name: The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.
        :param pulumi.Input[float] port: The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
        :param pulumi.Input[str] primary_endpoint_address: (Redis only) The address of the endpoint for the primary node in the replication group, if the cluster mode is disabled.
        :param pulumi.Input[str] replication_group_description: A user-created description for the replication group.
        :param pulumi.Input[str] replication_group_id: The replication group identifier. This parameter is stored as a lowercase string.
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_ids: One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_names: A list of cache security group names to associate with this replication group.
        :param pulumi.Input[List[pulumi.Input[str]]] snapshot_arns: A single-element string list containing an
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
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource
        :param pulumi.Input[bool] transit_encryption_enabled: Whether to enable encryption in transit.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["apply_immediately"] = apply_immediately
        __props__["at_rest_encryption_enabled"] = at_rest_encryption_enabled
        __props__["auth_token"] = auth_token
        __props__["auto_minor_version_upgrade"] = auto_minor_version_upgrade
        __props__["automatic_failover_enabled"] = automatic_failover_enabled
        __props__["availability_zones"] = availability_zones
        __props__["cluster_mode"] = cluster_mode
        __props__["configuration_endpoint_address"] = configuration_endpoint_address
        __props__["engine"] = engine
        __props__["engine_version"] = engine_version
        __props__["kms_key_id"] = kms_key_id
        __props__["maintenance_window"] = maintenance_window
        __props__["member_clusters"] = member_clusters
        __props__["node_type"] = node_type
        __props__["notification_topic_arn"] = notification_topic_arn
        __props__["number_cache_clusters"] = number_cache_clusters
        __props__["parameter_group_name"] = parameter_group_name
        __props__["port"] = port
        __props__["primary_endpoint_address"] = primary_endpoint_address
        __props__["replication_group_description"] = replication_group_description
        __props__["replication_group_id"] = replication_group_id
        __props__["security_group_ids"] = security_group_ids
        __props__["security_group_names"] = security_group_names
        __props__["snapshot_arns"] = snapshot_arns
        __props__["snapshot_name"] = snapshot_name
        __props__["snapshot_retention_limit"] = snapshot_retention_limit
        __props__["snapshot_window"] = snapshot_window
        __props__["subnet_group_name"] = subnet_group_name
        __props__["tags"] = tags
        __props__["transit_encryption_enabled"] = transit_encryption_enabled
        return ReplicationGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applyImmediately")
    def apply_immediately(self) -> bool:
        """
        Specifies whether any modifications are applied immediately, or during the next maintenance window. Default is `false`.
        """
        return pulumi.get(self, "apply_immediately")

    @property
    @pulumi.getter(name="atRestEncryptionEnabled")
    def at_rest_encryption_enabled(self) -> Optional[bool]:
        """
        Whether to enable encryption at rest.
        """
        return pulumi.get(self, "at_rest_encryption_enabled")

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> Optional[str]:
        """
        The password used to access a password protected server. Can be specified only if `transit_encryption_enabled = true`.
        """
        return pulumi.get(self, "auth_token")

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> Optional[bool]:
        """
        Specifies whether a minor engine upgrades will be applied automatically to the underlying Cache Cluster instances during the maintenance window. Defaults to `true`.
        """
        return pulumi.get(self, "auto_minor_version_upgrade")

    @property
    @pulumi.getter(name="automaticFailoverEnabled")
    def automatic_failover_enabled(self) -> Optional[bool]:
        """
        Specifies whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails. If true, Multi-AZ is enabled for this replication group. If false, Multi-AZ is disabled for this replication group. Must be enabled for Redis (cluster mode enabled) replication groups. Defaults to `false`.
        """
        return pulumi.get(self, "automatic_failover_enabled")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[List[str]]:
        """
        A list of EC2 availability zones in which the replication group's cache clusters will be created. The order of the availability zones in the list is not important.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter(name="clusterMode")
    def cluster_mode(self) -> 'outputs.ReplicationGroupClusterMode':
        """
        Create a native redis cluster. `automatic_failover_enabled` must be set to true. Cluster Mode documented below. Only 1 `cluster_mode` block is allowed.
        """
        return pulumi.get(self, "cluster_mode")

    @property
    @pulumi.getter(name="configurationEndpointAddress")
    def configuration_endpoint_address(self) -> str:
        """
        The address of the replication group configuration endpoint when cluster mode is enabled.
        """
        return pulumi.get(self, "configuration_endpoint_address")

    @property
    @pulumi.getter
    def engine(self) -> Optional[str]:
        """
        The name of the cache engine to be used for the clusters in this replication group. e.g. `redis`
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        """
        The version number of the cache engine to be used for the cache clusters in this replication group.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        """
        The ARN of the key that you wish to use if encrypting at rest. If not supplied, uses service managed encryption. Can be specified only if `at_rest_encryption_enabled = true`.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter(name="maintenanceWindow")
    def maintenance_window(self) -> str:
        """
        Specifies the weekly time range for when maintenance
        on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
        The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
        """
        return pulumi.get(self, "maintenance_window")

    @property
    @pulumi.getter(name="memberClusters")
    def member_clusters(self) -> List[str]:
        """
        The identifiers of all the nodes that are part of this replication group.
        """
        return pulumi.get(self, "member_clusters")

    @property
    @pulumi.getter(name="nodeType")
    def node_type(self) -> str:
        """
        The compute and memory capacity of the nodes in the node group.
        """
        return pulumi.get(self, "node_type")

    @property
    @pulumi.getter(name="notificationTopicArn")
    def notification_topic_arn(self) -> Optional[str]:
        """
        An Amazon Resource Name (ARN) of an
        SNS topic to send ElastiCache notifications to. Example:
        `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        """
        return pulumi.get(self, "notification_topic_arn")

    @property
    @pulumi.getter(name="numberCacheClusters")
    def number_cache_clusters(self) -> float:
        """
        The number of cache clusters (primary and replicas) this replication group will have. If Multi-AZ is enabled, the value of this parameter must be at least 2. Updates will occur before other modifications.
        """
        return pulumi.get(self, "number_cache_clusters")

    @property
    @pulumi.getter(name="parameterGroupName")
    def parameter_group_name(self) -> str:
        """
        The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.
        """
        return pulumi.get(self, "parameter_group_name")

    @property
    @pulumi.getter
    def port(self) -> Optional[float]:
        """
        The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="primaryEndpointAddress")
    def primary_endpoint_address(self) -> str:
        """
        (Redis only) The address of the endpoint for the primary node in the replication group, if the cluster mode is disabled.
        """
        return pulumi.get(self, "primary_endpoint_address")

    @property
    @pulumi.getter(name="replicationGroupDescription")
    def replication_group_description(self) -> str:
        """
        A user-created description for the replication group.
        """
        return pulumi.get(self, "replication_group_description")

    @property
    @pulumi.getter(name="replicationGroupId")
    def replication_group_id(self) -> str:
        """
        The replication group identifier. This parameter is stored as a lowercase string.
        """
        return pulumi.get(self, "replication_group_id")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> List[str]:
        """
        One or more Amazon VPC security groups associated with this replication group. Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="securityGroupNames")
    def security_group_names(self) -> List[str]:
        """
        A list of cache security group names to associate with this replication group.
        """
        return pulumi.get(self, "security_group_names")

    @property
    @pulumi.getter(name="snapshotArns")
    def snapshot_arns(self) -> Optional[List[str]]:
        """
        A single-element string list containing an
        Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
        Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
        """
        return pulumi.get(self, "snapshot_arns")

    @property
    @pulumi.getter(name="snapshotName")
    def snapshot_name(self) -> Optional[str]:
        """
        The name of a snapshot from which to restore data into the new node group. Changing the `snapshot_name` forces a new resource.
        """
        return pulumi.get(self, "snapshot_name")

    @property
    @pulumi.getter(name="snapshotRetentionLimit")
    def snapshot_retention_limit(self) -> Optional[float]:
        """
        The number of days for which ElastiCache will
        retain automatic cache cluster snapshots before deleting them. For example, if you set
        SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
        before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
        Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
        """
        return pulumi.get(self, "snapshot_retention_limit")

    @property
    @pulumi.getter(name="snapshotWindow")
    def snapshot_window(self) -> str:
        """
        The daily time range (in UTC) during which ElastiCache will
        begin taking a daily snapshot of your cache cluster. The minimum snapshot window is a 60 minute period. Example: `05:00-09:00`
        """
        return pulumi.get(self, "snapshot_window")

    @property
    @pulumi.getter(name="subnetGroupName")
    def subnet_group_name(self) -> str:
        """
        The name of the cache subnet group to be used for the replication group.
        """
        return pulumi.get(self, "subnet_group_name")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitEncryptionEnabled")
    def transit_encryption_enabled(self) -> Optional[bool]:
        """
        Whether to enable encryption in transit.
        """
        return pulumi.get(self, "transit_encryption_enabled")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

