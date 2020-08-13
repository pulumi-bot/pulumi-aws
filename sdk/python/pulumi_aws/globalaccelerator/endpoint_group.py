# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class EndpointGroup(pulumi.CustomResource):
    endpoint_configurations: pulumi.Output[list]
    """
    The list of endpoint objects. Fields documented below.

      * `endpoint_id` (`str`) - An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
      * `weight` (`float`) - The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify.
    """
    endpoint_group_region: pulumi.Output[str]
    """
    The name of the AWS Region where the endpoint group is located.
    """
    health_check_interval_seconds: pulumi.Output[float]
    """
    The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
    """
    health_check_path: pulumi.Output[str]
    """
    If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
    """
    health_check_port: pulumi.Output[float]
    """
    The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
    """
    health_check_protocol: pulumi.Output[str]
    """
    The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
    """
    listener_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the listener.
    """
    threshold_count: pulumi.Output[float]
    """
    The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
    """
    traffic_dial_percentage: pulumi.Output[float]
    """
    The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.
    """
    def __init__(__self__, resource_name, opts=None, endpoint_configurations=None, endpoint_group_region=None, health_check_interval_seconds=None, health_check_path=None, health_check_port=None, health_check_protocol=None, listener_arn=None, threshold_count=None, traffic_dial_percentage=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Global Accelerator endpoint group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.globalaccelerator.EndpointGroup("example",
            listener_arn=aws_globalaccelerator_listener["example"]["id"],
            endpoint_configurations=[{
                "endpoint_id": aws_lb["example"]["arn"],
                "weight": 100,
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] endpoint_configurations: The list of endpoint objects. Fields documented below.
        :param pulumi.Input[str] endpoint_group_region: The name of the AWS Region where the endpoint group is located.
        :param pulumi.Input[float] health_check_interval_seconds: The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        :param pulumi.Input[str] health_check_path: If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
        :param pulumi.Input[float] health_check_port: The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
        :param pulumi.Input[str] health_check_protocol: The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        :param pulumi.Input[str] listener_arn: The Amazon Resource Name (ARN) of the listener.
        :param pulumi.Input[float] threshold_count: The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        :param pulumi.Input[float] traffic_dial_percentage: The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.

        The **endpoint_configurations** object supports the following:

          * `endpoint_id` (`pulumi.Input[str]`) - An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
          * `weight` (`pulumi.Input[float]`) - The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify.
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

            __props__['endpointConfigurations'] = endpoint_configurations
            __props__['endpointGroupRegion'] = endpoint_group_region
            __props__['healthCheckIntervalSeconds'] = health_check_interval_seconds
            __props__['healthCheckPath'] = health_check_path
            __props__['healthCheckPort'] = health_check_port
            __props__['healthCheckProtocol'] = health_check_protocol
            if listener_arn is None:
                raise TypeError("Missing required property 'listener_arn'")
            __props__['listenerArn'] = listener_arn
            __props__['thresholdCount'] = threshold_count
            __props__['trafficDialPercentage'] = traffic_dial_percentage
        super(EndpointGroup, __self__).__init__(
            'aws:globalaccelerator/endpointGroup:EndpointGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, endpoint_configurations=None, endpoint_group_region=None, health_check_interval_seconds=None, health_check_path=None, health_check_port=None, health_check_protocol=None, listener_arn=None, threshold_count=None, traffic_dial_percentage=None):
        """
        Get an existing EndpointGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] endpoint_configurations: The list of endpoint objects. Fields documented below.
        :param pulumi.Input[str] endpoint_group_region: The name of the AWS Region where the endpoint group is located.
        :param pulumi.Input[float] health_check_interval_seconds: The time—10 seconds or 30 seconds—between each health check for an endpoint. The default value is 30.
        :param pulumi.Input[str] health_check_path: If the protocol is HTTP/S, then this specifies the path that is the destination for health check targets. The default value is slash (/).
        :param pulumi.Input[float] health_check_port: The port that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default port is the listener port that this endpoint group is associated with. If listener port is a list of ports, Global Accelerator uses the first port in the list.
        :param pulumi.Input[str] health_check_protocol: The protocol that AWS Global Accelerator uses to check the health of endpoints that are part of this endpoint group. The default value is TCP.
        :param pulumi.Input[str] listener_arn: The Amazon Resource Name (ARN) of the listener.
        :param pulumi.Input[float] threshold_count: The number of consecutive health checks required to set the state of a healthy endpoint to unhealthy, or to set an unhealthy endpoint to healthy. The default value is 3.
        :param pulumi.Input[float] traffic_dial_percentage: The percentage of traffic to send to an AWS Region. Additional traffic is distributed to other endpoint groups for this listener. The default value is 100.

        The **endpoint_configurations** object supports the following:

          * `endpoint_id` (`pulumi.Input[str]`) - An ID for the endpoint. If the endpoint is a Network Load Balancer or Application Load Balancer, this is the Amazon Resource Name (ARN) of the resource. If the endpoint is an Elastic IP address, this is the Elastic IP address allocation ID.
          * `weight` (`pulumi.Input[float]`) - The weight associated with the endpoint. When you add weights to endpoints, you configure AWS Global Accelerator to route traffic based on proportions that you specify.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["endpoint_configurations"] = endpoint_configurations
        __props__["endpoint_group_region"] = endpoint_group_region
        __props__["health_check_interval_seconds"] = health_check_interval_seconds
        __props__["health_check_path"] = health_check_path
        __props__["health_check_port"] = health_check_port
        __props__["health_check_protocol"] = health_check_protocol
        __props__["listener_arn"] = listener_arn
        __props__["threshold_count"] = threshold_count
        __props__["traffic_dial_percentage"] = traffic_dial_percentage
        return EndpointGroup(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
