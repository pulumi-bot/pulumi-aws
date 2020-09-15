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

__all__ = ['Fleet']


class Fleet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ec2_inbound_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FleetEc2InboundPermissionArgs']]]]] = None,
                 ec2_instance_type: Optional[pulumi.Input[str]] = None,
                 fleet_type: Optional[pulumi.Input[str]] = None,
                 instance_role_arn: Optional[pulumi.Input[str]] = None,
                 metric_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 new_game_session_protection_policy: Optional[pulumi.Input[str]] = None,
                 resource_creation_limit_policy: Optional[pulumi.Input[pulumi.InputType['FleetResourceCreationLimitPolicyArgs']]] = None,
                 runtime_configuration: Optional[pulumi.Input[pulumi.InputType['FleetRuntimeConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Gamelift Fleet resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.gamelift.Fleet("example",
            build_id=aws_gamelift_build["example"]["id"],
            ec2_instance_type="t2.micro",
            fleet_type="ON_DEMAND",
            runtime_configuration=aws.gamelift.FleetRuntimeConfigurationArgs(
                server_processes=[aws.gamelift.FleetRuntimeConfigurationServerProcessArgs(
                    concurrent_executions=1,
                    launch_path="C:\\game\\GomokuServer.exe",
                )],
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] build_id: ID of the Gamelift Build to be deployed on the fleet.
        :param pulumi.Input[str] description: Human-readable description of the fleet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FleetEc2InboundPermissionArgs']]]] ec2_inbound_permissions: Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        :param pulumi.Input[str] ec2_instance_type: Name of an EC2 instance type. e.g. `t2.micro`
        :param pulumi.Input[str] fleet_type: Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
        :param pulumi.Input[str] instance_role_arn: ARN of an IAM role that instances in the fleet can assume.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] metric_groups: List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        :param pulumi.Input[str] name: The name of the fleet.
        :param pulumi.Input[str] new_game_session_protection_policy: Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        :param pulumi.Input[pulumi.InputType['FleetResourceCreationLimitPolicyArgs']] resource_creation_limit_policy: Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        :param pulumi.Input[pulumi.InputType['FleetRuntimeConfigurationArgs']] runtime_configuration: Instructions for launching server processes on each instance in the fleet. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
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

            if build_id is None:
                raise TypeError("Missing required property 'build_id'")
            __props__['build_id'] = build_id
            __props__['description'] = description
            __props__['ec2_inbound_permissions'] = ec2_inbound_permissions
            if ec2_instance_type is None:
                raise TypeError("Missing required property 'ec2_instance_type'")
            __props__['ec2_instance_type'] = ec2_instance_type
            __props__['fleet_type'] = fleet_type
            __props__['instance_role_arn'] = instance_role_arn
            __props__['metric_groups'] = metric_groups
            __props__['name'] = name
            __props__['new_game_session_protection_policy'] = new_game_session_protection_policy
            __props__['resource_creation_limit_policy'] = resource_creation_limit_policy
            __props__['runtime_configuration'] = runtime_configuration
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['log_paths'] = None
            __props__['operating_system'] = None
        super(Fleet, __self__).__init__(
            'aws:gamelift/fleet:Fleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            build_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ec2_inbound_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FleetEc2InboundPermissionArgs']]]]] = None,
            ec2_instance_type: Optional[pulumi.Input[str]] = None,
            fleet_type: Optional[pulumi.Input[str]] = None,
            instance_role_arn: Optional[pulumi.Input[str]] = None,
            log_paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            metric_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            new_game_session_protection_policy: Optional[pulumi.Input[str]] = None,
            operating_system: Optional[pulumi.Input[str]] = None,
            resource_creation_limit_policy: Optional[pulumi.Input[pulumi.InputType['FleetResourceCreationLimitPolicyArgs']]] = None,
            runtime_configuration: Optional[pulumi.Input[pulumi.InputType['FleetRuntimeConfigurationArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Fleet':
        """
        Get an existing Fleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Fleet ARN.
        :param pulumi.Input[str] build_id: ID of the Gamelift Build to be deployed on the fleet.
        :param pulumi.Input[str] description: Human-readable description of the fleet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FleetEc2InboundPermissionArgs']]]] ec2_inbound_permissions: Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        :param pulumi.Input[str] ec2_instance_type: Name of an EC2 instance type. e.g. `t2.micro`
        :param pulumi.Input[str] fleet_type: Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
        :param pulumi.Input[str] instance_role_arn: ARN of an IAM role that instances in the fleet can assume.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] metric_groups: List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        :param pulumi.Input[str] name: The name of the fleet.
        :param pulumi.Input[str] new_game_session_protection_policy: Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        :param pulumi.Input[str] operating_system: Operating system of the fleet's computing resources.
        :param pulumi.Input[pulumi.InputType['FleetResourceCreationLimitPolicyArgs']] resource_creation_limit_policy: Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        :param pulumi.Input[pulumi.InputType['FleetRuntimeConfigurationArgs']] runtime_configuration: Instructions for launching server processes on each instance in the fleet. See below.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["build_id"] = build_id
        __props__["description"] = description
        __props__["ec2_inbound_permissions"] = ec2_inbound_permissions
        __props__["ec2_instance_type"] = ec2_instance_type
        __props__["fleet_type"] = fleet_type
        __props__["instance_role_arn"] = instance_role_arn
        __props__["log_paths"] = log_paths
        __props__["metric_groups"] = metric_groups
        __props__["name"] = name
        __props__["new_game_session_protection_policy"] = new_game_session_protection_policy
        __props__["operating_system"] = operating_system
        __props__["resource_creation_limit_policy"] = resource_creation_limit_policy
        __props__["runtime_configuration"] = runtime_configuration
        __props__["tags"] = tags
        return Fleet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Fleet ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="buildId")
    def build_id(self) -> pulumi.Output[str]:
        """
        ID of the Gamelift Build to be deployed on the fleet.
        """
        return pulumi.get(self, "build_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human-readable description of the fleet.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ec2InboundPermissions")
    def ec2_inbound_permissions(self) -> pulumi.Output[Optional[Sequence['outputs.FleetEc2InboundPermission']]]:
        """
        Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        """
        return pulumi.get(self, "ec2_inbound_permissions")

    @property
    @pulumi.getter(name="ec2InstanceType")
    def ec2_instance_type(self) -> pulumi.Output[str]:
        """
        Name of an EC2 instance type. e.g. `t2.micro`
        """
        return pulumi.get(self, "ec2_instance_type")

    @property
    @pulumi.getter(name="fleetType")
    def fleet_type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of fleet. This value must be `ON_DEMAND` or `SPOT`. Defaults to `ON_DEMAND`.
        """
        return pulumi.get(self, "fleet_type")

    @property
    @pulumi.getter(name="instanceRoleArn")
    def instance_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        ARN of an IAM role that instances in the fleet can assume.
        """
        return pulumi.get(self, "instance_role_arn")

    @property
    @pulumi.getter(name="logPaths")
    def log_paths(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "log_paths")

    @property
    @pulumi.getter(name="metricGroups")
    def metric_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        """
        return pulumi.get(self, "metric_groups")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the fleet.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="newGameSessionProtectionPolicy")
    def new_game_session_protection_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        """
        return pulumi.get(self, "new_game_session_protection_policy")

    @property
    @pulumi.getter(name="operatingSystem")
    def operating_system(self) -> pulumi.Output[str]:
        """
        Operating system of the fleet's computing resources.
        """
        return pulumi.get(self, "operating_system")

    @property
    @pulumi.getter(name="resourceCreationLimitPolicy")
    def resource_creation_limit_policy(self) -> pulumi.Output[Optional['outputs.FleetResourceCreationLimitPolicy']]:
        """
        Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        """
        return pulumi.get(self, "resource_creation_limit_policy")

    @property
    @pulumi.getter(name="runtimeConfiguration")
    def runtime_configuration(self) -> pulumi.Output[Optional['outputs.FleetRuntimeConfiguration']]:
        """
        Instructions for launching server processes on each instance in the fleet. See below.
        """
        return pulumi.get(self, "runtime_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

