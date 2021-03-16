# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['BucketPublicAccessBlockArgs', 'BucketPublicAccessBlock']

@pulumi.input_type
class BucketPublicAccessBlockArgs:
    def __init__(__self__, *,
                 bucket: pulumi.Input[str],
                 block_public_acls: Optional[pulumi.Input[bool]] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 ignore_public_acls: Optional[pulumi.Input[bool]] = None,
                 restrict_public_buckets: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BucketPublicAccessBlock resource.
        :param pulumi.Input[str] bucket: S3 Bucket to which this Public Access Block configuration should be applied.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls will fail if the request includes an object ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore public ACLs on this bucket and any objects that it contains.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
        """
        pulumi.set(__self__, "bucket", bucket)
        if block_public_acls is not None:
            pulumi.set(__self__, "block_public_acls", block_public_acls)
        if block_public_policy is not None:
            pulumi.set(__self__, "block_public_policy", block_public_policy)
        if ignore_public_acls is not None:
            pulumi.set(__self__, "ignore_public_acls", ignore_public_acls)
        if restrict_public_buckets is not None:
            pulumi.set(__self__, "restrict_public_buckets", restrict_public_buckets)

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Input[str]:
        """
        S3 Bucket to which this Public Access Block configuration should be applied.
        """
        return pulumi.get(self, "bucket")

    @bucket.setter
    def bucket(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket", value)

    @property
    @pulumi.getter(name="blockPublicAcls")
    def block_public_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
        * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
        * PUT Object calls will fail if the request includes an object ACL.
        """
        return pulumi.get(self, "block_public_acls")

    @block_public_acls.setter
    def block_public_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_public_acls", value)

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
        * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        """
        return pulumi.get(self, "block_public_policy")

    @block_public_policy.setter
    def block_public_policy(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_public_policy", value)

    @property
    @pulumi.getter(name="ignorePublicAcls")
    def ignore_public_acls(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
        * Ignore public ACLs on this bucket and any objects that it contains.
        """
        return pulumi.get(self, "ignore_public_acls")

    @ignore_public_acls.setter
    def ignore_public_acls(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_public_acls", value)

    @property
    @pulumi.getter(name="restrictPublicBuckets")
    def restrict_public_buckets(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
        * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
        """
        return pulumi.get(self, "restrict_public_buckets")

    @restrict_public_buckets.setter
    def restrict_public_buckets(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "restrict_public_buckets", value)


class BucketPublicAccessBlock(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_public_acls: Optional[pulumi.Input[bool]] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 ignore_public_acls: Optional[pulumi.Input[bool]] = None,
                 restrict_public_buckets: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages S3 bucket-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.Bucket("exampleBucket")
        example_bucket_public_access_block = aws.s3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock",
            bucket=example_bucket.id,
            block_public_acls=True,
            block_public_policy=True)
        ```

        ## Import

        `aws_s3_bucket_public_access_block` can be imported by using the bucket name, e.g.

        ```sh
         $ pulumi import aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock example my-bucket
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls will fail if the request includes an object ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[str] bucket: S3 Bucket to which this Public Access Block configuration should be applied.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore public ACLs on this bucket and any objects that it contains.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketPublicAccessBlockArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages S3 bucket-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.Bucket("exampleBucket")
        example_bucket_public_access_block = aws.s3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock",
            bucket=example_bucket.id,
            block_public_acls=True,
            block_public_policy=True)
        ```

        ## Import

        `aws_s3_bucket_public_access_block` can be imported by using the bucket name, e.g.

        ```sh
         $ pulumi import aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock example my-bucket
        ```

        :param str resource_name: The name of the resource.
        :param BucketPublicAccessBlockArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketPublicAccessBlockArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_public_acls: Optional[pulumi.Input[bool]] = None,
                 block_public_policy: Optional[pulumi.Input[bool]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 ignore_public_acls: Optional[pulumi.Input[bool]] = None,
                 restrict_public_buckets: Optional[pulumi.Input[bool]] = None,
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

            __props__['block_public_acls'] = block_public_acls
            __props__['block_public_policy'] = block_public_policy
            if bucket is None and not opts.urn:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['ignore_public_acls'] = ignore_public_acls
            __props__['restrict_public_buckets'] = restrict_public_buckets
        super(BucketPublicAccessBlock, __self__).__init__(
            'aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block_public_acls: Optional[pulumi.Input[bool]] = None,
            block_public_policy: Optional[pulumi.Input[bool]] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            ignore_public_acls: Optional[pulumi.Input[bool]] = None,
            restrict_public_buckets: Optional[pulumi.Input[bool]] = None) -> 'BucketPublicAccessBlock':
        """
        Get an existing BucketPublicAccessBlock resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] block_public_acls: Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
               * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
               * PUT Object calls will fail if the request includes an object ACL.
        :param pulumi.Input[bool] block_public_policy: Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
               * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        :param pulumi.Input[str] bucket: S3 Bucket to which this Public Access Block configuration should be applied.
        :param pulumi.Input[bool] ignore_public_acls: Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
               * Ignore public ACLs on this bucket and any objects that it contains.
        :param pulumi.Input[bool] restrict_public_buckets: Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
               * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["block_public_acls"] = block_public_acls
        __props__["block_public_policy"] = block_public_policy
        __props__["bucket"] = bucket
        __props__["ignore_public_acls"] = ignore_public_acls
        __props__["restrict_public_buckets"] = restrict_public_buckets
        return BucketPublicAccessBlock(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blockPublicAcls")
    def block_public_acls(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should block public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect existing policies or ACLs. When set to `true` causes the following behavior:
        * PUT Bucket acl and PUT Object acl calls will fail if the specified ACL allows public access.
        * PUT Object calls will fail if the request includes an object ACL.
        """
        return pulumi.get(self, "block_public_acls")

    @property
    @pulumi.getter(name="blockPublicPolicy")
    def block_public_policy(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should block public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the existing bucket policy. When set to `true` causes Amazon S3 to:
        * Reject calls to PUT Bucket policy if the specified bucket policy allows public access.
        """
        return pulumi.get(self, "block_public_policy")

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        S3 Bucket to which this Public Access Block configuration should be applied.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="ignorePublicAcls")
    def ignore_public_acls(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should ignore public ACLs for this bucket. Defaults to `false`. Enabling this setting does not affect the persistence of any existing ACLs and doesn't prevent new public ACLs from being set. When set to `true` causes Amazon S3 to:
        * Ignore public ACLs on this bucket and any objects that it contains.
        """
        return pulumi.get(self, "ignore_public_acls")

    @property
    @pulumi.getter(name="restrictPublicBuckets")
    def restrict_public_buckets(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Amazon S3 should restrict public bucket policies for this bucket. Defaults to `false`. Enabling this setting does not affect the previously stored bucket policy, except that public and cross-account access within the public bucket policy, including non-public delegation to specific accounts, is blocked. When set to `true`:
        * Only the bucket owner and AWS Services can access this buckets if it has a public policy.
        """
        return pulumi.get(self, "restrict_public_buckets")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

