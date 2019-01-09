# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Notification(pulumi.CustomResource):
    """
    Provides an AutoScaling Group with Notification support, via SNS Topics. Each of
    the `notifications` map to a [Notification Configuration][2] inside Amazon Web
    Services, and are applied to each AutoScaling Group you supply.
    """
    def __init__(__self__, __name__, __opts__=None, group_names=None, notifications=None, topic_arn=None):
        """Create a Notification resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not group_names:
            raise TypeError('Missing required property group_names')
        __props__['group_names'] = group_names

        if not notifications:
            raise TypeError('Missing required property notifications')
        __props__['notifications'] = notifications

        if not topic_arn:
            raise TypeError('Missing required property topic_arn')
        __props__['topic_arn'] = topic_arn

        super(Notification, __self__).__init__(
            'aws:autoscaling/notification:Notification',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

