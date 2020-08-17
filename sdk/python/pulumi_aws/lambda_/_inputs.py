# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AliasRoutingConfigArgs',
    'EventSourceMappingDestinationConfigArgs',
    'EventSourceMappingDestinationConfigOnFailureArgs',
    'FunctionDeadLetterConfigArgs',
    'FunctionEnvironmentArgs',
    'FunctionEventInvokeConfigDestinationConfigArgs',
    'FunctionEventInvokeConfigDestinationConfigOnFailureArgs',
    'FunctionEventInvokeConfigDestinationConfigOnSuccessArgs',
    'FunctionFileSystemConfigArgs',
    'FunctionTracingConfigArgs',
    'FunctionVpcConfigArgs',
]

@pulumi.input_type
class AliasRoutingConfigArgs:
    def __init__(__self__, *,
                 additional_version_weights: Optional[pulumi.Input[Mapping[str, pulumi.Input[float]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[float]]] additional_version_weights: A map that defines the proportion of events that should be sent to different versions of a lambda function.
        """
        if additional_version_weights is not None:
            pulumi.set(__self__, "additional_version_weights", additional_version_weights)

    @property
    @pulumi.getter(name="additionalVersionWeights")
    def additional_version_weights(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[float]]]]:
        """
        A map that defines the proportion of events that should be sent to different versions of a lambda function.
        """
        ...

    @additional_version_weights.setter
    def additional_version_weights(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[float]]]]):
        ...


@pulumi.input_type
class EventSourceMappingDestinationConfigArgs:
    def __init__(__self__, *,
                 on_failure: Optional[pulumi.Input['EventSourceMappingDestinationConfigOnFailureArgs']] = None):
        """
        :param pulumi.Input['EventSourceMappingDestinationConfigOnFailureArgs'] on_failure: The destination configuration for failed invocations. Detailed below.
        """
        if on_failure is not None:
            pulumi.set(__self__, "on_failure", on_failure)

    @property
    @pulumi.getter(name="onFailure")
    def on_failure(self) -> Optional[pulumi.Input['EventSourceMappingDestinationConfigOnFailureArgs']]:
        """
        The destination configuration for failed invocations. Detailed below.
        """
        ...

    @on_failure.setter
    def on_failure(self, value: Optional[pulumi.Input['EventSourceMappingDestinationConfigOnFailureArgs']]):
        ...


@pulumi.input_type
class EventSourceMappingDestinationConfigOnFailureArgs:
    def __init__(__self__, *,
                 destination_arn: pulumi.Input[str]):
        """
        :param pulumi.Input[str] destination_arn: The Amazon Resource Name (ARN) of the destination resource.
        """
        pulumi.set(__self__, "destination_arn", destination_arn)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the destination resource.
        """
        ...

    @destination_arn.setter
    def destination_arn(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class FunctionDeadLetterConfigArgs:
    def __init__(__self__, *,
                 target_arn: pulumi.Input[str]):
        """
        :param pulumi.Input[str] target_arn: The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this
               option is used, the function's IAM role must be granted suitable access to write to the target object,
               which means allowing either the `sns:Publish` or `sqs:SendMessage` action on this ARN, depending on
               which service is targeted.
        """
        pulumi.set(__self__, "target_arn", target_arn)

    @property
    @pulumi.getter(name="targetArn")
    def target_arn(self) -> pulumi.Input[str]:
        """
        The ARN of an SNS topic or SQS queue to notify when an invocation fails. If this
        option is used, the function's IAM role must be granted suitable access to write to the target object,
        which means allowing either the `sns:Publish` or `sqs:SendMessage` action on this ARN, depending on
        which service is targeted.
        """
        ...

    @target_arn.setter
    def target_arn(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class FunctionEnvironmentArgs:
    def __init__(__self__, *,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] variables: A map that defines environment variables for the Lambda function.
        """
        if variables is not None:
            pulumi.set(__self__, "variables", variables)

    @property
    @pulumi.getter
    def variables(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map that defines environment variables for the Lambda function.
        """
        ...

    @variables.setter
    def variables(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class FunctionEventInvokeConfigDestinationConfigArgs:
    def __init__(__self__, *,
                 on_failure: Optional[pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnFailureArgs']] = None,
                 on_success: Optional[pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnSuccessArgs']] = None):
        """
        :param pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnFailureArgs'] on_failure: Configuration block with destination configuration for failed asynchronous invocations. See below for details.
        :param pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnSuccessArgs'] on_success: Configuration block with destination configuration for successful asynchronous invocations. See below for details.
        """
        if on_failure is not None:
            pulumi.set(__self__, "on_failure", on_failure)
        if on_success is not None:
            pulumi.set(__self__, "on_success", on_success)

    @property
    @pulumi.getter(name="onFailure")
    def on_failure(self) -> Optional[pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnFailureArgs']]:
        """
        Configuration block with destination configuration for failed asynchronous invocations. See below for details.
        """
        ...

    @on_failure.setter
    def on_failure(self, value: Optional[pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnFailureArgs']]):
        ...

    @property
    @pulumi.getter(name="onSuccess")
    def on_success(self) -> Optional[pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnSuccessArgs']]:
        """
        Configuration block with destination configuration for successful asynchronous invocations. See below for details.
        """
        ...

    @on_success.setter
    def on_success(self, value: Optional[pulumi.Input['FunctionEventInvokeConfigDestinationConfigOnSuccessArgs']]):
        ...


@pulumi.input_type
class FunctionEventInvokeConfigDestinationConfigOnFailureArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str]):
        """
        :param pulumi.Input[str] destination: Amazon Resource Name (ARN) of the destination resource. See the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations) for acceptable resource types and associated IAM permissions.
        """
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the destination resource. See the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations) for acceptable resource types and associated IAM permissions.
        """
        ...

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class FunctionEventInvokeConfigDestinationConfigOnSuccessArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str]):
        """
        :param pulumi.Input[str] destination: Amazon Resource Name (ARN) of the destination resource. See the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations) for acceptable resource types and associated IAM permissions.
        """
        pulumi.set(__self__, "destination", destination)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        Amazon Resource Name (ARN) of the destination resource. See the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-async-destinations) for acceptable resource types and associated IAM permissions.
        """
        ...

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class FunctionFileSystemConfigArgs:
    def __init__(__self__, *,
                 arn: pulumi.Input[str],
                 local_mount_path: pulumi.Input[str]):
        """
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
        :param pulumi.Input[str] local_mount_path: The path where the function can access the file system, starting with /mnt/.
        """
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "local_mount_path", local_mount_path)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Input[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
        """
        ...

    @arn.setter
    def arn(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="localMountPath")
    def local_mount_path(self) -> pulumi.Input[str]:
        """
        The path where the function can access the file system, starting with /mnt/.
        """
        ...

    @local_mount_path.setter
    def local_mount_path(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class FunctionTracingConfigArgs:
    def __init__(__self__, *,
                 mode: pulumi.Input[str]):
        """
        :param pulumi.Input[str] mode: Can be either `PassThrough` or `Active`. If PassThrough, Lambda will only trace
               the request from an upstream service if it contains a tracing header with
               "sampled=1". If Active, Lambda will respect any tracing header it receives
               from an upstream service. If no tracing header is received, Lambda will call
               X-Ray for a tracing decision.
        """
        pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> pulumi.Input[str]:
        """
        Can be either `PassThrough` or `Active`. If PassThrough, Lambda will only trace
        the request from an upstream service if it contains a tracing header with
        "sampled=1". If Active, Lambda will respect any tracing header it receives
        from an upstream service. If no tracing header is received, Lambda will call
        X-Ray for a tracing decision.
        """
        ...

    @mode.setter
    def mode(self, value: pulumi.Input[str]):
        ...


@pulumi.input_type
class FunctionVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[List[pulumi.Input[str]]],
                 subnet_ids: pulumi.Input[List[pulumi.Input[str]]],
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_ids: A list of security group IDs associated with the Lambda function.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: A list of subnet IDs associated with the Lambda function.
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of security group IDs associated with the Lambda function.
        """
        ...

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        ...

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of subnet IDs associated with the Lambda function.
        """
        ...

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        ...

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        ...

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        ...


