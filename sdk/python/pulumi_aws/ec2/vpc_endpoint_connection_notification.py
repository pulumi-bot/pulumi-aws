# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

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
    def __init__(__self__, resource_name, opts=None, connection_events=None, connection_notification_arn=None, vpc_endpoint_id=None, vpc_endpoint_service_id=None, __name__=None, __opts__=None):
        """
        Provides a VPC Endpoint connection notification resource.
        Connection notifications notify subscribers of VPC Endpoint events.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] connection_events: One or more endpoint [events](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVpcEndpointConnectionNotification.html#API_CreateVpcEndpointConnectionNotification_RequestParameters) for which to receive notifications.
        :param pulumi.Input[str] connection_notification_arn: The ARN of the SNS topic for the notifications.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC Endpoint to receive notifications for.
        :param pulumi.Input[str] vpc_endpoint_service_id: The ID of the VPC Endpoint Service to receive notifications for.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_endpoint_connection_notification.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if connection_events is None:
            raise TypeError("Missing required property 'connection_events'")
        __props__['connection_events'] = connection_events
        if connection_notification_arn is None:
            raise TypeError("Missing required property 'connection_notification_arn'")
        __props__['connection_notification_arn'] = connection_notification_arn
        __props__['vpc_endpoint_id'] = vpc_endpoint_id
        __props__['vpc_endpoint_service_id'] = vpc_endpoint_service_id
        __props__['notification_type'] = None
        __props__['state'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(VpcEndpointConnectionNotification, __self__).__init__(
            'aws:ec2/vpcEndpointConnectionNotification:VpcEndpointConnectionNotification',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

