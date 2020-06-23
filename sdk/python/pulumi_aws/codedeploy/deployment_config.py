# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class DeploymentConfig(pulumi.CustomResource):
    compute_platform: pulumi.Output[str]
    """
    The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
    """
    deployment_config_id: pulumi.Output[str]
    """
    The AWS Assigned deployment config id
    """
    deployment_config_name: pulumi.Output[str]
    """
    The name of the deployment config.
    """
    minimum_healthy_hosts: pulumi.Output[dict]
    """
    A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.

      * `type` (`str`) - The type can either be `FLEET_PERCENT` or `HOST_COUNT`.
      * `value` (`float`) - The value when the type is `FLEET_PERCENT` represents the minimum number of healthy instances as
        a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the
        deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
        When the type is `HOST_COUNT`, the value represents the minimum number of healthy instances as an absolute value.
    """
    traffic_routing_config: pulumi.Output[dict]
    """
    A traffic_routing_config block. Traffic Routing Config is documented below.

      * `timeBasedCanary` (`dict`) - The time based canary configuration information. If `type` is `TimeBasedLinear`, use `time_based_linear` instead.
        * `interval` (`float`) - The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
        * `percentage` (`float`) - The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.

      * `timeBasedLinear` (`dict`) - The time based linear configuration information. If `type` is `TimeBasedCanary`, use `time_based_canary` instead.
        * `interval` (`float`) - The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
        * `percentage` (`float`) - The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.

      * `type` (`str`) - Type of traffic routing config. One of `TimeBasedCanary`, `TimeBasedLinear`, `AllAtOnce`.
    """
    def __init__(__self__, resource_name, opts=None, compute_platform=None, deployment_config_name=None, minimum_healthy_hosts=None, traffic_routing_config=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CodeDeploy deployment config for an application

        ## Example Usage
        ### Server Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_deployment_config = aws.codedeploy.DeploymentConfig("fooDeploymentConfig",
            deployment_config_name="test-deployment-config",
            minimum_healthy_hosts={
                "type": "HOST_COUNT",
                "value": 2,
            })
        foo_deployment_group = aws.codedeploy.DeploymentGroup("fooDeploymentGroup",
            alarm_configuration={
                "alarms": ["my-alarm-name"],
                "enabled": True,
            },
            app_name=aws_codedeploy_app["foo_app"]["name"],
            auto_rollback_configuration={
                "enabled": True,
                "events": ["DEPLOYMENT_FAILURE"],
            },
            deployment_config_name=foo_deployment_config.id,
            deployment_group_name="bar",
            ec2_tag_filters=[{
                "key": "filterkey",
                "type": "KEY_AND_VALUE",
                "value": "filtervalue",
            }],
            service_role_arn=aws_iam_role["foo_role"]["arn"],
            trigger_configurations=[{
                "triggerEvents": ["DeploymentFailure"],
                "triggerName": "foo-trigger",
                "triggerTargetArn": "foo-topic-arn",
            }])
        ```
        ### Lambda Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_deployment_config = aws.codedeploy.DeploymentConfig("fooDeploymentConfig",
            compute_platform="Lambda",
            deployment_config_name="test-deployment-config",
            traffic_routing_config={
                "timeBasedLinear": {
                    "interval": 10,
                    "percentage": 10,
                },
                "type": "TimeBasedLinear",
            })
        foo_deployment_group = aws.codedeploy.DeploymentGroup("fooDeploymentGroup",
            alarm_configuration={
                "alarms": ["my-alarm-name"],
                "enabled": True,
            },
            app_name=aws_codedeploy_app["foo_app"]["name"],
            auto_rollback_configuration={
                "enabled": True,
                "events": ["DEPLOYMENT_STOP_ON_ALARM"],
            },
            deployment_config_name=foo_deployment_config.id,
            deployment_group_name="bar",
            service_role_arn=aws_iam_role["foo_role"]["arn"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_platform: The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
        :param pulumi.Input[str] deployment_config_name: The name of the deployment config.
        :param pulumi.Input[dict] minimum_healthy_hosts: A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
        :param pulumi.Input[dict] traffic_routing_config: A traffic_routing_config block. Traffic Routing Config is documented below.

        The **minimum_healthy_hosts** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type can either be `FLEET_PERCENT` or `HOST_COUNT`.
          * `value` (`pulumi.Input[float]`) - The value when the type is `FLEET_PERCENT` represents the minimum number of healthy instances as
            a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the
            deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
            When the type is `HOST_COUNT`, the value represents the minimum number of healthy instances as an absolute value.

        The **traffic_routing_config** object supports the following:

          * `timeBasedCanary` (`pulumi.Input[dict]`) - The time based canary configuration information. If `type` is `TimeBasedLinear`, use `time_based_linear` instead.
            * `interval` (`pulumi.Input[float]`) - The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
            * `percentage` (`pulumi.Input[float]`) - The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.

          * `timeBasedLinear` (`pulumi.Input[dict]`) - The time based linear configuration information. If `type` is `TimeBasedCanary`, use `time_based_canary` instead.
            * `interval` (`pulumi.Input[float]`) - The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
            * `percentage` (`pulumi.Input[float]`) - The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.

          * `type` (`pulumi.Input[str]`) - Type of traffic routing config. One of `TimeBasedCanary`, `TimeBasedLinear`, `AllAtOnce`.
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

            __props__['compute_platform'] = compute_platform
            if deployment_config_name is None:
                raise TypeError("Missing required property 'deployment_config_name'")
            __props__['deployment_config_name'] = deployment_config_name
            __props__['minimum_healthy_hosts'] = minimum_healthy_hosts
            __props__['traffic_routing_config'] = traffic_routing_config
            __props__['deployment_config_id'] = None
        super(DeploymentConfig, __self__).__init__(
            'aws:codedeploy/deploymentConfig:DeploymentConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, compute_platform=None, deployment_config_id=None, deployment_config_name=None, minimum_healthy_hosts=None, traffic_routing_config=None):
        """
        Get an existing DeploymentConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] compute_platform: The compute platform can be `Server`, `Lambda`, or `ECS`. Default is `Server`.
        :param pulumi.Input[str] deployment_config_id: The AWS Assigned deployment config id
        :param pulumi.Input[str] deployment_config_name: The name of the deployment config.
        :param pulumi.Input[dict] minimum_healthy_hosts: A minimum_healthy_hosts block. Required for `Server` compute platform. Minimum Healthy Hosts are documented below.
        :param pulumi.Input[dict] traffic_routing_config: A traffic_routing_config block. Traffic Routing Config is documented below.

        The **minimum_healthy_hosts** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type can either be `FLEET_PERCENT` or `HOST_COUNT`.
          * `value` (`pulumi.Input[float]`) - The value when the type is `FLEET_PERCENT` represents the minimum number of healthy instances as
            a percentage of the total number of instances in the deployment. If you specify FLEET_PERCENT, at the start of the
            deployment, AWS CodeDeploy converts the percentage to the equivalent number of instance and rounds up fractional instances.
            When the type is `HOST_COUNT`, the value represents the minimum number of healthy instances as an absolute value.

        The **traffic_routing_config** object supports the following:

          * `timeBasedCanary` (`pulumi.Input[dict]`) - The time based canary configuration information. If `type` is `TimeBasedLinear`, use `time_based_linear` instead.
            * `interval` (`pulumi.Input[float]`) - The number of minutes between the first and second traffic shifts of a `TimeBasedCanary` deployment.
            * `percentage` (`pulumi.Input[float]`) - The percentage of traffic to shift in the first increment of a `TimeBasedCanary` deployment.

          * `timeBasedLinear` (`pulumi.Input[dict]`) - The time based linear configuration information. If `type` is `TimeBasedCanary`, use `time_based_canary` instead.
            * `interval` (`pulumi.Input[float]`) - The number of minutes between each incremental traffic shift of a `TimeBasedLinear` deployment.
            * `percentage` (`pulumi.Input[float]`) - The percentage of traffic that is shifted at the start of each increment of a `TimeBasedLinear` deployment.

          * `type` (`pulumi.Input[str]`) - Type of traffic routing config. One of `TimeBasedCanary`, `TimeBasedLinear`, `AllAtOnce`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["compute_platform"] = compute_platform
        __props__["deployment_config_id"] = deployment_config_id
        __props__["deployment_config_name"] = deployment_config_name
        __props__["minimum_healthy_hosts"] = minimum_healthy_hosts
        __props__["traffic_routing_config"] = traffic_routing_config
        return DeploymentConfig(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
