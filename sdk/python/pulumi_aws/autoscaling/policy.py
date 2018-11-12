# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Policy(pulumi.CustomResource):
    """
    Provides an AutoScaling Scaling Policy resource.
    
    ~> **NOTE:** You may want to omit `desired_capacity` attribute from attached `aws_autoscaling_group`
    when using autoscaling policies. It's good practice to pick either
    [manual](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-manual-scaling.html)
    or [dynamic](https://docs.aws.amazon.com/AutoScaling/latest/DeveloperGuide/as-scale-based-on-demand.html)
    (policy-based) scaling.
    """
    def __init__(__self__, __name__, __opts__=None, adjustment_type=None, autoscaling_group_name=None, cooldown=None, estimated_instance_warmup=None, metric_aggregation_type=None, min_adjustment_magnitude=None, min_adjustment_step=None, name=None, policy_type=None, scaling_adjustment=None, step_adjustments=None, target_tracking_configuration=None):
        """Create a Policy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['adjustment_type'] = adjustment_type

        if not autoscaling_group_name:
            raise TypeError('Missing required property autoscaling_group_name')
        __props__['autoscaling_group_name'] = autoscaling_group_name

        __props__['cooldown'] = cooldown

        __props__['estimated_instance_warmup'] = estimated_instance_warmup

        __props__['metric_aggregation_type'] = metric_aggregation_type

        __props__['min_adjustment_magnitude'] = min_adjustment_magnitude

        __props__['min_adjustment_step'] = min_adjustment_step

        __props__['name'] = name

        __props__['policy_type'] = policy_type

        __props__['scaling_adjustment'] = scaling_adjustment

        __props__['step_adjustments'] = step_adjustments

        __props__['target_tracking_configuration'] = target_tracking_configuration

        __props__['arn'] = None

        super(Policy, __self__).__init__(
            'aws:autoscaling/policy:Policy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

