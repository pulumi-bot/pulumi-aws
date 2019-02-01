# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Permission(pulumi.CustomResource):
    allow_ssh: pulumi.Output[bool]
    """
    Whether the user is allowed to use SSH to communicate with the instance
    """
    allow_sudo: pulumi.Output[bool]
    """
    Whether the user is allowed to use sudo to elevate privileges
    """
    level: pulumi.Output[str]
    """
    The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iam_only`
    """
    stack_id: pulumi.Output[str]
    """
    The stack to set the permissions for
    """
    user_arn: pulumi.Output[str]
    """
    The user's IAM ARN to set permissions for
    """
    def __init__(__self__, __name__, __opts__=None, allow_ssh=None, allow_sudo=None, level=None, stack_id=None, user_arn=None):
        """
        Provides an OpsWorks permission resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] allow_ssh: Whether the user is allowed to use SSH to communicate with the instance
        :param pulumi.Input[bool] allow_sudo: Whether the user is allowed to use sudo to elevate privileges
        :param pulumi.Input[str] level: The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iam_only`
        :param pulumi.Input[str] stack_id: The stack to set the permissions for
        :param pulumi.Input[str] user_arn: The user's IAM ARN to set permissions for
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allow_ssh'] = allow_ssh

        __props__['allow_sudo'] = allow_sudo

        __props__['level'] = level

        __props__['stack_id'] = stack_id

        if not user_arn:
            raise TypeError('Missing required property user_arn')
        __props__['user_arn'] = user_arn

        super(Permission, __self__).__init__(
            'aws:opsworks/permission:Permission',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

