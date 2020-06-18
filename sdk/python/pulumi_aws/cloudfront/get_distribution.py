# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDistributionResult:
    """
    A collection of values returned by getDistribution.
    """
    def __init__(__self__, arn=None, domain_name=None, enabled=None, etag=None, hosted_zone_id=None, id=None, in_progress_validation_batches=None, last_modified_time=None, status=None, tags=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        """
        The ARN (Amazon Resource Name) for the distribution. For example: arn:aws:cloudfront::123456789012:distribution/EDFDVBD632BHDS5, where 123456789012 is your AWS account ID.
        """
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        __self__.domain_name = domain_name
        """
        The domain name corresponding to the distribution. For
        example: `d604721fxaaqy9.cloudfront.net`.
        """
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        __self__.etag = etag
        """
        The current version of the distribution's information. For example:
        `E2QWRUHAPOMQZL`.
        """
        if hosted_zone_id and not isinstance(hosted_zone_id, str):
            raise TypeError("Expected argument 'hosted_zone_id' to be a str")
        __self__.hosted_zone_id = hosted_zone_id
        """
        The CloudFront Route 53 zone ID that can be used to
        route an [Alias Resource Record Set][7] to. This attribute is simply an
        alias for the zone ID `Z2FDTNDATAQYW2`.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The identifier for the distribution. For example: `EDFDVBD632BHDS5`.
        """
        if in_progress_validation_batches and not isinstance(in_progress_validation_batches, float):
            raise TypeError("Expected argument 'in_progress_validation_batches' to be a float")
        __self__.in_progress_validation_batches = in_progress_validation_batches
        """
        The number of invalidation batches
        currently in progress.
        """
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        __self__.last_modified_time = last_modified_time
        """
        The date and time the distribution was last modified.
        """
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        __self__.status = status
        """
        The current status of the distribution. `Deployed` if the
        distribution's information is fully propagated throughout the Amazon
        CloudFront system.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
class AwaitableGetDistributionResult(GetDistributionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDistributionResult(
            arn=self.arn,
            domain_name=self.domain_name,
            enabled=self.enabled,
            etag=self.etag,
            hosted_zone_id=self.hosted_zone_id,
            id=self.id,
            in_progress_validation_batches=self.in_progress_validation_batches,
            last_modified_time=self.last_modified_time,
            status=self.status,
            tags=self.tags)

def get_distribution(id=None,tags=None,opts=None):
    """
    Use this data source to retrieve information about a CloudFront distribution.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    test = aws.cloudfront.get_distribution(id="EDFDVBD632BHDS5")
    ```


    :param str id: The identifier for the distribution. For example: `EDFDVBD632BHDS5`.
    """
    __args__ = dict()


    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:cloudfront/getDistribution:getDistribution', __args__, opts=opts).value

    return AwaitableGetDistributionResult(
        arn=__ret__.get('arn'),
        domain_name=__ret__.get('domainName'),
        enabled=__ret__.get('enabled'),
        etag=__ret__.get('etag'),
        hosted_zone_id=__ret__.get('hostedZoneId'),
        id=__ret__.get('id'),
        in_progress_validation_batches=__ret__.get('inProgressValidationBatches'),
        last_modified_time=__ret__.get('lastModifiedTime'),
        status=__ret__.get('status'),
        tags=__ret__.get('tags'))
