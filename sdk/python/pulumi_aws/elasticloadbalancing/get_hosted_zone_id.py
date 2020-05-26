# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("aws.elasticloadbalancing.getHostedZoneId has been deprecated in favor of aws.elb.getHostedZoneId", DeprecationWarning)
class GetHostedZoneIdResult:
    """
    A collection of values returned by getHostedZoneId.
    """
    def __init__(__self__, id=None, region=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
class AwaitableGetHostedZoneIdResult(GetHostedZoneIdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostedZoneIdResult(
            id=self.id,
            region=self.region)

def get_hosted_zone_id(region=None,opts=None):
    """
    Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
    in a given region for the purpose of using in an AWS Route53 Alias.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_aws as aws

    main = aws.elb.get_hosted_zone_id()
    www = aws.route53.Record("www",
        aliases=[{
            "evaluateTargetHealth": True,
            "name": aws_elb["main"]["dns_name"],
            "zone_id": main.id,
        }],
        name="example.com",
        type="A",
        zone_id=aws_route53_zone["primary"]["zone_id"])
    ```



    :param str region: Name of the region whose AWS ELB HostedZoneId is desired.
           Defaults to the region from the AWS provider configuration.
    """
    pulumi.log.warn("get_hosted_zone_id is deprecated: aws.elasticloadbalancing.getHostedZoneId has been deprecated in favor of aws.elb.getHostedZoneId")
    __args__ = dict()


    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId', __args__, opts=opts).value

    return AwaitableGetHostedZoneIdResult(
        id=__ret__.get('id'),
        region=__ret__.get('region'))
