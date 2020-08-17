# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GroupInitialLifecycleHook',
    'GroupLaunchTemplate',
    'GroupMixedInstancesPolicy',
    'GroupMixedInstancesPolicyInstancesDistribution',
    'GroupMixedInstancesPolicyLaunchTemplate',
    'GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification',
    'GroupMixedInstancesPolicyLaunchTemplateOverride',
    'GroupTag',
    'PolicyStepAdjustment',
    'PolicyTargetTrackingConfiguration',
    'PolicyTargetTrackingConfigurationCustomizedMetricSpecification',
    'PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension',
    'PolicyTargetTrackingConfigurationPredefinedMetricSpecification',
]

@pulumi.output_type
class GroupInitialLifecycleHook(dict):
    def __init__(__self__, *,
                 lifecycle_transition: str,
                 name: str,
                 default_result: Optional[str] = None,
                 heartbeat_timeout: Optional[float] = None,
                 notification_metadata: Optional[str] = None,
                 notification_target_arn: Optional[str] = None,
                 role_arn: Optional[str] = None):
        """
        :param str name: The name of the auto scaling group. By default generated by this provider.
        """
        pulumi.set(__self__, "lifecycle_transition", lifecycle_transition)
        pulumi.set(__self__, "name", name)
        if default_result is not None:
            pulumi.set(__self__, "default_result", default_result)
        if heartbeat_timeout is not None:
            pulumi.set(__self__, "heartbeat_timeout", heartbeat_timeout)
        if notification_metadata is not None:
            pulumi.set(__self__, "notification_metadata", notification_metadata)
        if notification_target_arn is not None:
            pulumi.set(__self__, "notification_target_arn", notification_target_arn)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> str:
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the auto scaling group. By default generated by this provider.
        """
        ...

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> Optional[float]:
        ...

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="notificationTargetArn")
    def notification_target_arn(self) -> Optional[str]:
        ...

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GroupLaunchTemplate(dict):
    def __init__(__self__, *,
                 id: Optional[str] = None,
                 name: Optional[str] = None,
                 version: Optional[str] = None):
        """
        :param str id: The ID of the launch template. Conflicts with `name`.
        :param str name: The name of the auto scaling group. By default generated by this provider.
        :param str version: Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
        """
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of the launch template. Conflicts with `name`.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the auto scaling group. By default generated by this provider.
        """
        ...

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GroupMixedInstancesPolicy(dict):
    def __init__(__self__, *,
                 launch_template: 'outputs.GroupMixedInstancesPolicyLaunchTemplate',
                 instances_distribution: Optional['outputs.GroupMixedInstancesPolicyInstancesDistribution'] = None):
        """
        :param 'GroupMixedInstancesPolicyLaunchTemplateArgs' launch_template: Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
        :param 'GroupMixedInstancesPolicyInstancesDistributionArgs' instances_distribution: Nested argument containing settings on how to mix on-demand and Spot instances in the Auto Scaling group. Defined below.
        """
        pulumi.set(__self__, "launch_template", launch_template)
        if instances_distribution is not None:
            pulumi.set(__self__, "instances_distribution", instances_distribution)

    @property
    @pulumi.getter(name="launchTemplate")
    def launch_template(self) -> 'outputs.GroupMixedInstancesPolicyLaunchTemplate':
        """
        Nested argument containing launch template settings along with the overrides to specify multiple instance types and weights. Defined below.
        """
        ...

    @property
    @pulumi.getter(name="instancesDistribution")
    def instances_distribution(self) -> Optional['outputs.GroupMixedInstancesPolicyInstancesDistribution']:
        """
        Nested argument containing settings on how to mix on-demand and Spot instances in the Auto Scaling group. Defined below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GroupMixedInstancesPolicyInstancesDistribution(dict):
    def __init__(__self__, *,
                 on_demand_allocation_strategy: Optional[str] = None,
                 on_demand_base_capacity: Optional[float] = None,
                 on_demand_percentage_above_base_capacity: Optional[float] = None,
                 spot_allocation_strategy: Optional[str] = None,
                 spot_instance_pools: Optional[float] = None,
                 spot_max_price: Optional[str] = None):
        """
        :param str on_demand_allocation_strategy: Strategy to use when launching on-demand instances. Valid values: `prioritized`. Default: `prioritized`.
        :param float on_demand_base_capacity: Absolute minimum amount of desired capacity that must be fulfilled by on-demand instances. Default: `0`.
        :param float on_demand_percentage_above_base_capacity: Percentage split between on-demand and Spot instances above the base on-demand capacity. Default: `100`.
        :param str spot_allocation_strategy: How to allocate capacity across the Spot pools. Valid values: `lowest-price`, `capacity-optimized`. Default: `lowest-price`.
        :param float spot_instance_pools: Number of Spot pools per availability zone to allocate capacity. EC2 Auto Scaling selects the cheapest Spot pools and evenly allocates Spot capacity across the number of Spot pools that you specify. Default: `2`.
        :param str spot_max_price: Maximum price per unit hour that the user is willing to pay for the Spot instances. Default: an empty string which means the on-demand price.
        """
        if on_demand_allocation_strategy is not None:
            pulumi.set(__self__, "on_demand_allocation_strategy", on_demand_allocation_strategy)
        if on_demand_base_capacity is not None:
            pulumi.set(__self__, "on_demand_base_capacity", on_demand_base_capacity)
        if on_demand_percentage_above_base_capacity is not None:
            pulumi.set(__self__, "on_demand_percentage_above_base_capacity", on_demand_percentage_above_base_capacity)
        if spot_allocation_strategy is not None:
            pulumi.set(__self__, "spot_allocation_strategy", spot_allocation_strategy)
        if spot_instance_pools is not None:
            pulumi.set(__self__, "spot_instance_pools", spot_instance_pools)
        if spot_max_price is not None:
            pulumi.set(__self__, "spot_max_price", spot_max_price)

    @property
    @pulumi.getter(name="onDemandAllocationStrategy")
    def on_demand_allocation_strategy(self) -> Optional[str]:
        """
        Strategy to use when launching on-demand instances. Valid values: `prioritized`. Default: `prioritized`.
        """
        ...

    @property
    @pulumi.getter(name="onDemandBaseCapacity")
    def on_demand_base_capacity(self) -> Optional[float]:
        """
        Absolute minimum amount of desired capacity that must be fulfilled by on-demand instances. Default: `0`.
        """
        ...

    @property
    @pulumi.getter(name="onDemandPercentageAboveBaseCapacity")
    def on_demand_percentage_above_base_capacity(self) -> Optional[float]:
        """
        Percentage split between on-demand and Spot instances above the base on-demand capacity. Default: `100`.
        """
        ...

    @property
    @pulumi.getter(name="spotAllocationStrategy")
    def spot_allocation_strategy(self) -> Optional[str]:
        """
        How to allocate capacity across the Spot pools. Valid values: `lowest-price`, `capacity-optimized`. Default: `lowest-price`.
        """
        ...

    @property
    @pulumi.getter(name="spotInstancePools")
    def spot_instance_pools(self) -> Optional[float]:
        """
        Number of Spot pools per availability zone to allocate capacity. EC2 Auto Scaling selects the cheapest Spot pools and evenly allocates Spot capacity across the number of Spot pools that you specify. Default: `2`.
        """
        ...

    @property
    @pulumi.getter(name="spotMaxPrice")
    def spot_max_price(self) -> Optional[str]:
        """
        Maximum price per unit hour that the user is willing to pay for the Spot instances. Default: an empty string which means the on-demand price.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GroupMixedInstancesPolicyLaunchTemplate(dict):
    def __init__(__self__, *,
                 launch_template_specification: 'outputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification',
                 overrides: Optional[List['outputs.GroupMixedInstancesPolicyLaunchTemplateOverride']] = None):
        """
        :param 'GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecificationArgs' launch_template_specification: Nested argument defines the Launch Template. Defined below.
        :param List['GroupMixedInstancesPolicyLaunchTemplateOverrideArgs'] overrides: List of nested arguments provides the ability to specify multiple instance types. This will override the same parameter in the launch template. For on-demand instances, Auto Scaling considers the order of preference of instance types to launch based on the order specified in the overrides list. Defined below.
        """
        pulumi.set(__self__, "launch_template_specification", launch_template_specification)
        if overrides is not None:
            pulumi.set(__self__, "overrides", overrides)

    @property
    @pulumi.getter(name="launchTemplateSpecification")
    def launch_template_specification(self) -> 'outputs.GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification':
        """
        Nested argument defines the Launch Template. Defined below.
        """
        ...

    @property
    @pulumi.getter
    def overrides(self) -> Optional[List['outputs.GroupMixedInstancesPolicyLaunchTemplateOverride']]:
        """
        List of nested arguments provides the ability to specify multiple instance types. This will override the same parameter in the launch template. For on-demand instances, Auto Scaling considers the order of preference of instance types to launch based on the order specified in the overrides list. Defined below.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GroupMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification(dict):
    def __init__(__self__, *,
                 launch_template_id: Optional[str] = None,
                 launch_template_name: Optional[str] = None,
                 version: Optional[str] = None):
        """
        :param str launch_template_id: The ID of the launch template. Conflicts with `launch_template_name`.
        :param str launch_template_name: The name of the launch template. Conflicts with `launch_template_id`.
        :param str version: Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
        """
        if launch_template_id is not None:
            pulumi.set(__self__, "launch_template_id", launch_template_id)
        if launch_template_name is not None:
            pulumi.set(__self__, "launch_template_name", launch_template_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="launchTemplateId")
    def launch_template_id(self) -> Optional[str]:
        """
        The ID of the launch template. Conflicts with `launch_template_name`.
        """
        ...

    @property
    @pulumi.getter(name="launchTemplateName")
    def launch_template_name(self) -> Optional[str]:
        """
        The name of the launch template. Conflicts with `launch_template_id`.
        """
        ...

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GroupMixedInstancesPolicyLaunchTemplateOverride(dict):
    def __init__(__self__, *,
                 instance_type: Optional[str] = None,
                 weighted_capacity: Optional[str] = None):
        """
        :param str instance_type: Override the instance type in the Launch Template.
        :param str weighted_capacity: The number of capacity units, which gives the instance type a proportional weight to other instance types.
        """
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if weighted_capacity is not None:
            pulumi.set(__self__, "weighted_capacity", weighted_capacity)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[str]:
        """
        Override the instance type in the Launch Template.
        """
        ...

    @property
    @pulumi.getter(name="weightedCapacity")
    def weighted_capacity(self) -> Optional[str]:
        """
        The number of capacity units, which gives the instance type a proportional weight to other instance types.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GroupTag(dict):
    def __init__(__self__, *,
                 key: str,
                 propagate_at_launch: bool,
                 value: str):
        """
        :param str key: Key
        :param bool propagate_at_launch: Enables propagation of the tag to
               Amazon EC2 instances launched via this ASG
        :param str value: Value
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "propagate_at_launch", propagate_at_launch)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Key
        """
        ...

    @property
    @pulumi.getter(name="propagateAtLaunch")
    def propagate_at_launch(self) -> bool:
        """
        Enables propagation of the tag to
        Amazon EC2 instances launched via this ASG
        """
        ...

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Value
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyStepAdjustment(dict):
    def __init__(__self__, *,
                 scaling_adjustment: float,
                 metric_interval_lower_bound: Optional[str] = None,
                 metric_interval_upper_bound: Optional[str] = None):
        """
        :param float scaling_adjustment: The number of members by which to
               scale, when the adjustment bounds are breached. A positive value scales
               up. A negative value scales down.
        :param str metric_interval_lower_bound: The lower bound for the
               difference between the alarm threshold and the CloudWatch metric.
               Without a value, AWS will treat this bound as infinity.
        :param str metric_interval_upper_bound: The upper bound for the
               difference between the alarm threshold and the CloudWatch metric.
               Without a value, AWS will treat this bound as infinity. The upper bound
               must be greater than the lower bound.
        """
        pulumi.set(__self__, "scaling_adjustment", scaling_adjustment)
        if metric_interval_lower_bound is not None:
            pulumi.set(__self__, "metric_interval_lower_bound", metric_interval_lower_bound)
        if metric_interval_upper_bound is not None:
            pulumi.set(__self__, "metric_interval_upper_bound", metric_interval_upper_bound)

    @property
    @pulumi.getter(name="scalingAdjustment")
    def scaling_adjustment(self) -> float:
        """
        The number of members by which to
        scale, when the adjustment bounds are breached. A positive value scales
        up. A negative value scales down.
        """
        ...

    @property
    @pulumi.getter(name="metricIntervalLowerBound")
    def metric_interval_lower_bound(self) -> Optional[str]:
        """
        The lower bound for the
        difference between the alarm threshold and the CloudWatch metric.
        Without a value, AWS will treat this bound as infinity.
        """
        ...

    @property
    @pulumi.getter(name="metricIntervalUpperBound")
    def metric_interval_upper_bound(self) -> Optional[str]:
        """
        The upper bound for the
        difference between the alarm threshold and the CloudWatch metric.
        Without a value, AWS will treat this bound as infinity. The upper bound
        must be greater than the lower bound.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyTargetTrackingConfiguration(dict):
    def __init__(__self__, *,
                 target_value: float,
                 customized_metric_specification: Optional['outputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecification'] = None,
                 disable_scale_in: Optional[bool] = None,
                 predefined_metric_specification: Optional['outputs.PolicyTargetTrackingConfigurationPredefinedMetricSpecification'] = None):
        """
        :param float target_value: The target value for the metric.
        :param 'PolicyTargetTrackingConfigurationCustomizedMetricSpecificationArgs' customized_metric_specification: A customized metric. Conflicts with `predefined_metric_specification`.
        :param bool disable_scale_in: Indicates whether scale in by the target tracking policy is disabled.
        :param 'PolicyTargetTrackingConfigurationPredefinedMetricSpecificationArgs' predefined_metric_specification: A predefined metric. Conflicts with `customized_metric_specification`.
        """
        pulumi.set(__self__, "target_value", target_value)
        if customized_metric_specification is not None:
            pulumi.set(__self__, "customized_metric_specification", customized_metric_specification)
        if disable_scale_in is not None:
            pulumi.set(__self__, "disable_scale_in", disable_scale_in)
        if predefined_metric_specification is not None:
            pulumi.set(__self__, "predefined_metric_specification", predefined_metric_specification)

    @property
    @pulumi.getter(name="targetValue")
    def target_value(self) -> float:
        """
        The target value for the metric.
        """
        ...

    @property
    @pulumi.getter(name="customizedMetricSpecification")
    def customized_metric_specification(self) -> Optional['outputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecification']:
        """
        A customized metric. Conflicts with `predefined_metric_specification`.
        """
        ...

    @property
    @pulumi.getter(name="disableScaleIn")
    def disable_scale_in(self) -> Optional[bool]:
        """
        Indicates whether scale in by the target tracking policy is disabled.
        """
        ...

    @property
    @pulumi.getter(name="predefinedMetricSpecification")
    def predefined_metric_specification(self) -> Optional['outputs.PolicyTargetTrackingConfigurationPredefinedMetricSpecification']:
        """
        A predefined metric. Conflicts with `customized_metric_specification`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyTargetTrackingConfigurationCustomizedMetricSpecification(dict):
    def __init__(__self__, *,
                 metric_name: str,
                 namespace: str,
                 statistic: str,
                 metric_dimensions: Optional[List['outputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension']] = None,
                 unit: Optional[str] = None):
        """
        :param str metric_name: The name of the metric.
        :param str namespace: The namespace of the metric.
        :param str statistic: The statistic of the metric.
        :param List['PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimensionArgs'] metric_dimensions: The dimensions of the metric.
        :param str unit: The unit of the metric.
        """
        pulumi.set(__self__, "metric_name", metric_name)
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "statistic", statistic)
        if metric_dimensions is not None:
            pulumi.set(__self__, "metric_dimensions", metric_dimensions)
        if unit is not None:
            pulumi.set(__self__, "unit", unit)

    @property
    @pulumi.getter(name="metricName")
    def metric_name(self) -> str:
        """
        The name of the metric.
        """
        ...

    @property
    @pulumi.getter
    def namespace(self) -> str:
        """
        The namespace of the metric.
        """
        ...

    @property
    @pulumi.getter
    def statistic(self) -> str:
        """
        The statistic of the metric.
        """
        ...

    @property
    @pulumi.getter(name="metricDimensions")
    def metric_dimensions(self) -> Optional[List['outputs.PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension']]:
        """
        The dimensions of the metric.
        """
        ...

    @property
    @pulumi.getter
    def unit(self) -> Optional[str]:
        """
        The unit of the metric.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyTargetTrackingConfigurationCustomizedMetricSpecificationMetricDimension(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        :param str name: The name of the dimension.
        :param str value: The value of the dimension.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the dimension.
        """
        ...

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value of the dimension.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyTargetTrackingConfigurationPredefinedMetricSpecification(dict):
    def __init__(__self__, *,
                 predefined_metric_type: str,
                 resource_label: Optional[str] = None):
        """
        :param str predefined_metric_type: The metric type.
        :param str resource_label: Identifies the resource associated with the metric type.
        """
        pulumi.set(__self__, "predefined_metric_type", predefined_metric_type)
        if resource_label is not None:
            pulumi.set(__self__, "resource_label", resource_label)

    @property
    @pulumi.getter(name="predefinedMetricType")
    def predefined_metric_type(self) -> str:
        """
        The metric type.
        """
        ...

    @property
    @pulumi.getter(name="resourceLabel")
    def resource_label(self) -> Optional[str]:
        """
        Identifies the resource associated with the metric type.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


