# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class TopicSubscription(pulumi.CustomResource):
    """
      Provides a resource for subscribing to SNS topics. Requires that an SNS topic exist for the subscription to attach to.
    This resource allows you to automatically place messages sent to SNS topics in SQS queues, send them as HTTP(S) POST requests
    to a given endpoint, send SMS messages, or notify devices / applications. The most likely use case for Terraform users will
    probably be SQS queues.
    
    > **NOTE:** If the SNS topic and SQS queue are in different AWS regions, it is important for the "aws_sns_topic_subscription" to use an AWS provider that is in the same region of the SNS topic. If the "aws_sns_topic_subscription" is using a provider with a different region than the SNS topic, terraform will fail to create the subscription.
    
    > **NOTE:** Setup of cross-account subscriptions from SNS topics to SQS queues requires Terraform to have access to BOTH accounts.
    
    > **NOTE:** If SNS topic and SQS queue are in different AWS accounts but the same region it is important for the "aws_sns_topic_subscription" to use the AWS provider of the account with the SQS queue. If "aws_sns_topic_subscription" is using a Provider with a different account than the SNS topic, terraform creates the subscriptions but does not keep state and tries to re-create the subscription at every apply.
    
    > **NOTE:** If SNS topic and SQS queue are in different AWS accounts and different AWS regions it is important to recognize that the subscription needs to be initiated from the account with the SQS queue but in the region of the SNS topic.
    """
    def __init__(__self__, __name__, __opts__=None, confirmation_timeout_in_minutes=None, delivery_policy=None, endpoint=None, endpoint_auto_confirms=None, filter_policy=None, protocol=None, raw_message_delivery=None, topic=None):
        """Create a TopicSubscription resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['confirmation_timeout_in_minutes'] = confirmation_timeout_in_minutes

        __props__['delivery_policy'] = delivery_policy

        if not endpoint:
            raise TypeError('Missing required property endpoint')
        __props__['endpoint'] = endpoint

        __props__['endpoint_auto_confirms'] = endpoint_auto_confirms

        __props__['filter_policy'] = filter_policy

        if not protocol:
            raise TypeError('Missing required property protocol')
        __props__['protocol'] = protocol

        __props__['raw_message_delivery'] = raw_message_delivery

        if not topic:
            raise TypeError('Missing required property topic')
        __props__['topic'] = topic

        __props__['arn'] = None

        super(TopicSubscription, __self__).__init__(
            'aws:sns/topicSubscription:TopicSubscription',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

