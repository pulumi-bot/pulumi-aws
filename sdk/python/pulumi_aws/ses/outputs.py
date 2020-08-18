# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'EventDestinationCloudwatchDestination',
    'EventDestinationKinesisDestination',
    'EventDestinationSnsDestination',
    'ReceiptRuleAddHeaderAction',
    'ReceiptRuleBounceAction',
    'ReceiptRuleLambdaAction',
    'ReceiptRuleS3Action',
    'ReceiptRuleSnsAction',
    'ReceiptRuleStopAction',
    'ReceiptRuleWorkmailAction',
]

@pulumi.output_type
class EventDestinationCloudwatchDestination(dict):
    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> str:
        """
        The default value for the event
        """
        ...

    @property
    @pulumi.getter(name="dimensionName")
    def dimension_name(self) -> str:
        """
        The name for the dimension
        """
        ...

    @property
    @pulumi.getter(name="valueSource")
    def value_source(self) -> str:
        """
        The source for the value. It can be either `"messageTag"` or `"emailHeader"`
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EventDestinationKinesisDestination(dict):
    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> str:
        """
        The ARN of the role that has permissions to access the Kinesis Stream
        """
        ...

    @property
    @pulumi.getter(name="streamArn")
    def stream_arn(self) -> str:
        """
        The ARN of the Kinesis Stream
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class EventDestinationSnsDestination(dict):
    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> str:
        """
        The ARN of the SNS topic
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReceiptRuleAddHeaderAction(dict):
    @property
    @pulumi.getter(name="headerName")
    def header_name(self) -> str:
        """
        The name of the header to add
        """
        ...

    @property
    @pulumi.getter(name="headerValue")
    def header_value(self) -> str:
        """
        The value of the header to add
        """
        ...

    @property
    @pulumi.getter
    def position(self) -> float:
        """
        The position of the action in the receipt rule
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReceiptRuleBounceAction(dict):
    @property
    @pulumi.getter
    def message(self) -> str:
        """
        The message to send
        """
        ...

    @property
    @pulumi.getter
    def position(self) -> float:
        """
        The position of the action in the receipt rule
        """
        ...

    @property
    @pulumi.getter
    def sender(self) -> str:
        """
        The email address of the sender
        """
        ...

    @property
    @pulumi.getter(name="smtpReplyCode")
    def smtp_reply_code(self) -> str:
        """
        The RFC 5321 SMTP reply code
        """
        ...

    @property
    @pulumi.getter(name="statusCode")
    def status_code(self) -> Optional[str]:
        """
        The RFC 3463 SMTP enhanced status code
        """
        ...

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[str]:
        """
        The ARN of an SNS topic to notify
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReceiptRuleLambdaAction(dict):
    @property
    @pulumi.getter(name="functionArn")
    def function_arn(self) -> str:
        """
        The ARN of the Lambda function to invoke
        """
        ...

    @property
    @pulumi.getter(name="invocationType")
    def invocation_type(self) -> Optional[str]:
        """
        Event or RequestResponse
        """
        ...

    @property
    @pulumi.getter
    def position(self) -> float:
        """
        The position of the action in the receipt rule
        """
        ...

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[str]:
        """
        The ARN of an SNS topic to notify
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReceiptRuleS3Action(dict):
    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> str:
        """
        The name of the S3 bucket
        """
        ...

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> Optional[str]:
        """
        The ARN of the KMS key
        """
        ...

    @property
    @pulumi.getter(name="objectKeyPrefix")
    def object_key_prefix(self) -> Optional[str]:
        """
        The key prefix of the S3 bucket
        """
        ...

    @property
    @pulumi.getter
    def position(self) -> float:
        """
        The position of the action in the receipt rule
        """
        ...

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[str]:
        """
        The ARN of an SNS topic to notify
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReceiptRuleSnsAction(dict):
    @property
    @pulumi.getter
    def position(self) -> float:
        """
        The position of the action in the receipt rule
        """
        ...

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> str:
        """
        The ARN of an SNS topic to notify
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReceiptRuleStopAction(dict):
    @property
    @pulumi.getter
    def position(self) -> float:
        """
        The position of the action in the receipt rule
        """
        ...

    @property
    @pulumi.getter
    def scope(self) -> str:
        """
        The scope to apply
        """
        ...

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[str]:
        """
        The ARN of an SNS topic to notify
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ReceiptRuleWorkmailAction(dict):
    @property
    @pulumi.getter(name="organizationArn")
    def organization_arn(self) -> str:
        """
        The ARN of the WorkMail organization
        """
        ...

    @property
    @pulumi.getter
    def position(self) -> float:
        """
        The position of the action in the receipt rule
        """
        ...

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> Optional[str]:
        """
        The ARN of an SNS topic to notify
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


