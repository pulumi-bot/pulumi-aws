# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetHostedZoneIdResult:
    """
    A collection of values returned by getHostedZoneId.
    """
    def __init__(__self__, region=None, id=None):
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_hosted_zone_id(region=None,opts=None):
    """
    Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
    in a given region for the purpose of using in an AWS Route53 Alias.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/elb_hosted_zone_id_legacy.html.markdown.
    """
    __args__ = dict()

    __args__['region'] = region
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = await pulumi.runtime.invoke('aws:elasticloadbalancing/getHostedZoneId:getHostedZoneId', __args__, opts=opts)

    return GetHostedZoneIdResult(
        region=__ret__.get('region'),
        id=__ret__.get('id'))
