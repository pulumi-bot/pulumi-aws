# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class EventStream(pulumi.CustomResource):
    application_id: pulumi.Output[str]
    """
    The application ID.
    """
    destination_stream_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
    """
    role_arn: pulumi.Output[str]
    """
    The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
    """
    def __init__(__self__, resource_name, opts=None, application_id=None, destination_stream_arn=None, role_arn=None, __name__=None, __opts__=None):
        """
        Provides a Pinpoint Event Stream resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_id: The application ID.
        :param pulumi.Input[str] destination_stream_arn: The Amazon Resource Name (ARN) of the Amazon Kinesis stream or Firehose delivery stream to which you want to publish events.
        :param pulumi.Input[str] role_arn: The IAM role that authorizes Amazon Pinpoint to publish events to the stream in your account.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if application_id is None:
            raise TypeError("Missing required property 'application_id'")
        __props__['application_id'] = application_id

        if destination_stream_arn is None:
            raise TypeError("Missing required property 'destination_stream_arn'")
        __props__['destination_stream_arn'] = destination_stream_arn

        if role_arn is None:
            raise TypeError("Missing required property 'role_arn'")
        __props__['role_arn'] = role_arn

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(EventStream, __self__).__init__(
            'aws:pinpoint/eventStream:EventStream',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

