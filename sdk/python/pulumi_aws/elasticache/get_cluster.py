# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, arn=None, availability_zone=None, cache_nodes=None, cluster_address=None, cluster_id=None, configuration_endpoint=None, engine=None, engine_version=None, id=None, maintenance_window=None, node_type=None, notification_topic_arn=None, num_cache_nodes=None, parameter_group_name=None, port=None, replication_group_id=None, security_group_ids=None, security_group_names=None, snapshot_retention_limit=None, snapshot_window=None, subnet_group_name=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        __self__.availability_zone = availability_zone
        """
        The Availability Zone for the cache cluster.
        """
        if cache_nodes and not isinstance(cache_nodes, list):
            raise TypeError("Expected argument 'cache_nodes' to be a list")
        __self__.cache_nodes = cache_nodes
        """
        List of node objects including `id`, `address`, `port` and `availability_zone`.
        Referenceable e.g. as `${data.aws_elasticache_cluster.bar.cache_nodes.0.address}`
        """
        if cluster_address and not isinstance(cluster_address, str):
            raise TypeError("Expected argument 'cluster_address' to be a str")
        __self__.cluster_address = cluster_address
        """
        (Memcached only) The DNS name of the cache cluster without the port appended.
        """
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        __self__.cluster_id = cluster_id
        if configuration_endpoint and not isinstance(configuration_endpoint, str):
            raise TypeError("Expected argument 'configuration_endpoint' to be a str")
        __self__.configuration_endpoint = configuration_endpoint
        """
        (Memcached only) The configuration endpoint to allow host discovery.
        """
        if engine and not isinstance(engine, str):
            raise TypeError("Expected argument 'engine' to be a str")
        __self__.engine = engine
        """
        Name of the cache engine.
        """
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        __self__.engine_version = engine_version
        """
        Version number of the cache engine.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if maintenance_window and not isinstance(maintenance_window, str):
            raise TypeError("Expected argument 'maintenance_window' to be a str")
        __self__.maintenance_window = maintenance_window
        """
        Specifies the weekly time range for when maintenance
        on the cache cluster is performed.
        """
        if node_type and not isinstance(node_type, str):
            raise TypeError("Expected argument 'node_type' to be a str")
        __self__.node_type = node_type
        """
        The cluster node type.
        """
        if notification_topic_arn and not isinstance(notification_topic_arn, str):
            raise TypeError("Expected argument 'notification_topic_arn' to be a str")
        __self__.notification_topic_arn = notification_topic_arn
        """
        An Amazon Resource Name (ARN) of an
        SNS topic that ElastiCache notifications get sent to.
        """
        if num_cache_nodes and not isinstance(num_cache_nodes, float):
            raise TypeError("Expected argument 'num_cache_nodes' to be a float")
        __self__.num_cache_nodes = num_cache_nodes
        """
        The number of cache nodes that the cache cluster has.
        """
        if parameter_group_name and not isinstance(parameter_group_name, str):
            raise TypeError("Expected argument 'parameter_group_name' to be a str")
        __self__.parameter_group_name = parameter_group_name
        """
        Name of the parameter group associated with this cache cluster.
        """
        if port and not isinstance(port, float):
            raise TypeError("Expected argument 'port' to be a float")
        __self__.port = port
        """
        The port number on which each of the cache nodes will
        accept connections.
        """
        if replication_group_id and not isinstance(replication_group_id, str):
            raise TypeError("Expected argument 'replication_group_id' to be a str")
        __self__.replication_group_id = replication_group_id
        """
        The replication group to which this cache cluster belongs.
        """
        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError("Expected argument 'security_group_ids' to be a list")
        __self__.security_group_ids = security_group_ids
        """
        List VPC security groups associated with the cache cluster.
        """
        if security_group_names and not isinstance(security_group_names, list):
            raise TypeError("Expected argument 'security_group_names' to be a list")
        __self__.security_group_names = security_group_names
        """
        List of security group names associated with this cache cluster.
        """
        if snapshot_retention_limit and not isinstance(snapshot_retention_limit, float):
            raise TypeError("Expected argument 'snapshot_retention_limit' to be a float")
        __self__.snapshot_retention_limit = snapshot_retention_limit
        """
        The number of days for which ElastiCache will
        retain automatic cache cluster snapshots before deleting them.
        """
        if snapshot_window and not isinstance(snapshot_window, str):
            raise TypeError("Expected argument 'snapshot_window' to be a str")
        __self__.snapshot_window = snapshot_window
        """
        The daily time range (in UTC) during which ElastiCache will
        begin taking a daily snapshot of the cache cluster.
        """
        if subnet_group_name and not isinstance(subnet_group_name, str):
            raise TypeError("Expected argument 'subnet_group_name' to be a str")
        __self__.subnet_group_name = subnet_group_name
        """
        Name of the subnet group associated to the cache cluster.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        The tags assigned to the resource
        """
class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            arn=self.arn,
            availability_zone=self.availability_zone,
            cache_nodes=self.cache_nodes,
            cluster_address=self.cluster_address,
            cluster_id=self.cluster_id,
            configuration_endpoint=self.configuration_endpoint,
            engine=self.engine,
            engine_version=self.engine_version,
            id=self.id,
            maintenance_window=self.maintenance_window,
            node_type=self.node_type,
            notification_topic_arn=self.notification_topic_arn,
            num_cache_nodes=self.num_cache_nodes,
            parameter_group_name=self.parameter_group_name,
            port=self.port,
            replication_group_id=self.replication_group_id,
            security_group_ids=self.security_group_ids,
            security_group_names=self.security_group_names,
            snapshot_retention_limit=self.snapshot_retention_limit,
            snapshot_window=self.snapshot_window,
            subnet_group_name=self.subnet_group_name,
            tags=self.tags)

def get_cluster(cluster_id=None,tags=None,opts=None):
    """
    Use this data source to get information about an Elasticache Cluster

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elasticache_cluster.html.markdown.


    :param str cluster_id: Group identifier.
    """
    __args__ = dict()


    __args__['clusterId'] = cluster_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticache/getCluster:getCluster', __args__, opts=opts).value

    return AwaitableGetClusterResult(
        arn=__ret__.get('arn'),
        availability_zone=__ret__.get('availabilityZone'),
        cache_nodes=__ret__.get('cacheNodes'),
        cluster_address=__ret__.get('clusterAddress'),
        cluster_id=__ret__.get('clusterId'),
        configuration_endpoint=__ret__.get('configurationEndpoint'),
        engine=__ret__.get('engine'),
        engine_version=__ret__.get('engineVersion'),
        id=__ret__.get('id'),
        maintenance_window=__ret__.get('maintenanceWindow'),
        node_type=__ret__.get('nodeType'),
        notification_topic_arn=__ret__.get('notificationTopicArn'),
        num_cache_nodes=__ret__.get('numCacheNodes'),
        parameter_group_name=__ret__.get('parameterGroupName'),
        port=__ret__.get('port'),
        replication_group_id=__ret__.get('replicationGroupId'),
        security_group_ids=__ret__.get('securityGroupIds'),
        security_group_names=__ret__.get('securityGroupNames'),
        snapshot_retention_limit=__ret__.get('snapshotRetentionLimit'),
        snapshot_window=__ret__.get('snapshotWindow'),
        subnet_group_name=__ret__.get('subnetGroupName'),
        tags=__ret__.get('tags'))
