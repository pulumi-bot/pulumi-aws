# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
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
    def __init__(__self__, resource_name, opts=None, resource_arn=None, resource_share_arn=None, __name__=None, __opts__=None):
        """
        Manages a Resource Access Manager (RAM) Resource Association.
        
        > *NOTE:* Certain AWS resources (e.g. EC2 Subnets) can only be shared in an AWS account that is a member of an AWS Organizations organization with organization-wide Resource Access Manager functionality enabled. See the [Resource Access Manager User Guide](https://docs.aws.amazon.com/ram/latest/userguide/what-is.html) and AWS service specific documentation for additional information.
        
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
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if resource_arn is None:
            raise TypeError("Missing required property 'resource_arn'")
        __props__['resource_arn'] = resource_arn

        if resource_share_arn is None:
            raise TypeError("Missing required property 'resource_share_arn'")
        __props__['resource_share_arn'] = resource_share_arn

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(ResourceAssociation, __self__).__init__(
            'aws:ram/resourceAssociation:ResourceAssociation',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

