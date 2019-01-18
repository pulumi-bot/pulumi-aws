# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ThingPrincipalAttachment(pulumi.CustomResource):
    principal: pulumi.Output[str]
    """
    The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
    """
    thing: pulumi.Output[str]
    """
    The name of the thing.
    """
    def __init__(__self__, __name__, __opts__=None, principal=None, thing=None):
        """
        Attaches Principal to AWS IoT Thing.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] principal: The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
        :param pulumi.Input[str] thing: The name of the thing.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not principal:
            raise TypeError('Missing required property principal')
        __props__['principal'] = principal

        if not thing:
            raise TypeError('Missing required property thing')
        __props__['thing'] = thing

        super(ThingPrincipalAttachment, __self__).__init__(
            'aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

