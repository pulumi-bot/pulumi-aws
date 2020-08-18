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
    The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
    """

    ca_cert_identifier: pulumi.Output[str] = pulumi.property("caCertIdentifier")
    """
    (Optional) The identifier of the CA certificate for the DB instance.
    """

    cluster_identifier: pulumi.Output[str] = pulumi.property("clusterIdentifier")
    """
    The identifier of the `docdb.Cluster` in which to launch this instance.
    """

    db_subnet_group_name: pulumi.Output[str] = pulumi.property("dbSubnetGroupName")
    """
    The DB subnet group to associate with this DB instance.
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
    The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
    """

    engine_version: pulumi.Output[str] = pulumi.property("engineVersion")
    """
    The database engine version
    """

    identifier: pulumi.Output[str] = pulumi.property("identifier")
    """
    The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
    """

    identifier_prefix: pulumi.Output[str] = pulumi.property("identifierPrefix")
    """
    Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
    """

    instance_class: pulumi.Output[str] = pulumi.property("instanceClass")
    """
    The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
    supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
    - db.r4.large
    - db.r4.xlarge
    - db.r4.2xlarge
    - db.r4.4xlarge
    - db.r4.8xlarge
    - db.r4.16xlarge
    """

    kms_key_id: pulumi.Output[str] = pulumi.property("kmsKeyId")
    """
    The ARN for the KMS encryption key if one is set to the cluster.
    """

    port: pulumi.Output[float] = pulumi.property("port")
    """
    The database port
    """

    preferred_backup_window: pulumi.Output[str] = pulumi.property("preferredBackupWindow")
    """
    The daily time range during which automated backups are created if automated backups are enabled.
    """

    preferred_maintenance_window: pulumi.Output[str] = pulumi.property("preferredMaintenanceWindow")
    """
    The window to perform maintenance in.
    Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
    """

    promotion_tier: pulumi.Output[Optional[float]] = pulumi.property("promotionTier")
    """
    Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
    """

    publicly_accessible: pulumi.Output[bool] = pulumi.property("publiclyAccessible")

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
                 engine: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 instance_class: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 promotion_tier: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
        attributes that are specific to a single instance in a [DocDB Cluster](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html).

        You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
        Instances and DocDB manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/resources.html#count)
        meta-parameter to make multiple instances and join them all to the same DocDB
        Cluster, or you may specify different Cluster Instance resources with various
        `instance_class` sizes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.docdb.Cluster("default",
            cluster_identifier="docdb-cluster-demo",
            availability_zones=[
                "us-west-2a",
                "us-west-2b",
                "us-west-2c",
            ],
            master_username="foo",
            master_password="barbut8chars")
        cluster_instances = []
        for range in [{"value": i} for i in range(0, 2)]:
            cluster_instances.append(aws.docdb.ClusterInstance(f"clusterInstances-{range['value']}",
                identifier=f"docdb-cluster-demo-{range['value']}",
                cluster_identifier=default.id,
                instance_class="db.r5.large"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any database modifications
               are applied immediately, or during the next maintenance window. Default is`false`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        :param pulumi.Input[str] availability_zone: The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        :param pulumi.Input[str] ca_cert_identifier: (Optional) The identifier of the CA certificate for the DB instance.
        :param pulumi.Input[str] cluster_identifier: The identifier of the `docdb.Cluster` in which to launch this instance.
        :param pulumi.Input[str] engine: The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        :param pulumi.Input[str] identifier: The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
               supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
               - db.r4.large
               - db.r4.xlarge
               - db.r4.2xlarge
               - db.r4.4xlarge
               - db.r4.8xlarge
               - db.r4.16xlarge
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[float] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
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
            __props__['engine'] = engine
            __props__['identifier'] = identifier
            __props__['identifier_prefix'] = identifier_prefix
            if instance_class is None:
                raise TypeError("Missing required property 'instance_class'")
            __props__['instance_class'] = instance_class
            __props__['preferred_maintenance_window'] = preferred_maintenance_window
            __props__['promotion_tier'] = promotion_tier
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['db_subnet_group_name'] = None
            __props__['dbi_resource_id'] = None
            __props__['endpoint'] = None
            __props__['engine_version'] = None
            __props__['kms_key_id'] = None
            __props__['port'] = None
            __props__['preferred_backup_window'] = None
            __props__['publicly_accessible'] = None
            __props__['storage_encrypted'] = None
            __props__['writer'] = None
        super(ClusterInstance, __self__).__init__(
            'aws:docdb/clusterInstance:ClusterInstance',
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
            db_subnet_group_name: Optional[pulumi.Input[str]] = None,
            dbi_resource_id: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            identifier: Optional[pulumi.Input[str]] = None,
            identifier_prefix: Optional[pulumi.Input[str]] = None,
            instance_class: Optional[pulumi.Input[str]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
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
        :param pulumi.Input[str] availability_zone: The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        :param pulumi.Input[str] ca_cert_identifier: (Optional) The identifier of the CA certificate for the DB instance.
        :param pulumi.Input[str] cluster_identifier: The identifier of the `docdb.Cluster` in which to launch this instance.
        :param pulumi.Input[str] db_subnet_group_name: The DB subnet group to associate with this DB instance.
        :param pulumi.Input[str] dbi_resource_id: The region-unique, immutable identifier for the DB instance.
        :param pulumi.Input[str] endpoint: The DNS address for this instance. May not be writable
        :param pulumi.Input[str] engine: The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        :param pulumi.Input[str] engine_version: The database engine version
        :param pulumi.Input[str] identifier: The indentifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
               supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
               - db.r4.large
               - db.r4.xlarge
               - db.r4.2xlarge
               - db.r4.4xlarge
               - db.r4.8xlarge
               - db.r4.16xlarge
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key if one is set to the cluster.
        :param pulumi.Input[float] port: The database port
        :param pulumi.Input[str] preferred_backup_window: The daily time range during which automated backups are created if automated backups are enabled.
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[float] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
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
        __props__["db_subnet_group_name"] = db_subnet_group_name
        __props__["dbi_resource_id"] = dbi_resource_id
        __props__["endpoint"] = endpoint
        __props__["engine"] = engine
        __props__["engine_version"] = engine_version
        __props__["identifier"] = identifier
        __props__["identifier_prefix"] = identifier_prefix
        __props__["instance_class"] = instance_class
        __props__["kms_key_id"] = kms_key_id
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

