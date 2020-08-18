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

__all__ = ['Trail']


class Trail(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The Amazon Resource Name of the trail.
    """

    cloud_watch_logs_group_arn: pulumi.Output[Optional[str]] = pulumi.property("cloudWatchLogsGroupArn")
    """
    Specifies a log group name using an Amazon Resource Name (ARN),
    that represents the log group to which CloudTrail logs will be delivered.
    """

    cloud_watch_logs_role_arn: pulumi.Output[Optional[str]] = pulumi.property("cloudWatchLogsRoleArn")
    """
    Specifies the role for the CloudWatch Logs
    endpoint to assume to write to a user’s log group.
    """

    enable_log_file_validation: pulumi.Output[Optional[bool]] = pulumi.property("enableLogFileValidation")
    """
    Specifies whether log file integrity validation is enabled.
    Defaults to `false`.
    """

    enable_logging: pulumi.Output[Optional[bool]] = pulumi.property("enableLogging")
    """
    Enables logging for the trail. Defaults to `true`.
    Setting this to `false` will pause logging.
    """

    event_selectors: pulumi.Output[Optional[List['outputs.TrailEventSelector']]] = pulumi.property("eventSelectors")
    """
    Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
    """

    home_region: pulumi.Output[str] = pulumi.property("homeRegion")
    """
    The region in which the trail was created.
    """

    include_global_service_events: pulumi.Output[Optional[bool]] = pulumi.property("includeGlobalServiceEvents")
    """
    Specifies whether the trail is publishing events
    from global services such as IAM to the log files. Defaults to `true`.
    """

    is_multi_region_trail: pulumi.Output[Optional[bool]] = pulumi.property("isMultiRegionTrail")
    """
    Specifies whether the trail is created in the current
    region or in all regions. Defaults to `false`.
    """

    is_organization_trail: pulumi.Output[Optional[bool]] = pulumi.property("isOrganizationTrail")
    """
    Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
    """

    kms_key_id: pulumi.Output[Optional[str]] = pulumi.property("kmsKeyId")
    """
    Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    Specifies the name of the trail.
    """

    s3_bucket_name: pulumi.Output[str] = pulumi.property("s3BucketName")
    """
    Specifies the name of the S3 bucket designated for publishing log files.
    """

    s3_key_prefix: pulumi.Output[Optional[str]] = pulumi.property("s3KeyPrefix")
    """
    Specifies the S3 key prefix that follows
    the name of the bucket you have designated for log file delivery.
    """

    sns_topic_name: pulumi.Output[Optional[str]] = pulumi.property("snsTopicName")
    """
    Specifies the name of the Amazon SNS topic
    defined for notification of log file delivery.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A map of tags to assign to the trail
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloud_watch_logs_group_arn: Optional[pulumi.Input[str]] = None,
                 cloud_watch_logs_role_arn: Optional[pulumi.Input[str]] = None,
                 enable_log_file_validation: Optional[pulumi.Input[bool]] = None,
                 enable_logging: Optional[pulumi.Input[bool]] = None,
                 event_selectors: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TrailEventSelectorArgs']]]]] = None,
                 include_global_service_events: Optional[pulumi.Input[bool]] = None,
                 is_multi_region_trail: Optional[pulumi.Input[bool]] = None,
                 is_organization_trail: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 s3_bucket_name: Optional[pulumi.Input[str]] = None,
                 s3_key_prefix: Optional[pulumi.Input[str]] = None,
                 sns_topic_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudTrail resource.

        > *NOTE:* For a multi-region trail, this resource must be in the home region of the trail.

        > *NOTE:* For an organization trail, this resource must be in the master account of the organization.

        ## Example Usage
        ### Basic

        Enable CloudTrail to capture all compatible management events in region.
        For capturing events from services like IAM, `include_global_service_events` must be enabled.

        ```python
        import pulumi
        import pulumi_aws as aws

        current = aws.get_caller_identity()
        foo = aws.s3.Bucket("foo",
            force_destroy=True,
            policy=f\"\"\"{{
            "Version": "2012-10-17",
            "Statement": [
                {{
                    "Sid": "AWSCloudTrailAclCheck",
                    "Effect": "Allow",
                    "Principal": {{
                      "Service": "cloudtrail.amazonaws.com"
                    }},
                    "Action": "s3:GetBucketAcl",
                    "Resource": "arn:aws:s3:::tf-test-trail"
                }},
                {{
                    "Sid": "AWSCloudTrailWrite",
                    "Effect": "Allow",
                    "Principal": {{
                      "Service": "cloudtrail.amazonaws.com"
                    }},
                    "Action": "s3:PutObject",
                    "Resource": "arn:aws:s3:::tf-test-trail/prefix/AWSLogs/{current.account_id}/*",
                    "Condition": {{
                        "StringEquals": {{
                            "s3:x-amz-acl": "bucket-owner-full-control"
                        }}
                    }}
                }}
            ]
        }}
        \"\"\")
        foobar = aws.cloudtrail.Trail("foobar",
            s3_bucket_name=foo.id,
            s3_key_prefix="prefix",
            include_global_service_events=False)
        ```
        ### Data Event Logging

        CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the [CloudTrail API DataResource documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_DataResource.html).
        ### Logging All Lambda Function Invocations

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudtrail.Trail("example", event_selectors=[{
            "dataResources": [{
                "type": "AWS::Lambda::Function",
                "values": ["arn:aws:lambda"],
            }],
            "includeManagementEvents": True,
            "readWriteType": "All",
        }])
        ```
        ### Logging All S3 Bucket Object Events

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.cloudtrail.Trail("example", event_selectors=[{
            "dataResources": [{
                "type": "AWS::S3::Object",
                "values": ["arn:aws:s3:::"],
            }],
            "includeManagementEvents": True,
            "readWriteType": "All",
        }])
        ```
        ### Logging Individual S3 Bucket Events

        ```python
        import pulumi
        import pulumi_aws as aws

        important_bucket = aws.s3.get_bucket(bucket="important-bucket")
        example = aws.cloudtrail.Trail("example", event_selectors=[{
            "dataResources": [{
                "type": "AWS::S3::Object",
                "values": [f"{important_bucket.arn}/"],
            }],
            "includeManagementEvents": True,
            "readWriteType": "All",
        }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cloud_watch_logs_group_arn: Specifies a log group name using an Amazon Resource Name (ARN),
               that represents the log group to which CloudTrail logs will be delivered.
        :param pulumi.Input[str] cloud_watch_logs_role_arn: Specifies the role for the CloudWatch Logs
               endpoint to assume to write to a user’s log group.
        :param pulumi.Input[bool] enable_log_file_validation: Specifies whether log file integrity validation is enabled.
               Defaults to `false`.
        :param pulumi.Input[bool] enable_logging: Enables logging for the trail. Defaults to `true`.
               Setting this to `false` will pause logging.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TrailEventSelectorArgs']]]] event_selectors: Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
        :param pulumi.Input[bool] include_global_service_events: Specifies whether the trail is publishing events
               from global services such as IAM to the log files. Defaults to `true`.
        :param pulumi.Input[bool] is_multi_region_trail: Specifies whether the trail is created in the current
               region or in all regions. Defaults to `false`.
        :param pulumi.Input[bool] is_organization_trail: Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
        :param pulumi.Input[str] kms_key_id: Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
        :param pulumi.Input[str] name: Specifies the name of the trail.
        :param pulumi.Input[str] s3_bucket_name: Specifies the name of the S3 bucket designated for publishing log files.
        :param pulumi.Input[str] s3_key_prefix: Specifies the S3 key prefix that follows
               the name of the bucket you have designated for log file delivery.
        :param pulumi.Input[str] sns_topic_name: Specifies the name of the Amazon SNS topic
               defined for notification of log file delivery.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the trail
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
            if s3_bucket_name is None:
                raise TypeError("Missing required property 's3_bucket_name'")
            __props__['s3_bucket_name'] = s3_bucket_name
            __props__['s3_key_prefix'] = s3_key_prefix
            __props__['sns_topic_name'] = sns_topic_name
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['home_region'] = None
        super(Trail, __self__).__init__(
            'aws:cloudtrail/trail:Trail',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cloud_watch_logs_group_arn: Optional[pulumi.Input[str]] = None,
            cloud_watch_logs_role_arn: Optional[pulumi.Input[str]] = None,
            enable_log_file_validation: Optional[pulumi.Input[bool]] = None,
            enable_logging: Optional[pulumi.Input[bool]] = None,
            event_selectors: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TrailEventSelectorArgs']]]]] = None,
            home_region: Optional[pulumi.Input[str]] = None,
            include_global_service_events: Optional[pulumi.Input[bool]] = None,
            is_multi_region_trail: Optional[pulumi.Input[bool]] = None,
            is_organization_trail: Optional[pulumi.Input[bool]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            s3_bucket_name: Optional[pulumi.Input[str]] = None,
            s3_key_prefix: Optional[pulumi.Input[str]] = None,
            sns_topic_name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Trail':
        """
        Get an existing Trail resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name of the trail.
        :param pulumi.Input[str] cloud_watch_logs_group_arn: Specifies a log group name using an Amazon Resource Name (ARN),
               that represents the log group to which CloudTrail logs will be delivered.
        :param pulumi.Input[str] cloud_watch_logs_role_arn: Specifies the role for the CloudWatch Logs
               endpoint to assume to write to a user’s log group.
        :param pulumi.Input[bool] enable_log_file_validation: Specifies whether log file integrity validation is enabled.
               Defaults to `false`.
        :param pulumi.Input[bool] enable_logging: Enables logging for the trail. Defaults to `true`.
               Setting this to `false` will pause logging.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['TrailEventSelectorArgs']]]] event_selectors: Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these.
        :param pulumi.Input[str] home_region: The region in which the trail was created.
        :param pulumi.Input[bool] include_global_service_events: Specifies whether the trail is publishing events
               from global services such as IAM to the log files. Defaults to `true`.
        :param pulumi.Input[bool] is_multi_region_trail: Specifies whether the trail is created in the current
               region or in all regions. Defaults to `false`.
        :param pulumi.Input[bool] is_organization_trail: Specifies whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
        :param pulumi.Input[str] kms_key_id: Specifies the KMS key ARN to use to encrypt the logs delivered by CloudTrail.
        :param pulumi.Input[str] name: Specifies the name of the trail.
        :param pulumi.Input[str] s3_bucket_name: Specifies the name of the S3 bucket designated for publishing log files.
        :param pulumi.Input[str] s3_key_prefix: Specifies the S3 key prefix that follows
               the name of the bucket you have designated for log file delivery.
        :param pulumi.Input[str] sns_topic_name: Specifies the name of the Amazon SNS topic
               defined for notification of log file delivery.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the trail
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["cloud_watch_logs_group_arn"] = cloud_watch_logs_group_arn
        __props__["cloud_watch_logs_role_arn"] = cloud_watch_logs_role_arn
        __props__["enable_log_file_validation"] = enable_log_file_validation
        __props__["enable_logging"] = enable_logging
        __props__["event_selectors"] = event_selectors
        __props__["home_region"] = home_region
        __props__["include_global_service_events"] = include_global_service_events
        __props__["is_multi_region_trail"] = is_multi_region_trail
        __props__["is_organization_trail"] = is_organization_trail
        __props__["kms_key_id"] = kms_key_id
        __props__["name"] = name
        __props__["s3_bucket_name"] = s3_bucket_name
        __props__["s3_key_prefix"] = s3_key_prefix
        __props__["sns_topic_name"] = sns_topic_name
        __props__["tags"] = tags
        return Trail(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

