# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ClusterInstance']


class ClusterInstance(pulumi.CustomResource):
    apply_immediately: pulumi.Output[bool] = pulumi.property("applyImmediately")
    """
    Specifies whether any database modifications
    are applied immediately, or during the next maintenance window. Default is`false`.
    """

    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    Amazon Resource Name (ARN) of cluster instance
    """

    auto_minor_version_upgrade: pulumi.Output[Optional[bool]] = pulumi.property("autoMinorVersionUpgrade")
    """
    Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
    """

    availability_zone: pulumi.Output[str] = pulumi.property("availabilityZone")
    """
    The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
    """

    ca_cert_identifier: pulumi.Output[str] = pulumi.property("caCertIdentifier")
    """
    The identifier of the CA certificate for the DB instance.
    """

    cluster_identifier: pulumi.Output[str] = pulumi.property("clusterIdentifier")
    """
    The identifier of the `rds.Cluster` in which to launch this instance.
    """

    copy_tags_to_snapshot: pulumi.Output[Optional[bool]] = pulumi.property("copyTagsToSnapshot")
    """
    Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
    """

    db_parameter_group_name: pulumi.Output[str] = pulumi.property("dbParameterGroupName")
    """
    The name of the DB parameter group to associate with this instance.
    """

    db_subnet_group_name: pulumi.Output[str] = pulumi.property("dbSubnetGroupName")
    """
    A DB subnet group to associate with this DB instance. **NOTE:** This must match the `db_subnet_group_name` of the attached `rds.Cluster`.
    """

    dbi_resource_id: pulumi.Output[str] = pulumi.property("dbiResourceId")
    """
    The region-unique, immutable identifier for the DB instance.
    """

    endpoint: pulumi.Output[str] = pulumi.property("endpoint")
    """
    The DNS address for this instance. May not be writable
    """

    engine: pulumi.Output[Optional[str]] = pulumi.property("engine")
    """
    The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
    For information on the difference between the available Aurora MySQL engines
    see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
    in the Amazon RDS User Guide.
    """

    engine_version: pulumi.Output[str] = pulumi.property("engineVersion")
    """
    The database engine version.
    """

    identifier: pulumi.Output[str] = pulumi.property("identifier")
    """
    The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
    """

    identifier_prefix: pulumi.Output[str] = pulumi.property("identifierPrefix")
    """
    Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
    """

    instance_class: pulumi.Output[str] = pulumi.property("instanceClass")
    """
    The instance class to use. For details on CPU
    and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
    """

    kms_key_id: pulumi.Output[str] = pulumi.property("kmsKeyId")
    """
    The ARN for the KMS encryption key if one is set to the cluster.
    """

    monitoring_interval: pulumi.Output[Optional[float]] = pulumi.property("monitoringInterval")
    """
    The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
    """

    monitoring_role_arn: pulumi.Output[str] = pulumi.property("monitoringRoleArn")
    """
    The ARN for the IAM role that permits RDS to send
    enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
    what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
    """

    performance_insights_enabled: pulumi.Output[bool] = pulumi.property("performanceInsightsEnabled")
    """
    Specifies whether Performance Insights is enabled or not.
    """

    performance_insights_kms_key_id: pulumi.Output[str] = pulumi.property("performanceInsightsKmsKeyId")
    """
    The ARN for the KMS key to encrypt Performance Insights data. When specifying `performance_insights_kms_key_id`, `performance_insights_enabled` needs to be set to true.
    """

    port: pulumi.Output[float] = pulumi.property("port")
    """
    The database port
    """

    preferred_backup_window: pulumi.Output[str] = pulumi.property("preferredBackupWindow")
    """
    The daily time range during which automated backups are created if automated backups are enabled.
    Eg: "04:00-09:00"
    """

    preferred_maintenance_window: pulumi.Output[str] = pulumi.property("preferredMaintenanceWindow")
    """
    The window to perform maintenance in.
    Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
    """

    promotion_tier: pulumi.Output[Optional[float]] = pulumi.property("promotionTier")
    """
    Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
    """

    publicly_accessible: pulumi.Output[Optional[bool]] = pulumi.property("publiclyAccessible")
    """
    Bool to control if instance is publicly accessible.
    Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
    details on controlling this property.
    """

    storage_encrypted: pulumi.Output[bool] = pulumi.property("storageEncrypted")
    """
    Specifies whether the DB cluster is encrypted.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A map of tags to assign to the instance.
    """

    writer: pulumi.Output[bool] = pulumi.property("writer")
    """
    Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_immediately: Optional[pulumi.Input[bool]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 ca_cert_identifier: Optional[pulumi.Input[str]] = None,
                 cluster_identifier: Optional[pulumi.Input[str]] = None,
                 copy_tags_to_snapshot: Optional[pulumi.Input[bool]] = None,
                 db_parameter_group_name: Optional[pulumi.Input[str]] = None,
                 db_subnet_group_name: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 monitoring_interval: Optional[pulumi.Input[float]] = None,
                 monitoring_role_arn: Optional[pulumi.Input[str]] = None,
                 performance_insights_enabled: Optional[pulumi.Input[bool]] = None,
                 performance_insights_kms_key_id: Optional[pulumi.Input[str]] = None,
                 preferred_backup_window: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 promotion_tier: Optional[pulumi.Input[float]] = None,
                 publicly_accessible: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an RDS Cluster Instance Resource. A Cluster Instance Resource defines
        attributes that are specific to a single instance in a [RDS Cluster](https://www.terraform.io/docs/providers/aws/r/rds_cluster.html),
        specifically running Amazon Aurora.

        Unlike other RDS resources that support replication, with Amazon Aurora you do
        not designate a primary and subsequent replicas. Instead, you simply add RDS
        Instances and Aurora manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/resources.html#count)
        meta-parameter to make multiple instances and join them all to the same RDS
        Cluster, or you may specify different Cluster Instance resources with various
        `instance_class` sizes.

        For more information on Amazon Aurora, see [Aurora on Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Aurora.html) in the Amazon RDS User Guide.

        > **NOTE:** Deletion Protection from the RDS service can only be enabled at the cluster level, not for individual cluster instances. You can still add the [`protect` CustomResourceOption](https://www.pulumi.com/docs/intro/concepts/programming-model/#protect) to this resource configuration if you desire protection from accidental deletion.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.rds.Cluster("default",
            cluster_identifier="aurora-cluster-demo",
            availability_zones=[
                "us-west-2a",
                "us-west-2b",
                "us-west-2c",
            ],
            database_name="mydb",
            master_username="foo",
            master_password="barbut8chars")
        cluster_instances = []
        for range in [{"value": i} for i in range(0, 2)]:
            cluster_instances.append(aws.rds.ClusterInstance(f"clusterInstances-{range['value']}",
                identifier=f"aurora-cluster-demo-{range['value']}",
                cluster_identifier=default.id,
                instance_class="db.r4.large",
                engine=default.engine,
                engine_version=default.engine_version))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any database modifications
               are applied immediately, or during the next maintenance window. Default is`false`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        :param pulumi.Input[str] availability_zone: The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
        :param pulumi.Input[str] ca_cert_identifier: The identifier of the CA certificate for the DB instance.
        :param pulumi.Input[str] cluster_identifier: The identifier of the `rds.Cluster` in which to launch this instance.
        :param pulumi.Input[bool] copy_tags_to_snapshot: Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
        :param pulumi.Input[str] db_parameter_group_name: The name of the DB parameter group to associate with this instance.
        :param pulumi.Input[str] db_subnet_group_name: A DB subnet group to associate with this DB instance. **NOTE:** This must match the `db_subnet_group_name` of the attached `rds.Cluster`.
        :param pulumi.Input[str] engine: The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
               For information on the difference between the available Aurora MySQL engines
               see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
               in the Amazon RDS User Guide.
        :param pulumi.Input[str] engine_version: The database engine version.
        :param pulumi.Input[str] identifier: The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU
               and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
        :param pulumi.Input[float] monitoring_interval: The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
        :param pulumi.Input[str] monitoring_role_arn: The ARN for the IAM role that permits RDS to send
               enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
               what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
        :param pulumi.Input[bool] performance_insights_enabled: Specifies whether Performance Insights is enabled or not.
        :param pulumi.Input[str] performance_insights_kms_key_id: The ARN for the KMS key to encrypt Performance Insights data. When specifying `performance_insights_kms_key_id`, `performance_insights_enabled` needs to be set to true.
        :param pulumi.Input[str] preferred_backup_window: The daily time range during which automated backups are created if automated backups are enabled.
               Eg: "04:00-09:00"
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[float] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
        :param pulumi.Input[bool] publicly_accessible: Bool to control if instance is publicly accessible.
               Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
               details on controlling this property.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the instance.
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
            __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade
            __props__['availability_zone'] = availability_zone
            __props__['ca_cert_identifier'] = ca_cert_identifier
            if cluster_identifier is None:
                raise TypeError("Missing required property 'cluster_identifier'")
            __props__['cluster_identifier'] = cluster_identifier
            __props__['copy_tags_to_snapshot'] = copy_tags_to_snapshot
            __props__['db_parameter_group_name'] = db_parameter_group_name
            __props__['db_subnet_group_name'] = db_subnet_group_name
            __props__['engine'] = engine
            __props__['engine_version'] = engine_version
            __props__['identifier'] = identifier
            __props__['identifier_prefix'] = identifier_prefix
            if instance_class is None:
                raise TypeError("Missing required property 'instance_class'")
            __props__['instance_class'] = instance_class
            __props__['monitoring_interval'] = monitoring_interval
            __props__['monitoring_role_arn'] = monitoring_role_arn
            __props__['performance_insights_enabled'] = performance_insights_enabled
            __props__['performance_insights_kms_key_id'] = performance_insights_kms_key_id
            __props__['preferred_backup_window'] = preferred_backup_window
            __props__['preferred_maintenance_window'] = preferred_maintenance_window
            __props__['promotion_tier'] = promotion_tier
            __props__['publicly_accessible'] = publicly_accessible
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['dbi_resource_id'] = None
            __props__['endpoint'] = None
            __props__['kms_key_id'] = None
            __props__['port'] = None
            __props__['storage_encrypted'] = None
            __props__['writer'] = None
        super(ClusterInstance, __self__).__init__(
            'aws:rds/clusterInstance:ClusterInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            apply_immediately: Optional[pulumi.Input[bool]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            ca_cert_identifier: Optional[pulumi.Input[str]] = None,
            cluster_identifier: Optional[pulumi.Input[str]] = None,
            copy_tags_to_snapshot: Optional[pulumi.Input[bool]] = None,
            db_parameter_group_name: Optional[pulumi.Input[str]] = None,
            db_subnet_group_name: Optional[pulumi.Input[str]] = None,
            dbi_resource_id: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            identifier_prefix: Optional[pulumi.Input[str]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            monitoring_interval: Optional[pulumi.Input[float]] = None,
            monitoring_role_arn: Optional[pulumi.Input[str]] = None,
            performance_insights_enabled: Optional[pulumi.Input[bool]] = None,
            performance_insights_kms_key_id: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[float]] = None,
            preferred_backup_window: Optional[pulumi.Input[str]] = None,
            preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
            promotion_tier: Optional[pulumi.Input[float]] = None,
            publicly_accessible: Optional[pulumi.Input[bool]] = None,
            storage_encrypted: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            writer: Optional[pulumi.Input[bool]] = None) -> 'ClusterInstance':
        """
        Get an existing ClusterInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any database modifications
               are applied immediately, or during the next maintenance window. Default is`false`.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of cluster instance
        :param pulumi.Input[bool] auto_minor_version_upgrade: Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        :param pulumi.Input[str] availability_zone: The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) about the details.
        :param pulumi.Input[str] ca_cert_identifier: The identifier of the CA certificate for the DB instance.
        :param pulumi.Input[str] cluster_identifier: The identifier of the `rds.Cluster` in which to launch this instance.
        :param pulumi.Input[bool] copy_tags_to_snapshot: Indicates whether to copy all of the user-defined tags from the DB instance to snapshots of the DB instance. Default `false`.
        :param pulumi.Input[str] db_parameter_group_name: The name of the DB parameter group to associate with this instance.
        :param pulumi.Input[str] db_subnet_group_name: A DB subnet group to associate with this DB instance. **NOTE:** This must match the `db_subnet_group_name` of the attached `rds.Cluster`.
        :param pulumi.Input[str] dbi_resource_id: The region-unique, immutable identifier for the DB instance.
        :param pulumi.Input[str] endpoint: The DNS address for this instance. May not be writable
        :param pulumi.Input[str] engine: The name of the database engine to be used for the RDS instance. Defaults to `aurora`. Valid Values: `aurora`, `aurora-mysql`, `aurora-postgresql`.
               For information on the difference between the available Aurora MySQL engines
               see [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
               in the Amazon RDS User Guide.
        :param pulumi.Input[str] engine_version: The database engine version.
        :param pulumi.Input[str] identifier: The indentifier for the RDS instance, if omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU
               and memory, see [Scaling Aurora DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Aurora.Managing.html). Aurora uses `db.*` instance classes/types. Please see [AWS Documentation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) for currently available instance classes and complete details.
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key if one is set to the cluster.
        :param pulumi.Input[float] monitoring_interval: The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.
        :param pulumi.Input[str] monitoring_role_arn: The ARN for the IAM role that permits RDS to send
               enhanced monitoring metrics to CloudWatch Logs. You can find more information on the [AWS Documentation](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.html)
               what IAM permissions are needed to allow Enhanced Monitoring for RDS Instances.
        :param pulumi.Input[bool] performance_insights_enabled: Specifies whether Performance Insights is enabled or not.
        :param pulumi.Input[str] performance_insights_kms_key_id: The ARN for the KMS key to encrypt Performance Insights data. When specifying `performance_insights_kms_key_id`, `performance_insights_enabled` needs to be set to true.
        :param pulumi.Input[float] port: The database port
        :param pulumi.Input[str] preferred_backup_window: The daily time range during which automated backups are created if automated backups are enabled.
               Eg: "04:00-09:00"
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[float] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoted to writer.
        :param pulumi.Input[bool] publicly_accessible: Bool to control if instance is publicly accessible.
               Default `false`. See the documentation on [Creating DB Instances](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html) for more
               details on controlling this property.
        :param pulumi.Input[bool] storage_encrypted: Specifies whether the DB cluster is encrypted.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the instance.
        :param pulumi.Input[bool] writer: Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["apply_immediately"] = apply_immediately
        __props__["arn"] = arn
        __props__["auto_minor_version_upgrade"] = auto_minor_version_upgrade
        __props__["availability_zone"] = availability_zone
        __props__["ca_cert_identifier"] = ca_cert_identifier
        __props__["cluster_identifier"] = cluster_identifier
        __props__["copy_tags_to_snapshot"] = copy_tags_to_snapshot
        __props__["db_parameter_group_name"] = db_parameter_group_name
        __props__["db_subnet_group_name"] = db_subnet_group_name
        __props__["dbi_resource_id"] = dbi_resource_id
        __props__["endpoint"] = endpoint
        __props__["engine"] = engine
        __props__["engine_version"] = engine_version
        __props__["identifier"] = identifier
        __props__["identifier_prefix"] = identifier_prefix
        __props__["instance_class"] = instance_class
        __props__["kms_key_id"] = kms_key_id
        __props__["monitoring_interval"] = monitoring_interval
        __props__["monitoring_role_arn"] = monitoring_role_arn
        __props__["performance_insights_enabled"] = performance_insights_enabled
        __props__["performance_insights_kms_key_id"] = performance_insights_kms_key_id
        __props__["port"] = port
        __props__["preferred_backup_window"] = preferred_backup_window
        __props__["preferred_maintenance_window"] = preferred_maintenance_window
        __props__["promotion_tier"] = promotion_tier
        __props__["publicly_accessible"] = publicly_accessible
        __props__["storage_encrypted"] = storage_encrypted
        __props__["tags"] = tags
        __props__["writer"] = writer
        return ClusterInstance(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

