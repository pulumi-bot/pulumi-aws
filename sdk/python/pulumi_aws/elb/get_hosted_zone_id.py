# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetHostedZoneIdResult',
    'AwaitableGetHostedZoneIdResult',
    'get_hosted_zone_id',
]

@pulumi.output_type
class GetHostedZoneIdResult:
    """
    A collection of values returned by getHostedZoneId.
    """
    def __init__(__self__, id=None, region=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")


class AwaitableGetHostedZoneIdResult(GetHostedZoneIdResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostedZoneIdResult(
            id=self.id,
            region=self.region)


def get_hosted_zone_id(region: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostedZoneIdResult:
    """
    Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
    in a given region for the purpose of using in an AWS Route53 Alias.


    :param str region: Name of the region whose AWS ELB HostedZoneId is desired.
           Defaults to the region from the AWS provider configuration.
    """
    __args__ = dict()
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elb/getHostedZoneId:getHostedZoneId', __args__, opts=opts, typ=GetHostedZoneIdResult).value

    return AwaitableGetHostedZoneIdResult(
        id=__ret__.id,
        region=__ret__.region)
