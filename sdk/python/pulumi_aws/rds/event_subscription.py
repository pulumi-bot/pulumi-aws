# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventSubscription(pulumi.CustomResource):
    """
    Provides a DB event subscription resource.
    """
    def __init__(__self__, __name__, __opts__=None, enabled=None, event_categories=None, name=None, name_prefix=None, sns_topic=None, source_ids=None, source_type=None, tags=None):
        """Create a EventSubscription resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['enabled'] = enabled

        __props__['event_categories'] = event_categories

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        if not sns_topic:
            raise TypeError('Missing required property sns_topic')
        __props__['sns_topic'] = sns_topic

        __props__['source_ids'] = source_ids

        __props__['source_type'] = source_type

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['customer_aws_id'] = None

        super(EventSubscription, __self__).__init__(
            'aws:rds/eventSubscription:EventSubscription',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

