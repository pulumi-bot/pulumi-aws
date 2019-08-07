# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DeliveryChannel(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
    """
    s3_bucket_name: pulumi.Output[str]
    """
    The name of the S3 bucket used to store the configuration history.
    """
    s3_key_prefix: pulumi.Output[str]
    """
    The prefix for the specified S3 bucket.
    """
    snapshot_delivery_properties: pulumi.Output[dict]
    """
    Options for how AWS Config delivers configuration snapshots. See below
    """
    sns_topic_arn: pulumi.Output[str]
    """
    The ARN of the SNS topic that AWS Config delivers notifications to.
    """
    def __init__(__self__, resource_name, opts=None, name=None, s3_bucket_name=None, s3_key_prefix=None, snapshot_delivery_properties=None, sns_topic_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Config Delivery Channel.
        
        > **Note:** Delivery Channel requires a [Configuration Recorder](https://www.terraform.io/docs/providers/aws/r/config_configuration_recorder.html) to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
        :param pulumi.Input[str] s3_bucket_name: The name of the S3 bucket used to store the configuration history.
        :param pulumi.Input[str] s3_key_prefix: The prefix for the specified S3 bucket.
        :param pulumi.Input[dict] snapshot_delivery_properties: Options for how AWS Config delivers configuration snapshots. See below
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic that AWS Config delivers notifications to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_delivery_channel.html.markdown.
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
                raise TypeError("[__props__] should only be provided when [opts.id] was not [None].")
            __props__ = dict()

            __props__['name'] = name
            if s3_bucket_name is None:
                raise TypeError("Missing required property 's3_bucket_name'")
            __props__['s3_bucket_name'] = s3_bucket_name
            __props__['s3_key_prefix'] = s3_key_prefix
            __props__['snapshot_delivery_properties'] = snapshot_delivery_properties
            __props__['sns_topic_arn'] = sns_topic_arn
        super(DeliveryChannel, __self__).__init__(
            'aws:cfg/deliveryChannel:DeliveryChannel',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, s3_bucket_name=None, s3_key_prefix=None, snapshot_delivery_properties=None, sns_topic_arn=None):
        """
        Get an existing DeliveryChannel resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
        :param pulumi.Input[str] s3_bucket_name: The name of the S3 bucket used to store the configuration history.
        :param pulumi.Input[str] s3_key_prefix: The prefix for the specified S3 bucket.
        :param pulumi.Input[dict] snapshot_delivery_properties: Options for how AWS Config delivers configuration snapshots. See below
        :param pulumi.Input[str] sns_topic_arn: The ARN of the SNS topic that AWS Config delivers notifications to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/config_delivery_channel.html.markdown.
        """
        opts = pulumi.ResourceOptions(id=id) if opts is None else opts.merge(pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["name"] = name
        __props__["s3_bucket_name"] = s3_bucket_name
        __props__["s3_key_prefix"] = s3_key_prefix
        __props__["snapshot_delivery_properties"] = snapshot_delivery_properties
        __props__["sns_topic_arn"] = sns_topic_arn
        return DeliveryChannel(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

