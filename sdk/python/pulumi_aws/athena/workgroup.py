# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Workgroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the workgroup
    """
    configuration: pulumi.Output[dict]
    """
    Configuration block with various settings for the workgroup. Documented below.
    """
    description: pulumi.Output[str]
    """
    Description of the workgroup.
    """
    name: pulumi.Output[str]
    """
    Name of the workgroup.
    """
    state: pulumi.Output[str]
    """
    State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags for the workgroup.
    """
    def __init__(__self__, resource_name, opts=None, configuration=None, description=None, name=None, state=None, tags=None, __name__=None, __opts__=None):
        """
        Provides an Athena Workgroup.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] configuration: Configuration block with various settings for the workgroup. Documented below.
        :param pulumi.Input[str] description: Description of the workgroup.
        :param pulumi.Input[str] name: Name of the workgroup.
        :param pulumi.Input[str] state: State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags for the workgroup.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/athena_workgroup.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['configuration'] = configuration
        __props__['description'] = description
        __props__['name'] = name
        __props__['state'] = state
        __props__['tags'] = tags
        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Workgroup, __self__).__init__(
            'aws:athena/workgroup:Workgroup',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

