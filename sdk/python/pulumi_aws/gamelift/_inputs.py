# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'AliasRoutingStrategyArgs',
    'BuildStorageLocationArgs',
    'FleetEc2InboundPermissionArgs',
    'FleetResourceCreationLimitPolicyArgs',
    'FleetRuntimeConfigurationArgs',
    'FleetRuntimeConfigurationServerProcessArgs',
    'GameSessionQueuePlayerLatencyPolicyArgs',
]

@pulumi.input_type
class AliasRoutingStrategyArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 fleet_id: Optional[pulumi.Input[str]] = None,
                 message: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "type", type)
        if fleet_id is not None:
            pulumi.set(__self__, "fleet_id", fleet_id)
        if message is not None:
            pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="fleetId")
    def fleet_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "fleet_id")

    @fleet_id.setter
    def fleet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "fleet_id", value)

    @property
    @pulumi.getter
    def message(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message", value)


@pulumi.input_type
class BuildStorageLocationArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 key: pulumi.Input[str],
                 role_arn: pulumi.Input[str]):
        pulumi.set(__self__, "bucket", bucket)
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "role_arn", role_arn)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_arn", value)


@pulumi.input_type
class FleetEc2InboundPermissionArgs:
    def __init__(__self__, *,
                 from_port: pulumi.Input[float],
                 ip_range: pulumi.Input[str],
                 protocol: pulumi.Input[str],
                 to_port: pulumi.Input[float]):
        pulumi.set(__self__, "from_port", from_port)
        pulumi.set(__self__, "ip_range", ip_range)
        pulumi.set(__self__, "protocol", protocol)
        pulumi.set(__self__, "to_port", to_port)

    @property
    @pulumi.getter(name="fromPort")
    def from_port(self) -> pulumi.Input[float]:
        return pulumi.get(self, "from_port")

    @from_port.setter
    def from_port(self, value: pulumi.Input[float]):
        pulumi.set(self, "from_port", value)

    @property
    @pulumi.getter(name="ipRange")
    def ip_range(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ip_range")

    @ip_range.setter
    def ip_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_range", value)

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Input[str]:
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="toPort")
    def to_port(self) -> pulumi.Input[float]:
        return pulumi.get(self, "to_port")

    @to_port.setter
    def to_port(self, value: pulumi.Input[float]):
        pulumi.set(self, "to_port", value)


@pulumi.input_type
class FleetResourceCreationLimitPolicyArgs:
    def __init__(__self__, *,
                 new_game_sessions_per_creator: Optional[pulumi.Input[float]] = None,
                 policy_period_in_minutes: Optional[pulumi.Input[float]] = None):
        if new_game_sessions_per_creator is not None:
            pulumi.set(__self__, "new_game_sessions_per_creator", new_game_sessions_per_creator)
        if policy_period_in_minutes is not None:
            pulumi.set(__self__, "policy_period_in_minutes", policy_period_in_minutes)

    @property
    @pulumi.getter(name="newGameSessionsPerCreator")
    def new_game_sessions_per_creator(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "new_game_sessions_per_creator")

    @new_game_sessions_per_creator.setter
    def new_game_sessions_per_creator(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "new_game_sessions_per_creator", value)

    @property
    @pulumi.getter(name="policyPeriodInMinutes")
    def policy_period_in_minutes(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "policy_period_in_minutes")

    @policy_period_in_minutes.setter
    def policy_period_in_minutes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "policy_period_in_minutes", value)


@pulumi.input_type
class FleetRuntimeConfigurationArgs:
    def __init__(__self__, *,
                 game_session_activation_timeout_seconds: Optional[pulumi.Input[float]] = None,
                 max_concurrent_game_session_activations: Optional[pulumi.Input[float]] = None,
                 server_processes: Optional[pulumi.Input[List[pulumi.Input['FleetRuntimeConfigurationServerProcessArgs']]]] = None):
        if game_session_activation_timeout_seconds is not None:
            pulumi.set(__self__, "game_session_activation_timeout_seconds", game_session_activation_timeout_seconds)
        if max_concurrent_game_session_activations is not None:
            pulumi.set(__self__, "max_concurrent_game_session_activations", max_concurrent_game_session_activations)
        if server_processes is not None:
            pulumi.set(__self__, "server_processes", server_processes)

    @property
    @pulumi.getter(name="gameSessionActivationTimeoutSeconds")
    def game_session_activation_timeout_seconds(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "game_session_activation_timeout_seconds")

    @game_session_activation_timeout_seconds.setter
    def game_session_activation_timeout_seconds(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "game_session_activation_timeout_seconds", value)

    @property
    @pulumi.getter(name="maxConcurrentGameSessionActivations")
    def max_concurrent_game_session_activations(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "max_concurrent_game_session_activations")

    @max_concurrent_game_session_activations.setter
    def max_concurrent_game_session_activations(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_concurrent_game_session_activations", value)

    @property
    @pulumi.getter(name="serverProcesses")
    def server_processes(self) -> Optional[pulumi.Input[List[pulumi.Input['FleetRuntimeConfigurationServerProcessArgs']]]]:
        return pulumi.get(self, "server_processes")

    @server_processes.setter
    def server_processes(self, value: Optional[pulumi.Input[List[pulumi.Input['FleetRuntimeConfigurationServerProcessArgs']]]]):
        pulumi.set(self, "server_processes", value)


@pulumi.input_type
class FleetRuntimeConfigurationServerProcessArgs:
    def __init__(__self__, *,
                 concurrent_executions: pulumi.Input[float],
                 launch_path: pulumi.Input[str],
                 parameters: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "concurrent_executions", concurrent_executions)
        pulumi.set(__self__, "launch_path", launch_path)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter(name="concurrentExecutions")
    def concurrent_executions(self) -> pulumi.Input[float]:
        return pulumi.get(self, "concurrent_executions")

    @concurrent_executions.setter
    def concurrent_executions(self, value: pulumi.Input[float]):
        pulumi.set(self, "concurrent_executions", value)

    @property
    @pulumi.getter(name="launchPath")
    def launch_path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "launch_path")

    @launch_path.setter
    def launch_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "launch_path", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parameters", value)


@pulumi.input_type
class GameSessionQueuePlayerLatencyPolicyArgs:
    def __init__(__self__, *,
                 maximum_individual_player_latency_milliseconds: pulumi.Input[float],
                 policy_duration_seconds: Optional[pulumi.Input[float]] = None):
        pulumi.set(__self__, "maximum_individual_player_latency_milliseconds", maximum_individual_player_latency_milliseconds)
        if policy_duration_seconds is not None:
            pulumi.set(__self__, "policy_duration_seconds", policy_duration_seconds)

    @property
    @pulumi.getter(name="maximumIndividualPlayerLatencyMilliseconds")
    def maximum_individual_player_latency_milliseconds(self) -> pulumi.Input[float]:
        return pulumi.get(self, "maximum_individual_player_latency_milliseconds")

    @maximum_individual_player_latency_milliseconds.setter
    def maximum_individual_player_latency_milliseconds(self, value: pulumi.Input[float]):
        pulumi.set(self, "maximum_individual_player_latency_milliseconds", value)

    @property
    @pulumi.getter(name="policyDurationSeconds")
    def policy_duration_seconds(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "policy_duration_seconds")

    @policy_duration_seconds.setter
    def policy_duration_seconds(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "policy_duration_seconds", value)


