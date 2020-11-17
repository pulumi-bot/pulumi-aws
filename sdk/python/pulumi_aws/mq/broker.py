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

__all__ = ['Broker']


class Broker(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apply_immediately: Optional[pulumi.Input[bool]] = None,
                 auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
                 broker_name: Optional[pulumi.Input[str]] = None,
                 configuration: Optional[pulumi.Input[pulumi.InputType['BrokerConfigurationArgs']]] = None,
                 deployment_mode: Optional[pulumi.Input[str]] = None,
                 encryption_options: Optional[pulumi.Input[pulumi.InputType['BrokerEncryptionOptionsArgs']]] = None,
                 engine_type: Optional[pulumi.Input[str]] = None,
                 engine_version: Optional[pulumi.Input[str]] = None,
                 host_instance_type: Optional[pulumi.Input[str]] = None,
                 logs: Optional[pulumi.Input[pulumi.InputType['BrokerLogsArgs']]] = None,
                 maintenance_window_start_time: Optional[pulumi.Input[pulumi.InputType['BrokerMaintenanceWindowStartTimeArgs']]] = None,
                 publicly_accessible: Optional[pulumi.Input[bool]] = None,
                 security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BrokerUserArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an MQ Broker Resource. This resources also manages users for the broker.

        For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).

        Changes to an MQ Broker can occur when you change a
        parameter, such as `configuration` or `user`, and are reflected in the next maintenance
        window. Because of this, this provider may report a difference in its planning
        phase because a modification has not yet taken place. You can use the
        `apply_immediately` flag to instruct the service to apply the change immediately
        (see documentation below).

        > **Note:** using `apply_immediately` can result in a
        brief downtime as the broker reboots.

        > **Note:** All arguments including the username and password will be stored in the raw state as plain-text.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.mq.Broker("example",
            broker_name="example",
            configuration=aws.mq.BrokerConfigurationArgs(
                id=aws_mq_configuration["test"]["id"],
                revision=aws_mq_configuration["test"]["latest_revision"],
            ),
            engine_type="ActiveMQ",
            engine_version="5.15.0",
            host_instance_type="mq.t2.micro",
            security_groups=[aws_security_group["test"]["id"]],
            users=[aws.mq.BrokerUserArgs(
                username="ExampleUser",
                password="MindTheGap",
            )])
        ```

        ## Import

        MQ Brokers can be imported using their broker id, e.g.

        ```sh
         $ pulumi import aws:mq/broker:Broker example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any broker modifications
               are applied immediately, or during the next maintenance window. Default is `false`.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
        :param pulumi.Input[str] broker_name: The name of the broker.
        :param pulumi.Input[pulumi.InputType['BrokerConfigurationArgs']] configuration: Configuration of the broker. See below.
        :param pulumi.Input[str] deployment_mode: The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
        :param pulumi.Input[pulumi.InputType['BrokerEncryptionOptionsArgs']] encryption_options: Configuration block containing encryption options. See below.
        :param pulumi.Input[str] engine_type: The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
        :param pulumi.Input[str] engine_version: The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
        :param pulumi.Input[str] host_instance_type: The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
        :param pulumi.Input[pulumi.InputType['BrokerLogsArgs']] logs: Logging configuration of the broker. See below.
        :param pulumi.Input[pulumi.InputType['BrokerMaintenanceWindowStartTimeArgs']] maintenance_window_start_time: Maintenance window start time. See below.
        :param pulumi.Input[bool] publicly_accessible: Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: The list of security group IDs assigned to the broker.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BrokerUserArgs']]]] users: The list of all ActiveMQ usernames for the specified broker. See below.
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

            __props__['apply_immediately'] = apply_immediately
            __props__['auto_minor_version_upgrade'] = auto_minor_version_upgrade
            if broker_name is None:
                raise TypeError("Missing required property 'broker_name'")
            __props__['broker_name'] = broker_name
            __props__['configuration'] = configuration
            __props__['deployment_mode'] = deployment_mode
            __props__['encryption_options'] = encryption_options
            if engine_type is None:
                raise TypeError("Missing required property 'engine_type'")
            __props__['engine_type'] = engine_type
            if engine_version is None:
                raise TypeError("Missing required property 'engine_version'")
            __props__['engine_version'] = engine_version
            if host_instance_type is None:
                raise TypeError("Missing required property 'host_instance_type'")
            __props__['host_instance_type'] = host_instance_type
            __props__['logs'] = logs
            __props__['maintenance_window_start_time'] = maintenance_window_start_time
            __props__['publicly_accessible'] = publicly_accessible
            if security_groups is None:
                raise TypeError("Missing required property 'security_groups'")
            __props__['security_groups'] = security_groups
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            if users is None:
                raise TypeError("Missing required property 'users'")
            __props__['users'] = users
            __props__['arn'] = None
            __props__['instances'] = None
        super(Broker, __self__).__init__(
            'aws:mq/broker:Broker',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            apply_immediately: Optional[pulumi.Input[bool]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            auto_minor_version_upgrade: Optional[pulumi.Input[bool]] = None,
            broker_name: Optional[pulumi.Input[str]] = None,
            configuration: Optional[pulumi.Input[pulumi.InputType['BrokerConfigurationArgs']]] = None,
            deployment_mode: Optional[pulumi.Input[str]] = None,
            encryption_options: Optional[pulumi.Input[pulumi.InputType['BrokerEncryptionOptionsArgs']]] = None,
            engine_type: Optional[pulumi.Input[str]] = None,
            engine_version: Optional[pulumi.Input[str]] = None,
            host_instance_type: Optional[pulumi.Input[str]] = None,
            instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BrokerInstanceArgs']]]]] = None,
            logs: Optional[pulumi.Input[pulumi.InputType['BrokerLogsArgs']]] = None,
            maintenance_window_start_time: Optional[pulumi.Input[pulumi.InputType['BrokerMaintenanceWindowStartTimeArgs']]] = None,
            publicly_accessible: Optional[pulumi.Input[bool]] = None,
            security_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BrokerUserArgs']]]]] = None) -> 'Broker':
        """
        Get an existing Broker resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] apply_immediately: Specifies whether any broker modifications
               are applied immediately, or during the next maintenance window. Default is `false`.
        :param pulumi.Input[str] arn: The ARN of the broker.
        :param pulumi.Input[bool] auto_minor_version_upgrade: Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
        :param pulumi.Input[str] broker_name: The name of the broker.
        :param pulumi.Input[pulumi.InputType['BrokerConfigurationArgs']] configuration: Configuration of the broker. See below.
        :param pulumi.Input[str] deployment_mode: The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
        :param pulumi.Input[pulumi.InputType['BrokerEncryptionOptionsArgs']] encryption_options: Configuration block containing encryption options. See below.
        :param pulumi.Input[str] engine_type: The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
        :param pulumi.Input[str] engine_version: The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
        :param pulumi.Input[str] host_instance_type: The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BrokerInstanceArgs']]]] instances: A list of information about allocated brokers (both active & standby).
               * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
               * `instances.0.ip_address` - The IP Address of the broker.
               * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
               * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
               * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
               * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
               * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
               * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
        :param pulumi.Input[pulumi.InputType['BrokerLogsArgs']] logs: Logging configuration of the broker. See below.
        :param pulumi.Input[pulumi.InputType['BrokerMaintenanceWindowStartTimeArgs']] maintenance_window_start_time: Maintenance window start time. See below.
        :param pulumi.Input[bool] publicly_accessible: Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_groups: The list of security group IDs assigned to the broker.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BrokerUserArgs']]]] users: The list of all ActiveMQ usernames for the specified broker. See below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["apply_immediately"] = apply_immediately
        __props__["arn"] = arn
        __props__["auto_minor_version_upgrade"] = auto_minor_version_upgrade
        __props__["broker_name"] = broker_name
        __props__["configuration"] = configuration
        __props__["deployment_mode"] = deployment_mode
        __props__["encryption_options"] = encryption_options
        __props__["engine_type"] = engine_type
        __props__["engine_version"] = engine_version
        __props__["host_instance_type"] = host_instance_type
        __props__["instances"] = instances
        __props__["logs"] = logs
        __props__["maintenance_window_start_time"] = maintenance_window_start_time
        __props__["publicly_accessible"] = publicly_accessible
        __props__["security_groups"] = security_groups
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["users"] = users
        return Broker(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applyImmediately")
    def apply_immediately(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether any broker modifications
        are applied immediately, or during the next maintenance window. Default is `false`.
        """
        return pulumi.get(self, "apply_immediately")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the broker.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="autoMinorVersionUpgrade")
    def auto_minor_version_upgrade(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables automatic upgrades to new minor versions for brokers, as Apache releases the versions.
        """
        return pulumi.get(self, "auto_minor_version_upgrade")

    @property
    @pulumi.getter(name="brokerName")
    def broker_name(self) -> pulumi.Output[str]:
        """
        The name of the broker.
        """
        return pulumi.get(self, "broker_name")

    @property
    @pulumi.getter
    def configuration(self) -> pulumi.Output['outputs.BrokerConfiguration']:
        """
        Configuration of the broker. See below.
        """
        return pulumi.get(self, "configuration")

    @property
    @pulumi.getter(name="deploymentMode")
    def deployment_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The deployment mode of the broker. Supported: `SINGLE_INSTANCE` and `ACTIVE_STANDBY_MULTI_AZ`. Defaults to `SINGLE_INSTANCE`.
        """
        return pulumi.get(self, "deployment_mode")

    @property
    @pulumi.getter(name="encryptionOptions")
    def encryption_options(self) -> pulumi.Output[Optional['outputs.BrokerEncryptionOptions']]:
        """
        Configuration block containing encryption options. See below.
        """
        return pulumi.get(self, "encryption_options")

    @property
    @pulumi.getter(name="engineType")
    def engine_type(self) -> pulumi.Output[str]:
        """
        The type of broker engine. Currently, Amazon MQ supports only `ActiveMQ`.
        """
        return pulumi.get(self, "engine_type")

    @property
    @pulumi.getter(name="engineVersion")
    def engine_version(self) -> pulumi.Output[str]:
        """
        The version of the broker engine. See the [AmazonMQ Broker Engine docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) for supported versions.
        """
        return pulumi.get(self, "engine_version")

    @property
    @pulumi.getter(name="hostInstanceType")
    def host_instance_type(self) -> pulumi.Output[str]:
        """
        The broker's instance type. e.g. `mq.t2.micro` or `mq.m4.large`
        """
        return pulumi.get(self, "host_instance_type")

    @property
    @pulumi.getter
    def instances(self) -> pulumi.Output[Sequence['outputs.BrokerInstance']]:
        """
        A list of information about allocated brokers (both active & standby).
        * `instances.0.console_url` - The URL of the broker's [ActiveMQ Web Console](http://activemq.apache.org/web-console.html).
        * `instances.0.ip_address` - The IP Address of the broker.
        * `instances.0.endpoints` - The broker's wire-level protocol endpoints in the following order & format referenceable e.g. as `instances.0.endpoints.0` (SSL):
        * `ssl://broker-id.mq.us-west-2.amazonaws.com:61617`
        * `amqp+ssl://broker-id.mq.us-west-2.amazonaws.com:5671`
        * `stomp+ssl://broker-id.mq.us-west-2.amazonaws.com:61614`
        * `mqtt+ssl://broker-id.mq.us-west-2.amazonaws.com:8883`
        * `wss://broker-id.mq.us-west-2.amazonaws.com:61619`
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter
    def logs(self) -> pulumi.Output[Optional['outputs.BrokerLogs']]:
        """
        Logging configuration of the broker. See below.
        """
        return pulumi.get(self, "logs")

    @property
    @pulumi.getter(name="maintenanceWindowStartTime")
    def maintenance_window_start_time(self) -> pulumi.Output['outputs.BrokerMaintenanceWindowStartTime']:
        """
        Maintenance window start time. See below.
        """
        return pulumi.get(self, "maintenance_window_start_time")

    @property
    @pulumi.getter(name="publiclyAccessible")
    def publicly_accessible(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to enable connections from applications outside of the VPC that hosts the broker's subnets.
        """
        return pulumi.get(self, "publicly_accessible")

    @property
    @pulumi.getter(name="securityGroups")
    def security_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of security group IDs assigned to the broker.
        """
        return pulumi.get(self, "security_groups")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of subnet IDs in which to launch the broker. A `SINGLE_INSTANCE` deployment requires one subnet. An `ACTIVE_STANDBY_MULTI_AZ` deployment requires two subnets.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Sequence['outputs.BrokerUser']]:
        """
        The list of all ActiveMQ usernames for the specified broker. See below.
        """
        return pulumi.get(self, "users")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

