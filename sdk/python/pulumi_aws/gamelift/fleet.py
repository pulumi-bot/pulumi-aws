# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Fleet(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Fleet ARN.
    """
    build_id: pulumi.Output[str]
    """
    ID of the Gamelift Build to be deployed on the fleet.
    """
    description: pulumi.Output[str]
    """
    Human-readable description of the fleet.
    """
    ec2_inbound_permissions: pulumi.Output[list]
    """
    Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
    """
    ec2_instance_type: pulumi.Output[str]
    """
    Name of an EC2 instance type. e.g. `t2.micro`
    """
    log_paths: pulumi.Output[list]
    metric_groups: pulumi.Output[list]
    """
    List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
    """
    name: pulumi.Output[str]
    """
    The name of the fleet.
    """
    new_game_session_protection_policy: pulumi.Output[str]
    """
    Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
    """
    operating_system: pulumi.Output[str]
    """
    Operating system of the fleet's computing resources.
    """
    resource_creation_limit_policy: pulumi.Output[dict]
    """
    Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
    """
    runtime_configuration: pulumi.Output[dict]
    """
    Instructions for launching server processes on each instance in the fleet. See below.
    """
    def __init__(__self__, resource_name, opts=None, build_id=None, description=None, ec2_inbound_permissions=None, ec2_instance_type=None, metric_groups=None, name=None, new_game_session_protection_policy=None, resource_creation_limit_policy=None, runtime_configuration=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Gamelift Fleet resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] build_id: ID of the Gamelift Build to be deployed on the fleet.
        :param pulumi.Input[str] description: Human-readable description of the fleet.
        :param pulumi.Input[list] ec2_inbound_permissions: Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        :param pulumi.Input[str] ec2_instance_type: Name of an EC2 instance type. e.g. `t2.micro`
        :param pulumi.Input[list] metric_groups: List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        :param pulumi.Input[str] name: The name of the fleet.
        :param pulumi.Input[str] new_game_session_protection_policy: Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        :param pulumi.Input[dict] resource_creation_limit_policy: Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        :param pulumi.Input[dict] runtime_configuration: Instructions for launching server processes on each instance in the fleet. See below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_fleet.html.markdown.
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
            opts.version = utilities.get_version()
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
            __props__['metric_groups'] = metric_groups
            __props__['name'] = name
            __props__['new_game_session_protection_policy'] = new_game_session_protection_policy
            __props__['resource_creation_limit_policy'] = resource_creation_limit_policy
            __props__['runtime_configuration'] = runtime_configuration
            __props__['arn'] = None
            __props__['log_paths'] = None
            __props__['operating_system'] = None
        super(Fleet, __self__).__init__(
            'aws:gamelift/fleet:Fleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, build_id=None, description=None, ec2_inbound_permissions=None, ec2_instance_type=None, log_paths=None, metric_groups=None, name=None, new_game_session_protection_policy=None, operating_system=None, resource_creation_limit_policy=None, runtime_configuration=None):
        """
        Get an existing Fleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Fleet ARN.
        :param pulumi.Input[str] build_id: ID of the Gamelift Build to be deployed on the fleet.
        :param pulumi.Input[str] description: Human-readable description of the fleet.
        :param pulumi.Input[list] ec2_inbound_permissions: Range of IP addresses and port settings that permit inbound traffic to access server processes running on the fleet. See below.
        :param pulumi.Input[str] ec2_instance_type: Name of an EC2 instance type. e.g. `t2.micro`
        :param pulumi.Input[list] metric_groups: List of names of metric groups to add this fleet to. A metric group tracks metrics across all fleets in the group. Defaults to `default`.
        :param pulumi.Input[str] name: The name of the fleet.
        :param pulumi.Input[str] new_game_session_protection_policy: Game session protection policy to apply to all instances in this fleet. e.g. `FullProtection`. Defaults to `NoProtection`.
        :param pulumi.Input[str] operating_system: Operating system of the fleet's computing resources.
        :param pulumi.Input[dict] resource_creation_limit_policy: Policy that limits the number of game sessions an individual player can create over a span of time for this fleet. See below.
        :param pulumi.Input[dict] runtime_configuration: Instructions for launching server processes on each instance in the fleet. See below.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/gamelift_fleet.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["build_id"] = build_id
        __props__["description"] = description
        __props__["ec2_inbound_permissions"] = ec2_inbound_permissions
        __props__["ec2_instance_type"] = ec2_instance_type
        __props__["log_paths"] = log_paths
        __props__["metric_groups"] = metric_groups
        __props__["name"] = name
        __props__["new_game_session_protection_policy"] = new_game_session_protection_policy
        __props__["operating_system"] = operating_system
        __props__["resource_creation_limit_policy"] = resource_creation_limit_policy
        __props__["runtime_configuration"] = runtime_configuration
        return Fleet(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

