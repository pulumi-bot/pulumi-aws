# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Role']


class Role(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assume_role_policy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 force_detach_policies: Optional[pulumi.Input[bool]] = None,
                 max_session_duration: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 permissions_boundary: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Role resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            if assume_role_policy is None:
                raise TypeError("Missing required property 'assume_role_policy'")
            __props__['assume_role_policy'] = assume_role_policy
            __props__['description'] = description
            __props__['force_detach_policies'] = force_detach_policies
            __props__['max_session_duration'] = max_session_duration
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['path'] = path
            __props__['permissions_boundary'] = permissions_boundary
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['create_date'] = None
            __props__['unique_id'] = None
        super(Role, __self__).__init__(
            'aws:iam/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            assume_role_policy: Optional[pulumi.Input[str]] = None,
            create_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            force_detach_policies: Optional[pulumi.Input[bool]] = None,
            max_session_duration: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            permissions_boundary: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            unique_id: Optional[pulumi.Input[str]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["assume_role_policy"] = assume_role_policy
        __props__["create_date"] = create_date
        __props__["description"] = description
        __props__["force_detach_policies"] = force_detach_policies
        __props__["max_session_duration"] = max_session_duration
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["path"] = path
        __props__["permissions_boundary"] = permissions_boundary
        __props__["tags"] = tags
        __props__["unique_id"] = unique_id
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="assumeRolePolicy")
    def assume_role_policy(self) -> pulumi.Output[str]:
        return pulumi.get(self, "assume_role_policy")

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> pulumi.Output[str]:
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="forceDetachPolicies")
    def force_detach_policies(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "force_detach_policies")

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "max_session_duration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="permissionsBoundary")
    def permissions_boundary(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "permissions_boundary")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "unique_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

