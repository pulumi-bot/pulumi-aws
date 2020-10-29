# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Permission']


class Permission(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_ssh: Optional[pulumi.Input[bool]] = None,
                 allow_sudo: Optional[pulumi.Input[bool]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 stack_id: Optional[pulumi.Input[str]] = None,
                 user_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an OpsWorks permission resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_ssh: Whether the user is allowed to use SSH to communicate with the instance
        :param pulumi.Input[bool] allow_sudo: Whether the user is allowed to use sudo to elevate privileges
        :param pulumi.Input[str] level: The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iam_only`
        :param pulumi.Input[str] stack_id: The stack to set the permissions for
        :param pulumi.Input[str] user_arn: The user's IAM ARN to set permissions for
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

            __props__['allow_ssh'] = allow_ssh
            __props__['allow_sudo'] = allow_sudo
            __props__['level'] = level
            __props__['stack_id'] = stack_id
            if user_arn is None:
                raise TypeError("Missing required property 'user_arn'")
            __props__['user_arn'] = user_arn
        super(Permission, __self__).__init__(
            'aws:opsworks/permission:Permission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_ssh: Optional[pulumi.Input[bool]] = None,
            allow_sudo: Optional[pulumi.Input[bool]] = None,
            level: Optional[pulumi.Input[str]] = None,
            stack_id: Optional[pulumi.Input[str]] = None,
            user_arn: Optional[pulumi.Input[str]] = None) -> 'Permission':
        """
        Get an existing Permission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_ssh: Whether the user is allowed to use SSH to communicate with the instance
        :param pulumi.Input[bool] allow_sudo: Whether the user is allowed to use sudo to elevate privileges
        :param pulumi.Input[str] level: The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iam_only`
        :param pulumi.Input[str] stack_id: The stack to set the permissions for
        :param pulumi.Input[str] user_arn: The user's IAM ARN to set permissions for
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_ssh"] = allow_ssh
        __props__["allow_sudo"] = allow_sudo
        __props__["level"] = level
        __props__["stack_id"] = stack_id
        __props__["user_arn"] = user_arn
        return Permission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowSsh")
    def allow_ssh(self) -> pulumi.Output[bool]:
        """
        Whether the user is allowed to use SSH to communicate with the instance
        """
        return pulumi.get(self, "allow_ssh")

    @property
    @pulumi.getter(name="allowSudo")
    def allow_sudo(self) -> pulumi.Output[bool]:
        """
        Whether the user is allowed to use sudo to elevate privileges
        """
        return pulumi.get(self, "allow_sudo")

    @property
    @pulumi.getter
    def level(self) -> pulumi.Output[str]:
        """
        The users permission level. Mus be one of `deny`, `show`, `deploy`, `manage`, `iam_only`
        """
        return pulumi.get(self, "level")

    @property
    @pulumi.getter(name="stackId")
    def stack_id(self) -> pulumi.Output[str]:
        """
        The stack to set the permissions for
        """
        return pulumi.get(self, "stack_id")

    @property
    @pulumi.getter(name="userArn")
    def user_arn(self) -> pulumi.Output[str]:
        """
        The user's IAM ARN to set permissions for
        """
        return pulumi.get(self, "user_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

