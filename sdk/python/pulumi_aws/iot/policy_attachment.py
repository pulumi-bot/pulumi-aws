# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class PolicyAttachment(pulumi.CustomResource):
    policy: pulumi.Output[str]
    """
    The name of the policy to attach.
    """
    target: pulumi.Output[str]
    """
    The identity to which the policy is attached.
    """
    def __init__(__self__, __name__, __opts__=None, policy=None, target=None):
        """
        Provides an IoT policy attachment.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] policy: The name of the policy to attach.
        :param pulumi.Input[str] target: The identity to which the policy is attached.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not policy:
            raise TypeError('Missing required property policy')
        __props__['policy'] = policy

        if not target:
            raise TypeError('Missing required property target')
        __props__['target'] = target

        super(PolicyAttachment, __self__).__init__(
            'aws:iot/policyAttachment:PolicyAttachment',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

