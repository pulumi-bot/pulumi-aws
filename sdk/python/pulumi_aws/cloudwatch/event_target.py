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

__all__ = ['EventTarget']


class EventTarget(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 batch_target: Optional[pulumi.Input[pulumi.InputType['EventTargetBatchTargetArgs']]] = None,
                 ecs_target: Optional[pulumi.Input[pulumi.InputType['EventTargetEcsTargetArgs']]] = None,
                 input: Optional[pulumi.Input[str]] = None,
                 input_path: Optional[pulumi.Input[str]] = None,
                 input_transformer: Optional[pulumi.Input[pulumi.InputType['EventTargetInputTransformerArgs']]] = None,
                 kinesis_target: Optional[pulumi.Input[pulumi.InputType['EventTargetKinesisTargetArgs']]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 rule: Optional[pulumi.Input[str]] = None,
                 run_command_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EventTargetRunCommandTargetArgs']]]]] = None,
                 sqs_target: Optional[pulumi.Input[pulumi.InputType['EventTargetSqsTargetArgs']]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudWatch Event Target resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        console = aws.cloudwatch.EventRule("console",
            description="Capture all EC2 scaling events",
            event_pattern=\"\"\"{
          "source": [
            "aws.autoscaling"
          ],
          "detail-type": [
            "EC2 Instance Launch Successful",
            "EC2 Instance Terminate Successful",
            "EC2 Instance Launch Unsuccessful",
            "EC2 Instance Terminate Unsuccessful"
          ]
        }
        \"\"\")
        test_stream = aws.kinesis.Stream("testStream", shard_count=1)
        yada = aws.cloudwatch.EventTarget("yada",
            rule=console.name,
            arn=test_stream.arn,
            run_command_targets=[
                {
                    "key": "tag:Name",
                    "values": ["FooBar"],
                },
                {
                    "key": "InstanceIds",
                    "values": ["i-162058cd308bffec2"],
                },
            ])
        ```
        ## Example SSM Document Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        ssm_lifecycle_trust = aws.iam.get_policy_document(statements=[{
            "actions": ["sts:AssumeRole"],
            "principals": [{
                "type": "Service",
                "identifiers": ["events.amazonaws.com"],
            }],
        }])
        stop_instance = aws.ssm.Document("stopInstance",
            document_type="Command",
            content=\"\"\"  {
            "schemaVersion": "1.2",
            "description": "Stop an instance",
            "parameters": {

            },
            "runtimeConfig": {
              "aws:runShellScript": {
                "properties": [
                  {
                    "id": "0.aws:runShellScript",
                    "runCommand": ["halt"]
                  }
                ]
              }
            }
          }
        \"\"\")
        ssm_lifecycle_policy_document = stop_instance.arn.apply(lambda arn: aws.iam.get_policy_document(statements=[
            {
                "effect": "Allow",
                "actions": ["ssm:SendCommand"],
                "resources": ["arn:aws:ec2:eu-west-1:1234567890:instance/*"],
                "conditions": [{
                    "test": "StringEquals",
                    "variable": "ec2:ResourceTag/Terminate",
                    "values": ["*"],
                }],
            },
            {
                "effect": "Allow",
                "actions": ["ssm:SendCommand"],
                "resources": [arn],
            },
        ]))
        ssm_lifecycle_role = aws.iam.Role("ssmLifecycleRole", assume_role_policy=ssm_lifecycle_trust.json)
        ssm_lifecycle_policy = aws.iam.Policy("ssmLifecyclePolicy", policy=ssm_lifecycle_policy_document.json)
        stop_instances_event_rule = aws.cloudwatch.EventRule("stopInstancesEventRule",
            description="Stop instances nightly",
            schedule_expression="cron(0 0 * * ? *)")
        stop_instances_event_target = aws.cloudwatch.EventTarget("stopInstancesEventTarget",
            arn=stop_instance.arn,
            rule=stop_instances_event_rule.name,
            role_arn=ssm_lifecycle_role.arn,
            run_command_targets=[{
                "key": "tag:Terminate",
                "values": ["midnight"],
            }])
        ```

        ## Example RunCommand Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        stop_instances_event_rule = aws.cloudwatch.EventRule("stopInstancesEventRule",
            description="Stop instances nightly",
            schedule_expression="cron(0 0 * * ? *)")
        stop_instances_event_target = aws.cloudwatch.EventTarget("stopInstancesEventTarget",
            arn=f"arn:aws:ssm:{var['aws_region']}::document/AWS-RunShellScript",
            input="{\"commands\":[\"halt\"]}",
            rule=stop_instances_event_rule.name,
            role_arn=aws_iam_role["ssm_lifecycle"]["arn"],
            run_command_targets=[{
                "key": "tag:Terminate",
                "values": ["midnight"],
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) associated of the target.
        :param pulumi.Input[pulumi.InputType['EventTargetBatchTargetArgs']] batch_target: Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[pulumi.InputType['EventTargetEcsTargetArgs']] ecs_target: Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] input: Valid JSON text passed to the target.
        :param pulumi.Input[str] input_path: The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
               that is used for extracting part of the matched event when passing it to the target.
        :param pulumi.Input[pulumi.InputType['EventTargetInputTransformerArgs']] input_transformer: Parameters used when you are providing a custom input to a target based on certain event data.
        :param pulumi.Input[pulumi.InputType['EventTargetKinesisTargetArgs']] kinesis_target: Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
        :param pulumi.Input[str] rule: The name of the rule you want to add targets to.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EventTargetRunCommandTargetArgs']]]] run_command_targets: Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        :param pulumi.Input[pulumi.InputType['EventTargetSqsTargetArgs']] sqs_target: Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] target_id: The unique target assignment ID.  If missing, will generate a random, unique id.
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

            if arn is None:
                raise TypeError("Missing required property 'arn'")
            __props__['arn'] = arn
            __props__['batch_target'] = batch_target
            __props__['ecs_target'] = ecs_target
            __props__['input'] = input
            __props__['input_path'] = input_path
            __props__['input_transformer'] = input_transformer
            __props__['kinesis_target'] = kinesis_target
            __props__['role_arn'] = role_arn
            if rule is None:
                raise TypeError("Missing required property 'rule'")
            __props__['rule'] = rule
            __props__['run_command_targets'] = run_command_targets
            __props__['sqs_target'] = sqs_target
            __props__['target_id'] = target_id
        super(EventTarget, __self__).__init__(
            'aws:cloudwatch/eventTarget:EventTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            batch_target: Optional[pulumi.Input[pulumi.InputType['EventTargetBatchTargetArgs']]] = None,
            ecs_target: Optional[pulumi.Input[pulumi.InputType['EventTargetEcsTargetArgs']]] = None,
            input: Optional[pulumi.Input[str]] = None,
            input_path: Optional[pulumi.Input[str]] = None,
            input_transformer: Optional[pulumi.Input[pulumi.InputType['EventTargetInputTransformerArgs']]] = None,
            kinesis_target: Optional[pulumi.Input[pulumi.InputType['EventTargetKinesisTargetArgs']]] = None,
            role_arn: Optional[pulumi.Input[str]] = None,
            rule: Optional[pulumi.Input[str]] = None,
            run_command_targets: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EventTargetRunCommandTargetArgs']]]]] = None,
            sqs_target: Optional[pulumi.Input[pulumi.InputType['EventTargetSqsTargetArgs']]] = None,
            target_id: Optional[pulumi.Input[str]] = None) -> 'EventTarget':
        """
        Get an existing EventTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) associated of the target.
        :param pulumi.Input[pulumi.InputType['EventTargetBatchTargetArgs']] batch_target: Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[pulumi.InputType['EventTargetEcsTargetArgs']] ecs_target: Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] input: Valid JSON text passed to the target.
        :param pulumi.Input[str] input_path: The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
               that is used for extracting part of the matched event when passing it to the target.
        :param pulumi.Input[pulumi.InputType['EventTargetInputTransformerArgs']] input_transformer: Parameters used when you are providing a custom input to a target based on certain event data.
        :param pulumi.Input[pulumi.InputType['EventTargetKinesisTargetArgs']] kinesis_target: Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] role_arn: The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
        :param pulumi.Input[str] rule: The name of the rule you want to add targets to.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EventTargetRunCommandTargetArgs']]]] run_command_targets: Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        :param pulumi.Input[pulumi.InputType['EventTargetSqsTargetArgs']] sqs_target: Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        :param pulumi.Input[str] target_id: The unique target assignment ID.  If missing, will generate a random, unique id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["batch_target"] = batch_target
        __props__["ecs_target"] = ecs_target
        __props__["input"] = input
        __props__["input_path"] = input_path
        __props__["input_transformer"] = input_transformer
        __props__["kinesis_target"] = kinesis_target
        __props__["role_arn"] = role_arn
        __props__["rule"] = rule
        __props__["run_command_targets"] = run_command_targets
        __props__["sqs_target"] = sqs_target
        __props__["target_id"] = target_id
        return EventTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) associated of the target.
        """
        ...

    @property
    @pulumi.getter(name="batchTarget")
    def batch_target(self) -> Optional['outputs.EventTargetBatchTarget']:
        """
        Parameters used when you are using the rule to invoke an Amazon Batch Job. Documented below. A maximum of 1 are allowed.
        """
        ...

    @property
    @pulumi.getter(name="ecsTarget")
    def ecs_target(self) -> Optional['outputs.EventTargetEcsTarget']:
        """
        Parameters used when you are using the rule to invoke Amazon ECS Task. Documented below. A maximum of 1 are allowed.
        """
        ...

    @property
    @pulumi.getter
    def input(self) -> Optional[str]:
        """
        Valid JSON text passed to the target.
        """
        ...

    @property
    @pulumi.getter(name="inputPath")
    def input_path(self) -> Optional[str]:
        """
        The value of the [JSONPath](http://goessner.net/articles/JsonPath/)
        that is used for extracting part of the matched event when passing it to the target.
        """
        ...

    @property
    @pulumi.getter(name="inputTransformer")
    def input_transformer(self) -> Optional['outputs.EventTargetInputTransformer']:
        """
        Parameters used when you are providing a custom input to a target based on certain event data.
        """
        ...

    @property
    @pulumi.getter(name="kinesisTarget")
    def kinesis_target(self) -> Optional['outputs.EventTargetKinesisTarget']:
        """
        Parameters used when you are using the rule to invoke an Amazon Kinesis Stream. Documented below. A maximum of 1 are allowed.
        """
        ...

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        The Amazon Resource Name (ARN) of the IAM role to be used for this target when the rule is triggered. Required if `ecs_target` is used.
        """
        ...

    @property
    @pulumi.getter
    def rule(self) -> str:
        """
        The name of the rule you want to add targets to.
        """
        ...

    @property
    @pulumi.getter(name="runCommandTargets")
    def run_command_targets(self) -> Optional[List['outputs.EventTargetRunCommandTarget']]:
        """
        Parameters used when you are using the rule to invoke Amazon EC2 Run Command. Documented below. A maximum of 5 are allowed.
        """
        ...

    @property
    @pulumi.getter(name="sqsTarget")
    def sqs_target(self) -> Optional['outputs.EventTargetSqsTarget']:
        """
        Parameters used when you are using the rule to invoke an Amazon SQS Queue. Documented below. A maximum of 1 are allowed.
        """
        ...

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        The unique target assignment ID.  If missing, will generate a random, unique id.
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

