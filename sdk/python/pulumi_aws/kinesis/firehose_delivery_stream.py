# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['FirehoseDeliveryStream']


class FirehoseDeliveryStream(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 destination_id: Optional[pulumi.Input[str]] = None,
                 elasticsearch_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamElasticsearchConfigurationArgs']]] = None,
                 extended_s3_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamExtendedS3ConfigurationArgs']]] = None,
                 kinesis_source_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamKinesisSourceConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 redshift_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamRedshiftConfigurationArgs']]] = None,
                 s3_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamS3ConfigurationArgs']]] = None,
                 server_side_encryption: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamServerSideEncryptionArgs']]] = None,
                 splunk_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamSplunkConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 version_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Kinesis Firehose Delivery Stream resource. Amazon Kinesis Firehose is a fully managed, elastic service to easily deliver real-time data streams to destinations such as Amazon S3 and Amazon Redshift.

        For more details, see the [Amazon Kinesis Firehose Documentation](https://aws.amazon.com/documentation/firehose/).

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the Stream
        :param pulumi.Input[str] destination: This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamElasticsearchConfigurationArgs']] elasticsearch_configuration: Configuration options if elasticsearch is the destination. More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamExtendedS3ConfigurationArgs']] extended_s3_configuration: Enhanced configuration options for the s3 destination. More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamKinesisSourceConfigurationArgs']] kinesis_source_configuration: Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
        :param pulumi.Input[str] name: A name to identify the stream. This is unique to the
               AWS account and region the Stream is created in.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamRedshiftConfigurationArgs']] redshift_configuration: Configuration options if redshift is the destination.
               Using `redshift_configuration` requires the user to also specify a
               `s3_configuration` block. More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamS3ConfigurationArgs']] s3_configuration: Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
               is redshift). More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamServerSideEncryptionArgs']] server_side_encryption: Encrypt at rest options.
               Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] version_id: Specifies the table version for the output data schema. Defaults to `LATEST`.
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
            __props__['server_side_encryption'] = server_side_encryption
            __props__['splunk_configuration'] = splunk_configuration
            __props__['tags'] = tags
            __props__['version_id'] = version_id
        super(FirehoseDeliveryStream, __self__).__init__(
            'aws:kinesis/firehoseDeliveryStream:FirehoseDeliveryStream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            destination: Optional[pulumi.Input[str]] = None,
            destination_id: Optional[pulumi.Input[str]] = None,
            elasticsearch_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamElasticsearchConfigurationArgs']]] = None,
            extended_s3_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamExtendedS3ConfigurationArgs']]] = None,
            kinesis_source_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamKinesisSourceConfigurationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            redshift_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamRedshiftConfigurationArgs']]] = None,
            s3_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamS3ConfigurationArgs']]] = None,
            server_side_encryption: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamServerSideEncryptionArgs']]] = None,
            splunk_configuration: Optional[pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamSplunkConfigurationArgs']]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            version_id: Optional[pulumi.Input[str]] = None) -> 'FirehoseDeliveryStream':
        """
        Get an existing FirehoseDeliveryStream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the Stream
        :param pulumi.Input[str] destination: This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamElasticsearchConfigurationArgs']] elasticsearch_configuration: Configuration options if elasticsearch is the destination. More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamExtendedS3ConfigurationArgs']] extended_s3_configuration: Enhanced configuration options for the s3 destination. More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamKinesisSourceConfigurationArgs']] kinesis_source_configuration: Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
        :param pulumi.Input[str] name: A name to identify the stream. This is unique to the
               AWS account and region the Stream is created in.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamRedshiftConfigurationArgs']] redshift_configuration: Configuration options if redshift is the destination.
               Using `redshift_configuration` requires the user to also specify a
               `s3_configuration` block. More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamS3ConfigurationArgs']] s3_configuration: Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
               is redshift). More details are given below.
        :param pulumi.Input[pulumi.InputType['FirehoseDeliveryStreamServerSideEncryptionArgs']] server_side_encryption: Encrypt at rest options.
               Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] version_id: Specifies the table version for the output data schema. Defaults to `LATEST`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["destination"] = destination
        __props__["destination_id"] = destination_id
        __props__["elasticsearch_configuration"] = elasticsearch_configuration
        __props__["extended_s3_configuration"] = extended_s3_configuration
        __props__["kinesis_source_configuration"] = kinesis_source_configuration
        __props__["name"] = name
        __props__["redshift_configuration"] = redshift_configuration
        __props__["s3_configuration"] = s3_configuration
        __props__["server_side_encryption"] = server_side_encryption
        __props__["splunk_configuration"] = splunk_configuration
        __props__["tags"] = tags
        __props__["version_id"] = version_id
        return FirehoseDeliveryStream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) specifying the Stream
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[str]:
        """
        This is the destination to where the data is delivered. The only options are `s3` (Deprecated, use `extended_s3` instead), `extended_s3`, `redshift`, `elasticsearch`, and `splunk`.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="destinationId")
    def destination_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "destination_id")

    @property
    @pulumi.getter(name="elasticsearchConfiguration")
    def elasticsearch_configuration(self) -> pulumi.Output[Optional['outputs.FirehoseDeliveryStreamElasticsearchConfiguration']]:
        """
        Configuration options if elasticsearch is the destination. More details are given below.
        """
        return pulumi.get(self, "elasticsearch_configuration")

    @property
    @pulumi.getter(name="extendedS3Configuration")
    def extended_s3_configuration(self) -> pulumi.Output[Optional['outputs.FirehoseDeliveryStreamExtendedS3Configuration']]:
        """
        Enhanced configuration options for the s3 destination. More details are given below.
        """
        return pulumi.get(self, "extended_s3_configuration")

    @property
    @pulumi.getter(name="kinesisSourceConfiguration")
    def kinesis_source_configuration(self) -> pulumi.Output[Optional['outputs.FirehoseDeliveryStreamKinesisSourceConfiguration']]:
        """
        Allows the ability to specify the kinesis stream that is used as the source of the firehose delivery stream.
        """
        return pulumi.get(self, "kinesis_source_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A name to identify the stream. This is unique to the
        AWS account and region the Stream is created in.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="redshiftConfiguration")
    def redshift_configuration(self) -> pulumi.Output[Optional['outputs.FirehoseDeliveryStreamRedshiftConfiguration']]:
        """
        Configuration options if redshift is the destination.
        Using `redshift_configuration` requires the user to also specify a
        `s3_configuration` block. More details are given below.
        """
        return pulumi.get(self, "redshift_configuration")

    @property
    @pulumi.getter(name="s3Configuration")
    def s3_configuration(self) -> pulumi.Output[Optional['outputs.FirehoseDeliveryStreamS3Configuration']]:
        """
        Required for non-S3 destinations. For S3 destination, use `extended_s3_configuration` instead. Configuration options for the s3 destination (or the intermediate bucket if the destination
        is redshift). More details are given below.
        """
        return pulumi.get(self, "s3_configuration")

    @property
    @pulumi.getter(name="serverSideEncryption")
    def server_side_encryption(self) -> pulumi.Output[Optional['outputs.FirehoseDeliveryStreamServerSideEncryption']]:
        """
        Encrypt at rest options.
        Server-side encryption should not be enabled when a kinesis stream is configured as the source of the firehose delivery stream.
        """
        return pulumi.get(self, "server_side_encryption")

    @property
    @pulumi.getter(name="splunkConfiguration")
    def splunk_configuration(self) -> pulumi.Output[Optional['outputs.FirehoseDeliveryStreamSplunkConfiguration']]:
        return pulumi.get(self, "splunk_configuration")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> pulumi.Output[str]:
        """
        Specifies the table version for the output data schema. Defaults to `LATEST`.
        """
        return pulumi.get(self, "version_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

