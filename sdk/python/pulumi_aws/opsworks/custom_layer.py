# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CustomLayer(pulumi.CustomResource):
    """
    Provides an OpsWorks custom layer resource.
    """
    def __init__(__self__, __name__, __opts__=None, auto_assign_elastic_ips=None, auto_assign_public_ips=None, auto_healing=None, custom_configure_recipes=None, custom_deploy_recipes=None, custom_instance_profile_arn=None, custom_json=None, custom_security_group_ids=None, custom_setup_recipes=None, custom_shutdown_recipes=None, custom_undeploy_recipes=None, drain_elb_on_shutdown=None, ebs_volumes=None, elastic_load_balancer=None, install_updates_on_boot=None, instance_shutdown_timeout=None, name=None, short_name=None, stack_id=None, system_packages=None, use_ebs_optimized_instances=None):
        """Create a CustomLayer resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['auto_assign_elastic_ips'] = auto_assign_elastic_ips

        __props__['auto_assign_public_ips'] = auto_assign_public_ips

        __props__['auto_healing'] = auto_healing

        __props__['custom_configure_recipes'] = custom_configure_recipes

        __props__['custom_deploy_recipes'] = custom_deploy_recipes

        __props__['custom_instance_profile_arn'] = custom_instance_profile_arn

        __props__['custom_json'] = custom_json

        __props__['custom_security_group_ids'] = custom_security_group_ids

        __props__['custom_setup_recipes'] = custom_setup_recipes

        __props__['custom_shutdown_recipes'] = custom_shutdown_recipes

        __props__['custom_undeploy_recipes'] = custom_undeploy_recipes

        __props__['drain_elb_on_shutdown'] = drain_elb_on_shutdown

        __props__['ebs_volumes'] = ebs_volumes

        __props__['elastic_load_balancer'] = elastic_load_balancer

        __props__['install_updates_on_boot'] = install_updates_on_boot

        __props__['instance_shutdown_timeout'] = instance_shutdown_timeout

        __props__['name'] = name

        if not short_name:
            raise TypeError('Missing required property short_name')
        __props__['short_name'] = short_name

        if not stack_id:
            raise TypeError('Missing required property stack_id')
        __props__['stack_id'] = stack_id

        __props__['system_packages'] = system_packages

        __props__['use_ebs_optimized_instances'] = use_ebs_optimized_instances

        super(CustomLayer, __self__).__init__(
            'aws:opsworks/customLayer:CustomLayer',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

