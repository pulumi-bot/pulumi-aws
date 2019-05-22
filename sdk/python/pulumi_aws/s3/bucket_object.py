# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class BucketObject(pulumi.CustomResource):
    acl: pulumi.Output[str]
    """
    The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
    """
    bucket: pulumi.Output[str]
    """
    The name of the bucket to put the file in.
    """
    cache_control: pulumi.Output[str]
    """
    Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
    """
    content: pulumi.Output[str]
    """
    Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
    """
    content_base64: pulumi.Output[str]
    """
    Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
    """
    content_disposition: pulumi.Output[str]
    """
    Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
    """
    content_encoding: pulumi.Output[str]
    """
    Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
    """
    content_language: pulumi.Output[str]
    """
    The language the content is in e.g. en-US or en-GB.
    """
    content_type: pulumi.Output[str]
    """
    A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
    """
    etag: pulumi.Output[str]
    """
    Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (Terraform 0.11.12 or later) or `${md5(file("path/to/file"))}` (Terraform 0.11.11 or earlier).
    This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"`.
    """
    key: pulumi.Output[str]
    """
    The name of the object once it is in the bucket.
    """
    kms_key_id: pulumi.Output[str]
    """
    Specifies the AWS KMS Key ARN to use for object encryption.
    This value is a fully qualified **ARN** of the KMS Key. If using `aws_kms_key`,
    use the exported `arn` attribute:
    `kms_key_id = "${aws_kms_key.foo.arn}"`
    """
    server_side_encryption: pulumi.Output[str]
    """
    Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
    """
    source: pulumi.Output[pulumi.Asset]
    """
    The path to a file that will be read and uploaded as raw bytes for the object content.
    """
    storage_class: pulumi.Output[str]
    """
    Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
    for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the object.
    """
    version_id: pulumi.Output[str]
    """
    A unique version ID value for the object, if bucket versioning
    is enabled.
    """
    website_redirect: pulumi.Output[str]
    """
    Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
    """
    def __init__(__self__, resource_name, opts=None, acl=None, bucket=None, cache_control=None, content=None, content_base64=None, content_disposition=None, content_encoding=None, content_language=None, content_type=None, etag=None, key=None, kms_key_id=None, server_side_encryption=None, source=None, storage_class=None, tags=None, website_redirect=None, __name__=None, __opts__=None):
        """
        Provides a S3 bucket object resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl: The [canned ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) to apply. Defaults to "private".
        :param pulumi.Input[str] bucket: The name of the bucket to put the file in.
        :param pulumi.Input[str] cache_control: Specifies caching behavior along the request/reply chain Read [w3c cache_control](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9) for further details.
        :param pulumi.Input[str] content: Literal string value to use as the object content, which will be uploaded as UTF-8-encoded text.
        :param pulumi.Input[str] content_base64: Base64-encoded data that will be decoded and uploaded as raw bytes for the object content. This allows safely uploading non-UTF8 binary data, but is recommended only for small content such as the result of the `gzipbase64` function with small text strings. For larger objects, use `source` to stream the content from a disk file.
        :param pulumi.Input[str] content_disposition: Specifies presentational information for the object. Read [w3c content_disposition](http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1) for further information.
        :param pulumi.Input[str] content_encoding: Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. Read [w3c content encoding](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11) for further information.
        :param pulumi.Input[str] content_language: The language the content is in e.g. en-US or en-GB.
        :param pulumi.Input[str] content_type: A standard MIME type describing the format of the object data, e.g. application/octet-stream. All Valid MIME Types are valid for this input.
        :param pulumi.Input[str] etag: Used to trigger updates. The only meaningful value is `${filemd5("path/to/file")}` (Terraform 0.11.12 or later) or `${md5(file("path/to/file"))}` (Terraform 0.11.11 or earlier).
               This attribute is not compatible with KMS encryption, `kms_key_id` or `server_side_encryption = "aws:kms"`.
        :param pulumi.Input[str] key: The name of the object once it is in the bucket.
        :param pulumi.Input[str] kms_key_id: Specifies the AWS KMS Key ARN to use for object encryption.
               This value is a fully qualified **ARN** of the KMS Key. If using `aws_kms_key`,
               use the exported `arn` attribute:
               `kms_key_id = "${aws_kms_key.foo.arn}"`
        :param pulumi.Input[str] server_side_encryption: Specifies server-side encryption of the object in S3. Valid values are "`AES256`" and "`aws:kms`".
        :param pulumi.Input[pulumi.Asset] source: The path to a file that will be read and uploaded as raw bytes for the object content.
        :param pulumi.Input[str] storage_class: Specifies the desired [Storage Class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html)
               for the object. Can be either "`STANDARD`", "`REDUCED_REDUNDANCY`", "`ONEZONE_IA`", "`INTELLIGENT_TIERING`", "`GLACIER`", "`DEEP_ARCHIVE`", or "`STANDARD_IA`". Defaults to "`STANDARD`".
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the object.
        :param pulumi.Input[str] website_redirect: Specifies a target URL for [website redirect](http://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html).
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

        __props__['acl'] = acl

        if bucket is None:
            raise TypeError("Missing required property 'bucket'")
        __props__['bucket'] = bucket

        __props__['cache_control'] = cache_control

        __props__['content'] = content

        __props__['content_base64'] = content_base64

        __props__['content_disposition'] = content_disposition

        __props__['content_encoding'] = content_encoding

        __props__['content_language'] = content_language

        __props__['content_type'] = content_type

        __props__['etag'] = etag

        __props__['key'] = key

        __props__['kms_key_id'] = kms_key_id

        __props__['server_side_encryption'] = server_side_encryption

        __props__['source'] = source

        __props__['storage_class'] = storage_class

        __props__['tags'] = tags

        __props__['website_redirect'] = website_redirect

        __props__['version_id'] = None

        super(BucketObject, __self__).__init__(
            'aws:s3/bucketObject:BucketObject',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

