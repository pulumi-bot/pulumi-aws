# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RolePolicyAttachment(pulumi.CustomResource):
    policy_arn: pulumi.Output[str]
    """
    The ARN of the policy you want to apply
    """
    role: pulumi.Output[str]
    """
    The role the policy should be applied to
    """
    def __init__(__self__, __name__, __opts__=None, policy_arn=None, role=None):
        """
        Attaches a Managed IAM Policy to an IAM role
        
        > **NOTE:** The usage of this resource conflicts with the `aws_iam_policy_attachment` resource and will permanently show a difference if both are defined.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] policy_arn: The ARN of the policy you want to apply
        :param pulumi.Input[str] role: The role the policy should be applied to
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not policy_arn:
            raise TypeError('Missing required property policy_arn')
        __props__['policy_arn'] = policy_arn

        if not role:
            raise TypeError('Missing required property role')
        __props__['role'] = role

        super(RolePolicyAttachment, __self__).__init__(
            'aws:iam/rolePolicyAttachment:RolePolicyAttachment',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

