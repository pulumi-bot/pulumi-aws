# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ServiceLinkedRole(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) specifying the role.
    """
    aws_service_name: pulumi.Output[str]
    """
    The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
    """
    create_date: pulumi.Output[str]
    """
    The creation date of the IAM role.
    """
    custom_suffix: pulumi.Output[str]
    """
    Additional string appended to the role name. Not all AWS services support custom suffixes.
    """
    description: pulumi.Output[str]
    """
    The description of the role.
    """
    name: pulumi.Output[str]
    """
    The name of the role.
    """
    path: pulumi.Output[str]
    """
    The path of the role.
    """
    unique_id: pulumi.Output[str]
    """
    The stable and unique string identifying the role.
    """
    def __init__(__self__, resource_name, opts=None, aws_service_name=None, custom_suffix=None, description=None, __name__=None, __opts__=None):
        """
        Provides an [IAM service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html).
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_service_name: The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
        :param pulumi.Input[str] custom_suffix: Additional string appended to the role name. Not all AWS services support custom suffixes.
        :param pulumi.Input[str] description: The description of the role.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_service_linked_role.html.markdown.
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

        if aws_service_name is None:
            raise TypeError("Missing required property 'aws_service_name'")
        __props__['aws_service_name'] = aws_service_name

        __props__['custom_suffix'] = custom_suffix

        __props__['description'] = description

        __props__['arn'] = None
        __props__['create_date'] = None
        __props__['name'] = None
        __props__['path'] = None
        __props__['unique_id'] = None

        super(ServiceLinkedRole, __self__).__init__(
            'aws:iam/serviceLinkedRole:ServiceLinkedRole',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

