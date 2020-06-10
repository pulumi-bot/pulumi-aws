# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class InternetGateway(pulumi.CustomResource):
    owner_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the internet gateway.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    """
    The VPC ID to create in.
    """
    def __init__(__self__, resource_name, opts=None, tags=None, vpc_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a VPC Internet Gateway.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        gw = aws.ec2.InternetGateway("gw",
            tags={
                "Name": "main",
            },
            vpc_id=aws_vpc["main"]["id"])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID to create in.
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

            __props__['tags'] = tags
            __props__['vpc_id'] = vpc_id
            __props__['owner_id'] = None
        super(InternetGateway, __self__).__init__(
            'aws:ec2/internetGateway:InternetGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, owner_id=None, tags=None, vpc_id=None):
        """
        Get an existing InternetGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the internet gateway.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID to create in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["owner_id"] = owner_id
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return InternetGateway(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
