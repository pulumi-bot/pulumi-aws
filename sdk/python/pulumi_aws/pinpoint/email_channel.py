# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EmailChannel(pulumi.CustomResource):
    application_id: pulumi.Output[str]
    """
    The application ID.
    """
    enabled: pulumi.Output[bool]
    """
    Whether the channel is enabled or disabled. Defaults to `true`.
    """
    from_address: pulumi.Output[str]
    """
    The email address used to send emails from.
    """
    identity: pulumi.Output[str]
    """
    The ARN of an identity verified with SES.
    """
    messages_per_second: pulumi.Output[int]
    """
    Messages per second that can be sent.
    """
    role_arn: pulumi.Output[str]
    """
    The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
    """
    def __init__(__self__, __name__, __opts__=None, application_id=None, enabled=None, from_address=None, identity=None, role_arn=None):
        """
        Provides a Pinpoint SMS Channel resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[bool] enabled: Whether the channel is enabled or disabled. Defaults to `true`.
        :param pulumi.Input[str] from_address: The email address used to send emails from.
        :param pulumi.Input[str] identity: The ARN of an identity verified with SES.
        :param pulumi.Input[str] role_arn: The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if application_id is None:
            raise TypeError('Missing required property application_id')
        __props__['application_id'] = application_id

        __props__['enabled'] = enabled

        if from_address is None:
            raise TypeError('Missing required property from_address')
        __props__['from_address'] = from_address

        if identity is None:
            raise TypeError('Missing required property identity')
        __props__['identity'] = identity

        if role_arn is None:
            raise TypeError('Missing required property role_arn')
        __props__['role_arn'] = role_arn

        __props__['messages_per_second'] = None

        super(EmailChannel, __self__).__init__(
            'aws:pinpoint/emailChannel:EmailChannel',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

