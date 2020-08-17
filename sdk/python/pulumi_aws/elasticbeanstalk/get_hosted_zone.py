# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetHostedZoneResult',
    'AwaitableGetHostedZoneResult',
    'get_hosted_zone',
]



@pulumi.output_type
class GetHostedZoneResult:
    """
    A collection of values returned by getHostedZone.
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
        ...

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The region of the hosted zone.
        """
        ...



class AwaitableGetHostedZoneResult(GetHostedZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHostedZoneResult(
            id=self.id,
            region=self.region)


def get_hosted_zone(region: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHostedZoneResult:
    """
    Use this data source to get the ID of an [elastic beanstalk hosted zone](http://docs.aws.amazon.com/general/latest/gr/rande.html#elasticbeanstalk_region).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    current = aws.elasticbeanstalk.get_hosted_zone()
    ```


    :param str region: The region you'd like the zone for. By default, fetches the current region.
    """
    __args__ = dict()
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticbeanstalk/getHostedZone:getHostedZone', __args__, opts=opts, typ=GetHostedZoneResult).value

    return AwaitableGetHostedZoneResult(
        id=__ret__.id,
        region=__ret__.region)
