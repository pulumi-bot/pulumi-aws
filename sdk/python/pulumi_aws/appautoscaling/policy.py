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
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN assigned by AWS to the scaling policy.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the policy.
    """

    policy_type: pulumi.Output[Optional[str]] = pulumi.property("policyType")
    """
    The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
    """

    resource_id: pulumi.Output[str] = pulumi.property("resourceId")
    """
    The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
    """

    scalable_dimension: pulumi.Output[str] = pulumi.property("scalableDimension")
    """
    The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
    """

    service_namespace: pulumi.Output[str] = pulumi.property("serviceNamespace")
    """
    The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
    """

    step_scaling_policy_configuration: pulumi.Output[Optional['outputs.PolicyStepScalingPolicyConfiguration']] = pulumi.property("stepScalingPolicyConfiguration")
    """
    Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
    """

    target_tracking_scaling_policy_configuration: pulumi.Output[Optional['outputs.PolicyTargetTrackingScalingPolicyConfiguration']] = pulumi.property("targetTrackingScalingPolicyConfiguration")
    """
    A target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 scalable_dimension: Optional[pulumi.Input[str]] = None,
                 service_namespace: Optional[pulumi.Input[str]] = None,
                 step_scaling_policy_configuration: Optional[pulumi.Input[pulumi.InputType['PolicyStepScalingPolicyConfigurationArgs']]] = None,
                 target_tracking_scaling_policy_configuration: Optional[pulumi.Input[pulumi.InputType['PolicyTargetTrackingScalingPolicyConfigurationArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Application AutoScaling Policy resource.

        ## Example Usage
        ### DynamoDB Table Autoscaling

        ```python
        import pulumi
        import pulumi_aws as aws

        dynamodb_table_read_target = aws.appautoscaling.Target("dynamodbTableReadTarget",
            max_capacity=100,
            min_capacity=5,
            resource_id="table/tableName",
            scalable_dimension="dynamodb:table:ReadCapacityUnits",
            service_namespace="dynamodb")
        dynamodb_table_read_policy = aws.appautoscaling.Policy("dynamodbTableReadPolicy",
            policy_type="TargetTrackingScaling",
            resource_id=dynamodb_table_read_target.resource_id,
            scalable_dimension=dynamodb_table_read_target.scalable_dimension,
            service_namespace=dynamodb_table_read_target.service_namespace,
            target_tracking_scaling_policy_configuration={
                "predefinedMetricSpecification": {
                    "predefinedMetricType": "DynamoDBReadCapacityUtilization",
                },
                "targetValue": 70,
            })
        ```
        ### ECS Service Autoscaling

        ```python
        import pulumi
        import pulumi_aws as aws

        ecs_target = aws.appautoscaling.Target("ecsTarget",
            max_capacity=4,
            min_capacity=1,
            resource_id="service/clusterName/serviceName",
            scalable_dimension="ecs:service:DesiredCount",
            service_namespace="ecs")
        ecs_policy = aws.appautoscaling.Policy("ecsPolicy",
            policy_type="StepScaling",
            resource_id=ecs_target.resource_id,
            scalable_dimension=ecs_target.scalable_dimension,
            service_namespace=ecs_target.service_namespace,
            step_scaling_policy_configuration={
                "adjustment_type": "ChangeInCapacity",
                "cooldown": 60,
                "metric_aggregation_type": "Maximum",
                "step_adjustments": [{
                    "metricIntervalUpperBound": 0,
                    "scaling_adjustment": -1,
                }],
            })
        ```
        ### Preserve desired count when updating an autoscaled ECS Service

        ```python
        import pulumi
        import pulumi_aws as aws

        ecs_service = aws.ecs.Service("ecsService",
            cluster="clusterName",
            task_definition="taskDefinitionFamily:1",
            desired_count=2)
        ```
        ### Aurora Read Replica Autoscaling

        ```python
        import pulumi
        import pulumi_aws as aws

        replicas_target = aws.appautoscaling.Target("replicasTarget",
            service_namespace="rds",
            scalable_dimension="rds:cluster:ReadReplicaCount",
            resource_id=f"cluster:{aws_rds_cluster['example']['id']}",
            min_capacity=1,
            max_capacity=15)
        replicas_policy = aws.appautoscaling.Policy("replicasPolicy",
            service_namespace=replicas_target.service_namespace,
            scalable_dimension=replicas_target.scalable_dimension,
            resource_id=replicas_target.resource_id,
            policy_type="TargetTrackingScaling",
            target_tracking_scaling_policy_configuration={
                "predefinedMetricSpecification": {
                    "predefinedMetricType": "RDSReaderAverageCPUUtilization",
                },
                "targetValue": 75,
                "scaleInCooldown": 300,
                "scaleOutCooldown": 300,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[str] policy_type: The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
        :param pulumi.Input[str] resource_id: The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] scalable_dimension: The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] service_namespace: The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[pulumi.InputType['PolicyStepScalingPolicyConfigurationArgs']] step_scaling_policy_configuration: Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
        :param pulumi.Input[pulumi.InputType['PolicyTargetTrackingScalingPolicyConfigurationArgs']] target_tracking_scaling_policy_configuration: A target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.
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

            __props__['name'] = name
            __props__['policy_type'] = policy_type
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            if scalable_dimension is None:
                raise TypeError("Missing required property 'scalable_dimension'")
            __props__['scalable_dimension'] = scalable_dimension
            if service_namespace is None:
                raise TypeError("Missing required property 'service_namespace'")
            __props__['service_namespace'] = service_namespace
            __props__['step_scaling_policy_configuration'] = step_scaling_policy_configuration
            __props__['target_tracking_scaling_policy_configuration'] = target_tracking_scaling_policy_configuration
            __props__['arn'] = None
        super(Policy, __self__).__init__(
            'aws:appautoscaling/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            scalable_dimension: Optional[pulumi.Input[str]] = None,
            service_namespace: Optional[pulumi.Input[str]] = None,
            step_scaling_policy_configuration: Optional[pulumi.Input[pulumi.InputType['PolicyStepScalingPolicyConfigurationArgs']]] = None,
            target_tracking_scaling_policy_configuration: Optional[pulumi.Input[pulumi.InputType['PolicyTargetTrackingScalingPolicyConfigurationArgs']]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN assigned by AWS to the scaling policy.
        :param pulumi.Input[str] name: The name of the policy.
        :param pulumi.Input[str] policy_type: The policy type. Valid values are `StepScaling` and `TargetTrackingScaling`. Defaults to `StepScaling`. Certain services only support only one policy type. For more information see the [Target Tracking Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-target-tracking.html) and [Step Scaling Policies](https://docs.aws.amazon.com/autoscaling/application/userguide/application-auto-scaling-step-scaling-policies.html) documentation.
        :param pulumi.Input[str] resource_id: The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] scalable_dimension: The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] service_namespace: The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](http://docs.aws.amazon.com/ApplicationAutoScaling/latest/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[pulumi.InputType['PolicyStepScalingPolicyConfigurationArgs']] step_scaling_policy_configuration: Step scaling policy configuration, requires `policy_type = "StepScaling"` (default). See supported fields below.
        :param pulumi.Input[pulumi.InputType['PolicyTargetTrackingScalingPolicyConfigurationArgs']] target_tracking_scaling_policy_configuration: A target tracking policy, requires `policy_type = "TargetTrackingScaling"`. See supported fields below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["policy_type"] = policy_type
        __props__["resource_id"] = resource_id
        __props__["scalable_dimension"] = scalable_dimension
        __props__["service_namespace"] = service_namespace
        __props__["step_scaling_policy_configuration"] = step_scaling_policy_configuration
        __props__["target_tracking_scaling_policy_configuration"] = target_tracking_scaling_policy_configuration
        return Policy(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

