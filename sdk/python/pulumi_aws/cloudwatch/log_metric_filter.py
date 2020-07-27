# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class LogMetricFilter(pulumi.CustomResource):
    log_group_name: pulumi.Output[str]
    """
    The name of the log group to associate the metric filter with.
    """
    metric_transformation: pulumi.Output[dict]
    """
    A block defining collection of information
    needed to define how metric data gets emitted. See below.

      * `default_value` (`str`) - The value to emit when a filter pattern does not match a log event.
      * `name` (`str`) - The name of the CloudWatch metric to which the monitored log information should be published (e.g. `ErrorCount`)
      * `namespace` (`str`) - The destination namespace of the CloudWatch metric.
      * `value` (`str`) - What to publish to the metric. For example, if you're counting the occurrences of a particular term like "Error", the value will be "1" for each occurrence. If you're counting the bytes transferred the published value will be the value in the log event.
    """
    name: pulumi.Output[str]
    """
    A name for the metric filter.
    """
    pattern: pulumi.Output[str]
    """
    A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
    for extracting metric data out of ingested log events.
    """
    def __init__(__self__, resource_name, opts=None, log_group_name=None, metric_transformation=None, name=None, pattern=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CloudWatch Log Metric Filter resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        dada = aws.cloudwatch.LogGroup("dada")
        yada = aws.cloudwatch.LogMetricFilter("yada",
            log_group_name=dada.name,
            metric_transformation={
                "name": "EventCount",
                "namespace": "YourNamespace",
                "value": "1",
            },
            pattern="")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] log_group_name: The name of the log group to associate the metric filter with.
        :param pulumi.Input[dict] metric_transformation: A block defining collection of information
               needed to define how metric data gets emitted. See below.
        :param pulumi.Input[str] name: A name for the metric filter.
        :param pulumi.Input[str] pattern: A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
               for extracting metric data out of ingested log events.

        The **metric_transformation** object supports the following:

          * `default_value` (`pulumi.Input[str]`) - The value to emit when a filter pattern does not match a log event.
          * `name` (`pulumi.Input[str]`) - The name of the CloudWatch metric to which the monitored log information should be published (e.g. `ErrorCount`)
          * `namespace` (`pulumi.Input[str]`) - The destination namespace of the CloudWatch metric.
          * `value` (`pulumi.Input[str]`) - What to publish to the metric. For example, if you're counting the occurrences of a particular term like "Error", the value will be "1" for each occurrence. If you're counting the bytes transferred the published value will be the value in the log event.
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

            if log_group_name is None:
                raise TypeError("Missing required property 'log_group_name'")
            __props__['log_group_name'] = log_group_name
            if metric_transformation is None:
                raise TypeError("Missing required property 'metric_transformation'")
            __props__['metric_transformation'] = metric_transformation
            __props__['name'] = name
            if pattern is None:
                raise TypeError("Missing required property 'pattern'")
            __props__['pattern'] = pattern
        super(LogMetricFilter, __self__).__init__(
            'aws:cloudwatch/logMetricFilter:LogMetricFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, log_group_name=None, metric_transformation=None, name=None, pattern=None):
        """
        Get an existing LogMetricFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] log_group_name: The name of the log group to associate the metric filter with.
        :param pulumi.Input[dict] metric_transformation: A block defining collection of information
               needed to define how metric data gets emitted. See below.
        :param pulumi.Input[str] name: A name for the metric filter.
        :param pulumi.Input[str] pattern: A valid [CloudWatch Logs filter pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/FilterAndPatternSyntax.html)
               for extracting metric data out of ingested log events.

        The **metric_transformation** object supports the following:

          * `default_value` (`pulumi.Input[str]`) - The value to emit when a filter pattern does not match a log event.
          * `name` (`pulumi.Input[str]`) - The name of the CloudWatch metric to which the monitored log information should be published (e.g. `ErrorCount`)
          * `namespace` (`pulumi.Input[str]`) - The destination namespace of the CloudWatch metric.
          * `value` (`pulumi.Input[str]`) - What to publish to the metric. For example, if you're counting the occurrences of a particular term like "Error", the value will be "1" for each occurrence. If you're counting the bytes transferred the published value will be the value in the log event.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["log_group_name"] = log_group_name
        __props__["metric_transformation"] = metric_transformation
        __props__["name"] = name
        __props__["pattern"] = pattern
        return LogMetricFilter(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
