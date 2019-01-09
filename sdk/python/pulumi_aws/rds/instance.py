# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    """
    Provides an RDS instance resource.  A DB instance is an isolated database
    environment in the cloud.  A DB instance can contain multiple user-created
    databases.
    
    Changes to a DB instance can occur when you manually change a parameter, such as
    `allocated_storage`, and are reflected in the next maintenance window. Because
    of this, Terraform may report a difference in its planning phase because a
    modification has not yet taken place. You can use the `apply_immediately` flag
    to instruct the service to apply the change immediately (see documentation
    below).
    
    When upgrading the major version of an engine, `allow_major_version_upgrade`
    must be set to `true`.
    
    > **Note:** using `apply_immediately` can result in a brief downtime as the
    server reboots. See the AWS Docs on [RDS Maintenance][2] for more information.
    
    > **Note:** All arguments including the username and password will be stored in
    the raw state as plain-text. [Read more about sensitive data in
    state](https://www.terraform.io/docs/state/sensitive-data.html).
    """
    def __init__(__self__, __name__, __opts__=None, allocated_storage=None, allow_major_version_upgrade=None, apply_immediately=None, auto_minor_version_upgrade=None, availability_zone=None, backup_retention_period=None, backup_window=None, character_set_name=None, copy_tags_to_snapshot=None, db_subnet_group_name=None, deletion_protection=None, domain=None, domain_iam_role_name=None, enabled_cloudwatch_logs_exports=None, engine=None, engine_version=None, final_snapshot_identifier=None, iam_database_authentication_enabled=None, identifier=None, identifier_prefix=None, instance_class=None, iops=None, kms_key_id=None, license_model=None, maintenance_window=None, monitoring_interval=None, monitoring_role_arn=None, multi_az=None, name=None, option_group_name=None, parameter_group_name=None, password=None, port=None, publicly_accessible=None, replicate_source_db=None, s3_import=None, security_group_names=None, skip_final_snapshot=None, snapshot_identifier=None, storage_encrypted=None, storage_type=None, tags=None, timezone=None, username=None, vpc_security_group_ids=None):
        """Create a Instance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allocated_storage'] = allocated_storage

        __props__['allow_major_version_upgrade'] = allow_major_version_upgrade

        __props__['apply_immediately'] = apply_immediately

        __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade

        __props__['availability_zone'] = availability_zone

        __props__['backup_retention_period'] = backup_retention_period

        __props__['backup_window'] = backup_window

        __props__['character_set_name'] = character_set_name

        __props__['copy_tags_to_snapshot'] = copy_tags_to_snapshot

        __props__['db_subnet_group_name'] = db_subnet_group_name

        __props__['deletion_protection'] = deletion_protection

        __props__['domain'] = domain

        __props__['domain_iam_role_name'] = domain_iam_role_name

        __props__['enabled_cloudwatch_logs_exports'] = enabled_cloudwatch_logs_exports

        __props__['engine'] = engine

        __props__['engine_version'] = engine_version

        __props__['final_snapshot_identifier'] = final_snapshot_identifier

        __props__['iam_database_authentication_enabled'] = iam_database_authentication_enabled

        __props__['identifier'] = identifier

        __props__['identifier_prefix'] = identifier_prefix

        if not instance_class:
            raise TypeError('Missing required property instance_class')
        __props__['instance_class'] = instance_class

        __props__['iops'] = iops

        __props__['kms_key_id'] = kms_key_id

        __props__['license_model'] = license_model

        __props__['maintenance_window'] = maintenance_window

        __props__['monitoring_interval'] = monitoring_interval

        __props__['monitoring_role_arn'] = monitoring_role_arn

        __props__['multi_az'] = multi_az

        __props__['name'] = name

        __props__['option_group_name'] = option_group_name

        __props__['parameter_group_name'] = parameter_group_name

        __props__['password'] = password

        __props__['port'] = port

        __props__['publicly_accessible'] = publicly_accessible

        __props__['replicate_source_db'] = replicate_source_db

        __props__['s3_import'] = s3_import

        __props__['security_group_names'] = security_group_names

        __props__['skip_final_snapshot'] = skip_final_snapshot

        __props__['snapshot_identifier'] = snapshot_identifier

        __props__['storage_encrypted'] = storage_encrypted

        __props__['storage_type'] = storage_type

        __props__['tags'] = tags

        __props__['timezone'] = timezone

        __props__['username'] = username

        __props__['vpc_security_group_ids'] = vpc_security_group_ids

        __props__['address'] = None
        __props__['arn'] = None
        __props__['ca_cert_identifier'] = None
        __props__['endpoint'] = None
        __props__['hosted_zone_id'] = None
        __props__['replicas'] = None
        __props__['resource_id'] = None
        __props__['status'] = None

        super(Instance, __self__).__init__(
            'aws:rds/instance:Instance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

