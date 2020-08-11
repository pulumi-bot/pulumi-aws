# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


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
    across different markets and instance types. Conflicts with `launch_template_config`. At least one of `launch_specification` or `launch_template_config` is required.

      * `ami` (`str`)
      * `associate_public_ip_address` (`bool`)
      * `availability_zone` (`str`) - The availability zone in which to place the request.
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
      * `instance_type` (`str`) - The type of instance to request.
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

      * `spot_price` (`str`) - The maximum spot bid for this override request.
      * `subnet_id` (`str`) - The subnet in which to launch the requested instance.
      * `tags` (`dict`) - A map of tags to assign to the resource.
      * `user_data` (`str`)
      * `vpc_security_group_ids` (`list`)
      * `weightedCapacity` (`str`) - The capacity added to the fleet by a fulfilled request.
    """
    launch_template_configs: pulumi.Output[list]
    """
    Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launch_specification`. At least one of `launch_specification` or `launch_template_config` is required.

      * `launchTemplateSpecification` (`dict`) - Launch template specification. See Launch Template Specification below for more details.
        * `id` (`str`) - The ID of the launch template. Conflicts with `name`.
        * `name` (`str`) - The name of the launch template. Conflicts with `id`.
        * `version` (`str`) - Template version. Unlike the autoscaling equivalent, does not support `$Latest` or `$Default`, so use the launch_template resource's attribute, e.g. `"${aws_launch_template.foo.latest_version}"`. It will use the default version if omitted.

      * `overrides` (`list`) - One or more override configurations. See Overrides below for more details.
        * `availability_zone` (`str`) - The availability zone in which to place the request.
        * `instance_type` (`str`) - The type of instance to request.
        * `priority` (`float`) - The priority for the launch template override. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority.
        * `spot_price` (`str`) - The maximum spot bid for this override request.
        * `subnet_id` (`str`) - The subnet in which to launch the requested instance.
        * `weightedCapacity` (`float`) - The capacity added to the fleet by a fulfilled request.
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
    The maximum spot bid for this override request.
    """
    spot_request_state: pulumi.Output[str]
    """
    The state of the Spot fleet request.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
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
    The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
    """
    wait_for_fulfillment: pulumi.Output[bool]
    """
    If set, this provider will
    wait for the Spot Request to be fulfilled, and will throw an error if the
    timeout of 10m is reached.
    """
    def __init__(__self__, resource_name, opts=None, allocation_strategy=None, excess_capacity_termination_policy=None, fleet_type=None, iam_fleet_role=None, instance_interruption_behaviour=None, instance_pools_to_use_count=None, launch_specifications=None, launch_template_configs=None, load_balancers=None, replace_unhealthy_instances=None, spot_price=None, tags=None, target_capacity=None, target_group_arns=None, terminate_instances_with_expiration=None, valid_from=None, valid_until=None, wait_for_fulfillment=None, __props__=None, __name__=None, __opts__=None):
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
                {
                    "instance_type": "m4.10xlarge",
                    "ami": "ami-1234",
                    "spot_price": "2.793",
                    "placement_tenancy": "dedicated",
                    "iamInstanceProfileArn": aws_iam_instance_profile["example"]["arn"],
                },
                {
                    "instance_type": "m4.4xlarge",
                    "ami": "ami-5678",
                    "key_name": "my-key",
                    "spot_price": "1.117",
                    "iamInstanceProfileArn": aws_iam_instance_profile["example"]["arn"],
                    "availability_zone": "us-west-1a",
                    "subnet_id": "subnet-1234",
                    "weightedCapacity": "35",
                    "root_block_devices": [{
                        "volume_size": "300",
                        "volumeType": "gp2",
                    }],
                    "tags": {
                        "Name": "spot-fleet-example",
                    },
                },
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
            launch_template_configs=[{
                "launchTemplateSpecification": {
                    "id": foo_launch_template.id,
                    "version": foo_launch_template.latest_version,
                },
            }],
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
                {
                    "ami": "ami-d06a90b0",
                    "availability_zone": "us-west-2a",
                    "instance_type": "m1.small",
                    "key_name": "my-key",
                },
                {
                    "ami": "ami-d06a90b0",
                    "availability_zone": "us-west-2a",
                    "instance_type": "m5.large",
                    "key_name": "my-key",
                },
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
            launch_template_configs=[{
                "launchTemplateSpecification": {
                    "id": foo_launch_template.id,
                    "version": foo_launch_template.latest_version,
                },
                "overrides": [
                    {
                        "subnet_id": data["aws_subnets"]["example"]["ids"],
                    },
                    {
                        "subnet_id": data["aws_subnets"]["example"]["ids"],
                    },
                    {
                        "subnet_id": data["aws_subnets"]["example"]["ids"],
                    },
                ],
            }],
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
        :param pulumi.Input[float] instance_pools_to_use_count: The number of Spot pools across which to allocate your target Spot capacity.
               Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
               the cheapest Spot pools and evenly allocates your target Spot capacity across
               the number of Spot pools that you specify.
        :param pulumi.Input[list] launch_specifications: Used to define the launch configuration of the
               spot-fleet request. Can be specified multiple times to define different bids
               across different markets and instance types. Conflicts with `launch_template_config`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[list] launch_template_configs: Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launch_specification`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[list] load_balancers: A list of elastic load balancer names to add to the Spot fleet.
        :param pulumi.Input[bool] replace_unhealthy_instances: Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
        :param pulumi.Input[str] spot_price: The maximum spot bid for this override request.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] target_capacity: The number of units to request. You can choose to set the
               target capacity in terms of instances or a performance characteristic that is
               important to your application workload, such as vCPUs, memory, or I/O.
        :param pulumi.Input[list] target_group_arns: A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
        :param pulumi.Input[bool] terminate_instances_with_expiration: Indicates whether running Spot
               instances should be terminated when the Spot fleet request expires.
        :param pulumi.Input[str] valid_from: The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        :param pulumi.Input[str] valid_until: The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
        :param pulumi.Input[bool] wait_for_fulfillment: If set, this provider will
               wait for the Spot Request to be fulfilled, and will throw an error if the
               timeout of 10m is reached.

        The **launch_specifications** object supports the following:

          * `ami` (`pulumi.Input[str]`)
          * `associate_public_ip_address` (`pulumi.Input[bool]`)
          * `availability_zone` (`pulumi.Input[str]`) - The availability zone in which to place the request.
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
          * `instance_type` (`pulumi.Input[str]`) - The type of instance to request.
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

          * `spot_price` (`pulumi.Input[str]`) - The maximum spot bid for this override request.
          * `subnet_id` (`pulumi.Input[str]`) - The subnet in which to launch the requested instance.
          * `tags` (`pulumi.Input[dict]`) - A map of tags to assign to the resource.
          * `user_data` (`pulumi.Input[str]`)
          * `vpc_security_group_ids` (`pulumi.Input[list]`)
          * `weightedCapacity` (`pulumi.Input[str]`) - The capacity added to the fleet by a fulfilled request.

        The **launch_template_configs** object supports the following:

          * `launchTemplateSpecification` (`pulumi.Input[dict]`) - Launch template specification. See Launch Template Specification below for more details.
            * `id` (`pulumi.Input[str]`) - The ID of the launch template. Conflicts with `name`.
            * `name` (`pulumi.Input[str]`) - The name of the launch template. Conflicts with `id`.
            * `version` (`pulumi.Input[str]`) - Template version. Unlike the autoscaling equivalent, does not support `$Latest` or `$Default`, so use the launch_template resource's attribute, e.g. `"${aws_launch_template.foo.latest_version}"`. It will use the default version if omitted.

          * `overrides` (`pulumi.Input[list]`) - One or more override configurations. See Overrides below for more details.
            * `availability_zone` (`pulumi.Input[str]`) - The availability zone in which to place the request.
            * `instance_type` (`pulumi.Input[str]`) - The type of instance to request.
            * `priority` (`pulumi.Input[float]`) - The priority for the launch template override. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority.
            * `spot_price` (`pulumi.Input[str]`) - The maximum spot bid for this override request.
            * `subnet_id` (`pulumi.Input[str]`) - The subnet in which to launch the requested instance.
            * `weightedCapacity` (`pulumi.Input[float]`) - The capacity added to the fleet by a fulfilled request.
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
    def get(resource_name, id, opts=None, allocation_strategy=None, client_token=None, excess_capacity_termination_policy=None, fleet_type=None, iam_fleet_role=None, instance_interruption_behaviour=None, instance_pools_to_use_count=None, launch_specifications=None, launch_template_configs=None, load_balancers=None, replace_unhealthy_instances=None, spot_price=None, spot_request_state=None, tags=None, target_capacity=None, target_group_arns=None, terminate_instances_with_expiration=None, valid_from=None, valid_until=None, wait_for_fulfillment=None):
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
        :param pulumi.Input[float] instance_pools_to_use_count: The number of Spot pools across which to allocate your target Spot capacity.
               Valid only when `allocation_strategy` is set to `lowestPrice`. Spot Fleet selects
               the cheapest Spot pools and evenly allocates your target Spot capacity across
               the number of Spot pools that you specify.
        :param pulumi.Input[list] launch_specifications: Used to define the launch configuration of the
               spot-fleet request. Can be specified multiple times to define different bids
               across different markets and instance types. Conflicts with `launch_template_config`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[list] launch_template_configs: Launch template configuration block. See Launch Template Configs below for more details. Conflicts with `launch_specification`. At least one of `launch_specification` or `launch_template_config` is required.
        :param pulumi.Input[list] load_balancers: A list of elastic load balancer names to add to the Spot fleet.
        :param pulumi.Input[bool] replace_unhealthy_instances: Indicates whether Spot fleet should replace unhealthy instances. Default `false`.
        :param pulumi.Input[str] spot_price: The maximum spot bid for this override request.
        :param pulumi.Input[str] spot_request_state: The state of the Spot fleet request.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[float] target_capacity: The number of units to request. You can choose to set the
               target capacity in terms of instances or a performance characteristic that is
               important to your application workload, such as vCPUs, memory, or I/O.
        :param pulumi.Input[list] target_group_arns: A list of `alb.TargetGroup` ARNs, for use with Application Load Balancing.
        :param pulumi.Input[bool] terminate_instances_with_expiration: Indicates whether running Spot
               instances should be terminated when the Spot fleet request expires.
        :param pulumi.Input[str] valid_from: The start date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). The default is to start fulfilling the request immediately.
        :param pulumi.Input[str] valid_until: The end date and time of the request, in UTC [RFC3339](https://tools.ietf.org/html/rfc3339#section-5.8) format(for example, YYYY-MM-DDTHH:MM:SSZ). At this point, no new Spot instance requests are placed or enabled to fulfill the request.
        :param pulumi.Input[bool] wait_for_fulfillment: If set, this provider will
               wait for the Spot Request to be fulfilled, and will throw an error if the
               timeout of 10m is reached.

        The **launch_specifications** object supports the following:

          * `ami` (`pulumi.Input[str]`)
          * `associate_public_ip_address` (`pulumi.Input[bool]`)
          * `availability_zone` (`pulumi.Input[str]`) - The availability zone in which to place the request.
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
          * `instance_type` (`pulumi.Input[str]`) - The type of instance to request.
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

          * `spot_price` (`pulumi.Input[str]`) - The maximum spot bid for this override request.
          * `subnet_id` (`pulumi.Input[str]`) - The subnet in which to launch the requested instance.
          * `tags` (`pulumi.Input[dict]`) - A map of tags to assign to the resource.
          * `user_data` (`pulumi.Input[str]`)
          * `vpc_security_group_ids` (`pulumi.Input[list]`)
          * `weightedCapacity` (`pulumi.Input[str]`) - The capacity added to the fleet by a fulfilled request.

        The **launch_template_configs** object supports the following:

          * `launchTemplateSpecification` (`pulumi.Input[dict]`) - Launch template specification. See Launch Template Specification below for more details.
            * `id` (`pulumi.Input[str]`) - The ID of the launch template. Conflicts with `name`.
            * `name` (`pulumi.Input[str]`) - The name of the launch template. Conflicts with `id`.
            * `version` (`pulumi.Input[str]`) - Template version. Unlike the autoscaling equivalent, does not support `$Latest` or `$Default`, so use the launch_template resource's attribute, e.g. `"${aws_launch_template.foo.latest_version}"`. It will use the default version if omitted.

          * `overrides` (`pulumi.Input[list]`) - One or more override configurations. See Overrides below for more details.
            * `availability_zone` (`pulumi.Input[str]`) - The availability zone in which to place the request.
            * `instance_type` (`pulumi.Input[str]`) - The type of instance to request.
            * `priority` (`pulumi.Input[float]`) - The priority for the launch template override. The lower the number, the higher the priority. If no number is set, the launch template override has the lowest priority.
            * `spot_price` (`pulumi.Input[str]`) - The maximum spot bid for this override request.
            * `subnet_id` (`pulumi.Input[str]`) - The subnet in which to launch the requested instance.
            * `weightedCapacity` (`pulumi.Input[float]`) - The capacity added to the fleet by a fulfilled request.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
