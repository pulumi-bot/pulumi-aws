# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UserGroup(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    The description of the user group.
    """
    name: pulumi.Output[str]
    """
    The name of the user group.
    """
    precedence: pulumi.Output[int]
    """
    The precedence of the user group.
    """
    role_arn: pulumi.Output[str]
    """
    The ARN of the IAM role to be associated with the user group.
    """
    user_pool_id: pulumi.Output[str]
    """
    The user pool ID.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, precedence=None, role_arn=None, user_pool_id=None, __name__=None, __opts__=None):
        """
        Provides a Cognito User Group resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the user group.
        :param pulumi.Input[str] name: The name of the user group.
        :param pulumi.Input[int] precedence: The precedence of the user group.
        :param pulumi.Input[str] role_arn: The ARN of the IAM role to be associated with the user group.
        :param pulumi.Input[str] user_pool_id: The user pool ID.
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

        __props__['description'] = description

        __props__['name'] = name

        __props__['precedence'] = precedence

        __props__['role_arn'] = role_arn

        if user_pool_id is None:
            raise TypeError('Missing required property user_pool_id')
        __props__['user_pool_id'] = user_pool_id

        super(UserGroup, __self__).__init__(
            'aws:cognito/userGroup:UserGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

