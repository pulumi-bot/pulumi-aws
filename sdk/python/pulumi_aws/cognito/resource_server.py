# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ResourceServer(pulumi.CustomResource):
    identifier: pulumi.Output[str]
    """
    An identifier for the resource server.
    """
    name: pulumi.Output[str]
    """
    A name for the resource server.
    """
    scopes: pulumi.Output[list]
    """
    A list of Authorization Scope.
    """
    scope_identifiers: pulumi.Output[list]
    """
    A list of all scopes configured for this resource server in the format identifier/scope_name.
    """
    user_pool_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, identifier=None, name=None, scopes=None, user_pool_id=None, __name__=None, __opts__=None):
        """
        Provides a Cognito Resource Server.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identifier: An identifier for the resource server.
        :param pulumi.Input[str] name: A name for the resource server.
        :param pulumi.Input[list] scopes: A list of Authorization Scope.
        :param pulumi.Input[str] user_pool_id
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
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

