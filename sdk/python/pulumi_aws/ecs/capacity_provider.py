# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['CapacityProvider']


class CapacityProvider(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The Amazon Resource Name (ARN) that identifies the capacity provider.
    """

    auto_scaling_group_provider: pulumi.Output['outputs.CapacityProviderAutoScalingGroupProvider'] = pulumi.property("autoScalingGroupProvider")
    """
    Nested argument defining the provider for the ECS auto scaling group. Defined below.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the capacity provider.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    Key-value map of resource tags.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_scaling_group_provider: Optional[pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).

        > **NOTE:** Associating an ECS Capacity Provider to an Auto Scaling Group will automatically add the `AmazonECSManaged` tag to the Auto Scaling Group. This tag should be included in the `autoscaling.Group` resource configuration to prevent the provider from removing it in subsequent executions as well as ensuring the `AmazonECSManaged` tag is propagated to all EC2 Instances in the Auto Scaling Group if `min_size` is above 0 on creation. Any EC2 Instances in the Auto Scaling Group without this tag must be manually be updated, otherwise they may cause unexpected scaling behavior and metrics.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configuration, including potentially other tags ...
        test_group = aws.autoscaling.Group("testGroup", tags=[{
            "key": "AmazonECSManaged",
            "propagateAtLaunch": True,
        }])
        test_capacity_provider = aws.ecs.CapacityProvider("testCapacityProvider", auto_scaling_group_provider={
            "autoScalingGroupArn": test_group.arn,
            "managedTerminationProtection": "ENABLED",
            "managedScaling": {
                "maximumScalingStepSize": 1000,
                "minimumScalingStepSize": 1,
                "status": "ENABLED",
                "target_capacity": 10,
            },
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']] auto_scaling_group_provider: Nested argument defining the provider for the ECS auto scaling group. Defined below.
        :param pulumi.Input[str] name: The name of the capacity provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
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

            if auto_scaling_group_provider is None:
                raise TypeError("Missing required property 'auto_scaling_group_provider'")
            __props__['auto_scaling_group_provider'] = auto_scaling_group_provider
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
        super(CapacityProvider, __self__).__init__(
            'aws:ecs/capacityProvider:CapacityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            auto_scaling_group_provider: Optional[pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'CapacityProvider':
        """
        Get an existing CapacityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the capacity provider.
        :param pulumi.Input[pulumi.InputType['CapacityProviderAutoScalingGroupProviderArgs']] auto_scaling_group_provider: Nested argument defining the provider for the ECS auto scaling group. Defined below.
        :param pulumi.Input[str] name: The name of the capacity provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["auto_scaling_group_provider"] = auto_scaling_group_provider
        __props__["name"] = name
        __props__["tags"] = tags
        return CapacityProvider(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

