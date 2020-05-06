# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("aws.applicationloadbalancing.TargetGroupAttachment has been deprecated in favour of aws.alb.TargetGroupAttachment", DeprecationWarning)
class TargetGroupAttachment(pulumi.CustomResource):
    availability_zone: pulumi.Output[str]
    """
    The Availability Zone where the IP address of the target is to be registered.
    """
    port: pulumi.Output[float]
    """
    The port on which targets receive traffic.
    """
    target_group_arn: pulumi.Output[str]
    """
    The ARN of the target group with which to register targets
    """
    target_id: pulumi.Output[str]
    """
    The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
    """
    warnings.warn("aws.applicationloadbalancing.TargetGroupAttachment has been deprecated in favour of aws.alb.TargetGroupAttachment", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, availability_zone=None, port=None, target_group_arn=None, target_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the [`elb.Attachment` resource](https://www.terraform.io/docs/providers/aws/r/elb_attachment.html).

        > **Note:** `alb.TargetGroupAttachment` is known as `lb.TargetGroupAttachment`. The functionality is identical.


        ## Usage with lambda

        ```python
        import pulumi
        import pulumi_aws as aws

        test_target_group = aws.lb.TargetGroup("testTargetGroup", target_type="lambda")
        test_function = aws.lambda_.Function("testFunction")
        with_lb = aws.lambda_.Permission("withLb",
            action="lambda:InvokeFunction",
            function=test_function.arn,
            principal="elasticloadbalancing.amazonaws.com",
            source_arn=test_target_group.arn)
        test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
            target_group_arn=test_target_group.arn,
            target_id=test_function.arn)
        ```

        Deprecated: aws.applicationloadbalancing.TargetGroupAttachment has been deprecated in favour of aws.alb.TargetGroupAttachment

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone where the IP address of the target is to be registered.
        :param pulumi.Input[float] port: The port on which targets receive traffic.
        :param pulumi.Input[str] target_group_arn: The ARN of the target group with which to register targets
        :param pulumi.Input[str] target_id: The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        """
        pulumi.log.warn("TargetGroupAttachment is deprecated: aws.applicationloadbalancing.TargetGroupAttachment has been deprecated in favour of aws.alb.TargetGroupAttachment")
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

            __props__['availability_zone'] = availability_zone
            __props__['port'] = port
            if target_group_arn is None:
                raise TypeError("Missing required property 'target_group_arn'")
            __props__['target_group_arn'] = target_group_arn
            if target_id is None:
                raise TypeError("Missing required property 'target_id'")
            __props__['target_id'] = target_id
        super(TargetGroupAttachment, __self__).__init__(
            'aws:applicationloadbalancing/targetGroupAttachment:TargetGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, availability_zone=None, port=None, target_group_arn=None, target_id=None):
        """
        Get an existing TargetGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone where the IP address of the target is to be registered.
        :param pulumi.Input[float] port: The port on which targets receive traffic.
        :param pulumi.Input[str] target_group_arn: The ARN of the target group with which to register targets
        :param pulumi.Input[str] target_id: The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["availability_zone"] = availability_zone
        __props__["port"] = port
        __props__["target_group_arn"] = target_group_arn
        __props__["target_id"] = target_id
        return TargetGroupAttachment(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

