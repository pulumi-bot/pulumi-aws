# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class VpcEndpointConnectionNotification(pulumi.CustomResource):
    connection_events: pulumi.Output[list]
    """
    One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
    """
    connection_notification_arn: pulumi.Output[str]
    """
    The ARN of the SNS topic for the notifications.
    """
    notification_type: pulumi.Output[str]
    """
    The type of notification.
    """
    state: pulumi.Output[str]
    """
    The state of the notification.
    """
    vpc_endpoint_id: pulumi.Output[str]
    """
    The ID of the VPC Endpoint to receive notifications for.
    """
    vpc_endpoint_service_id: pulumi.Output[str]
    """
    The ID of the VPC Endpoint Service to receive notifications for.
    """
    def __init__(__self__, resource_name, opts=None, connection_events=None, connection_notification_arn=None, vpc_endpoint_id=None, vpc_endpoint_service_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a VPC Endpoint connection notification resource.
        Connection notifications notify subscribers of VPC Endpoint events.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        topic = aws.sns.Topic("topic", policy=\"\"\"{
            "Version":"2012-10-17",
            "Statement":[{
                "Effect": "Allow",
                "Principal": {
                    "Service": "vpce.amazonaws.com"
                },
                "Action": "SNS:Publish",
                "Resource": "arn:aws:sns:*:*:vpce-notification-topic"
            }]
        }
        \"\"\")
        foo_vpc_endpoint_service = aws.ec2.VpcEndpointService("fooVpcEndpointService",
            acceptance_required=False,
            network_load_balancer_arns=[aws_lb["test"]["arn"]])
        foo_vpc_endpoint_connection_notification = aws.ec2.VpcEndpointConnectionNotification("fooVpcEndpointConnectionNotification",
            vpc_endpoint_service_id=foo_vpc_endpoint_service.id,
            connection_notification_arn=topic.arn,
            connection_events=[
                "Accept",
                "Reject",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] connection_events: One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        :param pulumi.Input[str] connection_notification_arn: The ARN of the SNS topic for the notifications.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC Endpoint to receive notifications for.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC Endpoint Service to receive notifications for.
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

            if connection_events is None:
                raise TypeError("Missing required property 'connection_events'")
            __props__['connectionEvents'] = connection_events
            if connection_notification_arn is None:
                raise TypeError("Missing required property 'connection_notification_arn'")
            __props__['connectionNotificationArn'] = connection_notification_arn
            __props__['vpcEndpointId'] = vpc_endpoint_id
            __props__['vpcEndpointServiceId'] = vpc_endpoint_service_id
            __props__['notification_type'] = None
            __props__['state'] = None
        super(VpcEndpointConnectionNotification, __self__).__init__(
            'aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, connection_events=None, connection_notification_arn=None, notification_type=None, state=None, vpc_endpoint_id=None, vpc_endpoint_service_id=None):
        """
        Get an existing VpcEndpointConnectionNotification resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] connection_events: One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        :param pulumi.Input[str] connection_notification_arn: The ARN of the SNS topic for the notifications.
        :param pulumi.Input[str] notification_type: The type of notification.
        :param pulumi.Input[str] state: The state of the notification.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC Endpoint to receive notifications for.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC Endpoint Service to receive notifications for.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["connection_events"] = connection_events
        __props__["connection_notification_arn"] = connection_notification_arn
        __props__["notification_type"] = notification_type
        __props__["state"] = state
        __props__["vpc_endpoint_id"] = vpc_endpoint_id
        __props__["vpc_endpoint_service_id"] = vpc_endpoint_service_id
        return VpcEndpointConnectionNotification(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
