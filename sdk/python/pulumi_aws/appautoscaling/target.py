# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Target(pulumi.CustomResource):
    max_capacity: pulumi.Output[int]
    """
    The max capacity of the scalable target.
    """
    min_capacity: pulumi.Output[int]
    """
    The min capacity of the scalable target.
    """
    resource_id: pulumi.Output[str]
    """
    The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
    """
    role_arn: pulumi.Output[str]
    """
    The ARN of the IAM role that allows Application
    AutoScaling to modify your scalable target on your behalf.
    """
    scalable_dimension: pulumi.Output[str]
    """
    The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
    """
    service_namespace: pulumi.Output[str]
    """
    The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
    """
    def __init__(__self__, __name__, __opts__=None, max_capacity=None, min_capacity=None, resource_id=None, role_arn=None, scalable_dimension=None, service_namespace=None):
        """
        Provides an Application AutoScaling ScalableTarget resource. To manage policies which get attached to the target, see the [`aws_appautoscaling_policy` resource](https://www.terraform.io/docs/providers/aws/r/appautoscaling_policy.html).
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[int] max_capacity: The max capacity of the scalable target.
        :param pulumi.Input[int] min_capacity: The min capacity of the scalable target.
        :param pulumi.Input[str] resource_id: The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] role_arn: The ARN of the IAM role that allows Application
               AutoScaling to modify your scalable target on your behalf.
        :param pulumi.Input[str] scalable_dimension: The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] service_namespace: The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not max_capacity:
            raise TypeError('Missing required property max_capacity')
        __props__['max_capacity'] = max_capacity

        if not min_capacity:
            raise TypeError('Missing required property min_capacity')
        __props__['min_capacity'] = min_capacity

        if not resource_id:
            raise TypeError('Missing required property resource_id')
        __props__['resource_id'] = resource_id

        __props__['role_arn'] = role_arn

        if not scalable_dimension:
            raise TypeError('Missing required property scalable_dimension')
        __props__['scalable_dimension'] = scalable_dimension

        if not service_namespace:
            raise TypeError('Missing required property service_namespace')
        __props__['service_namespace'] = service_namespace

        super(Target, __self__).__init__(
            'aws:appautoscaling/target:Target',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

