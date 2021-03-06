# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class DeliveryChannel(pulumi.CustomResource):
    """
    Provides an AWS Config Delivery Channel.
    
    ~> **Note:** Delivery Channel requires a [Configuration Recorder](/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.
    """
    def __init__(__self__, __name__, __opts__=None, name=None, s3_bucket_name=None, s3_key_prefix=None, snapshot_delivery_properties=None, sns_topic_arn=None):
        """Create a DeliveryChannel resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
        """
        __props__['name'] = name

        if not s3_bucket_name:
            raise TypeError('Missing required property s3_bucket_name')
        elif not isinstance(s3_bucket_name, basestring):
            raise TypeError('Expected property s3_bucket_name to be a basestring')
        __self__.s3_bucket_name = s3_bucket_name
        """
        The name of the S3 bucket used to store the configuration history.
        """
        __props__['s3BucketName'] = s3_bucket_name

        if s3_key_prefix and not isinstance(s3_key_prefix, basestring):
            raise TypeError('Expected property s3_key_prefix to be a basestring')
        __self__.s3_key_prefix = s3_key_prefix
        """
        The prefix for the specified S3 bucket.
        """
        __props__['s3KeyPrefix'] = s3_key_prefix

        if snapshot_delivery_properties and not isinstance(snapshot_delivery_properties, dict):
            raise TypeError('Expected property snapshot_delivery_properties to be a dict')
        __self__.snapshot_delivery_properties = snapshot_delivery_properties
        """
        Options for how AWS Config delivers configuration snapshots. See below
        """
        __props__['snapshotDeliveryProperties'] = snapshot_delivery_properties

        if sns_topic_arn and not isinstance(sns_topic_arn, basestring):
            raise TypeError('Expected property sns_topic_arn to be a basestring')
        __self__.sns_topic_arn = sns_topic_arn
        """
        The ARN of the SNS topic that AWS Config delivers notifications to.
        """
        __props__['snsTopicArn'] = sns_topic_arn

        super(DeliveryChannel, __self__).__init__(
            'aws:cfg/deliveryChannel:DeliveryChannel',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'name' in outs:
            self.name = outs['name']
        if 's3BucketName' in outs:
            self.s3_bucket_name = outs['s3BucketName']
        if 's3KeyPrefix' in outs:
            self.s3_key_prefix = outs['s3KeyPrefix']
        if 'snapshotDeliveryProperties' in outs:
            self.snapshot_delivery_properties = outs['snapshotDeliveryProperties']
        if 'snsTopicArn' in outs:
            self.sns_topic_arn = outs['snsTopicArn']
