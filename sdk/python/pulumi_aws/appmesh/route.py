# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Route(pulumi.CustomResource):
    """
    Provides an AWS App Mesh route resource.
    """
    def __init__(__self__, __name__, __opts__=None, mesh_name=None, name=None, spec=None, virtual_router_name=None):
        """Create a Route resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not mesh_name:
            raise TypeError('Missing required property mesh_name')
        __props__['mesh_name'] = mesh_name

        __props__['name'] = name

        if not spec:
            raise TypeError('Missing required property spec')
        __props__['spec'] = spec

        if not virtual_router_name:
            raise TypeError('Missing required property virtual_router_name')
        __props__['virtual_router_name'] = virtual_router_name

        __props__['arn'] = None
        __props__['created_date'] = None
        __props__['last_updated_date'] = None

        super(Route, __self__).__init__(
            'aws:appmesh/route:Route',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

