# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'LifecyclePolicyPolicyDetailsArgs',
    'LifecyclePolicyPolicyDetailsScheduleArgs',
    'LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs',
    'LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs',
]

@pulumi.input_type
class LifecyclePolicyPolicyDetailsArgs:
    def __init__(__self__, *,
                 resource_types: pulumi.Input[List[pulumi.Input[str]]],
                 schedules: pulumi.Input[List[pulumi.Input['LifecyclePolicyPolicyDetailsScheduleArgs']]],
                 target_tags: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(__self__, "resource_types", resource_types)
        pulumi.set(__self__, "schedules", schedules)
        pulumi.set(__self__, "target_tags", target_tags)

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        return pulumi.get(self, "resource_types")

    @resource_types.setter
    def resource_types(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "resource_types", value)

    @property
    @pulumi.getter
    def schedules(self) -> pulumi.Input[List[pulumi.Input['LifecyclePolicyPolicyDetailsScheduleArgs']]]:
        return pulumi.get(self, "schedules")

    @schedules.setter
    def schedules(self, value: pulumi.Input[List[pulumi.Input['LifecyclePolicyPolicyDetailsScheduleArgs']]]):
        pulumi.set(self, "schedules", value)

    @property
    @pulumi.getter(name="targetTags")
    def target_tags(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        return pulumi.get(self, "target_tags")

    @target_tags.setter
    def target_tags(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "target_tags", value)


@pulumi.input_type
class LifecyclePolicyPolicyDetailsScheduleArgs:
    def __init__(__self__, *,
                 create_rule: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs'],
                 name: pulumi.Input[str],
                 retain_rule: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs'],
                 copy_tags: Optional[pulumi.Input[bool]] = None,
                 tags_to_add: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "create_rule", create_rule)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "retain_rule", retain_rule)
        if copy_tags is not None:
            pulumi.set(__self__, "copy_tags", copy_tags)
        if tags_to_add is not None:
            pulumi.set(__self__, "tags_to_add", tags_to_add)

    @property
    @pulumi.getter(name="createRule")
    def create_rule(self) -> pulumi.Input['LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs']:
        return pulumi.get(self, "create_rule")

    @create_rule.setter
    def create_rule(self, value: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs']):
        pulumi.set(self, "create_rule", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retainRule")
    def retain_rule(self) -> pulumi.Input['LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs']:
        return pulumi.get(self, "retain_rule")

    @retain_rule.setter
    def retain_rule(self, value: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs']):
        pulumi.set(self, "retain_rule", value)

    @property
    @pulumi.getter(name="copyTags")
    def copy_tags(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "copy_tags")

    @copy_tags.setter
    def copy_tags(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "copy_tags", value)

    @property
    @pulumi.getter(name="tagsToAdd")
    def tags_to_add(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_to_add")

    @tags_to_add.setter
    def tags_to_add(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_to_add", value)


@pulumi.input_type
class LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs:
    def __init__(__self__, *,
                 interval: pulumi.Input[float],
                 interval_unit: Optional[pulumi.Input[str]] = None,
                 times: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "interval", interval)
        if interval_unit is not None:
            pulumi.set(__self__, "interval_unit", interval_unit)
        if times is not None:
            pulumi.set(__self__, "times", times)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[float]:
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: pulumi.Input[float]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "interval_unit")

    @interval_unit.setter
    def interval_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interval_unit", value)

    @property
    @pulumi.getter
    def times(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "times")

    @times.setter
    def times(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "times", value)


@pulumi.input_type
class LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float]):
        pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: pulumi.Input[float]):
        pulumi.set(self, "count", value)


