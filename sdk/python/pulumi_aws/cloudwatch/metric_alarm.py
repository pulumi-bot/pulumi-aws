# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class MetricAlarm(pulumi.CustomResource):
    actions_enabled: pulumi.Output[bool]
    """
    Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
    """
    alarm_actions: pulumi.Output[list]
    """
    The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
    """
    alarm_description: pulumi.Output[str]
    """
    The description for the alarm.
    """
    name: pulumi.Output[str]
    """
    The descriptive name for the alarm. This name must be unique within the user's AWS account
    """
    arn: pulumi.Output[str]
    """
    The ARN of the cloudwatch metric alarm.
    """
    comparison_operator: pulumi.Output[str]
    """
    The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`.
    """
    datapoints_to_alarm: pulumi.Output[float]
    """
    The number of datapoints that must be breaching to trigger the alarm.
    """
    dimensions: pulumi.Output[dict]
    """
    The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
    """
    evaluate_low_sample_count_percentiles: pulumi.Output[str]
    """
    Used only for alarms
    based on percentiles. If you specify `ignore`, the alarm state will not
    change during periods with too few data points to be statistically significant.
    If you specify `evaluate` or omit this parameter, the alarm will always be
    evaluated and possibly change state no matter how many data points are available.
    The following values are supported: `ignore`, and `evaluate`.
    """
    evaluation_periods: pulumi.Output[float]
    """
    The number of periods over which data is compared to the specified threshold.
    """
    extended_statistic: pulumi.Output[str]
    """
    The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
    """
    insufficient_data_actions: pulumi.Output[list]
    """
    The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
    """
    metric_name: pulumi.Output[str]
    """
    The name for this metric.
    See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
    """
    metric_queries: pulumi.Output[list]
    """
    Enables you to create an alarm based on a metric math expression. You may specify at most 20.
    """
    namespace: pulumi.Output[str]
    """
    The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
    See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
    """
    ok_actions: pulumi.Output[list]
    """
    The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
    """
    period: pulumi.Output[float]
    """
    The period in seconds over which the specified `stat` is applied.
    """
    statistic: pulumi.Output[str]
    """
    The statistic to apply to the alarm's associated metric.
    Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    threshold: pulumi.Output[float]
    """
    The value against which the specified statistic is compared.
    """
    treat_missing_data: pulumi.Output[str]
    """
    Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
    """
    unit: pulumi.Output[str]
    """
    The unit for this metric.
    """
    def __init__(__self__, resource_name, opts=None, actions_enabled=None, alarm_actions=None, alarm_description=None, name=None, comparison_operator=None, datapoints_to_alarm=None, dimensions=None, evaluate_low_sample_count_percentiles=None, evaluation_periods=None, extended_statistic=None, insufficient_data_actions=None, metric_name=None, metric_queries=None, namespace=None, ok_actions=None, period=None, statistic=None, tags=None, threshold=None, treat_missing_data=None, unit=None, __name__=None, __opts__=None):
        """
        Provides a CloudWatch Metric Alarm resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
        :param pulumi.Input[list] alarm_actions: The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] alarm_description: The description for the alarm.
        :param pulumi.Input[str] name: The descriptive name for the alarm. This name must be unique within the user's AWS account
        :param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`.
        :param pulumi.Input[float] datapoints_to_alarm: The number of datapoints that must be breaching to trigger the alarm.
        :param pulumi.Input[dict] dimensions: The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[str] evaluate_low_sample_count_percentiles: Used only for alarms
               based on percentiles. If you specify `ignore`, the alarm state will not
               change during periods with too few data points to be statistically significant.
               If you specify `evaluate` or omit this parameter, the alarm will always be
               evaluated and possibly change state no matter how many data points are available.
               The following values are supported: `ignore`, and `evaluate`.
        :param pulumi.Input[float] evaluation_periods: The number of periods over which data is compared to the specified threshold.
        :param pulumi.Input[str] extended_statistic: The percentile statistic for the metric associated with the alarm. Specify a value between p0.0 and p100.
        :param pulumi.Input[list] insufficient_data_actions: The list of actions to execute when this alarm transitions into an INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] metric_name: The name for this metric.
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[list] metric_queries: Enables you to create an alarm based on a metric math expression. You may specify at most 20.
        :param pulumi.Input[str] namespace: The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[list] ok_actions: The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[float] period: The period in seconds over which the specified `stat` is applied.
        :param pulumi.Input[str] statistic: The statistic to apply to the alarm's associated metric.
               Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[float] threshold: The value against which the specified statistic is compared.
        :param pulumi.Input[str] treat_missing_data: Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
        :param pulumi.Input[str] unit: The unit for this metric.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/cloudwatch_metric_alarm.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['actions_enabled'] = actions_enabled

        __props__['alarm_actions'] = alarm_actions

        __props__['alarm_description'] = alarm_description

        __props__['name'] = name

        if comparison_operator is None:
            raise TypeError("Missing required property 'comparison_operator'")
        __props__['comparison_operator'] = comparison_operator

        __props__['datapoints_to_alarm'] = datapoints_to_alarm

        __props__['dimensions'] = dimensions

        __props__['evaluate_low_sample_count_percentiles'] = evaluate_low_sample_count_percentiles

        if evaluation_periods is None:
            raise TypeError("Missing required property 'evaluation_periods'")
        __props__['evaluation_periods'] = evaluation_periods

        __props__['extended_statistic'] = extended_statistic

        __props__['insufficient_data_actions'] = insufficient_data_actions

        __props__['metric_name'] = metric_name

        __props__['metric_queries'] = metric_queries

        __props__['namespace'] = namespace

        __props__['ok_actions'] = ok_actions

        __props__['period'] = period

        __props__['statistic'] = statistic

        __props__['tags'] = tags

        if threshold is None:
            raise TypeError("Missing required property 'threshold'")
        __props__['threshold'] = threshold

        __props__['treat_missing_data'] = treat_missing_data

        __props__['unit'] = unit

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(MetricAlarm, __self__).__init__(
            'aws:cloudwatch/metricAlarm:MetricAlarm',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

