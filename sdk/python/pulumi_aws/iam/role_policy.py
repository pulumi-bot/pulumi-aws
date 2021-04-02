# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['RolePolicyArgs', 'RolePolicy']

@pulumi.input_type
class RolePolicyArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input[str],
                 role: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RolePolicy resource.
        :param pulumi.Input[str] policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] role: The IAM role to attach to the policy.
        :param pulumi.Input[str] name: The name of the role policy. If omitted, this provider will
               assign a random, unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        """
        pulumi.set(__self__, "policy", policy)
        pulumi.set(__self__, "role", role)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if name_prefix is not None:
            pulumi.set(__self__, "name_prefix", name_prefix)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        The policy document. This is a JSON formatted string.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The IAM role to attach to the policy.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role policy. If omitted, this provider will
        assign a random, unique name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[pulumi.Input[str]]:
        """
        Creates a unique name beginning with the specified
        prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name_prefix", value)


class RolePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an IAM role inline policy.

        > **NOTE:** For a given role, this resource is incompatible with using the `iam.Role` resource `inline_policy` argument. When using that argument and this resource, both will attempt to manage the role's inline policies and the provider will show a permanent difference.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        test_role = aws.iam.Role("testRole", assume_role_policy=json.dumps({
            "Version": "2012-10-17",
            "Statement": [{
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Sid": "",
                "Principal": {
                    "Service": "ec2.amazonaws.com",
                },
            }],
        }))
        test_policy = aws.iam.RolePolicy("testPolicy",
            role=test_role.id,
            policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Action": ["ec2:Describe*"],
                    "Effect": "Allow",
                    "Resource": "*",
                }],
            }))
        ```

        ## Import

        IAM Role Policies can be imported using the `role_name:role_policy_name`, e.g.

        ```sh
         $ pulumi import aws:iam/rolePolicy:RolePolicy mypolicy role_of_mypolicy_name:mypolicy_name
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the role policy. If omitted, this provider will
               assign a random, unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[str] policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] role: The IAM role to attach to the policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RolePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an IAM role inline policy.

        > **NOTE:** For a given role, this resource is incompatible with using the `iam.Role` resource `inline_policy` argument. When using that argument and this resource, both will attempt to manage the role's inline policies and the provider will show a permanent difference.

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        test_role = aws.iam.Role("testRole", assume_role_policy=json.dumps({
            "Version": "2012-10-17",
            "Statement": [{
                "Action": "sts:AssumeRole",
                "Effect": "Allow",
                "Sid": "",
                "Principal": {
                    "Service": "ec2.amazonaws.com",
                },
            }],
        }))
        test_policy = aws.iam.RolePolicy("testPolicy",
            role=test_role.id,
            policy=json.dumps({
                "Version": "2012-10-17",
                "Statement": [{
                    "Action": ["ec2:Describe*"],
                    "Effect": "Allow",
                    "Resource": "*",
                }],
            }))
        ```

        ## Import

        IAM Role Policies can be imported using the `role_name:role_policy_name`, e.g.

        ```sh
         $ pulumi import aws:iam/rolePolicy:RolePolicy mypolicy role_of_mypolicy_name:mypolicy_name
        ```

        :param str resource_name: The name of the resource.
        :param RolePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RolePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
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

            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
        super(RolePolicy, __self__).__init__(
            'aws:iam/rolePolicy:RolePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'RolePolicy':
        """
        Get an existing RolePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the role policy. If omitted, this provider will
               assign a random, unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[str] policy: The policy document. This is a JSON formatted string.
        :param pulumi.Input[str] role: The IAM role to attach to the policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["policy"] = policy
        __props__["role"] = role
        return RolePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the role policy. If omitted, this provider will
        assign a random, unique name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Output[Optional[str]]:
        """
        Creates a unique name beginning with the specified
        prefix. Conflicts with `name`.
        """
        return pulumi.get(self, "name_prefix")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        The policy document. This is a JSON formatted string.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The IAM role to attach to the policy.
        """
        return pulumi.get(self, "role")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

