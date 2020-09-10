# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'CapacityProviderAutoScalingGroupProviderArgs',
    'CapacityProviderAutoScalingGroupProviderManagedScalingArgs',
    'ClusterDefaultCapacityProviderStrategyArgs',
    'ClusterSettingArgs',
    'ServiceCapacityProviderStrategyArgs',
    'ServiceDeploymentControllerArgs',
    'ServiceLoadBalancerArgs',
    'ServiceNetworkConfigurationArgs',
    'ServiceOrderedPlacementStrategyArgs',
    'ServicePlacementConstraintArgs',
    'ServiceServiceRegistriesArgs',
    'TaskDefinitionInferenceAcceleratorArgs',
    'TaskDefinitionPlacementConstraintArgs',
    'TaskDefinitionProxyConfigurationArgs',
    'TaskDefinitionVolumeArgs',
    'TaskDefinitionVolumeDockerVolumeConfigurationArgs',
    'TaskDefinitionVolumeEfsVolumeConfigurationArgs',
    'TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs',
]

@pulumi.input_type
class CapacityProviderAutoScalingGroupProviderArgs:
    def __init__(__self__, *,
                 auto_scaling_group_arn: pulumi.Input[str],
                 managed_scaling: Optional[pulumi.Input['CapacityProviderAutoScalingGroupProviderManagedScalingArgs']] = None,
                 managed_termination_protection: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "auto_scaling_group_arn", auto_scaling_group_arn)
        if managed_scaling is not None:
            pulumi.set(__self__, "managed_scaling", managed_scaling)
        if managed_termination_protection is not None:
            pulumi.set(__self__, "managed_termination_protection", managed_termination_protection)

    @property
    @pulumi.getter(name="autoScalingGroupArn")
    def auto_scaling_group_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "auto_scaling_group_arn")

    @auto_scaling_group_arn.setter
    def auto_scaling_group_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "auto_scaling_group_arn", value)

    @property
    @pulumi.getter(name="managedScaling")
    def managed_scaling(self) -> Optional[pulumi.Input['CapacityProviderAutoScalingGroupProviderManagedScalingArgs']]:
        return pulumi.get(self, "managed_scaling")

    @managed_scaling.setter
    def managed_scaling(self, value: Optional[pulumi.Input['CapacityProviderAutoScalingGroupProviderManagedScalingArgs']]):
        pulumi.set(self, "managed_scaling", value)

    @property
    @pulumi.getter(name="managedTerminationProtection")
    def managed_termination_protection(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "managed_termination_protection")

    @managed_termination_protection.setter
    def managed_termination_protection(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_termination_protection", value)


@pulumi.input_type
class CapacityProviderAutoScalingGroupProviderManagedScalingArgs:
    def __init__(__self__, *,
                 maximum_scaling_step_size: Optional[pulumi.Input[float]] = None,
                 minimum_scaling_step_size: Optional[pulumi.Input[float]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 target_capacity: Optional[pulumi.Input[float]] = None):
        if maximum_scaling_step_size is not None:
            pulumi.set(__self__, "maximum_scaling_step_size", maximum_scaling_step_size)
        if minimum_scaling_step_size is not None:
            pulumi.set(__self__, "minimum_scaling_step_size", minimum_scaling_step_size)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if target_capacity is not None:
            pulumi.set(__self__, "target_capacity", target_capacity)

    @property
    @pulumi.getter(name="maximumScalingStepSize")
    def maximum_scaling_step_size(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "maximum_scaling_step_size")

    @maximum_scaling_step_size.setter
    def maximum_scaling_step_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "maximum_scaling_step_size", value)

    @property
    @pulumi.getter(name="minimumScalingStepSize")
    def minimum_scaling_step_size(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "minimum_scaling_step_size")

    @minimum_scaling_step_size.setter
    def minimum_scaling_step_size(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "minimum_scaling_step_size", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="targetCapacity")
    def target_capacity(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "target_capacity")

    @target_capacity.setter
    def target_capacity(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "target_capacity", value)


@pulumi.input_type
class ClusterDefaultCapacityProviderStrategyArgs:
    def __init__(__self__, *,
                 capacity_provider: pulumi.Input[str],
                 base: Optional[pulumi.Input[float]] = None,
                 weight: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "capacity_provider", capacity_provider)
        if base is not None:
            pulumi.set(__self__, "base", base)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="capacityProvider")
    def capacity_provider(self) -> pulumi.Input[str]:
        return pulumi.get(self, "capacity_provider")

    @capacity_provider.setter
    def capacity_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "capacity_provider", value)

    @property
    @pulumi.getter
    def base(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "base")

    @base.setter
    def base(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "base", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class ClusterSettingArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ServiceCapacityProviderStrategyArgs:
    def __init__(__self__, *,
                 capacity_provider: pulumi.Input[str],
                 base: Optional[pulumi.Input[float]] = None,
                 weight: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "capacity_provider", capacity_provider)
        if base is not None:
            pulumi.set(__self__, "base", base)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="capacityProvider")
    def capacity_provider(self) -> pulumi.Input[str]:
        return pulumi.get(self, "capacity_provider")

    @capacity_provider.setter
    def capacity_provider(self, value: pulumi.Input[str]):
        pulumi.set(self, "capacity_provider", value)

    @property
    @pulumi.getter
    def base(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "base")

    @base.setter
    def base(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "base", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class ServiceDeploymentControllerArgs:
    def __init__(__self__, *,
                 type: Optional[pulumi.Input[str]] = None):
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ServiceLoadBalancerArgs:
    def __init__(__self__, *,
                 container_name: pulumi.Input[str],
                 container_port: pulumi.Input[float],
                 elb_name: Optional[pulumi.Input[str]] = None,
                 target_group_arn: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "container_name", container_name)
        pulumi.set(__self__, "container_port", container_port)
        if elb_name is not None:
            pulumi.set(__self__, "elb_name", elb_name)
        if target_group_arn is not None:
            pulumi.set(__self__, "target_group_arn", target_group_arn)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "container_name")

    @container_name.setter
    def container_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_name", value)

    @property
    @pulumi.getter(name="containerPort")
    def container_port(self) -> pulumi.Input[float]:
        return pulumi.get(self, "container_port")

    @container_port.setter
    def container_port(self, value: pulumi.Input[float]):
        pulumi.set(self, "container_port", value)

    @property
    @pulumi.getter(name="elbName")
    def elb_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "elb_name")

    @elb_name.setter
    def elb_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "elb_name", value)

    @property
    @pulumi.getter(name="targetGroupArn")
    def target_group_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_group_arn")

    @target_group_arn.setter
    def target_group_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_group_arn", value)


@pulumi.input_type
class ServiceNetworkConfigurationArgs:
    def __init__(__self__, *,
                 subnets: pulumi.Input[List[pulumi.Input[str]]],
                 assign_public_ip: Optional[pulumi.Input[bool]] = None,
                 security_groups: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "subnets", subnets)
        if assign_public_ip is not None:
            pulumi.set(__self__, "assign_public_ip", assign_public_ip)
        if security_groups is not None:
            pulumi.set(__self__, "security_groups", security_groups)

    @property
    @pulumi.getter
    def subnets(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        return pulumi.get(self, "subnets")

    @subnets.setter
    def subnets(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "subnets", value)

    @property
    @pulumi.getter(name="assignPublicIp")
    def assign_public_ip(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "assign_public_ip")

    @assign_public_ip.setter
    def assign_public_ip(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "assign_public_ip", value)

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_groups")

    @security_groups.setter
    def security_groups(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "security_groups", value)


@pulumi.input_type
class ServiceOrderedPlacementStrategyArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 field: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if field is not None:
            pulumi.set(__self__, "field", field)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def field(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "field")

    @field.setter
    def field(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "field", value)


@pulumi.input_type
class ServicePlacementConstraintArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 expression: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if expression is not None:
            pulumi.set(__self__, "expression", expression)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)


@pulumi.input_type
class ServiceServiceRegistriesArgs:
    def __init__(__self__, *,
                 registry_arn: pulumi.Input[str],
                 container_name: Optional[pulumi.Input[str]] = None,
                 container_port: Optional[pulumi.Input[float]] = None,
                 port: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "registry_arn", registry_arn)
        if container_name is not None:
            pulumi.set(__self__, "container_name", container_name)
        if container_port is not None:
            pulumi.set(__self__, "container_port", container_port)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="registryArn")
    def registry_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "registry_arn")

    @registry_arn.setter
    def registry_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry_arn", value)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "container_name")

    @container_name.setter
    def container_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_name", value)

    @property
    @pulumi.getter(name="containerPort")
    def container_port(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "container_port")

    @container_port.setter
    def container_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "container_port", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "port", value)


@pulumi.input_type
class TaskDefinitionInferenceAcceleratorArgs:
    def __init__(__self__, *,
                 device_name: pulumi.Input[str],
                 device_type: pulumi.Input[str]):
        pulumi.set(__self__, "device_name", device_name)
        pulumi.set(__self__, "device_type", device_type)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter(name="deviceType")
    def device_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "device_type")

    @device_type.setter
    def device_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_type", value)


@pulumi.input_type
class TaskDefinitionPlacementConstraintArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 expression: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if expression is not None:
            pulumi.set(__self__, "expression", expression)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def expression(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expression", value)


@pulumi.input_type
class TaskDefinitionProxyConfigurationArgs:
    def __init__(__self__, *,
                 container_name: pulumi.Input[str],
                 properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "container_name", container_name)
        if properties is not None:
            pulumi.set(__self__, "properties", properties)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "container_name")

    @container_name.setter
    def container_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_name", value)

    @property
    @pulumi.getter
    def properties(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class TaskDefinitionVolumeArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 docker_volume_configuration: Optional[pulumi.Input['TaskDefinitionVolumeDockerVolumeConfigurationArgs']] = None,
                 efs_volume_configuration: Optional[pulumi.Input['TaskDefinitionVolumeEfsVolumeConfigurationArgs']] = None,
                 host_path: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        if docker_volume_configuration is not None:
            pulumi.set(__self__, "docker_volume_configuration", docker_volume_configuration)
        if efs_volume_configuration is not None:
            pulumi.set(__self__, "efs_volume_configuration", efs_volume_configuration)
        if host_path is not None:
            pulumi.set(__self__, "host_path", host_path)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="dockerVolumeConfiguration")
    def docker_volume_configuration(self) -> Optional[pulumi.Input['TaskDefinitionVolumeDockerVolumeConfigurationArgs']]:
        return pulumi.get(self, "docker_volume_configuration")

    @docker_volume_configuration.setter
    def docker_volume_configuration(self, value: Optional[pulumi.Input['TaskDefinitionVolumeDockerVolumeConfigurationArgs']]):
        pulumi.set(self, "docker_volume_configuration", value)

    @property
    @pulumi.getter(name="efsVolumeConfiguration")
    def efs_volume_configuration(self) -> Optional[pulumi.Input['TaskDefinitionVolumeEfsVolumeConfigurationArgs']]:
        return pulumi.get(self, "efs_volume_configuration")

    @efs_volume_configuration.setter
    def efs_volume_configuration(self, value: Optional[pulumi.Input['TaskDefinitionVolumeEfsVolumeConfigurationArgs']]):
        pulumi.set(self, "efs_volume_configuration", value)

    @property
    @pulumi.getter(name="hostPath")
    def host_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "host_path")

    @host_path.setter
    def host_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "host_path", value)


@pulumi.input_type
class TaskDefinitionVolumeDockerVolumeConfigurationArgs:
    def __init__(__self__, *,
                 autoprovision: Optional[pulumi.Input[bool]] = None,
                 driver: Optional[pulumi.Input[str]] = None,
                 driver_opts: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        if autoprovision is not None:
            pulumi.set(__self__, "autoprovision", autoprovision)
        if driver is not None:
            pulumi.set(__self__, "driver", driver)
        if driver_opts is not None:
            pulumi.set(__self__, "driver_opts", driver_opts)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def autoprovision(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "autoprovision")

    @autoprovision.setter
    def autoprovision(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "autoprovision", value)

    @property
    @pulumi.getter
    def driver(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "driver")

    @driver.setter
    def driver(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "driver", value)

    @property
    @pulumi.getter(name="driverOpts")
    def driver_opts(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "driver_opts")

    @driver_opts.setter
    def driver_opts(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "driver_opts", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


@pulumi.input_type
class TaskDefinitionVolumeEfsVolumeConfigurationArgs:
    def __init__(__self__, *,
                 file_system_id: pulumi.Input[str],
                 authorization_config: Optional[pulumi.Input['TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs']] = None,
                 root_directory: Optional[pulumi.Input[str]] = None,
                 transit_encryption: Optional[pulumi.Input[str]] = None,
                 transit_encryption_port: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "file_system_id", file_system_id)
        if authorization_config is not None:
            pulumi.set(__self__, "authorization_config", authorization_config)
        if root_directory is not None:
            pulumi.set(__self__, "root_directory", root_directory)
        if transit_encryption is not None:
            pulumi.set(__self__, "transit_encryption", transit_encryption)
        if transit_encryption_port is not None:
            pulumi.set(__self__, "transit_encryption_port", transit_encryption_port)

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "file_system_id")

    @file_system_id.setter
    def file_system_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_system_id", value)

    @property
    @pulumi.getter(name="authorizationConfig")
    def authorization_config(self) -> Optional[pulumi.Input['TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs']]:
        return pulumi.get(self, "authorization_config")

    @authorization_config.setter
    def authorization_config(self, value: Optional[pulumi.Input['TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs']]):
        pulumi.set(self, "authorization_config", value)

    @property
    @pulumi.getter(name="rootDirectory")
    def root_directory(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "root_directory")

    @root_directory.setter
    def root_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_directory", value)

    @property
    @pulumi.getter(name="transitEncryption")
    def transit_encryption(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "transit_encryption")

    @transit_encryption.setter
    def transit_encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_encryption", value)

    @property
    @pulumi.getter(name="transitEncryptionPort")
    def transit_encryption_port(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "transit_encryption_port")

    @transit_encryption_port.setter
    def transit_encryption_port(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "transit_encryption_port", value)


@pulumi.input_type
class TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs:
    def __init__(__self__, *,
                 access_point_id: Optional[pulumi.Input[str]] = None,
                 iam: Optional[pulumi.Input[str]] = None):
        if access_point_id is not None:
            pulumi.set(__self__, "access_point_id", access_point_id)
        if iam is not None:
            pulumi.set(__self__, "iam", iam)

    @property
    @pulumi.getter(name="accessPointId")
    def access_point_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_point_id")

    @access_point_id.setter
    def access_point_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_point_id", value)

    @property
    @pulumi.getter
    def iam(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "iam")

    @iam.setter
    def iam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam", value)


