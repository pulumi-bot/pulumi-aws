# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Stream(pulumi.CustomResource):
    """
    Provides a Kinesis Stream resource. Amazon Kinesis is a managed service that
    scales elastically for real-time processing of streaming big data.
    
    For more details, see the [Amazon Kinesis Documentation][1].
    """
    def __init__(__self__, __name__, __opts__=None, arn=None, encryption_type=None, kms_key_id=None, name=None, retention_period=None, shard_count=None, shard_level_metrics=None, tags=None):
        """Create a Stream resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['arn'] = arn

        __props__['encryptionType'] = encryption_type

        __props__['kmsKeyId'] = kms_key_id

        __props__['name'] = name

        __props__['retentionPeriod'] = retention_period

        if not shard_count:
            raise TypeError('Missing required property shard_count')
        __props__['shardCount'] = shard_count

        __props__['shardLevelMetrics'] = shard_level_metrics

        __props__['tags'] = tags

        super(Stream, __self__).__init__(
            'aws:kinesis/stream:Stream',
            __name__,
            __props__,
            __opts__)

