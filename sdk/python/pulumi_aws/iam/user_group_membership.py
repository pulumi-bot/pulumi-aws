# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class UserGroupMembership(pulumi.CustomResource):
    groups: pulumi.Output[list]
    """
    A list of [IAM Groups][1] to add the user to
    """
    user: pulumi.Output[str]
    """
    The name of the [IAM User][2] to add to groups
    """
    def __init__(__self__, resource_name, opts=None, groups=None, user=None, __name__=None, __opts__=None):
        """
        Provides a resource for adding an [IAM User][2] to [IAM Groups][1]. This
        resource can be used multiple times with the same user for non-overlapping
        groups.
        
        To exclusively manage the users in a group, see the
        [`aws_iam_group_membership` resource][3].
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] groups: A list of [IAM Groups][1] to add the user to
        :param pulumi.Input[str] user: The name of the [IAM User][2] to add to groups

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/iam_user_group_membership.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if groups is None:
            raise TypeError("Missing required property 'groups'")
        __props__['groups'] = groups
        if user is None:
            raise TypeError("Missing required property 'user'")
        __props__['user'] = user
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(UserGroupMembership, __self__).__init__(
            'aws:iam/userGroupMembership:UserGroupMembership',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

