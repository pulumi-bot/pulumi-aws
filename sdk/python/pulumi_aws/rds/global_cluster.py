# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GlobalCluster(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    RDS Global Cluster Amazon Resource Name (ARN)
    """
    database_name: pulumi.Output[str]
    """
    Name for an automatically created database on cluster creation.
    """
    deletion_protection: pulumi.Output[bool]
    """
    If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
    """
    engine: pulumi.Output[str]
    """
    Name of the database engine to be used for this DB cluster. Valid values: `aurora`, `aurora-mysql`. Defaults to `aurora`.
    """
    engine_version: pulumi.Output[str]
    """
    Engine version of the Aurora global database.
    * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
    """
    global_cluster_identifier: pulumi.Output[str]
    """
    The global cluster identifier.
    """
    global_cluster_resource_id: pulumi.Output[str]
    """
    AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
    """
    storage_encrypted: pulumi.Output[bool]
    """
    Specifies whether the DB cluster is encrypted. The default is `false`.
    """
    def __init__(__self__, resource_name, opts=None, database_name=None, deletion_protection=None, engine=None, engine_version=None, global_cluster_identifier=None, storage_encrypted=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.

        More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_global_cluster.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: Name for an automatically created database on cluster creation.
        :param pulumi.Input[bool] deletion_protection: If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        :param pulumi.Input[str] engine: Name of the database engine to be used for this DB cluster. Valid values: `aurora`, `aurora-mysql`. Defaults to `aurora`.
        :param pulumi.Input[str] engine_version: Engine version of the Aurora global database.
               * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        :param pulumi.Input[str] global_cluster_identifier: The global cluster identifier.
        :param pulumi.Input[bool] storage_encrypted: Specifies whether the DB cluster is encrypted. The default is `false`.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['database_name'] = database_name
            __props__['deletion_protection'] = deletion_protection
            __props__['engine'] = engine
            __props__['engine_version'] = engine_version
            if global_cluster_identifier is None:
                raise TypeError("Missing required property 'global_cluster_identifier'")
            __props__['global_cluster_identifier'] = global_cluster_identifier
            __props__['storage_encrypted'] = storage_encrypted
            __props__['arn'] = None
            __props__['global_cluster_resource_id'] = None
        super(GlobalCluster, __self__).__init__(
            'aws:rds/globalCluster:GlobalCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, database_name=None, deletion_protection=None, engine=None, engine_version=None, global_cluster_identifier=None, global_cluster_resource_id=None, storage_encrypted=None):
        """
        Get an existing GlobalCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: RDS Global Cluster Amazon Resource Name (ARN)
        :param pulumi.Input[str] database_name: Name for an automatically created database on cluster creation.
        :param pulumi.Input[bool] deletion_protection: If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
        :param pulumi.Input[str] engine: Name of the database engine to be used for this DB cluster. Valid values: `aurora`, `aurora-mysql`. Defaults to `aurora`.
        :param pulumi.Input[str] engine_version: Engine version of the Aurora global database.
               * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
        :param pulumi.Input[str] global_cluster_identifier: The global cluster identifier.
        :param pulumi.Input[str] global_cluster_resource_id: AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
        :param pulumi.Input[bool] storage_encrypted: Specifies whether the DB cluster is encrypted. The default is `false`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["database_name"] = database_name
        __props__["deletion_protection"] = deletion_protection
        __props__["engine"] = engine
        __props__["engine_version"] = engine_version
        __props__["global_cluster_identifier"] = global_cluster_identifier
        __props__["global_cluster_resource_id"] = global_cluster_resource_id
        __props__["storage_encrypted"] = storage_encrypted
        return GlobalCluster(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

