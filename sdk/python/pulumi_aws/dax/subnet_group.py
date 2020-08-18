# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['SubnetGroup']


class SubnetGroup(pulumi.CustomResource):
    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    A description of the subnet group.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the subnet group.
    """

    subnet_ids: pulumi.Output[List[str]] = pulumi.property("subnetIds")
    """
    A list of VPC subnet IDs for the subnet group.
    """

    vpc_id: pulumi.Output[str] = pulumi.property("vpcId")
    """
    VPC ID of the subnet group.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DAX Subnet Group resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.dax.SubnetGroup("example", subnet_ids=[
            aws_subnet["example1"]["id"],
            aws_subnet["example2"]["id"],
        ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the subnet group.
        :param pulumi.Input[str] name: The name of the subnet group.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: A list of VPC subnet IDs for the subnet group.
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

            __props__['description'] = description
            __props__['name'] = name
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['vpc_id'] = None
        super(SubnetGroup, __self__).__init__(
            'aws:dax/subnetGroup:SubnetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'SubnetGroup':
        """
        Get an existing SubnetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the subnet group.
        :param pulumi.Input[str] name: The name of the subnet group.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: A list of VPC subnet IDs for the subnet group.
        :param pulumi.Input[str] vpc_id: VPC ID of the subnet group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["subnet_ids"] = subnet_ids
        __props__["vpc_id"] = vpc_id
        return SubnetGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

