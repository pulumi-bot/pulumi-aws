# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Stack(pulumi.CustomResource):
    """
    Provides an OpsWorks stack resource.
    """
    def __init__(__self__, __name__, __opts__=None, agent_version=None, berkshelf_version=None, color=None, configuration_manager_name=None, configuration_manager_version=None, custom_cookbooks_sources=None, custom_json=None, default_availability_zone=None, default_instance_profile_arn=None, default_os=None, default_root_device_type=None, default_ssh_key_name=None, default_subnet_id=None, hostname_theme=None, manage_berkshelf=None, name=None, region=None, service_role_arn=None, tags=None, use_custom_cookbooks=None, use_opsworks_security_groups=None, vpc_id=None):
        """Create a Stack resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['agent_version'] = agent_version

        __props__['berkshelf_version'] = berkshelf_version

        __props__['color'] = color

        __props__['configuration_manager_name'] = configuration_manager_name

        __props__['configuration_manager_version'] = configuration_manager_version

        __props__['custom_cookbooks_sources'] = custom_cookbooks_sources

        __props__['custom_json'] = custom_json

        __props__['default_availability_zone'] = default_availability_zone

        if not default_instance_profile_arn:
            raise TypeError('Missing required property default_instance_profile_arn')
        __props__['default_instance_profile_arn'] = default_instance_profile_arn

        __props__['default_os'] = default_os

        __props__['default_root_device_type'] = default_root_device_type

        __props__['default_ssh_key_name'] = default_ssh_key_name

        __props__['default_subnet_id'] = default_subnet_id

        __props__['hostname_theme'] = hostname_theme

        __props__['manage_berkshelf'] = manage_berkshelf

        __props__['name'] = name

        if not region:
            raise TypeError('Missing required property region')
        __props__['region'] = region

        if not service_role_arn:
            raise TypeError('Missing required property service_role_arn')
        __props__['service_role_arn'] = service_role_arn

        __props__['tags'] = tags

        __props__['use_custom_cookbooks'] = use_custom_cookbooks

        __props__['use_opsworks_security_groups'] = use_opsworks_security_groups

        __props__['vpc_id'] = vpc_id

        __props__['arn'] = None
        __props__['stack_endpoint'] = None

        super(Stack, __self__).__init__(
            'aws:opsworks/stack:Stack',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

