# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['User']


class User(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_account_id: Optional[pulumi.Input[str]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 iam_arn: Optional[pulumi.Input[str]] = None,
                 identity_type: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 session_name: Optional[pulumi.Input[str]] = None,
                 user_name: Optional[pulumi.Input[str]] = None,
                 user_role: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Resource for managing QuickSight User

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.quicksight.User("example",
            email="author@example.com",
            identity_type="IAM",
            user_name="an-author",
            user_role="AUTHOR")
        ```

        ## Import

        Importing is currently not supported on this resource.

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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['aws_account_id'] = aws_account_id
            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            __props__['iam_arn'] = iam_arn
            if identity_type is None and not opts.urn:
                raise TypeError("Missing required property 'identity_type'")
            __props__['identity_type'] = identity_type
            __props__['namespace'] = namespace
            __props__['session_name'] = session_name
            __props__['user_name'] = user_name
            if user_role is None and not opts.urn:
                raise TypeError("Missing required property 'user_role'")
            __props__['user_role'] = user_role
            __props__['arn'] = None
        super(User, __self__).__init__(
            'aws:quicksight/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_account_id: Optional[pulumi.Input[str]] = None,
            email: Optional[pulumi.Input[str]] = None,
            iam_arn: Optional[pulumi.Input[str]] = None,
            identity_type: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            session_name: Optional[pulumi.Input[str]] = None,
            user_name: Optional[pulumi.Input[str]] = None,
            user_role: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the user
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsAccountId")
    def aws_account_id(self) -> pulumi.Output[str]:
        """
        The ID for the AWS account that the user is in. Currently, you use the ID for the AWS account that contains your Amazon QuickSight account.
        """
        return pulumi.get(self, "aws_account_id")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The email address of the user that you want to register.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="iamArn")
    def iam_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of the IAM user or role that you are registering with Amazon QuickSight.
        """
        return pulumi.get(self, "iam_arn")

    @property
    @pulumi.getter(name="identityType")
    def identity_type(self) -> pulumi.Output[str]:
        """
        Amazon QuickSight supports several ways of managing the identity of users. This parameter accepts either  `IAM` or `QUICKSIGHT`.
        """
        return pulumi.get(self, "identity_type")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace. Currently, you should set this to `default`.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="sessionName")
    def session_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the IAM session to use when assuming roles that can embed QuickSight dashboards.
        """
        return pulumi.get(self, "session_name")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon QuickSight user name that you want to create for the user you are registering.
        """
        return pulumi.get(self, "user_name")

    @property
    @pulumi.getter(name="userRole")
    def user_role(self) -> pulumi.Output[str]:
        """
        The Amazon QuickSight role of the user. The user role can be one of the following: `READER`, `AUTHOR`, or `ADMIN`
        """
        return pulumi.get(self, "user_role")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

