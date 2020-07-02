# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetBrokerResult:
    """
    A collection of values returned by getBroker.
    """
    def __init__(__self__, arn=None, auto_minor_version_upgrade=None, broker_id=None, broker_name=None, configuration=None, deployment_mode=None, encryption_options=None, engine_type=None, engine_version=None, host_instance_type=None, id=None, instances=None, logs=None, maintenance_window_start_time=None, publicly_accessible=None, security_groups=None, subnet_ids=None, tags=None, users=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        if auto_minor_version_upgrade and not isinstance(auto_minor_version_upgrade, bool):
            raise TypeError("Expected argument 'auto_minor_version_upgrade' to be a bool")
        __self__.auto_minor_version_upgrade = auto_minor_version_upgrade
        if broker_id and not isinstance(broker_id, str):
            raise TypeError("Expected argument 'broker_id' to be a str")
        __self__.broker_id = broker_id
        if broker_name and not isinstance(broker_name, str):
            raise TypeError("Expected argument 'broker_name' to be a str")
        __self__.broker_name = broker_name
        if configuration and not isinstance(configuration, dict):
            raise TypeError("Expected argument 'configuration' to be a dict")
        __self__.configuration = configuration
        if deployment_mode and not isinstance(deployment_mode, str):
            raise TypeError("Expected argument 'deployment_mode' to be a str")
        __self__.deployment_mode = deployment_mode
        if encryption_options and not isinstance(encryption_options, list):
            raise TypeError("Expected argument 'encryption_options' to be a list")
        __self__.encryption_options = encryption_options
        if engine_type and not isinstance(engine_type, str):
            raise TypeError("Expected argument 'engine_type' to be a str")
        __self__.engine_type = engine_type
        if engine_version and not isinstance(engine_version, str):
            raise TypeError("Expected argument 'engine_version' to be a str")
        __self__.engine_version = engine_version
        if host_instance_type and not isinstance(host_instance_type, str):
            raise TypeError("Expected argument 'host_instance_type' to be a str")
        __self__.host_instance_type = host_instance_type
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if instances and not isinstance(instances, list):
            raise TypeError("Expected argument 'instances' to be a list")
        __self__.instances = instances
        if logs and not isinstance(logs, dict):
            raise TypeError("Expected argument 'logs' to be a dict")
        __self__.logs = logs
        if maintenance_window_start_time and not isinstance(maintenance_window_start_time, dict):
            raise TypeError("Expected argument 'maintenance_window_start_time' to be a dict")
        __self__.maintenance_window_start_time = maintenance_window_start_time
        if publicly_accessible and not isinstance(publicly_accessible, bool):
            raise TypeError("Expected argument 'publicly_accessible' to be a bool")
        __self__.publicly_accessible = publicly_accessible
        if security_groups and not isinstance(security_groups, list):
            raise TypeError("Expected argument 'security_groups' to be a list")
        __self__.security_groups = security_groups
        if subnet_ids and not isinstance(subnet_ids, list):
            raise TypeError("Expected argument 'subnet_ids' to be a list")
        __self__.subnet_ids = subnet_ids
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if users and not isinstance(users, list):
            raise TypeError("Expected argument 'users' to be a list")
        __self__.users = users


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


def get_broker(broker_id=None,broker_name=None,logs=None,tags=None,opts=None):
    """
    Provides information about a MQ Broker.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    broker_id = config.get("brokerId")
    if broker_id is None:
        broker_id = ""
    broker_name = config.get("brokerName")
    if broker_name is None:
        broker_name = ""
    by_id = aws.mq.get_broker(broker_id=broker_id)
    by_name = aws.mq.get_broker(broker_name=broker_name)
    ```


    :param str broker_id: The unique id of the mq broker.
    :param str broker_name: The unique name of the mq broker.

    The **logs** object supports the following:

      * `audit` (`bool`)
      * `general` (`bool`)
    """
    __args__ = dict()

    __args__['brokerId'] = broker_id
    __args__['brokerName'] = broker_name
    __args__['logs'] = logs
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:mq/getBroker:getBroker', __args__, opts=opts).value

    return AwaitableGetBrokerResult(
        arn=__ret__.get('arn'),
        auto_minor_version_upgrade=__ret__.get('autoMinorVersionUpgrade'),
        broker_id=__ret__.get('brokerId'),
        broker_name=__ret__.get('brokerName'),
        configuration=__ret__.get('configuration'),
        deployment_mode=__ret__.get('deploymentMode'),
        encryption_options=__ret__.get('encryptionOptions'),
        engine_type=__ret__.get('engineType'),
        engine_version=__ret__.get('engineVersion'),
        host_instance_type=__ret__.get('hostInstanceType'),
        id=__ret__.get('id'),
        instances=__ret__.get('instances'),
        logs=__ret__.get('logs'),
        maintenance_window_start_time=__ret__.get('maintenanceWindowStartTime'),
        publicly_accessible=__ret__.get('publiclyAccessible'),
        security_groups=__ret__.get('securityGroups'),
        subnet_ids=__ret__.get('subnetIds'),
        tags=__ret__.get('tags'),
        users=__ret__.get('users'))
