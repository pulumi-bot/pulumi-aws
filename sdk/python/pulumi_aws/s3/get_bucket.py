# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetBucketResult',
    'AwaitableGetBucketResult',
    'get_bucket',
]

@pulumi.output_type
class GetBucketResult:
    """
    A collection of values returned by getBucket.
    """
    def __init__(__self__, arn=None, bucket=None, bucket_domain_name=None, bucket_regional_domain_name=None, hosted_zone_id=None, id=None, region=None, website_domain=None, website_endpoint=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
        if bucket_domain_name and not isinstance(bucket_domain_name, str):
            raise TypeError("Expected argument 'bucket_domain_name' to be a str")
        pulumi.set(__self__, "bucket_domain_name", bucket_domain_name)
        if bucket_regional_domain_name and not isinstance(bucket_regional_domain_name, str):
            raise TypeError("Expected argument 'bucket_regional_domain_name' to be a str")
        pulumi.set(__self__, "bucket_regional_domain_name", bucket_regional_domain_name)
        if hosted_zone_id and not isinstance(hosted_zone_id, str):
            raise TypeError("Expected argument 'hosted_zone_id' to be a str")
        pulumi.set(__self__, "hosted_zone_id", hosted_zone_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if website_domain and not isinstance(website_domain, str):
            raise TypeError("Expected argument 'website_domain' to be a str")
        pulumi.set(__self__, "website_domain", website_domain)
        if website_endpoint and not isinstance(website_endpoint, str):
            raise TypeError("Expected argument 'website_endpoint' to be a str")
        pulumi.set(__self__, "website_endpoint", website_endpoint)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def bucket(self) -> str:
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="bucketDomainName")
    def bucket_domain_name(self) -> str:
        """
        The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
        """
        return pulumi.get(self, "bucket_domain_name")

    @property
    @pulumi.getter(name="bucketRegionalDomainName")
    def bucket_regional_domain_name(self) -> str:
        """
        The bucket region-specific domain name. The bucket domain name including the region name, please refer [here](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region) for format. Note: The AWS CloudFront allows specifying S3 region-specific endpoint when creating S3 origin, it will prevent [redirect issues](https://forums.aws.amazon.com/thread.jspa?threadID=216814) from CloudFront to S3 Origin URL.
        """
        return pulumi.get(self, "bucket_regional_domain_name")

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> str:
        """
        The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
        """
        return pulumi.get(self, "hosted_zone_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The AWS region this bucket resides in.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="websiteDomain")
    def website_domain(self) -> str:
        """
        The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
        """
        return pulumi.get(self, "website_domain")

    @property
    @pulumi.getter(name="websiteEndpoint")
    def website_endpoint(self) -> str:
        """
        The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
        """
        return pulumi.get(self, "website_endpoint")


class AwaitableGetBucketResult(GetBucketResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBucketResult(
            arn=self.arn,
            bucket=self.bucket,
            bucket_domain_name=self.bucket_domain_name,
            bucket_regional_domain_name=self.bucket_regional_domain_name,
            hosted_zone_id=self.hosted_zone_id,
            id=self.id,
            region=self.region,
            website_domain=self.website_domain,
            website_endpoint=self.website_endpoint)


def get_bucket(bucket: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBucketResult:
    """
    Provides details about a specific S3 bucket.

    This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
    Distribution.

    ## Example Usage


    :param str bucket: The name of the bucket
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:s3/getBucket:getBucket', __args__, opts=opts, typ=GetBucketResult).value

    return AwaitableGetBucketResult(
        arn=__ret__.arn,
        bucket=__ret__.bucket,
        bucket_domain_name=__ret__.bucket_domain_name,
        bucket_regional_domain_name=__ret__.bucket_regional_domain_name,
        hosted_zone_id=__ret__.hosted_zone_id,
        id=__ret__.id,
        region=__ret__.region,
        website_domain=__ret__.website_domain,
        website_endpoint=__ret__.website_endpoint)
