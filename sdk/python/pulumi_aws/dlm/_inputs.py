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
        """
        :param pulumi.Input[List[pulumi.Input[str]]] resource_types: A list of resource types that should be targeted by the lifecycle policy. `VOLUME` is currently the only allowed value.
        :param pulumi.Input[List[pulumi.Input['LifecyclePolicyPolicyDetailsScheduleArgs']]] schedules: See the `schedule` configuration block.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] target_tags: A map of tag keys and their values. Any resources that match the `resource_types` and are tagged with _any_ of these tags will be targeted.
        """
        pulumi.set(__self__, "resourceTypes", resource_types)
        pulumi.set(__self__, "schedules", schedules)
        pulumi.set(__self__, "targetTags", target_tags)

    @property
    @pulumi.getter(name="resourceTypes")
    def resource_types(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of resource types that should be targeted by the lifecycle policy. `VOLUME` is currently the only allowed value.
        """
        ...

    @resource_types.setter
    def resource_types(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        ...

    @property
    @pulumi.getter
    def schedules(self) -> pulumi.Input[List[pulumi.Input['LifecyclePolicyPolicyDetailsScheduleArgs']]]:
        """
        See the `schedule` configuration block.
        """
        ...

    @schedules.setter
    def schedules(self, value: pulumi.Input[List[pulumi.Input['LifecyclePolicyPolicyDetailsScheduleArgs']]]):
        ...

    @property
    @pulumi.getter(name="targetTags")
    def target_tags(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        A map of tag keys and their values. Any resources that match the `resource_types` and are tagged with _any_ of these tags will be targeted.
        """
        ...

    @target_tags.setter
    def target_tags(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        ...


@pulumi.input_type
class LifecyclePolicyPolicyDetailsScheduleArgs:
    def __init__(__self__, *,
                 create_rule: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs'],
                 name: pulumi.Input[str],
                 retain_rule: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs'],
                 copy_tags: Optional[pulumi.Input[bool]] = None,
                 tags_to_add: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input['LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs'] create_rule: See the `create_rule` block. Max of 1 per schedule.
        :param pulumi.Input[str] name: A name for the schedule.
        :param pulumi.Input['LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs'] retain_rule: See the `retain_rule` block. Max of 1 per schedule.
        :param pulumi.Input[bool] copy_tags: Copy all user-defined tags on a source volume to snapshots of the volume created by this policy.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_to_add: A map of tag keys and their values. DLM lifecycle policies will already tag the snapshot with the tags on the volume. This configuration adds extra tags on top of these.
        """
        pulumi.set(__self__, "createRule", create_rule)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "retainRule", retain_rule)
        pulumi.set(__self__, "copyTags", copy_tags)
        pulumi.set(__self__, "tagsToAdd", tags_to_add)

    @property
    @pulumi.getter(name="createRule")
    def create_rule(self) -> pulumi.Input['LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs']:
        """
        See the `create_rule` block. Max of 1 per schedule.
        """
        ...

    @create_rule.setter
    def create_rule(self, value: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs']):
        ...

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        A name for the schedule.
        """
        ...

    @name.setter
    def name(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="retainRule")
    def retain_rule(self) -> pulumi.Input['LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs']:
        """
        See the `retain_rule` block. Max of 1 per schedule.
        """
        ...

    @retain_rule.setter
    def retain_rule(self, value: pulumi.Input['LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs']):
        ...

    @property
    @pulumi.getter(name="copyTags")
    def copy_tags(self) -> Optional[pulumi.Input[bool]]:
        """
        Copy all user-defined tags on a source volume to snapshots of the volume created by this policy.
        """
        ...

    @copy_tags.setter
    def copy_tags(self, value: Optional[pulumi.Input[bool]]):
        ...

    @property
    @pulumi.getter(name="tagsToAdd")
    def tags_to_add(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tag keys and their values. DLM lifecycle policies will already tag the snapshot with the tags on the volume. This configuration adds extra tags on top of these.
        """
        ...

    @tags_to_add.setter
    def tags_to_add(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs:
    def __init__(__self__, *,
                 interval: pulumi.Input[float],
                 interval_unit: Optional[pulumi.Input[str]] = None,
                 times: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[float] interval: How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values.
        :param pulumi.Input[str] interval_unit: The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value.
        :param pulumi.Input[str] times: A list of times in 24 hour clock format that sets when the lifecycle policy should be evaluated. Max of 1.
        """
        pulumi.set(__self__, "interval", interval)
        pulumi.set(__self__, "intervalUnit", interval_unit)
        pulumi.set(__self__, "times", times)

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Input[float]:
        """
        How often this lifecycle policy should be evaluated. `1`, `2`,`3`,`4`,`6`,`8`,`12` or `24` are valid values.
        """
        ...

    @interval.setter
    def interval(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="intervalUnit")
    def interval_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The unit for how often the lifecycle policy should be evaluated. `HOURS` is currently the only allowed value and also the default value.
        """
        ...

    @interval_unit.setter
    def interval_unit(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def times(self) -> Optional[pulumi.Input[str]]:
        """
        A list of times in 24 hour clock format that sets when the lifecycle policy should be evaluated. Max of 1.
        """
        ...

    @times.setter
    def times(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs:
    def __init__(__self__, *,
                 count: pulumi.Input[float]):
        """
        :param pulumi.Input[float] count: How many snapshots to keep. Must be an integer between 1 and 1000.
        """
        pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter
    def count(self) -> pulumi.Input[float]:
        """
        How many snapshots to keep. Must be an integer between 1 and 1000.
        """
        ...

    @count.setter
    def count(self, value: pulumi.Input[float]):
        ...


