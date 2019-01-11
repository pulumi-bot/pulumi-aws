# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class MountTarget(pulumi.CustomResource):
    """
    Provides an Elastic File System (EFS) mount target.
    """
    def __init__(__self__, __name__, __opts__=None, file_system_id=None, ip_address=None, security_groups=None, subnet_id=None):
        """Create a MountTarget resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not file_system_id:
            raise TypeError('Missing required property file_system_id')
        __props__['file_system_id'] = file_system_id

        __props__['ip_address'] = ip_address

        __props__['security_groups'] = security_groups

        if not subnet_id:
            raise TypeError('Missing required property subnet_id')
        __props__['subnet_id'] = subnet_id

        __props__['dns_name'] = None
        __props__['file_system_arn'] = None
        __props__['network_interface_id'] = None

        super(MountTarget, __self__).__init__(
            'aws:efs/mountTarget:MountTarget',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

