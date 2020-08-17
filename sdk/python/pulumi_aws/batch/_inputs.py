# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ComputeEnvironmentComputeResourcesArgs',
    'ComputeEnvironmentComputeResourcesLaunchTemplateArgs',
    'JobDefinitionRetryStrategyArgs',
    'JobDefinitionTimeoutArgs',
]

@pulumi.input_type
class ComputeEnvironmentComputeResourcesArgs:
    def __init__(__self__, *,
                 instance_role: pulumi.Input[str],
                 instance_types: pulumi.Input[List[pulumi.Input[str]]],
                 max_vcpus: pulumi.Input[float],
                 min_vcpus: pulumi.Input[float],
                 security_group_ids: pulumi.Input[List[pulumi.Input[str]]],
                 subnets: pulumi.Input[List[pulumi.Input[str]]],
                 type: pulumi.Input[str],
                 allocation_strategy: Optional[pulumi.Input[str]] = None,
                 bid_percentage: Optional[pulumi.Input[float]] = None,
                 desired_vcpus: Optional[pulumi.Input[float]] = None,
                 ec2_key_pair: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 launch_template: Optional[pulumi.Input['ComputeEnvironmentComputeResourcesLaunchTemplateArgs']] = None,
                 spot_iam_fleet_role: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] instance_role: The Amazon ECS instance role applied to Amazon EC2 instances in a compute environment.
        :param pulumi.Input[List[pulumi.Input[str]]] instance_types: A list of instance types that may be launched.
        :param pulumi.Input[float] max_vcpus: The maximum number of EC2 vCPUs that an environment can reach.
        :param pulumi.Input[float] min_vcpus: The minimum number of EC2 vCPUs that an environment should maintain.
        :param pulumi.Input[List[pulumi.Input[str]]] security_group_ids: A list of EC2 security group that are associated with instances launched in the compute environment.
        :param pulumi.Input[List[pulumi.Input[str]]] subnets: A list of VPC subnets into which the compute resources are launched.
        :param pulumi.Input[str] type: The type of compute environment. Valid items are `EC2` or `SPOT`.
        :param pulumi.Input[str] allocation_strategy: The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated. Valid items are `BEST_FIT_PROGRESSIVE`, `SPOT_CAPACITY_OPTIMIZED` or `BEST_FIT`. Defaults to `BEST_FIT`. See [AWS docs](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) for details.
        :param pulumi.Input[float] bid_percentage: Integer of minimum percentage that a Spot Instance price must be when compared with the On-Demand price for that instance type before instances are launched. For example, if your bid percentage is 20% (`20`), then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. This parameter is required for SPOT compute environments.
        :param pulumi.Input[float] desired_vcpus: The desired number of EC2 vCPUS in the compute environment.
        :param pulumi.Input[str] ec2_key_pair: The EC2 key pair that is used for instances launched in the compute environment.
        :param pulumi.Input[str] image_id: The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
        :param pulumi.Input['ComputeEnvironmentComputeResourcesLaunchTemplateArgs'] launch_template: The launch template to use for your compute resources. See details below.
        :param pulumi.Input[str] spot_iam_fleet_role: The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment. This parameter is required for SPOT compute environments.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pair tags to be applied to resources that are launched in the compute environment.
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
    def instance_role(self) -> pulumi.Input[str]:
        """
        The Amazon ECS instance role applied to Amazon EC2 instances in a compute environment.
        """
        ...

    @instance_role.setter
    def instance_role(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="instanceTypes")
    def instance_types(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of instance types that may be launched.
        """
        ...

    @instance_types.setter
    def instance_types(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        ...

    @property
    @pulumi.getter(name="maxVcpus")
    def max_vcpus(self) -> pulumi.Input[float]:
        """
        The maximum number of EC2 vCPUs that an environment can reach.
        """
        ...

    @max_vcpus.setter
    def max_vcpus(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="minVcpus")
    def min_vcpus(self) -> pulumi.Input[float]:
        """
        The minimum number of EC2 vCPUs that an environment should maintain.
        """
        ...

    @min_vcpus.setter
    def min_vcpus(self, value: pulumi.Input[float]):
        ...

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of EC2 security group that are associated with instances launched in the compute environment.
        """
        ...

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        ...

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of VPC subnets into which the compute resources are launched.
        """
        ...

    @subnets.setter
    def subnets(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        ...

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of compute environment. Valid items are `EC2` or `SPOT`.
        """
        ...

    @type.setter
    def type(self, value: pulumi.Input[str]):
        ...

    @property
    @pulumi.getter(name="allocationStrategy")
    def allocation_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        The allocation strategy to use for the compute resource in case not enough instances of the best fitting instance type can be allocated. Valid items are `BEST_FIT_PROGRESSIVE`, `SPOT_CAPACITY_OPTIMIZED` or `BEST_FIT`. Defaults to `BEST_FIT`. See [AWS docs](https://docs.aws.amazon.com/batch/latest/userguide/allocation-strategies.html) for details.
        """
        ...

    @allocation_strategy.setter
    def allocation_strategy(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="bidPercentage")
    def bid_percentage(self) -> Optional[pulumi.Input[float]]:
        """
        Integer of minimum percentage that a Spot Instance price must be when compared with the On-Demand price for that instance type before instances are launched. For example, if your bid percentage is 20% (`20`), then the Spot price must be below 20% of the current On-Demand price for that EC2 instance. This parameter is required for SPOT compute environments.
        """
        ...

    @bid_percentage.setter
    def bid_percentage(self, value: Optional[pulumi.Input[float]]):
        ...

    @property
    @pulumi.getter(name="desiredVcpus")
    def desired_vcpus(self) -> Optional[pulumi.Input[float]]:
        """
        The desired number of EC2 vCPUS in the compute environment.
        """
        ...

    @desired_vcpus.setter
    def desired_vcpus(self, value: Optional[pulumi.Input[float]]):
        ...

    @property
    @pulumi.getter(name="ec2KeyPair")
    def ec2_key_pair(self) -> Optional[pulumi.Input[str]]:
        """
        The EC2 key pair that is used for instances launched in the compute environment.
        """
        ...

    @ec2_key_pair.setter
    def ec2_key_pair(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Machine Image (AMI) ID used for instances launched in the compute environment.
        """
        ...

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="launchTemplate")
    def launch_template(self) -> Optional[pulumi.Input['ComputeEnvironmentComputeResourcesLaunchTemplateArgs']]:
        """
        The launch template to use for your compute resources. See details below.
        """
        ...

    @launch_template.setter
    def launch_template(self, value: Optional[pulumi.Input['ComputeEnvironmentComputeResourcesLaunchTemplateArgs']]):
        ...

    @property
    @pulumi.getter(name="spotIamFleetRole")
    def spot_iam_fleet_role(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the Amazon EC2 Spot Fleet IAM role applied to a SPOT compute environment. This parameter is required for SPOT compute environments.
        """
        ...

    @spot_iam_fleet_role.setter
    def spot_iam_fleet_role(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value pair tags to be applied to resources that are launched in the compute environment.
        """
        ...

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        ...


@pulumi.input_type
class ComputeEnvironmentComputeResourcesLaunchTemplateArgs:
    def __init__(__self__, *,
                 launch_template_id: Optional[pulumi.Input[str]] = None,
                 launch_template_name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] launch_template_id: ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
        :param pulumi.Input[str] launch_template_name: Name of the launch template.
        :param pulumi.Input[str] version: The version number of the launch template. Default: The default version of the launch template.
        """
        if launch_template_id is not None:
            pulumi.set(__self__, "launch_template_id", launch_template_id)
        if launch_template_name is not None:
            pulumi.set(__self__, "launch_template_name", launch_template_name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="launchTemplateId")
    def launch_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the launch template. You must specify either the launch template ID or launch template name in the request, but not both.
        """
        ...

    @launch_template_id.setter
    def launch_template_id(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter(name="launchTemplateName")
    def launch_template_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the launch template.
        """
        ...

    @launch_template_name.setter
    def launch_template_name(self, value: Optional[pulumi.Input[str]]):
        ...

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version number of the launch template. Default: The default version of the launch template.
        """
        ...

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        ...


@pulumi.input_type
class JobDefinitionRetryStrategyArgs:
    def __init__(__self__, *,
                 attempts: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[float] attempts: The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
        """
        if attempts is not None:
            pulumi.set(__self__, "attempts", attempts)

    @property
    @pulumi.getter
    def attempts(self) -> Optional[pulumi.Input[float]]:
        """
        The number of times to move a job to the `RUNNABLE` status. You may specify between `1` and `10` attempts.
        """
        ...

    @attempts.setter
    def attempts(self, value: Optional[pulumi.Input[float]]):
        ...


@pulumi.input_type
class JobDefinitionTimeoutArgs:
    def __init__(__self__, *,
                 attempt_duration_seconds: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[float] attempt_duration_seconds: The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
        """
        if attempt_duration_seconds is not None:
            pulumi.set(__self__, "attempt_duration_seconds", attempt_duration_seconds)

    @property
    @pulumi.getter(name="attemptDurationSeconds")
    def attempt_duration_seconds(self) -> Optional[pulumi.Input[float]]:
        """
        The time duration in seconds after which AWS Batch terminates your jobs if they have not finished. The minimum value for the timeout is `60` seconds.
        """
        ...

    @attempt_duration_seconds.setter
    def attempt_duration_seconds(self, value: Optional[pulumi.Input[float]]):
        ...


