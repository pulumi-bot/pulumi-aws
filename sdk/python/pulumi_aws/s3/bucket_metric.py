# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['BucketMetricArgs', 'BucketMetric']

@pulumi.input_type
class BucketMetricArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 filter: Optional[pulumi.Input['BucketMetricFilterArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BucketMetric resource.
        :param pulumi.Input[str] bucket: The name of the bucket to put metric configuration.
        :param pulumi.Input['BucketMetricFilterArgs'] filter: [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        :param pulumi.Input[str] name: Unique identifier of the metrics configuration for the bucket.
        """
        pulumi.set(__self__, "bucket", bucket)
        if filter is not None:
            pulumi.set(__self__, "filter", filter)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        The name of the bucket to put metric configuration.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter
    def filter(self) -> Optional[pulumi.Input['BucketMetricFilterArgs']]:
        """
        [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        """
        return pulumi.get(self, "filter")

    @filter.setter
    def filter(self, value: Optional[pulumi.Input['BucketMetricFilterArgs']]):
        pulumi.set(self, "filter", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier of the metrics configuration for the bucket.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class BucketMetric(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[pulumi.InputType['BucketMetricFilterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.

        ## Example Usage
        ### Add metrics configuration for entire S3 bucket

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.Bucket("example")
        example_entire_bucket = aws.s3.BucketMetric("example-entire-bucket", bucket=example.bucket)
        ```
        ### Add metrics configuration with S3 bucket object filter

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.Bucket("example")
        example_filtered = aws.s3.BucketMetric("example-filtered",
            bucket=example.bucket,
            filter=aws.s3.BucketMetricFilterArgs(
                prefix="documents/",
                tags={
                    "priority": "high",
                    "class": "blue",
                },
            ))
        ```

        ## Import

        S3 bucket metric configurations can be imported using `bucket:metric`, e.g.

        ```sh
         $ pulumi import aws:s3/bucketMetric:BucketMetric my-bucket-entire-bucket my-bucket:EntireBucket
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket to put metric configuration.
        :param pulumi.Input[pulumi.InputType['BucketMetricFilterArgs']] filter: [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        :param pulumi.Input[str] name: Unique identifier of the metrics configuration for the bucket.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketMetricArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.

        ## Example Usage
        ### Add metrics configuration for entire S3 bucket

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.Bucket("example")
        example_entire_bucket = aws.s3.BucketMetric("example-entire-bucket", bucket=example.bucket)
        ```
        ### Add metrics configuration with S3 bucket object filter

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.Bucket("example")
        example_filtered = aws.s3.BucketMetric("example-filtered",
            bucket=example.bucket,
            filter=aws.s3.BucketMetricFilterArgs(
                prefix="documents/",
                tags={
                    "priority": "high",
                    "class": "blue",
                },
            ))
        ```

        ## Import

        S3 bucket metric configurations can be imported using `bucket:metric`, e.g.

        ```sh
         $ pulumi import aws:s3/bucketMetric:BucketMetric my-bucket-entire-bucket my-bucket:EntireBucket
        ```

        :param str resource_name: The name of the resource.
        :param BucketMetricArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketMetricArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[pulumi.InputType['BucketMetricFilterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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

            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['filter'] = filter
            __props__['name'] = name
        super(BucketMetric, __self__).__init__(
            'aws:s3/bucketMetric:BucketMetric',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[pulumi.InputType['BucketMetricFilterArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'BucketMetric':
        """
        Get an existing BucketMetric resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket to put metric configuration.
        :param pulumi.Input[pulumi.InputType['BucketMetricFilterArgs']] filter: [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        :param pulumi.Input[str] name: Unique identifier of the metrics configuration for the bucket.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket"] = bucket
        __props__["filter"] = filter
        __props__["name"] = name
        return BucketMetric(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket to put metric configuration.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional['outputs.BucketMetricFilter']]:
        """
        [Object filtering](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html#metrics-configurations-filter) that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique identifier of the metrics configuration for the bucket.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

