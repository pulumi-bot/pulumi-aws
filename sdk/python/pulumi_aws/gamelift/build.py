# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
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
    def __init__(__self__, __name__, __opts__=None, name=None, operating_system=None, storage_location=None, version=None):
        """
        Provides an Gamelift Build resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] name: Name of the build
        :param pulumi.Input[str] operating_system: Operating system that the game server binaries are built to run on. e.g. `WINDOWS_2012` or `AMAZON_LINUX`.
        :param pulumi.Input[dict] storage_location: Information indicating where your game build files are stored. See below.
        :param pulumi.Input[str] version: Version that is associated with this build.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        if not operating_system:
            raise TypeError('Missing required property operating_system')
        __props__['operating_system'] = operating_system

        if not storage_location:
            raise TypeError('Missing required property storage_location')
        __props__['storage_location'] = storage_location

        __props__['version'] = version

        super(Build, __self__).__init__(
            'aws:gamelift/build:Build',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

