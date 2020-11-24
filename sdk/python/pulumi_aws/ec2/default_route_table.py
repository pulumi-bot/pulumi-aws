# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DefaultRouteTable']


class DefaultRouteTable(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_route_table_id: Optional[pulumi.Input[str]] = None,
                 propagating_vgws: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultRouteTableRouteArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage a Default VPC Routing Table.

        Each VPC created in AWS comes with a Default Route Table that can be managed, but not
        destroyed. **This is an advanced resource**, and has special caveats to be aware
        of when using it. Please read this document in its entirety before using this
        resource. It is recommended you **do not** use both `ec2.DefaultRouteTable` to
        manage the default route table **and** use the `ec2.MainRouteTableAssociation`,
        due to possible conflict in routes.

        The `ec2.DefaultRouteTable` behaves differently from normal resources, in that
        this provider does not _create_ this resource, but instead attempts to "adopt" it
        into management. We can do this because each VPC created has a Default Route
        Table that cannot be destroyed, and is created with a single route.

        When this provider first adopts the Default Route Table, it **immediately removes all
        defined routes**. It then proceeds to create any routes specified in the
        configuration. This step is required so that only the routes specified in the
        configuration present in the Default Route Table.

        For more information about Route Tables, see the AWS Documentation on
        [Route Tables][aws-route-tables].

        For more information about managing normal Route Tables in this provider, see our
        documentation on [ec2.RouteTable][tf-route-tables].

        > **NOTE on Route Tables and Routes:** This provider currently
        provides both a standalone Route resource and a Route Table resource with routes
        defined in-line. At this time you cannot use a Route Table with in-line routes
        in conjunction with any Route resources. Doing so will cause
        a conflict of rule settings and will overwrite routes.

        ## Example Usage
        ### With Tags

        ```python
        import pulumi
        import pulumi_aws as aws

        default_route_table = aws.ec2.DefaultRouteTable("defaultRouteTable",
            default_route_table_id=aws_vpc["foo"]["default_route_table_id"],
            routes=[{}],
            tags={
                "Name": "default table",
            })
        ```

        ## Import

        Default VPC Routing tables can be imported using the `vpc_id`, e.g.

        ```sh
         $ pulumi import aws:ec2/defaultRouteTable:DefaultRouteTable example vpc-33cc44dd
        ```

         [aws-route-tables]http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Route_Tables.html#Route_Replacing_Main_Table [tf-route-tables]/docs/providers/aws/r/route_table.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_route_table_id: The ID of the Default Routing Table.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] propagating_vgws: A list of virtual gateways for propagation.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultRouteTableRouteArgs']]]] routes: A list of route objects. Their keys are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
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

            if default_route_table_id is None:
                raise TypeError("Missing required property 'default_route_table_id'")
            __props__['default_route_table_id'] = default_route_table_id
            __props__['propagating_vgws'] = propagating_vgws
            __props__['routes'] = routes
            __props__['tags'] = tags
            __props__['owner_id'] = None
            __props__['vpc_id'] = None
        super(DefaultRouteTable, __self__).__init__(
            'aws:ec2/defaultRouteTable:DefaultRouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_route_table_id: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            propagating_vgws: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            routes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultRouteTableRouteArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'DefaultRouteTable':
        """
        Get an existing DefaultRouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_route_table_id: The ID of the Default Routing Table.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the route table
        :param pulumi.Input[Sequence[pulumi.Input[str]]] propagating_vgws: A list of virtual gateways for propagation.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultRouteTableRouteArgs']]]] routes: A list of route objects. Their keys are documented below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A mapping of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["default_route_table_id"] = default_route_table_id
        __props__["owner_id"] = owner_id
        __props__["propagating_vgws"] = propagating_vgws
        __props__["routes"] = routes
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return DefaultRouteTable(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultRouteTableId")
    def default_route_table_id(self) -> pulumi.Output[str]:
        """
        The ID of the Default Routing Table.
        """
        return pulumi.get(self, "default_route_table_id")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The ID of the AWS account that owns the route table
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="propagatingVgws")
    def propagating_vgws(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of virtual gateways for propagation.
        """
        return pulumi.get(self, "propagating_vgws")

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Output[Sequence['outputs.DefaultRouteTableRoute']]:
        """
        A list of route objects. Their keys are documented below.
        """
        return pulumi.get(self, "routes")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A mapping of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

