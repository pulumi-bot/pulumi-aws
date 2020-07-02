# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, arn=None, availability_zones=None, default_cooldown=None, desired_capacity=None, health_check_grace_period=None, health_check_type=None, id=None, launch_configuration=None, load_balancers=None, max_size=None, min_size=None, name=None, new_instances_protected_from_scale_in=None, placement_group=None, service_linked_role_arn=None, status=None, target_group_arns=None, termination_policies=None, vpc_zone_identifier=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The Amazon Resource Name (ARN) of the Auto Scaling group.
        """
        if availability_zones and not isinstance(availability_zones, list):
            raise TypeError("Expected argument 'availability_zones' to be a list")
        __self__.availability_zones = availability_zones
        """
        One or more Availability Zones for the group.
        """
        if default_cooldown and not isinstance(default_cooldown, float):
            raise TypeError("Expected argument 'default_cooldown' to be a float")
        __self__.default_cooldown = default_cooldown
        if desired_capacity and not isinstance(desired_capacity, float):
            raise TypeError("Expected argument 'desired_capacity' to be a float")
        __self__.desired_capacity = desired_capacity
        """
        The desired size of the group.
        """
        if health_check_grace_period and not isinstance(health_check_grace_period, float):
            raise TypeError("Expected argument 'health_check_grace_period' to be a float")
        __self__.health_check_grace_period = health_check_grace_period
        """
        The amount of time, in seconds, that Amazon EC2 Auto Scaling waits before checking the health status of an EC2 instance that has come into service.
        """
        if health_check_type and not isinstance(health_check_type, str):
            raise TypeError("Expected argument 'health_check_type' to be a str")
        __self__.health_check_type = health_check_type
        """
        The service to use for the health checks. The valid values are EC2 and ELB.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if launch_configuration and not isinstance(launch_configuration, str):
            raise TypeError("Expected argument 'launch_configuration' to be a str")
        __self__.launch_configuration = launch_configuration
        """
        The name of the associated launch configuration.
        """
        if load_balancers and not isinstance(load_balancers, list):
            raise TypeError("Expected argument 'load_balancers' to be a list")
        __self__.load_balancers = load_balancers
        """
        One or more load balancers associated with the group.
        """
        if max_size and not isinstance(max_size, float):
            raise TypeError("Expected argument 'max_size' to be a float")
        __self__.max_size = max_size
        """
        The maximum size of the group.
        """
        if min_size and not isinstance(min_size, float):
            raise TypeError("Expected argument 'min_size' to be a float")
        __self__.min_size = min_size
        """
        The minimum size of the group.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the Auto Scaling group.
        """
        if new_instances_protected_from_scale_in and not isinstance(new_instances_protected_from_scale_in, bool):
            raise TypeError("Expected argument 'new_instances_protected_from_scale_in' to be a bool")
        __self__.new_instances_protected_from_scale_in = new_instances_protected_from_scale_in
        if placement_group and not isinstance(placement_group, str):
            raise TypeError("Expected argument 'placement_group' to be a str")
        __self__.placement_group = placement_group
        """
        The name of the placement group into which to launch your instances, if any. For more information, see Placement Groups (http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html) in the Amazon Elastic Compute Cloud User Guide.
        """
        if service_linked_role_arn and not isinstance(service_linked_role_arn, str):
            raise TypeError("Expected argument 'service_linked_role_arn' to be a str")
        __self__.service_linked_role_arn = service_linked_role_arn
        """
        The Amazon Resource Name (ARN) of the service-linked role that the Auto Scaling group uses to call other AWS services on your behalf.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        The current state of the group when DeleteAutoScalingGroup is in progress.
        """
        if target_group_arns and not isinstance(target_group_arns, list):
            raise TypeError("Expected argument 'target_group_arns' to be a list")
        __self__.target_group_arns = target_group_arns
        """
        The Amazon Resource Names (ARN) of the target groups for your load balancer.
        """
        if termination_policies and not isinstance(termination_policies, list):
            raise TypeError("Expected argument 'termination_policies' to be a list")
        __self__.termination_policies = termination_policies
        """
        The termination policies for the group.
        """
        if vpc_zone_identifier and not isinstance(vpc_zone_identifier, str):
            raise TypeError("Expected argument 'vpc_zone_identifier' to be a str")
        __self__.vpc_zone_identifier = vpc_zone_identifier
        """
        VPC ID for the group.
        """


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            arn=self.arn,
            availability_zones=self.availability_zones,
            default_cooldown=self.default_cooldown,
            desired_capacity=self.desired_capacity,
            health_check_grace_period=self.health_check_grace_period,
            health_check_type=self.health_check_type,
            id=self.id,
            launch_configuration=self.launch_configuration,
            load_balancers=self.load_balancers,
            max_size=self.max_size,
            min_size=self.min_size,
            name=self.name,
            new_instances_protected_from_scale_in=self.new_instances_protected_from_scale_in,
            placement_group=self.placement_group,
            service_linked_role_arn=self.service_linked_role_arn,
            status=self.status,
            target_group_arns=self.target_group_arns,
            termination_policies=self.termination_policies,
            vpc_zone_identifier=self.vpc_zone_identifier)


def get_group(name=None,opts=None):
    """
    Use this data source to get information on an existing autoscaling group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    foo = aws.autoscaling.get_group(name="foo")
    ```


    :param str name: Specify the exact name of the desired autoscaling group.
    """
    __args__ = dict()

    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:autoscaling/getGroup:getGroup', __args__, opts=opts).value

    return AwaitableGetGroupResult(
        arn=__ret__.get('arn'),
        availability_zones=__ret__.get('availabilityZones'),
        default_cooldown=__ret__.get('defaultCooldown'),
        desired_capacity=__ret__.get('desiredCapacity'),
        health_check_grace_period=__ret__.get('healthCheckGracePeriod'),
        health_check_type=__ret__.get('healthCheckType'),
        id=__ret__.get('id'),
        launch_configuration=__ret__.get('launchConfiguration'),
        load_balancers=__ret__.get('loadBalancers'),
        max_size=__ret__.get('maxSize'),
        min_size=__ret__.get('minSize'),
        name=__ret__.get('name'),
        new_instances_protected_from_scale_in=__ret__.get('newInstancesProtectedFromScaleIn'),
        placement_group=__ret__.get('placementGroup'),
        service_linked_role_arn=__ret__.get('serviceLinkedRoleArn'),
        status=__ret__.get('status'),
        target_group_arns=__ret__.get('targetGroupArns'),
        termination_policies=__ret__.get('terminationPolicies'),
        vpc_zone_identifier=__ret__.get('vpcZoneIdentifier'))
