# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ResourceAssociation(pulumi.CustomResource):
    resource_arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
    """
    resource_share_arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the RAM Resource Share.
    """
    def __init__(__self__, resource_name, opts=None, resource_arn=None, resource_share_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Resource Access Manager (RAM) Resource Association.

        > *NOTE:* Certain AWS resources (e.g. EC2 Subnets) can only be shared in an AWS account that is a member of an AWS Organizations organization with organization-wide Resource Access Manager functionality enabled. See the [Resource Access Manager User Guide](https://docs.aws.amazon.com/ram/latest/userguide/what-is.html) and AWS service specific documentation for additional information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ram.ResourceAssociation("example",
            resource_arn=aws_subnet["example"]["arn"],
            resource_share_arn=aws_ram_resource_share["example"]["arn"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
        :param pulumi.Input[str] resource_share_arn: Amazon Resource Name (ARN) of the RAM Resource Share.
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

            if resource_arn is None:
                raise TypeError("Missing required property 'resource_arn'")
            __props__['resource_arn'] = resource_arn
            if resource_share_arn is None:
                raise TypeError("Missing required property 'resource_share_arn'")
            __props__['resource_share_arn'] = resource_share_arn
        super(ResourceAssociation, __self__).__init__(
            'aws:ram/resourceAssociation:ResourceAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, resource_arn=None, resource_share_arn=None):
        """
        Get an existing ResourceAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_arn: Amazon Resource Name (ARN) of the resource to associate with the RAM Resource Share.
        :param pulumi.Input[str] resource_share_arn: Amazon Resource Name (ARN) of the RAM Resource Share.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["resource_arn"] = resource_arn
        __props__["resource_share_arn"] = resource_share_arn
        return ResourceAssociation(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
