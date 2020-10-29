# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SubnetGroup']


class SubnetGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an ElastiCache Subnet Group resource.

        > **NOTE:** ElastiCache Subnet Groups are only for use when working with an
        ElastiCache cluster **inside** of a VPC. If you are on EC2 Classic, see the
        ElastiCache Security Group resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for the cache subnet group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[str] name: Name for the cache subnet group. Elasticache converts this name to lowercase.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: List of VPC Subnet IDs for the cache subnet group
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            __props__['name'] = name
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
        super(SubnetGroup, __self__).__init__(
            'aws:elasticache/subnetGroup:SubnetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SubnetGroup':
        """
        Get an existing SubnetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for the cache subnet group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[str] name: Name for the cache subnet group. Elasticache converts this name to lowercase.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: List of VPC Subnet IDs for the cache subnet group
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["name"] = name
        __props__["subnet_ids"] = subnet_ids
        return SubnetGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description for the cache subnet group. Defaults to "Managed by Pulumi".
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name for the cache subnet group. Elasticache converts this name to lowercase.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        List of VPC Subnet IDs for the cache subnet group
        """
        return pulumi.get(self, "subnet_ids")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

