# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Policy(pulumi.CustomResource):
    adjustment_type: pulumi.Output[str]
    """
    Specifies whether the adjustment is an absolute number or a percentage of the current capacity. Valid values are `ChangeInCapacity`, `ExactCapacity`, and `PercentChangeInCapacity`.
    """
    arn: pulumi.Output[str]
    """
    The ARN assigned by AWS to the scaling policy.
    """
    autoscaling_group_name: pulumi.Output[str]
    """
    The name of the autoscaling group.
    """
    cooldown: pulumi.Output[float]
    """
    The amount of time, in seconds, after a scaling activity completes and before the next scaling activity can start.
    """
    estimated_instance_warmup: pulumi.Output[float]
    """
    The estimated time, in seconds, until a newly launched instance will contribute CloudWatch metrics. Without a value, AWS will default to the group's specified cooldown period.
    """
    metric_aggregation_type: pulumi.Output[str]
    """
    The aggregation type for the policy's metrics. Valid values are "Minimum", "Maximum", and "Average". Without a value, AWS will treat the aggregation type as "Average".
    """
    min_adjustment_magnitude: pulumi.Output[float]
    name: pulumi.Output[str]
    """
    The name of the dimension.
    """
    policy_type: pulumi.Output[str]
    """
    The policy type, either "SimpleScaling", "StepScaling" or "TargetTrackingScaling". If this value isn't provided, AWS will default to "SimpleScaling."
    """
    scaling_adjustment: pulumi.Output[float]
    """
    The number of members by which to
    scale, when the adjustment bounds are breached. A positive value scales
    up. A negative value scales down.
    """
    step_adjustments: pulumi.Output[list]
    target_tracking_configuration: pulumi.Output[dict]
    """
    A target tracking policy. These have the following structure:
    """
    def __init__(__self__, resource_name, opts=None, adjustment_type=None, autoscaling_group_name=None, cooldown=None, estimated_instance_warmup=None, metric_aggregation_type=None, min_adjustment_magnitude=None, name=None, policy_type=None, scaling_adjustment=None, step_adjustments=None, target_tracking_configuration=None, __name__=None, __opts__=None):
        """
        Provides an AutoScaling Scaling Policy resource.
        
        > **NOTE:** You may want to omit `desired_capacity` attribute from attached `autoscaling.Group`
        when using autoscaling policies. It's good practice to pick either
        [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
        or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
        (policy-based) scaling.
        
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
        :param pulumi.Input[dict] target_tracking_configuration: A target tracking policy. These have the following structure:

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/autoscaling_policy.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

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

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Policy, __self__).__init__(
            'aws:autoscaling/policy:Policy',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

