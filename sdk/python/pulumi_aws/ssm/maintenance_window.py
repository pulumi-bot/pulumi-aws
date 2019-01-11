# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class MaintenanceWindow(pulumi.CustomResource):
    """
    Provides an SSM Maintenance Window resource
    """
    def __init__(__self__, __name__, __opts__=None, allow_unassociated_targets=None, cutoff=None, duration=None, enabled=None, name=None, schedule=None):
        """Create a MaintenanceWindow resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allow_unassociated_targets'] = allow_unassociated_targets

        if not cutoff:
            raise TypeError('Missing required property cutoff')
        __props__['cutoff'] = cutoff

        if not duration:
            raise TypeError('Missing required property duration')
        __props__['duration'] = duration

        __props__['enabled'] = enabled

        __props__['name'] = name

        if not schedule:
            raise TypeError('Missing required property schedule')
        __props__['schedule'] = schedule

        super(MaintenanceWindow, __self__).__init__(
            'aws:ssm/maintenanceWindow:MaintenanceWindow',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

