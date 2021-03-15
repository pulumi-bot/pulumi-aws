# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['StreamArgs', 'Stream']

@pulumi.input_type
class StreamArgs:
    def __init__(__self__, *,
                 shard_count: pulumi.Input[int],
                 arn: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 enforce_consumer_deletion: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None,
                 shard_level_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Stream resource.
        :param pulumi.Input[int] shard_count: The number of shards that the stream will use.
               Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        :param pulumi.Input[str] encryption_type: The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        :param pulumi.Input[bool] enforce_consumer_deletion: A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] kms_key_id: The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        :param pulumi.Input[str] name: A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        :param pulumi.Input[int] retention_period: Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] shard_level_metrics: A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        pulumi.set(__self__, "shard_count", shard_count)
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if encryption_type is not None:
            pulumi.set(__self__, "encryption_type", encryption_type)
        if enforce_consumer_deletion is not None:
            pulumi.set(__self__, "enforce_consumer_deletion", enforce_consumer_deletion)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if retention_period is not None:
            pulumi.set(__self__, "retention_period", retention_period)
        if shard_level_metrics is not None:
            pulumi.set(__self__, "shard_level_metrics", shard_level_metrics)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> pulumi.Input[int]:
        """
        The number of shards that the stream will use.
        Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        """
        return pulumi.get(self, "shard_count")

    @shard_count.setter
    def shard_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "shard_count", value)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> Optional[pulumi.Input[str]]:
        """
        The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        """
        return pulumi.get(self, "encryption_type")

    @encryption_type.setter
    def encryption_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_type", value)

    @property
    @pulumi.getter(name="enforceConsumerDeletion")
    def enforce_consumer_deletion(self) -> Optional[pulumi.Input[bool]]:
        """
        A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        """
        return pulumi.get(self, "enforce_consumer_deletion")

    @enforce_consumer_deletion.setter
    def enforce_consumer_deletion(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_consumer_deletion", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        """
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> Optional[pulumi.Input[int]]:
        """
        Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
        """
        return pulumi.get(self, "retention_period")

    @retention_period.setter
    def retention_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "retention_period", value)

    @property
    @pulumi.getter(name="shardLevelMetrics")
    def shard_level_metrics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        """
        return pulumi.get(self, "shard_level_metrics")

    @shard_level_metrics.setter
    def shard_level_metrics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "shard_level_metrics", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class Stream(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StreamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        ## Import

        Kinesis Streams can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:kinesis/stream:Stream test_stream kinesis-test
        ```

         [1]https://aws.amazon.com/documentation/kinesis/ [2]https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html [3]https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html

        :param str resource_name: The name of the resource.
        :param StreamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 enforce_consumer_deletion: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None,
                 shard_level_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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

        ## Import

        Kinesis Streams can be imported using the `name`, e.g.

        ```sh
         $ pulumi import aws:kinesis/stream:Stream test_stream kinesis-test
        ```

         [1]https://aws.amazon.com/documentation/kinesis/ [2]https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html [3]https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        :param pulumi.Input[str] encryption_type: The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        :param pulumi.Input[bool] enforce_consumer_deletion: A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        :param pulumi.Input[str] kms_key_id: The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        :param pulumi.Input[str] name: A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        :param pulumi.Input[int] retention_period: Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
        :param pulumi.Input[int] shard_count: The number of shards that the stream will use.
               Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] shard_level_metrics: A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StreamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 arn: Optional[pulumi.Input[str]] = None,
                 encryption_type: Optional[pulumi.Input[str]] = None,
                 enforce_consumer_deletion: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 retention_period: Optional[pulumi.Input[int]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None,
                 shard_level_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
            if shard_count is None and not opts.urn:
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
            retention_period: Optional[pulumi.Input[int]] = None,
            shard_count: Optional[pulumi.Input[int]] = None,
            shard_level_metrics: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
        :param pulumi.Input[int] retention_period: Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
        :param pulumi.Input[int] shard_count: The number of shards that the stream will use.
               Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] shard_level_metrics: A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
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
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) specifying the Stream (same as `id`)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="encryptionType")
    def encryption_type(self) -> pulumi.Output[Optional[str]]:
        """
        The encryption type to use. The only acceptable values are `NONE` or `KMS`. The default value is `NONE`.
        """
        return pulumi.get(self, "encryption_type")

    @property
    @pulumi.getter(name="enforceConsumerDeletion")
    def enforce_consumer_deletion(self) -> pulumi.Output[Optional[bool]]:
        """
        A boolean that indicates all registered consumers should be deregistered from the stream so that the stream can be destroyed without error. The default value is `false`.
        """
        return pulumi.get(self, "enforce_consumer_deletion")

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        The GUID for the customer-managed KMS key to use for encryption. You can also use a Kinesis-owned master key by specifying the alias `alias/aws/kinesis`.
        """
        return pulumi.get(self, "kms_key_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A name to identify the stream. This is unique to the AWS account and region the Stream is created in.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionPeriod")
    def retention_period(self) -> pulumi.Output[Optional[int]]:
        """
        Length of time data records are accessible after they are added to the stream. The maximum value of a stream's retention period is 8760 hours. Minimum value is 24. Default is 24.
        """
        return pulumi.get(self, "retention_period")

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> pulumi.Output[int]:
        """
        The number of shards that the stream will use.
        Amazon has guidelines for specifying the Stream size that should be referenced when creating a Kinesis stream. See [Amazon Kinesis Streams](https://docs.aws.amazon.com/kinesis/latest/dev/amazon-kinesis-streams.html) for more.
        """
        return pulumi.get(self, "shard_count")

    @property
    @pulumi.getter(name="shardLevelMetrics")
    def shard_level_metrics(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of shard-level CloudWatch metrics which can be enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more. Note that the value ALL should not be used; instead you should provide an explicit list of metrics you wish to enable.
        """
        return pulumi.get(self, "shard_level_metrics")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

