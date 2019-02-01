# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BucketNotification(pulumi.CustomResource):
    bucket: pulumi.Output[str]
    """
    The name of the bucket to put notification configuration.
    """
    lambda_functions: pulumi.Output[list]
    """
    Used to configure notifications to a Lambda Function (documented below).
    """
    queues: pulumi.Output[list]
    """
    The notification configuration to SQS Queue (documented below).
    """
    topics: pulumi.Output[list]
    """
    The notification configuration to SNS Topic (documented below).
    """
    def __init__(__self__, __name__, __opts__=None, bucket=None, lambda_functions=None, queues=None, topics=None):
        """
        Provides a S3 bucket notification resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket to put notification configuration.
        :param pulumi.Input[list] lambda_functions: Used to configure notifications to a Lambda Function (documented below).
        :param pulumi.Input[list] queues: The notification configuration to SQS Queue (documented below).
        :param pulumi.Input[list] topics: The notification configuration to SNS Topic (documented below).
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not bucket:
            raise TypeError('Missing required property bucket')
        __props__['bucket'] = bucket

        __props__['lambda_functions'] = lambda_functions

        __props__['queues'] = queues

        __props__['topics'] = topics

        super(BucketNotification, __self__).__init__(
            'aws:s3/bucketNotification:BucketNotification',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

