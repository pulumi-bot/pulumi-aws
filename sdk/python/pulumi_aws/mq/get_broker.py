# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetBrokerResult',
    'AwaitableGetBrokerResult',
    'get_broker',
]

@pulumi.output_type
class GetBrokerResult:
    """
    A collection of values returned by getBroker.
    """
    def __init__(__self__, arn=None, auto_minor_version_upgrade=None, broker_id=None, broker_name=None, configuration=None, deployment_mode=None, encryption_options=None, engine_type=None, engine_version=None, host_instance_type=None, id=None, instances=None, logs=None, maintenance_window_start_time=None, publicly_accessible=None, security_groups=None, subnet_ids=None, tags=None, users=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if auto_minor_version_upgrade and not isinstance(auto_minor_version_upgrade, bool):
            raise TypeError("Expected argument 'auto_minor_version_upgrade' to be a bool")
        pulumi.set(__self__, "auto_minor_version_upgrade", auto_minor_version_upgrade)
        if broker_id and not isinstance(broker_id, str):
            raise TypeError("Expected argument 'broker_id' to be a str")
        pulumi.set(__self__, "broker_id", broker_id)
        if broker_name and not isinstance(broker_name, str):
            raise TypeError("Expected argument 'broker_name' to be a str")
        pulumi.set(__self__, "broker_name", broker_name)
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        pulumi.set(__self__, "configuration", configuration)
        if deployment_mode and not isinstance(deployment_mode, str):
            raise TypeError("Expected argument 'deployment_mode' to be a str")
        pulumi.set(__self__, "deployment_mode", deployment_mode)
        if encryption_options and not isinstance(encryption_options, list):
            raise TypeError("Expected argument 'encryption_options' to be a list")
        pulumi.set(__self__, "encryption_options", encryption_options)
        if engine_type and not isinstance(engine_type, str):
            raise TypeError("Expected argument 'engine_type' to be a str")
        pulumi.set(__self__, "engine_type", engine_type)
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        pulumi.set(__self__, "engine_version", engine_version)
        if host_instance_type and not isinstance(host_instance_type, str):
            raise TypeError("Expected argument 'host_instance_type' to be a str")
        pulumi.set(__self__, "host_instance_type", host_instance_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        pulumi.set(__self__, "instances", instances)
        if logs and not isinstance(logs, dict):
            raise TypeError("Expected argument 'logs' to be a dict")
        pulumi.set(__self__, "logs", logs)
        if maintenance_window_start_time and not isinstance(maintenance_window_start_time, dict):
            raise TypeError("Expected argument 'maintenance_window_start_time' to be a dict")
        pulumi.set(__self__, "maintenance_window_start_time", maintenance_window_start_time)
        if publicly_accessible and not isinstance(publicly_accessible, bool):
            raise TypeError("Expected argument 'publicly_accessible' to be a bool")
        pulumi.set(__self__, "publicly_accessible", publicly_accessible)
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        pulumi.set(__self__, "security_groups", security_groups)
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        pulumi.set(__self__, "users", users)

    @property
    @pulumi.getter
    def arn(self) -> str:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> bool:
        return pulumi.get(self, "auto_minor_version_upgrade")

    @property
    @pulumi.getter(name="brokerId")
    def broker_id(self) -> str:
        return pulumi.get(self, "broker_id")

    @property
    @pulumi.getter(name="brokerName")
    def broker_name(self) -> str:
        return pulumi.get(self, "broker_name")

    @property
    @pulumi.getter
    def configuration(self) -> 'outputs.GetBrokerConfigurationResult':
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="deploymentMode")
    def deployment_mode(self) -> str:
        return pulumi.get(self, "deployment_mode")

    @property
    @pulumi.getter(name="encryptionOptions")
    def encryption_options(self) -> Sequence['outputs.GetBrokerEncryptionOptionResult']:
        return pulumi.get(self, "encryption_options")

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> str:
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> str:
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="hostInstanceType")
    def host_instance_type(self) -> str:
        return pulumi.get(self, "host_instance_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def instances(self) -> Sequence['outputs.GetBrokerInstanceResult']:
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def logs(self) -> Optional['outputs.GetBrokerLogsResult']:
        return pulumi.get(self, "logs")

    @property
    @pulumi.getter(name="maintenanceWindowStartTime")
    def maintenance_window_start_time(self) -> 'outputs.GetBrokerMaintenanceWindowStartTimeResult':
        return pulumi.get(self, "maintenance_window_start_time")

    @property
    @pulumi.getter(name="publiclyAccessible")
    def publicly_accessible(self) -> bool:
        return pulumi.get(self, "publicly_accessible")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> Sequence[str]:
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Sequence[str]:
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def users(self) -> Sequence['outputs.GetBrokerUserResult']:
        return pulumi.get(self, "users")


class AwaitableGetBrokerResult(GetBrokerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBrokerResult(
            arn=self.arn,
            auto_minor_version_upgrade=self.auto_minor_version_upgrade,
            broker_id=self.broker_id,
            broker_name=self.broker_name,
            configuration=self.configuration,
            deployment_mode=self.deployment_mode,
            encryption_options=self.encryption_options,
            engine_type=self.engine_type,
            engine_version=self.engine_version,
            host_instance_type=self.host_instance_type,
            id=self.id,
            instances=self.instances,
            logs=self.logs,
            maintenance_window_start_time=self.maintenance_window_start_time,
            publicly_accessible=self.publicly_accessible,
            security_groups=self.security_groups,
            subnet_ids=self.subnet_ids,
            tags=self.tags,
            users=self.users)


def get_broker(broker_id: Optional[str] = None,
               broker_name: Optional[str] = None,
               logs: Optional[pulumi.InputType['GetBrokerLogsArgs']] = None,
               tags: Optional[Mapping[str, str]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBrokerResult:
    """
    Provides information about a MQ Broker.


    :param str broker_id: The unique id of the mq broker.
    :param str broker_name: The unique name of the mq broker.
    """
    __args__ = dict()
    __args__['brokerId'] = broker_id
    __args__['brokerName'] = broker_name
    __args__['logs'] = logs
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:mq/getBroker:getBroker', __args__, opts=opts, typ=GetBrokerResult).value

    return AwaitableGetBrokerResult(
        arn=__ret__.arn,
        auto_minor_version_upgrade=__ret__.auto_minor_version_upgrade,
        broker_id=__ret__.broker_id,
        broker_name=__ret__.broker_name,
        configuration=__ret__.configuration,
        deployment_mode=__ret__.deployment_mode,
        encryption_options=__ret__.encryption_options,
        engine_type=__ret__.engine_type,
        engine_version=__ret__.engine_version,
        host_instance_type=__ret__.host_instance_type,
        id=__ret__.id,
        instances=__ret__.instances,
        logs=__ret__.logs,
        maintenance_window_start_time=__ret__.maintenance_window_start_time,
        publicly_accessible=__ret__.publicly_accessible,
        security_groups=__ret__.security_groups,
        subnet_ids=__ret__.subnet_ids,
        tags=__ret__.tags,
        users=__ret__.users)
