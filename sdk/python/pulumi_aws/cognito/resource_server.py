# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ResourceServer(pulumi.CustomResource):
    """
    Provides a Cognito Resource Server.
    """
    def __init__(__self__, __name__, __opts__=None, identifier=None, name=None, scopes=None, user_pool_id=None):
        """Create a ResourceServer resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not identifier:
            raise TypeError('Missing required property identifier')
        __props__['identifier'] = identifier

        __props__['name'] = name

        __props__['scopes'] = scopes

        if not user_pool_id:
            raise TypeError('Missing required property user_pool_id')
        __props__['user_pool_id'] = user_pool_id

        __props__['scope_identifiers'] = None

        super(ResourceServer, __self__).__init__(
            'aws:cognito/resourceServer:ResourceServer',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

