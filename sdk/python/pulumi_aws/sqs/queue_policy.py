# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['QueuePolicy']


class QueuePolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 queue_url: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Allows you to set a policy of an SQS Queue
        while referencing ARN of the queue within the policy.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        queue = aws.sqs.Queue("queue")
        test = aws.sqs.QueuePolicy("test",
            queue_url=queue.id,
            policy=queue.arn.apply(lambda arn: f\"\"\"{{
          "Version": "2012-10-17",
          "Id": "sqspolicy",
          "Statement": [
            {{
              "Sid": "First",
              "Effect": "Allow",
              "Principal": "*",
              "Action": "sqs:SendMessage",
              "Resource": "{arn}",
              "Condition": {{
                "ArnEquals": {{
                  "aws:SourceArn": "{aws_sns_topic["example"]["arn"]}"
                }}
              }}
            }}
          ]
        }}
        \"\"\"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: The JSON policy for the SQS queue.
        :param pulumi.Input[str] queue_url: The URL of the SQS Queue to which to attach the policy
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

            if policy is None:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
            if queue_url is None:
                raise TypeError("Missing required property 'queue_url'")
            __props__['queue_url'] = queue_url
        super(QueuePolicy, __self__).__init__(
            'aws:sqs/queuePolicy:QueuePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            policy: Optional[pulumi.Input[str]] = None,
            queue_url: Optional[pulumi.Input[str]] = None) -> 'QueuePolicy':
        """
        Get an existing QueuePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy: The JSON policy for the SQS queue.
        :param pulumi.Input[str] queue_url: The URL of the SQS Queue to which to attach the policy
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["policy"] = policy
        __props__["queue_url"] = queue_url
        return QueuePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def policy(self) -> str:
        """
        The JSON policy for the SQS queue.
        """
        ...

    @property
    @pulumi.getter(name="queueUrl")
    def queue_url(self) -> str:
        """
        The URL of the SQS Queue to which to attach the policy
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

