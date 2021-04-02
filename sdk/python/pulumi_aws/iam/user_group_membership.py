# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['UserGroupMembershipArgs', 'UserGroupMembership']

@pulumi.input_type
class UserGroupMembershipArgs:
    def __init__(__self__, *,
                 groups: pulumi.Input[Sequence[pulumi.Input[str]]],
                 user: pulumi.Input[str]):
        """
        The set of arguments for constructing a UserGroupMembership resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: A list of IAM Groups to add the user to
        :param pulumi.Input[str] user: The name of the IAM User to add to groups
        """
        pulumi.set(__self__, "groups", groups)
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of IAM Groups to add the user to
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        The name of the IAM User to add to groups
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)


class UserGroupMembership(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource for adding an IAM User to IAM Groups. This
        resource can be used multiple times with the same user for non-overlapping
        groups.

        To exclusively manage the users in a group, see the
        `iam.GroupMembership` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        user1 = aws.iam.User("user1")
        group1 = aws.iam.Group("group1")
        group2 = aws.iam.Group("group2")
        example1 = aws.iam.UserGroupMembership("example1",
            user=user1.name,
            groups=[
                group1.name,
                group2.name,
            ])
        group3 = aws.iam.Group("group3")
        example2 = aws.iam.UserGroupMembership("example2",
            user=user1.name,
            groups=[group3.name])
        ```

        ## Import

        IAM user group membership can be imported using the user name and group names separated by `/`.

        ```sh
         $ pulumi import aws:iam/userGroupMembership:UserGroupMembership example1 user1/group1/group2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: A list of IAM Groups to add the user to
        :param pulumi.Input[str] user: The name of the IAM User to add to groups
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserGroupMembershipArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource for adding an IAM User to IAM Groups. This
        resource can be used multiple times with the same user for non-overlapping
        groups.

        To exclusively manage the users in a group, see the
        `iam.GroupMembership` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        user1 = aws.iam.User("user1")
        group1 = aws.iam.Group("group1")
        group2 = aws.iam.Group("group2")
        example1 = aws.iam.UserGroupMembership("example1",
            user=user1.name,
            groups=[
                group1.name,
                group2.name,
            ])
        group3 = aws.iam.Group("group3")
        example2 = aws.iam.UserGroupMembership("example2",
            user=user1.name,
            groups=[group3.name])
        ```

        ## Import

        IAM user group membership can be imported using the user name and group names separated by `/`.

        ```sh
         $ pulumi import aws:iam/userGroupMembership:UserGroupMembership example1 user1/group1/group2
        ```

        :param str resource_name: The name of the resource.
        :param UserGroupMembershipArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserGroupMembershipArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if groups is None and not opts.urn:
                raise TypeError("Missing required property 'groups'")
            __props__['groups'] = groups
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__['user'] = user
        super(UserGroupMembership, __self__).__init__(
            'aws:iam/userGroupMembership:UserGroupMembership',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'UserGroupMembership':
        """
        Get an existing UserGroupMembership resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: A list of IAM Groups to add the user to
        :param pulumi.Input[str] user: The name of the IAM User to add to groups
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["groups"] = groups
        __props__["user"] = user
        return UserGroupMembership(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of IAM Groups to add the user to
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        The name of the IAM User to add to groups
        """
        return pulumi.get(self, "user")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

