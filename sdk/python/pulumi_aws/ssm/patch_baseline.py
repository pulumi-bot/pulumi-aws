# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class PatchBaseline(pulumi.CustomResource):
    """
    Provides an SSM Patch Baseline resource
    
    > **NOTE on Patch Baselines:** The `approved_patches` and `approval_rule` are 
    both marked as optional fields, but the Patch Baseline requires that at least one
    of them is specified.
    """
    def __init__(__self__, __name__, __opts__=None, approval_rules=None, approved_patches=None, approved_patches_compliance_level=None, description=None, global_filters=None, name=None, operating_system=None, rejected_patches=None):
        """Create a PatchBaseline resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['approval_rules'] = approval_rules

        __props__['approved_patches'] = approved_patches

        __props__['approved_patches_compliance_level'] = approved_patches_compliance_level

        __props__['description'] = description

        __props__['global_filters'] = global_filters

        __props__['name'] = name

        __props__['operating_system'] = operating_system

        __props__['rejected_patches'] = rejected_patches

        super(PatchBaseline, __self__).__init__(
            'aws:ssm/patchBaseline:PatchBaseline',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

