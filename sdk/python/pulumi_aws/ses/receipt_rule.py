# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ReceiptRule(pulumi.CustomResource):
    add_header_actions: pulumi.Output[list]
    """
    A list of Add Header Action blocks. Documented below.

      * `headerName` (`str`) - The name of the header to add
      * `headerValue` (`str`) - The value of the header to add
      * `position` (`float`) - The position of the action in the receipt rule
    """
    after: pulumi.Output[str]
    """
    The name of the rule to place this rule after
    """
    bounce_actions: pulumi.Output[list]
    """
    A list of Bounce Action blocks. Documented below.

      * `message` (`str`) - The message to send
      * `position` (`float`) - The position of the action in the receipt rule
      * `sender` (`str`) - The email address of the sender
      * `smtpReplyCode` (`str`) - The RFC 5321 SMTP reply code
      * `status_code` (`str`) - The RFC 3463 SMTP enhanced status code
      * `topic_arn` (`str`) - The ARN of an SNS topic to notify
    """
    enabled: pulumi.Output[bool]
    """
    If true, the rule will be enabled
    """
    lambda_actions: pulumi.Output[list]
    """
    A list of Lambda Action blocks. Documented below.

      * `function_arn` (`str`) - The ARN of the Lambda function to invoke
      * `invocationType` (`str`) - Event or RequestResponse
      * `position` (`float`) - The position of the action in the receipt rule
      * `topic_arn` (`str`) - The ARN of an SNS topic to notify
    """
    name: pulumi.Output[str]
    """
    The name of the rule
    """
    recipients: pulumi.Output[list]
    """
    A list of email addresses
    """
    rule_set_name: pulumi.Output[str]
    """
    The name of the rule set
    """
    s3_actions: pulumi.Output[list]
    """
    A list of S3 Action blocks. Documented below.

      * `bucket_name` (`str`) - The name of the S3 bucket
      * `kms_key_arn` (`str`) - The ARN of the KMS key
      * `objectKeyPrefix` (`str`) - The key prefix of the S3 bucket
      * `position` (`float`) - The position of the action in the receipt rule
      * `topic_arn` (`str`) - The ARN of an SNS topic to notify
    """
    scan_enabled: pulumi.Output[bool]
    """
    If true, incoming emails will be scanned for spam and viruses
    """
    sns_actions: pulumi.Output[list]
    """
    A list of SNS Action blocks. Documented below.

      * `position` (`float`) - The position of the action in the receipt rule
      * `topic_arn` (`str`) - The ARN of an SNS topic to notify
    """
    stop_actions: pulumi.Output[list]
    """
    A list of Stop Action blocks. Documented below.

      * `position` (`float`) - The position of the action in the receipt rule
      * `scope` (`str`) - The scope to apply
      * `topic_arn` (`str`) - The ARN of an SNS topic to notify
    """
    tls_policy: pulumi.Output[str]
    """
    Require or Optional
    """
    workmail_actions: pulumi.Output[list]
    """
    A list of WorkMail Action blocks. Documented below.

      * `organizationArn` (`str`) - The ARN of the WorkMail organization
      * `position` (`float`) - The position of the action in the receipt rule
      * `topic_arn` (`str`) - The ARN of an SNS topic to notify
    """
    def __init__(__self__, resource_name, opts=None, add_header_actions=None, after=None, bounce_actions=None, enabled=None, lambda_actions=None, name=None, recipients=None, rule_set_name=None, s3_actions=None, scan_enabled=None, sns_actions=None, stop_actions=None, tls_policy=None, workmail_actions=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an SES receipt rule resource

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        # Add a header to the email and store it in S3
        store = aws.ses.ReceiptRule("store",
            add_header_actions=[{
                "headerName": "Custom-Header",
                "headerValue": "Added by SES",
                "position": 1,
            }],
            enabled=True,
            recipients=["karen@example.com"],
            rule_set_name="default-rule-set",
            s3_actions=[{
                "bucketName": "emails",
                "position": 2,
            }],
            scan_enabled=True)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] add_header_actions: A list of Add Header Action blocks. Documented below.
        :param pulumi.Input[str] after: The name of the rule to place this rule after
        :param pulumi.Input[list] bounce_actions: A list of Bounce Action blocks. Documented below.
        :param pulumi.Input[bool] enabled: If true, the rule will be enabled
        :param pulumi.Input[list] lambda_actions: A list of Lambda Action blocks. Documented below.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[list] recipients: A list of email addresses
        :param pulumi.Input[str] rule_set_name: The name of the rule set
        :param pulumi.Input[list] s3_actions: A list of S3 Action blocks. Documented below.
        :param pulumi.Input[bool] scan_enabled: If true, incoming emails will be scanned for spam and viruses
        :param pulumi.Input[list] sns_actions: A list of SNS Action blocks. Documented below.
        :param pulumi.Input[list] stop_actions: A list of Stop Action blocks. Documented below.
        :param pulumi.Input[str] tls_policy: Require or Optional
        :param pulumi.Input[list] workmail_actions: A list of WorkMail Action blocks. Documented below.

        The **add_header_actions** object supports the following:

          * `headerName` (`pulumi.Input[str]`) - The name of the header to add
          * `headerValue` (`pulumi.Input[str]`) - The value of the header to add
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule

        The **bounce_actions** object supports the following:

          * `message` (`pulumi.Input[str]`) - The message to send
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `sender` (`pulumi.Input[str]`) - The email address of the sender
          * `smtpReplyCode` (`pulumi.Input[str]`) - The RFC 5321 SMTP reply code
          * `status_code` (`pulumi.Input[str]`) - The RFC 3463 SMTP enhanced status code
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **lambda_actions** object supports the following:

          * `function_arn` (`pulumi.Input[str]`) - The ARN of the Lambda function to invoke
          * `invocationType` (`pulumi.Input[str]`) - Event or RequestResponse
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **s3_actions** object supports the following:

          * `bucket_name` (`pulumi.Input[str]`) - The name of the S3 bucket
          * `kms_key_arn` (`pulumi.Input[str]`) - The ARN of the KMS key
          * `objectKeyPrefix` (`pulumi.Input[str]`) - The key prefix of the S3 bucket
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **sns_actions** object supports the following:

          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **stop_actions** object supports the following:

          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `scope` (`pulumi.Input[str]`) - The scope to apply
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **workmail_actions** object supports the following:

          * `organizationArn` (`pulumi.Input[str]`) - The ARN of the WorkMail organization
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify
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

            __props__['add_header_actions'] = add_header_actions
            __props__['after'] = after
            __props__['bounce_actions'] = bounce_actions
            __props__['enabled'] = enabled
            __props__['lambda_actions'] = lambda_actions
            __props__['name'] = name
            __props__['recipients'] = recipients
            if rule_set_name is None:
                raise TypeError("Missing required property 'rule_set_name'")
            __props__['rule_set_name'] = rule_set_name
            __props__['s3_actions'] = s3_actions
            __props__['scan_enabled'] = scan_enabled
            __props__['sns_actions'] = sns_actions
            __props__['stop_actions'] = stop_actions
            __props__['tls_policy'] = tls_policy
            __props__['workmail_actions'] = workmail_actions
        super(ReceiptRule, __self__).__init__(
            'aws:ses/receiptRule:ReceiptRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_header_actions=None, after=None, bounce_actions=None, enabled=None, lambda_actions=None, name=None, recipients=None, rule_set_name=None, s3_actions=None, scan_enabled=None, sns_actions=None, stop_actions=None, tls_policy=None, workmail_actions=None):
        """
        Get an existing ReceiptRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] add_header_actions: A list of Add Header Action blocks. Documented below.
        :param pulumi.Input[str] after: The name of the rule to place this rule after
        :param pulumi.Input[list] bounce_actions: A list of Bounce Action blocks. Documented below.
        :param pulumi.Input[bool] enabled: If true, the rule will be enabled
        :param pulumi.Input[list] lambda_actions: A list of Lambda Action blocks. Documented below.
        :param pulumi.Input[str] name: The name of the rule
        :param pulumi.Input[list] recipients: A list of email addresses
        :param pulumi.Input[str] rule_set_name: The name of the rule set
        :param pulumi.Input[list] s3_actions: A list of S3 Action blocks. Documented below.
        :param pulumi.Input[bool] scan_enabled: If true, incoming emails will be scanned for spam and viruses
        :param pulumi.Input[list] sns_actions: A list of SNS Action blocks. Documented below.
        :param pulumi.Input[list] stop_actions: A list of Stop Action blocks. Documented below.
        :param pulumi.Input[str] tls_policy: Require or Optional
        :param pulumi.Input[list] workmail_actions: A list of WorkMail Action blocks. Documented below.

        The **add_header_actions** object supports the following:

          * `headerName` (`pulumi.Input[str]`) - The name of the header to add
          * `headerValue` (`pulumi.Input[str]`) - The value of the header to add
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule

        The **bounce_actions** object supports the following:

          * `message` (`pulumi.Input[str]`) - The message to send
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `sender` (`pulumi.Input[str]`) - The email address of the sender
          * `smtpReplyCode` (`pulumi.Input[str]`) - The RFC 5321 SMTP reply code
          * `status_code` (`pulumi.Input[str]`) - The RFC 3463 SMTP enhanced status code
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **lambda_actions** object supports the following:

          * `function_arn` (`pulumi.Input[str]`) - The ARN of the Lambda function to invoke
          * `invocationType` (`pulumi.Input[str]`) - Event or RequestResponse
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **s3_actions** object supports the following:

          * `bucket_name` (`pulumi.Input[str]`) - The name of the S3 bucket
          * `kms_key_arn` (`pulumi.Input[str]`) - The ARN of the KMS key
          * `objectKeyPrefix` (`pulumi.Input[str]`) - The key prefix of the S3 bucket
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **sns_actions** object supports the following:

          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **stop_actions** object supports the following:

          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `scope` (`pulumi.Input[str]`) - The scope to apply
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify

        The **workmail_actions** object supports the following:

          * `organizationArn` (`pulumi.Input[str]`) - The ARN of the WorkMail organization
          * `position` (`pulumi.Input[float]`) - The position of the action in the receipt rule
          * `topic_arn` (`pulumi.Input[str]`) - The ARN of an SNS topic to notify
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_header_actions"] = add_header_actions
        __props__["after"] = after
        __props__["bounce_actions"] = bounce_actions
        __props__["enabled"] = enabled
        __props__["lambda_actions"] = lambda_actions
        __props__["name"] = name
        __props__["recipients"] = recipients
        __props__["rule_set_name"] = rule_set_name
        __props__["s3_actions"] = s3_actions
        __props__["scan_enabled"] = scan_enabled
        __props__["sns_actions"] = sns_actions
        __props__["stop_actions"] = stop_actions
        __props__["tls_policy"] = tls_policy
        __props__["workmail_actions"] = workmail_actions
        return ReceiptRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

