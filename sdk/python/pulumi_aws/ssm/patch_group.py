# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class PatchGroup(pulumi.CustomResource):
    baseline_id: pulumi.Output[str]
    """
    The ID of the patch baseline to register the patch group with.
    """
    patch_group: pulumi.Output[str]
    """
    The name of the patch group that should be registered with the patch baseline.
    """
    def __init__(__self__, resource_name, opts=None, baseline_id=None, patch_group=None, __name__=None, __opts__=None):
        """
        Provides an SSM Patch Group resource
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] baseline_id: The ID of the patch baseline to register the patch group with.
        :param pulumi.Input[str] patch_group: The name of the patch group that should be registered with the patch baseline.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if baseline_id is None:
            raise TypeError('Missing required property baseline_id')
        __props__['baseline_id'] = baseline_id

        if patch_group is None:
            raise TypeError('Missing required property patch_group')
        __props__['patch_group'] = patch_group

        super(PatchGroup, __self__).__init__(
            'aws:ssm/patchGroup:PatchGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

