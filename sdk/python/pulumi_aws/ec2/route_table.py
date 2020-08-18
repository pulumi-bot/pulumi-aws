# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RouteTable']


class RouteTable(pulumi.CustomResource):
    owner_id: pulumi.Output[str] = pulumi.property("ownerId")
    """
    The ID of the AWS account that owns the route table.
    """

    propagating_vgws: pulumi.Output[List[str]] = pulumi.property("propagatingVgws")
    """
    A list of virtual gateways for propagation.
    """

    routes: pulumi.Output[List['outputs.RouteTableRoute']] = pulumi.property("routes")
    """
    A list of route objects. Their keys are documented below.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A mapping of tags to assign to the resource.
    """

    vpc_id: pulumi.Output[str] = pulumi.property("vpcId")
    """
    The VPC ID.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 propagating_vgws: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 routes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RouteTableRouteArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to create a VPC routing table.

        > **NOTE on Route Tables and Routes:** This provider currently
        provides both a standalone Route resource and a Route Table resource with routes
        defined in-line. At this time you cannot use a Route Table with in-line routes
        in conjunction with any Route resources. Doing so will cause
        a conflict of rule settings and will overwrite rules.

        > **NOTE on `gateway_id` and `nat_gateway_id`:** The AWS API is very forgiving with these two
        attributes and the `ec2.RouteTable` resource can be created with a NAT ID specified as a Gateway ID attribute.
        This _will_ lead to a permanent diff between your configuration and statefile, as the API returns the correct
        parameters in the returned route table. If you're experiencing constant diffs in your `ec2.RouteTable` resources,
        the first thing to check is whether or not you're specifying a NAT ID instead of a Gateway ID, or vice-versa.

        > **NOTE on `propagating_vgws` and the `ec2.VpnGatewayRoutePropagation` resource:**
        If the `propagating_vgws` argument is present, it's not supported to _also_
        define route propagations using `ec2.VpnGatewayRoutePropagation`, since
        this resource will delete any propagating gateways not explicitly listed in
        `propagating_vgws`. Omit this argument when defining route propagation using
        the separate resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        route_table = aws.ec2.RouteTable("routeTable",
            vpc_id=aws_vpc["default"]["id"],
            routes=[
                {
                    "cidr_block": "10.0.1.0/24",
                    "gateway_id": aws_internet_gateway["main"]["id"],
                },
                {
                    "ipv6_cidr_block": "::/0",
                    "egress_only_gateway_id": aws_egress_only_internet_gateway["foo"]["id"],
                },
            ],
            tags={
                "Name": "main",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] propagating_vgws: A list of virtual gateways for propagation.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RouteTableRouteArgs']]]] routes: A list of route objects. Their keys are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['propagating_vgws'] = propagating_vgws
            __props__['routes'] = routes
            __props__['tags'] = tags
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['owner_id'] = None
        super(RouteTable, __self__).__init__(
            'aws:ec2/routeTable:RouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            propagating_vgws: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            routes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['RouteTableRouteArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'RouteTable':
        """
        Get an existing RouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the route table.
        :param pulumi.Input[List[pulumi.Input[str]]] propagating_vgws: A list of virtual gateways for propagation.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['RouteTableRouteArgs']]]] routes: A list of route objects. Their keys are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["owner_id"] = owner_id
        __props__["propagating_vgws"] = propagating_vgws
        __props__["routes"] = routes
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return RouteTable(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

