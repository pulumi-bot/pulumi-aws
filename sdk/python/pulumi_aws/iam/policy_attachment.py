# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class PolicyAttachment(pulumi.CustomResource):
    groups: pulumi.Output[list]
    """
    The group(s) the policy should be applied to
    """
    name: pulumi.Output[str]
    """
    The name of the attachment. This cannot be an empty string.
    """
    policy_arn: pulumi.Output[str]
    """
    The ARN of the policy you want to apply
    """
    roles: pulumi.Output[list]
    """
    The role(s) the policy should be applied to
    """
    users: pulumi.Output[list]
    """
    The user(s) the policy should be applied to
    """
    def __init__(__self__, resource_name, opts=None, groups=None, name=None, policy_arn=None, roles=None, users=None, __name__=None, __opts__=None):
        """
        Attaches a Managed IAM Policy to user(s), role(s), and/or group(s)
        
        !> **WARNING:** The aws_iam_policy_attachment resource creates **exclusive** attachments of IAM policies. Across the entire AWS account, all of the users/roles/groups to which a single policy is attached must be declared by a single aws_iam_policy_attachment resource. This means that even any users/roles/groups that have the attached policy via any other mechanism (including other Terraform resources) will have that attached policy revoked by this resource. Consider `aws_iam_role_policy_attachment`, `aws_iam_user_policy_attachment`, or `aws_iam_group_policy_attachment` instead. These resources do not enforce exclusive attachment of an IAM policy.
        
        > **NOTE:** The usage of this resource conflicts with the `aws_iam_group_policy_attachment`, `aws_iam_role_policy_attachment`, and `aws_iam_user_policy_attachment` resources and will permanently show a difference if both are defined.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] groups: The group(s) the policy should be applied to
        :param pulumi.Input[str] name: The name of the attachment. This cannot be an empty string.
        :param pulumi.Input[str] policy_arn: The ARN of the policy you want to apply
        :param pulumi.Input[list] roles: The role(s) the policy should be applied to
        :param pulumi.Input[list] users: The user(s) the policy should be applied to
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['groups'] = groups

        __props__['name'] = name

        if not policy_arn:
            raise TypeError('Missing required property policy_arn')
        __props__['policy_arn'] = policy_arn

        __props__['roles'] = roles

        __props__['users'] = users

        super(PolicyAttachment, __self__).__init__(
            'aws:iam/policyAttachment:PolicyAttachment',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

