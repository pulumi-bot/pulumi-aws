# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecurityGroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The arn of the DB security group.
    """
    description: pulumi.Output[str]
    """
    The description of the DB security group. Defaults to "Managed by Pulumi".
    """
    ingress: pulumi.Output[list]
    """
    A list of ingress rules.
    """
    name: pulumi.Output[str]
    """
    The name of the DB security group.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, description=None, ingress=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an RDS security group resource. This is only for DB instances in the
        EC2-Classic Platform. For instances inside a VPC, use the
        [`aws_db_instance.vpc_security_group_ids`](https://www.terraform.io/docs/providers/aws/r/db_instance.html#vpc_security_group_ids)
        attribute instead.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the DB security group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[list] ingress: A list of ingress rules.
        :param pulumi.Input[str] name: The name of the DB security group.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_security_group.html.markdown.
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

            if description is None:
                description = 'Managed by Pulumi'
            __props__['description'] = description
            if ingress is None:
                raise TypeError("Missing required property 'ingress'")
            __props__['ingress'] = ingress
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(SecurityGroup, __self__).__init__(
            'aws:rds/securityGroup:SecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, description=None, ingress=None, name=None, tags=None):
        """
        Get an existing SecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The arn of the DB security group.
        :param pulumi.Input[str] description: The description of the DB security group. Defaults to "Managed by Pulumi".
        :param pulumi.Input[list] ingress: A list of ingress rules.
        :param pulumi.Input[str] name: The name of the DB security group.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/db_security_group.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["description"] = description
        __props__["ingress"] = ingress
        __props__["name"] = name
        __props__["tags"] = tags
        return SecurityGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

