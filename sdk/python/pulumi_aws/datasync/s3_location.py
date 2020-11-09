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

__all__ = ['S3Location']


class S3Location(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 s3_bucket_arn: Optional[pulumi.Input[str]] = None,
                 s3_config: Optional[pulumi.Input[pulumi.InputType['S3LocationS3ConfigArgs']]] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an S3 Location within AWS DataSync.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.datasync.S3Location("example",
            s3_bucket_arn=aws_s3_bucket["example"]["arn"],
            subdirectory="/example/prefix",
            s3_config=aws.datasync.S3LocationS3ConfigArgs(
                bucket_access_role_arn=aws_iam_role["example"]["arn"],
            ))
        ```

        ## Import

        `aws_datasync_location_s3` can be imported by using the DataSync Task Amazon Resource Name (ARN), e.g.

        ```sh
         $ pulumi import aws:datasync/s3Location:S3Location example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] s3_bucket_arn: Amazon Resource Name (ARN) of the S3 Bucket.
        :param pulumi.Input[pulumi.InputType['S3LocationS3ConfigArgs']] s3_config: Configuration block containing information for connecting to S3.
        :param pulumi.Input[str] subdirectory: Prefix to perform actions as source or destination.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of resource tags to assign to the DataSync Location.
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

            if s3_bucket_arn is None:
                raise TypeError("Missing required property 's3_bucket_arn'")
            __props__['s3_bucket_arn'] = s3_bucket_arn
            if s3_config is None:
                raise TypeError("Missing required property 's3_config'")
            __props__['s3_config'] = s3_config
            if subdirectory is None:
                raise TypeError("Missing required property 'subdirectory'")
            __props__['subdirectory'] = subdirectory
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['uri'] = None
        super(S3Location, __self__).__init__(
            'aws:datasync/s3Location:S3Location',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            s3_bucket_arn: Optional[pulumi.Input[str]] = None,
            s3_config: Optional[pulumi.Input[pulumi.InputType['S3LocationS3ConfigArgs']]] = None,
            subdirectory: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'S3Location':
        """
        Get an existing S3Location resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the DataSync Location.
        :param pulumi.Input[str] s3_bucket_arn: Amazon Resource Name (ARN) of the S3 Bucket.
        :param pulumi.Input[pulumi.InputType['S3LocationS3ConfigArgs']] s3_config: Configuration block containing information for connecting to S3.
        :param pulumi.Input[str] subdirectory: Prefix to perform actions as source or destination.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of resource tags to assign to the DataSync Location.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["s3_bucket_arn"] = s3_bucket_arn
        __props__["s3_config"] = s3_config
        __props__["subdirectory"] = subdirectory
        __props__["tags"] = tags
        __props__["uri"] = uri
        return S3Location(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the DataSync Location.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="s3BucketArn")
    def s3_bucket_arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the S3 Bucket.
        """
        return pulumi.get(self, "s3_bucket_arn")

    @property
    @pulumi.getter(name="s3Config")
    def s3_config(self) -> pulumi.Output['outputs.S3LocationS3Config']:
        """
        Configuration block containing information for connecting to S3.
        """
        return pulumi.get(self, "s3_config")

    @property
    @pulumi.getter
    def subdirectory(self) -> pulumi.Output[str]:
        """
        Prefix to perform actions as source or destination.
        """
        return pulumi.get(self, "subdirectory")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value pairs of resource tags to assign to the DataSync Location.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        return pulumi.get(self, "uri")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

