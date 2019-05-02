# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class SubnetGroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the docDB subnet group.
    """
    description: pulumi.Output[str]
    """
    The description of the docDB subnet group. Defaults to "Managed by Terraform".
    """
    name: pulumi.Output[str]
    """
    The name of the docDB subnet group. If omitted, Terraform will assign a random, unique name.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified prefix. Conflicts with `name`.
    """
    subnet_ids: pulumi.Output[list]
    """
    A list of VPC subnet IDs.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, name_prefix=None, subnet_ids=None, tags=None, __name__=None, __opts__=None):
        """
        Provides an DocumentDB subnet group resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the docDB subnet group. Defaults to "Managed by Terraform".
        :param pulumi.Input[str] name: The name of the docDB subnet group. If omitted, Terraform will assign a random, unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[list] subnet_ids: A list of VPC subnet IDs.
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

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        if subnet_ids is None:
            raise TypeError("Missing required property 'subnet_ids'")
        __props__['subnet_ids'] = subnet_ids

        __props__['tags'] = tags

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(SubnetGroup, __self__).__init__(
            'aws:docdb/subnetGroup:SubnetGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

