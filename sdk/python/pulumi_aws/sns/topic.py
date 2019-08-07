# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Topic(pulumi.CustomResource):
    application_failure_feedback_role_arn: pulumi.Output[str]
    """
    IAM role for failure feedback
    """
    application_success_feedback_role_arn: pulumi.Output[str]
    """
    The IAM role permitted to receive success feedback for this topic
    """
    application_success_feedback_sample_rate: pulumi.Output[float]
    """
    Percentage of success to sample
    """
    arn: pulumi.Output[str]
    """
    The ARN of the SNS topic, as a more obvious property (clone of id)
    """
    delivery_policy: pulumi.Output[str]
    """
    The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
    """
    display_name: pulumi.Output[str]
    """
    The display name for the SNS topic
    """
    http_failure_feedback_role_arn: pulumi.Output[str]
    """
    IAM role for failure feedback
    """
    http_success_feedback_role_arn: pulumi.Output[str]
    """
    The IAM role permitted to receive success feedback for this topic
    """
    http_success_feedback_sample_rate: pulumi.Output[float]
    """
    Percentage of success to sample
    """
    kms_master_key_id: pulumi.Output[str]
    """
    The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
    """
    lambda_failure_feedback_role_arn: pulumi.Output[str]
    """
    IAM role for failure feedback
    """
    lambda_success_feedback_role_arn: pulumi.Output[str]
    """
    The IAM role permitted to receive success feedback for this topic
    """
    lambda_success_feedback_sample_rate: pulumi.Output[float]
    """
    Percentage of success to sample
    """
    name: pulumi.Output[str]
    """
    The friendly name for the SNS topic. By default generated by this provider.
    """
    name_prefix: pulumi.Output[str]
    """
    The friendly name for the SNS topic. Conflicts with `name`.
    """
    policy: pulumi.Output[str]
    """
    The fully-formed AWS policy as JSON.
    """
    sqs_failure_feedback_role_arn: pulumi.Output[str]
    """
    IAM role for failure feedback
    """
    sqs_success_feedback_role_arn: pulumi.Output[str]
    """
    The IAM role permitted to receive success feedback for this topic
    """
    sqs_success_feedback_sample_rate: pulumi.Output[float]
    """
    Percentage of success to sample
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    def __init__(__self__, resource_name, opts=None, application_failure_feedback_role_arn=None, application_success_feedback_role_arn=None, application_success_feedback_sample_rate=None, delivery_policy=None, display_name=None, http_failure_feedback_role_arn=None, http_success_feedback_role_arn=None, http_success_feedback_sample_rate=None, kms_master_key_id=None, lambda_failure_feedback_role_arn=None, lambda_success_feedback_role_arn=None, lambda_success_feedback_sample_rate=None, name=None, name_prefix=None, policy=None, sqs_failure_feedback_role_arn=None, sqs_success_feedback_role_arn=None, sqs_success_feedback_sample_rate=None, tags=None, __name__=None, __opts__=None):
        """
        Provides an SNS topic resource
        
        ## Message Delivery Status Arguments
        
        The `<endpoint>_success_feedback_role_arn` and `<endpoint>_failure_feedback_role_arn` arguments are used to give Amazon SNS write access to use CloudWatch Logs on your behalf. The `<endpoint>_success_feedback_sample_rate` argument is for specifying the sample rate percentage (0-100) of successfully delivered messages. After you configure the  `<endpoint>_failure_feedback_role_arn` argument, then all failed message deliveries generate CloudWatch Logs.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_failure_feedback_role_arn: IAM role for failure feedback
        :param pulumi.Input[str] application_success_feedback_role_arn: The IAM role permitted to receive success feedback for this topic
        :param pulumi.Input[float] application_success_feedback_sample_rate: Percentage of success to sample
        :param pulumi.Input[str] delivery_policy: The SNS delivery policy. More on [AWS documentation](https://docs.aws.amazon.com/sns/latest/dg/DeliveryPolicies.html)
        :param pulumi.Input[str] display_name: The display name for the SNS topic
        :param pulumi.Input[str] http_failure_feedback_role_arn: IAM role for failure feedback
        :param pulumi.Input[str] http_success_feedback_role_arn: The IAM role permitted to receive success feedback for this topic
        :param pulumi.Input[float] http_success_feedback_sample_rate: Percentage of success to sample
        :param pulumi.Input[str] kms_master_key_id: The ID of an AWS-managed customer master key (CMK) for Amazon SNS or a custom CMK. For more information, see [Key Terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms)
        :param pulumi.Input[str] lambda_failure_feedback_role_arn: IAM role for failure feedback
        :param pulumi.Input[str] lambda_success_feedback_role_arn: The IAM role permitted to receive success feedback for this topic
        :param pulumi.Input[float] lambda_success_feedback_sample_rate: Percentage of success to sample
        :param pulumi.Input[str] name: The friendly name for the SNS topic. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The friendly name for the SNS topic. Conflicts with `name`.
        :param pulumi.Input[str] policy: The fully-formed AWS policy as JSON.
        :param pulumi.Input[str] sqs_failure_feedback_role_arn: IAM role for failure feedback
        :param pulumi.Input[str] sqs_success_feedback_role_arn: The IAM role permitted to receive success feedback for this topic
        :param pulumi.Input[float] sqs_success_feedback_sample_rate: Percentage of success to sample
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sns_topic.html.markdown.
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

        __props__['application_failure_feedback_role_arn'] = application_failure_feedback_role_arn
        __props__['application_success_feedback_role_arn'] = application_success_feedback_role_arn
        __props__['application_success_feedback_sample_rate'] = application_success_feedback_sample_rate
        __props__['delivery_policy'] = delivery_policy
        __props__['display_name'] = display_name
        __props__['http_failure_feedback_role_arn'] = http_failure_feedback_role_arn
        __props__['http_success_feedback_role_arn'] = http_success_feedback_role_arn
        __props__['http_success_feedback_sample_rate'] = http_success_feedback_sample_rate
        __props__['kms_master_key_id'] = kms_master_key_id
        __props__['lambda_failure_feedback_role_arn'] = lambda_failure_feedback_role_arn
        __props__['lambda_success_feedback_role_arn'] = lambda_success_feedback_role_arn
        __props__['lambda_success_feedback_sample_rate'] = lambda_success_feedback_sample_rate
        __props__['name'] = name
        __props__['name_prefix'] = name_prefix
        __props__['policy'] = policy
        __props__['sqs_failure_feedback_role_arn'] = sqs_failure_feedback_role_arn
        __props__['sqs_success_feedback_role_arn'] = sqs_success_feedback_role_arn
        __props__['sqs_success_feedback_sample_rate'] = sqs_success_feedback_sample_rate
        __props__['tags'] = tags
        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Topic, __self__).__init__(
            'aws:sns/topic:Topic',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

