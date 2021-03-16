# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['ClusterInstanceArgs', 'ClusterInstance']

@pulumi.input_type
class ClusterInstanceArgs:
    def __init__(__self__, *,
                 cluster_identifier: pulumi.Input[str],
                 instance_class: pulumi.Input[str],
                 apply_immediately: Optional[pulumi.Input[bool]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 ca_cert_identifier: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 identifier: Optional[pulumi.Input[str]] = None,
                 identifier_prefix: Optional[pulumi.Input[str]] = None,
                 preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
                 promotion_tier: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a ClusterInstance resource.
        :param pulumi.Input[str] cluster_identifier: The identifier of the `docdb.Cluster` in which to launch this instance.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
               supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
               - db.r5.large
               - db.r5.xlarge
               - db.r5.2xlarge
               - db.r5.4xlarge
               - db.r5.12xlarge
               - db.r5.24xlarge
               - db.r4.large
               - db.r4.xlarge
               - db.r4.2xlarge
               - db.r4.4xlarge
               - db.r4.8xlarge
               - db.r4.16xlarge
               - db.t3.medium
        :param pulumi.Input[bool] apply_immediately: Specifies whether any database modifications
               are applied immediately, or during the next maintenance window. Default is`false`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        :param pulumi.Input[str] availability_zone: The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        :param pulumi.Input[str] ca_cert_identifier: (Optional) The identifier of the CA certificate for the DB instance.
        :param pulumi.Input[str] engine: The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        :param pulumi.Input[str] identifier: The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[int] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the instance.
        """
        pulumi.set(__self__, "cluster_identifier", cluster_identifier)
        pulumi.set(__self__, "instance_class", instance_class)
        if apply_immediately is not None:
            pulumi.set(__self__, "apply_immediately", apply_immediately)
        if auto_minor_version_upgrade is not None:
            pulumi.set(__self__, "auto_minor_version_upgrade", auto_minor_version_upgrade)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if ca_cert_identifier is not None:
            pulumi.set(__self__, "ca_cert_identifier", ca_cert_identifier)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if identifier is not None:
            pulumi.set(__self__, "identifier", identifier)
        if identifier_prefix is not None:
            pulumi.set(__self__, "identifier_prefix", identifier_prefix)
        if preferred_maintenance_window is not None:
            pulumi.set(__self__, "preferred_maintenance_window", preferred_maintenance_window)
        if promotion_tier is not None:
            pulumi.set(__self__, "promotion_tier", promotion_tier)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Input[str]:
        """
        The identifier of the `docdb.Cluster` in which to launch this instance.
        """
        return pulumi.get(self, "cluster_identifier")

    @cluster_identifier.setter
    def cluster_identifier(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_identifier", value)

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Input[str]:
        """
        The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
        supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
        - db.r5.large
        - db.r5.xlarge
        - db.r5.2xlarge
        - db.r5.4xlarge
        - db.r5.12xlarge
        - db.r5.24xlarge
        - db.r4.large
        - db.r4.xlarge
        - db.r4.2xlarge
        - db.r4.4xlarge
        - db.r4.8xlarge
        - db.r4.16xlarge
        - db.t3.medium
        """
        return pulumi.get(self, "instance_class")

    @instance_class.setter
    def instance_class(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_class", value)

    @property
    @pulumi.getter(name="applyImmediately")
    def apply_immediately(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether any database modifications
        are applied immediately, or during the next maintenance window. Default is`false`.
        """
        return pulumi.get(self, "apply_immediately")

    @apply_immediately.setter
    def apply_immediately(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "apply_immediately", value)

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        """
        return pulumi.get(self, "auto_minor_version_upgrade")

    @auto_minor_version_upgrade.setter
    def auto_minor_version_upgrade(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_minor_version_upgrade", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="caCertIdentifier")
    def ca_cert_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) The identifier of the CA certificate for the DB instance.
        """
        return pulumi.get(self, "ca_cert_identifier")

    @ca_cert_identifier.setter
    def ca_cert_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert_identifier", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        """
        return pulumi.get(self, "identifier")

    @identifier.setter
    def identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier", value)

    @property
    @pulumi.getter(name="identifierPrefix")
    def identifier_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        """
        return pulumi.get(self, "identifier_prefix")

    @identifier_prefix.setter
    def identifier_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identifier_prefix", value)

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> Optional[pulumi.Input[str]]:
        """
        The window to perform maintenance in.
        Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        """
        return pulumi.get(self, "preferred_maintenance_window")

    @preferred_maintenance_window.setter
    def preferred_maintenance_window(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_maintenance_window", value)

    @property
    @pulumi.getter(name="promotionTier")
    def promotion_tier(self) -> Optional[pulumi.Input[int]]:
        """
        Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        """
        return pulumi.get(self, "promotion_tier")

    @promotion_tier.setter
    def promotion_tier(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "promotion_tier", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the instance.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class ClusterInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
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
                 promotion_tier: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
        attributes that are specific to a single instance in a [DocDB Cluster](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html).

        You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
        Instances and DocDB manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/meta-arguments/count.html)
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

        ## Import

        DocDB Cluster Instances can be imported using the `identifier`, e.g.

        ```sh
         $ pulumi import aws:docdb/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
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
        :param pulumi.Input[str] identifier: The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
               supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
               - db.r5.large
               - db.r5.xlarge
               - db.r5.2xlarge
               - db.r5.4xlarge
               - db.r5.12xlarge
               - db.r5.24xlarge
               - db.r4.large
               - db.r4.xlarge
               - db.r4.2xlarge
               - db.r4.4xlarge
               - db.r4.8xlarge
               - db.r4.16xlarge
               - db.t3.medium
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[int] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClusterInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
        attributes that are specific to a single instance in a [DocDB Cluster](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html).

        You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
        Instances and DocDB manages the replication. You can use the [count](https://www.terraform.io/docs/configuration/meta-arguments/count.html)
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

        ## Import

        DocDB Cluster Instances can be imported using the `identifier`, e.g.

        ```sh
         $ pulumi import aws:docdb/clusterInstance:ClusterInstance prod_instance_1 aurora-cluster-instance-1
        ```

        :param str resource_name: The name of the resource.
        :param ClusterInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
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
                 promotion_tier: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            if cluster_identifier is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_identifier'")
            __props__['cluster_identifier'] = cluster_identifier
            __props__['engine'] = engine
            __props__['identifier'] = identifier
            __props__['identifier_prefix'] = identifier_prefix
            if instance_class is None and not opts.urn:
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
            id: pulumi.Input[str],
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
            port: Optional[pulumi.Input[int]] = None,
            preferred_backup_window: Optional[pulumi.Input[str]] = None,
            preferred_maintenance_window: Optional[pulumi.Input[str]] = None,
            promotion_tier: Optional[pulumi.Input[int]] = None,
            publicly_accessible: Optional[pulumi.Input[bool]] = None,
            storage_encrypted: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            writer: Optional[pulumi.Input[bool]] = None) -> 'ClusterInstance':
        """
        Get an existing ClusterInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[str] identifier: The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
               supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
               - db.r5.large
               - db.r5.xlarge
               - db.r5.2xlarge
               - db.r5.4xlarge
               - db.r5.12xlarge
               - db.r5.24xlarge
               - db.r4.large
               - db.r4.xlarge
               - db.r4.2xlarge
               - db.r4.4xlarge
               - db.r4.8xlarge
               - db.r4.16xlarge
               - db.t3.medium
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key if one is set to the cluster.
        :param pulumi.Input[int] port: The database port
        :param pulumi.Input[str] preferred_backup_window: The daily time range during which automated backups are created if automated backups are enabled.
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[int] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
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

    @property
    @pulumi.getter(name="applyImmediately")
    def apply_immediately(self) -> pulumi.Output[bool]:
        """
        Specifies whether any database modifications
        are applied immediately, or during the next maintenance window. Default is`false`.
        """
        return pulumi.get(self, "apply_immediately")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of cluster instance
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        """
        return pulumi.get(self, "auto_minor_version_upgrade")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="caCertIdentifier")
    def ca_cert_identifier(self) -> pulumi.Output[str]:
        """
        (Optional) The identifier of the CA certificate for the DB instance.
        """
        return pulumi.get(self, "ca_cert_identifier")

    @property
    @pulumi.getter(name="clusterIdentifier")
    def cluster_identifier(self) -> pulumi.Output[str]:
        """
        The identifier of the `docdb.Cluster` in which to launch this instance.
        """
        return pulumi.get(self, "cluster_identifier")

    @property
    @pulumi.getter(name="dbSubnetGroupName")
    def db_subnet_group_name(self) -> pulumi.Output[str]:
        """
        The DB subnet group to associate with this DB instance.
        """
        return pulumi.get(self, "db_subnet_group_name")

    @property
    @pulumi.getter(name="dbiResourceId")
    def dbi_resource_id(self) -> pulumi.Output[str]:
        """
        The region-unique, immutable identifier for the DB instance.
        """
        return pulumi.get(self, "dbi_resource_id")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The DNS address for this instance. May not be writable
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        """
        The database engine version
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter
    def identifier(self) -> pulumi.Output[str]:
        """
        The identifier for the DocDB instance, if omitted, this provider will assign a random, unique identifier.
        """
        return pulumi.get(self, "identifier")

    @property
    @pulumi.getter(name="identifierPrefix")
    def identifier_prefix(self) -> pulumi.Output[str]:
        """
        Creates a unique identifier beginning with the specified prefix. Conflicts with `identifier`.
        """
        return pulumi.get(self, "identifier_prefix")

    @property
    @pulumi.getter(name="instanceClass")
    def instance_class(self) -> pulumi.Output[str]:
        """
        The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-manage-performance.html#db-cluster-manage-scaling-instance). DocDB currently
        supports the below instance classes. Please see [AWS Documentation](https://docs.aws.amazon.com/documentdb/latest/developerguide/db-instance-classes.html#db-instance-class-specs) for complete details.
        - db.r5.large
        - db.r5.xlarge
        - db.r5.2xlarge
        - db.r5.4xlarge
        - db.r5.12xlarge
        - db.r5.24xlarge
        - db.r4.large
        - db.r4.xlarge
        - db.r4.2xlarge
        - db.r4.4xlarge
        - db.r4.8xlarge
        - db.r4.16xlarge
        - db.t3.medium
        """
        return pulumi.get(self, "instance_class")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[str]:
        """
        The ARN for the KMS encryption key if one is set to the cluster.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The database port
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="preferredBackupWindow")
    def preferred_backup_window(self) -> pulumi.Output[str]:
        """
        The daily time range during which automated backups are created if automated backups are enabled.
        """
        return pulumi.get(self, "preferred_backup_window")

    @property
    @pulumi.getter(name="preferredMaintenanceWindow")
    def preferred_maintenance_window(self) -> pulumi.Output[str]:
        """
        The window to perform maintenance in.
        Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        """
        return pulumi.get(self, "preferred_maintenance_window")

    @property
    @pulumi.getter(name="promotionTier")
    def promotion_tier(self) -> pulumi.Output[Optional[int]]:
        """
        Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        """
        return pulumi.get(self, "promotion_tier")

    @property
    @pulumi.getter(name="publiclyAccessible")
    def publicly_accessible(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "publicly_accessible")

    @property
    @pulumi.getter(name="storageEncrypted")
    def storage_encrypted(self) -> pulumi.Output[bool]:
        """
        Specifies whether the DB cluster is encrypted.
        """
        return pulumi.get(self, "storage_encrypted")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the instance.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def writer(self) -> pulumi.Output[bool]:
        """
        Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
        """
        return pulumi.get(self, "writer")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

