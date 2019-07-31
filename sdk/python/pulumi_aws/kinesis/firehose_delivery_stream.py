# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class FirehoseDeliveryStream(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) specifying the Stream
    """
    destination: pulumi.Output[str]
    """
    This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
    """
    destination_id: pulumi.Output[str]
    elasticsearch_configuration: pulumi.Output[dict]
    extended_s3_configuration: pulumi.Output[dict]
    """
    Enhanced configuration options for the s3 destination. More details are given below.
    """
    kinesis_source_configuration: pulumi.Output[dict]
    """
    Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
    """
    name: pulumi.Output[str]
    """
    A name to identify the stream. This is unique to the
    AWS account and region the Stream is created in.
    """
    redshift_configuration: pulumi.Output[dict]
    """
    Configuration options if redshift is the destination.
    Using `redshift_configuration` requires the user to also specify a
    `s3_configuration` block. More details are given below.
    """
    s3_configuration: pulumi.Output[dict]
    """
    Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
    is redshift). More details are given below.
    """
    splunk_configuration: pulumi.Output[dict]
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    version_id: pulumi.Output[str]
    """
    Specifies the table version for the output data schema. Defaults to `LATEST`.
    """
    def __init__(__self__, resource_name, opts=None, arn=None, destination=None, destination_id=None, elasticsearch_configuration=None, extended_s3_configuration=None, kinesis_source_configuration=None, name=None, redshift_configuration=None, s3_configuration=None, splunk_configuration=None, tags=None, version_id=None, __name__=None, __opts__=None):
        """
        Provides a Kinesis Firehose Delivery Stream resource. Amazon Kinesis Firehose is a fully managed, elastic service to easily deliver real-time data streams to destinations such as Amazon S3 and Amazon Redshift.
        
        For more details, see the [Amazon Kinesis Firehose Documentation][1].
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the Stream
        :param pulumi.Input[str] destination: This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
        :param pulumi.Input[dict] extended_s3_configuration: Enhanced configuration options for the s3 destination. More details are given below.
        :param pulumi.Input[dict] kinesis_source_configuration: Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
        :param pulumi.Input[str] name: A name to identify the stream. This is unique to the
               AWS account and region the Stream is created in.
        :param pulumi.Input[dict] redshift_configuration: Configuration options if redshift is the destination.
               Using `redshift_configuration` requires the user to also specify a
               `s3_configuration` block. More details are given below.
        :param pulumi.Input[dict] s3_configuration: Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
               is redshift). More details are given below.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] version_id: Specifies the table version for the output data schema. Defaults to `LATEST`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kinesis_firehose_delivery_stream.html.markdown.
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

        __props__['arn'] = arn

        if destination is None:
            raise TypeError("Missing required property 'destination'")
        __props__['destination'] = destination

        __props__['destination_id'] = destination_id

        __props__['elasticsearch_configuration'] = elasticsearch_configuration

        __props__['extended_s3_configuration'] = extended_s3_configuration

        __props__['kinesis_source_configuration'] = kinesis_source_configuration

        __props__['name'] = name

        __props__['redshift_configuration'] = redshift_configuration

        __props__['s3_configuration'] = s3_configuration

        __props__['splunk_configuration'] = splunk_configuration

        __props__['tags'] = tags

        __props__['version_id'] = version_id

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(FirehoseDeliveryStream, __self__).__init__(
            'aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

