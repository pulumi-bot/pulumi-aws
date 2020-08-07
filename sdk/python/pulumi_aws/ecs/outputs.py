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
    'CapacityProviderAutoScalingGroupProvider',
    'CapacityProviderAutoScalingGroupProviderManagedScaling',
    'ClusterDefaultCapacityProviderStrategy',
    'ClusterSetting',
    'ServiceCapacityProviderStrategy',
    'ServiceDeploymentController',
    'ServiceLoadBalancer',
    'ServiceNetworkConfiguration',
    'ServiceOrderedPlacementStrategy',
    'ServicePlacementConstraint',
    'ServiceServiceRegistries',
    'TaskDefinitionInferenceAccelerator',
    'TaskDefinitionPlacementConstraint',
    'TaskDefinitionProxyConfiguration',
    'TaskDefinitionVolume',
    'TaskDefinitionVolumeDockerVolumeConfiguration',
    'TaskDefinitionVolumeEfsVolumeConfiguration',
    'TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig',
    'GetClusterSetting',
]

@pulumi.output_type
class CapacityProviderAutoScalingGroupProvider(dict):
    @property
    @pulumi.getter(name="autoScalingGroupArn")
    def auto_scaling_group_arn(self) -> str:
        """
        - The Amazon Resource Name (ARN) of the associated auto scaling group.
        """
        ...

    @property
    @pulumi.getter(name="managedScaling")
    def managed_scaling(self) -> Optional['outputs.CapacityProviderAutoScalingGroupProviderManagedScaling']:
        """
        - Nested argument defining the parameters of the auto scaling. Defined below.
        """
        ...

    @property
    @pulumi.getter(name="managedTerminationProtection")
    def managed_termination_protection(self) -> Optional[str]:
        """
        - Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are `ENABLED` and `DISABLED`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CapacityProviderAutoScalingGroupProviderManagedScaling(dict):
    @property
    @pulumi.getter(name="maximumScalingStepSize")
    def maximum_scaling_step_size(self) -> Optional[float]:
        """
        The maximum step adjustment size. A number between 1 and 10,000.
        """
        ...

    @property
    @pulumi.getter(name="minimumScalingStepSize")
    def minimum_scaling_step_size(self) -> Optional[float]:
        """
        The minimum step adjustment size. A number between 1 and 10,000.
        """
        ...

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Whether auto scaling is managed by ECS. Valid values are `ENABLED` and `DISABLED`.
        """
        ...

    @property
    @pulumi.getter(name="targetCapacity")
    def target_capacity(self) -> Optional[float]:
        """
        The target utilization for the capacity provider. A number between 1 and 100.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterDefaultCapacityProviderStrategy(dict):
    @property
    @pulumi.getter
    def base(self) -> Optional[float]:
        """
        The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
        """
        ...

    @property
    @pulumi.getter(name="capacityProvider")
    def capacity_provider(self) -> str:
        """
        The short name of the capacity provider.
        """
        ...

    @property
    @pulumi.getter
    def weight(self) -> Optional[float]:
        """
        The relative percentage of the total number of launched tasks that should use the specified capacity provider.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterSetting(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the setting to manage. Valid values: `containerInsights`.
        """
        ...

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        The value to assign to the setting. Value values are `enabled` and `disabled`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceCapacityProviderStrategy(dict):
    @property
    @pulumi.getter
    def base(self) -> Optional[float]:
        """
        The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
        """
        ...

    @property
    @pulumi.getter(name="capacityProvider")
    def capacity_provider(self) -> str:
        """
        The short name of the capacity provider.
        """
        ...

    @property
    @pulumi.getter
    def weight(self) -> Optional[float]:
        """
        The relative percentage of the total number of launched tasks that should use the specified capacity provider.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceDeploymentController(dict):
    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Type of deployment controller. Valid values: `CODE_DEPLOY`, `ECS`, `EXTERNAL`. Default: `ECS`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceLoadBalancer(dict):
    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> str:
        """
        The name of the container to associate with the load balancer (as it appears in a container definition).
        """
        ...

    @property
    @pulumi.getter(name="containerPort")
    def container_port(self) -> float:
        """
        The port on the container to associate with the load balancer.
        """
        ...

    @property
    @pulumi.getter(name="elbName")
    def elb_name(self) -> Optional[str]:
        """
        The name of the ELB (Classic) to associate with the service.
        """
        ...

    @property
    @pulumi.getter(name="targetGroupArn")
    def target_group_arn(self) -> Optional[str]:
        """
        The ARN of the Load Balancer target group to associate with the service.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceNetworkConfiguration(dict):
    @property
    @pulumi.getter(name="assignPublicIp")
    def assign_public_ip(self) -> Optional[bool]:
        """
        Assign a public IP address to the ENI (Fargate launch type only). Valid values are `true` or `false`. Default `false`.
        """
        ...

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Optional[List[str]]:
        """
        The security groups associated with the task or service. If you do not specify a security group, the default security group for the VPC is used.
        """
        ...

    @property
    @pulumi.getter
    def subnets(self) -> List[str]:
        """
        The subnets associated with the task or service.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceOrderedPlacementStrategy(dict):
    @property
    @pulumi.getter
    def field(self) -> Optional[str]:
        """
        For the `spread` placement strategy, valid values are `instanceId` (or `host`,
        which has the same effect), or any platform or custom attribute that is applied to a container instance.
        For the `binpack` type, valid values are `memory` and `cpu`. For the `random` type, this attribute is not
        needed. For more information, see [Placement Strategy](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_PlacementStrategy.html).
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of placement strategy. Must be one of: `binpack`, `random`, or `spread`
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServicePlacementConstraint(dict):
    @property
    @pulumi.getter
    def expression(self) -> Optional[str]:
        """
        Cluster Query Language expression to apply to the constraint. Does not need to be specified
        for the `distinctInstance` type.
        For more information, see [Cluster Query Language in the Amazon EC2 Container
        Service Developer
        Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of constraint. The only valid values at this time are `memberOf` and `distinctInstance`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceServiceRegistries(dict):
    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> Optional[str]:
        """
        The container name value, already specified in the task definition, to be used for your service discovery service.
        """
        ...

    @property
    @pulumi.getter(name="containerPort")
    def container_port(self) -> Optional[float]:
        """
        The port value, already specified in the task definition, to be used for your service discovery service.
        """
        ...

    @property
    @pulumi.getter
    def port(self) -> Optional[float]:
        """
        The port value used if your Service Discovery service specified an SRV record.
        """
        ...

    @property
    @pulumi.getter(name="registryArn")
    def registry_arn(self) -> str:
        """
        The ARN of the Service Registry. The currently supported service registry is Amazon Route 53 Auto Naming Service(`servicediscovery.Service`). For more information, see [Service](https://docs.aws.amazon.com/Route53/latest/APIReference/API_autonaming_Service.html)
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TaskDefinitionInferenceAccelerator(dict):
    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> str:
        """
        The Elastic Inference accelerator device name. The deviceName must also be referenced in a container definition as a ResourceRequirement.
        """
        ...

    @property
    @pulumi.getter(name="deviceType")
    def device_type(self) -> str:
        """
        The Elastic Inference accelerator type to use.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TaskDefinitionPlacementConstraint(dict):
    @property
    @pulumi.getter
    def expression(self) -> Optional[str]:
        """
        Cluster Query Language expression to apply to the constraint.
        For more information, see [Cluster Query Language in the Amazon EC2 Container
        Service Developer
        Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TaskDefinitionProxyConfiguration(dict):
    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> str:
        """
        The name of the container that will serve as the App Mesh proxy.
        """
        ...

    @property
    @pulumi.getter
    def properties(self) -> Optional[Mapping[str, str]]:
        """
        The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TaskDefinitionVolume(dict):
    @property
    @pulumi.getter(name="dockerVolumeConfiguration")
    def docker_volume_configuration(self) -> Optional['outputs.TaskDefinitionVolumeDockerVolumeConfiguration']:
        """
        Used to configure a docker volume
        """
        ...

    @property
    @pulumi.getter(name="efsVolumeConfiguration")
    def efs_volume_configuration(self) -> Optional['outputs.TaskDefinitionVolumeEfsVolumeConfiguration']:
        """
        Used to configure a EFS volume.
        """
        ...

    @property
    @pulumi.getter(name="hostPath")
    def host_path(self) -> Optional[str]:
        """
        The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the volume. This name is referenced in the `sourceVolume`
        parameter of container definition in the `mountPoints` section.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TaskDefinitionVolumeDockerVolumeConfiguration(dict):
    @property
    @pulumi.getter
    def autoprovision(self) -> Optional[bool]:
        """
        If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
        """
        ...

    @property
    @pulumi.getter
    def driver(self) -> Optional[str]:
        """
        The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
        """
        ...

    @property
    @pulumi.getter(name="driverOpts")
    def driver_opts(self) -> Optional[Mapping[str, str]]:
        """
        A map of Docker driver specific options.
        """
        ...

    @property
    @pulumi.getter
    def labels(self) -> Optional[Mapping[str, str]]:
        """
        A map of custom metadata to add to your Docker volume.
        """
        ...

    @property
    @pulumi.getter
    def scope(self) -> Optional[str]:
        """
        The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TaskDefinitionVolumeEfsVolumeConfiguration(dict):
    @property
    @pulumi.getter(name="authorizationConfig")
    def authorization_config(self) -> Optional['outputs.TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig']:
        """
        The authorization configuration details for the Amazon EFS file system.
        """
        ...

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        """
        The ID of the EFS File System.
        """
        ...

    @property
    @pulumi.getter(name="rootDirectory")
    def root_directory(self) -> Optional[str]:
        """
        The directory within the Amazon EFS file system to mount as the root directory inside the host. If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying / will have the same effect as omitting this parameter. This argument is ignored when using `authorization_config`.
        """
        ...

    @property
    @pulumi.getter(name="transitEncryption")
    def transit_encryption(self) -> Optional[str]:
        """
        Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server. Transit encryption must be enabled if Amazon EFS IAM authorization is used. Valid values: `ENABLED`, `DISABLED`. If this parameter is omitted, the default value of `DISABLED` is used.
        """
        ...

    @property
    @pulumi.getter(name="transitEncryptionPort")
    def transit_encryption_port(self) -> Optional[float]:
        """
        The port to use for transit encryption. If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig(dict):
    @property
    @pulumi.getter(name="accessPointId")
    def access_point_id(self) -> Optional[str]:
        """
        The access point ID to use. If an access point is specified, the root directory value will be relative to the directory set for the access point. If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
        """
        ...

    @property
    @pulumi.getter
    def iam(self) -> Optional[str]:
        """
        Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system. If enabled, transit encryption must be enabled in the EFSVolumeConfiguration. Valid values: `ENABLED`, `DISABLED`. If this parameter is omitted, the default value of `DISABLED` is used.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetClusterSetting(dict):
    @property
    @pulumi.getter
    def name(self) -> str:
        ...

    @property
    @pulumi.getter
    def value(self) -> str:
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


