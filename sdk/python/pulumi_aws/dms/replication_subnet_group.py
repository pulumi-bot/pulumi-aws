# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ReplicationSubnetGroup(pulumi.CustomResource):
    replication_subnet_group_arn: pulumi.Output[str]
    replication_subnet_group_description: pulumi.Output[str]
    """
    The description for the subnet group.
    """
    replication_subnet_group_id: pulumi.Output[str]
    """
    The name for the replication subnet group. This value is stored as a lowercase string.
    """
    subnet_ids: pulumi.Output[list]
    """
    A list of the EC2 subnet IDs for the subnet group.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The ID of the VPC the subnet group is in.
    """
    def __init__(__self__, resource_name, opts=None, replication_subnet_group_description=None, replication_subnet_group_id=None, subnet_ids=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new replication subnet group
        test = aws.dms.ReplicationSubnetGroup("test",
            replication_subnet_group_description="Test replication subnet group",
            replication_subnet_group_id="test-dms-replication-subnet-group-tf",
            subnet_ids=["subnet-12345678"],
            tags={
                "Name": "test",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] replication_subnet_group_description: The description for the subnet group.
        :param pulumi.Input[str] replication_subnet_group_id: The name for the replication subnet group. This value is stored as a lowercase string.
        :param pulumi.Input[list] subnet_ids: A list of the EC2 subnet IDs for the subnet group.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
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

            if replication_subnet_group_description is None:
                raise TypeError("Missing required property 'replication_subnet_group_description'")
            __props__['replication_subnet_group_description'] = replication_subnet_group_description
            if replication_subnet_group_id is None:
                raise TypeError("Missing required property 'replication_subnet_group_id'")
            __props__['replication_subnet_group_id'] = replication_subnet_group_id
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            __props__['replication_subnet_group_arn'] = None
            __props__['vpc_id'] = None
        super(ReplicationSubnetGroup, __self__).__init__(
            'aws:dms/replicationSubnetGroup:ReplicationSubnetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, replication_subnet_group_arn=None, replication_subnet_group_description=None, replication_subnet_group_id=None, subnet_ids=None, tags=None, vpc_id=None):
        """
        Get an existing ReplicationSubnetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] replication_subnet_group_description: The description for the subnet group.
        :param pulumi.Input[str] replication_subnet_group_id: The name for the replication subnet group. This value is stored as a lowercase string.
        :param pulumi.Input[list] subnet_ids: A list of the EC2 subnet IDs for the subnet group.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The ID of the VPC the subnet group is in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["replication_subnet_group_arn"] = replication_subnet_group_arn
        __props__["replication_subnet_group_description"] = replication_subnet_group_description
        __props__["replication_subnet_group_id"] = replication_subnet_group_id
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return ReplicationSubnetGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
