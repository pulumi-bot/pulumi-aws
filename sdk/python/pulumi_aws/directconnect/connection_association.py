# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ConnectionAssociation(pulumi.CustomResource):
    connection_id: pulumi.Output[str]
    """
    The ID of the connection.
    """
    lag_id: pulumi.Output[str]
    """
    The ID of the LAG with which to associate the connection.
    """
    def __init__(__self__, resource_name, opts=None, connection_id=None, lag_id=None, __name__=None, __opts__=None):
        """
        Associates a Direct Connect Connection with a LAG.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_id: The ID of the connection.
        :param pulumi.Input[str] lag_id: The ID of the LAG with which to associate the connection.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_connection_association.html.markdown.
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

        if connection_id is None:
            raise TypeError("Missing required property 'connection_id'")
        __props__['connection_id'] = connection_id

        if lag_id is None:
            raise TypeError("Missing required property 'lag_id'")
        __props__['lag_id'] = lag_id

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ConnectionAssociation, __self__).__init__(
            'aws:directconnect/connectionAssociation:ConnectionAssociation',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

