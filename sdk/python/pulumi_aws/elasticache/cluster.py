# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class Cluster(pulumi.CustomResource):
    """
    Provides an ElastiCache Cluster resource, which manages a Memcached cluster or Redis instance.
    For working with Redis (Cluster Mode Enabled) replication groups, see the
    [`aws_elasticache_replication_group` resource](https://www.terraform.io/docs/providers/aws/r/elasticache_replication_group.html).
    
    ~> **Note:** When you change an attribute, such as `node_type`, by default
    it is applied in the next maintenance window. Because of this, Terraform may report
    a difference in its planning phase because the actual modification has not yet taken
    place. You can use the `apply_immediately` flag to instruct the service to apply the
    change immediately. Using `apply_immediately` can result in a brief downtime as the server reboots.
    See the AWS Docs on [Modifying an ElastiCache Cache Cluster][2] for more information.
    """
    def __init__(__self__, __name__, __opts__=None, apply_immediately=None, availability_zone=None, availability_zones=None, az_mode=None, cluster_id=None, engine=None, engine_version=None, maintenance_window=None, node_type=None, notification_topic_arn=None, num_cache_nodes=None, parameter_group_name=None, port=None, preferred_availability_zones=None, replication_group_id=None, security_group_ids=None, security_group_names=None, snapshot_arns=None, snapshot_name=None, snapshot_retention_limit=None, snapshot_window=None, subnet_group_name=None, tags=None):
        """Create a Cluster resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if apply_immediately and not isinstance(apply_immediately, bool):
            raise TypeError('Expected property apply_immediately to be a bool')
        __self__.apply_immediately = apply_immediately
        """
        Specifies whether any database modifications
        are applied immediately, or during the next maintenance window. Default is
        `false`. See [Amazon ElastiCache Documentation for more information.][1]
        (Available since v0.6.0)
        """
        __props__['applyImmediately'] = apply_immediately

        if availability_zone and not isinstance(availability_zone, basestring):
            raise TypeError('Expected property availability_zone to be a basestring')
        __self__.availability_zone = availability_zone
        """
        The Availability Zone for the cache cluster. If you want to create cache nodes in multi-az, use `preferred_availability_zones` instead. Default: System chosen Availability Zone.
        """
        __props__['availabilityZone'] = availability_zone

        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError('Expected property availability_zones to be a list')
        __self__.availability_zones = availability_zones
        """
        Use `preferred_availability_zones` instead unless you want to create cache nodes in single-az, then use `availability_zone`. Set of Availability Zones in which the cache nodes will be created.
        """
        __props__['availabilityZones'] = availability_zones

        if az_mode and not isinstance(az_mode, basestring):
            raise TypeError('Expected property az_mode to be a basestring')
        __self__.az_mode = az_mode
        """
        Specifies whether the nodes in this Memcached node group are created in a single Availability Zone or created across multiple Availability Zones in the cluster's region. Valid values for this parameter are `single-az` or `cross-az`, default is `single-az`. If you want to choose `cross-az`, `num_cache_nodes` must be greater than `1`
        """
        __props__['azMode'] = az_mode

        if not cluster_id:
            raise TypeError('Missing required property cluster_id')
        elif not isinstance(cluster_id, basestring):
            raise TypeError('Expected property cluster_id to be a basestring')
        __self__.cluster_id = cluster_id
        """
        Group identifier. ElastiCache converts
        this name to lowercase
        """
        __props__['clusterId'] = cluster_id

        if engine and not isinstance(engine, basestring):
            raise TypeError('Expected property engine to be a basestring')
        __self__.engine = engine
        """
        Name of the cache engine to be used for this cache cluster.
        Valid values for this parameter are `memcached` or `redis`
        """
        __props__['engine'] = engine

        if engine_version and not isinstance(engine_version, basestring):
            raise TypeError('Expected property engine_version to be a basestring')
        __self__.engine_version = engine_version
        """
        Version number of the cache engine to be used.
        See [Describe Cache Engine Versions](https://docs.aws.amazon.com/cli/latest/reference/elasticache/describe-cache-engine-versions.html)
        in the AWS Documentation center for supported versions
        """
        __props__['engineVersion'] = engine_version

        if maintenance_window and not isinstance(maintenance_window, basestring):
            raise TypeError('Expected property maintenance_window to be a basestring')
        __self__.maintenance_window = maintenance_window
        """
        Specifies the weekly time range for when maintenance
        on the cache cluster is performed. The format is `ddd:hh24:mi-ddd:hh24:mi` (24H Clock UTC).
        The minimum maintenance window is a 60 minute period. Example: `sun:05:00-sun:09:00`
        """
        __props__['maintenanceWindow'] = maintenance_window

        if node_type and not isinstance(node_type, basestring):
            raise TypeError('Expected property node_type to be a basestring')
        __self__.node_type = node_type
        """
        The compute and memory capacity of the nodes. See
        [Available Cache Node Types](https://aws.amazon.com/elasticache/details#Available_Cache_Node_Types) for
        supported node types
        """
        __props__['nodeType'] = node_type

        if notification_topic_arn and not isinstance(notification_topic_arn, basestring):
            raise TypeError('Expected property notification_topic_arn to be a basestring')
        __self__.notification_topic_arn = notification_topic_arn
        """
        An Amazon Resource Name (ARN) of an
        SNS topic to send ElastiCache notifications to. Example:
        `arn:aws:sns:us-east-1:012345678999:my_sns_topic`
        """
        __props__['notificationTopicArn'] = notification_topic_arn

        if num_cache_nodes and not isinstance(num_cache_nodes, int):
            raise TypeError('Expected property num_cache_nodes to be a int')
        __self__.num_cache_nodes = num_cache_nodes
        """
        The initial number of cache nodes that the
        cache cluster will have. For Redis, this value must be 1. For Memcache, this
        value must be between 1 and 20. If this number is reduced on subsequent runs,
        the highest numbered nodes will be removed.
        """
        __props__['numCacheNodes'] = num_cache_nodes

        if parameter_group_name and not isinstance(parameter_group_name, basestring):
            raise TypeError('Expected property parameter_group_name to be a basestring')
        __self__.parameter_group_name = parameter_group_name
        """
        Name of the parameter group to associate
        with this cache cluster
        """
        __props__['parameterGroupName'] = parameter_group_name

        if port and not isinstance(port, int):
            raise TypeError('Expected property port to be a int')
        __self__.port = port
        """
        The port number on which each of the cache nodes will accept connections. For Memcache the default is 11211, and for Redis the default port is 6379. Cannot be provided with `replication_group_id`.
        """
        __props__['port'] = port

        if preferred_availability_zones and not isinstance(preferred_availability_zones, list):
            raise TypeError('Expected property preferred_availability_zones to be a list')
        __self__.preferred_availability_zones = preferred_availability_zones
        """
        A list of the Availability Zones in which cache nodes are created. If you are creating your cluster in an Amazon VPC you can only locate nodes in Availability Zones that are associated with the subnets in the selected subnet group. The number of Availability Zones listed must equal the value of `num_cache_nodes`. If you want all the nodes in the same Availability Zone, use `availability_zone` instead, or repeat the Availability Zone multiple times in the list. Default: System chosen Availability Zones. Detecting drift of existing node availability zone is not currently supported. Updating this argument by itself to migrate existing node availability zones is not currently supported and will show a perpetual difference.
        """
        __props__['preferredAvailabilityZones'] = preferred_availability_zones

        if replication_group_id and not isinstance(replication_group_id, basestring):
            raise TypeError('Expected property replication_group_id to be a basestring')
        __self__.replication_group_id = replication_group_id
        """
        The ID of the replication group to which this cluster should belong. If this parameter is specified, the cluster is added to the specified replication group as a read replica; otherwise, the cluster is a standalone primary that is not part of any replication group.
        """
        __props__['replicationGroupId'] = replication_group_id

        if security_group_ids and not isinstance(security_group_ids, list):
            raise TypeError('Expected property security_group_ids to be a list')
        __self__.security_group_ids = security_group_ids
        """
        One or more VPC security groups associated
        with the cache cluster
        """
        __props__['securityGroupIds'] = security_group_ids

        if security_group_names and not isinstance(security_group_names, list):
            raise TypeError('Expected property security_group_names to be a list')
        __self__.security_group_names = security_group_names
        """
        List of security group
        names to associate with this cache cluster
        """
        __props__['securityGroupNames'] = security_group_names

        if snapshot_arns and not isinstance(snapshot_arns, list):
            raise TypeError('Expected property snapshot_arns to be a list')
        __self__.snapshot_arns = snapshot_arns
        """
        A single-element string list containing an
        Amazon Resource Name (ARN) of a Redis RDB snapshot file stored in Amazon S3.
        Example: `arn:aws:s3:::my_bucket/snapshot1.rdb`
        """
        __props__['snapshotArns'] = snapshot_arns

        if snapshot_name and not isinstance(snapshot_name, basestring):
            raise TypeError('Expected property snapshot_name to be a basestring')
        __self__.snapshot_name = snapshot_name
        """
        The name of a snapshot from which to restore data into the new node group.  Changing the `snapshot_name` forces a new resource.
        """
        __props__['snapshotName'] = snapshot_name

        if snapshot_retention_limit and not isinstance(snapshot_retention_limit, int):
            raise TypeError('Expected property snapshot_retention_limit to be a int')
        __self__.snapshot_retention_limit = snapshot_retention_limit
        """
        The number of days for which ElastiCache will
        retain automatic cache cluster snapshots before deleting them. For example, if you set
        SnapshotRetentionLimit to 5, then a snapshot that was taken today will be retained for 5 days
        before being deleted. If the value of SnapshotRetentionLimit is set to zero (0), backups are turned off.
        Please note that setting a `snapshot_retention_limit` is not supported on cache.t1.micro or cache.t2.* cache nodes
        """
        __props__['snapshotRetentionLimit'] = snapshot_retention_limit

        if snapshot_window and not isinstance(snapshot_window, basestring):
            raise TypeError('Expected property snapshot_window to be a basestring')
        __self__.snapshot_window = snapshot_window
        """
        The daily time range (in UTC) during which ElastiCache will
        begin taking a daily snapshot of your cache cluster. Example: 05:00-09:00
        """
        __props__['snapshotWindow'] = snapshot_window

        if subnet_group_name and not isinstance(subnet_group_name, basestring):
            raise TypeError('Expected property subnet_group_name to be a basestring')
        __self__.subnet_group_name = subnet_group_name
        """
        Name of the subnet group to be used
        for the cache cluster.
        """
        __props__['subnetGroupName'] = subnet_group_name

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource
        """
        __props__['tags'] = tags

        __self__.cache_nodes = pulumi.runtime.UNKNOWN
        """
        List of node objects including `id`, `address`, `port` and `availability_zone`.
        Referenceable e.g. as `${aws_elasticache_cluster.bar.cache_nodes.0.address}`
        """
        __self__.cluster_address = pulumi.runtime.UNKNOWN
        """
        (Memcached only) The DNS name of the cache cluster without the port appended.
        """
        __self__.configuration_endpoint = pulumi.runtime.UNKNOWN
        """
        (Memcached only) The configuration endpoint to allow host discovery.
        """

        super(Cluster, __self__).__init__(
            'aws:elasticache/cluster:Cluster',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'applyImmediately' in outs:
            self.apply_immediately = outs['applyImmediately']
        if 'availabilityZone' in outs:
            self.availability_zone = outs['availabilityZone']
        if 'availabilityZones' in outs:
            self.availability_zones = outs['availabilityZones']
        if 'azMode' in outs:
            self.az_mode = outs['azMode']
        if 'cacheNodes' in outs:
            self.cache_nodes = outs['cacheNodes']
        if 'clusterAddress' in outs:
            self.cluster_address = outs['clusterAddress']
        if 'clusterId' in outs:
            self.cluster_id = outs['clusterId']
        if 'configurationEndpoint' in outs:
            self.configuration_endpoint = outs['configurationEndpoint']
        if 'engine' in outs:
            self.engine = outs['engine']
        if 'engineVersion' in outs:
            self.engine_version = outs['engineVersion']
        if 'maintenanceWindow' in outs:
            self.maintenance_window = outs['maintenanceWindow']
        if 'nodeType' in outs:
            self.node_type = outs['nodeType']
        if 'notificationTopicArn' in outs:
            self.notification_topic_arn = outs['notificationTopicArn']
        if 'numCacheNodes' in outs:
            self.num_cache_nodes = outs['numCacheNodes']
        if 'parameterGroupName' in outs:
            self.parameter_group_name = outs['parameterGroupName']
        if 'port' in outs:
            self.port = outs['port']
        if 'preferredAvailabilityZones' in outs:
            self.preferred_availability_zones = outs['preferredAvailabilityZones']
        if 'replicationGroupId' in outs:
            self.replication_group_id = outs['replicationGroupId']
        if 'securityGroupIds' in outs:
            self.security_group_ids = outs['securityGroupIds']
        if 'securityGroupNames' in outs:
            self.security_group_names = outs['securityGroupNames']
        if 'snapshotArns' in outs:
            self.snapshot_arns = outs['snapshotArns']
        if 'snapshotName' in outs:
            self.snapshot_name = outs['snapshotName']
        if 'snapshotRetentionLimit' in outs:
            self.snapshot_retention_limit = outs['snapshotRetentionLimit']
        if 'snapshotWindow' in outs:
            self.snapshot_window = outs['snapshotWindow']
        if 'subnetGroupName' in outs:
            self.subnet_group_name = outs['subnetGroupName']
        if 'tags' in outs:
            self.tags = outs['tags']
