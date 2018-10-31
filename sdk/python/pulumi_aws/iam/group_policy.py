# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class GroupPolicy(pulumi.CustomResource):
    """
    Provides an IAM policy attached to a group.
    """
    def __init__(__self__, __name__, __opts__=None, group=None, name=None, name_prefix=None, policy=None):
        """Create a GroupPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not group:
            raise TypeError('Missing required property group')
        __props__['group'] = group

        __props__['name'] = name

        __props__['namePrefix'] = name_prefix

        if not policy:
            raise TypeError('Missing required property policy')
        __props__['policy'] = policy

        super(GroupPolicy, __self__).__init__(
            'aws:iam/groupPolicy:GroupPolicy',
            __name__,
            __props__,
            __opts__)

