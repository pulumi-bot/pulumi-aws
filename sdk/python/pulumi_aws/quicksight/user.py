# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class User(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the user
    """
    aws_account_id: pulumi.Output[str]
    """
    The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
    """
    email: pulumi.Output[str]
    """
    The email address of the user that you want to register.
    """
    iam_arn: pulumi.Output[str]
    """
    The ARN of the IAM user or role that you are registering with Amazon QuickSight.
    """
    identity_type: pulumi.Output[str]
    """
    Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
    """
    namespace: pulumi.Output[str]
    """
    The namespace. Currently, you should set this to `default`.
    """
    session_name: pulumi.Output[str]
    """
    The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
    """
    user_name: pulumi.Output[str]
    """
    The Amazon QuickSight user name that you want to create for the user you are registering.
    """
    user_role: pulumi.Output[str]
    """
    The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
    """
    def __init__(__self__, resource_name, opts=None, aws_account_id=None, email=None, iam_arn=None, identity_type=None, namespace=None, session_name=None, user_name=None, user_role=None, __props__=None, __name__=None, __opts__=None):
        """
        Resource for managing QuickSight User

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/quicksight_user.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_account_id: The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
        :param pulumi.Input[str] email: The email address of the user that you want to register.
        :param pulumi.Input[str] iam_arn: The ARN of the IAM user or role that you are registering with Amazon QuickSight.
        :param pulumi.Input[str] identity_type: Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
        :param pulumi.Input[str] namespace: The namespace. Currently, you should set this to `default`.
        :param pulumi.Input[str] session_name: The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
        :param pulumi.Input[str] user_name: The Amazon QuickSight user name that you want to create for the user you are registering.
        :param pulumi.Input[str] user_role: The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
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

            __props__['aws_account_id'] = aws_account_id
            if email is None:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            __props__['iam_arn'] = iam_arn
            if identity_type is None:
                raise TypeError("Missing required property 'identity_type'")
            __props__['identity_type'] = identity_type
            __props__['namespace'] = namespace
            __props__['session_name'] = session_name
            __props__['user_name'] = user_name
            if user_role is None:
                raise TypeError("Missing required property 'user_role'")
            __props__['user_role'] = user_role
            __props__['arn'] = None
        super(User, __self__).__init__(
            'aws:quicksight/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, aws_account_id=None, email=None, iam_arn=None, identity_type=None, namespace=None, session_name=None, user_name=None, user_role=None):
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the user
        :param pulumi.Input[str] aws_account_id: The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
        :param pulumi.Input[str] email: The email address of the user that you want to register.
        :param pulumi.Input[str] iam_arn: The ARN of the IAM user or role that you are registering with Amazon QuickSight.
        :param pulumi.Input[str] identity_type: Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
        :param pulumi.Input[str] namespace: The namespace. Currently, you should set this to `default`.
        :param pulumi.Input[str] session_name: The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
        :param pulumi.Input[str] user_name: The Amazon QuickSight user name that you want to create for the user you are registering.
        :param pulumi.Input[str] user_role: The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["aws_account_id"] = aws_account_id
        __props__["email"] = email
        __props__["iam_arn"] = iam_arn
        __props__["identity_type"] = identity_type
        __props__["namespace"] = namespace
        __props__["session_name"] = session_name
        __props__["user_name"] = user_name
        __props__["user_role"] = user_role
        return User(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

