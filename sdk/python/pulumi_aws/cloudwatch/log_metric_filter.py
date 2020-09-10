# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['LogMetricFilter']


class LogMetricFilter(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 metric_transformation: Optional[pulumi.Input[pulumi.InputType['LogMetricFilterMetricTransformationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a LogMetricFilter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            log_group_name: Optional[pulumi.Input[str]] = None,
            metric_transformation: Optional[pulumi.Input[pulumi.InputType['LogMetricFilterMetricTransformationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pattern: Optional[pulumi.Input[str]] = None) -> 'LogMetricFilter':
        """
        Get an existing LogMetricFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["log_group_name"] = log_group_name
        __props__["metric_transformation"] = metric_transformation
        __props__["name"] = name
        __props__["pattern"] = pattern
        return LogMetricFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter(name="metricTransformation")
    def metric_transformation(self) -> pulumi.Output['outputs.LogMetricFilterMetricTransformation']:
        return pulumi.get(self, "metric_transformation")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[str]:
        return pulumi.get(self, "pattern")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

