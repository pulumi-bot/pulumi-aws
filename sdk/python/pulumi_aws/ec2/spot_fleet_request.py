# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SpotFleetRequest(pulumi.CustomResource):
    allocation_strategy: pulumi.Output[str]
    """
    Indicates how to allocate the target capacity across
    the Spot pools specified by the Spot fleet request. The default is
    `lowestPrice`.
    """
    client_token: pulumi.Output[str]
    excess_capacity_termination_policy: pulumi.Output[str]
    """
    Indicates whether running Spot
    instances should be terminated if the target capacity of the Spot fleet
    request is decreased below the current size of the Spot fleet.
    """
    fleet_type: pulumi.Output[str]
    """
    The type of fleet request. Indicates whether the Spot Fleet only requests the target
    capacity or also attempts to maintain it. Default is `maintain`.
    """
    iam_fleet_role: pulumi.Output[str]
    """
    Grants the Spot fleet permission to terminate
    Spot instances on your behalf when you cancel its Spot fleet request using
    CancelSpotFleetRequests or when the Spot fleet request expires, if you set
    terminateInstancesWithExpiration.
    """
    instance_interruption_behaviour: pulumi.Output[str]
    """
    Indicates whether a Spot
    instance stops or terminates when it is interrupted. Default is
    `terminate`.
    """
    instance_pools_to_use_count: pulumi.Output[float]
    """

    The number of Spot pools across which to allocate your target Spot capacity.
    Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
    the cheapest Spot pools and evenly allocates your target Spot capacity across
    the number of Spot pools that you specify.
    """
    launch_specifications: pulumi.Output[list]
    """
    Used to define the launch configuration of the
    spot-fleet request. Can be specified multiple times to define different bids
    across different markets and instance types.

      * `ami` (`str`)
      * `associate_public_ip_address` (`bool`)
      * `availability_zone` (`str`)
      * `ebs_block_devices` (`list`)
        * `deleteOnTermination` (`bool`)
        * `device_name` (`str`)
        * `encrypted` (`bool`)
        * `iops` (`float`)
        * `kms_key_id` (`str`)
        * `snapshot_id` (`str`)
        * `volume_size` (`float`)
        * `volumeType` (`str`)

      * `ebs_optimized` (`bool`)
      * `ephemeral_block_devices` (`list`)
        * `device_name` (`str`)
        * `virtualName` (`str`)

      * `iam_instance_profile` (`str`)
      * `iamInstanceProfileArn` (`str`)
      * `instance_type` (`str`)
      * `key_name` (`str`)
      * `monitoring` (`bool`)
      * `placement_group` (`str`)
      * `placement_tenancy` (`str`)
      * `root_block_devices` (`list`)
        * `deleteOnTermination` (`bool`)
        * `encrypted` (`bool`)
        * `iops` (`float`)
        * `kms_key_id` (`str`)
        * `volume_size` (`float`)
        * `volumeType` (`str`)

      * `spot_price` (`str`) - The maximum bid price per unit hour.
      * `subnet_id` (`str`)
      * `tags` (`dict`) - A mapping of tags to assign to the resource.
      * `user_data` (`str`)
      * `vpc_security_group_ids` (`list`)
      * `weightedCapacity` (`str`)
    """
    load_balancers: pulumi.Output[list]
    """
    A list of elastic load balancer names to add to the Spot fleet.
    """
    replace_unhealthy_instances: pulumi.Output[bool]
    """
    Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
    """
    spot_price: pulumi.Output[str]
    """
    The maximum bid price per unit hour.
    """
    spot_request_state: pulumi.Output[str]
    """
    The state of the Spot fleet request.
    """
    target_capacity: pulumi.Output[float]
    """
    The number of units to request. You can choose to set the
    target capacity in terms of instances or a performance characteristic that is
    important to your application workload, such as vCPUs, memory, or I/O.
    """
    target_group_arns: pulumi.Output[list]
    """
    A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
    """
    terminate_instances_with_expiration: pulumi.Output[bool]
    """
    Indicates whether running Spot
    instances should be terminated when the Spot fleet request expires.
    """
    valid_from: pulumi.Output[str]
    """
    The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
    """
    valid_until: pulumi.Output[str]
    """
    The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request. Defaults to 24 hours.
    """
    wait_for_fulfillment: pulumi.Output[bool]
    """
    If set, this provider will
    wait for the Spot Request to be fulfilled, and will throw an error if the
    timeout of 10m is reached.
    """
    def __init__(__self__, resource_name, opts=None, allocation_strategy=None, excess_capacity_termination_policy=None, fleet_type=None, iam_fleet_role=None, instance_interruption_behaviour=None, instance_pools_to_use_count=None, launch_specifications=None, load_balancers=None, replace_unhealthy_instances=None, spot_price=None, target_capacity=None, target_group_arns=None, terminate_instances_with_expiration=None, valid_from=None, valid_until=None, wait_for_fulfillment=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an EC2 Spot Fleet Request resource. This allows a fleet of Spot
        instances to be requested on the Spot market.

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/spot_fleet_request.html.markdown.

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
        :param pulumi.Input[float] instance_pools_to_use_count: 
               The number of Spot pools across which to allocate your target Spot capacity.
               Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
               the cheapest Spot pools and evenly allocates your target Spot capacity across
               the number of Spot pools that you specify.
        :param pulumi.Input[list] launch_specifications: Used to define the launch configuration of the
               spot-fleet request. Can be specified multiple times to define different bids
               across different markets and instance types.
        :param pulumi.Input[list] load_balancers: A list of elastic load balancer names to add to the Spot fleet.
        :param pulumi.Input[bool] replace_unhealthy_instances: Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
        :param pulumi.Input[str] spot_price: The maximum bid price per unit hour.
        :param pulumi.Input[float] target_capacity: The number of units to request. You can choose to set the
               target capacity in terms of instances or a performance characteristic that is
               important to your application workload, such as vCPUs, memory, or I/O.
        :param pulumi.Input[list] target_group_arns: A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
        :param pulumi.Input[bool] terminate_instances_with_expiration: Indicates whether running Spot
               instances should be terminated when the Spot fleet request expires.
        :param pulumi.Input[str] valid_from: The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        :param pulumi.Input[str] valid_until: The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request. Defaults to 24 hours.
        :param pulumi.Input[bool] wait_for_fulfillment: If set, this provider will
               wait for the Spot Request to be fulfilled, and will throw an error if the
               timeout of 10m is reached.

        The **launch_specifications** object supports the following:

          * `ami` (`pulumi.Input[str]`)
          * `associate_public_ip_address` (`pulumi.Input[bool]`)
          * `availability_zone` (`pulumi.Input[str]`)
          * `ebs_block_devices` (`pulumi.Input[list]`)
            * `deleteOnTermination` (`pulumi.Input[bool]`)
            * `device_name` (`pulumi.Input[str]`)
            * `encrypted` (`pulumi.Input[bool]`)
            * `iops` (`pulumi.Input[float]`)
            * `kms_key_id` (`pulumi.Input[str]`)
            * `snapshot_id` (`pulumi.Input[str]`)
            * `volume_size` (`pulumi.Input[float]`)
            * `volumeType` (`pulumi.Input[str]`)

          * `ebs_optimized` (`pulumi.Input[bool]`)
          * `ephemeral_block_devices` (`pulumi.Input[list]`)
            * `device_name` (`pulumi.Input[str]`)
            * `virtualName` (`pulumi.Input[str]`)

          * `iam_instance_profile` (`pulumi.Input[str]`)
          * `iamInstanceProfileArn` (`pulumi.Input[str]`)
          * `instance_type` (`pulumi.Input[str]`)
          * `key_name` (`pulumi.Input[str]`)
          * `monitoring` (`pulumi.Input[bool]`)
          * `placement_group` (`pulumi.Input[str]`)
          * `placement_tenancy` (`pulumi.Input[str]`)
          * `root_block_devices` (`pulumi.Input[list]`)
            * `deleteOnTermination` (`pulumi.Input[bool]`)
            * `encrypted` (`pulumi.Input[bool]`)
            * `iops` (`pulumi.Input[float]`)
            * `kms_key_id` (`pulumi.Input[str]`)
            * `volume_size` (`pulumi.Input[float]`)
            * `volumeType` (`pulumi.Input[str]`)

          * `spot_price` (`pulumi.Input[str]`) - The maximum bid price per unit hour.
          * `subnet_id` (`pulumi.Input[str]`)
          * `tags` (`pulumi.Input[dict]`) - A mapping of tags to assign to the resource.
          * `user_data` (`pulumi.Input[str]`)
          * `vpc_security_group_ids` (`pulumi.Input[list]`)
          * `weightedCapacity` (`pulumi.Input[str]`)
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

            __props__['allocation_strategy'] = allocation_strategy
            __props__['excess_capacity_termination_policy'] = excess_capacity_termination_policy
            __props__['fleet_type'] = fleet_type
            if iam_fleet_role is None:
                raise TypeError("Missing required property 'iam_fleet_role'")
            __props__['iam_fleet_role'] = iam_fleet_role
            __props__['instance_interruption_behaviour'] = instance_interruption_behaviour
            __props__['instance_pools_to_use_count'] = instance_pools_to_use_count
            if launch_specifications is None:
                raise TypeError("Missing required property 'launch_specifications'")
            __props__['launch_specifications'] = launch_specifications
            __props__['load_balancers'] = load_balancers
            __props__['replace_unhealthy_instances'] = replace_unhealthy_instances
            __props__['spot_price'] = spot_price
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
    def get(resource_name, id, opts=None, allocation_strategy=None, client_token=None, excess_capacity_termination_policy=None, fleet_type=None, iam_fleet_role=None, instance_interruption_behaviour=None, instance_pools_to_use_count=None, launch_specifications=None, load_balancers=None, replace_unhealthy_instances=None, spot_price=None, spot_request_state=None, target_capacity=None, target_group_arns=None, terminate_instances_with_expiration=None, valid_from=None, valid_until=None, wait_for_fulfillment=None):
        """
        Get an existing SpotFleetRequest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
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
        :param pulumi.Input[float] instance_pools_to_use_count: 
               The number of Spot pools across which to allocate your target Spot capacity.
               Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
               the cheapest Spot pools and evenly allocates your target Spot capacity across
               the number of Spot pools that you specify.
        :param pulumi.Input[list] launch_specifications: Used to define the launch configuration of the
               spot-fleet request. Can be specified multiple times to define different bids
               across different markets and instance types.
        :param pulumi.Input[list] load_balancers: A list of elastic load balancer names to add to the Spot fleet.
        :param pulumi.Input[bool] replace_unhealthy_instances: Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
        :param pulumi.Input[str] spot_price: The maximum bid price per unit hour.
        :param pulumi.Input[str] spot_request_state: The state of the Spot fleet request.
        :param pulumi.Input[float] target_capacity: The number of units to request. You can choose to set the
               target capacity in terms of instances or a performance characteristic that is
               important to your application workload, such as vCPUs, memory, or I/O.
        :param pulumi.Input[list] target_group_arns: A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
        :param pulumi.Input[bool] terminate_instances_with_expiration: Indicates whether running Spot
               instances should be terminated when the Spot fleet request expires.
        :param pulumi.Input[str] valid_from: The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        :param pulumi.Input[str] valid_until: The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request. Defaults to 24 hours.
        :param pulumi.Input[bool] wait_for_fulfillment: If set, this provider will
               wait for the Spot Request to be fulfilled, and will throw an error if the
               timeout of 10m is reached.

        The **launch_specifications** object supports the following:

          * `ami` (`pulumi.Input[str]`)
          * `associate_public_ip_address` (`pulumi.Input[bool]`)
          * `availability_zone` (`pulumi.Input[str]`)
          * `ebs_block_devices` (`pulumi.Input[list]`)
            * `deleteOnTermination` (`pulumi.Input[bool]`)
            * `device_name` (`pulumi.Input[str]`)
            * `encrypted` (`pulumi.Input[bool]`)
            * `iops` (`pulumi.Input[float]`)
            * `kms_key_id` (`pulumi.Input[str]`)
            * `snapshot_id` (`pulumi.Input[str]`)
            * `volume_size` (`pulumi.Input[float]`)
            * `volumeType` (`pulumi.Input[str]`)

          * `ebs_optimized` (`pulumi.Input[bool]`)
          * `ephemeral_block_devices` (`pulumi.Input[list]`)
            * `device_name` (`pulumi.Input[str]`)
            * `virtualName` (`pulumi.Input[str]`)

          * `iam_instance_profile` (`pulumi.Input[str]`)
          * `iamInstanceProfileArn` (`pulumi.Input[str]`)
          * `instance_type` (`pulumi.Input[str]`)
          * `key_name` (`pulumi.Input[str]`)
          * `monitoring` (`pulumi.Input[bool]`)
          * `placement_group` (`pulumi.Input[str]`)
          * `placement_tenancy` (`pulumi.Input[str]`)
          * `root_block_devices` (`pulumi.Input[list]`)
            * `deleteOnTermination` (`pulumi.Input[bool]`)
            * `encrypted` (`pulumi.Input[bool]`)
            * `iops` (`pulumi.Input[float]`)
            * `kms_key_id` (`pulumi.Input[str]`)
            * `volume_size` (`pulumi.Input[float]`)
            * `volumeType` (`pulumi.Input[str]`)

          * `spot_price` (`pulumi.Input[str]`) - The maximum bid price per unit hour.
          * `subnet_id` (`pulumi.Input[str]`)
          * `tags` (`pulumi.Input[dict]`) - A mapping of tags to assign to the resource.
          * `user_data` (`pulumi.Input[str]`)
          * `vpc_security_group_ids` (`pulumi.Input[list]`)
          * `weightedCapacity` (`pulumi.Input[str]`)
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
        __props__["load_balancers"] = load_balancers
        __props__["replace_unhealthy_instances"] = replace_unhealthy_instances
        __props__["spot_price"] = spot_price
        __props__["spot_request_state"] = spot_request_state
        __props__["target_capacity"] = target_capacity
        __props__["target_group_arns"] = target_group_arns
        __props__["terminate_instances_with_expiration"] = terminate_instances_with_expiration
        __props__["valid_from"] = valid_from
        __props__["valid_until"] = valid_until
        __props__["wait_for_fulfillment"] = wait_for_fulfillment
        return SpotFleetRequest(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

