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

__all__ = ['AnalyticsConfiguration']


class AnalyticsConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 filter: Optional[pulumi.Input[pulumi.InputType['AnalyticsConfigurationFilterArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 storage_class_analysis: Optional[pulumi.Input[pulumi.InputType['AnalyticsConfigurationStorageClassAnalysisArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a S3 bucket [analytics configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html) resource.

        ## Example Usage
        ### Add analytics configuration for entire S3 bucket and export results to a second S3 bucket

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.Bucket("example")
        analytics = aws.s3.Bucket("analytics")
        example_entire_bucket = aws.s3.AnalyticsConfiguration("example-entire-bucket",
            bucket=example.bucket,
            storage_class_analysis={
                "dataExport": {
                    "destination": {
                        "s3BucketDestination": {
                            "bucketArn": analytics.arn,
                        },
                    },
                },
            })
        ```
        ### Add analytics configuration with S3 bucket object filter

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.s3.Bucket("example")
        example_filtered = aws.s3.AnalyticsConfiguration("example-filtered",
            bucket=example.bucket,
            filter={
                "prefix": "documents/",
                "tags": {
                    "priority": "high",
                    "class": "blue",
                },
            })
        ```

        ## Import

        S3 bucket analytics configurations can be imported using `bucket:analytics`, e.g.

        ```sh
         $ pulumi import aws:s3/analyticsConfiguration:AnalyticsConfiguration my-bucket-entire-bucket my-bucket:EntireBucket
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket this analytics configuration is associated with.
        :param pulumi.Input[pulumi.InputType['AnalyticsConfigurationFilterArgs']] filter: Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        :param pulumi.Input[str] name: Unique identifier of the analytics configuration for the bucket.
        :param pulumi.Input[pulumi.InputType['AnalyticsConfigurationStorageClassAnalysisArgs']] storage_class_analysis: Configuration for the analytics data export (documented below).
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

            if bucket is None:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['filter'] = filter
            __props__['name'] = name
            __props__['storage_class_analysis'] = storage_class_analysis
        super(AnalyticsConfiguration, __self__).__init__(
            'aws:s3/analyticsConfiguration:AnalyticsConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            filter: Optional[pulumi.Input[pulumi.InputType['AnalyticsConfigurationFilterArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            storage_class_analysis: Optional[pulumi.Input[pulumi.InputType['AnalyticsConfigurationStorageClassAnalysisArgs']]] = None) -> 'AnalyticsConfiguration':
        """
        Get an existing AnalyticsConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket: The name of the bucket this analytics configuration is associated with.
        :param pulumi.Input[pulumi.InputType['AnalyticsConfigurationFilterArgs']] filter: Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        :param pulumi.Input[str] name: Unique identifier of the analytics configuration for the bucket.
        :param pulumi.Input[pulumi.InputType['AnalyticsConfigurationStorageClassAnalysisArgs']] storage_class_analysis: Configuration for the analytics data export (documented below).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["bucket"] = bucket
        __props__["filter"] = filter
        __props__["name"] = name
        __props__["storage_class_analysis"] = storage_class_analysis
        return AnalyticsConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket this analytics configuration is associated with.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter
    def filter(self) -> pulumi.Output[Optional['outputs.AnalyticsConfigurationFilter']]:
        """
        Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique identifier of the analytics configuration for the bucket.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="storageClassAnalysis")
    def storage_class_analysis(self) -> pulumi.Output[Optional['outputs.AnalyticsConfigurationStorageClassAnalysis']]:
        """
        Configuration for the analytics data export (documented below).
        """
        return pulumi.get(self, "storage_class_analysis")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

