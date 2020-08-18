# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['LogSubscriptionFilter']


class LogSubscriptionFilter(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_arn: Optional[pulumi.Input[str]] = None,
                 distribution: Optional[pulumi.Input[str]] = None,
                 filter_pattern: Optional[pulumi.Input[str]] = None,
                 log_group: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Logs subscription filter resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_lambdafunction_logfilter = aws.cloudwatch.LogSubscriptionFilter("testLambdafunctionLogfilter",
            role_arn=aws_iam_role["iam_for_lambda"]["arn"],
            log_group="/aws/lambda/example_lambda_name",
            filter_pattern="logtype test",
            destination_arn=aws_kinesis_stream["test_logstream"]["arn"],
            distribution="Random")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_arn: The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        :param pulumi.Input[str] distribution: The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        :param pulumi.Input[str] filter_pattern: A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        :param pulumi.Input[str] log_group: The name of the log group to associate the subscription filter with
        :param pulumi.Input[str] name: A name for the subscription filter
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
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
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_arn: Optional[pulumi.Input[str]] = None,
            distribution: Optional[pulumi.Input[str]] = None,
            filter_pattern: Optional[pulumi.Input[str]] = None,
            log_group: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None) -> 'LogSubscriptionFilter':
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
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["destination_arn"] = destination_arn
        __props__["distribution"] = distribution
        __props__["filter_pattern"] = filter_pattern
        __props__["log_group"] = log_group
        __props__["name"] = name
        __props__["role_arn"] = role_arn
        return LogSubscriptionFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> str:
        """
        The ARN of the destination to deliver matching log events to. Kinesis stream or Lambda function ARN.
        """
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter
    def distribution(self) -> Optional[str]:
        """
        The method used to distribute log data to the destination. By default log data is grouped by log stream, but the grouping can be set to random for a more even distribution. This property is only applicable when the destination is an Amazon Kinesis stream. Valid values are "Random" and "ByLogStream".
        """
        return pulumi.get(self, "distribution")

    @property
    @pulumi.getter(name="filterPattern")
    def filter_pattern(self) -> str:
        """
        A valid CloudWatch Logs filter pattern for subscribing to a filtered stream of log events.
        """
        return pulumi.get(self, "filter_pattern")

    @property
    @pulumi.getter(name="logGroup")
    def log_group(self) -> str:
        """
        The name of the log group to associate the subscription filter with
        """
        return pulumi.get(self, "log_group")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A name for the subscription filter
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        """
        The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to deliver ingested log events to the destination. If you use Lambda as a destination, you should skip this argument and use `lambda.Permission` resource for granting access from CloudWatch logs to the destination Lambda function.
        """
        return pulumi.get(self, "role_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

