# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Schedule']


class Schedule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 desired_capacity: Optional[pulumi.Input[float]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 max_size: Optional[pulumi.Input[float]] = None,
                 min_size: Optional[pulumi.Input[float]] = None,
                 recurrence: Optional[pulumi.Input[str]] = None,
                 scheduled_action_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AutoScaling Schedule resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foobar_group = aws.autoscaling.Group("foobarGroup",
            availability_zones=["us-west-2a"],
            max_size=1,
            min_size=1,
            health_check_grace_period=300,
            health_check_type="ELB",
            force_delete=True,
            termination_policies=["OldestInstance"])
        foobar_schedule = aws.autoscaling.Schedule("foobarSchedule",
            scheduled_action_name="foobar",
            min_size=0,
            max_size=1,
            desired_capacity=0,
            start_time="2016-12-11T18:00:00Z",
            end_time="2016-12-12T06:00:00Z",
            autoscaling_group_name=foobar_group.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscaling_group_name: The name or Amazon Resource Name (ARN) of the Auto Scaling group.
        :param pulumi.Input[float] desired_capacity: The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
        :param pulumi.Input[str] end_time: The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
               If you try to schedule your action in the past, Auto Scaling returns an error message.
        :param pulumi.Input[float] max_size: The maximum size for the Auto Scaling group. Default 0.
               Set to -1 if you don't want to change the maximum size at the scheduled time.
        :param pulumi.Input[float] min_size: The minimum size for the Auto Scaling group. Default 0.
               Set to -1 if you don't want to change the minimum size at the scheduled time.
        :param pulumi.Input[str] recurrence: The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
        :param pulumi.Input[str] scheduled_action_name: The name of this scaling action.
        :param pulumi.Input[str] start_time: The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
               If you try to schedule your action in the past, Auto Scaling returns an error message.
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

            if autoscaling_group_name is None:
                raise TypeError("Missing required property 'autoscaling_group_name'")
            __props__['autoscaling_group_name'] = autoscaling_group_name
            __props__['desired_capacity'] = desired_capacity
            __props__['end_time'] = end_time
            __props__['max_size'] = max_size
            __props__['min_size'] = min_size
            __props__['recurrence'] = recurrence
            if scheduled_action_name is None:
                raise TypeError("Missing required property 'scheduled_action_name'")
            __props__['scheduled_action_name'] = scheduled_action_name
            __props__['start_time'] = start_time
            __props__['arn'] = None
        super(Schedule, __self__).__init__(
            'aws:autoscaling/schedule:Schedule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            autoscaling_group_name: Optional[pulumi.Input[str]] = None,
            desired_capacity: Optional[pulumi.Input[float]] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            max_size: Optional[pulumi.Input[float]] = None,
            min_size: Optional[pulumi.Input[float]] = None,
            recurrence: Optional[pulumi.Input[str]] = None,
            scheduled_action_name: Optional[pulumi.Input[str]] = None,
            start_time: Optional[pulumi.Input[str]] = None) -> 'Schedule':
        """
        Get an existing Schedule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN assigned by AWS to the autoscaling schedule.
        :param pulumi.Input[str] autoscaling_group_name: The name or Amazon Resource Name (ARN) of the Auto Scaling group.
        :param pulumi.Input[float] desired_capacity: The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
        :param pulumi.Input[str] end_time: The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
               If you try to schedule your action in the past, Auto Scaling returns an error message.
        :param pulumi.Input[float] max_size: The maximum size for the Auto Scaling group. Default 0.
               Set to -1 if you don't want to change the maximum size at the scheduled time.
        :param pulumi.Input[float] min_size: The minimum size for the Auto Scaling group. Default 0.
               Set to -1 if you don't want to change the minimum size at the scheduled time.
        :param pulumi.Input[str] recurrence: The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
        :param pulumi.Input[str] scheduled_action_name: The name of this scaling action.
        :param pulumi.Input[str] start_time: The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
               If you try to schedule your action in the past, Auto Scaling returns an error message.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["autoscaling_group_name"] = autoscaling_group_name
        __props__["desired_capacity"] = desired_capacity
        __props__["end_time"] = end_time
        __props__["max_size"] = max_size
        __props__["min_size"] = min_size
        __props__["recurrence"] = recurrence
        __props__["scheduled_action_name"] = scheduled_action_name
        __props__["start_time"] = start_time
        return Schedule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN assigned by AWS to the autoscaling schedule.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoscalingGroupName")
    def autoscaling_group_name(self) -> str:
        """
        The name or Amazon Resource Name (ARN) of the Auto Scaling group.
        """
        return pulumi.get(self, "autoscaling_group_name")

    @property
    @pulumi.getter(name="desiredCapacity")
    def desired_capacity(self) -> float:
        """
        The number of EC2 instances that should be running in the group. Default 0.  Set to -1 if you don't want to change the desired capacity at the scheduled time.
        """
        return pulumi.get(self, "desired_capacity")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        The time for this action to end, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
        If you try to schedule your action in the past, Auto Scaling returns an error message.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> float:
        """
        The maximum size for the Auto Scaling group. Default 0.
        Set to -1 if you don't want to change the maximum size at the scheduled time.
        """
        return pulumi.get(self, "max_size")

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> float:
        """
        The minimum size for the Auto Scaling group. Default 0.
        Set to -1 if you don't want to change the minimum size at the scheduled time.
        """
        return pulumi.get(self, "min_size")

    @property
    @pulumi.getter
    def recurrence(self) -> str:
        """
        The time when recurring future actions will start. Start time is specified by the user following the Unix cron syntax format.
        """
        return pulumi.get(self, "recurrence")

    @property
    @pulumi.getter(name="scheduledActionName")
    def scheduled_action_name(self) -> str:
        """
        The name of this scaling action.
        """
        return pulumi.get(self, "scheduled_action_name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        The time for this action to start, in "YYYY-MM-DDThh:mm:ssZ" format in UTC/GMT only (for example, 2014-06-01T00:00:00Z ).
        If you try to schedule your action in the past, Auto Scaling returns an error message.
        """
        return pulumi.get(self, "start_time")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

