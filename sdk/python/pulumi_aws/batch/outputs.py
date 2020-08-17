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
    'ComputeEnvironmentComputeResources',
    'ComputeEnvironmentComputeResourcesLaunchTemplate',
    'JobDefinitionRetryStrategy',
    'JobDefinitionTimeout',
    'GetJobQueueComputeEnvironmentOrderResult',
]

@pulumi.output_type
class ComputeEnvironmentComputeResources(dict):
    def __init__(__self__, *,
                 instance_role: str,
                 instance_types: List[str],
                 max_vcpus: float,
                 min_vcpus: float,
                 security_group_ids: List[str],
                 subnets: List[str],
                 type: str,
                 allocation_strategy: Optional[str] = None,
                 bid_percentage: Optional[float] = None,
                 desired_vcpus: Optional[float] = None,
                 ec2_key_pair: Optional[str] = None,
                 image_id: Optional[str] = None,
                 launch_template: Optional['outputs.ComputeEnvironmentComputeResourcesLaunchTemplate'] = None,
                 spot_iam_fleet_role: Optional[str] = None,
                 tags: Optional[Mapping[str, str]] = None):
        """
        :param str instance_role: The Amazon ECS instance role applied to Amazon EC2 instances in a compute environment.
        :param List[str] instance_types: A list of instance types that may be launched.
        :param float max_vcpus: The maximum number of EC2 vCPUs that an environment can reach.
        :param float min_vcpus: The minimum number of EC2 vCPUs that an environment should maintain.
        :param List[str] security_group_ids: A list of EC2 security group that are associated with instances launched in the compute environment.
        :param List[str] subnets: A list of VPC subnets into which the compute resources are launched.
        :param str type: The type of compute environment. Valid items are `EC2` or `SPOT`.
        :param str allocation_strategy: The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated. Valid items are `BEST_FIT_PROGRESSIVE`, `SPOT_CAPACITY_OPTIMIZED` or `BEST_FIT`. Defaults to `BEST_FIT`. See [AWS docs](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) for details.
        :param float bid_percentage: Integer of minimum percentage that a Spot Instance price must be when compared with the On-Demand price for that instance type before instances are launched. For example, if your bid percentage is 20% (`20`), then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. This parameter is required for SPOT compute environments.
        :param float desired_vcpus: The desired number of EC2 vCPUS in the compute environment.
        :param str ec2_key_pair: The EC2 key pair that is used for instances launched in the compute environment.
        :param str image_id: The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
        :param 'ComputeEnvironmentComputeResourcesLaunchTemplateArgs' launch_template: The launch template to use for your compute resources. See details below.
        :param str spot_iam_fleet_role: The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment. This parameter is required for SPOT compute environments.
        :param Mapping[str, str] tags: Key-value pair tags to be applied to resources that are launched in the compute environment.
        """
        pulumi.set(__self__, "instance_role", instance_role)
        pulumi.set(__self__, "instance_types", instance_types)
        pulumi.set(__self__, "max_vcpus", max_vcpus)
        pulumi.set(__self__, "min_vcpus", min_vcpus)
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnets", subnets)
        pulumi.set(__self__, "type", type)
        if allocation_strategy is not None:
            pulumi.set(__self__, "allocation_strategy", allocation_strategy)
        if bid_percentage is not None:
            pulumi.set(__self__, "bid_percentage", bid_percentage)
        if desired_vcpus is not None:
            pulumi.set(__self__, "desired_vcpus", desired_vcpus)
        if ec2_key_pair is not None:
            pulumi.set(__self__, "ec2_key_pair", ec2_key_pair)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if launch_template is not None:
            pulumi.set(__self__, "launch_template", launch_template)
        if spot_iam_fleet_role is not None:
            pulumi.set(__self__, "spot_iam_fleet_role", spot_iam_fleet_role)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="instanceRole")
    def instance_role(self) -> str:
        """
        The Amazon ECS instance role applied to Amazon EC2 instances in a compute environment.
        """
        ...

    @property
    @pulumi.getter(name="instanceTypes")
    def instance_types(self) -> List[str]:
        """
        A list of instance types that may be launched.
        """
        ...

    @property
    @pulumi.getter(name="maxVcpus")
    def max_vcpus(self) -> float:
        """
        The maximum number of EC2 vCPUs that an environment can reach.
        """
        ...

    @property
    @pulumi.getter(name="minVcpus")
    def min_vcpus(self) -> float:
        """
        The minimum number of EC2 vCPUs that an environment should maintain.
        """
        ...

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> List[str]:
        """
        A list of EC2 security group that are associated with instances launched in the compute environment.
        """
        ...

    @property
    @pulumi.getter
    def subnets(self) -> List[str]:
        """
        A list of VPC subnets into which the compute resources are launched.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of compute environment. Valid items are `EC2` or `SPOT`.
        """
        ...

    @property
    @pulumi.getter(name="allocationStrategy")
    def allocation_strategy(self) -> Optional[str]:
        """
        The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated. Valid items are `BEST_FIT_PROGRESSIVE`, `SPOT_CAPACITY_OPTIMIZED` or `BEST_FIT`. Defaults to `BEST_FIT`. See [AWS docs](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) for details.
        """
        ...

    @property
    @pulumi.getter(name="bidPercentage")
    def bid_percentage(self) -> Optional[float]:
        """
        Integer of minimum percentage that a Spot Instance price must be when compared with the On-Demand price for that instance type before instances are launched. For example, if your bid percentage is 20% (`20`), then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. This parameter is required for SPOT compute environments.
        """
        ...

    @property
    @pulumi.getter(name="desiredVcpus")
    def desired_vcpus(self) -> Optional[float]:
        """
        The desired number of EC2 vCPUS in the compute environment.
        """
        ...

    @property
    @pulumi.getter(name="ec2KeyPair")
    def ec2_key_pair(self) -> Optional[str]:
        """
        The EC2 key pair that is used for instances launched in the compute environment.
        """
        ...

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[str]:
        """
        The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
        """
        ...

    @property
    @pulumi.getter(name="launchTemplate")
    def launch_template(self) -> Optional['outputs.ComputeEnvironmentComputeResourcesLaunchTemplate']:
        """
        The launch template to use for your compute resources. See details below.
        """
        ...

    @property
    @pulumi.getter(name="spotIamFleetRole")
    def spot_iam_fleet_role(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment. This parameter is required for SPOT compute environments.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Key-value pair tags to be applied to resources that are launched in the compute environment.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ComputeEnvironmentComputeResourcesLaunchTemplate(dict):
    def __init__(__self__, *,
                 launch_template_id: Optional[str] = None,
                 launch_template_name: Optional[str] = None,
                 version: Optional[str] = None):
        """
        :param str launch_template_id: ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
        :param str launch_template_name: Name of the launch template.
        :param str version: The version number of the launch template. Default: The default version of the launch template.
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
        ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
        """
        ...

    @property
    @pulumi.getter(name="launchTemplateName")
    def launch_template_name(self) -> Optional[str]:
        """
        Name of the launch template.
        """
        ...

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        The version number of the launch template. Default: The default version of the launch template.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobDefinitionRetryStrategy(dict):
    def __init__(__self__, *,
                 attempts: Optional[float] = None):
        """
        :param float attempts: The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
        """
        if attempts is not None:
            pulumi.set(__self__, "attempts", attempts)

    @property
    @pulumi.getter
    def attempts(self) -> Optional[float]:
        """
        The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class JobDefinitionTimeout(dict):
    def __init__(__self__, *,
                 attempt_duration_seconds: Optional[float] = None):
        """
        :param float attempt_duration_seconds: The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
        """
        if attempt_duration_seconds is not None:
            pulumi.set(__self__, "attempt_duration_seconds", attempt_duration_seconds)

    @property
    @pulumi.getter(name="attemptDurationSeconds")
    def attempt_duration_seconds(self) -> Optional[float]:
        """
        The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetJobQueueComputeEnvironmentOrderResult(dict):
    def __init__(__self__, *,
                 compute_environment: str,
                 order: float):
        pulumi.set(__self__, "compute_environment", compute_environment)
        pulumi.set(__self__, "order", order)

    @property
    @pulumi.getter(name="computeEnvironment")
    def compute_environment(self) -> str:
        ...

    @property
    @pulumi.getter
    def order(self) -> float:
        ...


