# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ResourceDataSync(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    Name for the configuration.
    """
    s3_destination: pulumi.Output[dict]
    """
    Amazon S3 configuration details for the sync.

      * `bucket_name` (`str`) - Name of S3 bucket where the aggregated data is stored.
      * `kms_key_arn` (`str`) - ARN of an encryption key for a destination in Amazon S3.
      * `prefix` (`str`) - Prefix for the bucket.
      * `region` (`str`) - Region with the bucket targeted by the Resource Data Sync.
      * `syncFormat` (`str`) - A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
    """
    def __init__(__self__, resource_name, opts=None, name=None, s3_destination=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a SSM resource data sync.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        hoge_bucket = aws.s3.Bucket("hogeBucket", region="us-east-1")
        hoge_bucket_policy = aws.s3.BucketPolicy("hogeBucketPolicy",
            bucket=hoge_bucket.bucket,
            policy=\"\"\"{
            "Version": "2012-10-17",
            "Statement": [
                {
                    "Sid": "SSMBucketPermissionsCheck",
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "ssm.amazonaws.com"
                    },
                    "Action": "s3:GetBucketAcl",
                    "Resource": "arn:aws:s3:::tf-test-bucket-1234"
                },
                {
                    "Sid": " SSMBucketDelivery",
                    "Effect": "Allow",
                    "Principal": {
                        "Service": "ssm.amazonaws.com"
                    },
                    "Action": "s3:PutObject",
                    "Resource": ["arn:aws:s3:::tf-test-bucket-1234/*"],
                    "Condition": {
                        "StringEquals": {
                            "s3:x-amz-acl": "bucket-owner-full-control"
                        }
                    }
                }
            ]
        }

        \"\"\")
        foo = aws.ssm.ResourceDataSync("foo", s3_destination={
            "bucket_name": hoge_bucket.bucket,
            "region": hoge_bucket.region,
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name for the configuration.
        :param pulumi.Input[dict] s3_destination: Amazon S3 configuration details for the sync.

        The **s3_destination** object supports the following:

          * `bucket_name` (`pulumi.Input[str]`) - Name of S3 bucket where the aggregated data is stored.
          * `kms_key_arn` (`pulumi.Input[str]`) - ARN of an encryption key for a destination in Amazon S3.
          * `prefix` (`pulumi.Input[str]`) - Prefix for the bucket.
          * `region` (`pulumi.Input[str]`) - Region with the bucket targeted by the Resource Data Sync.
          * `syncFormat` (`pulumi.Input[str]`) - A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
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
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['name'] = name
            if s3_destination is None:
                raise TypeError("Missing required property 's3_destination'")
            __props__['s3_destination'] = s3_destination
        super(ResourceDataSync, __self__).__init__(
            'aws:ssm/resourceDataSync:ResourceDataSync',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, s3_destination=None):
        """
        Get an existing ResourceDataSync resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name for the configuration.
        :param pulumi.Input[dict] s3_destination: Amazon S3 configuration details for the sync.

        The **s3_destination** object supports the following:

          * `bucket_name` (`pulumi.Input[str]`) - Name of S3 bucket where the aggregated data is stored.
          * `kms_key_arn` (`pulumi.Input[str]`) - ARN of an encryption key for a destination in Amazon S3.
          * `prefix` (`pulumi.Input[str]`) - Prefix for the bucket.
          * `region` (`pulumi.Input[str]`) - Region with the bucket targeted by the Resource Data Sync.
          * `syncFormat` (`pulumi.Input[str]`) - A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["s3_destination"] = s3_destination
        return ResourceDataSync(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
