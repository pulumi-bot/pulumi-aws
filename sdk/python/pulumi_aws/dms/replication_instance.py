# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ReplicationInstance(pulumi.CustomResource):
    """
    Provides a DMS (Data Migration Service) replication instance resource. DMS replication instances can be created, updated, deleted, and imported.
    """
    def __init__(__self__, __name__, __opts__=None, allocated_storage=None, apply_immediately=None, auto_minor_version_upgrade=None, availability_zone=None, engine_version=None, kms_key_arn=None, multi_az=None, preferred_maintenance_window=None, publicly_accessible=None, replication_instance_class=None, replication_instance_id=None, replication_subnet_group_id=None, tags=None, vpc_security_group_ids=None):
        """Create a ReplicationInstance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allocated_storage'] = allocated_storage

        __props__['apply_immediately'] = apply_immediately

        __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade

        __props__['availability_zone'] = availability_zone

        __props__['engine_version'] = engine_version

        __props__['kms_key_arn'] = kms_key_arn

        __props__['multi_az'] = multi_az

        __props__['preferred_maintenance_window'] = preferred_maintenance_window

        __props__['publicly_accessible'] = publicly_accessible

        if not replication_instance_class:
            raise TypeError('Missing required property replication_instance_class')
        __props__['replication_instance_class'] = replication_instance_class

        if not replication_instance_id:
            raise TypeError('Missing required property replication_instance_id')
        __props__['replication_instance_id'] = replication_instance_id

        __props__['replication_subnet_group_id'] = replication_subnet_group_id

        __props__['tags'] = tags

        __props__['vpc_security_group_ids'] = vpc_security_group_ids

        __props__['replication_instance_arn'] = None
        __props__['replication_instance_private_ips'] = None
        __props__['replication_instance_public_ips'] = None

        super(ReplicationInstance, __self__).__init__(
            'aws:dms/replicationInstance:ReplicationInstance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

