# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventSubscription(pulumi.CustomResource):
    arn: pulumi.Output[str]
    customer_aws_id: pulumi.Output[str]
    enabled: pulumi.Output[bool]
    """
    A boolean flag to enable/disable the subscription. Defaults to true.
    """
    event_categories: pulumi.Output[list]
    """
    A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
    """
    name: pulumi.Output[str]
    """
    The name of the Neptune event subscription. By default generated by Terraform.
    """
    name_prefix: pulumi.Output[str]
    """
    The name of the Neptune event subscription. Conflicts with `name`.
    """
    sns_topic_arn: pulumi.Output[str]
    """
    The ARN of the SNS topic to send events to.
    """
    source_ids: pulumi.Output[list]
    """
    A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
    """
    source_type: pulumi.Output[str]
    """
    The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    def __init__(__self__, __name__, __opts__=None, enabled=None, event_categories=None, name=None, name_prefix=None, sns_topic_arn=None, source_ids=None, source_type=None, tags=None):
        """
        Create a EventSubscription resource with the given unique name, props, and options.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[bool] enabled: A boolean flag to enable/disable the subscription. Defaults to true.
        :param pulumi.Input[list] event_categories: A list of event categories for a `source_type` that you want to subscribe to. Run `aws neptune describe-event-categories` to find all the event categories.
        :param pulumi.Input[str] name: The name of the Neptune event subscription. By default generated by Terraform.
        :param pulumi.Input[str] name_prefix: The name of the Neptune event subscription. Conflicts with `name`.
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic to send events to.
        :param pulumi.Input[list] source_ids: A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
        :param pulumi.Input[str] source_type: The type of source that will be generating the events. Valid options are `db-instance`, `db-security-group`, `db-parameter-group`, `db-snapshot`, `db-cluster` or `db-cluster-snapshot`. If not set, all sources will be subscribed to.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
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

        if not sns_topic_arn:
            raise TypeError('Missing required property sns_topic_arn')
        __props__['sns_topic_arn'] = sns_topic_arn

        __props__['source_ids'] = source_ids

        __props__['source_type'] = source_type

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['customer_aws_id'] = None

        super(EventSubscription, __self__).__init__(
            'aws:neptune/eventSubscription:EventSubscription',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

