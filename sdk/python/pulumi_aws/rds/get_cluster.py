# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetClusterResult:
    """
    A collection of values returned by getCluster.
    """
    def __init__(__self__, arn=None, availability_zones=None, backup_retention_period=None, cluster_identifier=None, cluster_members=None, cluster_resource_id=None, database_name=None, db_cluster_parameter_group_name=None, db_subnet_group_name=None, enabled_cloudwatch_logs_exports=None, endpoint=None, engine=None, engine_version=None, final_snapshot_identifier=None, iam_database_authentication_enabled=None, iam_roles=None, kms_key_id=None, master_username=None, port=None, preferred_backup_window=None, preferred_maintenance_window=None, reader_endpoint=None, replication_source_identifier=None, storage_encrypted=None, tags=None, vpc_security_group_ids=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError('Expected argument availability_zones to be a list')
        __self__.availability_zones = availability_zones
        if backup_retention_period and not isinstance(backup_retention_period, float):
            raise TypeError('Expected argument backup_retention_period to be a float')
        __self__.backup_retention_period = backup_retention_period
        if cluster_identifier and not isinstance(cluster_identifier, str):
            raise TypeError('Expected argument cluster_identifier to be a str')
        __self__.cluster_identifier = cluster_identifier
        if cluster_members and not isinstance(cluster_members, list):
            raise TypeError('Expected argument cluster_members to be a list')
        __self__.cluster_members = cluster_members
        if cluster_resource_id and not isinstance(cluster_resource_id, str):
            raise TypeError('Expected argument cluster_resource_id to be a str')
        __self__.cluster_resource_id = cluster_resource_id
        if database_name and not isinstance(database_name, str):
            raise TypeError('Expected argument database_name to be a str')
        __self__.database_name = database_name
        if db_cluster_parameter_group_name and not isinstance(db_cluster_parameter_group_name, str):
            raise TypeError('Expected argument db_cluster_parameter_group_name to be a str')
        __self__.db_cluster_parameter_group_name = db_cluster_parameter_group_name
        if db_subnet_group_name and not isinstance(db_subnet_group_name, str):
            raise TypeError('Expected argument db_subnet_group_name to be a str')
        __self__.db_subnet_group_name = db_subnet_group_name
        if enabled_cloudwatch_logs_exports and not isinstance(enabled_cloudwatch_logs_exports, list):
            raise TypeError('Expected argument enabled_cloudwatch_logs_exports to be a list')
        __self__.enabled_cloudwatch_logs_exports = enabled_cloudwatch_logs_exports
        if endpoint and not isinstance(endpoint, str):
            raise TypeError('Expected argument endpoint to be a str')
        __self__.endpoint = endpoint
        if engine and not isinstance(engine, str):
            raise TypeError('Expected argument engine to be a str')
        __self__.engine = engine
        if engine_version and not isinstance(engine_version, str):
            raise TypeError('Expected argument engine_version to be a str')
        __self__.engine_version = engine_version
        if final_snapshot_identifier and not isinstance(final_snapshot_identifier, str):
            raise TypeError('Expected argument final_snapshot_identifier to be a str')
        __self__.final_snapshot_identifier = final_snapshot_identifier
        if iam_database_authentication_enabled and not isinstance(iam_database_authentication_enabled, bool):
            raise TypeError('Expected argument iam_database_authentication_enabled to be a bool')
        __self__.iam_database_authentication_enabled = iam_database_authentication_enabled
        if iam_roles and not isinstance(iam_roles, list):
            raise TypeError('Expected argument iam_roles to be a list')
        __self__.iam_roles = iam_roles
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError('Expected argument kms_key_id to be a str')
        __self__.kms_key_id = kms_key_id
        if master_username and not isinstance(master_username, str):
            raise TypeError('Expected argument master_username to be a str')
        __self__.master_username = master_username
        if port and not isinstance(port, float):
            raise TypeError('Expected argument port to be a float')
        __self__.port = port
        if preferred_backup_window and not isinstance(preferred_backup_window, str):
            raise TypeError('Expected argument preferred_backup_window to be a str')
        __self__.preferred_backup_window = preferred_backup_window
        if preferred_maintenance_window and not isinstance(preferred_maintenance_window, str):
            raise TypeError('Expected argument preferred_maintenance_window to be a str')
        __self__.preferred_maintenance_window = preferred_maintenance_window
        if reader_endpoint and not isinstance(reader_endpoint, str):
            raise TypeError('Expected argument reader_endpoint to be a str')
        __self__.reader_endpoint = reader_endpoint
        if replication_source_identifier and not isinstance(replication_source_identifier, str):
            raise TypeError('Expected argument replication_source_identifier to be a str')
        __self__.replication_source_identifier = replication_source_identifier
        if storage_encrypted and not isinstance(storage_encrypted, bool):
            raise TypeError('Expected argument storage_encrypted to be a bool')
        __self__.storage_encrypted = storage_encrypted
        if tags and not isinstance(tags, dict):
            raise TypeError('Expected argument tags to be a dict')
        __self__.tags = tags
        if vpc_security_group_ids and not isinstance(vpc_security_group_ids, list):
            raise TypeError('Expected argument vpc_security_group_ids to be a list')
        __self__.vpc_security_group_ids = vpc_security_group_ids
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_cluster(cluster_identifier=None,tags=None,opts=None):
    """
    Provides information about a RDS cluster.
    """
    __args__ = dict()

    __args__['clusterIdentifier'] = cluster_identifier
    __args__['tags'] = tags
    __ret__ = await pulumi.runtime.invoke('aws:rds/getCluster:getCluster', __args__, opts=opts)

    return GetClusterResult(
        arn=__ret__.get('arn'),
        availability_zones=__ret__.get('availabilityZones'),
        backup_retention_period=__ret__.get('backupRetentionPeriod'),
        cluster_identifier=__ret__.get('clusterIdentifier'),
        cluster_members=__ret__.get('clusterMembers'),
        cluster_resource_id=__ret__.get('clusterResourceId'),
        database_name=__ret__.get('databaseName'),
        db_cluster_parameter_group_name=__ret__.get('dbClusterParameterGroupName'),
        db_subnet_group_name=__ret__.get('dbSubnetGroupName'),
        enabled_cloudwatch_logs_exports=__ret__.get('enabledCloudwatchLogsExports'),
        endpoint=__ret__.get('endpoint'),
        engine=__ret__.get('engine'),
        engine_version=__ret__.get('engineVersion'),
        final_snapshot_identifier=__ret__.get('finalSnapshotIdentifier'),
        iam_database_authentication_enabled=__ret__.get('iamDatabaseAuthenticationEnabled'),
        iam_roles=__ret__.get('iamRoles'),
        kms_key_id=__ret__.get('kmsKeyId'),
        master_username=__ret__.get('masterUsername'),
        port=__ret__.get('port'),
        preferred_backup_window=__ret__.get('preferredBackupWindow'),
        preferred_maintenance_window=__ret__.get('preferredMaintenanceWindow'),
        reader_endpoint=__ret__.get('readerEndpoint'),
        replication_source_identifier=__ret__.get('replicationSourceIdentifier'),
        storage_encrypted=__ret__.get('storageEncrypted'),
        tags=__ret__.get('tags'),
        vpc_security_group_ids=__ret__.get('vpcSecurityGroupIds'),
        id=__ret__.get('id'))
