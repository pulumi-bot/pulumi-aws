# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
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
    def __init__(__self__, resource_name, opts=None, destination_arn=None, distribution=None, filter_pattern=None, log_group=None, name=None, role_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CloudWatch Logs subscription filter resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_arn: The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        :param pulumi.Input[str] distribution: The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        :param pulumi.Input[str] filter_pattern: A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        :param pulumi.Input[str] log_group: The name of the log group to associate the subscription filter with
        :param pulumi.Input[str] name: A name for the subscription filter
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws_lambda_permission` resource for granting access from CloudWatch logs to the destination Lambda function. 

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            if destination_arn is None:
                raise TypeError("Missing required property 'destination_arn'")
            __props__['destination_arn'] = destination_arn
            __props__['distribution'] = distribution
            if filter_pattern is None:
                raise TypeError("Missing required property 'filter_pattern'")
            __props__['filter_pattern'] = filter_pattern
            if log_group is None:
                raise TypeError("Missing required property 'log_group'")
            __props__['log_group'] = log_group
            __props__['name'] = name
            __props__['role_arn'] = role_arn
        super(LogSubscriptionFilter, __self__).__init__(
            'aws:cloudwatch/logSubscriptionFilter:LogSubscriptionFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, destination_arn=None, distribution=None, filter_pattern=None, log_group=None, name=None, role_arn=None):
        """
        Get an existing LogSubscriptionFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_arn: The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        :param pulumi.Input[str] distribution: The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        :param pulumi.Input[str] filter_pattern: A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        :param pulumi.Input[str] log_group: The name of the log group to associate the subscription filter with
        :param pulumi.Input[str] name: A name for the subscription filter
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `aws_lambda_permission` resource for granting access from CloudWatch logs to the destination Lambda function. 

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_log_subscription_filter.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["destination_arn"] = destination_arn
        __props__["distribution"] = distribution
        __props__["filter_pattern"] = filter_pattern
        __props__["log_group"] = log_group
        __props__["name"] = name
        __props__["role_arn"] = role_arn
        return LogSubscriptionFilter(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

