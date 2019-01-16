# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BucketNotification(pulumi.CustomResource):
    """
    Provides a S3 bucket notification resource.
    """
    def __init__(__self__, __name__, __opts__=None, bucket=None, lambda_functions=None, queues=None, topics=None):
        """Create a BucketNotification resource with the given unique name, props, and options."""
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

