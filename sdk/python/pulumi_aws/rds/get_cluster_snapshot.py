# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetClusterSnapshotResult:
    """
    A collection of values returned by getClusterSnapshot.
    """
    def __init__(__self__, allocated_storage=None, availability_zones=None, db_cluster_snapshot_arn=None, engine=None, engine_version=None, kms_key_id=None, license_model=None, port=None, snapshot_create_time=None, source_db_cluster_snapshot_arn=None, status=None, storage_encrypted=None, vpc_id=None, id=None):
        if allocated_storage and not isinstance(allocated_storage, int):
            raise TypeError('Expected argument allocated_storage to be a int')
        __self__.allocated_storage = allocated_storage
        """
        Specifies the allocated storage size in gigabytes (GB).
        """
        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError('Expected argument availability_zones to be a list')
        __self__.availability_zones = availability_zones
        """
        List of EC2 Availability Zones that instances in the DB cluster snapshot can be restored in.
        """
        if db_cluster_snapshot_arn and not isinstance(db_cluster_snapshot_arn, str):
            raise TypeError('Expected argument db_cluster_snapshot_arn to be a str')
        __self__.db_cluster_snapshot_arn = db_cluster_snapshot_arn
        """
        The Amazon Resource Name (ARN) for the DB Cluster Snapshot.
        """
        if engine and not isinstance(engine, str):
            raise TypeError('Expected argument engine to be a str')
        __self__.engine = engine
        """
        Specifies the name of the database engine.
        """
        if engine_version and not isinstance(engine_version, str):
            raise TypeError('Expected argument engine_version to be a str')
        __self__.engine_version = engine_version
        """
        Version of the database engine for this DB cluster snapshot.
        """
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError('Expected argument kms_key_id to be a str')
        __self__.kms_key_id = kms_key_id
        """
        If storage_encrypted is true, the AWS KMS key identifier for the encrypted DB cluster snapshot.
        """
        if license_model and not isinstance(license_model, str):
            raise TypeError('Expected argument license_model to be a str')
        __self__.license_model = license_model
        """
        License model information for the restored DB cluster.
        """
        if port and not isinstance(port, int):
            raise TypeError('Expected argument port to be a int')
        __self__.port = port
        """
        Port that the DB cluster was listening on at the time of the snapshot.
        """
        if snapshot_create_time and not isinstance(snapshot_create_time, str):
            raise TypeError('Expected argument snapshot_create_time to be a str')
        __self__.snapshot_create_time = snapshot_create_time
        """
        Time when the snapshot was taken, in Universal Coordinated Time (UTC).
        """
        if source_db_cluster_snapshot_arn and not isinstance(source_db_cluster_snapshot_arn, str):
            raise TypeError('Expected argument source_db_cluster_snapshot_arn to be a str')
        __self__.source_db_cluster_snapshot_arn = source_db_cluster_snapshot_arn
        if status and not isinstance(status, str):
            raise TypeError('Expected argument status to be a str')
        __self__.status = status
        """
        The status of this DB Cluster Snapshot.
        """
        if storage_encrypted and not isinstance(storage_encrypted, bool):
            raise TypeError('Expected argument storage_encrypted to be a bool')
        __self__.storage_encrypted = storage_encrypted
        """
        Specifies whether the DB cluster snapshot is encrypted.
        """
        if vpc_id and not isinstance(vpc_id, str):
            raise TypeError('Expected argument vpc_id to be a str')
        __self__.vpc_id = vpc_id
        """
        The VPC ID associated with the DB cluster snapshot.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_cluster_snapshot(db_cluster_identifier=None,db_cluster_snapshot_identifier=None,include_public=None,include_shared=None,most_recent=None,snapshot_type=None,opts=None):
    """
    Use this data source to get information about a DB Cluster Snapshot for use when provisioning DB clusters.
    
    > **NOTE:** This data source does not apply to snapshots created on DB Instances. 
    See the [`aws_db_snapshot` data source](https://www.terraform.io/docs/providers/aws/d/db_snapshot.html) for DB Instance snapshots.
    """
    __args__ = dict()

    __args__['dbClusterIdentifier'] = db_cluster_identifier
    __args__['dbClusterSnapshotIdentifier'] = db_cluster_snapshot_identifier
    __args__['includePublic'] = include_public
    __args__['includeShared'] = include_shared
    __args__['mostRecent'] = most_recent
    __args__['snapshotType'] = snapshot_type
    __ret__ = await pulumi.runtime.invoke('aws:rds/getClusterSnapshot:getClusterSnapshot', __args__, opts=opts)

    return GetClusterSnapshotResult(
        allocated_storage=__ret__.get('allocatedStorage'),
        availability_zones=__ret__.get('availabilityZones'),
        db_cluster_snapshot_arn=__ret__.get('dbClusterSnapshotArn'),
        engine=__ret__.get('engine'),
        engine_version=__ret__.get('engineVersion'),
        kms_key_id=__ret__.get('kmsKeyId'),
        license_model=__ret__.get('licenseModel'),
        port=__ret__.get('port'),
        snapshot_create_time=__ret__.get('snapshotCreateTime'),
        source_db_cluster_snapshot_arn=__ret__.get('sourceDbClusterSnapshotArn'),
        status=__ret__.get('status'),
        storage_encrypted=__ret__.get('storageEncrypted'),
        vpc_id=__ret__.get('vpcId'),
        id=__ret__.get('id'))
