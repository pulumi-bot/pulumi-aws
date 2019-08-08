# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventPermission(pulumi.CustomResource):
    action: pulumi.Output[str]
    """
    The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
    """
    condition: pulumi.Output[dict]
    """
    Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
    """
    principal: pulumi.Output[str]
    """
    The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
    """
    statement_id: pulumi.Output[str]
    """
    An identifier string for the external account that you are granting permissions to.
    """
    def __init__(__self__, resource_name, opts=None, action=None, condition=None, principal=None, statement_id=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a CloudWatch Events permission to support cross-account events in the current account default event bus.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: The action that you are enabling the other account to perform. Defaults to `events:PutEvents`.
        :param pulumi.Input[dict] condition: Configuration block to limit the event bus permissions you are granting to only accounts that fulfill the condition. Specified below.
        :param pulumi.Input[str] principal: The 12-digit AWS account ID that you are permitting to put events to your default event bus. Specify `*` to permit any account to put events to your default event bus, optionally limited by `condition`.
        :param pulumi.Input[str] statement_id: An identifier string for the external account that you are granting permissions to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_event_permission.html.markdown.
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

        __props__['action'] = action
        __props__['condition'] = condition
        if principal is None:
            raise TypeError("Missing required property 'principal'")
        __props__['principal'] = principal
        if statement_id is None:
            raise TypeError("Missing required property 'statement_id'")
        __props__['statement_id'] = statement_id
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(EventPermission, __self__).__init__(
            'aws:cloudwatch/eventPermission:EventPermission',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

