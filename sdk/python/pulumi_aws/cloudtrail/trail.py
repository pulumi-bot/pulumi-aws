# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Trail(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name of the trail.
    """
    cloud_watch_logs_group_arn: pulumi.Output[str]
    """
    Specifies a log group name using an Amazon Resource Name (ARN),
    that represents the log group to which CloudTrail logs will be delivered.
    """
    cloud_watch_logs_role_arn: pulumi.Output[str]
    """
    Specifies the role for the CloudWatch Logs
    endpoint to assume to write to a user’s log group.
    """
    enable_log_file_validation: pulumi.Output[bool]
    """
    Specifies whether log file integrity validation is enabled.
    Defaults to `false`.
    """
    enable_logging: pulumi.Output[bool]
    """
    Enables logging for the trail. Defaults to `true`.
    Setting this to `false` will pause logging.
    """
    event_selectors: pulumi.Output[list]
    """
    Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
    """
    home_region: pulumi.Output[str]
    """
    The region in which the trail was created.
    """
    include_global_service_events: pulumi.Output[bool]
    """
    Specifies whether the trail is publishing events
    from global services such as IAM to the log files. Defaults to `true`.
    """
    is_multi_region_trail: pulumi.Output[bool]
    """
    Specifies whether the trail is created in the current
    region or in all regions. Defaults to `false`.
    """
    is_organization_trail: pulumi.Output[bool]
    """
    Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
    """
    kms_key_id: pulumi.Output[str]
    """
    Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the trail.
    """
    s3_bucket_name: pulumi.Output[str]
    """
    Specifies the name of the S3 bucket designated for publishing log files.
    """
    s3_key_prefix: pulumi.Output[str]
    """
    Specifies the S3 key prefix that precedes
    the name of the bucket you have designated for log file delivery.
    """
    sns_topic_name: pulumi.Output[str]
    """
    Specifies the name of the Amazon SNS topic
    defined for notification of log file delivery.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the trail
    """
    def __init__(__self__, __name__, __opts__=None, cloud_watch_logs_group_arn=None, cloud_watch_logs_role_arn=None, enable_log_file_validation=None, enable_logging=None, event_selectors=None, include_global_service_events=None, is_multi_region_trail=None, is_organization_trail=None, kms_key_id=None, name=None, s3_bucket_name=None, s3_key_prefix=None, sns_topic_name=None, tags=None):
        """
        Provides a CloudTrail resource.
        
        > *NOTE:* For a multi-region trail, this resource must be in the home region of the trail.
        
        > *NOTE:* For an organization trail, this resource must be in the master account of the organization.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] cloud_watch_logs_group_arn: Specifies a log group name using an Amazon Resource Name (ARN),
               that represents the log group to which CloudTrail logs will be delivered.
        :param pulumi.Input[str] cloud_watch_logs_role_arn: Specifies the role for the CloudWatch Logs
               endpoint to assume to write to a user’s log group.
        :param pulumi.Input[bool] enable_log_file_validation: Specifies whether log file integrity validation is enabled.
               Defaults to `false`.
        :param pulumi.Input[bool] enable_logging: Enables logging for the trail. Defaults to `true`.
               Setting this to `false` will pause logging.
        :param pulumi.Input[list] event_selectors: Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
        :param pulumi.Input[bool] include_global_service_events: Specifies whether the trail is publishing events
               from global services such as IAM to the log files. Defaults to `true`.
        :param pulumi.Input[bool] is_multi_region_trail: Specifies whether the trail is created in the current
               region or in all regions. Defaults to `false`.
        :param pulumi.Input[bool] is_organization_trail: Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
        :param pulumi.Input[str] kms_key_id: Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
        :param pulumi.Input[str] name: Specifies the name of the trail.
        :param pulumi.Input[str] s3_bucket_name: Specifies the name of the S3 bucket designated for publishing log files.
        :param pulumi.Input[str] s3_key_prefix: Specifies the S3 key prefix that precedes
               the name of the bucket you have designated for log file delivery.
        :param pulumi.Input[str] sns_topic_name: Specifies the name of the Amazon SNS topic
               defined for notification of log file delivery.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the trail
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['cloud_watch_logs_group_arn'] = cloud_watch_logs_group_arn

        __props__['cloud_watch_logs_role_arn'] = cloud_watch_logs_role_arn

        __props__['enable_log_file_validation'] = enable_log_file_validation

        __props__['enable_logging'] = enable_logging

        __props__['event_selectors'] = event_selectors

        __props__['include_global_service_events'] = include_global_service_events

        __props__['is_multi_region_trail'] = is_multi_region_trail

        __props__['is_organization_trail'] = is_organization_trail

        __props__['kms_key_id'] = kms_key_id

        __props__['name'] = name

        if not s3_bucket_name:
            raise TypeError('Missing required property s3_bucket_name')
        __props__['s3_bucket_name'] = s3_bucket_name

        __props__['s3_key_prefix'] = s3_key_prefix

        __props__['sns_topic_name'] = sns_topic_name

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['home_region'] = None

        super(Trail, __self__).__init__(
            'aws:cloudtrail/trail:Trail',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

