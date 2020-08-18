# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['EventRule']


class EventRule(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The Amazon Resource Name (ARN) of the rule.
    """

    description: pulumi.Output[Optional[str]] = pulumi.property("description")
    """
    The description of the rule.
    """

    event_pattern: pulumi.Output[Optional[str]] = pulumi.property("eventPattern")
    """
    Event pattern
    described a JSON object.
    See full documentation of [CloudWatch Events and Event Patterns](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CloudWatchEventsandEventPatterns.html) for details.
    """

    is_enabled: pulumi.Output[Optional[bool]] = pulumi.property("isEnabled")
    """
    Whether the rule should be enabled (defaults to `true`).
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The rule's name. By default generated by this provider.
    """

    name_prefix: pulumi.Output[Optional[str]] = pulumi.property("namePrefix")
    """
    The rule's name. Conflicts with `name`.
    """

    role_arn: pulumi.Output[Optional[str]] = pulumi.property("roleArn")
    """
    The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
    """

    schedule_expression: pulumi.Output[Optional[str]] = pulumi.property("scheduleExpression")
    """
    The scheduling expression.
    For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A map of tags to assign to the resource.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 event_pattern: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 name_prefix: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 schedule_expression: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Event Rule resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        console = aws.cloudwatch.EventRule("console",
            description="Capture each AWS Console Sign In",
            event_pattern=\"\"\"{
          "detail-type": [
            "AWS Console Sign In via CloudTrail"
          ]
        }
        \"\"\")
        aws_logins = aws.sns.Topic("awsLogins")
        sns = aws.cloudwatch.EventTarget("sns",
            rule=console.name,
            arn=aws_logins.arn)
        sns_topic_policy = aws_logins.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[{
            "effect": "Allow",
            "actions": ["SNS:Publish"],
            "principals": [{
                "type": "Service",
                "identifiers": ["events.amazonaws.com"],
            }],
            "resources": [arn],
        }]))
        default = aws.sns.TopicPolicy("default",
            arn=aws_logins.arn,
            policy=sns_topic_policy.json)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] event_pattern: Event pattern
               described a JSON object.
               See full documentation of [CloudWatch Events and Event Patterns](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CloudWatchEventsandEventPatterns.html) for details.
        :param pulumi.Input[bool] is_enabled: Whether the rule should be enabled (defaults to `true`).
        :param pulumi.Input[str] name: The rule's name. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The rule's name. Conflicts with `name`.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        :param pulumi.Input[str] schedule_expression: The scheduling expression.
               For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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

            __props__['description'] = description
            __props__['event_pattern'] = event_pattern
            __props__['is_enabled'] = is_enabled
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['role_arn'] = role_arn
            __props__['schedule_expression'] = schedule_expression
            __props__['tags'] = tags
            __props__['arn'] = None
        super(EventRule, __self__).__init__(
            'aws:cloudwatch/eventRule:EventRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            event_pattern: Optional[pulumi.Input[str]] = None,
            is_enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            name_prefix: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            schedule_expression: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'EventRule':
        """
        Get an existing EventRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the rule.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[str] event_pattern: Event pattern
               described a JSON object.
               See full documentation of [CloudWatch Events and Event Patterns](http://docs.aws.amazon.com/AmazonCloudWatch/latest/DeveloperGuide/CloudWatchEventsandEventPatterns.html) for details.
        :param pulumi.Input[bool] is_enabled: Whether the rule should be enabled (defaults to `true`).
        :param pulumi.Input[str] name: The rule's name. By default generated by this provider.
        :param pulumi.Input[str] name_prefix: The rule's name. Conflicts with `name`.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) associated with the role that is used for target invocation.
        :param pulumi.Input[str] schedule_expression: The scheduling expression.
               For example, `cron(0 20 * * ? *)` or `rate(5 minutes)`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["event_pattern"] = event_pattern
        __props__["is_enabled"] = is_enabled
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["role_arn"] = role_arn
        __props__["schedule_expression"] = schedule_expression
        __props__["tags"] = tags
        return EventRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

