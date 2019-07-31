# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Build(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Name of the build
    """
    operating_system: pulumi.Output[str]
    """
    Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
    """
    storage_location: pulumi.Output[dict]
    """
    Information indicating where your game build files are stored. See below.
    """
    version: pulumi.Output[str]
    """
    Version that is associated with this build.
    """
    def __init__(__self__, resource_name, opts=None, name=None, operating_system=None, storage_location=None, version=None, __name__=None, __opts__=None):
        """
        Provides an Gamelift Build resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the build
        :param pulumi.Input[str] operating_system: Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
        :param pulumi.Input[dict] storage_location: Information indicating where your game build files are stored. See below.
        :param pulumi.Input[str] version: Version that is associated with this build.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_build.html.markdown.
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

        if operating_system is None:
            raise TypeError("Missing required property 'operating_system'")
        __props__['operating_system'] = operating_system

        if storage_location is None:
            raise TypeError("Missing required property 'storage_location'")
        __props__['storage_location'] = storage_location

        __props__['version'] = version

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Build, __self__).__init__(
            'aws:gamelift/build:Build',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

