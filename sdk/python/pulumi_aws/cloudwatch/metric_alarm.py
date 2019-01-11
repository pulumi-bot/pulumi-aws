# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class MetricAlarm(pulumi.CustomResource):
    """
    Provides a CloudWatch Metric Alarm resource.
    """
    def __init__(__self__, __name__, __opts__=None, actions_enabled=None, alarm_actions=None, alarm_description=None, name=None, comparison_operator=None, datapoints_to_alarm=None, dimensions=None, evaluate_low_sample_count_percentiles=None, evaluation_periods=None, extended_statistic=None, insufficient_data_actions=None, metric_name=None, namespace=None, ok_actions=None, period=None, statistic=None, threshold=None, treat_missing_data=None, unit=None):
        """Create a MetricAlarm resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['actions_enabled'] = actions_enabled

        __props__['alarm_actions'] = alarm_actions

        __props__['alarm_description'] = alarm_description

        __props__['name'] = name

        if not comparison_operator:
            raise TypeError('Missing required property comparison_operator')
        __props__['comparison_operator'] = comparison_operator

        __props__['datapoints_to_alarm'] = datapoints_to_alarm

        __props__['dimensions'] = dimensions

        __props__['evaluate_low_sample_count_percentiles'] = evaluate_low_sample_count_percentiles

        if not evaluation_periods:
            raise TypeError('Missing required property evaluation_periods')
        __props__['evaluation_periods'] = evaluation_periods

        __props__['extended_statistic'] = extended_statistic

        __props__['insufficient_data_actions'] = insufficient_data_actions

        if not metric_name:
            raise TypeError('Missing required property metric_name')
        __props__['metric_name'] = metric_name

        if not namespace:
            raise TypeError('Missing required property namespace')
        __props__['namespace'] = namespace

        __props__['ok_actions'] = ok_actions

        if not period:
            raise TypeError('Missing required property period')
        __props__['period'] = period

        __props__['statistic'] = statistic

        if not threshold:
            raise TypeError('Missing required property threshold')
        __props__['threshold'] = threshold

        __props__['treat_missing_data'] = treat_missing_data

        __props__['unit'] = unit

        __props__['arn'] = None

        super(MetricAlarm, __self__).__init__(
            'aws:cloudwatch/metricAlarm:MetricAlarm',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

