# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ParameterGroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Neptune parameter group Amazon Resource Name (ARN).
    """
    description: pulumi.Output[str]
    """
    The description of the Neptune parameter group. Defaults to "Managed by Terraform".
    """
    family: pulumi.Output[str]
    """
    The family of the Neptune parameter group.
    """
    name: pulumi.Output[str]
    """
    The name of the Neptune parameter.
    """
    parameters: pulumi.Output[list]
    """
    A list of Neptune parameters to apply.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, family=None, name=None, parameters=None, tags=None, __name__=None, __opts__=None):
        """
        Manages a Neptune Parameter Group
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Neptune parameter group. Defaults to "Managed by Terraform".
        :param pulumi.Input[str] family: The family of the Neptune parameter group.
        :param pulumi.Input[str] name: The name of the Neptune parameter.
        :param pulumi.Input[list] parameters: A list of Neptune parameters to apply.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
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

        __props__['description'] = description

        if family is None:
            raise TypeError("Missing required property 'family'")
        __props__['family'] = family

        __props__['name'] = name

        __props__['parameters'] = parameters

        __props__['tags'] = tags

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ParameterGroup, __self__).__init__(
            'aws:neptune/parameterGroup:ParameterGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

