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

__all__ = ['SpotFleetRequest']


class SpotFleetRequest(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allocation_strategy: Optional[pulumi.Input[str]] = None,
                 excess_capacity_termination_policy: Optional[pulumi.Input[str]] = None,
                 fleet_type: Optional[pulumi.Input[str]] = None,
                 iam_fleet_role: Optional[pulumi.Input[str]] = None,
                 instance_interruption_behaviour: Optional[pulumi.Input[str]] = None,
                 instance_pools_to_use_count: Optional[pulumi.Input[int]] = None,
                 launch_specifications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchSpecificationArgs']]]]] = None,
                 launch_template_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchTemplateConfigArgs']]]]] = None,
                 load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 replace_unhealthy_instances: Optional[pulumi.Input[bool]] = None,
                 spot_price: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_capacity: Optional[pulumi.Input[int]] = None,
                 target_group_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 terminate_instances_with_expiration: Optional[pulumi.Input[bool]] = None,
                 valid_from: Optional[pulumi.Input[str]] = None,
                 valid_until: Optional[pulumi.Input[str]] = None,
                 wait_for_fulfillment: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an EC2 Spot Fleet Request resource. This allows a fleet of Spot
        instances to be requested on the Spot market.

        ## Example Usage
        ### Using launch specifications

        ```python
        import pulumi
        import pulumi_aws as aws

        # Request a Spot fleet
        cheap_compute = aws.ec2.SpotFleetRequest("cheapCompute",
            iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
            spot_price="0.03",
            allocation_strategy="diversified",
            target_capacity=6,
            valid_until="2019-11-04T20:44:20Z",
            launch_specifications=[
                aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
                    instance_type="m4.10xlarge",
                    ami="ami-1234",
                    spot_price="2.793",
                    placement_tenancy="dedicated",
                    iam_instance_profile_arn=aws_iam_instance_profile["example"]["arn"],
                ),
                aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
                    instance_type="m4.4xlarge",
                    ami="ami-5678",
                    key_name="my-key",
                    spot_price="1.117",
                    iam_instance_profile_arn=aws_iam_instance_profile["example"]["arn"],
                    availability_zone="us-west-1a",
                    subnet_id="subnet-1234",
                    weighted_capacity="35",
                    root_block_devices=[{
                        "volume_size": "300",
                        "volumeType": "gp2",
                    }],
                    tags={
                        "Name": "spot-fleet-example",
                    },
                ),
            ])
        ```
        ### Using launch templates

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_launch_template = aws.ec2.LaunchTemplate("fooLaunchTemplate",
            image_id="ami-516b9131",
            instance_type="m1.small",
            key_name="some-key")
        foo_spot_fleet_request = aws.ec2.SpotFleetRequest("fooSpotFleetRequest",
            iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
            spot_price="0.005",
            target_capacity=2,
            valid_until="2019-11-04T20:44:20Z",
            launch_template_configs=[aws.ec2.SpotFleetRequestLaunchTemplateConfigArgs(
                launch_template_specification=aws.ec2.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs(
                    id=foo_launch_template.id,
                    version=foo_launch_template.latest_version,
                ),
            )],
            opts=ResourceOptions(depends_on=[aws_iam_policy_attachment["test-attach"]]))
        ```

        > **NOTE:** This provider does not support the functionality where multiple `subnet_id` or `availability_zone` parameters can be specified in the same
        launch configuration block. If you want to specify multiple values, then separate launch configuration blocks should be used:
        ### Using multiple launch specifications

        ```python
        import pulumi
        import pulumi_aws as aws

        foo = aws.ec2.SpotFleetRequest("foo",
            iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
            launch_specifications=[
                aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
                    ami="ami-d06a90b0",
                    availability_zone="us-west-2a",
                    instance_type="m1.small",
                    key_name="my-key",
                ),
                aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
                    ami="ami-d06a90b0",
                    availability_zone="us-west-2a",
                    instance_type="m5.large",
                    key_name="my-key",
                ),
            ],
            spot_price="0.005",
            target_capacity=2,
            valid_until="2019-11-04T20:44:20Z")
        ```
        ### Using multiple launch configurations

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2.get_subnet_ids(vpc_id=var["vpc_id"])
        foo_launch_template = aws.ec2.LaunchTemplate("fooLaunchTemplate",
            image_id="ami-516b9131",
            instance_type="m1.small",
            key_name="some-key")
        foo_spot_fleet_request = aws.ec2.SpotFleetRequest("fooSpotFleetRequest",
            iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
            spot_price="0.005",
            target_capacity=2,
            valid_until="2019-11-04T20:44:20Z",
            launch_template_configs=[aws.ec2.SpotFleetRequestLaunchTemplateConfigArgs(
                launch_template_specification=aws.ec2.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs(
                    id=foo_launch_template.id,
                    version=foo_launch_template.latest_version,
                ),
                overrides=[
                    aws.ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs(
                        subnet_id=data["aws_subnets"]["example"]["ids"],
                    ),
                    aws.ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs(
                        subnet_id=data["aws_subnets"]["example"]["ids"],
                    ),
                    aws.ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs(
                        subnet_id=data["aws_subnets"]["example"]["ids"],
                    ),
                ],
            )],
            opts=ResourceOptions(depends_on=[aws_iam_policy_attachment["test-attach"]]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_strategy: Indicates how to allocate the target capacity across
               the Spot pools specified by the Spot fleet request. The default is
               `lowestPrice`.
        :param pulumi.Input[str] excess_capacity_termination_policy: Indicates whether running Spot
               instances should be terminated if the target capacity of the Spot fleet
               request is decreased below the current size of the Spot fleet.
        :param pulumi.Input[str] fleet_type: The type of fleet request. Indicates whether the Spot Fleet only requests the target
               capacity or also attempts to maintain it. Default is `maintain`.
        :param pulumi.Input[str] iam_fleet_role: Grants the Spot fleet permission to terminate
               Spot instances on your behalf when you cancel its Spot fleet request using
               CancelSpotFleetRequests or when the Spot fleet request expires, if you set
               terminateInstancesWithExpiration.
        :param pulumi.Input[str] instance_interruption_behaviour: Indicates whether a Spot
               instance stops or terminates when it is interrupted. Default is
               `terminate`.
        :param pulumi.Input[int] instance_pools_to_use_count: The number of Spot pools across which to allocate your target Spot capacity.
               Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
               the cheapest Spot pools and evenly allocates your target Spot capacity across
               the number of Spot pools that you specify.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchSpecificationArgs']]]] launch_specifications: Used to define the launch configuration of the
               spot-fleet request. Can be specified multiple times to define different bids
               across different markets and instance types. Conflicts with `launch_template_config`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchTemplateConfigArgs']]]] launch_template_configs: Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launch_specification`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] load_balancers: A list of elastic load balancer names to add to the Spot fleet.
        :param pulumi.Input[bool] replace_unhealthy_instances: Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
        :param pulumi.Input[str] spot_price: The maximum spot bid for this override request.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[int] target_capacity: The number of units to request. You can choose to set the
               target capacity in terms of instances or a performance characteristic that is
               important to your application workload, such as vCPUs, memory, or I/O.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_group_arns: A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
        :param pulumi.Input[bool] terminate_instances_with_expiration: Indicates whether running Spot
               instances should be terminated when the Spot fleet request expires.
        :param pulumi.Input[str] valid_from: The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        :param pulumi.Input[str] valid_until: The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
        :param pulumi.Input[bool] wait_for_fulfillment: If set, this provider will
               wait for the Spot Request to be fulfilled, and will throw an error if the
               timeout of 10m is reached.
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

            __props__['allocation_strategy'] = allocation_strategy
            __props__['excess_capacity_termination_policy'] = excess_capacity_termination_policy
            __props__['fleet_type'] = fleet_type
            if iam_fleet_role is None:
                raise TypeError("Missing required property 'iam_fleet_role'")
            __props__['iam_fleet_role'] = iam_fleet_role
            __props__['instance_interruption_behaviour'] = instance_interruption_behaviour
            __props__['instance_pools_to_use_count'] = instance_pools_to_use_count
            __props__['launch_specifications'] = launch_specifications
            __props__['launch_template_configs'] = launch_template_configs
            __props__['load_balancers'] = load_balancers
            __props__['replace_unhealthy_instances'] = replace_unhealthy_instances
            __props__['spot_price'] = spot_price
            __props__['tags'] = tags
            if target_capacity is None:
                raise TypeError("Missing required property 'target_capacity'")
            __props__['target_capacity'] = target_capacity
            __props__['target_group_arns'] = target_group_arns
            __props__['terminate_instances_with_expiration'] = terminate_instances_with_expiration
            __props__['valid_from'] = valid_from
            __props__['valid_until'] = valid_until
            __props__['wait_for_fulfillment'] = wait_for_fulfillment
            __props__['client_token'] = None
            __props__['spot_request_state'] = None
        super(SpotFleetRequest, __self__).__init__(
            'aws:ec2/spotFleetRequest:SpotFleetRequest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_strategy: Optional[pulumi.Input[str]] = None,
            client_token: Optional[pulumi.Input[str]] = None,
            excess_capacity_termination_policy: Optional[pulumi.Input[str]] = None,
            fleet_type: Optional[pulumi.Input[str]] = None,
            iam_fleet_role: Optional[pulumi.Input[str]] = None,
            instance_interruption_behaviour: Optional[pulumi.Input[str]] = None,
            instance_pools_to_use_count: Optional[pulumi.Input[int]] = None,
            launch_specifications: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchSpecificationArgs']]]]] = None,
            launch_template_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchTemplateConfigArgs']]]]] = None,
            load_balancers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            replace_unhealthy_instances: Optional[pulumi.Input[bool]] = None,
            spot_price: Optional[pulumi.Input[str]] = None,
            spot_request_state: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_capacity: Optional[pulumi.Input[int]] = None,
            target_group_arns: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            terminate_instances_with_expiration: Optional[pulumi.Input[bool]] = None,
            valid_from: Optional[pulumi.Input[str]] = None,
            valid_until: Optional[pulumi.Input[str]] = None,
            wait_for_fulfillment: Optional[pulumi.Input[bool]] = None) -> 'SpotFleetRequest':
        """
        Get an existing SpotFleetRequest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_strategy: Indicates how to allocate the target capacity across
               the Spot pools specified by the Spot fleet request. The default is
               `lowestPrice`.
        :param pulumi.Input[str] excess_capacity_termination_policy: Indicates whether running Spot
               instances should be terminated if the target capacity of the Spot fleet
               request is decreased below the current size of the Spot fleet.
        :param pulumi.Input[str] fleet_type: The type of fleet request. Indicates whether the Spot Fleet only requests the target
               capacity or also attempts to maintain it. Default is `maintain`.
        :param pulumi.Input[str] iam_fleet_role: Grants the Spot fleet permission to terminate
               Spot instances on your behalf when you cancel its Spot fleet request using
               CancelSpotFleetRequests or when the Spot fleet request expires, if you set
               terminateInstancesWithExpiration.
        :param pulumi.Input[str] instance_interruption_behaviour: Indicates whether a Spot
               instance stops or terminates when it is interrupted. Default is
               `terminate`.
        :param pulumi.Input[int] instance_pools_to_use_count: The number of Spot pools across which to allocate your target Spot capacity.
               Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
               the cheapest Spot pools and evenly allocates your target Spot capacity across
               the number of Spot pools that you specify.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchSpecificationArgs']]]] launch_specifications: Used to define the launch configuration of the
               spot-fleet request. Can be specified multiple times to define different bids
               across different markets and instance types. Conflicts with `launch_template_config`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SpotFleetRequestLaunchTemplateConfigArgs']]]] launch_template_configs: Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launch_specification`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] load_balancers: A list of elastic load balancer names to add to the Spot fleet.
        :param pulumi.Input[bool] replace_unhealthy_instances: Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
        :param pulumi.Input[str] spot_price: The maximum spot bid for this override request.
        :param pulumi.Input[str] spot_request_state: The state of the Spot fleet request.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[int] target_capacity: The number of units to request. You can choose to set the
               target capacity in terms of instances or a performance characteristic that is
               important to your application workload, such as vCPUs, memory, or I/O.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] target_group_arns: A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
        :param pulumi.Input[bool] terminate_instances_with_expiration: Indicates whether running Spot
               instances should be terminated when the Spot fleet request expires.
        :param pulumi.Input[str] valid_from: The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        :param pulumi.Input[str] valid_until: The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
        :param pulumi.Input[bool] wait_for_fulfillment: If set, this provider will
               wait for the Spot Request to be fulfilled, and will throw an error if the
               timeout of 10m is reached.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocation_strategy"] = allocation_strategy
        __props__["client_token"] = client_token
        __props__["excess_capacity_termination_policy"] = excess_capacity_termination_policy
        __props__["fleet_type"] = fleet_type
        __props__["iam_fleet_role"] = iam_fleet_role
        __props__["instance_interruption_behaviour"] = instance_interruption_behaviour
        __props__["instance_pools_to_use_count"] = instance_pools_to_use_count
        __props__["launch_specifications"] = launch_specifications
        __props__["launch_template_configs"] = launch_template_configs
        __props__["load_balancers"] = load_balancers
        __props__["replace_unhealthy_instances"] = replace_unhealthy_instances
        __props__["spot_price"] = spot_price
        __props__["spot_request_state"] = spot_request_state
        __props__["tags"] = tags
        __props__["target_capacity"] = target_capacity
        __props__["target_group_arns"] = target_group_arns
        __props__["terminate_instances_with_expiration"] = terminate_instances_with_expiration
        __props__["valid_from"] = valid_from
        __props__["valid_until"] = valid_until
        __props__["wait_for_fulfillment"] = wait_for_fulfillment
        return SpotFleetRequest(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationStrategy")
    def allocation_strategy(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates how to allocate the target capacity across
        the Spot pools specified by the Spot fleet request. The default is
        `lowestPrice`.
        """
        return pulumi.get(self, "allocation_strategy")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter(name="excessCapacityTerminationPolicy")
    def excess_capacity_termination_policy(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether running Spot
        instances should be terminated if the target capacity of the Spot fleet
        request is decreased below the current size of the Spot fleet.
        """
        return pulumi.get(self, "excess_capacity_termination_policy")

    @property
    @pulumi.getter(name="fleetType")
    def fleet_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of fleet request. Indicates whether the Spot Fleet only requests the target
        capacity or also attempts to maintain it. Default is `maintain`.
        """
        return pulumi.get(self, "fleet_type")

    @property
    @pulumi.getter(name="iamFleetRole")
    def iam_fleet_role(self) -> pulumi.Output[str]:
        """
        Grants the Spot fleet permission to terminate
        Spot instances on your behalf when you cancel its Spot fleet request using
        CancelSpotFleetRequests or when the Spot fleet request expires, if you set
        terminateInstancesWithExpiration.
        """
        return pulumi.get(self, "iam_fleet_role")

    @property
    @pulumi.getter(name="instanceInterruptionBehaviour")
    def instance_interruption_behaviour(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether a Spot
        instance stops or terminates when it is interrupted. Default is
        `terminate`.
        """
        return pulumi.get(self, "instance_interruption_behaviour")

    @property
    @pulumi.getter(name="instancePoolsToUseCount")
    def instance_pools_to_use_count(self) -> pulumi.Output[Optional[int]]:
        """
        The number of Spot pools across which to allocate your target Spot capacity.
        Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
        the cheapest Spot pools and evenly allocates your target Spot capacity across
        the number of Spot pools that you specify.
        """
        return pulumi.get(self, "instance_pools_to_use_count")

    @property
    @pulumi.getter(name="launchSpecifications")
    def launch_specifications(self) -> pulumi.Output[Optional[Sequence['outputs.SpotFleetRequestLaunchSpecification']]]:
        """
        Used to define the launch configuration of the
        spot-fleet request. Can be specified multiple times to define different bids
        across different markets and instance types. Conflicts with `launch_template_config`. At least one of `launch_specification` or `launch_template_config` is required.
        """
        return pulumi.get(self, "launch_specifications")

    @property
    @pulumi.getter(name="launchTemplateConfigs")
    def launch_template_configs(self) -> pulumi.Output[Optional[Sequence['outputs.SpotFleetRequestLaunchTemplateConfig']]]:
        """
        Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launch_specification`. At least one of `launch_specification` or `launch_template_config` is required.
        """
        return pulumi.get(self, "launch_template_configs")

    @property
    @pulumi.getter(name="loadBalancers")
    def load_balancers(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of elastic load balancer names to add to the Spot fleet.
        """
        return pulumi.get(self, "load_balancers")

    @property
    @pulumi.getter(name="replaceUnhealthyInstances")
    def replace_unhealthy_instances(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
        """
        return pulumi.get(self, "replace_unhealthy_instances")

    @property
    @pulumi.getter(name="spotPrice")
    def spot_price(self) -> pulumi.Output[Optional[str]]:
        """
        The maximum spot bid for this override request.
        """
        return pulumi.get(self, "spot_price")

    @property
    @pulumi.getter(name="spotRequestState")
    def spot_request_state(self) -> pulumi.Output[str]:
        """
        The state of the Spot fleet request.
        """
        return pulumi.get(self, "spot_request_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetCapacity")
    def target_capacity(self) -> pulumi.Output[int]:
        """
        The number of units to request. You can choose to set the
        target capacity in terms of instances or a performance characteristic that is
        important to your application workload, such as vCPUs, memory, or I/O.
        """
        return pulumi.get(self, "target_capacity")

    @property
    @pulumi.getter(name="targetGroupArns")
    def target_group_arns(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
        """
        return pulumi.get(self, "target_group_arns")

    @property
    @pulumi.getter(name="terminateInstancesWithExpiration")
    def terminate_instances_with_expiration(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether running Spot
        instances should be terminated when the Spot fleet request expires.
        """
        return pulumi.get(self, "terminate_instances_with_expiration")

    @property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> pulumi.Output[Optional[str]]:
        """
        The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        """
        return pulumi.get(self, "valid_from")

    @property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> pulumi.Output[Optional[str]]:
        """
        The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
        """
        return pulumi.get(self, "valid_until")

    @property
    @pulumi.getter(name="waitForFulfillment")
    def wait_for_fulfillment(self) -> pulumi.Output[Optional[bool]]:
        """
        If set, this provider will
        wait for the Spot Request to be fulfilled, and will throw an error if the
        timeout of 10m is reached.
        """
        return pulumi.get(self, "wait_for_fulfillment")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

