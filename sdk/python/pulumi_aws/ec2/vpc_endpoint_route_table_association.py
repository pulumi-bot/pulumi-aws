# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class VpcEndpointRouteTableAssociation(pulumi.CustomResource):
    route_table_id: pulumi.Output[str]
    """
    Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
    """
    vpc_endpoint_id: pulumi.Output[str]
    """
    Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
    """
    def __init__(__self__, resource_name, opts=None, route_table_id=None, vpc_endpoint_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a VPC Endpoint Route Table Association
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
        :param pulumi.Input[str] vpc_endpoint_id: Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_endpoint_route_table_association.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if route_table_id is None:
                raise TypeError("Missing required property 'route_table_id'")
            __props__['route_table_id'] = route_table_id
            if vpc_endpoint_id is None:
                raise TypeError("Missing required property 'vpc_endpoint_id'")
            __props__['vpc_endpoint_id'] = vpc_endpoint_id
        super(VpcEndpointRouteTableAssociation, __self__).__init__(
            'aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, route_table_id=None, vpc_endpoint_id=None):
        """
        Get an existing VpcEndpointRouteTableAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
        :param pulumi.Input[str] vpc_endpoint_id: Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_endpoint_route_table_association.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["route_table_id"] = route_table_id
        __props__["vpc_endpoint_id"] = vpc_endpoint_id
        return VpcEndpointRouteTableAssociation(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

