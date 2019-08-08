# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RouteTable(pulumi.CustomResource):
    default_association_route_table: pulumi.Output[bool]
    """
    Boolean whether this is the default association route table for the EC2 Transit Gateway.
    """
    default_propagation_route_table: pulumi.Output[bool]
    """
    Boolean whether this is the default propagation route table for the EC2 Transit Gateway.
    """
    tags: pulumi.Output[dict]
    """
    Key-value tags for the EC2 Transit Gateway Route Table.
    """
    transit_gateway_id: pulumi.Output[str]
    """
    Identifier of EC2 Transit Gateway.
    """
    def __init__(__self__, resource_name, opts=None, tags=None, transit_gateway_id=None, __name__=None, __opts__=None):
        """
        Manages an EC2 Transit Gateway Route Table.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] tags: Key-value tags for the EC2 Transit Gateway Route Table.
        :param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route_table.html.markdown.
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

        __props__['tags'] = tags
        if transit_gateway_id is None:
            raise TypeError("Missing required property 'transit_gateway_id'")
        __props__['transit_gateway_id'] = transit_gateway_id
        __props__['default_association_route_table'] = None
        __props__['default_propagation_route_table'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(RouteTable, __self__).__init__(
            'aws:ec2transitgateway/routeTable:RouteTable',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

