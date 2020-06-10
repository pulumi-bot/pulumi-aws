# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
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
    arn: pulumi.Output[str]
    """
    The ARN of the cloudwatch metric alarm.
    """
    comparison_operator: pulumi.Output[str]
    """
    The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
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

      * `expression` (`str`) - The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the id of the other metrics to refer to those metrics, and can also use the id of other expressions to use the result of those expressions. For more information about metric math expressions, see Metric Math Syntax and Functions in the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax).
      * `id` (`str`) - A short name used to tie this object to the results in the response. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
      * `label` (`str`) - A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents.
      * `metric` (`dict`) - The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
        * `dimensions` (`dict`) - The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        * `metric_name` (`str`) - The name for this metric.
          See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        * `namespace` (`str`) - The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
          See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        * `period` (`float`) - The period in seconds over which the specified `stat` is applied.
        * `stat` (`str`) - The statistic to apply to this metric.
          Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        * `unit` (`str`) - The unit for this metric.

      * `returnData` (`bool`) - Specify exactly one `metric_query` to be `true` to use that `metric_query` result as the alarm.
    """
    name: pulumi.Output[str]
    """
    The descriptive name for the alarm. This name must be unique within the user's AWS account
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
    A map of tags to assign to the resource.
    """
    threshold: pulumi.Output[float]
    """
    The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
    """
    threshold_metric_id: pulumi.Output[str]
    """
    If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
    """
    treat_missing_data: pulumi.Output[str]
    """
    Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
    """
    unit: pulumi.Output[str]
    """
    The unit for this metric.
    """
    def __init__(__self__, resource_name, opts=None, actions_enabled=None, alarm_actions=None, alarm_description=None, comparison_operator=None, datapoints_to_alarm=None, dimensions=None, evaluate_low_sample_count_percentiles=None, evaluation_periods=None, extended_statistic=None, insufficient_data_actions=None, metric_name=None, metric_queries=None, name=None, namespace=None, ok_actions=None, period=None, statistic=None, tags=None, threshold=None, threshold_metric_id=None, treat_missing_data=None, unit=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CloudWatch Metric Alarm resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        foobar = aws.cloudwatch.MetricAlarm("foobar",
            alarm_description="This metric monitors ec2 cpu utilization",
            comparison_operator="GreaterThanOrEqualToThreshold",
            evaluation_periods="2",
            insufficient_data_actions=[],
            metric_name="CPUUtilization",
            namespace="AWS/EC2",
            period="120",
            statistic="Average",
            threshold="80")
        ```

        ## Example in Conjunction with Scaling Policies

        ```python
        import pulumi
        import pulumi_aws as aws

        bat_policy = aws.autoscaling.Policy("batPolicy",
            adjustment_type="ChangeInCapacity",
            autoscaling_group_name=aws_autoscaling_group["bar"]["name"],
            cooldown=300,
            scaling_adjustment=4)
        bat_metric_alarm = aws.cloudwatch.MetricAlarm("batMetricAlarm",
            alarm_actions=[bat_policy.arn],
            alarm_description="This metric monitors ec2 cpu utilization",
            comparison_operator="GreaterThanOrEqualToThreshold",
            dimensions={
                "AutoScalingGroupName": aws_autoscaling_group["bar"]["name"],
            },
            evaluation_periods="2",
            metric_name="CPUUtilization",
            namespace="AWS/EC2",
            period="120",
            statistic="Average",
            threshold="80")
        ```

        ## Example with an Expression

        ```python
        import pulumi
        import pulumi_aws as aws

        foobar = aws.cloudwatch.MetricAlarm("foobar",
            alarm_description="Request error rate has exceeded 10%",
            comparison_operator="GreaterThanOrEqualToThreshold",
            evaluation_periods="2",
            insufficient_data_actions=[],
            metric_queries=[
                {
                    "expression": "m2/m1*100",
                    "id": "e1",
                    "label": "Error Rate",
                    "returnData": "true",
                },
                {
                    "id": "m1",
                    "metric": {
                        "dimensions": {
                            "LoadBalancer": "app/web",
                        },
                        "metric_name": "RequestCount",
                        "namespace": "AWS/ApplicationELB",
                        "period": "120",
                        "stat": "Sum",
                        "unit": "Count",
                    },
                },
                {
                    "id": "m2",
                    "metric": {
                        "dimensions": {
                            "LoadBalancer": "app/web",
                        },
                        "metric_name": "HTTPCode_ELB_5XX_Count",
                        "namespace": "AWS/ApplicationELB",
                        "period": "120",
                        "stat": "Sum",
                        "unit": "Count",
                    },
                },
            ],
            threshold="10")
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        xx_anomaly_detection = aws.cloudwatch.MetricAlarm("xxAnomalyDetection",
            alarm_description="This metric monitors ec2 cpu utilization",
            comparison_operator="GreaterThanUpperThreshold",
            evaluation_periods="2",
            insufficient_data_actions=[],
            metric_queries=[
                {
                    "expression": "ANOMALY_DETECTION_BAND(m1)",
                    "id": "e1",
                    "label": "CPUUtilization (Expected)",
                    "returnData": "true",
                },
                {
                    "id": "m1",
                    "metric": {
                        "dimensions": {
                            "InstanceId": "i-abc123",
                        },
                        "metric_name": "CPUUtilization",
                        "namespace": "AWS/EC2",
                        "period": "120",
                        "stat": "Average",
                        "unit": "Count",
                    },
                    "returnData": "true",
                },
            ],
            threshold_metric_id="e1")
        ```

        ## Example of monitoring Healthy Hosts on NLB using Target Group and NLB

        ```python
        import pulumi
        import pulumi_aws as aws

        xxx_nlb_healthyhosts = aws.cloudwatch.MetricAlarm("xxxNlbHealthyhosts",
            comparison_operator="LessThanThreshold",
            evaluation_periods="1",
            metric_name="HealthyHostCount",
            namespace="AWS/NetworkELB",
            period="60",
            statistic="Average",
            threshold=var["logstash_servers_count"],
            alarm_description="Number of XXXX nodes healthy in Target Group",
            actions_enabled="true",
            alarm_actions=[aws_sns_topic["sns"]["arn"]],
            ok_actions=[aws_sns_topic["sns"]["arn"]],
            dimensions={
                "TargetGroup": aws_lb_target_group["lb-tg"]["arn_suffix"],
                "LoadBalancer": aws_lb["lb"]["arn_suffix"],
            })
        ```

        > **NOTE:**  You cannot create a metric alarm consisting of both `statistic` and `extended_statistic` parameters.
        You must choose one or the other

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
        :param pulumi.Input[list] alarm_actions: The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] alarm_description: The description for the alarm.
        :param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
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
        :param pulumi.Input[str] name: The descriptive name for the alarm. This name must be unique within the user's AWS account
        :param pulumi.Input[str] namespace: The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[list] ok_actions: The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[float] period: The period in seconds over which the specified `stat` is applied.
        :param pulumi.Input[str] statistic: The statistic to apply to the alarm's associated metric.
               Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] threshold: The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
        :param pulumi.Input[str] threshold_metric_id: If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
        :param pulumi.Input[str] treat_missing_data: Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
        :param pulumi.Input[str] unit: The unit for this metric.

        The **metric_queries** object supports the following:

          * `expression` (`pulumi.Input[str]`) - The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the id of the other metrics to refer to those metrics, and can also use the id of other expressions to use the result of those expressions. For more information about metric math expressions, see Metric Math Syntax and Functions in the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax).
          * `id` (`pulumi.Input[str]`) - A short name used to tie this object to the results in the response. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
          * `label` (`pulumi.Input[str]`) - A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents.
          * `metric` (`pulumi.Input[dict]`) - The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
            * `dimensions` (`pulumi.Input[dict]`) - The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
            * `metric_name` (`pulumi.Input[str]`) - The name for this metric.
              See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
            * `namespace` (`pulumi.Input[str]`) - The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
              See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
            * `period` (`pulumi.Input[float]`) - The period in seconds over which the specified `stat` is applied.
            * `stat` (`pulumi.Input[str]`) - The statistic to apply to this metric.
              Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
            * `unit` (`pulumi.Input[str]`) - The unit for this metric.

          * `returnData` (`pulumi.Input[bool]`) - Specify exactly one `metric_query` to be `true` to use that `metric_query` result as the alarm.
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
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['actions_enabled'] = actions_enabled
            __props__['alarm_actions'] = alarm_actions
            __props__['alarm_description'] = alarm_description
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
            __props__['name'] = name
            __props__['namespace'] = namespace
            __props__['ok_actions'] = ok_actions
            __props__['period'] = period
            __props__['statistic'] = statistic
            __props__['tags'] = tags
            __props__['threshold'] = threshold
            __props__['threshold_metric_id'] = threshold_metric_id
            __props__['treat_missing_data'] = treat_missing_data
            __props__['unit'] = unit
            __props__['arn'] = None
        super(MetricAlarm, __self__).__init__(
            'aws:cloudwatch/metricAlarm:MetricAlarm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, actions_enabled=None, alarm_actions=None, alarm_description=None, arn=None, comparison_operator=None, datapoints_to_alarm=None, dimensions=None, evaluate_low_sample_count_percentiles=None, evaluation_periods=None, extended_statistic=None, insufficient_data_actions=None, metric_name=None, metric_queries=None, name=None, namespace=None, ok_actions=None, period=None, statistic=None, tags=None, threshold=None, threshold_metric_id=None, treat_missing_data=None, unit=None):
        """
        Get an existing MetricAlarm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] actions_enabled: Indicates whether or not actions should be executed during any changes to the alarm's state. Defaults to `true`.
        :param pulumi.Input[list] alarm_actions: The list of actions to execute when this alarm transitions into an ALARM state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[str] alarm_description: The description for the alarm.
        :param pulumi.Input[str] arn: The ARN of the cloudwatch metric alarm.
        :param pulumi.Input[str] comparison_operator: The arithmetic operation to use when comparing the specified Statistic and Threshold. The specified Statistic value is used as the first operand. Either of the following is supported: `GreaterThanOrEqualToThreshold`, `GreaterThanThreshold`, `LessThanThreshold`, `LessThanOrEqualToThreshold`. Additionally, the values  `LessThanLowerOrGreaterThanUpperThreshold`, `LessThanLowerThreshold`, and `GreaterThanUpperThreshold` are used only for alarms based on anomaly detection models.
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
        :param pulumi.Input[str] name: The descriptive name for the alarm. This name must be unique within the user's AWS account
        :param pulumi.Input[str] namespace: The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
               See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
        :param pulumi.Input[list] ok_actions: The list of actions to execute when this alarm transitions into an OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
        :param pulumi.Input[float] period: The period in seconds over which the specified `stat` is applied.
        :param pulumi.Input[str] statistic: The statistic to apply to the alarm's associated metric.
               Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] threshold: The value against which the specified statistic is compared. This parameter is required for alarms based on static thresholds, but should not be used for alarms based on anomaly detection models.
        :param pulumi.Input[str] threshold_metric_id: If this is an alarm based on an anomaly detection model, make this value match the ID of the ANOMALY_DETECTION_BAND function.
        :param pulumi.Input[str] treat_missing_data: Sets how this alarm is to handle missing data points. The following values are supported: `missing`, `ignore`, `breaching` and `notBreaching`. Defaults to `missing`.
        :param pulumi.Input[str] unit: The unit for this metric.

        The **metric_queries** object supports the following:

          * `expression` (`pulumi.Input[str]`) - The math expression to be performed on the returned data, if this object is performing a math expression. This expression can use the id of the other metrics to refer to those metrics, and can also use the id of other expressions to use the result of those expressions. For more information about metric math expressions, see Metric Math Syntax and Functions in the [Amazon CloudWatch User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html#metric-math-syntax).
          * `id` (`pulumi.Input[str]`) - A short name used to tie this object to the results in the response. If you are performing math expressions on this set of data, this name represents that data and can serve as a variable in the mathematical expression. The valid characters are letters, numbers, and underscore. The first character must be a lowercase letter.
          * `label` (`pulumi.Input[str]`) - A human-readable label for this metric or expression. This is especially useful if this is an expression, so that you know what the value represents.
          * `metric` (`pulumi.Input[dict]`) - The metric to be returned, along with statistics, period, and units. Use this parameter only if this object is retrieving a metric and not performing a math expression on returned data.
            * `dimensions` (`pulumi.Input[dict]`) - The dimensions for this metric.  For the list of available dimensions see the AWS documentation [here](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
            * `metric_name` (`pulumi.Input[str]`) - The name for this metric.
              See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
            * `namespace` (`pulumi.Input[str]`) - The namespace for this metric. See docs for the [list of namespaces](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/aws-namespaces.html).
              See docs for [supported metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CW_Support_For_AWS.html).
            * `period` (`pulumi.Input[float]`) - The period in seconds over which the specified `stat` is applied.
            * `stat` (`pulumi.Input[str]`) - The statistic to apply to this metric.
              Either of the following is supported: `SampleCount`, `Average`, `Sum`, `Minimum`, `Maximum`
            * `unit` (`pulumi.Input[str]`) - The unit for this metric.

          * `returnData` (`pulumi.Input[bool]`) - Specify exactly one `metric_query` to be `true` to use that `metric_query` result as the alarm.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions_enabled"] = actions_enabled
        __props__["alarm_actions"] = alarm_actions
        __props__["alarm_description"] = alarm_description
        __props__["arn"] = arn
        __props__["comparison_operator"] = comparison_operator
        __props__["datapoints_to_alarm"] = datapoints_to_alarm
        __props__["dimensions"] = dimensions
        __props__["evaluate_low_sample_count_percentiles"] = evaluate_low_sample_count_percentiles
        __props__["evaluation_periods"] = evaluation_periods
        __props__["extended_statistic"] = extended_statistic
        __props__["insufficient_data_actions"] = insufficient_data_actions
        __props__["metric_name"] = metric_name
        __props__["metric_queries"] = metric_queries
        __props__["name"] = name
        __props__["namespace"] = namespace
        __props__["ok_actions"] = ok_actions
        __props__["period"] = period
        __props__["statistic"] = statistic
        __props__["tags"] = tags
        __props__["threshold"] = threshold
        __props__["threshold_metric_id"] = threshold_metric_id
        __props__["treat_missing_data"] = treat_missing_data
        __props__["unit"] = unit
        return MetricAlarm(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
