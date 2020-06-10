# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class CapacityProvider(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) that identifies the capacity provider.
    """
    auto_scaling_group_provider: pulumi.Output[dict]
    """
    Nested argument defining the provider for the ECS auto scaling group. Defined below.

      * `autoScalingGroupArn` (`str`) - - The Amazon Resource Name (ARN) of the associated auto scaling group.
      * `managedScaling` (`dict`) - - Nested argument defining the parameters of the auto scaling. Defined below.
        * `maximumScalingStepSize` (`float`) - The maximum step adjustment size. A number between 1 and 10,000.
        * `minimumScalingStepSize` (`float`) - The minimum step adjustment size. A number between 1 and 10,000.
        * `status` (`str`) - Whether auto scaling is managed by ECS. Valid values are `ENABLED` and `DISABLED`.
        * `target_capacity` (`float`) - The target utilization for the capacity provider. A number between 1 and 100.

      * `managedTerminationProtection` (`str`) - - Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are `ENABLED` and `DISABLED`.
    """
    name: pulumi.Output[str]
    """
    The name of the capacity provider.
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of resource tags.
    """
    def __init__(__self__, resource_name, opts=None, auto_scaling_group_provider=None, name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an ECS cluster capacity provider. More information can be found on the [ECS Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).

        > **NOTE:** The AWS API does not currently support deleting ECS cluster capacity providers. Removing this resource will only remove the state for it.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.ecs.CapacityProvider("test", auto_scaling_group_provider={
            "autoScalingGroupArn": aws_autoscaling_group["test"]["arn"],
            "managedTerminationProtection": "ENABLED",
            "managed_scaling": {
                "maximumScalingStepSize": 1000,
                "minimumScalingStepSize": 1,
                "status": "ENABLED",
                "target_capacity": 10,
            },
        })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] auto_scaling_group_provider: Nested argument defining the provider for the ECS auto scaling group. Defined below.
        :param pulumi.Input[str] name: The name of the capacity provider.
        :param pulumi.Input[dict] tags: Key-value map of resource tags.

        The **auto_scaling_group_provider** object supports the following:

          * `autoScalingGroupArn` (`pulumi.Input[str]`) - - The Amazon Resource Name (ARN) of the associated auto scaling group.
          * `managedScaling` (`pulumi.Input[dict]`) - - Nested argument defining the parameters of the auto scaling. Defined below.
            * `maximumScalingStepSize` (`pulumi.Input[float]`) - The maximum step adjustment size. A number between 1 and 10,000.
            * `minimumScalingStepSize` (`pulumi.Input[float]`) - The minimum step adjustment size. A number between 1 and 10,000.
            * `status` (`pulumi.Input[str]`) - Whether auto scaling is managed by ECS. Valid values are `ENABLED` and `DISABLED`.
            * `target_capacity` (`pulumi.Input[float]`) - The target utilization for the capacity provider. A number between 1 and 100.

          * `managedTerminationProtection` (`pulumi.Input[str]`) - - Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are `ENABLED` and `DISABLED`.
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
    def get(resource_name, id, opts=None, arn=None, auto_scaling_group_provider=None, name=None, tags=None):
        """
        Get an existing CapacityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the capacity provider.
        :param pulumi.Input[dict] auto_scaling_group_provider: Nested argument defining the provider for the ECS auto scaling group. Defined below.
        :param pulumi.Input[str] name: The name of the capacity provider.
        :param pulumi.Input[dict] tags: Key-value map of resource tags.

        The **auto_scaling_group_provider** object supports the following:

          * `autoScalingGroupArn` (`pulumi.Input[str]`) - - The Amazon Resource Name (ARN) of the associated auto scaling group.
          * `managedScaling` (`pulumi.Input[dict]`) - - Nested argument defining the parameters of the auto scaling. Defined below.
            * `maximumScalingStepSize` (`pulumi.Input[float]`) - The maximum step adjustment size. A number between 1 and 10,000.
            * `minimumScalingStepSize` (`pulumi.Input[float]`) - The minimum step adjustment size. A number between 1 and 10,000.
            * `status` (`pulumi.Input[str]`) - Whether auto scaling is managed by ECS. Valid values are `ENABLED` and `DISABLED`.
            * `target_capacity` (`pulumi.Input[float]`) - The target utilization for the capacity provider. A number between 1 and 100.

          * `managedTerminationProtection` (`pulumi.Input[str]`) - - Enables or disables container-aware termination of instances in the auto scaling group when scale-in happens. Valid values are `ENABLED` and `DISABLED`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["auto_scaling_group_provider"] = auto_scaling_group_provider
        __props__["name"] = name
        __props__["tags"] = tags
        return CapacityProvider(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
