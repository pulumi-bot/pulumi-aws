# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Instance(pulumi.CustomResource):
    """
    Provides an OpsWorks instance resource.
    """
    def __init__(__self__, __name__, __opts__=None, agent_version=None, ami_id=None, architecture=None, auto_scaling_type=None, availability_zone=None, created_at=None, delete_ebs=None, delete_eip=None, ebs_block_devices=None, ebs_optimized=None, ecs_cluster_arn=None, elastic_ip=None, ephemeral_block_devices=None, hostname=None, infrastructure_class=None, install_updates_on_boot=None, instance_profile_arn=None, instance_type=None, last_service_error_id=None, layer_ids=None, os=None, platform=None, private_dns=None, private_ip=None, public_dns=None, public_ip=None, registered_by=None, reported_agent_version=None, reported_os_family=None, reported_os_name=None, reported_os_version=None, root_block_devices=None, root_device_type=None, root_device_volume_id=None, security_group_ids=None, ssh_host_dsa_key_fingerprint=None, ssh_host_rsa_key_fingerprint=None, ssh_key_name=None, stack_id=None, state=None, status=None, subnet_id=None, tenancy=None, virtualization_type=None):
        """Create a Instance resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['agent_version'] = agent_version

        __props__['ami_id'] = ami_id

        __props__['architecture'] = architecture

        __props__['auto_scaling_type'] = auto_scaling_type

        __props__['availability_zone'] = availability_zone

        __props__['created_at'] = created_at

        __props__['delete_ebs'] = delete_ebs

        __props__['delete_eip'] = delete_eip

        __props__['ebs_block_devices'] = ebs_block_devices

        __props__['ebs_optimized'] = ebs_optimized

        __props__['ecs_cluster_arn'] = ecs_cluster_arn

        __props__['elastic_ip'] = elastic_ip

        __props__['ephemeral_block_devices'] = ephemeral_block_devices

        __props__['hostname'] = hostname

        __props__['infrastructure_class'] = infrastructure_class

        __props__['install_updates_on_boot'] = install_updates_on_boot

        __props__['instance_profile_arn'] = instance_profile_arn

        __props__['instance_type'] = instance_type

        __props__['last_service_error_id'] = last_service_error_id

        if not layer_ids:
            raise TypeError('Missing required property layer_ids')
        __props__['layer_ids'] = layer_ids

        __props__['os'] = os

        __props__['platform'] = platform

        __props__['private_dns'] = private_dns

        __props__['private_ip'] = private_ip

        __props__['public_dns'] = public_dns

        __props__['public_ip'] = public_ip

        __props__['registered_by'] = registered_by

        __props__['reported_agent_version'] = reported_agent_version

        __props__['reported_os_family'] = reported_os_family

        __props__['reported_os_name'] = reported_os_name

        __props__['reported_os_version'] = reported_os_version

        __props__['root_block_devices'] = root_block_devices

        __props__['root_device_type'] = root_device_type

        __props__['root_device_volume_id'] = root_device_volume_id

        __props__['security_group_ids'] = security_group_ids

        __props__['ssh_host_dsa_key_fingerprint'] = ssh_host_dsa_key_fingerprint

        __props__['ssh_host_rsa_key_fingerprint'] = ssh_host_rsa_key_fingerprint

        __props__['ssh_key_name'] = ssh_key_name

        if not stack_id:
            raise TypeError('Missing required property stack_id')
        __props__['stack_id'] = stack_id

        __props__['state'] = state

        __props__['status'] = status

        __props__['subnet_id'] = subnet_id

        __props__['tenancy'] = tenancy

        __props__['virtualization_type'] = virtualization_type

        __props__['ec2_instance_id'] = None

        super(Instance, __self__).__init__(
            'aws:opsworks/instance:Instance',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

