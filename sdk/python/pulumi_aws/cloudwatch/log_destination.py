# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class LogDestination(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) specifying the log destination.
    """
    name: pulumi.Output[str]
    """
    A name for the log destination
    """
    role_arn: pulumi.Output[str]
    """
    The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
    """
    target_arn: pulumi.Output[str]
    """
    The ARN of the target Amazon Kinesis stream resource for the destination
    """
    def __init__(__self__, resource_name, opts=None, name=None, role_arn=None, target_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CloudWatch Logs destination resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_destination = aws.cloudwatch.LogDestination("testDestination",
            role_arn=aws_iam_role["iam_for_cloudwatch"]["arn"],
            target_arn=aws_kinesis_stream["kinesis_for_cloudwatch"]["arn"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: A name for the log destination
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
        :param pulumi.Input[str] target_arn: The ARN of the target Amazon Kinesis stream resource for the destination
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

            __props__['name'] = name
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['roleArn'] = role_arn
            if target_arn is None:
                raise TypeError("Missing required property 'target_arn'")
            __props__['targetArn'] = target_arn
            __props__['arn'] = None
        super(LogDestination, __self__).__init__(
            'aws:cloudwatch/logDestination:LogDestination',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, name=None, role_arn=None, target_arn=None):
        """
        Get an existing LogDestination resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the log destination.
        :param pulumi.Input[str] name: A name for the log destination
        :param pulumi.Input[str] role_arn: The ARN of an IAM role that grants Amazon CloudWatch Logs permissions to put data into the target
        :param pulumi.Input[str] target_arn: The ARN of the target Amazon Kinesis stream resource for the destination
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["role_arn"] = role_arn
        __props__["target_arn"] = target_arn
        return LogDestination(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
