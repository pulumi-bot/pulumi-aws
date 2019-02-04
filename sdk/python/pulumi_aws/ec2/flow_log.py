# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class FlowLog(pulumi.CustomResource):
    eni_id: pulumi.Output[str]
    """
    Elastic Network Interface ID to attach to
    """
    iam_role_arn: pulumi.Output[str]
    """
    The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
    """
    log_destination: pulumi.Output[str]
    """
    The ARN of the logging destination.
    """
    log_destination_type: pulumi.Output[str]
    """
    The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
    """
    log_group_name: pulumi.Output[str]
    """
    *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
    """
    subnet_id: pulumi.Output[str]
    """
    Subnet ID to attach to
    """
    traffic_type: pulumi.Output[str]
    """
    The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
    """
    vpc_id: pulumi.Output[str]
    """
    VPC ID to attach to
    """
    def __init__(__self__, resource_name, opts=None, eni_id=None, iam_role_arn=None, log_destination=None, log_destination_type=None, log_group_name=None, subnet_id=None, traffic_type=None, vpc_id=None, __name__=None, __opts__=None):
        """
        Provides a VPC/Subnet/ENI Flow Log to capture IP traffic for a specific network
        interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group or a S3 Bucket.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eni_id: Elastic Network Interface ID to attach to
        :param pulumi.Input[str] iam_role_arn: The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
        :param pulumi.Input[str] log_destination: The ARN of the logging destination.
        :param pulumi.Input[str] log_destination_type: The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
        :param pulumi.Input[str] log_group_name: *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
        :param pulumi.Input[str] subnet_id: Subnet ID to attach to
        :param pulumi.Input[str] traffic_type: The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
        :param pulumi.Input[str] vpc_id: VPC ID to attach to
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['eni_id'] = eni_id

        __props__['iam_role_arn'] = iam_role_arn

        __props__['log_destination'] = log_destination

        __props__['log_destination_type'] = log_destination_type

        __props__['log_group_name'] = log_group_name

        __props__['subnet_id'] = subnet_id

        if not traffic_type:
            raise TypeError('Missing required property traffic_type')
        __props__['traffic_type'] = traffic_type

        __props__['vpc_id'] = vpc_id

        super(FlowLog, __self__).__init__(
            'aws:ec2/flowLog:FlowLog',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

