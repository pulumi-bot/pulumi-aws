# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Stream']


class Stream(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 enforce_consumer_deletion: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[float]] = None,
                 shard_count: Optional[pulumi.Input[float]] = None,
                 shard_level_metrics: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Kinesis Stream resource. Amazon Kinesis is a managed service that
        scales elastically for real-time processing of streaming big data.

        For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_stream = aws.kinesis.Stream("testStream",
            retention_period=48,
            shard_count=1,
            shard_level_metrics=[
                "IncomingBytes",
                "OutgoingBytes",
            ],
            tags={
                "Environment": "test",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        :param pulumi.Input[str] encryption_type: The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        :param pulumi.Input[bool] enforce_consumer_deletion: A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] kms_key_id: The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        :param pulumi.Input[str] name: A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        :param pulumi.Input[float] retention_period: Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
        :param pulumi.Input[float] shard_count: The number of shards that the stream will use.
               Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        :param pulumi.Input[List[pulumi.Input[str]]] shard_level_metrics: A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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
            __props__['encryption_type'] = encryption_type
            __props__['enforce_consumer_deletion'] = enforce_consumer_deletion
            __props__['kms_key_id'] = kms_key_id
            __props__['name'] = name
            __props__['retention_period'] = retention_period
            if shard_count is None:
                raise TypeError("Missing required property 'shard_count'")
            __props__['shard_count'] = shard_count
            __props__['shard_level_metrics'] = shard_level_metrics
            __props__['tags'] = tags
        super(Stream, __self__).__init__(
            'aws:kinesis/stream:Stream',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            encryption_type: Optional[pulumi.Input[str]] = None,
            enforce_consumer_deletion: Optional[pulumi.Input[bool]] = None,
            kms_key_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            retention_period: Optional[pulumi.Input[float]] = None,
            shard_count: Optional[pulumi.Input[float]] = None,
            shard_level_metrics: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'Stream':
        """
        Get an existing Stream resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        :param pulumi.Input[str] encryption_type: The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        :param pulumi.Input[bool] enforce_consumer_deletion: A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] kms_key_id: The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        :param pulumi.Input[str] name: A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        :param pulumi.Input[float] retention_period: Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
        :param pulumi.Input[float] shard_count: The number of shards that the stream will use.
               Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        :param pulumi.Input[List[pulumi.Input[str]]] shard_level_metrics: A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["encryption_type"] = encryption_type
        __props__["enforce_consumer_deletion"] = enforce_consumer_deletion
        __props__["kms_key_id"] = kms_key_id
        __props__["name"] = name
        __props__["retention_period"] = retention_period
        __props__["shard_count"] = shard_count
        __props__["shard_level_metrics"] = shard_level_metrics
        __props__["tags"] = tags
        return Stream(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> Optional[str]:
        """
        The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        """
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter(name="enforceConsumerDeletion")
    def enforce_consumer_deletion(self) -> Optional[bool]:
        """
        A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        """
        return pulumi.get(self, "enforce_consumer_deletion")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[str]:
        """
        The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[float]:
        """
        Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 168 hours. Minimum value is 24. Default is 24.
        """
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> float:
        """
        The number of shards that the stream will use.
        Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        """
        return pulumi.get(self, "shard_count")

    @property
    @pulumi.getter(name="shardLevelMetrics")
    def shard_level_metrics(self) -> Optional[List[str]]:
        """
        A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        """
        return pulumi.get(self, "shard_level_metrics")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

