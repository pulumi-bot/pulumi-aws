# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ClusterInstance(pulumi.CustomResource):
    apply_immediately: pulumi.Output[bool]
    """
    Specifies whether any database modifications
    are applied immediately, or during the next maintenance window. Default is`false`.
    """
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of cluster instance
    """
    auto_minor_version_upgrade: pulumi.Output[bool]
    """
    Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
    """
    availability_zone: pulumi.Output[str]
    """
    The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
    """
    cluster_identifier: pulumi.Output[str]
    """
    The identifier of the [`aws_docdb_cluster`](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html) in which to launch this instance.
    """
    db_subnet_group_name: pulumi.Output[str]
    """
    The DB subnet group to associate with this DB instance.
    """
    dbi_resource_id: pulumi.Output[str]
    """
    The region-unique, immutable identifier for the DB instance.
    """
    endpoint: pulumi.Output[str]
    """
    The DNS address for this instance. May not be writable
    """
    engine: pulumi.Output[str]
    """
    The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
    """
    engine_version: pulumi.Output[str]
    """
    The database engine version
    """
    identifier: pulumi.Output[str]
    identifier_prefix: pulumi.Output[str]
    """
    Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
    """
    instance_class: pulumi.Output[str]
    """
    The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances][2]. DocDB currently
    supports the below instance classes. Please see [AWS Documentation][4] for complete details.
    - db.r4.large
    - db.r4.xlarge
    - db.r4.2xlarge
    - db.r4.4xlarge
    - db.r4.8xlarge
    - db.r4.16xlarge
    """
    kms_key_id: pulumi.Output[str]
    """
    The ARN for the KMS encryption key if one is set to the cluster.
    """
    port: pulumi.Output[float]
    """
    The database port
    """
    preferred_backup_window: pulumi.Output[str]
    """
    The daily time range during which automated backups are created if automated backups are enabled.
    """
    preferred_maintenance_window: pulumi.Output[str]
    """
    The window to perform maintenance in.
    Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
    """
    promotion_tier: pulumi.Output[float]
    """
    Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
    """
    publicly_accessible: pulumi.Output[bool]
    storage_encrypted: pulumi.Output[bool]
    """
    Specifies whether the DB cluster is encrypted.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the instance.
    """
    writer: pulumi.Output[bool]
    """
    Boolean indicating if this instance is writable. `False` indicates this instance is a read replica.
    """
    def __init__(__self__, resource_name, opts=None, apply_immediately=None, auto_minor_version_upgrade=None, availability_zone=None, cluster_identifier=None, engine=None, identifier=None, identifier_prefix=None, instance_class=None, preferred_maintenance_window=None, promotion_tier=None, tags=None, __name__=None, __opts__=None):
        """
        Provides an DocDB Cluster Resource Instance. A Cluster Instance Resource defines
        attributes that are specific to a single instance in a [DocDB Cluster][1].
        
        You do not designate a primary and subsequent replicas. Instead, you simply add DocDB
        Instances and DocDB manages the replication. You can use the [count][3]
        meta-parameter to make multiple instances and join them all to the same DocDB
        Cluster, or you may specify different Cluster Instance resources with various
        `instance_class` sizes.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any database modifications
               are applied immediately, or during the next maintenance window. Default is`false`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window. Default `true`.
        :param pulumi.Input[str] availability_zone: The EC2 Availability Zone that the DB instance is created in. See [docs](https://docs.aws.amazon.com/documentdb/latest/developerguide/API_CreateDBInstance.html) about the details.
        :param pulumi.Input[str] cluster_identifier: The identifier of the [`aws_docdb_cluster`](https://www.terraform.io/docs/providers/aws/r/docdb_cluster.html) in which to launch this instance.
        :param pulumi.Input[str] engine: The name of the database engine to be used for the DocDB instance. Defaults to `docdb`. Valid Values: `docdb`.
        :param pulumi.Input[str] identifier_prefix: Creates a unique identifier beginning with the specified prefix. Conflicts with `identifer`.
        :param pulumi.Input[str] instance_class: The instance class to use. For details on CPU and memory, see [Scaling for DocDB Instances][2]. DocDB currently
               supports the below instance classes. Please see [AWS Documentation][4] for complete details.
               - db.r4.large
               - db.r4.xlarge
               - db.r4.2xlarge
               - db.r4.4xlarge
               - db.r4.8xlarge
               - db.r4.16xlarge
        :param pulumi.Input[str] preferred_maintenance_window: The window to perform maintenance in.
               Syntax: "ddd:hh24:mi-ddd:hh24:mi". Eg: "Mon:00:00-Mon:03:00".
        :param pulumi.Input[float] promotion_tier: Default 0. Failover Priority setting on instance level. The reader who has lower tier has higher priority to get promoter to writer.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the instance.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/docdb_cluster_instance.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apply_immediately'] = apply_immediately

        __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade

        __props__['availability_zone'] = availability_zone

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

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ClusterInstance, __self__).__init__(
            'aws:docdb/clusterInstance:ClusterInstance',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

