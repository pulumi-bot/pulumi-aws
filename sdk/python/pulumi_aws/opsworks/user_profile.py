# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UserProfile(pulumi.CustomResource):
    allow_self_management: pulumi.Output[bool]
    """
    Whether users can specify their own SSH public key through the My Settings page
    """
    ssh_public_key: pulumi.Output[str]
    """
    The users public key
    """
    ssh_username: pulumi.Output[str]
    """
    The ssh username, with witch this user wants to log in
    """
    user_arn: pulumi.Output[str]
    """
    The user's IAM ARN
    """
    def __init__(__self__, __name__, __opts__=None, allow_self_management=None, ssh_public_key=None, ssh_username=None, user_arn=None):
        """
        Provides an OpsWorks User Profile resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] allow_self_management: Whether users can specify their own SSH public key through the My Settings page
        :param pulumi.Input[str] ssh_public_key: The users public key
        :param pulumi.Input[str] ssh_username: The ssh username, with witch this user wants to log in
        :param pulumi.Input[str] user_arn: The user's IAM ARN
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['allow_self_management'] = allow_self_management

        __props__['ssh_public_key'] = ssh_public_key

        if ssh_username is None:
            raise TypeError('Missing required property ssh_username')
        __props__['ssh_username'] = ssh_username

        if user_arn is None:
            raise TypeError('Missing required property user_arn')
        __props__['user_arn'] = user_arn

        super(UserProfile, __self__).__init__(
            'aws:opsworks/userProfile:UserProfile',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

