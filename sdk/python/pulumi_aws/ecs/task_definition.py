# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['TaskDefinitionArgs', 'TaskDefinition']

@pulumi.input_type
class TaskDefinitionArgs:
    def __init__(__self__, *,
                 container_definitions: pulumi.Input[str],
                 family: pulumi.Input[str],
                 cpu: Optional[pulumi.Input[str]] = None,
                 execution_role_arn: Optional[pulumi.Input[str]] = None,
                 inference_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionInferenceAcceleratorArgs']]]] = None,
                 ipc_mode: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[str]] = None,
                 network_mode: Optional[pulumi.Input[str]] = None,
                 pid_mode: Optional[pulumi.Input[str]] = None,
                 placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionPlacementConstraintArgs']]]] = None,
                 proxy_configuration: Optional[pulumi.Input['TaskDefinitionProxyConfigurationArgs']] = None,
                 requires_compatibilities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 task_role_arn: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionVolumeArgs']]]] = None):
        """
        The set of arguments for constructing a TaskDefinition resource.
        :param pulumi.Input[str] container_definitions: A list of valid [container
               definitions](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html)
               provided as a single valid JSON document. Please note that you should only
               provide values that are part of the container definition document. For a
               detailed description of what parameters are available, see the [Task Definition
               Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html)
               section from the official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
        :param pulumi.Input[str] family: A unique name for your task definition.
        :param pulumi.Input[str] cpu: The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] execution_role_arn: The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
        :param pulumi.Input[Sequence[pulumi.Input['TaskDefinitionInferenceAcceleratorArgs']]] inference_accelerators: Configuration block(s) with Inference Accelerators settings. Detailed below.
        :param pulumi.Input[str] ipc_mode: The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
        :param pulumi.Input[str] memory: The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] network_mode: The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
        :param pulumi.Input[str] pid_mode: The process namespace to use for the containers in the task. The valid values are `host` and `task`.
        :param pulumi.Input[Sequence[pulumi.Input['TaskDefinitionPlacementConstraintArgs']]] placement_constraints: A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
        :param pulumi.Input['TaskDefinitionProxyConfigurationArgs'] proxy_configuration: The proxy configuration details for the App Mesh proxy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requires_compatibilities: A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] task_role_arn: The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
        :param pulumi.Input[Sequence[pulumi.Input['TaskDefinitionVolumeArgs']]] volumes: A set of volume blocks that containers in your task may use.
        """
        pulumi.set(__self__, "container_definitions", container_definitions)
        pulumi.set(__self__, "family", family)
        if cpu is not None:
            pulumi.set(__self__, "cpu", cpu)
        if execution_role_arn is not None:
            pulumi.set(__self__, "execution_role_arn", execution_role_arn)
        if inference_accelerators is not None:
            pulumi.set(__self__, "inference_accelerators", inference_accelerators)
        if ipc_mode is not None:
            pulumi.set(__self__, "ipc_mode", ipc_mode)
        if memory is not None:
            pulumi.set(__self__, "memory", memory)
        if network_mode is not None:
            pulumi.set(__self__, "network_mode", network_mode)
        if pid_mode is not None:
            pulumi.set(__self__, "pid_mode", pid_mode)
        if placement_constraints is not None:
            pulumi.set(__self__, "placement_constraints", placement_constraints)
        if proxy_configuration is not None:
            pulumi.set(__self__, "proxy_configuration", proxy_configuration)
        if requires_compatibilities is not None:
            pulumi.set(__self__, "requires_compatibilities", requires_compatibilities)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if task_role_arn is not None:
            pulumi.set(__self__, "task_role_arn", task_role_arn)
        if volumes is not None:
            pulumi.set(__self__, "volumes", volumes)

    @property
    @pulumi.getter(name="containerDefinitions")
    def container_definitions(self) -> pulumi.Input[str]:
        """
        A list of valid [container
        definitions](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html)
        provided as a single valid JSON document. Please note that you should only
        provide values that are part of the container definition document. For a
        detailed description of what parameters are available, see the [Task Definition
        Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html)
        section from the official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
        """
        return pulumi.get(self, "container_definitions")

    @container_definitions.setter
    def container_definitions(self, value: pulumi.Input[str]):
        pulumi.set(self, "container_definitions", value)

    @property
    @pulumi.getter
    def family(self) -> pulumi.Input[str]:
        """
        A unique name for your task definition.
        """
        return pulumi.get(self, "family")

    @family.setter
    def family(self, value: pulumi.Input[str]):
        pulumi.set(self, "family", value)

    @property
    @pulumi.getter
    def cpu(self) -> Optional[pulumi.Input[str]]:
        """
        The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        """
        return pulumi.get(self, "cpu")

    @cpu.setter
    def cpu(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cpu", value)

    @property
    @pulumi.getter(name="executionRoleArn")
    def execution_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
        """
        return pulumi.get(self, "execution_role_arn")

    @execution_role_arn.setter
    def execution_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "execution_role_arn", value)

    @property
    @pulumi.getter(name="inferenceAccelerators")
    def inference_accelerators(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionInferenceAcceleratorArgs']]]]:
        """
        Configuration block(s) with Inference Accelerators settings. Detailed below.
        """
        return pulumi.get(self, "inference_accelerators")

    @inference_accelerators.setter
    def inference_accelerators(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionInferenceAcceleratorArgs']]]]):
        pulumi.set(self, "inference_accelerators", value)

    @property
    @pulumi.getter(name="ipcMode")
    def ipc_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
        """
        return pulumi.get(self, "ipc_mode")

    @ipc_mode.setter
    def ipc_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ipc_mode", value)

    @property
    @pulumi.getter
    def memory(self) -> Optional[pulumi.Input[str]]:
        """
        The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        """
        return pulumi.get(self, "memory")

    @memory.setter
    def memory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "memory", value)

    @property
    @pulumi.getter(name="networkMode")
    def network_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
        """
        return pulumi.get(self, "network_mode")

    @network_mode.setter
    def network_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_mode", value)

    @property
    @pulumi.getter(name="pidMode")
    def pid_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The process namespace to use for the containers in the task. The valid values are `host` and `task`.
        """
        return pulumi.get(self, "pid_mode")

    @pid_mode.setter
    def pid_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pid_mode", value)

    @property
    @pulumi.getter(name="placementConstraints")
    def placement_constraints(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionPlacementConstraintArgs']]]]:
        """
        A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
        """
        return pulumi.get(self, "placement_constraints")

    @placement_constraints.setter
    def placement_constraints(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionPlacementConstraintArgs']]]]):
        pulumi.set(self, "placement_constraints", value)

    @property
    @pulumi.getter(name="proxyConfiguration")
    def proxy_configuration(self) -> Optional[pulumi.Input['TaskDefinitionProxyConfigurationArgs']]:
        """
        The proxy configuration details for the App Mesh proxy.
        """
        return pulumi.get(self, "proxy_configuration")

    @proxy_configuration.setter
    def proxy_configuration(self, value: Optional[pulumi.Input['TaskDefinitionProxyConfigurationArgs']]):
        pulumi.set(self, "proxy_configuration", value)

    @property
    @pulumi.getter(name="requiresCompatibilities")
    def requires_compatibilities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
        """
        return pulumi.get(self, "requires_compatibilities")

    @requires_compatibilities.setter
    def requires_compatibilities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "requires_compatibilities", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="taskRoleArn")
    def task_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
        """
        return pulumi.get(self, "task_role_arn")

    @task_role_arn.setter
    def task_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "task_role_arn", value)

    @property
    @pulumi.getter
    def volumes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionVolumeArgs']]]]:
        """
        A set of volume blocks that containers in your task may use.
        """
        return pulumi.get(self, "volumes")

    @volumes.setter
    def volumes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TaskDefinitionVolumeArgs']]]]):
        pulumi.set(self, "volumes", value)


class TaskDefinition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_definitions: Optional[pulumi.Input[str]] = None,
                 cpu: Optional[pulumi.Input[str]] = None,
                 execution_role_arn: Optional[pulumi.Input[str]] = None,
                 family: Optional[pulumi.Input[str]] = None,
                 inference_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionInferenceAcceleratorArgs']]]]] = None,
                 ipc_mode: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[str]] = None,
                 network_mode: Optional[pulumi.Input[str]] = None,
                 pid_mode: Optional[pulumi.Input[str]] = None,
                 placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionPlacementConstraintArgs']]]]] = None,
                 proxy_configuration: Optional[pulumi.Input[pulumi.InputType['TaskDefinitionProxyConfigurationArgs']]] = None,
                 requires_compatibilities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 task_role_arn: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionVolumeArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a revision of an ECS task definition to be used in `ecs.Service`.

        ## Example Usage
        ### With AppMesh Proxy

        ```python
        import pulumi
        import pulumi_aws as aws

        service = aws.ecs.TaskDefinition("service",
            family="service",
            container_definitions=(lambda path: open(path).read())("task-definitions/service.json"),
            proxy_configuration=aws.ecs.TaskDefinitionProxyConfigurationArgs(
                type="APPMESH",
                container_name="applicationContainerName",
                properties={
                    "AppPorts": "8080",
                    "EgressIgnoredIPs": "169.254.170.2,169.254.169.254",
                    "IgnoredUID": "1337",
                    "ProxyEgressPort": "15001",
                    "ProxyIngressPort": "15000",
                },
            ))
        ```

        ## Import

        ECS Task Definitions can be imported via their Amazon Resource Name (ARN)

        ```sh
         $ pulumi import aws:ecs/taskDefinition:TaskDefinition example arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] container_definitions: A list of valid [container
               definitions](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html)
               provided as a single valid JSON document. Please note that you should only
               provide values that are part of the container definition document. For a
               detailed description of what parameters are available, see the [Task Definition
               Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html)
               section from the official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
        :param pulumi.Input[str] cpu: The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] execution_role_arn: The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
        :param pulumi.Input[str] family: A unique name for your task definition.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionInferenceAcceleratorArgs']]]] inference_accelerators: Configuration block(s) with Inference Accelerators settings. Detailed below.
        :param pulumi.Input[str] ipc_mode: The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
        :param pulumi.Input[str] memory: The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] network_mode: The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
        :param pulumi.Input[str] pid_mode: The process namespace to use for the containers in the task. The valid values are `host` and `task`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionPlacementConstraintArgs']]]] placement_constraints: A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
        :param pulumi.Input[pulumi.InputType['TaskDefinitionProxyConfigurationArgs']] proxy_configuration: The proxy configuration details for the App Mesh proxy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requires_compatibilities: A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] task_role_arn: The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionVolumeArgs']]]] volumes: A set of volume blocks that containers in your task may use.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaskDefinitionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a revision of an ECS task definition to be used in `ecs.Service`.

        ## Example Usage
        ### With AppMesh Proxy

        ```python
        import pulumi
        import pulumi_aws as aws

        service = aws.ecs.TaskDefinition("service",
            family="service",
            container_definitions=(lambda path: open(path).read())("task-definitions/service.json"),
            proxy_configuration=aws.ecs.TaskDefinitionProxyConfigurationArgs(
                type="APPMESH",
                container_name="applicationContainerName",
                properties={
                    "AppPorts": "8080",
                    "EgressIgnoredIPs": "169.254.170.2,169.254.169.254",
                    "IgnoredUID": "1337",
                    "ProxyEgressPort": "15001",
                    "ProxyIngressPort": "15000",
                },
            ))
        ```

        ## Import

        ECS Task Definitions can be imported via their Amazon Resource Name (ARN)

        ```sh
         $ pulumi import aws:ecs/taskDefinition:TaskDefinition example arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
        ```

        :param str resource_name: The name of the resource.
        :param TaskDefinitionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaskDefinitionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 container_definitions: Optional[pulumi.Input[str]] = None,
                 cpu: Optional[pulumi.Input[str]] = None,
                 execution_role_arn: Optional[pulumi.Input[str]] = None,
                 family: Optional[pulumi.Input[str]] = None,
                 inference_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionInferenceAcceleratorArgs']]]]] = None,
                 ipc_mode: Optional[pulumi.Input[str]] = None,
                 memory: Optional[pulumi.Input[str]] = None,
                 network_mode: Optional[pulumi.Input[str]] = None,
                 pid_mode: Optional[pulumi.Input[str]] = None,
                 placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionPlacementConstraintArgs']]]]] = None,
                 proxy_configuration: Optional[pulumi.Input[pulumi.InputType['TaskDefinitionProxyConfigurationArgs']]] = None,
                 requires_compatibilities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 task_role_arn: Optional[pulumi.Input[str]] = None,
                 volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionVolumeArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if container_definitions is None and not opts.urn:
                raise TypeError("Missing required property 'container_definitions'")
            __props__['container_definitions'] = container_definitions
            __props__['cpu'] = cpu
            __props__['execution_role_arn'] = execution_role_arn
            if family is None and not opts.urn:
                raise TypeError("Missing required property 'family'")
            __props__['family'] = family
            __props__['inference_accelerators'] = inference_accelerators
            __props__['ipc_mode'] = ipc_mode
            __props__['memory'] = memory
            __props__['network_mode'] = network_mode
            __props__['pid_mode'] = pid_mode
            __props__['placement_constraints'] = placement_constraints
            __props__['proxy_configuration'] = proxy_configuration
            __props__['requires_compatibilities'] = requires_compatibilities
            __props__['tags'] = tags
            __props__['task_role_arn'] = task_role_arn
            __props__['volumes'] = volumes
            __props__['arn'] = None
            __props__['revision'] = None
        super(TaskDefinition, __self__).__init__(
            'aws:ecs/taskDefinition:TaskDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            container_definitions: Optional[pulumi.Input[str]] = None,
            cpu: Optional[pulumi.Input[str]] = None,
            execution_role_arn: Optional[pulumi.Input[str]] = None,
            family: Optional[pulumi.Input[str]] = None,
            inference_accelerators: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionInferenceAcceleratorArgs']]]]] = None,
            ipc_mode: Optional[pulumi.Input[str]] = None,
            memory: Optional[pulumi.Input[str]] = None,
            network_mode: Optional[pulumi.Input[str]] = None,
            pid_mode: Optional[pulumi.Input[str]] = None,
            placement_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionPlacementConstraintArgs']]]]] = None,
            proxy_configuration: Optional[pulumi.Input[pulumi.InputType['TaskDefinitionProxyConfigurationArgs']]] = None,
            requires_compatibilities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            revision: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            task_role_arn: Optional[pulumi.Input[str]] = None,
            volumes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionVolumeArgs']]]]] = None) -> 'TaskDefinition':
        """
        Get an existing TaskDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Full ARN of the Task Definition (including both `family` and `revision`).
        :param pulumi.Input[str] container_definitions: A list of valid [container
               definitions](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html)
               provided as a single valid JSON document. Please note that you should only
               provide values that are part of the container definition document. For a
               detailed description of what parameters are available, see the [Task Definition
               Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html)
               section from the official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
        :param pulumi.Input[str] cpu: The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] execution_role_arn: The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
        :param pulumi.Input[str] family: A unique name for your task definition.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionInferenceAcceleratorArgs']]]] inference_accelerators: Configuration block(s) with Inference Accelerators settings. Detailed below.
        :param pulumi.Input[str] ipc_mode: The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
        :param pulumi.Input[str] memory: The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        :param pulumi.Input[str] network_mode: The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
        :param pulumi.Input[str] pid_mode: The process namespace to use for the containers in the task. The valid values are `host` and `task`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionPlacementConstraintArgs']]]] placement_constraints: A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
        :param pulumi.Input[pulumi.InputType['TaskDefinitionProxyConfigurationArgs']] proxy_configuration: The proxy configuration details for the App Mesh proxy.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] requires_compatibilities: A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
        :param pulumi.Input[int] revision: The revision of the task in a particular family.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] task_role_arn: The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TaskDefinitionVolumeArgs']]]] volumes: A set of volume blocks that containers in your task may use.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["container_definitions"] = container_definitions
        __props__["cpu"] = cpu
        __props__["execution_role_arn"] = execution_role_arn
        __props__["family"] = family
        __props__["inference_accelerators"] = inference_accelerators
        __props__["ipc_mode"] = ipc_mode
        __props__["memory"] = memory
        __props__["network_mode"] = network_mode
        __props__["pid_mode"] = pid_mode
        __props__["placement_constraints"] = placement_constraints
        __props__["proxy_configuration"] = proxy_configuration
        __props__["requires_compatibilities"] = requires_compatibilities
        __props__["revision"] = revision
        __props__["tags"] = tags
        __props__["task_role_arn"] = task_role_arn
        __props__["volumes"] = volumes
        return TaskDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Full ARN of the Task Definition (including both `family` and `revision`).
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="containerDefinitions")
    def container_definitions(self) -> pulumi.Output[str]:
        """
        A list of valid [container
        definitions](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html)
        provided as a single valid JSON document. Please note that you should only
        provide values that are part of the container definition document. For a
        detailed description of what parameters are available, see the [Task Definition
        Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html)
        section from the official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
        """
        return pulumi.get(self, "container_definitions")

    @property
    @pulumi.getter
    def cpu(self) -> pulumi.Output[Optional[str]]:
        """
        The number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter(name="executionRoleArn")
    def execution_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
        """
        return pulumi.get(self, "execution_role_arn")

    @property
    @pulumi.getter
    def family(self) -> pulumi.Output[str]:
        """
        A unique name for your task definition.
        """
        return pulumi.get(self, "family")

    @property
    @pulumi.getter(name="inferenceAccelerators")
    def inference_accelerators(self) -> pulumi.Output[Optional[Sequence['outputs.TaskDefinitionInferenceAccelerator']]]:
        """
        Configuration block(s) with Inference Accelerators settings. Detailed below.
        """
        return pulumi.get(self, "inference_accelerators")

    @property
    @pulumi.getter(name="ipcMode")
    def ipc_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
        """
        return pulumi.get(self, "ipc_mode")

    @property
    @pulumi.getter
    def memory(self) -> pulumi.Output[Optional[str]]:
        """
        The amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter(name="networkMode")
    def network_mode(self) -> pulumi.Output[str]:
        """
        The Docker networking mode to use for the containers in the task. The valid values are `none`, `bridge`, `awsvpc`, and `host`.
        """
        return pulumi.get(self, "network_mode")

    @property
    @pulumi.getter(name="pidMode")
    def pid_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The process namespace to use for the containers in the task. The valid values are `host` and `task`.
        """
        return pulumi.get(self, "pid_mode")

    @property
    @pulumi.getter(name="placementConstraints")
    def placement_constraints(self) -> pulumi.Output[Optional[Sequence['outputs.TaskDefinitionPlacementConstraint']]]:
        """
        A set of placement constraints rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`.
        """
        return pulumi.get(self, "placement_constraints")

    @property
    @pulumi.getter(name="proxyConfiguration")
    def proxy_configuration(self) -> pulumi.Output[Optional['outputs.TaskDefinitionProxyConfiguration']]:
        """
        The proxy configuration details for the App Mesh proxy.
        """
        return pulumi.get(self, "proxy_configuration")

    @property
    @pulumi.getter(name="requiresCompatibilities")
    def requires_compatibilities(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
        """
        return pulumi.get(self, "requires_compatibilities")

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Output[int]:
        """
        The revision of the task in a particular family.
        """
        return pulumi.get(self, "revision")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="taskRoleArn")
    def task_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
        """
        return pulumi.get(self, "task_role_arn")

    @property
    @pulumi.getter
    def volumes(self) -> pulumi.Output[Optional[Sequence['outputs.TaskDefinitionVolume']]]:
        """
        A set of volume blocks that containers in your task may use.
        """
        return pulumi.get(self, "volumes")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

