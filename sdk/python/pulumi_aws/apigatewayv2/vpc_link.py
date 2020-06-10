# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class VpcLink(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The VPC Link ARN.
    """
    name: pulumi.Output[str]
    """
    The name of the VPC Link.
    """
    security_group_ids: pulumi.Output[list]
    """
    Security group IDs for the VPC Link.
    """
    subnet_ids: pulumi.Output[list]
    """
    Subnet IDs for the VPC Link.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the VPC Link.
    """
    def __init__(__self__, resource_name, opts=None, name=None, security_group_ids=None, subnet_ids=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Amazon API Gateway Version 2 VPC Link.

        > **Note:** Amazon API Gateway Version 2 VPC Links enable private integrations that connect HTTP APIs to private resources in a VPC.
        To enable private integration for REST APIs, use the `Amazon API Gateway Version 1 VPC Link` resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.VpcLink("example",
            security_group_ids=[data["ec2.SecurityGroup"]["example"]["id"]],
            subnet_ids=data["ec2.getSubnetIds"]["example"]["ids"],
            tags={
                "Usage": "example",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the VPC Link.
        :param pulumi.Input[list] security_group_ids: Security group IDs for the VPC Link.
        :param pulumi.Input[list] subnet_ids: Subnet IDs for the VPC Link.
        :param pulumi.Input[dict] tags: A map of tags to assign to the VPC Link.
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

            __props__['name'] = name
            if security_group_ids is None:
                raise TypeError("Missing required property 'security_group_ids'")
            __props__['security_group_ids'] = security_group_ids
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            __props__['arn'] = None
        super(VpcLink, __self__).__init__(
            'aws:apigatewayv2/vpcLink:VpcLink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, name=None, security_group_ids=None, subnet_ids=None, tags=None):
        """
        Get an existing VpcLink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The VPC Link ARN.
        :param pulumi.Input[str] name: The name of the VPC Link.
        :param pulumi.Input[list] security_group_ids: Security group IDs for the VPC Link.
        :param pulumi.Input[list] subnet_ids: Subnet IDs for the VPC Link.
        :param pulumi.Input[dict] tags: A map of tags to assign to the VPC Link.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["security_group_ids"] = security_group_ids
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        return VpcLink(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
