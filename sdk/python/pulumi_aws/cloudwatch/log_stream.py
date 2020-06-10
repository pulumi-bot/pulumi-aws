# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class LogStream(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) specifying the log stream.
    """
    log_group_name: pulumi.Output[str]
    """
    The name of the log group under which the log stream is to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the log stream. Must not be longer than 512 characters and must not contain `:`
    """
    def __init__(__self__, resource_name, opts=None, log_group_name=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CloudWatch Log Stream resource.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        yada = aws.cloudwatch.LogGroup("yada")
        foo = aws.cloudwatch.LogStream("foo", log_group_name=yada.name)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] log_group_name: The name of the log group under which the log stream is to be created.
        :param pulumi.Input[str] name: The name of the log stream. Must not be longer than 512 characters and must not contain `:`
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

            if log_group_name is None:
                raise TypeError("Missing required property 'log_group_name'")
            __props__['log_group_name'] = log_group_name
            __props__['name'] = name
            __props__['arn'] = None
        super(LogStream, __self__).__init__(
            'aws:cloudwatch/logStream:LogStream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, log_group_name=None, name=None):
        """
        Get an existing LogStream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the log stream.
        :param pulumi.Input[str] log_group_name: The name of the log group under which the log stream is to be created.
        :param pulumi.Input[str] name: The name of the log stream. Must not be longer than 512 characters and must not contain `:`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["log_group_name"] = log_group_name
        __props__["name"] = name
        return LogStream(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
