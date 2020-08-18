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

__all__ = ['TopicRule']


class TopicRule(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cloudwatch_alarm: Optional[pulumi.Input[pulumi.InputType['TopicRuleCloudwatchAlarmArgs']]] = None,
                 cloudwatch_metric: Optional[pulumi.Input[pulumi.InputType['TopicRuleCloudwatchMetricArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dynamodb: Optional[pulumi.Input[pulumi.InputType['TopicRuleDynamodbArgs']]] = None,
                 dynamodbv2s: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleDynamodbv2Args']]]]] = None,
                 elasticsearch: Optional[pulumi.Input[pulumi.InputType['TopicRuleElasticsearchArgs']]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 error_action: Optional[pulumi.Input[pulumi.InputType['TopicRuleErrorActionArgs']]] = None,
                 firehose: Optional[pulumi.Input[pulumi.InputType['TopicRuleFirehoseArgs']]] = None,
                 iot_analytics: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleIotAnalyticArgs']]]]] = None,
                 iot_events: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleIotEventArgs']]]]] = None,
                 kinesis: Optional[pulumi.Input[pulumi.InputType['TopicRuleKinesisArgs']]] = None,
                 lambda_: Optional[pulumi.Input[pulumi.InputType['TopicRuleLambdaArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 republish: Optional[pulumi.Input[pulumi.InputType['TopicRuleRepublishArgs']]] = None,
                 s3: Optional[pulumi.Input[pulumi.InputType['TopicRuleS3Args']]] = None,
                 sns: Optional[pulumi.Input[pulumi.InputType['TopicRuleSnsArgs']]] = None,
                 sql: Optional[pulumi.Input[str]] = None,
                 sql_version: Optional[pulumi.Input[str]] = None,
                 sqs: Optional[pulumi.Input[pulumi.InputType['TopicRuleSqsArgs']]] = None,
                 step_functions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleStepFunctionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        mytopic = aws.sns.Topic("mytopic")
        myerrortopic = aws.sns.Topic("myerrortopic")
        role = aws.iam.Role("role", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": "iot.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }
        \"\"\")
        rule = aws.iot.TopicRule("rule",
            description="Example rule",
            enabled=True,
            sql="SELECT * FROM 'topic/test'",
            sql_version="2016-03-23",
            sns={
                "messageFormat": "RAW",
                "role_arn": role.arn,
                "target_arn": mytopic.arn,
            },
            error_action={
                "sns": {
                    "messageFormat": "RAW",
                    "role_arn": role.arn,
                    "target_arn": myerrortopic.arn,
                },
            })
        iam_policy_for_lambda = aws.iam.RolePolicy("iamPolicyForLambda",
            role=role.id,
            policy=mytopic.arn.apply(lambda arn: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
                "Effect": "Allow",
                "Action": [
                    "sns:Publish"
                ],
                "Resource": "{arn}"
            }}
          ]
        }}
        \"\"\"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[bool] enabled: Specifies whether the rule is enabled.
        :param pulumi.Input[pulumi.InputType['TopicRuleErrorActionArgs']] error_action: Configuration block with error action to be associated with the rule. See the documentation for `cloudwatch_alarm`, `cloudwatch_metric`, `dynamodb`, `dynamodbv2`, `elasticsearch`, `firehose`, `iot_analytics`, `iot_events`, `kinesis`, `lambda`, `republish`, `s3`, `step_functions`, `sns`, `sqs` configuration blocks for further configuration details.
        :param pulumi.Input[str] name: The name of the rule.
        :param pulumi.Input[str] sql: The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
        :param pulumi.Input[str] sql_version: The version of the SQL rules engine to use when evaluating the rule.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
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

            __props__['cloudwatch_alarm'] = cloudwatch_alarm
            __props__['cloudwatch_metric'] = cloudwatch_metric
            __props__['description'] = description
            __props__['dynamodb'] = dynamodb
            __props__['dynamodbv2s'] = dynamodbv2s
            __props__['elasticsearch'] = elasticsearch
            if enabled is None:
                raise TypeError("Missing required property 'enabled'")
            __props__['enabled'] = enabled
            __props__['error_action'] = error_action
            __props__['firehose'] = firehose
            __props__['iot_analytics'] = iot_analytics
            __props__['iot_events'] = iot_events
            __props__['kinesis'] = kinesis
            __props__['lambda_'] = lambda_
            __props__['name'] = name
            __props__['republish'] = republish
            __props__['s3'] = s3
            __props__['sns'] = sns
            if sql is None:
                raise TypeError("Missing required property 'sql'")
            __props__['sql'] = sql
            if sql_version is None:
                raise TypeError("Missing required property 'sql_version'")
            __props__['sql_version'] = sql_version
            __props__['sqs'] = sqs
            __props__['step_functions'] = step_functions
            __props__['tags'] = tags
            __props__['arn'] = None
        super(TopicRule, __self__).__init__(
            'aws:iot/topicRule:TopicRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cloudwatch_alarm: Optional[pulumi.Input[pulumi.InputType['TopicRuleCloudwatchAlarmArgs']]] = None,
            cloudwatch_metric: Optional[pulumi.Input[pulumi.InputType['TopicRuleCloudwatchMetricArgs']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dynamodb: Optional[pulumi.Input[pulumi.InputType['TopicRuleDynamodbArgs']]] = None,
            dynamodbv2s: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleDynamodbv2Args']]]]] = None,
            elasticsearch: Optional[pulumi.Input[pulumi.InputType['TopicRuleElasticsearchArgs']]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            error_action: Optional[pulumi.Input[pulumi.InputType['TopicRuleErrorActionArgs']]] = None,
            firehose: Optional[pulumi.Input[pulumi.InputType['TopicRuleFirehoseArgs']]] = None,
            iot_analytics: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleIotAnalyticArgs']]]]] = None,
            iot_events: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleIotEventArgs']]]]] = None,
            kinesis: Optional[pulumi.Input[pulumi.InputType['TopicRuleKinesisArgs']]] = None,
            lambda_: Optional[pulumi.Input[pulumi.InputType['TopicRuleLambdaArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            republish: Optional[pulumi.Input[pulumi.InputType['TopicRuleRepublishArgs']]] = None,
            s3: Optional[pulumi.Input[pulumi.InputType['TopicRuleS3Args']]] = None,
            sns: Optional[pulumi.Input[pulumi.InputType['TopicRuleSnsArgs']]] = None,
            sql: Optional[pulumi.Input[str]] = None,
            sql_version: Optional[pulumi.Input[str]] = None,
            sqs: Optional[pulumi.Input[pulumi.InputType['TopicRuleSqsArgs']]] = None,
            step_functions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['TopicRuleStepFunctionArgs']]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'TopicRule':
        """
        Get an existing TopicRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the topic rule
        :param pulumi.Input[str] description: The description of the rule.
        :param pulumi.Input[bool] enabled: Specifies whether the rule is enabled.
        :param pulumi.Input[pulumi.InputType['TopicRuleErrorActionArgs']] error_action: Configuration block with error action to be associated with the rule. See the documentation for `cloudwatch_alarm`, `cloudwatch_metric`, `dynamodb`, `dynamodbv2`, `elasticsearch`, `firehose`, `iot_analytics`, `iot_events`, `kinesis`, `lambda`, `republish`, `s3`, `step_functions`, `sns`, `sqs` configuration blocks for further configuration details.
        :param pulumi.Input[str] name: The name of the rule.
        :param pulumi.Input[str] sql: The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
        :param pulumi.Input[str] sql_version: The version of the SQL rules engine to use when evaluating the rule.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["cloudwatch_alarm"] = cloudwatch_alarm
        __props__["cloudwatch_metric"] = cloudwatch_metric
        __props__["description"] = description
        __props__["dynamodb"] = dynamodb
        __props__["dynamodbv2s"] = dynamodbv2s
        __props__["elasticsearch"] = elasticsearch
        __props__["enabled"] = enabled
        __props__["error_action"] = error_action
        __props__["firehose"] = firehose
        __props__["iot_analytics"] = iot_analytics
        __props__["iot_events"] = iot_events
        __props__["kinesis"] = kinesis
        __props__["lambda_"] = lambda_
        __props__["name"] = name
        __props__["republish"] = republish
        __props__["s3"] = s3
        __props__["sns"] = sns
        __props__["sql"] = sql
        __props__["sql_version"] = sql_version
        __props__["sqs"] = sqs
        __props__["step_functions"] = step_functions
        __props__["tags"] = tags
        return TopicRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the topic rule
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="cloudwatchAlarm")
    def cloudwatch_alarm(self) -> Optional['outputs.TopicRuleCloudwatchAlarm']:
        return pulumi.get(self, "cloudwatch_alarm")

    @property
    @pulumi.getter(name="cloudwatchMetric")
    def cloudwatch_metric(self) -> Optional['outputs.TopicRuleCloudwatchMetric']:
        return pulumi.get(self, "cloudwatch_metric")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def dynamodb(self) -> Optional['outputs.TopicRuleDynamodb']:
        return pulumi.get(self, "dynamodb")

    @property
    @pulumi.getter
    def dynamodbv2s(self) -> Optional[List['outputs.TopicRuleDynamodbv2']]:
        return pulumi.get(self, "dynamodbv2s")

    @property
    @pulumi.getter
    def elasticsearch(self) -> Optional['outputs.TopicRuleElasticsearch']:
        return pulumi.get(self, "elasticsearch")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Specifies whether the rule is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="errorAction")
    def error_action(self) -> Optional['outputs.TopicRuleErrorAction']:
        """
        Configuration block with error action to be associated with the rule. See the documentation for `cloudwatch_alarm`, `cloudwatch_metric`, `dynamodb`, `dynamodbv2`, `elasticsearch`, `firehose`, `iot_analytics`, `iot_events`, `kinesis`, `lambda`, `republish`, `s3`, `step_functions`, `sns`, `sqs` configuration blocks for further configuration details.
        """
        return pulumi.get(self, "error_action")

    @property
    @pulumi.getter
    def firehose(self) -> Optional['outputs.TopicRuleFirehose']:
        return pulumi.get(self, "firehose")

    @property
    @pulumi.getter(name="iotAnalytics")
    def iot_analytics(self) -> Optional[List['outputs.TopicRuleIotAnalytic']]:
        return pulumi.get(self, "iot_analytics")

    @property
    @pulumi.getter(name="iotEvents")
    def iot_events(self) -> Optional[List['outputs.TopicRuleIotEvent']]:
        return pulumi.get(self, "iot_events")

    @property
    @pulumi.getter
    def kinesis(self) -> Optional['outputs.TopicRuleKinesis']:
        return pulumi.get(self, "kinesis")

    @property
    @pulumi.getter(name="lambda")
    def lambda_(self) -> Optional['outputs.TopicRuleLambda']:
        return pulumi.get(self, "lambda_")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the rule.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def republish(self) -> Optional['outputs.TopicRuleRepublish']:
        return pulumi.get(self, "republish")

    @property
    @pulumi.getter
    def s3(self) -> Optional['outputs.TopicRuleS3']:
        return pulumi.get(self, "s3")

    @property
    @pulumi.getter
    def sns(self) -> Optional['outputs.TopicRuleSns']:
        return pulumi.get(self, "sns")

    @property
    @pulumi.getter
    def sql(self) -> str:
        """
        The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
        """
        return pulumi.get(self, "sql")

    @property
    @pulumi.getter(name="sqlVersion")
    def sql_version(self) -> str:
        """
        The version of the SQL rules engine to use when evaluating the rule.
        """
        return pulumi.get(self, "sql_version")

    @property
    @pulumi.getter
    def sqs(self) -> Optional['outputs.TopicRuleSqs']:
        return pulumi.get(self, "sqs")

    @property
    @pulumi.getter(name="stepFunctions")
    def step_functions(self) -> Optional[List['outputs.TopicRuleStepFunction']]:
        return pulumi.get(self, "step_functions")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

