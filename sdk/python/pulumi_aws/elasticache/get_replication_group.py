# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetReplicationGroupResult(object):
    """
    A collection of values returned by getReplicationGroup.
    """
    def __init__(__self__, auth_token_enabled=None, automatic_failover_enabled=None, configuration_endpoint_address=None, member_clusters=None, node_type=None, number_cache_clusters=None, port=None, primary_endpoint_address=None, replication_group_description=None, snapshot_retention_limit=None, snapshot_window=None, id=None):
        if auth_token_enabled and not isinstance(auth_token_enabled, bool):
            raise TypeError('Expected argument auth_token_enabled to be a bool')
        __self__.auth_token_enabled = auth_token_enabled
        """
        A flag that enables using an AuthToken (password) when issuing Redis commands.
        """
        if automatic_failover_enabled and not isinstance(automatic_failover_enabled, bool):
            raise TypeError('Expected argument automatic_failover_enabled to be a bool')
        __self__.automatic_failover_enabled = automatic_failover_enabled
        """
        A flag whether a read-only replica will be automatically promoted to read/write primary if the existing primary fails.
        """
        if configuration_endpoint_address and not isinstance(configuration_endpoint_address, str):
            raise TypeError('Expected argument configuration_endpoint_address to be a str')
        __self__.configuration_endpoint_address = configuration_endpoint_address
        """
        The configuration endpoint address to allow host discovery.
        """
        if member_clusters and not isinstance(member_clusters, list):
            raise TypeError('Expected argument member_clusters to be a list')
        __self__.member_clusters = member_clusters
        """
        The identifiers of all the nodes that are part of this replication group.
        """
        if node_type and not isinstance(node_type, str):
            raise TypeError('Expected argument node_type to be a str')
        __self__.node_type = node_type
        """
        The cluster node type.
        """
        if number_cache_clusters and not isinstance(number_cache_clusters, int):
            raise TypeError('Expected argument number_cache_clusters to be a int')
        __self__.number_cache_clusters = number_cache_clusters
        """
        The number of cache clusters that the replication group has.
        """
        if port and not isinstance(port, int):
            raise TypeError('Expected argument port to be a int')
        __self__.port = port
        """
        The port number on which the configuration endpoint will accept connections.
        """
        if primary_endpoint_address and not isinstance(primary_endpoint_address, str):
            raise TypeError('Expected argument primary_endpoint_address to be a str')
        __self__.primary_endpoint_address = primary_endpoint_address
        """
        The endpoint of the primary node in this node group (shard).
        """
        if replication_group_description and not isinstance(replication_group_description, str):
            raise TypeError('Expected argument replication_group_description to be a str')
        __self__.replication_group_description = replication_group_description
        """
        The description of the replication group.
        """
        if snapshot_retention_limit and not isinstance(snapshot_retention_limit, int):
            raise TypeError('Expected argument snapshot_retention_limit to be a int')
        __self__.snapshot_retention_limit = snapshot_retention_limit
        """
        The number of days for which ElastiCache retains automatic cache cluster snapshots before deleting them.
        """
        if snapshot_window and not isinstance(snapshot_window, str):
            raise TypeError('Expected argument snapshot_window to be a str')
        __self__.snapshot_window = snapshot_window
        """
        The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_replication_group(replication_group_id=None):
    """
    Use this data source to get information about an Elasticache Replication Group.
    """
    __args__ = dict()

    __args__['replicationGroupId'] = replication_group_id
    __ret__ = await pulumi.runtime.invoke('aws:elasticache/getReplicationGroup:getReplicationGroup', __args__)

    return GetReplicationGroupResult(
        auth_token_enabled=__ret__.get('authTokenEnabled'),
        automatic_failover_enabled=__ret__.get('automaticFailoverEnabled'),
        configuration_endpoint_address=__ret__.get('configurationEndpointAddress'),
        member_clusters=__ret__.get('memberClusters'),
        node_type=__ret__.get('nodeType'),
        number_cache_clusters=__ret__.get('numberCacheClusters'),
        port=__ret__.get('port'),
        primary_endpoint_address=__ret__.get('primaryEndpointAddress'),
        replication_group_description=__ret__.get('replicationGroupDescription'),
        snapshot_retention_limit=__ret__.get('snapshotRetentionLimit'),
        snapshot_window=__ret__.get('snapshotWindow'),
        id=__ret__.get('id'))
