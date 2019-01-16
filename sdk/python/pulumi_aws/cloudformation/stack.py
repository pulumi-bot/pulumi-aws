# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Stack(pulumi.CustomResource):
    """
    Provides a CloudFormation Stack resource.
    """
    def __init__(__self__, __name__, __opts__=None, capabilities=None, disable_rollback=None, iam_role_arn=None, name=None, notification_arns=None, on_failure=None, parameters=None, policy_body=None, policy_url=None, tags=None, template_body=None, template_url=None, timeout_in_minutes=None):
        """Create a Stack resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['capabilities'] = capabilities

        __props__['disable_rollback'] = disable_rollback

        __props__['iam_role_arn'] = iam_role_arn

        __props__['name'] = name

        __props__['notification_arns'] = notification_arns

        __props__['on_failure'] = on_failure

        __props__['parameters'] = parameters

        __props__['policy_body'] = policy_body

        __props__['policy_url'] = policy_url

        __props__['tags'] = tags

        __props__['template_body'] = template_body

        __props__['template_url'] = template_url

        __props__['timeout_in_minutes'] = timeout_in_minutes

        __props__['outputs'] = None

        super(Stack, __self__).__init__(
            'aws:cloudformation/stack:Stack',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

