# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class EventTarget(pulumi.CustomResource):
    """
    Provides a CloudWatch Event Target resource.
    """
    def __init__(__self__, __name__, __opts__=None, arn=None, batch_target=None, ecs_target=None, input=None, input_path=None, input_transformer=None, kinesis_target=None, role_arn=None, rule=None, run_command_targets=None, sqs_target=None, target_id=None):
        """Create a EventTarget resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not arn:
            raise TypeError('Missing required property arn')
        __props__['arn'] = arn

        __props__['batchTarget'] = batch_target

        __props__['ecsTarget'] = ecs_target

        __props__['input'] = input

        __props__['inputPath'] = input_path

        __props__['inputTransformer'] = input_transformer

        __props__['kinesisTarget'] = kinesis_target

        __props__['roleArn'] = role_arn

        if not rule:
            raise TypeError('Missing required property rule')
        __props__['rule'] = rule

        __props__['runCommandTargets'] = run_command_targets

        __props__['sqsTarget'] = sqs_target

        __props__['targetId'] = target_id

        super(EventTarget, __self__).__init__(
            'aws:cloudwatch/eventTarget:EventTarget',
            __name__,
            __props__,
            __opts__)

