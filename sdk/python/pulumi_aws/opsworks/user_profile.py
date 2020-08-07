# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['UserProfile']


class UserProfile(pulumi.CustomResource):
    allow_self_management: pulumi.Output[Optional[bool]] = pulumi.property("allowSelfManagement")
    """
    Whether users can specify their own SSH public key through the My Settings page
    """

    ssh_public_key: pulumi.Output[Optional[str]] = pulumi.property("sshPublicKey")
    """
    The users public key
    """

    ssh_username: pulumi.Output[str] = pulumi.property("sshUsername")
    """
    The ssh username, with witch this user wants to log in
    """

    user_arn: pulumi.Output[str] = pulumi.property("userArn")
    """
    The user's IAM ARN
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_self_management: Optional[pulumi.Input[bool]] = None,
                 ssh_public_key: Optional[pulumi.Input[str]] = None,
                 ssh_username: Optional[pulumi.Input[str]] = None,
                 user_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an OpsWorks User Profile resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        my_profile = aws.opsworks.UserProfile("myProfile",
            ssh_username="my_user",
            user_arn=aws_iam_user["user"]["arn"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_self_management: Whether users can specify their own SSH public key through the My Settings page
        :param pulumi.Input[str] ssh_public_key: The users public key
        :param pulumi.Input[str] ssh_username: The ssh username, with witch this user wants to log in
        :param pulumi.Input[str] user_arn: The user's IAM ARN
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

            __props__['allow_self_management'] = allow_self_management
            __props__['ssh_public_key'] = ssh_public_key
            if ssh_username is None:
                raise TypeError("Missing required property 'ssh_username'")
            __props__['ssh_username'] = ssh_username
            if user_arn is None:
                raise TypeError("Missing required property 'user_arn'")
            __props__['user_arn'] = user_arn
        super(UserProfile, __self__).__init__(
            'aws:opsworks/userProfile:UserProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_self_management: Optional[pulumi.Input[bool]] = None,
            ssh_public_key: Optional[pulumi.Input[str]] = None,
            ssh_username: Optional[pulumi.Input[str]] = None,
            user_arn: Optional[pulumi.Input[str]] = None) -> 'UserProfile':
        """
        Get an existing UserProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_self_management: Whether users can specify their own SSH public key through the My Settings page
        :param pulumi.Input[str] ssh_public_key: The users public key
        :param pulumi.Input[str] ssh_username: The ssh username, with witch this user wants to log in
        :param pulumi.Input[str] user_arn: The user's IAM ARN
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_self_management"] = allow_self_management
        __props__["ssh_public_key"] = ssh_public_key
        __props__["ssh_username"] = ssh_username
        __props__["user_arn"] = user_arn
        return UserProfile(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

