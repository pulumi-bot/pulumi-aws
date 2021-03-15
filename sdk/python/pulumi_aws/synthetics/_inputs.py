# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'CanaryRunConfigArgs',
    'CanaryScheduleArgs',
    'CanaryTimelineArgs',
    'CanaryVpcConfigArgs',
]

@pulumi.input_type
class CanaryRunConfigArgs:
    def __init__(__self__, *,
                 active_tracing: Optional[pulumi.Input[bool]] = None,
                 memory_in_mb: Optional[pulumi.Input[int]] = None,
                 timeout_in_seconds: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[bool] active_tracing: Whether this canary is to use active AWS X-Ray tracing when it runs. You can enable active tracing only for canaries that use version syn-nodejs-2.0 or later for their canary runtime.
        :param pulumi.Input[int] memory_in_mb: Maximum amount of memory available to the canary while it is running, in MB. The value you specify must be a multiple of 64.
        :param pulumi.Input[int] timeout_in_seconds: Number of seconds the canary is allowed to run before it must stop. If you omit this field, the frequency of the canary is used, up to a maximum of 840 (14 minutes).
        """
        if active_tracing is not None:
            pulumi.set(__self__, "active_tracing", active_tracing)
        if memory_in_mb is not None:
            pulumi.set(__self__, "memory_in_mb", memory_in_mb)
        if timeout_in_seconds is not None:
            pulumi.set(__self__, "timeout_in_seconds", timeout_in_seconds)

    @property
    @pulumi.getter(name="activeTracing")
    def active_tracing(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this canary is to use active AWS X-Ray tracing when it runs. You can enable active tracing only for canaries that use version syn-nodejs-2.0 or later for their canary runtime.
        """
        return pulumi.get(self, "active_tracing")

    @active_tracing.setter
    def active_tracing(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active_tracing", value)

    @property
    @pulumi.getter(name="memoryInMb")
    def memory_in_mb(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum amount of memory available to the canary while it is running, in MB. The value you specify must be a multiple of 64.
        """
        return pulumi.get(self, "memory_in_mb")

    @memory_in_mb.setter
    def memory_in_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_in_mb", value)

    @property
    @pulumi.getter(name="timeoutInSeconds")
    def timeout_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Number of seconds the canary is allowed to run before it must stop. If you omit this field, the frequency of the canary is used, up to a maximum of 840 (14 minutes).
        """
        return pulumi.get(self, "timeout_in_seconds")

    @timeout_in_seconds.setter
    def timeout_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "timeout_in_seconds", value)


@pulumi.input_type
class CanaryScheduleArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 duration_in_seconds: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] expression: Rate expression that defines how often the canary is to run. The syntax is rate(number unit). unit can be minute, minutes, or hour.
        :param pulumi.Input[int] duration_in_seconds: Duration in seconds, for the canary to continue making regular runs according to the schedule in the Expression value.
        """
        pulumi.set(__self__, "expression", expression)
        if duration_in_seconds is not None:
            pulumi.set(__self__, "duration_in_seconds", duration_in_seconds)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        """
        Rate expression that defines how often the canary is to run. The syntax is rate(number unit). unit can be minute, minutes, or hour.
        """
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter(name="durationInSeconds")
    def duration_in_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Duration in seconds, for the canary to continue making regular runs according to the schedule in the Expression value.
        """
        return pulumi.get(self, "duration_in_seconds")

    @duration_in_seconds.setter
    def duration_in_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration_in_seconds", value)


@pulumi.input_type
class CanaryTimelineArgs:
    def __init__(__self__, *,
                 created: Optional[pulumi.Input[str]] = None,
                 last_modified: Optional[pulumi.Input[str]] = None,
                 last_started: Optional[pulumi.Input[str]] = None,
                 last_stopped: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] created: Date and time the canary was created.
        :param pulumi.Input[str] last_modified: Date and time the canary was most recently modified.
        :param pulumi.Input[str] last_started: Date and time that the canary's most recent run started.
        :param pulumi.Input[str] last_stopped: Date and time that the canary's most recent run ended.
        """
        if created is not None:
            pulumi.set(__self__, "created", created)
        if last_modified is not None:
            pulumi.set(__self__, "last_modified", last_modified)
        if last_started is not None:
            pulumi.set(__self__, "last_started", last_started)
        if last_stopped is not None:
            pulumi.set(__self__, "last_stopped", last_stopped)

    @property
    @pulumi.getter
    def created(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time the canary was created.
        """
        return pulumi.get(self, "created")

    @created.setter
    def created(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created", value)

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time the canary was most recently modified.
        """
        return pulumi.get(self, "last_modified")

    @last_modified.setter
    def last_modified(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_modified", value)

    @property
    @pulumi.getter(name="lastStarted")
    def last_started(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time that the canary's most recent run started.
        """
        return pulumi.get(self, "last_started")

    @last_started.setter
    def last_started(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_started", value)

    @property
    @pulumi.getter(name="lastStopped")
    def last_stopped(self) -> Optional[pulumi.Input[str]]:
        """
        Date and time that the canary's most recent run ended.
        """
        return pulumi.get(self, "last_stopped")

    @last_stopped.setter
    def last_stopped(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_stopped", value)


@pulumi.input_type
class CanaryVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: IDs of the security groups for this canary.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: IDs of the subnets where this canary is to run.
        :param pulumi.Input[str] vpc_id: ID of the VPC where this canary is to run.
        """
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of the security groups for this canary.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IDs of the subnets where this canary is to run.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the VPC where this canary is to run.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


