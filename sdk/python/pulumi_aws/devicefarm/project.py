# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Project(pulumi.CustomResource):
    """
    Provides a resource to manage AWS Device Farm Projects. 
    Please keep in mind that this feature is only supported on the "us-west-2" region.
    This resource will error if you try to create a project in another region.
    
    For more information about Device Farm Projects, see the AWS Documentation on
    [Device Farm Projects][aws-get-project].
    """
    def __init__(__self__, __name__, __opts__=None, name=None):
        """Create a Project resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        __props__['arn'] = None

        super(Project, __self__).__init__(
            'aws:devicefarm/project:Project',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

