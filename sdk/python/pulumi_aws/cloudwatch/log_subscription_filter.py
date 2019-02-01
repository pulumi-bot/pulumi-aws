# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class LogSubscriptionFilter(pulumi.CustomResource):
    destination_arn: pulumi.Output[str]
    """
    The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
    """
    distribution: pulumi.Output[str]
    """
    The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
    """
    filter_pattern: pulumi.Output[str]
    """
    A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
    """
    log_group: pulumi.Output[str]
    """
    The name of the log group to associate the subscription filter with
    """
    name: pulumi.Output[str]
    """
    A name for the subscription filter
    """
    role_arn: pulumi.Output[str]
    """
    The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws_lambda_permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
    """
    def __init__(__self__, __name__, __opts__=None, destination_arn=None, distribution=None, filter_pattern=None, log_group=None, name=None, role_arn=None):
        """
        Provides a CloudWatch Logs subscription filter resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] destination_arn: The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        :param pulumi.Input[str] distribution: The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        :param pulumi.Input[str] filter_pattern: A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        :param pulumi.Input[str] log_group: The name of the log group to associate the subscription filter with
        :param pulumi.Input[str] name: A name for the subscription filter
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws_lambda_permission` resource for granting access from CloudWatch logs to the destination Lambda function. 
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not destination_arn:
            raise TypeError('Missing required property destination_arn')
        __props__['destination_arn'] = destination_arn

        __props__['distribution'] = distribution

        if not filter_pattern:
            raise TypeError('Missing required property filter_pattern')
        __props__['filter_pattern'] = filter_pattern

        if not log_group:
            raise TypeError('Missing required property log_group')
        __props__['log_group'] = log_group

        __props__['name'] = name

        __props__['role_arn'] = role_arn

        super(LogSubscriptionFilter, __self__).__init__(
            'aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

