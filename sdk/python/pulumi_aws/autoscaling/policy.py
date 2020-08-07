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

__all__ = ['Policy']


class Policy(pulumi.CustomResource):
    adjustment_type: pulumi.Output[Optional[str]] = pulumi.property("adjustmentType")
    """
    Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
    """

    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN assigned by AWS to the scaling policy.
    """

    autoscaling_group_name: pulumi.Output[str] = pulumi.property("autoscalingGroupName")
    """
    The name of the autoscaling group.
    """

    cooldown: pulumi.Output[Optional[float]] = pulumi.property("cooldown")
    """
    The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
    """

    estimated_instance_warmup: pulumi.Output[Optional[float]] = pulumi.property("estimatedInstanceWarmup")
    """
    The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
    """

    metric_aggregation_type: pulumi.Output[str] = pulumi.property("metricAggregationType")
    """
    The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
    """

    min_adjustment_magnitude: pulumi.Output[Optional[float]] = pulumi.property("minAdjustmentMagnitude")

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the dimension.
    """

    policy_type: pulumi.Output[Optional[str]] = pulumi.property("policyType")
    """
    The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
    """

    scaling_adjustment: pulumi.Output[Optional[float]] = pulumi.property("scalingAdjustment")
    """
    The number of members by which to
    scale, when the adjustment bounds are breached. A positive value scales
    up. A negative value scales down.
    """

    step_adjustments: pulumi.Output[Optional[List['outputs.PolicyStepAdjustment']]] = pulumi.property("stepAdjustments")
    """
    A set of adjustments that manage
    group scaling. These have the following structure:
    """

    target_tracking_configuration: pulumi.Output[Optional['outputs.PolicyTargetTrackingConfiguration']] = pulumi.property("targetTrackingConfiguration")
    """
    A target tracking policy. These have the following structure:
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 adjustment_type: Optional[pulumi.Input[str]] = None,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 cooldown: Optional[pulumi.Input[float]] = None,
                 estimated_instance_warmup: Optional[pulumi.Input[float]] = None,
                 metric_aggregation_type: Optional[pulumi.Input[str]] = None,
                 min_adjustment_magnitude: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 scaling_adjustment: Optional[pulumi.Input[float]] = None,
                 step_adjustments: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyStepAdjustmentArgs']]]]] = None,
                 target_tracking_configuration: Optional[pulumi.Input[pulumi.InputType['PolicyTargetTrackingConfigurationArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AutoScaling Scaling Policy resource.

        > **NOTE:** You may want to omit `desired_capacity` attribute from attached `autoscaling.Group`
        when using autoscaling policies. It's good practice to pick either
        [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
        or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
        (policy-based) scaling.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        bar = aws.autoscaling.Group("bar",
            availability_zones=["us-east-1a"],
            force_delete=True,
            health_check_grace_period=300,
            health_check_type="ELB",
            launch_configuration=aws_launch_configuration["foo"]["name"],
            max_size=5,
            min_size=2)
        bat = aws.autoscaling.Policy("bat",
            adjustment_type="ChangeInCapacity",
            autoscaling_group_name=bar.name,
            cooldown=300,
            scaling_adjustment=4)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adjustment_type: Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
        :param pulumi.Input[str] autoscaling_group_name: The name of the autoscaling group.
        :param pulumi.Input[float] cooldown: The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
        :param pulumi.Input[float] estimated_instance_warmup: The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
        :param pulumi.Input[str] metric_aggregation_type: The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
        :param pulumi.Input[str] name: The name of the dimension.
        :param pulumi.Input[str] policy_type: The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
        :param pulumi.Input[float] scaling_adjustment: The number of members by which to
               scale, when the adjustment bounds are breached. A positive value scales
               up. A negative value scales down.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyStepAdjustmentArgs']]]] step_adjustments: A set of adjustments that manage
               group scaling. These have the following structure:
        :param pulumi.Input[pulumi.InputType['PolicyTargetTrackingConfigurationArgs']] target_tracking_configuration: A target tracking policy. These have the following structure:
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

            __props__['adjustment_type'] = adjustment_type
            if autoscaling_group_name is None:
                raise TypeError("Missing required property 'autoscaling_group_name'")
            __props__['autoscaling_group_name'] = autoscaling_group_name
            __props__['cooldown'] = cooldown
            __props__['estimated_instance_warmup'] = estimated_instance_warmup
            __props__['metric_aggregation_type'] = metric_aggregation_type
            __props__['min_adjustment_magnitude'] = min_adjustment_magnitude
            __props__['name'] = name
            __props__['policy_type'] = policy_type
            __props__['scaling_adjustment'] = scaling_adjustment
            __props__['step_adjustments'] = step_adjustments
            __props__['target_tracking_configuration'] = target_tracking_configuration
            __props__['arn'] = None
        super(Policy, __self__).__init__(
            'aws:autoscaling/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            adjustment_type: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            autoscaling_group_name: Optional[pulumi.Input[str]] = None,
            cooldown: Optional[pulumi.Input[float]] = None,
            estimated_instance_warmup: Optional[pulumi.Input[float]] = None,
            metric_aggregation_type: Optional[pulumi.Input[str]] = None,
            min_adjustment_magnitude: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            scaling_adjustment: Optional[pulumi.Input[float]] = None,
            step_adjustments: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyStepAdjustmentArgs']]]]] = None,
            target_tracking_configuration: Optional[pulumi.Input[pulumi.InputType['PolicyTargetTrackingConfigurationArgs']]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] adjustment_type: Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
        :param pulumi.Input[str] arn: The ARN assigned by AWS to the scaling policy.
        :param pulumi.Input[str] autoscaling_group_name: The name of the autoscaling group.
        :param pulumi.Input[float] cooldown: The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
        :param pulumi.Input[float] estimated_instance_warmup: The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
        :param pulumi.Input[str] metric_aggregation_type: The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
        :param pulumi.Input[str] name: The name of the dimension.
        :param pulumi.Input[str] policy_type: The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
        :param pulumi.Input[float] scaling_adjustment: The number of members by which to
               scale, when the adjustment bounds are breached. A positive value scales
               up. A negative value scales down.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['PolicyStepAdjustmentArgs']]]] step_adjustments: A set of adjustments that manage
               group scaling. These have the following structure:
        :param pulumi.Input[pulumi.InputType['PolicyTargetTrackingConfigurationArgs']] target_tracking_configuration: A target tracking policy. These have the following structure:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["adjustment_type"] = adjustment_type
        __props__["arn"] = arn
        __props__["autoscaling_group_name"] = autoscaling_group_name
        __props__["cooldown"] = cooldown
        __props__["estimated_instance_warmup"] = estimated_instance_warmup
        __props__["metric_aggregation_type"] = metric_aggregation_type
        __props__["min_adjustment_magnitude"] = min_adjustment_magnitude
        __props__["name"] = name
        __props__["policy_type"] = policy_type
        __props__["scaling_adjustment"] = scaling_adjustment
        __props__["step_adjustments"] = step_adjustments
        __props__["target_tracking_configuration"] = target_tracking_configuration
        return Policy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

