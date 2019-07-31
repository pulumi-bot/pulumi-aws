# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Project(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name of this project
    """
    name: pulumi.Output[str]
    """
    The name of the project
    """
    def __init__(__self__, resource_name, opts=None, name=None, __name__=None, __opts__=None):
        """
        Provides a resource to manage AWS Device Farm Projects. 
        Please keep in mind that this feature is only supported on the "us-west-2" region.
        This resource will error if you try to create a project in another region.
        
        For more information about Device Farm Projects, see the AWS Documentation on
        [Device Farm Projects][aws-get-project].
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the project

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/devicefarm_project.html.markdown.
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

        __props__['name'] = name

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Project, __self__).__init__(
            'aws:devicefarm/project:Project',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

