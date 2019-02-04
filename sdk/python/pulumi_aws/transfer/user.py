# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class User(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of Transfer User
    """
    home_directory: pulumi.Output[str]
    """
    The landing directory (folder) for a user when they log in to the server using their SFTP client.
    """
    policy: pulumi.Output[str]
    """
    An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. Since the IAM variable syntax matches Terraform's interpolation syntax, they must be escaped inside Terraform configuration strings (`$${Transfer:UserName}`).
    """
    role: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
    """
    server_id: pulumi.Output[str]
    """
    The Server ID of the Transfer Server (e.g. `s-12345678`)
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    user_name: pulumi.Output[str]
    """
    The name used for log in to your SFTP server.
    """
    def __init__(__self__, __name__, __opts__=None, home_directory=None, policy=None, role=None, server_id=None, tags=None, user_name=None):
        """
        Provides a AWS Transfer User resource. Managing SSH keys can be accomplished with the [`aws_transfer_ssh_key` resource](https://www.terraform.io/docs/providers/aws/r/transfer_ssh_key.html).
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] home_directory: The landing directory (folder) for a user when they log in to the server using their SFTP client.
        :param pulumi.Input[str] policy: An IAM JSON policy document that scopes down user access to portions of their Amazon S3 bucket. IAM variables you can use inside this policy include `${Transfer:UserName}`, `${Transfer:HomeDirectory}`, and `${Transfer:HomeBucket}`. Since the IAM variable syntax matches Terraform's interpolation syntax, they must be escaped inside Terraform configuration strings (`$${Transfer:UserName}`).
        :param pulumi.Input[str] role: Amazon Resource Name (ARN) of an IAM role that allows the service to controls your user’s access to your Amazon S3 bucket.
        :param pulumi.Input[str] server_id: The Server ID of the Transfer Server (e.g. `s-12345678`)
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] user_name: The name used for log in to your SFTP server.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['home_directory'] = home_directory

        __props__['policy'] = policy

        if role is None:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        if server_id is None:
            raise TypeError('Missing required property server_id')
        __props__['server_id'] = server_id

        __props__['tags'] = tags

        if user_name is None:
            raise TypeError('Missing required property user_name')
        __props__['user_name'] = user_name

        __props__['arn'] = None

        super(User, __self__).__init__(
            'aws:transfer/user:User',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

