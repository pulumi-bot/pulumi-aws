# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Target']


class Target(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 max_capacity: Optional[pulumi.Input[float]] = None,
                 min_capacity: Optional[pulumi.Input[float]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 scalable_dimension: Optional[pulumi.Input[str]] = None,
                 service_namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Application AutoScaling ScalableTarget resource. To manage policies which get attached to the target, see the `appautoscaling.Policy` resource.

        > **NOTE:** The [Application Auto Scaling service automatically attempts to manage IAM Service-Linked Roles](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles) when registering certain service namespaces for the first time. To manually manage this role, see the `iam.ServiceLinkedRole` resource.

        ## Example Usage
        ### DynamoDB Table Autoscaling

        ```python
        import pulumi
        import pulumi_aws as aws

        dynamodb_table_read_target = aws.appautoscaling.Target("dynamodbTableReadTarget",
            max_capacity=100,
            min_capacity=5,
            resource_id=f"table/{aws_dynamodb_table['example']['name']}",
            scalable_dimension="dynamodb:table:ReadCapacityUnits",
            service_namespace="dynamodb")
        ```
        ### DynamoDB Index Autoscaling

        ```python
        import pulumi
        import pulumi_aws as aws

        dynamodb_index_read_target = aws.appautoscaling.Target("dynamodbIndexReadTarget",
            max_capacity=100,
            min_capacity=5,
            resource_id=f"table/{aws_dynamodb_table['example']['name']}/index/{var['index_name']}",
            scalable_dimension="dynamodb:index:ReadCapacityUnits",
            service_namespace="dynamodb")
        ```
        ### ECS Service Autoscaling

        ```python
        import pulumi
        import pulumi_aws as aws

        ecs_target = aws.appautoscaling.Target("ecsTarget",
            max_capacity=4,
            min_capacity=1,
            resource_id=f"service/{aws_ecs_cluster['example']['name']}/{aws_ecs_service['example']['name']}",
            scalable_dimension="ecs:service:DesiredCount",
            service_namespace="ecs")
        ```
        ### Aurora Read Replica Autoscaling

        ```python
        import pulumi
        import pulumi_aws as aws

        replicas = aws.appautoscaling.Target("replicas",
            max_capacity=15,
            min_capacity=1,
            resource_id=f"cluster:{aws_rds_cluster['example']['id']}",
            scalable_dimension="rds:cluster:ReadReplicaCount",
            service_namespace="rds")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] max_capacity: The max capacity of the scalable target.
        :param pulumi.Input[float] min_capacity: The min capacity of the scalable target.
        :param pulumi.Input[str] resource_id: The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] role_arn: The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf. This defaults to an IAM Service-Linked Role for most services and custom IAM Roles are ignored by the API for those namespaces. See the [AWS Application Auto Scaling documentation](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles) for more information about how this service interacts with IAM.
        :param pulumi.Input[str] scalable_dimension: The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] service_namespace: The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
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

            if max_capacity is None:
                raise TypeError("Missing required property 'max_capacity'")
            __props__['max_capacity'] = max_capacity
            if min_capacity is None:
                raise TypeError("Missing required property 'min_capacity'")
            __props__['min_capacity'] = min_capacity
            if resource_id is None:
                raise TypeError("Missing required property 'resource_id'")
            __props__['resource_id'] = resource_id
            __props__['role_arn'] = role_arn
            if scalable_dimension is None:
                raise TypeError("Missing required property 'scalable_dimension'")
            __props__['scalable_dimension'] = scalable_dimension
            if service_namespace is None:
                raise TypeError("Missing required property 'service_namespace'")
            __props__['service_namespace'] = service_namespace
        super(Target, __self__).__init__(
            'aws:appautoscaling/target:Target',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            max_capacity: Optional[pulumi.Input[float]] = None,
            min_capacity: Optional[pulumi.Input[float]] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            scalable_dimension: Optional[pulumi.Input[str]] = None,
            service_namespace: Optional[pulumi.Input[str]] = None) -> 'Target':
        """
        Get an existing Target resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] max_capacity: The max capacity of the scalable target.
        :param pulumi.Input[float] min_capacity: The min capacity of the scalable target.
        :param pulumi.Input[str] resource_id: The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] role_arn: The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf. This defaults to an IAM Service-Linked Role for most services and custom IAM Roles are ignored by the API for those namespaces. See the [AWS Application Auto Scaling documentation](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles) for more information about how this service interacts with IAM.
        :param pulumi.Input[str] scalable_dimension: The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        :param pulumi.Input[str] service_namespace: The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["max_capacity"] = max_capacity
        __props__["min_capacity"] = min_capacity
        __props__["resource_id"] = resource_id
        __props__["role_arn"] = role_arn
        __props__["scalable_dimension"] = scalable_dimension
        __props__["service_namespace"] = service_namespace
        return Target(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="maxCapacity")
    def max_capacity(self) -> float:
        """
        The max capacity of the scalable target.
        """
        return pulumi.get(self, "max_capacity")

    @property
    @pulumi.getter(name="minCapacity")
    def min_capacity(self) -> float:
        """
        The min capacity of the scalable target.
        """
        return pulumi.get(self, "min_capacity")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The resource type and unique identifier string for the resource associated with the scaling policy. Documentation can be found in the `ResourceId` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        """
        The ARN of the IAM role that allows Application AutoScaling to modify your scalable target on your behalf. This defaults to an IAM Service-Linked Role for most services and custom IAM Roles are ignored by the API for those namespaces. See the [AWS Application Auto Scaling documentation](https://docs.aws.amazon.com/autoscaling/application/userguide/security_iam_service-with-iam.html#security_iam_service-with-iam-roles) for more information about how this service interacts with IAM.
        """
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="scalableDimension")
    def scalable_dimension(self) -> str:
        """
        The scalable dimension of the scalable target. Documentation can be found in the `ScalableDimension` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        """
        return pulumi.get(self, "scalable_dimension")

    @property
    @pulumi.getter(name="serviceNamespace")
    def service_namespace(self) -> str:
        """
        The AWS service namespace of the scalable target. Documentation can be found in the `ServiceNamespace` parameter at: [AWS Application Auto Scaling API Reference](https://docs.aws.amazon.com/autoscaling/application/APIReference/API_RegisterScalableTarget.html#API_RegisterScalableTarget_RequestParameters)
        """
        return pulumi.get(self, "service_namespace")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

